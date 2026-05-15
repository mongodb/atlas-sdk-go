// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiSearchDeploymentRequestSpec struct for ApiSearchDeploymentRequestSpec
type ApiSearchDeploymentRequestSpec struct {
	// Cloud service provider that hosts the Search Nodes in this region. Required when a region is specified.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Hardware specification for the Search Node instance sizes.
	InstanceSize string `json:"instanceSize"`
	// Number of Search Nodes in this region. Optional; falls back to the request-level default when omitted.
	NodeCount *int `json:"nodeCount,omitempty"`
	// Cloud provider region where Search Nodes are provisioned. Required when the request configures more than one region; optional for single-region requests.
	RegionName *string `json:"regionName,omitempty"`
}

// NewApiSearchDeploymentRequestSpec instantiates a new ApiSearchDeploymentRequestSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchDeploymentRequestSpec(instanceSize string) *ApiSearchDeploymentRequestSpec {
	this := ApiSearchDeploymentRequestSpec{}
	this.InstanceSize = instanceSize
	return &this
}

// NewApiSearchDeploymentRequestSpecWithDefaults instantiates a new ApiSearchDeploymentRequestSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchDeploymentRequestSpecWithDefaults() *ApiSearchDeploymentRequestSpec {
	this := ApiSearchDeploymentRequestSpec{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *ApiSearchDeploymentRequestSpec) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentRequestSpec) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ApiSearchDeploymentRequestSpec) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ApiSearchDeploymentRequestSpec) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetInstanceSize returns the InstanceSize field value
func (o *ApiSearchDeploymentRequestSpec) GetInstanceSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentRequestSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceSize, true
}

// SetInstanceSize sets field value
func (o *ApiSearchDeploymentRequestSpec) SetInstanceSize(v string) {
	o.InstanceSize = v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise
func (o *ApiSearchDeploymentRequestSpec) GetNodeCount() int {
	if o == nil || IsNil(o.NodeCount) {
		var ret int
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentRequestSpec) GetNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}

	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ApiSearchDeploymentRequestSpec) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int and assigns it to the NodeCount field.
func (o *ApiSearchDeploymentRequestSpec) SetNodeCount(v int) {
	o.NodeCount = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *ApiSearchDeploymentRequestSpec) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentRequestSpec) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ApiSearchDeploymentRequestSpec) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ApiSearchDeploymentRequestSpec) SetRegionName(v string) {
	o.RegionName = &v
}
