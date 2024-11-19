# \RootApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemStatus**](RootApi.md#GetSystemStatus) | **Get** /api/atlas/v2 | Return the status of this MongoDB application
[**ReturnAllControlPlaneIpAddresses**](RootApi.md#ReturnAllControlPlaneIpAddresses) | **Get** /api/atlas/v2/unauth/controlPlaneIPAddresses | Return All Control Plane IP Addresses



## GetSystemStatus

> SystemStatus GetSystemStatus(ctx).Execute()

Return the status of this MongoDB application


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023002/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }


    resp, r, err := sdk.RootApi.GetSystemStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootApi.GetSystemStatus`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetSystemStatus`: SystemStatus
    fmt.Fprintf(os.Stdout, "Response from `RootApi.GetSystemStatus`: %v (%v)\n", resp, r)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemStatusRequest struct via the builder pattern


### Return type

[**SystemStatus**](SystemStatus.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllControlPlaneIpAddresses

> ControlPlaneIPAddresses ReturnAllControlPlaneIpAddresses(ctx).Execute()

Return All Control Plane IP Addresses


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023002/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }


    resp, r, err := sdk.RootApi.ReturnAllControlPlaneIpAddresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootApi.ReturnAllControlPlaneIpAddresses`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ReturnAllControlPlaneIpAddresses`: ControlPlaneIPAddresses
    fmt.Fprintf(os.Stdout, "Response from `RootApi.ReturnAllControlPlaneIpAddresses`: %v (%v)\n", resp, r)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllControlPlaneIpAddressesRequest struct via the builder pattern


### Return type

[**ControlPlaneIPAddresses**](ControlPlaneIPAddresses.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

