# PaginatedOrgUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | Pointer to [**[]OrgUserResponse**](OrgUserResponse.md) | List of returned documents that MongoDB Cloud provides when completing this request. | [optional] [readonly] 
**TotalCount** | Pointer to **int** | Total number of documents available. MongoDB Cloud omits this value if &#x60;includeCount&#x60; is set to &#x60;false&#x60;. The total number is an estimate and may not be exact. | [optional] [readonly] 

## Methods

### NewPaginatedOrgUser

`func NewPaginatedOrgUser() *PaginatedOrgUser`

NewPaginatedOrgUser instantiates a new PaginatedOrgUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOrgUserWithDefaults

`func NewPaginatedOrgUserWithDefaults() *PaginatedOrgUser`

NewPaginatedOrgUserWithDefaults instantiates a new PaginatedOrgUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedOrgUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedOrgUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedOrgUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedOrgUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetResults

`func (o *PaginatedOrgUser) GetResults() []OrgUserResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedOrgUser) GetResultsOk() (*[]OrgUserResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedOrgUser) SetResults(v []OrgUserResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedOrgUser) HasResults() bool`

HasResults returns a boolean if a field has been set.
### GetTotalCount

`func (o *PaginatedOrgUser) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedOrgUser) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedOrgUser) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedOrgUser) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


