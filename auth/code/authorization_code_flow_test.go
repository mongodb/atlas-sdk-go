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
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func authServerMetadataForConfig(c *Config, authorizationEndpoint, tokenEndpoint, revocationEndpoint string) *AuthServerMetadata {
	return &AuthServerMetadata{
		Issuer:                c.AuthServerURL.String(),
		AuthorizationEndpoint: authorizationEndpoint,
		TokenEndpoint:         tokenEndpoint,
		RevocationEndpoint:    revocationEndpoint,
	}
}

func TestGeneratePKCE(t *testing.T) {
	pkce, err := GeneratePKCE()
	if err != nil {
		t.Fatalf("GeneratePKCE() returned error: %v", err)
	}
	if len(pkce.CodeVerifier) != 43 {
		t.Errorf("verifier length = %d, expected 43", len(pkce.CodeVerifier))
	}
	hash := sha256.Sum256([]byte(pkce.CodeVerifier))
	wantChallenge := base64.RawURLEncoding.EncodeToString(hash[:])
	if pkce.CodeChallenge != wantChallenge {
		t.Errorf("challenge = %q, expected S256(verifier) %q", pkce.CodeChallenge, wantChallenge)
	}
	if len(pkce.CodeChallenge) != 43 {
		t.Errorf("challenge length = %d, expected 43", len(pkce.CodeChallenge))
	}

	other, err := GeneratePKCE()
	if err != nil {
		t.Fatalf("second GeneratePKCE() returned error: %v", err)
	}
	if pkce.CodeVerifier == other.CodeVerifier {
		t.Error("two calls generated the same code verifier")
	}
}

func TestGenerateState(t *testing.T) {
	state, err := GenerateState()
	if err != nil {
		t.Fatalf("GenerateState() returned error: %v", err)
	}
	if len(state) != 43 {
		t.Errorf("state length = %d, expected 43", len(state))
	}
	other, err := GenerateState()
	if err != nil {
		t.Fatalf("second GenerateState() returned error: %v", err)
	}
	if state == other {
		t.Error("two calls generated the same state")
	}
}

func TestConfig_DiscoverProductionAuthServer(t *testing.T) {
	t.Setenv(authServerURLEnvVar, "")
	const metadataJSON = `{
		"issuer": "https://authorize.mongodb.com",
		"authorization_endpoint": "https://cloud.mongodb.com/oauth/authorize",
		"token_endpoint": "https://authorize.mongodb.com/tokens",
		"revocation_endpoint": "https://authorize.mongodb.com/tokens/revoke",
		"code_challenge_methods_supported": ["S256"]
	}`
	var requestedURL string
	client := &http.Client{Transport: authServerRoundTripper(func(r *http.Request) (*http.Response, error) {
		requestedURL = r.URL.String()
		header := make(http.Header)
		header.Set("Cache-Control", "max-age=3600")
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     header,
			Body:       io.NopCloser(strings.NewReader(metadataJSON)),
			Request:    r,
		}, nil
	})}
	c := NewConfig(client)
	metadata, resp, err := c.DiscoverAuthServer(ctx)
	if err != nil {
		t.Fatalf("DiscoverAuthServer() returned error: %v", err)
	}
	if got, want := requestedURL, defaultAuthServerURL+"/.well-known/oauth-authorization-server"; got != want {
		t.Errorf("discovery URL = %q, expected %q", got, want)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("response = %v, expected HTTP 200", resp)
	}
	if metadata.Issuer != defaultAuthServerURL {
		t.Errorf("issuer = %q, expected %q", metadata.Issuer, defaultAuthServerURL)
	}
	if metadata.AuthorizationEndpoint != "https://cloud.mongodb.com/oauth/authorize" {
		t.Errorf("authorization endpoint = %q", metadata.AuthorizationEndpoint)
	}
	if metadata.TokenEndpoint != "https://authorize.mongodb.com/tokens" {
		t.Errorf("token endpoint = %q", metadata.TokenEndpoint)
	}
	if metadata.RevocationEndpoint != "https://authorize.mongodb.com/tokens/revoke" {
		t.Errorf("revocation endpoint = %q", metadata.RevocationEndpoint)
	}
	if time.Until(metadata.CacheExpiry) < 3590*time.Second || time.Until(metadata.CacheExpiry) > 3610*time.Second {
		t.Errorf("metadata cache expiry = %v, expected about one hour from now", metadata.CacheExpiry)
	}
}

