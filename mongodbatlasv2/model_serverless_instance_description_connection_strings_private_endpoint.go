/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint{}

// ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint Private endpoint connection string that you can use to connect to this serverless instance through a private endpoint.
type ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint struct {
	// List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].srvConnectionString**.
	Endpoints []ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint `json:"endpoints,omitempty"`
	// Private endpoint-aware connection string that uses the `mongodb+srv://` protocol to connect to MongoDB Cloud through a private endpoint. The `mongodb+srv` protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS).
	SrvConnectionString *string `json:"srvConnectionString,omitempty"`
	// MongoDB process type to which your application connects.
	Type *string `json:"type,omitempty"`
}

// NewServerlessInstanceDescriptionConnectionStringsPrivateEndpoint instantiates a new ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessInstanceDescriptionConnectionStringsPrivateEndpoint() *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint {
	this := ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint{}
	return &this
}

// NewServerlessInstanceDescriptionConnectionStringsPrivateEndpointWithDefaults instantiates a new ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessInstanceDescriptionConnectionStringsPrivateEndpointWithDefaults() *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint {
	this := ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint{}
	return &this
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetEndpoints() []ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint {
	if o == nil || IsNil(o.Endpoints) {
		var ret []ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetEndpointsOk() ([]ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint and assigns it to the Endpoints field.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) SetEndpoints(v []ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint) {
	o.Endpoints = v
}

// GetSrvConnectionString returns the SrvConnectionString field value if set, zero value otherwise.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetSrvConnectionString() string {
	if o == nil || IsNil(o.SrvConnectionString) {
		var ret string
		return ret
	}
	return *o.SrvConnectionString
}

// GetSrvConnectionStringOk returns a tuple with the SrvConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetSrvConnectionStringOk() (*string, bool) {
	if o == nil || IsNil(o.SrvConnectionString) {
		return nil, false
	}
	return o.SrvConnectionString, true
}

// HasSrvConnectionString returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) HasSrvConnectionString() bool {
	if o != nil && !IsNil(o.SrvConnectionString) {
		return true
	}

	return false
}

// SetSrvConnectionString gets a reference to the given string and assigns it to the SrvConnectionString field.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) SetSrvConnectionString(v string) {
	o.SrvConnectionString = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) SetType(v string) {
	o.Type = &v
}

func (o ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint struct {
	value *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint
	isSet bool
}

func (v NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) Get() *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint {
	return v.value
}

func (v *NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) Set(val *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint(val *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) *NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint {
	return &NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint{value: val, isSet: true}
}

func (v NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


