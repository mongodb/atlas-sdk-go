// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosTopologyResourceRegionConfig Dedicated mongos placement in one region.
type MongosTopologyResourceRegionConfig struct {
	// Number of dedicated mongos to run in this region.
	NodeCount *int `json:"nodeCount,omitempty"`
	// Cloud provider on which the dedicated mongos run.
	ProviderName *string `json:"providerName,omitempty"`
	// Region in which the dedicated mongos run.
	RegionName *string `json:"regionName,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *MongosTopologyResourceRegionConfig) MarshalJSON() ([]byte, error) {
	type noMethod MongosTopologyResourceRegionConfig
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewMongosTopologyResourceRegionConfig instantiates a new MongosTopologyResourceRegionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosTopologyResourceRegionConfig() *MongosTopologyResourceRegionConfig {
	this := MongosTopologyResourceRegionConfig{}
	return &this
}

// NewMongosTopologyResourceRegionConfigWithDefaults instantiates a new MongosTopologyResourceRegionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosTopologyResourceRegionConfigWithDefaults() *MongosTopologyResourceRegionConfig {
	this := MongosTopologyResourceRegionConfig{}
	return &this
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise
func (o *MongosTopologyResourceRegionConfig) GetNodeCount() int {
	if o == nil || IsNil(o.NodeCount) {
		var ret int
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResourceRegionConfig) GetNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}

	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *MongosTopologyResourceRegionConfig) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int and assigns it to the NodeCount field.
func (o *MongosTopologyResourceRegionConfig) SetNodeCount(v int) {
	o.NodeCount = &v
	o.NullFields = removeNullField(o.NullFields, "NodeCount")
}

// SetNodeCountNil sets NodeCount to an explicit JSON null when marshaled.
func (o *MongosTopologyResourceRegionConfig) SetNodeCountNil() {
	o.NodeCount = nil
	o.NullFields = addNullField(o.NullFields, "NodeCount")
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise
func (o *MongosTopologyResourceRegionConfig) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResourceRegionConfig) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}

	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *MongosTopologyResourceRegionConfig) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *MongosTopologyResourceRegionConfig) SetProviderName(v string) {
	o.ProviderName = &v
	o.NullFields = removeNullField(o.NullFields, "ProviderName")
}

// SetProviderNameNil sets ProviderName to an explicit JSON null when marshaled.
func (o *MongosTopologyResourceRegionConfig) SetProviderNameNil() {
	o.ProviderName = nil
	o.NullFields = addNullField(o.NullFields, "ProviderName")
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *MongosTopologyResourceRegionConfig) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResourceRegionConfig) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *MongosTopologyResourceRegionConfig) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *MongosTopologyResourceRegionConfig) SetRegionName(v string) {
	o.RegionName = &v
	o.NullFields = removeNullField(o.NullFields, "RegionName")
}

// SetRegionNameNil sets RegionName to an explicit JSON null when marshaled.
func (o *MongosTopologyResourceRegionConfig) SetRegionNameNil() {
	o.RegionName = nil
	o.NullFields = addNullField(o.NullFields, "RegionName")
}
