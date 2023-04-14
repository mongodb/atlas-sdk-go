/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the UpdateCustomDBRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomDBRole{}

// UpdateCustomDBRole struct for UpdateCustomDBRole
type UpdateCustomDBRole struct {
	// List of the individual privilege actions that the role grants.
	Actions []DBAction `json:"actions,omitempty"`
	// List of the built-in roles that this custom role inherits.
	InheritedRoles []InheritedRole `json:"inheritedRoles,omitempty"`
}

// NewUpdateCustomDBRole instantiates a new UpdateCustomDBRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomDBRole() *UpdateCustomDBRole {
	this := UpdateCustomDBRole{}
	return &this
}

// NewUpdateCustomDBRoleWithDefaults instantiates a new UpdateCustomDBRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomDBRoleWithDefaults() *UpdateCustomDBRole {
	this := UpdateCustomDBRole{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *UpdateCustomDBRole) GetActions() []DBAction {
	if o == nil || IsNil(o.Actions) {
		var ret []DBAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomDBRole) GetActionsOk() ([]DBAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *UpdateCustomDBRole) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []DBAction and assigns it to the Actions field.
func (o *UpdateCustomDBRole) SetActions(v []DBAction) {
	o.Actions = v
}

// GetInheritedRoles returns the InheritedRoles field value if set, zero value otherwise.
func (o *UpdateCustomDBRole) GetInheritedRoles() []InheritedRole {
	if o == nil || IsNil(o.InheritedRoles) {
		var ret []InheritedRole
		return ret
	}
	return o.InheritedRoles
}

// GetInheritedRolesOk returns a tuple with the InheritedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomDBRole) GetInheritedRolesOk() ([]InheritedRole, bool) {
	if o == nil || IsNil(o.InheritedRoles) {
		return nil, false
	}
	return o.InheritedRoles, true
}

// HasInheritedRoles returns a boolean if a field has been set.
func (o *UpdateCustomDBRole) HasInheritedRoles() bool {
	if o != nil && !IsNil(o.InheritedRoles) {
		return true
	}

	return false
}

// SetInheritedRoles gets a reference to the given []InheritedRole and assigns it to the InheritedRoles field.
func (o *UpdateCustomDBRole) SetInheritedRoles(v []InheritedRole) {
	o.InheritedRoles = v
}

func (o UpdateCustomDBRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o UpdateCustomDBRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.InheritedRoles) {
		toSerialize["inheritedRoles"] = o.InheritedRoles
	}
	return toSerialize, nil
}

type NullableUpdateCustomDBRole struct {
	value *UpdateCustomDBRole
	isSet bool
}

func (v NullableUpdateCustomDBRole) Get() *UpdateCustomDBRole {
	return v.value
}

func (v *NullableUpdateCustomDBRole) Set(val *UpdateCustomDBRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomDBRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomDBRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomDBRole(val *UpdateCustomDBRole) *NullableUpdateCustomDBRole {
	return &NullableUpdateCustomDBRole{value: val, isSet: true}
}

func (v NullableUpdateCustomDBRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomDBRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


