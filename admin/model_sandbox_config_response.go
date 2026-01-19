// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SandboxConfigResponse Configuration details for the sandbox feature of an organization.
type SandboxConfigResponse struct {
	ClusterTemplate *SandboxClusterTemplate `json:"clusterTemplate,omitempty"`
	// Whether sandbox mode is enabled for the organization.
	Enabled *bool `json:"enabled,omitempty"`
	// The id of the sandbox config.
	// Read only field.
	Id *string `json:"id,omitempty"`
}

// NewSandboxConfigResponse instantiates a new SandboxConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxConfigResponse() *SandboxConfigResponse {
	this := SandboxConfigResponse{}
	return &this
}

// NewSandboxConfigResponseWithDefaults instantiates a new SandboxConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxConfigResponseWithDefaults() *SandboxConfigResponse {
	this := SandboxConfigResponse{}
	return &this
}

// GetClusterTemplate returns the ClusterTemplate field value if set, zero value otherwise
func (o *SandboxConfigResponse) GetClusterTemplate() SandboxClusterTemplate {
	if o == nil || IsNil(o.ClusterTemplate) {
		var ret SandboxClusterTemplate
		return ret
	}
	return *o.ClusterTemplate
}

// GetClusterTemplateOk returns a tuple with the ClusterTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxConfigResponse) GetClusterTemplateOk() (*SandboxClusterTemplate, bool) {
	if o == nil || IsNil(o.ClusterTemplate) {
		return nil, false
	}

	return o.ClusterTemplate, true
}

// HasClusterTemplate returns a boolean if a field has been set.
func (o *SandboxConfigResponse) HasClusterTemplate() bool {
	if o != nil && !IsNil(o.ClusterTemplate) {
		return true
	}

	return false
}

// SetClusterTemplate gets a reference to the given SandboxClusterTemplate and assigns it to the ClusterTemplate field.
func (o *SandboxConfigResponse) SetClusterTemplate(v SandboxClusterTemplate) {
	o.ClusterTemplate = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *SandboxConfigResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxConfigResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SandboxConfigResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SandboxConfigResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *SandboxConfigResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxConfigResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SandboxConfigResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SandboxConfigResponse) SetId(v string) {
	o.Id = &v
}
