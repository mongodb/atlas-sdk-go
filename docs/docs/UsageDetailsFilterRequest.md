# UsageDetailsFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillEndDate** | Pointer to **string** | The inclusive billing start date for usage details filter. | [optional] 
**BillStartDate** | Pointer to **string** | The inclusive billing start date for usage details filter. | [optional] 
**ClusterIds** | Pointer to **[]string** | The list of unique cluster ids to be included in the Usage Details filter. | [optional] 
**GroupIds** | Pointer to **[]string** | The list of groups to be included in the Usage Details filter. | [optional] 
**IncludeZeroCentLineItems** | Pointer to **bool** | Whether zero cent line items should be included. | [optional] 
**SkuServices** | Pointer to **[]string** | The list of projects to be included in the Cost Explorer Query. | [optional] 
**UsageEndDate** | Pointer to **string** | The inclusive billing start date for usage details filter. | [optional] 
**UsageStartDate** | Pointer to **string** | The inclusive usage start date for usage details filter. | [optional] 

## Methods

### NewUsageDetailsFilterRequest

`func NewUsageDetailsFilterRequest() *UsageDetailsFilterRequest`

NewUsageDetailsFilterRequest instantiates a new UsageDetailsFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDetailsFilterRequestWithDefaults

`func NewUsageDetailsFilterRequestWithDefaults() *UsageDetailsFilterRequest`

NewUsageDetailsFilterRequestWithDefaults instantiates a new UsageDetailsFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillEndDate

`func (o *UsageDetailsFilterRequest) GetBillEndDate() string`

GetBillEndDate returns the BillEndDate field if non-nil, zero value otherwise.

### GetBillEndDateOk

`func (o *UsageDetailsFilterRequest) GetBillEndDateOk() (*string, bool)`

GetBillEndDateOk returns a tuple with the BillEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillEndDate

`func (o *UsageDetailsFilterRequest) SetBillEndDate(v string)`

SetBillEndDate sets BillEndDate field to given value.

### HasBillEndDate

`func (o *UsageDetailsFilterRequest) HasBillEndDate() bool`

HasBillEndDate returns a boolean if a field has been set.

### SetBillEndDateNil

`func (o *UsageDetailsFilterRequest) SetBillEndDateNil()`

SetBillEndDateNil sets BillEndDate to an explicit JSON null when marshaled, overriding any value previously set with SetBillEndDate. Calling SetBillEndDate again clears the null override.

### GetBillStartDate

`func (o *UsageDetailsFilterRequest) GetBillStartDate() string`

GetBillStartDate returns the BillStartDate field if non-nil, zero value otherwise.

### GetBillStartDateOk

`func (o *UsageDetailsFilterRequest) GetBillStartDateOk() (*string, bool)`

GetBillStartDateOk returns a tuple with the BillStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillStartDate

`func (o *UsageDetailsFilterRequest) SetBillStartDate(v string)`

SetBillStartDate sets BillStartDate field to given value.

### HasBillStartDate

`func (o *UsageDetailsFilterRequest) HasBillStartDate() bool`

HasBillStartDate returns a boolean if a field has been set.

### SetBillStartDateNil

`func (o *UsageDetailsFilterRequest) SetBillStartDateNil()`

SetBillStartDateNil sets BillStartDate to an explicit JSON null when marshaled, overriding any value previously set with SetBillStartDate. Calling SetBillStartDate again clears the null override.

### GetClusterIds

`func (o *UsageDetailsFilterRequest) GetClusterIds() []string`

GetClusterIds returns the ClusterIds field if non-nil, zero value otherwise.

### GetClusterIdsOk

`func (o *UsageDetailsFilterRequest) GetClusterIdsOk() (*[]string, bool)`

GetClusterIdsOk returns a tuple with the ClusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIds

`func (o *UsageDetailsFilterRequest) SetClusterIds(v []string)`

SetClusterIds sets ClusterIds field to given value.

### HasClusterIds

`func (o *UsageDetailsFilterRequest) HasClusterIds() bool`

HasClusterIds returns a boolean if a field has been set.

### SetClusterIdsNil

`func (o *UsageDetailsFilterRequest) SetClusterIdsNil()`

SetClusterIdsNil sets ClusterIds to an explicit JSON null when marshaled, overriding any value previously set with SetClusterIds. Calling SetClusterIds again clears the null override.

### GetGroupIds

