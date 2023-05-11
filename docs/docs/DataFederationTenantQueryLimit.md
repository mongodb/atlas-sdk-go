# DataFederationTenantQueryLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsage** | Pointer to **int64** | Amount that indicates the current usage of the limit. | [optional] [readonly] 
**DefaultLimit** | Pointer to **int64** | Default value of the limit. | [optional] [readonly] 
**LastModifiedDate** | Pointer to **time.Time** | Only used for Data Federation limits. Timestamp that indicates when this usage limit was last modified. This field uses the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**MaximumLimit** | Pointer to **int64** | Maximum value of the limit. | [optional] [readonly] 
**Name** | **string** | Human-readable label that identifies the user-managed limit to modify. | [readonly] 
**OverrunPolicy** | Pointer to **string** | Only used for Data Federation limits. Action to take when the usage limit is exceeded. If limit span is set to QUERY, this is ignored because MongoDB Cloud stops the query when it exceeds the usage limit. | [optional] 
**TenantName** | Pointer to **string** | Human-readable label that identifies the Federated Database Instance. If specified, the usage limit is for the specified federated database instance only. If omitted, the usage limit is for all federated database instances in the project. | [optional] [readonly] 
**Value** | **int64** | Amount to set the limit to. | 

## Methods

### NewDataFederationTenantQueryLimit

`func NewDataFederationTenantQueryLimit(name string, value int64, ) *DataFederationTenantQueryLimit`

NewDataFederationTenantQueryLimit instantiates a new DataFederationTenantQueryLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataFederationTenantQueryLimitWithDefaults

`func NewDataFederationTenantQueryLimitWithDefaults() *DataFederationTenantQueryLimit`

NewDataFederationTenantQueryLimitWithDefaults instantiates a new DataFederationTenantQueryLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsage

`func (o *DataFederationTenantQueryLimit) GetCurrentUsage() int64`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *DataFederationTenantQueryLimit) GetCurrentUsageOk() (*int64, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *DataFederationTenantQueryLimit) SetCurrentUsage(v int64)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *DataFederationTenantQueryLimit) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.

### GetDefaultLimit

`func (o *DataFederationTenantQueryLimit) GetDefaultLimit() int64`

GetDefaultLimit returns the DefaultLimit field if non-nil, zero value otherwise.

### GetDefaultLimitOk

`func (o *DataFederationTenantQueryLimit) GetDefaultLimitOk() (*int64, bool)`

GetDefaultLimitOk returns a tuple with the DefaultLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLimit

`func (o *DataFederationTenantQueryLimit) SetDefaultLimit(v int64)`

SetDefaultLimit sets DefaultLimit field to given value.

### HasDefaultLimit

`func (o *DataFederationTenantQueryLimit) HasDefaultLimit() bool`

HasDefaultLimit returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *DataFederationTenantQueryLimit) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *DataFederationTenantQueryLimit) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *DataFederationTenantQueryLimit) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *DataFederationTenantQueryLimit) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetMaximumLimit

`func (o *DataFederationTenantQueryLimit) GetMaximumLimit() int64`

GetMaximumLimit returns the MaximumLimit field if non-nil, zero value otherwise.

### GetMaximumLimitOk

`func (o *DataFederationTenantQueryLimit) GetMaximumLimitOk() (*int64, bool)`

GetMaximumLimitOk returns a tuple with the MaximumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLimit

`func (o *DataFederationTenantQueryLimit) SetMaximumLimit(v int64)`

SetMaximumLimit sets MaximumLimit field to given value.

### HasMaximumLimit

`func (o *DataFederationTenantQueryLimit) HasMaximumLimit() bool`

HasMaximumLimit returns a boolean if a field has been set.

### GetName

`func (o *DataFederationTenantQueryLimit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataFederationTenantQueryLimit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataFederationTenantQueryLimit) SetName(v string)`

SetName sets Name field to given value.


### GetOverrunPolicy

`func (o *DataFederationTenantQueryLimit) GetOverrunPolicy() string`

GetOverrunPolicy returns the OverrunPolicy field if non-nil, zero value otherwise.

### GetOverrunPolicyOk

`func (o *DataFederationTenantQueryLimit) GetOverrunPolicyOk() (*string, bool)`

GetOverrunPolicyOk returns a tuple with the OverrunPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrunPolicy

`func (o *DataFederationTenantQueryLimit) SetOverrunPolicy(v string)`

SetOverrunPolicy sets OverrunPolicy field to given value.

### HasOverrunPolicy

`func (o *DataFederationTenantQueryLimit) HasOverrunPolicy() bool`

HasOverrunPolicy returns a boolean if a field has been set.

### GetTenantName

`func (o *DataFederationTenantQueryLimit) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *DataFederationTenantQueryLimit) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *DataFederationTenantQueryLimit) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *DataFederationTenantQueryLimit) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetValue

`func (o *DataFederationTenantQueryLimit) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataFederationTenantQueryLimit) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataFederationTenantQueryLimit) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


