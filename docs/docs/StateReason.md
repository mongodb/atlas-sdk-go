# StateReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | Error code relating to state. | [optional] 
**Message** | Pointer to **string** | Message describing error or state. | [optional] 

## Methods

### NewStateReason

`func NewStateReason() *StateReason`

NewStateReason instantiates a new StateReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateReasonWithDefaults

`func NewStateReasonWithDefaults() *StateReason`

NewStateReasonWithDefaults instantiates a new StateReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *StateReason) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *StateReason) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *StateReason) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *StateReason) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *StateReason) SetErrorCodeNil()`

SetErrorCodeNil sets ErrorCode to an explicit JSON null when marshaled, overriding any value previously set with SetErrorCode. Calling SetErrorCode again clears the null override.

### GetMessage

`func (o *StateReason) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StateReason) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StateReason) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StateReason) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *StateReason) SetMessageNil()`

SetMessageNil sets Message to an explicit JSON null when marshaled, overriding any value previously set with SetMessage. Calling SetMessage again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


