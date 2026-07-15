# AiModelsRateLimitUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestsPerMinuteLimit** | **int** | The number of requests per minute allowed for this model group. Must be a positive integer. Cannot be more than the organization level limit for this group model. | 
**TokensPerMinuteLimit** | **int** | The number of tokens per minute allowed for this model group. Must be a positive integer. Cannot be more than the organization level limit for this group model. | 

## Methods

### NewAiModelsRateLimitUpdateRequest

`func NewAiModelsRateLimitUpdateRequest(requestsPerMinuteLimit int, tokensPerMinuteLimit int, ) *AiModelsRateLimitUpdateRequest`

NewAiModelsRateLimitUpdateRequest instantiates a new AiModelsRateLimitUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelsRateLimitUpdateRequestWithDefaults

`func NewAiModelsRateLimitUpdateRequestWithDefaults() *AiModelsRateLimitUpdateRequest`

NewAiModelsRateLimitUpdateRequestWithDefaults instantiates a new AiModelsRateLimitUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestsPerMinuteLimit

`func (o *AiModelsRateLimitUpdateRequest) GetRequestsPerMinuteLimit() int`

GetRequestsPerMinuteLimit returns the RequestsPerMinuteLimit field if non-nil, zero value otherwise.

### GetRequestsPerMinuteLimitOk

`func (o *AiModelsRateLimitUpdateRequest) GetRequestsPerMinuteLimitOk() (*int, bool)`

GetRequestsPerMinuteLimitOk returns a tuple with the RequestsPerMinuteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerMinuteLimit

`func (o *AiModelsRateLimitUpdateRequest) SetRequestsPerMinuteLimit(v int)`

SetRequestsPerMinuteLimit sets RequestsPerMinuteLimit field to given value.

### GetTokensPerMinuteLimit

`func (o *AiModelsRateLimitUpdateRequest) GetTokensPerMinuteLimit() int`

GetTokensPerMinuteLimit returns the TokensPerMinuteLimit field if non-nil, zero value otherwise.

### GetTokensPerMinuteLimitOk

`func (o *AiModelsRateLimitUpdateRequest) GetTokensPerMinuteLimitOk() (*int, bool)`

GetTokensPerMinuteLimitOk returns a tuple with the TokensPerMinuteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensPerMinuteLimit

`func (o *AiModelsRateLimitUpdateRequest) SetTokensPerMinuteLimit(v int)`

SetTokensPerMinuteLimit sets TokensPerMinuteLimit field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


