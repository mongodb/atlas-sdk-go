// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiVersion struct for ApiVersion
type ApiVersion struct {
	// Object representing a version of the Atlas Admin API.
	Version *string `json:"version,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiVersion) MarshalJSON() ([]byte, error) {
	type noMethod ApiVersion
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiVersion instantiates a new ApiVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiVersion() *ApiVersion {
	this := ApiVersion{}
	return &this
}

// NewApiVersionWithDefaults instantiates a new ApiVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiVersionWithDefaults() *ApiVersion {
	this := ApiVersion{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise
func (o *ApiVersion) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVersion) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}

	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApiVersion) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApiVersion) SetVersion(v string) {
	o.Version = &v
	o.NullFields = removeNullField(o.NullFields, "Version")
}

// SetVersionNil sets Version to an explicit JSON null when marshaled.
func (o *ApiVersion) SetVersionNil() {
	o.Version = nil
	o.NullFields = addNullField(o.NullFields, "Version")
}
