// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AuthenticatedUser Details about the MongoDB Cloud user that this request is authenticated as.
type AuthenticatedUser struct {
	// Email address that represents the username of the MongoDB Cloud user.
	// Read only field.
	Username *string `json:"username,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AuthenticatedUser) MarshalJSON() ([]byte, error) {
	type noMethod AuthenticatedUser
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAuthenticatedUser instantiates a new AuthenticatedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedUser() *AuthenticatedUser {
	this := AuthenticatedUser{}
	return &this
}

// NewAuthenticatedUserWithDefaults instantiates a new AuthenticatedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedUserWithDefaults() *AuthenticatedUser {
	this := AuthenticatedUser{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *AuthenticatedUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AuthenticatedUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AuthenticatedUser) SetUsername(v string) {
	o.Username = &v
	o.NullFields = removeNullField(o.NullFields, "Username")
}

// SetUsernameNil sets Username to an explicit JSON null when marshaled.
func (o *AuthenticatedUser) SetUsernameNil() {
	o.Username = nil
	o.NullFields = addNullField(o.NullFields, "Username")
}
