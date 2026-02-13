# DiskBackupSnapshotSchedule20240805

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoExportEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud automatically exports Cloud Backup Snapshots to the Export Bucket. | [optional] 
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the Snapshot you want to return. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster with the Snapshot you want to return. | [optional] [readonly] 
**CopyPolicyItemsEnabled** | Pointer to **bool** | Flag that indicates whether copy settings use &#x60;copyPolicyItems&#x60; instead of &#x60;frequencies&#x60;. When true, requests must supply &#x60;copyPolicyItems&#x60; and responses return &#x60;copyPolicyItems&#x60; only. When false or omitted, requests must supply &#x60;frequencies&#x60; and responses return &#x60;frequencies&#x60; only. | [optional] 
**CopySettings** | Pointer to [**[]DiskBackupCopySetting20240805**](DiskBackupCopySetting20240805.md) | List that contains a document for each copy setting item in the desired backup policy. | [optional] 
**DeleteCopiedBackups** | Pointer to [**[]DeleteCopiedBackups20240805**](DeleteCopiedBackups20240805.md) | List that contains a document for each deleted copy setting whose backup copies you want to delete. | [optional] 
**DeleteCopySnapshots** | Pointer to **bool** | Flag that indicates whether to delete Snapshot copies that MongoDB Cloud took previously when their associated &#x60;copyPolicyItems&#x60; are removed from a &#x60;copySetting&#x60;. This option requires &#x60;copyPolicyItemsEnabled&#x60; to be true. | [optional] 
**DeleteSnapshots** | Pointer to **bool** | Flag that indicates whether to delete Snapshots that MongoDB Cloud took previously when deleting the associated backup policy. | [optional] 
**Export** | Pointer to [**AutoExportPolicy**](AutoExportPolicy.md) |  | [optional] 
**ExtraRetentionSettings** | Pointer to [**[]ExtraRetentionSetting**](ExtraRetentionSetting.md) | List that contains a document for each extra retention setting item in the desired backup policy. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**NextSnapshot** | Pointer to **time.Time** | Date and time when MongoDB Cloud takes the next Snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Policies** | Pointer to [**[]AdvancedDiskBackupSnapshotSchedulePolicy**](AdvancedDiskBackupSnapshotSchedulePolicy.md) | Rules set for this backup schedule. | [optional] 
**ReferenceHourOfDay** | Pointer to **int** | Hour of day in Coordinated Universal Time (UTC) that represents when MongoDB Cloud takes the Snapshot. | [optional] 
**ReferenceMinuteOfHour** | Pointer to **int** | Minute of the &#x60;referenceHourOfDay&#x60; that represents when MongoDB Cloud takes the Snapshot. | [optional] 
**RestoreWindowDays** | Pointer to **int** | Number of previous days that you can restore back to with Continuous Cloud Backup accuracy. You must specify a positive, non-zero integer. This parameter applies to continuous Cloud Backups only. | [optional] 
**UpdateCopySnapshots** | Pointer to **bool** | Flag that indicates whether to apply the retention changes for updated copy policy items to Snapshot copies that MongoDB Cloud took previously. | [optional] 
**UpdateSnapshots** | Pointer to **bool** | Flag that indicates whether to apply the retention changes in the updated backup policy to Snapshots that MongoDB Cloud took previously. | [optional] 
**UseOrgAndGroupNamesInExportPrefix** | Pointer to **bool** | Flag that indicates whether to use organization and project names instead of organization and project UUIDs in the path to the metadata files that MongoDB Cloud uploads to your Export Bucket. | [optional] 

## Methods

### NewDiskBackupSnapshotSchedule20240805

`func NewDiskBackupSnapshotSchedule20240805() *DiskBackupSnapshotSchedule20240805`

NewDiskBackupSnapshotSchedule20240805 instantiates a new DiskBackupSnapshotSchedule20240805 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotSchedule20240805WithDefaults

`func NewDiskBackupSnapshotSchedule20240805WithDefaults() *DiskBackupSnapshotSchedule20240805`

NewDiskBackupSnapshotSchedule20240805WithDefaults instantiates a new DiskBackupSnapshotSchedule20240805 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoExportEnabled

`func (o *DiskBackupSnapshotSchedule20240805) GetAutoExportEnabled() bool`

GetAutoExportEnabled returns the AutoExportEnabled field if non-nil, zero value otherwise.

### GetAutoExportEnabledOk

`func (o *DiskBackupSnapshotSchedule20240805) GetAutoExportEnabledOk() (*bool, bool)`

GetAutoExportEnabledOk returns a tuple with the AutoExportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExportEnabled

`func (o *DiskBackupSnapshotSchedule20240805) SetAutoExportEnabled(v bool)`

