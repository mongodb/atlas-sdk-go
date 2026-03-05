# \ServiceAccountsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessList**](ServiceAccountsApi.md#CreateAccessList) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList | Add Access List Entries for One Project Service Account
[**CreateGroupSecret**](ServiceAccountsApi.md#CreateGroupSecret) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets | Create One Project Service Account Secret
[**CreateGroupServiceAccount**](ServiceAccountsApi.md#CreateGroupServiceAccount) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts | Create One Project Service Account
[**CreateOrgAccessList**](ServiceAccountsApi.md#CreateOrgAccessList) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList | Add Access List Entries for One Organization Service Account
[**CreateOrgSecret**](ServiceAccountsApi.md#CreateOrgSecret) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets | Create One Organization Service Account Secret
[**CreateOrgServiceAccount**](ServiceAccountsApi.md#CreateOrgServiceAccount) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts | Create One Organization Service Account
[**DeleteGroupAccessEntry**](ServiceAccountsApi.md#DeleteGroupAccessEntry) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList/{ipAddress} | Remove One Access List Entry from One Project Service Account
[**DeleteGroupSecret**](ServiceAccountsApi.md#DeleteGroupSecret) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets/{secretId} | Delete One Project Service Account Secret
[**DeleteGroupServiceAccount**](ServiceAccountsApi.md#DeleteGroupServiceAccount) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Remove One Project Service Account
[**DeleteOrgAccessEntry**](ServiceAccountsApi.md#DeleteOrgAccessEntry) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList/{ipAddress} | Remove One Access List Entry from One Organization Service Account
[**DeleteOrgSecret**](ServiceAccountsApi.md#DeleteOrgSecret) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets/{secretId} | Delete One Organization Service Account Secret
[**DeleteOrgServiceAccount**](ServiceAccountsApi.md#DeleteOrgServiceAccount) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Delete One Organization Service Account
[**GetGroupServiceAccount**](ServiceAccountsApi.md#GetGroupServiceAccount) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Return One Project Service Account
[**GetOrgServiceAccount**](ServiceAccountsApi.md#GetOrgServiceAccount) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Return One Organization Service Account
[**GetServiceAccountGroups**](ServiceAccountsApi.md#GetServiceAccountGroups) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/groups | Return All Service Account Project Assignments
[**InviteGroupServiceAccount**](ServiceAccountsApi.md#InviteGroupServiceAccount) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}:invite | Assign One Service Account to One Project
[**ListAccessList**](ServiceAccountsApi.md#ListAccessList) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList | Return All Access List Entries for One Project Service Account
[**ListGroupServiceAccounts**](ServiceAccountsApi.md#ListGroupServiceAccounts) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts | Return All Project Service Accounts
[**ListOrgAccessList**](ServiceAccountsApi.md#ListOrgAccessList) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList | Return All Access List Entries for One Organization Service Account
[**ListOrgServiceAccounts**](ServiceAccountsApi.md#ListOrgServiceAccounts) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts | Return All Organization Service Accounts
[**UpdateGroupServiceAccount**](ServiceAccountsApi.md#UpdateGroupServiceAccount) | **Patch** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Update One Project Service Account
[**UpdateOrgServiceAccount**](ServiceAccountsApi.md#UpdateOrgServiceAccount) | **Patch** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Update One Organization Service Account



## CreateAccessList

> PaginatedServiceAccountIPAccessEntry CreateAccessList(ctx, groupId, clientId, serviceAccountIPAccessListEntry []ServiceAccountIPAccessListEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Add Access List Entries for One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    serviceAccountIPAccessListEntry := []admin.ServiceAccountIPAccessListEntry{*admin.NewServiceAccountIPAccessListEntry()} // []ServiceAccountIPAccessListEntry | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.CreateAccessList(context.Background(), groupId, clientId, &serviceAccountIPAccessListEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateAccessList`: PaginatedServiceAccountIPAccessEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAccountIPAccessListEntry** | [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | A list of access list entries to add to the access list of the specified Service Account for the project. | 
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedServiceAccountIPAccessEntry**](PaginatedServiceAccountIPAccessEntry.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupSecret

> ServiceAccountSecret CreateGroupSecret(ctx, groupId, clientId, serviceAccountSecretRequest ServiceAccountSecretRequest).Execute()

Create One Project Service Account Secret


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    serviceAccountSecretRequest := *admin.NewServiceAccountSecretRequest(int(8)) // ServiceAccountSecretRequest | 

    resp, r, err := sdk.ServiceAccountsApi.CreateGroupSecret(context.Background(), groupId, clientId, &serviceAccountSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateGroupSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupSecret`: ServiceAccountSecret
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateGroupSecret`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAccountSecretRequest** | [**ServiceAccountSecretRequest**](ServiceAccountSecretRequest.md) | Details for the new secret. | 

### Return type

[**ServiceAccountSecret**](ServiceAccountSecret.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupServiceAccount

> GroupServiceAccount CreateGroupServiceAccount(ctx, groupId, groupServiceAccountRequest GroupServiceAccountRequest).Execute()

Create One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    groupServiceAccountRequest := *admin.NewGroupServiceAccountRequest("Description_example", "Name_example", []string{"Roles_example"}, int(8)) // GroupServiceAccountRequest | 

    resp, r, err := sdk.ServiceAccountsApi.CreateGroupServiceAccount(context.Background(), groupId, &groupServiceAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateGroupServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupServiceAccount`: GroupServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateGroupServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupServiceAccountRequest** | [**GroupServiceAccountRequest**](GroupServiceAccountRequest.md) | Details of the new Service Account. | 

### Return type

[**GroupServiceAccount**](GroupServiceAccount.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgAccessList

> PaginatedServiceAccountIPAccessEntry CreateOrgAccessList(ctx, orgId, clientId, serviceAccountIPAccessListEntry []ServiceAccountIPAccessListEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Add Access List Entries for One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    serviceAccountIPAccessListEntry := []admin.ServiceAccountIPAccessListEntry{*admin.NewServiceAccountIPAccessListEntry()} // []ServiceAccountIPAccessListEntry | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.CreateOrgAccessList(context.Background(), orgId, clientId, &serviceAccountIPAccessListEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateOrgAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgAccessList`: PaginatedServiceAccountIPAccessEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateOrgAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAccountIPAccessListEntry** | [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | A list of access list entries to add to the access list of the specified Service Account for the organization. | 
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedServiceAccountIPAccessEntry**](PaginatedServiceAccountIPAccessEntry.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgSecret

> ServiceAccountSecret CreateOrgSecret(ctx, orgId, clientId, serviceAccountSecretRequest ServiceAccountSecretRequest).Execute()

Create One Organization Service Account Secret


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    serviceAccountSecretRequest := *admin.NewServiceAccountSecretRequest(int(8)) // ServiceAccountSecretRequest | 

    resp, r, err := sdk.ServiceAccountsApi.CreateOrgSecret(context.Background(), orgId, clientId, &serviceAccountSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateOrgSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgSecret`: ServiceAccountSecret
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateOrgSecret`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAccountSecretRequest** | [**ServiceAccountSecretRequest**](ServiceAccountSecretRequest.md) | Details for the new secret. | 

### Return type

[**ServiceAccountSecret**](ServiceAccountSecret.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgServiceAccount

> OrgServiceAccount CreateOrgServiceAccount(ctx, orgId, orgServiceAccountRequest OrgServiceAccountRequest).Execute()

Create One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    orgServiceAccountRequest := *admin.NewOrgServiceAccountRequest("Description_example", "Name_example", []string{"Roles_example"}, int(8)) // OrgServiceAccountRequest | 

    resp, r, err := sdk.ServiceAccountsApi.CreateOrgServiceAccount(context.Background(), orgId, &orgServiceAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateOrgServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgServiceAccount`: OrgServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateOrgServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgServiceAccountRequest** | [**OrgServiceAccountRequest**](OrgServiceAccountRequest.md) | Details of the new Service Account. | 

### Return type

[**OrgServiceAccount**](OrgServiceAccount.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupAccessEntry

> DeleteGroupAccessEntry(ctx, groupId, clientId, ipAddress).Execute()

Remove One Access List Entry from One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    ipAddress := "192.0.2.0%2F24" // string | 

    r, err := sdk.ServiceAccountsApi.DeleteGroupAccessEntry(context.Background(), groupId, clientId, ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteGroupAccessEntry`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 
**ipAddress** | **string** | One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupAccessEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupSecret

> DeleteGroupSecret(ctx, clientId, secretId, groupId).Execute()

Delete One Project Service Account Secret


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    secretId := "secretId_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.ServiceAccountsApi.DeleteGroupSecret(context.Background(), clientId, secretId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteGroupSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**secretId** | **string** | Unique 24-hexadecimal digit string that identifies the secret. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupServiceAccount

> DeleteGroupServiceAccount(ctx, clientId, groupId).Execute()

Remove One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.ServiceAccountsApi.DeleteGroupServiceAccount(context.Background(), clientId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteGroupServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgAccessEntry

> DeleteOrgAccessEntry(ctx, orgId, clientId, ipAddress).Execute()

Remove One Access List Entry from One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    ipAddress := "192.0.2.0%2F24" // string | 

    r, err := sdk.ServiceAccountsApi.DeleteOrgAccessEntry(context.Background(), orgId, clientId, ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteOrgAccessEntry`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 
**ipAddress** | **string** | One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgAccessEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgSecret

> DeleteOrgSecret(ctx, clientId, secretId, orgId).Execute()

Delete One Organization Service Account Secret


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    secretId := "secretId_example" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 

    r, err := sdk.ServiceAccountsApi.DeleteOrgSecret(context.Background(), clientId, secretId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteOrgSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**secretId** | **string** | Unique 24-hexadecimal digit string that identifies the secret. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgServiceAccount

> DeleteOrgServiceAccount(ctx, clientId, orgId).Execute()

Delete One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 

    r, err := sdk.ServiceAccountsApi.DeleteOrgServiceAccount(context.Background(), clientId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteOrgServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupServiceAccount

> GroupServiceAccount GetGroupServiceAccount(ctx, groupId, clientId).Execute()

Return One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 

    resp, r, err := sdk.ServiceAccountsApi.GetGroupServiceAccount(context.Background(), groupId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetGroupServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupServiceAccount`: GroupServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.GetGroupServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupServiceAccount**](GroupServiceAccount.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgServiceAccount

> OrgServiceAccount GetOrgServiceAccount(ctx, orgId, clientId).Execute()

Return One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 

    resp, r, err := sdk.ServiceAccountsApi.GetOrgServiceAccount(context.Background(), orgId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetOrgServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgServiceAccount`: OrgServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.GetOrgServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgServiceAccount**](OrgServiceAccount.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountGroups

> PaginatedServiceAccountGroup GetServiceAccountGroups(ctx, orgId, clientId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Service Account Project Assignments


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.GetServiceAccountGroups(context.Background(), orgId, clientId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetServiceAccountGroups`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetServiceAccountGroups`: PaginatedServiceAccountGroup
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.GetServiceAccountGroups`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedServiceAccountGroup**](PaginatedServiceAccountGroup.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteGroupServiceAccount

> GroupServiceAccount InviteGroupServiceAccount(ctx, clientId, groupId, groupServiceAccountRoleAssignment GroupServiceAccountRoleAssignment).Execute()

Assign One Service Account to One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    groupServiceAccountRoleAssignment := *admin.NewGroupServiceAccountRoleAssignment([]string{"Roles_example"}) // GroupServiceAccountRoleAssignment | 

    resp, r, err := sdk.ServiceAccountsApi.InviteGroupServiceAccount(context.Background(), clientId, groupId, &groupServiceAccountRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.InviteGroupServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `InviteGroupServiceAccount`: GroupServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.InviteGroupServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupServiceAccountRoleAssignment** | [**GroupServiceAccountRoleAssignment**](GroupServiceAccountRoleAssignment.md) | The Project permissions for the Service Account in the specified Project. | 

### Return type

[**GroupServiceAccount**](GroupServiceAccount.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessList

> PaginatedServiceAccountIPAccessEntry ListAccessList(ctx, groupId, clientId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Access List Entries for One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.ListAccessList(context.Background(), groupId, clientId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListAccessList`: PaginatedServiceAccountIPAccessEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedServiceAccountIPAccessEntry**](PaginatedServiceAccountIPAccessEntry.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupServiceAccounts

> PaginatedGroupServiceAccounts ListGroupServiceAccounts(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Project Service Accounts


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.ListGroupServiceAccounts(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListGroupServiceAccounts`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupServiceAccounts`: PaginatedGroupServiceAccounts
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListGroupServiceAccounts`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedGroupServiceAccounts**](PaginatedGroupServiceAccounts.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgAccessList

> PaginatedServiceAccountIPAccessEntry ListOrgAccessList(ctx, orgId, clientId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Access List Entries for One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.ListOrgAccessList(context.Background(), orgId, clientId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListOrgAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgAccessList`: PaginatedServiceAccountIPAccessEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListOrgAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedServiceAccountIPAccessEntry**](PaginatedServiceAccountIPAccessEntry.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgServiceAccounts

> PaginatedOrgServiceAccounts ListOrgServiceAccounts(ctx, orgId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Organization Service Accounts


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.ListOrgServiceAccounts(context.Background(), orgId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListOrgServiceAccounts`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgServiceAccounts`: PaginatedOrgServiceAccounts
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListOrgServiceAccounts`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedOrgServiceAccounts**](PaginatedOrgServiceAccounts.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupServiceAccount

> GroupServiceAccount UpdateGroupServiceAccount(ctx, clientId, groupId, groupServiceAccountUpdateRequest GroupServiceAccountUpdateRequest).Execute()

Update One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    groupServiceAccountUpdateRequest := *admin.NewGroupServiceAccountUpdateRequest() // GroupServiceAccountUpdateRequest | 

    resp, r, err := sdk.ServiceAccountsApi.UpdateGroupServiceAccount(context.Background(), clientId, groupId, &groupServiceAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateGroupServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupServiceAccount`: GroupServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateGroupServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupServiceAccountUpdateRequest** | [**GroupServiceAccountUpdateRequest**](GroupServiceAccountUpdateRequest.md) | The new details for the Service Account. | 

### Return type

[**GroupServiceAccount**](GroupServiceAccount.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgServiceAccount

> OrgServiceAccount UpdateOrgServiceAccount(ctx, clientId, orgId, orgServiceAccountUpdateRequest OrgServiceAccountUpdateRequest).Execute()

Update One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clientId := "mdb_sa_id_1234567890abcdef12345678" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 
    orgServiceAccountUpdateRequest := *admin.NewOrgServiceAccountUpdateRequest() // OrgServiceAccountUpdateRequest | 

    resp, r, err := sdk.ServiceAccountsApi.UpdateOrgServiceAccount(context.Background(), clientId, orgId, &orgServiceAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateOrgServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateOrgServiceAccount`: OrgServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateOrgServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgServiceAccountUpdateRequest** | [**OrgServiceAccountUpdateRequest**](OrgServiceAccountUpdateRequest.md) | The new details for the Service Account. | 

### Return type

[**OrgServiceAccount**](OrgServiceAccount.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

