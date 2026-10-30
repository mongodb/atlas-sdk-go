// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AutoEmbeddingInferenceScopeResponse Configured and effective inference scope for Auto-Embedding indexes on one cluster.
type AutoEmbeddingInferenceScopeResponse struct {
	Configured AutoEmbeddingConfiguredInferenceScope `json:"configured"`
	Effective  AutoEmbeddingEffectiveInferenceScope  `json:"effective"`
	// Flag that indicates whether the inference scope can be configured.
	// Read only field.
	Mutable bool `json:"mutable"`
	// Reason code that explains why the inference scope cannot be configured.
	// Read only field.
	Reason *string `json:"reason,omitempty"`
	// Source from which Atlas derived the effective inference scope.
	// Read only field.
	Source string `json:"source"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AutoEmbeddingInferenceScopeResponse) MarshalJSON() ([]byte, error) {
	type noMethod AutoEmbeddingInferenceScopeResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAutoEmbeddingInferenceScopeResponse instantiates a new AutoEmbeddingInferenceScopeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoEmbeddingInferenceScopeResponse(configured AutoEmbeddingConfiguredInferenceScope, effective AutoEmbeddingEffectiveInferenceScope, mutable bool, source string) *AutoEmbeddingInferenceScopeResponse {
	this := AutoEmbeddingInferenceScopeResponse{}
	this.Configured = configured
	this.Effective = effective
	this.Mutable = mutable
	this.Source = source
	return &this
}

// NewAutoEmbeddingInferenceScopeResponseWithDefaults instantiates a new AutoEmbeddingInferenceScopeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoEmbeddingInferenceScopeResponseWithDefaults() *AutoEmbeddingInferenceScopeResponse {
	this := AutoEmbeddingInferenceScopeResponse{}
	return &this
}

// GetConfigured returns the Configured field value
func (o *AutoEmbeddingInferenceScopeResponse) GetConfigured() AutoEmbeddingConfiguredInferenceScope {
	if o == nil {
		var ret AutoEmbeddingConfiguredInferenceScope
		return ret
	}

	return o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingInferenceScopeResponse) GetConfiguredOk() (*AutoEmbeddingConfiguredInferenceScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configured, true
}

// SetConfigured sets field value
func (o *AutoEmbeddingInferenceScopeResponse) SetConfigured(v AutoEmbeddingConfiguredInferenceScope) {
	o.Configured = v
}

// GetEffective returns the Effective field value
func (o *AutoEmbeddingInferenceScopeResponse) GetEffective() AutoEmbeddingEffectiveInferenceScope {
	if o == nil {
		var ret AutoEmbeddingEffectiveInferenceScope
		return ret
	}

	return o.Effective
}

// GetEffectiveOk returns a tuple with the Effective field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingInferenceScopeResponse) GetEffectiveOk() (*AutoEmbeddingEffectiveInferenceScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effective, true
}

// SetEffective sets field value
func (o *AutoEmbeddingInferenceScopeResponse) SetEffective(v AutoEmbeddingEffectiveInferenceScope) {
	o.Effective = v
}

// GetMutable returns the Mutable field value
func (o *AutoEmbeddingInferenceScopeResponse) GetMutable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mutable
}

// GetMutableOk returns a tuple with the Mutable field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingInferenceScopeResponse) GetMutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mutable, true
}

// SetMutable sets field value
func (o *AutoEmbeddingInferenceScopeResponse) SetMutable(v bool) {
	o.Mutable = v
}

// GetReason returns the Reason field value if set, zero value otherwise
func (o *AutoEmbeddingInferenceScopeResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingInferenceScopeResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}

	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AutoEmbeddingInferenceScopeResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AutoEmbeddingInferenceScopeResponse) SetReason(v string) {
	o.Reason = &v
	o.NullFields = removeNullField(o.NullFields, "Reason")
}

// SetReasonNil sets Reason to an explicit JSON null when marshaled.
func (o *AutoEmbeddingInferenceScopeResponse) SetReasonNil() {
	o.Reason = nil
	o.NullFields = addNullField(o.NullFields, "Reason")
}

// GetSource returns the Source field value
func (o *AutoEmbeddingInferenceScopeResponse) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingInferenceScopeResponse) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AutoEmbeddingInferenceScopeResponse) SetSource(v string) {
	o.Source = v
}
