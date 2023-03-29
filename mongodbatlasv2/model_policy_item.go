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

// checks if the PolicyItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyItem{}

// PolicyItem Specifications for one policy.
type PolicyItem struct {
	// Number that indicates the frequency interval for a set of snapshots. A value of `1` specifies the first instance of the corresponding `frequencyType`.  - In a monthly policy item, `1` indicates that the monthly snapshot occurs on the first day of the month and `40` indicates the last day of the month.  - In a weekly policy item, `1` indicates that the weekly snapshot occurs on Monday and `7` indicates Sunday.  - In an hourly policy item, you can set the frequency interval to `1`, `2`, `4`, `6`, `8`, or `12`. For hourly policy items for NVMe clusters, MongoDB Cloud accepts only `12` as the frequency interval value.   MongoDB Cloud ignores this setting for non-hourly policy items in Backup Compliance Policy settings.
	FrequencyInterval int32 `json:"frequencyInterval"`
	// Human-readable label that identifies the frequency type associated with the backup policy.
	FrequencyType string `json:"frequencyType"`
	// Unique 24-hexadecimal digit string that identifies this backup policy item.
	Id *string `json:"id,omitempty"`
	// Unit of time in which MongoDB Cloud measures snapshot retention.
	RetentionUnit string `json:"retentionUnit"`
	// Duration in days, weeks, or months that MongoDB Cloud retains the snapshot. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items.  For example: If the hourly policy item specifies a retention of two days, you must specify two days or greater for the retention of the weekly policy item.
	RetentionValue int32 `json:"retentionValue"`
}

// NewPolicyItem instantiates a new PolicyItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyItem(frequencyInterval int32, frequencyType string, retentionUnit string, retentionValue int32) *PolicyItem {
	this := PolicyItem{}
	this.FrequencyInterval = frequencyInterval
	this.FrequencyType = frequencyType
	this.RetentionUnit = retentionUnit
	this.RetentionValue = retentionValue
	return &this
}

// NewPolicyItemWithDefaults instantiates a new PolicyItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyItemWithDefaults() *PolicyItem {
	this := PolicyItem{}
	return &this
}

// GetFrequencyInterval returns the FrequencyInterval field value
func (o *PolicyItem) GetFrequencyInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FrequencyInterval
}

// GetFrequencyIntervalOk returns a tuple with the FrequencyInterval field value
// and a boolean to check if the value has been set.
func (o *PolicyItem) GetFrequencyIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrequencyInterval, true
}

// SetFrequencyInterval sets field value
func (o *PolicyItem) SetFrequencyInterval(v int32) {
	o.FrequencyInterval = v
}

// GetFrequencyType returns the FrequencyType field value
func (o *PolicyItem) GetFrequencyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value
// and a boolean to check if the value has been set.
func (o *PolicyItem) GetFrequencyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrequencyType, true
}

// SetFrequencyType sets field value
func (o *PolicyItem) SetFrequencyType(v string) {
	o.FrequencyType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PolicyItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PolicyItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PolicyItem) SetId(v string) {
	o.Id = &v
}

// GetRetentionUnit returns the RetentionUnit field value
func (o *PolicyItem) GetRetentionUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetentionUnit
}

// GetRetentionUnitOk returns a tuple with the RetentionUnit field value
// and a boolean to check if the value has been set.
func (o *PolicyItem) GetRetentionUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionUnit, true
}

// SetRetentionUnit sets field value
func (o *PolicyItem) SetRetentionUnit(v string) {
	o.RetentionUnit = v
}

// GetRetentionValue returns the RetentionValue field value
func (o *PolicyItem) GetRetentionValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetentionValue
}

// GetRetentionValueOk returns a tuple with the RetentionValue field value
// and a boolean to check if the value has been set.
func (o *PolicyItem) GetRetentionValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionValue, true
}

// SetRetentionValue sets field value
func (o *PolicyItem) SetRetentionValue(v int32) {
	o.RetentionValue = v
}

func (o PolicyItem) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PolicyItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["frequencyInterval"] = o.FrequencyInterval
	toSerialize["frequencyType"] = o.FrequencyType
	toSerialize["retentionUnit"] = o.RetentionUnit
	toSerialize["retentionValue"] = o.RetentionValue
	return toSerialize, nil
}

type NullablePolicyItem struct {
	value *PolicyItem
	isSet bool
}

func (v NullablePolicyItem) Get() *PolicyItem {
	return v.value
}

func (v *NullablePolicyItem) Set(val *PolicyItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyItem(val *PolicyItem) *NullablePolicyItem {
	return &NullablePolicyItem{value: val, isSet: true}
}

func (v NullablePolicyItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


