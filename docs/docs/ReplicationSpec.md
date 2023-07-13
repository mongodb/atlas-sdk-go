# ReplicationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Multi-Cloud Cluster. If you include existing zones in the request, you must specify this parameter. If you add a new zone to an existing Multi-Cloud Cluster, you may specify this parameter. The request deletes any existing zones in the Multi-Cloud Cluster that you exclude from the request. | [optional] [readonly] 
**NumShards** | Pointer to **int** | Positive integer that specifies the number of shards to deploy in each specified zone. If you set this value to &#x60;1&#x60; and **clusterType** is &#x60;SHARDED&#x60;, MongoDB Cloud deploys a single-shard sharded cluster. Don&#39;t create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don&#39;t provide the same benefits as multi-shard configurations. | [optional] 
**RegionConfigs** | Pointer to [**[]CloudRegionConfig**](CloudRegionConfig.md) | Hardware specifications for nodes set for a given region. Each **regionConfigs** object describes the region&#39;s priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each **regionConfigs** object must have either an **analyticsSpecs** object, **electableSpecs** object, or **readOnlySpecs** object. Tenant clusters only require **electableSpecs. Dedicated** clusters can specify any of these specifications, but must have at least one **electableSpecs** object within a **replicationSpec**. Every hardware specification must use the same **instanceSize**.  **Example:**  If you set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize\&quot; : \&quot;M30\&quot;&#x60;, set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\&quot; : &#x60;\&quot;M30\&quot;&#x60; if you have electable nodes and &#x60;\&quot;replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize\&quot; : &#x60;\&quot;M30\&quot;&#x60; if you have read-only nodes. | [optional] 
**ZoneName** | Pointer to **string** | Human-readable label that identifies the zone in a Global Cluster. Provide this value only if &#x60;\&quot;clusterType\&quot; : \&quot;GEOSHARDED\&quot;&#x60;. | [optional] 

## Methods

### NewReplicationSpec

`func NewReplicationSpec() *ReplicationSpec`

NewReplicationSpec instantiates a new ReplicationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationSpecWithDefaults

`func NewReplicationSpecWithDefaults() *ReplicationSpec`

NewReplicationSpecWithDefaults instantiates a new ReplicationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicationSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicationSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicationSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReplicationSpec) HasId() bool`

HasId returns a boolean if a field has been set.
### GetNumShards

`func (o *ReplicationSpec) GetNumShards() int`

GetNumShards returns the NumShards field if non-nil, zero value otherwise.

### GetNumShardsOk

`func (o *ReplicationSpec) GetNumShardsOk() (*int, bool)`

GetNumShardsOk returns a tuple with the NumShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumShards

`func (o *ReplicationSpec) SetNumShards(v int)`

SetNumShards sets NumShards field to given value.

### HasNumShards

`func (o *ReplicationSpec) HasNumShards() bool`

HasNumShards returns a boolean if a field has been set.
### GetRegionConfigs

`func (o *ReplicationSpec) GetRegionConfigs() []CloudRegionConfig`

GetRegionConfigs returns the RegionConfigs field if non-nil, zero value otherwise.

### GetRegionConfigsOk

`func (o *ReplicationSpec) GetRegionConfigsOk() (*[]CloudRegionConfig, bool)`

GetRegionConfigsOk returns a tuple with the RegionConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionConfigs

`func (o *ReplicationSpec) SetRegionConfigs(v []CloudRegionConfig)`

SetRegionConfigs sets RegionConfigs field to given value.

### HasRegionConfigs

`func (o *ReplicationSpec) HasRegionConfigs() bool`

HasRegionConfigs returns a boolean if a field has been set.
### GetZoneName

`func (o *ReplicationSpec) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ReplicationSpec) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ReplicationSpec) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ReplicationSpec) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


