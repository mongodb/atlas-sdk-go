# BaseNetworkPeeringConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering connection. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Cloud service provider that serves the requested network peering connection. | [optional] 
**AccepterRegionName** | Pointer to **string** | Amazon Web Services (AWS) region where the Virtual Peering Connection (VPC) that you peered with the MongoDB Cloud VPC resides. The resource returns &#x60;null&#x60; if your VPC and the MongoDB Cloud VPC reside in the same region. | [optional] 
**AwsAccountId** | Pointer to **string** | Unique twelve-digit string that identifies the Amazon Web Services (AWS) account that owns the VPC that you peered with the MongoDB Cloud VPC. | [optional] 
**ConnectionId** | Pointer to **string** | Unique string that identifies the peering connection on AWS. | [optional] [readonly] 
**ErrorStateName** | Pointer to **string** | Type of error that can be returned when requesting an Amazon Web Services (AWS) peering connection. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**RouteTableCidrBlock** | Pointer to **string** | Internet Protocol (IP) addresses expressed in Classless Inter-Domain Routing (CIDR) notation of the VPC&#39;s subnet that you want to peer with the MongoDB Cloud VPC. | [optional] 
**StatusName** | Pointer to **string** | State of the network peering connection at the time you made the request. | [optional] [readonly] 
**VpcId** | Pointer to **string** | Unique string that identifies the VPC on Amazon Web Services (AWS) that you want to peer with the MongoDB Cloud VPC. | [optional] 
**AzureDirectoryId** | Pointer to **string** | Unique string that identifies the Azure AD directory in which the VNet peered with the MongoDB Cloud VNet resides. | [optional] 
**AzureSubscriptionId** | Pointer to **string** | Unique string that identifies the Azure subscription in which the VNet you peered with the MongoDB Cloud VNet resides. | [optional] 
**ErrorState** | Pointer to **string** | Error message returned when a requested Azure network peering resource returns &#x60;\&quot;status\&quot; : \&quot;FAILED\&quot;&#x60;. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**ResourceGroupName** | Pointer to **string** | Human-readable label that identifies the resource group in which the VNet to peer with the MongoDB Cloud VNet resides. | [optional] 
**Status** | Pointer to **string** | State of the network peering connection at the time you made the request. | [optional] [readonly] 
**VnetName** | Pointer to **string** | Human-readable label that identifies the VNet that you want to peer with the MongoDB Cloud VNet. | [optional] 
**ErrorMessage** | Pointer to **string** | Details of the error returned when requesting a GCP network peering resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**GcpProjectId** | Pointer to **string** | Human-readable label that identifies the GCP project that contains the network that you want to peer with the MongoDB Cloud VPC. | [optional] 
**NetworkName** | Pointer to **string** | Human-readable label that identifies the network to peer with the MongoDB Cloud VPC. | [optional] 

## Methods

### NewBaseNetworkPeeringConnectionSettings

`func NewBaseNetworkPeeringConnectionSettings(containerId string, ) *BaseNetworkPeeringConnectionSettings`

NewBaseNetworkPeeringConnectionSettings instantiates a new BaseNetworkPeeringConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNetworkPeeringConnectionSettingsWithDefaults

`func NewBaseNetworkPeeringConnectionSettingsWithDefaults() *BaseNetworkPeeringConnectionSettings`

NewBaseNetworkPeeringConnectionSettingsWithDefaults instantiates a new BaseNetworkPeeringConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *BaseNetworkPeeringConnectionSettings) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *BaseNetworkPeeringConnectionSettings) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *BaseNetworkPeeringConnectionSettings) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### GetId

`func (o *BaseNetworkPeeringConnectionSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseNetworkPeeringConnectionSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseNetworkPeeringConnectionSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseNetworkPeeringConnectionSettings) HasId() bool`

HasId returns a boolean if a field has been set.
### GetProviderName

`func (o *BaseNetworkPeeringConnectionSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *BaseNetworkPeeringConnectionSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *BaseNetworkPeeringConnectionSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *BaseNetworkPeeringConnectionSettings) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.
### GetAccepterRegionName

`func (o *BaseNetworkPeeringConnectionSettings) GetAccepterRegionName() string`

GetAccepterRegionName returns the AccepterRegionName field if non-nil, zero value otherwise.

### GetAccepterRegionNameOk

`func (o *BaseNetworkPeeringConnectionSettings) GetAccepterRegionNameOk() (*string, bool)`

GetAccepterRegionNameOk returns a tuple with the AccepterRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterRegionName

`func (o *BaseNetworkPeeringConnectionSettings) SetAccepterRegionName(v string)`

SetAccepterRegionName sets AccepterRegionName field to given value.

### HasAccepterRegionName

`func (o *BaseNetworkPeeringConnectionSettings) HasAccepterRegionName() bool`

HasAccepterRegionName returns a boolean if a field has been set.
### GetAwsAccountId

`func (o *BaseNetworkPeeringConnectionSettings) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *BaseNetworkPeeringConnectionSettings) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *BaseNetworkPeeringConnectionSettings) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.

### HasAwsAccountId

`func (o *BaseNetworkPeeringConnectionSettings) HasAwsAccountId() bool`

HasAwsAccountId returns a boolean if a field has been set.
### GetConnectionId

