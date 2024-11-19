# FlexBackupSettings20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether backups are performed for this flex cluster. Backup uses flex cluster backups. | [optional] [readonly] [default to true]

## Methods

### NewFlexBackupSettings20241113

`func NewFlexBackupSettings20241113() *FlexBackupSettings20241113`

NewFlexBackupSettings20241113 instantiates a new FlexBackupSettings20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexBackupSettings20241113WithDefaults

`func NewFlexBackupSettings20241113WithDefaults() *FlexBackupSettings20241113`

NewFlexBackupSettings20241113WithDefaults instantiates a new FlexBackupSettings20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *FlexBackupSettings20241113) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FlexBackupSettings20241113) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FlexBackupSettings20241113) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FlexBackupSettings20241113) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


