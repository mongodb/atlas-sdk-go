/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the AppServiceMetricAlertConfigViewForNdsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppServiceMetricAlertConfigViewForNdsGroup{}

// AppServiceMetricAlertConfigViewForNdsGroup App Services metric alert configuration allows to select which app service metrics trigger alerts and how users are notified.
type AppServiceMetricAlertConfigViewForNdsGroup struct {
	// Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Flag that indicates whether someone enabled this alert configuration for the specified project.
	Enabled *bool `json:"enabled,omitempty"`
	EventTypeName AppServiceEventTypeViewAlertableWithThreshold `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert configuration.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster.
	Matchers []AppServiceMetricMatcher `json:"matchers,omitempty"`
	MetricThreshold *AppServiceMetricThreshold `json:"metricThreshold,omitempty"`
	// List that contains the targets that MongoDB Cloud sends notifications.
	Notifications []NotificationViewForNdsGroup `json:"notifications,omitempty"`
	// Date and time when someone last updated this alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewAppServiceMetricAlertConfigViewForNdsGroup instantiates a new AppServiceMetricAlertConfigViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppServiceMetricAlertConfigViewForNdsGroup(eventTypeName AppServiceEventTypeViewAlertableWithThreshold) *AppServiceMetricAlertConfigViewForNdsGroup {
	this := AppServiceMetricAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	this.EventTypeName = eventTypeName
	return &this
}

// NewAppServiceMetricAlertConfigViewForNdsGroupWithDefaults instantiates a new AppServiceMetricAlertConfigViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppServiceMetricAlertConfigViewForNdsGroupWithDefaults() *AppServiceMetricAlertConfigViewForNdsGroup {
	this := AppServiceMetricAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypeName returns the EventTypeName field value
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetEventTypeName() AppServiceEventTypeViewAlertableWithThreshold {
	if o == nil {
		var ret AppServiceEventTypeViewAlertableWithThreshold
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*AppServiceEventTypeViewAlertableWithThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetEventTypeName(v AppServiceEventTypeViewAlertableWithThreshold) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetMatchers() []AppServiceMetricMatcher {
	if o == nil || IsNil(o.Matchers) {
		var ret []AppServiceMetricMatcher
		return ret
	}
	return o.Matchers
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetMatchersOk() ([]AppServiceMetricMatcher, bool) {
	if o == nil || IsNil(o.Matchers) {
		return nil, false
	}
	return o.Matchers, true
}

// HasMatchers returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasMatchers() bool {
	if o != nil && !IsNil(o.Matchers) {
		return true
	}

	return false
}

// SetMatchers gets a reference to the given []AppServiceMetricMatcher and assigns it to the Matchers field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetMatchers(v []AppServiceMetricMatcher) {
	o.Matchers = v
}

// GetMetricThreshold returns the MetricThreshold field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetMetricThreshold() AppServiceMetricThreshold {
	if o == nil || IsNil(o.MetricThreshold) {
		var ret AppServiceMetricThreshold
		return ret
	}
	return *o.MetricThreshold
}

// GetMetricThresholdOk returns a tuple with the MetricThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetMetricThresholdOk() (*AppServiceMetricThreshold, bool) {
	if o == nil || IsNil(o.MetricThreshold) {
		return nil, false
	}
	return o.MetricThreshold, true
}

// HasMetricThreshold returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasMetricThreshold() bool {
	if o != nil && !IsNil(o.MetricThreshold) {
		return true
	}

	return false
}

// SetMetricThreshold gets a reference to the given AppServiceMetricThreshold and assigns it to the MetricThreshold field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetMetricThreshold(v AppServiceMetricThreshold) {
	o.MetricThreshold = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup {
	if o == nil || IsNil(o.Notifications) {
		var ret []NotificationViewForNdsGroup
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetNotificationsOk() ([]NotificationViewForNdsGroup, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []NotificationViewForNdsGroup and assigns it to the Notifications field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup) {
	o.Notifications = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o AppServiceMetricAlertConfigViewForNdsGroup) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AppServiceMetricAlertConfigViewForNdsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["eventTypeName"] = o.EventTypeName
	if !IsNil(o.Matchers) {
		toSerialize["matchers"] = o.Matchers
	}
	if !IsNil(o.MetricThreshold) {
		toSerialize["metricThreshold"] = o.MetricThreshold
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	return toSerialize, nil
}

type NullableAppServiceMetricAlertConfigViewForNdsGroup struct {
	value *AppServiceMetricAlertConfigViewForNdsGroup
	isSet bool
}

func (v NullableAppServiceMetricAlertConfigViewForNdsGroup) Get() *AppServiceMetricAlertConfigViewForNdsGroup {
	return v.value
}

func (v *NullableAppServiceMetricAlertConfigViewForNdsGroup) Set(val *AppServiceMetricAlertConfigViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAppServiceMetricAlertConfigViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAppServiceMetricAlertConfigViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppServiceMetricAlertConfigViewForNdsGroup(val *AppServiceMetricAlertConfigViewForNdsGroup) *NullableAppServiceMetricAlertConfigViewForNdsGroup {
	return &NullableAppServiceMetricAlertConfigViewForNdsGroup{value: val, isSet: true}
}

func (v NullableAppServiceMetricAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppServiceMetricAlertConfigViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


