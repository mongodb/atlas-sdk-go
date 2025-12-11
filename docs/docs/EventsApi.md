# \EventsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupEvent**](EventsApi.md#GetGroupEvent) | **Get** /api/atlas/v2/groups/{groupId}/events/{eventId} | Return One Event from One Project
[**GetOrgEvent**](EventsApi.md#GetOrgEvent) | **Get** /api/atlas/v2/orgs/{orgId}/events/{eventId} | Return One Event from One Organization
[**ListEventTypes**](EventsApi.md#ListEventTypes) | **Get** /api/atlas/v2/eventTypes | Return All Event Types
[**ListGroupEvents**](EventsApi.md#ListGroupEvents) | **Get** /api/atlas/v2/groups/{groupId}/events | Return Events from One Project
[**ListOrgEvents**](EventsApi.md#ListOrgEvents) | **Get** /api/atlas/v2/orgs/{orgId}/events | Return Events from One Organization



## GetGroupEvent

> EventViewForNdsGroup GetGroupEvent(ctx, groupId, eventId).IncludeRaw(includeRaw).Execute()

Return One Event from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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
    eventId := "eventId_example" // string | 
    includeRaw := true // bool |  (optional) (default to false)

    resp, r, err := sdk.EventsApi.GetGroupEvent(context.Background(), groupId, eventId).IncludeRaw(includeRaw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetGroupEvent`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupEvent`: EventViewForNdsGroup
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetGroupEvent`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**eventId** | **string** | Unique 24-hexadecimal digit string that identifies the event that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]

### Return type

[**EventViewForNdsGroup**](EventViewForNdsGroup.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgEvent

> EventViewForOrg GetOrgEvent(ctx, orgId, eventId).IncludeRaw(includeRaw).Execute()

Return One Event from One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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
    eventId := "eventId_example" // string | 
    includeRaw := true // bool |  (optional) (default to false)

    resp, r, err := sdk.EventsApi.GetOrgEvent(context.Background(), orgId, eventId).IncludeRaw(includeRaw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetOrgEvent`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgEvent`: EventViewForOrg
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetOrgEvent`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**eventId** | **string** | Unique 24-hexadecimal digit string that identifies the event that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]

### Return type

[**EventViewForOrg**](EventViewForOrg.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventTypes

> PaginatedEventTypeDetailsResponse ListEventTypes(ctx).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Event Types


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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

    resp, r, err := sdk.EventsApi.ListEventTypes(context.Background()).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEventTypes`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListEventTypes`: PaginatedEventTypeDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEventTypes`: %v (%v)\n", resp, r)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedEventTypeDetailsResponse**](PaginatedEventTypeDetailsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupEvents

> GroupPaginatedEvent ListGroupEvents(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ClusterNames(clusterNames).EventType(eventType).ExcludedEventType(excludedEventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()

Return Events from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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
    clusterNames := []string{"Inner_example"} // []string |  (optional)
    eventType := []string{"Inner_example"} // []string |  (optional)
    excludedEventType := []string{"Inner_example"} // []string |  (optional)
    includeRaw := true // bool |  (optional) (default to false)
    maxDate := time.Now() // time.Time |  (optional)
    minDate := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.EventsApi.ListGroupEvents(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ClusterNames(clusterNames).EventType(eventType).ExcludedEventType(excludedEventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListGroupEvents`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupEvents`: GroupPaginatedEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListGroupEvents`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **clusterNames** | **[]string** | Human-readable label that identifies the cluster. | 
 **eventType** | **[]string** | Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently. | 
 **excludedEventType** | **[]string** | Category of event that you would like to exclude from query results, such as CLUSTER_CREATED  **IMPORTANT**: Event type names change frequently. Verify that you specify the event type correctly by checking the complete list of event types. | 
 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]
 **maxDate** | **time.Time** | Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC. | 
 **minDate** | **time.Time** | Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC. | 

### Return type

[**GroupPaginatedEvent**](GroupPaginatedEvent.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgEvents

> OrgPaginatedEvent ListOrgEvents(ctx, orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()

Return Events from One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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
    eventType := []string{"Inner_example"} // []string |  (optional)
    includeRaw := true // bool |  (optional) (default to false)
    maxDate := time.Now() // time.Time |  (optional)
    minDate := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.EventsApi.ListOrgEvents(context.Background(), orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListOrgEvents`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgEvents`: OrgPaginatedEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListOrgEvents`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **eventType** | **[]string** | Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently. | 
 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]
 **maxDate** | **time.Time** | Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC. | 
 **minDate** | **time.Time** | Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC. | 

### Return type

[**OrgPaginatedEvent**](OrgPaginatedEvent.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

