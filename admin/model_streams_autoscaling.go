// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsAutoscaling Autoscaling configuration for a stream processor.
type StreamsAutoscaling struct {
	// Flag that indicates whether autoscaling is enabled.  - **Omitted, `null`, or `false`:**   - On `CREATE`: a no-op, there is no persisted setting yet to disable or clear.   - On `MODIFY` or `:startWith`: omitted preserves the current setting. `null` or `false` disables autoscaling and clears its configuration. - **`true`** on `CREATE`, `MODIFY`, or `:startWith`: enables autoscaling.
	Enabled *bool `json:"enabled,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Tier ceiling for autoscaling (scale-up limit).  - **Omitted:**   - On `CREATE`: falls back to the workspace max tier (there is no current bound to preserve).   - On `MODIFY` or `:startWith`: the current bound is preserved. - **`null`** on `CREATE`, `MODIFY`, or `:startWith`: resets the bound to the workspace max tier. - **A tier value** on `CREATE`, `MODIFY`, or `:startWith`: sets the bound to that tier.
	MaxTier *string `json:"maxTier,omitempty"`
	// Tier floor for autoscaling (scale-down limit).  - **Omitted:**   - On `CREATE`: falls back to the workspace default tier (there is no current bound to preserve).   - On `MODIFY` or `:startWith`: the current bound is preserved. - **`null`** on `CREATE`, `MODIFY`, or `:startWith`: resets the bound to the workspace default tier. - **A tier value** on `CREATE`, `MODIFY`, or `:startWith`: sets the bound to that tier.
	MinTier *string `json:"minTier,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsAutoscaling) MarshalJSON() ([]byte, error) {
	type noMethod StreamsAutoscaling
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsAutoscaling instantiates a new StreamsAutoscaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsAutoscaling() *StreamsAutoscaling {
	this := StreamsAutoscaling{}
	return &this
}

// NewStreamsAutoscalingWithDefaults instantiates a new StreamsAutoscaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsAutoscalingWithDefaults() *StreamsAutoscaling {
	this := StreamsAutoscaling{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *StreamsAutoscaling) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsAutoscaling) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StreamsAutoscaling) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StreamsAutoscaling) SetEnabled(v bool) {
	o.Enabled = &v
	o.NullFields = removeNullField(o.NullFields, "Enabled")
}

// SetEnabledNil sets Enabled to an explicit JSON null when marshaled.
func (o *StreamsAutoscaling) SetEnabledNil() {
	o.Enabled = nil
	o.NullFields = addNullField(o.NullFields, "Enabled")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsAutoscaling) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsAutoscaling) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsAutoscaling) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsAutoscaling) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsAutoscaling) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetMaxTier returns the MaxTier field value if set, zero value otherwise
func (o *StreamsAutoscaling) GetMaxTier() string {
	if o == nil || IsNil(o.MaxTier) {
		var ret string
		return ret
	}
	return *o.MaxTier
}

// GetMaxTierOk returns a tuple with the MaxTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsAutoscaling) GetMaxTierOk() (*string, bool) {
	if o == nil || IsNil(o.MaxTier) {
		return nil, false
	}

	return o.MaxTier, true
}

// HasMaxTier returns a boolean if a field has been set.
func (o *StreamsAutoscaling) HasMaxTier() bool {
	if o != nil && !IsNil(o.MaxTier) {
		return true
	}

	return false
}

// SetMaxTier gets a reference to the given string and assigns it to the MaxTier field.
func (o *StreamsAutoscaling) SetMaxTier(v string) {
	o.MaxTier = &v
	o.NullFields = removeNullField(o.NullFields, "MaxTier")
}

// SetMaxTierNil sets MaxTier to an explicit JSON null when marshaled.
func (o *StreamsAutoscaling) SetMaxTierNil() {
	o.MaxTier = nil
	o.NullFields = addNullField(o.NullFields, "MaxTier")
}

// GetMinTier returns the MinTier field value if set, zero value otherwise
func (o *StreamsAutoscaling) GetMinTier() string {
	if o == nil || IsNil(o.MinTier) {
		var ret string
		return ret
	}
	return *o.MinTier
}

// GetMinTierOk returns a tuple with the MinTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsAutoscaling) GetMinTierOk() (*string, bool) {
	if o == nil || IsNil(o.MinTier) {
		return nil, false
	}

	return o.MinTier, true
}

// HasMinTier returns a boolean if a field has been set.
func (o *StreamsAutoscaling) HasMinTier() bool {
	if o != nil && !IsNil(o.MinTier) {
		return true
	}

	return false
}

// SetMinTier gets a reference to the given string and assigns it to the MinTier field.
func (o *StreamsAutoscaling) SetMinTier(v string) {
	o.MinTier = &v
	o.NullFields = removeNullField(o.NullFields, "MinTier")
}

// SetMinTierNil sets MinTier to an explicit JSON null when marshaled.
func (o *StreamsAutoscaling) SetMinTierNil() {
	o.MinTier = nil
	o.NullFields = addNullField(o.NullFields, "MinTier")
}
