// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsSnowflakeAuthentication Snowflake service user and key pair used to authenticate to Snowflake.
type StreamsSnowflakeAuthentication struct {
	// Snowflake account identifier.
	Account string `json:"account"`
	// Private key of the service user's key pair. Never returned in responses; omit or leave blank on update to keep the stored key.
	// Write only field.
	PrivateKey string `json:"privateKey"`
	// Snowflake role the service user assumes.
	Role string `json:"role"`
	// Snowflake service user.
	User string `json:"user"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsSnowflakeAuthentication) MarshalJSON() ([]byte, error) {
	type noMethod StreamsSnowflakeAuthentication
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsSnowflakeAuthentication instantiates a new StreamsSnowflakeAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsSnowflakeAuthentication(account string, privateKey string, role string, user string) *StreamsSnowflakeAuthentication {
	this := StreamsSnowflakeAuthentication{}
	this.Account = account
	this.PrivateKey = privateKey
	this.Role = role
	this.User = user
	return &this
}

// NewStreamsSnowflakeAuthenticationWithDefaults instantiates a new StreamsSnowflakeAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsSnowflakeAuthenticationWithDefaults() *StreamsSnowflakeAuthentication {
	this := StreamsSnowflakeAuthentication{}
	return &this
}

// GetAccount returns the Account field value
func (o *StreamsSnowflakeAuthentication) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *StreamsSnowflakeAuthentication) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *StreamsSnowflakeAuthentication) SetAccount(v string) {
	o.Account = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *StreamsSnowflakeAuthentication) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *StreamsSnowflakeAuthentication) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *StreamsSnowflakeAuthentication) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetRole returns the Role field value
func (o *StreamsSnowflakeAuthentication) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *StreamsSnowflakeAuthentication) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *StreamsSnowflakeAuthentication) SetRole(v string) {
	o.Role = v
}

// GetUser returns the User field value
func (o *StreamsSnowflakeAuthentication) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *StreamsSnowflakeAuthentication) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *StreamsSnowflakeAuthentication) SetUser(v string) {
	o.User = v
}
