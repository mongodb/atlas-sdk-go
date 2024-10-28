package credentials

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.mongodb.org/atlas-sdk/v20240805005/internal/core"

	"github.com/stretchr/testify/assert"
)

// TokenResponse represents a mock OAuth Token response.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// MockOAuthTokenEndpoint creates a mock OAuth Token endpoint that returns a fake Token.
func MockOAuthTokenEndpoint(token string) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request uses POST and contains proper grant_type
		if r.Method != http.MethodPost || r.FormValue("grant_type") != "client_credentials" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		// Create mock Token response
		tokenResponse := TokenResponse{
			AccessToken: token,
			TokenType:   "Bearer",
			ExpiresIn:   3600,
		}

		// Encode and return the JSON response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(tokenResponse)
		if err != nil {
			log.Fatal("Failed to encode mocked response")
		}
	})

	// Create a new mock server
	return httptest.NewServer(handler)
}

// Mock Token generation for testing
func generateMockJWT(expiration time.Time) string {
	header := `{"alg":"HS256","typ":"JWT"}`
	payload := `{"exp":` + json.Number(fmt.Sprintf("%d", expiration.Unix())) + `}`

	return base64.RawURLEncoding.EncodeToString([]byte(header)) + "." +
		base64.RawURLEncoding.EncodeToString([]byte(payload)) + ".signature"
}

// MockTokenCache implements the LocalTokenCache interface for testing
type MockTokenCache struct {
	token string
}

// Retrieve returns the stored Token
func (m *MockTokenCache) RetrieveToken(_ context.Context) (*string, error) {
	return &m.token, nil
}

// Save saves the Token
func (m *MockTokenCache) SaveToken(_ context.Context, token string) error {
	m.token = token
	return nil
}

// Test for providing valid Token and parsing the expiration date
func TestNewServiceAccountOAuthClientWithTokenSource_WithValidToken(t *testing.T) {
	// Generate a mock JWT Token with future expiration
	expiration := time.Now().Add(1 * time.Hour)
	token := generateMockJWT(expiration)

	mockCache := &MockTokenCache{token: token}
	tokenSource := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   mockCache,
	})

	_, err := tokenSource.GetValidToken() // Get or refresh the Token
	assert.Nil(t, err)
	oAuthTokenSource := tokenSource.(*OAuthTokenSource)

	assert.NotNil(t, oAuthTokenSource)
	assert.Equal(t, token, oAuthTokenSource.token.AccessToken)
	assert.True(t, oAuthTokenSource.token.Valid())
	assert.WithinDuration(t, expiration, oAuthTokenSource.token.Expiry, time.Second)
}

// Test for providing valid Token and parsing the expiration date
func TestNewServiceAccountOAuthClientWithTokenSource_WithWithExpiredToken(t *testing.T) {
	// Generate a mock JWT Token with past expiration
	expiration := time.Now().Add(-4 * time.Hour)
	token := generateMockJWT(expiration)

	remoteToken := "TestTokenRefreshedOnServer"
	// Start the mock server and return expired Token
	mockServer := MockOAuthTokenEndpoint(remoteToken)
	defer mockServer.Close()

	mockCache := &MockTokenCache{token: token}
	tokenSource := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   mockCache,
		BaseURL:      &mockServer.URL,
	})

	oAuthTokenSource := tokenSource.(*OAuthTokenSource)
	_, err := tokenSource.GetValidToken() // Get or refresh the Token
	assert.Nil(t, err)
	assert.Equal(t, remoteToken, oAuthTokenSource.token.AccessToken)
}

// Test for providing valid Token and parsing the expiration date
func TestNewServiceAccountOAuthClientWithTokenSource_WithWithAlmostExpiredToken(t *testing.T) {
	// Generate a mock JWT Token that is going to expire in 5 seconds
	expiration := time.Now().Add(-5 * time.Second)
	token := generateMockJWT(expiration)

	remoteToken := "TestTokenRefreshedOnServer"
	// Start the mock server and return expired Token
	mockServer := MockOAuthTokenEndpoint(remoteToken)
	defer mockServer.Close()

	mockCache := &MockTokenCache{token: token}
	tokenSource := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   mockCache,
		BaseURL:      &mockServer.URL,
	})

	oAuthTokenSource := tokenSource.(*OAuthTokenSource)
	_, err := tokenSource.GetValidToken() // Get or refresh the Token
	assert.Nil(t, err)
	assert.Equal(t, remoteToken, oAuthTokenSource.token.AccessToken)
}

