// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosTopologyResponse Configuration of the dedicated mongos routers for a sharded cluster, identified by the name of its associated cluster.
type MongosTopologyResponse struct {
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
	// Lifecycle status of the dedicated mongos topology resource.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o MongosTopologyResponse) MarshalJSON() ([]byte, error) {
	type noMethod MongosTopologyResponse
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewMongosTopologyResponse instantiates a new MongosTopologyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosTopologyResponse() *MongosTopologyResponse {
	this := MongosTopologyResponse{}
	return &this
}

// NewMongosTopologyResponseWithDefaults instantiates a new MongosTopologyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosTopologyResponseWithDefaults() *MongosTopologyResponse {
	this := MongosTopologyResponse{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise
func (o *MongosTopologyResponse) GetAutoScaling() MongosAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret MongosAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResponse) GetAutoScalingOk() (*MongosAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}

	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *MongosTopologyResponse) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given MongosAutoScaling and assigns it to the AutoScaling field.
func (o *MongosTopologyResponse) SetAutoScaling(v MongosAutoScaling) {
	o.AutoScaling = &v
	o.NullFields = removeNullField(o.NullFields, "AutoScaling")
}

// SetAutoScalingNil sets AutoScaling to an explicit JSON null when marshaled.
func (o *MongosTopologyResponse) SetAutoScalingNil() {
	o.AutoScaling = nil
	o.NullFields = addNullField(o.NullFields, "AutoScaling")
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *MongosTopologyResponse) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResponse) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *MongosTopologyResponse) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *MongosTopologyResponse) SetClusterName(v string) {
	o.ClusterName = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterName")
}

// SetClusterNameNil sets ClusterName to an explicit JSON null when marshaled.
func (o *MongosTopologyResponse) SetClusterNameNil() {
	o.ClusterName = nil
	o.NullFields = addNullField(o.NullFields, "ClusterName")
}

// GetEffectiveAutoScaling returns the EffectiveAutoScaling field value if set, zero value otherwise
func (o *MongosTopologyResponse) GetEffectiveAutoScaling() MongosEffectiveAutoScaling {
	if o == nil || IsNil(o.EffectiveAutoScaling) {
		var ret MongosEffectiveAutoScaling
		return ret
	}
	return *o.EffectiveAutoScaling
}

// GetEffectiveAutoScalingOk returns a tuple with the EffectiveAutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResponse) GetEffectiveAutoScalingOk() (*MongosEffectiveAutoScaling, bool) {
	if o == nil || IsNil(o.EffectiveAutoScaling) {
		return nil, false
	}

	return o.EffectiveAutoScaling, true
}

// HasEffectiveAutoScaling returns a boolean if a field has been set.
func (o *MongosTopologyResponse) HasEffectiveAutoScaling() bool {
	if o != nil && !IsNil(o.EffectiveAutoScaling) {
		return true
	}

	return false
}

// SetEffectiveAutoScaling gets a reference to the given MongosEffectiveAutoScaling and assigns it to the EffectiveAutoScaling field.
func (o *MongosTopologyResponse) SetEffectiveAutoScaling(v MongosEffectiveAutoScaling) {
	o.EffectiveAutoScaling = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveAutoScaling")
}

// SetEffectiveAutoScalingNil sets EffectiveAutoScaling to an explicit JSON null when marshaled.
func (o *MongosTopologyResponse) SetEffectiveAutoScalingNil() {
	o.EffectiveAutoScaling = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveAutoScaling")
}

// GetEffectiveInstanceSize returns the EffectiveInstanceSize field value if set, zero value otherwise
func (o *MongosTopologyResponse) GetEffectiveInstanceSize() string {
	if o == nil || IsNil(o.EffectiveInstanceSize) {
		var ret string
		return ret
	}
	return *o.EffectiveInstanceSize
}

// GetEffectiveInstanceSizeOk returns a tuple with the EffectiveInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResponse) GetEffectiveInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveInstanceSize) {
		return nil, false
	}

	return o.EffectiveInstanceSize, true
}

// HasEffectiveInstanceSize returns a boolean if a field has been set.
func (o *MongosTopologyResponse) HasEffectiveInstanceSize() bool {
	if o != nil && !IsNil(o.EffectiveInstanceSize) {
		return true
	}

	return false
}

