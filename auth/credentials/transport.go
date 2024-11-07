package credentials

import (
	"go.mongodb.org/atlas-sdk/v20241023002/internal/core"
	"net/http"
)

// OAuthTransport provides a way to cache OAuth Token
type OAuthTransport struct {
	underlyingTransport http.RoundTripper
	userAgent           string
	tokenCache          *func(token string) error
}

func (t *OAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token := req.Header.Get("Authorization")
	if t.tokenCache != nil && token != "" {
		err := (*t.tokenCache)(token)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Set(userAgent, t.userAgent)
	return t.underlyingTransport.RoundTrip(req)
}

// NewOAuthCacheTransport http client that adds the User-Agent header and TokenCache
func NewOAuthCacheTransport(transport http.RoundTripper, userAgent *string, tokenCache *func(token string) error) http.RoundTripper {
	if userAgent != nil {
		defaultUserAgent := core.DefaultUserAgent
		userAgent = &defaultUserAgent
	}
	return &OAuthTransport{
		underlyingTransport: transport,
		userAgent:           *userAgent,
		tokenCache:          tokenCache,
	}
}