func TestConfig_DiscoverAuthServerFromEnvironment(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, expected GET", r.Method)
		}
		if r.URL.Path != "/.well-known/oauth-authorization-server" {
			t.Errorf("path = %q, expected OAuth well-known path", r.URL.Path)
		}
		if got := r.Header.Get("Accept"); got != "application/json" {
			t.Errorf("Accept = %q, expected application/json", got)
		}
		w.Header().Set("Cache-Control", "max-age=1800")
		fmt.Fprintf(w, `{
			"issuer": %q,
			"authorization_endpoint": %q,
			"token_endpoint": %q,
			"code_challenge_methods_supported": ["S256"]
		}`, server.URL, "https://auth.example.com/authorize", "https://auth.example.com/token")
	}))
	defer server.Close()

	t.Setenv(authServerURLEnvVar, server.URL)
	c := NewConfig(server.Client())
	metadata, resp, err := c.DiscoverAuthServer(ctx)
	if err != nil {
		t.Fatalf("DiscoverAuthServer() returned error: %v", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("response = %v, expected HTTP 200", resp)
	}
	if metadata.Issuer != server.URL {
		t.Errorf("issuer = %q, expected %q", metadata.Issuer, server.URL)
	}
	if metadata.AuthorizationEndpoint != "https://auth.example.com/authorize" {
		t.Errorf("authorization endpoint = %q", metadata.AuthorizationEndpoint)
	}
	if metadata.TokenEndpoint != "https://auth.example.com/token" {
		t.Errorf("token endpoint = %q", metadata.TokenEndpoint)
	}
	if time.Until(metadata.CacheExpiry) < 1790*time.Second || time.Until(metadata.CacheExpiry) > 1810*time.Second {
		t.Errorf("metadata cache expiry = %v, expected about 30 minutes from now", metadata.CacheExpiry)
	}
	if len(metadata.CodeChallengeMethodsSupported) != 1 || metadata.CodeChallengeMethodsSupported[0] != "S256" {
		t.Errorf("code challenge methods = %v, expected [S256]", metadata.CodeChallengeMethodsSupported)
	}
}

func TestAuthServerMetadataCacheExpiryFallbacks(t *testing.T) {
	t.Run("no cache headers", func(t *testing.T) {
		before := time.Now()
		expiry := authServerMetadataCacheExpiry(make(http.Header))
		if expiry.Before(before.Add(defaultAuthServerMetadataCacheTTL-time.Second)) || expiry.After(before.Add(defaultAuthServerMetadataCacheTTL+time.Second)) {
			t.Errorf("cache expiry = %v, expected about one week from now", expiry)
		}
	})

	t.Run("no-store", func(t *testing.T) {
		header := make(http.Header)
		header.Set("Cache-Control", "no-store")
		if expiry := authServerMetadataCacheExpiry(header); time.Since(expiry) > time.Second {
			t.Errorf("no-store cache expiry = %v, expected now", expiry)
		}
	})
}

func TestAuthServerMetadataCacheExpiry(t *testing.T) {
	t.Run("defaults to one week", func(t *testing.T) {
		before := time.Now()
		expiry := authServerMetadataCacheExpiry(make(http.Header))
		if expiry.Before(before.Add(defaultAuthServerMetadataCacheTTL-time.Second)) || expiry.After(before.Add(defaultAuthServerMetadataCacheTTL+time.Second)) {
			t.Errorf("cache expiry = %v, expected about one week from now", expiry)
		}
	})

	t.Run("no-store is immediately stale", func(t *testing.T) {
		header := make(http.Header)
		header.Set("Cache-Control", "no-store")
		if expiry := authServerMetadataCacheExpiry(header); time.Since(expiry) > time.Second {
			t.Errorf("no-store cache expiry = %v, expected now", expiry)
		}
	})
}

