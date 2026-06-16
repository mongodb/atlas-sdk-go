// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiSearchDeploymentRequest struct for ApiSearchDeploymentRequest
type ApiSearchDeploymentRequest struct {
	// Default number of Search Nodes per region. Applied to a region without an explicit override.
	DefaultNodeCount *int `json:"defaultNodeCount,omitempty"`
	// List of settings that configure the Search Nodes for your cluster. Provide one element per region when configuring asymmetric deployments; a single element applies to all regions.
	Specs []ApiSearchDeploymentRequestSpec `json:"specs"`
}

// NewApiSearchDeploymentRequest instantiates a new ApiSearchDeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchDeploymentRequest(specs []ApiSearchDeploymentRequestSpec) *ApiSearchDeploymentRequest {
	this := ApiSearchDeploymentRequest{}
	this.Specs = specs
	return &this
}

// NewApiSearchDeploymentRequestWithDefaults instantiates a new ApiSearchDeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchDeploymentRequestWithDefaults() *ApiSearchDeploymentRequest {
	this := ApiSearchDeploymentRequest{}
	return &this
}

// GetDefaultNodeCount returns the DefaultNodeCount field value if set, zero value otherwise
func (o *ApiSearchDeploymentRequest) GetDefaultNodeCount() int {
	if o == nil || IsNil(o.DefaultNodeCount) {
		var ret int
		return ret
	}
	return *o.DefaultNodeCount
}

// GetDefaultNodeCountOk returns a tuple with the DefaultNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentRequest) GetDefaultNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.DefaultNodeCount) {
		return nil, false
	}

	return o.DefaultNodeCount, true
}

// HasDefaultNodeCount returns a boolean if a field has been set.
func (o *ApiSearchDeploymentRequest) HasDefaultNodeCount() bool {
	if o != nil && !IsNil(o.DefaultNodeCount) {
		return true
	}

	return false
}

// SetDefaultNodeCount gets a reference to the given int and assigns it to the DefaultNodeCount field.
func (o *ApiSearchDeploymentRequest) SetDefaultNodeCount(v int) {
	o.DefaultNodeCount = &v
}

// GetSpecs returns the Specs field value
func (o *ApiSearchDeploymentRequest) GetSpecs() []ApiSearchDeploymentRequestSpec {
	if o == nil {
		var ret []ApiSearchDeploymentRequestSpec
		return ret
	}

	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentRequest) GetSpecsOk() (*[]ApiSearchDeploymentRequestSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specs, true
}

// SetSpecs sets field value
func (o *ApiSearchDeploymentRequest) SetSpecs(v []ApiSearchDeploymentRequestSpec) {
	o.Specs = v
}
