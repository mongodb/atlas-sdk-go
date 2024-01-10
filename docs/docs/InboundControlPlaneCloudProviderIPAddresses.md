# InboundControlPlaneCloudProviderIPAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to **map[string][]string** | Control plane IP addresses in AWS. Each key identifies an Amazon Web Services (AWS) region. Each value identifies control plane IP addresses in the AWS region. | [optional] [readonly] 
**Azure** | Pointer to **map[string][]string** | Control plane IP addresses in Azure. Each key identifies an Azure region. Each value identifies control plane IP addresses in the Azure region. | [optional] [readonly] 
**Gcp** | Pointer to **map[string][]string** | Control plane IP addresses in GCP. Each key identifies a Google Cloud (GCP) region. Each value identifies control plane IP addresses in the GCP region. | [optional] [readonly] 

## Methods

### NewInboundControlPlaneCloudProviderIPAddresses

`func NewInboundControlPlaneCloudProviderIPAddresses() *InboundControlPlaneCloudProviderIPAddresses`

NewInboundControlPlaneCloudProviderIPAddresses instantiates a new InboundControlPlaneCloudProviderIPAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundControlPlaneCloudProviderIPAddressesWithDefaults

`func NewInboundControlPlaneCloudProviderIPAddressesWithDefaults() *InboundControlPlaneCloudProviderIPAddresses`

NewInboundControlPlaneCloudProviderIPAddressesWithDefaults instantiates a new InboundControlPlaneCloudProviderIPAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *InboundControlPlaneCloudProviderIPAddresses) GetAws() map[string][]string`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *InboundControlPlaneCloudProviderIPAddresses) GetAwsOk() (*map[string][]string, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *InboundControlPlaneCloudProviderIPAddresses) SetAws(v map[string][]string)`

SetAws sets Aws field to given value.

### HasAws

`func (o *InboundControlPlaneCloudProviderIPAddresses) HasAws() bool`

HasAws returns a boolean if a field has been set.
### GetAzure

`func (o *InboundControlPlaneCloudProviderIPAddresses) GetAzure() map[string][]string`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *InboundControlPlaneCloudProviderIPAddresses) GetAzureOk() (*map[string][]string, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *InboundControlPlaneCloudProviderIPAddresses) SetAzure(v map[string][]string)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *InboundControlPlaneCloudProviderIPAddresses) HasAzure() bool`

HasAzure returns a boolean if a field has been set.
### GetGcp

`func (o *InboundControlPlaneCloudProviderIPAddresses) GetGcp() map[string][]string`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *InboundControlPlaneCloudProviderIPAddresses) GetGcpOk() (*map[string][]string, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *InboundControlPlaneCloudProviderIPAddresses) SetGcp(v map[string][]string)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *InboundControlPlaneCloudProviderIPAddresses) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


