# ApiPublicUsageDetailsQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**UsageDetailsFilterRequest**](UsageDetailsFilterRequest.md) |  | [optional] 
**SortField** | Pointer to **string** | Specify the field used to specify how to sort query results. Default to bill date. | [optional] 
**SortOrder** | Pointer to **string** | Specify the sort order (ascending / descending) used to specify how to sort query results. Defaults to descending. | [optional] 

## Methods

### NewApiPublicUsageDetailsQueryRequest

`func NewApiPublicUsageDetailsQueryRequest() *ApiPublicUsageDetailsQueryRequest`

NewApiPublicUsageDetailsQueryRequest instantiates a new ApiPublicUsageDetailsQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPublicUsageDetailsQueryRequestWithDefaults

`func NewApiPublicUsageDetailsQueryRequestWithDefaults() *ApiPublicUsageDetailsQueryRequest`

NewApiPublicUsageDetailsQueryRequestWithDefaults instantiates a new ApiPublicUsageDetailsQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ApiPublicUsageDetailsQueryRequest) GetFilters() UsageDetailsFilterRequest`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiPublicUsageDetailsQueryRequest) GetFiltersOk() (*UsageDetailsFilterRequest, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiPublicUsageDetailsQueryRequest) SetFilters(v UsageDetailsFilterRequest)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ApiPublicUsageDetailsQueryRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.
### GetSortField

`func (o *ApiPublicUsageDetailsQueryRequest) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *ApiPublicUsageDetailsQueryRequest) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *ApiPublicUsageDetailsQueryRequest) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *ApiPublicUsageDetailsQueryRequest) HasSortField() bool`

HasSortField returns a boolean if a field has been set.
### GetSortOrder

`func (o *ApiPublicUsageDetailsQueryRequest) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ApiPublicUsageDetailsQueryRequest) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ApiPublicUsageDetailsQueryRequest) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ApiPublicUsageDetailsQueryRequest) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


