# \EncryptionAtRestUsingCustomerKeyManagementApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncryptionAtRest**](EncryptionAtRestUsingCustomerKeyManagementApi.md#GetEncryptionAtRest) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project
[**UpdateEncryptionAtRest**](EncryptionAtRestUsingCustomerKeyManagementApi.md#UpdateEncryptionAtRest) | **Patch** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project



## GetEncryptionAtRest

> EncryptionAtRest GetEncryptionAtRest(ctx, groupId).Execute()

Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project



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
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetEncryptionAtRest`: EncryptionAtRest
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionAtRestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EncryptionAtRest**](EncryptionAtRest.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEncryptionAtRest

> EncryptionAtRest UpdateEncryptionAtRest(ctx, groupId).EncryptionAtRest(encryptionAtRest).Execute()

Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project



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
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    encryptionAtRest := *openapiclient.NewEncryptionAtRest() // EncryptionAtRest | 

    resp, r, err := sdk.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest(context.Background(), groupId).EncryptionAtRest(encryptionAtRest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateEncryptionAtRest`: EncryptionAtRest
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEncryptionAtRestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **encryptionAtRest** | [**EncryptionAtRest**](EncryptionAtRest.md) | Required parameters depend on whether someone has enabled Encryption at Rest using Customer Key Management:  If you have enabled Encryption at Rest using Customer Key Management (CMK), Atlas requires all of the parameters for the desired encryption provider.  - To use AWS Key Management Service (KMS), MongoDB Cloud requires all the fields in the **awsKms** object. - To use Azure Key Vault, MongoDB Cloud requires all the fields in the **azureKeyVault** object. - To use Google Cloud Key Management Service (KMS), MongoDB Cloud requires all the fields in the **googleCloudKms** object.  If you enabled Encryption at Rest using Customer Key  Management, administrators can pass only the changed fields for the **awsKms**, **azureKeyVault**, or **googleCloudKms** object to update the configuration to this endpoint. | 

### Return type

[**EncryptionAtRest**](EncryptionAtRest.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

