// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrganizationInvitationUpdateRequest struct for OrganizationInvitationUpdateRequest
type OrganizationInvitationUpdateRequest struct {
	// List of projects that the user will be added to when they accept their invitation to the organization.
	GroupRoleAssignments *[]OrganizationInvitationGroupRoleAssignmentsRequest `json:"groupRoleAssignments,omitempty"`
	// One or more organization level roles to assign to the MongoDB Cloud user.
	Roles *[]string `json:"roles,omitempty"`
	// List of teams to which you want to invite the desired MongoDB Cloud user.
	TeamIds *[]string `json:"teamIds,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OrganizationInvitationUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod OrganizationInvitationUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOrganizationInvitationUpdateRequest instantiates a new OrganizationInvitationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitationUpdateRequest() *OrganizationInvitationUpdateRequest {
	this := OrganizationInvitationUpdateRequest{}
	return &this
}

// NewOrganizationInvitationUpdateRequestWithDefaults instantiates a new OrganizationInvitationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationUpdateRequestWithDefaults() *OrganizationInvitationUpdateRequest {
	this := OrganizationInvitationUpdateRequest{}
	return &this
}

// GetGroupRoleAssignments returns the GroupRoleAssignments field value if set, zero value otherwise
func (o *OrganizationInvitationUpdateRequest) GetGroupRoleAssignments() []OrganizationInvitationGroupRoleAssignmentsRequest {
	if o == nil || IsNil(o.GroupRoleAssignments) {
		var ret []OrganizationInvitationGroupRoleAssignmentsRequest
		return ret
	}
	return *o.GroupRoleAssignments
}

// GetGroupRoleAssignmentsOk returns a tuple with the GroupRoleAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationUpdateRequest) GetGroupRoleAssignmentsOk() (*[]OrganizationInvitationGroupRoleAssignmentsRequest, bool) {
	if o == nil || IsNil(o.GroupRoleAssignments) {
		return nil, false
	}

	return o.GroupRoleAssignments, true
}

// HasGroupRoleAssignments returns a boolean if a field has been set.
func (o *OrganizationInvitationUpdateRequest) HasGroupRoleAssignments() bool {
	if o != nil && !IsNil(o.GroupRoleAssignments) {
		return true
	}

	return false
}

// SetGroupRoleAssignments gets a reference to the given []OrganizationInvitationGroupRoleAssignmentsRequest and assigns it to the GroupRoleAssignments field.
func (o *OrganizationInvitationUpdateRequest) SetGroupRoleAssignments(v []OrganizationInvitationGroupRoleAssignmentsRequest) {
	o.GroupRoleAssignments = &v
	o.NullFields = removeNullField(o.NullFields, "GroupRoleAssignments")
}

// SetGroupRoleAssignmentsNil sets GroupRoleAssignments to an explicit JSON null when marshaled.
func (o *OrganizationInvitationUpdateRequest) SetGroupRoleAssignmentsNil() {
	o.GroupRoleAssignments = nil
	o.NullFields = addNullField(o.NullFields, "GroupRoleAssignments")
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *OrganizationInvitationUpdateRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationUpdateRequest) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrganizationInvitationUpdateRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrganizationInvitationUpdateRequest) SetRoles(v []string) {
	o.Roles = &v
	o.NullFields = removeNullField(o.NullFields, "Roles")
}

// SetRolesNil sets Roles to an explicit JSON null when marshaled.
func (o *OrganizationInvitationUpdateRequest) SetRolesNil() {
	o.Roles = nil
	o.NullFields = addNullField(o.NullFields, "Roles")
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise
func (o *OrganizationInvitationUpdateRequest) GetTeamIds() []string {
	if o == nil || IsNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return *o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationUpdateRequest) GetTeamIdsOk() (*[]string, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}

	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *OrganizationInvitationUpdateRequest) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *OrganizationInvitationUpdateRequest) SetTeamIds(v []string) {
	o.TeamIds = &v
	o.NullFields = removeNullField(o.NullFields, "TeamIds")
}

// SetTeamIdsNil sets TeamIds to an explicit JSON null when marshaled.
func (o *OrganizationInvitationUpdateRequest) SetTeamIdsNil() {
	o.TeamIds = nil
	o.NullFields = addNullField(o.NullFields, "TeamIds")
}
