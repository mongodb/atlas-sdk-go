// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingInsightsBalancingWindow Balancing window configured on the sharded cluster. When configured, the balancer moves data only during the daily interval between `startTime` and `endTime`.
type ShardingInsightsBalancingWindow struct {
	// Time of day at which the balancing window ends, in 24-hour `HH:MM` format as configured on the cluster. Present only when `status` is `CONFIGURED`.
	// Read only field.
	EndTime *string `json:"endTime,omitempty"`
	// Time of day at which the balancing window starts, in 24-hour `HH:MM` format as configured on the cluster. Present only when `status` is `CONFIGURED`.
	// Read only field.
	StartTime *string `json:"startTime,omitempty"`
	// State of the balancing window. `CONFIGURED` means the cluster restricts balancing to the interval between `startTime` and `endTime`. `NOT_CONFIGURED` means the balancer may run at any time. `FAILED_INITIAL_METRIC_PENDING` means MongoDB Cloud has not yet collected balancer metrics for a recently created or converted cluster. `FAILED_TO_DETERMINE` means the window could not be resolved.
	// Read only field.
	Status string `json:"status"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingInsightsBalancingWindow) MarshalJSON() ([]byte, error) {
	type noMethod ShardingInsightsBalancingWindow
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingInsightsBalancingWindow instantiates a new ShardingInsightsBalancingWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingInsightsBalancingWindow(status string) *ShardingInsightsBalancingWindow {
	this := ShardingInsightsBalancingWindow{}
	this.Status = status
	return &this
}

// NewShardingInsightsBalancingWindowWithDefaults instantiates a new ShardingInsightsBalancingWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingInsightsBalancingWindowWithDefaults() *ShardingInsightsBalancingWindow {
	this := ShardingInsightsBalancingWindow{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise
func (o *ShardingInsightsBalancingWindow) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingInsightsBalancingWindow) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}

	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ShardingInsightsBalancingWindow) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *ShardingInsightsBalancingWindow) SetEndTime(v string) {
	o.EndTime = &v
	o.NullFields = removeNullField(o.NullFields, "EndTime")
}

// SetEndTimeNil sets EndTime to an explicit JSON null when marshaled.
func (o *ShardingInsightsBalancingWindow) SetEndTimeNil() {
	o.EndTime = nil
	o.NullFields = addNullField(o.NullFields, "EndTime")
}

// GetStartTime returns the StartTime field value if set, zero value otherwise
func (o *ShardingInsightsBalancingWindow) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingInsightsBalancingWindow) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}

	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ShardingInsightsBalancingWindow) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ShardingInsightsBalancingWindow) SetStartTime(v string) {
	o.StartTime = &v
	o.NullFields = removeNullField(o.NullFields, "StartTime")
}

// SetStartTimeNil sets StartTime to an explicit JSON null when marshaled.
func (o *ShardingInsightsBalancingWindow) SetStartTimeNil() {
	o.StartTime = nil
	o.NullFields = addNullField(o.NullFields, "StartTime")
}

// GetStatus returns the Status field value
func (o *ShardingInsightsBalancingWindow) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ShardingInsightsBalancingWindow) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ShardingInsightsBalancingWindow) SetStatus(v string) {
	o.Status = v
}
