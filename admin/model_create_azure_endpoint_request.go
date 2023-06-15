// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CreateAzureEndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAzureEndpointRequest{}

// CreateAzureEndpointRequest Group of Private Endpoint settings.
type CreateAzureEndpointRequest struct {
	// Unique string that identifies the private endpoint's network interface that someone added to this private endpoint service.
	Id string `json:"id"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIPAddress string `json:"privateEndpointIPAddress"`
}

// NewCreateAzureEndpointRequest instantiates a new CreateAzureEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAzureEndpointRequest(id string, privateEndpointIPAddress string) *CreateAzureEndpointRequest {
	this := CreateAzureEndpointRequest{}
	this.Id = id
	this.PrivateEndpointIPAddress = privateEndpointIPAddress
	return &this
}

// NewCreateAzureEndpointRequestWithDefaults instantiates a new CreateAzureEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAzureEndpointRequestWithDefaults() *CreateAzureEndpointRequest {
	this := CreateAzureEndpointRequest{}
	return &this
}

// GetId returns the Id field value
func (o *CreateAzureEndpointRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateAzureEndpointRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateAzureEndpointRequest) SetId(v string) {
	o.Id = v
}

// GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field value
func (o *CreateAzureEndpointRequest) GetPrivateEndpointIPAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateEndpointIPAddress
}

// GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field value
// and a boolean to check if the value has been set.
func (o *CreateAzureEndpointRequest) GetPrivateEndpointIPAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateEndpointIPAddress, true
}

// SetPrivateEndpointIPAddress sets field value
func (o *CreateAzureEndpointRequest) SetPrivateEndpointIPAddress(v string) {
	o.PrivateEndpointIPAddress = v
}

func (o CreateAzureEndpointRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateAzureEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["privateEndpointIPAddress"] = o.PrivateEndpointIPAddress
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
