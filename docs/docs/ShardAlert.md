# ShardAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Human-readable description of the condition that triggered this alert. | [readonly] 
**Severity** | **string** | Severity that MongoDB Cloud assigned to this alert. | [readonly] 

## Methods

### NewShardAlert

`func NewShardAlert(message string, severity string, ) *ShardAlert`

NewShardAlert instantiates a new ShardAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardAlertWithDefaults

`func NewShardAlertWithDefaults() *ShardAlert`

NewShardAlertWithDefaults instantiates a new ShardAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ShardAlert) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ShardAlert) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ShardAlert) SetMessage(v string)`

SetMessage sets Message field to given value.

### GetSeverity

`func (o *ShardAlert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ShardAlert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ShardAlert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


