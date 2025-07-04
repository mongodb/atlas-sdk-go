package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mongodb/atlas-sdk-go/auth"

	"github.com/mongodb/atlas-sdk-go/admin"
	"github.com/mongodb/atlas-sdk-go/auth/clientcredentials"
)

// Variable provided as example.
var cachedToken []byte

// Example caching a Service Account OAuth tokens
// to memory and reloading it for reuse

// Required env variables to run example:
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
func main() {
	host := os.Getenv("MONGODB_ATLAS_BASE_URL")

	// Fetch clientID and clientSecret from environment variables
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing required environment variables")
	}

	// 1. Create instance of OAuth config
	conf := clientcredentials.NewConfig(clientID, clientSecret)
	if host != "" {
		conf.TokenURL = strings.Trim(host, "/") + clientcredentials.TokenAPIPath
		conf.RevokeURL = strings.Trim(host, "/") + clientcredentials.RevokeAPIPath
	}

	// 2. Read token from storage.
	ctx := context.Background()
	src := readTokenFromDisk(ctx, conf)
	defer func() {
		if err := saveTokenToDisk(src); err != nil {
			log.Fatalln(err.Error())
		}
	}()
	sdk, err := admin.NewClient(
		admin.UseBaseURL(host),
		admin.UseDebug(true),
		admin.UseHTTPClient(auth.NewClient(ctx, src)),
	)

	if err != nil {
		log.Fatalln(err.Error())
	}
	projects, _, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("Projects size: %d\n", projects.GetTotalCount())
}

func saveTokenToDisk(tokenSource auth.TokenSource) error {
	token, err := tokenSource.Token()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Fetch OAuth Token from memory
	val, err := json.Marshal(token)
	if err != nil {
		return err
	}

	fmt.Println("Saving token")
	// TODO example uses memory storage.
	// Please adjust method for your own use cases.
	// For example use WriteFile method
	// os.WriteFile("token", val, 0600)
	cachedToken = val
	return nil
}

func readTokenFromDisk(ctx context.Context, conf *clientcredentials.Config) auth.TokenSource {
	// TODO in real cases  replace with os.ReadFile("yourTokenCache")
	data := cachedToken // data, err := os.ReadFile("token")
	var token *auth.Token
	if data == nil {
		fmt.Println("No cached token detected")
		// No cached token return token source
		return conf.TokenSource(ctx)
	}
	fmt.Println("Cached token detected")
	if err := json.Unmarshal(data, &token); err != nil {
		log.Fatalln(err.Error())
	}
	return auth.ReuseTokenSource(token, conf.TokenSource(ctx))
}
