// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosRegionConfig Dedicated mongos placement in one region.
type MongosRegionConfig struct {
	// Number of dedicated mongos to run in this region.
	NodeCount *int `json:"nodeCount,omitempty"`
	// Cloud provider on which the dedicated mongos run.
	ProviderName *string `json:"providerName,omitempty"`
	// Region in which the dedicated mongos run.
	RegionName *string `json:"regionName,omitempty"`
}

// NewMongosRegionConfig instantiates a new MongosRegionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosRegionConfig() *MongosRegionConfig {
	this := MongosRegionConfig{}
	return &this
}

// NewMongosRegionConfigWithDefaults instantiates a new MongosRegionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosRegionConfigWithDefaults() *MongosRegionConfig {
	this := MongosRegionConfig{}
	return &this
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise
func (o *MongosRegionConfig) GetNodeCount() int {
	if o == nil || IsNil(o.NodeCount) {
		var ret int
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosRegionConfig) GetNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}

	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *MongosRegionConfig) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int and assigns it to the NodeCount field.
func (o *MongosRegionConfig) SetNodeCount(v int) {
	o.NodeCount = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise
func (o *MongosRegionConfig) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosRegionConfig) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}

	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *MongosRegionConfig) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *MongosRegionConfig) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *MongosRegionConfig) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosRegionConfig) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *MongosRegionConfig) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *MongosRegionConfig) SetRegionName(v string) {
	o.RegionName = &v
}
