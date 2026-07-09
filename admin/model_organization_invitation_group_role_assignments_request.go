// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrganizationInvitationGroupRoleAssignmentsRequest struct for OrganizationInvitationGroupRoleAssignmentsRequest
type OrganizationInvitationGroupRoleAssignmentsRequest struct {
	// Unique 24-hexadecimal digit string that identifies the project to which these roles belong.
	GroupId *string `json:"groupId,omitempty"`
	// One or more project-level roles to assign to the MongoDB Cloud user.
	Roles *[]string `json:"roles,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) MarshalJSON() ([]byte, error) {
	type noMethod OrganizationInvitationGroupRoleAssignmentsRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOrganizationInvitationGroupRoleAssignmentsRequest instantiates a new OrganizationInvitationGroupRoleAssignmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitationGroupRoleAssignmentsRequest() *OrganizationInvitationGroupRoleAssignmentsRequest {
	this := OrganizationInvitationGroupRoleAssignmentsRequest{}
	return &this
}

// NewOrganizationInvitationGroupRoleAssignmentsRequestWithDefaults instantiates a new OrganizationInvitationGroupRoleAssignmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationGroupRoleAssignmentsRequestWithDefaults() *OrganizationInvitationGroupRoleAssignmentsRequest {
	this := OrganizationInvitationGroupRoleAssignmentsRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = append(o.NullFields, "GroupId")
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) SetRoles(v []string) {
	o.Roles = &v
}
