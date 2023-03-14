/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
)

// checks if the CreateAWSEndpointRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAWSEndpointRequestAllOf{}

// CreateAWSEndpointRequestAllOf struct for CreateAWSEndpointRequestAllOf
type CreateAWSEndpointRequestAllOf struct {
	// Unique string that identifies the private endpoint's network interface that someone added to this private endpoint service.
	Id *string `json:"id,omitempty"`
}

// NewCreateAWSEndpointRequestAllOf instantiates a new CreateAWSEndpointRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAWSEndpointRequestAllOf() *CreateAWSEndpointRequestAllOf {
	this := CreateAWSEndpointRequestAllOf{}
	return &this
}

// NewCreateAWSEndpointRequestAllOfWithDefaults instantiates a new CreateAWSEndpointRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAWSEndpointRequestAllOfWithDefaults() *CreateAWSEndpointRequestAllOf {
	this := CreateAWSEndpointRequestAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateAWSEndpointRequestAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAWSEndpointRequestAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateAWSEndpointRequestAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateAWSEndpointRequestAllOf) SetId(v string) {
	o.Id = &v
}

func (o CreateAWSEndpointRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAWSEndpointRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateAWSEndpointRequestAllOf struct {
	value *CreateAWSEndpointRequestAllOf
	isSet bool
}

func (v NullableCreateAWSEndpointRequestAllOf) Get() *CreateAWSEndpointRequestAllOf {
	return v.value
}

func (v *NullableCreateAWSEndpointRequestAllOf) Set(val *CreateAWSEndpointRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAWSEndpointRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAWSEndpointRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAWSEndpointRequestAllOf(val *CreateAWSEndpointRequestAllOf) *NullableCreateAWSEndpointRequestAllOf {
	return &NullableCreateAWSEndpointRequestAllOf{value: val, isSet: true}
}

func (v NullableCreateAWSEndpointRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAWSEndpointRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


