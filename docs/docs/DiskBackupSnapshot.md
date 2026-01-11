# DiskBackupSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud took the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Description** | Pointer to **string** | Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when &#x60;\&quot;status\&quot;: \&quot;onDemand\&quot;&#x60;. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FrequencyType** | Pointer to **string** | Human-readable label that identifies how often this snapshot triggers. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MasterKeyUUID** | Pointer to **string** | Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot. The resource returns this value when &#x60;\&quot;encryptionEnabled\&quot; : true&#x60;. | [optional] [readonly] 
**MongodVersion** | Pointer to **string** | Version of the MongoDB host that this snapshot backs up. | [optional] [readonly] 
**PolicyItems** | Pointer to **[]string** | List that contains unique identifiers for the policy items. | [optional] [readonly] 
**SnapshotType** | Pointer to **string** | Human-readable label that identifies when this snapshot triggers. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the stage of the backup process for this snapshot. | [optional] [readonly] 
**StorageSizeBytes** | Pointer to **int64** | Number of bytes taken to store the backup at time of snapshot. | [optional] [readonly] 
**Type** | Pointer to **string** | Human-readable label that categorizes the cluster as a replica set or sharded cluster. | [optional] [readonly] 
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider. | [optional] [readonly] 
**CopyRegions** | Pointer to **[]string** | List that identifies the regions to which MongoDB Cloud copies the snapshot. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set from which MongoDB Cloud took this snapshot. The resource returns this parameter when &#x60;\&quot;type\&quot;: \&quot;replicaSet\&quot;&#x60;. | [optional] [readonly] 
**ConfigServerType** | Pointer to **string** | Describes a sharded cluster&#39;s config server type. | [optional] [readonly] 
**Members** | Pointer to [**[]DiskBackupShardedClusterSnapshotMember**](DiskBackupShardedClusterSnapshotMember.md) | List that includes the snapshots and the cloud provider that stores the snapshots. The resource returns this parameter when &#x60;\&quot;type\&quot; : \&quot;SHARDED_CLUSTER\&quot;&#x60;. | [optional] [readonly] 
**SnapshotIds** | Pointer to **[]string** | List that contains the unique identifiers of the snapshots created for the shards and config host for a sharded cluster. The resource returns this parameter when &#x60;\&quot;type\&quot;: \&quot;SHARDED_CLUSTER\&quot;&#x60;. These identifiers should match the ones specified in the **members[n].id** parameters. This allows you to map a snapshot to its shard or config host name. | [optional] [readonly] 

## Methods

### NewDiskBackupSnapshot

`func NewDiskBackupSnapshot() *DiskBackupSnapshot`

NewDiskBackupSnapshot instantiates a new DiskBackupSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotWithDefaults

`func NewDiskBackupSnapshotWithDefaults() *DiskBackupSnapshot`

NewDiskBackupSnapshotWithDefaults instantiates a new DiskBackupSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DiskBackupSnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DiskBackupSnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DiskBackupSnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DiskBackupSnapshot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetDescription

`func (o *DiskBackupSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiskBackupSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiskBackupSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiskBackupSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetExpiresAt

`func (o *DiskBackupSnapshot) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DiskBackupSnapshot) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DiskBackupSnapshot) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DiskBackupSnapshot) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.
### GetFrequencyType

`func (o *DiskBackupSnapshot) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *DiskBackupSnapshot) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *DiskBackupSnapshot) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### HasFrequencyType

`func (o *DiskBackupSnapshot) HasFrequencyType() bool`

HasFrequencyType returns a boolean if a field has been set.
### GetId

`func (o *DiskBackupSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskBackupSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLinks

`func (o *DiskBackupSnapshot) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshot) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshot) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshot) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMasterKeyUUID

`func (o *DiskBackupSnapshot) GetMasterKeyUUID() string`

GetMasterKeyUUID returns the MasterKeyUUID field if non-nil, zero value otherwise.

### GetMasterKeyUUIDOk

`func (o *DiskBackupSnapshot) GetMasterKeyUUIDOk() (*string, bool)`

GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyUUID

`func (o *DiskBackupSnapshot) SetMasterKeyUUID(v string)`

SetMasterKeyUUID sets MasterKeyUUID field to given value.

### HasMasterKeyUUID

`func (o *DiskBackupSnapshot) HasMasterKeyUUID() bool`

HasMasterKeyUUID returns a boolean if a field has been set.
### GetMongodVersion

`func (o *DiskBackupSnapshot) GetMongodVersion() string`

GetMongodVersion returns the MongodVersion field if non-nil, zero value otherwise.

### GetMongodVersionOk

`func (o *DiskBackupSnapshot) GetMongodVersionOk() (*string, bool)`

GetMongodVersionOk returns a tuple with the MongodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodVersion

`func (o *DiskBackupSnapshot) SetMongodVersion(v string)`

SetMongodVersion sets MongodVersion field to given value.

### HasMongodVersion

`func (o *DiskBackupSnapshot) HasMongodVersion() bool`

HasMongodVersion returns a boolean if a field has been set.
### GetPolicyItems

`func (o *DiskBackupSnapshot) GetPolicyItems() []string`

GetPolicyItems returns the PolicyItems field if non-nil, zero value otherwise.

### GetPolicyItemsOk

`func (o *DiskBackupSnapshot) GetPolicyItemsOk() (*[]string, bool)`

GetPolicyItemsOk returns a tuple with the PolicyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyItems

`func (o *DiskBackupSnapshot) SetPolicyItems(v []string)`

SetPolicyItems sets PolicyItems field to given value.

### HasPolicyItems

`func (o *DiskBackupSnapshot) HasPolicyItems() bool`

HasPolicyItems returns a boolean if a field has been set.
### GetSnapshotType

`func (o *DiskBackupSnapshot) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *DiskBackupSnapshot) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *DiskBackupSnapshot) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *DiskBackupSnapshot) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.
### GetStatus

