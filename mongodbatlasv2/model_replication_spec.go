/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ReplicationSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicationSpec{}

// ReplicationSpec Details that explain how MongoDB Cloud replicates data on the specified MongoDB database.
type ReplicationSpec struct {
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Multi-Cloud Cluster. If you include existing zones in the request, you must specify this parameter. If you add a new zone to an existing Multi-Cloud Cluster, you may specify this parameter. The request deletes any existing zones in the Multi-Cloud Cluster that you exclude from the request.
	Id *string `json:"id,omitempty"`
	// Positive integer that specifies the number of shards to deploy in each specified zone. If you set this value to `1` and **clusterType** is `SHARDED`, MongoDB Cloud deploys a single-shard sharded cluster. Don't create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don't provide the same benefits as multi-shard configurations.
	NumShards *int32 `json:"numShards,omitempty"`
	// Hardware specifications for nodes set for a given region. Each **regionConfigs** object describes the region's priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each **regionConfigs** object must have either an **analyticsSpecs** object, **electableSpecs** object, or **readOnlySpecs** object. Tenant clusters only require **electableSpecs. Dedicated** clusters can specify any of these specifications, but must have at least one **electableSpecs** object within a **replicationSpec**. Every hardware specification must use the same **instanceSize**.  **Example:**  If you set `\"replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize\" : \"M30\"`, set `\"replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\" : `\"M30\"` if you have electable nodes and `\"replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize\" : `\"M30\"` if you have read-only nodes.
	RegionConfigs []RegionConfig `json:"regionConfigs,omitempty"`
	// Human-readable label that identifies the zone in a Global Cluster. Provide this value only if `\"clusterType\" : \"GEOSHARDED\"`.
	ZoneName *string `json:"zoneName,omitempty"`
}

// NewReplicationSpec instantiates a new ReplicationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationSpec() *ReplicationSpec {
	this := ReplicationSpec{}
	return &this
}

// NewReplicationSpecWithDefaults instantiates a new ReplicationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationSpecWithDefaults() *ReplicationSpec {
	this := ReplicationSpec{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReplicationSpec) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSpec) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReplicationSpec) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReplicationSpec) SetId(v string) {
	o.Id = &v
}

// GetNumShards returns the NumShards field value if set, zero value otherwise.
func (o *ReplicationSpec) GetNumShards() int32 {
	if o == nil || IsNil(o.NumShards) {
		var ret int32
		return ret
	}
	return *o.NumShards
}

// GetNumShardsOk returns a tuple with the NumShards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSpec) GetNumShardsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumShards) {
		return nil, false
	}
	return o.NumShards, true
}

// HasNumShards returns a boolean if a field has been set.
func (o *ReplicationSpec) HasNumShards() bool {
	if o != nil && !IsNil(o.NumShards) {
		return true
	}

	return false
}

// SetNumShards gets a reference to the given int32 and assigns it to the NumShards field.
func (o *ReplicationSpec) SetNumShards(v int32) {
	o.NumShards = &v
}

// GetRegionConfigs returns the RegionConfigs field value if set, zero value otherwise.
func (o *ReplicationSpec) GetRegionConfigs() []RegionConfig {
	if o == nil || IsNil(o.RegionConfigs) {
		var ret []RegionConfig
		return ret
	}
	return o.RegionConfigs
}

// GetRegionConfigsOk returns a tuple with the RegionConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSpec) GetRegionConfigsOk() ([]RegionConfig, bool) {
	if o == nil || IsNil(o.RegionConfigs) {
		return nil, false
	}
	return o.RegionConfigs, true
}

// HasRegionConfigs returns a boolean if a field has been set.
func (o *ReplicationSpec) HasRegionConfigs() bool {
	if o != nil && !IsNil(o.RegionConfigs) {
		return true
	}

	return false
}

// SetRegionConfigs gets a reference to the given []RegionConfig and assigns it to the RegionConfigs field.
func (o *ReplicationSpec) SetRegionConfigs(v []RegionConfig) {
	o.RegionConfigs = v
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise.
func (o *ReplicationSpec) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName) {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSpec) GetZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneName) {
		return nil, false
	}
	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *ReplicationSpec) HasZoneName() bool {
	if o != nil && !IsNil(o.ZoneName) {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *ReplicationSpec) SetZoneName(v string) {
	o.ZoneName = &v
}

func (o ReplicationSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ReplicationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumShards) {
		toSerialize["numShards"] = o.NumShards
	}
	if !IsNil(o.RegionConfigs) {
		toSerialize["regionConfigs"] = o.RegionConfigs
	}
	if !IsNil(o.ZoneName) {
		toSerialize["zoneName"] = o.ZoneName
	}
	return toSerialize, nil
}

type NullableReplicationSpec struct {
	value *ReplicationSpec
	isSet bool
}

func (v NullableReplicationSpec) Get() *ReplicationSpec {
	return v.value
}

func (v *NullableReplicationSpec) Set(val *ReplicationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationSpec(val *ReplicationSpec) *NullableReplicationSpec {
	return &NullableReplicationSpec{value: val, isSet: true}
}

func (v NullableReplicationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


