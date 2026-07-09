// Code based on the AtlasAPI V2 OpenAPI file

package admin

// FieldTransformation Field Transformations during ingestion of a Data Lake Pipeline.
type FieldTransformation struct {
	// Key in the document.
	Field *string `json:"field,omitempty"`
	// Type of transformation applied during the export of the namespace in a Data Lake Pipeline.
	Type *string `json:"type,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *FieldTransformation) MarshalJSON() ([]byte, error) {
	type noMethod FieldTransformation
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewFieldTransformation instantiates a new FieldTransformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldTransformation() *FieldTransformation {
	this := FieldTransformation{}
	return &this
}

// NewFieldTransformationWithDefaults instantiates a new FieldTransformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldTransformationWithDefaults() *FieldTransformation {
	this := FieldTransformation{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise
func (o *FieldTransformation) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldTransformation) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}

	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *FieldTransformation) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *FieldTransformation) SetField(v string) {
	o.Field = &v
}

// SetFieldNil sets Field to an explicit JSON null when marshaled.
func (o *FieldTransformation) SetFieldNil() {
	o.Field = nil
	o.NullFields = append(o.NullFields, "Field")
}

// GetType returns the Type field value if set, zero value otherwise
func (o *FieldTransformation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldTransformation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FieldTransformation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FieldTransformation) SetType(v string) {
	o.Type = &v
}

// SetTypeNil sets Type to an explicit JSON null when marshaled.
func (o *FieldTransformation) SetTypeNil() {
	o.Type = nil
	o.NullFields = append(o.NullFields, "Type")
}
