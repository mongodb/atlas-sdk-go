// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiLifecycleManagementCriteria Criteria that determines which data is eligible for the action.
type ApiLifecycleManagementCriteria struct {
	// Type of criteria.
	Type *string `json:"type,omitempty"`
	// Name of the date field used to evaluate transition eligibility.
	DateField *string `json:"dateField,omitempty"`
	// Number of days after the date field value when data becomes eligible for transition.
	TransitionAfterDays *int `json:"transitionAfterDays,omitempty"`
	// MongoDB find() based filter string in JSON format to determine data transition eligibility.
	Query *string `json:"query,omitempty"`
}

// NewApiLifecycleManagementCriteria instantiates a new ApiLifecycleManagementCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLifecycleManagementCriteria() *ApiLifecycleManagementCriteria {
	this := ApiLifecycleManagementCriteria{}
	return &this
}

// NewApiLifecycleManagementCriteriaWithDefaults instantiates a new ApiLifecycleManagementCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLifecycleManagementCriteriaWithDefaults() *ApiLifecycleManagementCriteria {
	this := ApiLifecycleManagementCriteria{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise
func (o *ApiLifecycleManagementCriteria) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementCriteria) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiLifecycleManagementCriteria) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiLifecycleManagementCriteria) SetType(v string) {
	o.Type = &v
}

// GetDateField returns the DateField field value if set, zero value otherwise
func (o *ApiLifecycleManagementCriteria) GetDateField() string {
	if o == nil || IsNil(o.DateField) {
		var ret string
		return ret
	}
	return *o.DateField
}

// GetDateFieldOk returns a tuple with the DateField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementCriteria) GetDateFieldOk() (*string, bool) {
	if o == nil || IsNil(o.DateField) {
		return nil, false
	}

	return o.DateField, true
}

// HasDateField returns a boolean if a field has been set.
func (o *ApiLifecycleManagementCriteria) HasDateField() bool {
	if o != nil && !IsNil(o.DateField) {
		return true
	}

	return false
}

// SetDateField gets a reference to the given string and assigns it to the DateField field.
func (o *ApiLifecycleManagementCriteria) SetDateField(v string) {
	o.DateField = &v
}

// GetTransitionAfterDays returns the TransitionAfterDays field value if set, zero value otherwise
func (o *ApiLifecycleManagementCriteria) GetTransitionAfterDays() int {
	if o == nil || IsNil(o.TransitionAfterDays) {
		var ret int
		return ret
	}
	return *o.TransitionAfterDays
}

// GetTransitionAfterDaysOk returns a tuple with the TransitionAfterDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementCriteria) GetTransitionAfterDaysOk() (*int, bool) {
	if o == nil || IsNil(o.TransitionAfterDays) {
		return nil, false
	}

	return o.TransitionAfterDays, true
}

// HasTransitionAfterDays returns a boolean if a field has been set.
func (o *ApiLifecycleManagementCriteria) HasTransitionAfterDays() bool {
	if o != nil && !IsNil(o.TransitionAfterDays) {
		return true
	}

	return false
}

// SetTransitionAfterDays gets a reference to the given int and assigns it to the TransitionAfterDays field.
func (o *ApiLifecycleManagementCriteria) SetTransitionAfterDays(v int) {
	o.TransitionAfterDays = &v
}

// GetQuery returns the Query field value if set, zero value otherwise
func (o *ApiLifecycleManagementCriteria) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementCriteria) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}

	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ApiLifecycleManagementCriteria) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ApiLifecycleManagementCriteria) SetQuery(v string) {
	o.Query = &v
}
