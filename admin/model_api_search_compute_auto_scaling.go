// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiSearchComputeAutoScaling Settings that control Search Node tier scaling based on load.
type ApiSearchComputeAutoScaling struct {
	// Flag that indicates whether Atlas raises the Search Node tier when the nodes are under sustained load. If set to `true`, you must also set `maxInstanceTier`.
	Enabled *bool `json:"enabled,omitempty"`
	// Highest Search Node tier that Atlas can scale up to. Required when `enabled` is `true`.
	MaxInstanceTier *string `json:"maxInstanceTier,omitempty"`
	// Lowest Search Node tier that Atlas can scale down to. Required when `scaleDownEnabled` is `true`. Scaling down is not supported yet, so setting this returns an error.
	MinInstanceTier *string `json:"minInstanceTier,omitempty"`
	// Flag that indicates whether Atlas lowers the Search Node tier when load drops. Takes effect only when `enabled` is `true`. If set to `true`, you must also set `minInstanceTier`. Scaling down is not supported yet, so setting this to `true` returns an error.
	ScaleDownEnabled *bool `json:"scaleDownEnabled,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiSearchComputeAutoScaling) MarshalJSON() ([]byte, error) {
	type noMethod ApiSearchComputeAutoScaling
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiSearchComputeAutoScaling instantiates a new ApiSearchComputeAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchComputeAutoScaling() *ApiSearchComputeAutoScaling {
	this := ApiSearchComputeAutoScaling{}
	var enabled bool = false
	this.Enabled = &enabled
	var scaleDownEnabled bool = false
	this.ScaleDownEnabled = &scaleDownEnabled
	return &this
}

// NewApiSearchComputeAutoScalingWithDefaults instantiates a new ApiSearchComputeAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchComputeAutoScalingWithDefaults() *ApiSearchComputeAutoScaling {
	this := ApiSearchComputeAutoScaling{}
	var enabled bool = false
	this.Enabled = &enabled
	var scaleDownEnabled bool = false
	this.ScaleDownEnabled = &scaleDownEnabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *ApiSearchComputeAutoScaling) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchComputeAutoScaling) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiSearchComputeAutoScaling) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiSearchComputeAutoScaling) SetEnabled(v bool) {
	o.Enabled = &v
	o.NullFields = removeNullField(o.NullFields, "Enabled")
}

// SetEnabledNil sets Enabled to an explicit JSON null when marshaled.
func (o *ApiSearchComputeAutoScaling) SetEnabledNil() {
	o.Enabled = nil
	o.NullFields = addNullField(o.NullFields, "Enabled")
}

// GetMaxInstanceTier returns the MaxInstanceTier field value if set, zero value otherwise
func (o *ApiSearchComputeAutoScaling) GetMaxInstanceTier() string {
	if o == nil || IsNil(o.MaxInstanceTier) {
		var ret string
		return ret
	}
	return *o.MaxInstanceTier
}

// GetMaxInstanceTierOk returns a tuple with the MaxInstanceTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchComputeAutoScaling) GetMaxInstanceTierOk() (*string, bool) {
	if o == nil || IsNil(o.MaxInstanceTier) {
		return nil, false
	}

	return o.MaxInstanceTier, true
}

// HasMaxInstanceTier returns a boolean if a field has been set.
func (o *ApiSearchComputeAutoScaling) HasMaxInstanceTier() bool {
	if o != nil && !IsNil(o.MaxInstanceTier) {
		return true
	}

	return false
}

// SetMaxInstanceTier gets a reference to the given string and assigns it to the MaxInstanceTier field.
func (o *ApiSearchComputeAutoScaling) SetMaxInstanceTier(v string) {
	o.MaxInstanceTier = &v
	o.NullFields = removeNullField(o.NullFields, "MaxInstanceTier")
}

// SetMaxInstanceTierNil sets MaxInstanceTier to an explicit JSON null when marshaled.
func (o *ApiSearchComputeAutoScaling) SetMaxInstanceTierNil() {
	o.MaxInstanceTier = nil
	o.NullFields = addNullField(o.NullFields, "MaxInstanceTier")
}

// GetMinInstanceTier returns the MinInstanceTier field value if set, zero value otherwise
func (o *ApiSearchComputeAutoScaling) GetMinInstanceTier() string {
	if o == nil || IsNil(o.MinInstanceTier) {
		var ret string
		return ret
	}
	return *o.MinInstanceTier
}

// GetMinInstanceTierOk returns a tuple with the MinInstanceTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchComputeAutoScaling) GetMinInstanceTierOk() (*string, bool) {
	if o == nil || IsNil(o.MinInstanceTier) {
		return nil, false
	}

	return o.MinInstanceTier, true
}

// HasMinInstanceTier returns a boolean if a field has been set.
func (o *ApiSearchComputeAutoScaling) HasMinInstanceTier() bool {
	if o != nil && !IsNil(o.MinInstanceTier) {
		return true
	}

	return false
}

// SetMinInstanceTier gets a reference to the given string and assigns it to the MinInstanceTier field.
func (o *ApiSearchComputeAutoScaling) SetMinInstanceTier(v string) {
	o.MinInstanceTier = &v
	o.NullFields = removeNullField(o.NullFields, "MinInstanceTier")
}

// SetMinInstanceTierNil sets MinInstanceTier to an explicit JSON null when marshaled.
func (o *ApiSearchComputeAutoScaling) SetMinInstanceTierNil() {
	o.MinInstanceTier = nil
	o.NullFields = addNullField(o.NullFields, "MinInstanceTier")
}

// GetScaleDownEnabled returns the ScaleDownEnabled field value if set, zero value otherwise
func (o *ApiSearchComputeAutoScaling) GetScaleDownEnabled() bool {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		var ret bool
		return ret
	}
	return *o.ScaleDownEnabled
}

// GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchComputeAutoScaling) GetScaleDownEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		return nil, false
	}

	return o.ScaleDownEnabled, true
}

// HasScaleDownEnabled returns a boolean if a field has been set.
func (o *ApiSearchComputeAutoScaling) HasScaleDownEnabled() bool {
	if o != nil && !IsNil(o.ScaleDownEnabled) {
		return true
	}

	return false
}

// SetScaleDownEnabled gets a reference to the given bool and assigns it to the ScaleDownEnabled field.
func (o *ApiSearchComputeAutoScaling) SetScaleDownEnabled(v bool) {
	o.ScaleDownEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "ScaleDownEnabled")
}

// SetScaleDownEnabledNil sets ScaleDownEnabled to an explicit JSON null when marshaled.
func (o *ApiSearchComputeAutoScaling) SetScaleDownEnabledNil() {
	o.ScaleDownEnabled = nil
	o.NullFields = addNullField(o.NullFields, "ScaleDownEnabled")
}
