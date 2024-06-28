# DiskBackupExportMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the the Cloud Backup snapshot export job for each shard in a sharded cluster. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set on the sharded cluster. | [optional] [readonly] 

## Methods

### NewDiskBackupExportMember

`func NewDiskBackupExportMember() *DiskBackupExportMember`

NewDiskBackupExportMember instantiates a new DiskBackupExportMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupExportMemberWithDefaults

`func NewDiskBackupExportMemberWithDefaults() *DiskBackupExportMember`

NewDiskBackupExportMemberWithDefaults instantiates a new DiskBackupExportMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportId

`func (o *DiskBackupExportMember) GetExportId() string`

GetExportId returns the ExportId field if non-nil, zero value otherwise.

### GetExportIdOk

`func (o *DiskBackupExportMember) GetExportIdOk() (*string, bool)`

GetExportIdOk returns a tuple with the ExportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportId

`func (o *DiskBackupExportMember) SetExportId(v string)`

SetExportId sets ExportId field to given value.

### HasExportId

`func (o *DiskBackupExportMember) HasExportId() bool`

HasExportId returns a boolean if a field has been set.
### GetReplicaSetName

`func (o *DiskBackupExportMember) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *DiskBackupExportMember) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *DiskBackupExportMember) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *DiskBackupExportMember) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


