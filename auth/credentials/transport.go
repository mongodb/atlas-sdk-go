package credentials

import (
	"go.mongodb.org/atlas-sdk/v20241023002/internal/core"
	"net/http"
)

// oAuthTransport supplies custom user agent to token requests
type oAuthTransport struct {
	underlyingTransport http.RoundTripper
	userAgent           string
}

func (t *oAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(userAgent, t.userAgent)
	return t.underlyingTransport.RoundTrip(req)
}

// NewOAuthCacheTransport http client that adds the User-Agent header
func NewOAuthCacheTransport(transport http.RoundTripper, userAgent *string) http.RoundTripper {
	if userAgent != nil {
		defaultUserAgent := core.DefaultUserAgent
		userAgent = &defaultUserAgent
	}
	return &oAuthTransport{
		underlyingTransport: transport,
		userAgent:           *userAgent,
	}
}
