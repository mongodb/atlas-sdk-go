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

// NewTokenSource initializes OAuth Token Source that provides a way to get valid OAuth Tokens.
// See credentials.NewTokenSourceWithOptions for advanced use cases.
//func NewTokenSource(clientID, clientSecret string) RevocableTokenSource {
//	return NewTokenSourceWithOptions(AtlasTokenSourceOptions{
//		ClientID:     clientID,
//		ClientSecret: clientSecret,
//		UserAgent:    core.DefaultUserAgent,
//	})
//}

func NewConfig(clientID, clientSecret string) *Config {
	c := &Config{}
	c.ClientID = clientID
	c.ClientSecret = clientSecret
	c.RevokeURL = revokeAPIPath
	c.TokenURL = serverTokenURL
	c.AuthStyle = oauth2.AuthStyleInHeader
	return c
}
