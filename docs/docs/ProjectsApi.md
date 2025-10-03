# \ProjectsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupUser**](ProjectsApi.md#AddGroupUser) | **Post** /api/atlas/v2/groups/{groupId}/access | Add One MongoDB Cloud User to One Project
[**CreateGroup**](ProjectsApi.md#CreateGroup) | **Post** /api/atlas/v2/groups | Create One Project
[**CreateGroupInvite**](ProjectsApi.md#CreateGroupInvite) | **Post** /api/atlas/v2/groups/{groupId}/invites | Create Invitation for One MongoDB Cloud User in One Project
[**DeleteGroup**](ProjectsApi.md#DeleteGroup) | **Delete** /api/atlas/v2/groups/{groupId} | Remove One Project
[**DeleteGroupInvite**](ProjectsApi.md#DeleteGroupInvite) | **Delete** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Remove One Invitation from One Project
[**DeleteGroupLimit**](ProjectsApi.md#DeleteGroupLimit) | **Delete** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Remove One Project Limit
[**GetGroup**](ProjectsApi.md#GetGroup) | **Get** /api/atlas/v2/groups/{groupId} | Return One Project
[**GetGroupByName**](ProjectsApi.md#GetGroupByName) | **Get** /api/atlas/v2/groups/byName/{groupName} | Return One Project by Name
[**GetGroupInvite**](ProjectsApi.md#GetGroupInvite) | **Get** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Return One Invitation in One Project by Invitation ID
[**GetGroupIpAddresses**](ProjectsApi.md#GetGroupIpAddresses) | **Get** /api/atlas/v2/groups/{groupId}/ipAddresses | Return All IP Addresses for One Project
[**GetGroupLimit**](ProjectsApi.md#GetGroupLimit) | **Get** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Return One Limit for One Project
[**GetGroupSettings**](ProjectsApi.md#GetGroupSettings) | **Get** /api/atlas/v2/groups/{groupId}/settings | Return Project Settings
[**GetMongoDbVersions**](ProjectsApi.md#GetMongoDbVersions) | **Get** /api/atlas/v2/groups/{groupId}/mongoDBVersions | Return All Available MongoDB LTS Versions for Clusters in One Project
[**ListGroupInvites**](ProjectsApi.md#ListGroupInvites) | **Get** /api/atlas/v2/groups/{groupId}/invites | Return All Invitations in One Project
[**ListGroupLimits**](ProjectsApi.md#ListGroupLimits) | **Get** /api/atlas/v2/groups/{groupId}/limits | Return All Limits for One Project
[**ListGroups**](ProjectsApi.md#ListGroups) | **Get** /api/atlas/v2/groups | Return All Projects
[**MigrateGroup**](ProjectsApi.md#MigrateGroup) | **Post** /api/atlas/v2/groups/{groupId}:migrate | Migrate One Project to Another Organization
[**SetGroupLimit**](ProjectsApi.md#SetGroupLimit) | **Patch** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Set One Project Limit
[**UpdateGroup**](ProjectsApi.md#UpdateGroup) | **Patch** /api/atlas/v2/groups/{groupId} | Update One Project
[**UpdateGroupInvites**](ProjectsApi.md#UpdateGroupInvites) | **Patch** /api/atlas/v2/groups/{groupId}/invites | Update One Invitation in One Project
[**UpdateGroupSettings**](ProjectsApi.md#UpdateGroupSettings) | **Patch** /api/atlas/v2/groups/{groupId}/settings | Update Project Settings
[**UpdateGroupUserRoles**](ProjectsApi.md#UpdateGroupUserRoles) | **Put** /api/atlas/v2/groups/{groupId}/users/{userId}/roles | Update Project Roles for One MongoDB Cloud User
[**UpdateInviteById**](ProjectsApi.md#UpdateInviteById) | **Patch** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Update One Invitation in One Project by Invitation ID



## AddGroupUser

> OrganizationInvitation AddGroupUser(ctx, groupId, groupInvitationRequest GroupInvitationRequest).Execute()

Add One MongoDB Cloud User to One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    groupInvitationRequest := *openapiclient.NewGroupInvitationRequest() // GroupInvitationRequest | 

    resp, r, err := sdk.ProjectsApi.AddGroupUser(context.Background(), groupId, &groupInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.AddGroupUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `AddGroupUser`: OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.AddGroupUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupInvitationRequest** | [**GroupInvitationRequest**](GroupInvitationRequest.md) | Adds one MongoDB Cloud user to the specified project. | 

### Return type

[**OrganizationInvitation**](OrganizationInvitation.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> Group CreateGroup(ctx, group Group).ProjectOwnerId(projectOwnerId).Execute()

Create One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    group := *openapiclient.NewGroup(int64(123), time.Now(), "Name_example", "32b6e34b3d91647abb20e7b8") // Group | 
    projectOwnerId := "projectOwnerId_example" // string |  (optional)

    resp, r, err := sdk.ProjectsApi.CreateGroup(context.Background(), &group).ProjectOwnerId(projectOwnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateGroup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateGroup`: %v (%v)\n", resp, r)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**Group**](Group.md) | Creates one project. | 
 **projectOwnerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user to whom to grant the Project Owner role on the specified project. If you set this parameter, it overrides the default value of the oldest Organization Owner. | 

### Return type

[**Group**](Group.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupInvite

> GroupInvitation CreateGroupInvite(ctx, groupId, groupInvitationRequest GroupInvitationRequest).Execute()

Create Invitation for One MongoDB Cloud User in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    groupInvitationRequest := *openapiclient.NewGroupInvitationRequest() // GroupInvitationRequest | 

    resp, r, err := sdk.ProjectsApi.CreateGroupInvite(context.Background(), groupId, &groupInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateGroupInvite`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupInvite`: GroupInvitation
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateGroupInvite`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupInvitationRequest** | [**GroupInvitationRequest**](GroupInvitationRequest.md) | Invites one MongoDB Cloud user to join the specified project. | 

### Return type

[**GroupInvitation**](GroupInvitation.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()

Remove One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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

    r, err := sdk.ProjectsApi.DeleteGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteGroup`: %v (%v)\n", err, r)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupInvite

> DeleteGroupInvite(ctx, groupId, invitationId).Execute()

Remove One Invitation from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    invitationId := "invitationId_example" // string | 

    r, err := sdk.ProjectsApi.DeleteGroupInvite(context.Background(), groupId, invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteGroupInvite`: %v (%v)\n", err, r)
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
**invitationId** | **string** | Unique 24-hexadecimal digit string that identifies the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupLimit

> DeleteGroupLimit(ctx, limitName, groupId).Execute()

Remove One Project Limit


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    limitName := "limitName_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.ProjectsApi.DeleteGroupLimit(context.Background(), limitName, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteGroupLimit`: %v (%v)\n", err, r)
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
**limitName** | **string** | Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 100 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 100 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|  | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> Group GetGroup(ctx, groupId).Execute()

Return One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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

    resp, r, err := sdk.ProjectsApi.GetGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetGroup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetGroup`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupByName

> Group GetGroupByName(ctx, groupName).Execute()

Return One Project by Name


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupName := "groupName_example" // string | 

    resp, r, err := sdk.ProjectsApi.GetGroupByName(context.Background(), groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetGroupByName`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupByName`: Group
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetGroupByName`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | Human-readable label that identifies this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupInvite

> GroupInvitation GetGroupInvite(ctx, groupId, invitationId).Execute()

Return One Invitation in One Project by Invitation ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    invitationId := "invitationId_example" // string | 

    resp, r, err := sdk.ProjectsApi.GetGroupInvite(context.Background(), groupId, invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetGroupInvite`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupInvite`: GroupInvitation
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetGroupInvite`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**invitationId** | **string** | Unique 24-hexadecimal digit string that identifies the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupInvitation**](GroupInvitation.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupIpAddresses

> GroupIPAddresses GetGroupIpAddresses(ctx, groupId).Execute()

Return All IP Addresses for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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

    resp, r, err := sdk.ProjectsApi.GetGroupIpAddresses(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetGroupIpAddresses`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupIpAddresses`: GroupIPAddresses
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetGroupIpAddresses`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupIpAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupIPAddresses**](GroupIPAddresses.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupLimit

> DataFederationLimit GetGroupLimit(ctx, limitName, groupId).Execute()

Return One Limit for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    limitName := "limitName_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.ProjectsApi.GetGroupLimit(context.Background(), limitName, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetGroupLimit`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupLimit`: DataFederationLimit
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetGroupLimit`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitName** | **string** | Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 100 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 100 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|  | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataFederationLimit**](DataFederationLimit.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupSettings

> GroupSettings GetGroupSettings(ctx, groupId).Execute()

Return Project Settings


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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

    resp, r, err := sdk.ProjectsApi.GetGroupSettings(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetGroupSettings`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupSettings`: GroupSettings
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetGroupSettings`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupSettings**](GroupSettings.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMongoDbVersions

> PaginatedAvailableVersion GetMongoDbVersions(ctx, groupId).CloudProvider(cloudProvider).InstanceSize(instanceSize).DefaultStatus(defaultStatus).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Available MongoDB LTS Versions for Clusters in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    cloudProvider := "cloudProvider_example" // string |  (optional)
    instanceSize := "M10" // string |  (optional)
    defaultStatus := "defaultStatus_example" // string |  (optional)
    itemsPerPage := int64(100) // int64 |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ProjectsApi.GetMongoDbVersions(context.Background(), groupId).CloudProvider(cloudProvider).InstanceSize(instanceSize).DefaultStatus(defaultStatus).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetMongoDbVersions`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetMongoDbVersions`: PaginatedAvailableVersion
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetMongoDbVersions`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMongoDbVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudProvider** | **string** | Filter results to only one cloud provider. | 
 **instanceSize** | **string** | Filter results to only one instance size. | 
 **defaultStatus** | **string** | Filter results to only the default values per tier. This value must be DEFAULT. | 
 **itemsPerPage** | **int64** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAvailableVersion**](PaginatedAvailableVersion.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupInvites

> []GroupInvitation ListGroupInvites(ctx, groupId).Username(username).Execute()

Return All Invitations in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    username := "username_example" // string |  (optional)

    resp, r, err := sdk.ProjectsApi.ListGroupInvites(context.Background(), groupId).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ListGroupInvites`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupInvites`: []GroupInvitation
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ListGroupInvites`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | Email address of the user account invited to this project. | 

### Return type

[**[]GroupInvitation**](GroupInvitation.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupLimits

> []DataFederationLimit ListGroupLimits(ctx, groupId).Execute()

Return All Limits for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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

    resp, r, err := sdk.ProjectsApi.ListGroupLimits(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ListGroupLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupLimits`: []DataFederationLimit
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ListGroupLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DataFederationLimit**](DataFederationLimit.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> PaginatedAtlasGroup ListGroups(ctx).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Projects


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ProjectsApi.ListGroups(context.Background()).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ListGroups`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroups`: PaginatedAtlasGroup
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ListGroups`: %v (%v)\n", resp, r)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAtlasGroup**](PaginatedAtlasGroup.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateGroup

> Group MigrateGroup(ctx, groupId, groupMigrationRequest GroupMigrationRequest).Execute()

Migrate One Project to Another Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    groupMigrationRequest := *openapiclient.NewGroupMigrationRequest() // GroupMigrationRequest | 

    resp, r, err := sdk.ProjectsApi.MigrateGroup(context.Background(), groupId, &groupMigrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.MigrateGroup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `MigrateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.MigrateGroup`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupMigrationRequest** | [**GroupMigrationRequest**](GroupMigrationRequest.md) | Migrates a project from its current organization to another organization. | 

### Return type

[**Group**](Group.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupLimit

> DataFederationLimit SetGroupLimit(ctx, limitName, groupId, dataFederationLimit DataFederationLimit).Execute()

Set One Project Limit


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    limitName := "limitName_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    dataFederationLimit := *openapiclient.NewDataFederationLimit("Name_example", int64(123)) // DataFederationLimit | 

    resp, r, err := sdk.ProjectsApi.SetGroupLimit(context.Background(), limitName, groupId, &dataFederationLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.SetGroupLimit`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `SetGroupLimit`: DataFederationLimit
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.SetGroupLimit`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitName** | **string** | Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 100 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 100 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|  | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataFederationLimit** | [**DataFederationLimit**](DataFederationLimit.md) | Limit to update. | 

### Return type

[**DataFederationLimit**](DataFederationLimit.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, groupId, groupUpdate GroupUpdate).Execute()

Update One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    groupUpdate := *openapiclient.NewGroupUpdate() // GroupUpdate | 

    resp, r, err := sdk.ProjectsApi.UpdateGroup(context.Background(), groupId, &groupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateGroup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateGroup`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdate** | [**GroupUpdate**](GroupUpdate.md) | Project to update. | 

### Return type

[**Group**](Group.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupInvites

> GroupInvitation UpdateGroupInvites(ctx, groupId, groupInvitationRequest GroupInvitationRequest).Execute()

Update One Invitation in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    groupInvitationRequest := *openapiclient.NewGroupInvitationRequest() // GroupInvitationRequest | 

    resp, r, err := sdk.ProjectsApi.UpdateGroupInvites(context.Background(), groupId, &groupInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateGroupInvites`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupInvites`: GroupInvitation
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateGroupInvites`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupInvitationRequest** | [**GroupInvitationRequest**](GroupInvitationRequest.md) | Updates the details of one pending invitation to the specified project. | 

### Return type

[**GroupInvitation**](GroupInvitation.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupSettings

> GroupSettings UpdateGroupSettings(ctx, groupId, groupSettings GroupSettings).Execute()

Update Project Settings


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    groupSettings := *openapiclient.NewGroupSettings() // GroupSettings | 

    resp, r, err := sdk.ProjectsApi.UpdateGroupSettings(context.Background(), groupId, &groupSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateGroupSettings`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupSettings`: GroupSettings
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateGroupSettings`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupSettings** | [**GroupSettings**](GroupSettings.md) | Settings to update. | 

### Return type

[**GroupSettings**](GroupSettings.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupUserRoles

> UpdateGroupRolesForUser UpdateGroupUserRoles(ctx, groupId, userId, updateGroupRolesForUser UpdateGroupRolesForUser).Execute()

Update Project Roles for One MongoDB Cloud User


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    userId := "userId_example" // string | 
    updateGroupRolesForUser := *openapiclient.NewUpdateGroupRolesForUser() // UpdateGroupRolesForUser | 

    resp, r, err := sdk.ProjectsApi.UpdateGroupUserRoles(context.Background(), groupId, userId, &updateGroupRolesForUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateGroupUserRoles`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupUserRoles`: UpdateGroupRolesForUser
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateGroupUserRoles`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the user to modify. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupRolesForUser** | [**UpdateGroupRolesForUser**](UpdateGroupRolesForUser.md) | Roles to update for the specified user. | 

### Return type

[**UpdateGroupRolesForUser**](UpdateGroupRolesForUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInviteById

> GroupInvitation UpdateInviteById(ctx, groupId, invitationId, groupInvitationUpdateRequest GroupInvitationUpdateRequest).Execute()

Update One Invitation in One Project by Invitation ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    invitationId := "invitationId_example" // string | 
    groupInvitationUpdateRequest := *openapiclient.NewGroupInvitationUpdateRequest() // GroupInvitationUpdateRequest | 

    resp, r, err := sdk.ProjectsApi.UpdateInviteById(context.Background(), groupId, invitationId, &groupInvitationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateInviteById`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateInviteById`: GroupInvitation
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateInviteById`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**invitationId** | **string** | Unique 24-hexadecimal digit string that identifies the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInviteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupInvitationUpdateRequest** | [**GroupInvitationUpdateRequest**](GroupInvitationUpdateRequest.md) | Updates the details of one pending invitation to the specified project. | 

### Return type

[**GroupInvitation**](GroupInvitation.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

