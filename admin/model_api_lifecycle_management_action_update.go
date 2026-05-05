// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiLifecycleManagementActionUpdate Partial update for an action. Only provided fields are applied.
type ApiLifecycleManagementActionUpdate struct {
	// Type of action. Must match existing action type.
	Type     *string                         `json:"type,omitempty"`
	Criteria *ApiLifecycleManagementCriteria `json:"criteria,omitempty"`
	// Name of the destination collection. If omitted, existing value is preserved.
	DestinationCollName *string `json:"destinationCollName,omitempty"`
}

// NewApiLifecycleManagementActionUpdate instantiates a new ApiLifecycleManagementActionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLifecycleManagementActionUpdate() *ApiLifecycleManagementActionUpdate {
	this := ApiLifecycleManagementActionUpdate{}
	return &this
}

// NewApiLifecycleManagementActionUpdateWithDefaults instantiates a new ApiLifecycleManagementActionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLifecycleManagementActionUpdateWithDefaults() *ApiLifecycleManagementActionUpdate {
	this := ApiLifecycleManagementActionUpdate{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise
func (o *ApiLifecycleManagementActionUpdate) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementActionUpdate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiLifecycleManagementActionUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiLifecycleManagementActionUpdate) SetType(v string) {
	o.Type = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise
func (o *ApiLifecycleManagementActionUpdate) GetCriteria() ApiLifecycleManagementCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret ApiLifecycleManagementCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementActionUpdate) GetCriteriaOk() (*ApiLifecycleManagementCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}

	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ApiLifecycleManagementActionUpdate) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ApiLifecycleManagementCriteria and assigns it to the Criteria field.
func (o *ApiLifecycleManagementActionUpdate) SetCriteria(v ApiLifecycleManagementCriteria) {
	o.Criteria = &v
}

// GetDestinationCollName returns the DestinationCollName field value if set, zero value otherwise
func (o *ApiLifecycleManagementActionUpdate) GetDestinationCollName() string {
	if o == nil || IsNil(o.DestinationCollName) {
		var ret string
		return ret
	}
	return *o.DestinationCollName
}

// GetDestinationCollNameOk returns a tuple with the DestinationCollName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementActionUpdate) GetDestinationCollNameOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationCollName) {
		return nil, false
	}

	return o.DestinationCollName, true
}

// HasDestinationCollName returns a boolean if a field has been set.
func (o *ApiLifecycleManagementActionUpdate) HasDestinationCollName() bool {
	if o != nil && !IsNil(o.DestinationCollName) {
		return true
	}

	return false
}

// SetDestinationCollName gets a reference to the given string and assigns it to the DestinationCollName field.
func (o *ApiLifecycleManagementActionUpdate) SetDestinationCollName(v string) {
	o.DestinationCollName = &v
}
