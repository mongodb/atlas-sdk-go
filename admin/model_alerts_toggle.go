// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AlertsToggle Enables or disables the specified alert configuration in the specified project.
type AlertsToggle struct {
	// Flag that indicates whether to enable or disable the specified alert configuration in the specified project.
	Enabled *bool `json:"enabled,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AlertsToggle) MarshalJSON() ([]byte, error) {
	type noMethod AlertsToggle
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAlertsToggle instantiates a new AlertsToggle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsToggle() *AlertsToggle {
	this := AlertsToggle{}
	return &this
}

// NewAlertsToggleWithDefaults instantiates a new AlertsToggle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsToggleWithDefaults() *AlertsToggle {
	this := AlertsToggle{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *AlertsToggle) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsToggle) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertsToggle) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertsToggle) SetEnabled(v bool) {
	o.Enabled = &v
	o.NullFields = removeNullField(o.NullFields, "Enabled")
}

// SetEnabledNil sets Enabled to an explicit JSON null when marshaled.
func (o *AlertsToggle) SetEnabledNil() {
	o.Enabled = nil
	o.NullFields = addNullField(o.NullFields, "Enabled")
}
