# \CloudMigrationServiceApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLinkToken**](CloudMigrationServiceApi.md#CreateLinkToken) | **Post** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Create One Link-Token
[**CreatePushMigration**](CloudMigrationServiceApi.md#CreatePushMigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations | Migrate One Local Managed Cluster to MongoDB Atlas
[**CutoverMigration**](CloudMigrationServiceApi.md#CutoverMigration) | **Put** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId}/cutover | Cut Over One Migrated Cluster
[**DeleteLinkToken**](CloudMigrationServiceApi.md#DeleteLinkToken) | **Delete** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Remove One Link-Token
[**GetPushMigration**](CloudMigrationServiceApi.md#GetPushMigration) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId} | Return One Migration Job
[**GetValidationStatus**](CloudMigrationServiceApi.md#GetValidationStatus) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/validate/{validationId} | Return One Migration Validation Job
[**ListSourceProjects**](CloudMigrationServiceApi.md#ListSourceProjects) | **Get** /api/atlas/v2/orgs/{orgId}/liveMigrations/availableProjects | Return All Projects Available for Migration
[**ValidateMigration**](CloudMigrationServiceApi.md#ValidateMigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations/validate | Validate One Migration Request



## CreateLinkToken

> TargetOrg CreateLinkToken(ctx, orgId, targetOrgRequest TargetOrgRequest).Execute()

Create One Link-Token


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    targetOrgRequest := *openapiclient.NewTargetOrgRequest() // TargetOrgRequest | 

    resp, r, err := sdk.CloudMigrationServiceApi.CreateLinkToken(context.Background(), orgId, &targetOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.CreateLinkToken`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateLinkToken`: TargetOrg
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.CreateLinkToken`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetOrgRequest** | [**TargetOrgRequest**](TargetOrgRequest.md) | IP address access list entries associated with the migration. | 

### Return type

[**TargetOrg**](TargetOrg.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePushMigration

> LiveMigrationResponse CreatePushMigration(ctx, groupId, liveMigrationRequest20240530 LiveMigrationRequest20240530).Execute()

Migrate One Local Managed Cluster to MongoDB Atlas


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    liveMigrationRequest20240530 := *openapiclient.NewLiveMigrationRequest20240530(*openapiclient.NewDestination("ClusterName_example", "32b6e34b3d91647abb20e7b8", "HostnameSchemaType_example"), *openapiclient.NewSource("ClusterName_example", "32b6e34b3d91647abb20e7b8", false, false)) // LiveMigrationRequest20240530 | 

    resp, r, err := sdk.CloudMigrationServiceApi.CreatePushMigration(context.Background(), groupId, &liveMigrationRequest20240530).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.CreatePushMigration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreatePushMigration`: LiveMigrationResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.CreatePushMigration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveMigrationRequest20240530** | [**LiveMigrationRequest20240530**](LiveMigrationRequest20240530.md) | One migration to be created. | 

### Return type

[**LiveMigrationResponse**](LiveMigrationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CutoverMigration

> CutoverMigration(ctx, groupId, liveMigrationId).Execute()

Cut Over One Migrated Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    liveMigrationId := "6296fb4c7c7aa997cf94e9a8" // string | 

    r, err := sdk.CloudMigrationServiceApi.CutoverMigration(context.Background(), groupId, liveMigrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.CutoverMigration`: %v (%v)\n", err, r)
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
**liveMigrationId** | **string** | Unique 24-hexadecimal digit string that identifies the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCutoverMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLinkToken

> DeleteLinkToken(ctx, orgId).Execute()

Remove One Link-Token


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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

    r, err := sdk.CloudMigrationServiceApi.DeleteLinkToken(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.DeleteLinkToken`: %v (%v)\n", err, r)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushMigration

> LiveMigrationResponse GetPushMigration(ctx, groupId, liveMigrationId).Execute()

Return One Migration Job


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    liveMigrationId := "6296fb4c7c7aa997cf94e9a8" // string | 

    resp, r, err := sdk.CloudMigrationServiceApi.GetPushMigration(context.Background(), groupId, liveMigrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.GetPushMigration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPushMigration`: LiveMigrationResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.GetPushMigration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**liveMigrationId** | **string** | Unique 24-hexadecimal digit string that identifies the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LiveMigrationResponse**](LiveMigrationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidationStatus

> LiveImportValidation GetValidationStatus(ctx, groupId, validationId).Execute()

Return One Migration Validation Job


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    validationId := "507f1f77bcf86cd799439011" // string | 

    resp, r, err := sdk.CloudMigrationServiceApi.GetValidationStatus(context.Background(), groupId, validationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.GetValidationStatus`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetValidationStatus`: LiveImportValidation
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.GetValidationStatus`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**validationId** | **string** | Unique 24-hexadecimal digit string that identifies the validation job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LiveImportValidation**](LiveImportValidation.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceProjects

> []LiveImportAvailableProject ListSourceProjects(ctx, orgId).Execute()

Return All Projects Available for Migration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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

    resp, r, err := sdk.CloudMigrationServiceApi.ListSourceProjects(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.ListSourceProjects`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListSourceProjects`: []LiveImportAvailableProject
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.ListSourceProjects`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LiveImportAvailableProject**](LiveImportAvailableProject.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMigration

> LiveImportValidation ValidateMigration(ctx, groupId, liveMigrationRequest20240530 LiveMigrationRequest20240530).Execute()

Validate One Migration Request


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    liveMigrationRequest20240530 := *openapiclient.NewLiveMigrationRequest20240530(*openapiclient.NewDestination("ClusterName_example", "32b6e34b3d91647abb20e7b8", "HostnameSchemaType_example"), *openapiclient.NewSource("ClusterName_example", "32b6e34b3d91647abb20e7b8", false, false)) // LiveMigrationRequest20240530 | 

    resp, r, err := sdk.CloudMigrationServiceApi.ValidateMigration(context.Background(), groupId, &liveMigrationRequest20240530).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.ValidateMigration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ValidateMigration`: LiveImportValidation
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.ValidateMigration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveMigrationRequest20240530** | [**LiveMigrationRequest20240530**](LiveMigrationRequest20240530.md) | One migration to be validated. | 

### Return type

[**LiveImportValidation**](LiveImportValidation.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

