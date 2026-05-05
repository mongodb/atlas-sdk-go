// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiLifecycleManagementPolicyCreateRequest Request body for creating a lifecycle management policy.
type ApiLifecycleManagementPolicyCreateRequest struct {
	Action ApiLifecycleManagementAction `json:"action"`
	// Name of the database this policy applies to.
	DbName   string                          `json:"dbName"`
	Schedule *ApiLifecycleManagementSchedule `json:"schedule,omitempty"`
	// Name of the source collection for data transition.
	TargetCollection string `json:"targetCollection"`
}

// NewApiLifecycleManagementPolicyCreateRequest instantiates a new ApiLifecycleManagementPolicyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLifecycleManagementPolicyCreateRequest(action ApiLifecycleManagementAction, dbName string, targetCollection string) *ApiLifecycleManagementPolicyCreateRequest {
	this := ApiLifecycleManagementPolicyCreateRequest{}
	this.Action = action
	this.DbName = dbName
	this.TargetCollection = targetCollection
	return &this
}

// NewApiLifecycleManagementPolicyCreateRequestWithDefaults instantiates a new ApiLifecycleManagementPolicyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLifecycleManagementPolicyCreateRequestWithDefaults() *ApiLifecycleManagementPolicyCreateRequest {
	this := ApiLifecycleManagementPolicyCreateRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *ApiLifecycleManagementPolicyCreateRequest) GetAction() ApiLifecycleManagementAction {
	if o == nil {
		var ret ApiLifecycleManagementAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyCreateRequest) GetActionOk() (*ApiLifecycleManagementAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ApiLifecycleManagementPolicyCreateRequest) SetAction(v ApiLifecycleManagementAction) {
	o.Action = v
}

// GetDbName returns the DbName field value
func (o *ApiLifecycleManagementPolicyCreateRequest) GetDbName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyCreateRequest) GetDbNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value
func (o *ApiLifecycleManagementPolicyCreateRequest) SetDbName(v string) {
	o.DbName = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise
func (o *ApiLifecycleManagementPolicyCreateRequest) GetSchedule() ApiLifecycleManagementSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret ApiLifecycleManagementSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyCreateRequest) GetScheduleOk() (*ApiLifecycleManagementSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}

	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ApiLifecycleManagementPolicyCreateRequest) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ApiLifecycleManagementSchedule and assigns it to the Schedule field.
func (o *ApiLifecycleManagementPolicyCreateRequest) SetSchedule(v ApiLifecycleManagementSchedule) {
	o.Schedule = &v
}

// GetTargetCollection returns the TargetCollection field value
func (o *ApiLifecycleManagementPolicyCreateRequest) GetTargetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetCollection
}

// GetTargetCollectionOk returns a tuple with the TargetCollection field value
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementPolicyCreateRequest) GetTargetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetCollection, true
}

// SetTargetCollection sets field value
func (o *ApiLifecycleManagementPolicyCreateRequest) SetTargetCollection(v string) {
	o.TargetCollection = v
}
