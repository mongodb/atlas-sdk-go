# \DefaultApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnAllControlPlaneIPAddresses**](DefaultApi.md#ReturnAllControlPlaneIPAddresses) | **Get** /api/atlas/v2/unauth/controlPlaneIPAddresses | Return All Control Plane IP Addresses



## ReturnAllControlPlaneIPAddresses

> ControlPlaneIPAddresses ReturnAllControlPlaneIPAddresses(ctx).Execute()

Return All Control Plane IP Addresses


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

    "go.mongodb.org/atlas-sdk/v20231115004/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))


    resp, r, err := sdk.DefaultApi.ReturnAllControlPlaneIPAddresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReturnAllControlPlaneIPAddresses``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ReturnAllControlPlaneIPAddresses`: ControlPlaneIPAddresses
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReturnAllControlPlaneIPAddresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllControlPlaneIPAddressesRequest struct via the builder pattern


### Return type

[**ControlPlaneIPAddresses**](ControlPlaneIPAddresses.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

