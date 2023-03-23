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

// checks if the SnapshotRetention type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotRetention{}

// SnapshotRetention struct for SnapshotRetention
type SnapshotRetention struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Quantity of time in which MongoDB Cloud measures snapshot retention.
	RetentionUnit string `json:"retentionUnit"`
	// Number that indicates the amount of days, weeks, or months that MongoDB Cloud retains the snapshot. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items. If the hourly policy item specifies a retention of two days, specify two days or greater for the retention of the weekly policy item.
	RetentionValue int32 `json:"retentionValue"`
}

// NewSnapshotRetention instantiates a new SnapshotRetention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRetention(retentionUnit string, retentionValue int32) *SnapshotRetention {
	this := SnapshotRetention{}
	this.RetentionUnit = retentionUnit
	this.RetentionValue = retentionValue
	return &this
}

// NewSnapshotRetentionWithDefaults instantiates a new SnapshotRetention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRetentionWithDefaults() *SnapshotRetention {
	this := SnapshotRetention{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SnapshotRetention) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRetention) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SnapshotRetention) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *SnapshotRetention) SetLinks(v []Link) {
	o.Links = v
}

// GetRetentionUnit returns the RetentionUnit field value
func (o *SnapshotRetention) GetRetentionUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetentionUnit
}

// GetRetentionUnitOk returns a tuple with the RetentionUnit field value
// and a boolean to check if the value has been set.
func (o *SnapshotRetention) GetRetentionUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionUnit, true
}

// SetRetentionUnit sets field value
func (o *SnapshotRetention) SetRetentionUnit(v string) {
	o.RetentionUnit = v
}

// GetRetentionValue returns the RetentionValue field value
func (o *SnapshotRetention) GetRetentionValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetentionValue
}

// GetRetentionValueOk returns a tuple with the RetentionValue field value
// and a boolean to check if the value has been set.
func (o *SnapshotRetention) GetRetentionValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionValue, true
}

// SetRetentionValue sets field value
func (o *SnapshotRetention) SetRetentionValue(v int32) {
	o.RetentionValue = v
}

func (o SnapshotRetention) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o SnapshotRetention) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["retentionUnit"] = o.RetentionUnit
	toSerialize["retentionValue"] = o.RetentionValue
	return toSerialize, nil
}

type NullableSnapshotRetention struct {
	value *SnapshotRetention
	isSet bool
}

func (v NullableSnapshotRetention) Get() *SnapshotRetention {
	return v.value
}

func (v *NullableSnapshotRetention) Set(val *SnapshotRetention) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRetention) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRetention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRetention(val *SnapshotRetention) *NullableSnapshotRetention {
	return &NullableSnapshotRetention{value: val, isSet: true}
}

func (v NullableSnapshotRetention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotRetention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

