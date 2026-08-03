// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GroupMcpConfigUpdateRequest struct for GroupMcpConfigUpdateRequest
type GroupMcpConfigUpdateRequest struct {
	// List of IP access list entries that define allowed source addresses for this MCP configuration. If provided, replaces the existing IP access list.
	IpAccessList *[]ServiceAccountIPAccessListEntry `json:"ipAccessList,omitempty"`
	// Updated human-readable name for this MCP configuration.
	McpConfigName *string `json:"mcpConfigName,omitempty"`
	// List of project roles associated with this MCP configuration. If provided, replaces the existing list of roles.
	Roles *[]string `json:"roles,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *GroupMcpConfigUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod GroupMcpConfigUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGroupMcpConfigUpdateRequest instantiates a new GroupMcpConfigUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMcpConfigUpdateRequest() *GroupMcpConfigUpdateRequest {
	this := GroupMcpConfigUpdateRequest{}
	return &this
}

// NewGroupMcpConfigUpdateRequestWithDefaults instantiates a new GroupMcpConfigUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMcpConfigUpdateRequestWithDefaults() *GroupMcpConfigUpdateRequest {
	this := GroupMcpConfigUpdateRequest{}
	return &this
}

// GetIpAccessList returns the IpAccessList field value if set, zero value otherwise
func (o *GroupMcpConfigUpdateRequest) GetIpAccessList() []ServiceAccountIPAccessListEntry {
	if o == nil || IsNil(o.IpAccessList) {
		var ret []ServiceAccountIPAccessListEntry
		return ret
	}
	return *o.IpAccessList
}

// GetIpAccessListOk returns a tuple with the IpAccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigUpdateRequest) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool) {
	if o == nil || IsNil(o.IpAccessList) {
		return nil, false
	}

	return o.IpAccessList, true
}

// HasIpAccessList returns a boolean if a field has been set.
func (o *GroupMcpConfigUpdateRequest) HasIpAccessList() bool {
	if o != nil && !IsNil(o.IpAccessList) {
		return true
	}

	return false
}

// SetIpAccessList gets a reference to the given []ServiceAccountIPAccessListEntry and assigns it to the IpAccessList field.
func (o *GroupMcpConfigUpdateRequest) SetIpAccessList(v []ServiceAccountIPAccessListEntry) {
	o.IpAccessList = &v
	o.NullFields = removeNullField(o.NullFields, "IpAccessList")
}

// SetIpAccessListNil sets IpAccessList to an explicit JSON null when marshaled.
func (o *GroupMcpConfigUpdateRequest) SetIpAccessListNil() {
	o.IpAccessList = nil
	o.NullFields = addNullField(o.NullFields, "IpAccessList")
}

// GetMcpConfigName returns the McpConfigName field value if set, zero value otherwise
func (o *GroupMcpConfigUpdateRequest) GetMcpConfigName() string {
	if o == nil || IsNil(o.McpConfigName) {
		var ret string
		return ret
	}
	return *o.McpConfigName
}

// GetMcpConfigNameOk returns a tuple with the McpConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigUpdateRequest) GetMcpConfigNameOk() (*string, bool) {
	if o == nil || IsNil(o.McpConfigName) {
		return nil, false
	}

	return o.McpConfigName, true
}

// HasMcpConfigName returns a boolean if a field has been set.
func (o *GroupMcpConfigUpdateRequest) HasMcpConfigName() bool {
	if o != nil && !IsNil(o.McpConfigName) {
		return true
	}

	return false
}

// SetMcpConfigName gets a reference to the given string and assigns it to the McpConfigName field.
func (o *GroupMcpConfigUpdateRequest) SetMcpConfigName(v string) {
	o.McpConfigName = &v
	o.NullFields = removeNullField(o.NullFields, "McpConfigName")
}

// SetMcpConfigNameNil sets McpConfigName to an explicit JSON null when marshaled.
func (o *GroupMcpConfigUpdateRequest) SetMcpConfigNameNil() {
	o.McpConfigName = nil
	o.NullFields = addNullField(o.NullFields, "McpConfigName")
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *GroupMcpConfigUpdateRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigUpdateRequest) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GroupMcpConfigUpdateRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GroupMcpConfigUpdateRequest) SetRoles(v []string) {
	o.Roles = &v
	o.NullFields = removeNullField(o.NullFields, "Roles")
}

// SetRolesNil sets Roles to an explicit JSON null when marshaled.
func (o *GroupMcpConfigUpdateRequest) SetRolesNil() {
	o.Roles = nil
	o.NullFields = addNullField(o.NullFields, "Roles")
}
