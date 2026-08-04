// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsFailoverConnection Settings that define a failover connection to an external data store.
type StreamsFailoverConnection struct {
	// Unique identifier of the connection.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Human-readable label that identifies the stream connection.
	Name *string `json:"name,omitempty"`
	// The connection region.
	Region *string `json:"region,omitempty"`
	// The connection state.
	// Read only field.
	State *string `json:"state,omitempty"`
	// The connection type.
	Type *string `json:"type,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that contains the configured cluster. Required if the ID does not match the project containing the streams workspace. You must first enable the organization setting.
	ClusterGroupId *string `json:"clusterGroupId,omitempty"`
	// Name of the cluster configured for this connection.
	ClusterName     *string                     `json:"clusterName,omitempty"`
	DbRoleToExecute *DBRoleToExecute            `json:"dbRoleToExecute,omitempty"`
	Authentication  *StreamsKafkaAuthentication `json:"authentication,omitempty"`
	// Comma separated list of server addresses.
	BootstrapServers *string `json:"bootstrapServers,omitempty"`
	// A map of Kafka key-value pairs for optional configuration. This is a flat object, and keys can have '.' characters.
	Config     *map[string]string      `json:"config,omitempty"`
	Networking *StreamsKafkaNetworking `json:"networking,omitempty"`
	Security   *StreamsKafkaSecurity   `json:"security,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsFailoverConnection) MarshalJSON() ([]byte, error) {
	type noMethod StreamsFailoverConnection
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsFailoverConnection instantiates a new StreamsFailoverConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsFailoverConnection() *StreamsFailoverConnection {
	this := StreamsFailoverConnection{}
	return &this
}

// NewStreamsFailoverConnectionWithDefaults instantiates a new StreamsFailoverConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsFailoverConnectionWithDefaults() *StreamsFailoverConnection {
	this := StreamsFailoverConnection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StreamsFailoverConnection) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsFailoverConnection) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetName returns the Name field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StreamsFailoverConnection) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StreamsFailoverConnection) SetRegion(v string) {
	o.Region = &v
	o.NullFields = removeNullField(o.NullFields, "Region")
}

// SetRegionNil sets Region to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetRegionNil() {
	o.Region = nil
	o.NullFields = addNullField(o.NullFields, "Region")
}

// GetState returns the State field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StreamsFailoverConnection) SetState(v string) {
	o.State = &v
	o.NullFields = removeNullField(o.NullFields, "State")
}

// SetStateNil sets State to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetStateNil() {
	o.State = nil
	o.NullFields = addNullField(o.NullFields, "State")
}

// GetType returns the Type field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StreamsFailoverConnection) SetType(v string) {
	o.Type = &v
	o.NullFields = removeNullField(o.NullFields, "Type")
}

// SetTypeNil sets Type to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetTypeNil() {
	o.Type = nil
	o.NullFields = addNullField(o.NullFields, "Type")
}

// GetClusterGroupId returns the ClusterGroupId field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetClusterGroupId() string {
	if o == nil || IsNil(o.ClusterGroupId) {
		var ret string
		return ret
	}
	return *o.ClusterGroupId
}

// GetClusterGroupIdOk returns a tuple with the ClusterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetClusterGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterGroupId) {
		return nil, false
	}

	return o.ClusterGroupId, true
}

// HasClusterGroupId returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasClusterGroupId() bool {
	if o != nil && !IsNil(o.ClusterGroupId) {
		return true
	}

	return false
}

// SetClusterGroupId gets a reference to the given string and assigns it to the ClusterGroupId field.
func (o *StreamsFailoverConnection) SetClusterGroupId(v string) {
	o.ClusterGroupId = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterGroupId")
}

// SetClusterGroupIdNil sets ClusterGroupId to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetClusterGroupIdNil() {
	o.ClusterGroupId = nil
	o.NullFields = addNullField(o.NullFields, "ClusterGroupId")
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *StreamsFailoverConnection) SetClusterName(v string) {
	o.ClusterName = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterName")
}

// SetClusterNameNil sets ClusterName to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetClusterNameNil() {
	o.ClusterName = nil
	o.NullFields = addNullField(o.NullFields, "ClusterName")
}

// GetDbRoleToExecute returns the DbRoleToExecute field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetDbRoleToExecute() DBRoleToExecute {
	if o == nil || IsNil(o.DbRoleToExecute) {
		var ret DBRoleToExecute
		return ret
	}
	return *o.DbRoleToExecute
}

// GetDbRoleToExecuteOk returns a tuple with the DbRoleToExecute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetDbRoleToExecuteOk() (*DBRoleToExecute, bool) {
	if o == nil || IsNil(o.DbRoleToExecute) {
		return nil, false
	}

	return o.DbRoleToExecute, true
}

