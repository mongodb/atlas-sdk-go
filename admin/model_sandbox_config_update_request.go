// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SandboxConfigUpdateRequest Settings to update an existing sandbox for an organization.
type SandboxConfigUpdateRequest struct {
	ClusterTemplate *SandboxClusterTemplate `json:"clusterTemplate,omitempty"`
	// Whether sandbox mode is enabled for the organization.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSandboxConfigUpdateRequest instantiates a new SandboxConfigUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxConfigUpdateRequest() *SandboxConfigUpdateRequest {
	this := SandboxConfigUpdateRequest{}
	return &this
}

// NewSandboxConfigUpdateRequestWithDefaults instantiates a new SandboxConfigUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxConfigUpdateRequestWithDefaults() *SandboxConfigUpdateRequest {
	this := SandboxConfigUpdateRequest{}
	return &this
}

// GetClusterTemplate returns the ClusterTemplate field value if set, zero value otherwise
func (o *SandboxConfigUpdateRequest) GetClusterTemplate() SandboxClusterTemplate {
	if o == nil || IsNil(o.ClusterTemplate) {
		var ret SandboxClusterTemplate
		return ret
	}
	return *o.ClusterTemplate
}

// GetClusterTemplateOk returns a tuple with the ClusterTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxConfigUpdateRequest) GetClusterTemplateOk() (*SandboxClusterTemplate, bool) {
	if o == nil || IsNil(o.ClusterTemplate) {
		return nil, false
	}

	return o.ClusterTemplate, true
}

// HasClusterTemplate returns a boolean if a field has been set.
func (o *SandboxConfigUpdateRequest) HasClusterTemplate() bool {
	if o != nil && !IsNil(o.ClusterTemplate) {
		return true
	}

	return false
}

// SetClusterTemplate gets a reference to the given SandboxClusterTemplate and assigns it to the ClusterTemplate field.
func (o *SandboxConfigUpdateRequest) SetClusterTemplate(v SandboxClusterTemplate) {
	o.ClusterTemplate = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *SandboxConfigUpdateRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxConfigUpdateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SandboxConfigUpdateRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SandboxConfigUpdateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}
