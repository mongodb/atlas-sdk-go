package credentials

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Mock OAuth2 server simulating token issuance
func MockOAuthServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token":"mock_access_token","token_type":"Bearer","expires_in":3600}`))
	}))
}

// Test the token fetching logic
func TestOAuthClient_FetchToken(t *testing.T) {
	// Mock OAuth2 server
	server := MockOAuthServer()
	defer server.Close()

	// Initialize OAuthClient with the mocked server
	client := NewServiceAccountOAuthClient("test-client-id", "test-client-secret", &server.URL)

	token, err := client.GetAccessToken()
	if err != nil {
		t.Fatalf("Failed to fetch token: %v", err)
	}

	// Check that the token is properly populated
	if token.AccessToken != "mock_access_token" {
		t.Errorf("Unexpected access token: got %v, want %v", token.AccessToken, "mock_access_token")
	}
	if token.TokenType != "Bearer" {
		t.Errorf("Unexpected token type: got %v, want %v", token.TokenType, "Bearer")
	}
}

// Test token expiry logic
func TestToken_Expired(t *testing.T) {
	// Test token that expires in 2 seconds
	token := &Token{
		AccessToken: "mock_access_token",
		Expiry:      time.Now().Add(12 * time.Second),
	}

	if token.expired() {
		t.Errorf("Token shouldn't be expired yet")
	}

	// Wait for token to expire
	time.Sleep(12 * time.Second)

	if !token.expired() {
		t.Errorf("Token should be expired now")
	}
}

// Test that the custom transport injects the Authorization header
func TestCustomTransport_RoundTrip(t *testing.T) {
	// Mock OAuth2 server for token issuance
	server := MockOAuthServer()
	defer server.Close()

	// Mock API server to check the Authorization header
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer mock_access_token" {
			t.Errorf("Unexpected Authorization header: got %v, want %v", authHeader, "Bearer mock_access_token")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	// Initialize OAuthClient with the mocked OAuth server
	client := NewServiceAccountOAuthClient("test-client-id", "test-client-secret", &server.URL)

	// Create an HTTP client with the custom transport
	httpClient := NewHTTPClientWithServiceAccountAuth(client)

	// Make a request to the mock API server
	resp, err := httpClient.Get(apiServer.URL)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check that the request succeeded
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

// Test OAuthClient with an invalid token response
func TestOAuthClient_InvalidTokenResponse(t *testing.T) {
	// Mock OAuth2 server returning an invalid response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token": "", "expires_in": 3600}`)) // Invalid token
	}))
	defer server.Close()

	// Initialize OAuthClient with the mocked server
	client := NewServiceAccountOAuthClient("test-client-id", "test-client-secret", &server.URL)

	_, err := client.GetAccessToken()
	if err == nil {
		t.Fatalf("Expected an error due to invalid token, but got none")
	}
}

// Test the fetchToken error path (e.g., non-200 status code from token server)
func TestOAuthClient_FetchTokenError(t *testing.T) {
	// Mock OAuth2 server that returns an error status
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	// Initialize OAuthClient with the mocked server
	client := NewServiceAccountOAuthClient("test-client-id", "test-client-secret", &server.URL)

	_, err := client.GetAccessToken()
	if err == nil {
		t.Fatalf("Expected an error due to invalid response, but got none")
	}
}
