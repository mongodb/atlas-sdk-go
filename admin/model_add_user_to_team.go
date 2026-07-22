// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AddUserToTeam struct for AddUserToTeam
type AddUserToTeam struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user.
	Id string `json:"id"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AddUserToTeam) MarshalJSON() ([]byte, error) {
	type noMethod AddUserToTeam
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAddUserToTeam instantiates a new AddUserToTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserToTeam(id string) *AddUserToTeam {
	this := AddUserToTeam{}
	this.Id = id
	return &this
}

// NewAddUserToTeamWithDefaults instantiates a new AddUserToTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserToTeamWithDefaults() *AddUserToTeam {
	this := AddUserToTeam{}
	return &this
}

// GetId returns the Id field value
func (o *AddUserToTeam) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddUserToTeam) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddUserToTeam) SetId(v string) {
	o.Id = v
}
