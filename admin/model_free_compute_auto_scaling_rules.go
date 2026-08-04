// Code based on the AtlasAPI V2 OpenAPI file

package admin

// FreeComputeAutoScalingRules Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down.
type FreeComputeAutoScalingRules struct {
	// Maximum instance size to which your cluster can automatically scale.
	MaxInstanceSize *string `json:"maxInstanceSize,omitempty"`
	// Minimum instance size to which your cluster can automatically scale.
	MinInstanceSize *string `json:"minInstanceSize,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *FreeComputeAutoScalingRules) MarshalJSON() ([]byte, error) {
	type noMethod FreeComputeAutoScalingRules
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewFreeComputeAutoScalingRules instantiates a new FreeComputeAutoScalingRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeComputeAutoScalingRules() *FreeComputeAutoScalingRules {
	this := FreeComputeAutoScalingRules{}
	return &this
}

// NewFreeComputeAutoScalingRulesWithDefaults instantiates a new FreeComputeAutoScalingRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeComputeAutoScalingRulesWithDefaults() *FreeComputeAutoScalingRules {
	this := FreeComputeAutoScalingRules{}
	return &this
}

// GetMaxInstanceSize returns the MaxInstanceSize field value if set, zero value otherwise
func (o *FreeComputeAutoScalingRules) GetMaxInstanceSize() string {
	if o == nil || IsNil(o.MaxInstanceSize) {
		var ret string
		return ret
	}
	return *o.MaxInstanceSize
}

// GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeComputeAutoScalingRules) GetMaxInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxInstanceSize) {
		return nil, false
	}

	return o.MaxInstanceSize, true
}

// HasMaxInstanceSize returns a boolean if a field has been set.
func (o *FreeComputeAutoScalingRules) HasMaxInstanceSize() bool {
	if o != nil && !IsNil(o.MaxInstanceSize) {
		return true
	}

	return false
}

// SetMaxInstanceSize gets a reference to the given string and assigns it to the MaxInstanceSize field.
func (o *FreeComputeAutoScalingRules) SetMaxInstanceSize(v string) {
	o.MaxInstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "MaxInstanceSize")
}

// SetMaxInstanceSizeNil sets MaxInstanceSize to an explicit JSON null when marshaled.
func (o *FreeComputeAutoScalingRules) SetMaxInstanceSizeNil() {
	o.MaxInstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "MaxInstanceSize")
}

// GetMinInstanceSize returns the MinInstanceSize field value if set, zero value otherwise
func (o *FreeComputeAutoScalingRules) GetMinInstanceSize() string {
	if o == nil || IsNil(o.MinInstanceSize) {
		var ret string
		return ret
	}
	return *o.MinInstanceSize
}

// GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeComputeAutoScalingRules) GetMinInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinInstanceSize) {
		return nil, false
	}

	return o.MinInstanceSize, true
}

// HasMinInstanceSize returns a boolean if a field has been set.
func (o *FreeComputeAutoScalingRules) HasMinInstanceSize() bool {
	if o != nil && !IsNil(o.MinInstanceSize) {
		return true
	}

	return false
}

// SetMinInstanceSize gets a reference to the given string and assigns it to the MinInstanceSize field.
func (o *FreeComputeAutoScalingRules) SetMinInstanceSize(v string) {
	o.MinInstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "MinInstanceSize")
}

// SetMinInstanceSizeNil sets MinInstanceSize to an explicit JSON null when marshaled.
func (o *FreeComputeAutoScalingRules) SetMinInstanceSizeNil() {
	o.MinInstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "MinInstanceSize")
}
