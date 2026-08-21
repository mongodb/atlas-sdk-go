// Code based on the AtlasAPI V2 OpenAPI file

package admin

// RedactedHeader HTTP header with a redacted value.
type RedactedHeader struct {
	// Header name.
	// Read only field.
	Name string `json:"name,omitempty"`
	// Redacted header value.
	// Read only field.
	Value string `json:"value,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *RedactedHeader) MarshalJSON() ([]byte, error) {
	type noMethod RedactedHeader
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewRedactedHeader instantiates a new RedactedHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedactedHeader(name string, value string) *RedactedHeader {
	this := RedactedHeader{}
	this.Name = name
	this.Value = value
	return &this
}

// NewRedactedHeaderWithDefaults instantiates a new RedactedHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedactedHeaderWithDefaults() *RedactedHeader {
	this := RedactedHeader{}
	return &this
}

// GetName returns the Name field value
func (o *RedactedHeader) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RedactedHeader) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RedactedHeader) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *RedactedHeader) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RedactedHeader) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RedactedHeader) SetValue(v string) {
	o.Value = v
}
