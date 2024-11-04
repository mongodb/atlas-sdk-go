package credentials

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"go.mongodb.org/atlas-sdk/v20241023001/internal/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// tokenAPIPath for obtaining OAuth Access Token from server
// nolint:gosec //url only
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
func NewTokenSourceWithOptions(opts AtlasTokenSourceOptions) oauth2.TokenSource {
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

	config := &clientcredentials.Config{
		ClientID:     opts.ClientID,
		ClientSecret: opts.ClientSecret,
		TokenURL:     tokenURL,
		AuthStyle:    oauth2.AuthStyleInHeader,
		Scopes:       []string{}, // No support for scopes.
	}

	// oauth2 token source for creating tokens
	tokenSource := config.TokenSource(ctx)
	if opts.TokenCache != nil {
		tokenSource = ReuseWrapperCacheTokenSource(tokenSource, opts.TokenCache)
	}

	return &OAuthTokenSource{
		tokenSource:  tokenSource,
		tokenCache:   opts.TokenCache,
		revokeURL:    revokeUrl,
		userAgent:    userAgent,
		ctx:          ctx,
		clientID:     opts.ClientID,
		clientSecret: opts.ClientSecret,
	}
}

// NewTokenSource initializes OAuth Token Source that provides a way to obtain valid OAuth Tokens.
// See credentials.NewTokenSourceWithOptions for advanced use cases.
func NewTokenSource(clientID, clientSecret string) oauth2.TokenSource {
	return NewTokenSourceWithOptions(AtlasTokenSourceOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		UserAgent:    core.DefaultUserAgent,
	})
}

// NewHTTPClientWithOAuthToken helper method for creating HTTP client with OAuth authentication support.
// Use this method for performing requests using http.DefaultTransport.
// For more advanced use cases, please create your own credentials.OAuthCustomHTTPTransport.
func NewHTTPClientWithOAuthToken(tokenSource oauth2.TokenSource) *http.Client {
	return &http.Client{
		Transport: &OAuthCustomHTTPTransport{
			UnderlyingTransport: http.DefaultTransport,
			TokenSource:         tokenSource,
		},
	}
}

// RevokeToken checks if TokenSource is a RevocableTokenSource and revokes Token.
func RevokeToken(tokenSource oauth2.TokenSource) error {
	var revocableTokenSource, ok = tokenSource.(RevocableTokenSource)
	if !ok {
		return errors.New("cannot revoke OAuth Token. tokenSource is not of type RevocableTokenSource")
	}
	return revocableTokenSource.RevokeToken()
}
