// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiSearchDeploymentEffectiveSpec struct for ApiSearchDeploymentEffectiveSpec
type ApiSearchDeploymentEffectiveSpec struct {
	// Cloud service provider on which Search Nodes are provisioned.
	// Read only field.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Hardware specification for the Search Node instance sizes.
	// Read only field.
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Number of Search Nodes in this region.
	// Read only field.
	NodeCount *int `json:"nodeCount,omitempty"`
	// Cloud provider region where Search Nodes are provisioned.
	// Read only field.
	RegionName *string `json:"regionName,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiSearchDeploymentEffectiveSpec) MarshalJSON() ([]byte, error) {
	type noMethod ApiSearchDeploymentEffectiveSpec
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiSearchDeploymentEffectiveSpec instantiates a new ApiSearchDeploymentEffectiveSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchDeploymentEffectiveSpec() *ApiSearchDeploymentEffectiveSpec {
	this := ApiSearchDeploymentEffectiveSpec{}
	return &this
}

// NewApiSearchDeploymentEffectiveSpecWithDefaults instantiates a new ApiSearchDeploymentEffectiveSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchDeploymentEffectiveSpecWithDefaults() *ApiSearchDeploymentEffectiveSpec {
	this := ApiSearchDeploymentEffectiveSpec{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *ApiSearchDeploymentEffectiveSpec) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentEffectiveSpec) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ApiSearchDeploymentEffectiveSpec) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ApiSearchDeploymentEffectiveSpec) SetCloudProvider(v string) {
	o.CloudProvider = &v
	o.NullFields = removeNullField(o.NullFields, "CloudProvider")
}

// SetCloudProviderNil sets CloudProvider to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentEffectiveSpec) SetCloudProviderNil() {
	o.CloudProvider = nil
	o.NullFields = addNullField(o.NullFields, "CloudProvider")
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise
func (o *ApiSearchDeploymentEffectiveSpec) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentEffectiveSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}

	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *ApiSearchDeploymentEffectiveSpec) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *ApiSearchDeploymentEffectiveSpec) SetInstanceSize(v string) {
	o.InstanceSize = &v
	o.NullFields = removeNullField(o.NullFields, "InstanceSize")
}

// SetInstanceSizeNil sets InstanceSize to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentEffectiveSpec) SetInstanceSizeNil() {
	o.InstanceSize = nil
	o.NullFields = addNullField(o.NullFields, "InstanceSize")
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise
func (o *ApiSearchDeploymentEffectiveSpec) GetNodeCount() int {
	if o == nil || IsNil(o.NodeCount) {
		var ret int
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentEffectiveSpec) GetNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}

	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ApiSearchDeploymentEffectiveSpec) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int and assigns it to the NodeCount field.
func (o *ApiSearchDeploymentEffectiveSpec) SetNodeCount(v int) {
	o.NodeCount = &v
	o.NullFields = removeNullField(o.NullFields, "NodeCount")
}

// SetNodeCountNil sets NodeCount to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentEffectiveSpec) SetNodeCountNil() {
	o.NodeCount = nil
	o.NullFields = addNullField(o.NullFields, "NodeCount")
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *ApiSearchDeploymentEffectiveSpec) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentEffectiveSpec) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ApiSearchDeploymentEffectiveSpec) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ApiSearchDeploymentEffectiveSpec) SetRegionName(v string) {
	o.RegionName = &v
	o.NullFields = removeNullField(o.NullFields, "RegionName")
}

// SetRegionNameNil sets RegionName to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentEffectiveSpec) SetRegionNameNil() {
	o.RegionName = nil
	o.NullFields = addNullField(o.NullFields, "RegionName")
}
