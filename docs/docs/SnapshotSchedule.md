# SnapshotSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCheckpointIntervalMin** | **int** | Quantity of time expressed in minutes between successive cluster checkpoints. This parameter applies only to sharded clusters. This number determines the granularity of continuous cloud backups for sharded clusters. | 
**ClusterId** | **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return. | 
**DailySnapshotRetentionDays** | **int** | Quantity of time to keep daily snapshots. MongoDB Cloud expresses this value in days. Set this value to &#x60;0&#x60; to disable daily snapshot retention. | 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project that contains the cluster. | [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MonthlySnapshotRetentionMonths** | **int** | Number of months that MongoDB Cloud must keep monthly snapshots. Set this value to &#x60;0&#x60; to disable monthly snapshot retention. | 
**PointInTimeWindowHours** | **int** | Number of hours before the current time from which MongoDB Cloud can create a Continuous Cloud Backup snapshot. | 
**SnapshotIntervalHours** | **int** | Number of hours that must elapse before taking another snapshot. | 
**SnapshotRetentionDays** | **int** | Number of days that MongoDB Cloud must keep recent snapshots. | 
**WeeklySnapshotRetentionWeeks** | **int** | Number of weeks that MongoDB Cloud must keep weekly snapshots. Set this value to &#x60;0&#x60; to disable weekly snapshot retention. | 

## Methods

### NewSnapshotSchedule

`func NewSnapshotSchedule(clusterCheckpointIntervalMin int, clusterId string, dailySnapshotRetentionDays int, groupId string, monthlySnapshotRetentionMonths int, pointInTimeWindowHours int, snapshotIntervalHours int, snapshotRetentionDays int, weeklySnapshotRetentionWeeks int, ) *SnapshotSchedule`

NewSnapshotSchedule instantiates a new SnapshotSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotScheduleWithDefaults

`func NewSnapshotScheduleWithDefaults() *SnapshotSchedule`

NewSnapshotScheduleWithDefaults instantiates a new SnapshotSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCheckpointIntervalMin

`func (o *SnapshotSchedule) GetClusterCheckpointIntervalMin() int`

GetClusterCheckpointIntervalMin returns the ClusterCheckpointIntervalMin field if non-nil, zero value otherwise.

### GetClusterCheckpointIntervalMinOk

`func (o *SnapshotSchedule) GetClusterCheckpointIntervalMinOk() (*int, bool)`

GetClusterCheckpointIntervalMinOk returns a tuple with the ClusterCheckpointIntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCheckpointIntervalMin

`func (o *SnapshotSchedule) SetClusterCheckpointIntervalMin(v int)`

SetClusterCheckpointIntervalMin sets ClusterCheckpointIntervalMin field to given value.


### GetClusterId

`func (o *SnapshotSchedule) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SnapshotSchedule) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SnapshotSchedule) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetDailySnapshotRetentionDays

`func (o *SnapshotSchedule) GetDailySnapshotRetentionDays() int`

GetDailySnapshotRetentionDays returns the DailySnapshotRetentionDays field if non-nil, zero value otherwise.

### GetDailySnapshotRetentionDaysOk

`func (o *SnapshotSchedule) GetDailySnapshotRetentionDaysOk() (*int, bool)`

GetDailySnapshotRetentionDaysOk returns a tuple with the DailySnapshotRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySnapshotRetentionDays

`func (o *SnapshotSchedule) SetDailySnapshotRetentionDays(v int)`

SetDailySnapshotRetentionDays sets DailySnapshotRetentionDays field to given value.


### GetGroupId

`func (o *SnapshotSchedule) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SnapshotSchedule) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SnapshotSchedule) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLinks

`func (o *SnapshotSchedule) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SnapshotSchedule) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SnapshotSchedule) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SnapshotSchedule) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMonthlySnapshotRetentionMonths

`func (o *SnapshotSchedule) GetMonthlySnapshotRetentionMonths() int`

GetMonthlySnapshotRetentionMonths returns the MonthlySnapshotRetentionMonths field if non-nil, zero value otherwise.

### GetMonthlySnapshotRetentionMonthsOk

`func (o *SnapshotSchedule) GetMonthlySnapshotRetentionMonthsOk() (*int, bool)`

GetMonthlySnapshotRetentionMonthsOk returns a tuple with the MonthlySnapshotRetentionMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlySnapshotRetentionMonths

`func (o *SnapshotSchedule) SetMonthlySnapshotRetentionMonths(v int)`

SetMonthlySnapshotRetentionMonths sets MonthlySnapshotRetentionMonths field to given value.


### GetPointInTimeWindowHours

`func (o *SnapshotSchedule) GetPointInTimeWindowHours() int`

GetPointInTimeWindowHours returns the PointInTimeWindowHours field if non-nil, zero value otherwise.

### GetPointInTimeWindowHoursOk

`func (o *SnapshotSchedule) GetPointInTimeWindowHoursOk() (*int, bool)`

GetPointInTimeWindowHoursOk returns a tuple with the PointInTimeWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeWindowHours

`func (o *SnapshotSchedule) SetPointInTimeWindowHours(v int)`

SetPointInTimeWindowHours sets PointInTimeWindowHours field to given value.


### GetSnapshotIntervalHours

`func (o *SnapshotSchedule) GetSnapshotIntervalHours() int`

GetSnapshotIntervalHours returns the SnapshotIntervalHours field if non-nil, zero value otherwise.

### GetSnapshotIntervalHoursOk

`func (o *SnapshotSchedule) GetSnapshotIntervalHoursOk() (*int, bool)`

GetSnapshotIntervalHoursOk returns a tuple with the SnapshotIntervalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIntervalHours

`func (o *SnapshotSchedule) SetSnapshotIntervalHours(v int)`

SetSnapshotIntervalHours sets SnapshotIntervalHours field to given value.


### GetSnapshotRetentionDays

`func (o *SnapshotSchedule) GetSnapshotRetentionDays() int`

GetSnapshotRetentionDays returns the SnapshotRetentionDays field if non-nil, zero value otherwise.

### GetSnapshotRetentionDaysOk

`func (o *SnapshotSchedule) GetSnapshotRetentionDaysOk() (*int, bool)`

GetSnapshotRetentionDaysOk returns a tuple with the SnapshotRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionDays

`func (o *SnapshotSchedule) SetSnapshotRetentionDays(v int)`

SetSnapshotRetentionDays sets SnapshotRetentionDays field to given value.


### GetWeeklySnapshotRetentionWeeks

`func (o *SnapshotSchedule) GetWeeklySnapshotRetentionWeeks() int`

GetWeeklySnapshotRetentionWeeks returns the WeeklySnapshotRetentionWeeks field if non-nil, zero value otherwise.

### GetWeeklySnapshotRetentionWeeksOk

`func (o *SnapshotSchedule) GetWeeklySnapshotRetentionWeeksOk() (*int, bool)`

GetWeeklySnapshotRetentionWeeksOk returns a tuple with the WeeklySnapshotRetentionWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklySnapshotRetentionWeeks

`func (o *SnapshotSchedule) SetWeeklySnapshotRetentionWeeks(v int)`

SetWeeklySnapshotRetentionWeeks sets WeeklySnapshotRetentionWeeks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