SetAutoExportEnabled sets AutoExportEnabled field to given value.

### HasAutoExportEnabled

`func (o *DiskBackupSnapshotSchedule20240805) HasAutoExportEnabled() bool`

HasAutoExportEnabled returns a boolean if a field has been set.
### GetClusterId

`func (o *DiskBackupSnapshotSchedule20240805) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DiskBackupSnapshotSchedule20240805) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DiskBackupSnapshotSchedule20240805) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *DiskBackupSnapshotSchedule20240805) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.
### GetClusterName

`func (o *DiskBackupSnapshotSchedule20240805) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DiskBackupSnapshotSchedule20240805) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DiskBackupSnapshotSchedule20240805) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *DiskBackupSnapshotSchedule20240805) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetCopyPolicyItemsEnabled

`func (o *DiskBackupSnapshotSchedule20240805) GetCopyPolicyItemsEnabled() bool`

GetCopyPolicyItemsEnabled returns the CopyPolicyItemsEnabled field if non-nil, zero value otherwise.

### GetCopyPolicyItemsEnabledOk

`func (o *DiskBackupSnapshotSchedule20240805) GetCopyPolicyItemsEnabledOk() (*bool, bool)`

GetCopyPolicyItemsEnabledOk returns a tuple with the CopyPolicyItemsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPolicyItemsEnabled

`func (o *DiskBackupSnapshotSchedule20240805) SetCopyPolicyItemsEnabled(v bool)`

SetCopyPolicyItemsEnabled sets CopyPolicyItemsEnabled field to given value.

### HasCopyPolicyItemsEnabled

`func (o *DiskBackupSnapshotSchedule20240805) HasCopyPolicyItemsEnabled() bool`

HasCopyPolicyItemsEnabled returns a boolean if a field has been set.
### GetCopySettings

`func (o *DiskBackupSnapshotSchedule20240805) GetCopySettings() []DiskBackupCopySetting20240805`

GetCopySettings returns the CopySettings field if non-nil, zero value otherwise.

### GetCopySettingsOk

`func (o *DiskBackupSnapshotSchedule20240805) GetCopySettingsOk() (*[]DiskBackupCopySetting20240805, bool)`

GetCopySettingsOk returns a tuple with the CopySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySettings

`func (o *DiskBackupSnapshotSchedule20240805) SetCopySettings(v []DiskBackupCopySetting20240805)`

SetCopySettings sets CopySettings field to given value.

### HasCopySettings

`func (o *DiskBackupSnapshotSchedule20240805) HasCopySettings() bool`

HasCopySettings returns a boolean if a field has been set.
### GetDeleteCopiedBackups

`func (o *DiskBackupSnapshotSchedule20240805) GetDeleteCopiedBackups() []DeleteCopiedBackups20240805`

GetDeleteCopiedBackups returns the DeleteCopiedBackups field if non-nil, zero value otherwise.

### GetDeleteCopiedBackupsOk

`func (o *DiskBackupSnapshotSchedule20240805) GetDeleteCopiedBackupsOk() (*[]DeleteCopiedBackups20240805, bool)`

GetDeleteCopiedBackupsOk returns a tuple with the DeleteCopiedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCopiedBackups

`func (o *DiskBackupSnapshotSchedule20240805) SetDeleteCopiedBackups(v []DeleteCopiedBackups20240805)`

SetDeleteCopiedBackups sets DeleteCopiedBackups field to given value.

### HasDeleteCopiedBackups

`func (o *DiskBackupSnapshotSchedule20240805) HasDeleteCopiedBackups() bool`

HasDeleteCopiedBackups returns a boolean if a field has been set.
### GetDeleteCopySnapshots

`func (o *DiskBackupSnapshotSchedule20240805) GetDeleteCopySnapshots() bool`

GetDeleteCopySnapshots returns the DeleteCopySnapshots field if non-nil, zero value otherwise.

### GetDeleteCopySnapshotsOk

`func (o *DiskBackupSnapshotSchedule20240805) GetDeleteCopySnapshotsOk() (*bool, bool)`

GetDeleteCopySnapshotsOk returns a tuple with the DeleteCopySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCopySnapshots

`func (o *DiskBackupSnapshotSchedule20240805) SetDeleteCopySnapshots(v bool)`

SetDeleteCopySnapshots sets DeleteCopySnapshots field to given value.

### HasDeleteCopySnapshots

`func (o *DiskBackupSnapshotSchedule20240805) HasDeleteCopySnapshots() bool`

HasDeleteCopySnapshots returns a boolean if a field has been set.
### GetDeleteSnapshots

