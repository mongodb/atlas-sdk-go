// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AddOrRemoveOrgRole struct for AddOrRemoveOrgRole
type AddOrRemoveOrgRole struct {
	// Organization-level role.
	OrgRole string `json:"orgRole"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AddOrRemoveOrgRole) MarshalJSON() ([]byte, error) {
	type noMethod AddOrRemoveOrgRole
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAddOrRemoveOrgRole instantiates a new AddOrRemoveOrgRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOrRemoveOrgRole(orgRole string) *AddOrRemoveOrgRole {
	this := AddOrRemoveOrgRole{}
	this.OrgRole = orgRole
	return &this
}

// NewAddOrRemoveOrgRoleWithDefaults instantiates a new AddOrRemoveOrgRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOrRemoveOrgRoleWithDefaults() *AddOrRemoveOrgRole {
	this := AddOrRemoveOrgRole{}
	return &this
}

// GetOrgRole returns the OrgRole field value
func (o *AddOrRemoveOrgRole) GetOrgRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgRole
}

// GetOrgRoleOk returns a tuple with the OrgRole field value
// and a boolean to check if the value has been set.
func (o *AddOrRemoveOrgRole) GetOrgRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgRole, true
}

// SetOrgRole sets field value
func (o *AddOrRemoveOrgRole) SetOrgRole(v string) {
	o.OrgRole = v
}
