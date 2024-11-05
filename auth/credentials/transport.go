package credentials

import (
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

// OAuthCustomHTTPTransport is custom HTTP transport that injects the OAuth2 Token into requests.
// Transport can be extended by supplying UnderlyingTransport
// TokenSource can be created by credentials.NewTokenSource
type OAuthCustomHTTPTransport struct {
	// UnderlyingTransport is the base RoundTripper used to make HTTP requests.
	UnderlyingTransport http.RoundTripper
	// TokenSource used to get valid OAuth Token
	TokenSource oauth2.TokenSource
}

// RoundTrip implements the RoundTripper interface.
func (t *OAuthCustomHTTPTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Get a valid Token (refreshing it if necessary)
	token, err := t.TokenSource.Token() // Get or refresh the Token
	if err != nil {
		return nil, fmt.Errorf("failed to inject access Token: %w", err)
	}

	// Inject the Token into the request
	t.setAuthHeader(req, token)

	// Proceed with the underlying transport
	return t.UnderlyingTransport.RoundTrip(req)
}

// SetAuthHeader sets the Authorization header with the access Token.
func (t *OAuthCustomHTTPTransport) setAuthHeader(r *http.Request, token *oauth2.Token) {
	r.Header.Set("Authorization", "Bearer "+token.AccessToken)
}

// userAgentTransport that adds the User-Agent header
type userAgentTransport struct {
	underlyingTransport http.RoundTripper
	userAgent           string
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent) // Set the custom User-Agent
	return t.underlyingTransport.RoundTrip(req)
}

// newUserAgentHTTPClient http client that adds the User-Agent header
func newUserAgentHTTPClient(userAgent string) *http.Client {
	return &http.Client{
		Transport: &userAgentTransport{
			underlyingTransport: http.DefaultTransport,
			userAgent:           userAgent,
		},
	}
}
