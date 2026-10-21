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
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

const (
	testCallbackPath   = "/test-app/callback"
	testCallbackIssuer = "https://auth.example.com"
)

func callbackServerTestMetadata(issSupported bool) *AuthServerMetadata {
	return &AuthServerMetadata{
		Issuer: testCallbackIssuer,
		AuthorizationResponseIssParameterSupported: issSupported,
	}
}

func TestStartCallbackServerUsesEphemeralLoopbackURI(t *testing.T) {
	callback, err := StartCallbackServer("expected-state", testCallbackPath, callbackServerTestMetadata(false))
	if err != nil {
		t.Fatalf("StartCallbackServer() returned error: %v", err)
	}
	defer callback.Close()

	u, err := url.Parse(callback.RedirectURI())
	if err != nil {
		t.Fatalf("RedirectURI() returned invalid URL: %v", err)
	}
	if u.Scheme != "http" || u.Hostname() != "127.0.0.1" || u.Path != testCallbackPath {
		t.Errorf("redirect URI = %q, expected loopback HTTP callback at %s", u, testCallbackPath)
	}
	if callback.Port() == 0 || u.Port() == "" {
		t.Errorf("callback port = %d, expected an OS-assigned port", callback.Port())
	}
}

func TestStartCallbackServerRejectsInvalidCallbackPaths(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{name: "empty", path: ""},
		{name: "relative", path: "callback"},
		{name: "query", path: "/callback?query=1"},
		{name: "fragment", path: "/callback#fragment"},
		{name: "absolute URL", path: "https://example.com/callback"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := StartCallbackServer("expected-state", test.path, callbackServerTestMetadata(false)); err == nil {
				t.Errorf("StartCallbackServer(%q) returned nil error, expected invalid callback path", test.path)
			}
		})
	}
}

func TestCallbackServerValidatesIssuerWhenPresentOrRequired(t *testing.T) {
	tests := []struct {
		name         string
		issSupported bool
		issuers      []string
		wantErr      string
	}{
		{name: "issuer omitted when not advertised"},
		{name: "matching issuer when not advertised", issuers: []string{testCallbackIssuer}},
		{name: "mismatched issuer when not advertised", issuers: []string{"https://other.example.com"}, wantErr: "does not match discovered issuer"},
		{name: "empty issuer when not advertised", issuers: []string{""}, wantErr: "include exactly one issuer"},
		{name: "duplicate issuer when not advertised", issuers: []string{testCallbackIssuer, testCallbackIssuer}, wantErr: "include exactly one issuer"},
		{name: "matching issuer when advertised", issSupported: true, issuers: []string{testCallbackIssuer}},
		{name: "missing issuer when advertised", issSupported: true, wantErr: "include exactly one issuer"},
		{name: "mismatched issuer when advertised", issSupported: true, issuers: []string{"https://other.example.com"}, wantErr: "does not match discovered issuer"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			callback, err := StartCallbackServer("expected-state", testCallbackPath, callbackServerTestMetadata(test.issSupported))
			if err != nil {
				t.Fatalf("StartCallbackServer() returned error: %v", err)
			}
			defer callback.Close()

			query := url.Values{
				"state": {"expected-state"},
				"code":  {"authorization-code"},
			}
			for _, issuer := range test.issuers {
				query.Add("iss", issuer)
			}
			response := sendCallbackRequest(t, callback, query)
			_ = response.Body.Close()

			code, err := callback.WaitForCallback(context.Background())
			if test.wantErr != "" {
				if response.StatusCode != http.StatusBadRequest {
					t.Errorf("callback status = %d, expected HTTP 400", response.StatusCode)
				}
				if code != "" || err == nil || !strings.Contains(err.Error(), test.wantErr) {
					t.Errorf("WaitForCallback() = (%q, %v), expected issuer validation error containing %q", code, err, test.wantErr)
				}
				return
			}
			if response.StatusCode != http.StatusOK {
				t.Errorf("callback status = %d, expected HTTP 200", response.StatusCode)
			}
			if code != "authorization-code" || err != nil {
				t.Errorf("WaitForCallback() = (%q, %v), expected authorization code", code, err)
			}
		})
	}
}

func TestCallbackServerRejectsMismatchedStateWithoutConsumingFlow(t *testing.T) {
	callback, err := StartCallbackServer("expected-state", testCallbackPath, callbackServerTestMetadata(false))
	if err != nil {
		t.Fatalf("StartCallbackServer() returned error: %v", err)
	}
	defer callback.Close()

	badResponse := sendCallbackRequest(t, callback, url.Values{
		"state":             {"wrong-state"},
		"code":              {"attacker-code"},
		"error":             {"access_denied"},
		"error_description": {"attacker description"},
		"error_uri":         {"https://attacker.example.com"},
	})
	if badResponse.StatusCode != http.StatusBadRequest {
		t.Errorf("mismatched-state status = %d, expected HTTP 400", badResponse.StatusCode)
	}
	badBody, err := io.ReadAll(badResponse.Body)
	if err != nil {
		t.Fatalf("reading mismatched-state response: %v", err)
	}
	_ = badResponse.Body.Close()
	if strings.Contains(string(badBody), "attacker") {
		t.Errorf("mismatched-state response exposed callback data: %q", badBody)
	}

	response := sendCallbackRequest(t, callback, url.Values{
		"state": {"expected-state"},
		"code":  {"authorization-code"},
	})
	if response.StatusCode != http.StatusOK {
		t.Errorf("valid callback status = %d, expected HTTP 200", response.StatusCode)
	}
	_ = response.Body.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	code, err := callback.WaitForCallback(ctx)
	if err != nil {
		t.Fatalf("WaitForCallback() returned error: %v", err)
	}
	if code != "authorization-code" {
		t.Errorf("authorization code = %q, expected authorization-code", code)
	}
}

