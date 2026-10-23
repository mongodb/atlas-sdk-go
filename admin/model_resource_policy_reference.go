// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ResourcePolicyReference Atlas Resource Policy that caused the request to be denied.
type ResourcePolicyReference struct {
	// Unique 24-hexadecimal character string that identifies the Atlas Resource Policy.
	// Read only field.
	Id string `json:"id"`
	// Human-readable label that identifies the Atlas Resource Policy.
	// Read only field.
	Name string `json:"name"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ResourcePolicyReference) MarshalJSON() ([]byte, error) {
	type noMethod ResourcePolicyReference
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewResourcePolicyReference instantiates a new ResourcePolicyReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePolicyReference(id string, name string) *ResourcePolicyReference {
	this := ResourcePolicyReference{}
	this.Id = id
	this.Name = name
	return &this
}

// NewResourcePolicyReferenceWithDefaults instantiates a new ResourcePolicyReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePolicyReferenceWithDefaults() *ResourcePolicyReference {
	this := ResourcePolicyReference{}
	return &this
}

// GetId returns the Id field value
func (o *ResourcePolicyReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourcePolicyReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourcePolicyReference) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResourcePolicyReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourcePolicyReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourcePolicyReference) SetName(v string) {
	o.Name = v
}
