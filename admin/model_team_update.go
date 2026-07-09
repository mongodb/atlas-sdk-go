// Code based on the AtlasAPI V2 OpenAPI file

package admin

// TeamUpdate struct for TeamUpdate
type TeamUpdate struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Human-readable label that identifies the team.
	// Write only field.
	Name string `json:"name"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *TeamUpdate) MarshalJSON() ([]byte, error) {
	type noMethod TeamUpdate
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewTeamUpdate instantiates a new TeamUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamUpdate(name string) *TeamUpdate {
	this := TeamUpdate{}
	this.Name = name
	return &this
}

// NewTeamUpdateWithDefaults instantiates a new TeamUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamUpdateWithDefaults() *TeamUpdate {
	this := TeamUpdate{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *TeamUpdate) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdate) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TeamUpdate) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *TeamUpdate) SetLinks(v []Link) {
	o.Links = &v
}

// GetName returns the Name field value
func (o *TeamUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamUpdate) SetName(v string) {
	o.Name = v
}
