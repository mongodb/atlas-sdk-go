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

// checks if the DatadogNotificationView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatadogNotificationView{}

// DatadogNotificationView Datadog notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type DatadogNotificationView struct {
	// Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog. You can find this API key in the Datadog dashboard. The resource requires this parameter when `\"notifications.[n].typeName\" : \"DATADOG\"`.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	DatadogApiKey *string `json:"datadogApiKey,omitempty"`
	// Datadog region that indicates which API Uniform Resource Locator (URL) to use. The resource requires this parameter when `\"notifications.[n].typeName\" : \"DATADOG\"`.  To learn more about Datadog's regions, see <a href=\"https://docs.datadoghq.com/getting_started/site/\" target=\"_blank\" rel=\"noopener noreferrer\">Datadog Sites</a>.
	DatadogRegion *string `json:"datadogRegion,omitempty"`
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewDatadogNotificationView instantiates a new DatadogNotificationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadogNotificationView(typeName string) *DatadogNotificationView {
	this := DatadogNotificationView{}
	var datadogRegion string = "US"
	this.DatadogRegion = &datadogRegion
	this.TypeName = typeName
	return &this
}

// NewDatadogNotificationViewWithDefaults instantiates a new DatadogNotificationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadogNotificationViewWithDefaults() *DatadogNotificationView {
	this := DatadogNotificationView{}
	var datadogRegion string = "US"
	this.DatadogRegion = &datadogRegion
	return &this
}

// GetDatadogApiKey returns the DatadogApiKey field value if set, zero value otherwise.
func (o *DatadogNotificationView) GetDatadogApiKey() string {
	if o == nil || IsNil(o.DatadogApiKey) {
		var ret string
		return ret
	}
	return *o.DatadogApiKey
}

// GetDatadogApiKeyOk returns a tuple with the DatadogApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogNotificationView) GetDatadogApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DatadogApiKey) {
		return nil, false
	}
	return o.DatadogApiKey, true
}

// HasDatadogApiKey returns a boolean if a field has been set.
func (o *DatadogNotificationView) HasDatadogApiKey() bool {
	if o != nil && !IsNil(o.DatadogApiKey) {
		return true
	}

	return false
}

// SetDatadogApiKey gets a reference to the given string and assigns it to the DatadogApiKey field.
func (o *DatadogNotificationView) SetDatadogApiKey(v string) {
	o.DatadogApiKey = &v
}

// GetDatadogRegion returns the DatadogRegion field value if set, zero value otherwise.
func (o *DatadogNotificationView) GetDatadogRegion() string {
	if o == nil || IsNil(o.DatadogRegion) {
		var ret string
		return ret
	}
	return *o.DatadogRegion
}

// GetDatadogRegionOk returns a tuple with the DatadogRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogNotificationView) GetDatadogRegionOk() (*string, bool) {
	if o == nil || IsNil(o.DatadogRegion) {
		return nil, false
	}
	return o.DatadogRegion, true
}

// HasDatadogRegion returns a boolean if a field has been set.
func (o *DatadogNotificationView) HasDatadogRegion() bool {
	if o != nil && !IsNil(o.DatadogRegion) {
		return true
	}

	return false
}

// SetDatadogRegion gets a reference to the given string and assigns it to the DatadogRegion field.
func (o *DatadogNotificationView) SetDatadogRegion(v string) {
	o.DatadogRegion = &v
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *DatadogNotificationView) GetDelayMin() int32 {
	if o == nil || IsNil(o.DelayMin) {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogNotificationView) GetDelayMinOk() (*int32, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *DatadogNotificationView) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *DatadogNotificationView) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *DatadogNotificationView) GetIntervalMin() int32 {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogNotificationView) GetIntervalMinOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *DatadogNotificationView) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *DatadogNotificationView) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetTypeName returns the TypeName field value
func (o *DatadogNotificationView) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *DatadogNotificationView) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *DatadogNotificationView) SetTypeName(v string) {
	o.TypeName = v
}

func (o DatadogNotificationView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatadogNotificationView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatadogApiKey) {
		toSerialize["datadogApiKey"] = o.DatadogApiKey
	}
	if !IsNil(o.DatadogRegion) {
		toSerialize["datadogRegion"] = o.DatadogRegion
	}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	toSerialize["typeName"] = o.TypeName
	return toSerialize, nil
}

type NullableDatadogNotificationView struct {
	value *DatadogNotificationView
	isSet bool
}

func (v NullableDatadogNotificationView) Get() *DatadogNotificationView {
	return v.value
}

func (v *NullableDatadogNotificationView) Set(val *DatadogNotificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadogNotificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadogNotificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadogNotificationView(val *DatadogNotificationView) *NullableDatadogNotificationView {
	return &NullableDatadogNotificationView{value: val, isSet: true}
}

func (v NullableDatadogNotificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadogNotificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


