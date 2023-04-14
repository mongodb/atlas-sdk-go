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

// checks if the CreateGCPEndpointGroupRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGCPEndpointGroupRequestAllOf{}

// CreateGCPEndpointGroupRequestAllOf struct for CreateGCPEndpointGroupRequestAllOf
type CreateGCPEndpointGroupRequestAllOf struct {
	// Human-readable label that identifies a set of endpoints.
	EndpointGroupName *string `json:"endpointGroupName,omitempty"`
	// List of individual private endpoints that comprise this endpoint group.
	Endpoints []CreateGCPForwardingRuleRequest `json:"endpoints,omitempty"`
	// Unique string that identifies the Google Cloud project in which you created the endpoints.
	GcpProjectId *string `json:"gcpProjectId,omitempty"`
}

// NewCreateGCPEndpointGroupRequestAllOf instantiates a new CreateGCPEndpointGroupRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGCPEndpointGroupRequestAllOf() *CreateGCPEndpointGroupRequestAllOf {
	this := CreateGCPEndpointGroupRequestAllOf{}
	return &this
}

// NewCreateGCPEndpointGroupRequestAllOfWithDefaults instantiates a new CreateGCPEndpointGroupRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGCPEndpointGroupRequestAllOfWithDefaults() *CreateGCPEndpointGroupRequestAllOf {
	this := CreateGCPEndpointGroupRequestAllOf{}
	return &this
}

// GetEndpointGroupName returns the EndpointGroupName field value if set, zero value otherwise.
func (o *CreateGCPEndpointGroupRequestAllOf) GetEndpointGroupName() string {
	if o == nil || IsNil(o.EndpointGroupName) {
		var ret string
		return ret
	}
	return *o.EndpointGroupName
}

// GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequestAllOf) GetEndpointGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointGroupName) {
		return nil, false
	}
	return o.EndpointGroupName, true
}

// HasEndpointGroupName returns a boolean if a field has been set.
func (o *CreateGCPEndpointGroupRequestAllOf) HasEndpointGroupName() bool {
	if o != nil && !IsNil(o.EndpointGroupName) {
		return true
	}

	return false
}

// SetEndpointGroupName gets a reference to the given string and assigns it to the EndpointGroupName field.
func (o *CreateGCPEndpointGroupRequestAllOf) SetEndpointGroupName(v string) {
	o.EndpointGroupName = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *CreateGCPEndpointGroupRequestAllOf) GetEndpoints() []CreateGCPForwardingRuleRequest {
	if o == nil || IsNil(o.Endpoints) {
		var ret []CreateGCPForwardingRuleRequest
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequestAllOf) GetEndpointsOk() ([]CreateGCPForwardingRuleRequest, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *CreateGCPEndpointGroupRequestAllOf) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []CreateGCPForwardingRuleRequest and assigns it to the Endpoints field.
func (o *CreateGCPEndpointGroupRequestAllOf) SetEndpoints(v []CreateGCPForwardingRuleRequest) {
	o.Endpoints = v
}

// GetGcpProjectId returns the GcpProjectId field value if set, zero value otherwise.
func (o *CreateGCPEndpointGroupRequestAllOf) GetGcpProjectId() string {
	if o == nil || IsNil(o.GcpProjectId) {
		var ret string
		return ret
	}
	return *o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequestAllOf) GetGcpProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectId) {
		return nil, false
	}
	return o.GcpProjectId, true
}

// HasGcpProjectId returns a boolean if a field has been set.
func (o *CreateGCPEndpointGroupRequestAllOf) HasGcpProjectId() bool {
	if o != nil && !IsNil(o.GcpProjectId) {
		return true
	}

	return false
}

// SetGcpProjectId gets a reference to the given string and assigns it to the GcpProjectId field.
func (o *CreateGCPEndpointGroupRequestAllOf) SetGcpProjectId(v string) {
	o.GcpProjectId = &v
}

func (o CreateGCPEndpointGroupRequestAllOf) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateGCPEndpointGroupRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointGroupName) {
		toSerialize["endpointGroupName"] = o.EndpointGroupName
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.GcpProjectId) {
		toSerialize["gcpProjectId"] = o.GcpProjectId
	}
	return toSerialize, nil
}

type NullableCreateGCPEndpointGroupRequestAllOf struct {
	value *CreateGCPEndpointGroupRequestAllOf
	isSet bool
}

func (v NullableCreateGCPEndpointGroupRequestAllOf) Get() *CreateGCPEndpointGroupRequestAllOf {
	return v.value
}

func (v *NullableCreateGCPEndpointGroupRequestAllOf) Set(val *CreateGCPEndpointGroupRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGCPEndpointGroupRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGCPEndpointGroupRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGCPEndpointGroupRequestAllOf(val *CreateGCPEndpointGroupRequestAllOf) *NullableCreateGCPEndpointGroupRequestAllOf {
	return &NullableCreateGCPEndpointGroupRequestAllOf{value: val, isSet: true}
}

func (v NullableCreateGCPEndpointGroupRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGCPEndpointGroupRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


