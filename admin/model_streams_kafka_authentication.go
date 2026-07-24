// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsKafkaAuthentication User credentials required to connect to a Kafka Cluster. Includes the authentication type, as well as the parameters for that authentication mode.
type StreamsKafkaAuthentication struct {
	// OIDC client identifier for authentication to the Kafka cluster.
	ClientId *string `json:"clientId,omitempty"`
	// OIDC client secret for authentication to the Kafka cluster.
	// Write only field.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Style of authentication. Can be one of PLAIN, SCRAM-256, SCRAM-512, or OAUTHBEARER.
	Mechanism *string `json:"mechanism,omitempty"`
	// SASL OAUTHBEARER authentication method. Currently, only OIDC is supported.
	Method *string `json:"method,omitempty"`
	// Password of the account to connect to the Kafka cluster.
	// Write only field.
	Password *string `json:"password,omitempty"`
	// SASL OAUTHBEARER extensions parameter for additional OAuth2 configuration.
	SaslOauthbearerExtensions *string `json:"saslOauthbearerExtensions,omitempty"`
	// OIDC scope parameter defining the access permissions requested.
	Scope *string `json:"scope,omitempty"`
	// SSL certificate for client authentication to Kafka.
	SslCertificate *string `json:"sslCertificate,omitempty"`
	// SSL key for client authentication to Kafka.
	// Write only field.
	SslKey *string `json:"sslKey,omitempty"`
	// Password for the SSL key, if it is password protected.
	// Write only field.
	SslKeyPassword *string `json:"sslKeyPassword,omitempty"`
	// OIDC token endpoint URL for obtaining access tokens.
	TokenEndpointUrl *string `json:"tokenEndpointUrl,omitempty"`
	// Username of the account to connect to the Kafka cluster.
	Username *string `json:"username,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsKafkaAuthentication) MarshalJSON() ([]byte, error) {
	type noMethod StreamsKafkaAuthentication
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsKafkaAuthentication instantiates a new StreamsKafkaAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsKafkaAuthentication() *StreamsKafkaAuthentication {
	this := StreamsKafkaAuthentication{}
	return &this
}

// NewStreamsKafkaAuthenticationWithDefaults instantiates a new StreamsKafkaAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsKafkaAuthenticationWithDefaults() *StreamsKafkaAuthentication {
	this := StreamsKafkaAuthentication{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}

	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *StreamsKafkaAuthentication) SetClientId(v string) {
	o.ClientId = &v
	o.NullFields = removeNullField(o.NullFields, "ClientId")
}

// SetClientIdNil sets ClientId to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetClientIdNil() {
	o.ClientId = nil
	o.NullFields = addNullField(o.NullFields, "ClientId")
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}

	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *StreamsKafkaAuthentication) SetClientSecret(v string) {
	o.ClientSecret = &v
	o.NullFields = removeNullField(o.NullFields, "ClientSecret")
}

// SetClientSecretNil sets ClientSecret to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetClientSecretNil() {
	o.ClientSecret = nil
	o.NullFields = addNullField(o.NullFields, "ClientSecret")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsKafkaAuthentication) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetMechanism returns the Mechanism field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetMechanism() string {
	if o == nil || IsNil(o.Mechanism) {
		var ret string
		return ret
	}
	return *o.Mechanism
}

// GetMechanismOk returns a tuple with the Mechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetMechanismOk() (*string, bool) {
	if o == nil || IsNil(o.Mechanism) {
		return nil, false
	}

	return o.Mechanism, true
}

// HasMechanism returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasMechanism() bool {
	if o != nil && !IsNil(o.Mechanism) {
		return true
	}

	return false
}

// SetMechanism gets a reference to the given string and assigns it to the Mechanism field.
func (o *StreamsKafkaAuthentication) SetMechanism(v string) {
	o.Mechanism = &v
	o.NullFields = removeNullField(o.NullFields, "Mechanism")
}

// SetMechanismNil sets Mechanism to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetMechanismNil() {
	o.Mechanism = nil
	o.NullFields = addNullField(o.NullFields, "Mechanism")
}

// GetMethod returns the Method field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}

	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *StreamsKafkaAuthentication) SetMethod(v string) {
	o.Method = &v
	o.NullFields = removeNullField(o.NullFields, "Method")
}

// SetMethodNil sets Method to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetMethodNil() {
	o.Method = nil
	o.NullFields = addNullField(o.NullFields, "Method")
}

// GetPassword returns the Password field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}

	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *StreamsKafkaAuthentication) SetPassword(v string) {
	o.Password = &v
	o.NullFields = removeNullField(o.NullFields, "Password")
}

// SetPasswordNil sets Password to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetPasswordNil() {
	o.Password = nil
	o.NullFields = addNullField(o.NullFields, "Password")
}

// GetSaslOauthbearerExtensions returns the SaslOauthbearerExtensions field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetSaslOauthbearerExtensions() string {
	if o == nil || IsNil(o.SaslOauthbearerExtensions) {
		var ret string
		return ret
	}
	return *o.SaslOauthbearerExtensions
}

// GetSaslOauthbearerExtensionsOk returns a tuple with the SaslOauthbearerExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetSaslOauthbearerExtensionsOk() (*string, bool) {
	if o == nil || IsNil(o.SaslOauthbearerExtensions) {
		return nil, false
	}

	return o.SaslOauthbearerExtensions, true
}

// HasSaslOauthbearerExtensions returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasSaslOauthbearerExtensions() bool {
	if o != nil && !IsNil(o.SaslOauthbearerExtensions) {
		return true
	}

	return false
}

// SetSaslOauthbearerExtensions gets a reference to the given string and assigns it to the SaslOauthbearerExtensions field.
func (o *StreamsKafkaAuthentication) SetSaslOauthbearerExtensions(v string) {
	o.SaslOauthbearerExtensions = &v
	o.NullFields = removeNullField(o.NullFields, "SaslOauthbearerExtensions")
}

// SetSaslOauthbearerExtensionsNil sets SaslOauthbearerExtensions to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetSaslOauthbearerExtensionsNil() {
	o.SaslOauthbearerExtensions = nil
	o.NullFields = addNullField(o.NullFields, "SaslOauthbearerExtensions")
}

// GetScope returns the Scope field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}

	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *StreamsKafkaAuthentication) SetScope(v string) {
	o.Scope = &v
	o.NullFields = removeNullField(o.NullFields, "Scope")
}

// SetScopeNil sets Scope to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetScopeNil() {
	o.Scope = nil
	o.NullFields = addNullField(o.NullFields, "Scope")
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetSslCertificate() string {
	if o == nil || IsNil(o.SslCertificate) {
		var ret string
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetSslCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.SslCertificate) {
		return nil, false
	}

	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasSslCertificate() bool {
	if o != nil && !IsNil(o.SslCertificate) {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given string and assigns it to the SslCertificate field.
func (o *StreamsKafkaAuthentication) SetSslCertificate(v string) {
	o.SslCertificate = &v
	o.NullFields = removeNullField(o.NullFields, "SslCertificate")
}

// SetSslCertificateNil sets SslCertificate to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetSslCertificateNil() {
	o.SslCertificate = nil
	o.NullFields = addNullField(o.NullFields, "SslCertificate")
}

// GetSslKey returns the SslKey field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetSslKey() string {
	if o == nil || IsNil(o.SslKey) {
		var ret string
		return ret
	}
	return *o.SslKey
}

// GetSslKeyOk returns a tuple with the SslKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetSslKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SslKey) {
		return nil, false
	}

	return o.SslKey, true
}

// HasSslKey returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasSslKey() bool {
	if o != nil && !IsNil(o.SslKey) {
		return true
	}

	return false
}

// SetSslKey gets a reference to the given string and assigns it to the SslKey field.
func (o *StreamsKafkaAuthentication) SetSslKey(v string) {
	o.SslKey = &v
	o.NullFields = removeNullField(o.NullFields, "SslKey")
}

// SetSslKeyNil sets SslKey to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetSslKeyNil() {
	o.SslKey = nil
	o.NullFields = addNullField(o.NullFields, "SslKey")
}

// GetSslKeyPassword returns the SslKeyPassword field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetSslKeyPassword() string {
	if o == nil || IsNil(o.SslKeyPassword) {
		var ret string
		return ret
	}
	return *o.SslKeyPassword
}

// GetSslKeyPasswordOk returns a tuple with the SslKeyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetSslKeyPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SslKeyPassword) {
		return nil, false
	}

	return o.SslKeyPassword, true
}

// HasSslKeyPassword returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasSslKeyPassword() bool {
	if o != nil && !IsNil(o.SslKeyPassword) {
		return true
	}

	return false
}

// SetSslKeyPassword gets a reference to the given string and assigns it to the SslKeyPassword field.
func (o *StreamsKafkaAuthentication) SetSslKeyPassword(v string) {
	o.SslKeyPassword = &v
	o.NullFields = removeNullField(o.NullFields, "SslKeyPassword")
}

// SetSslKeyPasswordNil sets SslKeyPassword to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetSslKeyPasswordNil() {
	o.SslKeyPassword = nil
	o.NullFields = addNullField(o.NullFields, "SslKeyPassword")
}

// GetTokenEndpointUrl returns the TokenEndpointUrl field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetTokenEndpointUrl() string {
	if o == nil || IsNil(o.TokenEndpointUrl) {
		var ret string
		return ret
	}
	return *o.TokenEndpointUrl
}

// GetTokenEndpointUrlOk returns a tuple with the TokenEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetTokenEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpointUrl) {
		return nil, false
	}

	return o.TokenEndpointUrl, true
}

// HasTokenEndpointUrl returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasTokenEndpointUrl() bool {
	if o != nil && !IsNil(o.TokenEndpointUrl) {
		return true
	}

	return false
}

// SetTokenEndpointUrl gets a reference to the given string and assigns it to the TokenEndpointUrl field.
func (o *StreamsKafkaAuthentication) SetTokenEndpointUrl(v string) {
	o.TokenEndpointUrl = &v
	o.NullFields = removeNullField(o.NullFields, "TokenEndpointUrl")
}

// SetTokenEndpointUrlNil sets TokenEndpointUrl to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetTokenEndpointUrlNil() {
	o.TokenEndpointUrl = nil
	o.NullFields = addNullField(o.NullFields, "TokenEndpointUrl")
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *StreamsKafkaAuthentication) SetUsername(v string) {
	o.Username = &v
	o.NullFields = removeNullField(o.NullFields, "Username")
}

// SetUsernameNil sets Username to an explicit JSON null when marshaled.
func (o *StreamsKafkaAuthentication) SetUsernameNil() {
	o.Username = nil
	o.NullFields = addNullField(o.NullFields, "Username")
}
