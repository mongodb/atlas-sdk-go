// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsConnection Settings that define a connection to an external data store.
type StreamsConnection struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Human-readable label that identifies the stream connection. In the case of the Sample type, this is the name of the sample source.
	Name *string `json:"name,omitempty"`
	// Type of the connection.
	Type *string `json:"type,omitempty"`
	// The id of the group that the cluster belongs to.
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
	// A map of key-value pairs that will be passed as headers for the request.
	Headers *map[string]string `json:"headers,omitempty"`
	// The url to be used for the request.
	Url *string                     `json:"url,omitempty"`
	Aws *StreamsAWSConnectionConfig `json:"aws,omitempty"`
}

// NewStreamsConnection instantiates a new StreamsConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsConnection() *StreamsConnection {
	this := StreamsConnection{}
	return &this
}

// NewStreamsConnectionWithDefaults instantiates a new StreamsConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsConnectionWithDefaults() *StreamsConnection {
	this := StreamsConnection{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsConnection) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsConnection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsConnection) SetLinks(v []Link) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *StreamsConnection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StreamsConnection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StreamsConnection) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *StreamsConnection) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StreamsConnection) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StreamsConnection) SetType(v string) {
	o.Type = &v
}

// GetClusterGroupId returns the ClusterGroupId field value if set, zero value otherwise
func (o *StreamsConnection) GetClusterGroupId() string {
	if o == nil || IsNil(o.ClusterGroupId) {
		var ret string
		return ret
	}
	return *o.ClusterGroupId
}

// GetClusterGroupIdOk returns a tuple with the ClusterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetClusterGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterGroupId) {
		return nil, false
	}

	return o.ClusterGroupId, true
}

// HasClusterGroupId returns a boolean if a field has been set.
func (o *StreamsConnection) HasClusterGroupId() bool {
	if o != nil && !IsNil(o.ClusterGroupId) {
		return true
	}

	return false
}

// SetClusterGroupId gets a reference to the given string and assigns it to the ClusterGroupId field.
func (o *StreamsConnection) SetClusterGroupId(v string) {
	o.ClusterGroupId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *StreamsConnection) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *StreamsConnection) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *StreamsConnection) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDbRoleToExecute returns the DbRoleToExecute field value if set, zero value otherwise
func (o *StreamsConnection) GetDbRoleToExecute() DBRoleToExecute {
	if o == nil || IsNil(o.DbRoleToExecute) {
		var ret DBRoleToExecute
		return ret
	}
	return *o.DbRoleToExecute
}

// GetDbRoleToExecuteOk returns a tuple with the DbRoleToExecute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetDbRoleToExecuteOk() (*DBRoleToExecute, bool) {
	if o == nil || IsNil(o.DbRoleToExecute) {
		return nil, false
	}

	return o.DbRoleToExecute, true
}

// HasDbRoleToExecute returns a boolean if a field has been set.
func (o *StreamsConnection) HasDbRoleToExecute() bool {
	if o != nil && !IsNil(o.DbRoleToExecute) {
		return true
	}

	return false
}

// SetDbRoleToExecute gets a reference to the given DBRoleToExecute and assigns it to the DbRoleToExecute field.
func (o *StreamsConnection) SetDbRoleToExecute(v DBRoleToExecute) {
	o.DbRoleToExecute = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise
func (o *StreamsConnection) GetAuthentication() StreamsKafkaAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret StreamsKafkaAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetAuthenticationOk() (*StreamsKafkaAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}

	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *StreamsConnection) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given StreamsKafkaAuthentication and assigns it to the Authentication field.
func (o *StreamsConnection) SetAuthentication(v StreamsKafkaAuthentication) {
	o.Authentication = &v
}

// GetBootstrapServers returns the BootstrapServers field value if set, zero value otherwise
func (o *StreamsConnection) GetBootstrapServers() string {
	if o == nil || IsNil(o.BootstrapServers) {
		var ret string
		return ret
	}
	return *o.BootstrapServers
}

// GetBootstrapServersOk returns a tuple with the BootstrapServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetBootstrapServersOk() (*string, bool) {
	if o == nil || IsNil(o.BootstrapServers) {
		return nil, false
	}

	return o.BootstrapServers, true
}

// HasBootstrapServers returns a boolean if a field has been set.
func (o *StreamsConnection) HasBootstrapServers() bool {
	if o != nil && !IsNil(o.BootstrapServers) {
		return true
	}

	return false
}

// SetBootstrapServers gets a reference to the given string and assigns it to the BootstrapServers field.
func (o *StreamsConnection) SetBootstrapServers(v string) {
	o.BootstrapServers = &v
}

// GetConfig returns the Config field value if set, zero value otherwise
func (o *StreamsConnection) GetConfig() map[string]string {
	if o == nil || IsNil(o.Config) {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetConfigOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}

	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *StreamsConnection) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *StreamsConnection) SetConfig(v map[string]string) {
	o.Config = &v
}

// GetNetworking returns the Networking field value if set, zero value otherwise
func (o *StreamsConnection) GetNetworking() StreamsKafkaNetworking {
	if o == nil || IsNil(o.Networking) {
		var ret StreamsKafkaNetworking
		return ret
	}
	return *o.Networking
}

// GetNetworkingOk returns a tuple with the Networking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetNetworkingOk() (*StreamsKafkaNetworking, bool) {
	if o == nil || IsNil(o.Networking) {
		return nil, false
	}

	return o.Networking, true
}

// HasNetworking returns a boolean if a field has been set.
func (o *StreamsConnection) HasNetworking() bool {
	if o != nil && !IsNil(o.Networking) {
		return true
	}

	return false
}

// SetNetworking gets a reference to the given StreamsKafkaNetworking and assigns it to the Networking field.
func (o *StreamsConnection) SetNetworking(v StreamsKafkaNetworking) {
	o.Networking = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise
func (o *StreamsConnection) GetSecurity() StreamsKafkaSecurity {
	if o == nil || IsNil(o.Security) {
		var ret StreamsKafkaSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetSecurityOk() (*StreamsKafkaSecurity, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}

	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *StreamsConnection) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given StreamsKafkaSecurity and assigns it to the Security field.
func (o *StreamsConnection) SetSecurity(v StreamsKafkaSecurity) {
	o.Security = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise
func (o *StreamsConnection) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}

	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *StreamsConnection) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *StreamsConnection) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetUrl returns the Url field value if set, zero value otherwise
func (o *StreamsConnection) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}

	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *StreamsConnection) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *StreamsConnection) SetUrl(v string) {
	o.Url = &v
}

// GetAws returns the Aws field value if set, zero value otherwise
func (o *StreamsConnection) GetAws() StreamsAWSConnectionConfig {
	if o == nil || IsNil(o.Aws) {
		var ret StreamsAWSConnectionConfig
		return ret
	}
	return *o.Aws
}

// GetAwsOk returns a tuple with the Aws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsConnection) GetAwsOk() (*StreamsAWSConnectionConfig, bool) {
	if o == nil || IsNil(o.Aws) {
		return nil, false
	}

	return o.Aws, true
}

// HasAws returns a boolean if a field has been set.
func (o *StreamsConnection) HasAws() bool {
	if o != nil && !IsNil(o.Aws) {
		return true
	}

	return false
}

// SetAws gets a reference to the given StreamsAWSConnectionConfig and assigns it to the Aws field.
func (o *StreamsConnection) SetAws(v StreamsAWSConnectionConfig) {
	o.Aws = &v
}
