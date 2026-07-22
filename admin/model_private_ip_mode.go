// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PrivateIPMode struct for PrivateIPMode
type PrivateIPMode struct {
	// Flag that indicates whether someone enabled **Connect via Peering Only** mode for the specified project.
	Enabled bool `json:"enabled"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PrivateIPMode) MarshalJSON() ([]byte, error) {
	type noMethod PrivateIPMode
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPrivateIPMode instantiates a new PrivateIPMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateIPMode(enabled bool) *PrivateIPMode {
	this := PrivateIPMode{}
	this.Enabled = enabled
	return &this
}

// NewPrivateIPModeWithDefaults instantiates a new PrivateIPMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateIPModeWithDefaults() *PrivateIPMode {
	this := PrivateIPMode{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *PrivateIPMode) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PrivateIPMode) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PrivateIPMode) SetEnabled(v bool) {
	o.Enabled = v
}
