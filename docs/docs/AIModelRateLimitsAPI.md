# \AIModelRateLimitsAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupModelLimits**](AIModelRateLimitsAPI.md#GetGroupModelLimits) | **Get** /api/atlas/v2/groups/{groupId}/aiModelApiRateLimits | Return AI Model Rate Limits for One Group
[**GetGroupRateLimits**](AIModelRateLimitsAPI.md#GetGroupRateLimits) | **Get** /api/atlas/v2/groups/{groupId}/aiModelApiClouds/{cloud}/geographies/{geography}/modelGroupNames/{modelGroupName}/rateLimits | Return Single AI Model Rate Limit for One Group
[**GetOrgModelLimits**](AIModelRateLimitsAPI.md#GetOrgModelLimits) | **Get** /api/atlas/v2/orgs/{orgId}/aiModelApiRateLimits | Return AI Model Rate Limits for One Organization
[**GetOrgRateLimits**](AIModelRateLimitsAPI.md#GetOrgRateLimits) | **Get** /api/atlas/v2/orgs/{orgId}/aiModelApiClouds/{cloud}/geographies/{geography}/modelGroupNames/{modelGroupName}/rateLimits | Return Single AI Model Rate Limit for One Organization
[**ResetGroupModelLimits**](AIModelRateLimitsAPI.md#ResetGroupModelLimits) | **Post** /api/atlas/v2/groups/{groupId}/aiModelApiClouds/{cloud}/geographies/{geography}/modelGroupNames/{modelGroupName}/rateLimits:reset | Reset AI Model Rate Limit for One Model Group
[**ResetGroupRateLimits**](AIModelRateLimitsAPI.md#ResetGroupRateLimits) | **Post** /api/atlas/v2/groups/{groupId}/aiModelApiRateLimits:reset | Reset AI Model Rate Limits for Group
[**UpdateGroupRateLimits**](AIModelRateLimitsAPI.md#UpdateGroupRateLimits) | **Patch** /api/atlas/v2/groups/{groupId}/aiModelApiClouds/{cloud}/geographies/{geography}/modelGroupNames/{modelGroupName}/rateLimits | Update AI Model Rate Limit



## GetGroupModelLimits

> PaginatedAtlasAiModelRateLimitsResponse GetGroupModelLimits(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return AI Model Rate Limits for One Group


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312023/admin"
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

    resp, r, err := sdk.AIModelRateLimitsAPI.GetGroupModelLimits(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelRateLimitsAPI.GetGroupModelLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupModelLimits`: PaginatedAtlasAiModelRateLimitsResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelRateLimitsAPI.GetGroupModelLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupModelLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAtlasAiModelRateLimitsResponse**](PaginatedAtlasAiModelRateLimitsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRateLimits

> AiModelRateLimitResponse GetGroupRateLimits(ctx, groupId, cloud, geography, modelGroupName).Execute()

Return Single AI Model Rate Limit for One Group


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312023/admin"
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
    cloud := "cloud_example" // string | 
    geography := "geography_example" // string | 
    modelGroupName := "modelGroupName_example" // string | 

    resp, r, err := sdk.AIModelRateLimitsAPI.GetGroupRateLimits(context.Background(), groupId, cloud, geography, modelGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelRateLimitsAPI.GetGroupRateLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupRateLimits`: AiModelRateLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelRateLimitsAPI.GetGroupRateLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloud** | **string** | Cloud provider scope. Must be \&quot;ANY\&quot;. Additional values will be supported in future API versions. | 
**geography** | **string** | Geography scope. Must be \&quot;ANY\&quot;. Additional values will be supported in future API versions. | 
**modelGroupName** | **string** | The name of the model group to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AiModelRateLimitResponse**](AiModelRateLimitResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgModelLimits

> PaginatedAtlasAiModelRateLimitsResponse GetOrgModelLimits(ctx, orgId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return AI Model Rate Limits for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312023/admin"
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

    resp, r, err := sdk.AIModelRateLimitsAPI.GetOrgModelLimits(context.Background(), orgId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelRateLimitsAPI.GetOrgModelLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgModelLimits`: PaginatedAtlasAiModelRateLimitsResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelRateLimitsAPI.GetOrgModelLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgModelLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAtlasAiModelRateLimitsResponse**](PaginatedAtlasAiModelRateLimitsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgRateLimits

> AiModelRateLimitResponse GetOrgRateLimits(ctx, orgId, cloud, geography, modelGroupName).Execute()

Return Single AI Model Rate Limit for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312023/admin"
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
    cloud := "cloud_example" // string | 
    geography := "geography_example" // string | 
    modelGroupName := "modelGroupName_example" // string | 

    resp, r, err := sdk.AIModelRateLimitsAPI.GetOrgRateLimits(context.Background(), orgId, cloud, geography, modelGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelRateLimitsAPI.GetOrgRateLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgRateLimits`: AiModelRateLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelRateLimitsAPI.GetOrgRateLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**cloud** | **string** | Cloud provider scope. Must be \&quot;ANY\&quot;. Additional values will be supported in future API versions. | 
**geography** | **string** | Geography scope. Must be \&quot;ANY\&quot;. Additional values will be supported in future API versions. | 
**modelGroupName** | **string** | The name of the model group to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AiModelRateLimitResponse**](AiModelRateLimitResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetGroupModelLimits

> AiModelRateLimitResponse ResetGroupModelLimits(ctx, groupId, cloud, geography, modelGroupName).Execute()

Reset AI Model Rate Limit for One Model Group


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312023/admin"
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
    cloud := "cloud_example" // string | 
    geography := "geography_example" // string | 
    modelGroupName := "modelGroupName_example" // string | 

    resp, r, err := sdk.AIModelRateLimitsAPI.ResetGroupModelLimits(context.Background(), groupId, cloud, geography, modelGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelRateLimitsAPI.ResetGroupModelLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ResetGroupModelLimits`: AiModelRateLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelRateLimitsAPI.ResetGroupModelLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloud** | **string** | Cloud provider scope. Must be \&quot;ANY\&quot;. Additional values will be supported in future API versions. | 
**geography** | **string** | Geography scope. Must be \&quot;ANY\&quot;. Additional values will be supported in future API versions. | 
**modelGroupName** | **string** | The name of the model group to be reset to default rate limits. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetGroupModelLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AiModelRateLimitResponse**](AiModelRateLimitResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetGroupRateLimits

> PaginatedAtlasAiModelRateLimitsResponse ResetGroupRateLimits(ctx, groupId).Execute()

Reset AI Model Rate Limits for Group


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312023/admin"
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

    resp, r, err := sdk.AIModelRateLimitsAPI.ResetGroupRateLimits(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelRateLimitsAPI.ResetGroupRateLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ResetGroupRateLimits`: PaginatedAtlasAiModelRateLimitsResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelRateLimitsAPI.ResetGroupRateLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetGroupRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedAtlasAiModelRateLimitsResponse**](PaginatedAtlasAiModelRateLimitsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupRateLimits

> AiModelRateLimitResponse UpdateGroupRateLimits(ctx, groupId, cloud, geography, modelGroupName, aiModelsRateLimitUpdateRequest AiModelsRateLimitUpdateRequest).Execute()

Update AI Model Rate Limit


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312023/admin"
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
    cloud := "cloud_example" // string | 
    geography := "geography_example" // string | 
    modelGroupName := "modelGroupName_example" // string | 
    aiModelsRateLimitUpdateRequest := *admin.NewAiModelsRateLimitUpdateRequest(int(123), int(123)) // AiModelsRateLimitUpdateRequest | 

    resp, r, err := sdk.AIModelRateLimitsAPI.UpdateGroupRateLimits(context.Background(), groupId, cloud, geography, modelGroupName, &aiModelsRateLimitUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelRateLimitsAPI.UpdateGroupRateLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupRateLimits`: AiModelRateLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelRateLimitsAPI.UpdateGroupRateLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloud** | **string** | Cloud provider scope. Must be \&quot;ANY\&quot;. Additional values will be supported in future API versions. | 
**geography** | **string** | Geography scope. Must be \&quot;ANY\&quot;. Additional values will be supported in future API versions. | 
**modelGroupName** | **string** | The name of the model group to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **aiModelsRateLimitUpdateRequest** | [**AiModelsRateLimitUpdateRequest**](AiModelsRateLimitUpdateRequest.md) | A request containing the new rate limits for the given model group. | 

### Return type

[**AiModelRateLimitResponse**](AiModelRateLimitResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

