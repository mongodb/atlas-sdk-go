// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ServiceAccountProjectAccessRole Group access roles
type ServiceAccountProjectAccessRole struct {
	// Id of the group.
	GroupId   string   `json:"groupId"`
	RoleNames []string `json:"roleNames,omitempty"`
}

// NewServiceAccountProjectAccessRole instantiates a new ServiceAccountProjectAccessRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountProjectAccessRole(groupId string) *ServiceAccountProjectAccessRole {
	this := ServiceAccountProjectAccessRole{}
	this.GroupId = groupId
	return &this
}

// NewServiceAccountProjectAccessRoleWithDefaults instantiates a new ServiceAccountProjectAccessRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountProjectAccessRoleWithDefaults() *ServiceAccountProjectAccessRole {
	this := ServiceAccountProjectAccessRole{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ServiceAccountProjectAccessRole) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountProjectAccessRole) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ServiceAccountProjectAccessRole) SetGroupId(v string) {
	o.GroupId = v
}

// GetRoleNames returns the RoleNames field value if set, zero value otherwise
func (o *ServiceAccountProjectAccessRole) GetRoleNames() []string {
	if o == nil || IsNil(o.RoleNames) {
		var ret []string
		return ret
	}
	return o.RoleNames
}

// GetRoleNamesOk returns a tuple with the RoleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountProjectAccessRole) GetRoleNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleNames) {
		return nil, false
	}

	return o.RoleNames, true
}

// HasRoleNames returns a boolean if a field has been set.
func (o *ServiceAccountProjectAccessRole) HasRoleNames() bool {
	if o != nil && !IsNil(o.RoleNames) {
		return true
	}

	return false
}

// SetRoleNames gets a reference to the given []string and assigns it to the RoleNames field.
func (o *ServiceAccountProjectAccessRole) SetRoleNames(v []string) {
	o.RoleNames = v
}

func (o ServiceAccountProjectAccessRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountProjectAccessRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	if !IsNil(o.RoleNames) {
		toSerialize["roleNames"] = o.RoleNames
	}
	return toSerialize, nil
}
