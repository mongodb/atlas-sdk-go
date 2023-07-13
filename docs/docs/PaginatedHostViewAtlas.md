# PaginatedHostViewAtlas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]LinkAtlas**](LinkAtlas.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | Pointer to [**[]ApiHostViewAtlas**](ApiHostViewAtlas.md) | List of returned documents that MongoDB Cloud providers when completing this request. | [optional] [readonly] 
**TotalCount** | Pointer to **int** | Number of documents returned in this response. | [optional] [readonly] 

## Methods

### NewPaginatedHostViewAtlas

`func NewPaginatedHostViewAtlas() *PaginatedHostViewAtlas`

NewPaginatedHostViewAtlas instantiates a new PaginatedHostViewAtlas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedHostViewAtlasWithDefaults

`func NewPaginatedHostViewAtlasWithDefaults() *PaginatedHostViewAtlas`

NewPaginatedHostViewAtlasWithDefaults instantiates a new PaginatedHostViewAtlas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedHostViewAtlas) GetLinks() []LinkAtlas`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedHostViewAtlas) GetLinksOk() (*[]LinkAtlas, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedHostViewAtlas) SetLinks(v []LinkAtlas)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedHostViewAtlas) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetResults

`func (o *PaginatedHostViewAtlas) GetResults() []ApiHostViewAtlas`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedHostViewAtlas) GetResultsOk() (*[]ApiHostViewAtlas, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedHostViewAtlas) SetResults(v []ApiHostViewAtlas)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedHostViewAtlas) HasResults() bool`

HasResults returns a boolean if a field has been set.
### GetTotalCount

`func (o *PaginatedHostViewAtlas) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedHostViewAtlas) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedHostViewAtlas) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedHostViewAtlas) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


