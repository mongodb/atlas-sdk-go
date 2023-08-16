# OnlineArchiveSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of schedule. | 
**EndHour** | Pointer to **int** | Hour of the day when the scheduled window to run one online archive ends. | [optional] 
**EndMinute** | Pointer to **int** | Minute of the hour when the scheduled window to run one online archive ends. | [optional] 
**StartHour** | Pointer to **int** | Hour of the day when the when the scheduled window to run one online archive starts. | [optional] 
**StartMinute** | Pointer to **int** | Minute of the hour when the scheduled window to run one online archive starts. | [optional] 
**DayOfWeek** | Pointer to **int** | Day of the week when the scheduled archive starts. The week starts with Monday (&#x60;1&#x60;) and ends with Sunday (&#x60;7&#x60;). | [optional] 
**DayOfMonth** | Pointer to **int** | Day of the month when the scheduled archive starts. | [optional] 

## Methods

### NewOnlineArchiveSchedule

`func NewOnlineArchiveSchedule(type_ string, ) *OnlineArchiveSchedule`

NewOnlineArchiveSchedule instantiates a new OnlineArchiveSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnlineArchiveScheduleWithDefaults

`func NewOnlineArchiveScheduleWithDefaults() *OnlineArchiveSchedule`

NewOnlineArchiveScheduleWithDefaults instantiates a new OnlineArchiveSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OnlineArchiveSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnlineArchiveSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnlineArchiveSchedule) SetType(v string)`

SetType sets Type field to given value.

### GetEndHour

`func (o *OnlineArchiveSchedule) GetEndHour() int`

GetEndHour returns the EndHour field if non-nil, zero value otherwise.

### GetEndHourOk

`func (o *OnlineArchiveSchedule) GetEndHourOk() (*int, bool)`

GetEndHourOk returns a tuple with the EndHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHour

`func (o *OnlineArchiveSchedule) SetEndHour(v int)`

SetEndHour sets EndHour field to given value.

### HasEndHour

`func (o *OnlineArchiveSchedule) HasEndHour() bool`

HasEndHour returns a boolean if a field has been set.
### GetEndMinute

`func (o *OnlineArchiveSchedule) GetEndMinute() int`

GetEndMinute returns the EndMinute field if non-nil, zero value otherwise.

### GetEndMinuteOk

`func (o *OnlineArchiveSchedule) GetEndMinuteOk() (*int, bool)`

GetEndMinuteOk returns a tuple with the EndMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMinute

`func (o *OnlineArchiveSchedule) SetEndMinute(v int)`

SetEndMinute sets EndMinute field to given value.

### HasEndMinute

`func (o *OnlineArchiveSchedule) HasEndMinute() bool`

HasEndMinute returns a boolean if a field has been set.
### GetStartHour

`func (o *OnlineArchiveSchedule) GetStartHour() int`

GetStartHour returns the StartHour field if non-nil, zero value otherwise.

### GetStartHourOk

`func (o *OnlineArchiveSchedule) GetStartHourOk() (*int, bool)`

GetStartHourOk returns a tuple with the StartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHour

`func (o *OnlineArchiveSchedule) SetStartHour(v int)`

SetStartHour sets StartHour field to given value.

### HasStartHour

`func (o *OnlineArchiveSchedule) HasStartHour() bool`

HasStartHour returns a boolean if a field has been set.
### GetStartMinute

`func (o *OnlineArchiveSchedule) GetStartMinute() int`

GetStartMinute returns the StartMinute field if non-nil, zero value otherwise.

### GetStartMinuteOk

`func (o *OnlineArchiveSchedule) GetStartMinuteOk() (*int, bool)`

GetStartMinuteOk returns a tuple with the StartMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMinute

`func (o *OnlineArchiveSchedule) SetStartMinute(v int)`

SetStartMinute sets StartMinute field to given value.

### HasStartMinute

`func (o *OnlineArchiveSchedule) HasStartMinute() bool`

HasStartMinute returns a boolean if a field has been set.
### GetDayOfWeek

`func (o *OnlineArchiveSchedule) GetDayOfWeek() int`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *OnlineArchiveSchedule) GetDayOfWeekOk() (*int, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *OnlineArchiveSchedule) SetDayOfWeek(v int)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *OnlineArchiveSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.
### GetDayOfMonth

`func (o *OnlineArchiveSchedule) GetDayOfMonth() int`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *OnlineArchiveSchedule) GetDayOfMonthOk() (*int, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *OnlineArchiveSchedule) SetDayOfMonth(v int)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *OnlineArchiveSchedule) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


