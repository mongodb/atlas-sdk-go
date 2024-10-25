package credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// tokenRemoteResponse represents successful response from token endpoint
var tokenRemoteResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
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
