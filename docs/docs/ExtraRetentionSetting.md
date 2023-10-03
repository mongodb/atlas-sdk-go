# ExtraRetentionSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrequencyType** | Pointer to **string** | The frequency type for the extra retention settings for the cluster. | [optional] 
**RetentionDays** | Pointer to **int** | The number of extra retention days for the cluster. | [optional] 

## Methods

### NewExtraRetentionSetting

`func NewExtraRetentionSetting() *ExtraRetentionSetting`

NewExtraRetentionSetting instantiates a new ExtraRetentionSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraRetentionSettingWithDefaults

`func NewExtraRetentionSettingWithDefaults() *ExtraRetentionSetting`

NewExtraRetentionSettingWithDefaults instantiates a new ExtraRetentionSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequencyType

`func (o *ExtraRetentionSetting) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *ExtraRetentionSetting) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *ExtraRetentionSetting) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### HasFrequencyType

`func (o *ExtraRetentionSetting) HasFrequencyType() bool`

HasFrequencyType returns a boolean if a field has been set.
### GetRetentionDays

`func (o *ExtraRetentionSetting) GetRetentionDays() int`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *ExtraRetentionSetting) GetRetentionDaysOk() (*int, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *ExtraRetentionSetting) SetRetentionDays(v int)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *ExtraRetentionSetting) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


