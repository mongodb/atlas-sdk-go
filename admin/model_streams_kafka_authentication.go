// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsKafkaAuthentication User credentials required to connect to a Kafka Cluster. Includes the authentication type, as well as the parameters for that authentication mode.
type StreamsKafkaAuthentication struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Style of authentication. Can be one of PLAIN, SCRAM-256, or SCRAM-512.
	Mechanism *string `json:"mechanism,omitempty"`
	// Password of the account to connect to the Kafka cluster.
	// Write only field.
	Password *string `json:"password,omitempty"`
	// SSL certificate for client authentication to Kafka.
	SslCertificate *string `json:"sslCertificate,omitempty"`
	// SSL key for client authentication to Kafka.
	// Write only field.
	SslKey *string `json:"sslKey,omitempty"`
	// Password for the SSL key, if it is password protected.
	// Write only field.
	SslKeyPassword *string `json:"sslKeyPassword,omitempty"`
	// Username of the account to connect to the Kafka cluster.
	Username *string `json:"username,omitempty"`
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
}
