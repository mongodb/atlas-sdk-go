// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasPolicyCreate struct for ApiAtlasPolicyCreate
type ApiAtlasPolicyCreate struct {
	// A string that defines the permissions for the policy. The syntax used is the Cedar Policy language.
	Body string `json:"body"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiAtlasPolicyCreate) MarshalJSON() ([]byte, error) {
	type noMethod ApiAtlasPolicyCreate
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiAtlasPolicyCreate instantiates a new ApiAtlasPolicyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasPolicyCreate(body string) *ApiAtlasPolicyCreate {
	this := ApiAtlasPolicyCreate{}
	this.Body = body
	return &this
}

// NewApiAtlasPolicyCreateWithDefaults instantiates a new ApiAtlasPolicyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasPolicyCreateWithDefaults() *ApiAtlasPolicyCreate {
	this := ApiAtlasPolicyCreate{}
	return &this
}

// GetBody returns the Body field value
func (o *ApiAtlasPolicyCreate) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasPolicyCreate) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *ApiAtlasPolicyCreate) SetBody(v string) {
	o.Body = v
}
