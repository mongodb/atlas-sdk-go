// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GroupRoleAssignment struct for GroupRoleAssignment
type GroupRoleAssignment struct {
	// Unique 24-hexadecimal digit string that identifies the project to which these roles belong.
	GroupId *string `json:"groupId,omitempty"`
	// One or more project-level roles assigned to the MongoDB Cloud user.
	GroupRoles *[]string `json:"groupRoles,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *GroupRoleAssignment) MarshalJSON() ([]byte, error) {
	type noMethod GroupRoleAssignment
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGroupRoleAssignment instantiates a new GroupRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRoleAssignment() *GroupRoleAssignment {
	this := GroupRoleAssignment{}
	return &this
}

// NewGroupRoleAssignmentWithDefaults instantiates a new GroupRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRoleAssignmentWithDefaults() *GroupRoleAssignment {
	this := GroupRoleAssignment{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *GroupRoleAssignment) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRoleAssignment) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupRoleAssignment) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupRoleAssignment) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *GroupRoleAssignment) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetGroupRoles returns the GroupRoles field value if set, zero value otherwise
func (o *GroupRoleAssignment) GetGroupRoles() []string {
	if o == nil || IsNil(o.GroupRoles) {
		var ret []string
		return ret
	}
	return *o.GroupRoles
}

// GetGroupRolesOk returns a tuple with the GroupRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRoleAssignment) GetGroupRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.GroupRoles) {
		return nil, false
	}

	return o.GroupRoles, true
}

// HasGroupRoles returns a boolean if a field has been set.
func (o *GroupRoleAssignment) HasGroupRoles() bool {
	if o != nil && !IsNil(o.GroupRoles) {
		return true
	}

	return false
}

// SetGroupRoles gets a reference to the given []string and assigns it to the GroupRoles field.
func (o *GroupRoleAssignment) SetGroupRoles(v []string) {
	o.GroupRoles = &v
	o.NullFields = removeNullField(o.NullFields, "GroupRoles")
}

// SetGroupRolesNil sets GroupRoles to an explicit JSON null when marshaled.
func (o *GroupRoleAssignment) SetGroupRolesNil() {
	o.GroupRoles = nil
	o.NullFields = addNullField(o.NullFields, "GroupRoles")
}
