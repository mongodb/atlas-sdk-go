# DiskBackupCopyPolicyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrequencyType** | **string** | Human-readable label that identifies the frequency type associated with the copy policy. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this copy policy item. | [optional] [readonly] 
**RetentionUnit** | Pointer to **string** | Unit of time in which MongoDB Cloud measures snapshot copy retention. | [optional] 
**RetentionValue** | Pointer to **int** | Duration in days, weeks, months, or years that MongoDB Cloud retains the snapshot copy. | [optional] 

## Methods

### NewDiskBackupCopyPolicyItem

`func NewDiskBackupCopyPolicyItem(frequencyType string, ) *DiskBackupCopyPolicyItem`

NewDiskBackupCopyPolicyItem instantiates a new DiskBackupCopyPolicyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupCopyPolicyItemWithDefaults

`func NewDiskBackupCopyPolicyItemWithDefaults() *DiskBackupCopyPolicyItem`

NewDiskBackupCopyPolicyItemWithDefaults instantiates a new DiskBackupCopyPolicyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequencyType

`func (o *DiskBackupCopyPolicyItem) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *DiskBackupCopyPolicyItem) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *DiskBackupCopyPolicyItem) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### GetId

`func (o *DiskBackupCopyPolicyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupCopyPolicyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupCopyPolicyItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskBackupCopyPolicyItem) HasId() bool`

HasId returns a boolean if a field has been set.
### GetRetentionUnit

`func (o *DiskBackupCopyPolicyItem) GetRetentionUnit() string`

GetRetentionUnit returns the RetentionUnit field if non-nil, zero value otherwise.

### GetRetentionUnitOk

`func (o *DiskBackupCopyPolicyItem) GetRetentionUnitOk() (*string, bool)`

GetRetentionUnitOk returns a tuple with the RetentionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionUnit

`func (o *DiskBackupCopyPolicyItem) SetRetentionUnit(v string)`

SetRetentionUnit sets RetentionUnit field to given value.

### HasRetentionUnit

`func (o *DiskBackupCopyPolicyItem) HasRetentionUnit() bool`

HasRetentionUnit returns a boolean if a field has been set.
### GetRetentionValue

`func (o *DiskBackupCopyPolicyItem) GetRetentionValue() int`

GetRetentionValue returns the RetentionValue field if non-nil, zero value otherwise.

### GetRetentionValueOk

`func (o *DiskBackupCopyPolicyItem) GetRetentionValueOk() (*int, bool)`

GetRetentionValueOk returns a tuple with the RetentionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionValue

`func (o *DiskBackupCopyPolicyItem) SetRetentionValue(v int)`

SetRetentionValue sets RetentionValue field to given value.

### HasRetentionValue

`func (o *DiskBackupCopyPolicyItem) HasRetentionValue() bool`

HasRetentionValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


