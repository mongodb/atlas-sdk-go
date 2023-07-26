// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// CreateAtlasProjectApiKey struct for CreateAtlasProjectApiKey
type CreateAtlasProjectApiKey struct {
	// Purpose or explanation provided when someone created this project API key.
	Desc *string `json:"desc,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project.
	Roles []string `json:"roles,omitempty"`
}

// NewCreateAtlasProjectApiKey instantiates a new CreateAtlasProjectApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAtlasProjectApiKey() *CreateAtlasProjectApiKey {
	this := CreateAtlasProjectApiKey{}
	return &this
}

// NewCreateAtlasProjectApiKeyWithDefaults instantiates a new CreateAtlasProjectApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAtlasProjectApiKeyWithDefaults() *CreateAtlasProjectApiKey {
	this := CreateAtlasProjectApiKey{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise
func (o *CreateAtlasProjectApiKey) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAtlasProjectApiKey) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}

	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *CreateAtlasProjectApiKey) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *CreateAtlasProjectApiKey) SetDesc(v string) {
	o.Desc = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *CreateAtlasProjectApiKey) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAtlasProjectApiKey) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateAtlasProjectApiKey) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateAtlasProjectApiKey) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateAtlasProjectApiKey) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateAtlasProjectApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}
