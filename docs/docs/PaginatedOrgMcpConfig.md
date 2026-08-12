# PaginatedOrgMcpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | [**[]OrgMcpConfigResponse**](OrgMcpConfigResponse.md) | List of returned documents that MongoDB Cloud provides when completing this request. | [readonly] 
**TotalCount** | Pointer to **int** | Total number of documents available. MongoDB Cloud omits this value if &#x60;includeCount&#x60; is set to &#x60;false&#x60;. The total number is an estimate and may not be exact. | [optional] [readonly] 

## Methods

### NewPaginatedOrgMcpConfig

`func NewPaginatedOrgMcpConfig(results []OrgMcpConfigResponse, ) *PaginatedOrgMcpConfig`

NewPaginatedOrgMcpConfig instantiates a new PaginatedOrgMcpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOrgMcpConfigWithDefaults

`func NewPaginatedOrgMcpConfigWithDefaults() *PaginatedOrgMcpConfig`

NewPaginatedOrgMcpConfigWithDefaults instantiates a new PaginatedOrgMcpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedOrgMcpConfig) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedOrgMcpConfig) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedOrgMcpConfig) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedOrgMcpConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *PaginatedOrgMcpConfig) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetResults

`func (o *PaginatedOrgMcpConfig) GetResults() []OrgMcpConfigResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedOrgMcpConfig) GetResultsOk() (*[]OrgMcpConfigResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedOrgMcpConfig) SetResults(v []OrgMcpConfigResponse)`

SetResults sets Results field to given value.

### GetTotalCount

`func (o *PaginatedOrgMcpConfig) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedOrgMcpConfig) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedOrgMcpConfig) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedOrgMcpConfig) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *PaginatedOrgMcpConfig) SetTotalCountNil()`

SetTotalCountNil sets TotalCount to an explicit JSON null when marshaled, overriding any value previously set with SetTotalCount. Calling SetTotalCount again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


