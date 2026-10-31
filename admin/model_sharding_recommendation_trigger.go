// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendationTrigger One monitored condition that fired, observed on one entity.
type ShardingRecommendationTrigger struct {
	// Identifier of the monitored condition that fired.
	// Read only field.
	Name string `json:"name"`
	// Entity the condition was observed on: a namespace in `<database>.<collection>` form, a host name, or a replica set (shard) name, as identified by `targetType`. Present if and only if `targetType` is present; both are absent for cluster-wide conditions, such as `UNSHARDED_COLLECTION_COUNT`, which are not tied to a single entity. The cluster is already identified in the request path.
	// Read only field.
	Target *string `json:"target,omitempty"`
	// Kind of entity that `target` identifies. Present if and only if `target` is present; both are absent for cluster-wide conditions.
	// Read only field.
	TargetType *string `json:"targetType,omitempty"`
	// Unit of the measured trigger value.
	// Read only field.
	Unit string `json:"unit"`
	// Measured value of the trigger metric for this entity. This parameter is absent when no finite measurement was available.
	// Read only field.
	Value *float64 `json:"value,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingRecommendationTrigger) MarshalJSON() ([]byte, error) {
	type noMethod ShardingRecommendationTrigger
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingRecommendationTrigger instantiates a new ShardingRecommendationTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingRecommendationTrigger(name string, unit string) *ShardingRecommendationTrigger {
	this := ShardingRecommendationTrigger{}
	this.Name = name
	this.Unit = unit
	return &this
}

// NewShardingRecommendationTriggerWithDefaults instantiates a new ShardingRecommendationTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingRecommendationTriggerWithDefaults() *ShardingRecommendationTrigger {
	this := ShardingRecommendationTrigger{}
	return &this
}

// GetName returns the Name field value
func (o *ShardingRecommendationTrigger) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTrigger) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShardingRecommendationTrigger) SetName(v string) {
	o.Name = v
}

// GetTarget returns the Target field value if set, zero value otherwise
func (o *ShardingRecommendationTrigger) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTrigger) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}

	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ShardingRecommendationTrigger) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ShardingRecommendationTrigger) SetTarget(v string) {
	o.Target = &v
	o.NullFields = removeNullField(o.NullFields, "Target")
}

// SetTargetNil sets Target to an explicit JSON null when marshaled.
func (o *ShardingRecommendationTrigger) SetTargetNil() {
	o.Target = nil
	o.NullFields = addNullField(o.NullFields, "Target")
}

// GetTargetType returns the TargetType field value if set, zero value otherwise
func (o *ShardingRecommendationTrigger) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTrigger) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}

	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *ShardingRecommendationTrigger) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *ShardingRecommendationTrigger) SetTargetType(v string) {
	o.TargetType = &v
	o.NullFields = removeNullField(o.NullFields, "TargetType")
}

// SetTargetTypeNil sets TargetType to an explicit JSON null when marshaled.
func (o *ShardingRecommendationTrigger) SetTargetTypeNil() {
	o.TargetType = nil
	o.NullFields = addNullField(o.NullFields, "TargetType")
}

// GetUnit returns the Unit field value
func (o *ShardingRecommendationTrigger) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTrigger) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ShardingRecommendationTrigger) SetUnit(v string) {
	o.Unit = v
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *ShardingRecommendationTrigger) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTrigger) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ShardingRecommendationTrigger) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *ShardingRecommendationTrigger) SetValue(v float64) {
	o.Value = &v
	o.NullFields = removeNullField(o.NullFields, "Value")
}

// SetValueNil sets Value to an explicit JSON null when marshaled.
func (o *ShardingRecommendationTrigger) SetValueNil() {
	o.Value = nil
	o.NullFields = addNullField(o.NullFields, "Value")
}
