// Code based on the AtlasAPI V2 OpenAPI file

package admin

// TeamRole struct for TeamRole
type TeamRole struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// One or more project-level roles to assign to the team.
	RoleNames *[]string `json:"roleNames,omitempty"`
	// Unique 24-hexadecimal character string that identifies the team.
	TeamId *string `json:"teamId,omitempty"`
}

// NewTeamRole instantiates a new TeamRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamRole() *TeamRole {
	this := TeamRole{}
	return &this
}

// NewTeamRoleWithDefaults instantiates a new TeamRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamRoleWithDefaults() *TeamRole {
	this := TeamRole{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *TeamRole) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRole) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TeamRole) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *TeamRole) SetLinks(v []Link) {
	o.Links = &v
}

// GetRoleNames returns the RoleNames field value if set, zero value otherwise
func (o *TeamRole) GetRoleNames() []string {
	if o == nil || IsNil(o.RoleNames) {
		var ret []string
		return ret
	}
	return *o.RoleNames
}

// GetRoleNamesOk returns a tuple with the RoleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRole) GetRoleNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.RoleNames) {
		return nil, false
	}

	return o.RoleNames, true
}

// HasRoleNames returns a boolean if a field has been set.
func (o *TeamRole) HasRoleNames() bool {
	if o != nil && !IsNil(o.RoleNames) {
		return true
	}

	return false
}

// SetRoleNames gets a reference to the given []string and assigns it to the RoleNames field.
func (o *TeamRole) SetRoleNames(v []string) {
	o.RoleNames = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise
func (o *TeamRole) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRole) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}

	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *TeamRole) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *TeamRole) SetTeamId(v string) {
	o.TeamId = &v
}
