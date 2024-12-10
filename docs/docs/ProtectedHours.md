# ProtectedHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndHourOfDay** | Pointer to **int** | Zero-based integer that represents the end hour of the of the day that the maintenance will not begin in. | [optional] 
**StartHourOfDay** | Pointer to **int** | Zero-based integer that represents the beginning hour of the of the day that the maintenance will not begin in. | [optional] 

## Methods

### NewProtectedHours

`func NewProtectedHours() *ProtectedHours`

NewProtectedHours instantiates a new ProtectedHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedHoursWithDefaults

`func NewProtectedHoursWithDefaults() *ProtectedHours`

NewProtectedHoursWithDefaults instantiates a new ProtectedHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndHourOfDay

`func (o *ProtectedHours) GetEndHourOfDay() int`

GetEndHourOfDay returns the EndHourOfDay field if non-nil, zero value otherwise.

### GetEndHourOfDayOk

`func (o *ProtectedHours) GetEndHourOfDayOk() (*int, bool)`

GetEndHourOfDayOk returns a tuple with the EndHourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHourOfDay

`func (o *ProtectedHours) SetEndHourOfDay(v int)`

SetEndHourOfDay sets EndHourOfDay field to given value.

### HasEndHourOfDay

`func (o *ProtectedHours) HasEndHourOfDay() bool`

HasEndHourOfDay returns a boolean if a field has been set.
### GetStartHourOfDay

`func (o *ProtectedHours) GetStartHourOfDay() int`

GetStartHourOfDay returns the StartHourOfDay field if non-nil, zero value otherwise.

### GetStartHourOfDayOk

`func (o *ProtectedHours) GetStartHourOfDayOk() (*int, bool)`

GetStartHourOfDayOk returns a tuple with the StartHourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHourOfDay

`func (o *ProtectedHours) SetStartHourOfDay(v int)`

SetStartHourOfDay sets StartHourOfDay field to given value.

### HasStartHourOfDay

`func (o *ProtectedHours) HasStartHourOfDay() bool`

HasStartHourOfDay returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


