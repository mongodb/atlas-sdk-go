// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasCollectionRestoreCollectionStateResponse Collection-level state within a collection restore job.
type ApiAtlasCollectionRestoreCollectionStateResponse struct {
	// Actual target namespace after restore (e.g. after conflict rename).
	// Read only field.
	EffectiveTargetNamespace *string                               `json:"effectiveTargetNamespace,omitempty"`
	IndexStatus              *ApiAtlasCollectionRestoreIndexStatus `json:"indexStatus,omitempty"`
	// Number of documents restored so far.
	// Read only field.
	RestoredDocuments *int64 `json:"restoredDocuments,omitempty"`
	// Source namespace that was requested to restore.
	// Read only field.
	SourceNamespace *string `json:"sourceNamespace,omitempty"`
	// Current state of this collection within the restore job.
	// Read only field.
	State *string `json:"state,omitempty"`
	// Requested target namespace for the restored collection.
	// Read only field.
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	// Total document count for this collection.
	// Read only field.
	TotalDocuments *int64 `json:"totalDocuments,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) MarshalJSON() ([]byte, error) {
	type noMethod ApiAtlasCollectionRestoreCollectionStateResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiAtlasCollectionRestoreCollectionStateResponse instantiates a new ApiAtlasCollectionRestoreCollectionStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCollectionRestoreCollectionStateResponse() *ApiAtlasCollectionRestoreCollectionStateResponse {
	this := ApiAtlasCollectionRestoreCollectionStateResponse{}
	return &this
}

// NewApiAtlasCollectionRestoreCollectionStateResponseWithDefaults instantiates a new ApiAtlasCollectionRestoreCollectionStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCollectionRestoreCollectionStateResponseWithDefaults() *ApiAtlasCollectionRestoreCollectionStateResponse {
	this := ApiAtlasCollectionRestoreCollectionStateResponse{}
	return &this
}

// GetEffectiveTargetNamespace returns the EffectiveTargetNamespace field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetEffectiveTargetNamespace() string {
	if o == nil || IsNil(o.EffectiveTargetNamespace) {
		var ret string
		return ret
	}
	return *o.EffectiveTargetNamespace
}

// GetEffectiveTargetNamespaceOk returns a tuple with the EffectiveTargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetEffectiveTargetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveTargetNamespace) {
		return nil, false
	}

	return o.EffectiveTargetNamespace, true
}

// HasEffectiveTargetNamespace returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasEffectiveTargetNamespace() bool {
	if o != nil && !IsNil(o.EffectiveTargetNamespace) {
		return true
	}

	return false
}

// SetEffectiveTargetNamespace gets a reference to the given string and assigns it to the EffectiveTargetNamespace field.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetEffectiveTargetNamespace(v string) {
	o.EffectiveTargetNamespace = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveTargetNamespace")
}

// SetEffectiveTargetNamespaceNil sets EffectiveTargetNamespace to an explicit JSON null when marshaled.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetEffectiveTargetNamespaceNil() {
	o.EffectiveTargetNamespace = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveTargetNamespace")
}

// GetIndexStatus returns the IndexStatus field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetIndexStatus() ApiAtlasCollectionRestoreIndexStatus {
	if o == nil || IsNil(o.IndexStatus) {
		var ret ApiAtlasCollectionRestoreIndexStatus
		return ret
	}
	return *o.IndexStatus
}

// GetIndexStatusOk returns a tuple with the IndexStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetIndexStatusOk() (*ApiAtlasCollectionRestoreIndexStatus, bool) {
	if o == nil || IsNil(o.IndexStatus) {
		return nil, false
	}

	return o.IndexStatus, true
}

// HasIndexStatus returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasIndexStatus() bool {
	if o != nil && !IsNil(o.IndexStatus) {
		return true
	}

	return false
}

// SetIndexStatus gets a reference to the given ApiAtlasCollectionRestoreIndexStatus and assigns it to the IndexStatus field.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetIndexStatus(v ApiAtlasCollectionRestoreIndexStatus) {
	o.IndexStatus = &v
	o.NullFields = removeNullField(o.NullFields, "IndexStatus")
}

// SetIndexStatusNil sets IndexStatus to an explicit JSON null when marshaled.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetIndexStatusNil() {
	o.IndexStatus = nil
	o.NullFields = addNullField(o.NullFields, "IndexStatus")
}

// GetRestoredDocuments returns the RestoredDocuments field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetRestoredDocuments() int64 {
	if o == nil || IsNil(o.RestoredDocuments) {
		var ret int64
		return ret
	}
	return *o.RestoredDocuments
}

// GetRestoredDocumentsOk returns a tuple with the RestoredDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetRestoredDocumentsOk() (*int64, bool) {
	if o == nil || IsNil(o.RestoredDocuments) {
		return nil, false
	}

	return o.RestoredDocuments, true
}

// HasRestoredDocuments returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasRestoredDocuments() bool {
	if o != nil && !IsNil(o.RestoredDocuments) {
		return true
	}

	return false
}

// SetRestoredDocuments gets a reference to the given int64 and assigns it to the RestoredDocuments field.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetRestoredDocuments(v int64) {
	o.RestoredDocuments = &v
	o.NullFields = removeNullField(o.NullFields, "RestoredDocuments")
}

// SetRestoredDocumentsNil sets RestoredDocuments to an explicit JSON null when marshaled.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetRestoredDocumentsNil() {
	o.RestoredDocuments = nil
	o.NullFields = addNullField(o.NullFields, "RestoredDocuments")
}

// GetSourceNamespace returns the SourceNamespace field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetSourceNamespace() string {
	if o == nil || IsNil(o.SourceNamespace) {
		var ret string
		return ret
	}
	return *o.SourceNamespace
}

// GetSourceNamespaceOk returns a tuple with the SourceNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetSourceNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.SourceNamespace) {
		return nil, false
	}

	return o.SourceNamespace, true
}

// HasSourceNamespace returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasSourceNamespace() bool {
	if o != nil && !IsNil(o.SourceNamespace) {
		return true
	}

	return false
}

// SetSourceNamespace gets a reference to the given string and assigns it to the SourceNamespace field.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetSourceNamespace(v string) {
	o.SourceNamespace = &v
	o.NullFields = removeNullField(o.NullFields, "SourceNamespace")
}

// SetSourceNamespaceNil sets SourceNamespace to an explicit JSON null when marshaled.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetSourceNamespaceNil() {
	o.SourceNamespace = nil
	o.NullFields = addNullField(o.NullFields, "SourceNamespace")
}

// GetState returns the State field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetState(v string) {
	o.State = &v
	o.NullFields = removeNullField(o.NullFields, "State")
}

// SetStateNil sets State to an explicit JSON null when marshaled.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetStateNil() {
	o.State = nil
	o.NullFields = addNullField(o.NullFields, "State")
}

// GetTargetNamespace returns the TargetNamespace field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetTargetNamespace() string {
	if o == nil || IsNil(o.TargetNamespace) {
		var ret string
		return ret
	}
	return *o.TargetNamespace
}

// GetTargetNamespaceOk returns a tuple with the TargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetTargetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNamespace) {
		return nil, false
	}

	return o.TargetNamespace, true
}

// HasTargetNamespace returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasTargetNamespace() bool {
	if o != nil && !IsNil(o.TargetNamespace) {
		return true
	}

	return false
}

// SetTargetNamespace gets a reference to the given string and assigns it to the TargetNamespace field.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetTargetNamespace(v string) {
	o.TargetNamespace = &v
	o.NullFields = removeNullField(o.NullFields, "TargetNamespace")
}

// SetTargetNamespaceNil sets TargetNamespace to an explicit JSON null when marshaled.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetTargetNamespaceNil() {
	o.TargetNamespace = nil
	o.NullFields = addNullField(o.NullFields, "TargetNamespace")
}

// GetTotalDocuments returns the TotalDocuments field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetTotalDocuments() int64 {
	if o == nil || IsNil(o.TotalDocuments) {
		var ret int64
		return ret
	}
	return *o.TotalDocuments
}

// GetTotalDocumentsOk returns a tuple with the TotalDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetTotalDocumentsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDocuments) {
		return nil, false
	}

	return o.TotalDocuments, true
}

// HasTotalDocuments returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasTotalDocuments() bool {
	if o != nil && !IsNil(o.TotalDocuments) {
		return true
	}

	return false
}

// SetTotalDocuments gets a reference to the given int64 and assigns it to the TotalDocuments field.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetTotalDocuments(v int64) {
	o.TotalDocuments = &v
	o.NullFields = removeNullField(o.NullFields, "TotalDocuments")
}

// SetTotalDocumentsNil sets TotalDocuments to an explicit JSON null when marshaled.
func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetTotalDocumentsNil() {
	o.TotalDocuments = nil
	o.NullFields = addNullField(o.NullFields, "TotalDocuments")
}
