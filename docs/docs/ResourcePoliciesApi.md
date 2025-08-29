# \ResourcePoliciesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgResourcePolicy**](ResourcePoliciesApi.md#CreateOrgResourcePolicy) | **Post** /api/atlas/v2/orgs/{orgId}/resourcePolicies | Create One Atlas Resource Policy
[**DeleteOrgResourcePolicy**](ResourcePoliciesApi.md#DeleteOrgResourcePolicy) | **Delete** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Delete One Atlas Resource Policy
[**GetNonCompliantResources**](ResourcePoliciesApi.md#GetNonCompliantResources) | **Get** /api/atlas/v2/orgs/{orgId}/nonCompliantResources | Return All Non-Compliant Resources
[**GetOrgResourcePolicy**](ResourcePoliciesApi.md#GetOrgResourcePolicy) | **Get** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Return One Atlas Resource Policy
[**ListOrgResourcePolicies**](ResourcePoliciesApi.md#ListOrgResourcePolicies) | **Get** /api/atlas/v2/orgs/{orgId}/resourcePolicies | Return All Atlas Resource Policies
[**UpdateOrgResourcePolicy**](ResourcePoliciesApi.md#UpdateOrgResourcePolicy) | **Patch** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Update One Atlas Resource Policy
[**ValidateResourcePolicies**](ResourcePoliciesApi.md#ValidateResourcePolicies) | **Post** /api/atlas/v2/orgs/{orgId}/resourcePolicies:validate | Validate One Atlas Resource Policy



## CreateOrgResourcePolicy

> ApiAtlasResourcePolicy CreateOrgResourcePolicy(ctx, orgId, apiAtlasResourcePolicyCreate ApiAtlasResourcePolicyCreate).Execute()

Create One Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.ResourcePoliciesApi.CreateOrgResourcePolicy(context.Background(), orgId, &apiAtlasResourcePolicyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.CreateOrgResourcePolicy`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgResourcePolicy`: ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.CreateOrgResourcePolicy`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgResourcePolicyRequest struct via the builder pattern


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


## DeleteOrgResourcePolicy

> DeleteOrgResourcePolicy(ctx, orgId, resourcePolicyId).Execute()

Delete One Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    r, err := sdk.ResourcePoliciesApi.DeleteOrgResourcePolicy(context.Background(), orgId, resourcePolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.DeleteOrgResourcePolicy`: %v (%v)\n", err, r)
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
**resourcePolicyId** | **string** | Unique 24-hexadecimal digit string that identifies an atlas resource policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgResourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonCompliantResources

> []ApiAtlasNonCompliantResource GetNonCompliantResources(ctx, orgId).Execute()

Return All Non-Compliant Resources


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.ResourcePoliciesApi.GetNonCompliantResources(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.GetNonCompliantResources`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetNonCompliantResources`: []ApiAtlasNonCompliantResource
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.GetNonCompliantResources`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonCompliantResourcesRequest struct via the builder pattern


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


## GetOrgResourcePolicy

> ApiAtlasResourcePolicy GetOrgResourcePolicy(ctx, orgId, resourcePolicyId).Execute()

Return One Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.ResourcePoliciesApi.GetOrgResourcePolicy(context.Background(), orgId, resourcePolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.GetOrgResourcePolicy`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgResourcePolicy`: ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.GetOrgResourcePolicy`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**resourcePolicyId** | **string** | Unique 24-hexadecimal digit string that identifies an atlas resource policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgResourcePolicyRequest struct via the builder pattern


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


## ListOrgResourcePolicies

> []ApiAtlasResourcePolicy ListOrgResourcePolicies(ctx, orgId).Execute()

Return All Atlas Resource Policies


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.ResourcePoliciesApi.ListOrgResourcePolicies(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.ListOrgResourcePolicies`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgResourcePolicies`: []ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.ListOrgResourcePolicies`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgResourcePoliciesRequest struct via the builder pattern


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


## UpdateOrgResourcePolicy

> ApiAtlasResourcePolicy UpdateOrgResourcePolicy(ctx, orgId, resourcePolicyId, apiAtlasResourcePolicyEdit ApiAtlasResourcePolicyEdit).Execute()

Update One Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.ResourcePoliciesApi.UpdateOrgResourcePolicy(context.Background(), orgId, resourcePolicyId, &apiAtlasResourcePolicyEdit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.UpdateOrgResourcePolicy`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateOrgResourcePolicy`: ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.UpdateOrgResourcePolicy`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**resourcePolicyId** | **string** | Unique 24-hexadecimal digit string that identifies an atlas resource policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgResourcePolicyRequest struct via the builder pattern


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


## ValidateResourcePolicies

> ApiAtlasResourcePolicy ValidateResourcePolicies(ctx, orgId, apiAtlasResourcePolicyCreate ApiAtlasResourcePolicyCreate).Execute()

Validate One Atlas Resource Policy


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.ResourcePoliciesApi.ValidateResourcePolicies(context.Background(), orgId, &apiAtlasResourcePolicyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoliciesApi.ValidateResourcePolicies`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ValidateResourcePolicies`: ApiAtlasResourcePolicy
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoliciesApi.ValidateResourcePolicies`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateResourcePoliciesRequest struct via the builder pattern


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