`func (o *DiskBackupSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiskBackupSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiskBackupSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiskBackupSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetStorageSizeBytes

`func (o *DiskBackupSnapshot) GetStorageSizeBytes() int64`

GetStorageSizeBytes returns the StorageSizeBytes field if non-nil, zero value otherwise.

### GetStorageSizeBytesOk

`func (o *DiskBackupSnapshot) GetStorageSizeBytesOk() (*int64, bool)`

GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeBytes

`func (o *DiskBackupSnapshot) SetStorageSizeBytes(v int64)`

SetStorageSizeBytes sets StorageSizeBytes field to given value.

### HasStorageSizeBytes

`func (o *DiskBackupSnapshot) HasStorageSizeBytes() bool`

HasStorageSizeBytes returns a boolean if a field has been set.
### GetType

`func (o *DiskBackupSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiskBackupSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiskBackupSnapshot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiskBackupSnapshot) HasType() bool`

HasType returns a boolean if a field has been set.
### GetCloudProvider

`func (o *DiskBackupSnapshot) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupSnapshot) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupSnapshot) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DiskBackupSnapshot) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetCopyRegions

`func (o *DiskBackupSnapshot) GetCopyRegions() []string`

GetCopyRegions returns the CopyRegions field if non-nil, zero value otherwise.

### GetCopyRegionsOk

`func (o *DiskBackupSnapshot) GetCopyRegionsOk() (*[]string, bool)`

GetCopyRegionsOk returns a tuple with the CopyRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRegions

`func (o *DiskBackupSnapshot) SetCopyRegions(v []string)`

SetCopyRegions sets CopyRegions field to given value.

### HasCopyRegions

`func (o *DiskBackupSnapshot) HasCopyRegions() bool`

HasCopyRegions returns a boolean if a field has been set.
### GetReplicaSetName

`func (o *DiskBackupSnapshot) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *DiskBackupSnapshot) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *DiskBackupSnapshot) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *DiskBackupSnapshot) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.
### GetConfigServerType

`func (o *DiskBackupSnapshot) GetConfigServerType() string`

GetConfigServerType returns the ConfigServerType field if non-nil, zero value otherwise.

### GetConfigServerTypeOk

`func (o *DiskBackupSnapshot) GetConfigServerTypeOk() (*string, bool)`

GetConfigServerTypeOk returns a tuple with the ConfigServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigServerType

`func (o *DiskBackupSnapshot) SetConfigServerType(v string)`

SetConfigServerType sets ConfigServerType field to given value.

### HasConfigServerType

`func (o *DiskBackupSnapshot) HasConfigServerType() bool`

HasConfigServerType returns a boolean if a field has been set.
### GetMembers

`func (o *DiskBackupSnapshot) GetMembers() []DiskBackupShardedClusterSnapshotMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DiskBackupSnapshot) GetMembersOk() (*[]DiskBackupShardedClusterSnapshotMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DiskBackupSnapshot) SetMembers(v []DiskBackupShardedClusterSnapshotMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DiskBackupSnapshot) HasMembers() bool`

HasMembers returns a boolean if a field has been set.
### GetSnapshotIds

`func (o *DiskBackupSnapshot) GetSnapshotIds() []string`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *DiskBackupSnapshot) GetSnapshotIdsOk() (*[]string, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIds

`func (o *DiskBackupSnapshot) SetSnapshotIds(v []string)`

SetSnapshotIds sets SnapshotIds field to given value.

### HasSnapshotIds

`func (o *DiskBackupSnapshot) HasSnapshotIds() bool`

HasSnapshotIds returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


