# DailySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndHour** | Pointer to **int** | Hour of the day when the scheduled window to run one online archive ends. | [optional] 
**EndMinute** | Pointer to **int** | Minute of the hour when the scheduled window to run one online archive ends. | [optional] 
**StartHour** | Pointer to **int** | Hour of the day when the when the scheduled window to run one online archive starts. | [optional] 
**StartMinute** | Pointer to **int** | Minute of the hour when the scheduled window to run one online archive starts. | [optional] 
**Type** | **string** |  | 

## Methods

### NewDailySchedule

`func NewDailySchedule(type_ string, ) *DailySchedule`

NewDailySchedule instantiates a new DailySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyScheduleWithDefaults

`func NewDailyScheduleWithDefaults() *DailySchedule`

NewDailyScheduleWithDefaults instantiates a new DailySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndHour

`func (o *DailySchedule) GetEndHour() int`

GetEndHour returns the EndHour field if non-nil, zero value otherwise.

### GetEndHourOk

`func (o *DailySchedule) GetEndHourOk() (*int, bool)`

GetEndHourOk returns a tuple with the EndHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHour

`func (o *DailySchedule) SetEndHour(v int)`

SetEndHour sets EndHour field to given value.

### HasEndHour

`func (o *DailySchedule) HasEndHour() bool`

HasEndHour returns a boolean if a field has been set.

### GetEndMinute

`func (o *DailySchedule) GetEndMinute() int`

GetEndMinute returns the EndMinute field if non-nil, zero value otherwise.

### GetEndMinuteOk

`func (o *DailySchedule) GetEndMinuteOk() (*int, bool)`

GetEndMinuteOk returns a tuple with the EndMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMinute

`func (o *DailySchedule) SetEndMinute(v int)`

SetEndMinute sets EndMinute field to given value.

### HasEndMinute

`func (o *DailySchedule) HasEndMinute() bool`

HasEndMinute returns a boolean if a field has been set.

### GetStartHour

`func (o *DailySchedule) GetStartHour() int`

GetStartHour returns the StartHour field if non-nil, zero value otherwise.

### GetStartHourOk

`func (o *DailySchedule) GetStartHourOk() (*int, bool)`

GetStartHourOk returns a tuple with the StartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHour

`func (o *DailySchedule) SetStartHour(v int)`

SetStartHour sets StartHour field to given value.

### HasStartHour

`func (o *DailySchedule) HasStartHour() bool`

HasStartHour returns a boolean if a field has been set.

### GetStartMinute

`func (o *DailySchedule) GetStartMinute() int`

GetStartMinute returns the StartMinute field if non-nil, zero value otherwise.

### GetStartMinuteOk

`func (o *DailySchedule) GetStartMinuteOk() (*int, bool)`

GetStartMinuteOk returns a tuple with the StartMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMinute

`func (o *DailySchedule) SetStartMinute(v int)`

SetStartMinute sets StartMinute field to given value.

### HasStartMinute

`func (o *DailySchedule) HasStartMinute() bool`

HasStartMinute returns a boolean if a field has been set.

### GetType

`func (o *DailySchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DailySchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DailySchedule) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


