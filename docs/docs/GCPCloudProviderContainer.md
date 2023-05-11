# GCPCloudProviderContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasCidrBlock** | **string** | IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project&#39;s clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. GCP further limits the block to a lower bound of the &#x60;/18&#x60; range.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Google Cloud (GCP) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of &#x60;/24&#x60; equals 27 three-node replica sets. | 
**GcpProjectId** | Pointer to **string** | Unique string that identifies the GCP project in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container. | [optional] [readonly] 
**NetworkName** | Pointer to **string** | Human-readable label that identifies the network in which MongoDB Cloud clusters in this network peering container exist. MongoDB Cloud returns **null** if no clusters exist in this network peering container. | [optional] [readonly] 
**Regions** | Pointer to **[]string** | List of GCP regions to which you want to deploy this MongoDB Cloud network peering container.  In this MongoDB Cloud project, you can deploy clusters only to the GCP regions in this list. To deploy MongoDB Cloud clusters to other GCP regions, create additional projects. | [optional] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering container. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Cloud service provider that serves the requested network peering containers. | [optional] 
**Provisioned** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container. | [optional] [readonly] 

## Methods

### NewGCPCloudProviderContainer

`func NewGCPCloudProviderContainer(atlasCidrBlock string, ) *GCPCloudProviderContainer`

NewGCPCloudProviderContainer instantiates a new GCPCloudProviderContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPCloudProviderContainerWithDefaults

`func NewGCPCloudProviderContainerWithDefaults() *GCPCloudProviderContainer`

NewGCPCloudProviderContainerWithDefaults instantiates a new GCPCloudProviderContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasCidrBlock

`func (o *GCPCloudProviderContainer) GetAtlasCidrBlock() string`

GetAtlasCidrBlock returns the AtlasCidrBlock field if non-nil, zero value otherwise.

### GetAtlasCidrBlockOk

`func (o *GCPCloudProviderContainer) GetAtlasCidrBlockOk() (*string, bool)`

GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasCidrBlock

`func (o *GCPCloudProviderContainer) SetAtlasCidrBlock(v string)`

SetAtlasCidrBlock sets AtlasCidrBlock field to given value.


### GetGcpProjectId

`func (o *GCPCloudProviderContainer) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *GCPCloudProviderContainer) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *GCPCloudProviderContainer) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *GCPCloudProviderContainer) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.

### GetNetworkName

`func (o *GCPCloudProviderContainer) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *GCPCloudProviderContainer) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *GCPCloudProviderContainer) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *GCPCloudProviderContainer) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetRegions

`func (o *GCPCloudProviderContainer) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GCPCloudProviderContainer) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GCPCloudProviderContainer) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GCPCloudProviderContainer) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetId

`func (o *GCPCloudProviderContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GCPCloudProviderContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GCPCloudProviderContainer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GCPCloudProviderContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *GCPCloudProviderContainer) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *GCPCloudProviderContainer) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *GCPCloudProviderContainer) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *GCPCloudProviderContainer) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetProvisioned

`func (o *GCPCloudProviderContainer) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *GCPCloudProviderContainer) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *GCPCloudProviderContainer) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *GCPCloudProviderContainer) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


