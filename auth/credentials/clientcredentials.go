package credentials

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"go.mongodb.org/atlas-sdk/v20241023001/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// Config describes a 2-legged OAuth2 flow, with both the
// client application information and the server's endpoint URLs.
// NOTE: clientcredentials.Config values are used only internally
// and should not be overridden by clients
type Config struct {
	clientcredentials.Config
	RevokeURL string
	userAgent string
}

// tokenSource manages the OAuth Token fetching and refreshing using a LocalTokenCache.
type tokenSource struct {
	ctx  context.Context
	conf *Config
}

// RevokeToken revokes the Access Token while also removing it from the Access Token Cache.
// When the Access Token is expired or missing, revoke will return without any action.
func (c *Config) RevokeToken(ctx context.Context, t *auth.Token) error {
	if c.RevokeURL == "" {
		return errors.New("endpoint missing RevokeURL")
	}
	if !t.Valid() {
		return nil // nothing to do
	}
	v := url.Values{
		"token":           {t.AccessToken},
		"token_type_hint": {"access_token"},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.RevokeURL, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(url.QueryEscape(c.ClientID), url.QueryEscape(c.ClientSecret))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(userAgent, c.userAgent)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			msg, _ := io.ReadAll(resp.Body)
			formattedMessage := fmt.Sprintf("%s %s: HTTP %v Detail: %v Reason: %v",
				http.MethodPost, c.RevokeURL, resp.StatusCode,
				"Token Revocation request was rate limited", string(msg))
			return errors.New(formattedMessage)
		}
		formattedMessage := fmt.Sprintf("%s %s: HTTP %v Detail: %v Reason: %v",
			http.MethodPost, c.RevokeURL, resp.StatusCode,
			"Failed to revoke Access Token when fetching new OAuth Token from remote server",
			resp.Header.Get("www-authenticate"))
		return errors.New(formattedMessage)
	}
	return nil

}

func (c *tokenSource) Token() (*auth.Token, error) {
	return c.conf.Token(c.ctx)
}
