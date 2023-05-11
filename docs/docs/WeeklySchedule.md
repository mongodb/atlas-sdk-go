# WeeklySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **int** | Day of the week when the scheduled archive starts. The week starts with Monday (&#x60;1&#x60;) and ends with Sunday (&#x60;7&#x60;). | [optional] 
**EndHour** | Pointer to **int** | Hour of the day when the scheduled window to run one online archive ends. | [optional] 
**EndMinute** | Pointer to **int** | Minute of the hour when the scheduled window to run one online archive ends. | [optional] 
**StartHour** | Pointer to **int** | Hour of the day when the when the scheduled window to run one online archive starts. | [optional] 
**StartMinute** | Pointer to **int** | Minute of the hour when the scheduled window to run one online archive starts. | [optional] 
**Type** | **string** |  | 

## Methods

### NewWeeklySchedule

`func NewWeeklySchedule(type_ string, ) *WeeklySchedule`

NewWeeklySchedule instantiates a new WeeklySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeeklyScheduleWithDefaults

`func NewWeeklyScheduleWithDefaults() *WeeklySchedule`

NewWeeklyScheduleWithDefaults instantiates a new WeeklySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *WeeklySchedule) GetDayOfWeek() int`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *WeeklySchedule) GetDayOfWeekOk() (*int, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *WeeklySchedule) SetDayOfWeek(v int)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *WeeklySchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetEndHour

`func (o *WeeklySchedule) GetEndHour() int`

GetEndHour returns the EndHour field if non-nil, zero value otherwise.

### GetEndHourOk

`func (o *WeeklySchedule) GetEndHourOk() (*int, bool)`

GetEndHourOk returns a tuple with the EndHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHour

`func (o *WeeklySchedule) SetEndHour(v int)`

SetEndHour sets EndHour field to given value.

### HasEndHour

`func (o *WeeklySchedule) HasEndHour() bool`

HasEndHour returns a boolean if a field has been set.

### GetEndMinute

`func (o *WeeklySchedule) GetEndMinute() int`

GetEndMinute returns the EndMinute field if non-nil, zero value otherwise.

### GetEndMinuteOk

`func (o *WeeklySchedule) GetEndMinuteOk() (*int, bool)`

GetEndMinuteOk returns a tuple with the EndMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMinute

`func (o *WeeklySchedule) SetEndMinute(v int)`

SetEndMinute sets EndMinute field to given value.

### HasEndMinute

`func (o *WeeklySchedule) HasEndMinute() bool`

HasEndMinute returns a boolean if a field has been set.

### GetStartHour

`func (o *WeeklySchedule) GetStartHour() int`

GetStartHour returns the StartHour field if non-nil, zero value otherwise.

### GetStartHourOk

`func (o *WeeklySchedule) GetStartHourOk() (*int, bool)`

GetStartHourOk returns a tuple with the StartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHour

`func (o *WeeklySchedule) SetStartHour(v int)`

SetStartHour sets StartHour field to given value.

### HasStartHour

`func (o *WeeklySchedule) HasStartHour() bool`

HasStartHour returns a boolean if a field has been set.

### GetStartMinute

`func (o *WeeklySchedule) GetStartMinute() int`

GetStartMinute returns the StartMinute field if non-nil, zero value otherwise.

### GetStartMinuteOk

`func (o *WeeklySchedule) GetStartMinuteOk() (*int, bool)`

GetStartMinuteOk returns a tuple with the StartMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMinute

`func (o *WeeklySchedule) SetStartMinute(v int)`

SetStartMinute sets StartMinute field to given value.

### HasStartMinute

`func (o *WeeklySchedule) HasStartMinute() bool`

HasStartMinute returns a boolean if a field has been set.

### GetType

`func (o *WeeklySchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WeeklySchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WeeklySchedule) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


