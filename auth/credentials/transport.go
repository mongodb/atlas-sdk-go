package credentials

import (
	"fmt"
	"net/http"
)

// OAuthCustomHTTPTransport is an custom HTTP transport that injects the OAuth2 Token into requests.
// Transport can be extended by supplying UnderlyingTransport
// TokenSource can be created by credentials.NewTokenSource
type OAuthCustomHTTPTransport struct {
	// UnderlyingTransport is the base RoundTripper used to make HTTP requests.
	UnderlyingTransport http.RoundTripper
	// TokenSource used to obtain valid OAuth Token
	TokenSource TokenSource
}

// RoundTrip implements the RoundTripper interface.
func (t *OAuthCustomHTTPTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Get a valid Token (refreshing it if necessary)
	token, err := t.TokenSource.GetValidToken() // Get or refresh the Token
	if err != nil {
		return nil, fmt.Errorf("failed to inject access Token: %w", err)
	}

	// Inject the Token into the request
	t.setAuthHeader(req, token)

	// Proceed with the underlying transport
	return t.UnderlyingTransport.RoundTrip(req)
}

// SetAuthHeader sets the Authorization header with the access Token.
func (t *OAuthCustomHTTPTransport) setAuthHeader(r *http.Request, token *Token) {
	r.Header.Set("Authorization", "Bearer "+token.AccessToken)
}