func TestCallbackServerPreservesOAuthErrorAfterStateValidation(t *testing.T) {
	callback, err := StartCallbackServer("expected-state", testCallbackPath, callbackServerTestMetadata(false))
	if err != nil {
		t.Fatalf("StartCallbackServer() returned error: %v", err)
	}
	defer callback.Close()

	response := sendCallbackRequest(t, callback, url.Values{
		"state":             {"expected-state"},
		"error":             {"access_denied"},
		"error_description": {"user denied access"},
		"error_uri":         {"https://auth.example.com/errors/access-denied"},
	})
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("OAuth error callback status = %d, expected HTTP 400", response.StatusCode)
	}
	_ = response.Body.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = callback.WaitForCallback(ctx)
	var oauthErr *OAuthError
	if !errors.As(err, &oauthErr) {
		t.Fatalf("WaitForCallback() error = %v, expected *OAuthError", err)
	}
	if oauthErr.Operation != "authorization" || oauthErr.ErrorCode != "access_denied" ||
		oauthErr.ErrorDescription != "user denied access" || oauthErr.ErrorURI != "https://auth.example.com/errors/access-denied" {
		t.Errorf("OAuthError = %+v", oauthErr)
	}
	for _, want := range []string{"access_denied", "user denied access", "https://auth.example.com/errors/access-denied"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not contain %q", err, want)
		}
	}
}

func TestCallbackServerWaitsForResponseFlush(t *testing.T) {
	callback, err := StartCallbackServer("expected-state", testCallbackPath, callbackServerTestMetadata(false))
	if err != nil {
		t.Fatalf("StartCallbackServer() returned error: %v", err)
	}
	defer callback.Close()

	allowFlush := make(chan struct{}, 1)
	defer func() {
		select {
		case allowFlush <- struct{}{}:
		default:
		}
	}()

	recorder := httptest.NewRecorder()
	responseWriter := &blockingCallbackResponseWriter{
		ResponseRecorder: recorder,
		flushStarted:     make(chan struct{}),
		allowFlush:       allowFlush,
	}
	requestURL := callback.RedirectURI() + "?" + url.Values{
		"state": {"expected-state"},
		"code":  {"authorization-code"},
	}.Encode()
	request := httptest.NewRequest(http.MethodGet, requestURL, http.NoBody)
	handlerDone := make(chan struct{})
	go func() {
		callback.handleCallback(responseWriter, request)
		close(handlerDone)
	}()

	select {
	case <-responseWriter.flushStarted:
	case <-time.After(time.Second):
		t.Fatal("callback handler did not flush its response")
	}
	select {
	case <-callback.done:
		t.Fatal("callback result was published before the response was flushed")
	default:
	}

	allowFlush <- struct{}{}
	select {
	case <-callback.done:
	case <-time.After(time.Second):
		t.Fatal("callback result was not published after the response was flushed")
	}
	if code, err := callback.WaitForCallback(context.Background()); err != nil || code != "authorization-code" {
		t.Errorf("WaitForCallback() = (%q, %v), expected authorization code", code, err)
	}
	if recorder.Code != http.StatusOK || recorder.Body.String() != "Authorization successful. You may close this window." {
		t.Errorf("callback response = (%d, %q)", recorder.Code, recorder.Body.String())
	}
	select {
	case <-handlerDone:
	case <-time.After(time.Second):
		t.Fatal("callback handler did not finish")
	}
}

func TestCallbackServerWaitHonorsContextCancellation(t *testing.T) {
	callback, err := StartCallbackServer("expected-state", testCallbackPath, callbackServerTestMetadata(false))
	if err != nil {
		t.Fatalf("StartCallbackServer() returned error: %v", err)
	}
	defer callback.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := callback.WaitForCallback(ctx); !errors.Is(err, context.Canceled) {
		t.Errorf("WaitForCallback() error = %v, expected context.Canceled", err)
	}
}

func TestCallbackServerCloseReleasesListener(t *testing.T) {
	callback, err := StartCallbackServer("expected-state", testCallbackPath, callbackServerTestMetadata(false))
	if err != nil {
		t.Fatalf("StartCallbackServer() returned error: %v", err)
	}
	address := callback.listener.Addr().String()
	if err := callback.Close(); err != nil {
		t.Fatalf("Close() returned error: %v", err)
	}
	if _, err := callback.WaitForCallback(context.Background()); !errors.Is(err, ErrCallbackServerClosed) {
		t.Errorf("WaitForCallback() error = %v, expected ErrCallbackServerClosed", err)
	}

	conn, err := net.DialTimeout("tcp", address, 100*time.Millisecond)
	if err == nil {
		_ = conn.Close()
		t.Fatalf("callback listener at %s accepted a connection after Close()", address)
	}
}

func sendCallbackRequest(t *testing.T, callback *CallbackServer, query url.Values) *http.Response {
	t.Helper()
	requestURL := callback.RedirectURI() + "?" + query.Encode()
	client := &http.Client{Timeout: time.Second}
	resp, err := client.Get(requestURL)
	if err != nil {
		t.Fatalf("callback request failed: %v", err)
	}
	return resp
}

type blockingCallbackResponseWriter struct {
	*httptest.ResponseRecorder
	flushStarted chan struct{}
	allowFlush   <-chan struct{}
}

func (w *blockingCallbackResponseWriter) Flush() {
	close(w.flushStarted)
	<-w.allowFlush
	w.ResponseRecorder.Flush()
}
