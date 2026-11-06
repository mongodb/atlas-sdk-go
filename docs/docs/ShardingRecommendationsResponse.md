# ShardingRecommendationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recommendations** | [**[]ShardingRecommendation**](ShardingRecommendation.md) | Sharding recommendations for the cluster. Each recommendation identifies the recommendation type that produced it in its &#x60;type&#x60; property. | [readonly] 

## Methods

### NewShardingRecommendationsResponse

`func NewShardingRecommendationsResponse(recommendations []ShardingRecommendation, ) *ShardingRecommendationsResponse`

NewShardingRecommendationsResponse instantiates a new ShardingRecommendationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardingRecommendationsResponseWithDefaults

`func NewShardingRecommendationsResponseWithDefaults() *ShardingRecommendationsResponse`

NewShardingRecommendationsResponseWithDefaults instantiates a new ShardingRecommendationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecommendations

`func (o *ShardingRecommendationsResponse) GetRecommendations() []ShardingRecommendation`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *ShardingRecommendationsResponse) GetRecommendationsOk() (*[]ShardingRecommendation, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *ShardingRecommendationsResponse) SetRecommendations(v []ShardingRecommendation)`

SetRecommendations sets Recommendations field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


