# LegacyReplicationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Global Cluster.  - If you include existing zones in the request, you must specify this parameter.  - If you add a new zone to an existing Global Cluster, you may specify this parameter. The request deletes any existing zones in a Global Cluster that you exclude from the request. | [optional] 
**NumShards** | Pointer to **int** | Positive integer that specifies the number of shards to deploy in each specified zone If you set this value to &#x60;1&#x60; and **clusterType** is &#x60;SHARDED&#x60;, MongoDB Cloud deploys a single-shard sharded cluster. Don&#39;t create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don&#39;t provide the same benefits as multi-shard configurations. | [optional] [default to 1]
**RegionsConfig** | Pointer to [**map[string]RegionSpec**](RegionSpec.md) | Physical location where MongoDB Cloud provisions cluster nodes. | [optional] 
**ZoneName** | Pointer to **string** | Human-readable label that identifies the zone in a Global Cluster. Provide this value only if **clusterType** is &#x60;GEOSHARDED&#x60;. | [optional] 

## Methods

### NewLegacyReplicationSpec

`func NewLegacyReplicationSpec() *LegacyReplicationSpec`

NewLegacyReplicationSpec instantiates a new LegacyReplicationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyReplicationSpecWithDefaults

`func NewLegacyReplicationSpecWithDefaults() *LegacyReplicationSpec`

NewLegacyReplicationSpecWithDefaults instantiates a new LegacyReplicationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LegacyReplicationSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacyReplicationSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacyReplicationSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegacyReplicationSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumShards

`func (o *LegacyReplicationSpec) GetNumShards() int`

GetNumShards returns the NumShards field if non-nil, zero value otherwise.

### GetNumShardsOk

`func (o *LegacyReplicationSpec) GetNumShardsOk() (*int, bool)`

GetNumShardsOk returns a tuple with the NumShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumShards

`func (o *LegacyReplicationSpec) SetNumShards(v int)`

SetNumShards sets NumShards field to given value.

### HasNumShards

`func (o *LegacyReplicationSpec) HasNumShards() bool`

HasNumShards returns a boolean if a field has been set.

### GetRegionsConfig

`func (o *LegacyReplicationSpec) GetRegionsConfig() map[string]RegionSpec`

GetRegionsConfig returns the RegionsConfig field if non-nil, zero value otherwise.

### GetRegionsConfigOk

`func (o *LegacyReplicationSpec) GetRegionsConfigOk() (*map[string]RegionSpec, bool)`

GetRegionsConfigOk returns a tuple with the RegionsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionsConfig

`func (o *LegacyReplicationSpec) SetRegionsConfig(v map[string]RegionSpec)`

SetRegionsConfig sets RegionsConfig field to given value.

### HasRegionsConfig

`func (o *LegacyReplicationSpec) HasRegionsConfig() bool`

HasRegionsConfig returns a boolean if a field has been set.

### GetZoneName

`func (o *LegacyReplicationSpec) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *LegacyReplicationSpec) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *LegacyReplicationSpec) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *LegacyReplicationSpec) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


