// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasRestoreNamespace Source and optional target namespace for a restore.
type ApiAtlasRestoreNamespace struct {
	// Namespace requested to restore (e.g. database name or `database.collection`).
	SourceNamespace string `json:"sourceNamespace"`
	// Requested target namespace for the restored data; if empty, source namespace is used.
	TargetNamespace *string `json:"targetNamespace,omitempty"`
}

// NewApiAtlasRestoreNamespace instantiates a new ApiAtlasRestoreNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasRestoreNamespace(sourceNamespace string) *ApiAtlasRestoreNamespace {
	this := ApiAtlasRestoreNamespace{}
	this.SourceNamespace = sourceNamespace
	return &this
}

// NewApiAtlasRestoreNamespaceWithDefaults instantiates a new ApiAtlasRestoreNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasRestoreNamespaceWithDefaults() *ApiAtlasRestoreNamespace {
	this := ApiAtlasRestoreNamespace{}
	return &this
}

// GetSourceNamespace returns the SourceNamespace field value
func (o *ApiAtlasRestoreNamespace) GetSourceNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceNamespace
}

// GetSourceNamespaceOk returns a tuple with the SourceNamespace field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasRestoreNamespace) GetSourceNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceNamespace, true
}

// SetSourceNamespace sets field value
func (o *ApiAtlasRestoreNamespace) SetSourceNamespace(v string) {
	o.SourceNamespace = v
}

// GetTargetNamespace returns the TargetNamespace field value if set, zero value otherwise
func (o *ApiAtlasRestoreNamespace) GetTargetNamespace() string {
	if o == nil || IsNil(o.TargetNamespace) {
		var ret string
		return ret
	}
	return *o.TargetNamespace
}

// GetTargetNamespaceOk returns a tuple with the TargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasRestoreNamespace) GetTargetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNamespace) {
		return nil, false
	}

	return o.TargetNamespace, true
}

// HasTargetNamespace returns a boolean if a field has been set.
func (o *ApiAtlasRestoreNamespace) HasTargetNamespace() bool {
	if o != nil && !IsNil(o.TargetNamespace) {
		return true
	}

	return false
}

// SetTargetNamespace gets a reference to the given string and assigns it to the TargetNamespace field.
func (o *ApiAtlasRestoreNamespace) SetTargetNamespace(v string) {
	o.TargetNamespace = &v
}
