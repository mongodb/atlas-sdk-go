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
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestStartCallbackServerUsesEphemeralLoopbackURI(t *testing.T) {
	callback, err := StartCallbackServer("expected-state")
	if err != nil {
		t.Fatalf("StartCallbackServer() returned error: %v", err)
	}
	defer callback.Close()

	u, err := url.Parse(callback.RedirectURI())
	if err != nil {
		t.Fatalf("RedirectURI() returned invalid URL: %v", err)
	}
	if u.Scheme != "http" || u.Hostname() != "127.0.0.1" || u.Path != callbackPath {
		t.Errorf("redirect URI = %q, expected loopback HTTP callback at %s", u, callbackPath)
	}
	if callback.Port() == 0 || u.Port() == "" {
		t.Errorf("callback port = %d, expected an OS-assigned port", callback.Port())
	}
}

func TestCallbackServerRejectsMismatchedStateWithoutConsumingFlow(t *testing.T) {
	callback, err := StartCallbackServer("expected-state")
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
	callback, err := StartCallbackServer("expected-state")
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

func TestCallbackServerWaitHonorsContextCancellation(t *testing.T) {
	callback, err := StartCallbackServer("expected-state")
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
	callback, err := StartCallbackServer("expected-state")
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
