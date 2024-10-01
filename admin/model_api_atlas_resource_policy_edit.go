// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ApiAtlasResourcePolicyEdit struct for ApiAtlasResourcePolicyEdit
type ApiAtlasResourcePolicyEdit struct {
	// Human-readable label that describes the atlas resource policy.
	Name *string `json:"name,omitempty"`
	// List of policies that make up the atlas resource policy.
	Policies *[]ApiAtlasPolicyCreate `json:"policies,omitempty"`
}

// NewApiAtlasResourcePolicyEdit instantiates a new ApiAtlasResourcePolicyEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasResourcePolicyEdit() *ApiAtlasResourcePolicyEdit {
	this := ApiAtlasResourcePolicyEdit{}
	return &this
}

// NewApiAtlasResourcePolicyEditWithDefaults instantiates a new ApiAtlasResourcePolicyEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasResourcePolicyEditWithDefaults() *ApiAtlasResourcePolicyEdit {
	this := ApiAtlasResourcePolicyEdit{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise
func (o *ApiAtlasResourcePolicyEdit) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyEdit) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAtlasResourcePolicyEdit) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAtlasResourcePolicyEdit) SetName(v string) {
	o.Name = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise
func (o *ApiAtlasResourcePolicyEdit) GetPolicies() []ApiAtlasPolicyCreate {
	if o == nil || IsNil(o.Policies) {
		var ret []ApiAtlasPolicyCreate
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyEdit) GetPoliciesOk() (*[]ApiAtlasPolicyCreate, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}

	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *ApiAtlasResourcePolicyEdit) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []ApiAtlasPolicyCreate and assigns it to the Policies field.
func (o *ApiAtlasResourcePolicyEdit) SetPolicies(v []ApiAtlasPolicyCreate) {
	o.Policies = &v
}

func (o ApiAtlasResourcePolicyEdit) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiAtlasResourcePolicyEdit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return toSerialize, nil
}