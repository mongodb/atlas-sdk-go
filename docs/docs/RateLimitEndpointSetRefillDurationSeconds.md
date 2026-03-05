# RateLimitEndpointSetRefillDurationSeconds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **int64** | The default rate limit refill duration, in seconds, of the endpoint set. Returned if there is a rate limit refill duration override set for the requested entity. | [optional] 
**Value** | Pointer to **int64** | The applied rate limit refill duration of the endpoint set. | [optional] 

## Methods

### NewRateLimitEndpointSetRefillDurationSeconds

`func NewRateLimitEndpointSetRefillDurationSeconds() *RateLimitEndpointSetRefillDurationSeconds`

NewRateLimitEndpointSetRefillDurationSeconds instantiates a new RateLimitEndpointSetRefillDurationSeconds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitEndpointSetRefillDurationSecondsWithDefaults

`func NewRateLimitEndpointSetRefillDurationSecondsWithDefaults() *RateLimitEndpointSetRefillDurationSeconds`

NewRateLimitEndpointSetRefillDurationSecondsWithDefaults instantiates a new RateLimitEndpointSetRefillDurationSeconds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *RateLimitEndpointSetRefillDurationSeconds) GetDefaultValue() int64`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *RateLimitEndpointSetRefillDurationSeconds) GetDefaultValueOk() (*int64, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *RateLimitEndpointSetRefillDurationSeconds) SetDefaultValue(v int64)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *RateLimitEndpointSetRefillDurationSeconds) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.
### GetValue

`func (o *RateLimitEndpointSetRefillDurationSeconds) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RateLimitEndpointSetRefillDurationSeconds) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RateLimitEndpointSetRefillDurationSeconds) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *RateLimitEndpointSetRefillDurationSeconds) HasValue() bool`

HasValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


