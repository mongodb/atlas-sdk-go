// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OpenApiInfo struct for OpenApiInfo
type OpenApiInfo struct {
	Info *Info `json:"info,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OpenApiInfo) MarshalJSON() ([]byte, error) {
	type noMethod OpenApiInfo
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOpenApiInfo instantiates a new OpenApiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenApiInfo() *OpenApiInfo {
	this := OpenApiInfo{}
	return &this
}

// NewOpenApiInfoWithDefaults instantiates a new OpenApiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenApiInfoWithDefaults() *OpenApiInfo {
	this := OpenApiInfo{}
	return &this
}

// GetInfo returns the Info field value if set, zero value otherwise
func (o *OpenApiInfo) GetInfo() Info {
	if o == nil || IsNil(o.Info) {
		var ret Info
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiInfo) GetInfoOk() (*Info, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}

	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *OpenApiInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given Info and assigns it to the Info field.
func (o *OpenApiInfo) SetInfo(v Info) {
	o.Info = &v
	o.NullFields = removeNullField(o.NullFields, "Info")
}

// SetInfoNil sets Info to an explicit JSON null when marshaled.
func (o *OpenApiInfo) SetInfoNil() {
	o.Info = nil
	o.NullFields = addNullField(o.NullFields, "Info")
}
