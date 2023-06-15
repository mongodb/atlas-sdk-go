// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CreateOrganizationKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationKey{}

// CreateOrganizationKey struct for CreateOrganizationKey
type CreateOrganizationKey struct {
	// Purpose or explanation provided when someone created this organization API key.
	Desc *string `json:"desc,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this organization.
	Roles []string `json:"roles,omitempty"`
}

// NewCreateOrganizationKey instantiates a new CreateOrganizationKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationKey() *CreateOrganizationKey {
	this := CreateOrganizationKey{}
	return &this
}

// NewCreateOrganizationKeyWithDefaults instantiates a new CreateOrganizationKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationKeyWithDefaults() *CreateOrganizationKey {
	this := CreateOrganizationKey{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *CreateOrganizationKey) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationKey) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *CreateOrganizationKey) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *CreateOrganizationKey) SetDesc(v string) {
	o.Desc = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateOrganizationKey) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationKey) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateOrganizationKey) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateOrganizationKey) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateOrganizationKey) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateOrganizationKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableCreateOrganizationKey struct {
	value *CreateOrganizationKey
	isSet bool
}

func (v NullableCreateOrganizationKey) Get() *CreateOrganizationKey {
	return v.value
}

func (v *NullableCreateOrganizationKey) Set(val *CreateOrganizationKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationKey(val *CreateOrganizationKey) *NullableCreateOrganizationKey {
	return &NullableCreateOrganizationKey{value: val, isSet: true}
}

func (v NullableCreateOrganizationKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
