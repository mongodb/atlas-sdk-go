# \SharedTierRestoreJobsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupTenantRestore**](SharedTierRestoreJobsApi.md#CreateBackupTenantRestore) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore | Create One Restore Job for One M2 or M5 Cluster
[**GetBackupTenantRestore**](SharedTierRestoreJobsApi.md#GetBackupTenantRestore) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId} | Return One Restore Job for One M2 or M5 Cluster
[**ListBackupTenantRestores**](SharedTierRestoreJobsApi.md#ListBackupTenantRestores) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores | Return All Restore Jobs for One M2 or M5 Cluster



## CreateBackupTenantRestore

> TenantRestore CreateBackupTenantRestore(ctx, clusterName, groupId, tenantRestore TenantRestore).Execute()

Create One Restore Job for One M2 or M5 Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clusterName := "clusterName_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantRestore := *openapiclient.NewTenantRestore("32b6e34b3d91647abb20e7b8", "TargetDeploymentItemName_example") // TenantRestore | 

    resp, r, err := sdk.SharedTierRestoreJobsApi.CreateBackupTenantRestore(context.Background(), clusterName, groupId, &tenantRestore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierRestoreJobsApi.CreateBackupTenantRestore`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateBackupTenantRestore`: TenantRestore
    fmt.Fprintf(os.Stdout, "Response from `SharedTierRestoreJobsApi.CreateBackupTenantRestore`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupTenantRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantRestore** | [**TenantRestore**](TenantRestore.md) | The restore job details. | 

### Return type

[**TenantRestore**](TenantRestore.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupTenantRestore

> TenantRestore GetBackupTenantRestore(ctx, clusterName, groupId, restoreId).Execute()

Return One Restore Job for One M2 or M5 Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clusterName := "clusterName_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    restoreId := "restoreId_example" // string | 

    resp, r, err := sdk.SharedTierRestoreJobsApi.GetBackupTenantRestore(context.Background(), clusterName, groupId, restoreId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierRestoreJobsApi.GetBackupTenantRestore`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetBackupTenantRestore`: TenantRestore
    fmt.Fprintf(os.Stdout, "Response from `SharedTierRestoreJobsApi.GetBackupTenantRestore`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**restoreId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupTenantRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TenantRestore**](TenantRestore.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupTenantRestores

> PaginatedTenantRestore ListBackupTenantRestores(ctx, clusterName, groupId).Execute()

Return All Restore Jobs for One M2 or M5 Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clusterName := "clusterName_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.SharedTierRestoreJobsApi.ListBackupTenantRestores(context.Background(), clusterName, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierRestoreJobsApi.ListBackupTenantRestores`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListBackupTenantRestores`: PaginatedTenantRestore
    fmt.Fprintf(os.Stdout, "Response from `SharedTierRestoreJobsApi.ListBackupTenantRestores`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupTenantRestoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedTenantRestore**](PaginatedTenantRestore.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

