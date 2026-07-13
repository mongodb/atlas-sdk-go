// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosTier Hardware spec for the dedicated mongos tier.
type MongosTier struct {
	// Instance size of the dedicated mongos hosts (for example, `M30`).
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Per-region placement of the dedicated mongos hosts.
	RegionConfigs *[]MongosRegionConfig `json:"regionConfigs,omitempty"`
}

// NewMongosTier instantiates a new MongosTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosTier() *MongosTier {
	this := MongosTier{}
	return &this
}

// NewMongosTierWithDefaults instantiates a new MongosTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosTierWithDefaults() *MongosTier {
	this := MongosTier{}
	return &this
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise
func (o *MongosTier) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTier) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}

	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *MongosTier) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *MongosTier) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

// GetRegionConfigs returns the RegionConfigs field value if set, zero value otherwise
func (o *MongosTier) GetRegionConfigs() []MongosRegionConfig {
	if o == nil || IsNil(o.RegionConfigs) {
		var ret []MongosRegionConfig
		return ret
	}
	return *o.RegionConfigs
}

// GetRegionConfigsOk returns a tuple with the RegionConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTier) GetRegionConfigsOk() (*[]MongosRegionConfig, bool) {
	if o == nil || IsNil(o.RegionConfigs) {
		return nil, false
	}

	return o.RegionConfigs, true
}

// HasRegionConfigs returns a boolean if a field has been set.
func (o *MongosTier) HasRegionConfigs() bool {
	if o != nil && !IsNil(o.RegionConfigs) {
		return true
	}

	return false
}

// SetRegionConfigs gets a reference to the given []MongosRegionConfig and assigns it to the RegionConfigs field.
func (o *MongosTier) SetRegionConfigs(v []MongosRegionConfig) {
	o.RegionConfigs = &v
}
