# \MongoDBCloudUsersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](MongoDBCloudUsersApi.md#CreateUser) | **Post** /api/atlas/v2/users | Create One MongoDB Cloud User
[**GetUser**](MongoDBCloudUsersApi.md#GetUser) | **Get** /api/atlas/v2/users/{userId} | Return One MongoDB Cloud User using Its ID
[**GetUserByUsername**](MongoDBCloudUsersApi.md#GetUserByUsername) | **Get** /api/atlas/v2/users/byName/{userName} | Return One MongoDB Cloud User using Their Username



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

    "go.mongodb.org/atlas-sdk/v20240530005/admin"
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

    "go.mongodb.org/atlas-sdk/v20240530005/admin"
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

    "go.mongodb.org/atlas-sdk/v20240530005/admin"
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

