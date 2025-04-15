# DataProtectionSettings20231001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedEmail** | **string** | Email address of the user who authorized to update the Backup Compliance Policy  settings. | 
**AuthorizedUserFirstName** | **string** | First name of the user who authorized to updated the Backup Compliance Policy  settings. | 
**AuthorizedUserLastName** | **string** | Last name of the user who authorized to updated the Backup Compliance Policy  settings. | 
**CopyProtectionEnabled** | Pointer to **bool** | Flag that indicates whether to prevent cluster users from deleting backups copied to other regions, even if those additional snapshot regions are removed. If unspecified, this value defaults to false. | [optional] [default to false]
**Deletable** | Pointer to **bool** | Flag that indicates whether the Backup Compliance Policy is allowed to be disabled. It is default to false and a support ticket needs to be filed to request setting to true. | [optional] [readonly] [default to false]
**EncryptionAtRestEnabled** | Pointer to **bool** | Flag that indicates whether Encryption at Rest using Customer Key  Management is required for all clusters with a Backup Compliance Policy. If unspecified, this value defaults to false. | [optional] [default to false]
**OnDemandPolicyItem** | Pointer to [**BackupComplianceOnDemandPolicyItem**](BackupComplianceOnDemandPolicyItem.md) |  | [optional] 
**PitEnabled** | Pointer to **bool** | Flag that indicates whether the cluster uses Continuous Cloud Backups with a Backup Compliance Policy. If unspecified, this value defaults to false. | [optional] [default to false]
**ProjectId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project for the Backup Compliance Policy. | [optional] 
**RestoreWindowDays** | Pointer to **int** | Number of previous days that you can restore back to with Continuous Cloud Backup with a Backup Compliance Policy. You must specify a positive, non-zero integer, and the maximum retention window can&#39;t exceed the hourly retention time. This parameter applies only to Continuous Cloud Backups with a Backup Compliance Policy. | [optional] 
**ScheduledPolicyItems** | Pointer to [**[]BackupComplianceScheduledPolicyItem**](BackupComplianceScheduledPolicyItem.md) | List that contains the specifications for one scheduled policy. | [optional] 
**State** | Pointer to **string** | Label that indicates the state of the Backup Compliance Policy settings. MongoDB Cloud ignores this setting when you enable or update the Backup Compliance Policy settings. | [optional] [readonly] 
**UpdatedDate** | Pointer to **time.Time** | ISO 8601 timestamp format in UTC that indicates when the user updated the Data Protection Policy settings. MongoDB Cloud ignores this setting when you enable or update the Backup Compliance Policy settings. | [optional] [readonly] 
**UpdatedUser** | Pointer to **string** | Email address that identifies the user who updated the Backup Compliance Policy settings. MongoDB Cloud ignores this email setting when you enable or update the Backup Compliance Policy settings. | [optional] [readonly] 

## Methods

### NewDataProtectionSettings20231001

`func NewDataProtectionSettings20231001(authorizedEmail string, authorizedUserFirstName string, authorizedUserLastName string, ) *DataProtectionSettings20231001`

NewDataProtectionSettings20231001 instantiates a new DataProtectionSettings20231001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataProtectionSettings20231001WithDefaults

`func NewDataProtectionSettings20231001WithDefaults() *DataProtectionSettings20231001`

NewDataProtectionSettings20231001WithDefaults instantiates a new DataProtectionSettings20231001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedEmail

`func (o *DataProtectionSettings20231001) GetAuthorizedEmail() string`

GetAuthorizedEmail returns the AuthorizedEmail field if non-nil, zero value otherwise.

### GetAuthorizedEmailOk

`func (o *DataProtectionSettings20231001) GetAuthorizedEmailOk() (*string, bool)`

GetAuthorizedEmailOk returns a tuple with the AuthorizedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedEmail

`func (o *DataProtectionSettings20231001) SetAuthorizedEmail(v string)`

SetAuthorizedEmail sets AuthorizedEmail field to given value.

