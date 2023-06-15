// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AzurePeerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzurePeerNetwork{}

// AzurePeerNetwork Group of Network Peering connection settings.
type AzurePeerNetwork struct {
	// Unique string that identifies the Azure AD directory in which the VNet peered with the MongoDB Cloud VNet resides.
	AzureDirectoryId string `json:"azureDirectoryId"`
	// Unique string that identifies the Azure subscription in which the VNet you peered with the MongoDB Cloud VNet resides.
	AzureSubscriptionId string `json:"azureSubscriptionId"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
	ContainerId string `json:"containerId"`
	// Error message returned when a requested Azure network peering resource returns `\"status\" : \"FAILED\"`. The resource returns `null` if the request succeeded.
	ErrorState *string `json:"errorState,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the network peering connection.
	Id *string `json:"id,omitempty"`
	// Cloud service provider that serves the requested network peering connection.
	ProviderName *string `json:"providerName,omitempty"`
	// Human-readable label that identifies the resource group in which the VNet to peer with the MongoDB Cloud VNet resides.
	ResourceGroupName string `json:"resourceGroupName"`
	// State of the network peering connection at the time you made the request.
	Status *string `json:"status,omitempty"`
	// Human-readable label that identifies the VNet that you want to peer with the MongoDB Cloud VNet.
	VnetName string `json:"vnetName"`
}

// NewAzurePeerNetwork instantiates a new AzurePeerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurePeerNetwork(azureDirectoryId string, azureSubscriptionId string, containerId string, resourceGroupName string, vnetName string) *AzurePeerNetwork {
	this := AzurePeerNetwork{}
	this.AzureDirectoryId = azureDirectoryId
	this.AzureSubscriptionId = azureSubscriptionId
	this.ContainerId = containerId
	this.ResourceGroupName = resourceGroupName
	this.VnetName = vnetName
	return &this
}

// NewAzurePeerNetworkWithDefaults instantiates a new AzurePeerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurePeerNetworkWithDefaults() *AzurePeerNetwork {
	this := AzurePeerNetwork{}
	return &this
}

// GetAzureDirectoryId returns the AzureDirectoryId field value
func (o *AzurePeerNetwork) GetAzureDirectoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureDirectoryId
}

// GetAzureDirectoryIdOk returns a tuple with the AzureDirectoryId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetAzureDirectoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureDirectoryId, true
}

// SetAzureDirectoryId sets field value
func (o *AzurePeerNetwork) SetAzureDirectoryId(v string) {
	o.AzureDirectoryId = v
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value
func (o *AzurePeerNetwork) GetAzureSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureSubscriptionId
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureSubscriptionId, true
}

// SetAzureSubscriptionId sets field value
func (o *AzurePeerNetwork) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId = v
}

// GetContainerId returns the ContainerId field value
func (o *AzurePeerNetwork) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *AzurePeerNetwork) SetContainerId(v string) {
	o.ContainerId = v
}

// GetErrorState returns the ErrorState field value if set, zero value otherwise.
func (o *AzurePeerNetwork) GetErrorState() string {
	if o == nil || IsNil(o.ErrorState) {
		var ret string
		return ret
	}
	return *o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetErrorStateOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorState) {
		return nil, false
	}
	return o.ErrorState, true
}

// HasErrorState returns a boolean if a field has been set.
func (o *AzurePeerNetwork) HasErrorState() bool {
	if o != nil && !IsNil(o.ErrorState) {
		return true
	}

	return false
}

// SetErrorState gets a reference to the given string and assigns it to the ErrorState field.
func (o *AzurePeerNetwork) SetErrorState(v string) {
	o.ErrorState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzurePeerNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzurePeerNetwork) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzurePeerNetwork) SetId(v string) {
	o.Id = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *AzurePeerNetwork) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *AzurePeerNetwork) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *AzurePeerNetwork) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetResourceGroupName returns the ResourceGroupName field value
func (o *AzurePeerNetwork) GetResourceGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGroupName
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetResourceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroupName, true
}

// SetResourceGroupName sets field value
func (o *AzurePeerNetwork) SetResourceGroupName(v string) {
	o.ResourceGroupName = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AzurePeerNetwork) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AzurePeerNetwork) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AzurePeerNetwork) SetStatus(v string) {
	o.Status = &v
}

// GetVnetName returns the VnetName field value
func (o *AzurePeerNetwork) GetVnetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VnetName
}

// GetVnetNameOk returns a tuple with the VnetName field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetVnetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VnetName, true
}

// SetVnetName sets field value
func (o *AzurePeerNetwork) SetVnetName(v string) {
	o.VnetName = v
}

func (o AzurePeerNetwork) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzurePeerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["azureDirectoryId"] = o.AzureDirectoryId
	toSerialize["azureSubscriptionId"] = o.AzureSubscriptionId
	toSerialize["containerId"] = o.ContainerId
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	toSerialize["resourceGroupName"] = o.ResourceGroupName
	toSerialize["vnetName"] = o.VnetName
	return toSerialize, nil
}

type NullableAzurePeerNetwork struct {
	value *AzurePeerNetwork
	isSet bool
}

func (v NullableAzurePeerNetwork) Get() *AzurePeerNetwork {
	return v.value
}

func (v *NullableAzurePeerNetwork) Set(val *AzurePeerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurePeerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurePeerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurePeerNetwork(val *AzurePeerNetwork) *NullableAzurePeerNetwork {
	return &NullableAzurePeerNetwork{value: val, isSet: true}
}

func (v NullableAzurePeerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurePeerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
