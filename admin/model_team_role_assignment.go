// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// TeamRoleAssignment struct for TeamRoleAssignment
type TeamRoleAssignment struct {
	Roles  []string `json:"roles,omitempty"`
	TeamId *string  `json:"teamId,omitempty"`
}

// NewTeamRoleAssignment instantiates a new TeamRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamRoleAssignment() *TeamRoleAssignment {
	this := TeamRoleAssignment{}
	return &this
}

// NewTeamRoleAssignmentWithDefaults instantiates a new TeamRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamRoleAssignmentWithDefaults() *TeamRoleAssignment {
	this := TeamRoleAssignment{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *TeamRoleAssignment) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoleAssignment) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *TeamRoleAssignment) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *TeamRoleAssignment) SetRoles(v []string) {
	o.Roles = v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise
func (o *TeamRoleAssignment) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoleAssignment) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}

	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *TeamRoleAssignment) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *TeamRoleAssignment) SetTeamId(v string) {
	o.TeamId = &v
}

func (o TeamRoleAssignment) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TeamRoleAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.TeamId) {
		toSerialize["teamId"] = o.TeamId
	}
	return toSerialize, nil
}
