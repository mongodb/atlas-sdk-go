// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SkuServiceResponse struct for SkuServiceResponse
type SkuServiceResponse struct {
	// Human-readable name of the SKU service as shown on invoices.
	// Read only field.
	Name *string `json:"name,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *SkuServiceResponse) MarshalJSON() ([]byte, error) {
	type noMethod SkuServiceResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewSkuServiceResponse instantiates a new SkuServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuServiceResponse() *SkuServiceResponse {
	this := SkuServiceResponse{}
	return &this
}

// NewSkuServiceResponseWithDefaults instantiates a new SkuServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuServiceResponseWithDefaults() *SkuServiceResponse {
	this := SkuServiceResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise
func (o *SkuServiceResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuServiceResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SkuServiceResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SkuServiceResponse) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *SkuServiceResponse) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}
