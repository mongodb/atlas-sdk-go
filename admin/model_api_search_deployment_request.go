// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ApiSearchDeploymentRequest struct for ApiSearchDeploymentRequest
type ApiSearchDeploymentRequest struct {
	// List of settings that configure the search nodes for your cluster.
	Specs []ApiSearchDeploymentSpec `json:"specs"`
}

// NewApiSearchDeploymentRequest instantiates a new ApiSearchDeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchDeploymentRequest() *ApiSearchDeploymentRequest {
	this := ApiSearchDeploymentRequest{}
	return &this
}

// NewApiSearchDeploymentRequestWithDefaults instantiates a new ApiSearchDeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchDeploymentRequestWithDefaults() *ApiSearchDeploymentRequest {
	this := ApiSearchDeploymentRequest{}
	return &this
}

// GetSpecs returns the Specs field value if set, zero value otherwise
func (o *ApiSearchDeploymentRequest) GetSpecs() []ApiSearchDeploymentSpec {
	if o == nil || IsNil(o.Specs) {
		var ret []ApiSearchDeploymentSpec
		return ret
	}
	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentRequest) GetSpecsOk() ([]ApiSearchDeploymentSpec, bool) {
	if o == nil || IsNil(o.Specs) {
		return nil, false
	}

	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *ApiSearchDeploymentRequest) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given []ApiSearchDeploymentSpec and assigns it to the Specs field.
func (o *ApiSearchDeploymentRequest) SetSpecs(v []ApiSearchDeploymentSpec) {
	o.Specs = v
}

func (o ApiSearchDeploymentRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiSearchDeploymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Specs) {
		toSerialize["specs"] = o.Specs
	}
	return toSerialize, nil
}
