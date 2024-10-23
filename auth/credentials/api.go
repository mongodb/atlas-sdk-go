package credentials

import (
	"context"
	"net/http"
)

// tokenAPIPath for obtaining OAuth Access Token from server
//
//nolint:gosec //url only
const tokenAPIPath = "/api/oauth/token"

// serverURL for atlas API
const serverURL = "https://cloud.mongodb.com" + tokenAPIPath

// AtlasTokenSourceOptions provides set of input arguments
// for creation of credentials.TokenSource interface
type AtlasTokenSourceOptions struct {
	ClientID     string
	ClientSecret string
	// Custom Token source. InMemoryTokenCache being default
	TokenCache LocalTokenCache

	// Custom context. context.Background() will be used by default
	Context *context.Context
	// Custom base url for fetching Token using TokenSource. Reserved for internal use.
	BaseURL *string
}

// NewTokenSourceWithOptions initializes an OAuthTokenSource with advanced credentials.AtlasTokenSourceOptions
// Use this method to initialize custom OAuth Token Cache (filesystem).
func NewTokenSourceWithOptions(opts AtlasTokenSourceOptions) TokenSource {
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

	return &OAuthTokenSource{
		clientID:     opts.ClientID,
		clientSecret: opts.ClientSecret,
		tokenURL:     tokenURL,
		tokenCache:   opts.TokenCache,
		ctx:          ctx,
	}
}

// NewTokenSource initializes OAuth Token Source that provides a way to obtain valid OAuth Tokens.
// See credentials.NewTokenSourceWithOptions for advanced use cases.
func NewTokenSource(clientID, clientSecret string) TokenSource {
	return NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenCache:   &InMemoryTokenCache{},
	})
}

// NewHTTPClientWithOAuthToken helper method for creating HTTP client with OAuth authentication support.
// Use this method for performing requests using http.DefaultTransport.
// For more advanced use cases please create your own credentials.OAuthCustomHTTPTransport.
func NewHTTPClientWithOAuthToken(client TokenSource) *http.Client {
	return &http.Client{
		Transport: &OAuthCustomHTTPTransport{
			UnderlyingTransport: http.DefaultTransport,
			TokenSource:         client,
		},
	}
}
