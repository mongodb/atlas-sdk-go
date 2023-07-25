# DiskBackupCopySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider that stores the snapshot copy. | [optional] 
**Frequencies** | Pointer to **[]string** | List that describes which types of snapshots to copy. | [optional] 
**RegionName** | Pointer to **string** | Target region to copy snapshots belonging to replicationSpecId to. Please supply the &#39;Atlas Region&#39; which can be found under [Cloud Providers](https://www.mongodb.com/docs/atlas/reference/cloud-providers/) &#39;regions&#39; link. | [optional] 
**ReplicationSpecId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array [Return One Cluster in One Project](#operation/getLegacyCluster). | [optional] 
**ShouldCopyOplogs** | Pointer to **bool** | Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores. | [optional] 

## Methods

### NewDiskBackupCopySetting

`func NewDiskBackupCopySetting() *DiskBackupCopySetting`

NewDiskBackupCopySetting instantiates a new DiskBackupCopySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupCopySettingWithDefaults

`func NewDiskBackupCopySettingWithDefaults() *DiskBackupCopySetting`

NewDiskBackupCopySettingWithDefaults instantiates a new DiskBackupCopySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DiskBackupCopySetting) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupCopySetting) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupCopySetting) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DiskBackupCopySetting) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetFrequencies

`func (o *DiskBackupCopySetting) GetFrequencies() []string`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *DiskBackupCopySetting) GetFrequenciesOk() (*[]string, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *DiskBackupCopySetting) SetFrequencies(v []string)`

SetFrequencies sets Frequencies field to given value.

### HasFrequencies

`func (o *DiskBackupCopySetting) HasFrequencies() bool`

HasFrequencies returns a boolean if a field has been set.
### GetRegionName

`func (o *DiskBackupCopySetting) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *DiskBackupCopySetting) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *DiskBackupCopySetting) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *DiskBackupCopySetting) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetReplicationSpecId

`func (o *DiskBackupCopySetting) GetReplicationSpecId() string`

GetReplicationSpecId returns the ReplicationSpecId field if non-nil, zero value otherwise.

### GetReplicationSpecIdOk

`func (o *DiskBackupCopySetting) GetReplicationSpecIdOk() (*string, bool)`

GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecId

`func (o *DiskBackupCopySetting) SetReplicationSpecId(v string)`

SetReplicationSpecId sets ReplicationSpecId field to given value.

### HasReplicationSpecId

`func (o *DiskBackupCopySetting) HasReplicationSpecId() bool`

HasReplicationSpecId returns a boolean if a field has been set.
### GetShouldCopyOplogs

`func (o *DiskBackupCopySetting) GetShouldCopyOplogs() bool`

GetShouldCopyOplogs returns the ShouldCopyOplogs field if non-nil, zero value otherwise.

### GetShouldCopyOplogsOk

`func (o *DiskBackupCopySetting) GetShouldCopyOplogsOk() (*bool, bool)`

GetShouldCopyOplogsOk returns a tuple with the ShouldCopyOplogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCopyOplogs

`func (o *DiskBackupCopySetting) SetShouldCopyOplogs(v bool)`

SetShouldCopyOplogs sets ShouldCopyOplogs field to given value.

### HasShouldCopyOplogs

`func (o *DiskBackupCopySetting) HasShouldCopyOplogs() bool`

HasShouldCopyOplogs returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


