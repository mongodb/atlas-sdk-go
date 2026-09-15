// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiSearchDeploymentResponse struct for ApiSearchDeploymentResponse
type ApiSearchDeploymentResponse struct {
	// Autoscaling settings configured for this Search deployment.
	// Read only field.
	AutoScaling *ApiSearchAutoScaling `json:"autoScaling,omitempty"`
	// List of starting settings for the Search Nodes in each cluster region.
	// Read only field.
	BaselineSpecs *[]ApiSearchDeploymentEffectiveSpec `json:"baselineSpecs,omitempty"`
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
	// Deprecated. `specs` will be removed in a future release. We strongly recommend that you use `effectiveSpecs` instead.
	// Read only field.
	// Deprecated
	Specs *[]ApiSearchDeploymentSpec `json:"specs,omitempty"`
	// Human-readable label that indicates the current operating condition of this search deployment.
	// Read only field.
	StateName *string `json:"stateName,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiSearchDeploymentResponse) MarshalJSON() ([]byte, error) {
	type noMethod ApiSearchDeploymentResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiSearchDeploymentResponse instantiates a new ApiSearchDeploymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchDeploymentResponse() *ApiSearchDeploymentResponse {
	this := ApiSearchDeploymentResponse{}
	return &this
}

// NewApiSearchDeploymentResponseWithDefaults instantiates a new ApiSearchDeploymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchDeploymentResponseWithDefaults() *ApiSearchDeploymentResponse {
	this := ApiSearchDeploymentResponse{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise
func (o *ApiSearchDeploymentResponse) GetAutoScaling() ApiSearchAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret ApiSearchAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentResponse) GetAutoScalingOk() (*ApiSearchAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}

	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *ApiSearchDeploymentResponse) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given ApiSearchAutoScaling and assigns it to the AutoScaling field.
func (o *ApiSearchDeploymentResponse) SetAutoScaling(v ApiSearchAutoScaling) {
	o.AutoScaling = &v
	o.NullFields = removeNullField(o.NullFields, "AutoScaling")
}

// SetAutoScalingNil sets AutoScaling to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentResponse) SetAutoScalingNil() {
	o.AutoScaling = nil
	o.NullFields = addNullField(o.NullFields, "AutoScaling")
}

// GetBaselineSpecs returns the BaselineSpecs field value if set, zero value otherwise
func (o *ApiSearchDeploymentResponse) GetBaselineSpecs() []ApiSearchDeploymentEffectiveSpec {
	if o == nil || IsNil(o.BaselineSpecs) {
		var ret []ApiSearchDeploymentEffectiveSpec
		return ret
	}
	return *o.BaselineSpecs
}

// GetBaselineSpecsOk returns a tuple with the BaselineSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentResponse) GetBaselineSpecsOk() (*[]ApiSearchDeploymentEffectiveSpec, bool) {
	if o == nil || IsNil(o.BaselineSpecs) {
		return nil, false
	}

	return o.BaselineSpecs, true
}

// HasBaselineSpecs returns a boolean if a field has been set.
func (o *ApiSearchDeploymentResponse) HasBaselineSpecs() bool {
	if o != nil && !IsNil(o.BaselineSpecs) {
		return true
	}

	return false
}

// SetBaselineSpecs gets a reference to the given []ApiSearchDeploymentEffectiveSpec and assigns it to the BaselineSpecs field.
func (o *ApiSearchDeploymentResponse) SetBaselineSpecs(v []ApiSearchDeploymentEffectiveSpec) {
	o.BaselineSpecs = &v
	o.NullFields = removeNullField(o.NullFields, "BaselineSpecs")
}

// SetBaselineSpecsNil sets BaselineSpecs to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentResponse) SetBaselineSpecsNil() {
	o.BaselineSpecs = nil
	o.NullFields = addNullField(o.NullFields, "BaselineSpecs")
}

// GetEffectiveSpecs returns the EffectiveSpecs field value if set, zero value otherwise
func (o *ApiSearchDeploymentResponse) GetEffectiveSpecs() []ApiSearchDeploymentEffectiveSpec {
	if o == nil || IsNil(o.EffectiveSpecs) {
		var ret []ApiSearchDeploymentEffectiveSpec
		return ret
	}
	return *o.EffectiveSpecs
}

// GetEffectiveSpecsOk returns a tuple with the EffectiveSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentResponse) GetEffectiveSpecsOk() (*[]ApiSearchDeploymentEffectiveSpec, bool) {
	if o == nil || IsNil(o.EffectiveSpecs) {
		return nil, false
	}

	return o.EffectiveSpecs, true
}

// HasEffectiveSpecs returns a boolean if a field has been set.
func (o *ApiSearchDeploymentResponse) HasEffectiveSpecs() bool {
	if o != nil && !IsNil(o.EffectiveSpecs) {
		return true
	}

	return false
}

// SetEffectiveSpecs gets a reference to the given []ApiSearchDeploymentEffectiveSpec and assigns it to the EffectiveSpecs field.
func (o *ApiSearchDeploymentResponse) SetEffectiveSpecs(v []ApiSearchDeploymentEffectiveSpec) {
	o.EffectiveSpecs = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveSpecs")
}

// SetEffectiveSpecsNil sets EffectiveSpecs to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentResponse) SetEffectiveSpecsNil() {
	o.EffectiveSpecs = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveSpecs")
}

// GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field value if set, zero value otherwise
func (o *ApiSearchDeploymentResponse) GetEncryptionAtRestProvider() string {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		var ret string
		return ret
	}
	return *o.EncryptionAtRestProvider
}

// GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentResponse) GetEncryptionAtRestProviderOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		return nil, false
	}

	return o.EncryptionAtRestProvider, true
}

// HasEncryptionAtRestProvider returns a boolean if a field has been set.
func (o *ApiSearchDeploymentResponse) HasEncryptionAtRestProvider() bool {
	if o != nil && !IsNil(o.EncryptionAtRestProvider) {
		return true
	}

	return false
}

// SetEncryptionAtRestProvider gets a reference to the given string and assigns it to the EncryptionAtRestProvider field.
func (o *ApiSearchDeploymentResponse) SetEncryptionAtRestProvider(v string) {
	o.EncryptionAtRestProvider = &v
	o.NullFields = removeNullField(o.NullFields, "EncryptionAtRestProvider")
}

// SetEncryptionAtRestProviderNil sets EncryptionAtRestProvider to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentResponse) SetEncryptionAtRestProviderNil() {
	o.EncryptionAtRestProvider = nil
	o.NullFields = addNullField(o.NullFields, "EncryptionAtRestProvider")
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *ApiSearchDeploymentResponse) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentResponse) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiSearchDeploymentResponse) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiSearchDeploymentResponse) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentResponse) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ApiSearchDeploymentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiSearchDeploymentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiSearchDeploymentResponse) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentResponse) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetSpecs returns the Specs field value if set, zero value otherwise
// Deprecated
func (o *ApiSearchDeploymentResponse) GetSpecs() []ApiSearchDeploymentSpec {
	if o == nil || IsNil(o.Specs) {
		var ret []ApiSearchDeploymentSpec
		return ret
	}
	return *o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApiSearchDeploymentResponse) GetSpecsOk() (*[]ApiSearchDeploymentSpec, bool) {
	if o == nil || IsNil(o.Specs) {
		return nil, false
	}

	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *ApiSearchDeploymentResponse) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given []ApiSearchDeploymentSpec and assigns it to the Specs field.
// Deprecated
func (o *ApiSearchDeploymentResponse) SetSpecs(v []ApiSearchDeploymentSpec) {
	o.Specs = &v
	o.NullFields = removeNullField(o.NullFields, "Specs")
}

// SetSpecsNil sets Specs to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentResponse) SetSpecsNil() {
	o.Specs = nil
	o.NullFields = addNullField(o.NullFields, "Specs")
}

// GetStateName returns the StateName field value if set, zero value otherwise
func (o *ApiSearchDeploymentResponse) GetStateName() string {
	if o == nil || IsNil(o.StateName) {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentResponse) GetStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.StateName) {
		return nil, false
	}

	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *ApiSearchDeploymentResponse) HasStateName() bool {
	if o != nil && !IsNil(o.StateName) {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *ApiSearchDeploymentResponse) SetStateName(v string) {
	o.StateName = &v
	o.NullFields = removeNullField(o.NullFields, "StateName")
}

// SetStateNameNil sets StateName to an explicit JSON null when marshaled.
func (o *ApiSearchDeploymentResponse) SetStateNameNil() {
	o.StateName = nil
	o.NullFields = addNullField(o.NullFields, "StateName")
}
