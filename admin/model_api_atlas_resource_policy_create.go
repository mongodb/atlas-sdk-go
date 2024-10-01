// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ApiAtlasResourcePolicyCreate struct for ApiAtlasResourcePolicyCreate
type ApiAtlasResourcePolicyCreate struct {
	// Human-readable label that describes the atlas resource policy.
	Name string `json:"name"`
	// List of policies that make up the atlas resource policy.
	Policies []ApiAtlasPolicyCreate `json:"policies"`
}

// NewApiAtlasResourcePolicyCreate instantiates a new ApiAtlasResourcePolicyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasResourcePolicyCreate(name string, policies []ApiAtlasPolicyCreate) *ApiAtlasResourcePolicyCreate {
	this := ApiAtlasResourcePolicyCreate{}
	this.Name = name
	this.Policies = policies
	return &this
}

// NewApiAtlasResourcePolicyCreateWithDefaults instantiates a new ApiAtlasResourcePolicyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasResourcePolicyCreateWithDefaults() *ApiAtlasResourcePolicyCreate {
	this := ApiAtlasResourcePolicyCreate{}
	return &this
}

// GetName returns the Name field value
func (o *ApiAtlasResourcePolicyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiAtlasResourcePolicyCreate) SetName(v string) {
	o.Name = v
}

// GetPolicies returns the Policies field value
func (o *ApiAtlasResourcePolicyCreate) GetPolicies() []ApiAtlasPolicyCreate {
	if o == nil {
		var ret []ApiAtlasPolicyCreate
		return ret
	}

	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyCreate) GetPoliciesOk() (*[]ApiAtlasPolicyCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policies, true
}

// SetPolicies sets field value
func (o *ApiAtlasResourcePolicyCreate) SetPolicies(v []ApiAtlasPolicyCreate) {
	o.Policies = v
}

func (o *ApiAtlasResourcePolicyCreate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o *ApiAtlasResourcePolicyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["policies"] = o.Policies
	return toSerialize, nil
}
