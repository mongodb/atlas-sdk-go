// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AdditionalData Additional metadata associated with the line item.
type AdditionalData struct {
	// Identifier of the stream processor associated with the line item.
	ProcessorId *string `json:"processorId,omitempty"`
	// Name of the stream processor associated with the line item.
	ProcessorName *string `json:"processorName,omitempty"`
	// Workspace associated with the line item.
	Workspace *string `json:"workspace,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AdditionalData) MarshalJSON() ([]byte, error) {
	type noMethod AdditionalData
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAdditionalData instantiates a new AdditionalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalData() *AdditionalData {
	this := AdditionalData{}
	return &this
}

// NewAdditionalDataWithDefaults instantiates a new AdditionalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataWithDefaults() *AdditionalData {
	this := AdditionalData{}
	return &this
}

// GetProcessorId returns the ProcessorId field value if set, zero value otherwise
func (o *AdditionalData) GetProcessorId() string {
	if o == nil || IsNil(o.ProcessorId) {
		var ret string
		return ret
	}
	return *o.ProcessorId
}

// GetProcessorIdOk returns a tuple with the ProcessorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData) GetProcessorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorId) {
		return nil, false
	}

	return o.ProcessorId, true
}

// HasProcessorId returns a boolean if a field has been set.
func (o *AdditionalData) HasProcessorId() bool {
	if o != nil && !IsNil(o.ProcessorId) {
		return true
	}

	return false
}

// SetProcessorId gets a reference to the given string and assigns it to the ProcessorId field.
func (o *AdditionalData) SetProcessorId(v string) {
	o.ProcessorId = &v
	o.NullFields = removeNullField(o.NullFields, "ProcessorId")
}

// SetProcessorIdNil sets ProcessorId to an explicit JSON null when marshaled.
func (o *AdditionalData) SetProcessorIdNil() {
	o.ProcessorId = nil
	o.NullFields = addNullField(o.NullFields, "ProcessorId")
}

// GetProcessorName returns the ProcessorName field value if set, zero value otherwise
func (o *AdditionalData) GetProcessorName() string {
	if o == nil || IsNil(o.ProcessorName) {
		var ret string
		return ret
	}
	return *o.ProcessorName
}

// GetProcessorNameOk returns a tuple with the ProcessorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData) GetProcessorNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorName) {
		return nil, false
	}

	return o.ProcessorName, true
}

// HasProcessorName returns a boolean if a field has been set.
func (o *AdditionalData) HasProcessorName() bool {
	if o != nil && !IsNil(o.ProcessorName) {
		return true
	}

	return false
}

// SetProcessorName gets a reference to the given string and assigns it to the ProcessorName field.
func (o *AdditionalData) SetProcessorName(v string) {
	o.ProcessorName = &v
	o.NullFields = removeNullField(o.NullFields, "ProcessorName")
}

// SetProcessorNameNil sets ProcessorName to an explicit JSON null when marshaled.
func (o *AdditionalData) SetProcessorNameNil() {
	o.ProcessorName = nil
	o.NullFields = addNullField(o.NullFields, "ProcessorName")
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise
func (o *AdditionalData) GetWorkspace() string {
	if o == nil || IsNil(o.Workspace) {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData) GetWorkspaceOk() (*string, bool) {
	if o == nil || IsNil(o.Workspace) {
		return nil, false
	}

	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *AdditionalData) HasWorkspace() bool {
	if o != nil && !IsNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *AdditionalData) SetWorkspace(v string) {
	o.Workspace = &v
	o.NullFields = removeNullField(o.NullFields, "Workspace")
}

// SetWorkspaceNil sets Workspace to an explicit JSON null when marshaled.
func (o *AdditionalData) SetWorkspaceNil() {
	o.Workspace = nil
	o.NullFields = addNullField(o.NullFields, "Workspace")
}
