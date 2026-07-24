// Code based on the AtlasAPI V2 OpenAPI file

package admin

// FlexConnectionStrings20241113 Collection of Uniform Resource Locators that point to the MongoDB database.
type FlexConnectionStrings20241113 struct {
	// Public connection string that you can use to connect to this cluster. This connection string uses the `mongodb://` protocol.
	// Read only field.
	Standard *string `json:"standard,omitempty"`
	// Public connection string that you can use to connect to this flex cluster. This connection string uses the `mongodb+srv://` protocol.
	// Read only field.
	StandardSrv *string `json:"standardSrv,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *FlexConnectionStrings20241113) MarshalJSON() ([]byte, error) {
	type noMethod FlexConnectionStrings20241113
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewFlexConnectionStrings20241113 instantiates a new FlexConnectionStrings20241113 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexConnectionStrings20241113() *FlexConnectionStrings20241113 {
	this := FlexConnectionStrings20241113{}
	return &this
}

// NewFlexConnectionStrings20241113WithDefaults instantiates a new FlexConnectionStrings20241113 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexConnectionStrings20241113WithDefaults() *FlexConnectionStrings20241113 {
	this := FlexConnectionStrings20241113{}
	return &this
}

// GetStandard returns the Standard field value if set, zero value otherwise
func (o *FlexConnectionStrings20241113) GetStandard() string {
	if o == nil || IsNil(o.Standard) {
		var ret string
		return ret
	}
	return *o.Standard
}

// GetStandardOk returns a tuple with the Standard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexConnectionStrings20241113) GetStandardOk() (*string, bool) {
	if o == nil || IsNil(o.Standard) {
		return nil, false
	}

	return o.Standard, true
}

// HasStandard returns a boolean if a field has been set.
func (o *FlexConnectionStrings20241113) HasStandard() bool {
	if o != nil && !IsNil(o.Standard) {
		return true
	}

	return false
}

// SetStandard gets a reference to the given string and assigns it to the Standard field.
func (o *FlexConnectionStrings20241113) SetStandard(v string) {
	o.Standard = &v
	o.NullFields = removeNullField(o.NullFields, "Standard")
}

// SetStandardNil sets Standard to an explicit JSON null when marshaled.
func (o *FlexConnectionStrings20241113) SetStandardNil() {
	o.Standard = nil
	o.NullFields = addNullField(o.NullFields, "Standard")
}

// GetStandardSrv returns the StandardSrv field value if set, zero value otherwise
func (o *FlexConnectionStrings20241113) GetStandardSrv() string {
	if o == nil || IsNil(o.StandardSrv) {
		var ret string
		return ret
	}
	return *o.StandardSrv
}

// GetStandardSrvOk returns a tuple with the StandardSrv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexConnectionStrings20241113) GetStandardSrvOk() (*string, bool) {
	if o == nil || IsNil(o.StandardSrv) {
		return nil, false
	}

	return o.StandardSrv, true
}

// HasStandardSrv returns a boolean if a field has been set.
func (o *FlexConnectionStrings20241113) HasStandardSrv() bool {
	if o != nil && !IsNil(o.StandardSrv) {
		return true
	}

	return false
}

// SetStandardSrv gets a reference to the given string and assigns it to the StandardSrv field.
func (o *FlexConnectionStrings20241113) SetStandardSrv(v string) {
	o.StandardSrv = &v
	o.NullFields = removeNullField(o.NullFields, "StandardSrv")
}

// SetStandardSrvNil sets StandardSrv to an explicit JSON null when marshaled.
func (o *FlexConnectionStrings20241113) SetStandardSrvNil() {
	o.StandardSrv = nil
	o.NullFields = addNullField(o.NullFields, "StandardSrv")
}
