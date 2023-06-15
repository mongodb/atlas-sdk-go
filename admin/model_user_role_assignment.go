// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the UserRoleAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRoleAssignment{}

// UserRoleAssignment struct for UserRoleAssignment
type UserRoleAssignment struct {
	// Unique 24-hexadecimal digit string that identifies the organization API key.
	ApiUserId *string `json:"apiUserId,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project.
	Roles []string `json:"roles,omitempty"`
}

// NewUserRoleAssignment instantiates a new UserRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRoleAssignment() *UserRoleAssignment {
	this := UserRoleAssignment{}
	return &this
}

// NewUserRoleAssignmentWithDefaults instantiates a new UserRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRoleAssignmentWithDefaults() *UserRoleAssignment {
	this := UserRoleAssignment{}
	return &this
}

// GetApiUserId returns the ApiUserId field value if set, zero value otherwise.
func (o *UserRoleAssignment) GetApiUserId() string {
	if o == nil || IsNil(o.ApiUserId) {
		var ret string
		return ret
	}
	return *o.ApiUserId
}

// GetApiUserIdOk returns a tuple with the ApiUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRoleAssignment) GetApiUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiUserId) {
		return nil, false
	}
	return o.ApiUserId, true
}

// HasApiUserId returns a boolean if a field has been set.
func (o *UserRoleAssignment) HasApiUserId() bool {
	if o != nil && !IsNil(o.ApiUserId) {
		return true
	}

	return false
}

// SetApiUserId gets a reference to the given string and assigns it to the ApiUserId field.
func (o *UserRoleAssignment) SetApiUserId(v string) {
	o.ApiUserId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserRoleAssignment) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRoleAssignment) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserRoleAssignment) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserRoleAssignment) SetRoles(v []string) {
	o.Roles = v
}

func (o UserRoleAssignment) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o UserRoleAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableUserRoleAssignment struct {
	value *UserRoleAssignment
	isSet bool
}

func (v NullableUserRoleAssignment) Get() *UserRoleAssignment {
	return v.value
}

func (v *NullableUserRoleAssignment) Set(val *UserRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRoleAssignment(val *UserRoleAssignment) *NullableUserRoleAssignment {
	return &NullableUserRoleAssignment{value: val, isSet: true}
}

func (v NullableUserRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
