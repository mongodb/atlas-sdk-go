# \ResourcePoliciesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAtlasResourcePolicy**](ResourcePoliciesApi.md#CreateAtlasResourcePolicy) | **Post** /api/atlas/v2/orgs/{orgId}/resourcePolicies | Create one Atlas Resource Policy
[**DeleteAtlasResourcePolicy**](ResourcePoliciesApi.md#DeleteAtlasResourcePolicy) | **Delete** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Delete one Atlas Resource Policy
[**GetAtlasResourcePolicies**](ResourcePoliciesApi.md#GetAtlasResourcePolicies) | **Get** /api/atlas/v2/orgs/{orgId}/resourcePolicies | Return all Atlas Resource Policies
[**GetAtlasResourcePolicy**](ResourcePoliciesApi.md#GetAtlasResourcePolicy) | **Get** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Return one Atlas Resource Policy
[**GetResourcesNonCompliant**](ResourcePoliciesApi.md#GetResourcesNonCompliant) | **Get** /api/atlas/v2/orgs/{orgId}/nonCompliantResources | Return all non-compliant resources
[**UpdateAtlasResourcePolicy**](ResourcePoliciesApi.md#UpdateAtlasResourcePolicy) | **Patch** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Update one Atlas Resource Policy
[**ValidateAtlasResourcePolicy**](ResourcePoliciesApi.md#ValidateAtlasResourcePolicy) | **Post** /api/atlas/v2/orgs/{orgId}/resourcePolicies:validate | Validate one Atlas Resource Policy



## CreateAtlasResourcePolicy

> ApiAtlasResourcePolicy CreateAtlasResourcePolicy(ctx, orgId, apiAtlasResourcePolicyCreate ApiAtlasResourcePolicyCreate).Execute()

Create one Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
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
    apiAtlasResourcePolicyCreate := *openapiclient.NewApiAtlasResourcePolicyCreate("Name_example", []openapiclient.ApiAtlasPolicyCreate{*openapiclient.NewApiAtlasPolicyCreate("  forbid (
    principal,
    action == cloud::Action::"cluster.createEdit",
    resource
  ) when {
   context.cluster.regions.contains(cloud::region::"aws:us-east-1")
  };
")}) // ApiAtlasResourcePolicyCreate | 

    resp, r, err := sdk.ResourcePoliciesApi.CreateAtlasResourcePolicy(context.Background(), orgId, &apiAtlasResourcePolicyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.CreateAtlasResourcePolicy`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateAtlasResourcePolicy`: ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.CreateAtlasResourcePolicy`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAtlasResourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasResourcePolicyCreate** | [**ApiAtlasResourcePolicyCreate**](ApiAtlasResourcePolicyCreate.md) | Atlas Resource Policy to create. | 

### Return type

[**ApiAtlasResourcePolicy**](ApiAtlasResourcePolicy.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAtlasResourcePolicy

> any DeleteAtlasResourcePolicy(ctx, orgId, resourcePolicyId).Execute()

Delete one Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
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
    resourcePolicyId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.ResourcePoliciesApi.DeleteAtlasResourcePolicy(context.Background(), orgId, resourcePolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.DeleteAtlasResourcePolicy`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteAtlasResourcePolicy`: any
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.DeleteAtlasResourcePolicy`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**resourcePolicyId** | **string** | Unique 24-hexadecimal digit string that identifies an atlas resource policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAtlasResourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAtlasResourcePolicies

> []ApiAtlasResourcePolicy GetAtlasResourcePolicies(ctx, orgId).Execute()

Return all Atlas Resource Policies


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
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

    resp, r, err := sdk.ResourcePoliciesApi.GetAtlasResourcePolicies(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.GetAtlasResourcePolicies`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAtlasResourcePolicies`: []ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.GetAtlasResourcePolicies`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasResourcePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApiAtlasResourcePolicy**](ApiAtlasResourcePolicy.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAtlasResourcePolicy

> ApiAtlasResourcePolicy GetAtlasResourcePolicy(ctx, orgId, resourcePolicyId).Execute()

Return one Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
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
    resourcePolicyId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.ResourcePoliciesApi.GetAtlasResourcePolicy(context.Background(), orgId, resourcePolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.GetAtlasResourcePolicy`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAtlasResourcePolicy`: ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.GetAtlasResourcePolicy`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**resourcePolicyId** | **string** | Unique 24-hexadecimal digit string that identifies an atlas resource policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasResourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiAtlasResourcePolicy**](ApiAtlasResourcePolicy.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourcesNonCompliant

> []ApiAtlasNonCompliantResource GetResourcesNonCompliant(ctx, orgId).Execute()

Return all non-compliant resources


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
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

    resp, r, err := sdk.ResourcePoliciesApi.GetResourcesNonCompliant(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.GetResourcesNonCompliant`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetResourcesNonCompliant`: []ApiAtlasNonCompliantResource
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.GetResourcesNonCompliant`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcesNonCompliantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApiAtlasNonCompliantResource**](ApiAtlasNonCompliantResource.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAtlasResourcePolicy

> ApiAtlasResourcePolicy UpdateAtlasResourcePolicy(ctx, orgId, resourcePolicyId, apiAtlasResourcePolicyEdit ApiAtlasResourcePolicyEdit).Execute()

Update one Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
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
    resourcePolicyId := "32b6e34b3d91647abb20e7b8" // string | 
    apiAtlasResourcePolicyEdit := *openapiclient.NewApiAtlasResourcePolicyEdit() // ApiAtlasResourcePolicyEdit | 

    resp, r, err := sdk.ResourcePoliciesApi.UpdateAtlasResourcePolicy(context.Background(), orgId, resourcePolicyId, &apiAtlasResourcePolicyEdit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.UpdateAtlasResourcePolicy`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateAtlasResourcePolicy`: ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.UpdateAtlasResourcePolicy`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**resourcePolicyId** | **string** | Unique 24-hexadecimal digit string that identifies an atlas resource policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAtlasResourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasResourcePolicyEdit** | [**ApiAtlasResourcePolicyEdit**](ApiAtlasResourcePolicyEdit.md) | Atlas Resource Policy to update. | 

### Return type

[**ApiAtlasResourcePolicy**](ApiAtlasResourcePolicy.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAtlasResourcePolicy

> ApiAtlasResourcePolicy ValidateAtlasResourcePolicy(ctx, orgId, apiAtlasResourcePolicyCreate ApiAtlasResourcePolicyCreate).Execute()

Validate one Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
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
    apiAtlasResourcePolicyCreate := *openapiclient.NewApiAtlasResourcePolicyCreate("Name_example", []openapiclient.ApiAtlasPolicyCreate{*openapiclient.NewApiAtlasPolicyCreate("  forbid (
    principal,
    action == cloud::Action::"cluster.createEdit",
    resource
  ) when {
   context.cluster.regions.contains(cloud::region::"aws:us-east-1")
  };
")}) // ApiAtlasResourcePolicyCreate | 

    resp, r, err := sdk.ResourcePoliciesApi.ValidateAtlasResourcePolicy(context.Background(), orgId, &apiAtlasResourcePolicyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.ValidateAtlasResourcePolicy`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ValidateAtlasResourcePolicy`: ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.ValidateAtlasResourcePolicy`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAtlasResourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasResourcePolicyCreate** | [**ApiAtlasResourcePolicyCreate**](ApiAtlasResourcePolicyCreate.md) | Atlas Resource Policy to create. | 

### Return type

[**ApiAtlasResourcePolicy**](ApiAtlasResourcePolicy.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

