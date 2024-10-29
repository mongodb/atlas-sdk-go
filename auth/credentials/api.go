package credentials

import (
	"context"
	"net/http"
	"strings"

	"go.mongodb.org/atlas-sdk/v20240805005/internal/core"
)

// tokenAPIPath for obtaining OAuth Access Token from server
//
//nolint:gosec //url only
const tokenAPIPath = "/api/oauth/token"

// revokeAPIPath for revoking OAuth Access Token from server
const revokeAPIPath = "/api/oauth/revoke"

// serverURL for Token Atlas API
const serverTokenURL = core.DefaultCloudURL + tokenAPIPath

// serverURL for Revoke Atlas API
const serverRevokeURL = core.DefaultCloudURL + revokeAPIPath

// AtlasTokenSourceOptions provides set of input arguments
// for creation of credentials.TokenSource interface
type AtlasTokenSourceOptions struct {
	ClientID     string
	ClientSecret string
	// Custom user agent. core.DefaultUserAgent being default
	UserAgent string
	// Custom Token source. InMemoryTokenCache being default
	TokenCache LocalTokenCache

	// Custom context. context.Background() will be used by default
	Context *context.Context
	// Custom base url for fetching Token using TokenSource. Reserved for internal use.
	BaseURL *string
}

// NewTokenSourceWithOptions initializes an OAuthTokenSource with advanced credentials.AtlasTokenSourceOptions
func NewTokenSourceWithOptions(opts AtlasTokenSourceOptions) TokenSource {
	var tokenURL string
	var revokeUrl string
	if opts.BaseURL != nil {
		baseUrlNoSuffix := strings.TrimSuffix(*opts.BaseURL, "/")
		tokenURL = baseUrlNoSuffix + tokenAPIPath
		revokeUrl = baseUrlNoSuffix + revokeAPIPath
	} else {
		tokenURL = serverTokenURL
		revokeUrl = serverRevokeURL
	}
	var userAgent string
	if opts.UserAgent != "" {
		userAgent = opts.UserAgent
	} else {
		userAgent = core.DefaultUserAgent
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
		userAgent:    userAgent,
		tokenURL:     tokenURL,
		revokeURL:    revokeUrl,
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
		UserAgent:    core.DefaultUserAgent,
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
