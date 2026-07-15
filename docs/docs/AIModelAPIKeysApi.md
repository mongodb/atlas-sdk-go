# \AIModelAPIKeysApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupModelKey**](AIModelAPIKeysApi.md#CreateGroupModelKey) | **Post** /api/atlas/v2/groups/{groupId}/aiModelApiKeys | Create New AI Model API Key
[**DeleteGroupModelKey**](AIModelAPIKeysApi.md#DeleteGroupModelKey) | **Delete** /api/atlas/v2/groups/{groupId}/aiModelApiKeys/{apiKeyId} | Delete Existing AI Model API Key
[**GetGroupModelKey**](AIModelAPIKeysApi.md#GetGroupModelKey) | **Get** /api/atlas/v2/groups/{groupId}/aiModelApiKeys/{apiKeyId} | Return Single AI Model API Key for One Group
[**GetOrgModelKey**](AIModelAPIKeysApi.md#GetOrgModelKey) | **Get** /api/atlas/v2/orgs/{orgId}/aiModelApiKeys/{apiKeyId} | Return Single AI Model API Key for One Organization
[**ListGroupModelKeys**](AIModelAPIKeysApi.md#ListGroupModelKeys) | **Get** /api/atlas/v2/groups/{groupId}/aiModelApiKeys | Return AI Model API Keys for One Group
[**ListOrgModelKeys**](AIModelAPIKeysApi.md#ListOrgModelKeys) | **Get** /api/atlas/v2/orgs/{orgId}/aiModelApiKeys | Return AI Model API Keys for One Organization
[**UpdateGroupModelKey**](AIModelAPIKeysApi.md#UpdateGroupModelKey) | **Patch** /api/atlas/v2/groups/{groupId}/aiModelApiKeys/{apiKeyId} | Update Existing AI Model API Key



## CreateGroupModelKey

> AiModelApiKeyResponse CreateGroupModelKey(ctx, groupId, aiModelApiKeyCreateRequest AiModelApiKeyCreateRequest).Execute()

Create New AI Model API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    aiModelApiKeyCreateRequest := *admin.NewAiModelApiKeyCreateRequest("ANY", "ANY", "Name_example") // AiModelApiKeyCreateRequest | 

    resp, r, err := sdk.AIModelAPIKeysApi.CreateGroupModelKey(context.Background(), groupId, &aiModelApiKeyCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelAPIKeysApi.CreateGroupModelKey`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupModelKey`: AiModelApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelAPIKeysApi.CreateGroupModelKey`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupModelKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aiModelApiKeyCreateRequest** | [**AiModelApiKeyCreateRequest**](AiModelApiKeyCreateRequest.md) | A request containing the name, cloud, and geography of the new API key. | 

### Return type

[**AiModelApiKeyResponse**](AiModelApiKeyResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupModelKey

> DeleteGroupModelKey(ctx, groupId, apiKeyId).Execute()

Delete Existing AI Model API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    apiKeyId := "apiKeyId_example" // string | 

    r, err := sdk.AIModelAPIKeysApi.DeleteGroupModelKey(context.Background(), groupId, apiKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelAPIKeysApi.DeleteGroupModelKey`: %v (%v)\n", err, r)
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
**apiKeyId** | **string** | The id of the API key to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupModelKeyRequest struct via the builder pattern


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


## GetGroupModelKey

> AiModelApiKeyResponse GetGroupModelKey(ctx, groupId, apiKeyId).Execute()

Return Single AI Model API Key for One Group


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    apiKeyId := "apiKeyId_example" // string | 

    resp, r, err := sdk.AIModelAPIKeysApi.GetGroupModelKey(context.Background(), groupId, apiKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelAPIKeysApi.GetGroupModelKey`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupModelKey`: AiModelApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelAPIKeysApi.GetGroupModelKey`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**apiKeyId** | **string** | The id of the API key to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupModelKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AiModelApiKeyResponse**](AiModelApiKeyResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgModelKey

> AiModelApiKeyResponse GetOrgModelKey(ctx, orgId, apiKeyId).Execute()

Return Single AI Model API Key for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    apiKeyId := "apiKeyId_example" // string | 

    resp, r, err := sdk.AIModelAPIKeysApi.GetOrgModelKey(context.Background(), orgId, apiKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelAPIKeysApi.GetOrgModelKey`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgModelKey`: AiModelApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelAPIKeysApi.GetOrgModelKey`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**apiKeyId** | **string** | The id of the API key to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgModelKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AiModelApiKeyResponse**](AiModelApiKeyResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupModelKeys

> PaginatedAtlasAiModelApiKeysResponse ListGroupModelKeys(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return AI Model API Keys for One Group


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    resp, r, err := sdk.AIModelAPIKeysApi.ListGroupModelKeys(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelAPIKeysApi.ListGroupModelKeys`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupModelKeys`: PaginatedAtlasAiModelApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelAPIKeysApi.ListGroupModelKeys`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupModelKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAtlasAiModelApiKeysResponse**](PaginatedAtlasAiModelApiKeysResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgModelKeys

> PaginatedAtlasAiModelApiKeysResponse ListOrgModelKeys(ctx, orgId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return AI Model API Keys for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    resp, r, err := sdk.AIModelAPIKeysApi.ListOrgModelKeys(context.Background(), orgId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelAPIKeysApi.ListOrgModelKeys`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgModelKeys`: PaginatedAtlasAiModelApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelAPIKeysApi.ListOrgModelKeys`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgModelKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAtlasAiModelApiKeysResponse**](PaginatedAtlasAiModelApiKeysResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupModelKey

> AiModelApiKeyResponse UpdateGroupModelKey(ctx, groupId, apiKeyId, aiModelApiKeyUpdateRequest AiModelApiKeyUpdateRequest).Execute()

Update Existing AI Model API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    apiKeyId := "apiKeyId_example" // string | 
    aiModelApiKeyUpdateRequest := *admin.NewAiModelApiKeyUpdateRequest("Name_example") // AiModelApiKeyUpdateRequest | 

    resp, r, err := sdk.AIModelAPIKeysApi.UpdateGroupModelKey(context.Background(), groupId, apiKeyId, &aiModelApiKeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelAPIKeysApi.UpdateGroupModelKey`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupModelKey`: AiModelApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIModelAPIKeysApi.UpdateGroupModelKey`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**apiKeyId** | **string** | The id of the API key to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupModelKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aiModelApiKeyUpdateRequest** | [**AiModelApiKeyUpdateRequest**](AiModelApiKeyUpdateRequest.md) | A request containing the new name for the API key. | 

### Return type

[**AiModelApiKeyResponse**](AiModelApiKeyResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