type authServerRoundTripper func(*http.Request) (*http.Response, error)

func (f authServerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestConfig_DiscoverAuthServerErrors(t *testing.T) {
	t.Run("issuer mismatch", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprint(w, `{"issuer":"https://unexpected.example.com","authorization_endpoint":"https://auth.example.com/authorize","token_endpoint":"https://auth.example.com/token"}`)
		}))
		defer server.Close()

		c, _ := NewConfigWithOptions(server.Client(), SetAuthServerURL(server.URL))
		_, _, err := c.DiscoverAuthServer(ctx)
		if err == nil || !strings.Contains(err.Error(), "does not match configured issuer") {
			t.Fatalf("DiscoverAuthServer() error = %v, expected issuer mismatch", err)
		}
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		c, _ := NewConfigWithOptions(server.Client(), SetAuthServerURL(server.URL))
		_, resp, err := c.DiscoverAuthServer(ctx)
		if err == nil || !strings.Contains(err.Error(), "HTTP 500") {
			t.Fatalf("DiscoverAuthServer() error = %v, expected HTTP 500", err)
		}
		if resp == nil || resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("response = %v, expected HTTP 500", resp)
		}
	})
}

func TestConfig_DiscoverAuthServerReportsInvalidEnvironmentURL(t *testing.T) {
	t.Setenv(authServerURLEnvVar, "https://[invalid")
	c := NewConfig(nil)
	_, _, err := c.DiscoverAuthServer(ctx)
	if err == nil || !strings.Contains(err.Error(), authServerURLEnvVar) {
		t.Fatalf("DiscoverAuthServer() error = %v, expected invalid environment URL error", err)
	}
}

func TestConfig_RejectsHTTPAuthServerURLs(t *testing.T) {
	t.Setenv(authServerURLEnvVar, "http://auth.example.com")
	c := NewConfig(nil)
	if _, _, err := c.DiscoverAuthServer(ctx); err == nil || !strings.Contains(err.Error(), "HTTPS") {
		t.Fatalf("DiscoverAuthServer() error = %v, expected HTTP issuer to be rejected", err)
	}
	c.ClientID = "test-client"
	pkce, err := GeneratePKCE()
	if err != nil {
		t.Fatalf("GeneratePKCE() returned error: %v", err)
	}
	metadata := &AuthServerMetadata{
		Issuer:                c.AuthServerURL.String(),
		AuthorizationEndpoint: "https://auth.example.com/authorize",
		TokenEndpoint:         "https://auth.example.com/tokens",
	}
	if _, err := c.AuthorizationURL(metadata, "http://127.0.0.1/callback", "state", pkce); err == nil || !strings.Contains(err.Error(), "HTTPS") {
		t.Errorf("AuthorizationURL() error = %v, expected HTTP issuer to be rejected", err)
	}

	if _, err := NewConfigWithOptions(nil, SetAuthServerURL("http://auth.example.com")); err == nil || !strings.Contains(err.Error(), "HTTPS") {
		t.Errorf("SetAuthServerURL() error = %v, expected HTTP issuer to be rejected", err)
	}
}

