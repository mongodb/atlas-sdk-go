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

// checks if the CreateEndpointServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEndpointServiceRequest{}

// CreateEndpointServiceRequest struct for CreateEndpointServiceRequest
type CreateEndpointServiceRequest struct {
	// Human-readable label that identifies the cloud service provider for which you want to create the private endpoint service.
	ProviderName string `json:"providerName"`
	// Cloud provider region in which you want to create the private endpoint service. Regions accepted as values differ for [Amazon Web Services](https://docs.atlas.mongodb.com/reference/amazon-aws/), [Google Cloud Platform](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Microsoft Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
	Region string `json:"region"`
}

// NewCreateEndpointServiceRequest instantiates a new CreateEndpointServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEndpointServiceRequest(providerName string, region string) *CreateEndpointServiceRequest {
	this := CreateEndpointServiceRequest{}
	this.ProviderName = providerName
	this.Region = region
	return &this
}

// NewCreateEndpointServiceRequestWithDefaults instantiates a new CreateEndpointServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEndpointServiceRequestWithDefaults() *CreateEndpointServiceRequest {
	this := CreateEndpointServiceRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *CreateEndpointServiceRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *CreateEndpointServiceRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *CreateEndpointServiceRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetRegion returns the Region field value
func (o *CreateEndpointServiceRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CreateEndpointServiceRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CreateEndpointServiceRequest) SetRegion(v string) {
	o.Region = v
}

func (o CreateEndpointServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEndpointServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providerName"] = o.ProviderName
	toSerialize["region"] = o.Region
	return toSerialize, nil
}

type NullableCreateEndpointServiceRequest struct {
	value *CreateEndpointServiceRequest
	isSet bool
}

func (v NullableCreateEndpointServiceRequest) Get() *CreateEndpointServiceRequest {
	return v.value
}

func (v *NullableCreateEndpointServiceRequest) Set(val *CreateEndpointServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEndpointServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEndpointServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEndpointServiceRequest(val *CreateEndpointServiceRequest) *NullableCreateEndpointServiceRequest {
	return &NullableCreateEndpointServiceRequest{value: val, isSet: true}
}

func (v NullableCreateEndpointServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEndpointServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


