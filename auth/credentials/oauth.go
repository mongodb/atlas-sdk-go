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
)

// Token represents the internal OAuth2 Token structure
type Token struct {
	AccessToken string    `json:"access_token"`
	Expiry      time.Time `json:"expiry,omitempty"`
	ExpiresIn   int       `json:"expires_in"`
}

// TokenSource interface allows to fetch and revoke valid OAuth Access Tokens.
type TokenSource interface {
	// GetValidToken retrieves the valid Token, refreshing it if necessary.
	// Implementation will check if token exist in the cache,
	// otherwise we will fetch new token from server and cache it.
	// If cached token expired it will be automatically removed and new OAuth Token will be requested
	GetValidToken() (*Token, error)

	// RevokeToken revokes the Access Token while also removing it from the Access Token Cache.
	// When Access Token is expired or missing revoke will return without any action.
	RevokeToken() error
}

// OAuthTokenSource manages the OAuth Token fetching and refreshing using a LocalTokenCache.
type OAuthTokenSource struct {
	clientID     string
	clientSecret string
	userAgent    string
	tokenURL     string
	revokeURL    string
	token        *Token
	tokenCache   LocalTokenCache
	ctx          context.Context
}

func (c *OAuthTokenSource) RevokeToken() error {
	tokenString, err := c.tokenCache.RetrieveToken(c.ctx)
	if err != nil {
		return err
	}
	if tokenString != nil && *tokenString != "" {
		err := c.revokeTokenInRemoteServer(*tokenString)
		if err != nil {
			return err
		}

		// Remove token from cache
		err = c.tokenCache.SaveToken(c.ctx, "")
		if err != nil {
			return err
		}
	}
	// No revocation needed for empty tokens.
	return nil
}

// GetValidToken retrieves the valid Token, refreshing it if necessary.
func (c *OAuthTokenSource) GetValidToken() (*Token, error) {
	// Try to retrieve the Token string from the Token source
	tokenString, err := c.tokenCache.RetrieveToken(c.ctx)
	if err != nil || tokenString == nil {
		return c.refreshToken()
	}

	// Parse the Token string into the Token structure (mock parse operation)
	c.token, err = parseToken(*tokenString)
	if err != nil || c.token.expired() {
		// Token is invalid or expired, refresh it
		return c.refreshToken()
	}

	return c.token, nil
}

// refreshToken fetches a new Token and saves it using the Token source.
func (c *OAuthTokenSource) refreshToken() (*Token, error) {
	newToken, err := c.fetchTokenFromRemoteServer()
	if err != nil {
		return nil, err
	}

	// Save the access Token string to the Token source
	err = c.tokenCache.SaveToken(c.ctx, newToken.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to save Token: %w", err)
	}

	c.token = newToken
	return newToken, nil
}

// ExpiryDelta Additional time for Access Tokens to not expire.
const ExpiryDelta = 10 * time.Second

// expired checks if the Token is close to expiring.
func (t *Token) expired() bool {
	if t.Expiry.IsZero() {
		return false
	}
	return t.Expiry.Round(0).Add(-ExpiryDelta).Before(time.Now())
}

// Valid checks if the Token is still valid (present and not expired)
func (t *Token) Valid() bool {
	return t != nil && t.AccessToken != "" && !t.expired()
}

// ParseToken extracts expiry details from JWT Token
func parseToken(accessToken string) (*Token, error) {
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

	return &Token{
		AccessToken: accessToken,
		Expiry:      expiry,
		ExpiresIn:   int(time.Until(expiry).Seconds()),
	}, nil
}

// fetchToken makes a manual POST request to Server (tokenUrl) to fetch the access Token.
func (c *OAuthTokenSource) fetchTokenFromRemoteServer() (*Token, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", c.tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", c.userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			msg, _ := io.ReadAll(resp.Body)
			formattedMessage := fmt.Sprintf("%v %v: HTTP %v Detail: %v Reason: %v",
				"POST", c.tokenURL, resp.StatusCode,
				"Token request was rate limited", string(msg))
			return nil, errors.New(formattedMessage)
		}
		formattedMessage := fmt.Sprintf("%v %v: HTTP %v Detail: %v Reason: %v",
			"POST", c.tokenURL, resp.StatusCode,
			"Failed to obtain Access Token when fetching new OAuth Token from remote server",
			resp.Header.Get("www-authenticate"))
		return nil, errors.New(formattedMessage)
	}
	// tokenRemoteResponse represents successful response from token endpoint
	var tokenRemoteResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenRemoteResponse); err != nil {
		return nil, err
	}

	// Construct the Token with expiry time
	token := &Token{
		AccessToken: tokenRemoteResponse.AccessToken,
		ExpiresIn:   tokenRemoteResponse.ExpiresIn,
		Expiry:      time.Now().Add(time.Duration(tokenRemoteResponse.ExpiresIn) * time.Second),
	}
	return token, nil
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
				"POST", c.tokenURL, resp.StatusCode,
				"Token Revocation request was rate limited", string(msg))
			return errors.New(formattedMessage)
		}
		formattedMessage := fmt.Sprintf("%v %v: HTTP %v Detail: %v Reason: %v",
			"POST", c.tokenURL, resp.StatusCode,
			"Failed to revoke Access Token when fetching new OAuth Token from remote server",
			resp.Header.Get("www-authenticate"))
		return errors.New(formattedMessage)
	}
	return nil
}
