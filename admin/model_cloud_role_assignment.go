// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudRoleAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRoleAssignment{}

// CloudRoleAssignment MongoDB Cloud user's roles and the corresponding organization or project to which that role applies. Each role can apply to one organization or one project but not both.
type CloudRoleAssignment struct {
	// Unique 24-hexadecimal digit string that identifies the project to which this role belongs. You can set a value for this parameter or **orgId** but not both in the same request.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which this role belongs. You can set a value for this parameter or **groupId** but not both in the same request.
	OrgId *string `json:"orgId,omitempty"`
	// Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team. These roles include organization- and project-level roles.  Organization Roles  * ORG_OWNER * ORG_MEMBER * ORG_GROUP_CREATOR * ORG_BILLING_ADMIN * ORG_READ_ONLY  Project Roles  * GROUP_CLUSTER_MANAGER * GROUP_DATA_ACCESS_ADMIN * GROUP_DATA_ACCESS_READ_ONLY * GROUP_DATA_ACCESS_READ_WRITE * GROUP_OWNER * GROUP_READ_ONLY * GROUP_SEARCH_INDEX_EDITOR
	RoleName *string `json:"roleName,omitempty"`
}

// NewCloudRoleAssignment instantiates a new CloudRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRoleAssignment() *CloudRoleAssignment {
	this := CloudRoleAssignment{}
	return &this
}

// NewCloudRoleAssignmentWithDefaults instantiates a new CloudRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRoleAssignmentWithDefaults() *CloudRoleAssignment {
	this := CloudRoleAssignment{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CloudRoleAssignment) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRoleAssignment) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CloudRoleAssignment) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CloudRoleAssignment) SetGroupId(v string) {
	o.GroupId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *CloudRoleAssignment) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRoleAssignment) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *CloudRoleAssignment) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *CloudRoleAssignment) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *CloudRoleAssignment) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRoleAssignment) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *CloudRoleAssignment) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *CloudRoleAssignment) SetRoleName(v string) {
	o.RoleName = &v
}

func (o CloudRoleAssignment) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudRoleAssignment) ToMap() (map[string]interface{}, error) {
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

type NullableCloudRoleAssignment struct {
	value *CloudRoleAssignment
	isSet bool
}

func (v NullableCloudRoleAssignment) Get() *CloudRoleAssignment {
	return v.value
}

func (v *NullableCloudRoleAssignment) Set(val *CloudRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRoleAssignment(val *CloudRoleAssignment) *NullableCloudRoleAssignment {
	return &NullableCloudRoleAssignment{value: val, isSet: true}
}

func (v NullableCloudRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