func TestConfig_RejectsHTTPAuthorizationServerEndpoints(t *testing.T) {
	t.Setenv(authServerURLEnvVar, "")
	c, err := NewConfigWithOptions(nil, SetClientID("test-client"))
	if err != nil {
		t.Fatalf("NewConfigWithOptions() returned error: %v", err)
	}
	pkce, err := GeneratePKCE()
	if err != nil {
		t.Fatalf("GeneratePKCE() returned error: %v", err)
	}
	metadata := authServerMetadataForConfig(
		c,
		"http://auth.example.com/authorize",
		"http://auth.example.com/tokens",
		"http://auth.example.com/tokens/revoke",
	)

	if _, err := c.AuthorizationURL(metadata, "http://127.0.0.1/callback", "state", pkce); err == nil || !strings.Contains(err.Error(), "HTTPS") {
		t.Errorf("AuthorizationURL() error = %v, expected HTTP endpoint to be rejected", err)
	}
	if _, _, err := c.ExchangeCode(ctx, metadata, "code", "http://127.0.0.1/callback", pkce.CodeVerifier); err == nil || !strings.Contains(err.Error(), "HTTPS") {
		t.Errorf("ExchangeCode() error = %v, expected HTTP endpoint to be rejected", err)
	}
	if _, _, err := c.RefreshAccessToken(ctx, metadata, "refresh-token"); err == nil || !strings.Contains(err.Error(), "HTTPS") {
		t.Errorf("RefreshAccessToken() error = %v, expected HTTP endpoint to be rejected", err)
	}
	if _, err := c.RevokeAuthServerToken(ctx, metadata, "refresh-token", "refresh_token"); err == nil || !strings.Contains(err.Error(), "HTTPS") {
		t.Errorf("RevokeAuthServerToken() error = %v, expected HTTP endpoint to be rejected", err)
	}
}

func TestConfig_AuthorizationURL(t *testing.T) {
	t.Setenv(authServerURLEnvVar, "")
	c, err := NewConfigWithOptions(nil,
		SetClientID("test-client"),
		SetScopes([]string{"atlas", "openid"}),
	)
	if err != nil {
		t.Fatalf("NewConfigWithOptions() returned error: %v", err)
	}

	pkce, err := GeneratePKCE()
	if err != nil {
		t.Fatalf("GeneratePKCE() returned error: %v", err)
	}
	metadata := authServerMetadataForConfig(c,
		"https://auth.example.com/authorize?tenant=atlas",
		"https://auth.example.com/token",
		"https://auth.example.com/revoke",
	)
	got, err := c.AuthorizationURL(
		metadata,
		"http://127.0.0.1:1234/callback",
		"test-state",
		pkce,
	)
	if err != nil {
		t.Fatalf("AuthorizationURL() returned error: %v", err)
	}

	u, err := url.Parse(got)
	if err != nil {
		t.Fatalf("AuthorizationURL() returned invalid URL %q: %v", got, err)
	}
	want := map[string]string{
		"tenant":                "atlas",
		"response_type":         "code",
		"client_id":             "test-client",
		"redirect_uri":          "http://127.0.0.1:1234/callback",
		"scope":                 "atlas openid",
		"state":                 "test-state",
		"code_challenge":        pkce.CodeChallenge,
		"code_challenge_method": "S256",
	}
	for key, value := range want {
		if got := u.Query().Get(key); got != value {
			t.Errorf("query[%q] = %q, expected %q", key, got, value)
		}
	}
}

func TestConfig_AuthorizationURLRequiresStateAndPKCE(t *testing.T) {
	t.Setenv(authServerURLEnvVar, "")
	c, _ := NewConfigWithOptions(nil, SetClientID("test-client"))
	metadata := authServerMetadataForConfig(c,
		"https://auth.example.com/authorize",
		"https://auth.example.com/token",
		"https://auth.example.com/revoke",
	)
	pkce := &PKCEParams{CodeVerifier: "verifier", CodeChallenge: "challenge"}

	for _, test := range []struct {
		name  string
		state string
		pkce  *PKCEParams
	}{
		{name: "missing state", state: "", pkce: pkce},
		{name: "missing PKCE", state: "state", pkce: nil},
		{name: "missing verifier", state: "state", pkce: &PKCEParams{CodeChallenge: "challenge"}},
		{name: "mismatched PKCE", state: "state", pkce: &PKCEParams{CodeVerifier: "verifier", CodeChallenge: "challenge"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := c.AuthorizationURL(metadata, "http://127.0.0.1/callback", test.state, test.pkce)
			if err == nil {
				t.Fatal("AuthorizationURL() returned nil error")
			}
		})
	}
}

