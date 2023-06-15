// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CreateProjectKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectKey{}

// CreateProjectKey struct for CreateProjectKey
type CreateProjectKey struct {
	// Purpose or explanation provided when someone created this project API key.
	Desc *string `json:"desc,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project.
	Roles []string `json:"roles,omitempty"`
}

// NewCreateProjectKey instantiates a new CreateProjectKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectKey() *CreateProjectKey {
	this := CreateProjectKey{}
	return &this
}

// NewCreateProjectKeyWithDefaults instantiates a new CreateProjectKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectKeyWithDefaults() *CreateProjectKey {
	this := CreateProjectKey{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *CreateProjectKey) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectKey) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *CreateProjectKey) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *CreateProjectKey) SetDesc(v string) {
	o.Desc = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateProjectKey) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectKey) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateProjectKey) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateProjectKey) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateProjectKey) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateProjectKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableCreateProjectKey struct {
	value *CreateProjectKey
	isSet bool
}

func (v NullableCreateProjectKey) Get() *CreateProjectKey {
	return v.value
}

func (v *NullableCreateProjectKey) Set(val *CreateProjectKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectKey(val *CreateProjectKey) *NullableCreateProjectKey {
	return &NullableCreateProjectKey{value: val, isSet: true}
}

func (v NullableCreateProjectKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
