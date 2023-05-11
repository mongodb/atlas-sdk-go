# AWSCloudProviderContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasCidrBlock** | Pointer to **string** | IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project&#39;s clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. AWS and Azure further limit the block to between the &#x60;/24&#x60; and  &#x60;/21&#x60; ranges.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Amazon Web Services (AWS) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of &#x60;/24&#x60; equals 27 three-node replica sets. | [optional] 
**RegionName** | **string** | Geographic area that Amazon Web Services (AWS) defines to which MongoDB Cloud deployed this network peering container. | 
**VpcId** | Pointer to **string** | Unique string that identifies the MongoDB Cloud VPC on AWS. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering container. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Cloud service provider that serves the requested network peering containers. | [optional] 
**Provisioned** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container. | [optional] [readonly] 

## Methods

### NewAWSCloudProviderContainer

`func NewAWSCloudProviderContainer(regionName string, ) *AWSCloudProviderContainer`

NewAWSCloudProviderContainer instantiates a new AWSCloudProviderContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCloudProviderContainerWithDefaults

`func NewAWSCloudProviderContainerWithDefaults() *AWSCloudProviderContainer`

NewAWSCloudProviderContainerWithDefaults instantiates a new AWSCloudProviderContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasCidrBlock

`func (o *AWSCloudProviderContainer) GetAtlasCidrBlock() string`

GetAtlasCidrBlock returns the AtlasCidrBlock field if non-nil, zero value otherwise.

### GetAtlasCidrBlockOk

`func (o *AWSCloudProviderContainer) GetAtlasCidrBlockOk() (*string, bool)`

GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasCidrBlock

`func (o *AWSCloudProviderContainer) SetAtlasCidrBlock(v string)`

SetAtlasCidrBlock sets AtlasCidrBlock field to given value.

### HasAtlasCidrBlock

`func (o *AWSCloudProviderContainer) HasAtlasCidrBlock() bool`

HasAtlasCidrBlock returns a boolean if a field has been set.

### GetRegionName

`func (o *AWSCloudProviderContainer) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AWSCloudProviderContainer) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AWSCloudProviderContainer) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetVpcId

`func (o *AWSCloudProviderContainer) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AWSCloudProviderContainer) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AWSCloudProviderContainer) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AWSCloudProviderContainer) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetId

`func (o *AWSCloudProviderContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSCloudProviderContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSCloudProviderContainer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AWSCloudProviderContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *AWSCloudProviderContainer) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AWSCloudProviderContainer) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AWSCloudProviderContainer) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *AWSCloudProviderContainer) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetProvisioned

`func (o *AWSCloudProviderContainer) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *AWSCloudProviderContainer) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *AWSCloudProviderContainer) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *AWSCloudProviderContainer) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


