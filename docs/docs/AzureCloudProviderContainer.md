# AzureCloudProviderContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasCidrBlock** | **string** | IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project&#39;s clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. AWS and Azure further limit the block to between the &#x60;/24&#x60; and  &#x60;/21&#x60; ranges.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Amazon Web Services (AWS) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of &#x60;/24&#x60; equals 27 three-node replica sets. | 
**AzureSubscriptionId** | Pointer to **string** | Unique string that identifies the Azure subscription in which the MongoDB Cloud VNet resides. | [optional] [readonly] 
**Region** | **string** | Azure region to which MongoDB Cloud deployed this network peering container. | 
**VnetName** | Pointer to **string** | Unique string that identifies the Azure VNet in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering container. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Cloud service provider that serves the requested network peering containers. | [optional] 
**Provisioned** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container. | [optional] [readonly] 

## Methods

### NewAzureCloudProviderContainer

`func NewAzureCloudProviderContainer(atlasCidrBlock string, region string, ) *AzureCloudProviderContainer`

NewAzureCloudProviderContainer instantiates a new AzureCloudProviderContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCloudProviderContainerWithDefaults

`func NewAzureCloudProviderContainerWithDefaults() *AzureCloudProviderContainer`

NewAzureCloudProviderContainerWithDefaults instantiates a new AzureCloudProviderContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasCidrBlock

`func (o *AzureCloudProviderContainer) GetAtlasCidrBlock() string`

GetAtlasCidrBlock returns the AtlasCidrBlock field if non-nil, zero value otherwise.

### GetAtlasCidrBlockOk

`func (o *AzureCloudProviderContainer) GetAtlasCidrBlockOk() (*string, bool)`

GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasCidrBlock

`func (o *AzureCloudProviderContainer) SetAtlasCidrBlock(v string)`

SetAtlasCidrBlock sets AtlasCidrBlock field to given value.


### GetAzureSubscriptionId

`func (o *AzureCloudProviderContainer) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *AzureCloudProviderContainer) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *AzureCloudProviderContainer) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *AzureCloudProviderContainer) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### GetRegion

`func (o *AzureCloudProviderContainer) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureCloudProviderContainer) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureCloudProviderContainer) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVnetName

`func (o *AzureCloudProviderContainer) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *AzureCloudProviderContainer) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *AzureCloudProviderContainer) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.

### HasVnetName

`func (o *AzureCloudProviderContainer) HasVnetName() bool`

HasVnetName returns a boolean if a field has been set.

### GetId

`func (o *AzureCloudProviderContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureCloudProviderContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureCloudProviderContainer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureCloudProviderContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *AzureCloudProviderContainer) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AzureCloudProviderContainer) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AzureCloudProviderContainer) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *AzureCloudProviderContainer) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetProvisioned

`func (o *AzureCloudProviderContainer) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *AzureCloudProviderContainer) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *AzureCloudProviderContainer) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *AzureCloudProviderContainer) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