`func (o *DiskBackupSnapshotSchedule20240805) GetDeleteSnapshots() bool`

GetDeleteSnapshots returns the DeleteSnapshots field if non-nil, zero value otherwise.

### GetDeleteSnapshotsOk

`func (o *DiskBackupSnapshotSchedule20240805) GetDeleteSnapshotsOk() (*bool, bool)`

GetDeleteSnapshotsOk returns a tuple with the DeleteSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSnapshots

`func (o *DiskBackupSnapshotSchedule20240805) SetDeleteSnapshots(v bool)`

SetDeleteSnapshots sets DeleteSnapshots field to given value.

### HasDeleteSnapshots

`func (o *DiskBackupSnapshotSchedule20240805) HasDeleteSnapshots() bool`

HasDeleteSnapshots returns a boolean if a field has been set.
### GetExport

`func (o *DiskBackupSnapshotSchedule20240805) GetExport() AutoExportPolicy`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *DiskBackupSnapshotSchedule20240805) GetExportOk() (*AutoExportPolicy, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *DiskBackupSnapshotSchedule20240805) SetExport(v AutoExportPolicy)`

SetExport sets Export field to given value.

### HasExport

`func (o *DiskBackupSnapshotSchedule20240805) HasExport() bool`

HasExport returns a boolean if a field has been set.
### GetExtraRetentionSettings

`func (o *DiskBackupSnapshotSchedule20240805) GetExtraRetentionSettings() []ExtraRetentionSetting`

GetExtraRetentionSettings returns the ExtraRetentionSettings field if non-nil, zero value otherwise.

### GetExtraRetentionSettingsOk

`func (o *DiskBackupSnapshotSchedule20240805) GetExtraRetentionSettingsOk() (*[]ExtraRetentionSetting, bool)`

GetExtraRetentionSettingsOk returns a tuple with the ExtraRetentionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRetentionSettings

`func (o *DiskBackupSnapshotSchedule20240805) SetExtraRetentionSettings(v []ExtraRetentionSetting)`

SetExtraRetentionSettings sets ExtraRetentionSettings field to given value.

### HasExtraRetentionSettings

`func (o *DiskBackupSnapshotSchedule20240805) HasExtraRetentionSettings() bool`

HasExtraRetentionSettings returns a boolean if a field has been set.
### GetLinks

`func (o *DiskBackupSnapshotSchedule20240805) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshotSchedule20240805) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshotSchedule20240805) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshotSchedule20240805) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetNextSnapshot

`func (o *DiskBackupSnapshotSchedule20240805) GetNextSnapshot() time.Time`

GetNextSnapshot returns the NextSnapshot field if non-nil, zero value otherwise.

### GetNextSnapshotOk

`func (o *DiskBackupSnapshotSchedule20240805) GetNextSnapshotOk() (*time.Time, bool)`

GetNextSnapshotOk returns a tuple with the NextSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSnapshot

`func (o *DiskBackupSnapshotSchedule20240805) SetNextSnapshot(v time.Time)`

SetNextSnapshot sets NextSnapshot field to given value.

### HasNextSnapshot

`func (o *DiskBackupSnapshotSchedule20240805) HasNextSnapshot() bool`

HasNextSnapshot returns a boolean if a field has been set.
### GetPolicies

