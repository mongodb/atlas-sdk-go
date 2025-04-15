# \LDAPConfigurationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLdapConfiguration**](LDAPConfigurationApi.md#DeleteLdapConfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping | Remove the Current LDAP User to DN Mapping
[**GetLdapConfiguration**](LDAPConfigurationApi.md#GetLdapConfiguration) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity | Return the Current LDAP or X.509 Configuration
[**GetLdapConfigurationStatus**](LDAPConfigurationApi.md#GetLdapConfigurationStatus) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return the Status of One Verify LDAP Configuration Request
[**SaveLdapConfiguration**](LDAPConfigurationApi.md#SaveLdapConfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/userSecurity | Edit the LDAP or X.509 Configuration
[**VerifyLdapConfiguration**](LDAPConfigurationApi.md#VerifyLdapConfiguration) | **Post** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify | Verify the LDAP Configuration in One Project



## DeleteLdapConfiguration

> UserSecurity DeleteLdapConfiguration(ctx, groupId).Execute()

Remove the Current LDAP User to DN Mapping


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

    resp, r, err := sdk.LDAPConfigurationApi.DeleteLdapConfiguration(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.DeleteLdapConfiguration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteLdapConfiguration`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.DeleteLdapConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLdapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSecurity**](UserSecurity.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapConfiguration

> UserSecurity GetLdapConfiguration(ctx, groupId).Execute()

Return the Current LDAP or X.509 Configuration


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

    resp, r, err := sdk.LDAPConfigurationApi.GetLdapConfiguration(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.GetLdapConfiguration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetLdapConfiguration`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.GetLdapConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSecurity**](UserSecurity.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapConfigurationStatus

> LDAPVerifyConnectivityJobRequest GetLdapConfigurationStatus(ctx, groupId, requestId).Execute()

Return the Status of One Verify LDAP Configuration Request


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
    requestId := "requestId_example" // string | 

    resp, r, err := sdk.LDAPConfigurationApi.GetLdapConfigurationStatus(context.Background(), groupId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.GetLdapConfigurationStatus``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetLdapConfigurationStatus`: LDAPVerifyConnectivityJobRequest
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.GetLdapConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**requestId** | **string** | Unique string that identifies the request to verify an Lightweight Directory Access Protocol (LDAP) configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LDAPVerifyConnectivityJobRequest**](LDAPVerifyConnectivityJobRequest.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveLdapConfiguration

> UserSecurity SaveLdapConfiguration(ctx, groupId, userSecurity UserSecurity).Execute()

Edit the LDAP or X.509 Configuration


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
    userSecurity := *openapiclient.NewUserSecurity() // UserSecurity | 

    resp, r, err := sdk.LDAPConfigurationApi.SaveLdapConfiguration(context.Background(), groupId, &userSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.SaveLdapConfiguration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `SaveLdapConfiguration`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.SaveLdapConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveLdapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSecurity** | [**UserSecurity**](UserSecurity.md) | Updates the LDAP configuration for the specified project. | 

### Return type

[**UserSecurity**](UserSecurity.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyLdapConfiguration

> LDAPVerifyConnectivityJobRequest VerifyLdapConfiguration(ctx, groupId, lDAPVerifyConnectivityJobRequestParams LDAPVerifyConnectivityJobRequestParams).Execute()

Verify the LDAP Configuration in One Project


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
    lDAPVerifyConnectivityJobRequestParams := *openapiclient.NewLDAPVerifyConnectivityJobRequestParams("BindPassword_example", "CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com", "Hostname_example", int(123)) // LDAPVerifyConnectivityJobRequestParams | 

    resp, r, err := sdk.LDAPConfigurationApi.VerifyLdapConfiguration(context.Background(), groupId, &lDAPVerifyConnectivityJobRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.VerifyLdapConfiguration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `VerifyLdapConfiguration`: LDAPVerifyConnectivityJobRequest
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.VerifyLdapConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyLdapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPVerifyConnectivityJobRequestParams** | [**LDAPVerifyConnectivityJobRequestParams**](LDAPVerifyConnectivityJobRequestParams.md) | The LDAP configuration for the specified project that you want to verify. | 

### Return type

[**LDAPVerifyConnectivityJobRequest**](LDAPVerifyConnectivityJobRequest.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

