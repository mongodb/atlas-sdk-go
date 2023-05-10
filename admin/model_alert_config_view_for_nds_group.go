// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"time"
)

// checks if the AlertConfigViewForNdsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertConfigViewForNdsGroup{}

// AlertConfigViewForNdsGroup struct for AlertConfigViewForNdsGroup
type AlertConfigViewForNdsGroup struct {
	// Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Flag that indicates whether someone enabled this alert configuration for the specified project.
	Enabled       *bool                             `json:"enabled,omitempty"`
	EventTypeName *ServerlessEventTypeViewAlertable `json:"eventTypeName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert configuration.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// No matchers are available for these alert types. The list is always empty.
	Matchers []map[string]interface{} `json:"matchers,omitempty"`
	// List that contains the targets that MongoDB Cloud sends notifications.
	Notifications []NotificationViewForNdsGroup `json:"notifications,omitempty"`
	// Date and time when someone last updated this alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated         *time.Time                 `json:"updated,omitempty"`
	MetricThreshold *ServerlessMetricThreshold `json:"metricThreshold,omitempty"`
	Threshold       *ThresholdViewInteger      `json:"threshold,omitempty"`
}

// NewAlertConfigViewForNdsGroup instantiates a new AlertConfigViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertConfigViewForNdsGroup() *AlertConfigViewForNdsGroup {
	this := AlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewAlertConfigViewForNdsGroupWithDefaults instantiates a new AlertConfigViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertConfigViewForNdsGroupWithDefaults() *AlertConfigViewForNdsGroup {
	this := AlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AlertConfigViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertConfigViewForNdsGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypeName returns the EventTypeName field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetEventTypeName() ServerlessEventTypeViewAlertable {
	if o == nil || IsNil(o.EventTypeName) {
		var ret ServerlessEventTypeViewAlertable
		return ret
	}
	return *o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetEventTypeNameOk() (*ServerlessEventTypeViewAlertable, bool) {
	if o == nil || IsNil(o.EventTypeName) {
		return nil, false
	}
	return o.EventTypeName, true
}

// HasEventTypeName returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasEventTypeName() bool {
	if o != nil && !IsNil(o.EventTypeName) {
		return true
	}

	return false
}

// SetEventTypeName gets a reference to the given ServerlessEventTypeViewAlertable and assigns it to the EventTypeName field.
func (o *AlertConfigViewForNdsGroup) SetEventTypeName(v ServerlessEventTypeViewAlertable) {
	o.EventTypeName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AlertConfigViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertConfigViewForNdsGroup) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *AlertConfigViewForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetMatchers() []map[string]interface{} {
	if o == nil || IsNil(o.Matchers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Matchers
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetMatchersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Matchers) {
		return nil, false
	}
	return o.Matchers, true
}

// HasMatchers returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasMatchers() bool {
	if o != nil && !IsNil(o.Matchers) {
		return true
	}

	return false
}

// SetMatchers gets a reference to the given []map[string]interface{} and assigns it to the Matchers field.
func (o *AlertConfigViewForNdsGroup) SetMatchers(v []map[string]interface{}) {
	o.Matchers = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup {
	if o == nil || IsNil(o.Notifications) {
		var ret []NotificationViewForNdsGroup
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetNotificationsOk() ([]NotificationViewForNdsGroup, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []NotificationViewForNdsGroup and assigns it to the Notifications field.
func (o *AlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup) {
	o.Notifications = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *AlertConfigViewForNdsGroup) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetMetricThreshold returns the MetricThreshold field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetMetricThreshold() ServerlessMetricThreshold {
	if o == nil || IsNil(o.MetricThreshold) {
		var ret ServerlessMetricThreshold
		return ret
	}
	return *o.MetricThreshold
}

// GetMetricThresholdOk returns a tuple with the MetricThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetMetricThresholdOk() (*ServerlessMetricThreshold, bool) {
	if o == nil || IsNil(o.MetricThreshold) {
		return nil, false
	}
	return o.MetricThreshold, true
}

// HasMetricThreshold returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasMetricThreshold() bool {
	if o != nil && !IsNil(o.MetricThreshold) {
		return true
	}

	return false
}

// SetMetricThreshold gets a reference to the given ServerlessMetricThreshold and assigns it to the MetricThreshold field.
func (o *AlertConfigViewForNdsGroup) SetMetricThreshold(v ServerlessMetricThreshold) {
	o.MetricThreshold = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *AlertConfigViewForNdsGroup) GetThreshold() ThresholdViewInteger {
	if o == nil || IsNil(o.Threshold) {
		var ret ThresholdViewInteger
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfigViewForNdsGroup) GetThresholdOk() (*ThresholdViewInteger, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *AlertConfigViewForNdsGroup) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given ThresholdViewInteger and assigns it to the Threshold field.
func (o *AlertConfigViewForNdsGroup) SetThreshold(v ThresholdViewInteger) {
	o.Threshold = &v
}

func (o AlertConfigViewForNdsGroup) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AlertConfigViewForNdsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EventTypeName) {
		toSerialize["eventTypeName"] = o.EventTypeName
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.MetricThreshold) {
		toSerialize["metricThreshold"] = o.MetricThreshold
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	return toSerialize, nil
}

type NullableAlertConfigViewForNdsGroup struct {
	value *AlertConfigViewForNdsGroup
	isSet bool
}

func (v NullableAlertConfigViewForNdsGroup) Get() *AlertConfigViewForNdsGroup {
	return v.value
}

func (v *NullableAlertConfigViewForNdsGroup) Set(val *AlertConfigViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertConfigViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertConfigViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertConfigViewForNdsGroup(val *AlertConfigViewForNdsGroup) *NullableAlertConfigViewForNdsGroup {
	return &NullableAlertConfigViewForNdsGroup{value: val, isSet: true}
}

func (v NullableAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertConfigViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}