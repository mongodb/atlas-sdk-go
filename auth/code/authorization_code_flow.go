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
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/atlas-sdk/v20250312026/internal/core"
)

const wellKnownPath = ".well-known/oauth-authorization-server"

const defaultAuthServerMetadataCacheTTL = 7 * 24 * time.Hour

// AuthServerMetadata holds discovered OAuth endpoints and cache expiry.
type AuthServerMetadata struct {
	Issuer                        string   `json:"issuer"`
	AuthorizationEndpoint         string   `json:"authorization_endpoint"`
	TokenEndpoint                 string   `json:"token_endpoint"`
	RevocationEndpoint            string   `json:"revocation_endpoint,omitempty"`
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported,omitempty"`
	// CacheExpiry is derived from HTTP cache headers and is not part of the AS response.
	CacheExpiry time.Time `json:"-"`
}

// PKCEParams holds an RFC 7636 verifier and its S256 challenge.
type PKCEParams struct {
	CodeVerifier  string
	CodeChallenge string
}

// OAuthError describes an authorization-server error response.
type OAuthError struct {
	Operation        string
	StatusCode       int
	ErrorCode        string
	ErrorDescription string
	ErrorURI         string
}

func (e *OAuthError) Error() string {
	operation := e.Operation
	if operation == "" {
		operation = "request"
	}
	message := fmt.Sprintf("OAuth %s failed", operation)
	if e.StatusCode > 0 {
		message += fmt.Sprintf(" (HTTP %d)", e.StatusCode)
	}
	if e.ErrorCode != "" {
		message += ": " + e.ErrorCode
	}
	if e.ErrorDescription != "" {
		message += ": " + e.ErrorDescription
	}
	if e.ErrorURI != "" {
		message += " (see " + e.ErrorURI + ")"
	}
	return message
}

// GeneratePKCE creates a random verifier and its S256 challenge.
func GeneratePKCE() (*PKCEParams, error) {
	verifierBytes := make([]byte, 32)
	if _, err := rand.Read(verifierBytes); err != nil {
		return nil, fmt.Errorf("generating PKCE code verifier: %w", err)
	}

	verifier := base64.RawURLEncoding.EncodeToString(verifierBytes)
	return &PKCEParams{
		CodeVerifier:  verifier,
		CodeChallenge: s256CodeChallenge(verifier),
	}, nil
}

