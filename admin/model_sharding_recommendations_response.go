// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendationsResponse Sharding recommendations generated for one cluster.
type ShardingRecommendationsResponse struct {
	// Recommendation types that were evaluated for the cluster. A type that appears here but has no entry in `recommendations` was evaluated and produced nothing to recommend. A type is absent when it does not apply to the cluster topology, was excluded by the `recommendationTypes` query parameter, or failed to generate (see `failures`).
	// Read only field.
	EvaluatedTypes []string `json:"evaluatedTypes"`
	// Recommendation types whose generation failed. Empty when every evaluated recommendation type completed.
	// Read only field.
	Failures []ShardingRecommendationFailure `json:"failures"`
	// Sharding recommendations for the cluster. Each recommendation identifies the recommendation type that produced it in its `type` property.
	// Read only field.
	Recommendations []ShardingRecommendation `json:"recommendations"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingRecommendationsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ShardingRecommendationsResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingRecommendationsResponse instantiates a new ShardingRecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingRecommendationsResponse(evaluatedTypes []string, failures []ShardingRecommendationFailure, recommendations []ShardingRecommendation) *ShardingRecommendationsResponse {
	this := ShardingRecommendationsResponse{}
	this.EvaluatedTypes = evaluatedTypes
	this.Failures = failures
	this.Recommendations = recommendations
	return &this
}

// NewShardingRecommendationsResponseWithDefaults instantiates a new ShardingRecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingRecommendationsResponseWithDefaults() *ShardingRecommendationsResponse {
	this := ShardingRecommendationsResponse{}
	return &this
}

// GetEvaluatedTypes returns the EvaluatedTypes field value
func (o *ShardingRecommendationsResponse) GetEvaluatedTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EvaluatedTypes
}

// GetEvaluatedTypesOk returns a tuple with the EvaluatedTypes field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationsResponse) GetEvaluatedTypesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluatedTypes, true
}

// SetEvaluatedTypes sets field value
func (o *ShardingRecommendationsResponse) SetEvaluatedTypes(v []string) {
	o.EvaluatedTypes = v
}

// GetFailures returns the Failures field value
func (o *ShardingRecommendationsResponse) GetFailures() []ShardingRecommendationFailure {
	if o == nil {
		var ret []ShardingRecommendationFailure
		return ret
	}

	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationsResponse) GetFailuresOk() (*[]ShardingRecommendationFailure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failures, true
}

// SetFailures sets field value
func (o *ShardingRecommendationsResponse) SetFailures(v []ShardingRecommendationFailure) {
	o.Failures = v
}

// GetRecommendations returns the Recommendations field value
func (o *ShardingRecommendationsResponse) GetRecommendations() []ShardingRecommendation {
	if o == nil {
		var ret []ShardingRecommendation
		return ret
	}

	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationsResponse) GetRecommendationsOk() (*[]ShardingRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommendations, true
}

// SetRecommendations sets field value
func (o *ShardingRecommendationsResponse) SetRecommendations(v []ShardingRecommendation) {
	o.Recommendations = v
}
