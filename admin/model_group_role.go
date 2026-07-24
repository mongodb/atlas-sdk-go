// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GroupRole struct for GroupRole
type GroupRole struct {
	// Unique 24-hexadecimal digit string that identifies the project to which this role belongs.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team. These roles include project-level roles.
	GroupRole *string `json:"groupRole,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *GroupRole) MarshalJSON() ([]byte, error) {
	type noMethod GroupRole
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGroupRole instantiates a new GroupRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRole() *GroupRole {
	this := GroupRole{}
	return &this
}

// NewGroupRoleWithDefaults instantiates a new GroupRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRoleWithDefaults() *GroupRole {
	this := GroupRole{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *GroupRole) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupRole) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupRole) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *GroupRole) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetGroupRole returns the GroupRole field value if set, zero value otherwise
func (o *GroupRole) GetGroupRole() string {
	if o == nil || IsNil(o.GroupRole) {
		var ret string
		return ret
	}
	return *o.GroupRole
}

// GetGroupRoleOk returns a tuple with the GroupRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetGroupRoleOk() (*string, bool) {
	if o == nil || IsNil(o.GroupRole) {
		return nil, false
	}

	return o.GroupRole, true
}

// HasGroupRole returns a boolean if a field has been set.
func (o *GroupRole) HasGroupRole() bool {
	if o != nil && !IsNil(o.GroupRole) {
		return true
	}

	return false
}

// SetGroupRole gets a reference to the given string and assigns it to the GroupRole field.
func (o *GroupRole) SetGroupRole(v string) {
	o.GroupRole = &v
	o.NullFields = removeNullField(o.NullFields, "GroupRole")
}

// SetGroupRoleNil sets GroupRole to an explicit JSON null when marshaled.
func (o *GroupRole) SetGroupRoleNil() {
	o.GroupRole = nil
	o.NullFields = addNullField(o.NullFields, "GroupRole")
}