// GenerateState creates random OAuth state for the authorization request.
func GenerateState() (string, error) {
	stateBytes := make([]byte, 32)
	if _, err := rand.Read(stateBytes); err != nil {
		return "", fmt.Errorf("generating OAuth state: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(stateBytes), nil
}

// DiscoverAuthServer fetches and validates metadata for the configured issuer.
func (c *Config) DiscoverAuthServer(ctx context.Context) (*AuthServerMetadata, *core.Response, error) {
	if c.AuthServerURL == nil {
		if c.authServerURLError != nil {
			return nil, nil, fmt.Errorf("invalid %s: %w", authServerURLEnvVar, c.authServerURLError)
		}
		return nil, nil, errors.New("OAuth Authorization Server issuer URL is not configured")
	}
	if _, err := parseOAuthEndpoint(c.AuthServerURL.String(), "Authorization Server issuer"); err != nil {
		return nil, nil, err
	}

	discoveryURL := c.AuthServerURL.ResolveReference(&url.URL{Path: wellKnownPath})
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, discoveryURL.String(), http.NoBody)
	if err != nil {
		return nil, nil, fmt.Errorf("creating OAuth discovery request: %w", err)
	}
	c.setOAuthRequestHeaders(req)

	httpResp, err := c.doOAuthRequest(req)
	if err != nil {
		if ctx.Err() != nil {
			return nil, nil, ctx.Err()
		}
		return nil, nil, fmt.Errorf("OAuth discovery request failed: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &core.Response{Response: httpResp}
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, resp, fmt.Errorf("reading OAuth discovery response: %w", err)
	}
	if c.withRaw {
		resp.Raw = body
	}
	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		return nil, resp, fmt.Errorf("OAuth discovery endpoint returned HTTP %d", httpResp.StatusCode)
	}

	var metadata AuthServerMetadata
	if err := json.Unmarshal(body, &metadata); err != nil {
		return nil, resp, fmt.Errorf("decoding OAuth discovery response: %w", err)
	}
	if err := c.validateAuthServerMetadata(&metadata); err != nil {
		return nil, resp, err
	}
	metadata.CacheExpiry = authServerMetadataCacheExpiry(httpResp.Header)

	return &metadata, resp, nil
}

func (c *Config) validateAuthServerMetadata(metadata *AuthServerMetadata) error {
	if c.AuthServerURL == nil {
		if c.authServerURLError != nil {
			return fmt.Errorf("invalid %s: %w", authServerURLEnvVar, c.authServerURLError)
		}
		return errors.New("OAuth Authorization Server issuer URL is not configured")
	}
	if metadata == nil {
		return errors.New("OAuth Authorization Server metadata is required")
	}
	if _, err := parseOAuthEndpoint(c.AuthServerURL.String(), "Authorization Server issuer"); err != nil {
		return err
	}
	if metadata.Issuer == "" {
		return errors.New("OAuth discovery response is missing issuer")
	}
	if metadata.Issuer != c.AuthServerURL.String() {
		return fmt.Errorf("OAuth discovery issuer %q does not match configured issuer %q", metadata.Issuer, c.AuthServerURL)
	}
	if _, err := parseOAuthEndpoint(metadata.AuthorizationEndpoint, "authorization endpoint"); err != nil {
		return err
	}
	if _, err := parseOAuthEndpoint(metadata.TokenEndpoint, "token endpoint"); err != nil {
		return err
	}
	if metadata.RevocationEndpoint != "" {
		if _, err := parseOAuthEndpoint(metadata.RevocationEndpoint, "revocation endpoint"); err != nil {
			return err
		}
	}
	return nil
}

// AuthorizationURL builds a request from discovered metadata, state, and PKCE.
// Retain and validate state, and keep the verifier until code exchange.
func (c *Config) AuthorizationURL(metadata *AuthServerMetadata, redirectURI, state string, pkce *PKCEParams) (string, error) {
	if err := c.validateAuthServerMetadata(metadata); err != nil {
		return "", err
	}
	u, err := parseOAuthEndpoint(metadata.AuthorizationEndpoint, "authorization endpoint")
	if err != nil {
		return "", err
	}
	if c.ClientID == "" {
		return "", errors.New("OAuth client ID is not configured")
	}
	if redirectURI == "" {
		return "", errors.New("OAuth redirect URI is required")
	}
	if state == "" {
		return "", errors.New("OAuth state is required")
	}
	if pkce == nil || pkce.CodeVerifier == "" || pkce.CodeChallenge == "" {
		return "", errors.New("PKCE code verifier and challenge are required")
	}
	if pkce.CodeChallenge != s256CodeChallenge(pkce.CodeVerifier) {
		return "", errors.New("PKCE code challenge does not match the S256 code verifier")
	}

	query := u.Query()
	query.Set("response_type", "code")
	query.Set("client_id", c.ClientID)
	query.Set("redirect_uri", redirectURI)
	if len(c.Scopes) > 0 {
		query.Set("scope", strings.Join(c.Scopes, " "))
	}
	query.Set("state", state)
	query.Set("code_challenge", pkce.CodeChallenge)
	query.Set("code_challenge_method", "S256")
	u.RawQuery = query.Encode()

	return u.String(), nil
}

// ExchangeCode exchanges an authorization code using its redirect URI and verifier.
func (c *Config) ExchangeCode(ctx context.Context, metadata *AuthServerMetadata, authorizationCode, redirectURI, codeVerifier string) (*Token, *core.Response, error) {
	if err := c.validateAuthServerMetadata(metadata); err != nil {
		return nil, nil, err
	}
	if c.ClientID == "" {
		return nil, nil, errors.New("OAuth client ID is not configured")
	}
	if authorizationCode == "" {
		return nil, nil, errors.New("authorization code is required")
	}
	if redirectURI == "" {
		return nil, nil, errors.New("OAuth redirect URI is required")
	}
	if codeVerifier == "" {
		return nil, nil, errors.New("PKCE code verifier is required")
	}

	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {authorizationCode},
		"redirect_uri":  {redirectURI},
		"client_id":     {c.ClientID},
		"code_verifier": {codeVerifier},
	}
	resp, body, err := c.postOAuthForm(ctx, metadata.TokenEndpoint, "token exchange", form)
	if err != nil {
		return nil, resp, err
	}
	token, err := decodeOAuthToken(body, "token exchange")
	if err != nil {
		return nil, resp, err
	}
	return token, resp, nil
}

// RefreshAccessToken exchanges a refresh token and retains it if not rotated.
func (c *Config) RefreshAccessToken(ctx context.Context, metadata *AuthServerMetadata, refreshToken string) (*Token, *core.Response, error) {
	if err := c.validateAuthServerMetadata(metadata); err != nil {
		return nil, nil, err
	}
	if c.ClientID == "" {
		return nil, nil, errors.New("OAuth client ID is not configured")
	}
	if refreshToken == "" {
		return nil, nil, errors.New("OAuth refresh token is required")
	}

	form := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {c.ClientID},
		"refresh_token": {refreshToken},
	}
	resp, body, err := c.postOAuthForm(ctx, metadata.TokenEndpoint, "token refresh", form)
	if err != nil {
		return nil, resp, err
	}
	token, err := decodeOAuthToken(body, "token refresh")
	if err != nil {
		return nil, resp, err
	}
	if token.RefreshToken == "" {
		token.RefreshToken = refreshToken
	}
	return token, resp, nil
}

