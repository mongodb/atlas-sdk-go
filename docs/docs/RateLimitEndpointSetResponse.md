# RateLimitEndpointSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to [**RateLimitEndpointSetCapacity**](RateLimitEndpointSetCapacity.md) |  | [optional] 
**EndpointSetId** | Pointer to **string** | The ID of the endpoint set. | [optional] 
**EndpointSetName** | Pointer to **string** | The endpoint set name. | [optional] 
**Endpoints** | Pointer to [**[]RateLimitEndpointSetEndpoint**](RateLimitEndpointSetEndpoint.md) | A list of endpoints associated with the specified endpoint set. | [optional] 
**RefillDurationSeconds** | Pointer to [**RateLimitEndpointSetRefillDurationSeconds**](RateLimitEndpointSetRefillDurationSeconds.md) |  | [optional] 
**RefillRate** | Pointer to [**RateLimitEndpointSetRefillRate**](RateLimitEndpointSetRefillRate.md) |  | [optional] 
**Scope** | Pointer to **string** | The scope of the endpoint set. | [optional] 

## Methods

### NewRateLimitEndpointSetResponse

`func NewRateLimitEndpointSetResponse() *RateLimitEndpointSetResponse`

NewRateLimitEndpointSetResponse instantiates a new RateLimitEndpointSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitEndpointSetResponseWithDefaults

`func NewRateLimitEndpointSetResponseWithDefaults() *RateLimitEndpointSetResponse`

NewRateLimitEndpointSetResponseWithDefaults instantiates a new RateLimitEndpointSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *RateLimitEndpointSetResponse) GetCapacity() RateLimitEndpointSetCapacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *RateLimitEndpointSetResponse) GetCapacityOk() (*RateLimitEndpointSetCapacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *RateLimitEndpointSetResponse) SetCapacity(v RateLimitEndpointSetCapacity)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *RateLimitEndpointSetResponse) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.
### GetEndpointSetId

`func (o *RateLimitEndpointSetResponse) GetEndpointSetId() string`

GetEndpointSetId returns the EndpointSetId field if non-nil, zero value otherwise.

### GetEndpointSetIdOk

`func (o *RateLimitEndpointSetResponse) GetEndpointSetIdOk() (*string, bool)`

GetEndpointSetIdOk returns a tuple with the EndpointSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSetId

`func (o *RateLimitEndpointSetResponse) SetEndpointSetId(v string)`

SetEndpointSetId sets EndpointSetId field to given value.

### HasEndpointSetId

`func (o *RateLimitEndpointSetResponse) HasEndpointSetId() bool`

HasEndpointSetId returns a boolean if a field has been set.
### GetEndpointSetName

`func (o *RateLimitEndpointSetResponse) GetEndpointSetName() string`

GetEndpointSetName returns the EndpointSetName field if non-nil, zero value otherwise.

### GetEndpointSetNameOk

`func (o *RateLimitEndpointSetResponse) GetEndpointSetNameOk() (*string, bool)`

GetEndpointSetNameOk returns a tuple with the EndpointSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSetName

`func (o *RateLimitEndpointSetResponse) SetEndpointSetName(v string)`

SetEndpointSetName sets EndpointSetName field to given value.

### HasEndpointSetName

`func (o *RateLimitEndpointSetResponse) HasEndpointSetName() bool`

HasEndpointSetName returns a boolean if a field has been set.
### GetEndpoints

`func (o *RateLimitEndpointSetResponse) GetEndpoints() []RateLimitEndpointSetEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *RateLimitEndpointSetResponse) GetEndpointsOk() (*[]RateLimitEndpointSetEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *RateLimitEndpointSetResponse) SetEndpoints(v []RateLimitEndpointSetEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *RateLimitEndpointSetResponse) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.
### GetRefillDurationSeconds

`func (o *RateLimitEndpointSetResponse) GetRefillDurationSeconds() RateLimitEndpointSetRefillDurationSeconds`

GetRefillDurationSeconds returns the RefillDurationSeconds field if non-nil, zero value otherwise.

### GetRefillDurationSecondsOk

`func (o *RateLimitEndpointSetResponse) GetRefillDurationSecondsOk() (*RateLimitEndpointSetRefillDurationSeconds, bool)`

GetRefillDurationSecondsOk returns a tuple with the RefillDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefillDurationSeconds

`func (o *RateLimitEndpointSetResponse) SetRefillDurationSeconds(v RateLimitEndpointSetRefillDurationSeconds)`

SetRefillDurationSeconds sets RefillDurationSeconds field to given value.

### HasRefillDurationSeconds

`func (o *RateLimitEndpointSetResponse) HasRefillDurationSeconds() bool`

HasRefillDurationSeconds returns a boolean if a field has been set.
### GetRefillRate

`func (o *RateLimitEndpointSetResponse) GetRefillRate() RateLimitEndpointSetRefillRate`

GetRefillRate returns the RefillRate field if non-nil, zero value otherwise.

### GetRefillRateOk

`func (o *RateLimitEndpointSetResponse) GetRefillRateOk() (*RateLimitEndpointSetRefillRate, bool)`

GetRefillRateOk returns a tuple with the RefillRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefillRate

`func (o *RateLimitEndpointSetResponse) SetRefillRate(v RateLimitEndpointSetRefillRate)`

SetRefillRate sets RefillRate field to given value.

### HasRefillRate

`func (o *RateLimitEndpointSetResponse) HasRefillRate() bool`

HasRefillRate returns a boolean if a field has been set.
### GetScope

`func (o *RateLimitEndpointSetResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RateLimitEndpointSetResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RateLimitEndpointSetResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RateLimitEndpointSetResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


