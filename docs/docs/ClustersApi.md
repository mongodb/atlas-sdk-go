# \ClustersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterAdvancedConfiguration**](ClustersApi.md#GetClusterAdvancedConfiguration) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Return One Advanced Configuration Options for One Cluster
[**GetClusterStatus**](ClustersApi.md#GetClusterStatus) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/status | Return Status of All Cluster Operations
[**GetSampleDatasetLoadStatus**](ClustersApi.md#GetSampleDatasetLoadStatus) | **Get** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId} | Check Status of Cluster Sample Dataset Request
[**ListCloudProviderRegions**](ClustersApi.md#ListCloudProviderRegions) | **Get** /api/atlas/v2/groups/{groupId}/clusters/provider/regions | Return All Cloud Provider Regions
[**ListClustersForAllProjects**](ClustersApi.md#ListClustersForAllProjects) | **Get** /api/atlas/v2/clusters | Return All Authorized Clusters in All Projects
[**LoadSampleDataset**](ClustersApi.md#LoadSampleDataset) | **Post** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{name} | Load Sample Dataset Request into Cluster
[**UpdateClusterAdvancedConfiguration**](ClustersApi.md#UpdateClusterAdvancedConfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Update Advanced Configuration Options for One Cluster
[**UpgradeSharedCluster**](ClustersApi.md#UpgradeSharedCluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgrade | Upgrade One Shared-tier Cluster
[**UpgradeSharedClusterToServerless**](ClustersApi.md#UpgradeSharedClusterToServerless) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgradeToServerless | Upgrades One Shared-Tier Cluster to the Serverless Instance



## GetClusterAdvancedConfiguration

> ClusterDescriptionProcessArgs GetClusterAdvancedConfiguration(ctx, groupId, clusterName).Execute()

Return One Advanced Configuration Options for One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.ClustersApi.GetClusterAdvancedConfiguration(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterAdvancedConfiguration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetClusterAdvancedConfiguration`: ClusterDescriptionProcessArgs
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterAdvancedConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterAdvancedConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterDescriptionProcessArgs**](ClusterDescriptionProcessArgs.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatus

> ClusterStatus GetClusterStatus(ctx, groupId, clusterName).Execute()

Return Status of All Cluster Operations


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.ClustersApi.GetClusterStatus(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterStatus``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetClusterStatus`: ClusterStatus
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSampleDatasetLoadStatus

> SampleDatasetStatus GetSampleDatasetLoadStatus(ctx, groupId, sampleDatasetId).Execute()

Check Status of Cluster Sample Dataset Request


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    sampleDatasetId := "sampleDatasetId_example" // string | 

    resp, r, err := sdk.ClustersApi.GetSampleDatasetLoadStatus(context.Background(), groupId, sampleDatasetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetSampleDatasetLoadStatus``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetSampleDatasetLoadStatus`: SampleDatasetStatus
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetSampleDatasetLoadStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**sampleDatasetId** | **string** | Unique 24-hexadecimal digit string that identifies the loaded sample dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSampleDatasetLoadStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SampleDatasetStatus**](SampleDatasetStatus.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviderRegions

> PaginatedApiAtlasProviderRegions ListCloudProviderRegions(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Providers(providers).Tier(tier).Execute()

Return All Cloud Provider Regions


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    providers := []string{"Inner_example"} // []string |  (optional)
    tier := "tier_example" // string |  (optional)

    resp, r, err := sdk.ClustersApi.ListCloudProviderRegions(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Providers(providers).Tier(tier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListCloudProviderRegions``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListCloudProviderRegions`: PaginatedApiAtlasProviderRegions
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListCloudProviderRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProviderRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **providers** | **[]string** | Cloud providers whose regions to retrieve. When you specify multiple providers, the response can return only tiers and regions that support multi-cloud clusters. | 
 **tier** | **string** | Cluster tier for which to retrieve the regions. | 

### Return type

[**PaginatedApiAtlasProviderRegions**](PaginatedApiAtlasProviderRegions.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClustersForAllProjects

> PaginatedOrgGroup ListClustersForAllProjects(ctx).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Authorized Clusters in All Projects


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.ClustersApi.ListClustersForAllProjects(context.Background()).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClustersForAllProjects``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListClustersForAllProjects`: PaginatedOrgGroup
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClustersForAllProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClustersForAllProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedOrgGroup**](PaginatedOrgGroup.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadSampleDataset

> SampleDatasetStatus LoadSampleDataset(ctx, groupId, name).Execute()

Load Sample Dataset Request into Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    name := "name_example" // string | 

    resp, r, err := sdk.ClustersApi.LoadSampleDataset(context.Background(), groupId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.LoadSampleDataset``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `LoadSampleDataset`: SampleDatasetStatus
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.LoadSampleDataset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the cluster into which you load the sample dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadSampleDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SampleDatasetStatus**](SampleDatasetStatus.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterAdvancedConfiguration

> ClusterDescriptionProcessArgs UpdateClusterAdvancedConfiguration(ctx, groupId, clusterName, clusterDescriptionProcessArgs ClusterDescriptionProcessArgs).Execute()

Update Advanced Configuration Options for One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    clusterDescriptionProcessArgs := *openapiclient.NewClusterDescriptionProcessArgs() // ClusterDescriptionProcessArgs | 

    resp, r, err := sdk.ClustersApi.UpdateClusterAdvancedConfiguration(context.Background(), groupId, clusterName, &clusterDescriptionProcessArgs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateClusterAdvancedConfiguration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateClusterAdvancedConfiguration`: ClusterDescriptionProcessArgs
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateClusterAdvancedConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterAdvancedConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterDescriptionProcessArgs** | [**ClusterDescriptionProcessArgs**](ClusterDescriptionProcessArgs.md) | Advanced configuration details to add for one cluster in the specified project. | 

### Return type

[**ClusterDescriptionProcessArgs**](ClusterDescriptionProcessArgs.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSharedCluster

> LegacyAtlasCluster UpgradeSharedCluster(ctx, groupId, legacyAtlasCluster LegacyAtlasCluster).Execute()

Upgrade One Shared-tier Cluster


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    legacyAtlasCluster := *openapiclient.NewLegacyAtlasCluster() // LegacyAtlasCluster | 

    resp, r, err := sdk.ClustersApi.UpgradeSharedCluster(context.Background(), groupId, &legacyAtlasCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpgradeSharedCluster``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpgradeSharedCluster`: LegacyAtlasCluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpgradeSharedCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSharedClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **legacyAtlasCluster** | [**LegacyAtlasCluster**](LegacyAtlasCluster.md) | Details of the shared-tier cluster upgrade in the specified project. | 

### Return type

[**LegacyAtlasCluster**](LegacyAtlasCluster.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSharedClusterToServerless

> ServerlessInstanceDescription UpgradeSharedClusterToServerless(ctx, groupId, serverlessInstanceDescription ServerlessInstanceDescription).Execute()

Upgrades One Shared-Tier Cluster to the Serverless Instance


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    serverlessInstanceDescription := *openapiclient.NewServerlessInstanceDescription(*openapiclient.NewServerlessProviderSettings("BackingProviderName_example", "RegionName_example")) // ServerlessInstanceDescription | 

    resp, r, err := sdk.ClustersApi.UpgradeSharedClusterToServerless(context.Background(), groupId, &serverlessInstanceDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpgradeSharedClusterToServerless``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpgradeSharedClusterToServerless`: ServerlessInstanceDescription
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpgradeSharedClusterToServerless`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSharedClusterToServerlessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverlessInstanceDescription** | [**ServerlessInstanceDescription**](ServerlessInstanceDescription.md) | Details of the shared-tier cluster upgrade in the specified project. | 

### Return type

[**ServerlessInstanceDescription**](ServerlessInstanceDescription.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

