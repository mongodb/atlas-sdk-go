# DataFederationLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsage** | Pointer to **int64** | Amount that indicates the current usage of the limit. | [optional] [readonly] 
**DefaultLimit** | Pointer to **int64** | Default value of the limit. | [optional] [readonly] 
**MaximumLimit** | Pointer to **int64** | Maximum value of the limit. | [optional] [readonly] 
**Name** | **string** | Human-readable label that identifies the user-managed limit to modify. | [readonly] 
**Value** | **int64** | Amount to set the limit to. | 

## Methods

### NewDataFederationLimit

`func NewDataFederationLimit(name string, value int64, ) *DataFederationLimit`

NewDataFederationLimit instantiates a new DataFederationLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataFederationLimitWithDefaults

`func NewDataFederationLimitWithDefaults() *DataFederationLimit`

NewDataFederationLimitWithDefaults instantiates a new DataFederationLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsage

`func (o *DataFederationLimit) GetCurrentUsage() int64`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *DataFederationLimit) GetCurrentUsageOk() (*int64, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *DataFederationLimit) SetCurrentUsage(v int64)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *DataFederationLimit) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.
### GetDefaultLimit

`func (o *DataFederationLimit) GetDefaultLimit() int64`

GetDefaultLimit returns the DefaultLimit field if non-nil, zero value otherwise.

### GetDefaultLimitOk

`func (o *DataFederationLimit) GetDefaultLimitOk() (*int64, bool)`

GetDefaultLimitOk returns a tuple with the DefaultLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLimit

`func (o *DataFederationLimit) SetDefaultLimit(v int64)`

SetDefaultLimit sets DefaultLimit field to given value.

### HasDefaultLimit

`func (o *DataFederationLimit) HasDefaultLimit() bool`

HasDefaultLimit returns a boolean if a field has been set.
### GetMaximumLimit

`func (o *DataFederationLimit) GetMaximumLimit() int64`

GetMaximumLimit returns the MaximumLimit field if non-nil, zero value otherwise.

### GetMaximumLimitOk

`func (o *DataFederationLimit) GetMaximumLimitOk() (*int64, bool)`

GetMaximumLimitOk returns a tuple with the MaximumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLimit

`func (o *DataFederationLimit) SetMaximumLimit(v int64)`

SetMaximumLimit sets MaximumLimit field to given value.

### HasMaximumLimit

`func (o *DataFederationLimit) HasMaximumLimit() bool`

HasMaximumLimit returns a boolean if a field has been set.
### GetName

`func (o *DataFederationLimit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataFederationLimit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataFederationLimit) SetName(v string)`

SetName sets Name field to given value.

### GetValue

`func (o *DataFederationLimit) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataFederationLimit) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataFederationLimit) SetValue(v int64)`

SetValue sets Value field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


