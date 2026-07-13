// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgMcpConfigUpdateRequest struct for OrgMcpConfigUpdateRequest
type OrgMcpConfigUpdateRequest struct {
	// List of IP access list entries that define allowed source addresses for this MCP configuration. If provided, replaces the existing IP access list.
	IpAccessList *[]ServiceAccountIPAccessListEntry `json:"ipAccessList,omitempty"`
	// Updated human-readable name for this MCP configuration.
	McpConfigName *string `json:"mcpConfigName,omitempty"`
	// List of organization roles associated with this MCP configuration. If provided, replaces the existing list of roles.
	Roles *[]string `json:"roles,omitempty"`
}

// NewOrgMcpConfigUpdateRequest instantiates a new OrgMcpConfigUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgMcpConfigUpdateRequest() *OrgMcpConfigUpdateRequest {
	this := OrgMcpConfigUpdateRequest{}
	return &this
}

// NewOrgMcpConfigUpdateRequestWithDefaults instantiates a new OrgMcpConfigUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgMcpConfigUpdateRequestWithDefaults() *OrgMcpConfigUpdateRequest {
	this := OrgMcpConfigUpdateRequest{}
	return &this
}

// GetIpAccessList returns the IpAccessList field value if set, zero value otherwise
func (o *OrgMcpConfigUpdateRequest) GetIpAccessList() []ServiceAccountIPAccessListEntry {
	if o == nil || IsNil(o.IpAccessList) {
		var ret []ServiceAccountIPAccessListEntry
		return ret
	}
	return *o.IpAccessList
}

// GetIpAccessListOk returns a tuple with the IpAccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigUpdateRequest) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool) {
	if o == nil || IsNil(o.IpAccessList) {
		return nil, false
	}

	return o.IpAccessList, true
}

// HasIpAccessList returns a boolean if a field has been set.
func (o *OrgMcpConfigUpdateRequest) HasIpAccessList() bool {
	if o != nil && !IsNil(o.IpAccessList) {
		return true
	}

	return false
}

// SetIpAccessList gets a reference to the given []ServiceAccountIPAccessListEntry and assigns it to the IpAccessList field.
func (o *OrgMcpConfigUpdateRequest) SetIpAccessList(v []ServiceAccountIPAccessListEntry) {
	o.IpAccessList = &v
}

// GetMcpConfigName returns the McpConfigName field value if set, zero value otherwise
func (o *OrgMcpConfigUpdateRequest) GetMcpConfigName() string {
	if o == nil || IsNil(o.McpConfigName) {
		var ret string
		return ret
	}
	return *o.McpConfigName
}

// GetMcpConfigNameOk returns a tuple with the McpConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigUpdateRequest) GetMcpConfigNameOk() (*string, bool) {
	if o == nil || IsNil(o.McpConfigName) {
		return nil, false
	}

	return o.McpConfigName, true
}

// HasMcpConfigName returns a boolean if a field has been set.
func (o *OrgMcpConfigUpdateRequest) HasMcpConfigName() bool {
	if o != nil && !IsNil(o.McpConfigName) {
		return true
	}

	return false
}

// SetMcpConfigName gets a reference to the given string and assigns it to the McpConfigName field.
func (o *OrgMcpConfigUpdateRequest) SetMcpConfigName(v string) {
	o.McpConfigName = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *OrgMcpConfigUpdateRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigUpdateRequest) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrgMcpConfigUpdateRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrgMcpConfigUpdateRequest) SetRoles(v []string) {
	o.Roles = &v
}
