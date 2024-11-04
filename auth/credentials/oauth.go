package credentials

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

// OAuthTokenSource manages the OAuth Token fetching and refreshing using a LocalTokenCache.
type OAuthTokenSource struct {
	// Wrapped TokenSource that will be used to obtain OAuth Tokens
	tokenSource oauth2.TokenSource
	tokenCache  LocalTokenCache

	// Customization
	ctx       context.Context
	userAgent string

	// Revocation
	revokeURL    string
	clientID     string
	clientSecret string
}

// RevokeToken revokes the Access Token while also removing it from the Access Token Cache.
// When the Access Token is expired or missing, revoke will return without any action.
func (c *OAuthTokenSource) RevokeToken() error {
	// We want to revoke token from cache only (avoid fetching new token to revoke it)
	tokenString, err := c.tokenCache.RetrieveToken(c.ctx)
	if err != nil {
		return err
	}
	if tokenString != "" {
		err := c.revokeTokenInRemoteServer(tokenString)
		if err != nil {
			return err
		}

		// Revoked token can be removed from the cache
		err = c.tokenCache.SaveToken(c.ctx, "")
		if err != nil {
			return err
		}
	}
	// No revocation needed for empty tokens.
	return nil
}

func (c *OAuthTokenSource) Token() (*oauth2.Token, error) {
	return c.tokenSource.Token()
}

// ReuseWrapperCacheTokenSource wraps a TokenSource to reuse a token until it expires.
func ReuseWrapperCacheTokenSource(src oauth2.TokenSource, cache LocalTokenCache) oauth2.TokenSource {
	return oauth2.ReuseTokenSource(nil, &tokenSourceWrapper{src: src, cache: cache})
}

type tokenSourceWrapper struct {
	src   oauth2.TokenSource
	cache LocalTokenCache
}

func (w *tokenSourceWrapper) Token() (*oauth2.Token, error) {
	tokenString, err := w.cache.RetrieveToken(context.Background())
	if err == nil && tokenString != "" {
		token, err := parseToken(tokenString)
		if err == nil && !isTokenExpired(token) {
			return &oauth2.Token{
				AccessToken: token.AccessToken,
				Expiry:      token.Expiry,
			}, nil
		}
	}

	token, err := w.src.Token()
	if err != nil {
		return nil, err
	}

	err = w.cache.SaveToken(context.Background(), token.AccessToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// ExpiryDelta is the Additional time for Access Tokens to not expire.
const ExpiryDelta = 10 * time.Second

// expired checks if the Token is close to expiring.
func isTokenExpired(token *oauth2.Token) bool {
	if token.Expiry.IsZero() {
		return false
	}
	return token.Expiry.Round(0).Add(-ExpiryDelta).Before(time.Now())
}

// parseToken extracts expiry details from JWT Token
func parseToken(accessToken string) (*oauth2.Token, error) {
	parts := strings.Split(accessToken, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid access Token format")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	var tokenData struct {
		Exp int64 `json:"exp"`
	}
	if err := json.Unmarshal(payload, &tokenData); err != nil {
		return nil, err
	}
	expiry := time.Unix(tokenData.Exp, 0)
	if time.Now().After(expiry) {
		return nil, errors.New("atlas cloud Access Token has expired")
	}
	return &oauth2.Token{
		AccessToken: accessToken,
		Expiry:      expiry,
	}, nil
}

// revokeToken revokes the provided access token by making a POST request to the OAuth revoke endpoint.
func (c *OAuthTokenSource) revokeTokenInRemoteServer(token string) error {
	revokeUrl := c.revokeURL
	data := url.Values{}
	data.Set("token", token)
	data.Set("token_type_hint", "access_token")

	req, err := http.NewRequest("POST", revokeUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", c.userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			msg, _ := io.ReadAll(resp.Body)
			formattedMessage := fmt.Sprintf("%v %v: HTTP %v Detail: %v Reason: %v",
				"POST", c.revokeURL, resp.StatusCode,
				"Token Revocation request was rate limited", string(msg))
			return errors.New(formattedMessage)
		}
		formattedMessage := fmt.Sprintf("%v %v: HTTP %v Detail: %v Reason: %v",
			"POST", c.revokeURL, resp.StatusCode,
			"Failed to revoke Access Token when fetching new OAuth Token from remote server",
			resp.Header.Get("www-authenticate"))
		return errors.New(formattedMessage)
	}
	return nil
}
