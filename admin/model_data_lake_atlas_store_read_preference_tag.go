// Code based on the AtlasAPI V2 OpenAPI file

package admin

// DataLakeAtlasStoreReadPreferenceTag struct for DataLakeAtlasStoreReadPreferenceTag
type DataLakeAtlasStoreReadPreferenceTag struct {
	// Human-readable label of the tag.
	Name *string `json:"name,omitempty"`
	// Value of the tag.
	Value *string `json:"value,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *DataLakeAtlasStoreReadPreferenceTag) MarshalJSON() ([]byte, error) {
	type noMethod DataLakeAtlasStoreReadPreferenceTag
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewDataLakeAtlasStoreReadPreferenceTag instantiates a new DataLakeAtlasStoreReadPreferenceTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeAtlasStoreReadPreferenceTag() *DataLakeAtlasStoreReadPreferenceTag {
	this := DataLakeAtlasStoreReadPreferenceTag{}
	return &this
}

// NewDataLakeAtlasStoreReadPreferenceTagWithDefaults instantiates a new DataLakeAtlasStoreReadPreferenceTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeAtlasStoreReadPreferenceTagWithDefaults() *DataLakeAtlasStoreReadPreferenceTag {
	this := DataLakeAtlasStoreReadPreferenceTag{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise
func (o *DataLakeAtlasStoreReadPreferenceTag) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreReadPreferenceTag) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreReadPreferenceTag) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataLakeAtlasStoreReadPreferenceTag) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *DataLakeAtlasStoreReadPreferenceTag) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *DataLakeAtlasStoreReadPreferenceTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreReadPreferenceTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreReadPreferenceTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DataLakeAtlasStoreReadPreferenceTag) SetValue(v string) {
	o.Value = &v
	o.NullFields = removeNullField(o.NullFields, "Value")
}

// SetValueNil sets Value to an explicit JSON null when marshaled.
func (o *DataLakeAtlasStoreReadPreferenceTag) SetValueNil() {
	o.Value = nil
	o.NullFields = addNullField(o.NullFields, "Value")
}
