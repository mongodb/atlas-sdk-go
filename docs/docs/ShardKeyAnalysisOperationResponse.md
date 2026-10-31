# ShardKeyAnalysisOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Date and time when MongoDB Cloud accepted this operation (ISO 8601 format in UTC). | [readonly] 
**Error** | Pointer to [**ShardKeyAnalysisError**](ShardKeyAnalysisError.md) |  | [optional] 
**ExpiresAt** | **time.Time** | Date and time when MongoDB Cloud stops returning this operation (ISO 8601 format in UTC). Reads after this time return an HTTP 404 status. This value derives from &#x60;updatedAt&#x60;, so it advances with each state change while the operation is running and settles only once the operation is terminal. | [readonly] 
**OperationId** | **string** | Unique 24-hexadecimal digit string that identifies this operation. When the operation succeeds, the analysis it created carries the same identifier. | [readonly] 
**OperationType** | **string** | Kind of change the operation makes. Submitting a shard key analysis creates an analysis resource, so the operation type is always &#x60;CREATE&#x60;. | [readonly] 
**ResultHref** | Pointer to **string** | URI of the analysis this operation created. MongoDB Cloud returns this parameter only when &#x60;status&#x60; is &#x60;SUCCEEDED&#x60;. | [optional] [readonly] 
**RetryAfterSeconds** | Pointer to **int** | Number of seconds to wait before polling this operation again. MongoDB Cloud returns this parameter only while &#x60;status&#x60; is not terminal. | [optional] [readonly] 
**Status** | **string** | State of the shard key analysis operation. &#x60;PENDING&#x60; and &#x60;IN_PROGRESS&#x60; are non-terminal; &#x60;SUCCEEDED&#x60;, &#x60;FAILED&#x60;, &#x60;CANCELED&#x60; and &#x60;SUPERSEDED&#x60; are terminal. MongoDB Cloud currently reports only &#x60;PENDING&#x60;, &#x60;SUCCEEDED&#x60; and &#x60;FAILED&#x60;. The remaining values are reserved and may be reported in the future without a new resource version. | [readonly] 
**UpdatedAt** | **time.Time** | Date and time when this operation last changed state (ISO 8601 format in UTC). Once the operation is terminal, this is when it finished. | [readonly] 

## Methods

### NewShardKeyAnalysisOperationResponse

`func NewShardKeyAnalysisOperationResponse(createdAt time.Time, expiresAt time.Time, operationId string, operationType string, status string, updatedAt time.Time, ) *ShardKeyAnalysisOperationResponse`

NewShardKeyAnalysisOperationResponse instantiates a new ShardKeyAnalysisOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardKeyAnalysisOperationResponseWithDefaults

`func NewShardKeyAnalysisOperationResponseWithDefaults() *ShardKeyAnalysisOperationResponse`

NewShardKeyAnalysisOperationResponseWithDefaults instantiates a new ShardKeyAnalysisOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ShardKeyAnalysisOperationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShardKeyAnalysisOperationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShardKeyAnalysisOperationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetError

`func (o *ShardKeyAnalysisOperationResponse) GetError() ShardKeyAnalysisError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ShardKeyAnalysisOperationResponse) GetErrorOk() (*ShardKeyAnalysisError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ShardKeyAnalysisOperationResponse) SetError(v ShardKeyAnalysisError)`

SetError sets Error field to given value.

### HasError

`func (o *ShardKeyAnalysisOperationResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ShardKeyAnalysisOperationResponse) SetErrorNil()`

SetErrorNil sets Error to an explicit JSON null when marshaled, overriding any value previously set with SetError. Calling SetError again clears the null override.

### GetExpiresAt

`func (o *ShardKeyAnalysisOperationResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ShardKeyAnalysisOperationResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ShardKeyAnalysisOperationResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### GetOperationId

`func (o *ShardKeyAnalysisOperationResponse) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *ShardKeyAnalysisOperationResponse) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *ShardKeyAnalysisOperationResponse) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### GetOperationType

`func (o *ShardKeyAnalysisOperationResponse) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ShardKeyAnalysisOperationResponse) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ShardKeyAnalysisOperationResponse) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### GetResultHref

`func (o *ShardKeyAnalysisOperationResponse) GetResultHref() string`

GetResultHref returns the ResultHref field if non-nil, zero value otherwise.

### GetResultHrefOk

`func (o *ShardKeyAnalysisOperationResponse) GetResultHrefOk() (*string, bool)`

GetResultHrefOk returns a tuple with the ResultHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultHref

`func (o *ShardKeyAnalysisOperationResponse) SetResultHref(v string)`

SetResultHref sets ResultHref field to given value.

### HasResultHref

`func (o *ShardKeyAnalysisOperationResponse) HasResultHref() bool`

HasResultHref returns a boolean if a field has been set.

### SetResultHrefNil

`func (o *ShardKeyAnalysisOperationResponse) SetResultHrefNil()`

SetResultHrefNil sets ResultHref to an explicit JSON null when marshaled, overriding any value previously set with SetResultHref. Calling SetResultHref again clears the null override.

### GetRetryAfterSeconds

`func (o *ShardKeyAnalysisOperationResponse) GetRetryAfterSeconds() int`

GetRetryAfterSeconds returns the RetryAfterSeconds field if non-nil, zero value otherwise.

### GetRetryAfterSecondsOk

`func (o *ShardKeyAnalysisOperationResponse) GetRetryAfterSecondsOk() (*int, bool)`

GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterSeconds

`func (o *ShardKeyAnalysisOperationResponse) SetRetryAfterSeconds(v int)`

SetRetryAfterSeconds sets RetryAfterSeconds field to given value.

### HasRetryAfterSeconds

`func (o *ShardKeyAnalysisOperationResponse) HasRetryAfterSeconds() bool`

HasRetryAfterSeconds returns a boolean if a field has been set.

### SetRetryAfterSecondsNil

`func (o *ShardKeyAnalysisOperationResponse) SetRetryAfterSecondsNil()`

SetRetryAfterSecondsNil sets RetryAfterSeconds to an explicit JSON null when marshaled, overriding any value previously set with SetRetryAfterSeconds. Calling SetRetryAfterSeconds again clears the null override.

### GetStatus

`func (o *ShardKeyAnalysisOperationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShardKeyAnalysisOperationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShardKeyAnalysisOperationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### GetUpdatedAt

`func (o *ShardKeyAnalysisOperationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ShardKeyAnalysisOperationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ShardKeyAnalysisOperationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


