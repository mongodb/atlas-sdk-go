package credentials

import (
	"context"
	"net/http"
)

// tokenAPIPath for obtaining OAuth Access token from server
//
//nolint:gosec //url only
const tokenAPIPath = "/api/oauth/token"

// serverURL for atlas API
const serverURL = "https://cloud.mongodb.com" + tokenAPIPath

type ServiceAccountOAuthClientWithTokenSource struct {
	ClientID     string
	ClientSecret string
	// Custom token source. InMemoryTokenSource being default
	TokenSource TokenSource

	// Custom context
	Context *context.Context
	// Custom base url for change
	BaseURL *string
}

// NewServiceAccountOAuthClientWithTokenSource initializes an OAuthClient with client credentials and custom TokenSource
// Use this method to initialze custom token storage (filesystem)
func NewServiceAccountOAuthClientWithTokenSource(opts ServiceAccountOAuthClientWithTokenSource) *OAuthClient {
	var tokenURL string
	if opts.BaseURL != nil {
		tokenURL = *opts.BaseURL + tokenAPIPath
	} else {
		tokenURL = serverURL
	}
	var ctx context.Context
	if opts.Context == nil {
		ctx = context.Background()
	} else {
		ctx = *opts.Context
	}

	return &OAuthClient{
		clientID:     opts.ClientID,
		clientSecret: opts.ClientSecret,
		tokenURL:     tokenURL,
		tokenSource:  opts.TokenSource,
		ctx:          ctx,
	}
}

// NewHTTPClientWithServiceAccountAuth creates an HTTP client with the custom transport.
func NewHTTPClientWithServiceAccountAuth(client *OAuthClient) *http.Client {
	return &http.Client{
		Transport: &OAuthCustomHTTPTransport{
			underlyingTransport: http.DefaultTransport,
			client:              client,
		},
	}
}

// NewServiceAccountOAuthClient initializes an OAuthClient with client credentials with default InMemoryTokenSource.
func NewServiceAccountOAuthClient(clientID, clientSecret string) *OAuthClient {
	return NewServiceAccountOAuthClientWithTokenSource(ServiceAccountOAuthClientWithTokenSource{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenSource:  &InMemoryTokenSource{},
	})
}
