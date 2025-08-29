# \FlexRestoreJobsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlexRestoreJob**](FlexRestoreJobsApi.md#CreateFlexRestoreJob) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs | Restore One Snapshot of One Flex Cluster
[**GetFlexRestoreJob**](FlexRestoreJobsApi.md#GetFlexRestoreJob) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Flex Cluster
[**ListFlexRestoreJobs**](FlexRestoreJobsApi.md#ListFlexRestoreJobs) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs | Return All Restore Jobs for One Flex Cluster



## CreateFlexRestoreJob

> FlexBackupRestoreJob20241113 CreateFlexRestoreJob(ctx, groupId, name, flexBackupRestoreJobCreate20241113 FlexBackupRestoreJobCreate20241113).Execute()

Restore One Snapshot of One Flex Cluster


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

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    name := "name_example" // string | 
    flexBackupRestoreJobCreate20241113 := *openapiclient.NewFlexBackupRestoreJobCreate20241113("32b6e34b3d91647abb20e7b8", "TargetDeploymentItemName_example") // FlexBackupRestoreJobCreate20241113 | 

    resp, r, err := sdk.FlexRestoreJobsApi.CreateFlexRestoreJob(context.Background(), groupId, name, &flexBackupRestoreJobCreate20241113).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexRestoreJobsApi.CreateFlexRestoreJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateFlexRestoreJob`: FlexBackupRestoreJob20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexRestoreJobsApi.CreateFlexRestoreJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the flex cluster whose snapshot you want to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlexRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flexBackupRestoreJobCreate20241113** | [**FlexBackupRestoreJobCreate20241113**](FlexBackupRestoreJobCreate20241113.md) | Restores one snapshot of one flex cluster from the specified project. | 

### Return type

[**FlexBackupRestoreJob20241113**](FlexBackupRestoreJob20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-11-13+json
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlexRestoreJob

> FlexBackupRestoreJob20241113 GetFlexRestoreJob(ctx, groupId, name, restoreJobId).Execute()

Return One Restore Job for One Flex Cluster


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

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    name := "name_example" // string | 
    restoreJobId := "restoreJobId_example" // string | 

    resp, r, err := sdk.FlexRestoreJobsApi.GetFlexRestoreJob(context.Background(), groupId, name, restoreJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexRestoreJobsApi.GetFlexRestoreJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetFlexRestoreJob`: FlexBackupRestoreJob20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexRestoreJobsApi.GetFlexRestoreJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the flex cluster. | 
**restoreJobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlexRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlexBackupRestoreJob20241113**](FlexBackupRestoreJob20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlexRestoreJobs

> PaginatedApiAtlasFlexBackupRestoreJob20241113 ListFlexRestoreJobs(ctx, groupId, name).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Restore Jobs for One Flex Cluster


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

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    name := "name_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.FlexRestoreJobsApi.ListFlexRestoreJobs(context.Background(), groupId, name).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexRestoreJobsApi.ListFlexRestoreJobs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListFlexRestoreJobs`: PaginatedApiAtlasFlexBackupRestoreJob20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexRestoreJobsApi.ListFlexRestoreJobs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the flex cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFlexRestoreJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiAtlasFlexBackupRestoreJob20241113**](PaginatedApiAtlasFlexBackupRestoreJob20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