`func (o *UsageDetailsFilterRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *UsageDetailsFilterRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *UsageDetailsFilterRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *UsageDetailsFilterRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### SetGroupIdsNil

`func (o *UsageDetailsFilterRequest) SetGroupIdsNil()`

SetGroupIdsNil sets GroupIds to an explicit JSON null when marshaled, overriding any value previously set with SetGroupIds. Calling SetGroupIds again clears the null override.

### GetIncludeZeroCentLineItems

`func (o *UsageDetailsFilterRequest) GetIncludeZeroCentLineItems() bool`

GetIncludeZeroCentLineItems returns the IncludeZeroCentLineItems field if non-nil, zero value otherwise.

### GetIncludeZeroCentLineItemsOk

`func (o *UsageDetailsFilterRequest) GetIncludeZeroCentLineItemsOk() (*bool, bool)`

GetIncludeZeroCentLineItemsOk returns a tuple with the IncludeZeroCentLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZeroCentLineItems

`func (o *UsageDetailsFilterRequest) SetIncludeZeroCentLineItems(v bool)`

SetIncludeZeroCentLineItems sets IncludeZeroCentLineItems field to given value.

### HasIncludeZeroCentLineItems

`func (o *UsageDetailsFilterRequest) HasIncludeZeroCentLineItems() bool`

HasIncludeZeroCentLineItems returns a boolean if a field has been set.

### SetIncludeZeroCentLineItemsNil

`func (o *UsageDetailsFilterRequest) SetIncludeZeroCentLineItemsNil()`

SetIncludeZeroCentLineItemsNil sets IncludeZeroCentLineItems to an explicit JSON null when marshaled, overriding any value previously set with SetIncludeZeroCentLineItems. Calling SetIncludeZeroCentLineItems again clears the null override.

### GetSkuServices

`func (o *UsageDetailsFilterRequest) GetSkuServices() []string`

GetSkuServices returns the SkuServices field if non-nil, zero value otherwise.

### GetSkuServicesOk

`func (o *UsageDetailsFilterRequest) GetSkuServicesOk() (*[]string, bool)`

GetSkuServicesOk returns a tuple with the SkuServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuServices

`func (o *UsageDetailsFilterRequest) SetSkuServices(v []string)`

SetSkuServices sets SkuServices field to given value.

### HasSkuServices

`func (o *UsageDetailsFilterRequest) HasSkuServices() bool`

HasSkuServices returns a boolean if a field has been set.

### SetSkuServicesNil

`func (o *UsageDetailsFilterRequest) SetSkuServicesNil()`

SetSkuServicesNil sets SkuServices to an explicit JSON null when marshaled, overriding any value previously set with SetSkuServices. Calling SetSkuServices again clears the null override.

### GetUsageEndDate

`func (o *UsageDetailsFilterRequest) GetUsageEndDate() string`

GetUsageEndDate returns the UsageEndDate field if non-nil, zero value otherwise.

### GetUsageEndDateOk

`func (o *UsageDetailsFilterRequest) GetUsageEndDateOk() (*string, bool)`

GetUsageEndDateOk returns a tuple with the UsageEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEndDate

`func (o *UsageDetailsFilterRequest) SetUsageEndDate(v string)`

SetUsageEndDate sets UsageEndDate field to given value.

### HasUsageEndDate

`func (o *UsageDetailsFilterRequest) HasUsageEndDate() bool`

HasUsageEndDate returns a boolean if a field has been set.

### SetUsageEndDateNil

`func (o *UsageDetailsFilterRequest) SetUsageEndDateNil()`

SetUsageEndDateNil sets UsageEndDate to an explicit JSON null when marshaled, overriding any value previously set with SetUsageEndDate. Calling SetUsageEndDate again clears the null override.

### GetUsageStartDate

`func (o *UsageDetailsFilterRequest) GetUsageStartDate() string`

GetUsageStartDate returns the UsageStartDate field if non-nil, zero value otherwise.

### GetUsageStartDateOk

`func (o *UsageDetailsFilterRequest) GetUsageStartDateOk() (*string, bool)`

GetUsageStartDateOk returns a tuple with the UsageStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStartDate

`func (o *UsageDetailsFilterRequest) SetUsageStartDate(v string)`

SetUsageStartDate sets UsageStartDate field to given value.

### HasUsageStartDate

`func (o *UsageDetailsFilterRequest) HasUsageStartDate() bool`

HasUsageStartDate returns a boolean if a field has been set.

### SetUsageStartDateNil

`func (o *UsageDetailsFilterRequest) SetUsageStartDateNil()`

SetUsageStartDateNil sets UsageStartDate to an explicit JSON null when marshaled, overriding any value previously set with SetUsageStartDate. Calling SetUsageStartDate again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


