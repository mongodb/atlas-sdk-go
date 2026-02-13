# \LDAPConfigurationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLdapUserMapping**](LDAPConfigurationApi.md#DeleteLdapUserMapping) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping | Remove LDAP User to DN Mapping
[**GetUserSecurity**](LDAPConfigurationApi.md#GetUserSecurity) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity | Return LDAP or X.509 Configuration
[**GetUserSecurityVerify**](LDAPConfigurationApi.md#GetUserSecurityVerify) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return Status of LDAP Configuration Verification in One Project
[**UpdateUserSecurity**](LDAPConfigurationApi.md#UpdateUserSecurity) | **Patch** /api/atlas/v2/groups/{groupId}/userSecurity | Update LDAP or X.509 Configuration
[**VerifyUserSecurityLdap**](LDAPConfigurationApi.md#VerifyUserSecurityLdap) | **Post** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify | Verify LDAP Configuration in One Project



## DeleteLdapUserMapping

> UserSecurity DeleteLdapUserMapping(ctx, groupId).Execute()

Remove LDAP User to DN Mapping


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312014/admin"
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

    resp, r, err := sdk.LDAPConfigurationApi.DeleteLdapUserMapping(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.DeleteLdapUserMapping`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteLdapUserMapping`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.DeleteLdapUserMapping`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLdapUserMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSecurity**](UserSecurity.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSecurity

> UserSecurity GetUserSecurity(ctx, groupId).Execute()

Return LDAP or X.509 Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312014/admin"
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

    resp, r, err := sdk.LDAPConfigurationApi.GetUserSecurity(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.GetUserSecurity`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetUserSecurity`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.GetUserSecurity`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSecurity**](UserSecurity.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSecurityVerify

> LDAPVerifyConnectivityJobRequest GetUserSecurityVerify(ctx, groupId, requestId).Execute()

Return Status of LDAP Configuration Verification in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312014/admin"
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

    resp, r, err := sdk.LDAPConfigurationApi.GetUserSecurityVerify(context.Background(), groupId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.GetUserSecurityVerify`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetUserSecurityVerify`: LDAPVerifyConnectivityJobRequest
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.GetUserSecurityVerify`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**requestId** | **string** | Unique string that identifies the request to verify an Lightweight Directory Access Protocol (LDAP) configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSecurityVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LDAPVerifyConnectivityJobRequest**](LDAPVerifyConnectivityJobRequest.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSecurity

> UserSecurity UpdateUserSecurity(ctx, groupId, userSecurity UserSecurity).Execute()

Update LDAP or X.509 Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312014/admin"
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
    userSecurity := *admin.NewUserSecurity() // UserSecurity | 

    resp, r, err := sdk.LDAPConfigurationApi.UpdateUserSecurity(context.Background(), groupId, &userSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.UpdateUserSecurity`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateUserSecurity`: UserSecurity
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.UpdateUserSecurity`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSecurity** | [**UserSecurity**](UserSecurity.md) | Updates the LDAP configuration for the specified project. | 

### Return type

[**UserSecurity**](UserSecurity.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyUserSecurityLdap

> LDAPVerifyConnectivityJobRequest VerifyUserSecurityLdap(ctx, groupId, lDAPVerifyConnectivityJobRequestParams LDAPVerifyConnectivityJobRequestParams).Execute()

Verify LDAP Configuration in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312014/admin"
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
    lDAPVerifyConnectivityJobRequestParams := *admin.NewLDAPVerifyConnectivityJobRequestParams("BindPassword_example", "CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com", "Hostname_example", int(123)) // LDAPVerifyConnectivityJobRequestParams | 

    resp, r, err := sdk.LDAPConfigurationApi.VerifyUserSecurityLdap(context.Background(), groupId, &lDAPVerifyConnectivityJobRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.VerifyUserSecurityLdap`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `VerifyUserSecurityLdap`: LDAPVerifyConnectivityJobRequest
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.VerifyUserSecurityLdap`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserSecurityLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPVerifyConnectivityJobRequestParams** | [**LDAPVerifyConnectivityJobRequestParams**](LDAPVerifyConnectivityJobRequestParams.md) | The LDAP configuration for the specified project that you want to verify. | 

### Return type

[**LDAPVerifyConnectivityJobRequest**](LDAPVerifyConnectivityJobRequest.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

