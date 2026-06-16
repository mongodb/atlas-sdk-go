// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterConfigurationValidation struct for ClusterConfigurationValidation
type ClusterConfigurationValidation struct {
	AdvancedConfiguration *ClusterDescriptionProcessArgs20240805 `json:"advancedConfiguration,omitempty"`
	ClusterDescription    ClusterDescription20240805             `json:"clusterDescription"`
	// Name of the existing cluster to validate against when editing is true.
	ClusterName *string `json:"clusterName,omitempty"`
	// When true, validates the configuration as an update to an existing cluster. When false, validates as a new cluster create.
	Editing              bool                     `json:"editing"`
	SearchDeploymentSpec *ApiSearchDeploymentSpec `json:"searchDeploymentSpec,omitempty"`
	// Optional list of resource tags.
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

// NewClusterConfigurationValidation instantiates a new ClusterConfigurationValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfigurationValidation(clusterDescription ClusterDescription20240805, editing bool) *ClusterConfigurationValidation {
	this := ClusterConfigurationValidation{}
	this.ClusterDescription = clusterDescription
	this.Editing = editing
	return &this
}

// NewClusterConfigurationValidationWithDefaults instantiates a new ClusterConfigurationValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigurationValidationWithDefaults() *ClusterConfigurationValidation {
	this := ClusterConfigurationValidation{}
	return &this
}

// GetAdvancedConfiguration returns the AdvancedConfiguration field value if set, zero value otherwise
func (o *ClusterConfigurationValidation) GetAdvancedConfiguration() ClusterDescriptionProcessArgs20240805 {
	if o == nil || IsNil(o.AdvancedConfiguration) {
		var ret ClusterDescriptionProcessArgs20240805
		return ret
	}
	return *o.AdvancedConfiguration
}

// GetAdvancedConfigurationOk returns a tuple with the AdvancedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidation) GetAdvancedConfigurationOk() (*ClusterDescriptionProcessArgs20240805, bool) {
	if o == nil || IsNil(o.AdvancedConfiguration) {
		return nil, false
	}

	return o.AdvancedConfiguration, true
}

// HasAdvancedConfiguration returns a boolean if a field has been set.
func (o *ClusterConfigurationValidation) HasAdvancedConfiguration() bool {
	if o != nil && !IsNil(o.AdvancedConfiguration) {
		return true
	}

	return false
}

// SetAdvancedConfiguration gets a reference to the given ClusterDescriptionProcessArgs20240805 and assigns it to the AdvancedConfiguration field.
func (o *ClusterConfigurationValidation) SetAdvancedConfiguration(v ClusterDescriptionProcessArgs20240805) {
	o.AdvancedConfiguration = &v
}

// GetClusterDescription returns the ClusterDescription field value
func (o *ClusterConfigurationValidation) GetClusterDescription() ClusterDescription20240805 {
	if o == nil {
		var ret ClusterDescription20240805
		return ret
	}

	return o.ClusterDescription
}

// GetClusterDescriptionOk returns a tuple with the ClusterDescription field value
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidation) GetClusterDescriptionOk() (*ClusterDescription20240805, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterDescription, true
}

// SetClusterDescription sets field value
func (o *ClusterConfigurationValidation) SetClusterDescription(v ClusterDescription20240805) {
	o.ClusterDescription = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *ClusterConfigurationValidation) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidation) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ClusterConfigurationValidation) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ClusterConfigurationValidation) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetEditing returns the Editing field value
func (o *ClusterConfigurationValidation) GetEditing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Editing
}

// GetEditingOk returns a tuple with the Editing field value
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidation) GetEditingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editing, true
}

// SetEditing sets field value
func (o *ClusterConfigurationValidation) SetEditing(v bool) {
	o.Editing = v
}

// GetSearchDeploymentSpec returns the SearchDeploymentSpec field value if set, zero value otherwise
func (o *ClusterConfigurationValidation) GetSearchDeploymentSpec() ApiSearchDeploymentSpec {
	if o == nil || IsNil(o.SearchDeploymentSpec) {
		var ret ApiSearchDeploymentSpec
		return ret
	}
	return *o.SearchDeploymentSpec
}

// GetSearchDeploymentSpecOk returns a tuple with the SearchDeploymentSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidation) GetSearchDeploymentSpecOk() (*ApiSearchDeploymentSpec, bool) {
	if o == nil || IsNil(o.SearchDeploymentSpec) {
		return nil, false
	}

	return o.SearchDeploymentSpec, true
}

// HasSearchDeploymentSpec returns a boolean if a field has been set.
func (o *ClusterConfigurationValidation) HasSearchDeploymentSpec() bool {
	if o != nil && !IsNil(o.SearchDeploymentSpec) {
		return true
	}

	return false
}

// SetSearchDeploymentSpec gets a reference to the given ApiSearchDeploymentSpec and assigns it to the SearchDeploymentSpec field.
func (o *ClusterConfigurationValidation) SetSearchDeploymentSpec(v ApiSearchDeploymentSpec) {
	o.SearchDeploymentSpec = &v
}

// GetTags returns the Tags field value if set, zero value otherwise
func (o *ClusterConfigurationValidation) GetTags() []ResourceTag {
	if o == nil || IsNil(o.Tags) {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidation) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}

	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ClusterConfigurationValidation) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *ClusterConfigurationValidation) SetTags(v []ResourceTag) {
	o.Tags = &v
}
