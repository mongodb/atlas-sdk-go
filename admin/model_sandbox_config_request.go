// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SandboxConfigRequest Settings to configure the sandbox feature for an organization.
type SandboxConfigRequest struct {
	ClusterTemplate *SandboxClusterTemplate `json:"clusterTemplate,omitempty"`
	// Whether sandbox mode is enabled for the organization.
	Enabled *bool `json:"enabled,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *SandboxConfigRequest) MarshalJSON() ([]byte, error) {
	type noMethod SandboxConfigRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewSandboxConfigRequest instantiates a new SandboxConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxConfigRequest() *SandboxConfigRequest {
	this := SandboxConfigRequest{}
	return &this
}

// NewSandboxConfigRequestWithDefaults instantiates a new SandboxConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxConfigRequestWithDefaults() *SandboxConfigRequest {
	this := SandboxConfigRequest{}
	return &this
}

// GetClusterTemplate returns the ClusterTemplate field value if set, zero value otherwise
func (o *SandboxConfigRequest) GetClusterTemplate() SandboxClusterTemplate {
	if o == nil || IsNil(o.ClusterTemplate) {
		var ret SandboxClusterTemplate
		return ret
	}
	return *o.ClusterTemplate
}

// GetClusterTemplateOk returns a tuple with the ClusterTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxConfigRequest) GetClusterTemplateOk() (*SandboxClusterTemplate, bool) {
	if o == nil || IsNil(o.ClusterTemplate) {
		return nil, false
	}

	return o.ClusterTemplate, true
}

// HasClusterTemplate returns a boolean if a field has been set.
func (o *SandboxConfigRequest) HasClusterTemplate() bool {
	if o != nil && !IsNil(o.ClusterTemplate) {
		return true
	}

	return false
}

// SetClusterTemplate gets a reference to the given SandboxClusterTemplate and assigns it to the ClusterTemplate field.
func (o *SandboxConfigRequest) SetClusterTemplate(v SandboxClusterTemplate) {
	o.ClusterTemplate = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterTemplate")
}

// SetClusterTemplateNil sets ClusterTemplate to an explicit JSON null when marshaled.
func (o *SandboxConfigRequest) SetClusterTemplateNil() {
	o.ClusterTemplate = nil
	o.NullFields = addNullField(o.NullFields, "ClusterTemplate")
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *SandboxConfigRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxConfigRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SandboxConfigRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SandboxConfigRequest) SetEnabled(v bool) {
	o.Enabled = &v
	o.NullFields = removeNullField(o.NullFields, "Enabled")
}

// SetEnabledNil sets Enabled to an explicit JSON null when marshaled.
func (o *SandboxConfigRequest) SetEnabledNil() {
	o.Enabled = nil
	o.NullFields = addNullField(o.NullFields, "Enabled")
}
