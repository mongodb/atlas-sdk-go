package credentials

import (
	"go.mongodb.org/atlas-sdk/v20241023001/internal/core"
	"golang.org/x/oauth2"
)

// tokenAPIPath for getting OAuth Access Token from server
//
//nolint:gosec //url only
const tokenAPIPath = "/api/oauth/token"

// revokeAPIPath for revoking OAuth Access Token from server
const revokeAPIPath = "/api/oauth/revoke"

// serverURL for Token Atlas API
const serverTokenURL = core.DefaultCloudURL + tokenAPIPath

// serverURL for Revoke Atlas API
const serverRevokeURL = core.DefaultCloudURL + revokeAPIPath

func NewConfig(clientID, clientSecret string) *Config {
	c := &Config{}
	c.ClientID = clientID
	c.ClientSecret = clientSecret
	c.RevokeURL = serverRevokeURL
	c.TokenURL = serverTokenURL
	c.AuthStyle = oauth2.AuthStyleInHeader
	c.userAgent = core.DefaultUserAgent
	return c
}
