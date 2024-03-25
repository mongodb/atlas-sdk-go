// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GroupServiceAccountRoleAssignment struct for GroupServiceAccountRoleAssignment
type GroupServiceAccountRoleAssignment struct {
	// Group access roles.
	RoleNames *[]string `json:"roleNames,omitempty"`
}

// NewGroupServiceAccountRoleAssignment instantiates a new GroupServiceAccountRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupServiceAccountRoleAssignment() *GroupServiceAccountRoleAssignment {
	this := GroupServiceAccountRoleAssignment{}
	return &this
}

// NewGroupServiceAccountRoleAssignmentWithDefaults instantiates a new GroupServiceAccountRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupServiceAccountRoleAssignmentWithDefaults() *GroupServiceAccountRoleAssignment {
	this := GroupServiceAccountRoleAssignment{}
	return &this
}

// GetRoleNames returns the RoleNames field value if set, zero value otherwise
func (o *GroupServiceAccountRoleAssignment) GetRoleNames() []string {
	if o == nil || IsNil(o.RoleNames) {
		var ret []string
		return ret
	}
	return *o.RoleNames
}

// GetRoleNamesOk returns a tuple with the RoleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRoleAssignment) GetRoleNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.RoleNames) {
		return nil, false
	}

	return o.RoleNames, true
}

// HasRoleNames returns a boolean if a field has been set.
func (o *GroupServiceAccountRoleAssignment) HasRoleNames() bool {
	if o != nil && !IsNil(o.RoleNames) {
		return true
	}

	return false
}

// SetRoleNames gets a reference to the given []string and assigns it to the RoleNames field.
func (o *GroupServiceAccountRoleAssignment) SetRoleNames(v []string) {
	o.RoleNames = &v
}

func (o GroupServiceAccountRoleAssignment) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupServiceAccountRoleAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleNames) {
		toSerialize["roleNames"] = o.RoleNames
	}
	return toSerialize, nil
}
