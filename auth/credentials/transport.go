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
    // Get a valid token (refreshing it if necessary)
    token, err := t.client.getValidToken() // Get or refresh the token
	if err != nil {
		return nil, fmt.Errorf("failed to inject access token: %w", err)
	}

	// Inject the token into the request
	req.Header.Set("Authorization", "Bearer " + token.AccessToken)

	// Proceed with the underlying transport
	return t.underlyingTransport.RoundTrip(req)
}
