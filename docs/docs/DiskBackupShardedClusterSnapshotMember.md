# DiskBackupShardedClusterSnapshotMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Human-readable label that identifies the cloud provider. | [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [readonly] 
**ReplicaSetName** | **string** | Human-readable label that identifies the shard or config host from which MongoDB Cloud took this snapshot. | [readonly] 

## Methods

### NewDiskBackupShardedClusterSnapshotMember

`func NewDiskBackupShardedClusterSnapshotMember(cloudProvider string, id string, replicaSetName string, ) *DiskBackupShardedClusterSnapshotMember`

NewDiskBackupShardedClusterSnapshotMember instantiates a new DiskBackupShardedClusterSnapshotMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupShardedClusterSnapshotMemberWithDefaults

`func NewDiskBackupShardedClusterSnapshotMemberWithDefaults() *DiskBackupShardedClusterSnapshotMember`

NewDiskBackupShardedClusterSnapshotMemberWithDefaults instantiates a new DiskBackupShardedClusterSnapshotMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DiskBackupShardedClusterSnapshotMember) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupShardedClusterSnapshotMember) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupShardedClusterSnapshotMember) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetId

`func (o *DiskBackupShardedClusterSnapshotMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupShardedClusterSnapshotMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupShardedClusterSnapshotMember) SetId(v string)`

SetId sets Id field to given value.

### GetReplicaSetName

`func (o *DiskBackupShardedClusterSnapshotMember) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *DiskBackupShardedClusterSnapshotMember) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *DiskBackupShardedClusterSnapshotMember) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


