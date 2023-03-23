/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the GroupInvitationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupInvitationRequest{}

// GroupInvitationRequest struct for GroupInvitationRequest
type GroupInvitationRequest struct {
	// One or more organization or project level roles to assign to the MongoDB Cloud user.
	Roles []string `json:"roles,omitempty"`
	// Email address of the MongoDB Cloud user invited to the specified project.
	Username *string `json:"username,omitempty"`
}

// NewGroupInvitationRequest instantiates a new GroupInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupInvitationRequest() *GroupInvitationRequest {
	this := GroupInvitationRequest{}
	return &this
}

// NewGroupInvitationRequestWithDefaults instantiates a new GroupInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupInvitationRequestWithDefaults() *GroupInvitationRequest {
	this := GroupInvitationRequest{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *GroupInvitationRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInvitationRequest) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GroupInvitationRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GroupInvitationRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GroupInvitationRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInvitationRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GroupInvitationRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GroupInvitationRequest) SetUsername(v string) {
	o.Username = &v
}

func (o GroupInvitationRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupInvitationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableGroupInvitationRequest struct {
	value *GroupInvitationRequest
	isSet bool
}

func (v NullableGroupInvitationRequest) Get() *GroupInvitationRequest {
	return v.value
}

func (v *NullableGroupInvitationRequest) Set(val *GroupInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupInvitationRequest(val *GroupInvitationRequest) *NullableGroupInvitationRequest {
	return &NullableGroupInvitationRequest{value: val, isSet: true}
}

func (v NullableGroupInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

