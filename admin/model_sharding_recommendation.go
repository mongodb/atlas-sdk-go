// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendation One sharding recommendation for the cluster or one of its collections.
type ShardingRecommendation struct {
	// Level at which the recommendation applies.
	// Read only field.
	Scope   string                         `json:"scope"`
	Summary *ShardingRecommendationSummary `json:"summary,omitempty"`
	// Entity the recommendation applies to: the cluster name when scope is `CLUSTER`, or the namespace in `database.collection` form when scope is `COLLECTION`.
	// Read only field.
	Target string `json:"target"`
	// Monitored conditions that fired and produced this recommendation.
	// Read only field.
	Triggers []ShardingRecommendationFiredTrigger `json:"triggers"`
	// Recommendation type that produced this recommendation.
	// Read only field.
	Type string `json:"type"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingRecommendation) MarshalJSON() ([]byte, error) {
	type noMethod ShardingRecommendation
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingRecommendation instantiates a new ShardingRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingRecommendation(scope string, target string, triggers []ShardingRecommendationFiredTrigger, type_ string) *ShardingRecommendation {
	this := ShardingRecommendation{}
	this.Scope = scope
	this.Target = target
	this.Triggers = triggers
	this.Type = type_
	return &this
}

// NewShardingRecommendationWithDefaults instantiates a new ShardingRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingRecommendationWithDefaults() *ShardingRecommendation {
	this := ShardingRecommendation{}
	return &this
}

// GetScope returns the Scope field value
func (o *ShardingRecommendation) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendation) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *ShardingRecommendation) SetScope(v string) {
	o.Scope = v
}

// GetSummary returns the Summary field value if set, zero value otherwise
func (o *ShardingRecommendation) GetSummary() ShardingRecommendationSummary {
	if o == nil || IsNil(o.Summary) {
		var ret ShardingRecommendationSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendation) GetSummaryOk() (*ShardingRecommendationSummary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}

	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ShardingRecommendation) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given ShardingRecommendationSummary and assigns it to the Summary field.
func (o *ShardingRecommendation) SetSummary(v ShardingRecommendationSummary) {
	o.Summary = &v
	o.NullFields = removeNullField(o.NullFields, "Summary")
}

// SetSummaryNil sets Summary to an explicit JSON null when marshaled.
func (o *ShardingRecommendation) SetSummaryNil() {
	o.Summary = nil
	o.NullFields = addNullField(o.NullFields, "Summary")
}

// GetTarget returns the Target field value
func (o *ShardingRecommendation) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendation) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *ShardingRecommendation) SetTarget(v string) {
	o.Target = v
}

// GetTriggers returns the Triggers field value
func (o *ShardingRecommendation) GetTriggers() []ShardingRecommendationFiredTrigger {
	if o == nil {
		var ret []ShardingRecommendationFiredTrigger
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendation) GetTriggersOk() (*[]ShardingRecommendationFiredTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Triggers, true
}

// SetTriggers sets field value
func (o *ShardingRecommendation) SetTriggers(v []ShardingRecommendationFiredTrigger) {
	o.Triggers = v
}

// GetType returns the Type field value
func (o *ShardingRecommendation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShardingRecommendation) SetType(v string) {
	o.Type = v
}