// Test error handling when fetching a Token fails (e.g., invalid response)
func TestOAuthClient_FetchToken_Failure(t *testing.T) {
	// Mock OAuth server failure
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte(`{"error":"invalid_client"}`))
		assert.NoError(t, err)
	}))
	defer mockServer.Close()

	mockCache := &MockTokenCache{}
	finalUrl := mockServer.URL + "/"
	tokenSource := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   mockCache,
		BaseURL:      &finalUrl,
	})

	// Call GetAccessToken expecting an error
	_, err := tokenSource.GetValidToken()
	assert.Error(t, err)
	oAuthTokenSource := tokenSource.(*OAuthTokenSource)
	assert.Equal(t, mockServer.URL+tokenAPIPath, oAuthTokenSource.tokenURL)
	assert.Contains(t, err.Error(), "Failed to obtain Access Token when fetching new OAuth Token from remote server")
}

// Test error handling when fetching a Token fails (e.g., invalid response)
func TestOAuthClient_FetchToken_FailureRateLimit(t *testing.T) {
	// Mock OAuth server failure
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, err := w.Write([]byte(`Rate Limit reached`))
		assert.NoError(t, err)
	}))
	defer mockServer.Close()

	mockCache := &MockTokenCache{}
	client := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   mockCache,
		BaseURL:      &mockServer.URL,
	})

	// Call GetAccessToken expecting an error
	_, err := client.GetValidToken()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Rate Limit reached")
	assert.Contains(t, err.Error(), "Token request was rate limited")
}

// Test error handling when fetching a Token fails (e.g., invalid response)
func TestOAuthClient_FetchToken_FailureHtmlContentType(t *testing.T) {
	// Mock OAuth server failure
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte(`<html>Wrong</html>`))
		assert.NoError(t, err)
	}))
	defer mockServer.Close()

	mockCache := &MockTokenCache{}
	client := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   mockCache,
		BaseURL:      &mockServer.URL,
	})

	// Call GetAccessToken expecting an error
	_, err := client.GetValidToken()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to obtain Access Token when fetching new OAuth Token from remote server")
}

// Test CustomTransport to ensure Token is injected into requests
func TestCustomTransport_RoundTrip(t *testing.T) {
	// Mock the OAuth TokenCache
	expiration := time.Now().Add(1 * time.Hour)
	token := generateMockJWT(expiration)
	mockCache := &MockTokenCache{token: token}

	ctx := context.Background()
	oauthClient := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   mockCache,
		Context:      &ctx,
	})

	httpClientWithTransport := NewHTTPClientWithOAuthToken(oauthClient)
	assert.NotNil(t, httpClientWithTransport)
	// Try using transport directly without http TokenCache
	transport := &OAuthCustomHTTPTransport{
		UnderlyingTransport: http.DefaultTransport,
		TokenSource:         oauthClient,
	}

	// Mock the HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the Authorization header is set correctly
		assert.Contains(t, r.Header.Get("Authorization"), "Bearer "+token)
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	// Make a request using the custom transport
	req, _ := http.NewRequest("GET", mockServer.URL, http.NoBody)
	client := &http.Client{Transport: transport}
	//nolint:all
	resp, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Test default User Agent
func TestOAuthClient_DefaultUserAgent(t *testing.T) {
	// Assert default User Agent
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotNil(t, r.Header.Get("User-Agent"))
		assert.Equal(t, core.DefaultUserAgent, r.Header.Get("User-Agent"))
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	mockCache := &MockTokenCache{}
	tokenSource := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   mockCache,
		BaseURL:      &mockServer.URL,
	})

	_, _ = tokenSource.GetValidToken()
}

