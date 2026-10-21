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

	"go.mongodb.org/atlas-sdk/v20250312027/admin"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

	sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
	if err != nil {
		log.Fatalf("Error instantiating new client: %v", err)
	}
	projects, response, err := sdk.ProjectsAPI.ListProjects(ctx).Execute()
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
For more information about the feature, please refer to [Service Account Public documentation.](https://www.mongodb.com/docs/atlas/api/service-accounts-overview/)

## OAuth Authentication

### Authenticating with OAuth ClientID and ClientSecret

### Admin API Authentication using Service Accounts

Authenticating using Service Accounts

```go
package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20250312027/admin"
)

func main() {
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing CLIENT_ID or CLIENT_SECRET environment variables")
	}

	// Using ClientID and ClientSecret. No cache supported (nil).
	sdk, err := admin.NewClient(admin.UseOAuthAuth(context.Background(), clientID, clientSecret))
	// Make API calls
}
```

For complete reference implementation,
please refer to [service account example](https://github.com/mongodb/atlas-sdk-go/tree/main/examples/sevice_account_management)


### Revocation

Revocation invalidates an OAuth token before its expiration date. 
This effectively "logs out" the current OAuth client and allows you to configure a new client.

```go
// Sounding code omitted for brevity 
revokeConfig := credentials.NewConfig(clientID, clientSecret)
revokeConfig.RevokeToken(context.Background(), &auth.Token{
	AccessToken: "yourTokenHere"
});
```
### OAuth Token Cache

Service Account OAuth Access Tokens validity is expressed as a number of seconds in the  `expires_in` field. 
Clients can cache these Access Tokens to mitigate rate limiting and adhere to [Access Token limits](https://www.mongodb.com/docs/manual/reference/limits/#mongodb-atlas-service-account-limits).

For an example on how to cache and reuse OAuth tokens in the Atlas SDK for Go, see the Service Account OAuth Token Cache 
[Service Account OAuth Token Cache example](https://github.com/mongodb/atlas-sdk-go/tree/main/examples/service_account_token_store).

## User Account Authentication with `auth/code`

The `auth/code` package exposes two user-authentication flows through a shared
`code.Config`. The Device Authorization Grant is retained for legacy consumers.
New interactive consumers should use Authorization Code + PKCE.

### Legacy Device Authorization Grant

The device flow uses `Config.AuthURL`, configurable with `code.SetAuthURL`. Its
methods are `RequestCode`, `GetToken`, `PollToken`, `RefreshToken`, `RevokeToken`,
and `RegistrationConfig`.

```go
// Assume context, fmt, and code are imported. config has a client ID and scopes.
func pollDeviceFlow(ctx context.Context, config *code.Config) error {
	deviceCode, _, err := config.RequestCode(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Open %s and enter %s\n", deviceCode.VerificationURI, deviceCode.UserCode)
	_, _, err = config.PollToken(ctx, deviceCode)
	if err != nil {
		return err
	}
	fmt.Println("Device authorization complete.")
	return nil
}
```

### Authorization Code + PKCE

The Authorization Code flow uses `Config.AuthServerURL`, configurable with
`code.SetAuthServerURL`. Use `DiscoverAuthServer` to obtain endpoint metadata,
`AuthorizationURL` to create the PKCE authorization URL, `ExchangeCode` to
exchange the callback code, and `RefreshAccessToken` or `RevokeAuthServerToken`
to manage tokens. The OAuth client ID must be recognized by the authorization
server for the requested resource. Register the loopback redirect path with the
client and pass the same path to `StartCallbackServer`; the listener selects an
ephemeral loopback port. The callback server validates `state`, validates any
present `iss` against the discovered issuer, and requires `iss` when discovery
metadata advertises `authorization_response_iss_parameter_supported`.

The production issuer advertises
`client_id_metadata_document_supported`, which indicates support for
Client ID Metadata Documents (CIMD). With CIMD, a client can use the HTTPS URL
of a JSON document containing its client metadata as its `client_id`; the
authorization server retrieves that document instead of relying on a manually
assigned client ID. The discovery flag indicates server capability; whether a
specific client can use CIMD depends on the authorization server's policy. See
[the OAuth Client ID Metadata Document draft](https://datatracker.ietf.org/doc/draft-ietf-oauth-client-id-metadata-document/)
for the document format.

`ExchangeCode` returns a `*code.Token`, distinct from the `auth.Token` used by
Service Account OAuth. Its `SetAuthHeader` method adds the access token as a
Bearer header. A complete Admin API example is in
[examples/auth_code](https://github.com/mongodb/atlas-sdk-go/tree/main/examples/auth_code).

```go
// Assume context and net/http are imported. token comes from ExchangeCode. apiURL is an Atlas endpoint.
func requestWithToken(ctx context.Context, token *code.Token, apiURL string) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, http.NoBody)
	if err != nil {
		return err
	}
	token.SetAuthHeader(request)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}
```
