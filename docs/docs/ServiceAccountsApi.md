# \ServiceAccountsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProjectServiceAccount**](ServiceAccountsApi.md#AddProjectServiceAccount) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}:invite | Assign One Service Account to One Project
[**CreateProjectServiceAccount**](ServiceAccountsApi.md#CreateProjectServiceAccount) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts | Create One Project Service Account
[**CreateProjectServiceAccountAccessList**](ServiceAccountsApi.md#CreateProjectServiceAccountAccessList) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList | Add Access List Entries for One Project Service Account
[**CreateProjectServiceAccountSecret**](ServiceAccountsApi.md#CreateProjectServiceAccountSecret) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets | Create One Project Service Account Secret
[**CreateServiceAccount**](ServiceAccountsApi.md#CreateServiceAccount) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts | Create One Organization Service Account
[**CreateServiceAccountAccessList**](ServiceAccountsApi.md#CreateServiceAccountAccessList) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList | Add Access List Entries for One Organization Service Account
[**CreateServiceAccountSecret**](ServiceAccountsApi.md#CreateServiceAccountSecret) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets | Create One Organization Service Account Secret
[**DeleteProjectServiceAccount**](ServiceAccountsApi.md#DeleteProjectServiceAccount) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Remove One Project Service Account
[**DeleteProjectServiceAccountAccessListEntry**](ServiceAccountsApi.md#DeleteProjectServiceAccountAccessListEntry) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList/{ipAddress} | Remove One Access List Entry from One Project Service Account
[**DeleteProjectServiceAccountSecret**](ServiceAccountsApi.md#DeleteProjectServiceAccountSecret) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets/{secretId} | Delete One Project Service Account Secret
[**DeleteServiceAccount**](ServiceAccountsApi.md#DeleteServiceAccount) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Delete One Organization Service Account
[**DeleteServiceAccountAccessListEntry**](ServiceAccountsApi.md#DeleteServiceAccountAccessListEntry) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList/{ipAddress} | Remove One Access List Entry from One Organization Service Account
[**DeleteServiceAccountSecret**](ServiceAccountsApi.md#DeleteServiceAccountSecret) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets/{secretId} | Delete One Organization Service Account Secret
[**GetProjectServiceAccount**](ServiceAccountsApi.md#GetProjectServiceAccount) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Return One Project Service Account
[**GetServiceAccount**](ServiceAccountsApi.md#GetServiceAccount) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Return One Organization Service Account
[**ListProjectServiceAccountAccessList**](ServiceAccountsApi.md#ListProjectServiceAccountAccessList) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList | Return All Access List Entries for One Project Service Account
[**ListProjectServiceAccounts**](ServiceAccountsApi.md#ListProjectServiceAccounts) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts | Return All Project Service Accounts
[**ListServiceAccountAccessList**](ServiceAccountsApi.md#ListServiceAccountAccessList) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList | Return All Access List Entries for One Organization Service Account
[**ListServiceAccountProjects**](ServiceAccountsApi.md#ListServiceAccountProjects) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/groups | Return All Service Account Project Assignments
[**ListServiceAccounts**](ServiceAccountsApi.md#ListServiceAccounts) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts | Return All Organization Service Accounts
[**UpdateProjectServiceAccount**](ServiceAccountsApi.md#UpdateProjectServiceAccount) | **Patch** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Update One Project Service Account
[**UpdateServiceAccount**](ServiceAccountsApi.md#UpdateServiceAccount) | **Patch** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Update One Organization Service Account



## AddProjectServiceAccount

> GroupServiceAccount AddProjectServiceAccount(ctx, clientId, groupId, groupServiceAccountRoleAssignment GroupServiceAccountRoleAssignment).Execute()

Assign One Service Account to One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    groupServiceAccountRoleAssignment := *openapiclient.NewGroupServiceAccountRoleAssignment([]string{"Roles_example"}) // GroupServiceAccountRoleAssignment | 

    resp, r, err := sdk.ServiceAccountsApi.AddProjectServiceAccount(context.Background(), clientId, groupId, &groupServiceAccountRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.AddProjectServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `AddProjectServiceAccount`: GroupServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.AddProjectServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectServiceAccountRequest struct via the builder pattern


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


## CreateProjectServiceAccount

> GroupServiceAccount CreateProjectServiceAccount(ctx, groupId, groupServiceAccountRequest GroupServiceAccountRequest).Execute()

Create One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    groupServiceAccountRequest := *openapiclient.NewGroupServiceAccountRequest("Description_example", "Name_example", []string{"Roles_example"}, int(8)) // GroupServiceAccountRequest | 

    resp, r, err := sdk.ServiceAccountsApi.CreateProjectServiceAccount(context.Background(), groupId, &groupServiceAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateProjectServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateProjectServiceAccount`: GroupServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateProjectServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectServiceAccountRequest struct via the builder pattern


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


## CreateProjectServiceAccountAccessList

> PaginatedServiceAccountIPAccessEntry CreateProjectServiceAccountAccessList(ctx, groupId, clientId, serviceAccountIPAccessListEntry []ServiceAccountIPAccessListEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Add Access List Entries for One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    serviceAccountIPAccessListEntry := []openapiclient.ServiceAccountIPAccessListEntry{*openapiclient.NewServiceAccountIPAccessListEntry()} // []ServiceAccountIPAccessListEntry | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.CreateProjectServiceAccountAccessList(context.Background(), groupId, clientId, &serviceAccountIPAccessListEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateProjectServiceAccountAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateProjectServiceAccountAccessList`: PaginatedServiceAccountIPAccessEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateProjectServiceAccountAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectServiceAccountAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAccountIPAccessListEntry** | [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | A list of access list entries to add to the access list of the specified Service Account for the project. | 
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
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


## CreateProjectServiceAccountSecret

> ServiceAccountSecret CreateProjectServiceAccountSecret(ctx, groupId, clientId, serviceAccountSecretRequest ServiceAccountSecretRequest).Execute()

Create One Project Service Account Secret


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    serviceAccountSecretRequest := *openapiclient.NewServiceAccountSecretRequest(int(8)) // ServiceAccountSecretRequest | 

    resp, r, err := sdk.ServiceAccountsApi.CreateProjectServiceAccountSecret(context.Background(), groupId, clientId, &serviceAccountSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateProjectServiceAccountSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateProjectServiceAccountSecret`: ServiceAccountSecret
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateProjectServiceAccountSecret`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectServiceAccountSecretRequest struct via the builder pattern


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


## CreateServiceAccount

> OrgServiceAccount CreateServiceAccount(ctx, orgId, orgServiceAccountRequest OrgServiceAccountRequest).Execute()

Create One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    orgServiceAccountRequest := *openapiclient.NewOrgServiceAccountRequest("Description_example", "Name_example", []string{"Roles_example"}, int(8)) // OrgServiceAccountRequest | 

    resp, r, err := sdk.ServiceAccountsApi.CreateServiceAccount(context.Background(), orgId, &orgServiceAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateServiceAccount`: OrgServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


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


## CreateServiceAccountAccessList

> PaginatedServiceAccountIPAccessEntry CreateServiceAccountAccessList(ctx, orgId, clientId, serviceAccountIPAccessListEntry []ServiceAccountIPAccessListEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Add Access List Entries for One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    serviceAccountIPAccessListEntry := []openapiclient.ServiceAccountIPAccessListEntry{*openapiclient.NewServiceAccountIPAccessListEntry()} // []ServiceAccountIPAccessListEntry | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServiceAccountsApi.CreateServiceAccountAccessList(context.Background(), orgId, clientId, &serviceAccountIPAccessListEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccountAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateServiceAccountAccessList`: PaginatedServiceAccountIPAccessEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccountAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAccountIPAccessListEntry** | [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | A list of access list entries to add to the access list of the specified Service Account for the organization. | 
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
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


## CreateServiceAccountSecret

> ServiceAccountSecret CreateServiceAccountSecret(ctx, orgId, clientId, serviceAccountSecretRequest ServiceAccountSecretRequest).Execute()

Create One Organization Service Account Secret


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    serviceAccountSecretRequest := *openapiclient.NewServiceAccountSecretRequest(int(8)) // ServiceAccountSecretRequest | 

    resp, r, err := sdk.ServiceAccountsApi.CreateServiceAccountSecret(context.Background(), orgId, clientId, &serviceAccountSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccountSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateServiceAccountSecret`: ServiceAccountSecret
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccountSecret`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountSecretRequest struct via the builder pattern


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


## DeleteProjectServiceAccount

> DeleteProjectServiceAccount(ctx, clientId, groupId).Execute()

Remove One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    r, err := sdk.ServiceAccountsApi.DeleteProjectServiceAccount(context.Background(), clientId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteProjectServiceAccount`: %v (%v)\n", err, r)
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

Other parameters are passed through a pointer to a apiDeleteProjectServiceAccountRequest struct via the builder pattern


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


## DeleteProjectServiceAccountAccessListEntry

> DeleteProjectServiceAccountAccessListEntry(ctx, groupId, clientId, ipAddress).Execute()

Remove One Access List Entry from One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    r, err := sdk.ServiceAccountsApi.DeleteProjectServiceAccountAccessListEntry(context.Background(), groupId, clientId, ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteProjectServiceAccountAccessListEntry`: %v (%v)\n", err, r)
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

Other parameters are passed through a pointer to a apiDeleteProjectServiceAccountAccessListEntryRequest struct via the builder pattern


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


## DeleteProjectServiceAccountSecret

> DeleteProjectServiceAccountSecret(ctx, clientId, secretId, groupId).Execute()

Delete One Project Service Account Secret


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    r, err := sdk.ServiceAccountsApi.DeleteProjectServiceAccountSecret(context.Background(), clientId, secretId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteProjectServiceAccountSecret`: %v (%v)\n", err, r)
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

Other parameters are passed through a pointer to a apiDeleteProjectServiceAccountSecretRequest struct via the builder pattern


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


## DeleteServiceAccount

> DeleteServiceAccount(ctx, clientId, orgId).Execute()

Delete One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    r, err := sdk.ServiceAccountsApi.DeleteServiceAccount(context.Background(), clientId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteServiceAccount`: %v (%v)\n", err, r)
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
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


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


## DeleteServiceAccountAccessListEntry

> DeleteServiceAccountAccessListEntry(ctx, orgId, clientId, ipAddress).Execute()

Remove One Access List Entry from One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    r, err := sdk.ServiceAccountsApi.DeleteServiceAccountAccessListEntry(context.Background(), orgId, clientId, ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteServiceAccountAccessListEntry`: %v (%v)\n", err, r)
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
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 
**ipAddress** | **string** | One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountAccessListEntryRequest struct via the builder pattern


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


## DeleteServiceAccountSecret

> DeleteServiceAccountSecret(ctx, clientId, secretId, orgId).Execute()

Delete One Organization Service Account Secret


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    r, err := sdk.ServiceAccountsApi.DeleteServiceAccountSecret(context.Background(), clientId, secretId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteServiceAccountSecret`: %v (%v)\n", err, r)
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
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountSecretRequest struct via the builder pattern


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


## GetProjectServiceAccount

> GroupServiceAccount GetProjectServiceAccount(ctx, groupId, clientId).Execute()

Return One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    resp, r, err := sdk.ServiceAccountsApi.GetProjectServiceAccount(context.Background(), groupId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetProjectServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetProjectServiceAccount`: GroupServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.GetProjectServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectServiceAccountRequest struct via the builder pattern


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


## GetServiceAccount

> OrgServiceAccount GetServiceAccount(ctx, orgId, clientId).Execute()

Return One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    resp, r, err := sdk.ServiceAccountsApi.GetServiceAccount(context.Background(), orgId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetServiceAccount`: OrgServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.GetServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountRequest struct via the builder pattern


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


## ListProjectServiceAccountAccessList

> PaginatedServiceAccountIPAccessEntry ListProjectServiceAccountAccessList(ctx, groupId, clientId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Access List Entries for One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    resp, r, err := sdk.ServiceAccountsApi.ListProjectServiceAccountAccessList(context.Background(), groupId, clientId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListProjectServiceAccountAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListProjectServiceAccountAccessList`: PaginatedServiceAccountIPAccessEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListProjectServiceAccountAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectServiceAccountAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
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


## ListProjectServiceAccounts

> PaginatedGroupServiceAccounts ListProjectServiceAccounts(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Project Service Accounts


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    resp, r, err := sdk.ServiceAccountsApi.ListProjectServiceAccounts(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListProjectServiceAccounts`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListProjectServiceAccounts`: PaginatedGroupServiceAccounts
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListProjectServiceAccounts`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectServiceAccountsRequest struct via the builder pattern


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


## ListServiceAccountAccessList

> PaginatedServiceAccountIPAccessEntry ListServiceAccountAccessList(ctx, orgId, clientId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Access List Entries for One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    resp, r, err := sdk.ServiceAccountsApi.ListServiceAccountAccessList(context.Background(), orgId, clientId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListServiceAccountAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListServiceAccountAccessList`: PaginatedServiceAccountIPAccessEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListServiceAccountAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
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


## ListServiceAccountProjects

> PaginatedServiceAccountGroup ListServiceAccountProjects(ctx, orgId, clientId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Service Account Project Assignments


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    resp, r, err := sdk.ServiceAccountsApi.ListServiceAccountProjects(context.Background(), orgId, clientId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListServiceAccountProjects`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListServiceAccountProjects`: PaginatedServiceAccountGroup
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListServiceAccountProjects`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**clientId** | **string** | The Client ID of the Service Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountProjectsRequest struct via the builder pattern


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


## ListServiceAccounts

> PaginatedOrgServiceAccounts ListServiceAccounts(ctx, orgId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Organization Service Accounts


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    resp, r, err := sdk.ServiceAccountsApi.ListServiceAccounts(context.Background(), orgId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListServiceAccounts`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListServiceAccounts`: PaginatedOrgServiceAccounts
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListServiceAccounts`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountsRequest struct via the builder pattern


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


## UpdateProjectServiceAccount

> GroupServiceAccount UpdateProjectServiceAccount(ctx, clientId, groupId, groupServiceAccountUpdateRequest GroupServiceAccountUpdateRequest).Execute()

Update One Project Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    groupServiceAccountUpdateRequest := *openapiclient.NewGroupServiceAccountUpdateRequest() // GroupServiceAccountUpdateRequest | 

    resp, r, err := sdk.ServiceAccountsApi.UpdateProjectServiceAccount(context.Background(), clientId, groupId, &groupServiceAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateProjectServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateProjectServiceAccount`: GroupServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateProjectServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectServiceAccountRequest struct via the builder pattern


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


## UpdateServiceAccount

> OrgServiceAccount UpdateServiceAccount(ctx, clientId, orgId, orgServiceAccountUpdateRequest OrgServiceAccountUpdateRequest).Execute()

Update One Organization Service Account


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    orgServiceAccountUpdateRequest := *openapiclient.NewOrgServiceAccountUpdateRequest() // OrgServiceAccountUpdateRequest | 

    resp, r, err := sdk.ServiceAccountsApi.UpdateServiceAccount(context.Background(), clientId, orgId, &orgServiceAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateServiceAccount`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateServiceAccount`: OrgServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateServiceAccount`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The Client ID of the Service Account. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountRequest struct via the builder pattern


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

