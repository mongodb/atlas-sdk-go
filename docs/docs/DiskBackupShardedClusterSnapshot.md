# DiskBackupShardedClusterSnapshot

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
**Members** | Pointer to [**[]DiskBackupShardedClusterSnapshotMember**](DiskBackupShardedClusterSnapshotMember.md) | List that includes the snapshots and the cloud provider that stores the snapshots. The resource returns this parameter when &#x60;\&quot;type\&quot; : \&quot;SHARDED_CLUSTER\&quot;&#x60;. | [optional] [readonly] 
**MongodVersion** | Pointer to **string** | Version of the MongoDB host that this snapshot backs up. | [optional] [readonly] 
**PolicyItems** | Pointer to **[]string** | List that contains unique identifiers for the policy items. | [optional] [readonly] 
**SnapshotIds** | Pointer to **[]string** | List that contains the unique identifiers of the snapshots created for the shards and config host for a sharded cluster. The resource returns this parameter when &#x60;\&quot;type\&quot;: \&quot;SHARDED_CLUSTER\&quot;&#x60;. These identifiers should match the ones specified in the **members[n].id** parameters. This allows you to map a snapshot to its shard or config host name. | [optional] [readonly] 
**SnapshotType** | Pointer to **string** | Human-readable label that identifies when this snapshot triggers. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the stage of the backup process for this snapshot. | [optional] [readonly] 
**StorageSizeBytes** | Pointer to **int64** | Number of bytes taken to store the backup snapshot. | [optional] [readonly] 
**Type** | Pointer to **string** | Human-readable label that categorizes the cluster as a replica set or sharded cluster. | [optional] [readonly] 

## Methods

### NewDiskBackupShardedClusterSnapshot

`func NewDiskBackupShardedClusterSnapshot() *DiskBackupShardedClusterSnapshot`

NewDiskBackupShardedClusterSnapshot instantiates a new DiskBackupShardedClusterSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupShardedClusterSnapshotWithDefaults

`func NewDiskBackupShardedClusterSnapshotWithDefaults() *DiskBackupShardedClusterSnapshot`

NewDiskBackupShardedClusterSnapshotWithDefaults instantiates a new DiskBackupShardedClusterSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DiskBackupShardedClusterSnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DiskBackupShardedClusterSnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DiskBackupShardedClusterSnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DiskBackupShardedClusterSnapshot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DiskBackupShardedClusterSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiskBackupShardedClusterSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiskBackupShardedClusterSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiskBackupShardedClusterSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DiskBackupShardedClusterSnapshot) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DiskBackupShardedClusterSnapshot) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DiskBackupShardedClusterSnapshot) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DiskBackupShardedClusterSnapshot) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFrequencyType

`func (o *DiskBackupShardedClusterSnapshot) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *DiskBackupShardedClusterSnapshot) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *DiskBackupShardedClusterSnapshot) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### HasFrequencyType

`func (o *DiskBackupShardedClusterSnapshot) HasFrequencyType() bool`

HasFrequencyType returns a boolean if a field has been set.

### GetId

`func (o *DiskBackupShardedClusterSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupShardedClusterSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupShardedClusterSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskBackupShardedClusterSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *DiskBackupShardedClusterSnapshot) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupShardedClusterSnapshot) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupShardedClusterSnapshot) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupShardedClusterSnapshot) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMasterKeyUUID

`func (o *DiskBackupShardedClusterSnapshot) GetMasterKeyUUID() string`

GetMasterKeyUUID returns the MasterKeyUUID field if non-nil, zero value otherwise.

### GetMasterKeyUUIDOk

`func (o *DiskBackupShardedClusterSnapshot) GetMasterKeyUUIDOk() (*string, bool)`

GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyUUID

`func (o *DiskBackupShardedClusterSnapshot) SetMasterKeyUUID(v string)`

SetMasterKeyUUID sets MasterKeyUUID field to given value.

### HasMasterKeyUUID

