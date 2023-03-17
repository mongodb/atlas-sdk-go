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

// checks if the NamespaceObj type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceObj{}

// NamespaceObj Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as `<database>.<collection>`.
type NamespaceObj struct {
	// Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as `<database>.<collection>`.
	Namespace *string `json:"namespace,omitempty"`
	// Human-readable label that identifies the type of namespace.
	Type *string `json:"type,omitempty"`
}

// NewNamespaceObj instantiates a new NamespaceObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceObj() *NamespaceObj {
	this := NamespaceObj{}
	return &this
}

// NewNamespaceObjWithDefaults instantiates a new NamespaceObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceObjWithDefaults() *NamespaceObj {
	this := NamespaceObj{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *NamespaceObj) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceObj) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *NamespaceObj) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *NamespaceObj) SetNamespace(v string) {
	o.Namespace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NamespaceObj) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceObj) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NamespaceObj) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NamespaceObj) SetType(v string) {
	o.Type = &v
}

func (o NamespaceObj) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceObj) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: namespace is readOnly
	// skip: type is readOnly
	return toSerialize, nil
}

type NullableNamespaceObj struct {
	value *NamespaceObj
	isSet bool
}

func (v NullableNamespaceObj) Get() *NamespaceObj {
	return v.value
}

func (v *NullableNamespaceObj) Set(val *NamespaceObj) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceObj) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceObj(val *NamespaceObj) *NullableNamespaceObj {
	return &NullableNamespaceObj{value: val, isSet: true}
}

func (v NullableNamespaceObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


