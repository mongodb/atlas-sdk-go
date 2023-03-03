/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the SnapshotSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotSchedule{}

// SnapshotSchedule struct for SnapshotSchedule
type SnapshotSchedule struct {
	// Quantity of time expressed in minutes between successive cluster checkpoints. This parameter applies only to sharded clusters. This number determines the granularity of continuous cloud backups for sharded clusters.
	ClusterCheckpointIntervalMin int32 `json:"clusterCheckpointIntervalMin"`
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return.
	ClusterId string `json:"clusterId"`
	// Quantity of time to keep daily snapshots. MongoDB Cloud expresses this value in days. Set this value to `0` to disable daily snapshot retention.
	DailySnapshotRetentionDays int32 `json:"dailySnapshotRetentionDays"`
	// Unique 24-hexadecimal digit string that identifies the project that contains the cluster.
	GroupId string `json:"groupId"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Number of months that MongoDB Cloud must keep monthly snapshots. Set this value to `0` to disable monthly snapshot retention.
	MonthlySnapshotRetentionMonths int32 `json:"monthlySnapshotRetentionMonths"`
	// Number of hours before the current time from which MongoDB Cloud can create a Continuous Cloud Backup snapshot.
	PointInTimeWindowHours int32 `json:"pointInTimeWindowHours"`
	// Number of hours that must elapse before taking another snapshot.
	SnapshotIntervalHours int32 `json:"snapshotIntervalHours"`
	// Number of days that MongoDB Cloud must keep recent snapshots.
	SnapshotRetentionDays int32 `json:"snapshotRetentionDays"`
	// Number of weeks that MongoDB Cloud must keep weekly snapshots. Set this value to `0` to disable weekly snapshot retention.
	WeeklySnapshotRetentionWeeks int32 `json:"weeklySnapshotRetentionWeeks"`
}

// NewSnapshotSchedule instantiates a new SnapshotSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotSchedule(clusterCheckpointIntervalMin int32, clusterId string, dailySnapshotRetentionDays int32, groupId string, monthlySnapshotRetentionMonths int32, pointInTimeWindowHours int32, snapshotIntervalHours int32, snapshotRetentionDays int32, weeklySnapshotRetentionWeeks int32) *SnapshotSchedule {
	this := SnapshotSchedule{}
	this.ClusterCheckpointIntervalMin = clusterCheckpointIntervalMin
	this.ClusterId = clusterId
	this.DailySnapshotRetentionDays = dailySnapshotRetentionDays
	this.GroupId = groupId
	this.MonthlySnapshotRetentionMonths = monthlySnapshotRetentionMonths
	this.PointInTimeWindowHours = pointInTimeWindowHours
	this.SnapshotIntervalHours = snapshotIntervalHours
	this.SnapshotRetentionDays = snapshotRetentionDays
	this.WeeklySnapshotRetentionWeeks = weeklySnapshotRetentionWeeks
	return &this
}

// NewSnapshotScheduleWithDefaults instantiates a new SnapshotSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotScheduleWithDefaults() *SnapshotSchedule {
	this := SnapshotSchedule{}
	return &this
}

// GetClusterCheckpointIntervalMin returns the ClusterCheckpointIntervalMin field value
func (o *SnapshotSchedule) GetClusterCheckpointIntervalMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClusterCheckpointIntervalMin
}

// GetClusterCheckpointIntervalMinOk returns a tuple with the ClusterCheckpointIntervalMin field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetClusterCheckpointIntervalMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterCheckpointIntervalMin, true
}

// SetClusterCheckpointIntervalMin sets field value
func (o *SnapshotSchedule) SetClusterCheckpointIntervalMin(v int32) {
	o.ClusterCheckpointIntervalMin = v
}

// GetClusterId returns the ClusterId field value
func (o *SnapshotSchedule) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *SnapshotSchedule) SetClusterId(v string) {
	o.ClusterId = v
}

// GetDailySnapshotRetentionDays returns the DailySnapshotRetentionDays field value
func (o *SnapshotSchedule) GetDailySnapshotRetentionDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DailySnapshotRetentionDays
}

// GetDailySnapshotRetentionDaysOk returns a tuple with the DailySnapshotRetentionDays field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetDailySnapshotRetentionDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DailySnapshotRetentionDays, true
}

// SetDailySnapshotRetentionDays sets field value
func (o *SnapshotSchedule) SetDailySnapshotRetentionDays(v int32) {
	o.DailySnapshotRetentionDays = v
}

// GetGroupId returns the GroupId field value
func (o *SnapshotSchedule) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *SnapshotSchedule) SetGroupId(v string) {
	o.GroupId = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SnapshotSchedule) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SnapshotSchedule) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *SnapshotSchedule) SetLinks(v []Link) {
	o.Links = v
}

// GetMonthlySnapshotRetentionMonths returns the MonthlySnapshotRetentionMonths field value
func (o *SnapshotSchedule) GetMonthlySnapshotRetentionMonths() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MonthlySnapshotRetentionMonths
}

// GetMonthlySnapshotRetentionMonthsOk returns a tuple with the MonthlySnapshotRetentionMonths field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetMonthlySnapshotRetentionMonthsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlySnapshotRetentionMonths, true
}

// SetMonthlySnapshotRetentionMonths sets field value
func (o *SnapshotSchedule) SetMonthlySnapshotRetentionMonths(v int32) {
	o.MonthlySnapshotRetentionMonths = v
}

// GetPointInTimeWindowHours returns the PointInTimeWindowHours field value
func (o *SnapshotSchedule) GetPointInTimeWindowHours() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PointInTimeWindowHours
}

// GetPointInTimeWindowHoursOk returns a tuple with the PointInTimeWindowHours field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetPointInTimeWindowHoursOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PointInTimeWindowHours, true
}

// SetPointInTimeWindowHours sets field value
func (o *SnapshotSchedule) SetPointInTimeWindowHours(v int32) {
	o.PointInTimeWindowHours = v
}

// GetSnapshotIntervalHours returns the SnapshotIntervalHours field value
func (o *SnapshotSchedule) GetSnapshotIntervalHours() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SnapshotIntervalHours
}

// GetSnapshotIntervalHoursOk returns a tuple with the SnapshotIntervalHours field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetSnapshotIntervalHoursOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotIntervalHours, true
}

// SetSnapshotIntervalHours sets field value
func (o *SnapshotSchedule) SetSnapshotIntervalHours(v int32) {
	o.SnapshotIntervalHours = v
}

// GetSnapshotRetentionDays returns the SnapshotRetentionDays field value
func (o *SnapshotSchedule) GetSnapshotRetentionDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SnapshotRetentionDays
}

// GetSnapshotRetentionDaysOk returns a tuple with the SnapshotRetentionDays field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetSnapshotRetentionDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotRetentionDays, true
}

// SetSnapshotRetentionDays sets field value
func (o *SnapshotSchedule) SetSnapshotRetentionDays(v int32) {
	o.SnapshotRetentionDays = v
}

// GetWeeklySnapshotRetentionWeeks returns the WeeklySnapshotRetentionWeeks field value
func (o *SnapshotSchedule) GetWeeklySnapshotRetentionWeeks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WeeklySnapshotRetentionWeeks
}

// GetWeeklySnapshotRetentionWeeksOk returns a tuple with the WeeklySnapshotRetentionWeeks field value
// and a boolean to check if the value has been set.
func (o *SnapshotSchedule) GetWeeklySnapshotRetentionWeeksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeeklySnapshotRetentionWeeks, true
}

// SetWeeklySnapshotRetentionWeeks sets field value
func (o *SnapshotSchedule) SetWeeklySnapshotRetentionWeeks(v int32) {
	o.WeeklySnapshotRetentionWeeks = v
}

func (o SnapshotSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterCheckpointIntervalMin"] = o.ClusterCheckpointIntervalMin
	toSerialize["clusterId"] = o.ClusterId
	toSerialize["dailySnapshotRetentionDays"] = o.DailySnapshotRetentionDays
	// skip: groupId is readOnly
	// skip: links is readOnly
	toSerialize["monthlySnapshotRetentionMonths"] = o.MonthlySnapshotRetentionMonths
	toSerialize["pointInTimeWindowHours"] = o.PointInTimeWindowHours
	toSerialize["snapshotIntervalHours"] = o.SnapshotIntervalHours
	toSerialize["snapshotRetentionDays"] = o.SnapshotRetentionDays
	toSerialize["weeklySnapshotRetentionWeeks"] = o.WeeklySnapshotRetentionWeeks
	return toSerialize, nil
}

type NullableSnapshotSchedule struct {
	value *SnapshotSchedule
	isSet bool
}

func (v NullableSnapshotSchedule) Get() *SnapshotSchedule {
	return v.value
}

func (v *NullableSnapshotSchedule) Set(val *SnapshotSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotSchedule(val *SnapshotSchedule) *NullableSnapshotSchedule {
	return &NullableSnapshotSchedule{value: val, isSet: true}
}

func (v NullableSnapshotSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


