// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasRestoreCollectionNamespace Source and optional target collection for a restore.
type ApiAtlasRestoreCollectionNamespace struct {
	// Collection requested to restore, as `database.collection`.
	SourceNamespace string `json:"sourceNamespace"`
	// Requested target collection as `database.collection`; if empty, source namespace is used.
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiAtlasRestoreCollectionNamespace) MarshalJSON() ([]byte, error) {
	type noMethod ApiAtlasRestoreCollectionNamespace
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiAtlasRestoreCollectionNamespace instantiates a new ApiAtlasRestoreCollectionNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasRestoreCollectionNamespace(sourceNamespace string) *ApiAtlasRestoreCollectionNamespace {
	this := ApiAtlasRestoreCollectionNamespace{}
	this.SourceNamespace = sourceNamespace
	return &this
}

// NewApiAtlasRestoreCollectionNamespaceWithDefaults instantiates a new ApiAtlasRestoreCollectionNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasRestoreCollectionNamespaceWithDefaults() *ApiAtlasRestoreCollectionNamespace {
	this := ApiAtlasRestoreCollectionNamespace{}
	return &this
}

// GetSourceNamespace returns the SourceNamespace field value
func (o *ApiAtlasRestoreCollectionNamespace) GetSourceNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceNamespace
}

// GetSourceNamespaceOk returns a tuple with the SourceNamespace field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasRestoreCollectionNamespace) GetSourceNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceNamespace, true
}

// SetSourceNamespace sets field value
func (o *ApiAtlasRestoreCollectionNamespace) SetSourceNamespace(v string) {
	o.SourceNamespace = v
}

// GetTargetNamespace returns the TargetNamespace field value if set, zero value otherwise
func (o *ApiAtlasRestoreCollectionNamespace) GetTargetNamespace() string {
	if o == nil || IsNil(o.TargetNamespace) {
		var ret string
		return ret
	}
	return *o.TargetNamespace
}

// GetTargetNamespaceOk returns a tuple with the TargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasRestoreCollectionNamespace) GetTargetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNamespace) {
		return nil, false
	}

	return o.TargetNamespace, true
}

// HasTargetNamespace returns a boolean if a field has been set.
func (o *ApiAtlasRestoreCollectionNamespace) HasTargetNamespace() bool {
	if o != nil && !IsNil(o.TargetNamespace) {
		return true
	}

	return false
}

// SetTargetNamespace gets a reference to the given string and assigns it to the TargetNamespace field.
func (o *ApiAtlasRestoreCollectionNamespace) SetTargetNamespace(v string) {
	o.TargetNamespace = &v
	o.NullFields = removeNullField(o.NullFields, "TargetNamespace")
}

// SetTargetNamespaceNil sets TargetNamespace to an explicit JSON null when marshaled.
func (o *ApiAtlasRestoreCollectionNamespace) SetTargetNamespaceNil() {
	o.TargetNamespace = nil
	o.NullFields = addNullField(o.NullFields, "TargetNamespace")
}
