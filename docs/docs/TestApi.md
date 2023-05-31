# \TestApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionedExample**](TestApi.md#VersionedExample) | **Get** /api/atlas/v2/example/info | Example resource info for versioning of the Atlas API



## VersionedExample

> ExampleResourceResponseView20230201 VersionedExample(ctx).AdditionalInfo(additionalInfo).Execute()

Example resource info for versioning of the Atlas API



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

    additionalInfo := true // bool |  (optional) (default to false)

    resp, r, err := sdk.TestApi.VersionedExample(context.Background()).AdditionalInfo(additionalInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestApi.VersionedExample``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `VersionedExample`: ExampleResourceResponseView20230201
    fmt.Fprintf(os.Stdout, "Response from `TestApi.VersionedExample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVersionedExampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **additionalInfo** | **bool** |  | [default to false]

### Return type

[**ExampleResourceResponseView20230201**](ExampleResourceResponseView20230201.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

