// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgMcpConfigResponse struct for OrgMcpConfigResponse
type OrgMcpConfigResponse struct {
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
	// List of organization roles associated with this MCP configuration.
	Roles *[]string `json:"roles,omitempty"`
}

// NewOrgMcpConfigResponse instantiates a new OrgMcpConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgMcpConfigResponse() *OrgMcpConfigResponse {
	this := OrgMcpConfigResponse{}
	return &this
}

// NewOrgMcpConfigResponseWithDefaults instantiates a new OrgMcpConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgMcpConfigResponseWithDefaults() *OrgMcpConfigResponse {
	this := OrgMcpConfigResponse{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise
func (o *OrgMcpConfigResponse) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigResponse) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}

	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OrgMcpConfigResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OrgMcpConfigResponse) SetClientId(v string) {
	o.ClientId = &v
}

// GetEgressClientId returns the EgressClientId field value if set, zero value otherwise
func (o *OrgMcpConfigResponse) GetEgressClientId() string {
	if o == nil || IsNil(o.EgressClientId) {
		var ret string
		return ret
	}
	return *o.EgressClientId
}

// GetEgressClientIdOk returns a tuple with the EgressClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigResponse) GetEgressClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.EgressClientId) {
		return nil, false
	}

	return o.EgressClientId, true
}

// HasEgressClientId returns a boolean if a field has been set.
func (o *OrgMcpConfigResponse) HasEgressClientId() bool {
	if o != nil && !IsNil(o.EgressClientId) {
		return true
	}

	return false
}

// SetEgressClientId gets a reference to the given string and assigns it to the EgressClientId field.
func (o *OrgMcpConfigResponse) SetEgressClientId(v string) {
	o.EgressClientId = &v
}

// GetIpAccessList returns the IpAccessList field value if set, zero value otherwise
func (o *OrgMcpConfigResponse) GetIpAccessList() []ServiceAccountIPAccessListEntry {
	if o == nil || IsNil(o.IpAccessList) {
		var ret []ServiceAccountIPAccessListEntry
		return ret
	}
	return *o.IpAccessList
}

// GetIpAccessListOk returns a tuple with the IpAccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigResponse) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool) {
	if o == nil || IsNil(o.IpAccessList) {
		return nil, false
	}

	return o.IpAccessList, true
}

// HasIpAccessList returns a boolean if a field has been set.
func (o *OrgMcpConfigResponse) HasIpAccessList() bool {
	if o != nil && !IsNil(o.IpAccessList) {
		return true
	}

	return false
}

// SetIpAccessList gets a reference to the given []ServiceAccountIPAccessListEntry and assigns it to the IpAccessList field.
func (o *OrgMcpConfigResponse) SetIpAccessList(v []ServiceAccountIPAccessListEntry) {
	o.IpAccessList = &v
}

// GetMcpConfigId returns the McpConfigId field value if set, zero value otherwise
func (o *OrgMcpConfigResponse) GetMcpConfigId() string {
	if o == nil || IsNil(o.McpConfigId) {
		var ret string
		return ret
	}
	return *o.McpConfigId
}

// GetMcpConfigIdOk returns a tuple with the McpConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigResponse) GetMcpConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.McpConfigId) {
		return nil, false
	}

	return o.McpConfigId, true
}

// HasMcpConfigId returns a boolean if a field has been set.
func (o *OrgMcpConfigResponse) HasMcpConfigId() bool {
	if o != nil && !IsNil(o.McpConfigId) {
		return true
	}

	return false
}

// SetMcpConfigId gets a reference to the given string and assigns it to the McpConfigId field.
func (o *OrgMcpConfigResponse) SetMcpConfigId(v string) {
	o.McpConfigId = &v
}

// GetMcpConfigName returns the McpConfigName field value if set, zero value otherwise
func (o *OrgMcpConfigResponse) GetMcpConfigName() string {
	if o == nil || IsNil(o.McpConfigName) {
		var ret string
		return ret
	}
	return *o.McpConfigName
}

// GetMcpConfigNameOk returns a tuple with the McpConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigResponse) GetMcpConfigNameOk() (*string, bool) {
	if o == nil || IsNil(o.McpConfigName) {
		return nil, false
	}

	return o.McpConfigName, true
}

// HasMcpConfigName returns a boolean if a field has been set.
func (o *OrgMcpConfigResponse) HasMcpConfigName() bool {
	if o != nil && !IsNil(o.McpConfigName) {
		return true
	}

	return false
}

// SetMcpConfigName gets a reference to the given string and assigns it to the McpConfigName field.
func (o *OrgMcpConfigResponse) SetMcpConfigName(v string) {
	o.McpConfigName = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *OrgMcpConfigResponse) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMcpConfigResponse) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrgMcpConfigResponse) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrgMcpConfigResponse) SetRoles(v []string) {
	o.Roles = &v
}
