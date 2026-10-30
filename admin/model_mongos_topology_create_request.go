// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosTopologyCreateRequest Dedicated mongos topology to attach to a sharded cluster.
type MongosTopologyCreateRequest struct {
	// Human-readable label that identifies the sharded cluster this topology belongs to.
	ClusterName *string `json:"clusterName,omitempty"`
	// Instance size of the dedicated mongos hosts (for example, `M30`).
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Per-region placement of the dedicated mongos hosts.
	RegionPlacements *[]MongosTopologyResourceRegionConfig `json:"regionPlacements,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *MongosTopologyCreateRequest) MarshalJSON() ([]byte, error) {
	type noMethod MongosTopologyCreateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewMongosTopologyCreateRequest instantiates a new MongosTopologyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosTopologyCreateRequest() *MongosTopologyCreateRequest {
	this := MongosTopologyCreateRequest{}
	return &this
}

// NewMongosTopologyCreateRequestWithDefaults instantiates a new MongosTopologyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosTopologyCreateRequestWithDefaults() *MongosTopologyCreateRequest {
	this := MongosTopologyCreateRequest{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *MongosTopologyCreateRequest) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyCreateRequest) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *MongosTopologyCreateRequest) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *MongosTopologyCreateRequest) SetClusterName(v string) {
	o.ClusterName = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterName")
}

// SetClusterNameNil sets ClusterName to an explicit JSON null when marshaled.
func (o *MongosTopologyCreateRequest) SetClusterNameNil() {
	o.ClusterName = nil
	o.NullFields = addNullField(o.NullFields, "ClusterName")
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise
func (o *MongosTopologyCreateRequest) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyCreateRequest) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}

	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *MongosTopologyCreateRequest) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *MongosTopologyCreateRequest) SetInstanceSize(v string) {
	o.InstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "InstanceSize")
}

// SetInstanceSizeNil sets InstanceSize to an explicit JSON null when marshaled.
func (o *MongosTopologyCreateRequest) SetInstanceSizeNil() {
	o.InstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "InstanceSize")
}

// GetRegionPlacements returns the RegionPlacements field value if set, zero value otherwise
func (o *MongosTopologyCreateRequest) GetRegionPlacements() []MongosTopologyResourceRegionConfig {
	if o == nil || IsNil(o.RegionPlacements) {
		var ret []MongosTopologyResourceRegionConfig
		return ret
	}
	return *o.RegionPlacements
}

// GetRegionPlacementsOk returns a tuple with the RegionPlacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyCreateRequest) GetRegionPlacementsOk() (*[]MongosTopologyResourceRegionConfig, bool) {
	if o == nil || IsNil(o.RegionPlacements) {
		return nil, false
	}

	return o.RegionPlacements, true
}

// HasRegionPlacements returns a boolean if a field has been set.
func (o *MongosTopologyCreateRequest) HasRegionPlacements() bool {
	if o != nil && !IsNil(o.RegionPlacements) {
		return true
	}

	return false
}

// SetRegionPlacements gets a reference to the given []MongosTopologyResourceRegionConfig and assigns it to the RegionPlacements field.
func (o *MongosTopologyCreateRequest) SetRegionPlacements(v []MongosTopologyResourceRegionConfig) {
	o.RegionPlacements = &v
	o.NullFields = removeNullField(o.NullFields, "RegionPlacements")
}

// SetRegionPlacementsNil sets RegionPlacements to an explicit JSON null when marshaled.
func (o *MongosTopologyCreateRequest) SetRegionPlacementsNil() {
	o.RegionPlacements = nil
	o.NullFields = addNullField(o.NullFields, "RegionPlacements")
}
