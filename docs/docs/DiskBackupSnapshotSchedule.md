# DiskBackupSnapshotSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoExportEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud automatically exports cloud backup snapshots to the AWS bucket. | [optional] 
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | [optional] [readonly] 
**CopySettings** | Pointer to [**[]DiskBackupCopySetting**](DiskBackupCopySetting.md) | List that contains a document for each copy setting item in the desired backup policy. | [optional] 
**DeleteCopiedBackups** | Pointer to [**[]DeleteCopiedBackups**](DeleteCopiedBackups.md) | List that contains a document for each deleted copy setting whose backup copies you want to delete. | [optional] 
**Export** | Pointer to [**AutoExportPolicy**](AutoExportPolicy.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**NextSnapshot** | Pointer to **time.Time** | Date and time when MongoDB Cloud takes the next snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Policies** | Pointer to [**[]Policy**](Policy.md) | Rules set for this backup schedule. | [optional] 
**ReferenceHourOfDay** | Pointer to **int** | Hour of day in Coordinated Universal Time (UTC) that represents when MongoDB Cloud takes the snapshot. | [optional] 
**ReferenceMinuteOfHour** | Pointer to **int** | Minute of the **referenceHourOfDay** that represents when MongoDB Cloud takes the snapshot. | [optional] 
**RestoreWindowDays** | Pointer to **int** | Number of previous days that you can restore back to with Continuous Cloud Backup accuracy. You must specify a positive, non-zero integer. This parameter applies to continuous cloud backups only. | [optional] 
**UpdateSnapshots** | Pointer to **bool** | Flag that indicates whether to apply the retention changes in the updated backup policy to snapshots that MongoDB Cloud took previously. | [optional] 
**UseOrgAndGroupNamesInExportPrefix** | Pointer to **bool** | Flag that indicates whether to use organization and project names instead of organization and project UUIDs in the path to the metadata files that MongoDB Cloud uploads to your AWS bucket. | [optional] 

## Methods

### NewDiskBackupSnapshotSchedule

`func NewDiskBackupSnapshotSchedule() *DiskBackupSnapshotSchedule`

NewDiskBackupSnapshotSchedule instantiates a new DiskBackupSnapshotSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotScheduleWithDefaults

`func NewDiskBackupSnapshotScheduleWithDefaults() *DiskBackupSnapshotSchedule`

NewDiskBackupSnapshotScheduleWithDefaults instantiates a new DiskBackupSnapshotSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoExportEnabled

`func (o *DiskBackupSnapshotSchedule) GetAutoExportEnabled() bool`

GetAutoExportEnabled returns the AutoExportEnabled field if non-nil, zero value otherwise.

### GetAutoExportEnabledOk

`func (o *DiskBackupSnapshotSchedule) GetAutoExportEnabledOk() (*bool, bool)`

GetAutoExportEnabledOk returns a tuple with the AutoExportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExportEnabled

`func (o *DiskBackupSnapshotSchedule) SetAutoExportEnabled(v bool)`

SetAutoExportEnabled sets AutoExportEnabled field to given value.

### HasAutoExportEnabled

`func (o *DiskBackupSnapshotSchedule) HasAutoExportEnabled() bool`

HasAutoExportEnabled returns a boolean if a field has been set.

### GetClusterId

`func (o *DiskBackupSnapshotSchedule) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DiskBackupSnapshotSchedule) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DiskBackupSnapshotSchedule) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *DiskBackupSnapshotSchedule) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterName

`func (o *DiskBackupSnapshotSchedule) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DiskBackupSnapshotSchedule) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DiskBackupSnapshotSchedule) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *DiskBackupSnapshotSchedule) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCopySettings

`func (o *DiskBackupSnapshotSchedule) GetCopySettings() []DiskBackupCopySetting`

GetCopySettings returns the CopySettings field if non-nil, zero value otherwise.

### GetCopySettingsOk

`func (o *DiskBackupSnapshotSchedule) GetCopySettingsOk() (*[]DiskBackupCopySetting, bool)`

GetCopySettingsOk returns a tuple with the CopySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySettings

`func (o *DiskBackupSnapshotSchedule) SetCopySettings(v []DiskBackupCopySetting)`

SetCopySettings sets CopySettings field to given value.

### HasCopySettings

`func (o *DiskBackupSnapshotSchedule) HasCopySettings() bool`

HasCopySettings returns a boolean if a field has been set.

### GetDeleteCopiedBackups

`func (o *DiskBackupSnapshotSchedule) GetDeleteCopiedBackups() []DeleteCopiedBackups`

GetDeleteCopiedBackups returns the DeleteCopiedBackups field if non-nil, zero value otherwise.

### GetDeleteCopiedBackupsOk

`func (o *DiskBackupSnapshotSchedule) GetDeleteCopiedBackupsOk() (*[]DeleteCopiedBackups, bool)`

GetDeleteCopiedBackupsOk returns a tuple with the DeleteCopiedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCopiedBackups

`func (o *DiskBackupSnapshotSchedule) SetDeleteCopiedBackups(v []DeleteCopiedBackups)`

SetDeleteCopiedBackups sets DeleteCopiedBackups field to given value.

### HasDeleteCopiedBackups

`func (o *DiskBackupSnapshotSchedule) HasDeleteCopiedBackups() bool`

HasDeleteCopiedBackups returns a boolean if a field has been set.

### GetExport

`func (o *DiskBackupSnapshotSchedule) GetExport() AutoExportPolicy`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *DiskBackupSnapshotSchedule) GetExportOk() (*AutoExportPolicy, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *DiskBackupSnapshotSchedule) SetExport(v AutoExportPolicy)`

SetExport sets Export field to given value.

### HasExport

`func (o *DiskBackupSnapshotSchedule) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetLinks

