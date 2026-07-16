# \DataLakePipelinesAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePipeline**](DataLakePipelinesAPI.md#CreatePipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines | Create One Data Lake Pipeline
[**DeletePipeline**](DataLakePipelinesAPI.md#DeletePipeline) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Remove One Data Lake Pipeline
[**DeletePipelineRun**](DataLakePipelinesAPI.md#DeletePipelineRun) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Delete One Pipeline Run Dataset
[**GetAvailablePipelineSchedules**](DataLakePipelinesAPI.md#GetAvailablePipelineSchedules) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSchedules | Return All Ingestion Schedules for One Data Lake Pipeline
[**GetAvailablePipelineSnapshots**](DataLakePipelinesAPI.md#GetAvailablePipelineSnapshots) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSnapshots | Return All Backup Snapshots for One Data Lake Pipeline
[**GetPipeline**](DataLakePipelinesAPI.md#GetPipeline) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Return One Data Lake Pipeline
[**GetPipelineRun**](DataLakePipelinesAPI.md#GetPipelineRun) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Return One Data Lake Pipeline Run
[**ListPipelineRuns**](DataLakePipelinesAPI.md#ListPipelineRuns) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs | Return All Data Lake Pipeline Runs in One Project
[**ListPipelines**](DataLakePipelinesAPI.md#ListPipelines) | **Get** /api/atlas/v2/groups/{groupId}/pipelines | Return All Data Lake Pipelines in One Project
[**PausePipeline**](DataLakePipelinesAPI.md#PausePipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/pause | Pause One Data Lake Pipeline
[**ResumePipeline**](DataLakePipelinesAPI.md#ResumePipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/resume | Resume One Data Lake Pipeline
[**TriggerPipeline**](DataLakePipelinesAPI.md#TriggerPipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/trigger | Trigger On-Demand Snapshot Ingestion
[**UpdatePipeline**](DataLakePipelinesAPI.md#UpdatePipeline) | **Patch** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Update One Data Lake Pipeline



## CreatePipeline

> DataLakeIngestionPipeline CreatePipeline(ctx, groupId, dataLakeIngestionPipeline DataLakeIngestionPipeline).Execute()

Create One Data Lake Pipeline


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    dataLakeIngestionPipeline := *admin.NewDataLakeIngestionPipeline() // DataLakeIngestionPipeline | 

    resp, r, err := sdk.DataLakePipelinesAPI.CreatePipeline(context.Background(), groupId, &dataLakeIngestionPipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.CreatePipeline`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreatePipeline`: DataLakeIngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.CreatePipeline`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataLakeIngestionPipeline** | [**DataLakeIngestionPipeline**](DataLakeIngestionPipeline.md) | Creates one Data Lake Pipeline. | 

### Return type

[**DataLakeIngestionPipeline**](DataLakeIngestionPipeline.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipeline

> DeletePipeline(ctx, groupId, pipelineName).Execute()

Remove One Data Lake Pipeline


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 

    r, err := sdk.DataLakePipelinesAPI.DeletePipeline(context.Background(), groupId, pipelineName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.DeletePipeline`: %v (%v)\n", err, r)
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
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineRequest struct via the builder pattern


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


## DeletePipelineRun

> any DeletePipelineRun(ctx, groupId, pipelineName, pipelineRunId).Execute()

Delete One Pipeline Run Dataset


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 
    pipelineRunId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.DataLakePipelinesAPI.DeletePipelineRun(context.Background(), groupId, pipelineName, pipelineRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.DeletePipelineRun`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeletePipelineRun`: any
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.DeletePipelineRun`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 
**pipelineRunId** | **string** | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailablePipelineSchedules

> []DiskBackupApiPolicyItem GetAvailablePipelineSchedules(ctx, groupId, pipelineName).Execute()

Return All Ingestion Schedules for One Data Lake Pipeline


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 

    resp, r, err := sdk.DataLakePipelinesAPI.GetAvailablePipelineSchedules(context.Background(), groupId, pipelineName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.GetAvailablePipelineSchedules`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAvailablePipelineSchedules`: []DiskBackupApiPolicyItem
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.GetAvailablePipelineSchedules`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePipelineSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DiskBackupApiPolicyItem**](DiskBackupApiPolicyItem.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailablePipelineSnapshots

> PaginatedBackupSnapshot GetAvailablePipelineSnapshots(ctx, groupId, pipelineName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).CompletedAfter(completedAfter).Execute()

Return All Backup Snapshots for One Data Lake Pipeline


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    completedAfter := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.DataLakePipelinesAPI.GetAvailablePipelineSnapshots(context.Background(), groupId, pipelineName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).CompletedAfter(completedAfter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.GetAvailablePipelineSnapshots`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAvailablePipelineSnapshots`: PaginatedBackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.GetAvailablePipelineSnapshots`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePipelineSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **completedAfter** | **time.Time** | Date and time after which MongoDB Cloud created the snapshot. If specified, MongoDB Cloud returns available backup snapshots created after this time and date only. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | 

### Return type

[**PaginatedBackupSnapshot**](PaginatedBackupSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipeline

> DataLakeIngestionPipeline GetPipeline(ctx, groupId, pipelineName).Execute()

Return One Data Lake Pipeline


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 

    resp, r, err := sdk.DataLakePipelinesAPI.GetPipeline(context.Background(), groupId, pipelineName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.GetPipeline`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPipeline`: DataLakeIngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.GetPipeline`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataLakeIngestionPipeline**](DataLakeIngestionPipeline.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineRun

> IngestionPipelineRun GetPipelineRun(ctx, groupId, pipelineName, pipelineRunId).Execute()

Return One Data Lake Pipeline Run


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 
    pipelineRunId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.DataLakePipelinesAPI.GetPipelineRun(context.Background(), groupId, pipelineName, pipelineRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.GetPipelineRun`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPipelineRun`: IngestionPipelineRun
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.GetPipelineRun`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 
**pipelineRunId** | **string** | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IngestionPipelineRun**](IngestionPipelineRun.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelineRuns

> PaginatedPipelineRun ListPipelineRuns(ctx, groupId, pipelineName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).CreatedBefore(createdBefore).Execute()

Return All Data Lake Pipeline Runs in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    createdBefore := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.DataLakePipelinesAPI.ListPipelineRuns(context.Background(), groupId, pipelineName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).CreatedBefore(createdBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.ListPipelineRuns`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPipelineRuns`: PaginatedPipelineRun
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.ListPipelineRuns`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **createdBefore** | **time.Time** | If specified, Atlas returns only Data Lake Pipeline runs initiated before this time and date. | 

### Return type

[**PaginatedPipelineRun**](PaginatedPipelineRun.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelines

> []DataLakeIngestionPipeline ListPipelines(ctx, groupId).Execute()

Return All Data Lake Pipelines in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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

    resp, r, err := sdk.DataLakePipelinesAPI.ListPipelines(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.ListPipelines`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPipelines`: []DataLakeIngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.ListPipelines`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DataLakeIngestionPipeline**](DataLakeIngestionPipeline.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PausePipeline

> DataLakeIngestionPipeline PausePipeline(ctx, groupId, pipelineName).Execute()

Pause One Data Lake Pipeline


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 

    resp, r, err := sdk.DataLakePipelinesAPI.PausePipeline(context.Background(), groupId, pipelineName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.PausePipeline`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `PausePipeline`: DataLakeIngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.PausePipeline`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPausePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataLakeIngestionPipeline**](DataLakeIngestionPipeline.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumePipeline

> DataLakeIngestionPipeline ResumePipeline(ctx, groupId, pipelineName).Execute()

Resume One Data Lake Pipeline


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 

    resp, r, err := sdk.DataLakePipelinesAPI.ResumePipeline(context.Background(), groupId, pipelineName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.ResumePipeline`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ResumePipeline`: DataLakeIngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.ResumePipeline`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataLakeIngestionPipeline**](DataLakeIngestionPipeline.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerPipeline

> IngestionPipelineRun TriggerPipeline(ctx, groupId, pipelineName, triggerIngestionPipelineRequest TriggerIngestionPipelineRequest).Execute()

Trigger On-Demand Snapshot Ingestion


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 
    triggerIngestionPipelineRequest := *admin.NewTriggerIngestionPipelineRequest("32b6e34b3d91647abb20e7b8") // TriggerIngestionPipelineRequest | 

    resp, r, err := sdk.DataLakePipelinesAPI.TriggerPipeline(context.Background(), groupId, pipelineName, &triggerIngestionPipelineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.TriggerPipeline`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `TriggerPipeline`: IngestionPipelineRun
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.TriggerPipeline`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **triggerIngestionPipelineRequest** | [**TriggerIngestionPipelineRequest**](TriggerIngestionPipelineRequest.md) | Triggers a single ingestion run of a snapshot. | 

### Return type

[**IngestionPipelineRun**](IngestionPipelineRun.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipeline

> DataLakeIngestionPipeline UpdatePipeline(ctx, groupId, pipelineName, dataLakeIngestionPipeline DataLakeIngestionPipeline).Execute()

Update One Data Lake Pipeline


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312022/admin"
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
    pipelineName := "pipelineName_example" // string | 
    dataLakeIngestionPipeline := *admin.NewDataLakeIngestionPipeline() // DataLakeIngestionPipeline | 

    resp, r, err := sdk.DataLakePipelinesAPI.UpdatePipeline(context.Background(), groupId, pipelineName, &dataLakeIngestionPipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesAPI.UpdatePipeline`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdatePipeline`: DataLakeIngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesAPI.UpdatePipeline`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataLakeIngestionPipeline** | [**DataLakeIngestionPipeline**](DataLakeIngestionPipeline.md) | Updates one Data Lake Pipeline. | 

### Return type

[**DataLakeIngestionPipeline**](DataLakeIngestionPipeline.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