// RevokeAuthServerToken revokes a token at the discovered endpoint.
func (c *Config) RevokeAuthServerToken(ctx context.Context, metadata *AuthServerMetadata, token, tokenTypeHint string) (*core.Response, error) {
	if err := c.validateAuthServerMetadata(metadata); err != nil {
		return nil, err
	}
	if c.ClientID == "" {
		return nil, errors.New("OAuth client ID is not configured")
	}
	if token == "" {
		return nil, errors.New("OAuth token to revoke is required")
	}
	if metadata.RevocationEndpoint == "" {
		return nil, errors.New("OAuth discovery response is missing revocation endpoint")
	}

	form := url.Values{
		"client_id": {c.ClientID},
		"token":     {token},
	}
	if tokenTypeHint != "" {
		form.Set("token_type_hint", tokenTypeHint)
	}
	resp, _, err := c.postOAuthForm(ctx, metadata.RevocationEndpoint, "token revocation", form)
	return resp, err
}

func (c *Config) postOAuthForm(ctx context.Context, endpoint, operation string, form url.Values) (*core.Response, []byte, error) {
	u, err := parseOAuthEndpoint(endpoint, operation+" endpoint")
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, fmt.Errorf("creating OAuth %s request: %w", operation, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c.setOAuthRequestHeaders(req)

	httpResp, err := c.doOAuthRequest(req)
	if err != nil {
		if ctx.Err() != nil {
			return nil, nil, ctx.Err()
		}
		return nil, nil, fmt.Errorf("OAuth %s request failed: %w", operation, err)
	}
	defer httpResp.Body.Close()

	resp := &core.Response{Response: httpResp}
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return resp, nil, fmt.Errorf("reading OAuth %s response: %w", operation, err)
	}
	if c.withRaw {
		resp.Raw = body
	}
	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		return resp, body, parseOAuthError(httpResp.StatusCode, operation, body)
	}
	return resp, body, nil
}

func parseOAuthError(statusCode int, operation string, body []byte) error {
	var oauthErr struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		ErrorURI         string `json:"error_uri"`
	}
	if err := json.Unmarshal(body, &oauthErr); err == nil && oauthErr.Error != "" {
		return &OAuthError{
			Operation:        operation,
			StatusCode:       statusCode,
			ErrorCode:        oauthErr.Error,
			ErrorDescription: oauthErr.ErrorDescription,
			ErrorURI:         oauthErr.ErrorURI,
		}
	}
	return fmt.Errorf("OAuth %s endpoint returned HTTP %d", operation, statusCode)
}

func decodeOAuthToken(body []byte, operation string) (*Token, error) {
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, fmt.Errorf("decoding OAuth %s response: %w", operation, err)
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("decoding OAuth %s response: access_token is missing", operation)
	}
	if token.ExpiresIn <= 0 {
		return nil, fmt.Errorf("decoding OAuth %s response: expires_in must be positive", operation)
	}
	token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

	return &token, nil
}

func (c *Config) doOAuthRequest(req *http.Request) (*http.Response, error) {
	client := *c.client
	client.CheckRedirect = func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}
	// #nosec G704 -- OAuth request URLs are configured issuers or HTTPS endpoints from their validated discovery metadata.
	return client.Do(req)
}

func authServerMetadataCacheExpiry(header http.Header) time.Time {
	now := time.Now()
	for _, directive := range strings.Split(header.Get("Cache-Control"), ",") {
		parts := strings.SplitN(strings.TrimSpace(directive), "=", 2)
		if len(parts) == 1 && (strings.EqualFold(parts[0], "no-cache") || strings.EqualFold(parts[0], "no-store")) {
			return now
		}
		if len(parts) != 2 || !strings.EqualFold(strings.TrimSpace(parts[0]), "max-age") {
			continue
		}
		seconds, err := strconv.ParseInt(strings.Trim(strings.TrimSpace(parts[1]), `"`), 10, 64)
		maxCacheAgeSeconds := int64(^uint64(0)>>1) / int64(time.Second)
		if err == nil && seconds >= 0 && seconds <= maxCacheAgeSeconds {
			return now.Add(time.Duration(seconds) * time.Second)
		}
	}

	if expires := header.Get("Expires"); expires != "" {
		if expiry, err := http.ParseTime(expires); err == nil {
			return expiry
		}
	}
	return now.Add(defaultAuthServerMetadataCacheTTL)
}

func (c *Config) setOAuthRequestHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
}

func parseOAuthEndpoint(raw, name string) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, fmt.Errorf("invalid OAuth %s: %w", name, err)
	}
	if u.Scheme != "https" || u.Host == "" {
		return nil, fmt.Errorf("OAuth %s must be an absolute HTTPS URL", name)
	}
	return u, nil
}

func s256CodeChallenge(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}
