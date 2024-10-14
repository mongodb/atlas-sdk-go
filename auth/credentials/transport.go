package credentials

import (
	"fmt"
	"net/http"
)

// CustomTransport is a custom HTTP transport that injects the OAuth2 token into requests.
type CustomTransport struct {
	underlyingTransport http.RoundTripper
	client              *OAuthClient
}

// RoundTrip implements the RoundTripper interface.
func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Ensure the token is valid
	token, err := t.client.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	// Inject the token into the request
	token.SetAuthHeader(req)

	// Proceed with the underlying transport
	return t.underlyingTransport.RoundTrip(req)
}
