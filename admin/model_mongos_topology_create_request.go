// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosTopologyCreateRequest Dedicated mongos topology to attach to a sharded cluster.
type MongosTopologyCreateRequest struct {
	AutoScaling *MongosAutoScaling `json:"autoScaling,omitempty"`
	// Human-readable label that identifies the sharded cluster this topology belongs to.
	ClusterName *string `json:"clusterName,omitempty"`
	// Autoscaling configuration in effect for the dedicated mongos tier. Resolves the defaults omitted from `autoScaling`.
	// Read only field.
	EffectiveAutoScaling *MongosEffectiveAutoScaling `json:"effectiveAutoScaling,omitempty"`
	// Instance size the dedicated mongos hosts are running at. Can differ from `instanceSize` when autoscaling is enabled.
	// Read only field.
	EffectiveInstanceSize *string `json:"effectiveInstanceSize,omitempty"`
	// Per-region placement of the dedicated mongos hosts in effect. Resolves the defaults omitted from `regionPlacements`.
	// Read only field.
	EffectiveRegionPlacements *[]MongosRegionPlacement `json:"effectiveRegionPlacements,omitempty"`
	// Instance size of the dedicated mongos hosts (for example, `M30`).
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Per-region placement of the dedicated mongos hosts.
	RegionPlacements *[]MongosRegionPlacement `json:"regionPlacements,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o MongosTopologyCreateRequest) MarshalJSON() ([]byte, error) {
	type noMethod MongosTopologyCreateRequest
	return marshalWithNullFields(noMethod(o), o.NullFields)
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

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise
func (o *MongosTopologyCreateRequest) GetAutoScaling() MongosAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret MongosAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyCreateRequest) GetAutoScalingOk() (*MongosAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}

	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *MongosTopologyCreateRequest) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given MongosAutoScaling and assigns it to the AutoScaling field.
func (o *MongosTopologyCreateRequest) SetAutoScaling(v MongosAutoScaling) {
	o.AutoScaling = &v
	o.NullFields = removeNullField(o.NullFields, "AutoScaling")
}

// SetAutoScalingNil sets AutoScaling to an explicit JSON null when marshaled.
func (o *MongosTopologyCreateRequest) SetAutoScalingNil() {
	o.AutoScaling = nil
	o.NullFields = addNullField(o.NullFields, "AutoScaling")
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

// GetEffectiveAutoScaling returns the EffectiveAutoScaling field value if set, zero value otherwise
func (o *MongosTopologyCreateRequest) GetEffectiveAutoScaling() MongosEffectiveAutoScaling {
	if o == nil || IsNil(o.EffectiveAutoScaling) {
		var ret MongosEffectiveAutoScaling
		return ret
	}
	return *o.EffectiveAutoScaling
}

// GetEffectiveAutoScalingOk returns a tuple with the EffectiveAutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyCreateRequest) GetEffectiveAutoScalingOk() (*MongosEffectiveAutoScaling, bool) {
	if o == nil || IsNil(o.EffectiveAutoScaling) {
		return nil, false
	}

	return o.EffectiveAutoScaling, true
}

// HasEffectiveAutoScaling returns a boolean if a field has been set.
func (o *MongosTopologyCreateRequest) HasEffectiveAutoScaling() bool {
	if o != nil && !IsNil(o.EffectiveAutoScaling) {
		return true
	}

	return false
}

// SetEffectiveAutoScaling gets a reference to the given MongosEffectiveAutoScaling and assigns it to the EffectiveAutoScaling field.
func (o *MongosTopologyCreateRequest) SetEffectiveAutoScaling(v MongosEffectiveAutoScaling) {
	o.EffectiveAutoScaling = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveAutoScaling")
}

// SetEffectiveAutoScalingNil sets EffectiveAutoScaling to an explicit JSON null when marshaled.
func (o *MongosTopologyCreateRequest) SetEffectiveAutoScalingNil() {
	o.EffectiveAutoScaling = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveAutoScaling")
}

// GetEffectiveInstanceSize returns the EffectiveInstanceSize field value if set, zero value otherwise
func (o *MongosTopologyCreateRequest) GetEffectiveInstanceSize() string {
	if o == nil || IsNil(o.EffectiveInstanceSize) {
		var ret string
		return ret
	}
	return *o.EffectiveInstanceSize
}

// GetEffectiveInstanceSizeOk returns a tuple with the EffectiveInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyCreateRequest) GetEffectiveInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveInstanceSize) {
		return nil, false
	}

	return o.EffectiveInstanceSize, true
}

// HasEffectiveInstanceSize returns a boolean if a field has been set.
func (o *MongosTopologyCreateRequest) HasEffectiveInstanceSize() bool {
	if o != nil && !IsNil(o.EffectiveInstanceSize) {
		return true
	}

	return false
}

// SetEffectiveInstanceSize gets a reference to the given string and assigns it to the EffectiveInstanceSize field.
func (o *MongosTopologyCreateRequest) SetEffectiveInstanceSize(v string) {
	o.EffectiveInstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveInstanceSize")
}

// SetEffectiveInstanceSizeNil sets EffectiveInstanceSize to an explicit JSON null when marshaled.
func (o *MongosTopologyCreateRequest) SetEffectiveInstanceSizeNil() {
	o.EffectiveInstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveInstanceSize")
}

// GetEffectiveRegionPlacements returns the EffectiveRegionPlacements field value if set, zero value otherwise
func (o *MongosTopologyCreateRequest) GetEffectiveRegionPlacements() []MongosRegionPlacement {
	if o == nil || IsNil(o.EffectiveRegionPlacements) {
		var ret []MongosRegionPlacement
		return ret
	}
	return *o.EffectiveRegionPlacements
}

// GetEffectiveRegionPlacementsOk returns a tuple with the EffectiveRegionPlacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyCreateRequest) GetEffectiveRegionPlacementsOk() (*[]MongosRegionPlacement, bool) {
	if o == nil || IsNil(o.EffectiveRegionPlacements) {
		return nil, false
	}

	return o.EffectiveRegionPlacements, true
}

// HasEffectiveRegionPlacements returns a boolean if a field has been set.
func (o *MongosTopologyCreateRequest) HasEffectiveRegionPlacements() bool {
	if o != nil && !IsNil(o.EffectiveRegionPlacements) {
		return true
	}

	return false
}

// SetEffectiveRegionPlacements gets a reference to the given []MongosRegionPlacement and assigns it to the EffectiveRegionPlacements field.
func (o *MongosTopologyCreateRequest) SetEffectiveRegionPlacements(v []MongosRegionPlacement) {
	o.EffectiveRegionPlacements = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveRegionPlacements")
}

// SetEffectiveRegionPlacementsNil sets EffectiveRegionPlacements to an explicit JSON null when marshaled.
func (o *MongosTopologyCreateRequest) SetEffectiveRegionPlacementsNil() {
	o.EffectiveRegionPlacements = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveRegionPlacements")
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
func (o *MongosTopologyCreateRequest) GetRegionPlacements() []MongosRegionPlacement {
	if o == nil || IsNil(o.RegionPlacements) {
		var ret []MongosRegionPlacement
		return ret
	}
	return *o.RegionPlacements
}

// GetRegionPlacementsOk returns a tuple with the RegionPlacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyCreateRequest) GetRegionPlacementsOk() (*[]MongosRegionPlacement, bool) {
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

// SetRegionPlacements gets a reference to the given []MongosRegionPlacement and assigns it to the RegionPlacements field.
func (o *MongosTopologyCreateRequest) SetRegionPlacements(v []MongosRegionPlacement) {
	o.RegionPlacements = &v
	o.NullFields = removeNullField(o.NullFields, "RegionPlacements")
}

// SetRegionPlacementsNil sets RegionPlacements to an explicit JSON null when marshaled.
func (o *MongosTopologyCreateRequest) SetRegionPlacementsNil() {
	o.RegionPlacements = nil
	o.NullFields = addNullField(o.NullFields, "RegionPlacements")
}
