// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsDatabricksConnectionConfig Databricks settings for a connection.
type StreamsDatabricksConnectionConfig struct {
	Authentication StreamsDatabricksAuthentication `json:"authentication"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsDatabricksConnectionConfig) MarshalJSON() ([]byte, error) {
	type noMethod StreamsDatabricksConnectionConfig
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsDatabricksConnectionConfig instantiates a new StreamsDatabricksConnectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsDatabricksConnectionConfig(authentication StreamsDatabricksAuthentication) *StreamsDatabricksConnectionConfig {
	this := StreamsDatabricksConnectionConfig{}
	this.Authentication = authentication
	return &this
}

// NewStreamsDatabricksConnectionConfigWithDefaults instantiates a new StreamsDatabricksConnectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsDatabricksConnectionConfigWithDefaults() *StreamsDatabricksConnectionConfig {
	this := StreamsDatabricksConnectionConfig{}
	return &this
}

// GetAuthentication returns the Authentication field value
func (o *StreamsDatabricksConnectionConfig) GetAuthentication() StreamsDatabricksAuthentication {
	if o == nil {
		var ret StreamsDatabricksAuthentication
		return ret
	}

	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *StreamsDatabricksConnectionConfig) GetAuthenticationOk() (*StreamsDatabricksAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value
func (o *StreamsDatabricksConnectionConfig) SetAuthentication(v StreamsDatabricksAuthentication) {
	o.Authentication = v
}
