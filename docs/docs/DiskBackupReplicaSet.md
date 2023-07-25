# DiskBackupReplicaSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider that stores this snapshot. The resource returns this parameter when &#x60;\&quot;type\&quot;: \&quot;replicaSet\&quot;&#x60;. | [optional] [readonly] 
**CopyRegions** | Pointer to **[]string** | List that identifies the regions to which MongoDB Cloud copies the snapshot. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud took the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Description** | Pointer to **string** | Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when &#x60;\&quot;status\&quot;: \&quot;onDemand\&quot;&#x60;. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FrequencyType** | Pointer to **string** | Human-readable label that identifies how often this snapshot triggers. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MasterKeyUUID** | Pointer to **string** | Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot. The resource returns this value when &#x60;\&quot;encryptionEnabled\&quot; : true&#x60;. | [optional] [readonly] 
**MongodVersion** | Pointer to **string** | Version of the MongoDB host that this snapshot backs up. | [optional] [readonly] 
**PolicyItems** | Pointer to **[]string** | List that contains unique identifiers for the policy items. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set from which MongoDB Cloud took this snapshot. The resource returns this parameter when &#x60;\&quot;type\&quot;: \&quot;replicaSet\&quot;&#x60;. | [optional] [readonly] 
**SnapshotType** | Pointer to **string** | Human-readable label that identifies when this snapshot triggers. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the stage of the backup process for this snapshot. | [optional] [readonly] 
**StorageSizeBytes** | Pointer to **int64** | Number of bytes taken to store the backup snapshot. | [optional] [readonly] 
**Type** | Pointer to **string** | Human-readable label that categorizes the cluster as a replica set or sharded cluster. | [optional] [readonly] 

## Methods

### NewDiskBackupReplicaSet

`func NewDiskBackupReplicaSet() *DiskBackupReplicaSet`

NewDiskBackupReplicaSet instantiates a new DiskBackupReplicaSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupReplicaSetWithDefaults

`func NewDiskBackupReplicaSetWithDefaults() *DiskBackupReplicaSet`

NewDiskBackupReplicaSetWithDefaults instantiates a new DiskBackupReplicaSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DiskBackupReplicaSet) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupReplicaSet) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupReplicaSet) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DiskBackupReplicaSet) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetCopyRegions

`func (o *DiskBackupReplicaSet) GetCopyRegions() []string`

GetCopyRegions returns the CopyRegions field if non-nil, zero value otherwise.

### GetCopyRegionsOk

`func (o *DiskBackupReplicaSet) GetCopyRegionsOk() (*[]string, bool)`

GetCopyRegionsOk returns a tuple with the CopyRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRegions

`func (o *DiskBackupReplicaSet) SetCopyRegions(v []string)`

SetCopyRegions sets CopyRegions field to given value.

### HasCopyRegions

`func (o *DiskBackupReplicaSet) HasCopyRegions() bool`

HasCopyRegions returns a boolean if a field has been set.
### GetCreatedAt

`func (o *DiskBackupReplicaSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DiskBackupReplicaSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DiskBackupReplicaSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DiskBackupReplicaSet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetDescription

`func (o *DiskBackupReplicaSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiskBackupReplicaSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiskBackupReplicaSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiskBackupReplicaSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetExpiresAt

`func (o *DiskBackupReplicaSet) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DiskBackupReplicaSet) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DiskBackupReplicaSet) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DiskBackupReplicaSet) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.
### GetFrequencyType

`func (o *DiskBackupReplicaSet) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *DiskBackupReplicaSet) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *DiskBackupReplicaSet) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### HasFrequencyType

`func (o *DiskBackupReplicaSet) HasFrequencyType() bool`

HasFrequencyType returns a boolean if a field has been set.
### GetId

`func (o *DiskBackupReplicaSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupReplicaSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupReplicaSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskBackupReplicaSet) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLinks

`func (o *DiskBackupReplicaSet) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupReplicaSet) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupReplicaSet) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupReplicaSet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMasterKeyUUID

`func (o *DiskBackupReplicaSet) GetMasterKeyUUID() string`

GetMasterKeyUUID returns the MasterKeyUUID field if non-nil, zero value otherwise.

### GetMasterKeyUUIDOk

`func (o *DiskBackupReplicaSet) GetMasterKeyUUIDOk() (*string, bool)`

GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyUUID

`func (o *DiskBackupReplicaSet) SetMasterKeyUUID(v string)`

SetMasterKeyUUID sets MasterKeyUUID field to given value.

### HasMasterKeyUUID

`func (o *DiskBackupReplicaSet) HasMasterKeyUUID() bool`

HasMasterKeyUUID returns a boolean if a field has been set.
### GetMongodVersion

`func (o *DiskBackupReplicaSet) GetMongodVersion() string`

GetMongodVersion returns the MongodVersion field if non-nil, zero value otherwise.

### GetMongodVersionOk

`func (o *DiskBackupReplicaSet) GetMongodVersionOk() (*string, bool)`

GetMongodVersionOk returns a tuple with the MongodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodVersion

`func (o *DiskBackupReplicaSet) SetMongodVersion(v string)`

SetMongodVersion sets MongodVersion field to given value.

### HasMongodVersion

`func (o *DiskBackupReplicaSet) HasMongodVersion() bool`

HasMongodVersion returns a boolean if a field has been set.
### GetPolicyItems

`func (o *DiskBackupReplicaSet) GetPolicyItems() []string`

GetPolicyItems returns the PolicyItems field if non-nil, zero value otherwise.

### GetPolicyItemsOk

`func (o *DiskBackupReplicaSet) GetPolicyItemsOk() (*[]string, bool)`

GetPolicyItemsOk returns a tuple with the PolicyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyItems

`func (o *DiskBackupReplicaSet) SetPolicyItems(v []string)`

SetPolicyItems sets PolicyItems field to given value.

### HasPolicyItems

`func (o *DiskBackupReplicaSet) HasPolicyItems() bool`

HasPolicyItems returns a boolean if a field has been set.
### GetReplicaSetName

`func (o *DiskBackupReplicaSet) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *DiskBackupReplicaSet) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *DiskBackupReplicaSet) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *DiskBackupReplicaSet) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.
### GetSnapshotType

`func (o *DiskBackupReplicaSet) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *DiskBackupReplicaSet) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *DiskBackupReplicaSet) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *DiskBackupReplicaSet) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.
### GetStatus

`func (o *DiskBackupReplicaSet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiskBackupReplicaSet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiskBackupReplicaSet) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiskBackupReplicaSet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetStorageSizeBytes

`func (o *DiskBackupReplicaSet) GetStorageSizeBytes() int64`

GetStorageSizeBytes returns the StorageSizeBytes field if non-nil, zero value otherwise.

### GetStorageSizeBytesOk

`func (o *DiskBackupReplicaSet) GetStorageSizeBytesOk() (*int64, bool)`

GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeBytes

`func (o *DiskBackupReplicaSet) SetStorageSizeBytes(v int64)`

SetStorageSizeBytes sets StorageSizeBytes field to given value.

### HasStorageSizeBytes

`func (o *DiskBackupReplicaSet) HasStorageSizeBytes() bool`

HasStorageSizeBytes returns a boolean if a field has been set.
### GetType

`func (o *DiskBackupReplicaSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiskBackupReplicaSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiskBackupReplicaSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiskBackupReplicaSet) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


