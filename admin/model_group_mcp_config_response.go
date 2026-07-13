// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GroupMcpConfigResponse struct for GroupMcpConfigResponse
type GroupMcpConfigResponse struct {
	// Unique identifier for the Service Account client associated with this MCP configuration. Use this Service Account to connect to the Atlas Remote MCP.
	// Read only field.
	ClientId *string `json:"clientId,omitempty"`
	// Unique identifier for the egress Service Account client associated with this MCP configuration. This Service Account is managed by MongoDB Atlas.
	// Read only field.
	EgressClientId *string `json:"egressClientId,omitempty"`
	// List of IP access list entries that define allowed source addresses for this MCP configuration.
	IpAccessList *[]ServiceAccountIPAccessListEntry `json:"ipAccessList,omitempty"`
	// Unique identifier that identifies this MCP configuration.
	// Read only field.
	McpConfigId *string `json:"mcpConfigId,omitempty"`
	// Human-readable name that identifies this MCP configuration.
	McpConfigName *string `json:"mcpConfigName,omitempty"`
	// List of project roles associated with this MCP configuration.
	Roles *[]string `json:"roles,omitempty"`
}

// NewGroupMcpConfigResponse instantiates a new GroupMcpConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMcpConfigResponse() *GroupMcpConfigResponse {
	this := GroupMcpConfigResponse{}
	return &this
}

// NewGroupMcpConfigResponseWithDefaults instantiates a new GroupMcpConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMcpConfigResponseWithDefaults() *GroupMcpConfigResponse {
	this := GroupMcpConfigResponse{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise
func (o *GroupMcpConfigResponse) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigResponse) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}

	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GroupMcpConfigResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GroupMcpConfigResponse) SetClientId(v string) {
	o.ClientId = &v
}

// GetEgressClientId returns the EgressClientId field value if set, zero value otherwise
func (o *GroupMcpConfigResponse) GetEgressClientId() string {
	if o == nil || IsNil(o.EgressClientId) {
		var ret string
		return ret
	}
	return *o.EgressClientId
}

// GetEgressClientIdOk returns a tuple with the EgressClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigResponse) GetEgressClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.EgressClientId) {
		return nil, false
	}

	return o.EgressClientId, true
}

// HasEgressClientId returns a boolean if a field has been set.
func (o *GroupMcpConfigResponse) HasEgressClientId() bool {
	if o != nil && !IsNil(o.EgressClientId) {
		return true
	}

	return false
}

// SetEgressClientId gets a reference to the given string and assigns it to the EgressClientId field.
func (o *GroupMcpConfigResponse) SetEgressClientId(v string) {
	o.EgressClientId = &v
}

// GetIpAccessList returns the IpAccessList field value if set, zero value otherwise
func (o *GroupMcpConfigResponse) GetIpAccessList() []ServiceAccountIPAccessListEntry {
	if o == nil || IsNil(o.IpAccessList) {
		var ret []ServiceAccountIPAccessListEntry
		return ret
	}
	return *o.IpAccessList
}

// GetIpAccessListOk returns a tuple with the IpAccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigResponse) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool) {
	if o == nil || IsNil(o.IpAccessList) {
		return nil, false
	}

	return o.IpAccessList, true
}

// HasIpAccessList returns a boolean if a field has been set.
func (o *GroupMcpConfigResponse) HasIpAccessList() bool {
	if o != nil && !IsNil(o.IpAccessList) {
		return true
	}

	return false
}

// SetIpAccessList gets a reference to the given []ServiceAccountIPAccessListEntry and assigns it to the IpAccessList field.
func (o *GroupMcpConfigResponse) SetIpAccessList(v []ServiceAccountIPAccessListEntry) {
	o.IpAccessList = &v
}

// GetMcpConfigId returns the McpConfigId field value if set, zero value otherwise
func (o *GroupMcpConfigResponse) GetMcpConfigId() string {
	if o == nil || IsNil(o.McpConfigId) {
		var ret string
		return ret
	}
	return *o.McpConfigId
}

// GetMcpConfigIdOk returns a tuple with the McpConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigResponse) GetMcpConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.McpConfigId) {
		return nil, false
	}

	return o.McpConfigId, true
}

// HasMcpConfigId returns a boolean if a field has been set.
func (o *GroupMcpConfigResponse) HasMcpConfigId() bool {
	if o != nil && !IsNil(o.McpConfigId) {
		return true
	}

	return false
}

// SetMcpConfigId gets a reference to the given string and assigns it to the McpConfigId field.
func (o *GroupMcpConfigResponse) SetMcpConfigId(v string) {
	o.McpConfigId = &v
}

// GetMcpConfigName returns the McpConfigName field value if set, zero value otherwise
func (o *GroupMcpConfigResponse) GetMcpConfigName() string {
	if o == nil || IsNil(o.McpConfigName) {
		var ret string
		return ret
	}
	return *o.McpConfigName
}

// GetMcpConfigNameOk returns a tuple with the McpConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigResponse) GetMcpConfigNameOk() (*string, bool) {
	if o == nil || IsNil(o.McpConfigName) {
		return nil, false
	}

	return o.McpConfigName, true
}

// HasMcpConfigName returns a boolean if a field has been set.
func (o *GroupMcpConfigResponse) HasMcpConfigName() bool {
	if o != nil && !IsNil(o.McpConfigName) {
		return true
	}

	return false
}

// SetMcpConfigName gets a reference to the given string and assigns it to the McpConfigName field.
func (o *GroupMcpConfigResponse) SetMcpConfigName(v string) {
	o.McpConfigName = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *GroupMcpConfigResponse) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMcpConfigResponse) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GroupMcpConfigResponse) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GroupMcpConfigResponse) SetRoles(v []string) {
	o.Roles = &v
}
