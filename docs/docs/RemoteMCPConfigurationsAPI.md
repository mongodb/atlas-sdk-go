# \RemoteMCPConfigurationsAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupMcpConfig**](RemoteMCPConfigurationsAPI.md#CreateGroupMcpConfig) | **Post** /api/atlas/v2/groups/{groupId}/mcpConfigs | Create One MCP Configuration for One Project
[**CreateGroupMcpSecret**](RemoteMCPConfigurationsAPI.md#CreateGroupMcpSecret) | **Post** /api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}/secrets | Create One Secret for One Project MCP Configuration
[**CreateOrgMcpConfig**](RemoteMCPConfigurationsAPI.md#CreateOrgMcpConfig) | **Post** /api/atlas/v2/orgs/{orgId}/mcpConfigs | Create One MCP Configuration for One Organization
[**CreateOrgMcpSecret**](RemoteMCPConfigurationsAPI.md#CreateOrgMcpSecret) | **Post** /api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}/secrets | Create One Secret for One Organization MCP Configuration
[**DeleteGroupMcpConfig**](RemoteMCPConfigurationsAPI.md#DeleteGroupMcpConfig) | **Delete** /api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId} | Delete One MCP Configuration for One Project
[**DeleteGroupMcpSecret**](RemoteMCPConfigurationsAPI.md#DeleteGroupMcpSecret) | **Delete** /api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}/secrets/{secretId} | Delete One Secret for One Project MCP Configuration
[**DeleteOrgMcpConfig**](RemoteMCPConfigurationsAPI.md#DeleteOrgMcpConfig) | **Delete** /api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId} | Delete One MCP Configuration for One Organization
[**DeleteOrgMcpSecret**](RemoteMCPConfigurationsAPI.md#DeleteOrgMcpSecret) | **Delete** /api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}/secrets/{secretId} | Delete One Secret for One Organization MCP Configuration
[**GetGroupMcpConfig**](RemoteMCPConfigurationsAPI.md#GetGroupMcpConfig) | **Get** /api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId} | Return One MCP Configuration for One Project
[**GetGroupMcpSecret**](RemoteMCPConfigurationsAPI.md#GetGroupMcpSecret) | **Get** /api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}/secrets/{secretId} | Return One Secret for One Project MCP Configuration
[**GetOrgMcpConfig**](RemoteMCPConfigurationsAPI.md#GetOrgMcpConfig) | **Get** /api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId} | Return One MCP Configuration for One Organization
[**GetOrgMcpSecret**](RemoteMCPConfigurationsAPI.md#GetOrgMcpSecret) | **Get** /api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}/secrets/{secretId} | Return One Secret for One Organization MCP Configuration
[**ListGroupMcpConfigs**](RemoteMCPConfigurationsAPI.md#ListGroupMcpConfigs) | **Get** /api/atlas/v2/groups/{groupId}/mcpConfigs | Return All MCP Configurations for One Project
[**ListGroupMcpSecrets**](RemoteMCPConfigurationsAPI.md#ListGroupMcpSecrets) | **Get** /api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}/secrets | Return All Secrets for One Project MCP Configuration
[**ListOrgMcpConfigs**](RemoteMCPConfigurationsAPI.md#ListOrgMcpConfigs) | **Get** /api/atlas/v2/orgs/{orgId}/mcpConfigs | Return All MCP Configurations for One Organization
[**ListOrgMcpSecrets**](RemoteMCPConfigurationsAPI.md#ListOrgMcpSecrets) | **Get** /api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}/secrets | Return All Secrets for One Organization MCP Configuration
[**UpdateGroupMcpConfig**](RemoteMCPConfigurationsAPI.md#UpdateGroupMcpConfig) | **Patch** /api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId} | Update One MCP Configuration for One Project
[**UpdateOrgMcpConfig**](RemoteMCPConfigurationsAPI.md#UpdateOrgMcpConfig) | **Patch** /api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId} | Update One MCP Configuration for One Organization



## CreateGroupMcpConfig

> GroupMcpConfigResponse CreateGroupMcpConfig(ctx, groupId, createGroupMcpConfigRequest CreateGroupMcpConfigRequest).Execute()

Create One MCP Configuration for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    createGroupMcpConfigRequest := *admin.NewCreateGroupMcpConfigRequest("McpConfigName_example", []string{"Roles_example"}) // CreateGroupMcpConfigRequest | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.CreateGroupMcpConfig(context.Background(), groupId, &createGroupMcpConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.CreateGroupMcpConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupMcpConfig`: GroupMcpConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.CreateGroupMcpConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupMcpConfigRequest** | [**CreateGroupMcpConfigRequest**](CreateGroupMcpConfigRequest.md) | MCP configuration to create. | 

### Return type

[**GroupMcpConfigResponse**](GroupMcpConfigResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupMcpSecret

> ServiceAccountSecret CreateGroupMcpSecret(ctx, groupId, mcpConfigId, serviceAccountSecretRequest ServiceAccountSecretRequest).Execute()

Create One Secret for One Project MCP Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "32b6e34b3d91647adeabc012" // string | 
    serviceAccountSecretRequest := *admin.NewServiceAccountSecretRequest(int(8)) // ServiceAccountSecretRequest | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.CreateGroupMcpSecret(context.Background(), groupId, mcpConfigId, &serviceAccountSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.CreateGroupMcpSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupMcpSecret`: ServiceAccountSecret
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.CreateGroupMcpSecret`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupMcpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAccountSecretRequest** | [**ServiceAccountSecretRequest**](ServiceAccountSecretRequest.md) | Secret creation request. | 

### Return type

[**ServiceAccountSecret**](ServiceAccountSecret.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgMcpConfig

> OrgMcpConfigResponse CreateOrgMcpConfig(ctx, orgId, createOrgMcpConfigRequest CreateOrgMcpConfigRequest).Execute()

Create One MCP Configuration for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    createOrgMcpConfigRequest := *admin.NewCreateOrgMcpConfigRequest("McpConfigName_example", []string{"Roles_example"}) // CreateOrgMcpConfigRequest | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.CreateOrgMcpConfig(context.Background(), orgId, &createOrgMcpConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.CreateOrgMcpConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgMcpConfig`: OrgMcpConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.CreateOrgMcpConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrgMcpConfigRequest** | [**CreateOrgMcpConfigRequest**](CreateOrgMcpConfigRequest.md) | MCP configuration to create. | 

### Return type

[**OrgMcpConfigResponse**](OrgMcpConfigResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgMcpSecret

> ServiceAccountSecret CreateOrgMcpSecret(ctx, orgId, mcpConfigId, serviceAccountSecretRequest ServiceAccountSecretRequest).Execute()

Create One Secret for One Organization MCP Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "32b6e34b3d91647adeabc012" // string | 
    serviceAccountSecretRequest := *admin.NewServiceAccountSecretRequest(int(8)) // ServiceAccountSecretRequest | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.CreateOrgMcpSecret(context.Background(), orgId, mcpConfigId, &serviceAccountSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.CreateOrgMcpSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgMcpSecret`: ServiceAccountSecret
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.CreateOrgMcpSecret`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgMcpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAccountSecretRequest** | [**ServiceAccountSecretRequest**](ServiceAccountSecretRequest.md) | Secret creation request. | 

### Return type

[**ServiceAccountSecret**](ServiceAccountSecret.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupMcpConfig

> DeleteGroupMcpConfig(ctx, groupId, mcpConfigId).Cascading(cascading).Execute()

Delete One MCP Configuration for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "b9254bc4-d6cc-4325-abf4-fb9d2a9de00a" // string | 
    cascading := true // bool |  (optional) (default to false)

    r, err := sdk.RemoteMCPConfigurationsAPI.DeleteGroupMcpConfig(context.Background(), groupId, mcpConfigId).Cascading(cascading).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.DeleteGroupMcpConfig`: %v (%v)\n", err, r)
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
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cascading** | **bool** | Flag that indicates whether to delete the MCP configuration even if it has active secrets. If false and active secrets exist, the request returns an error. Defaults to false. | [default to false]

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupMcpSecret

> DeleteGroupMcpSecret(ctx, groupId, mcpConfigId, secretId).Execute()

Delete One Secret for One Project MCP Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "32b6e34b3d91647adeabc012" // string | 
    secretId := "32b6e34b3d91647adeabc013" // string | 

    r, err := sdk.RemoteMCPConfigurationsAPI.DeleteGroupMcpSecret(context.Background(), groupId, mcpConfigId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.DeleteGroupMcpSecret`: %v (%v)\n", err, r)
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
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 
**secretId** | **string** | Unique 24-hexadecimal digit string that identifies the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupMcpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgMcpConfig

> DeleteOrgMcpConfig(ctx, orgId, mcpConfigId).Cascading(cascading).Execute()

Delete One MCP Configuration for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "b9254bc4-d6cc-4325-abf4-fb9d2a9de00a" // string | 
    cascading := true // bool |  (optional) (default to false)

    r, err := sdk.RemoteMCPConfigurationsAPI.DeleteOrgMcpConfig(context.Background(), orgId, mcpConfigId).Cascading(cascading).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.DeleteOrgMcpConfig`: %v (%v)\n", err, r)
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
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cascading** | **bool** | Flag that indicates whether to delete the MCP configuration even if it has active secrets. If false and active secrets exist, the request returns an error. Defaults to false. | [default to false]

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgMcpSecret

> DeleteOrgMcpSecret(ctx, orgId, mcpConfigId, secretId).Execute()

Delete One Secret for One Organization MCP Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "32b6e34b3d91647adeabc012" // string | 
    secretId := "32b6e34b3d91647adeabc013" // string | 

    r, err := sdk.RemoteMCPConfigurationsAPI.DeleteOrgMcpSecret(context.Background(), orgId, mcpConfigId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.DeleteOrgMcpSecret`: %v (%v)\n", err, r)
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
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 
**secretId** | **string** | Unique 24-hexadecimal digit string that identifies the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMcpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMcpConfig

> GroupMcpConfigResponse GetGroupMcpConfig(ctx, groupId, mcpConfigId).Execute()

Return One MCP Configuration for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "b9254bc4-d6cc-4325-abf4-fb9d2a9de00a" // string | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.GetGroupMcpConfig(context.Background(), groupId, mcpConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.GetGroupMcpConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupMcpConfig`: GroupMcpConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.GetGroupMcpConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupMcpConfigResponse**](GroupMcpConfigResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMcpSecret

> ServiceAccountSecret GetGroupMcpSecret(ctx, groupId, mcpConfigId, secretId).Execute()

Return One Secret for One Project MCP Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "32b6e34b3d91647adeabc012" // string | 
    secretId := "32b6e34b3d91647adeabc013" // string | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.GetGroupMcpSecret(context.Background(), groupId, mcpConfigId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.GetGroupMcpSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupMcpSecret`: ServiceAccountSecret
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.GetGroupMcpSecret`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 
**secretId** | **string** | Unique 24-hexadecimal digit string that identifies the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMcpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceAccountSecret**](ServiceAccountSecret.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgMcpConfig

> OrgMcpConfigResponse GetOrgMcpConfig(ctx, orgId, mcpConfigId).Execute()

Return One MCP Configuration for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "b9254bc4-d6cc-4325-abf4-fb9d2a9de00a" // string | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.GetOrgMcpConfig(context.Background(), orgId, mcpConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.GetOrgMcpConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgMcpConfig`: OrgMcpConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.GetOrgMcpConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgMcpConfigResponse**](OrgMcpConfigResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgMcpSecret

> ServiceAccountSecret GetOrgMcpSecret(ctx, orgId, mcpConfigId, secretId).Execute()

Return One Secret for One Organization MCP Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "32b6e34b3d91647adeabc012" // string | 
    secretId := "32b6e34b3d91647adeabc013" // string | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.GetOrgMcpSecret(context.Background(), orgId, mcpConfigId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.GetOrgMcpSecret`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgMcpSecret`: ServiceAccountSecret
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.GetOrgMcpSecret`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 
**secretId** | **string** | Unique 24-hexadecimal digit string that identifies the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMcpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceAccountSecret**](ServiceAccountSecret.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupMcpConfigs

> PaginatedGroupMcpConfig ListGroupMcpConfigs(ctx, groupId).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Execute()

Return All MCP Configurations for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    includeCount := true // bool |  (optional) (default to true)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.ListGroupMcpConfigs(context.Background(), groupId).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.ListGroupMcpConfigs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupMcpConfigs`: PaginatedGroupMcpConfig
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.ListGroupMcpConfigs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupMcpConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedGroupMcpConfig**](PaginatedGroupMcpConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupMcpSecrets

> PaginatedMcpConfigSecret ListGroupMcpSecrets(ctx, groupId, mcpConfigId).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Execute()

Return All Secrets for One Project MCP Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "32b6e34b3d91647adeabc012" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    includeCount := true // bool |  (optional) (default to true)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.ListGroupMcpSecrets(context.Background(), groupId, mcpConfigId).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.ListGroupMcpSecrets`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupMcpSecrets`: PaginatedMcpConfigSecret
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.ListGroupMcpSecrets`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupMcpSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedMcpConfigSecret**](PaginatedMcpConfigSecret.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMcpConfigs

> PaginatedOrgMcpConfig ListOrgMcpConfigs(ctx, orgId).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Execute()

Return All MCP Configurations for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    includeCount := true // bool |  (optional) (default to true)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.ListOrgMcpConfigs(context.Background(), orgId).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.ListOrgMcpConfigs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgMcpConfigs`: PaginatedOrgMcpConfig
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.ListOrgMcpConfigs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMcpConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedOrgMcpConfig**](PaginatedOrgMcpConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMcpSecrets

> PaginatedMcpConfigSecret ListOrgMcpSecrets(ctx, orgId, mcpConfigId).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Execute()

Return All Secrets for One Organization MCP Configuration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "32b6e34b3d91647adeabc012" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    includeCount := true // bool |  (optional) (default to true)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.ListOrgMcpSecrets(context.Background(), orgId, mcpConfigId).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.ListOrgMcpSecrets`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgMcpSecrets`: PaginatedMcpConfigSecret
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.ListOrgMcpSecrets`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMcpSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedMcpConfigSecret**](PaginatedMcpConfigSecret.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupMcpConfig

> GroupMcpConfigResponse UpdateGroupMcpConfig(ctx, groupId, mcpConfigId, groupMcpConfigUpdateRequest GroupMcpConfigUpdateRequest).Execute()

Update One MCP Configuration for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "b9254bc4-d6cc-4325-abf4-fb9d2a9de00a" // string | 
    groupMcpConfigUpdateRequest := *admin.NewGroupMcpConfigUpdateRequest() // GroupMcpConfigUpdateRequest | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.UpdateGroupMcpConfig(context.Background(), groupId, mcpConfigId, &groupMcpConfigUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.UpdateGroupMcpConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupMcpConfig`: GroupMcpConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.UpdateGroupMcpConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupMcpConfigUpdateRequest** | [**GroupMcpConfigUpdateRequest**](GroupMcpConfigUpdateRequest.md) | MCP configuration fields to update. | 

### Return type

[**GroupMcpConfigResponse**](GroupMcpConfigResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgMcpConfig

> OrgMcpConfigResponse UpdateOrgMcpConfig(ctx, orgId, mcpConfigId, orgMcpConfigUpdateRequest OrgMcpConfigUpdateRequest).Execute()

Update One MCP Configuration for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312024/admin"
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
    mcpConfigId := "b9254bc4-d6cc-4325-abf4-fb9d2a9de00a" // string | 
    orgMcpConfigUpdateRequest := *admin.NewOrgMcpConfigUpdateRequest() // OrgMcpConfigUpdateRequest | 

    resp, r, err := sdk.RemoteMCPConfigurationsAPI.UpdateOrgMcpConfig(context.Background(), orgId, mcpConfigId, &orgMcpConfigUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteMCPConfigurationsAPI.UpdateOrgMcpConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateOrgMcpConfig`: OrgMcpConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `RemoteMCPConfigurationsAPI.UpdateOrgMcpConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**mcpConfigId** | **string** | Unique identifier of the MCP configuration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgMcpConfigUpdateRequest** | [**OrgMcpConfigUpdateRequest**](OrgMcpConfigUpdateRequest.md) | MCP configuration fields to update. | 

### Return type

[**OrgMcpConfigResponse**](OrgMcpConfigResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

