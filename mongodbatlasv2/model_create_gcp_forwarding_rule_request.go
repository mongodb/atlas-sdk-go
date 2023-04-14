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

// checks if the CreateGCPForwardingRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGCPForwardingRuleRequest{}

// CreateGCPForwardingRuleRequest struct for CreateGCPForwardingRuleRequest
type CreateGCPForwardingRuleRequest struct {
	// Human-readable label that identifies the Google Cloud consumer forwarding rule that you created.
	EndpointName *string `json:"endpointName,omitempty"`
	// One Private Internet Protocol version 4 (IPv4) address to which this Google Cloud consumer forwarding rule resolves.
	IpAddress *string `json:"ipAddress,omitempty"`
}

// NewCreateGCPForwardingRuleRequest instantiates a new CreateGCPForwardingRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGCPForwardingRuleRequest() *CreateGCPForwardingRuleRequest {
	this := CreateGCPForwardingRuleRequest{}
	return &this
}

// NewCreateGCPForwardingRuleRequestWithDefaults instantiates a new CreateGCPForwardingRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGCPForwardingRuleRequestWithDefaults() *CreateGCPForwardingRuleRequest {
	this := CreateGCPForwardingRuleRequest{}
	return &this
}

// GetEndpointName returns the EndpointName field value if set, zero value otherwise.
func (o *CreateGCPForwardingRuleRequest) GetEndpointName() string {
	if o == nil || IsNil(o.EndpointName) {
		var ret string
		return ret
	}
	return *o.EndpointName
}

// GetEndpointNameOk returns a tuple with the EndpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPForwardingRuleRequest) GetEndpointNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointName) {
		return nil, false
	}
	return o.EndpointName, true
}

// HasEndpointName returns a boolean if a field has been set.
func (o *CreateGCPForwardingRuleRequest) HasEndpointName() bool {
	if o != nil && !IsNil(o.EndpointName) {
		return true
	}

	return false
}

// SetEndpointName gets a reference to the given string and assigns it to the EndpointName field.
func (o *CreateGCPForwardingRuleRequest) SetEndpointName(v string) {
	o.EndpointName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *CreateGCPForwardingRuleRequest) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPForwardingRuleRequest) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CreateGCPForwardingRuleRequest) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *CreateGCPForwardingRuleRequest) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o CreateGCPForwardingRuleRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateGCPForwardingRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointName) {
		toSerialize["endpointName"] = o.EndpointName
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	return toSerialize, nil
}

type NullableCreateGCPForwardingRuleRequest struct {
	value *CreateGCPForwardingRuleRequest
	isSet bool
}

func (v NullableCreateGCPForwardingRuleRequest) Get() *CreateGCPForwardingRuleRequest {
	return v.value
}

func (v *NullableCreateGCPForwardingRuleRequest) Set(val *CreateGCPForwardingRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGCPForwardingRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGCPForwardingRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGCPForwardingRuleRequest(val *CreateGCPForwardingRuleRequest) *NullableCreateGCPForwardingRuleRequest {
	return &NullableCreateGCPForwardingRuleRequest{value: val, isSet: true}
}

func (v NullableCreateGCPForwardingRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGCPForwardingRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


