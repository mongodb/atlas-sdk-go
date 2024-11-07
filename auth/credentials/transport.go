package credentials

import (
	"go.mongodb.org/atlas-sdk/v20241023001/auth"
	"net/http"
)

// OAuthTransport provides a way to cache OAuth Token
type OAuthTransport struct {
	tokenSource         auth.TokenSource
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
	req.Header.Set("User-Agent", t.userAgent)
	return t.underlyingTransport.RoundTrip(req)
}

// NewOAuthTransport http client that adds the User-Agent header
func NewOAuthTransport(transport http.RoundTripper, userAgent string, tokenCache *func(token string) error) http.RoundTripper {
	return &OAuthTransport{
		underlyingTransport: transport,
		userAgent:           userAgent,
		tokenCache:          tokenCache,
	}
}
