// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CreateAWSEndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAWSEndpointRequest{}

// CreateAWSEndpointRequest Group of Private Endpoint settings.
type CreateAWSEndpointRequest struct {
	// Unique string that identifies the private endpoint's network interface that someone added to this private endpoint service.
	Id string `json:"id"`
}

// NewCreateAWSEndpointRequest instantiates a new CreateAWSEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAWSEndpointRequest(id string) *CreateAWSEndpointRequest {
	this := CreateAWSEndpointRequest{}
	this.Id = id
	return &this
}

// NewCreateAWSEndpointRequestWithDefaults instantiates a new CreateAWSEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAWSEndpointRequestWithDefaults() *CreateAWSEndpointRequest {
	this := CreateAWSEndpointRequest{}
	return &this
}

// GetId returns the Id field value
func (o *CreateAWSEndpointRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateAWSEndpointRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateAWSEndpointRequest) SetId(v string) {
	o.Id = v
}

func (o CreateAWSEndpointRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateAWSEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCreateAWSEndpointRequest struct {
	value *CreateAWSEndpointRequest
	isSet bool
}

func (v NullableCreateAWSEndpointRequest) Get() *CreateAWSEndpointRequest {
	return v.value
}

func (v *NullableCreateAWSEndpointRequest) Set(val *CreateAWSEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAWSEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAWSEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAWSEndpointRequest(val *CreateAWSEndpointRequest) *NullableCreateAWSEndpointRequest {
	return &NullableCreateAWSEndpointRequest{value: val, isSet: true}
}

func (v NullableCreateAWSEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAWSEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
