# \ActivityFeedApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupActivityFeed**](ActivityFeedApi.md#GetGroupActivityFeed) | **Get** /api/atlas/v2/groups/{groupId}/activityFeed | Return Pre-Filtered Activity Feed Link for One Project
[**GetOrgActivityFeed**](ActivityFeedApi.md#GetOrgActivityFeed) | **Get** /api/atlas/v2/orgs/{orgId}/activityFeed | Return Pre-Filtered Activity Feed Link for One Organization



## GetGroupActivityFeed

> ActivityFeedLinkResponse GetGroupActivityFeed(ctx, groupId).EventType(eventType).MaxDate(maxDate).MinDate(minDate).ClusterName(clusterName).Execute()

Return Pre-Filtered Activity Feed Link for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312013/admin"
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
    eventType := []string{"Inner_example"} // []string |  (optional)
    maxDate := time.Now() // time.Time |  (optional)
    minDate := time.Now() // time.Time |  (optional)
    clusterName := []string{"Inner_example"} // []string |  (optional)

    resp, r, err := sdk.ActivityFeedApi.GetGroupActivityFeed(context.Background(), groupId).EventType(eventType).MaxDate(maxDate).MinDate(minDate).ClusterName(clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedApi.GetGroupActivityFeed`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupActivityFeed`: ActivityFeedLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityFeedApi.GetGroupActivityFeed`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupActivityFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventType** | **[]string** | Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently. | 
 **maxDate** | **time.Time** | End date and time for events to include in the activity feed link. ISO 8601 timestamp format in UTC. | 
 **minDate** | **time.Time** | Start date and time for events to include in the activity feed link. ISO 8601 timestamp format in UTC. | 
 **clusterName** | **[]string** | Human-readable label that identifies the cluster. | 

### Return type

[**ActivityFeedLinkResponse**](ActivityFeedLinkResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgActivityFeed

> ActivityFeedLinkResponse GetOrgActivityFeed(ctx, orgId).EventType(eventType).MaxDate(maxDate).MinDate(minDate).Execute()

Return Pre-Filtered Activity Feed Link for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312013/admin"
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
    eventType := []string{"Inner_example"} // []string |  (optional)
    maxDate := time.Now() // time.Time |  (optional)
    minDate := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.ActivityFeedApi.GetOrgActivityFeed(context.Background(), orgId).EventType(eventType).MaxDate(maxDate).MinDate(minDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedApi.GetOrgActivityFeed`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgActivityFeed`: ActivityFeedLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityFeedApi.GetOrgActivityFeed`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgActivityFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventType** | **[]string** | Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently. | 
 **maxDate** | **time.Time** | End date and time for events to include in the activity feed link. ISO 8601 timestamp format in UTC. | 
 **minDate** | **time.Time** | Start date and time for events to include in the activity feed link. ISO 8601 timestamp format in UTC. | 

### Return type

[**ActivityFeedLinkResponse**](ActivityFeedLinkResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

