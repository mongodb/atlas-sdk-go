package credentials

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Token represents the OAuth2 token structure
type Token struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	Expiry      time.Time `json:"expiry,omitempty"`
	ExpiresIn   int       `json:"expires_in"`
}

// OAuthClient manages the OAuth token fetching and refreshing.
type OAuthClient struct {
	clientID     string
	clientSecret string
	tokenURL     string
	token        *Token
	ctx          context.Context
}

// GetAccessToken checks if the token is valid, and refreshes it if necessary.
func (c *OAuthClient) GetAccessToken() (*Token, error) {
	// If token is valid, return the existing token
	if c.token != nil && c.token.Valid() {
		return c.token, nil
	}

	// Fetch new token if none is available or it is expired
	token, err := c.fetchToken()
	if err != nil {
		return nil, err
	}
	c.token = token
	return token, nil
}

// fetchToken makes a manual POST request to the OAuth server to get the access token.
func (c *OAuthClient) fetchToken() (*Token, error) {
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
	token := &Token{
		AccessToken: tokenResp.AccessToken,
		TokenType:   tokenResp.TokenType,
		ExpiresIn:   tokenResp.ExpiresIn,
		Expiry:      time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second),
	}
	return token, nil
}

// CustomTransport is a custom HTTP transport that injects the OAuth2 token into requests.
type CustomTransport struct {
	underlyingTransport http.RoundTripper
	client              *OAuthClient
}

// RoundTrip implements the RoundTripper interface.
func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Ensure the token is valid
	token, err := t.client.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %v", err)
	}

	// Inject the token into the request
	token.SetAuthHeader(req)

	// Proceed with the underlying transport
	return t.underlyingTransport.RoundTrip(req)
}

// SetAuthHeader sets the Authorization header with the access token.
func (t *Token) SetAuthHeader(r *http.Request) {
	r.Header.Set("Authorization", "Bearer " + t.AccessToken)
}

const expiryDelta = 10 * time.Second

// expired checks if the token is close to expiring.
func (t *Token) expired() bool {
	if t.Expiry.IsZero() {
		return false
	}
	return t.Expiry.Round(0).Add(-expiryDelta).Before(time.Now())
}

// Valid checks if the token is still valid.
func (t *Token) Valid() bool {
	return t != nil && t.AccessToken != "" && !t.expired()
}

// NewServiceAccountClient initializes an OAuthClient with client credentials.
func NewServiceAccountOAuthClient(clientID, clientSecret string, baseURL *string) *OAuthClient {
	tokenURL := "https://cloud-dev.mongodb.com/api/oauth/token"
	if baseURL != nil {
		tokenURL = *baseURL + "/api/oauth/token";
	}
	
	return &OAuthClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		tokenURL:     tokenURL,
		token: &Token{
			// TODO add token storage
		},
		ctx: context.Background(),
	}
}

// NewHTTPClient creates an HTTP client with the custom transport.
func NewHTTPClientWithServiceAccountAuth(client *OAuthClient) *http.Client {
	return &http.Client{
		Transport: &CustomTransport{
			underlyingTransport: http.DefaultTransport,
			client:              client,
		},
	}
}
