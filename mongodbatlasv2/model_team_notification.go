/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the TeamNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamNotification{}

// TeamNotification Team notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type TeamNotification struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Flag that indicates whether MongoDB Cloud should send email notifications. The resource requires this parameter when one of the following values have been set:  - `\"notifications.[n].typeName\" : \"ORG\"` - `\"notifications.[n].typeName\" : \"GROUP\"` - `\"notifications.[n].typeName\" : \"USER\"`
	EmailEnabled *bool `json:"emailEnabled,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Flag that indicates whether MongoDB Cloud should send text message notifications. The resource requires this parameter when one of the following values have been set:  - `\"notifications.[n].typeName\" : \"ORG\"` - `\"notifications.[n].typeName\" : \"GROUP\"` - `\"notifications.[n].typeName\" : \"USER\"`
	SmsEnabled *bool `json:"smsEnabled,omitempty"`
	// Unique 24-hexadecimal digit string that identifies one MongoDB Cloud team. The resource requires this parameter when `\"notifications.[n].typeName\" : \"TEAM\"`.
	TeamId *string `json:"teamId,omitempty"`
	// Name of the MongoDB Cloud team that receives this notification. The resource requires this parameter when `\"notifications.[n].typeName\" : \"TEAM\"`.
	TeamName *string `json:"teamName,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewTeamNotification instantiates a new TeamNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamNotification(typeName string) *TeamNotification {
	this := TeamNotification{}
	this.TypeName = typeName
	return &this
}

// NewTeamNotificationWithDefaults instantiates a new TeamNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamNotificationWithDefaults() *TeamNotification {
	this := TeamNotification{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *TeamNotification) GetDelayMin() int32 {
	if o == nil || IsNil(o.DelayMin) {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotification) GetDelayMinOk() (*int32, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *TeamNotification) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *TeamNotification) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetEmailEnabled returns the EmailEnabled field value if set, zero value otherwise.
func (o *TeamNotification) GetEmailEnabled() bool {
	if o == nil || IsNil(o.EmailEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailEnabled
}

// GetEmailEnabledOk returns a tuple with the EmailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotification) GetEmailEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailEnabled) {
		return nil, false
	}
	return o.EmailEnabled, true
}

// HasEmailEnabled returns a boolean if a field has been set.
func (o *TeamNotification) HasEmailEnabled() bool {
	if o != nil && !IsNil(o.EmailEnabled) {
		return true
	}

	return false
}

// SetEmailEnabled gets a reference to the given bool and assigns it to the EmailEnabled field.
func (o *TeamNotification) SetEmailEnabled(v bool) {
	o.EmailEnabled = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *TeamNotification) GetIntervalMin() int32 {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotification) GetIntervalMinOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *TeamNotification) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *TeamNotification) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetSmsEnabled returns the SmsEnabled field value if set, zero value otherwise.
func (o *TeamNotification) GetSmsEnabled() bool {
	if o == nil || IsNil(o.SmsEnabled) {
		var ret bool
		return ret
	}
	return *o.SmsEnabled
}

// GetSmsEnabledOk returns a tuple with the SmsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotification) GetSmsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SmsEnabled) {
		return nil, false
	}
	return o.SmsEnabled, true
}

// HasSmsEnabled returns a boolean if a field has been set.
func (o *TeamNotification) HasSmsEnabled() bool {
	if o != nil && !IsNil(o.SmsEnabled) {
		return true
	}

	return false
}

// SetSmsEnabled gets a reference to the given bool and assigns it to the SmsEnabled field.
func (o *TeamNotification) SetSmsEnabled(v bool) {
	o.SmsEnabled = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *TeamNotification) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotification) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *TeamNotification) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *TeamNotification) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *TeamNotification) GetTeamName() string {
	if o == nil || IsNil(o.TeamName) {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotification) GetTeamNameOk() (*string, bool) {
	if o == nil || IsNil(o.TeamName) {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *TeamNotification) HasTeamName() bool {
	if o != nil && !IsNil(o.TeamName) {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *TeamNotification) SetTeamName(v string) {
	o.TeamName = &v
}

// GetTypeName returns the TypeName field value
func (o *TeamNotification) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *TeamNotification) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *TeamNotification) SetTypeName(v string) {
	o.TypeName = v
}

func (o TeamNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.EmailEnabled) {
		toSerialize["emailEnabled"] = o.EmailEnabled
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if !IsNil(o.SmsEnabled) {
		toSerialize["smsEnabled"] = o.SmsEnabled
	}
	if !IsNil(o.TeamId) {
		toSerialize["teamId"] = o.TeamId
	}
	if !IsNil(o.TeamName) {
		toSerialize["teamName"] = o.TeamName
	}
	toSerialize["typeName"] = o.TypeName
	return toSerialize, nil
}

type NullableTeamNotification struct {
	value *TeamNotification
	isSet bool
}

func (v NullableTeamNotification) Get() *TeamNotification {
	return v.value
}

func (v *NullableTeamNotification) Set(val *TeamNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamNotification(val *TeamNotification) *NullableTeamNotification {
	return &NullableTeamNotification{value: val, isSet: true}
}

func (v NullableTeamNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


