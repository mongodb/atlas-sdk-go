# \RateLimitingApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRateLimit**](RateLimitingApi.md#GetRateLimit) | **Get** /api/atlas/v2/rateLimits/{endpointSetId} | Return One Rate Limit
[**ListRateLimits**](RateLimitingApi.md#ListRateLimits) | **Get** /api/atlas/v2/rateLimits | Return All Rate Limits



## GetRateLimit

> RateLimitEndpointSetResponse GetRateLimit(ctx, endpointSetId).GroupId(groupId).OrgId(orgId).UserId(userId).IpAddress(ipAddress).Execute()

Return One Rate Limit


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312016/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    endpointSetId := "rateLimitsInspection_group" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string |  (optional)
    orgId := "32b6e34b3d91647abb20e7b8" // string |  (optional)
    userId := "mdb_sa_id_1234567890abcdef12345678" // string |  (optional)
    ipAddress := "127.0.0.1" // string |  (optional)

    resp, r, err := sdk.RateLimitingApi.GetRateLimit(context.Background(), endpointSetId).GroupId(groupId).OrgId(orgId).UserId(userId).IpAddress(ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.GetRateLimit`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetRateLimit`: RateLimitEndpointSetResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.GetRateLimit`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointSetId** | **string** | The ID of the rate limit endpoint set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Project to request rate limits for. When this parameter is provided, the limits returned are applicable to the specified project. The requesting user must have the Project Read Only role for the specified project. | 
 **orgId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Organization to request rate limits for. When this parameter is provided, the limits returned are applicable to the specified organization. The requesting user must have the Organization Read Only role for the specified organization. | 
 **userId** | **string** | A string that identifies the Atlas user to request rate limits for. The ID can for example be the Service Account Client ID or the API Public Key. When this parameter is provided, the limits returned are applicable to the specified  user. The requesting user must be the same as the specified user. | 
 **ipAddress** | **string** | An IP address to request rate limits for. When this parameter is provided, the limits returned are applicable to the specified IP address. The requesting user must have the same IP address as the one provided in the request. | 

### Return type

[**RateLimitEndpointSetResponse**](RateLimitEndpointSetResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRateLimits

> PaginatedRateLimitEndpointSets ListRateLimits(ctx).ItemsPerPage(itemsPerPage).PageNum(pageNum).GroupId(groupId).OrgId(orgId).UserId(userId).IpAddress(ipAddress).Name(name).EndpointPath(endpointPath).Execute()

Return All Rate Limits


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312016/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    groupId := "32b6e34b3d91647abb20e7b8" // string |  (optional)
    orgId := "32b6e34b3d91647abb20e7b8" // string |  (optional)
    userId := "mdb_sa_id_1234567890abcdef12345678" // string |  (optional)
    ipAddress := "127.0.0.1" // string |  (optional)
    name := "Rate Limits Inspection" // string |  (optional)
    endpointPath := "/api/atlas/v2/clusters" // string |  (optional)

    resp, r, err := sdk.RateLimitingApi.ListRateLimits(context.Background()).ItemsPerPage(itemsPerPage).PageNum(pageNum).GroupId(groupId).OrgId(orgId).UserId(userId).IpAddress(ipAddress).Name(name).EndpointPath(endpointPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.ListRateLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListRateLimits`: PaginatedRateLimitEndpointSets
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.ListRateLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **groupId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Project to request rate limits for. When this parameter is provided, only group scoped endpoint sets are returned and the limits returned are applicable to the specified project. The requesting user must have the Project Read Only role for the specified project. | 
 **orgId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Organization to request rate limits for. When this parameter is provided, only organization scoped endpoint sets are returned and the limits returned are applicable to the specified organization. The requesting user must have the Organization Read Only role for the specified organization. | 
 **userId** | **string** | A string that identifies the Atlas user to request rate limits for. The ID can for example be the Service Account Client ID or the API Public Key. When this parameter is provided, only user scoped endpoint sets are returned and the limits returned are applicable to the specified user. The requesting user must be the same as the specified user. | 
 **ipAddress** | **string** | An IP address to request rate limits for. When this parameter is provided, only IP scoped endpoint sets are returned and the limits returned are applicable to the specified IP address. The requesting user must have the same IP address as the one provided in the request. | 
 **name** | **string** | Filters the returned endpoint sets by the provided endpoint set name. Multiple names may be provided, for example &#x60;/rateLimits?name&#x3D;Name1&amp;name&#x3D;Name2&#x60;. For names that use spaces, replace the space with its URL-encoded value (&#x60;%20&#x60;). | 
 **endpointPath** | **string** | Filters the returned endpoint sets by the provided endpoint path. Multiple paths may be provided, for example &#x60;/rateLimits?endpointPath&#x3D;%2Fapi%2Fatlas%2Fv2%2Fclusters&amp;endpointPath&#x3D;%2Fapi%2Fatlas%2Fv2%2Fgroups%2F%7BgroupId%7D%2F&#x60;. Replace &#x60;/&#x60;, &#x60;{&#x60; and &#x60;}&#x60; with their URL-encoded values (&#x60;%2F&#x60;, &#x60;%7B&#x60; and &#x60;%7D&#x60; respectively). | 

### Return type

[**PaginatedRateLimitEndpointSets**](PaginatedRateLimitEndpointSets.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

