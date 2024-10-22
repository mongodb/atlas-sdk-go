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

	"github.com/stretchr/testify/assert"
)

// TokenResponse represents a mock OAuth token response.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// MockOAuthTokenEndpoint creates a mock OAuth token endpoint that returns a fake token.
func MockOAuthTokenEndpoint(token string) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request uses POST and contains proper grant_type
		if r.Method != http.MethodPost || r.FormValue("grant_type") != "client_credentials" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		// Create mock token response
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

// Mock token generation for testing
func generateMockJWT(expiration time.Time) string {
	header := `{"alg":"HS256","typ":"JWT"}`
	payload := `{"exp":` + json.Number(fmt.Sprintf("%d", expiration.Unix())) + `}`

	return base64.RawURLEncoding.EncodeToString([]byte(header)) + "." +
		base64.RawURLEncoding.EncodeToString([]byte(payload)) + ".signature"
}

// MockTokenSource implements the TokenSource interface for testing
type MockTokenSource struct {
	token string
}

// Retrieve returns the stored token
func (m *MockTokenSource) RetrieveToken() (*string, error) {
	return &m.token, nil
}

// Save saves the token
func (m *MockTokenSource) SaveToken(token string) error {
	m.token = token
	return nil
}

// Test for providing valid token and parsing the expiration date
func TestNewServiceAccountOAuthClientWithTokenSource_WithValidToken(t *testing.T) {
	// Generate a mock JWT token with future expiration
	expiration := time.Now().Add(1 * time.Hour)
	token := generateMockJWT(expiration)

	mockSource := &MockTokenSource{token: token}
	client := NewServiceAccountOAuthClientWithTokenSource(ServiceAccountOAuthClientWithTokenSource{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenSource:  mockSource,
	})

	_, err := client.getValidToken() // Get or refresh the token
	assert.Nil(t, err)

	assert.NotNil(t, client)
	assert.Equal(t, token, client.token.AccessToken)
	assert.True(t, client.token.Valid())
	assert.WithinDuration(t, expiration, client.token.Expiry, time.Second)
}

// Test for providing valid token and parsing the expiration date
func TestNewServiceAccountOAuthClientWithTokenSource_WithWithExpiredToken(t *testing.T) {
	// Generate a mock JWT token with future expiration
	expiration := time.Now().Add(-4 * time.Hour)
	token := generateMockJWT(expiration)

	remoteToken := "TestTokenRefreshedOnServer"
	// Start the mock server and return expired token
	mockServer := MockOAuthTokenEndpoint(remoteToken)
	defer mockServer.Close()

	mockSource := &MockTokenSource{token: token}
	client := NewServiceAccountOAuthClientWithTokenSource(ServiceAccountOAuthClientWithTokenSource{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenSource:  mockSource,
		BaseURL:      &mockServer.URL,
	})

	_, err := client.getValidToken() // Get or refresh the token
	assert.Nil(t, err)
	assert.Equal(t, remoteToken, client.token.AccessToken)
}

// Test error handling when fetching a token fails (e.g., invalid response)
func TestOAuthClient_FetchToken_Failure(t *testing.T) {
	// Mock OAuth server failure
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte(`{"error":"invalid_client"}`))
		assert.NoError(t, err)
	}))
	defer mockServer.Close()

	mockSource := &MockTokenSource{}
	client := NewServiceAccountOAuthClientWithTokenSource(ServiceAccountOAuthClientWithTokenSource{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenSource:  mockSource,
		BaseURL:      &mockServer.URL,
	})

	// Call GetAccessToken expecting an error
	_, err := client.getValidToken()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to obtain token")
}

// // Test CustomTransport to ensure token is injected into requests
func TestCustomTransport_RoundTrip(t *testing.T) {
	// Mock the OAuth client
	expiration := time.Now().Add(1 * time.Hour)
	token := generateMockJWT(expiration)
	mockSource := &MockTokenSource{token: token}

	ctx := context.Background()
	oauthClient := NewServiceAccountOAuthClientWithTokenSource(ServiceAccountOAuthClientWithTokenSource{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		TokenSource:  mockSource,
		Context:      &ctx,
	})

	httpClientWithTransport := NewHTTPClientWithServiceAccountAuth(oauthClient)
	assert.NotNil(t, httpClientWithTransport)
	// Try using transport directly without http client
	transport := &OAuthCustomHTTPTransport{
		underlyingTransport: http.DefaultTransport,
		client:              oauthClient,
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
