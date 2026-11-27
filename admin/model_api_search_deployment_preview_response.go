// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiSearchDeploymentPreviewResponse struct for ApiSearchDeploymentPreviewResponse
type ApiSearchDeploymentPreviewResponse struct {
	// Autoscaling settings configured for this Search deployment.
	// Read only field.
	AutoScaling *ApiSearchAutoScaling `json:"autoScaling,omitempty"`
	// List of settings that configure the Search Nodes for your cluster. Each entry describes one region or, when `shardId` is present, one shard in one region.
	// Read only field.
	EffectiveSpecs *[]ApiSearchDeploymentEffectiveSpec `json:"effectiveSpecs,omitempty"`
	// Cloud service provider that manages your customer keys to provide an additional layer of Encryption At Rest for the cluster.
	// Read only field.
	EncryptionAtRestProvider *string `json:"encryptionAtRestProvider,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the search deployment.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Resolved baseline settings for the Search Nodes in each cluster region. Autoscaling changes are reflected in `effectiveSpecs`.
	// Read only field.
	Specs *[]ApiSearchDeploymentSpec `json:"specs,omitempty"`
	// Human-readable label that indicates the current operating condition of this search deployment.
	// Read only field.
	StateName *string `json:"stateName,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o ApiSearchDeploymentPreviewResponse) MarshalJSON() ([]byte, error) {
	type noMethod ApiSearchDeploymentPreviewResponse
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewApiSearchDeploymentPreviewResponse instantiates a new ApiSearchDeploymentPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchDeploymentPreviewResponse() *ApiSearchDeploymentPreviewResponse {
	this := ApiSearchDeploymentPreviewResponse{}
	return &this
}

// NewApiSearchDeploymentPreviewResponseWithDefaults instantiates a new ApiSearchDeploymentPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchDeploymentPreviewResponseWithDefaults() *ApiSearchDeploymentPreviewResponse {
	this := ApiSearchDeploymentPreviewResponse{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise
func (o *ApiSearchDeploymentPreviewResponse) GetAutoScaling() ApiSearchAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret ApiSearchAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentPreviewResponse) GetAutoScalingOk() (*ApiSearchAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}

	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *ApiSearchDeploymentPreviewResponse) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given ApiSearchAutoScaling and assigns it to the AutoScaling field.
func (o *ApiSearchDeploymentPreviewResponse) SetAutoScaling(v ApiSearchAutoScaling) {
	o.AutoScaling = &v
	o.NullFields = removeNullField(o.NullFields, "AutoScaling")
}

// SetAutoScalingNil sets AutoScaling to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentPreviewResponse) SetAutoScalingNil() {
	o.AutoScaling = nil
	o.NullFields = addNullField(o.NullFields, "AutoScaling")
}

// GetEffectiveSpecs returns the EffectiveSpecs field value if set, zero value otherwise
func (o *ApiSearchDeploymentPreviewResponse) GetEffectiveSpecs() []ApiSearchDeploymentEffectiveSpec {
	if o == nil || IsNil(o.EffectiveSpecs) {
		var ret []ApiSearchDeploymentEffectiveSpec
		return ret
	}
	return *o.EffectiveSpecs
}

// GetEffectiveSpecsOk returns a tuple with the EffectiveSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentPreviewResponse) GetEffectiveSpecsOk() (*[]ApiSearchDeploymentEffectiveSpec, bool) {
	if o == nil || IsNil(o.EffectiveSpecs) {
		return nil, false
	}

	return o.EffectiveSpecs, true
}

// HasEffectiveSpecs returns a boolean if a field has been set.
func (o *ApiSearchDeploymentPreviewResponse) HasEffectiveSpecs() bool {
	if o != nil && !IsNil(o.EffectiveSpecs) {
		return true
	}

	return false
}

