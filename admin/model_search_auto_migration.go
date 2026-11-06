// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SearchAutoMigration Configuration for automatically adding Dedicated Search Nodes on eligible Atlas Infinite clusters. This setting has no effect on Atlas Core clusters.
type SearchAutoMigration struct {
	// Flag that disables future automatic Search Node provisioning. Defaults to false on creation. Omission on update preserves the stored choice. Disabling retains existing Search Nodes and allows in-flight migrations to continue. Enabling does not bypass resource policies; if indexes exceed the migration threshold, the next eligible evaluation may start provisioning.
	Disabled *bool `json:"disabled,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *SearchAutoMigration) MarshalJSON() ([]byte, error) {
	type noMethod SearchAutoMigration
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewSearchAutoMigration instantiates a new SearchAutoMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAutoMigration() *SearchAutoMigration {
	this := SearchAutoMigration{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewSearchAutoMigrationWithDefaults instantiates a new SearchAutoMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAutoMigrationWithDefaults() *SearchAutoMigration {
	this := SearchAutoMigration{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise
func (o *SearchAutoMigration) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAutoMigration) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}

	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *SearchAutoMigration) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *SearchAutoMigration) SetDisabled(v bool) {
	o.Disabled = &v
	o.NullFields = removeNullField(o.NullFields, "Disabled")
}

// SetDisabledNil sets Disabled to an explicit JSON null when marshaled.
func (o *SearchAutoMigration) SetDisabledNil() {
	o.Disabled = nil
	o.NullFields = addNullField(o.NullFields, "Disabled")
}
