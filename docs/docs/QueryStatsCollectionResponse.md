# QueryStatsCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Method used to collect &#x60;queryStats&#x60; entries on the cluster&#39;s processes. | [readonly] 
**SampleRate** | Pointer to **float64** | Fraction of operations recorded. Dividing reported counts by this rate yields an unbiased estimate of the true totals. | [optional] [readonly] 
**RateLimitPerSecond** | Pointer to **int** | Maximum read queries recorded per second. | [optional] [readonly] 

## Methods

### NewQueryStatsCollectionResponse

`func NewQueryStatsCollectionResponse(mode string, ) *QueryStatsCollectionResponse`

NewQueryStatsCollectionResponse instantiates a new QueryStatsCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryStatsCollectionResponseWithDefaults

`func NewQueryStatsCollectionResponseWithDefaults() *QueryStatsCollectionResponse`

NewQueryStatsCollectionResponseWithDefaults instantiates a new QueryStatsCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *QueryStatsCollectionResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *QueryStatsCollectionResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *QueryStatsCollectionResponse) SetMode(v string)`

SetMode sets Mode field to given value.

### GetSampleRate

`func (o *QueryStatsCollectionResponse) GetSampleRate() float64`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *QueryStatsCollectionResponse) GetSampleRateOk() (*float64, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *QueryStatsCollectionResponse) SetSampleRate(v float64)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *QueryStatsCollectionResponse) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### SetSampleRateNil

`func (o *QueryStatsCollectionResponse) SetSampleRateNil()`

SetSampleRateNil sets SampleRate to an explicit JSON null when marshaled, overriding any value previously set with SetSampleRate. Calling SetSampleRate again clears the null override.

### GetRateLimitPerSecond

`func (o *QueryStatsCollectionResponse) GetRateLimitPerSecond() int`

GetRateLimitPerSecond returns the RateLimitPerSecond field if non-nil, zero value otherwise.

### GetRateLimitPerSecondOk

`func (o *QueryStatsCollectionResponse) GetRateLimitPerSecondOk() (*int, bool)`

GetRateLimitPerSecondOk returns a tuple with the RateLimitPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerSecond

`func (o *QueryStatsCollectionResponse) SetRateLimitPerSecond(v int)`

SetRateLimitPerSecond sets RateLimitPerSecond field to given value.

### HasRateLimitPerSecond

`func (o *QueryStatsCollectionResponse) HasRateLimitPerSecond() bool`

HasRateLimitPerSecond returns a boolean if a field has been set.

### SetRateLimitPerSecondNil

`func (o *QueryStatsCollectionResponse) SetRateLimitPerSecondNil()`

SetRateLimitPerSecondNil sets RateLimitPerSecond to an explicit JSON null when marshaled, overriding any value previously set with SetRateLimitPerSecond. Calling SetRateLimitPerSecond again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


