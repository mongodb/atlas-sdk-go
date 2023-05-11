# GroupPaginatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | Pointer to [**[]EventViewForNdsGroup**](EventViewForNdsGroup.md) | List of returned documents that MongoDB Cloud providers when completing this request. | [optional] [readonly] 
**TotalCount** | Pointer to **int** | Number of documents returned in this response if **includeCount** query param is true. | [optional] [readonly] 

## Methods

### NewGroupPaginatedEvent

`func NewGroupPaginatedEvent() *GroupPaginatedEvent`

NewGroupPaginatedEvent instantiates a new GroupPaginatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPaginatedEventWithDefaults

`func NewGroupPaginatedEventWithDefaults() *GroupPaginatedEvent`

NewGroupPaginatedEventWithDefaults instantiates a new GroupPaginatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GroupPaginatedEvent) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupPaginatedEvent) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupPaginatedEvent) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupPaginatedEvent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *GroupPaginatedEvent) GetResults() []EventViewForNdsGroup`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GroupPaginatedEvent) GetResultsOk() (*[]EventViewForNdsGroup, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GroupPaginatedEvent) SetResults(v []EventViewForNdsGroup)`

SetResults sets Results field to given value.

### HasResults

`func (o *GroupPaginatedEvent) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *GroupPaginatedEvent) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GroupPaginatedEvent) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GroupPaginatedEvent) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GroupPaginatedEvent) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


