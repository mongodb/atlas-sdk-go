package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/atlas-sdk/v20250312027/admin"
	"go.mongodb.org/atlas-sdk/v20250312027/auth/code"
)

const callbackPathEnv = "MONGODB_ATLAS_CALLBACK_PATH"

// accessTokenTransport adds a Bearer token only for the Atlas API host.
type accessTokenTransport struct {
	base  http.RoundTripper
	token *code.Token
}

func (t *accessTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	request := req.Clone(req.Context())
	if request.URL.Scheme != "https" || !strings.EqualFold(request.URL.Host, "cloud.mongodb.com") {
		return nil, fmt.Errorf("refusing to send Atlas access token to %q", request.URL.Host)
	}
	t.token.SetAuthHeader(request)
	return t.base.RoundTrip(request)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	if clientID == "" {
		return errors.New("set MONGODB_ATLAS_CLIENT_ID to a registered OAuth client ID")
	}
	callbackPath := os.Getenv(callbackPathEnv)
	if callbackPath == "" {
		return errors.New("set MONGODB_ATLAS_CALLBACK_PATH to the path registered for the OAuth client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	config, err := code.NewConfigWithOptions(nil,
		code.SetClientID(clientID),
		code.SetScopes([]string{"atlas"}),
	)
	if err != nil {
		return fmt.Errorf("creating auth config: %w", err)
	}
	metadata, _, err := config.DiscoverAuthServer(ctx)
	if err != nil {
		return fmt.Errorf("discovering authorization server: %w", err)
	}

	state, err := code.GenerateState()
	if err != nil {
		return fmt.Errorf("generating OAuth state: %w", err)
	}
	codeVerifier := code.GenerateVerifier()
	callback, err := code.StartCallbackServer(state, callbackPath, metadata)
	if err != nil {
		return fmt.Errorf("starting callback server: %w", err)
	}
	defer callback.Close()

	redirectURI := callback.RedirectURI()
	authorizationURL, err := config.AuthorizationURL(metadata, redirectURI, state, codeVerifier)
	if err != nil {
		return fmt.Errorf("creating authorization URL: %w", err)
	}
	fmt.Printf("Open this URL in a browser:\n%s\n", authorizationURL)

	authorizationCode, err := callback.WaitForCallback(ctx)
	if err != nil {
		return fmt.Errorf("waiting for authorization callback: %w", err)
	}
	token, _, err := config.ExchangeCode(ctx, metadata, authorizationCode, redirectURI, codeVerifier)
	if err != nil {
		return fmt.Errorf("exchanging authorization code: %w", err)
	}
	if !token.Valid() {
		return errors.New("authorization returned an invalid or expired token")
	}

	// This transport adds the code.Token Bearer header to Admin API requests.
	client := &http.Client{Transport: &accessTokenTransport{
		base:  http.DefaultTransport,
		token: token,
	}}
	atlas, err := admin.NewClient(admin.UseHTTPClient(client))
	if err != nil {
		return fmt.Errorf("creating Atlas Admin API client: %w", err)
	}
	projects, response, err := atlas.ProjectsAPI.ListGroups(ctx).Execute()
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("listing projects: %w", err)
	}
	fmt.Printf("Authorized. Visible projects: %d\n", projects.GetTotalCount())
	return nil
}
