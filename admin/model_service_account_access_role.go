// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ServiceAccountAccessRole struct for ServiceAccountAccessRole
type ServiceAccountAccessRole struct {
	// Project Id.
	GroupId *string `json:"groupId,omitempty"`
	// Organization Id.
	OrgId *string `json:"orgId,omitempty"`
	// Access Role.
	RoleName *string `json:"roleName,omitempty"`
}

// NewServiceAccountAccessRole instantiates a new ServiceAccountAccessRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountAccessRole() *ServiceAccountAccessRole {
	this := ServiceAccountAccessRole{}
	return &this
}

// NewServiceAccountAccessRoleWithDefaults instantiates a new ServiceAccountAccessRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountAccessRoleWithDefaults() *ServiceAccountAccessRole {
	this := ServiceAccountAccessRole{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *ServiceAccountAccessRole) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessRole) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ServiceAccountAccessRole) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ServiceAccountAccessRole) SetGroupId(v string) {
	o.GroupId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise
func (o *ServiceAccountAccessRole) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessRole) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}

	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ServiceAccountAccessRole) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ServiceAccountAccessRole) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise
func (o *ServiceAccountAccessRole) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessRole) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}

	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *ServiceAccountAccessRole) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *ServiceAccountAccessRole) SetRoleName(v string) {
	o.RoleName = &v
}

func (o ServiceAccountAccessRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountAccessRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !IsNil(o.RoleName) {
		toSerialize["roleName"] = o.RoleName
	}
	return toSerialize, nil
}