### GetAuthorizedUserFirstName

`func (o *DataProtectionSettings20231001) GetAuthorizedUserFirstName() string`

GetAuthorizedUserFirstName returns the AuthorizedUserFirstName field if non-nil, zero value otherwise.

### GetAuthorizedUserFirstNameOk

`func (o *DataProtectionSettings20231001) GetAuthorizedUserFirstNameOk() (*string, bool)`

GetAuthorizedUserFirstNameOk returns a tuple with the AuthorizedUserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUserFirstName

`func (o *DataProtectionSettings20231001) SetAuthorizedUserFirstName(v string)`

SetAuthorizedUserFirstName sets AuthorizedUserFirstName field to given value.

### GetAuthorizedUserLastName

`func (o *DataProtectionSettings20231001) GetAuthorizedUserLastName() string`

GetAuthorizedUserLastName returns the AuthorizedUserLastName field if non-nil, zero value otherwise.

### GetAuthorizedUserLastNameOk

`func (o *DataProtectionSettings20231001) GetAuthorizedUserLastNameOk() (*string, bool)`

GetAuthorizedUserLastNameOk returns a tuple with the AuthorizedUserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUserLastName

`func (o *DataProtectionSettings20231001) SetAuthorizedUserLastName(v string)`

SetAuthorizedUserLastName sets AuthorizedUserLastName field to given value.

### GetCopyProtectionEnabled

`func (o *DataProtectionSettings20231001) GetCopyProtectionEnabled() bool`

GetCopyProtectionEnabled returns the CopyProtectionEnabled field if non-nil, zero value otherwise.

### GetCopyProtectionEnabledOk

`func (o *DataProtectionSettings20231001) GetCopyProtectionEnabledOk() (*bool, bool)`

GetCopyProtectionEnabledOk returns a tuple with the CopyProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyProtectionEnabled

`func (o *DataProtectionSettings20231001) SetCopyProtectionEnabled(v bool)`

SetCopyProtectionEnabled sets CopyProtectionEnabled field to given value.

### HasCopyProtectionEnabled

`func (o *DataProtectionSettings20231001) HasCopyProtectionEnabled() bool`

HasCopyProtectionEnabled returns a boolean if a field has been set.
### GetDeletable

`func (o *DataProtectionSettings20231001) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *DataProtectionSettings20231001) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *DataProtectionSettings20231001) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *DataProtectionSettings20231001) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.
### GetEncryptionAtRestEnabled

`func (o *DataProtectionSettings20231001) GetEncryptionAtRestEnabled() bool`

GetEncryptionAtRestEnabled returns the EncryptionAtRestEnabled field if non-nil, zero value otherwise.

### GetEncryptionAtRestEnabledOk

`func (o *DataProtectionSettings20231001) GetEncryptionAtRestEnabledOk() (*bool, bool)`

GetEncryptionAtRestEnabledOk returns a tuple with the EncryptionAtRestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestEnabled

`func (o *DataProtectionSettings20231001) SetEncryptionAtRestEnabled(v bool)`

SetEncryptionAtRestEnabled sets EncryptionAtRestEnabled field to given value.

### HasEncryptionAtRestEnabled

`func (o *DataProtectionSettings20231001) HasEncryptionAtRestEnabled() bool`

HasEncryptionAtRestEnabled returns a boolean if a field has been set.
### GetOnDemandPolicyItem

`func (o *DataProtectionSettings20231001) GetOnDemandPolicyItem() BackupComplianceOnDemandPolicyItem`

GetOnDemandPolicyItem returns the OnDemandPolicyItem field if non-nil, zero value otherwise.

### GetOnDemandPolicyItemOk

`func (o *DataProtectionSettings20231001) GetOnDemandPolicyItemOk() (*BackupComplianceOnDemandPolicyItem, bool)`

GetOnDemandPolicyItemOk returns a tuple with the OnDemandPolicyItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandPolicyItem

`func (o *DataProtectionSettings20231001) SetOnDemandPolicyItem(v BackupComplianceOnDemandPolicyItem)`

SetOnDemandPolicyItem sets OnDemandPolicyItem field to given value.

### HasOnDemandPolicyItem

`func (o *DataProtectionSettings20231001) HasOnDemandPolicyItem() bool`

HasOnDemandPolicyItem returns a boolean if a field has been set.
### GetPitEnabled

`func (o *DataProtectionSettings20231001) GetPitEnabled() bool`

GetPitEnabled returns the PitEnabled field if non-nil, zero value otherwise.

### GetPitEnabledOk

`func (o *DataProtectionSettings20231001) GetPitEnabledOk() (*bool, bool)`

GetPitEnabledOk returns a tuple with the PitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitEnabled

`func (o *DataProtectionSettings20231001) SetPitEnabled(v bool)`

SetPitEnabled sets PitEnabled field to given value.

### HasPitEnabled

`func (o *DataProtectionSettings20231001) HasPitEnabled() bool`

HasPitEnabled returns a boolean if a field has been set.
### GetProjectId

`func (o *DataProtectionSettings20231001) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DataProtectionSettings20231001) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DataProtectionSettings20231001) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DataProtectionSettings20231001) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.
### GetRestoreWindowDays

