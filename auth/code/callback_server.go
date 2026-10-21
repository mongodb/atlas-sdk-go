// Copyright 2026 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package code

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// ErrCallbackServerClosed is returned when the server closes before receiving a callback.
var ErrCallbackServerClosed = errors.New("OAuth callback server closed before receiving a response")

type callbackResult struct {
	code string
	err  error
}

// CallbackServer accepts one state-validated loopback authorization callback,
// validating its issuer when present and requiring it when the authorization
// server advertises support.
type CallbackServer struct {
	listener      net.Listener
	server        *http.Server
	state         string
	callbackPath  string
	issuer        string
	requireIssuer bool
	done          chan struct{}
	result        callbackResult
	once          sync.Once
}

// StartCallbackServer listens on 127.0.0.1 using an ephemeral port and serves
// callbacks at callbackPath. callbackPath must be an absolute URL path without
// a query or fragment, and must match the path registered for the OAuth client.
// Any iss parameter in the callback must match metadata.Issuer. When metadata
// advertises authorization response issuer support, the callback must include
// iss. metadata must be the validated authorization-server discovery response.
func StartCallbackServer(expectedState, callbackPath string, metadata *AuthServerMetadata) (*CallbackServer, error) {
	if expectedState == "" {
		return nil, errors.New("expected OAuth state is required")
	}
	if metadata == nil {
		return nil, errors.New("OAuth Authorization Server metadata is required")
	}
	if _, err := parseOAuthEndpoint(metadata.Issuer, "Authorization Server issuer"); err != nil {
		return nil, err
	}
	callbackPath, err := normalizeCallbackPath(callbackPath)
	if err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("starting OAuth callback listener: %w", err)
	}

	callback := &CallbackServer{
		listener:      listener,
		state:         expectedState,
		callbackPath:  callbackPath,
		issuer:        metadata.Issuer,
		requireIssuer: metadata.AuthorizationResponseIssParameterSupported,
		done:          make(chan struct{}),
	}
	callback.server = &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.EscapedPath() != callback.callbackPath {
				http.NotFound(w, r)
				return
			}
			callback.handleCallback(w, r)
		}),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		if err := callback.server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			callback.complete(callbackResult{err: fmt.Errorf("serving OAuth callback: %w", err)})
		}
	}()
	return callback, nil
}

func normalizeCallbackPath(callbackPath string) (string, error) {
	if !strings.HasPrefix(callbackPath, "/") {
		return "", errors.New("OAuth callback path must be an absolute URL path")
	}
	if strings.ContainsAny(callbackPath, "?#") {
		return "", errors.New("OAuth callback path must not include a query or fragment")
	}

	parsedPath, err := url.ParseRequestURI(callbackPath)
	if err != nil {
		return "", fmt.Errorf("invalid OAuth callback path: %w", err)
	}
	if parsedPath.Path == "" || !strings.HasPrefix(parsedPath.Path, "/") ||
		parsedPath.IsAbs() || parsedPath.Host != "" || parsedPath.Opaque != "" ||
		parsedPath.RawQuery != "" || parsedPath.ForceQuery || parsedPath.Fragment != "" {
		return "", errors.New("OAuth callback path must not include a host, query, or fragment")
	}

	return parsedPath.EscapedPath(), nil
}

// RedirectURI returns the loopback callback URI.
func (s *CallbackServer) RedirectURI() string {
	return fmt.Sprintf("http://127.0.0.1:%d%s", s.Port(), s.callbackPath)
}

// Port returns the callback listener's port.
func (s *CallbackServer) Port() int {
	return s.listener.Addr().(*net.TCPAddr).Port
}

// WaitForCallback waits for a callback or context cancellation.
func (s *CallbackServer) WaitForCallback(ctx context.Context) (string, error) {
	select {
	case <-s.done:
		return s.result.code, s.result.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// Close shuts down the callback server and releases its listener.
func (s *CallbackServer) Close() error {
	s.complete(callbackResult{err: ErrCallbackServerClosed})
	serverErr := s.server.Close()
	listenerErr := s.listener.Close()
	if errors.Is(serverErr, http.ErrServerClosed) || errors.Is(serverErr, net.ErrClosed) {
		serverErr = nil
	}
	if errors.Is(listenerErr, net.ErrClosed) {
		listenerErr = nil
	}
	return errors.Join(serverErr, listenerErr)
}

func (s *CallbackServer) complete(result callbackResult) {
	s.completeWithResponse(result, nil)
}

func (s *CallbackServer) completeWithResponse(result callbackResult, writeResponse func()) bool {
	completed := false
	s.once.Do(func() {
		s.result = result
		if writeResponse != nil {
			writeResponse()
		}
		close(s.done)
		completed = true
	})
	return completed
}

func (s *CallbackServer) validateCallbackIssuer(query url.Values) error {
	issuers, present := query["iss"]
	if !present {
		if s.requireIssuer {
			return errors.New("OAuth callback response must include exactly one issuer")
		}
		return nil
	}
	if len(issuers) != 1 || issuers[0] == "" {
		return errors.New("OAuth callback response must include exactly one issuer")
	}
	if issuers[0] != s.issuer {
		return errors.New("OAuth callback issuer does not match discovered issuer")
	}
	return nil
}

func (s *CallbackServer) handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeCallbackResponse(w, http.StatusMethodNotAllowed, "Method not allowed.")
		return
	}

	query := r.URL.Query()
	if query.Get("state") != s.state {
		// Reject unsolicited callbacks without consuming the one valid result.
		writeCallbackResponse(w, http.StatusBadRequest, "Authorization failed: state mismatch.")
		return
	}

	if err := s.validateCallbackIssuer(query); err != nil {
		if !s.completeWithResponse(callbackResult{err: err}, func() {
			writeCallbackResponse(w, http.StatusBadRequest, "Authorization failed: issuer validation failed.")
		}) {
			writeCallbackResponse(w, http.StatusGone, "Authorization response already received.")
			return
		}
		return
	}

	if errorCode := query.Get("error"); errorCode != "" {
		err := &OAuthError{
			Operation:        "authorization",
			ErrorCode:        errorCode,
			ErrorDescription: query.Get("error_description"),
			ErrorURI:         query.Get("error_uri"),
		}
		if !s.completeWithResponse(callbackResult{err: err}, func() {
			writeCallbackResponse(w, http.StatusBadRequest, "Authorization failed. You may close this window.")
		}) {
			writeCallbackResponse(w, http.StatusGone, "Authorization response already received.")
			return
		}
		return
	}

	code := query.Get("code")
	if code == "" {
		if !s.completeWithResponse(callbackResult{err: errors.New("OAuth callback response did not include an authorization code")}, func() {
			writeCallbackResponse(w, http.StatusBadRequest, "Authorization failed: no code received.")
		}) {
			writeCallbackResponse(w, http.StatusGone, "Authorization response already received.")
			return
		}
		return
	}

	if !s.completeWithResponse(callbackResult{code: code}, func() {
		writeCallbackResponse(w, http.StatusOK, "Authorization successful. You may close this window.")
	}) {
		writeCallbackResponse(w, http.StatusGone, "Authorization response already received.")
		return
	}
}

func writeCallbackResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	_, _ = io.WriteString(w, message)
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}
