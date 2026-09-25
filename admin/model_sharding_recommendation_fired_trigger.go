// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendationFiredTrigger One monitored condition that fired and contributed to the recommendation.
type ShardingRecommendationFiredTrigger struct {
	// Identifier of the monitored condition that fired, in upper snake case. Conditions are added over time, so treat this as an open set and tolerate identifiers that your client does not recognize.
	// Read only field.
	Name string `json:"name"`
	// Entities for which the trigger fired.
	// Read only field.
	Targets []ShardingRecommendationTriggerTarget `json:"targets"`
	// Unit of the measured trigger values.
	// Read only field.
	Unit string `json:"unit"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingRecommendationFiredTrigger) MarshalJSON() ([]byte, error) {
	type noMethod ShardingRecommendationFiredTrigger
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingRecommendationFiredTrigger instantiates a new ShardingRecommendationFiredTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingRecommendationFiredTrigger(name string, targets []ShardingRecommendationTriggerTarget, unit string) *ShardingRecommendationFiredTrigger {
	this := ShardingRecommendationFiredTrigger{}
	this.Name = name
	this.Targets = targets
	this.Unit = unit
	return &this
}

// NewShardingRecommendationFiredTriggerWithDefaults instantiates a new ShardingRecommendationFiredTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingRecommendationFiredTriggerWithDefaults() *ShardingRecommendationFiredTrigger {
	this := ShardingRecommendationFiredTrigger{}
	return &this
}

// GetName returns the Name field value
func (o *ShardingRecommendationFiredTrigger) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationFiredTrigger) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShardingRecommendationFiredTrigger) SetName(v string) {
	o.Name = v
}

// GetTargets returns the Targets field value
func (o *ShardingRecommendationFiredTrigger) GetTargets() []ShardingRecommendationTriggerTarget {
	if o == nil {
		var ret []ShardingRecommendationTriggerTarget
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationFiredTrigger) GetTargetsOk() (*[]ShardingRecommendationTriggerTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value
func (o *ShardingRecommendationFiredTrigger) SetTargets(v []ShardingRecommendationTriggerTarget) {
	o.Targets = v
}

// GetUnit returns the Unit field value
func (o *ShardingRecommendationFiredTrigger) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationFiredTrigger) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ShardingRecommendationFiredTrigger) SetUnit(v string) {
	o.Unit = v
}
