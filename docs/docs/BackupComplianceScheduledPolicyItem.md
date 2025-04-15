# BackupComplianceScheduledPolicyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrequencyInterval** | **int** | Number that indicates the frequency interval for a set of Snapshots. A value of &#x60;1&#x60; specifies the first instance of the corresponding &#x60;frequencyType&#x60;.  - In a yearly policy item, &#x60;1&#x60; indicates that the yearly Snapshot occurs on the first day of January and &#x60;12&#x60; indicates the first day of December.  - In a monthly policy item, &#x60;1&#x60; indicates that the monthly Snapshot occurs on the first day of the month and &#x60;40&#x60; indicates the last day of the month.  - In a weekly policy item, &#x60;1&#x60; indicates that the weekly Snapshot occurs on Monday and &#x60;7&#x60; indicates Sunday.  - In an hourly policy item, you can set the frequency interval to &#x60;1&#x60;, &#x60;2&#x60;, &#x60;4&#x60;, &#x60;6&#x60;, &#x60;8&#x60;, or &#x60;12&#x60;. For hourly policy items for NVMe clusters, MongoDB Cloud accepts only &#x60;12&#x60; as the frequency interval value.   MongoDB Cloud ignores this setting for non-hourly policy items in Backup Compliance Policy settings. | 
**FrequencyType** | **string** | Human-readable label that identifies the frequency type associated with the backup policy. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this backup policy item. | [optional] [readonly] 
**RetentionUnit** | **string** | Unit of time in which MongoDB Cloud measures Snapshot retention. | 
**RetentionValue** | **int** | Duration in days, weeks, months, or years that MongoDB Cloud retains the Snapshot. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items.  For example: If the hourly policy item specifies a retention of two days, you must specify two days or greater for the retention of the weekly policy item. | 

## Methods

### NewBackupComplianceScheduledPolicyItem

`func NewBackupComplianceScheduledPolicyItem(frequencyInterval int, frequencyType string, retentionUnit string, retentionValue int, ) *BackupComplianceScheduledPolicyItem`

NewBackupComplianceScheduledPolicyItem instantiates a new BackupComplianceScheduledPolicyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupComplianceScheduledPolicyItemWithDefaults

`func NewBackupComplianceScheduledPolicyItemWithDefaults() *BackupComplianceScheduledPolicyItem`

NewBackupComplianceScheduledPolicyItemWithDefaults instantiates a new BackupComplianceScheduledPolicyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequencyInterval

`func (o *BackupComplianceScheduledPolicyItem) GetFrequencyInterval() int`

GetFrequencyInterval returns the FrequencyInterval field if non-nil, zero value otherwise.

### GetFrequencyIntervalOk

`func (o *BackupComplianceScheduledPolicyItem) GetFrequencyIntervalOk() (*int, bool)`

GetFrequencyIntervalOk returns a tuple with the FrequencyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyInterval

`func (o *BackupComplianceScheduledPolicyItem) SetFrequencyInterval(v int)`

SetFrequencyInterval sets FrequencyInterval field to given value.

### GetFrequencyType

`func (o *BackupComplianceScheduledPolicyItem) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *BackupComplianceScheduledPolicyItem) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *BackupComplianceScheduledPolicyItem) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### GetId

`func (o *BackupComplianceScheduledPolicyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupComplianceScheduledPolicyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupComplianceScheduledPolicyItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupComplianceScheduledPolicyItem) HasId() bool`

HasId returns a boolean if a field has been set.
### GetRetentionUnit

`func (o *BackupComplianceScheduledPolicyItem) GetRetentionUnit() string`

GetRetentionUnit returns the RetentionUnit field if non-nil, zero value otherwise.

### GetRetentionUnitOk

`func (o *BackupComplianceScheduledPolicyItem) GetRetentionUnitOk() (*string, bool)`

GetRetentionUnitOk returns a tuple with the RetentionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionUnit

`func (o *BackupComplianceScheduledPolicyItem) SetRetentionUnit(v string)`

SetRetentionUnit sets RetentionUnit field to given value.

### GetRetentionValue

`func (o *BackupComplianceScheduledPolicyItem) GetRetentionValue() int`

GetRetentionValue returns the RetentionValue field if non-nil, zero value otherwise.

### GetRetentionValueOk

`func (o *BackupComplianceScheduledPolicyItem) GetRetentionValueOk() (*int, bool)`

GetRetentionValueOk returns a tuple with the RetentionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionValue

`func (o *BackupComplianceScheduledPolicyItem) SetRetentionValue(v int)`

SetRetentionValue sets RetentionValue field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


