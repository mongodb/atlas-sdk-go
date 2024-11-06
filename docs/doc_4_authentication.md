# Authenticating with the Atlas Go SDK

The `atlas-sdk-go` library utilizes Digest authentication as its default authentication method.
You can [create an API key](https://www.mongodb.com/docs/atlas/configure-api-access/#create-an-api-key-in-an-organization) through the Atlas UI or the Atlas CLI.

To learn more about API authentication, refer to [Atlas Administration API Authentication](https://www.mongodb.com/docs/atlas/api/api-authentication).

## Using the Atlas Go SDK in Your Code with Digest Authentication

To access different parts of the Atlas Admin API, construct a new Atlas SDK client and use its services. For example:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20241023002/admin"
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

## (Preview) Using the Atlas Go SDK with Service Account Authentication

Atlas SDK Go provides OAuth Authentication using Service Accounts (currently available as a [Preview](https://www.mongodb.com/resources/beta-programs) feature)
A Service Account implements an OAuth [client_credentials](https://oauth.net/2/grant-types/client-credentials) grant.
For more information about feature please refer to [Service Account Public documentation.](https://www.mongodb.com/docs/atlas/api/service-accounts-overview/)

## OAuth Authentication

### Authenticating with OAuth ClientID and ClientSecret

### Admin API Authentication using Service Accounts

Authenticating using Service Accounts

```go
package main

import (
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20241023002/admin"
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

> NOTE: This method of initialization provides **No Ability to Revoke API**: once a token is issued, it cannot be invalidated until it expires.

### Specifying a Token Cache

In this example, we will demonstrate how to use the OAuth Client Credentials flow with a custom token cache.
The cache allows you to store OAuth tokens for reuse across application restarts, improving efficiency by reducing the number of token requests to the authorization server. This reduction in requests also minimizes the impact of OAuth Token Limits and Rate Limiting.

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	"go.mongodb.org/atlas-sdk/v20241023002/auth/credentials"
	"go.mongodb.org/atlas-sdk/v20241023002/admin"
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
	// Using ClientID and ClientSecret. No cache supported (nil).
	httpClient, err := admin.NewClient(admin.UseOAuthAuth(clientID, clientSecret, fileTokenCache))
	sdk, err := admin.NewClient(admin.UseHTTPClient(httpClient))

	// Make requests to the API
	// ...
	// Revoke Token
	_ = tokenSource.RevokeToken()
}
```

## Advanced Use Cases

### Revocation

Revocation invalidates an OAuth token before its expiration date. This effectively "logs out" the current OAuth client and allows you to configure a new client.

```go
// Sounding code omitted for brevity 
err := tokenSource.RevokeToken()
if err != nil {
	log.Fatalf("Error: %v", err)
}
```

### Creating a Custom Transport

You can create a custom transport to inject the OAuth token into HTTP requests. 
The `OAuthCustomHTTPTransport` provides an `UnderlyingTransport` field that specifies the transport to use for requests.

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

You can override the default user agent by specifying an agent of your choice in the `AtlasTokenSourceOptions`.

```go
tokenSource := credentials.NewTokenSourceWithOptions(credentials.AtlasTokenSourceOptions{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	UserAgent:    "CustomUserAgent/1.0",
})
```