// HasDbRoleToExecute returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasDbRoleToExecute() bool {
	if o != nil && !IsNil(o.DbRoleToExecute) {
		return true
	}

	return false
}

// SetDbRoleToExecute gets a reference to the given DBRoleToExecute and assigns it to the DbRoleToExecute field.
func (o *StreamsFailoverConnection) SetDbRoleToExecute(v DBRoleToExecute) {
	o.DbRoleToExecute = &v
	o.NullFields = removeNullField(o.NullFields, "DbRoleToExecute")
}

// SetDbRoleToExecuteNil sets DbRoleToExecute to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetDbRoleToExecuteNil() {
	o.DbRoleToExecute = nil
	o.NullFields = addNullField(o.NullFields, "DbRoleToExecute")
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetAuthentication() StreamsKafkaAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret StreamsKafkaAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetAuthenticationOk() (*StreamsKafkaAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}

	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given StreamsKafkaAuthentication and assigns it to the Authentication field.
func (o *StreamsFailoverConnection) SetAuthentication(v StreamsKafkaAuthentication) {
	o.Authentication = &v
	o.NullFields = removeNullField(o.NullFields, "Authentication")
}

// SetAuthenticationNil sets Authentication to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetAuthenticationNil() {
	o.Authentication = nil
	o.NullFields = addNullField(o.NullFields, "Authentication")
}

// GetBootstrapServers returns the BootstrapServers field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetBootstrapServers() string {
	if o == nil || IsNil(o.BootstrapServers) {
		var ret string
		return ret
	}
	return *o.BootstrapServers
}

// GetBootstrapServersOk returns a tuple with the BootstrapServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetBootstrapServersOk() (*string, bool) {
	if o == nil || IsNil(o.BootstrapServers) {
		return nil, false
	}

	return o.BootstrapServers, true
}

// HasBootstrapServers returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasBootstrapServers() bool {
	if o != nil && !IsNil(o.BootstrapServers) {
		return true
	}

	return false
}

// SetBootstrapServers gets a reference to the given string and assigns it to the BootstrapServers field.
func (o *StreamsFailoverConnection) SetBootstrapServers(v string) {
	o.BootstrapServers = &v
	o.NullFields = removeNullField(o.NullFields, "BootstrapServers")
}

// SetBootstrapServersNil sets BootstrapServers to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetBootstrapServersNil() {
	o.BootstrapServers = nil
	o.NullFields = addNullField(o.NullFields, "BootstrapServers")
}

// GetConfig returns the Config field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetConfig() map[string]string {
	if o == nil || IsNil(o.Config) {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetConfigOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}

	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *StreamsFailoverConnection) SetConfig(v map[string]string) {
	o.Config = &v
	o.NullFields = removeNullField(o.NullFields, "Config")
}

// SetConfigNil sets Config to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetConfigNil() {
	o.Config = nil
	o.NullFields = addNullField(o.NullFields, "Config")
}

// GetNetworking returns the Networking field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetNetworking() StreamsKafkaNetworking {
	if o == nil || IsNil(o.Networking) {
		var ret StreamsKafkaNetworking
		return ret
	}
	return *o.Networking
}

// GetNetworkingOk returns a tuple with the Networking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetNetworkingOk() (*StreamsKafkaNetworking, bool) {
	if o == nil || IsNil(o.Networking) {
		return nil, false
	}

	return o.Networking, true
}

// HasNetworking returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasNetworking() bool {
	if o != nil && !IsNil(o.Networking) {
		return true
	}

	return false
}

// SetNetworking gets a reference to the given StreamsKafkaNetworking and assigns it to the Networking field.
func (o *StreamsFailoverConnection) SetNetworking(v StreamsKafkaNetworking) {
	o.Networking = &v
	o.NullFields = removeNullField(o.NullFields, "Networking")
}

// SetNetworkingNil sets Networking to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetNetworkingNil() {
	o.Networking = nil
	o.NullFields = addNullField(o.NullFields, "Networking")
}

// GetSecurity returns the Security field value if set, zero value otherwise
func (o *StreamsFailoverConnection) GetSecurity() StreamsKafkaSecurity {
	if o == nil || IsNil(o.Security) {
		var ret StreamsKafkaSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsFailoverConnection) GetSecurityOk() (*StreamsKafkaSecurity, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}

	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *StreamsFailoverConnection) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given StreamsKafkaSecurity and assigns it to the Security field.
func (o *StreamsFailoverConnection) SetSecurity(v StreamsKafkaSecurity) {
	o.Security = &v
	o.NullFields = removeNullField(o.NullFields, "Security")
}

// SetSecurityNil sets Security to an explicit JSON null when marshaled.
func (o *StreamsFailoverConnection) SetSecurityNil() {
	o.Security = nil
	o.NullFields = addNullField(o.NullFields, "Security")
}