func TestConfig_AuthorizationURLRejectsMetadataFromDifferentIssuer(t *testing.T) {
	t.Setenv(authServerURLEnvVar, "")
	c, _ := NewConfigWithOptions(nil, SetClientID("test-client"))
	pkce, err := GeneratePKCE()
	if err != nil {
		t.Fatalf("GeneratePKCE() returned error: %v", err)
	}
	metadata := &AuthServerMetadata{
		Issuer:                "https://other.example.com",
		AuthorizationEndpoint: "https://auth.example.com/authorize",
		TokenEndpoint:         "https://auth.example.com/tokens",
	}
	if _, err := c.AuthorizationURL(metadata, "http://127.0.0.1/callback", "state", pkce); err == nil || !strings.Contains(err.Error(), "does not match configured issuer") {
		t.Fatalf("AuthorizationURL() error = %v, expected issuer mismatch", err)
	}
}

func TestConfig_ExchangeCode(t *testing.T) {
	var gotForm url.Values
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, expected POST", r.Method)
		}
		if got := r.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
			t.Errorf("Content-Type = %q", got)
		}
		if got := r.Header.Get("Accept"); got != "application/json" {
			t.Errorf("Accept = %q", got)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("ParseForm() returned error: %v", err)
		}
		gotForm = r.PostForm
		fmt.Fprint(w, `{"access_token":"access-token","refresh_token":"refresh-token","scope":"atlas","id_token":"id-token","token_type":"Bearer","expires_in":3600}`)
	}))
	defer server.Close()

	c, err := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	if err != nil {
		t.Fatalf("NewConfigWithOptions() returned error: %v", err)
	}
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL, server.URL+"/revoke")
	before := time.Now()
	token, resp, err := c.ExchangeCode(ctx, metadata, "authorization-code", "http://127.0.0.1/callback", "code-verifier")
	if err != nil {
		t.Fatalf("ExchangeCode() returned error: %v", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("response = %v, expected HTTP 200", resp)
	}
	if token.AccessToken != "access-token" || token.RefreshToken != "refresh-token" || token.IDToken != "id-token" {
		t.Errorf("token = %+v", token)
	}
	if token.Expiry.Before(before.Add(3590*time.Second)) || token.Expiry.After(before.Add(3610*time.Second)) {
		t.Errorf("token expiry %v outside expected range", token.Expiry)
	}
	wantForm := map[string]string{
		"grant_type":    "authorization_code",
		"code":          "authorization-code",
		"redirect_uri":  "http://127.0.0.1/callback",
		"client_id":     "test-client",
		"code_verifier": "code-verifier",
	}
	for key, value := range wantForm {
		if got := gotForm.Get(key); got != value {
			t.Errorf("form[%q] = %q, expected %q", key, got, value)
		}
	}
}

