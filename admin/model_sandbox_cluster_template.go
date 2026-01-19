// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SandboxClusterTemplate Cluster template configuration for sandbox clusters.
type SandboxClusterTemplate struct {
	// The provider for a cluster.
	Provider *string `json:"provider,omitempty"`
	// The region for a cluster.
	Region *string `json:"region,omitempty"`
	// The tier for a cluster.
	Tier *string `json:"tier,omitempty"`
}

// NewSandboxClusterTemplate instantiates a new SandboxClusterTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxClusterTemplate() *SandboxClusterTemplate {
	this := SandboxClusterTemplate{}
	return &this
}

// NewSandboxClusterTemplateWithDefaults instantiates a new SandboxClusterTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxClusterTemplateWithDefaults() *SandboxClusterTemplate {
	this := SandboxClusterTemplate{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise
func (o *SandboxClusterTemplate) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxClusterTemplate) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}

	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SandboxClusterTemplate) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *SandboxClusterTemplate) SetProvider(v string) {
	o.Provider = &v
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *SandboxClusterTemplate) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxClusterTemplate) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SandboxClusterTemplate) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *SandboxClusterTemplate) SetRegion(v string) {
	o.Region = &v
}

// GetTier returns the Tier field value if set, zero value otherwise
func (o *SandboxClusterTemplate) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxClusterTemplate) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}

	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *SandboxClusterTemplate) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *SandboxClusterTemplate) SetTier(v string) {
	o.Tier = &v
}
