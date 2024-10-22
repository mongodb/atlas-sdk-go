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

// token represents the internal OAuth2 token structure
type token struct {
	AccessToken string    `json:"access_token"`
	Expiry      time.Time `json:"expiry,omitempty"`
	ExpiresIn   int       `json:"expires_in"`
}

// OAuthClient manages the OAuth token fetching and refreshing using a TokenSource.
type OAuthClient struct {
	clientID     string
	clientSecret string
	tokenURL     string
	token        *token
	tokenSource  TokenSource
	ctx          context.Context
}

// getValidToken retrieves the valid token, refreshing it if necessary.
func (c *OAuthClient) getValidToken() (*token, error) {
	// Try to retrieve the token string from the token source
	tokenString, err := c.tokenSource.RetrieveToken(c.ctx)
	if err != nil || tokenString == nil {
		return c.refreshToken()
	}

	// Parse the token string into the token structure (mock parse operation)
	c.token, err = parseToken(*tokenString)
	if err != nil || c.token.expired() {
		// Token is invalid or expired, refresh it
		return c.refreshToken()
	}

	return c.token, nil
}

// refreshToken fetches a new token and saves it using the token source.
func (c *OAuthClient) refreshToken() (*token, error) {
	newToken, err := c.fetchToken()
	if err != nil {
		return nil, err
	}

	// Save the access token string to the token source
	err = c.tokenSource.SaveToken(c.ctx, newToken.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to save token: %w", err)
	}

	c.token = newToken
	return newToken, nil
}

// fetchToken makes a manual POST request to Server (tokenUrl) to fetch the access token.
func (c *OAuthClient) fetchToken() (*token, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", c.tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to obtain token, status: " + resp.Status)
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	// Construct the token with expiry time
	token := &token{
		AccessToken: tokenResp.AccessToken,
		ExpiresIn:   tokenResp.ExpiresIn,
		Expiry:      time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second),
	}
	return token, nil
}

// SetAuthHeader sets the Authorization header with the access token.
func (t *token) SetAuthHeader(r *http.Request) {
	r.Header.Set("Authorization", "Bearer "+t.AccessToken)
}

// Additional time for Access Tokens to not expire.
const expiryDelta = 10 * time.Second

// expired checks if the token is close to expiring.
func (t *token) expired() bool {
	if t.Expiry.IsZero() {
		return false
	}
	return t.Expiry.Round(0).Add(-expiryDelta).Before(time.Now())
}

// Valid checks if the token is still valid (present and not expired)
func (t *token) Valid() bool {
	return t != nil && t.AccessToken != "" && !t.expired()
}

// ParseToken extracts expiry details from JWT token
func parseToken(accessToken string) (*token, error) {
	parts := strings.Split(accessToken, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid access token format")
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
		return nil, errors.New("token has expired")
	}

	return &token{
		AccessToken: accessToken,
		Expiry:      expiry,
		ExpiresIn:   int(time.Until(expiry).Seconds()),
	}, nil
}
