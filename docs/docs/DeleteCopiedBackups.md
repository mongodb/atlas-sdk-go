# DeleteCopiedBackups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete. | [optional] 
**RegionName** | Pointer to **string** | Target region for the deleted copy setting whose backup copies you want to delete. Please supply the &#39;Atlas Region&#39; which can be found under [Cloud Providers](https://www.mongodb.com/docs/atlas/reference/cloud-providers/) &#39;regions&#39; link. | [optional] 
**ReplicationSpecId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica setclusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array [Return One Cluster in One Project](#operation/getLegacyCluster). | [optional] 
**ZoneId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Zone Id, do a GET request to Return One Cluster from One Project and consult the replicationSpecs array [Return One Cluster from One Project](#operation/getCluster). | [optional] 

## Methods

### NewDeleteCopiedBackups

`func NewDeleteCopiedBackups() *DeleteCopiedBackups`

NewDeleteCopiedBackups instantiates a new DeleteCopiedBackups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCopiedBackupsWithDefaults

`func NewDeleteCopiedBackupsWithDefaults() *DeleteCopiedBackups`

NewDeleteCopiedBackupsWithDefaults instantiates a new DeleteCopiedBackups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DeleteCopiedBackups) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DeleteCopiedBackups) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DeleteCopiedBackups) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DeleteCopiedBackups) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetRegionName

`func (o *DeleteCopiedBackups) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *DeleteCopiedBackups) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *DeleteCopiedBackups) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *DeleteCopiedBackups) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetReplicationSpecId

`func (o *DeleteCopiedBackups) GetReplicationSpecId() string`

GetReplicationSpecId returns the ReplicationSpecId field if non-nil, zero value otherwise.

### GetReplicationSpecIdOk

`func (o *DeleteCopiedBackups) GetReplicationSpecIdOk() (*string, bool)`

GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecId

`func (o *DeleteCopiedBackups) SetReplicationSpecId(v string)`

SetReplicationSpecId sets ReplicationSpecId field to given value.

### HasReplicationSpecId

`func (o *DeleteCopiedBackups) HasReplicationSpecId() bool`

HasReplicationSpecId returns a boolean if a field has been set.
### GetZoneId

`func (o *DeleteCopiedBackups) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DeleteCopiedBackups) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DeleteCopiedBackups) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DeleteCopiedBackups) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


