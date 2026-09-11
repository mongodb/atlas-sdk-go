// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasRestoreDatabaseNamespace Source and optional target database for a restore.
type ApiAtlasRestoreDatabaseNamespace struct {
	// Database name requested to restore.
	SourceNamespace string `json:"sourceNamespace"`
	// Requested target database name; if empty, source database name is used.
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiAtlasRestoreDatabaseNamespace) MarshalJSON() ([]byte, error) {
	type noMethod ApiAtlasRestoreDatabaseNamespace
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiAtlasRestoreDatabaseNamespace instantiates a new ApiAtlasRestoreDatabaseNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasRestoreDatabaseNamespace(sourceNamespace string) *ApiAtlasRestoreDatabaseNamespace {
	this := ApiAtlasRestoreDatabaseNamespace{}
	this.SourceNamespace = sourceNamespace
	return &this
}

// NewApiAtlasRestoreDatabaseNamespaceWithDefaults instantiates a new ApiAtlasRestoreDatabaseNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasRestoreDatabaseNamespaceWithDefaults() *ApiAtlasRestoreDatabaseNamespace {
	this := ApiAtlasRestoreDatabaseNamespace{}
	return &this
}

// GetSourceNamespace returns the SourceNamespace field value
func (o *ApiAtlasRestoreDatabaseNamespace) GetSourceNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceNamespace
}

// GetSourceNamespaceOk returns a tuple with the SourceNamespace field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasRestoreDatabaseNamespace) GetSourceNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceNamespace, true
}

// SetSourceNamespace sets field value
func (o *ApiAtlasRestoreDatabaseNamespace) SetSourceNamespace(v string) {
	o.SourceNamespace = v
}

// GetTargetNamespace returns the TargetNamespace field value if set, zero value otherwise
func (o *ApiAtlasRestoreDatabaseNamespace) GetTargetNamespace() string {
	if o == nil || IsNil(o.TargetNamespace) {
		var ret string
		return ret
	}
	return *o.TargetNamespace
}

// GetTargetNamespaceOk returns a tuple with the TargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasRestoreDatabaseNamespace) GetTargetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNamespace) {
		return nil, false
	}

	return o.TargetNamespace, true
}

// HasTargetNamespace returns a boolean if a field has been set.
func (o *ApiAtlasRestoreDatabaseNamespace) HasTargetNamespace() bool {
	if o != nil && !IsNil(o.TargetNamespace) {
		return true
	}

	return false
}

// SetTargetNamespace gets a reference to the given string and assigns it to the TargetNamespace field.
func (o *ApiAtlasRestoreDatabaseNamespace) SetTargetNamespace(v string) {
	o.TargetNamespace = &v
	o.NullFields = removeNullField(o.NullFields, "TargetNamespace")
}

// SetTargetNamespaceNil sets TargetNamespace to an explicit JSON null when marshaled.
func (o *ApiAtlasRestoreDatabaseNamespace) SetTargetNamespaceNil() {
	o.TargetNamespace = nil
	o.NullFields = addNullField(o.NullFields, "TargetNamespace")
}
