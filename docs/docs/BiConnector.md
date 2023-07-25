# BiConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether MongoDB Connector for Business Intelligence is enabled on the specified cluster. | [optional] 
**ReadPreference** | Pointer to **string** | Data source node designated for the MongoDB Connector for Business Intelligence on MongoDB Cloud. The MongoDB Connector for Business Intelligence on MongoDB Cloud reads data from the primary, secondary, or analytics node based on your read preferences. Defaults to &#x60;ANALYTICS&#x60; node, or &#x60;SECONDARY&#x60; if there are no &#x60;ANALYTICS&#x60; nodes. | [optional] 

## Methods

### NewBiConnector

`func NewBiConnector() *BiConnector`

NewBiConnector instantiates a new BiConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiConnectorWithDefaults

`func NewBiConnectorWithDefaults() *BiConnector`

NewBiConnectorWithDefaults instantiates a new BiConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BiConnector) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BiConnector) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BiConnector) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BiConnector) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.
### GetReadPreference

`func (o *BiConnector) GetReadPreference() string`

GetReadPreference returns the ReadPreference field if non-nil, zero value otherwise.

### GetReadPreferenceOk

`func (o *BiConnector) GetReadPreferenceOk() (*string, bool)`

GetReadPreferenceOk returns a tuple with the ReadPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPreference

`func (o *BiConnector) SetReadPreference(v string)`

SetReadPreference sets ReadPreference field to given value.

### HasReadPreference

`func (o *BiConnector) HasReadPreference() bool`

HasReadPreference returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


