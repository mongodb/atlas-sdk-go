// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosTopology Configuration of the mongos routers for a sharded cluster: colocated on the shard nodes or running on a dedicated tier.
type MongosTopology struct {
	// Desired mongos topology. `COLOCATED` runs mongos on the shard nodes; `DEDICATED` runs mongos on a dedicated tier described by `tier`.
	State *string     `json:"state,omitempty"`
	Tier  *MongosTier `json:"tier,omitempty"`
}

// NewMongosTopology instantiates a new MongosTopology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosTopology() *MongosTopology {
	this := MongosTopology{}
	return &this
}

// NewMongosTopologyWithDefaults instantiates a new MongosTopology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosTopologyWithDefaults() *MongosTopology {
	this := MongosTopology{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise
func (o *MongosTopology) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopology) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MongosTopology) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MongosTopology) SetState(v string) {
	o.State = &v
}

// GetTier returns the Tier field value if set, zero value otherwise
func (o *MongosTopology) GetTier() MongosTier {
	if o == nil || IsNil(o.Tier) {
		var ret MongosTier
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopology) GetTierOk() (*MongosTier, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}

	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *MongosTopology) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given MongosTier and assigns it to the Tier field.
func (o *MongosTopology) SetTier(v MongosTier) {
	o.Tier = &v
}
