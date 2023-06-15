# PaginatedAlertConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | Pointer to [**[]GroupAlertsConfig**](GroupAlertsConfig.md) | List of returned documents that MongoDB Cloud providers when completing this request. | [optional] [readonly] 
**TotalCount** | Pointer to **int** | Number of documents returned in this response if **includeCount** query param is true. | [optional] [readonly] 

## Methods

### NewPaginatedAlertConfig

`func NewPaginatedAlertConfig() *PaginatedAlertConfig`

NewPaginatedAlertConfig instantiates a new PaginatedAlertConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAlertConfigWithDefaults

`func NewPaginatedAlertConfigWithDefaults() *PaginatedAlertConfig`

NewPaginatedAlertConfigWithDefaults instantiates a new PaginatedAlertConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedAlertConfig) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedAlertConfig) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedAlertConfig) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedAlertConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedAlertConfig) GetResults() []GroupAlertsConfig`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAlertConfig) GetResultsOk() (*[]GroupAlertsConfig, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAlertConfig) SetResults(v []GroupAlertsConfig)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedAlertConfig) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginatedAlertConfig) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedAlertConfig) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedAlertConfig) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedAlertConfig) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