// Test custom User Agent
func TestOAuthClient_CustomUserAgent(t *testing.T) {
	const customUserAgent = "/testing"

	// Assert custom User Agent
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotNil(t, r.Header.Get("User-Agent"))
		assert.Equal(t, customUserAgent, r.Header.Get("User-Agent"))
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	mockCache := &MockTokenCache{}
	tokenSource := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		UserAgent:    customUserAgent,
		TokenCache:   mockCache,
		BaseURL:      &mockServer.URL,
	})

	_, _ = tokenSource.GetValidToken()
}

// MockOAuthRevokeEndpoint creates a mock OAuth revoke endpoint that simulates token revocation responses.
func MockOAuthRevokeEndpoint(statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.FormValue("token") == "" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		w.WriteHeader(statusCode)
	})
	return httptest.NewServer(handler)
}

// TestOAuthTokenSource_RevokeToken_Success tests successful token revocation.
func TestOAuthTokenSource_RevokeToken_Success(t *testing.T) {
	mockServer := MockOAuthRevokeEndpoint(http.StatusOK)
	defer mockServer.Close()

	tokenSource := &OAuthTokenSource{
		clientID:     "clientID",
		clientSecret: "clientSecret",
		userAgent:    core.DefaultUserAgent,
		tokenCache:   &MockTokenCache{token: "test"},
		revokeURL:    mockServer.URL,
	}

	err := tokenSource.RevokeToken()
	assert.NoError(t, err)
}

func TestOAuthTokenSource_TokenAndRevokeUrls_Default(t *testing.T) {
	tokenSource := NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenCache:   &MockTokenCache{token: "test"},
	})

	oAuthTokenSource := tokenSource.(*OAuthTokenSource)
	assert.NotNil(t, oAuthTokenSource.tokenURL)
	assert.NotNil(t, oAuthTokenSource.revokeURL)
	assert.Equal(t, oAuthTokenSource.tokenURL, serverTokenURL)
	assert.Equal(t, oAuthTokenSource.revokeURL, serverRevokeURL)
}

// TestOAuthTokenSource_RevokeToken_Failure tests token revocation failure due to unauthorized access.
func TestOAuthTokenSource_RevokeToken_Failure(t *testing.T) {
	mockServer := MockOAuthRevokeEndpoint(http.StatusUnauthorized)
	defer mockServer.Close()

	tokenSource := &OAuthTokenSource{
		clientID:     "clientID",
		clientSecret: "clientSecret",
		userAgent:    core.DefaultUserAgent,
		tokenCache:   &MockTokenCache{token: "test"},
		revokeURL:    mockServer.URL,
	}

	err := tokenSource.RevokeToken()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to revoke Access Token when fetching new OAuth Token from remote server")
}

// TestOAuthTokenSource_RevokeToken_RateLimit tests handling of rate-limited responses during token revocation.
func TestOAuthTokenSource_RevokeToken_RateLimit(t *testing.T) {
	mockServer := MockOAuthRevokeEndpoint(http.StatusTooManyRequests)
	defer mockServer.Close()

	tokenSource := &OAuthTokenSource{
		clientID:     "clientID",
		clientSecret: "clientSecret",
		userAgent:    core.DefaultUserAgent,
		tokenCache:   &MockTokenCache{token: "test"},
		revokeURL:    mockServer.URL,
	}

	err := tokenSource.RevokeToken()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rate limited")
}

// TestOAuthTokenSource_RevokeToken_InvalidRequest tests handling of invalid request scenarios.
func TestOAuthTokenSource_RevokeToken_InvalidRequest(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}))
	defer mockServer.Close()

	tokenSource := &OAuthTokenSource{
		clientID:     "clientID",
		clientSecret: "clientSecret",
		userAgent:    core.DefaultUserAgent,
		tokenCache:   &MockTokenCache{token: "test"},
		revokeURL:    mockServer.URL,
	}

	err := tokenSource.RevokeToken()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to revoke Access Token when fetching new OAuth Token from remote server")
}
