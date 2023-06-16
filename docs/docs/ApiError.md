# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **int** | HTTP status code returned with this error. | [optional] 
**ErrorCode** | Pointer to **string** | Application error code returned with this error. | [optional] 
**Parameters** | Pointer to **[]interface{}** | Parameter uses to give more information about the error. | [optional] 
**Reason** | Pointer to **string** | Application error message returned with this error. | [optional] 

## Methods

### NewApiError

`func NewApiError() *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ApiError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ApiError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ApiError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ApiError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetError

`func (o *ApiError) GetError() int`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiError) GetErrorOk() (*int, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiError) SetError(v int)`

SetError sets Error field to given value.

### HasError

`func (o *ApiError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApiError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApiError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApiError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApiError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetParameters

`func (o *ApiError) GetParameters() []interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ApiError) GetParametersOk() (*[]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ApiError) SetParameters(v []interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ApiError) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReason

`func (o *ApiError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ApiError) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


