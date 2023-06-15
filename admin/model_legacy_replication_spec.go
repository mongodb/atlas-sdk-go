// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the LegacyReplicationSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegacyReplicationSpec{}

// LegacyReplicationSpec struct for LegacyReplicationSpec
type LegacyReplicationSpec struct {
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Global Cluster.  - If you include existing zones in the request, you must specify this parameter.  - If you add a new zone to an existing Global Cluster, you may specify this parameter. The request deletes any existing zones in a Global Cluster that you exclude from the request.
	Id *string `json:"id,omitempty"`
	// Positive integer that specifies the number of shards to deploy in each specified zone If you set this value to `1` and **clusterType** is `SHARDED`, MongoDB Cloud deploys a single-shard sharded cluster. Don't create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don't provide the same benefits as multi-shard configurations.
	NumShards *int `json:"numShards,omitempty"`
	// Physical location where MongoDB Cloud provisions cluster nodes.
	RegionsConfig *map[string]RegionSpec `json:"regionsConfig,omitempty"`
	// Human-readable label that identifies the zone in a Global Cluster. Provide this value only if **clusterType** is `GEOSHARDED`.
	ZoneName *string `json:"zoneName,omitempty"`
}

// NewLegacyReplicationSpec instantiates a new LegacyReplicationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyReplicationSpec() *LegacyReplicationSpec {
	this := LegacyReplicationSpec{}
	var numShards int = 1
	this.NumShards = &numShards
	return &this
}

// NewLegacyReplicationSpecWithDefaults instantiates a new LegacyReplicationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyReplicationSpecWithDefaults() *LegacyReplicationSpec {
	this := LegacyReplicationSpec{}
	var numShards int = 1
	this.NumShards = &numShards
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LegacyReplicationSpec) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyReplicationSpec) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LegacyReplicationSpec) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LegacyReplicationSpec) SetId(v string) {
	o.Id = &v
}

// GetNumShards returns the NumShards field value if set, zero value otherwise.
func (o *LegacyReplicationSpec) GetNumShards() int {
	if o == nil || IsNil(o.NumShards) {
		var ret int
		return ret
	}
	return *o.NumShards
}

// GetNumShardsOk returns a tuple with the NumShards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyReplicationSpec) GetNumShardsOk() (*int, bool) {
	if o == nil || IsNil(o.NumShards) {
		return nil, false
	}
	return o.NumShards, true
}

// HasNumShards returns a boolean if a field has been set.
func (o *LegacyReplicationSpec) HasNumShards() bool {
	if o != nil && !IsNil(o.NumShards) {
		return true
	}

	return false
}

// SetNumShards gets a reference to the given int and assigns it to the NumShards field.
func (o *LegacyReplicationSpec) SetNumShards(v int) {
	o.NumShards = &v
}

// GetRegionsConfig returns the RegionsConfig field value if set, zero value otherwise.
func (o *LegacyReplicationSpec) GetRegionsConfig() map[string]RegionSpec {
	if o == nil || IsNil(o.RegionsConfig) {
		var ret map[string]RegionSpec
		return ret
	}
	return *o.RegionsConfig
}

// GetRegionsConfigOk returns a tuple with the RegionsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyReplicationSpec) GetRegionsConfigOk() (*map[string]RegionSpec, bool) {
	if o == nil || IsNil(o.RegionsConfig) {
		return nil, false
	}
	return o.RegionsConfig, true
}

// HasRegionsConfig returns a boolean if a field has been set.
func (o *LegacyReplicationSpec) HasRegionsConfig() bool {
	if o != nil && !IsNil(o.RegionsConfig) {
		return true
	}

	return false
}

// SetRegionsConfig gets a reference to the given map[string]RegionSpec and assigns it to the RegionsConfig field.
func (o *LegacyReplicationSpec) SetRegionsConfig(v map[string]RegionSpec) {
	o.RegionsConfig = &v
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise.
func (o *LegacyReplicationSpec) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName) {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyReplicationSpec) GetZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneName) {
		return nil, false
	}
	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *LegacyReplicationSpec) HasZoneName() bool {
	if o != nil && !IsNil(o.ZoneName) {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *LegacyReplicationSpec) SetZoneName(v string) {
	o.ZoneName = &v
}

func (o LegacyReplicationSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o LegacyReplicationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NumShards) {
		toSerialize["numShards"] = o.NumShards
	}
	if !IsNil(o.RegionsConfig) {
		toSerialize["regionsConfig"] = o.RegionsConfig
	}
	if !IsNil(o.ZoneName) {
		toSerialize["zoneName"] = o.ZoneName
	}
	return toSerialize, nil
}

type NullableLegacyReplicationSpec struct {
	value *LegacyReplicationSpec
	isSet bool
}

func (v NullableLegacyReplicationSpec) Get() *LegacyReplicationSpec {
	return v.value
}

func (v *NullableLegacyReplicationSpec) Set(val *LegacyReplicationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyReplicationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyReplicationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyReplicationSpec(val *LegacyReplicationSpec) *NullableLegacyReplicationSpec {
	return &NullableLegacyReplicationSpec{value: val, isSet: true}
}

func (v NullableLegacyReplicationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyReplicationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
