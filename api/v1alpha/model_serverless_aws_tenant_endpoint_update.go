/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the ServerlessAWSTenantEndpointUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerlessAWSTenantEndpointUpdate{}

// ServerlessAWSTenantEndpointUpdate Updates to a serverless AWS tenant endpoint.
type ServerlessAWSTenantEndpointUpdate struct {
	// Unique string that identifies the private endpoint's network interface.
	CloudProviderEndpointId *string `json:"cloudProviderEndpointId,omitempty"`
	// Human-readable comment associated with the private endpoint.
	Comment *string `json:"comment,omitempty"`
	ProviderName string `json:"providerName"`
}

// NewServerlessAWSTenantEndpointUpdate instantiates a new ServerlessAWSTenantEndpointUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessAWSTenantEndpointUpdate(providerName string) *ServerlessAWSTenantEndpointUpdate {
	this := ServerlessAWSTenantEndpointUpdate{}
	this.ProviderName = providerName
	return &this
}

// NewServerlessAWSTenantEndpointUpdateWithDefaults instantiates a new ServerlessAWSTenantEndpointUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessAWSTenantEndpointUpdateWithDefaults() *ServerlessAWSTenantEndpointUpdate {
	this := ServerlessAWSTenantEndpointUpdate{}
	return &this
}

// GetCloudProviderEndpointId returns the CloudProviderEndpointId field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpointUpdate) GetCloudProviderEndpointId() string {
	if o == nil || IsNil(o.CloudProviderEndpointId) {
		var ret string
		return ret
	}
	return *o.CloudProviderEndpointId
}

// GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpointUpdate) GetCloudProviderEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProviderEndpointId) {
		return nil, false
	}
	return o.CloudProviderEndpointId, true
}

// HasCloudProviderEndpointId returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpointUpdate) HasCloudProviderEndpointId() bool {
	if o != nil && !IsNil(o.CloudProviderEndpointId) {
		return true
	}

	return false
}

// SetCloudProviderEndpointId gets a reference to the given string and assigns it to the CloudProviderEndpointId field.
func (o *ServerlessAWSTenantEndpointUpdate) SetCloudProviderEndpointId(v string) {
	o.CloudProviderEndpointId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpointUpdate) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpointUpdate) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpointUpdate) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ServerlessAWSTenantEndpointUpdate) SetComment(v string) {
	o.Comment = &v
}

// GetProviderName returns the ProviderName field value
func (o *ServerlessAWSTenantEndpointUpdate) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpointUpdate) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *ServerlessAWSTenantEndpointUpdate) SetProviderName(v string) {
	o.ProviderName = v
}

func (o ServerlessAWSTenantEndpointUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerlessAWSTenantEndpointUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProviderEndpointId) {
		toSerialize["cloudProviderEndpointId"] = o.CloudProviderEndpointId
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableServerlessAWSTenantEndpointUpdate struct {
	value *ServerlessAWSTenantEndpointUpdate
	isSet bool
}

func (v NullableServerlessAWSTenantEndpointUpdate) Get() *ServerlessAWSTenantEndpointUpdate {
	return v.value
}

func (v *NullableServerlessAWSTenantEndpointUpdate) Set(val *ServerlessAWSTenantEndpointUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessAWSTenantEndpointUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessAWSTenantEndpointUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessAWSTenantEndpointUpdate(val *ServerlessAWSTenantEndpointUpdate) *NullableServerlessAWSTenantEndpointUpdate {
	return &NullableServerlessAWSTenantEndpointUpdate{value: val, isSet: true}
}

func (v NullableServerlessAWSTenantEndpointUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessAWSTenantEndpointUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


