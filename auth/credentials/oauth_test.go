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

	"go.mongodb.org/atlas-sdk/v20241023001/internal/core"

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
func (m *MockTokenCache) RetrieveToken(_ context.Context) (string, error) {
	return m.token, nil
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

	validToken, err := tokenSource.Token() // Get or refresh the Token
	assert.Nil(t, err)
	assert.NotNil(t, validToken)
	assert.Equal(t, token, validToken.AccessToken)
	assert.True(t, validToken.Valid())
	assert.WithinDuration(t, expiration, validToken.Expiry, time.Second)
}

// Test for providing valid Token and parsing the expiration date
func TestNewServiceAccountOAuthClientWithTokenSource_WithWithExpiredToken(t *testing.T) {
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

	validToken, err := tokenSource.Token() // Get or refresh the Token
	assert.Nil(t, err)
	assert.Equal(t, remoteToken, validToken.AccessToken)
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
	validToken, err := tokenSource.Token() // Get or refresh the Token
	assert.Nil(t, err)
	assert.Equal(t, remoteToken, validToken.AccessToken)
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

	// Call GetValidToken expecting an error
	_, err := tokenSource.Token()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "oauth2: cannot fetch token: 401 Unauthorized")
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

	// Call GetValidToken expecting an error
	_, err := client.Token()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Rate Limit reached")
	assert.Contains(t, err.Error(), "auth2: cannot fetch token: 429 Too Many Requests")
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

	// Call GetValidToken expecting an error
	_, err := client.Token()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "oauth2: cannot fetch token: 403 Forbidden")
}

// Test OAuthTokenSource_RevokeToken_Success tests successful token revocation.
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
