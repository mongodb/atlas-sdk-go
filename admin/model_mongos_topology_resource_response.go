// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosTopologyResourceResponse Configuration of the dedicated mongos routers for a sharded cluster, identified by the name of its associated cluster.
type MongosTopologyResourceResponse struct {
	// Human-readable label that identifies the sharded cluster this topology belongs to.
	ClusterName *string `json:"clusterName,omitempty"`
	// Instance size of the dedicated mongos hosts (for example, `M30`).
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Per-region placement of the dedicated mongos hosts.
	RegionConfigs *[]MongosTopologyResourceRegionConfig `json:"regionConfigs,omitempty"`
	// Lifecycle status of the dedicated mongos topology resource.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *MongosTopologyResourceResponse) MarshalJSON() ([]byte, error) {
	type noMethod MongosTopologyResourceResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewMongosTopologyResourceResponse instantiates a new MongosTopologyResourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosTopologyResourceResponse() *MongosTopologyResourceResponse {
	this := MongosTopologyResourceResponse{}
	return &this
}

// NewMongosTopologyResourceResponseWithDefaults instantiates a new MongosTopologyResourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosTopologyResourceResponseWithDefaults() *MongosTopologyResourceResponse {
	this := MongosTopologyResourceResponse{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *MongosTopologyResourceResponse) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResourceResponse) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *MongosTopologyResourceResponse) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *MongosTopologyResourceResponse) SetClusterName(v string) {
	o.ClusterName = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterName")
}

// SetClusterNameNil sets ClusterName to an explicit JSON null when marshaled.
func (o *MongosTopologyResourceResponse) SetClusterNameNil() {
	o.ClusterName = nil
	o.NullFields = addNullField(o.NullFields, "ClusterName")
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise
func (o *MongosTopologyResourceResponse) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResourceResponse) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}

	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *MongosTopologyResourceResponse) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *MongosTopologyResourceResponse) SetInstanceSize(v string) {
	o.InstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "InstanceSize")
}

// SetInstanceSizeNil sets InstanceSize to an explicit JSON null when marshaled.
func (o *MongosTopologyResourceResponse) SetInstanceSizeNil() {
	o.InstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "InstanceSize")
}

// GetRegionConfigs returns the RegionConfigs field value if set, zero value otherwise
func (o *MongosTopologyResourceResponse) GetRegionConfigs() []MongosTopologyResourceRegionConfig {
	if o == nil || IsNil(o.RegionConfigs) {
		var ret []MongosTopologyResourceRegionConfig
		return ret
	}
	return *o.RegionConfigs
}

// GetRegionConfigsOk returns a tuple with the RegionConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResourceResponse) GetRegionConfigsOk() (*[]MongosTopologyResourceRegionConfig, bool) {
	if o == nil || IsNil(o.RegionConfigs) {
		return nil, false
	}

	return o.RegionConfigs, true
}

// HasRegionConfigs returns a boolean if a field has been set.
func (o *MongosTopologyResourceResponse) HasRegionConfigs() bool {
	if o != nil && !IsNil(o.RegionConfigs) {
		return true
	}

	return false
}

// SetRegionConfigs gets a reference to the given []MongosTopologyResourceRegionConfig and assigns it to the RegionConfigs field.
func (o *MongosTopologyResourceResponse) SetRegionConfigs(v []MongosTopologyResourceRegionConfig) {
	o.RegionConfigs = &v
	o.NullFields = removeNullField(o.NullFields, "RegionConfigs")
}

// SetRegionConfigsNil sets RegionConfigs to an explicit JSON null when marshaled.
func (o *MongosTopologyResourceResponse) SetRegionConfigsNil() {
	o.RegionConfigs = nil
	o.NullFields = addNullField(o.NullFields, "RegionConfigs")
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *MongosTopologyResourceResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResourceResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MongosTopologyResourceResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MongosTopologyResourceResponse) SetStatus(v string) {
	o.Status = &v
	o.NullFields = removeNullField(o.NullFields, "Status")
}

// SetStatusNil sets Status to an explicit JSON null when marshaled.
func (o *MongosTopologyResourceResponse) SetStatusNil() {
	o.Status = nil
	o.NullFields = addNullField(o.NullFields, "Status")
}
