// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosRegionAutoScaling Node-count bounds for dedicated mongos autoscaling in this region. Default enabled. To disable autoscaling for a certain region, you can set the `minNodeCount` and `maxNodeCount` to equal `nodeCount`.
type MongosRegionAutoScaling struct {
	// Maximum number of dedicated mongos the tier may scale out to in this region. Defaults to `10`.
	MaxNodeCount *int `json:"maxNodeCount,omitempty"`
	// Minimum number of dedicated mongos the tier may scale in to in this region. Defaults to `2`.
	MinNodeCount *int `json:"minNodeCount,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o MongosRegionAutoScaling) MarshalJSON() ([]byte, error) {
	type noMethod MongosRegionAutoScaling
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewMongosRegionAutoScaling instantiates a new MongosRegionAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosRegionAutoScaling() *MongosRegionAutoScaling {
	this := MongosRegionAutoScaling{}
	return &this
}

// NewMongosRegionAutoScalingWithDefaults instantiates a new MongosRegionAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosRegionAutoScalingWithDefaults() *MongosRegionAutoScaling {
	this := MongosRegionAutoScaling{}
	return &this
}

// GetMaxNodeCount returns the MaxNodeCount field value if set, zero value otherwise
func (o *MongosRegionAutoScaling) GetMaxNodeCount() int {
	if o == nil || IsNil(o.MaxNodeCount) {
		var ret int
		return ret
	}
	return *o.MaxNodeCount
}

// GetMaxNodeCountOk returns a tuple with the MaxNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosRegionAutoScaling) GetMaxNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.MaxNodeCount) {
		return nil, false
	}

	return o.MaxNodeCount, true
}

// HasMaxNodeCount returns a boolean if a field has been set.
func (o *MongosRegionAutoScaling) HasMaxNodeCount() bool {
	if o != nil && !IsNil(o.MaxNodeCount) {
		return true
	}

	return false
}

// SetMaxNodeCount gets a reference to the given int and assigns it to the MaxNodeCount field.
func (o *MongosRegionAutoScaling) SetMaxNodeCount(v int) {
	o.MaxNodeCount = &v
	o.NullFields = removeNullField(o.NullFields, "MaxNodeCount")
}

// SetMaxNodeCountNil sets MaxNodeCount to an explicit JSON null when marshaled.
func (o *MongosRegionAutoScaling) SetMaxNodeCountNil() {
	o.MaxNodeCount = nil
	o.NullFields = addNullField(o.NullFields, "MaxNodeCount")
}

// GetMinNodeCount returns the MinNodeCount field value if set, zero value otherwise
func (o *MongosRegionAutoScaling) GetMinNodeCount() int {
	if o == nil || IsNil(o.MinNodeCount) {
		var ret int
		return ret
	}
	return *o.MinNodeCount
}

// GetMinNodeCountOk returns a tuple with the MinNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosRegionAutoScaling) GetMinNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.MinNodeCount) {
		return nil, false
	}

	return o.MinNodeCount, true
}

// HasMinNodeCount returns a boolean if a field has been set.
func (o *MongosRegionAutoScaling) HasMinNodeCount() bool {
	if o != nil && !IsNil(o.MinNodeCount) {
		return true
	}

	return false
}

// SetMinNodeCount gets a reference to the given int and assigns it to the MinNodeCount field.
func (o *MongosRegionAutoScaling) SetMinNodeCount(v int) {
	o.MinNodeCount = &v
	o.NullFields = removeNullField(o.NullFields, "MinNodeCount")
}

// SetMinNodeCountNil sets MinNodeCount to an explicit JSON null when marshaled.
func (o *MongosRegionAutoScaling) SetMinNodeCountNil() {
	o.MinNodeCount = nil
	o.NullFields = addNullField(o.NullFields, "MinNodeCount")
}
