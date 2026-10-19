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
	"sync"
	"time"
)

const callbackPath = "/atlas-cli/callback"

// ErrCallbackServerClosed is returned when the server closes before receiving a callback.
var ErrCallbackServerClosed = errors.New("OAuth callback server closed before receiving a response")

type callbackResult struct {
	code string
	err  error
}

// CallbackServer accepts one state-validated loopback authorization callback.
type CallbackServer struct {
	listener net.Listener
	server   *http.Server
	state    string
	done     chan struct{}
	result   callbackResult
	once     sync.Once
}

// StartCallbackServer listens on 127.0.0.1 using an ephemeral port.
func StartCallbackServer(expectedState string) (*CallbackServer, error) {
	if expectedState == "" {
		return nil, errors.New("expected OAuth state is required")
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("starting OAuth callback listener: %w", err)
	}

	callback := &CallbackServer{
		listener: listener,
		state:    expectedState,
		done:     make(chan struct{}),
	}
	mux := http.NewServeMux()
	mux.HandleFunc(callbackPath, callback.handleCallback)
	callback.server = &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		if err := callback.server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			callback.complete(callbackResult{err: fmt.Errorf("serving OAuth callback: %w", err)})
		}
	}()
	return callback, nil
}

// RedirectURI returns the loopback callback URI.
func (s *CallbackServer) RedirectURI() string {
	return fmt.Sprintf("http://127.0.0.1:%d%s", s.Port(), callbackPath)
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

func (s *CallbackServer) complete(result callbackResult) bool {
	completed := false
	s.once.Do(func() {
		s.result = result
		close(s.done)
		completed = true
	})
	return completed
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

	if errorCode := query.Get("error"); errorCode != "" {
		err := &OAuthError{
			Operation:        "authorization",
			ErrorCode:        errorCode,
			ErrorDescription: query.Get("error_description"),
			ErrorURI:         query.Get("error_uri"),
		}
		if !s.complete(callbackResult{err: err}) {
			writeCallbackResponse(w, http.StatusGone, "Authorization response already received.")
			return
		}
		writeCallbackResponse(w, http.StatusBadRequest, "Authorization failed. You may close this window.")
		return
	}

	code := query.Get("code")
	if code == "" {
		if !s.complete(callbackResult{err: errors.New("OAuth callback response did not include an authorization code")}) {
			writeCallbackResponse(w, http.StatusGone, "Authorization response already received.")
			return
		}
		writeCallbackResponse(w, http.StatusBadRequest, "Authorization failed: no code received.")
		return
	}

	if !s.complete(callbackResult{code: code}) {
		writeCallbackResponse(w, http.StatusGone, "Authorization response already received.")
		return
	}
	writeCallbackResponse(w, http.StatusOK, "Authorization successful. You may close this window.")
}

func writeCallbackResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	_, _ = io.WriteString(w, message)
}