`func (o *DataProtectionSettings20231001) GetRestoreWindowDays() int`

GetRestoreWindowDays returns the RestoreWindowDays field if non-nil, zero value otherwise.

### GetRestoreWindowDaysOk

`func (o *DataProtectionSettings20231001) GetRestoreWindowDaysOk() (*int, bool)`

GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreWindowDays

`func (o *DataProtectionSettings20231001) SetRestoreWindowDays(v int)`

SetRestoreWindowDays sets RestoreWindowDays field to given value.

### HasRestoreWindowDays

`func (o *DataProtectionSettings20231001) HasRestoreWindowDays() bool`

HasRestoreWindowDays returns a boolean if a field has been set.
### GetScheduledPolicyItems

`func (o *DataProtectionSettings20231001) GetScheduledPolicyItems() []BackupComplianceScheduledPolicyItem`

GetScheduledPolicyItems returns the ScheduledPolicyItems field if non-nil, zero value otherwise.

### GetScheduledPolicyItemsOk

`func (o *DataProtectionSettings20231001) GetScheduledPolicyItemsOk() (*[]BackupComplianceScheduledPolicyItem, bool)`

GetScheduledPolicyItemsOk returns a tuple with the ScheduledPolicyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPolicyItems

`func (o *DataProtectionSettings20231001) SetScheduledPolicyItems(v []BackupComplianceScheduledPolicyItem)`

SetScheduledPolicyItems sets ScheduledPolicyItems field to given value.

### HasScheduledPolicyItems

`func (o *DataProtectionSettings20231001) HasScheduledPolicyItems() bool`

HasScheduledPolicyItems returns a boolean if a field has been set.
### GetState

`func (o *DataProtectionSettings20231001) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DataProtectionSettings20231001) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DataProtectionSettings20231001) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DataProtectionSettings20231001) HasState() bool`

HasState returns a boolean if a field has been set.
### GetUpdatedDate

`func (o *DataProtectionSettings20231001) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *DataProtectionSettings20231001) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *DataProtectionSettings20231001) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *DataProtectionSettings20231001) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.
### GetUpdatedUser

`func (o *DataProtectionSettings20231001) GetUpdatedUser() string`

GetUpdatedUser returns the UpdatedUser field if non-nil, zero value otherwise.

### GetUpdatedUserOk

`func (o *DataProtectionSettings20231001) GetUpdatedUserOk() (*string, bool)`

GetUpdatedUserOk returns a tuple with the UpdatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedUser

`func (o *DataProtectionSettings20231001) SetUpdatedUser(v string)`

SetUpdatedUser sets UpdatedUser field to given value.

### HasUpdatedUser

`func (o *DataProtectionSettings20231001) HasUpdatedUser() bool`

HasUpdatedUser returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


