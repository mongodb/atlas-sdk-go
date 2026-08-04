// Code based on the AtlasAPI V2 OpenAPI file

package admin

// CreateAtlasProjectApiKey struct for CreateAtlasProjectApiKey
type CreateAtlasProjectApiKey struct {
	// Purpose or explanation provided when someone created this project API key.
	Desc string `json:"desc"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project.
	Roles []string `json:"roles"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *CreateAtlasProjectApiKey) MarshalJSON() ([]byte, error) {
	type noMethod CreateAtlasProjectApiKey
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewCreateAtlasProjectApiKey instantiates a new CreateAtlasProjectApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAtlasProjectApiKey(desc string, roles []string) *CreateAtlasProjectApiKey {
	this := CreateAtlasProjectApiKey{}
	this.Desc = desc
	this.Roles = roles
	return &this
}

// NewCreateAtlasProjectApiKeyWithDefaults instantiates a new CreateAtlasProjectApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAtlasProjectApiKeyWithDefaults() *CreateAtlasProjectApiKey {
	this := CreateAtlasProjectApiKey{}
	return &this
}

// GetDesc returns the Desc field value
func (o *CreateAtlasProjectApiKey) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *CreateAtlasProjectApiKey) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *CreateAtlasProjectApiKey) SetDesc(v string) {
	o.Desc = v
}

// GetRoles returns the Roles field value
func (o *CreateAtlasProjectApiKey) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CreateAtlasProjectApiKey) GetRolesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *CreateAtlasProjectApiKey) SetRoles(v []string) {
	o.Roles = v
}
