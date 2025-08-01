# \PrivateEndpointServicesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivateEndpoint**](PrivateEndpointServicesApi.md#CreatePrivateEndpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint | Create One Private Endpoint for One Provider
[**CreatePrivateEndpointService**](PrivateEndpointServicesApi.md#CreatePrivateEndpointService) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/endpointService | Create One Private Endpoint Service for One Provider
[**DeletePrivateEndpoint**](PrivateEndpointServicesApi.md#DeletePrivateEndpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Remove One Private Endpoint for One Provider
[**DeletePrivateEndpointService**](PrivateEndpointServicesApi.md#DeletePrivateEndpointService) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Remove One Private Endpoint Service for One Provider
[**GetPrivateEndpoint**](PrivateEndpointServicesApi.md#GetPrivateEndpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Return One Private Endpoint for One Provider
[**GetPrivateEndpointService**](PrivateEndpointServicesApi.md#GetPrivateEndpointService) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Return One Private Endpoint Service for One Provider
[**GetRegionalizedPrivateEndpointSetting**](PrivateEndpointServicesApi.md#GetRegionalizedPrivateEndpointSetting) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Return Regionalized Private Endpoint Status
[**ListPrivateEndpointServices**](PrivateEndpointServicesApi.md#ListPrivateEndpointServices) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService | Return All Private Endpoint Services for One Provider
[**ToggleRegionalizedPrivateEndpointSetting**](PrivateEndpointServicesApi.md#ToggleRegionalizedPrivateEndpointSetting) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Toggle Regionalized Private Endpoint Status



## CreatePrivateEndpoint

> PrivateLinkEndpoint CreatePrivateEndpoint(ctx, groupId, cloudProvider, endpointServiceId, createEndpointRequest CreateEndpointRequest).Execute()

Create One Private Endpoint for One Provider


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
    cloudProvider := "cloudProvider_example" // string |  (default to "AWS")
    endpointServiceId := "endpointServiceId_example" // string | 
    createEndpointRequest := *openapiclient.NewCreateEndpointRequest() // CreateEndpointRequest | 

    resp, r, err := sdk.PrivateEndpointServicesApi.CreatePrivateEndpoint(context.Background(), groupId, cloudProvider, endpointServiceId, &createEndpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.CreatePrivateEndpoint`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreatePrivateEndpoint`: PrivateLinkEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.CreatePrivateEndpoint`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint. | [default to &quot;AWS&quot;]
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createEndpointRequest** | [**CreateEndpointRequest**](CreateEndpointRequest.md) | Creates one private endpoint for the specified cloud service provider. | 

### Return type

[**PrivateLinkEndpoint**](PrivateLinkEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateEndpointService

> EndpointService CreatePrivateEndpointService(ctx, groupId, cloudProviderEndpointServiceRequest CloudProviderEndpointServiceRequest).Execute()

Create One Private Endpoint Service for One Provider


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
    cloudProviderEndpointServiceRequest := *openapiclient.NewCloudProviderEndpointServiceRequest("ProviderName_example", "Region_example") // CloudProviderEndpointServiceRequest | 

    resp, r, err := sdk.PrivateEndpointServicesApi.CreatePrivateEndpointService(context.Background(), groupId, &cloudProviderEndpointServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.CreatePrivateEndpointService`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreatePrivateEndpointService`: EndpointService
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.CreatePrivateEndpointService`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateEndpointServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudProviderEndpointServiceRequest** | [**CloudProviderEndpointServiceRequest**](CloudProviderEndpointServiceRequest.md) | Creates one private endpoint for the specified cloud service provider. | 

### Return type

[**EndpointService**](EndpointService.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateEndpoint

> DeletePrivateEndpoint(ctx, groupId, cloudProvider, endpointId, endpointServiceId).Execute()

Remove One Private Endpoint for One Provider


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
    cloudProvider := "cloudProvider_example" // string |  (default to "AWS")
    endpointId := "endpointId_example" // string | 
    endpointServiceId := "endpointServiceId_example" // string | 

    r, err := sdk.PrivateEndpointServicesApi.DeletePrivateEndpoint(context.Background(), groupId, cloudProvider, endpointId, endpointServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.DeletePrivateEndpoint`: %v (%v)\n", err, r)
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
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint. | [default to &quot;AWS&quot;]
**endpointId** | **string** | Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints. | 
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateEndpointRequest struct via the builder pattern


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


## DeletePrivateEndpointService

> DeletePrivateEndpointService(ctx, groupId, cloudProvider, endpointServiceId).Execute()

Remove One Private Endpoint Service for One Provider


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
    cloudProvider := "cloudProvider_example" // string |  (default to "AWS")
    endpointServiceId := "endpointServiceId_example" // string | 

    r, err := sdk.PrivateEndpointServicesApi.DeletePrivateEndpointService(context.Background(), groupId, cloudProvider, endpointServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.DeletePrivateEndpointService`: %v (%v)\n", err, r)
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
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint service. | [default to &quot;AWS&quot;]
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateEndpointServiceRequest struct via the builder pattern


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


## GetPrivateEndpoint

> PrivateLinkEndpoint GetPrivateEndpoint(ctx, groupId, cloudProvider, endpointId, endpointServiceId).Execute()

Return One Private Endpoint for One Provider


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
    cloudProvider := "cloudProvider_example" // string |  (default to "AWS")
    endpointId := "endpointId_example" // string | 
    endpointServiceId := "endpointServiceId_example" // string | 

    resp, r, err := sdk.PrivateEndpointServicesApi.GetPrivateEndpoint(context.Background(), groupId, cloudProvider, endpointId, endpointServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.GetPrivateEndpoint`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPrivateEndpoint`: PrivateLinkEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.GetPrivateEndpoint`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint. | [default to &quot;AWS&quot;]
**endpointId** | **string** | Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints. | 
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PrivateLinkEndpoint**](PrivateLinkEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateEndpointService

> EndpointService GetPrivateEndpointService(ctx, groupId, cloudProvider, endpointServiceId).Execute()

Return One Private Endpoint Service for One Provider


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
    cloudProvider := "cloudProvider_example" // string |  (default to "AWS")
    endpointServiceId := "endpointServiceId_example" // string | 

    resp, r, err := sdk.PrivateEndpointServicesApi.GetPrivateEndpointService(context.Background(), groupId, cloudProvider, endpointServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.GetPrivateEndpointService`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPrivateEndpointService`: EndpointService
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.GetPrivateEndpointService`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint service. | [default to &quot;AWS&quot;]
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateEndpointServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EndpointService**](EndpointService.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionalizedPrivateEndpointSetting

> ProjectSettingItem GetRegionalizedPrivateEndpointSetting(ctx, groupId).Execute()

Return Regionalized Private Endpoint Status


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

    resp, r, err := sdk.PrivateEndpointServicesApi.GetRegionalizedPrivateEndpointSetting(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.GetRegionalizedPrivateEndpointSetting`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetRegionalizedPrivateEndpointSetting`: ProjectSettingItem
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.GetRegionalizedPrivateEndpointSetting`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionalizedPrivateEndpointSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectSettingItem**](ProjectSettingItem.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateEndpointServices

> []EndpointService ListPrivateEndpointServices(ctx, groupId, cloudProvider).Execute()

Return All Private Endpoint Services for One Provider


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
    cloudProvider := "cloudProvider_example" // string |  (default to "AWS")

    resp, r, err := sdk.PrivateEndpointServicesApi.ListPrivateEndpointServices(context.Background(), groupId, cloudProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ListPrivateEndpointServices`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPrivateEndpointServices`: []EndpointService
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ListPrivateEndpointServices`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint service. | [default to &quot;AWS&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateEndpointServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]EndpointService**](EndpointService.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleRegionalizedPrivateEndpointSetting

> ProjectSettingItem ToggleRegionalizedPrivateEndpointSetting(ctx, groupId, projectSettingItem ProjectSettingItem).Execute()

Toggle Regionalized Private Endpoint Status


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
    projectSettingItem := *openapiclient.NewProjectSettingItem(false) // ProjectSettingItem | 

    resp, r, err := sdk.PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointSetting(context.Background(), groupId, &projectSettingItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointSetting`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ToggleRegionalizedPrivateEndpointSetting`: ProjectSettingItem
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointSetting`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleRegionalizedPrivateEndpointSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectSettingItem** | [**ProjectSettingItem**](ProjectSettingItem.md) | Enables or disables the ability to create multiple private endpoints per region in all cloud service providers in one project. | 

### Return type

[**ProjectSettingItem**](ProjectSettingItem.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

