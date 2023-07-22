# GroupMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeferOnceEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them. | [optional] 
**DayOfWeek** | **int** | One-based integer that represents the day of the week that the maintenance window starts.  | Value | Day of Week | |---|---| | &#x60;1&#x60; | Sunday | | &#x60;2&#x60; | Monday | | &#x60;3&#x60; | Tuesday | | &#x60;4&#x60; | Wednesday | | &#x60;5&#x60; | Thursday | | &#x60;6&#x60; | Friday | | &#x60;7&#x60; | Saturday |  | 
**HourOfDay** | **int** | Zero-based integer that represents the hour of the of the day that the maintenance window starts according to a 24-hour clock. Use &#x60;0&#x60; for midnight and &#x60;12&#x60; for noon. | 
**StartASAP** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud starts the maintenance window immediately upon receiving this request. To start the maintenance window immediately for your project, MongoDB Cloud must have maintenance scheduled and you must set a maintenance window. This flag resets to &#x60;false&#x60; after MongoDB Cloud completes maintenance. | [optional] 

## Methods

### NewGroupMaintenanceWindow

`func NewGroupMaintenanceWindow(dayOfWeek int, hourOfDay int, ) *GroupMaintenanceWindow`

NewGroupMaintenanceWindow instantiates a new GroupMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMaintenanceWindowWithDefaults

`func NewGroupMaintenanceWindowWithDefaults() *GroupMaintenanceWindow`

NewGroupMaintenanceWindowWithDefaults instantiates a new GroupMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeferOnceEnabled

`func (o *GroupMaintenanceWindow) GetAutoDeferOnceEnabled() bool`

GetAutoDeferOnceEnabled returns the AutoDeferOnceEnabled field if non-nil, zero value otherwise.

### GetAutoDeferOnceEnabledOk

`func (o *GroupMaintenanceWindow) GetAutoDeferOnceEnabledOk() (*bool, bool)`

GetAutoDeferOnceEnabledOk returns a tuple with the AutoDeferOnceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeferOnceEnabled

`func (o *GroupMaintenanceWindow) SetAutoDeferOnceEnabled(v bool)`

SetAutoDeferOnceEnabled sets AutoDeferOnceEnabled field to given value.

### HasAutoDeferOnceEnabled

`func (o *GroupMaintenanceWindow) HasAutoDeferOnceEnabled() bool`

HasAutoDeferOnceEnabled returns a boolean if a field has been set.
### GetDayOfWeek

`func (o *GroupMaintenanceWindow) GetDayOfWeek() int`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *GroupMaintenanceWindow) GetDayOfWeekOk() (*int, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *GroupMaintenanceWindow) SetDayOfWeek(v int)`

SetDayOfWeek sets DayOfWeek field to given value.

### GetHourOfDay

`func (o *GroupMaintenanceWindow) GetHourOfDay() int`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *GroupMaintenanceWindow) GetHourOfDayOk() (*int, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *GroupMaintenanceWindow) SetHourOfDay(v int)`

SetHourOfDay sets HourOfDay field to given value.

### GetStartASAP

`func (o *GroupMaintenanceWindow) GetStartASAP() bool`

GetStartASAP returns the StartASAP field if non-nil, zero value otherwise.

### GetStartASAPOk

`func (o *GroupMaintenanceWindow) GetStartASAPOk() (*bool, bool)`

GetStartASAPOk returns a tuple with the StartASAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartASAP

`func (o *GroupMaintenanceWindow) SetStartASAP(v bool)`

SetStartASAP sets StartASAP field to given value.

### HasStartASAP

`func (o *GroupMaintenanceWindow) HasStartASAP() bool`

HasStartASAP returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


