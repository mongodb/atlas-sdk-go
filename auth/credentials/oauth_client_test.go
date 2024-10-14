package credentials

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Mock token generation for testing
func generateMockJWT(expiration time.Time) string {
	header := `{"alg":"HS256","typ":"JWT"}`
	payload := `{"exp":` + json.Number(fmt.Sprintf("%d", expiration.Unix())) + `}`

	return base64.RawURLEncoding.EncodeToString([]byte(header)) + "." +
		base64.RawURLEncoding.EncodeToString([]byte(payload)) + ".signature"
}

// Test for providing valid token and parsing the expiration date
func TestNewServiceAccountOAuthClient_WithValidToken(t *testing.T) {
	// Generate a mock JWT token with future expiration
	expiration := time.Now().Add(1 * time.Hour)
	token := generateMockJWT(expiration)

	client, err := NewServiceAccountOAuthClient("clientID", "clientSecret", token, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, token, client.token.AccessToken)
	assert.True(t, client.token.Valid())
	assert.WithinDuration(t, expiration, client.token.Expiry, time.Second)
}

// Test for providing an expired token and ensuring it's invalid
func TestNewServiceAccountOAuthClient_WithExpiredToken(t *testing.T) {
	// Generate a mock JWT token with past expiration
	expiration := time.Now().Add(-1 * time.Hour)
	token := generateMockJWT(expiration)

	client, err := NewServiceAccountOAuthClient("clientID", "clientSecret", token, nil)
	assert.Error(t, err)
	assert.NotNil(t, client)
}

// Test token fetching when no valid token is provided (simulate server response)
func TestOAuthClient_FetchNewToken(t *testing.T) {
	// Mock the OAuth server response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Contains(t, r.Header.Get("Authorization"), "Basic")

		resp := `{"access_token":"new_access_token","token_type":"bearer","expires_in":3600}`
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer mockServer.Close()

	client, err := NewServiceAccountOAuthClient("clientID", "clientSecret", "", &mockServer.URL)
	assert.NoError(t, err)

	// Call GetAccessToken (which should fetch a new token)
	token, err := client.GetAccessToken()
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, "new_access_token", token.AccessToken)
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

	client, err := NewServiceAccountOAuthClient("clientID", "clientSecret", "", &mockServer.URL)
	assert.NoError(t, err)

	// Call GetAccessToken expecting an error
	_, err = client.GetAccessToken()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to obtain token")
}

// Test CustomTransport to ensure token is injected into requests
func TestCustomTransport_RoundTrip(t *testing.T) {
	// Mock the OAuth client
	expiration := time.Now().Add(1 * time.Hour)
	token := generateMockJWT(expiration)
	oauthClient, err := NewServiceAccountOAuthClient("clientID", "clientSecret", token, nil)
	assert.NoError(t, err)

	transport := &CustomTransport{
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
	/* trunk-ignore(golangci-lint/bodyclose) */
	resp, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Test token expiration and automatic refresh
func TestOAuthClient_AutoRefreshExpiredToken(t *testing.T) {
	// Mock the OAuth server response for refreshing token
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `{"access_token":"refreshed_token","token_type":"bearer","expires_in":3600}`
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer mockServer.Close()

	// Initialize client with expired token
	expiration := time.Now().Add(-1 * time.Hour)
	expiredToken := generateMockJWT(expiration)

	client, err := NewServiceAccountOAuthClient("clientID", "clientSecret", expiredToken, &mockServer.URL)
	// Token expired but we can still continue using client
	assert.Error(t, err)

	// Fetch new token after expiration
	token, err := client.GetAccessToken()
	assert.NoError(t, err)
	assert.Equal(t, "refreshed_token", token.AccessToken)
	assert.True(t, token.Valid())
}
