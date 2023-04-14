/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the VictorOpsNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VictorOpsNotification{}

// VictorOpsNotification VictorOps notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type VictorOpsNotification struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
	// API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when `\"notifications.[n].typeName\" : \"VICTOR_OPS\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	VictorOpsApiKey *string `json:"victorOpsApiKey,omitempty"`
	// Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when `\"notifications.[n].typeName\" : \"VICTOR_OPS\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	VictorOpsRoutingKey *string `json:"victorOpsRoutingKey,omitempty"`
}

// NewVictorOpsNotification instantiates a new VictorOpsNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVictorOpsNotification(typeName string) *VictorOpsNotification {
	this := VictorOpsNotification{}
	this.TypeName = typeName
	return &this
}

// NewVictorOpsNotificationWithDefaults instantiates a new VictorOpsNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVictorOpsNotificationWithDefaults() *VictorOpsNotification {
	this := VictorOpsNotification{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *VictorOpsNotification) GetDelayMin() int32 {
	if o == nil || IsNil(o.DelayMin) {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsNotification) GetDelayMinOk() (*int32, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *VictorOpsNotification) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *VictorOpsNotification) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *VictorOpsNotification) GetIntervalMin() int32 {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsNotification) GetIntervalMinOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *VictorOpsNotification) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *VictorOpsNotification) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetTypeName returns the TypeName field value
func (o *VictorOpsNotification) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *VictorOpsNotification) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *VictorOpsNotification) SetTypeName(v string) {
	o.TypeName = v
}

// GetVictorOpsApiKey returns the VictorOpsApiKey field value if set, zero value otherwise.
func (o *VictorOpsNotification) GetVictorOpsApiKey() string {
	if o == nil || IsNil(o.VictorOpsApiKey) {
		var ret string
		return ret
	}
	return *o.VictorOpsApiKey
}

// GetVictorOpsApiKeyOk returns a tuple with the VictorOpsApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsNotification) GetVictorOpsApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.VictorOpsApiKey) {
		return nil, false
	}
	return o.VictorOpsApiKey, true
}

// HasVictorOpsApiKey returns a boolean if a field has been set.
func (o *VictorOpsNotification) HasVictorOpsApiKey() bool {
	if o != nil && !IsNil(o.VictorOpsApiKey) {
		return true
	}

	return false
}

// SetVictorOpsApiKey gets a reference to the given string and assigns it to the VictorOpsApiKey field.
func (o *VictorOpsNotification) SetVictorOpsApiKey(v string) {
	o.VictorOpsApiKey = &v
}

// GetVictorOpsRoutingKey returns the VictorOpsRoutingKey field value if set, zero value otherwise.
func (o *VictorOpsNotification) GetVictorOpsRoutingKey() string {
	if o == nil || IsNil(o.VictorOpsRoutingKey) {
		var ret string
		return ret
	}
	return *o.VictorOpsRoutingKey
}

// GetVictorOpsRoutingKeyOk returns a tuple with the VictorOpsRoutingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsNotification) GetVictorOpsRoutingKeyOk() (*string, bool) {
	if o == nil || IsNil(o.VictorOpsRoutingKey) {
		return nil, false
	}
	return o.VictorOpsRoutingKey, true
}

// HasVictorOpsRoutingKey returns a boolean if a field has been set.
func (o *VictorOpsNotification) HasVictorOpsRoutingKey() bool {
	if o != nil && !IsNil(o.VictorOpsRoutingKey) {
		return true
	}

	return false
}

// SetVictorOpsRoutingKey gets a reference to the given string and assigns it to the VictorOpsRoutingKey field.
func (o *VictorOpsNotification) SetVictorOpsRoutingKey(v string) {
	o.VictorOpsRoutingKey = &v
}

func (o VictorOpsNotification) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o VictorOpsNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	toSerialize["typeName"] = o.TypeName
	if !IsNil(o.VictorOpsApiKey) {
		toSerialize["victorOpsApiKey"] = o.VictorOpsApiKey
	}
	if !IsNil(o.VictorOpsRoutingKey) {
		toSerialize["victorOpsRoutingKey"] = o.VictorOpsRoutingKey
	}
	return toSerialize, nil
}

type NullableVictorOpsNotification struct {
	value *VictorOpsNotification
	isSet bool
}

func (v NullableVictorOpsNotification) Get() *VictorOpsNotification {
	return v.value
}

func (v *NullableVictorOpsNotification) Set(val *VictorOpsNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableVictorOpsNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableVictorOpsNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVictorOpsNotification(val *VictorOpsNotification) *NullableVictorOpsNotification {
	return &NullableVictorOpsNotification{value: val, isSet: true}
}

func (v NullableVictorOpsNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVictorOpsNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