`func (o *DiskBackupSnapshotSchedule20240805) GetPolicies() []AdvancedDiskBackupSnapshotSchedulePolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *DiskBackupSnapshotSchedule20240805) GetPoliciesOk() (*[]AdvancedDiskBackupSnapshotSchedulePolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *DiskBackupSnapshotSchedule20240805) SetPolicies(v []AdvancedDiskBackupSnapshotSchedulePolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *DiskBackupSnapshotSchedule20240805) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.
### GetReferenceHourOfDay

`func (o *DiskBackupSnapshotSchedule20240805) GetReferenceHourOfDay() int`

GetReferenceHourOfDay returns the ReferenceHourOfDay field if non-nil, zero value otherwise.

### GetReferenceHourOfDayOk

`func (o *DiskBackupSnapshotSchedule20240805) GetReferenceHourOfDayOk() (*int, bool)`

GetReferenceHourOfDayOk returns a tuple with the ReferenceHourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceHourOfDay

`func (o *DiskBackupSnapshotSchedule20240805) SetReferenceHourOfDay(v int)`

SetReferenceHourOfDay sets ReferenceHourOfDay field to given value.

### HasReferenceHourOfDay

`func (o *DiskBackupSnapshotSchedule20240805) HasReferenceHourOfDay() bool`

HasReferenceHourOfDay returns a boolean if a field has been set.
### GetReferenceMinuteOfHour

`func (o *DiskBackupSnapshotSchedule20240805) GetReferenceMinuteOfHour() int`

GetReferenceMinuteOfHour returns the ReferenceMinuteOfHour field if non-nil, zero value otherwise.

### GetReferenceMinuteOfHourOk

`func (o *DiskBackupSnapshotSchedule20240805) GetReferenceMinuteOfHourOk() (*int, bool)`

GetReferenceMinuteOfHourOk returns a tuple with the ReferenceMinuteOfHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMinuteOfHour

`func (o *DiskBackupSnapshotSchedule20240805) SetReferenceMinuteOfHour(v int)`

SetReferenceMinuteOfHour sets ReferenceMinuteOfHour field to given value.

### HasReferenceMinuteOfHour

`func (o *DiskBackupSnapshotSchedule20240805) HasReferenceMinuteOfHour() bool`

HasReferenceMinuteOfHour returns a boolean if a field has been set.
### GetRestoreWindowDays

`func (o *DiskBackupSnapshotSchedule20240805) GetRestoreWindowDays() int`

GetRestoreWindowDays returns the RestoreWindowDays field if non-nil, zero value otherwise.

### GetRestoreWindowDaysOk

`func (o *DiskBackupSnapshotSchedule20240805) GetRestoreWindowDaysOk() (*int, bool)`

GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreWindowDays

`func (o *DiskBackupSnapshotSchedule20240805) SetRestoreWindowDays(v int)`

SetRestoreWindowDays sets RestoreWindowDays field to given value.

### HasRestoreWindowDays

`func (o *DiskBackupSnapshotSchedule20240805) HasRestoreWindowDays() bool`

HasRestoreWindowDays returns a boolean if a field has been set.
### GetUpdateCopySnapshots

`func (o *DiskBackupSnapshotSchedule20240805) GetUpdateCopySnapshots() bool`

GetUpdateCopySnapshots returns the UpdateCopySnapshots field if non-nil, zero value otherwise.

### GetUpdateCopySnapshotsOk

`func (o *DiskBackupSnapshotSchedule20240805) GetUpdateCopySnapshotsOk() (*bool, bool)`

GetUpdateCopySnapshotsOk returns a tuple with the UpdateCopySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCopySnapshots

`func (o *DiskBackupSnapshotSchedule20240805) SetUpdateCopySnapshots(v bool)`

SetUpdateCopySnapshots sets UpdateCopySnapshots field to given value.

### HasUpdateCopySnapshots

`func (o *DiskBackupSnapshotSchedule20240805) HasUpdateCopySnapshots() bool`

HasUpdateCopySnapshots returns a boolean if a field has been set.
### GetUpdateSnapshots

`func (o *DiskBackupSnapshotSchedule20240805) GetUpdateSnapshots() bool`

GetUpdateSnapshots returns the UpdateSnapshots field if non-nil, zero value otherwise.

### GetUpdateSnapshotsOk

`func (o *DiskBackupSnapshotSchedule20240805) GetUpdateSnapshotsOk() (*bool, bool)`

GetUpdateSnapshotsOk returns a tuple with the UpdateSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSnapshots

`func (o *DiskBackupSnapshotSchedule20240805) SetUpdateSnapshots(v bool)`

SetUpdateSnapshots sets UpdateSnapshots field to given value.

### HasUpdateSnapshots

`func (o *DiskBackupSnapshotSchedule20240805) HasUpdateSnapshots() bool`

HasUpdateSnapshots returns a boolean if a field has been set.
### GetUseOrgAndGroupNamesInExportPrefix

`func (o *DiskBackupSnapshotSchedule20240805) GetUseOrgAndGroupNamesInExportPrefix() bool`

GetUseOrgAndGroupNamesInExportPrefix returns the UseOrgAndGroupNamesInExportPrefix field if non-nil, zero value otherwise.

### GetUseOrgAndGroupNamesInExportPrefixOk

`func (o *DiskBackupSnapshotSchedule20240805) GetUseOrgAndGroupNamesInExportPrefixOk() (*bool, bool)`

GetUseOrgAndGroupNamesInExportPrefixOk returns a tuple with the UseOrgAndGroupNamesInExportPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOrgAndGroupNamesInExportPrefix

`func (o *DiskBackupSnapshotSchedule20240805) SetUseOrgAndGroupNamesInExportPrefix(v bool)`

SetUseOrgAndGroupNamesInExportPrefix sets UseOrgAndGroupNamesInExportPrefix field to given value.

### HasUseOrgAndGroupNamesInExportPrefix

`func (o *DiskBackupSnapshotSchedule20240805) HasUseOrgAndGroupNamesInExportPrefix() bool`

HasUseOrgAndGroupNamesInExportPrefix returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