`func (o *BaseNetworkPeeringConnectionSettings) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseNetworkPeeringConnectionSettings) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseNetworkPeeringConnectionSettings) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *BaseNetworkPeeringConnectionSettings) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.
### GetErrorStateName

`func (o *BaseNetworkPeeringConnectionSettings) GetErrorStateName() string`

GetErrorStateName returns the ErrorStateName field if non-nil, zero value otherwise.

### GetErrorStateNameOk

`func (o *BaseNetworkPeeringConnectionSettings) GetErrorStateNameOk() (*string, bool)`

GetErrorStateNameOk returns a tuple with the ErrorStateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStateName

`func (o *BaseNetworkPeeringConnectionSettings) SetErrorStateName(v string)`

SetErrorStateName sets ErrorStateName field to given value.

### HasErrorStateName

`func (o *BaseNetworkPeeringConnectionSettings) HasErrorStateName() bool`

HasErrorStateName returns a boolean if a field has been set.
### GetRouteTableCidrBlock

`func (o *BaseNetworkPeeringConnectionSettings) GetRouteTableCidrBlock() string`

GetRouteTableCidrBlock returns the RouteTableCidrBlock field if non-nil, zero value otherwise.

### GetRouteTableCidrBlockOk

`func (o *BaseNetworkPeeringConnectionSettings) GetRouteTableCidrBlockOk() (*string, bool)`

GetRouteTableCidrBlockOk returns a tuple with the RouteTableCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableCidrBlock

`func (o *BaseNetworkPeeringConnectionSettings) SetRouteTableCidrBlock(v string)`

SetRouteTableCidrBlock sets RouteTableCidrBlock field to given value.

### HasRouteTableCidrBlock

`func (o *BaseNetworkPeeringConnectionSettings) HasRouteTableCidrBlock() bool`

HasRouteTableCidrBlock returns a boolean if a field has been set.
### GetStatusName

`func (o *BaseNetworkPeeringConnectionSettings) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *BaseNetworkPeeringConnectionSettings) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *BaseNetworkPeeringConnectionSettings) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *BaseNetworkPeeringConnectionSettings) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.
### GetVpcId

`func (o *BaseNetworkPeeringConnectionSettings) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BaseNetworkPeeringConnectionSettings) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BaseNetworkPeeringConnectionSettings) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BaseNetworkPeeringConnectionSettings) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.
### GetAzureDirectoryId

`func (o *BaseNetworkPeeringConnectionSettings) GetAzureDirectoryId() string`

GetAzureDirectoryId returns the AzureDirectoryId field if non-nil, zero value otherwise.

### GetAzureDirectoryIdOk

`func (o *BaseNetworkPeeringConnectionSettings) GetAzureDirectoryIdOk() (*string, bool)`

GetAzureDirectoryIdOk returns a tuple with the AzureDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDirectoryId

`func (o *BaseNetworkPeeringConnectionSettings) SetAzureDirectoryId(v string)`

SetAzureDirectoryId sets AzureDirectoryId field to given value.

### HasAzureDirectoryId

`func (o *BaseNetworkPeeringConnectionSettings) HasAzureDirectoryId() bool`

HasAzureDirectoryId returns a boolean if a field has been set.
### GetAzureSubscriptionId

`func (o *BaseNetworkPeeringConnectionSettings) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *BaseNetworkPeeringConnectionSettings) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *BaseNetworkPeeringConnectionSettings) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *BaseNetworkPeeringConnectionSettings) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.
### GetErrorState

`func (o *BaseNetworkPeeringConnectionSettings) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *BaseNetworkPeeringConnectionSettings) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *BaseNetworkPeeringConnectionSettings) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *BaseNetworkPeeringConnectionSettings) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.
### GetResourceGroupName

`func (o *BaseNetworkPeeringConnectionSettings) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *BaseNetworkPeeringConnectionSettings) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *BaseNetworkPeeringConnectionSettings) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *BaseNetworkPeeringConnectionSettings) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.
### GetStatus

`func (o *BaseNetworkPeeringConnectionSettings) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseNetworkPeeringConnectionSettings) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseNetworkPeeringConnectionSettings) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseNetworkPeeringConnectionSettings) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetVnetName

`func (o *BaseNetworkPeeringConnectionSettings) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *BaseNetworkPeeringConnectionSettings) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *BaseNetworkPeeringConnectionSettings) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.

### HasVnetName

`func (o *BaseNetworkPeeringConnectionSettings) HasVnetName() bool`

HasVnetName returns a boolean if a field has been set.
### GetErrorMessage

`func (o *BaseNetworkPeeringConnectionSettings) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BaseNetworkPeeringConnectionSettings) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BaseNetworkPeeringConnectionSettings) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BaseNetworkPeeringConnectionSettings) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetGcpProjectId

`func (o *BaseNetworkPeeringConnectionSettings) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *BaseNetworkPeeringConnectionSettings) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *BaseNetworkPeeringConnectionSettings) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *BaseNetworkPeeringConnectionSettings) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.
### GetNetworkName

`func (o *BaseNetworkPeeringConnectionSettings) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *BaseNetworkPeeringConnectionSettings) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *BaseNetworkPeeringConnectionSettings) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *BaseNetworkPeeringConnectionSettings) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


