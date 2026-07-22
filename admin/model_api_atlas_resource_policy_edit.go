// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasResourcePolicyEdit struct for ApiAtlasResourcePolicyEdit
type ApiAtlasResourcePolicyEdit struct {
	// Description of the atlas resource policy.
	Description *string `json:"description,omitempty"`
	// Human-readable label that describes the atlas resource policy.
	Name *string `json:"name,omitempty"`
	// List of policies that make up the atlas resource policy.
	Policies *[]ApiAtlasPolicyCreate `json:"policies,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiAtlasResourcePolicyEdit) MarshalJSON() ([]byte, error) {
	type noMethod ApiAtlasResourcePolicyEdit
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiAtlasResourcePolicyEdit instantiates a new ApiAtlasResourcePolicyEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasResourcePolicyEdit() *ApiAtlasResourcePolicyEdit {
	this := ApiAtlasResourcePolicyEdit{}
	return &this
}

// NewApiAtlasResourcePolicyEditWithDefaults instantiates a new ApiAtlasResourcePolicyEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasResourcePolicyEditWithDefaults() *ApiAtlasResourcePolicyEdit {
	this := ApiAtlasResourcePolicyEdit{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *ApiAtlasResourcePolicyEdit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyEdit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiAtlasResourcePolicyEdit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiAtlasResourcePolicyEdit) SetDescription(v string) {
	o.Description = &v
	o.NullFields = removeNullField(o.NullFields, "Description")
}

// SetDescriptionNil sets Description to an explicit JSON null when marshaled.
func (o *ApiAtlasResourcePolicyEdit) SetDescriptionNil() {
	o.Description = nil
	o.NullFields = addNullField(o.NullFields, "Description")
}

// GetName returns the Name field value if set, zero value otherwise
func (o *ApiAtlasResourcePolicyEdit) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyEdit) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAtlasResourcePolicyEdit) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAtlasResourcePolicyEdit) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *ApiAtlasResourcePolicyEdit) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}

// GetPolicies returns the Policies field value if set, zero value otherwise
func (o *ApiAtlasResourcePolicyEdit) GetPolicies() []ApiAtlasPolicyCreate {
	if o == nil || IsNil(o.Policies) {
		var ret []ApiAtlasPolicyCreate
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyEdit) GetPoliciesOk() (*[]ApiAtlasPolicyCreate, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}

	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *ApiAtlasResourcePolicyEdit) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []ApiAtlasPolicyCreate and assigns it to the Policies field.
func (o *ApiAtlasResourcePolicyEdit) SetPolicies(v []ApiAtlasPolicyCreate) {
	o.Policies = &v
	o.NullFields = removeNullField(o.NullFields, "Policies")
}

// SetPoliciesNil sets Policies to an explicit JSON null when marshaled.
func (o *ApiAtlasResourcePolicyEdit) SetPoliciesNil() {
	o.Policies = nil
	o.NullFields = addNullField(o.NullFields, "Policies")
}
