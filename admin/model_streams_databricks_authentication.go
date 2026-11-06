// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsDatabricksAuthentication Databricks service principal credentials used to obtain an OAuth token.
type StreamsDatabricksAuthentication struct {
	// OAuth client ID of the Databricks service principal.
	ClientId string `json:"clientId"`
	// OAuth client secret of the Databricks service principal. Never returned in responses; omit or leave blank on update to keep the stored secret.
	// Write only field.
	ClientSecret string `json:"clientSecret"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsDatabricksAuthentication) MarshalJSON() ([]byte, error) {
	type noMethod StreamsDatabricksAuthentication
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsDatabricksAuthentication instantiates a new StreamsDatabricksAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsDatabricksAuthentication(clientId string, clientSecret string) *StreamsDatabricksAuthentication {
	this := StreamsDatabricksAuthentication{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewStreamsDatabricksAuthenticationWithDefaults instantiates a new StreamsDatabricksAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsDatabricksAuthenticationWithDefaults() *StreamsDatabricksAuthentication {
	this := StreamsDatabricksAuthentication{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *StreamsDatabricksAuthentication) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *StreamsDatabricksAuthentication) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *StreamsDatabricksAuthentication) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *StreamsDatabricksAuthentication) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *StreamsDatabricksAuthentication) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *StreamsDatabricksAuthentication) SetClientSecret(v string) {
	o.ClientSecret = v
}
