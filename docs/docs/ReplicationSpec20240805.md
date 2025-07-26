# ReplicationSpec20240805

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a shard in a Cluster. If you include existing shard replication configurations in the request, you must specify this parameter. If you add a new shard to an existing Cluster, you may specify this parameter. The request deletes any existing shards  in the Cluster that you exclude from the request. This corresponds to Shard ID displayed in the UI. | [optional] [readonly] 
**RegionConfigs** | Pointer to [**[]CloudRegionConfig20240805**](CloudRegionConfig20240805.md) | Hardware specifications for nodes set for a given region. Each **regionConfigs** object must be unique by region and cloud provider within the **replicationSpec**. Each **regionConfigs** object describes the region&#39;s priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each **regionConfigs** object must have either an **analyticsSpecs** object, **electableSpecs** object, or **readOnlySpecs** object. Tenant clusters only require **electableSpecs. Dedicated** clusters can specify any of these specifications, but must have at least one **electableSpecs** object within a **replicationSpec**.  **Example:**  If you set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize\&quot; : \&quot;M30\&quot;&#x60;, set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\&quot; : &#x60;\&quot;M30\&quot;&#x60; if you have electable nodes and &#x60;\&quot;replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize\&quot; : &#x60;\&quot;M30\&quot;&#x60; if you have read-only nodes. | [optional] 
**ZoneId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the zone in a Global Cluster. This value can be used to configure Global Cluster backup policies. | [optional] [readonly] 
**ZoneName** | Pointer to **string** | Human-readable label that describes the zone this shard belongs to in a Global Cluster. Provide this value only if \&quot;clusterType\&quot; : \&quot;GEOSHARDED\&quot; but not \&quot;selfManagedSharding\&quot; : true. | [optional] 

## Methods

### NewReplicationSpec20240805

`func NewReplicationSpec20240805() *ReplicationSpec20240805`

NewReplicationSpec20240805 instantiates a new ReplicationSpec20240805 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationSpec20240805WithDefaults

`func NewReplicationSpec20240805WithDefaults() *ReplicationSpec20240805`

NewReplicationSpec20240805WithDefaults instantiates a new ReplicationSpec20240805 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicationSpec20240805) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicationSpec20240805) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicationSpec20240805) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReplicationSpec20240805) HasId() bool`

HasId returns a boolean if a field has been set.
### GetRegionConfigs

`func (o *ReplicationSpec20240805) GetRegionConfigs() []CloudRegionConfig20240805`

GetRegionConfigs returns the RegionConfigs field if non-nil, zero value otherwise.

### GetRegionConfigsOk

`func (o *ReplicationSpec20240805) GetRegionConfigsOk() (*[]CloudRegionConfig20240805, bool)`

GetRegionConfigsOk returns a tuple with the RegionConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionConfigs

`func (o *ReplicationSpec20240805) SetRegionConfigs(v []CloudRegionConfig20240805)`

SetRegionConfigs sets RegionConfigs field to given value.

### HasRegionConfigs

`func (o *ReplicationSpec20240805) HasRegionConfigs() bool`

HasRegionConfigs returns a boolean if a field has been set.
### GetZoneId

`func (o *ReplicationSpec20240805) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ReplicationSpec20240805) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ReplicationSpec20240805) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ReplicationSpec20240805) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.
### GetZoneName

`func (o *ReplicationSpec20240805) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ReplicationSpec20240805) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ReplicationSpec20240805) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ReplicationSpec20240805) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


