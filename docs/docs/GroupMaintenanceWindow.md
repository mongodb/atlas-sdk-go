# GroupMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeferOnceEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them. This setting controls the same underlying auto-deferral feature as the &#x60;/maintenanceWindow/autoDefer&#x60; endpoint. Use either this field (to set a specific value) or that endpoint (to toggle the current value). For most use cases, this field in the PATCH request is preferred because it allows setting an explicit value rather than toggling. | [optional] 
**DayOfWeek** | **int** | One-based integer that represents the day of the week that the maintenance window starts.  - &#x60;1&#x60;: Sunday. - &#x60;2&#x60;: Monday. - &#x60;3&#x60;: Tuesday. - &#x60;4&#x60;: Wednesday. - &#x60;5&#x60;: Thursday. - &#x60;6&#x60;: Friday. - &#x60;7&#x60;: Saturday. | 
**HourOfDay** | Pointer to **int** | Zero-based integer that represents the hour of the of the day that the maintenance window starts according to a 24-hour clock. Use &#x60;0&#x60; for midnight and &#x60;12&#x60; for noon. | [optional] 
**NumberOfDeferrals** | Pointer to **int** | Number of times the current maintenance event for this project has been deferred. | [optional] [readonly] 
**ProtectedHours** | Pointer to [**ProtectedHours**](ProtectedHours.md) |  | [optional] 
**StartASAP** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud starts the maintenance window immediately upon receiving this request. To start the maintenance window immediately for your project, MongoDB Cloud must have maintenance scheduled and you must set a maintenance window. This flag resets to &#x60;false&#x60; after MongoDB Cloud completes maintenance. | [optional] 
**TimeZoneId** | Pointer to **string** | Identifier for the current time zone of the maintenance window. This can only be updated via the Project Settings UI. | [optional] [readonly] 

## Methods

### NewGroupMaintenanceWindow

`func NewGroupMaintenanceWindow(dayOfWeek int, ) *GroupMaintenanceWindow`

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

### HasHourOfDay

`func (o *GroupMaintenanceWindow) HasHourOfDay() bool`

HasHourOfDay returns a boolean if a field has been set.
### GetNumberOfDeferrals

`func (o *GroupMaintenanceWindow) GetNumberOfDeferrals() int`

GetNumberOfDeferrals returns the NumberOfDeferrals field if non-nil, zero value otherwise.

### GetNumberOfDeferralsOk

`func (o *GroupMaintenanceWindow) GetNumberOfDeferralsOk() (*int, bool)`

GetNumberOfDeferralsOk returns a tuple with the NumberOfDeferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDeferrals

`func (o *GroupMaintenanceWindow) SetNumberOfDeferrals(v int)`

SetNumberOfDeferrals sets NumberOfDeferrals field to given value.

### HasNumberOfDeferrals

`func (o *GroupMaintenanceWindow) HasNumberOfDeferrals() bool`

HasNumberOfDeferrals returns a boolean if a field has been set.
### GetProtectedHours

`func (o *GroupMaintenanceWindow) GetProtectedHours() ProtectedHours`

GetProtectedHours returns the ProtectedHours field if non-nil, zero value otherwise.

### GetProtectedHoursOk

`func (o *GroupMaintenanceWindow) GetProtectedHoursOk() (*ProtectedHours, bool)`

GetProtectedHoursOk returns a tuple with the ProtectedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedHours

`func (o *GroupMaintenanceWindow) SetProtectedHours(v ProtectedHours)`

SetProtectedHours sets ProtectedHours field to given value.

### HasProtectedHours

`func (o *GroupMaintenanceWindow) HasProtectedHours() bool`

HasProtectedHours returns a boolean if a field has been set.
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
### GetTimeZoneId

`func (o *GroupMaintenanceWindow) GetTimeZoneId() string`

GetTimeZoneId returns the TimeZoneId field if non-nil, zero value otherwise.

### GetTimeZoneIdOk

`func (o *GroupMaintenanceWindow) GetTimeZoneIdOk() (*string, bool)`

GetTimeZoneIdOk returns a tuple with the TimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneId

`func (o *GroupMaintenanceWindow) SetTimeZoneId(v string)`

SetTimeZoneId sets TimeZoneId field to given value.

### HasTimeZoneId

`func (o *GroupMaintenanceWindow) HasTimeZoneId() bool`

HasTimeZoneId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


