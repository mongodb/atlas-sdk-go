# \MongoDBCloudUsersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](MongoDBCloudUsersApi.md#CreateUser) | **Post** /api/atlas/v2/users | Create One MongoDB Cloud User
[**GetUser**](MongoDBCloudUsersApi.md#GetUser) | **Get** /api/atlas/v2/users/{userId} | Return One MongoDB Cloud User using Its ID
[**GetUserByUsername**](MongoDBCloudUsersApi.md#GetUserByUsername) | **Get** /api/atlas/v2/users/byName/{userName} | Return One MongoDB Cloud User using Their Username
[**ListOrganizationUsers**](MongoDBCloudUsersApi.md#ListOrganizationUsers) | **Get** /api/atlas/v2/orgs/{orgId}/users | Return All MongoDB Cloud Users in One Organization
[**ListProjectUsers**](MongoDBCloudUsersApi.md#ListProjectUsers) | **Get** /api/atlas/v2/groups/{groupId}/users | Return All MongoDB Cloud Users in One Project
[**ListTeamUsers**](MongoDBCloudUsersApi.md#ListTeamUsers) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Return All MongoDB Cloud Users Assigned to One Team
[**RemoveOrganizationUser**](MongoDBCloudUsersApi.md#RemoveOrganizationUser) | **Delete** /api/atlas/v2/orgs/{orgId}/users/{userId} | Remove One MongoDB Cloud User From One Organization
[**RemoveProjectUser**](MongoDBCloudUsersApi.md#RemoveProjectUser) | **Delete** /api/atlas/v2/groups/{groupId}/users/{userId} | Remove One MongoDB Cloud User from One Project



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

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> CloudAppUser GetUser(ctx, userId).Execute()

Return One MongoDB Cloud User using Its ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByUsername

> CloudAppUser GetUserByUsername(ctx, userName).Execute()

Return One MongoDB Cloud User using Their Username


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationUsers

> PaginatedAppUser ListOrganizationUsers(ctx, orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Username(username).OrgMembershipStatus(orgMembershipStatus).Execute()

Return All MongoDB Cloud Users in One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
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
    // response from `ListOrganizationUsers`: PaginatedAppUser
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

[**PaginatedAppUser**](PaginatedAppUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectUsers

> PaginatedAppUser ListProjectUsers(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).FlattenTeams(flattenTeams).IncludeOrgUsers(includeOrgUsers).OrgMembershipStatus(orgMembershipStatus).Username(username).Execute()

Return All MongoDB Cloud Users in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
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
    // response from `ListProjectUsers`: PaginatedAppUser
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

[**PaginatedAppUser**](PaginatedAppUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamUsers

> PaginatedAppUser ListTeamUsers(ctx, orgId, teamId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Username(username).OrgMembershipStatus(orgMembershipStatus).Execute()

Return All MongoDB Cloud Users Assigned to One Team


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
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
    // response from `ListTeamUsers`: PaginatedAppUser
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

[**PaginatedAppUser**](PaginatedAppUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationUser

> any RemoveOrganizationUser(ctx, orgId, userId).Execute()

Remove One MongoDB Cloud User From One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
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

    resp, r, err := sdk.MongoDBCloudUsersApi.RemoveOrganizationUser(context.Background(), orgId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.RemoveOrganizationUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `RemoveOrganizationUser`: any
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.RemoveOrganizationUser`: %v (%v)\n", resp, r)
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

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

