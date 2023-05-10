// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the HipChatNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HipChatNotification{}

// HipChatNotification HipChat notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type HipChatNotification struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat. The resource requires this parameter when `\"notifications.[n].typeName\" : \"HIP_CHAT\"`\". If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	NotificationToken *string `json:"notificationToken,omitempty"`
	// HipChat API room name to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.[n].typeName\" : \"HIP_CHAT\"`\".
	RoomName *string `json:"roomName,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewHipChatNotification instantiates a new HipChatNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHipChatNotification(typeName string) *HipChatNotification {
	this := HipChatNotification{}
	this.TypeName = typeName
	return &this
}

// NewHipChatNotificationWithDefaults instantiates a new HipChatNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHipChatNotificationWithDefaults() *HipChatNotification {
	this := HipChatNotification{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *HipChatNotification) GetDelayMin() int32 {
	if o == nil || IsNil(o.DelayMin) {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChatNotification) GetDelayMinOk() (*int32, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *HipChatNotification) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *HipChatNotification) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *HipChatNotification) GetIntervalMin() int32 {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChatNotification) GetIntervalMinOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *HipChatNotification) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *HipChatNotification) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetNotificationToken returns the NotificationToken field value if set, zero value otherwise.
func (o *HipChatNotification) GetNotificationToken() string {
	if o == nil || IsNil(o.NotificationToken) {
		var ret string
		return ret
	}
	return *o.NotificationToken
}

// GetNotificationTokenOk returns a tuple with the NotificationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChatNotification) GetNotificationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationToken) {
		return nil, false
	}
	return o.NotificationToken, true
}

// HasNotificationToken returns a boolean if a field has been set.
func (o *HipChatNotification) HasNotificationToken() bool {
	if o != nil && !IsNil(o.NotificationToken) {
		return true
	}

	return false
}

// SetNotificationToken gets a reference to the given string and assigns it to the NotificationToken field.
func (o *HipChatNotification) SetNotificationToken(v string) {
	o.NotificationToken = &v
}

// GetRoomName returns the RoomName field value if set, zero value otherwise.
func (o *HipChatNotification) GetRoomName() string {
	if o == nil || IsNil(o.RoomName) {
		var ret string
		return ret
	}
	return *o.RoomName
}

// GetRoomNameOk returns a tuple with the RoomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChatNotification) GetRoomNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoomName) {
		return nil, false
	}
	return o.RoomName, true
}

// HasRoomName returns a boolean if a field has been set.
func (o *HipChatNotification) HasRoomName() bool {
	if o != nil && !IsNil(o.RoomName) {
		return true
	}

	return false
}

// SetRoomName gets a reference to the given string and assigns it to the RoomName field.
func (o *HipChatNotification) SetRoomName(v string) {
	o.RoomName = &v
}

// GetTypeName returns the TypeName field value
func (o *HipChatNotification) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *HipChatNotification) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *HipChatNotification) SetTypeName(v string) {
	o.TypeName = v
}

func (o HipChatNotification) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o HipChatNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if !IsNil(o.NotificationToken) {
		toSerialize["notificationToken"] = o.NotificationToken
	}
	if !IsNil(o.RoomName) {
		toSerialize["roomName"] = o.RoomName
	}
	toSerialize["typeName"] = o.TypeName
	return toSerialize, nil
}

type NullableHipChatNotification struct {
	value *HipChatNotification
	isSet bool
}

func (v NullableHipChatNotification) Get() *HipChatNotification {
	return v.value
}

func (v *NullableHipChatNotification) Set(val *HipChatNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableHipChatNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableHipChatNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHipChatNotification(val *HipChatNotification) *NullableHipChatNotification {
	return &NullableHipChatNotification{value: val, isSet: true}
}

func (v NullableHipChatNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHipChatNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
