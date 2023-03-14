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

// checks if the CreateAzureEndpointRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAzureEndpointRequestAllOf{}

// CreateAzureEndpointRequestAllOf struct for CreateAzureEndpointRequestAllOf
type CreateAzureEndpointRequestAllOf struct {
	// Unique string that identifies the private endpoint's network interface that someone added to this private endpoint service.
	Id *string `json:"id,omitempty"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIPAddress *string `json:"privateEndpointIPAddress,omitempty"`
}

// NewCreateAzureEndpointRequestAllOf instantiates a new CreateAzureEndpointRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAzureEndpointRequestAllOf() *CreateAzureEndpointRequestAllOf {
	this := CreateAzureEndpointRequestAllOf{}
	return &this
}

// NewCreateAzureEndpointRequestAllOfWithDefaults instantiates a new CreateAzureEndpointRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAzureEndpointRequestAllOfWithDefaults() *CreateAzureEndpointRequestAllOf {
	this := CreateAzureEndpointRequestAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateAzureEndpointRequestAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAzureEndpointRequestAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateAzureEndpointRequestAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateAzureEndpointRequestAllOf) SetId(v string) {
	o.Id = &v
}

// GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field value if set, zero value otherwise.
func (o *CreateAzureEndpointRequestAllOf) GetPrivateEndpointIPAddress() string {
	if o == nil || IsNil(o.PrivateEndpointIPAddress) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointIPAddress
}

// GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAzureEndpointRequestAllOf) GetPrivateEndpointIPAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointIPAddress) {
		return nil, false
	}
	return o.PrivateEndpointIPAddress, true
}

// HasPrivateEndpointIPAddress returns a boolean if a field has been set.
func (o *CreateAzureEndpointRequestAllOf) HasPrivateEndpointIPAddress() bool {
	if o != nil && !IsNil(o.PrivateEndpointIPAddress) {
		return true
	}

	return false
}

// SetPrivateEndpointIPAddress gets a reference to the given string and assigns it to the PrivateEndpointIPAddress field.
func (o *CreateAzureEndpointRequestAllOf) SetPrivateEndpointIPAddress(v string) {
	o.PrivateEndpointIPAddress = &v
}

func (o CreateAzureEndpointRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAzureEndpointRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PrivateEndpointIPAddress) {
		toSerialize["privateEndpointIPAddress"] = o.PrivateEndpointIPAddress
	}
	return toSerialize, nil
}

type NullableCreateAzureEndpointRequestAllOf struct {
	value *CreateAzureEndpointRequestAllOf
	isSet bool
}

func (v NullableCreateAzureEndpointRequestAllOf) Get() *CreateAzureEndpointRequestAllOf {
	return v.value
}

func (v *NullableCreateAzureEndpointRequestAllOf) Set(val *CreateAzureEndpointRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAzureEndpointRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAzureEndpointRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAzureEndpointRequestAllOf(val *CreateAzureEndpointRequestAllOf) *NullableCreateAzureEndpointRequestAllOf {
	return &NullableCreateAzureEndpointRequestAllOf{value: val, isSet: true}
}

func (v NullableCreateAzureEndpointRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAzureEndpointRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


