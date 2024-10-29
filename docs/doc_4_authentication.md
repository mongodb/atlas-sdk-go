# Authenticating with the Atlas Go SDK

The `atlas-sdk-go` library utilizes Digest authentication. You can [create an API key](https://www.mongodb.com/docs/atlas/configure-api-access/#create-an-api-key-in-an-organization) through the Atlas UI or the Atlas CLI.

To learn more about API authentication, refer to the [Atlas Administration API Authentication](https://www.mongodb.com/docs/atlas/api/api-authentication).

## Using the Atlas Go SDK in Your Code with Digest Authentication

To access different parts of the Atlas Admin API, construct a new Atlas SDK client and use its services. For example:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

	sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
	if err != nil {
		log.Fatalf("Error instantiating new client: %v", err)
	}
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
	if err != nil {
		log.Fatalf("Could not fetch projects: %v", err)
	}
	fmt.Printf("Response status: %v\n", response.Status)
	fmt.Printf("Projects: %v\n", projects)
}
```

## (Preview) Using the Atlas Go SDK with OAuth Authentication

The `go.mongodb.org/atlas-sdk/v20240805005/auth/credentials` package provides an OAuth [client_credentials](https://oauth.net/2/grant-types/client-credentials) implementation. 
The Atlas Client Credentials Grant OAuth Authentication allows applications to authenticate and authorize themselves programmatically, 
rather than on behalf of a user.

You can use the credentials package in the Atlas SDK Go to perform Atlas Admin API requests. 
Package supports usage outside the SDK for making requests directly via an Go HTTP client.

## OAuth Authentication

### Authenticating with OAuth ClientID and ClientSecret

To authenticate using OAuth, you can use the `NewTokenSource` function, which requires a `ClientID` and `ClientSecret`.

```go
package main

import (
	"context"
	"log"
	"os"
	"go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing CLIENT_ID or CLIENT_SECRET environment variables")
	}

	// Using ClientID and ClientSecret. No cache supported (nil).
	sdk, err := admin.NewClient(admin.UseOAuthAuth(clientID, clientSecret, nil))

	// Make API calls
}
```

This method of initialization provides **No Ability to Revoke API**: once a token is issued, it cannot be invalidated until it expires.

### Specifying a Token Cache

In this example, we will demonstrate how to use the OAuth Client Credentials flow with a custom token cache. The cache allows you to store OAuth tokens for reuse across application restarts, improving efficiency by reducing the number of token requests to the authorization server.

**Note:** A token cache is recommended for clients that run for a very short time. Requesting multiple tokens can be subject to OAuth token limits and rate limiting.

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"
	"go.mongodb.org/atlas-sdk/v20240805005/admin"
	"go.mongodb.org/atlas-sdk/v20240805005/auth/credentials"
)

type MyTokenCache struct {
	fileContent []byte
	mu          sync.Mutex
}

func (s *MyTokenCache) RetrieveToken(_ context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var tkn string
	err := json.Unmarshal(s.fileContent, &tkn)
	if err != nil {
		return "", err
	}

	return tkn, nil
}

func (s *MyTokenCache) SaveToken(_ context.Context, tkn string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	fileData, err := json.Marshal(tkn)
	if err != nil {
		return err
	}
	s.fileContent = fileData
	return nil
}

func main() {
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing CLIENT_ID or CLIENT_SECRET environment variables")
	}

	fileTokenCache := MyTokenCache{}
	tokenSource := credentials.NewTokenSourceWithOptions(credentials.AtlasTokenSourceOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenCache:   &fileTokenCache,
		BaseURL:      admin.PtrString("https://cloud.mongodb.com"),
	})

	httpClient := credentials.NewHTTPClientWithOAuthToken(tokenSource)
	ctx := context.Background()

	sdk, err := admin.NewClient(
		admin.UseHTTPClient(httpClient),
		admin.UseBaseURL("https://cloud.mongodb.com"),
		admin.UseDebug(true))

	// Make requests to the API
	// ...
	// Revoke Token
	_ = tokenSource.RevokeToken()
}
```

## Advanced Use Cases

### Revocation

Revocation invalidates the token before its expiration.
It should be used when "logging out" the current OAuth client and configuring a new one.

```go
err := tokenSource.RevokeToken()
if err != nil {
	log.Fatalf("Error: %v", err)
}
```

### Creating a Custom Transport

You can create a custom transport to inject the OAuth token into HTTP requests. 
The `OAuthCustomHTTPTransport` provides an `UnderlyingTransport` field that specifies the transport to be used for requests.

```go
tokenSource := credentials.NewTokenSourceWithOptions(credentials.AtlasTokenSourceOptions{
	ClientID:     clientID,
	ClientSecret: clientSecret,
})

transport := OAuthCustomHTTPTransport{
	UnderlyingTransport: <your_transport>,
	TokenSource:        tokenSource,
}
// Use transport with your own HTTP client or create a new one.
```

### Overriding the User Agent

You can override the default user agent by specifying it in the `AtlasTokenSourceOptions`.

```go
tokenSource := credentials.NewTokenSourceWithOptions(credentials.AtlasTokenSourceOptions{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	UserAgent:    "CustomUserAgent/1.0",
})
```
