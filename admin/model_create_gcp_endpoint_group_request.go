// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CreateGCPEndpointGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGCPEndpointGroupRequest{}

// CreateGCPEndpointGroupRequest Group of Private Endpoint settings.
type CreateGCPEndpointGroupRequest struct {
	// Human-readable label that identifies a set of endpoints.
	EndpointGroupName string `json:"endpointGroupName"`
	// List of individual private endpoints that comprise this endpoint group.
	Endpoints []CreateGCPForwardingRuleRequest `json:"endpoints,omitempty"`
	// Unique string that identifies the Google Cloud project in which you created the endpoints.
	GcpProjectId string `json:"gcpProjectId"`
}

// NewCreateGCPEndpointGroupRequest instantiates a new CreateGCPEndpointGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGCPEndpointGroupRequest(endpointGroupName string, gcpProjectId string) *CreateGCPEndpointGroupRequest {
	this := CreateGCPEndpointGroupRequest{}
	this.EndpointGroupName = endpointGroupName
	this.GcpProjectId = gcpProjectId
	return &this
}

// NewCreateGCPEndpointGroupRequestWithDefaults instantiates a new CreateGCPEndpointGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGCPEndpointGroupRequestWithDefaults() *CreateGCPEndpointGroupRequest {
	this := CreateGCPEndpointGroupRequest{}
	return &this
}

// GetEndpointGroupName returns the EndpointGroupName field value
func (o *CreateGCPEndpointGroupRequest) GetEndpointGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointGroupName
}

// GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field value
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequest) GetEndpointGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointGroupName, true
}

// SetEndpointGroupName sets field value
func (o *CreateGCPEndpointGroupRequest) SetEndpointGroupName(v string) {
	o.EndpointGroupName = v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *CreateGCPEndpointGroupRequest) GetEndpoints() []CreateGCPForwardingRuleRequest {
	if o == nil || IsNil(o.Endpoints) {
		var ret []CreateGCPForwardingRuleRequest
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequest) GetEndpointsOk() ([]CreateGCPForwardingRuleRequest, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *CreateGCPEndpointGroupRequest) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []CreateGCPForwardingRuleRequest and assigns it to the Endpoints field.
func (o *CreateGCPEndpointGroupRequest) SetEndpoints(v []CreateGCPForwardingRuleRequest) {
	o.Endpoints = v
}

// GetGcpProjectId returns the GcpProjectId field value
func (o *CreateGCPEndpointGroupRequest) GetGcpProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequest) GetGcpProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpProjectId, true
}

// SetGcpProjectId sets field value
func (o *CreateGCPEndpointGroupRequest) SetGcpProjectId(v string) {
	o.GcpProjectId = v
}

func (o CreateGCPEndpointGroupRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateGCPEndpointGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpointGroupName"] = o.EndpointGroupName
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	toSerialize["gcpProjectId"] = o.GcpProjectId
	return toSerialize, nil
}

type NullableCreateGCPEndpointGroupRequest struct {
	value *CreateGCPEndpointGroupRequest
	isSet bool
}

func (v NullableCreateGCPEndpointGroupRequest) Get() *CreateGCPEndpointGroupRequest {
	return v.value
}

func (v *NullableCreateGCPEndpointGroupRequest) Set(val *CreateGCPEndpointGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGCPEndpointGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGCPEndpointGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGCPEndpointGroupRequest(val *CreateGCPEndpointGroupRequest) *NullableCreateGCPEndpointGroupRequest {
	return &NullableCreateGCPEndpointGroupRequest{value: val, isSet: true}
}

func (v NullableCreateGCPEndpointGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGCPEndpointGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
