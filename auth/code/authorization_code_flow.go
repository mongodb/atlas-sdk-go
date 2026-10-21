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
	"bytes"
	"context"
	"crypto/rand"
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

	"golang.org/x/oauth2"

	"go.mongodb.org/atlas-sdk/v20250312027/internal/core"
)

const wellKnownPath = ".well-known/oauth-authorization-server"

const defaultAuthServerMetadataCacheTTL = 7 * 24 * time.Hour

func authServerDiscoveryURL(issuer *url.URL) (*url.URL, error) {
	discoveryPath := "/" + wellKnownPath
	issuerPath := issuer.EscapedPath()
	if issuerPath != "" && issuerPath != "/" {
		discoveryPath += issuerPath
	}
	path, err := url.PathUnescape(discoveryPath)
	if err != nil {
		return nil, fmt.Errorf("constructing OAuth discovery path: %w", err)
	}

	discoveryURL := *issuer
	discoveryURL.Path = path
	discoveryURL.RawPath = discoveryPath
	discoveryURL.RawQuery = ""
	discoveryURL.Fragment = ""
	discoveryURL.RawFragment = ""
	return &discoveryURL, nil
}

// AuthServerMetadata holds discovered OAuth endpoints and cache expiry.
type AuthServerMetadata struct {
	Issuer                                     string   `json:"issuer"`
	AuthorizationEndpoint                      string   `json:"authorization_endpoint"`
	TokenEndpoint                              string   `json:"token_endpoint"`
	RevocationEndpoint                         string   `json:"revocation_endpoint,omitempty"`
	CodeChallengeMethodsSupported              []string `json:"code_challenge_methods_supported,omitempty"`
	// AuthorizationResponseIssParameterSupported reports whether authorization responses include an iss parameter.
	AuthorizationResponseIssParameterSupported bool     `json:"authorization_response_iss_parameter_supported,omitempty"`
	// CacheExpiry is derived from HTTP cache headers and is not part of the discovery document.
	CacheExpiry time.Time `json:"-"`
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

// GenerateState creates random OAuth state for the authorization request.
func GenerateState() (string, error) {
	stateBytes := make([]byte, 32)
	if _, err := rand.Read(stateBytes); err != nil {
		return "", fmt.Errorf("generating OAuth state: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(stateBytes), nil
}

// GenerateVerifier creates a random PKCE code verifier.
func GenerateVerifier() string {
	return oauth2.GenerateVerifier()
}

// DiscoverAuthServer fetches and validates metadata for the configured issuer.
func (c *Config) DiscoverAuthServer(ctx context.Context) (*AuthServerMetadata, *core.Response, error) {
	if c.AuthServerURL == nil {
		return nil, nil, errors.New("OAuth Authorization Server issuer URL is not configured")
	}
	if _, err := parseOAuthEndpoint(c.AuthServerURL.String(), "Authorization Server issuer"); err != nil {
		return nil, nil, err
	}

	discoveryURL, err := authServerDiscoveryURL(c.AuthServerURL)
	if err != nil {
		return nil, nil, err
	}
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

// AuthorizationURL returns an authorization URL with the supplied state and PKCE verifier.
// Keep both values for callback validation and code exchange.
func (c *Config) AuthorizationURL(metadata *AuthServerMetadata, redirectURI, state, codeVerifier string) (string, error) {
	oauthConfig, err := c.newOAuth2Config(metadata)
	if err != nil {
		return "", err
	}
	if redirectURI == "" {
		return "", errors.New("OAuth redirect URI is required")
	}
	if state == "" {
		return "", errors.New("OAuth state is required")
	}
	if codeVerifier == "" {
		return "", errors.New("PKCE code verifier is required")
	}

	oauthConfig.RedirectURL = redirectURI
	return oauthConfig.AuthCodeURL(state, oauth2.S256ChallengeOption(codeVerifier)), nil
}

// ExchangeCode exchanges an authorization code using its redirect URI and verifier.
func (c *Config) ExchangeCode(ctx context.Context, metadata *AuthServerMetadata, authorizationCode, redirectURI, codeVerifier string) (*Token, *core.Response, error) {
	oauthConfig, err := c.newOAuth2Config(metadata)
	if err != nil {
		return nil, nil, err
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

	oauthConfig.RedirectURL = redirectURI
	oauthCtx, responseCapture := c.oauth2Context(ctx)
	oauthToken, err := oauthConfig.Exchange(oauthCtx, authorizationCode, oauth2.VerifierOption(codeVerifier))
	resp := c.oauth2Response(responseCapture)
	if err != nil {
		if ctx.Err() != nil {
			return nil, resp, ctx.Err()
		}
		return nil, resp, oauth2TokenError(err, "token exchange")
	}
	token, err := tokenFromOAuth2(oauthToken, "token exchange")
	if err != nil {
		return nil, resp, err
	}
	return token, resp, nil
}

// RefreshAccessToken exchanges a refresh token and retains it if not rotated.
func (c *Config) RefreshAccessToken(ctx context.Context, metadata *AuthServerMetadata, refreshToken string) (*Token, *core.Response, error) {
	oauthConfig, err := c.newOAuth2Config(metadata)
	if err != nil {
		return nil, nil, err
	}
	if refreshToken == "" {
		return nil, nil, errors.New("OAuth refresh token is required")
	}

	oauthCtx, responseCapture := c.oauth2Context(ctx)
	tokenSource := oauthConfig.TokenSource(oauthCtx, &oauth2.Token{RefreshToken: refreshToken})
	oauthToken, err := tokenSource.Token()
	resp := c.oauth2Response(responseCapture)
	if err != nil {
		if ctx.Err() != nil {
			return nil, resp, ctx.Err()
		}
		return nil, resp, oauth2TokenError(err, "token refresh")
	}
	token, err := tokenFromOAuth2(oauthToken, "token refresh")
	if err != nil {
		return nil, resp, err
	}
	return token, resp, nil
}

func (c *Config) newOAuth2Config(metadata *AuthServerMetadata) (*oauth2.Config, error) {
	if err := c.validateAuthServerMetadata(metadata); err != nil {
		return nil, err
	}
	if c.ClientID == "" {
		return nil, errors.New("OAuth client ID is not configured")
	}
	return &oauth2.Config{
		ClientID: c.ClientID,
		Scopes:   append([]string(nil), c.Scopes...),
		Endpoint: oauth2.Endpoint{
			AuthURL:   metadata.AuthorizationEndpoint,
			TokenURL:  metadata.TokenEndpoint,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}, nil
}

func oauth2TokenError(err error, operation string) error {
	var retrieveErr *oauth2.RetrieveError
	if errors.As(err, &retrieveErr) {
		statusCode := 0
		if retrieveErr.Response != nil {
			statusCode = retrieveErr.Response.StatusCode
		}
		if retrieveErr.ErrorCode != "" || retrieveErr.ErrorDescription != "" || retrieveErr.ErrorURI != "" {
			return &OAuthError{
				Operation:        operation,
				StatusCode:       statusCode,
				ErrorCode:        retrieveErr.ErrorCode,
				ErrorDescription: retrieveErr.ErrorDescription,
				ErrorURI:         retrieveErr.ErrorURI,
			}
		}
		if statusCode > 0 {
			return fmt.Errorf("OAuth %s endpoint returned HTTP %d", operation, statusCode)
		}
	}
	return fmt.Errorf("OAuth %s request failed: %w", operation, err)
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

func tokenFromOAuth2(oauthToken *oauth2.Token, operation string) (*Token, error) {
	if oauthToken == nil || oauthToken.AccessToken == "" {
		return nil, fmt.Errorf("decoding OAuth %s response: access_token is missing", operation)
	}
	if oauthToken.ExpiresIn <= 0 {
		return nil, fmt.Errorf("decoding OAuth %s response: expires_in must be positive", operation)
	}

	return &Token{
		AccessToken:  oauthToken.AccessToken,
		RefreshToken: oauthToken.RefreshToken,
		Scope:        oauth2StringExtra(oauthToken, "scope"),
		IDToken:      oauth2StringExtra(oauthToken, "id_token"),
		TokenType:    oauthToken.TokenType,
		ExpiresIn:    int(oauthToken.ExpiresIn),
		Expiry:       oauthToken.Expiry,
	}, nil
}

func oauth2StringExtra(token *oauth2.Token, key string) string {
	value, _ := token.Extra(key).(string)
	return value
}

func (c *Config) oauth2HTTPClient() *http.Client {
	client := *c.client
	client.CheckRedirect = func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &client
}

func (c *Config) doOAuthRequest(req *http.Request) (*http.Response, error) {
	// #nosec G704 -- OAuth requests use the configured issuer or validated HTTPS endpoints from its metadata.
	return c.oauth2HTTPClient().Do(req)
}

const oauth2ResponseBodyLimit = 1 << 20

type oauth2ResponseCapture struct {
	response *http.Response
	body     []byte
}

// oauth2ResponseTransport adds SDK headers and captures responses for core.Response.
type oauth2ResponseTransport struct {
	base        http.RoundTripper
	capture     *oauth2ResponseCapture
	captureBody bool
	userAgent   string
}

func (t *oauth2ResponseTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	request := req.Clone(req.Context())
	request.Header = req.Header.Clone()
	if request.Header == nil {
		request.Header = make(http.Header)
	}
	request.Header.Set("Accept", "application/json")
	if t.userAgent != "" {
		request.Header.Set("User-Agent", t.userAgent)
	}

	// #nosec G704 -- OAuth requests are sent to validated HTTPS endpoints from the configured issuer.
	response, err := t.base.RoundTrip(request)
	if err != nil || response == nil {
		return response, err
	}
	t.capture.response = response
	if response.Body == nil || !t.captureBody {
		return response, nil
	}

	body, err := io.ReadAll(io.LimitReader(response.Body, oauth2ResponseBodyLimit))
	if err != nil {
		_ = response.Body.Close()
		return nil, err
	}
	_ = response.Body.Close()
	t.capture.body = body
	response.Body = io.NopCloser(bytes.NewReader(body))
	return response, nil
}

func (c *Config) oauth2Context(ctx context.Context) (context.Context, *oauth2ResponseCapture) {
	client := c.oauth2HTTPClient()
	baseTransport := client.Transport
	if baseTransport == nil {
		baseTransport = http.DefaultTransport
	}
	capture := &oauth2ResponseCapture{}
	client.Transport = &oauth2ResponseTransport{
		base:        baseTransport,
		capture:     capture,
		captureBody: c.withRaw,
		userAgent:   c.UserAgent,
	}
	return context.WithValue(ctx, oauth2.HTTPClient, client), capture
}

func (c *Config) oauth2Response(capture *oauth2ResponseCapture) *core.Response {
	if capture == nil || capture.response == nil {
		return nil
	}
	response := &core.Response{Response: capture.response}
	if c.withRaw {
		response.Raw = capture.body
	}
	return response
}

func authServerMetadataCacheExpiry(header http.Header) time.Time {
	now := time.Now()
	var maxAgeSeconds int64
	hasMaxAge := false
	cacheControl := strings.Join(header.Values("Cache-Control"), ",")
	for _, directive := range strings.Split(cacheControl, ",") {
		parts := strings.SplitN(strings.TrimSpace(directive), "=", 2)
		name := strings.TrimSpace(parts[0])
		if strings.EqualFold(name, "no-cache") || strings.EqualFold(name, "no-store") {
			return now
		}
		if len(parts) != 2 || !strings.EqualFold(name, "max-age") {
			continue
		}
		seconds, err := strconv.ParseInt(strings.Trim(strings.TrimSpace(parts[1]), `"`), 10, 64)
		maxCacheAgeSeconds := int64(^uint64(0)>>1) / int64(time.Second)
		if err == nil && seconds >= 0 && seconds <= maxCacheAgeSeconds {
			maxAgeSeconds = seconds
			hasMaxAge = true
		}
	}
	if hasMaxAge {
		return now.Add(time.Duration(maxAgeSeconds) * time.Second)
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
