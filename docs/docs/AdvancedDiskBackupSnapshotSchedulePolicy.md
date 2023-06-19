# AdvancedDiskBackupSnapshotSchedulePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this backup policy. | [optional] 
**PolicyItems** | Pointer to [**[]DiskBackupApiPolicyItem**](DiskBackupApiPolicyItem.md) | List that contains the specifications for one policy. | [optional] 

## Methods

### NewAdvancedDiskBackupSnapshotSchedulePolicy

`func NewAdvancedDiskBackupSnapshotSchedulePolicy() *AdvancedDiskBackupSnapshotSchedulePolicy`

NewAdvancedDiskBackupSnapshotSchedulePolicy instantiates a new AdvancedDiskBackupSnapshotSchedulePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedDiskBackupSnapshotSchedulePolicyWithDefaults

`func NewAdvancedDiskBackupSnapshotSchedulePolicyWithDefaults() *AdvancedDiskBackupSnapshotSchedulePolicy`

NewAdvancedDiskBackupSnapshotSchedulePolicyWithDefaults instantiates a new AdvancedDiskBackupSnapshotSchedulePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdvancedDiskBackupSnapshotSchedulePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedDiskBackupSnapshotSchedulePolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedDiskBackupSnapshotSchedulePolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedDiskBackupSnapshotSchedulePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyItems

`func (o *AdvancedDiskBackupSnapshotSchedulePolicy) GetPolicyItems() []DiskBackupApiPolicyItem`

GetPolicyItems returns the PolicyItems field if non-nil, zero value otherwise.

### GetPolicyItemsOk

`func (o *AdvancedDiskBackupSnapshotSchedulePolicy) GetPolicyItemsOk() (*[]DiskBackupApiPolicyItem, bool)`

GetPolicyItemsOk returns a tuple with the PolicyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyItems

`func (o *AdvancedDiskBackupSnapshotSchedulePolicy) SetPolicyItems(v []DiskBackupApiPolicyItem)`

SetPolicyItems sets PolicyItems field to given value.

### HasPolicyItems

`func (o *AdvancedDiskBackupSnapshotSchedulePolicy) HasPolicyItems() bool`

HasPolicyItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


