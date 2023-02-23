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

// checks if the AzurePeerNetworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzurePeerNetworkRequest{}

// AzurePeerNetworkRequest struct for AzurePeerNetworkRequest
type AzurePeerNetworkRequest struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
	ContainerId string `json:"containerId"`
	// Cloud service provider that determines the needed settings to configure the network connection for a virtual private connection.
	ProviderName string `json:"providerName"`
	// Unique string that identifies the Azure AD directory in which the VNet peered with the MongoDB Cloud VNet resides.
	AzureDirectoryId string `json:"azureDirectoryId"`
	// Unique string that identifies the Azure subscription in which the VNet you peered with the MongoDB Cloud VNet resides.
	AzureSubscriptionId string `json:"azureSubscriptionId"`
	// Error message returned when a requested Azure network peering resource returns `\"status\" : \"FAILED\"`. The resource returns `null` if the request succeeded.
	ErrorState *string `json:"errorState,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the network peering connection.
	Id *string `json:"id,omitempty"`
	// Human-readable label that identifies the resource group in which the VNet to peer with the MongoDB Cloud VNet resides.
	ResourceGroupName string `json:"resourceGroupName"`
	// State of the network peering connection at the time you made the request.
	Status *string `json:"status,omitempty"`
	// Human-readable label that identifies the VNet that you want to peer with the MongoDB Cloud VNet.
	VnetName string `json:"vnetName"`
}

// NewAzurePeerNetworkRequest instantiates a new AzurePeerNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurePeerNetworkRequest() *AzurePeerNetworkRequest {
	this := AzurePeerNetworkRequest{}
	return &this
}

// NewAzurePeerNetworkRequestWithDefaults instantiates a new AzurePeerNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurePeerNetworkRequestWithDefaults() *AzurePeerNetworkRequest {
	this := AzurePeerNetworkRequest{}
	return &this
}

// GetContainerId returns the ContainerId field value
func (o *AzurePeerNetworkRequest) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *AzurePeerNetworkRequest) SetContainerId(v string) {
	o.ContainerId = v
}

// GetProviderName returns the ProviderName field value
func (o *AzurePeerNetworkRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AzurePeerNetworkRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetAzureDirectoryId returns the AzureDirectoryId field value
func (o *AzurePeerNetworkRequest) GetAzureDirectoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureDirectoryId
}

// GetAzureDirectoryIdOk returns a tuple with the AzureDirectoryId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetAzureDirectoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureDirectoryId, true
}

// SetAzureDirectoryId sets field value
func (o *AzurePeerNetworkRequest) SetAzureDirectoryId(v string) {
	o.AzureDirectoryId = v
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value
func (o *AzurePeerNetworkRequest) GetAzureSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureSubscriptionId
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureSubscriptionId, true
}

// SetAzureSubscriptionId sets field value
func (o *AzurePeerNetworkRequest) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId = v
}

// GetErrorState returns the ErrorState field value if set, zero value otherwise.
func (o *AzurePeerNetworkRequest) GetErrorState() string {
	if o == nil || IsNil(o.ErrorState) {
		var ret string
		return ret
	}
	return *o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetErrorStateOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorState) {
		return nil, false
	}
	return o.ErrorState, true
}

// HasErrorState returns a boolean if a field has been set.
func (o *AzurePeerNetworkRequest) HasErrorState() bool {
	if o != nil && !IsNil(o.ErrorState) {
		return true
	}

	return false
}

// SetErrorState gets a reference to the given string and assigns it to the ErrorState field.
func (o *AzurePeerNetworkRequest) SetErrorState(v string) {
	o.ErrorState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzurePeerNetworkRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzurePeerNetworkRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzurePeerNetworkRequest) SetId(v string) {
	o.Id = &v
}

// GetResourceGroupName returns the ResourceGroupName field value
func (o *AzurePeerNetworkRequest) GetResourceGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGroupName
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetResourceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroupName, true
}

// SetResourceGroupName sets field value
func (o *AzurePeerNetworkRequest) SetResourceGroupName(v string) {
	o.ResourceGroupName = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AzurePeerNetworkRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AzurePeerNetworkRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AzurePeerNetworkRequest) SetStatus(v string) {
	o.Status = &v
}

// GetVnetName returns the VnetName field value
func (o *AzurePeerNetworkRequest) GetVnetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VnetName
}

// GetVnetNameOk returns a tuple with the VnetName field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetworkRequest) GetVnetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VnetName, true
}

// SetVnetName sets field value
func (o *AzurePeerNetworkRequest) SetVnetName(v string) {
	o.VnetName = v
}

func (o AzurePeerNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzurePeerNetworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerId"] = o.ContainerId
	toSerialize["providerName"] = o.ProviderName
	toSerialize["azureDirectoryId"] = o.AzureDirectoryId
	toSerialize["azureSubscriptionId"] = o.AzureSubscriptionId
	// skip: errorState is readOnly
	// skip: id is readOnly
	toSerialize["resourceGroupName"] = o.ResourceGroupName
	// skip: status is readOnly
	toSerialize["vnetName"] = o.VnetName
	return toSerialize, nil
}

type NullableAzurePeerNetworkRequest struct {
	value *AzurePeerNetworkRequest
	isSet bool
}

func (v NullableAzurePeerNetworkRequest) Get() *AzurePeerNetworkRequest {
	return v.value
}

func (v *NullableAzurePeerNetworkRequest) Set(val *AzurePeerNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurePeerNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurePeerNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurePeerNetworkRequest(val *AzurePeerNetworkRequest) *NullableAzurePeerNetworkRequest {
	return &NullableAzurePeerNetworkRequest{value: val, isSet: true}
}

func (v NullableAzurePeerNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurePeerNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


