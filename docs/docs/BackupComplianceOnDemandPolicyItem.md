# BackupComplianceOnDemandPolicyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrequencyInterval** | **int** | Number that indicates the frequency interval for a set of snapshots. MongoDB Cloud ignores this setting for non-hourly policy items in Backup Compliance Policy settings. | 
**FrequencyType** | **string** | Human-readable label that identifies the frequency type associated with the backup policy. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this backup policy item. | [optional] [readonly] 
**RetentionUnit** | **string** | Unit of time in which MongoDB Cloud measures snapshot retention. | 
**RetentionValue** | **int** | Duration in days, weeks, or months that MongoDB Cloud retains the snapshot. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items.  For example: If the hourly policy item specifies a retention of two days, you must specify two days or greater for the retention of the weekly policy item. | 

## Methods

### NewBackupComplianceOnDemandPolicyItem

`func NewBackupComplianceOnDemandPolicyItem(frequencyInterval int, frequencyType string, retentionUnit string, retentionValue int, ) *BackupComplianceOnDemandPolicyItem`

NewBackupComplianceOnDemandPolicyItem instantiates a new BackupComplianceOnDemandPolicyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupComplianceOnDemandPolicyItemWithDefaults

`func NewBackupComplianceOnDemandPolicyItemWithDefaults() *BackupComplianceOnDemandPolicyItem`

NewBackupComplianceOnDemandPolicyItemWithDefaults instantiates a new BackupComplianceOnDemandPolicyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequencyInterval

`func (o *BackupComplianceOnDemandPolicyItem) GetFrequencyInterval() int`

GetFrequencyInterval returns the FrequencyInterval field if non-nil, zero value otherwise.

### GetFrequencyIntervalOk

`func (o *BackupComplianceOnDemandPolicyItem) GetFrequencyIntervalOk() (*int, bool)`

GetFrequencyIntervalOk returns a tuple with the FrequencyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyInterval

`func (o *BackupComplianceOnDemandPolicyItem) SetFrequencyInterval(v int)`

SetFrequencyInterval sets FrequencyInterval field to given value.

### GetFrequencyType

`func (o *BackupComplianceOnDemandPolicyItem) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *BackupComplianceOnDemandPolicyItem) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *BackupComplianceOnDemandPolicyItem) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### GetId

`func (o *BackupComplianceOnDemandPolicyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupComplianceOnDemandPolicyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupComplianceOnDemandPolicyItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupComplianceOnDemandPolicyItem) HasId() bool`

HasId returns a boolean if a field has been set.
### GetRetentionUnit

`func (o *BackupComplianceOnDemandPolicyItem) GetRetentionUnit() string`

GetRetentionUnit returns the RetentionUnit field if non-nil, zero value otherwise.

### GetRetentionUnitOk

`func (o *BackupComplianceOnDemandPolicyItem) GetRetentionUnitOk() (*string, bool)`

GetRetentionUnitOk returns a tuple with the RetentionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionUnit

`func (o *BackupComplianceOnDemandPolicyItem) SetRetentionUnit(v string)`

SetRetentionUnit sets RetentionUnit field to given value.

### GetRetentionValue

`func (o *BackupComplianceOnDemandPolicyItem) GetRetentionValue() int`

GetRetentionValue returns the RetentionValue field if non-nil, zero value otherwise.

### GetRetentionValueOk

`func (o *BackupComplianceOnDemandPolicyItem) GetRetentionValueOk() (*int, bool)`

GetRetentionValueOk returns a tuple with the RetentionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionValue

`func (o *BackupComplianceOnDemandPolicyItem) SetRetentionValue(v int)`

SetRetentionValue sets RetentionValue field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


