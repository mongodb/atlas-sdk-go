// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsSnowflakeConnectionConfig Snowflake settings for a connection.
type StreamsSnowflakeConnectionConfig struct {
	Authentication StreamsSnowflakeAuthentication `json:"authentication"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsSnowflakeConnectionConfig) MarshalJSON() ([]byte, error) {
	type noMethod StreamsSnowflakeConnectionConfig
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsSnowflakeConnectionConfig instantiates a new StreamsSnowflakeConnectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsSnowflakeConnectionConfig(authentication StreamsSnowflakeAuthentication) *StreamsSnowflakeConnectionConfig {
	this := StreamsSnowflakeConnectionConfig{}
	this.Authentication = authentication
	return &this
}

// NewStreamsSnowflakeConnectionConfigWithDefaults instantiates a new StreamsSnowflakeConnectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsSnowflakeConnectionConfigWithDefaults() *StreamsSnowflakeConnectionConfig {
	this := StreamsSnowflakeConnectionConfig{}
	return &this
}

// GetAuthentication returns the Authentication field value
func (o *StreamsSnowflakeConnectionConfig) GetAuthentication() StreamsSnowflakeAuthentication {
	if o == nil {
		var ret StreamsSnowflakeAuthentication
		return ret
	}

	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *StreamsSnowflakeConnectionConfig) GetAuthenticationOk() (*StreamsSnowflakeAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value
func (o *StreamsSnowflakeConnectionConfig) SetAuthentication(v StreamsSnowflakeAuthentication) {
	o.Authentication = v
}
