// Code based on the AtlasAPI V2 OpenAPI file

package admin

// LinkAtlas struct for LinkAtlas
type LinkAtlas struct {
	// Uniform Resource Locator (URL) that points another API resource to which this response has some relationship. This URL often begins with `https://cloud.mongodb.com/api/atlas`.
	Href *string `json:"href,omitempty"`
	// Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource. This URL often begins with `https://cloud.mongodb.com/api/atlas`.
	Rel *string `json:"rel,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *LinkAtlas) MarshalJSON() ([]byte, error) {
	type noMethod LinkAtlas
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewLinkAtlas instantiates a new LinkAtlas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkAtlas() *LinkAtlas {
	this := LinkAtlas{}
	return &this
}

// NewLinkAtlasWithDefaults instantiates a new LinkAtlas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkAtlasWithDefaults() *LinkAtlas {
	this := LinkAtlas{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise
func (o *LinkAtlas) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkAtlas) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}

	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LinkAtlas) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LinkAtlas) SetHref(v string) {
	o.Href = &v
}

// SetHrefNil sets Href to an explicit JSON null when marshaled.
func (o *LinkAtlas) SetHrefNil() {
	o.Href = nil
	o.NullFields = append(o.NullFields, "Href")
}

// GetRel returns the Rel field value if set, zero value otherwise
func (o *LinkAtlas) GetRel() string {
	if o == nil || IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkAtlas) GetRelOk() (*string, bool) {
	if o == nil || IsNil(o.Rel) {
		return nil, false
	}

	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *LinkAtlas) HasRel() bool {
	if o != nil && !IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *LinkAtlas) SetRel(v string) {
	o.Rel = &v
}

// SetRelNil sets Rel to an explicit JSON null when marshaled.
func (o *LinkAtlas) SetRelNil() {
	o.Rel = nil
	o.NullFields = append(o.NullFields, "Rel")
}
