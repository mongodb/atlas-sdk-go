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

// checks if the AzureCloudProviderContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureCloudProviderContainer{}

// AzureCloudProviderContainer Collection of settings that configures the network container for a virtual private connection on Amazon Web Services.
type AzureCloudProviderContainer struct {
	// IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project's clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. AWS and Azure further limit the block to between the `/24` and  `/21` ranges.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Amazon Web Services (AWS) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of `/24` equals 27 three-node replica sets.
	AtlasCidrBlock string `json:"atlasCidrBlock"`
	// Unique string that identifies the Azure subscription in which the MongoDB Cloud VNet resides.
	AzureSubscriptionId *string `json:"azureSubscriptionId,omitempty"`
	// Azure region to which MongoDB Cloud deployed this network peering container.
	Region string `json:"region"`
	// Unique string that identifies the Azure VNet in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container.
	VnetName *string `json:"vnetName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the network peering container.
	Id *string `json:"id,omitempty"`
	// Cloud service provider that serves the requested network peering containers.
	ProviderName *string `json:"providerName,omitempty"`
	// Flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container.
	Provisioned *bool `json:"provisioned,omitempty"`
}

// NewAzureCloudProviderContainer instantiates a new AzureCloudProviderContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCloudProviderContainer(atlasCidrBlock string, region string) *AzureCloudProviderContainer {
	this := AzureCloudProviderContainer{}
	this.AtlasCidrBlock = atlasCidrBlock
	this.Region = region
	return &this
}

// NewAzureCloudProviderContainerWithDefaults instantiates a new AzureCloudProviderContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCloudProviderContainerWithDefaults() *AzureCloudProviderContainer {
	this := AzureCloudProviderContainer{}
	return &this
}

// GetAtlasCidrBlock returns the AtlasCidrBlock field value
func (o *AzureCloudProviderContainer) GetAtlasCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AtlasCidrBlock
}

// GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field value
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetAtlasCidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AtlasCidrBlock, true
}

// SetAtlasCidrBlock sets field value
func (o *AzureCloudProviderContainer) SetAtlasCidrBlock(v string) {
	o.AtlasCidrBlock = v
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value if set, zero value otherwise.
func (o *AzureCloudProviderContainer) GetAzureSubscriptionId() string {
	if o == nil || IsNil(o.AzureSubscriptionId) {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionId
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AzureSubscriptionId) {
		return nil, false
	}
	return o.AzureSubscriptionId, true
}

// HasAzureSubscriptionId returns a boolean if a field has been set.
func (o *AzureCloudProviderContainer) HasAzureSubscriptionId() bool {
	if o != nil && !IsNil(o.AzureSubscriptionId) {
		return true
	}

	return false
}

// SetAzureSubscriptionId gets a reference to the given string and assigns it to the AzureSubscriptionId field.
func (o *AzureCloudProviderContainer) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId = &v
}

// GetRegion returns the Region field value
func (o *AzureCloudProviderContainer) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AzureCloudProviderContainer) SetRegion(v string) {
	o.Region = v
}

// GetVnetName returns the VnetName field value if set, zero value otherwise.
func (o *AzureCloudProviderContainer) GetVnetName() string {
	if o == nil || IsNil(o.VnetName) {
		var ret string
		return ret
	}
	return *o.VnetName
}

// GetVnetNameOk returns a tuple with the VnetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetVnetNameOk() (*string, bool) {
	if o == nil || IsNil(o.VnetName) {
		return nil, false
	}
	return o.VnetName, true
}

// HasVnetName returns a boolean if a field has been set.
func (o *AzureCloudProviderContainer) HasVnetName() bool {
	if o != nil && !IsNil(o.VnetName) {
		return true
	}

	return false
}

// SetVnetName gets a reference to the given string and assigns it to the VnetName field.
func (o *AzureCloudProviderContainer) SetVnetName(v string) {
	o.VnetName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureCloudProviderContainer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureCloudProviderContainer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureCloudProviderContainer) SetId(v string) {
	o.Id = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *AzureCloudProviderContainer) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *AzureCloudProviderContainer) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *AzureCloudProviderContainer) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetProvisioned returns the Provisioned field value if set, zero value otherwise.
func (o *AzureCloudProviderContainer) GetProvisioned() bool {
	if o == nil || IsNil(o.Provisioned) {
		var ret bool
		return ret
	}
	return *o.Provisioned
}

// GetProvisionedOk returns a tuple with the Provisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetProvisionedOk() (*bool, bool) {
	if o == nil || IsNil(o.Provisioned) {
		return nil, false
	}
	return o.Provisioned, true
}

// HasProvisioned returns a boolean if a field has been set.
func (o *AzureCloudProviderContainer) HasProvisioned() bool {
	if o != nil && !IsNil(o.Provisioned) {
		return true
	}

	return false
}

// SetProvisioned gets a reference to the given bool and assigns it to the Provisioned field.
func (o *AzureCloudProviderContainer) SetProvisioned(v bool) {
	o.Provisioned = &v
}

func (o AzureCloudProviderContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureCloudProviderContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["atlasCidrBlock"] = o.AtlasCidrBlock
	// skip: azureSubscriptionId is readOnly
	toSerialize["region"] = o.Region
	// skip: vnetName is readOnly
	// skip: id is readOnly
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	// skip: provisioned is readOnly
	return toSerialize, nil
}

type NullableAzureCloudProviderContainer struct {
	value *AzureCloudProviderContainer
	isSet bool
}

func (v NullableAzureCloudProviderContainer) Get() *AzureCloudProviderContainer {
	return v.value
}

func (v *NullableAzureCloudProviderContainer) Set(val *AzureCloudProviderContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCloudProviderContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCloudProviderContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCloudProviderContainer(val *AzureCloudProviderContainer) *NullableAzureCloudProviderContainer {
	return &NullableAzureCloudProviderContainer{value: val, isSet: true}
}

func (v NullableAzureCloudProviderContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCloudProviderContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


