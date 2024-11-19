# BackupSnapshotPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return. | [optional] [readonly] 
**CompletedTime** | Pointer to **time.Time** | Date and time when the snapshot completed. | [optional] [readonly] 
**CompressionSetting** | Pointer to **string** | Human-readable label that identifies the method of compression for the snapshot. | [optional] [readonly] 
**DataSizeBytes** | Pointer to **int64** | Total size of the data stored on each node in the cluster. This parameter expresses its value in bytes. | [optional] [readonly] 
**EncryptionEnabled** | Pointer to **bool** | Flag that indicates whether someone encrypted this snapshot. | [optional] [readonly] 
**Fcv** | Pointer to **string** | Number that indicates the feature compatibility version of MongoDB that the replica set primary ran when MongoDB Cloud created the snapshot. | [optional] [readonly] 
**FileSizeBytes** | Pointer to **int64** | Number that indicates the total size of the data files in bytes. | [optional] [readonly] 
**MachineId** | Pointer to **string** | Hostname and port that indicate the node on which MongoDB Cloud created the snapshot. | [optional] [readonly] 
**MasterKeyUUID** | Pointer to **string** | Unique string that identifies the Key Management Interoperability (KMIP) master key used to encrypt the snapshot data. The resource returns this parameter when &#x60;\&quot;parts.encryptionEnabled\&quot; : true&#x60;. | [optional] [readonly] 
**MongodVersion** | Pointer to **string** | Number that indicates the version of MongoDB that the replica set primary ran when MongoDB Cloud created the snapshot. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set. | [optional] [readonly] 
**ReplicaState** | Pointer to **string** | The node&#39;s role at the time when snapshot process began. | [optional] [readonly] 
**StorageSizeBytes** | Pointer to **int64** | Number that indicates the total size of space allocated for document storage. | [optional] [readonly] 
**TypeName** | Pointer to **string** | Human-readable label that identifies the type of server from which MongoDB Cloud took this snapshot. | [optional] [readonly] 

## Methods

### NewBackupSnapshotPart

`func NewBackupSnapshotPart() *BackupSnapshotPart`

NewBackupSnapshotPart instantiates a new BackupSnapshotPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSnapshotPartWithDefaults

`func NewBackupSnapshotPartWithDefaults() *BackupSnapshotPart`

NewBackupSnapshotPartWithDefaults instantiates a new BackupSnapshotPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *BackupSnapshotPart) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BackupSnapshotPart) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BackupSnapshotPart) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *BackupSnapshotPart) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.
### GetCompletedTime

`func (o *BackupSnapshotPart) GetCompletedTime() time.Time`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *BackupSnapshotPart) GetCompletedTimeOk() (*time.Time, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *BackupSnapshotPart) SetCompletedTime(v time.Time)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *BackupSnapshotPart) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.
### GetCompressionSetting

`func (o *BackupSnapshotPart) GetCompressionSetting() string`

GetCompressionSetting returns the CompressionSetting field if non-nil, zero value otherwise.

### GetCompressionSettingOk

`func (o *BackupSnapshotPart) GetCompressionSettingOk() (*string, bool)`

GetCompressionSettingOk returns a tuple with the CompressionSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionSetting

`func (o *BackupSnapshotPart) SetCompressionSetting(v string)`

SetCompressionSetting sets CompressionSetting field to given value.

### HasCompressionSetting

`func (o *BackupSnapshotPart) HasCompressionSetting() bool`

HasCompressionSetting returns a boolean if a field has been set.
### GetDataSizeBytes

`func (o *BackupSnapshotPart) GetDataSizeBytes() int64`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *BackupSnapshotPart) GetDataSizeBytesOk() (*int64, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *BackupSnapshotPart) SetDataSizeBytes(v int64)`

SetDataSizeBytes sets DataSizeBytes field to given value.

### HasDataSizeBytes

`func (o *BackupSnapshotPart) HasDataSizeBytes() bool`

HasDataSizeBytes returns a boolean if a field has been set.
### GetEncryptionEnabled

`func (o *BackupSnapshotPart) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *BackupSnapshotPart) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *BackupSnapshotPart) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *BackupSnapshotPart) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.
### GetFcv

