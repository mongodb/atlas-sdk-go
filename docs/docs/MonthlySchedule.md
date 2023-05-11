# MonthlySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int** | Day of the month when the scheduled archive starts. | [optional] 
**EndHour** | Pointer to **int** | Hour of the day when the scheduled window to run one online archive ends. | [optional] 
**EndMinute** | Pointer to **int** | Minute of the hour when the scheduled window to run one online archive ends. | [optional] 
**StartHour** | Pointer to **int** | Hour of the day when the when the scheduled window to run one online archive starts. | [optional] 
**StartMinute** | Pointer to **int** | Minute of the hour when the scheduled window to run one online archive starts. | [optional] 
**Type** | **string** |  | 

## Methods

### NewMonthlySchedule

`func NewMonthlySchedule(type_ string, ) *MonthlySchedule`

NewMonthlySchedule instantiates a new MonthlySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyScheduleWithDefaults

`func NewMonthlyScheduleWithDefaults() *MonthlySchedule`

NewMonthlyScheduleWithDefaults instantiates a new MonthlySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *MonthlySchedule) GetDayOfMonth() int`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *MonthlySchedule) GetDayOfMonthOk() (*int, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *MonthlySchedule) SetDayOfMonth(v int)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *MonthlySchedule) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetEndHour

`func (o *MonthlySchedule) GetEndHour() int`

GetEndHour returns the EndHour field if non-nil, zero value otherwise.

### GetEndHourOk

`func (o *MonthlySchedule) GetEndHourOk() (*int, bool)`

GetEndHourOk returns a tuple with the EndHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHour

`func (o *MonthlySchedule) SetEndHour(v int)`

SetEndHour sets EndHour field to given value.

### HasEndHour

`func (o *MonthlySchedule) HasEndHour() bool`

HasEndHour returns a boolean if a field has been set.

### GetEndMinute

`func (o *MonthlySchedule) GetEndMinute() int`

GetEndMinute returns the EndMinute field if non-nil, zero value otherwise.

### GetEndMinuteOk

`func (o *MonthlySchedule) GetEndMinuteOk() (*int, bool)`

GetEndMinuteOk returns a tuple with the EndMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMinute

`func (o *MonthlySchedule) SetEndMinute(v int)`

SetEndMinute sets EndMinute field to given value.

### HasEndMinute

`func (o *MonthlySchedule) HasEndMinute() bool`

HasEndMinute returns a boolean if a field has been set.

### GetStartHour

`func (o *MonthlySchedule) GetStartHour() int`

GetStartHour returns the StartHour field if non-nil, zero value otherwise.

### GetStartHourOk

`func (o *MonthlySchedule) GetStartHourOk() (*int, bool)`

GetStartHourOk returns a tuple with the StartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHour

`func (o *MonthlySchedule) SetStartHour(v int)`

SetStartHour sets StartHour field to given value.

### HasStartHour

`func (o *MonthlySchedule) HasStartHour() bool`

HasStartHour returns a boolean if a field has been set.

### GetStartMinute

`func (o *MonthlySchedule) GetStartMinute() int`

GetStartMinute returns the StartMinute field if non-nil, zero value otherwise.

### GetStartMinuteOk

`func (o *MonthlySchedule) GetStartMinuteOk() (*int, bool)`

GetStartMinuteOk returns a tuple with the StartMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMinute

`func (o *MonthlySchedule) SetStartMinute(v int)`

SetStartMinute sets StartMinute field to given value.

### HasStartMinute

`func (o *MonthlySchedule) HasStartMinute() bool`

HasStartMinute returns a boolean if a field has been set.

### GetType

`func (o *MonthlySchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonthlySchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonthlySchedule) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


