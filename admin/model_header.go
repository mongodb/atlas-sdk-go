// Code based on the AtlasAPI V2 OpenAPI file

package admin

// Header HTTP header with name and value.
type Header struct {
	// Header name.
	Name string `json:"name"`
	// Header value.
	Value string `json:"value"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *Header) MarshalJSON() ([]byte, error) {
	type noMethod Header
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewHeader instantiates a new Header object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeader(name string, value string) *Header {
	this := Header{}
	this.Name = name
	this.Value = value
	return &this
}

// NewHeaderWithDefaults instantiates a new Header object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderWithDefaults() *Header {
	this := Header{}
	return &this
}

// GetName returns the Name field value
func (o *Header) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Header) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Header) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *Header) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Header) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Header) SetValue(v string) {
	o.Value = v
}