// SetEffectiveSpecs gets a reference to the given []ApiSearchDeploymentEffectiveSpec and assigns it to the EffectiveSpecs field.
func (o *ApiSearchDeploymentPreviewResponse) SetEffectiveSpecs(v []ApiSearchDeploymentEffectiveSpec) {
	o.EffectiveSpecs = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveSpecs")
}

// SetEffectiveSpecsNil sets EffectiveSpecs to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentPreviewResponse) SetEffectiveSpecsNil() {
	o.EffectiveSpecs = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveSpecs")
}

// GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field value if set, zero value otherwise
func (o *ApiSearchDeploymentPreviewResponse) GetEncryptionAtRestProvider() string {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		var ret string
		return ret
	}
	return *o.EncryptionAtRestProvider
}

// GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentPreviewResponse) GetEncryptionAtRestProviderOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		return nil, false
	}

	return o.EncryptionAtRestProvider, true
}

// HasEncryptionAtRestProvider returns a boolean if a field has been set.
func (o *ApiSearchDeploymentPreviewResponse) HasEncryptionAtRestProvider() bool {
	if o != nil && !IsNil(o.EncryptionAtRestProvider) {
		return true
	}

	return false
}

// SetEncryptionAtRestProvider gets a reference to the given string and assigns it to the EncryptionAtRestProvider field.
func (o *ApiSearchDeploymentPreviewResponse) SetEncryptionAtRestProvider(v string) {
	o.EncryptionAtRestProvider = &v
	o.NullFields = removeNullField(o.NullFields, "EncryptionAtRestProvider")
}

// SetEncryptionAtRestProviderNil sets EncryptionAtRestProvider to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentPreviewResponse) SetEncryptionAtRestProviderNil() {
	o.EncryptionAtRestProvider = nil
	o.NullFields = addNullField(o.NullFields, "EncryptionAtRestProvider")
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *ApiSearchDeploymentPreviewResponse) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentPreviewResponse) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiSearchDeploymentPreviewResponse) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiSearchDeploymentPreviewResponse) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentPreviewResponse) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ApiSearchDeploymentPreviewResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentPreviewResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiSearchDeploymentPreviewResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiSearchDeploymentPreviewResponse) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentPreviewResponse) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetSpecs returns the Specs field value if set, zero value otherwise
func (o *ApiSearchDeploymentPreviewResponse) GetSpecs() []ApiSearchDeploymentSpec {
	if o == nil || IsNil(o.Specs) {
		var ret []ApiSearchDeploymentSpec
		return ret
	}
	return *o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentPreviewResponse) GetSpecsOk() (*[]ApiSearchDeploymentSpec, bool) {
	if o == nil || IsNil(o.Specs) {
		return nil, false
	}

	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *ApiSearchDeploymentPreviewResponse) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given []ApiSearchDeploymentSpec and assigns it to the Specs field.
func (o *ApiSearchDeploymentPreviewResponse) SetSpecs(v []ApiSearchDeploymentSpec) {
	o.Specs = &v
	o.NullFields = removeNullField(o.NullFields, "Specs")
}

// SetSpecsNil sets Specs to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentPreviewResponse) SetSpecsNil() {
	o.Specs = nil
	o.NullFields = addNullField(o.NullFields, "Specs")
}

// GetStateName returns the StateName field value if set, zero value otherwise
func (o *ApiSearchDeploymentPreviewResponse) GetStateName() string {
	if o == nil || IsNil(o.StateName) {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentPreviewResponse) GetStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.StateName) {
		return nil, false
	}

	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *ApiSearchDeploymentPreviewResponse) HasStateName() bool {
	if o != nil && !IsNil(o.StateName) {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *ApiSearchDeploymentPreviewResponse) SetStateName(v string) {
	o.StateName = &v
	o.NullFields = removeNullField(o.NullFields, "StateName")
}

// SetStateNameNil sets StateName to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentPreviewResponse) SetStateNameNil() {
	o.StateName = nil
	o.NullFields = addNullField(o.NullFields, "StateName")
}
