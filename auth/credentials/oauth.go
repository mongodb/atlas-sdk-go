package credentials

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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

// TokenSource interface allows to fetch valid OAuth Token.
type TokenSource interface {
	// GetValidToken retrieves the valid Token, refreshing it if necessary.
	GetValidToken() (*Token, error)
}

// OAuthTokenSource manages the OAuth Token fetching and refreshing using a LocalTokenCache.
type OAuthTokenSource struct {
	clientID     string
	clientSecret string
	userAgent    string
	tokenURL     string
	token        *Token
	tokenCache   LocalTokenCache
	ctx          context.Context
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
	newToken, err := c.fetchToken()
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

// fetchToken makes a manual POST request to Server (tokenUrl) to fetch the access Token.
func (c *OAuthTokenSource) fetchToken() (*Token, error) {
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
		return nil, errors.New("failed to obtain Token, status: " + resp.Status)
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	// Construct the Token with expiry time
	token := &Token{
		AccessToken: tokenResp.AccessToken,
		ExpiresIn:   tokenResp.ExpiresIn,
		Expiry:      time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second),
	}
	return token, nil
}

// Additional time for Access Tokens to not expire.
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
		return nil, errors.New("Token has expired")
	}

	return &Token{
		AccessToken: accessToken,
		Expiry:      expiry,
		ExpiresIn:   int(time.Until(expiry).Seconds()),
	}, nil
}
