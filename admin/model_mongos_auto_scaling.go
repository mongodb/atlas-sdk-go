// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongosAutoScaling Autoscaling configuration for the dedicated mongos tier. Default enabled.
type MongosAutoScaling struct {
	// Flag that indicates whether autoscaling is enabled for the tier.
	Enabled *bool `json:"enabled,omitempty"`
	// Maximum instance size the tier may scale up to (for example, `M80`).
	MaxInstanceSize *string `json:"maxInstanceSize,omitempty"`
	// Minimum instance size the tier may scale down to (for example, `M40`).
	MinInstanceSize *string `json:"minInstanceSize,omitempty"`
	// Utilization, from 0 to 1, of CPU and connections below which the tier scales in. Defaults to `0.40`.
	ScaleInThreshold *float64 `json:"scaleInThreshold,omitempty"`
	// Utilization, from 0 to 1, of CPU or connections above which the tier scales out. Defaults to `0.75`.
	ScaleOutThreshold *float64 `json:"scaleOutThreshold,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o MongosAutoScaling) MarshalJSON() ([]byte, error) {
	type noMethod MongosAutoScaling
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewMongosAutoScaling instantiates a new MongosAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongosAutoScaling() *MongosAutoScaling {
	this := MongosAutoScaling{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewMongosAutoScalingWithDefaults instantiates a new MongosAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongosAutoScalingWithDefaults() *MongosAutoScaling {
	this := MongosAutoScaling{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *MongosAutoScaling) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosAutoScaling) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MongosAutoScaling) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MongosAutoScaling) SetEnabled(v bool) {
	o.Enabled = &v
	o.NullFields = removeNullField(o.NullFields, "Enabled")
}

// SetEnabledNil sets Enabled to an explicit JSON null when marshaled.
func (o *MongosAutoScaling) SetEnabledNil() {
	o.Enabled = nil
	o.NullFields = addNullField(o.NullFields, "Enabled")
}

// GetMaxInstanceSize returns the MaxInstanceSize field value if set, zero value otherwise
func (o *MongosAutoScaling) GetMaxInstanceSize() string {
	if o == nil || IsNil(o.MaxInstanceSize) {
		var ret string
		return ret
	}
	return *o.MaxInstanceSize
}

// GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosAutoScaling) GetMaxInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxInstanceSize) {
		return nil, false
	}

	return o.MaxInstanceSize, true
}

// HasMaxInstanceSize returns a boolean if a field has been set.
func (o *MongosAutoScaling) HasMaxInstanceSize() bool {
	if o != nil && !IsNil(o.MaxInstanceSize) {
		return true
	}

	return false
}

// SetMaxInstanceSize gets a reference to the given string and assigns it to the MaxInstanceSize field.
func (o *MongosAutoScaling) SetMaxInstanceSize(v string) {
	o.MaxInstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "MaxInstanceSize")
}

// SetMaxInstanceSizeNil sets MaxInstanceSize to an explicit JSON null when marshaled.
func (o *MongosAutoScaling) SetMaxInstanceSizeNil() {
	o.MaxInstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "MaxInstanceSize")
}

// GetMinInstanceSize returns the MinInstanceSize field value if set, zero value otherwise
func (o *MongosAutoScaling) GetMinInstanceSize() string {
	if o == nil || IsNil(o.MinInstanceSize) {
		var ret string
		return ret
	}
	return *o.MinInstanceSize
}

// GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosAutoScaling) GetMinInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinInstanceSize) {
		return nil, false
	}

	return o.MinInstanceSize, true
}

// HasMinInstanceSize returns a boolean if a field has been set.
func (o *MongosAutoScaling) HasMinInstanceSize() bool {
	if o != nil && !IsNil(o.MinInstanceSize) {
		return true
	}

	return false
}

// SetMinInstanceSize gets a reference to the given string and assigns it to the MinInstanceSize field.
func (o *MongosAutoScaling) SetMinInstanceSize(v string) {
	o.MinInstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "MinInstanceSize")
}

// SetMinInstanceSizeNil sets MinInstanceSize to an explicit JSON null when marshaled.
func (o *MongosAutoScaling) SetMinInstanceSizeNil() {
	o.MinInstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "MinInstanceSize")
}

// GetScaleInThreshold returns the ScaleInThreshold field value if set, zero value otherwise
func (o *MongosAutoScaling) GetScaleInThreshold() float64 {
	if o == nil || IsNil(o.ScaleInThreshold) {
		var ret float64
		return ret
	}
	return *o.ScaleInThreshold
}

// GetScaleInThresholdOk returns a tuple with the ScaleInThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosAutoScaling) GetScaleInThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.ScaleInThreshold) {
		return nil, false
	}

	return o.ScaleInThreshold, true
}

// HasScaleInThreshold returns a boolean if a field has been set.
func (o *MongosAutoScaling) HasScaleInThreshold() bool {
	if o != nil && !IsNil(o.ScaleInThreshold) {
		return true
	}

	return false
}

// SetScaleInThreshold gets a reference to the given float64 and assigns it to the ScaleInThreshold field.
func (o *MongosAutoScaling) SetScaleInThreshold(v float64) {
	o.ScaleInThreshold = &v
	o.NullFields = removeNullField(o.NullFields, "ScaleInThreshold")
}

// SetScaleInThresholdNil sets ScaleInThreshold to an explicit JSON null when marshaled.
func (o *MongosAutoScaling) SetScaleInThresholdNil() {
	o.ScaleInThreshold = nil
	o.NullFields = addNullField(o.NullFields, "ScaleInThreshold")
}

// GetScaleOutThreshold returns the ScaleOutThreshold field value if set, zero value otherwise
func (o *MongosAutoScaling) GetScaleOutThreshold() float64 {
	if o == nil || IsNil(o.ScaleOutThreshold) {
		var ret float64
		return ret
	}
	return *o.ScaleOutThreshold
}

// GetScaleOutThresholdOk returns a tuple with the ScaleOutThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongosAutoScaling) GetScaleOutThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.ScaleOutThreshold) {
		return nil, false
	}

	return o.ScaleOutThreshold, true
}

// HasScaleOutThreshold returns a boolean if a field has been set.
func (o *MongosAutoScaling) HasScaleOutThreshold() bool {
	if o != nil && !IsNil(o.ScaleOutThreshold) {
		return true
	}

	return false
}

// SetScaleOutThreshold gets a reference to the given float64 and assigns it to the ScaleOutThreshold field.
func (o *MongosAutoScaling) SetScaleOutThreshold(v float64) {
	o.ScaleOutThreshold = &v
	o.NullFields = removeNullField(o.NullFields, "ScaleOutThreshold")
}

// SetScaleOutThresholdNil sets ScaleOutThreshold to an explicit JSON null when marshaled.
func (o *MongosAutoScaling) SetScaleOutThresholdNil() {
	o.ScaleOutThreshold = nil
	o.NullFields = addNullField(o.NullFields, "ScaleOutThreshold")
}
