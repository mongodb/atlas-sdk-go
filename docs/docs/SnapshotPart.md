# SnapshotPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return. | [optional] [readonly] 
**CompressionSetting** | Pointer to **string** | Human-readable label that identifies the method of compression for the snapshot. | [optional] [readonly] 
**DataSizeBytes** | Pointer to **int64** | Total size of the data stored on each node in the cluster. This parameter expresses its value in bytes. | [optional] [readonly] 
**EncryptionEnabled** | Pointer to **bool** | Flag that indicates whether someone encrypted this snapshot. | [optional] [readonly] 
**FileSizeBytes** | Pointer to **int64** | Number that indicates the total size of the data files in bytes. | [optional] [readonly] 
**MasterKeyUUID** | Pointer to **string** | Unique string that identifies the Key Management Interoperability (KMIP) master key used to encrypt the snapshot data. The resource returns this parameter when &#x60;\&quot;parts.encryptionEnabled\&quot; : true&#x60;. | [optional] [readonly] 
**MongodVersion** | Pointer to **string** | Number that indicates the version of MongoDB that the replica set primary ran when MongoDB Cloud created the snapshot. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set. | [optional] [readonly] 
**StorageSizeBytes** | Pointer to **int64** | Number that indicates the total size of space allocated for document storage. | [optional] [readonly] 
**TypeName** | Pointer to **string** | Human-readable label that identifies the type of server from which MongoDB Cloud took this snapshot. | [optional] [readonly] 

## Methods

### NewSnapshotPart

`func NewSnapshotPart() *SnapshotPart`

NewSnapshotPart instantiates a new SnapshotPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotPartWithDefaults

`func NewSnapshotPartWithDefaults() *SnapshotPart`

NewSnapshotPartWithDefaults instantiates a new SnapshotPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *SnapshotPart) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SnapshotPart) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SnapshotPart) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *SnapshotPart) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCompressionSetting

`func (o *SnapshotPart) GetCompressionSetting() string`

GetCompressionSetting returns the CompressionSetting field if non-nil, zero value otherwise.

### GetCompressionSettingOk

`func (o *SnapshotPart) GetCompressionSettingOk() (*string, bool)`

GetCompressionSettingOk returns a tuple with the CompressionSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionSetting

`func (o *SnapshotPart) SetCompressionSetting(v string)`

SetCompressionSetting sets CompressionSetting field to given value.

### HasCompressionSetting

`func (o *SnapshotPart) HasCompressionSetting() bool`

HasCompressionSetting returns a boolean if a field has been set.

### GetDataSizeBytes

`func (o *SnapshotPart) GetDataSizeBytes() int64`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *SnapshotPart) GetDataSizeBytesOk() (*int64, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *SnapshotPart) SetDataSizeBytes(v int64)`

SetDataSizeBytes sets DataSizeBytes field to given value.

### HasDataSizeBytes

`func (o *SnapshotPart) HasDataSizeBytes() bool`

HasDataSizeBytes returns a boolean if a field has been set.

### GetEncryptionEnabled

`func (o *SnapshotPart) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *SnapshotPart) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *SnapshotPart) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *SnapshotPart) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### GetFileSizeBytes

`func (o *SnapshotPart) GetFileSizeBytes() int64`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *SnapshotPart) GetFileSizeBytesOk() (*int64, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *SnapshotPart) SetFileSizeBytes(v int64)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *SnapshotPart) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### GetMasterKeyUUID

`func (o *SnapshotPart) GetMasterKeyUUID() string`

GetMasterKeyUUID returns the MasterKeyUUID field if non-nil, zero value otherwise.

### GetMasterKeyUUIDOk

`func (o *SnapshotPart) GetMasterKeyUUIDOk() (*string, bool)`

GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyUUID

`func (o *SnapshotPart) SetMasterKeyUUID(v string)`

SetMasterKeyUUID sets MasterKeyUUID field to given value.

### HasMasterKeyUUID

`func (o *SnapshotPart) HasMasterKeyUUID() bool`

HasMasterKeyUUID returns a boolean if a field has been set.

### GetMongodVersion

`func (o *SnapshotPart) GetMongodVersion() string`

GetMongodVersion returns the MongodVersion field if non-nil, zero value otherwise.

### GetMongodVersionOk

`func (o *SnapshotPart) GetMongodVersionOk() (*string, bool)`

GetMongodVersionOk returns a tuple with the MongodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodVersion

`func (o *SnapshotPart) SetMongodVersion(v string)`

SetMongodVersion sets MongodVersion field to given value.

### HasMongodVersion

`func (o *SnapshotPart) HasMongodVersion() bool`

HasMongodVersion returns a boolean if a field has been set.

### GetReplicaSetName

`func (o *SnapshotPart) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *SnapshotPart) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *SnapshotPart) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *SnapshotPart) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

### GetStorageSizeBytes

`func (o *SnapshotPart) GetStorageSizeBytes() int64`

GetStorageSizeBytes returns the StorageSizeBytes field if non-nil, zero value otherwise.

### GetStorageSizeBytesOk

`func (o *SnapshotPart) GetStorageSizeBytesOk() (*int64, bool)`

GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeBytes

`func (o *SnapshotPart) SetStorageSizeBytes(v int64)`

SetStorageSizeBytes sets StorageSizeBytes field to given value.

### HasStorageSizeBytes

`func (o *SnapshotPart) HasStorageSizeBytes() bool`

HasStorageSizeBytes returns a boolean if a field has been set.

### GetTypeName

`func (o *SnapshotPart) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *SnapshotPart) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *SnapshotPart) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *SnapshotPart) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


