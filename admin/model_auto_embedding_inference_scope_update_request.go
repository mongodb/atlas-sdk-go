// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AutoEmbeddingInferenceScopeUpdateRequest struct for AutoEmbeddingInferenceScopeUpdateRequest
type AutoEmbeddingInferenceScopeUpdateRequest struct {
	Configured *AutoEmbeddingInferenceScopeUpdateRequestConfigured `json:"configured,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AutoEmbeddingInferenceScopeUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod AutoEmbeddingInferenceScopeUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAutoEmbeddingInferenceScopeUpdateRequest instantiates a new AutoEmbeddingInferenceScopeUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoEmbeddingInferenceScopeUpdateRequest() *AutoEmbeddingInferenceScopeUpdateRequest {
	this := AutoEmbeddingInferenceScopeUpdateRequest{}
	return &this
}

// NewAutoEmbeddingInferenceScopeUpdateRequestWithDefaults instantiates a new AutoEmbeddingInferenceScopeUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoEmbeddingInferenceScopeUpdateRequestWithDefaults() *AutoEmbeddingInferenceScopeUpdateRequest {
	this := AutoEmbeddingInferenceScopeUpdateRequest{}
	return &this
}

// GetConfigured returns the Configured field value if set, zero value otherwise
func (o *AutoEmbeddingInferenceScopeUpdateRequest) GetConfigured() AutoEmbeddingInferenceScopeUpdateRequestConfigured {
	if o == nil || IsNil(o.Configured) {
		var ret AutoEmbeddingInferenceScopeUpdateRequestConfigured
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingInferenceScopeUpdateRequest) GetConfiguredOk() (*AutoEmbeddingInferenceScopeUpdateRequestConfigured, bool) {
	if o == nil || IsNil(o.Configured) {
		return nil, false
	}

	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *AutoEmbeddingInferenceScopeUpdateRequest) HasConfigured() bool {
	if o != nil && !IsNil(o.Configured) {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given AutoEmbeddingInferenceScopeUpdateRequestConfigured and assigns it to the Configured field.
func (o *AutoEmbeddingInferenceScopeUpdateRequest) SetConfigured(v AutoEmbeddingInferenceScopeUpdateRequestConfigured) {
	o.Configured = &v
	o.NullFields = removeNullField(o.NullFields, "Configured")
}

// SetConfiguredNil sets Configured to an explicit JSON null when marshaled.
func (o *AutoEmbeddingInferenceScopeUpdateRequest) SetConfiguredNil() {
	o.Configured = nil
	o.NullFields = addNullField(o.NullFields, "Configured")
}