`func (o *BackupSnapshotPart) GetFcv() string`

GetFcv returns the Fcv field if non-nil, zero value otherwise.

### GetFcvOk

`func (o *BackupSnapshotPart) GetFcvOk() (*string, bool)`

GetFcvOk returns a tuple with the Fcv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcv

`func (o *BackupSnapshotPart) SetFcv(v string)`

SetFcv sets Fcv field to given value.

### HasFcv

`func (o *BackupSnapshotPart) HasFcv() bool`

HasFcv returns a boolean if a field has been set.
### GetFileSizeBytes

`func (o *BackupSnapshotPart) GetFileSizeBytes() int64`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *BackupSnapshotPart) GetFileSizeBytesOk() (*int64, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *BackupSnapshotPart) SetFileSizeBytes(v int64)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *BackupSnapshotPart) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.
### GetMachineId

`func (o *BackupSnapshotPart) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *BackupSnapshotPart) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *BackupSnapshotPart) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *BackupSnapshotPart) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.
### GetMasterKeyUUID

`func (o *BackupSnapshotPart) GetMasterKeyUUID() string`

GetMasterKeyUUID returns the MasterKeyUUID field if non-nil, zero value otherwise.

### GetMasterKeyUUIDOk

`func (o *BackupSnapshotPart) GetMasterKeyUUIDOk() (*string, bool)`

GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyUUID

`func (o *BackupSnapshotPart) SetMasterKeyUUID(v string)`

SetMasterKeyUUID sets MasterKeyUUID field to given value.

### HasMasterKeyUUID

`func (o *BackupSnapshotPart) HasMasterKeyUUID() bool`

HasMasterKeyUUID returns a boolean if a field has been set.
### GetMongodVersion

`func (o *BackupSnapshotPart) GetMongodVersion() string`

GetMongodVersion returns the MongodVersion field if non-nil, zero value otherwise.

### GetMongodVersionOk

`func (o *BackupSnapshotPart) GetMongodVersionOk() (*string, bool)`

GetMongodVersionOk returns a tuple with the MongodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodVersion

`func (o *BackupSnapshotPart) SetMongodVersion(v string)`

SetMongodVersion sets MongodVersion field to given value.

### HasMongodVersion

`func (o *BackupSnapshotPart) HasMongodVersion() bool`

HasMongodVersion returns a boolean if a field has been set.
### GetReplicaSetName

`func (o *BackupSnapshotPart) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *BackupSnapshotPart) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *BackupSnapshotPart) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *BackupSnapshotPart) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.
### GetReplicaState

`func (o *BackupSnapshotPart) GetReplicaState() string`

GetReplicaState returns the ReplicaState field if non-nil, zero value otherwise.

### GetReplicaStateOk

`func (o *BackupSnapshotPart) GetReplicaStateOk() (*string, bool)`

GetReplicaStateOk returns a tuple with the ReplicaState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaState

`func (o *BackupSnapshotPart) SetReplicaState(v string)`

SetReplicaState sets ReplicaState field to given value.

### HasReplicaState

`func (o *BackupSnapshotPart) HasReplicaState() bool`

HasReplicaState returns a boolean if a field has been set.
### GetStorageSizeBytes

`func (o *BackupSnapshotPart) GetStorageSizeBytes() int64`

GetStorageSizeBytes returns the StorageSizeBytes field if non-nil, zero value otherwise.

### GetStorageSizeBytesOk

`func (o *BackupSnapshotPart) GetStorageSizeBytesOk() (*int64, bool)`

GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeBytes

`func (o *BackupSnapshotPart) SetStorageSizeBytes(v int64)`

SetStorageSizeBytes sets StorageSizeBytes field to given value.

### HasStorageSizeBytes

`func (o *BackupSnapshotPart) HasStorageSizeBytes() bool`

HasStorageSizeBytes returns a boolean if a field has been set.
### GetTypeName

`func (o *BackupSnapshotPart) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *BackupSnapshotPart) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *BackupSnapshotPart) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *BackupSnapshotPart) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


