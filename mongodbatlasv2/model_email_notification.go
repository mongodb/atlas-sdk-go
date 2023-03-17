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

// checks if the EmailNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailNotification{}

// EmailNotification Email notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type EmailNotification struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Email address to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.[n].typeName\" : \"EMAIL\"`. You don’t need to set this value to send emails to individual or groups of MongoDB Cloud users including:  - specific MongoDB Cloud users (`\"notifications.[n].typeName\" : \"USER\"`) - MongoDB Cloud users with specific project roles (`\"notifications.[n].typeName\" : \"GROUP\"`) - MongoDB Cloud users with specific organization roles (`\"notifications.[n].typeName\" : \"ORG\"`) - MongoDB Cloud teams (`\"notifications.[n].typeName\" : \"TEAM\"`)  To send emails to one MongoDB Cloud user or grouping of users, set the `notifications.[n].emailEnabled` parameter.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewEmailNotification instantiates a new EmailNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotification(typeName string) *EmailNotification {
	this := EmailNotification{}
	this.TypeName = typeName
	return &this
}

// NewEmailNotificationWithDefaults instantiates a new EmailNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationWithDefaults() *EmailNotification {
	this := EmailNotification{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *EmailNotification) GetDelayMin() int32 {
	if o == nil || IsNil(o.DelayMin) {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotification) GetDelayMinOk() (*int32, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *EmailNotification) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *EmailNotification) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *EmailNotification) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotification) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *EmailNotification) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *EmailNotification) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *EmailNotification) GetIntervalMin() int32 {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotification) GetIntervalMinOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *EmailNotification) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *EmailNotification) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetTypeName returns the TypeName field value
func (o *EmailNotification) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *EmailNotification) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *EmailNotification) SetTypeName(v string) {
	o.TypeName = v
}

func (o EmailNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	toSerialize["typeName"] = o.TypeName
	return toSerialize, nil
}

type NullableEmailNotification struct {
	value *EmailNotification
	isSet bool
}

func (v NullableEmailNotification) Get() *EmailNotification {
	return v.value
}

func (v *NullableEmailNotification) Set(val *EmailNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotification(val *EmailNotification) *NullableEmailNotification {
	return &NullableEmailNotification{value: val, isSet: true}
}

func (v NullableEmailNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


