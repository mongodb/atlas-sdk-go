package clientcredentials

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/atlas-sdk/v20250312018/auth"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// mockOAuthRevokeEndpoint creates a mock OAuth revoke endpoint,
// that simulates token revocation responses.
func mockOAuthRevokeEndpoint(statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB limit
		if r.Method != http.MethodPost || r.FormValue("token") == "" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		w.WriteHeader(statusCode)
	})
	return httptest.NewServer(handler)
}

// Test OAuthTokenSource_RevokeToken_Success tests successful token revocation.
func TestOAuthTokenSource_RevokeToken_Success(t *testing.T) {
	mockServer := mockOAuthRevokeEndpoint(http.StatusOK)
	defer mockServer.Close()

	config := NewConfig("clientID", "clientSecret")
	config.RevokeURL = mockServer.URL
	expiry := time.Now().Add(1 * time.Hour)
	err := config.RevokeToken(context.Background(), &auth.Token{
		AccessToken: "test",
		Expiry:      expiry,
		ExpiresIn:   expiry.Unix(),
	})
	assert.NoError(t, err)
}

func TestRevokeToken_UsesContextHTTPClient(t *testing.T) {
	var receivedHeader atomic.Value
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeader.Store(r.Header.Get("X-Custom-Transport"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	customTransport := roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("X-Custom-Transport", "true")
		return http.DefaultTransport.RoundTrip(req)
	})
	customClient := &http.Client{Transport: customTransport}

	config := NewConfig("clientID", "clientSecret")
	config.RevokeURL = server.URL
	expiry := time.Now().Add(1 * time.Hour)

	ctx := context.WithValue(context.Background(), auth.HTTPClient, customClient)
	err := config.RevokeToken(ctx, &auth.Token{
		AccessToken: "test-token",
		Expiry:      expiry,
		ExpiresIn:   expiry.Unix(),
	})
	assert.NoError(t, err)
	assert.Equal(t, "true", receivedHeader.Load(), "RevokeToken should use the HTTP client from the context")
}

func TestRevokeToken_FallsBackToDefaultClient(t *testing.T) {
	var called atomic.Bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		called.Store(true)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewConfig("clientID", "clientSecret")
	config.RevokeURL = server.URL
	expiry := time.Now().Add(1 * time.Hour)

	err := config.RevokeToken(context.Background(), &auth.Token{
		AccessToken: "test-token",
		Expiry:      expiry,
		ExpiresIn:   expiry.Unix(),
	})
	assert.NoError(t, err)
	assert.True(t, called.Load(), "RevokeToken should fall back to default client and still reach the server")
}

// TestOAuthTokenSource_RevokeToken_Failure tests token revocation failure due to unauthorized access.
func TestOAuthTokenSource_RevokeToken_Failure(t *testing.T) {
	mockServer := mockOAuthRevokeEndpoint(http.StatusUnauthorized)
	defer mockServer.Close()

	config := NewConfig("clientID", "clientSecret")
	expiry := time.Now().Add(1 * time.Hour)
	err := config.RevokeToken(context.Background(), &auth.Token{
		AccessToken: "test",
		Expiry:      expiry,
		ExpiresIn:   expiry.Unix(),
	})
	assert.Error(t, err)
	assert.ErrorContains(t, err, "Failed to revoke Access Token when fetching new OAuth Token from remote server")
}
