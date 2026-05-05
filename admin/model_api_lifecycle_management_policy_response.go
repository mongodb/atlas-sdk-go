// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiLifecycleManagementPolicyResponse Lifecycle management policy for automated data transition.
type ApiLifecycleManagementPolicyResponse struct {
	Action ApiLifecycleManagementAction `json:"action"`
	// Name of the database this policy applies to.
	// Read only field.
	DbName string `json:"dbName"`
	// Current effective state of the policy.
	// Read only field.
	EffectiveState string `json:"effectiveState"`
	// Unique identifier of the policy.
	// Read only field.
	Id       string                         `json:"id"`
	Schedule ApiLifecycleManagementSchedule `json:"schedule"`
	// Current state of the policy.
	// Read only field.
	State string `json:"state"`
	// Name of the target collection for data transition.
	// Read only field.
	TargetCollection string `json:"targetCollection"`
}

// NewApiLifecycleManagementPolicyResponse instantiates a new ApiLifecycleManagementPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLifecycleManagementPolicyResponse(action ApiLifecycleManagementAction, dbName string, effectiveState string, id string, schedule ApiLifecycleManagementSchedule, state string, targetCollection string) *ApiLifecycleManagementPolicyResponse {
	this := ApiLifecycleManagementPolicyResponse{}
	this.Action = action
	this.DbName = dbName
	this.EffectiveState = effectiveState
	this.Id = id
	this.Schedule = schedule
	this.State = state
	this.TargetCollection = targetCollection
	return &this
}

// NewApiLifecycleManagementPolicyResponseWithDefaults instantiates a new ApiLifecycleManagementPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLifecycleManagementPolicyResponseWithDefaults() *ApiLifecycleManagementPolicyResponse {
	this := ApiLifecycleManagementPolicyResponse{}
	return &this
}

// GetAction returns the Action field value
func (o *ApiLifecycleManagementPolicyResponse) GetAction() ApiLifecycleManagementAction {
	if o == nil {
		var ret ApiLifecycleManagementAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyResponse) GetActionOk() (*ApiLifecycleManagementAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ApiLifecycleManagementPolicyResponse) SetAction(v ApiLifecycleManagementAction) {
	o.Action = v
}

// GetDbName returns the DbName field value
func (o *ApiLifecycleManagementPolicyResponse) GetDbName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyResponse) GetDbNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value
func (o *ApiLifecycleManagementPolicyResponse) SetDbName(v string) {
	o.DbName = v
}

// GetEffectiveState returns the EffectiveState field value
func (o *ApiLifecycleManagementPolicyResponse) GetEffectiveState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectiveState
}

// GetEffectiveStateOk returns a tuple with the EffectiveState field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyResponse) GetEffectiveStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveState, true
}

// SetEffectiveState sets field value
func (o *ApiLifecycleManagementPolicyResponse) SetEffectiveState(v string) {
	o.EffectiveState = v
}

// GetId returns the Id field value
func (o *ApiLifecycleManagementPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiLifecycleManagementPolicyResponse) SetId(v string) {
	o.Id = v
}

// GetSchedule returns the Schedule field value
func (o *ApiLifecycleManagementPolicyResponse) GetSchedule() ApiLifecycleManagementSchedule {
	if o == nil {
		var ret ApiLifecycleManagementSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyResponse) GetScheduleOk() (*ApiLifecycleManagementSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *ApiLifecycleManagementPolicyResponse) SetSchedule(v ApiLifecycleManagementSchedule) {
	o.Schedule = v
}

// GetState returns the State field value
func (o *ApiLifecycleManagementPolicyResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ApiLifecycleManagementPolicyResponse) SetState(v string) {
	o.State = v
}

// GetTargetCollection returns the TargetCollection field value
func (o *ApiLifecycleManagementPolicyResponse) GetTargetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetCollection
}

// GetTargetCollectionOk returns a tuple with the TargetCollection field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyResponse) GetTargetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetCollection, true
}

// SetTargetCollection sets field value
func (o *ApiLifecycleManagementPolicyResponse) SetTargetCollection(v string) {
	o.TargetCollection = v
}
