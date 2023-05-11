# CloudProviderContainer

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
**GcpProjectId** | Pointer to **string** | Unique string that identifies the GCP project in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container. | [optional] [readonly] 
**NetworkName** | Pointer to **string** | Human-readable label that identifies the network in which MongoDB Cloud clusters in this network peering container exist. MongoDB Cloud returns **null** if no clusters exist in this network peering container. | [optional] [readonly] 
**Regions** | Pointer to **[]string** | List of GCP regions to which you want to deploy this MongoDB Cloud network peering container.  In this MongoDB Cloud project, you can deploy clusters only to the GCP regions in this list. To deploy MongoDB Cloud clusters to other GCP regions, create additional projects. | [optional] 
**RegionName** | **string** | Geographic area that Amazon Web Services (AWS) defines to which MongoDB Cloud deployed this network peering container. | 
**VpcId** | Pointer to **string** | Unique string that identifies the MongoDB Cloud VPC on AWS. | [optional] [readonly] 

## Methods

### NewCloudProviderContainer

`func NewCloudProviderContainer(atlasCidrBlock string, region string, regionName string, ) *CloudProviderContainer`

NewCloudProviderContainer instantiates a new CloudProviderContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderContainerWithDefaults

`func NewCloudProviderContainerWithDefaults() *CloudProviderContainer`

NewCloudProviderContainerWithDefaults instantiates a new CloudProviderContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasCidrBlock

`func (o *CloudProviderContainer) GetAtlasCidrBlock() string`

GetAtlasCidrBlock returns the AtlasCidrBlock field if non-nil, zero value otherwise.

### GetAtlasCidrBlockOk

`func (o *CloudProviderContainer) GetAtlasCidrBlockOk() (*string, bool)`

GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasCidrBlock

`func (o *CloudProviderContainer) SetAtlasCidrBlock(v string)`

SetAtlasCidrBlock sets AtlasCidrBlock field to given value.


### GetAzureSubscriptionId

`func (o *CloudProviderContainer) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *CloudProviderContainer) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *CloudProviderContainer) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *CloudProviderContainer) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### GetRegion

`func (o *CloudProviderContainer) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudProviderContainer) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudProviderContainer) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVnetName

`func (o *CloudProviderContainer) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *CloudProviderContainer) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *CloudProviderContainer) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.

### HasVnetName

`func (o *CloudProviderContainer) HasVnetName() bool`

HasVnetName returns a boolean if a field has been set.

### GetId

`func (o *CloudProviderContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProviderContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProviderContainer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProviderContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *CloudProviderContainer) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CloudProviderContainer) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CloudProviderContainer) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *CloudProviderContainer) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetProvisioned

`func (o *CloudProviderContainer) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *CloudProviderContainer) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *CloudProviderContainer) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *CloudProviderContainer) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *CloudProviderContainer) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *CloudProviderContainer) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *CloudProviderContainer) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *CloudProviderContainer) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.

### GetNetworkName

`func (o *CloudProviderContainer) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *CloudProviderContainer) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *CloudProviderContainer) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *CloudProviderContainer) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetRegions

`func (o *CloudProviderContainer) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CloudProviderContainer) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CloudProviderContainer) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CloudProviderContainer) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRegionName

`func (o *CloudProviderContainer) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *CloudProviderContainer) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *CloudProviderContainer) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetVpcId

`func (o *CloudProviderContainer) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CloudProviderContainer) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CloudProviderContainer) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *CloudProviderContainer) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


