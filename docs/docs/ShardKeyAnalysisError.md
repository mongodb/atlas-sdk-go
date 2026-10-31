# ShardKeyAnalysisError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Reason that the shard key analysis failed. MongoDB Cloud may report additional values in the future without a new resource version, so treat a value this description does not list as a generic failure. | [readonly] 
**Message** | **string** | Human-readable explanation of the failure. This message describes the classified failure and is not the raw error text from the cluster. | [readonly] 
**RetryStrategy** | Pointer to **string** | How to retry a failed shard key analysis. &#x60;BACKOFF&#x60; indicates that resubmitting the same request after a delay may succeed. &#x60;NONE&#x60; indicates that the failure describes the collection or the candidate shard key and that resubmitting the same request will fail the same way. | [optional] [readonly] 
**Retryable** | **bool** | Flag that indicates whether submitting the same request again may succeed. This flag is &#x60;false&#x60; when the failure describes the collection or the candidate shard key, because the same request then fails the same way. | [readonly] 

## Methods

### NewShardKeyAnalysisError

`func NewShardKeyAnalysisError(code string, message string, retryable bool, ) *ShardKeyAnalysisError`

NewShardKeyAnalysisError instantiates a new ShardKeyAnalysisError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardKeyAnalysisErrorWithDefaults

`func NewShardKeyAnalysisErrorWithDefaults() *ShardKeyAnalysisError`

NewShardKeyAnalysisErrorWithDefaults instantiates a new ShardKeyAnalysisError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ShardKeyAnalysisError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ShardKeyAnalysisError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ShardKeyAnalysisError) SetCode(v string)`

SetCode sets Code field to given value.

### GetMessage

`func (o *ShardKeyAnalysisError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ShardKeyAnalysisError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ShardKeyAnalysisError) SetMessage(v string)`

SetMessage sets Message field to given value.

### GetRetryStrategy

`func (o *ShardKeyAnalysisError) GetRetryStrategy() string`

GetRetryStrategy returns the RetryStrategy field if non-nil, zero value otherwise.

### GetRetryStrategyOk

`func (o *ShardKeyAnalysisError) GetRetryStrategyOk() (*string, bool)`

GetRetryStrategyOk returns a tuple with the RetryStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryStrategy

`func (o *ShardKeyAnalysisError) SetRetryStrategy(v string)`

SetRetryStrategy sets RetryStrategy field to given value.

### HasRetryStrategy

`func (o *ShardKeyAnalysisError) HasRetryStrategy() bool`

HasRetryStrategy returns a boolean if a field has been set.

### SetRetryStrategyNil

`func (o *ShardKeyAnalysisError) SetRetryStrategyNil()`

SetRetryStrategyNil sets RetryStrategy to an explicit JSON null when marshaled, overriding any value previously set with SetRetryStrategy. Calling SetRetryStrategy again clears the null override.

### GetRetryable

`func (o *ShardKeyAnalysisError) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *ShardKeyAnalysisError) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *ShardKeyAnalysisError) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