func TestConfig_ExchangeCodeErrorResponse(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"invalid_grant","error_description":"authorization code expired","error_uri":"https://auth.example.com/errors"}`)
	}))
	defer server.Close()

	c, _ := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL, server.URL+"/revoke")
	_, resp, err := c.ExchangeCode(ctx, metadata, "expired-code", "http://127.0.0.1/callback", "verifier")
	if err == nil {
		t.Fatal("ExchangeCode() returned nil error for OAuth error response")
	}
	if resp == nil || resp.StatusCode != http.StatusBadRequest {
		t.Errorf("response = %v, expected HTTP 400", resp)
	}
	var oauthErr *OAuthError
	if !errors.As(err, &oauthErr) {
		t.Fatalf("error type = %T, expected *OAuthError", err)
	}
	if oauthErr.Operation != "token exchange" || oauthErr.ErrorCode != "invalid_grant" || oauthErr.ErrorDescription != "authorization code expired" || oauthErr.ErrorURI != "https://auth.example.com/errors" {
		t.Errorf("OAuthError = %+v", oauthErr)
	}
	for _, want := range []string{"invalid_grant", "authorization code expired", "https://auth.example.com/errors"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not include %q", err, want)
		}
	}
}

func TestConfig_ExchangeCodeUnexpectedErrorResponse(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "upstream unavailable")
	}))
	defer server.Close()

	c, _ := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL, server.URL+"/revoke")
	_, _, err := c.ExchangeCode(ctx, metadata, "authorization-code", "http://127.0.0.1/callback", "verifier")
	if err == nil || !strings.Contains(err.Error(), "HTTP 502") {
		t.Fatalf("ExchangeCode() error = %v, expected HTTP 502 error", err)
	}
	var oauthErr *OAuthError
	if errors.As(err, &oauthErr) {
		t.Errorf("error = %T, expected generic status error", err)
	}
}

func TestConfig_ExchangeCodeDoesNotFollowRedirects(t *testing.T) {
	var redirected atomic.Bool
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/redirected" {
			redirected.Store(true)
			fmt.Fprint(w, `{"access_token":"unexpected"}`)
			return
		}
		http.Redirect(w, r, "/redirected", http.StatusTemporaryRedirect)
	}))
	defer server.Close()

	c, err := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	if err != nil {
		t.Fatalf("NewConfigWithOptions() returned error: %v", err)
	}
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL+"/token", server.URL+"/revoke")
	_, resp, err := c.ExchangeCode(ctx, metadata, "authorization-code", "http://127.0.0.1/callback", "secret-verifier")
	if err == nil {
		t.Fatal("ExchangeCode() followed a redirect from the token endpoint")
	}
	if resp == nil || resp.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("response = %v, expected HTTP 307", resp)
	}
	if redirected.Load() {
		t.Error("OAuth token request body was replayed to the redirect target")
	}
}

func TestDecodeOAuthTokenRequiresAccessTokenAndExpiry(t *testing.T) {
	for _, test := range []struct {
		name string
		body string
	}{
		{name: "missing access token", body: `{"expires_in":3600}`},
		{name: "missing expiry", body: `{"access_token":"at"}`},
		{name: "zero expiry", body: `{"access_token":"at","expires_in":0}`},
		{name: "negative expiry", body: `{"access_token":"at","expires_in":-1}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := decodeOAuthToken([]byte(test.body), "test exchange"); err == nil {
				t.Fatal("decodeOAuthToken() returned nil error for an incomplete token response")
			}
		})
	}
}

func TestConfig_RefreshAccessToken(t *testing.T) {
	var gotForm url.Values
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, expected POST", r.Method)
		}
		if got := r.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
			t.Errorf("Content-Type = %q", got)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("ParseForm() returned error: %v", err)
		}
		gotForm = r.PostForm
		fmt.Fprint(w, `{"access_token":"new-access-token","refresh_token":"rotated-refresh-token","token_type":"Bearer","expires_in":1800}`)
	}))
	defer server.Close()

	c, err := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	if err != nil {
		t.Fatalf("NewConfigWithOptions() returned error: %v", err)
	}
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL, server.URL+"/revoke")
	before := time.Now()
	token, resp, err := c.RefreshAccessToken(ctx, metadata, "old-refresh-token")
	if err != nil {
		t.Fatalf("RefreshAccessToken() returned error: %v", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("response = %v, expected HTTP 200", resp)
	}
	if token.AccessToken != "new-access-token" || token.RefreshToken != "rotated-refresh-token" || token.TokenType != "Bearer" {
		t.Errorf("token = %+v", token)
	}
	if token.Expiry.Before(before.Add(1790*time.Second)) || token.Expiry.After(before.Add(1810*time.Second)) {
		t.Errorf("token expiry %v outside expected range", token.Expiry)
	}
	wantForm := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     "test-client",
		"refresh_token": "old-refresh-token",
	}
	for key, value := range wantForm {
		if got := gotForm.Get(key); got != value {
			t.Errorf("form[%q] = %q, expected %q", key, got, value)
		}
	}
}

