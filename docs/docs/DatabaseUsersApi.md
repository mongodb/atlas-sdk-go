# \DatabaseUsersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabaseUser**](DatabaseUsersApi.md#CreateDatabaseUser) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers | Create One Database User in One Project
[**DeleteDatabaseUser**](DatabaseUsersApi.md#DeleteDatabaseUser) | **Delete** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Remove One Database User from One Project
[**GetDatabaseUser**](DatabaseUsersApi.md#GetDatabaseUser) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Return One Database User from One Project
[**ListDatabaseUsers**](DatabaseUsersApi.md#ListDatabaseUsers) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers | Return All Database Users from One Project
[**UpdateDatabaseUser**](DatabaseUsersApi.md#UpdateDatabaseUser) | **Patch** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Update One Database User in One Project



## CreateDatabaseUser

> CloudDatabaseUser CreateDatabaseUser(ctx, groupId, cloudDatabaseUser CloudDatabaseUser).Execute()

Create One Database User in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312001/admin"
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
    cloudDatabaseUser := *openapiclient.NewCloudDatabaseUser("DatabaseName_example", "GroupId_example", "Username_example") // CloudDatabaseUser | 

    resp, r, err := sdk.DatabaseUsersApi.CreateDatabaseUser(context.Background(), groupId, &cloudDatabaseUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.CreateDatabaseUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateDatabaseUser`: CloudDatabaseUser
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.CreateDatabaseUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudDatabaseUser** | [**CloudDatabaseUser**](CloudDatabaseUser.md) | Creates one database user in the specified project. | 

### Return type

[**CloudDatabaseUser**](CloudDatabaseUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseUser

> any DeleteDatabaseUser(ctx, groupId, databaseName, username).Execute()

Remove One Database User from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312001/admin"
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
    databaseName := "databaseName_example" // string | 
    username := "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName" // string | 

    resp, r, err := sdk.DatabaseUsersApi.DeleteDatabaseUser(context.Background(), groupId, databaseName, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.DeleteDatabaseUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteDatabaseUser`: any
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.DeleteDatabaseUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**databaseName** | **string** | The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be &#x60;admin&#x60;. | 
**username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsIAMType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsIAMType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | OIDC Workforce | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP group name | | OIDC Workload | oidcAuthType | USER | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP user name | | SCRAM-SHA | awsIAMType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseUser

> CloudDatabaseUser GetDatabaseUser(ctx, groupId, databaseName, username).Execute()

Return One Database User from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312001/admin"
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
    databaseName := "databaseName_example" // string | 
    username := "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName" // string | 

    resp, r, err := sdk.DatabaseUsersApi.GetDatabaseUser(context.Background(), groupId, databaseName, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.GetDatabaseUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetDatabaseUser`: CloudDatabaseUser
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.GetDatabaseUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**databaseName** | **string** | The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be &#x60;admin&#x60;. | 
**username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsIAMType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsIAMType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | OIDC Workforce | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP group name | | OIDC Workload | oidcAuthType | USER | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP user name | | SCRAM-SHA | awsIAMType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CloudDatabaseUser**](CloudDatabaseUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseUsers

> PaginatedApiAtlasDatabaseUser ListDatabaseUsers(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Database Users from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312001/admin"
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

    resp, r, err := sdk.DatabaseUsersApi.ListDatabaseUsers(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.ListDatabaseUsers`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListDatabaseUsers`: PaginatedApiAtlasDatabaseUser
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.ListDatabaseUsers`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiAtlasDatabaseUser**](PaginatedApiAtlasDatabaseUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabaseUser

> CloudDatabaseUser UpdateDatabaseUser(ctx, groupId, databaseName, username, cloudDatabaseUser CloudDatabaseUser).Execute()

Update One Database User in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312001/admin"
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
    databaseName := "databaseName_example" // string | 
    username := "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName" // string | 
    cloudDatabaseUser := *openapiclient.NewCloudDatabaseUser("DatabaseName_example", "GroupId_example", "Username_example") // CloudDatabaseUser | 

    resp, r, err := sdk.DatabaseUsersApi.UpdateDatabaseUser(context.Background(), groupId, databaseName, username, &cloudDatabaseUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.UpdateDatabaseUser`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateDatabaseUser`: CloudDatabaseUser
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.UpdateDatabaseUser`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**databaseName** | **string** | The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be &#x60;admin&#x60;. | 
**username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsIAMType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsIAMType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | OIDC Workforce | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP group name | | OIDC Workload | oidcAuthType | USER | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP user name | | SCRAM-SHA | awsIAMType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cloudDatabaseUser** | [**CloudDatabaseUser**](CloudDatabaseUser.md) | Updates one database user that belongs to the specified project. | 

### Return type

[**CloudDatabaseUser**](CloudDatabaseUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

