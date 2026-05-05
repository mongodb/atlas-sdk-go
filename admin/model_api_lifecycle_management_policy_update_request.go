// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiLifecycleManagementPolicyUpdateRequest Request body for updating a lifecycle management policy.
type ApiLifecycleManagementPolicyUpdateRequest struct {
	Action   *ApiLifecycleManagementActionUpdate `json:"action,omitempty"`
	Schedule *ApiLifecycleManagementSchedule     `json:"schedule,omitempty"`
	// Requested state transition. Valid values: ACTIVE, PAUSED.
	State *string `json:"state,omitempty"`
}

// NewApiLifecycleManagementPolicyUpdateRequest instantiates a new ApiLifecycleManagementPolicyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLifecycleManagementPolicyUpdateRequest() *ApiLifecycleManagementPolicyUpdateRequest {
	this := ApiLifecycleManagementPolicyUpdateRequest{}
	return &this
}

// NewApiLifecycleManagementPolicyUpdateRequestWithDefaults instantiates a new ApiLifecycleManagementPolicyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLifecycleManagementPolicyUpdateRequestWithDefaults() *ApiLifecycleManagementPolicyUpdateRequest {
	this := ApiLifecycleManagementPolicyUpdateRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise
func (o *ApiLifecycleManagementPolicyUpdateRequest) GetAction() ApiLifecycleManagementActionUpdate {
	if o == nil || IsNil(o.Action) {
		var ret ApiLifecycleManagementActionUpdate
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyUpdateRequest) GetActionOk() (*ApiLifecycleManagementActionUpdate, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}

	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ApiLifecycleManagementPolicyUpdateRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given ApiLifecycleManagementActionUpdate and assigns it to the Action field.
func (o *ApiLifecycleManagementPolicyUpdateRequest) SetAction(v ApiLifecycleManagementActionUpdate) {
	o.Action = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise
func (o *ApiLifecycleManagementPolicyUpdateRequest) GetSchedule() ApiLifecycleManagementSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret ApiLifecycleManagementSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyUpdateRequest) GetScheduleOk() (*ApiLifecycleManagementSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}

	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ApiLifecycleManagementPolicyUpdateRequest) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ApiLifecycleManagementSchedule and assigns it to the Schedule field.
func (o *ApiLifecycleManagementPolicyUpdateRequest) SetSchedule(v ApiLifecycleManagementSchedule) {
	o.Schedule = &v
}

// GetState returns the State field value if set, zero value otherwise
func (o *ApiLifecycleManagementPolicyUpdateRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyUpdateRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiLifecycleManagementPolicyUpdateRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiLifecycleManagementPolicyUpdateRequest) SetState(v string) {
	o.State = &v
}
