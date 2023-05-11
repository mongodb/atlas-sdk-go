# \EventsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationEvent**](EventsApi.md#GetOrganizationEvent) | **Get** /api/atlas/v2/orgs/{orgId}/events/{eventId} | Return One Event from One Organization
[**GetProjectEvent**](EventsApi.md#GetProjectEvent) | **Get** /api/atlas/v2/groups/{groupId}/events/{eventId} | Return One Event from One Project
[**ListOrganizationEvents**](EventsApi.md#ListOrganizationEvents) | **Get** /api/atlas/v2/orgs/{orgId}/events | Return All Events from One Organization
[**ListProjectEvents**](EventsApi.md#ListProjectEvents) | **Get** /api/atlas/v2/groups/{groupId}/events | Return All Events from One Project



## GetOrganizationEvent

> EventViewForOrg GetOrganizationEvent(ctx, orgId, eventId).IncludeRaw(includeRaw).Execute()

Return One Event from One Organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    orgId := "4888442a3354817a7320eb61" // string | 
    eventId := "eventId_example" // string | 
    includeRaw := true // bool |  (optional) (default to false)

    resp, r, err := sdk.EventsApi.GetOrganizationEvent(context.Background(), orgId, eventId).IncludeRaw(includeRaw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetOrganizationEvent``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetOrganizationEvent`: EventViewForOrg
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetOrganizationEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**eventId** | **string** | Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listOrganizationEvents) endpoint to retrieve all events to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]

### Return type

[**EventViewForOrg**](EventViewForOrg.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectEvent

> EventViewForNdsGroup GetProjectEvent(ctx, groupId, eventId).IncludeRaw(includeRaw).Execute()

Return One Event from One Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    eventId := "eventId_example" // string | 
    includeRaw := true // bool |  (optional) (default to false)

    resp, r, err := sdk.EventsApi.GetProjectEvent(context.Background(), groupId, eventId).IncludeRaw(includeRaw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetProjectEvent``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetProjectEvent`: EventViewForNdsGroup
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetProjectEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**eventId** | **string** | Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listProjectEvents) endpoint to retrieve all events to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]

### Return type

[**EventViewForNdsGroup**](EventViewForNdsGroup.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationEvents

> OrgPaginatedEvent ListOrganizationEvents(ctx, orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()

Return All Events from One Organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    orgId := "4888442a3354817a7320eb61" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    eventType := openapiclient.EventTypeForOrg("ALERT_ACKNOWLEDGED_AUDIT") // EventTypeForOrg |  (optional)
    includeRaw := true // bool |  (optional) (default to false)
    maxDate := time.Now() // time.Time |  (optional)
    minDate := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.EventsApi.ListOrganizationEvents(context.Background(), orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListOrganizationEvents``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListOrganizationEvents`: OrgPaginatedEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListOrganizationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **eventType** | [**EventTypeForOrg**](EventTypeForOrg.md) | Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently. | 
 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]
 **maxDate** | **time.Time** | Date and time from when MongoDB Cloud stops returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | 
 **minDate** | **time.Time** | Date and time from when MongoDB Cloud starts returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | 

### Return type

[**OrgPaginatedEvent**](OrgPaginatedEvent.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectEvents

> GroupPaginatedEvent ListProjectEvents(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ClusterNames(clusterNames).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()

Return All Events from One Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    clusterNames := []string{"Inner_example"} // []string |  (optional)
    eventType := openapiclient.EventTypeForNdsGroup("ALERT_ACKNOWLEDGED_AUDIT") // EventTypeForNdsGroup |  (optional)
    includeRaw := true // bool |  (optional) (default to false)
    maxDate := time.Now() // time.Time |  (optional)
    minDate := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.EventsApi.ListProjectEvents(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ClusterNames(clusterNames).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListProjectEvents``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListProjectEvents`: GroupPaginatedEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListProjectEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **clusterNames** | **[]string** | Human-readable label that identifies the cluster. | 
 **eventType** | [**EventTypeForNdsGroup**](EventTypeForNdsGroup.md) | Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently. | 
 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]
 **maxDate** | **time.Time** | Date and time from when MongoDB Cloud stops returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | 
 **minDate** | **time.Time** | Date and time from when MongoDB Cloud starts returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | 

### Return type

[**GroupPaginatedEvent**](GroupPaginatedEvent.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

