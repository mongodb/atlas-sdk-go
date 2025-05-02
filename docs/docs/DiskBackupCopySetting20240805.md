# DiskBackupCopySetting20240805

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider that stores the snapshot copy. | [optional] 
**Frequencies** | Pointer to **[]string** | List that describes which types of snapshots to copy. | [optional] 
**RegionName** | Pointer to **string** | Target region to copy snapshots belonging to zoneId. Please supply the &#39;Atlas Region&#39;. | [optional] 
**ShouldCopyOplogs** | Pointer to **bool** | Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores. | [optional] 
**ZoneId** | **string** | Unique 24-hexadecimal digit string that identifies the zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Zone Id, do a GET request to Return One Cluster from One Project and consult the replicationSpecs array. | 

## Methods

### NewDiskBackupCopySetting20240805

`func NewDiskBackupCopySetting20240805(zoneId string, ) *DiskBackupCopySetting20240805`

NewDiskBackupCopySetting20240805 instantiates a new DiskBackupCopySetting20240805 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupCopySetting20240805WithDefaults

`func NewDiskBackupCopySetting20240805WithDefaults() *DiskBackupCopySetting20240805`

NewDiskBackupCopySetting20240805WithDefaults instantiates a new DiskBackupCopySetting20240805 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DiskBackupCopySetting20240805) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupCopySetting20240805) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupCopySetting20240805) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DiskBackupCopySetting20240805) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetFrequencies

`func (o *DiskBackupCopySetting20240805) GetFrequencies() []string`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *DiskBackupCopySetting20240805) GetFrequenciesOk() (*[]string, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *DiskBackupCopySetting20240805) SetFrequencies(v []string)`

SetFrequencies sets Frequencies field to given value.

### HasFrequencies

`func (o *DiskBackupCopySetting20240805) HasFrequencies() bool`

HasFrequencies returns a boolean if a field has been set.
### GetRegionName

`func (o *DiskBackupCopySetting20240805) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *DiskBackupCopySetting20240805) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *DiskBackupCopySetting20240805) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *DiskBackupCopySetting20240805) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetShouldCopyOplogs

`func (o *DiskBackupCopySetting20240805) GetShouldCopyOplogs() bool`

GetShouldCopyOplogs returns the ShouldCopyOplogs field if non-nil, zero value otherwise.

### GetShouldCopyOplogsOk

`func (o *DiskBackupCopySetting20240805) GetShouldCopyOplogsOk() (*bool, bool)`

GetShouldCopyOplogsOk returns a tuple with the ShouldCopyOplogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCopyOplogs

`func (o *DiskBackupCopySetting20240805) SetShouldCopyOplogs(v bool)`

SetShouldCopyOplogs sets ShouldCopyOplogs field to given value.

### HasShouldCopyOplogs

`func (o *DiskBackupCopySetting20240805) HasShouldCopyOplogs() bool`

HasShouldCopyOplogs returns a boolean if a field has been set.
### GetZoneId

`func (o *DiskBackupCopySetting20240805) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DiskBackupCopySetting20240805) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DiskBackupCopySetting20240805) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


