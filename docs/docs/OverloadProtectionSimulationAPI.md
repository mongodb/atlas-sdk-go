# \OverloadProtectionSimulationAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterOverloadSimulation**](OverloadProtectionSimulationAPI.md#CreateClusterOverloadSimulation) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/overloadSimulations | Create One Overload Protection Simulation
[**DeleteClusterOverloadSimulation**](OverloadProtectionSimulationAPI.md#DeleteClusterOverloadSimulation) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/overloadSimulations/{simulationId} | Delete One Overload Protection Simulation
[**GetClusterOverloadSimulation**](OverloadProtectionSimulationAPI.md#GetClusterOverloadSimulation) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/overloadSimulations/{simulationId} | Return One Overload Protection Simulation
[**ListClusterOverloadSimulations**](OverloadProtectionSimulationAPI.md#ListClusterOverloadSimulations) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/overloadSimulations | Return All Overload Protection Simulations



## CreateClusterOverloadSimulation

> OverloadProtectionSimulationResponse CreateClusterOverloadSimulation(ctx, groupId, clusterName, overloadProtectionSimulationRequest OverloadProtectionSimulationRequest).Execute()

Create One Overload Protection Simulation


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    clusterName := "clusterName_example" // string | 
    overloadProtectionSimulationRequest := *admin.NewOverloadProtectionSimulationRequest(int(123)) // OverloadProtectionSimulationRequest | 

    resp, r, err := sdk.OverloadProtectionSimulationAPI.CreateClusterOverloadSimulation(context.Background(), groupId, clusterName, &overloadProtectionSimulationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverloadProtectionSimulationAPI.CreateClusterOverloadSimulation`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateClusterOverloadSimulation`: OverloadProtectionSimulationResponse
    fmt.Fprintf(os.Stdout, "Response from `OverloadProtectionSimulationAPI.CreateClusterOverloadSimulation`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster on which to start the overload protection simulation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterOverloadSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **overloadProtectionSimulationRequest** | [**OverloadProtectionSimulationRequest**](OverloadProtectionSimulationRequest.md) | Details of the overload protection simulation to start. Valid durations (in seconds): 900, 3600, 28800, 86400. | 

### Return type

[**OverloadProtectionSimulationResponse**](OverloadProtectionSimulationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterOverloadSimulation

> DeleteClusterOverloadSimulation(ctx, groupId, clusterName, simulationId).Execute()

Delete One Overload Protection Simulation


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    clusterName := "clusterName_example" // string | 
    simulationId := "simulationId_example" // string | 

    r, err := sdk.OverloadProtectionSimulationAPI.DeleteClusterOverloadSimulation(context.Background(), groupId, clusterName, simulationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverloadProtectionSimulationAPI.DeleteClusterOverloadSimulation`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Human-readable label that identifies the cluster on which the overload protection simulation is running. | 
**simulationId** | **string** | Unique 24-hexadecimal digit string that identifies the overload protection simulation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterOverloadSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterOverloadSimulation

> OverloadProtectionSimulationResponse GetClusterOverloadSimulation(ctx, groupId, clusterName, simulationId).Execute()

Return One Overload Protection Simulation


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    clusterName := "clusterName_example" // string | 
    simulationId := "simulationId_example" // string | 

    resp, r, err := sdk.OverloadProtectionSimulationAPI.GetClusterOverloadSimulation(context.Background(), groupId, clusterName, simulationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverloadProtectionSimulationAPI.GetClusterOverloadSimulation`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetClusterOverloadSimulation`: OverloadProtectionSimulationResponse
    fmt.Fprintf(os.Stdout, "Response from `OverloadProtectionSimulationAPI.GetClusterOverloadSimulation`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster on which the overload protection simulation is running. | 
**simulationId** | **string** | Unique 24-hexadecimal digit string that identifies the overload protection simulation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterOverloadSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OverloadProtectionSimulationResponse**](OverloadProtectionSimulationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterOverloadSimulations

> PaginatedOverloadProtectionSimulationResponse ListClusterOverloadSimulations(ctx, groupId, clusterName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Overload Protection Simulations


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    clusterName := "clusterName_example" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.OverloadProtectionSimulationAPI.ListClusterOverloadSimulations(context.Background(), groupId, clusterName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverloadProtectionSimulationAPI.ListClusterOverloadSimulations`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListClusterOverloadSimulations`: PaginatedOverloadProtectionSimulationResponse
    fmt.Fprintf(os.Stdout, "Response from `OverloadProtectionSimulationAPI.ListClusterOverloadSimulations`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster on which the overload protection simulations are running. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterOverloadSimulationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedOverloadProtectionSimulationResponse**](PaginatedOverloadProtectionSimulationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

