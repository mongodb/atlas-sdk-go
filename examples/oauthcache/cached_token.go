package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/atlas-sdk/v20241023002/auth"
	"golang.org/x/oauth2"
	"log"
	"os"
	"strings"

	"go.mongodb.org/atlas-sdk/v20241023002/admin"
	"go.mongodb.org/atlas-sdk/v20241023002/auth/clientcredentials"
)

const fileName = "token.test"

// Example for saving a token to disk and reloading it for reuse

// Required env variables to run example:
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
func main() {
	host := os.Getenv("MONGODB_ATLAS_BASE_URL")
	if host == "" {
		host = "https://cloud.mongodb.com"
	}

	// Fetch clientID and clientSecret from environment variables
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing required environment variables")
	}

	ctx := context.Background()
	conf := clientcredentials.NewConfig(clientID, clientSecret)
	if host != "https://cloud.mongodb.com" {
		conf.TokenURL = strings.Trim(host, "/") + clientcredentials.TokenAPIPath
		conf.RevokeURL = strings.Trim(host, "/") + clientcredentials.RevokeAPIPath
	}
	// uncomment to see in action
	// WARNING: handle the generated file with the same care as a password
	//token, err := conf.Token(ctx)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//if err := saveTokenToDisk(token); err != nil {
	//	log.Fatalln(err.Error())
	//}
	src := tokenFromDisk(ctx, conf)
	sdk, err := admin.NewClient(
		admin.UseBaseURL(host),
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

func saveTokenToDisk(token *oauth2.Token) error {
	val, err := json.Marshal(token)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, val, 0600)
}

func tokenFromDisk(ctx context.Context, conf *clientcredentials.Config) oauth2.TokenSource {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var token *auth.Token
	if err := json.Unmarshal(data, &token); err != nil {
		log.Fatalln(err.Error())
	}
	return auth.ReuseTokenSource(token, conf.TokenSource(ctx))
}