`func (o *DiskBackupShardedClusterSnapshot) HasMasterKeyUUID() bool`

HasMasterKeyUUID returns a boolean if a field has been set.

### GetMembers

`func (o *DiskBackupShardedClusterSnapshot) GetMembers() []DiskBackupShardedClusterSnapshotMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DiskBackupShardedClusterSnapshot) GetMembersOk() (*[]DiskBackupShardedClusterSnapshotMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DiskBackupShardedClusterSnapshot) SetMembers(v []DiskBackupShardedClusterSnapshotMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DiskBackupShardedClusterSnapshot) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMongodVersion

`func (o *DiskBackupShardedClusterSnapshot) GetMongodVersion() string`

GetMongodVersion returns the MongodVersion field if non-nil, zero value otherwise.

### GetMongodVersionOk

`func (o *DiskBackupShardedClusterSnapshot) GetMongodVersionOk() (*string, bool)`

GetMongodVersionOk returns a tuple with the MongodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodVersion

`func (o *DiskBackupShardedClusterSnapshot) SetMongodVersion(v string)`

SetMongodVersion sets MongodVersion field to given value.

### HasMongodVersion

`func (o *DiskBackupShardedClusterSnapshot) HasMongodVersion() bool`

HasMongodVersion returns a boolean if a field has been set.

### GetPolicyItems

`func (o *DiskBackupShardedClusterSnapshot) GetPolicyItems() []string`

GetPolicyItems returns the PolicyItems field if non-nil, zero value otherwise.

### GetPolicyItemsOk

`func (o *DiskBackupShardedClusterSnapshot) GetPolicyItemsOk() (*[]string, bool)`

GetPolicyItemsOk returns a tuple with the PolicyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyItems

`func (o *DiskBackupShardedClusterSnapshot) SetPolicyItems(v []string)`

SetPolicyItems sets PolicyItems field to given value.

### HasPolicyItems

`func (o *DiskBackupShardedClusterSnapshot) HasPolicyItems() bool`

HasPolicyItems returns a boolean if a field has been set.

### GetSnapshotIds

`func (o *DiskBackupShardedClusterSnapshot) GetSnapshotIds() []string`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *DiskBackupShardedClusterSnapshot) GetSnapshotIdsOk() (*[]string, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIds

`func (o *DiskBackupShardedClusterSnapshot) SetSnapshotIds(v []string)`

SetSnapshotIds sets SnapshotIds field to given value.

### HasSnapshotIds

`func (o *DiskBackupShardedClusterSnapshot) HasSnapshotIds() bool`

HasSnapshotIds returns a boolean if a field has been set.

### GetSnapshotType

`func (o *DiskBackupShardedClusterSnapshot) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *DiskBackupShardedClusterSnapshot) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *DiskBackupShardedClusterSnapshot) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *DiskBackupShardedClusterSnapshot) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetStatus

`func (o *DiskBackupShardedClusterSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiskBackupShardedClusterSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiskBackupShardedClusterSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiskBackupShardedClusterSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageSizeBytes

`func (o *DiskBackupShardedClusterSnapshot) GetStorageSizeBytes() int64`

GetStorageSizeBytes returns the StorageSizeBytes field if non-nil, zero value otherwise.

### GetStorageSizeBytesOk

`func (o *DiskBackupShardedClusterSnapshot) GetStorageSizeBytesOk() (*int64, bool)`

GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeBytes

`func (o *DiskBackupShardedClusterSnapshot) SetStorageSizeBytes(v int64)`

SetStorageSizeBytes sets StorageSizeBytes field to given value.

### HasStorageSizeBytes

`func (o *DiskBackupShardedClusterSnapshot) HasStorageSizeBytes() bool`

HasStorageSizeBytes returns a boolean if a field has been set.

### GetType

`func (o *DiskBackupShardedClusterSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiskBackupShardedClusterSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiskBackupShardedClusterSnapshot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiskBackupShardedClusterSnapshot) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


