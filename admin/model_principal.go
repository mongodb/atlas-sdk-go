// Code based on the AtlasAPI V2 OpenAPI file

package admin

// Principal Information about a principal, such as an OAuth application, that triggered the event through delegated access.
type Principal struct {
	// The identifier of this principal.
	Id *string `json:"id,omitempty"`
	// The human-readable name of this principal.
	Name       *string    `json:"name,omitempty"`
	OnBehalfOf *Principal `json:"onBehalfOf,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *Principal) MarshalJSON() ([]byte, error) {
	type noMethod Principal
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPrincipal instantiates a new Principal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipal() *Principal {
	this := Principal{}
	return &this
}

// NewPrincipalWithDefaults instantiates a new Principal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalWithDefaults() *Principal {
	this := Principal{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *Principal) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Principal) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Principal) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *Principal) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetName returns the Name field value if set, zero value otherwise
func (o *Principal) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Principal) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Principal) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *Principal) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}

// GetOnBehalfOf returns the OnBehalfOf field value if set, zero value otherwise
func (o *Principal) GetOnBehalfOf() Principal {
	if o == nil || IsNil(o.OnBehalfOf) {
		var ret Principal
		return ret
	}
	return *o.OnBehalfOf
}

// GetOnBehalfOfOk returns a tuple with the OnBehalfOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetOnBehalfOfOk() (*Principal, bool) {
	if o == nil || IsNil(o.OnBehalfOf) {
		return nil, false
	}

	return o.OnBehalfOf, true
}

// HasOnBehalfOf returns a boolean if a field has been set.
func (o *Principal) HasOnBehalfOf() bool {
	if o != nil && !IsNil(o.OnBehalfOf) {
		return true
	}

	return false
}

// SetOnBehalfOf gets a reference to the given Principal and assigns it to the OnBehalfOf field.
func (o *Principal) SetOnBehalfOf(v Principal) {
	o.OnBehalfOf = &v
	o.NullFields = removeNullField(o.NullFields, "OnBehalfOf")
}

// SetOnBehalfOfNil sets OnBehalfOf to an explicit JSON null when marshaled.
func (o *Principal) SetOnBehalfOfNil() {
	o.OnBehalfOf = nil
	o.NullFields = addNullField(o.NullFields, "OnBehalfOf")
}
