/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
	"time"
)

// checks if the HostMetricAlertConfigViewForNdsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostMetricAlertConfigViewForNdsGroup{}

// HostMetricAlertConfigViewForNdsGroup Host metric alert configuration allows to select which mongod host metrics trigger alerts and how users are notified.
type HostMetricAlertConfigViewForNdsGroup struct {
	// Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Flag that indicates whether someone enabled this alert configuration for the specified project.
	Enabled *bool `json:"enabled,omitempty"`
	EventTypeName HostMetricEventTypeViewAlertable `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert configuration.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster.
	Matchers []HostMatcher `json:"matchers,omitempty"`
	MetricThreshold *HostMetricThreshold `json:"metricThreshold,omitempty"`
	// List that contains the targets that MongoDB Cloud sends notifications.
	Notifications []NotificationViewForNdsGroup `json:"notifications,omitempty"`
	// Date and time when someone last updated this alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewHostMetricAlertConfigViewForNdsGroup instantiates a new HostMetricAlertConfigViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMetricAlertConfigViewForNdsGroup(eventTypeName HostMetricEventTypeViewAlertable) *HostMetricAlertConfigViewForNdsGroup {
	this := HostMetricAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	this.EventTypeName = eventTypeName
	return &this
}

// NewHostMetricAlertConfigViewForNdsGroupWithDefaults instantiates a new HostMetricAlertConfigViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMetricAlertConfigViewForNdsGroupWithDefaults() *HostMetricAlertConfigViewForNdsGroup {
	this := HostMetricAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypeName returns the EventTypeName field value
func (o *HostMetricAlertConfigViewForNdsGroup) GetEventTypeName() HostMetricEventTypeViewAlertable {
	if o == nil {
		var ret HostMetricEventTypeViewAlertable
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*HostMetricEventTypeViewAlertable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *HostMetricAlertConfigViewForNdsGroup) SetEventTypeName(v HostMetricEventTypeViewAlertable) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetMatchers() []HostMatcher {
	if o == nil || IsNil(o.Matchers) {
		var ret []HostMatcher
		return ret
	}
	return o.Matchers
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetMatchersOk() ([]HostMatcher, bool) {
	if o == nil || IsNil(o.Matchers) {
		return nil, false
	}
	return o.Matchers, true
}

// HasMatchers returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasMatchers() bool {
	if o != nil && !IsNil(o.Matchers) {
		return true
	}

	return false
}

// SetMatchers gets a reference to the given []HostMatcher and assigns it to the Matchers field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetMatchers(v []HostMatcher) {
	o.Matchers = v
}

// GetMetricThreshold returns the MetricThreshold field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetMetricThreshold() HostMetricThreshold {
	if o == nil || IsNil(o.MetricThreshold) {
		var ret HostMetricThreshold
		return ret
	}
	return *o.MetricThreshold
}

// GetMetricThresholdOk returns a tuple with the MetricThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetMetricThresholdOk() (*HostMetricThreshold, bool) {
	if o == nil || IsNil(o.MetricThreshold) {
		return nil, false
	}
	return o.MetricThreshold, true
}

// HasMetricThreshold returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasMetricThreshold() bool {
	if o != nil && !IsNil(o.MetricThreshold) {
		return true
	}

	return false
}

// SetMetricThreshold gets a reference to the given HostMetricThreshold and assigns it to the MetricThreshold field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetMetricThreshold(v HostMetricThreshold) {
	o.MetricThreshold = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup {
	if o == nil || IsNil(o.Notifications) {
		var ret []NotificationViewForNdsGroup
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetNotificationsOk() ([]NotificationViewForNdsGroup, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []NotificationViewForNdsGroup and assigns it to the Notifications field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup) {
	o.Notifications = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *HostMetricAlertConfigViewForNdsGroup) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *HostMetricAlertConfigViewForNdsGroup) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *HostMetricAlertConfigViewForNdsGroup) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o HostMetricAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostMetricAlertConfigViewForNdsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created is readOnly
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["eventTypeName"] = o.EventTypeName
	// skip: groupId is readOnly
	// skip: id is readOnly
	// skip: links is readOnly
	if !IsNil(o.Matchers) {
		toSerialize["matchers"] = o.Matchers
	}
	if !IsNil(o.MetricThreshold) {
		toSerialize["metricThreshold"] = o.MetricThreshold
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	// skip: updated is readOnly
	return toSerialize, nil
}

type NullableHostMetricAlertConfigViewForNdsGroup struct {
	value *HostMetricAlertConfigViewForNdsGroup
	isSet bool
}

func (v NullableHostMetricAlertConfigViewForNdsGroup) Get() *HostMetricAlertConfigViewForNdsGroup {
	return v.value
}

func (v *NullableHostMetricAlertConfigViewForNdsGroup) Set(val *HostMetricAlertConfigViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMetricAlertConfigViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableHostMetricAlertConfigViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMetricAlertConfigViewForNdsGroup(val *HostMetricAlertConfigViewForNdsGroup) *NullableHostMetricAlertConfigViewForNdsGroup {
	return &NullableHostMetricAlertConfigViewForNdsGroup{value: val, isSet: true}
}

func (v NullableHostMetricAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMetricAlertConfigViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


