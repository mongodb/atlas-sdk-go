package credentials

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20241023001/auth"
)

// MockOAuthRevokeEndpoint creates a mock OAuth revoke endpoint,
// that simulates token revocation responses.
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

// Test OAuthTokenSource_RevokeToken_Success tests successful token revocation.
func TestOAuthTokenSource_RevokeToken_Success(t *testing.T) {
	mockServer := MockOAuthRevokeEndpoint(http.StatusOK)
	defer mockServer.Close()

	config := *NewConfig("clientID", "clientSecret")
	config.RevokeURL = mockServer.URL
	expiry := time.Now().Add(1 * time.Hour)
	err := config.RevokeToken(context.Background(), &auth.Token{
		AccessToken: "test",
		Expiry:      expiry,
		ExpiresIn:   expiry.Unix(),
	})
	assert.NoError(t, err)
}

// TestOAuthTokenSource_RevokeToken_Failure tests token revocation failure due to unauthorized access.
func TestOAuthTokenSource_RevokeToken_Failure(t *testing.T) {
	mockServer := MockOAuthRevokeEndpoint(http.StatusUnauthorized)
	defer mockServer.Close()

	config := *NewConfig("clientID", "clientSecret")
	expiry := time.Now().Add(1 * time.Hour)
	err := config.RevokeToken(context.Background(), &auth.Token{
		AccessToken: "test",
		Expiry:      expiry,
		ExpiresIn:   expiry.Unix(),
	})
	config.RevokeURL = mockServer.URL
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to revoke Access Token when fetching new OAuth Token from remote server")
}
