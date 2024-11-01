# PaginatedServiceAccountIPAccessEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | Pointer to [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | List of returned documents that MongoDB Cloud provides when completing this request. | [optional] [readonly] 
**TotalCount** | Pointer to **int** | Total number of documents available. MongoDB Cloud omits this value if &#x60;includeCount&#x60; is set to &#x60;false&#x60;. | [optional] [readonly] 

## Methods

### NewPaginatedServiceAccountIPAccessEntry

`func NewPaginatedServiceAccountIPAccessEntry() *PaginatedServiceAccountIPAccessEntry`

NewPaginatedServiceAccountIPAccessEntry instantiates a new PaginatedServiceAccountIPAccessEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedServiceAccountIPAccessEntryWithDefaults

`func NewPaginatedServiceAccountIPAccessEntryWithDefaults() *PaginatedServiceAccountIPAccessEntry`

NewPaginatedServiceAccountIPAccessEntryWithDefaults instantiates a new PaginatedServiceAccountIPAccessEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedServiceAccountIPAccessEntry) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedServiceAccountIPAccessEntry) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedServiceAccountIPAccessEntry) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedServiceAccountIPAccessEntry) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetResults

`func (o *PaginatedServiceAccountIPAccessEntry) GetResults() []ServiceAccountIPAccessListEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedServiceAccountIPAccessEntry) GetResultsOk() (*[]ServiceAccountIPAccessListEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedServiceAccountIPAccessEntry) SetResults(v []ServiceAccountIPAccessListEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedServiceAccountIPAccessEntry) HasResults() bool`

HasResults returns a boolean if a field has been set.
### GetTotalCount

`func (o *PaginatedServiceAccountIPAccessEntry) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedServiceAccountIPAccessEntry) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedServiceAccountIPAccessEntry) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedServiceAccountIPAccessEntry) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

