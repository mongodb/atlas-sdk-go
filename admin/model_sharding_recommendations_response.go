// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendationsResponse Sharding recommendations generated for one cluster.
type ShardingRecommendationsResponse struct {
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
func NewShardingRecommendationsResponse(recommendations []ShardingRecommendation) *ShardingRecommendationsResponse {
	this := ShardingRecommendationsResponse{}
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
