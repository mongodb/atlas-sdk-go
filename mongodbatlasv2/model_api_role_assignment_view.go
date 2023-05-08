/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ApiRoleAssignmentView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRoleAssignmentView{}

// ApiRoleAssignmentView MongoDB Cloud user's roles and the corresponding organization or project to which that role applies. Each role can apply to one organization or one project but not both.
type ApiRoleAssignmentView struct {
	// Unique 24-hexadecimal digit string that identifies the project to which this role belongs. You can set a value for this parameter or **orgId** but not both in the same request.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which this role belongs. You can set a value for this parameter or **groupId** but not both in the same request.
	OrgId *string `json:"orgId,omitempty"`
	// Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team. These roles include organization- and project-level roles.  Organization Roles  * ORG_OWNER * ORG_MEMBER * ORG_GROUP_CREATOR * ORG_BILLING_ADMIN * ORG_READ_ONLY  Project Roles  * GROUP_CLUSTER_MANAGER * GROUP_DATA_ACCESS_ADMIN * GROUP_DATA_ACCESS_READ_ONLY * GROUP_DATA_ACCESS_READ_WRITE * GROUP_OWNER * GROUP_READ_ONLY * GROUP_SEARCH_INDEX_EDITOR  
	RoleName *string `json:"roleName,omitempty"`
}

// NewApiRoleAssignmentView instantiates a new ApiRoleAssignmentView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRoleAssignmentView() *ApiRoleAssignmentView {
	this := ApiRoleAssignmentView{}
	return &this
}

// NewApiRoleAssignmentViewWithDefaults instantiates a new ApiRoleAssignmentView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRoleAssignmentViewWithDefaults() *ApiRoleAssignmentView {
	this := ApiRoleAssignmentView{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiRoleAssignmentView) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleAssignmentView) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiRoleAssignmentView) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiRoleAssignmentView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ApiRoleAssignmentView) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleAssignmentView) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ApiRoleAssignmentView) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ApiRoleAssignmentView) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *ApiRoleAssignmentView) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleAssignmentView) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *ApiRoleAssignmentView) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *ApiRoleAssignmentView) SetRoleName(v string) {
	o.RoleName = &v
}

func (o ApiRoleAssignmentView) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiRoleAssignmentView) ToMap() (map[string]interface{}, error) {
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

type NullableApiRoleAssignmentView struct {
	value *ApiRoleAssignmentView
	isSet bool
}

func (v NullableApiRoleAssignmentView) Get() *ApiRoleAssignmentView {
	return v.value
}

func (v *NullableApiRoleAssignmentView) Set(val *ApiRoleAssignmentView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRoleAssignmentView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRoleAssignmentView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRoleAssignmentView(val *ApiRoleAssignmentView) *NullableApiRoleAssignmentView {
	return &NullableApiRoleAssignmentView{value: val, isSet: true}
}

func (v NullableApiRoleAssignmentView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRoleAssignmentView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


