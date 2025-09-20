# CustomSessionTimeouts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteSessionTimeoutInSeconds** | Pointer to **int** | Specifies the absolute session timeout duration in seconds. When set to null, the field&#39;s value is unset, and the default value of 43,200 seconds (12 hours) is applied. Accepted values range between a minimum of 3,600 seconds (1 hour) and a maximum of 43,200 seconds (12 hours). | [optional] 
**IdleSessionTimeoutInSeconds** | Pointer to **int** | Specifies the idle session timeout duration in seconds. When set to null, the field&#39;s value is unset, and the default behavior depends on the context: no timeout for Atlas Commercial, and 600 seconds (10 minutes) for Atlas for Government. Accepted values start at a minimum of 300 seconds (5 minutes). For Atlas Commercial, the maximum value cannot exceed the configured absolute session timeout. For Atlas for Government, the maximum value is capped at 600 seconds (10 minutes). | [optional] 

## Methods

### NewCustomSessionTimeouts

`func NewCustomSessionTimeouts() *CustomSessionTimeouts`

NewCustomSessionTimeouts instantiates a new CustomSessionTimeouts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomSessionTimeoutsWithDefaults

`func NewCustomSessionTimeoutsWithDefaults() *CustomSessionTimeouts`

NewCustomSessionTimeoutsWithDefaults instantiates a new CustomSessionTimeouts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteSessionTimeoutInSeconds

`func (o *CustomSessionTimeouts) GetAbsoluteSessionTimeoutInSeconds() int`

GetAbsoluteSessionTimeoutInSeconds returns the AbsoluteSessionTimeoutInSeconds field if non-nil, zero value otherwise.

### GetAbsoluteSessionTimeoutInSecondsOk

`func (o *CustomSessionTimeouts) GetAbsoluteSessionTimeoutInSecondsOk() (*int, bool)`

GetAbsoluteSessionTimeoutInSecondsOk returns a tuple with the AbsoluteSessionTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteSessionTimeoutInSeconds

`func (o *CustomSessionTimeouts) SetAbsoluteSessionTimeoutInSeconds(v int)`

SetAbsoluteSessionTimeoutInSeconds sets AbsoluteSessionTimeoutInSeconds field to given value.

### HasAbsoluteSessionTimeoutInSeconds

`func (o *CustomSessionTimeouts) HasAbsoluteSessionTimeoutInSeconds() bool`

HasAbsoluteSessionTimeoutInSeconds returns a boolean if a field has been set.
### GetIdleSessionTimeoutInSeconds

`func (o *CustomSessionTimeouts) GetIdleSessionTimeoutInSeconds() int`

GetIdleSessionTimeoutInSeconds returns the IdleSessionTimeoutInSeconds field if non-nil, zero value otherwise.

### GetIdleSessionTimeoutInSecondsOk

`func (o *CustomSessionTimeouts) GetIdleSessionTimeoutInSecondsOk() (*int, bool)`

GetIdleSessionTimeoutInSecondsOk returns a tuple with the IdleSessionTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleSessionTimeoutInSeconds

`func (o *CustomSessionTimeouts) SetIdleSessionTimeoutInSeconds(v int)`

SetIdleSessionTimeoutInSeconds sets IdleSessionTimeoutInSeconds field to given value.

### HasIdleSessionTimeoutInSeconds

`func (o *CustomSessionTimeouts) HasIdleSessionTimeoutInSeconds() bool`

HasIdleSessionTimeoutInSeconds returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


