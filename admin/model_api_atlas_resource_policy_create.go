// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasResourcePolicyCreate struct for ApiAtlasResourcePolicyCreate
type ApiAtlasResourcePolicyCreate struct {
	// Description of the atlas resource policy.
	Description *string `json:"description,omitempty"`
	// Human-readable label that describes the atlas resource policy.
	Name string `json:"name"`
	// List of policies that make up the atlas resource policy.
	Policies []ApiAtlasPolicyCreate `json:"policies"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiAtlasResourcePolicyCreate) MarshalJSON() ([]byte, error) {
	type noMethod ApiAtlasResourcePolicyCreate
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiAtlasResourcePolicyCreate instantiates a new ApiAtlasResourcePolicyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasResourcePolicyCreate(name string, policies []ApiAtlasPolicyCreate) *ApiAtlasResourcePolicyCreate {
	this := ApiAtlasResourcePolicyCreate{}
	this.Name = name
	this.Policies = policies
	return &this
}

// NewApiAtlasResourcePolicyCreateWithDefaults instantiates a new ApiAtlasResourcePolicyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasResourcePolicyCreateWithDefaults() *ApiAtlasResourcePolicyCreate {
	this := ApiAtlasResourcePolicyCreate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *ApiAtlasResourcePolicyCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiAtlasResourcePolicyCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiAtlasResourcePolicyCreate) SetDescription(v string) {
	o.Description = &v
	o.NullFields = removeNullField(o.NullFields, "Description")
}

// SetDescriptionNil sets Description to an explicit JSON null when marshaled.
func (o *ApiAtlasResourcePolicyCreate) SetDescriptionNil() {
	o.Description = nil
	o.NullFields = addNullField(o.NullFields, "Description")
}

// GetName returns the Name field value
func (o *ApiAtlasResourcePolicyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiAtlasResourcePolicyCreate) SetName(v string) {
	o.Name = v
}

// GetPolicies returns the Policies field value
func (o *ApiAtlasResourcePolicyCreate) GetPolicies() []ApiAtlasPolicyCreate {
	if o == nil {
		var ret []ApiAtlasPolicyCreate
		return ret
	}

	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasResourcePolicyCreate) GetPoliciesOk() (*[]ApiAtlasPolicyCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policies, true
}

// SetPolicies sets field value
func (o *ApiAtlasResourcePolicyCreate) SetPolicies(v []ApiAtlasPolicyCreate) {
	o.Policies = v
}
