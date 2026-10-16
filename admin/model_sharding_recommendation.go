// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendation One sharding recommendation for the cluster or one of its collections.
type ShardingRecommendation struct {
	// Monitored conditions that fired and produced this recommendation.
	// Read only field.
	Triggers []ShardingRecommendationTrigger `json:"triggers"`
	// Recommendation type that produced this recommendation.
	// Read only field.
	Type string `json:"type"`
	// Recommended total number of shards for the cluster.
	// Read only field.
	RecommendedShardCount *int `json:"recommendedShardCount,omitempty"`
	// Namespace of the collection to shard, in `<database>.<collection>` form.
	// Read only field.
	Namespace *string `json:"namespace,omitempty"`
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
func NewShardingRecommendation(triggers []ShardingRecommendationTrigger, type_ string) *ShardingRecommendation {
	this := ShardingRecommendation{}
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

// GetTriggers returns the Triggers field value
func (o *ShardingRecommendation) GetTriggers() []ShardingRecommendationTrigger {
	if o == nil {
		var ret []ShardingRecommendationTrigger
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendation) GetTriggersOk() (*[]ShardingRecommendationTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Triggers, true
}

// SetTriggers sets field value
func (o *ShardingRecommendation) SetTriggers(v []ShardingRecommendationTrigger) {
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

// GetRecommendedShardCount returns the RecommendedShardCount field value if set, zero value otherwise
func (o *ShardingRecommendation) GetRecommendedShardCount() int {
	if o == nil || IsNil(o.RecommendedShardCount) {
		var ret int
		return ret
	}
	return *o.RecommendedShardCount
}

// GetRecommendedShardCountOk returns a tuple with the RecommendedShardCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendation) GetRecommendedShardCountOk() (*int, bool) {
	if o == nil || IsNil(o.RecommendedShardCount) {
		return nil, false
	}

	return o.RecommendedShardCount, true
}

// HasRecommendedShardCount returns a boolean if a field has been set.
func (o *ShardingRecommendation) HasRecommendedShardCount() bool {
	if o != nil && !IsNil(o.RecommendedShardCount) {
		return true
	}

	return false
}

// SetRecommendedShardCount gets a reference to the given int and assigns it to the RecommendedShardCount field.
func (o *ShardingRecommendation) SetRecommendedShardCount(v int) {
	o.RecommendedShardCount = &v
	o.NullFields = removeNullField(o.NullFields, "RecommendedShardCount")
}

// SetRecommendedShardCountNil sets RecommendedShardCount to an explicit JSON null when marshaled.
func (o *ShardingRecommendation) SetRecommendedShardCountNil() {
	o.RecommendedShardCount = nil
	o.NullFields = addNullField(o.NullFields, "RecommendedShardCount")
}

// GetNamespace returns the Namespace field value if set, zero value otherwise
func (o *ShardingRecommendation) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendation) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}

	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ShardingRecommendation) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ShardingRecommendation) SetNamespace(v string) {
	o.Namespace = &v
	o.NullFields = removeNullField(o.NullFields, "Namespace")
}

// SetNamespaceNil sets Namespace to an explicit JSON null when marshaled.
func (o *ShardingRecommendation) SetNamespaceNil() {
	o.Namespace = nil
	o.NullFields = addNullField(o.NullFields, "Namespace")
}
