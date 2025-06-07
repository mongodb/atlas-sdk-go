# \MongoDBCloudUsersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrganizationRole**](MongoDBCloudUsersApi.md#AddOrganizationRole) | **Post** /api/atlas/v2/orgs/{orgId}/users/{userId}:addRole | Add One Organization Role to One MongoDB Cloud User
[**AddProjectRole**](MongoDBCloudUsersApi.md#AddProjectRole) | **Post** /api/atlas/v2/groups/{groupId}/users/{userId}:addRole | Add One Project Role to One MongoDB Cloud User
[**AddProjectUser**](MongoDBCloudUsersApi.md#AddProjectUser) | **Post** /api/atlas/v2/groups/{groupId}/users | Add One MongoDB Cloud User to One Project
[**AddUserToTeam**](MongoDBCloudUsersApi.md#AddUserToTeam) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}:addUser | Add One MongoDB Cloud User to One Team
[**CreateOrganizationUser**](MongoDBCloudUsersApi.md#CreateOrganizationUser) | **Post** /api/atlas/v2/orgs/{orgId}/users | Add One MongoDB Cloud User to One Organization
[**CreateUser**](MongoDBCloudUsersApi.md#CreateUser) | **Post** /api/atlas/v2/users | Create One MongoDB Cloud User
[**GetOrganizationUser**](MongoDBCloudUsersApi.md#GetOrganizationUser) | **Get** /api/atlas/v2/orgs/{orgId}/users/{userId} | Return One MongoDB Cloud User in One Organization
[**GetProjectUser**](MongoDBCloudUsersApi.md#GetProjectUser) | **Get** /api/atlas/v2/groups/{groupId}/users/{userId} | Return One MongoDB Cloud User in One Project
[**GetUser**](MongoDBCloudUsersApi.md#GetUser) | **Get** /api/atlas/v2/users/{userId} | Return One MongoDB Cloud User by ID
[**GetUserByUsername**](MongoDBCloudUsersApi.md#GetUserByUsername) | **Get** /api/atlas/v2/users/byName/{userName} | Return One MongoDB Cloud User by Username
[**ListOrganizationUsers**](MongoDBCloudUsersApi.md#ListOrganizationUsers) | **Get** /api/atlas/v2/orgs/{orgId}/users | Return All MongoDB Cloud Users in One Organization
[**ListProjectUsers**](MongoDBCloudUsersApi.md#ListProjectUsers) | **Get** /api/atlas/v2/groups/{groupId}/users | Return All MongoDB Cloud Users in One Project
[**ListTeamUsers**](MongoDBCloudUsersApi.md#ListTeamUsers) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Return All MongoDB Cloud Users Assigned to One Team
[**RemoveOrganizationRole**](MongoDBCloudUsersApi.md#RemoveOrganizationRole) | **Post** /api/atlas/v2/orgs/{orgId}/users/{userId}:removeRole | Remove One Organization Role from One MongoDB Cloud User
[**RemoveOrganizationUser**](MongoDBCloudUsersApi.md#RemoveOrganizationUser) | **Delete** /api/atlas/v2/orgs/{orgId}/users/{userId} | Remove One MongoDB Cloud User from One Organization
[**RemoveProjectRole**](MongoDBCloudUsersApi.md#RemoveProjectRole) | **Post** /api/atlas/v2/groups/{groupId}/users/{userId}:removeRole | Remove One Project Role from One MongoDB Cloud User
[**RemoveProjectUser**](MongoDBCloudUsersApi.md#RemoveProjectUser) | **Delete** /api/atlas/v2/groups/{groupId}/users/{userId} | Remove One MongoDB Cloud User from One Project
[**RemoveUserFromTeam**](MongoDBCloudUsersApi.md#RemoveUserFromTeam) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}:removeUser | Remove One MongoDB Cloud User from One Team
[**UpdateOrganizationUser**](MongoDBCloudUsersApi.md#UpdateOrganizationUser) | **Patch** /api/atlas/v2/orgs/{orgId}/users/{userId} | Update One MongoDB Cloud User in One Organization



## AddOrganizationRole

> OrgUserResponse AddOrganizationRole(ctx, orgId, userId, addOrRemoveOrgRole AddOrRemoveOrgRole).Execute()

Add One Organization Role to One MongoDB Cloud User


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    userId := "userId_example" // string | 
    addOrRemoveOrgRole := *openapiclient.NewAddOrRemoveOrgRole("OrgRole_example") // AddOrRemoveOrgRole | 

    resp, r, err := sdk.MongoDBCloudUsersApi.AddOrganizationRole(context.Background(), orgId, userId, &addOrRemoveOrgRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.AddOrganizationRole`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `AddOrganizationRole`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.AddOrganizationRole`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addOrRemoveOrgRole** | [**AddOrRemoveOrgRole**](AddOrRemoveOrgRole.md) | Organization-level role to assign to the MongoDB Cloud user. | 

### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddProjectRole

> OrgUserResponse AddProjectRole(ctx, groupId, userId, addOrRemoveGroupRole AddOrRemoveGroupRole).Execute()

Add One Project Role to One MongoDB Cloud User


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    addOrRemoveGroupRole := *openapiclient.NewAddOrRemoveGroupRole("GroupRole_example") // AddOrRemoveGroupRole | 

    resp, r, err := sdk.MongoDBCloudUsersApi.AddProjectRole(context.Background(), groupId, userId, &addOrRemoveGroupRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.AddProjectRole`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `AddProjectRole`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.AddProjectRole`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addOrRemoveGroupRole** | [**AddOrRemoveGroupRole**](AddOrRemoveGroupRole.md) | Project-level role to assign to the MongoDB Cloud user. | 

### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddProjectUser

> GroupUserResponse AddProjectUser(ctx, groupId, groupUserRequest GroupUserRequest).Execute()

Add One MongoDB Cloud User to One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    groupUserRequest := *openapiclient.NewGroupUserRequest([]string{"Roles_example"}, "Username_example") // GroupUserRequest | 

    resp, r, err := sdk.MongoDBCloudUsersApi.AddProjectUser(context.Background(), groupId, &groupUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.AddProjectUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `AddProjectUser`: GroupUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.AddProjectUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUserRequest** | [**GroupUserRequest**](GroupUserRequest.md) | The active or pending MongoDB Cloud user that you want to add to the specified project. | 

### Return type

[**GroupUserResponse**](GroupUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserToTeam

> OrgUserResponse AddUserToTeam(ctx, orgId, teamId, addOrRemoveUserFromTeam AddOrRemoveUserFromTeam).Execute()

Add One MongoDB Cloud User to One Team


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    teamId := "teamId_example" // string | 
    addOrRemoveUserFromTeam := *openapiclient.NewAddOrRemoveUserFromTeam("32b6e34b3d91647abb20e7b8") // AddOrRemoveUserFromTeam | 

    resp, r, err := sdk.MongoDBCloudUsersApi.AddUserToTeam(context.Background(), orgId, teamId, &addOrRemoveUserFromTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.AddUserToTeam`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `AddUserToTeam`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.AddUserToTeam`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team to add the MongoDB Cloud user to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addOrRemoveUserFromTeam** | [**AddOrRemoveUserFromTeam**](AddOrRemoveUserFromTeam.md) | The active or pending MongoDB Cloud user that you want to add to the specified team. | 

### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationUser

> OrgUserResponse CreateOrganizationUser(ctx, orgId, orgUserRequest OrgUserRequest).Execute()

Add One MongoDB Cloud User to One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    orgUserRequest := *openapiclient.NewOrgUserRequest(*openapiclient.NewOrgUserRolesRequest([]string{"OrgRoles_example"}), "Username_example") // OrgUserRequest | 

    resp, r, err := sdk.MongoDBCloudUsersApi.CreateOrganizationUser(context.Background(), orgId, &orgUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.CreateOrganizationUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrganizationUser`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.CreateOrganizationUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgUserRequest** | [**OrgUserRequest**](OrgUserRequest.md) | Represents the MongoDB Cloud user to be created within the organization. | 

### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> CloudAppUser CreateUser(ctx, cloudAppUser CloudAppUser).Execute()

Create One MongoDB Cloud User


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    cloudAppUser := *openapiclient.NewCloudAppUser("Country_example", "EmailAddress_example", "FirstName_example", "LastName_example", "MobileNumber_example", "Password_example", "Username_example") // CloudAppUser | 

    resp, r, err := sdk.MongoDBCloudUsersApi.CreateUser(context.Background(), &cloudAppUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.CreateUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateUser`: CloudAppUser
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.CreateUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAppUser** | [**CloudAppUser**](CloudAppUser.md) | MongoDB Cloud user account to create. | 

### Return type

[**CloudAppUser**](CloudAppUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationUser

> OrgUserResponse GetOrganizationUser(ctx, orgId, userId).Execute()

Return One MongoDB Cloud User in One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    userId := "userId_example" // string | 

    resp, r, err := sdk.MongoDBCloudUsersApi.GetOrganizationUser(context.Background(), orgId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.GetOrganizationUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrganizationUser`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.GetOrganizationUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectUser

> GroupUserResponse GetProjectUser(ctx, groupId, userId).Execute()

Return One MongoDB Cloud User in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    resp, r, err := sdk.MongoDBCloudUsersApi.GetProjectUser(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.GetProjectUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetProjectUser`: GroupUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.GetProjectUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupUserResponse**](GroupUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> CloudAppUser GetUser(ctx, userId).Execute()

Return One MongoDB Cloud User by ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    userId := "userId_example" // string | 

    resp, r, err := sdk.MongoDBCloudUsersApi.GetUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.GetUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetUser`: CloudAppUser
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.GetUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Unique 24-hexadecimal digit string that identifies this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAppUser**](CloudAppUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByUsername

> CloudAppUser GetUserByUsername(ctx, userName).Execute()

Return One MongoDB Cloud User by Username


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    userName := "userName_example" // string | 

    resp, r, err := sdk.MongoDBCloudUsersApi.GetUserByUsername(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.GetUserByUsername`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetUserByUsername`: CloudAppUser
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.GetUserByUsername`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAppUser**](CloudAppUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationUsers

> PaginatedOrgUser ListOrganizationUsers(ctx, orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Username(username).OrgMembershipStatus(orgMembershipStatus).Execute()

Return All MongoDB Cloud Users in One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    username := "username_example" // string |  (optional)
    orgMembershipStatus := "ACTIVE" // string |  (optional)

    resp, r, err := sdk.MongoDBCloudUsersApi.ListOrganizationUsers(context.Background(), orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Username(username).OrgMembershipStatus(orgMembershipStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.ListOrganizationUsers`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrganizationUsers`: PaginatedOrgUser
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.ListOrganizationUsers`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **username** | **string** | Email address to filter users by. Not supported in deprecated versions. | 
 **orgMembershipStatus** | **string** | Organization membership status to filter users by. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions. | 

### Return type

[**PaginatedOrgUser**](PaginatedOrgUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectUsers

> PaginatedGroupUser ListProjectUsers(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).FlattenTeams(flattenTeams).IncludeOrgUsers(includeOrgUsers).OrgMembershipStatus(orgMembershipStatus).Username(username).Execute()

Return All MongoDB Cloud Users in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    flattenTeams := true // bool |  (optional) (default to false)
    includeOrgUsers := true // bool |  (optional) (default to false)
    orgMembershipStatus := "ACTIVE" // string |  (optional)
    username := "username_example" // string |  (optional)

    resp, r, err := sdk.MongoDBCloudUsersApi.ListProjectUsers(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).FlattenTeams(flattenTeams).IncludeOrgUsers(includeOrgUsers).OrgMembershipStatus(orgMembershipStatus).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.ListProjectUsers`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListProjectUsers`: PaginatedGroupUser
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.ListProjectUsers`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **flattenTeams** | **bool** | Flag that indicates whether the returned list should include users who belong to a team with a role in this project. You might not have assigned the individual users a role in this project. If &#x60;\&quot;flattenTeams\&quot; : false&#x60;, this resource returns only users with a role in the project.  If &#x60;\&quot;flattenTeams\&quot; : true&#x60;, this resource returns both users with roles in the project and users who belong to teams with roles in the project. | [default to false]
 **includeOrgUsers** | **bool** | Flag that indicates whether the returned list should include users with implicit access to the project, the Organization Owner or Organization Read Only role. You might not have assigned the individual users a role in this project. If &#x60;\&quot;includeOrgUsers\&quot;: false&#x60;, this resource returns only users with a role in the project. If &#x60;\&quot;includeOrgUsers\&quot;: true&#x60;, this resource returns both users with roles in the project and users who have implicit access to the project through their organization role. | [default to false]
 **orgMembershipStatus** | **string** | Flag that indicates whether to filter the returned list by users organization membership status. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions. | 
 **username** | **string** | Email address to filter users by. Not supported in deprecated versions. | 

### Return type

[**PaginatedGroupUser**](PaginatedGroupUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamUsers

> PaginatedOrgUser ListTeamUsers(ctx, orgId, teamId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Username(username).OrgMembershipStatus(orgMembershipStatus).Execute()

Return All MongoDB Cloud Users Assigned to One Team


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    teamId := "teamId_example" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    username := "username_example" // string |  (optional)
    orgMembershipStatus := "ACTIVE" // string |  (optional)

    resp, r, err := sdk.MongoDBCloudUsersApi.ListTeamUsers(context.Background(), orgId, teamId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Username(username).OrgMembershipStatus(orgMembershipStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.ListTeamUsers`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListTeamUsers`: PaginatedOrgUser
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.ListTeamUsers`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team whose application users you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **username** | **string** | Email address to filter users by. Not supported in deprecated versions. | 
 **orgMembershipStatus** | **string** | Organization membership status to filter users by. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions. | 

### Return type

[**PaginatedOrgUser**](PaginatedOrgUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationRole

> OrgUserResponse RemoveOrganizationRole(ctx, orgId, userId, addOrRemoveOrgRole AddOrRemoveOrgRole).Execute()

Remove One Organization Role from One MongoDB Cloud User


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    userId := "userId_example" // string | 
    addOrRemoveOrgRole := *openapiclient.NewAddOrRemoveOrgRole("OrgRole_example") // AddOrRemoveOrgRole | 

    resp, r, err := sdk.MongoDBCloudUsersApi.RemoveOrganizationRole(context.Background(), orgId, userId, &addOrRemoveOrgRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.RemoveOrganizationRole`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `RemoveOrganizationRole`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.RemoveOrganizationRole`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addOrRemoveOrgRole** | [**AddOrRemoveOrgRole**](AddOrRemoveOrgRole.md) | Organization-level role to remove from the MongoDB Cloud user. | 

### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationUser

> RemoveOrganizationUser(ctx, orgId, userId).Execute()

Remove One MongoDB Cloud User from One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    userId := "userId_example" // string | 

    r, err := sdk.MongoDBCloudUsersApi.RemoveOrganizationUser(context.Background(), orgId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.RemoveOrganizationUser`: %v (%v)\n", err, r)
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
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the [Return All MongoDB Cloud Users in One Organization](#tag/MongoDB-Cloud-Users/operation/listOrganizationUsers) resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProjectRole

> OrgUserResponse RemoveProjectRole(ctx, groupId, userId, addOrRemoveGroupRole AddOrRemoveGroupRole).Execute()

Remove One Project Role from One MongoDB Cloud User


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    addOrRemoveGroupRole := *openapiclient.NewAddOrRemoveGroupRole("GroupRole_example") // AddOrRemoveGroupRole | 

    resp, r, err := sdk.MongoDBCloudUsersApi.RemoveProjectRole(context.Background(), groupId, userId, &addOrRemoveGroupRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.RemoveProjectRole`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `RemoveProjectRole`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.RemoveProjectRole`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addOrRemoveGroupRole** | [**AddOrRemoveGroupRole**](AddOrRemoveGroupRole.md) | Project-level role to remove from the MongoDB Cloud user. | 

### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProjectUser

> RemoveProjectUser(ctx, groupId, userId).Execute()

Remove One MongoDB Cloud User from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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

    r, err := sdk.MongoDBCloudUsersApi.RemoveProjectUser(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.RemoveProjectUser`: %v (%v)\n", err, r)
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
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the [Return All MongoDB Cloud Users in One Project](#tag/MongoDB-Cloud-Users/operation/listProjectUsers) resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromTeam

> OrgUserResponse RemoveUserFromTeam(ctx, orgId, teamId, addOrRemoveUserFromTeam AddOrRemoveUserFromTeam).Execute()

Remove One MongoDB Cloud User from One Team


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    teamId := "teamId_example" // string | 
    addOrRemoveUserFromTeam := *openapiclient.NewAddOrRemoveUserFromTeam("32b6e34b3d91647abb20e7b8") // AddOrRemoveUserFromTeam | 

    resp, r, err := sdk.MongoDBCloudUsersApi.RemoveUserFromTeam(context.Background(), orgId, teamId, &addOrRemoveUserFromTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.RemoveUserFromTeam`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `RemoveUserFromTeam`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.RemoveUserFromTeam`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team to remove the MongoDB user from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addOrRemoveUserFromTeam** | [**AddOrRemoveUserFromTeam**](AddOrRemoveUserFromTeam.md) | The id of the active or pending MongoDB Cloud user that you want to remove from the specified team. | 

### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationUser

> OrgUserResponse UpdateOrganizationUser(ctx, orgId, userId, orgUserUpdateRequest OrgUserUpdateRequest).Execute()

Update One MongoDB Cloud User in One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312004/admin"
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
    userId := "userId_example" // string | 
    orgUserUpdateRequest := *openapiclient.NewOrgUserUpdateRequest() // OrgUserUpdateRequest | 

    resp, r, err := sdk.MongoDBCloudUsersApi.UpdateOrganizationUser(context.Background(), orgId, userId, &orgUserUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.UpdateOrganizationUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateOrganizationUser`: OrgUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.UpdateOrganizationUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user&#39;s userId or verify a user&#39;s status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgUserUpdateRequest** | [**OrgUserUpdateRequest**](OrgUserUpdateRequest.md) | Represents the roles and teams to assign the MongoDB Cloud user. | 

### Return type

[**OrgUserResponse**](OrgUserResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-02-19+json
- **Accept**: application/vnd.atlas.2025-02-19+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

