// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SearchTypeSets Type sets for an Atlas Search index definition.
type SearchTypeSets struct {
	// Label that identifies the type set name. Each `typeSets.name` must be unique within the same index definition.
	Name string `json:"name"`
	// List of types associated with the type set. Each type definition must include a `type` field specifying the search field type (`autocomplete`, `boolean`, `date`, `geo`, `number`, `objectId`, `string`, `token`, or `uuid`) and may include additional configuration properties specific to that type.
	Types *[]any `json:"types,omitempty"`
}

// NewSearchTypeSets instantiates a new SearchTypeSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchTypeSets(name string) *SearchTypeSets {
	this := SearchTypeSets{}
	this.Name = name
	return &this
}

// NewSearchTypeSetsWithDefaults instantiates a new SearchTypeSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchTypeSetsWithDefaults() *SearchTypeSets {
	this := SearchTypeSets{}
	return &this
}

// GetName returns the Name field value
func (o *SearchTypeSets) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SearchTypeSets) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SearchTypeSets) SetName(v string) {
	o.Name = v
}

// GetTypes returns the Types field value if set, zero value otherwise
func (o *SearchTypeSets) GetTypes() []any {
	if o == nil || IsNil(o.Types) {
		var ret []any
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchTypeSets) GetTypesOk() (*[]any, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}

	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *SearchTypeSets) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []any and assigns it to the Types field.
func (o *SearchTypeSets) SetTypes(v []any) {
	o.Types = &v
}