func TestConfig_RefreshAccessTokenRetainsRefreshTokenWhenNotRotated(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `{"access_token":"new-access-token","token_type":"Bearer","expires_in":1800}`)
	}))
	defer server.Close()

	c, _ := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL, server.URL+"/revoke")
	token, _, err := c.RefreshAccessToken(ctx, metadata, "existing-refresh-token")
	if err != nil {
		t.Fatalf("RefreshAccessToken() returned error: %v", err)
	}
	if token.RefreshToken != "existing-refresh-token" {
		t.Errorf("refresh token = %q, expected existing refresh token to be retained", token.RefreshToken)
	}
}

func TestConfig_RefreshAccessTokenErrorResponse(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"invalid_grant","error_description":"refresh token expired"}`)
	}))
	defer server.Close()

	c, _ := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL, server.URL+"/revoke")
	_, resp, err := c.RefreshAccessToken(ctx, metadata, "expired-refresh-token")
	if err == nil {
		t.Fatal("RefreshAccessToken() returned nil error for OAuth error response")
	}
	if resp == nil || resp.StatusCode != http.StatusBadRequest {
		t.Errorf("response = %v, expected HTTP 400", resp)
	}
	var oauthErr *OAuthError
	if !errors.As(err, &oauthErr) {
		t.Fatalf("error type = %T, expected *OAuthError", err)
	}
	if oauthErr.Operation != "token refresh" || oauthErr.ErrorCode != "invalid_grant" || oauthErr.ErrorDescription != "refresh token expired" {
		t.Errorf("OAuthError = %+v", oauthErr)
	}
}

func TestConfig_RevokeAuthServerToken(t *testing.T) {
	var gotForm url.Values
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, expected POST", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("ParseForm() returned error: %v", err)
		}
		gotForm = r.PostForm
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c, err := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	if err != nil {
		t.Fatalf("NewConfigWithOptions() returned error: %v", err)
	}
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL+"/token", server.URL)
	resp, err := c.RevokeAuthServerToken(ctx, metadata, "refresh-token", "refresh_token")
	if err != nil {
		t.Fatalf("RevokeAuthServerToken() returned error: %v", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("response = %v, expected HTTP 200", resp)
	}
	wantForm := map[string]string{
		"client_id":       "test-client",
		"token":           "refresh-token",
		"token_type_hint": "refresh_token",
	}
	for key, value := range wantForm {
		if got := gotForm.Get(key); got != value {
			t.Errorf("form[%q] = %q, expected %q", key, got, value)
		}
	}
}

func TestConfig_RevokeAuthServerTokenErrorResponse(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"invalid_request","error_description":"token is required"}`)
	}))
	defer server.Close()

	c, _ := NewConfigWithOptions(server.Client(), SetClientID("test-client"), SetAuthServerURL(server.URL))
	metadata := authServerMetadataForConfig(c, server.URL+"/authorize", server.URL+"/token", server.URL)
	resp, err := c.RevokeAuthServerToken(ctx, metadata, "refresh-token", "refresh_token")
	if err == nil {
		t.Fatal("RevokeAuthServerToken() returned nil error for OAuth error response")
	}
	if resp == nil || resp.StatusCode != http.StatusBadRequest {
		t.Errorf("response = %v, expected HTTP 400", resp)
	}
	var oauthErr *OAuthError
	if !errors.As(err, &oauthErr) {
		t.Fatalf("error type = %T, expected *OAuthError", err)
	}
	if oauthErr.Operation != "token revocation" || oauthErr.ErrorCode != "invalid_request" || oauthErr.ErrorDescription != "token is required" {
		t.Errorf("OAuthError = %+v", oauthErr)
	}
}