`func (o *DiskBackupSnapshotSchedule) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshotSchedule) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshotSchedule) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshotSchedule) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetNextSnapshot

`func (o *DiskBackupSnapshotSchedule) GetNextSnapshot() time.Time`

GetNextSnapshot returns the NextSnapshot field if non-nil, zero value otherwise.

### GetNextSnapshotOk

`func (o *DiskBackupSnapshotSchedule) GetNextSnapshotOk() (*time.Time, bool)`

GetNextSnapshotOk returns a tuple with the NextSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSnapshot

`func (o *DiskBackupSnapshotSchedule) SetNextSnapshot(v time.Time)`

SetNextSnapshot sets NextSnapshot field to given value.

### HasNextSnapshot

`func (o *DiskBackupSnapshotSchedule) HasNextSnapshot() bool`

HasNextSnapshot returns a boolean if a field has been set.

### GetPolicies

`func (o *DiskBackupSnapshotSchedule) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *DiskBackupSnapshotSchedule) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *DiskBackupSnapshotSchedule) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *DiskBackupSnapshotSchedule) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetReferenceHourOfDay

`func (o *DiskBackupSnapshotSchedule) GetReferenceHourOfDay() int`

GetReferenceHourOfDay returns the ReferenceHourOfDay field if non-nil, zero value otherwise.

### GetReferenceHourOfDayOk

`func (o *DiskBackupSnapshotSchedule) GetReferenceHourOfDayOk() (*int, bool)`

GetReferenceHourOfDayOk returns a tuple with the ReferenceHourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceHourOfDay

`func (o *DiskBackupSnapshotSchedule) SetReferenceHourOfDay(v int)`

SetReferenceHourOfDay sets ReferenceHourOfDay field to given value.

### HasReferenceHourOfDay

`func (o *DiskBackupSnapshotSchedule) HasReferenceHourOfDay() bool`

HasReferenceHourOfDay returns a boolean if a field has been set.

### GetReferenceMinuteOfHour

`func (o *DiskBackupSnapshotSchedule) GetReferenceMinuteOfHour() int`

GetReferenceMinuteOfHour returns the ReferenceMinuteOfHour field if non-nil, zero value otherwise.

### GetReferenceMinuteOfHourOk

`func (o *DiskBackupSnapshotSchedule) GetReferenceMinuteOfHourOk() (*int, bool)`

GetReferenceMinuteOfHourOk returns a tuple with the ReferenceMinuteOfHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMinuteOfHour

`func (o *DiskBackupSnapshotSchedule) SetReferenceMinuteOfHour(v int)`

SetReferenceMinuteOfHour sets ReferenceMinuteOfHour field to given value.

### HasReferenceMinuteOfHour

`func (o *DiskBackupSnapshotSchedule) HasReferenceMinuteOfHour() bool`

HasReferenceMinuteOfHour returns a boolean if a field has been set.

### GetRestoreWindowDays

`func (o *DiskBackupSnapshotSchedule) GetRestoreWindowDays() int`

GetRestoreWindowDays returns the RestoreWindowDays field if non-nil, zero value otherwise.

### GetRestoreWindowDaysOk

`func (o *DiskBackupSnapshotSchedule) GetRestoreWindowDaysOk() (*int, bool)`

GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreWindowDays

`func (o *DiskBackupSnapshotSchedule) SetRestoreWindowDays(v int)`

SetRestoreWindowDays sets RestoreWindowDays field to given value.

### HasRestoreWindowDays

`func (o *DiskBackupSnapshotSchedule) HasRestoreWindowDays() bool`

HasRestoreWindowDays returns a boolean if a field has been set.

### GetUpdateSnapshots

`func (o *DiskBackupSnapshotSchedule) GetUpdateSnapshots() bool`

GetUpdateSnapshots returns the UpdateSnapshots field if non-nil, zero value otherwise.

### GetUpdateSnapshotsOk

`func (o *DiskBackupSnapshotSchedule) GetUpdateSnapshotsOk() (*bool, bool)`

GetUpdateSnapshotsOk returns a tuple with the UpdateSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSnapshots

`func (o *DiskBackupSnapshotSchedule) SetUpdateSnapshots(v bool)`

SetUpdateSnapshots sets UpdateSnapshots field to given value.

### HasUpdateSnapshots

`func (o *DiskBackupSnapshotSchedule) HasUpdateSnapshots() bool`

HasUpdateSnapshots returns a boolean if a field has been set.

### GetUseOrgAndGroupNamesInExportPrefix

`func (o *DiskBackupSnapshotSchedule) GetUseOrgAndGroupNamesInExportPrefix() bool`

GetUseOrgAndGroupNamesInExportPrefix returns the UseOrgAndGroupNamesInExportPrefix field if non-nil, zero value otherwise.

### GetUseOrgAndGroupNamesInExportPrefixOk

`func (o *DiskBackupSnapshotSchedule) GetUseOrgAndGroupNamesInExportPrefixOk() (*bool, bool)`

GetUseOrgAndGroupNamesInExportPrefixOk returns a tuple with the UseOrgAndGroupNamesInExportPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOrgAndGroupNamesInExportPrefix

`func (o *DiskBackupSnapshotSchedule) SetUseOrgAndGroupNamesInExportPrefix(v bool)`

SetUseOrgAndGroupNamesInExportPrefix sets UseOrgAndGroupNamesInExportPrefix field to given value.

### HasUseOrgAndGroupNamesInExportPrefix

`func (o *DiskBackupSnapshotSchedule) HasUseOrgAndGroupNamesInExportPrefix() bool`

HasUseOrgAndGroupNamesInExportPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