// SetEffectiveInstanceSize gets a reference to the given string and assigns it to the EffectiveInstanceSize field.
func (o *MongosTopologyResponse) SetEffectiveInstanceSize(v string) {
	o.EffectiveInstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveInstanceSize")
}

// SetEffectiveInstanceSizeNil sets EffectiveInstanceSize to an explicit JSON null when marshaled.
func (o *MongosTopologyResponse) SetEffectiveInstanceSizeNil() {
	o.EffectiveInstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveInstanceSize")
}

// GetEffectiveRegionPlacements returns the EffectiveRegionPlacements field value if set, zero value otherwise
func (o *MongosTopologyResponse) GetEffectiveRegionPlacements() []MongosRegionPlacement {
	if o == nil || IsNil(o.EffectiveRegionPlacements) {
		var ret []MongosRegionPlacement
		return ret
	}
	return *o.EffectiveRegionPlacements
}

// GetEffectiveRegionPlacementsOk returns a tuple with the EffectiveRegionPlacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResponse) GetEffectiveRegionPlacementsOk() (*[]MongosRegionPlacement, bool) {
	if o == nil || IsNil(o.EffectiveRegionPlacements) {
		return nil, false
	}

	return o.EffectiveRegionPlacements, true
}

// HasEffectiveRegionPlacements returns a boolean if a field has been set.
func (o *MongosTopologyResponse) HasEffectiveRegionPlacements() bool {
	if o != nil && !IsNil(o.EffectiveRegionPlacements) {
		return true
	}

	return false
}

// SetEffectiveRegionPlacements gets a reference to the given []MongosRegionPlacement and assigns it to the EffectiveRegionPlacements field.
func (o *MongosTopologyResponse) SetEffectiveRegionPlacements(v []MongosRegionPlacement) {
	o.EffectiveRegionPlacements = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveRegionPlacements")
}

// SetEffectiveRegionPlacementsNil sets EffectiveRegionPlacements to an explicit JSON null when marshaled.
func (o *MongosTopologyResponse) SetEffectiveRegionPlacementsNil() {
	o.EffectiveRegionPlacements = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveRegionPlacements")
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise
func (o *MongosTopologyResponse) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResponse) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}

	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *MongosTopologyResponse) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *MongosTopologyResponse) SetInstanceSize(v string) {
	o.InstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "InstanceSize")
}

// SetInstanceSizeNil sets InstanceSize to an explicit JSON null when marshaled.
func (o *MongosTopologyResponse) SetInstanceSizeNil() {
	o.InstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "InstanceSize")
}

// GetRegionPlacements returns the RegionPlacements field value if set, zero value otherwise
func (o *MongosTopologyResponse) GetRegionPlacements() []MongosRegionPlacement {
	if o == nil || IsNil(o.RegionPlacements) {
		var ret []MongosRegionPlacement
		return ret
	}
	return *o.RegionPlacements
}

// GetRegionPlacementsOk returns a tuple with the RegionPlacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResponse) GetRegionPlacementsOk() (*[]MongosRegionPlacement, bool) {
	if o == nil || IsNil(o.RegionPlacements) {
		return nil, false
	}

	return o.RegionPlacements, true
}

// HasRegionPlacements returns a boolean if a field has been set.
func (o *MongosTopologyResponse) HasRegionPlacements() bool {
	if o != nil && !IsNil(o.RegionPlacements) {
		return true
	}

	return false
}

// SetRegionPlacements gets a reference to the given []MongosRegionPlacement and assigns it to the RegionPlacements field.
func (o *MongosTopologyResponse) SetRegionPlacements(v []MongosRegionPlacement) {
	o.RegionPlacements = &v
	o.NullFields = removeNullField(o.NullFields, "RegionPlacements")
}

// SetRegionPlacementsNil sets RegionPlacements to an explicit JSON null when marshaled.
func (o *MongosTopologyResponse) SetRegionPlacementsNil() {
	o.RegionPlacements = nil
	o.NullFields = addNullField(o.NullFields, "RegionPlacements")
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *MongosTopologyResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosTopologyResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MongosTopologyResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MongosTopologyResponse) SetStatus(v string) {
	o.Status = &v
	o.NullFields = removeNullField(o.NullFields, "Status")
}

// SetStatusNil sets Status to an explicit JSON null when marshaled.
func (o *MongosTopologyResponse) SetStatusNil() {
	o.Status = nil
	o.NullFields = addNullField(o.NullFields, "Status")
}
