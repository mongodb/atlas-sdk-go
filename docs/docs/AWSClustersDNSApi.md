# \AWSClustersDNSApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAwsCustomDns**](AWSClustersDNSApi.md#GetAwsCustomDns) | **Get** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Return One Custom DNS Configuration for Atlas Clusters on AWS
[**ToggleAwsCustomDns**](AWSClustersDNSApi.md#ToggleAwsCustomDns) | **Patch** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS



## GetAwsCustomDns

> AWSCustomDNSEnabled GetAwsCustomDns(ctx, groupId).Execute()

Return One Custom DNS Configuration for Atlas Clusters on AWS


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.AWSClustersDNSApi.GetAwsCustomDns(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSClustersDNSApi.GetAwsCustomDns``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetAwsCustomDns`: AWSCustomDNSEnabled
    fmt.Fprintf(os.Stdout, "Response from `AWSClustersDNSApi.GetAwsCustomDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsCustomDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AWSCustomDNSEnabled**](AWSCustomDNSEnabled.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleAwsCustomDns

> AWSCustomDNSEnabled ToggleAwsCustomDns(ctx, groupId, aWSCustomDNSEnabled AWSCustomDNSEnabled).Execute()

Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    aWSCustomDNSEnabled := *openapiclient.NewAWSCustomDNSEnabled(false) // AWSCustomDNSEnabled | 

    resp, r, err := sdk.AWSClustersDNSApi.ToggleAwsCustomDns(context.Background(), groupId, &aWSCustomDNSEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSClustersDNSApi.ToggleAwsCustomDns``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ToggleAwsCustomDns`: AWSCustomDNSEnabled
    fmt.Fprintf(os.Stdout, "Response from `AWSClustersDNSApi.ToggleAwsCustomDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleAwsCustomDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aWSCustomDNSEnabled** | [**AWSCustomDNSEnabled**](AWSCustomDNSEnabled.md) | Enables or disables the custom DNS configuration for AWS clusters in the specified project. | 

### Return type

[**AWSCustomDNSEnabled**](AWSCustomDNSEnabled.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

