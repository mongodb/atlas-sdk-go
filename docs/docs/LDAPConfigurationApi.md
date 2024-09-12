# \LDAPConfigurationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLDAPConfiguration**](LDAPConfigurationApi.md#DeleteLDAPConfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping | Remove the Current LDAP User to DN Mapping
[**GetLDAPConfiguration**](LDAPConfigurationApi.md#GetLDAPConfiguration) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity | Return the Current LDAP or X.509 Configuration
[**GetLDAPConfigurationStatus**](LDAPConfigurationApi.md#GetLDAPConfigurationStatus) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return the Status of One Verify LDAP Configuration Request
[**SaveLDAPConfiguration**](LDAPConfigurationApi.md#SaveLDAPConfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/userSecurity | Edit the LDAP or X.509 Configuration
[**VerifyLDAPConfiguration**](LDAPConfigurationApi.md#VerifyLDAPConfiguration) | **Post** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify | Verify the LDAP Configuration in One Project



## DeleteLDAPConfiguration

> UserSecurity DeleteLDAPConfiguration(ctx, groupId).Execute()

Remove the Current LDAP User to DN Mapping


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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

    resp, r, err := sdk.LDAPConfigurationApi.DeleteLDAPConfiguration(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.DeleteLDAPConfiguration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteLDAPConfiguration`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.DeleteLDAPConfiguration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLDAPConfigurationRequest struct via the builder pattern


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


## GetLDAPConfiguration

> UserSecurity GetLDAPConfiguration(ctx, groupId).Execute()

Return the Current LDAP or X.509 Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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

    resp, r, err := sdk.LDAPConfigurationApi.GetLDAPConfiguration(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.GetLDAPConfiguration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetLDAPConfiguration`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.GetLDAPConfiguration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLDAPConfigurationRequest struct via the builder pattern


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


## GetLDAPConfigurationStatus

> LDAPVerifyConnectivityJobRequest GetLDAPConfigurationStatus(ctx, groupId, requestId).Execute()

Return the Status of One Verify LDAP Configuration Request


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    requestId := "requestId_example" // string | 

    resp, r, err := sdk.LDAPConfigurationApi.GetLDAPConfigurationStatus(context.Background(), groupId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.GetLDAPConfigurationStatus`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetLDAPConfigurationStatus`: LDAPVerifyConnectivityJobRequest
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.GetLDAPConfigurationStatus`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**requestId** | **string** | Unique string that identifies the request to verify an &lt;abbr title&#x3D;\&quot;Lightweight Directory Access Protocol\&quot;&gt;LDAP&lt;/abbr&gt; configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLDAPConfigurationStatusRequest struct via the builder pattern


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


## SaveLDAPConfiguration

> UserSecurity SaveLDAPConfiguration(ctx, groupId, userSecurity UserSecurity).Execute()

Edit the LDAP or X.509 Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    userSecurity := *openapiclient.NewUserSecurity() // UserSecurity | 

    resp, r, err := sdk.LDAPConfigurationApi.SaveLDAPConfiguration(context.Background(), groupId, &userSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.SaveLDAPConfiguration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `SaveLDAPConfiguration`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.SaveLDAPConfiguration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveLDAPConfigurationRequest struct via the builder pattern


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


## VerifyLDAPConfiguration

> LDAPVerifyConnectivityJobRequest VerifyLDAPConfiguration(ctx, groupId, lDAPVerifyConnectivityJobRequestParams LDAPVerifyConnectivityJobRequestParams).Execute()

Verify the LDAP Configuration in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    lDAPVerifyConnectivityJobRequestParams := *openapiclient.NewLDAPVerifyConnectivityJobRequestParams("BindPassword_example", "CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com", "Hostname_example", int(123)) // LDAPVerifyConnectivityJobRequestParams | 

    resp, r, err := sdk.LDAPConfigurationApi.VerifyLDAPConfiguration(context.Background(), groupId, &lDAPVerifyConnectivityJobRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.VerifyLDAPConfiguration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `VerifyLDAPConfiguration`: LDAPVerifyConnectivityJobRequest
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.VerifyLDAPConfiguration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyLDAPConfigurationRequest struct via the builder pattern


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

