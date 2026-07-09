// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SkuResponse struct for SkuResponse
type SkuResponse struct {
	// Human-readable short summary of what this SKU represents.
	// Read only field.
	Description *string `json:"description,omitempty"`
	// Unique string that identifies the SKU.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *SkuResponse) MarshalJSON() ([]byte, error) {
	type noMethod SkuResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewSkuResponse instantiates a new SkuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuResponse() *SkuResponse {
	this := SkuResponse{}
	return &this
}

// NewSkuResponseWithDefaults instantiates a new SkuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuResponseWithDefaults() *SkuResponse {
	this := SkuResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *SkuResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SkuResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SkuResponse) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionNil sets Description to an explicit JSON null when marshaled.
func (o *SkuResponse) SetDescriptionNil() {
	o.Description = nil
	o.NullFields = append(o.NullFields, "Description")
}

// GetId returns the Id field value if set, zero value otherwise
func (o *SkuResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SkuResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SkuResponse) SetId(v string) {
	o.Id = &v
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *SkuResponse) SetIdNil() {
	o.Id = nil
	o.NullFields = append(o.NullFields, "Id")
}
