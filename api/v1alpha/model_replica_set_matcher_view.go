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

// checks if the ReplicaSetMatcherView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicaSetMatcherView{}

// ReplicaSetMatcherView Rules to apply when comparing an replica set against this alert configuration.
type ReplicaSetMatcherView struct {
	FieldName *ReplicaSetMatcherField `json:"fieldName,omitempty"`
	// Comparison operator to apply when checking the current metric value against **matcher[n].value**.
	Operator *string `json:"operator,omitempty"`
	// Value to match or exceed using the specified **matchers.operator**.
	Value *string `json:"value,omitempty"`
}

// NewReplicaSetMatcherView instantiates a new ReplicaSetMatcherView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaSetMatcherView() *ReplicaSetMatcherView {
	this := ReplicaSetMatcherView{}
	return &this
}

// NewReplicaSetMatcherViewWithDefaults instantiates a new ReplicaSetMatcherView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaSetMatcherViewWithDefaults() *ReplicaSetMatcherView {
	this := ReplicaSetMatcherView{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *ReplicaSetMatcherView) GetFieldName() ReplicaSetMatcherField {
	if o == nil || IsNil(o.FieldName) {
		var ret ReplicaSetMatcherField
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetMatcherView) GetFieldNameOk() (*ReplicaSetMatcherField, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *ReplicaSetMatcherView) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given ReplicaSetMatcherField and assigns it to the FieldName field.
func (o *ReplicaSetMatcherView) SetFieldName(v ReplicaSetMatcherField) {
	o.FieldName = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ReplicaSetMatcherView) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetMatcherView) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ReplicaSetMatcherView) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ReplicaSetMatcherView) SetOperator(v string) {
	o.Operator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ReplicaSetMatcherView) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetMatcherView) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ReplicaSetMatcherView) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ReplicaSetMatcherView) SetValue(v string) {
	o.Value = &v
}

func (o ReplicaSetMatcherView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicaSetMatcherView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldName) {
		toSerialize["fieldName"] = o.FieldName
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableReplicaSetMatcherView struct {
	value *ReplicaSetMatcherView
	isSet bool
}

func (v NullableReplicaSetMatcherView) Get() *ReplicaSetMatcherView {
	return v.value
}

func (v *NullableReplicaSetMatcherView) Set(val *ReplicaSetMatcherView) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaSetMatcherView) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaSetMatcherView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaSetMatcherView(val *ReplicaSetMatcherView) *NullableReplicaSetMatcherView {
	return &NullableReplicaSetMatcherView{value: val, isSet: true}
}

func (v NullableReplicaSetMatcherView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaSetMatcherView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


