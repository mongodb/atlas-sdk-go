// Code based on the AtlasAPI V2 OpenAPI file

package admin

// CreateOrgMcpConfigRequest struct for CreateOrgMcpConfigRequest
type CreateOrgMcpConfigRequest struct {
	// List of IP access list entries that define allowed source addresses for this MCP configuration.
	IpAccessList *[]ServiceAccountIPAccessListEntry `json:"ipAccessList,omitempty"`
	// Human-readable name that identifies this MCP configuration.
	McpConfigName string `json:"mcpConfigName"`
	// List of organization roles to assign to this MCP configuration.
	Roles []string `json:"roles"`
}

// NewCreateOrgMcpConfigRequest instantiates a new CreateOrgMcpConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrgMcpConfigRequest(mcpConfigName string, roles []string) *CreateOrgMcpConfigRequest {
	this := CreateOrgMcpConfigRequest{}
	this.McpConfigName = mcpConfigName
	this.Roles = roles
	return &this
}

// NewCreateOrgMcpConfigRequestWithDefaults instantiates a new CreateOrgMcpConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrgMcpConfigRequestWithDefaults() *CreateOrgMcpConfigRequest {
	this := CreateOrgMcpConfigRequest{}
	return &this
}

// GetIpAccessList returns the IpAccessList field value if set, zero value otherwise
func (o *CreateOrgMcpConfigRequest) GetIpAccessList() []ServiceAccountIPAccessListEntry {
	if o == nil || IsNil(o.IpAccessList) {
		var ret []ServiceAccountIPAccessListEntry
		return ret
	}
	return *o.IpAccessList
}

// GetIpAccessListOk returns a tuple with the IpAccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrgMcpConfigRequest) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool) {
	if o == nil || IsNil(o.IpAccessList) {
		return nil, false
	}

	return o.IpAccessList, true
}

// HasIpAccessList returns a boolean if a field has been set.
func (o *CreateOrgMcpConfigRequest) HasIpAccessList() bool {
	if o != nil && !IsNil(o.IpAccessList) {
		return true
	}

	return false
}

// SetIpAccessList gets a reference to the given []ServiceAccountIPAccessListEntry and assigns it to the IpAccessList field.
func (o *CreateOrgMcpConfigRequest) SetIpAccessList(v []ServiceAccountIPAccessListEntry) {
	o.IpAccessList = &v
}

// GetMcpConfigName returns the McpConfigName field value
func (o *CreateOrgMcpConfigRequest) GetMcpConfigName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.McpConfigName
}

// GetMcpConfigNameOk returns a tuple with the McpConfigName field value
// and a boolean to check if the value has been set.
func (o *CreateOrgMcpConfigRequest) GetMcpConfigNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.McpConfigName, true
}

// SetMcpConfigName sets field value
func (o *CreateOrgMcpConfigRequest) SetMcpConfigName(v string) {
	o.McpConfigName = v
}

// GetRoles returns the Roles field value
func (o *CreateOrgMcpConfigRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CreateOrgMcpConfigRequest) GetRolesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *CreateOrgMcpConfigRequest) SetRoles(v []string) {
	o.Roles = v
}
