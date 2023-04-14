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

// checks if the CreateAzureEndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAzureEndpointRequest{}

// CreateAzureEndpointRequest Group of Private Endpoint settings.
type CreateAzureEndpointRequest struct {
	// Unique string that identifies the private endpoint's network interface that someone added to this private endpoint service.
	Id *string `json:"id,omitempty"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIPAddress *string `json:"privateEndpointIPAddress,omitempty"`
}

// NewCreateAzureEndpointRequest instantiates a new CreateAzureEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAzureEndpointRequest() *CreateAzureEndpointRequest {
	this := CreateAzureEndpointRequest{}
	return &this
}

// NewCreateAzureEndpointRequestWithDefaults instantiates a new CreateAzureEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAzureEndpointRequestWithDefaults() *CreateAzureEndpointRequest {
	this := CreateAzureEndpointRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateAzureEndpointRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAzureEndpointRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateAzureEndpointRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateAzureEndpointRequest) SetId(v string) {
	o.Id = &v
}

// GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field value if set, zero value otherwise.
func (o *CreateAzureEndpointRequest) GetPrivateEndpointIPAddress() string {
	if o == nil || IsNil(o.PrivateEndpointIPAddress) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointIPAddress
}

// GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAzureEndpointRequest) GetPrivateEndpointIPAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointIPAddress) {
		return nil, false
	}
	return o.PrivateEndpointIPAddress, true
}

// HasPrivateEndpointIPAddress returns a boolean if a field has been set.
func (o *CreateAzureEndpointRequest) HasPrivateEndpointIPAddress() bool {
	if o != nil && !IsNil(o.PrivateEndpointIPAddress) {
		return true
	}

	return false
}

// SetPrivateEndpointIPAddress gets a reference to the given string and assigns it to the PrivateEndpointIPAddress field.
func (o *CreateAzureEndpointRequest) SetPrivateEndpointIPAddress(v string) {
	o.PrivateEndpointIPAddress = &v
}

func (o CreateAzureEndpointRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateAzureEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PrivateEndpointIPAddress) {
		toSerialize["privateEndpointIPAddress"] = o.PrivateEndpointIPAddress
	}
	return toSerialize, nil
}

type NullableCreateAzureEndpointRequest struct {
	value *CreateAzureEndpointRequest
	isSet bool
}

func (v NullableCreateAzureEndpointRequest) Get() *CreateAzureEndpointRequest {
	return v.value
}

func (v *NullableCreateAzureEndpointRequest) Set(val *CreateAzureEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAzureEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAzureEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAzureEndpointRequest(val *CreateAzureEndpointRequest) *NullableCreateAzureEndpointRequest {
	return &NullableCreateAzureEndpointRequest{value: val, isSet: true}
}

func (v NullableCreateAzureEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAzureEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


