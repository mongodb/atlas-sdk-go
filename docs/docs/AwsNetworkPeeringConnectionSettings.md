# AwsNetworkPeeringConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccepterRegionName** | **string** | Amazon Web Services (AWS) region where the Virtual Peering Connection (VPC) that you peered with the MongoDB Cloud VPC resides. The resource returns &#x60;null&#x60; if your VPC and the MongoDB Cloud VPC reside in the same region. | 
**AwsAccountId** | **string** | Unique twelve-digit string that identifies the Amazon Web Services (AWS) account that owns the VPC that you peered with the MongoDB Cloud VPC. | 
**ConnectionId** | Pointer to **string** | Unique string that identifies the peering connection on AWS. | [optional] [readonly] 
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
**ErrorStateName** | Pointer to **string** | Type of error that can be returned when requesting an Amazon Web Services (AWS) peering connection. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering connection. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Cloud service provider that serves the requested network peering connection. | [optional] 
**RouteTableCidrBlock** | **string** | Internet Protocol (IP) addresses expressed in Classless Inter-Domain Routing (CIDR) notation of the VPC&#39;s subnet that you want to peer with the MongoDB Cloud VPC. | 
**StatusName** | Pointer to **string** | State of the network peering connection at the time you made the request. | [optional] [readonly] 
**VpcId** | **string** | Unique string that identifies the VPC on Amazon Web Services (AWS) that you want to peer with the MongoDB Cloud VPC. | 

## Methods

### NewAwsNetworkPeeringConnectionSettings

`func NewAwsNetworkPeeringConnectionSettings(accepterRegionName string, awsAccountId string, containerId string, routeTableCidrBlock string, vpcId string, ) *AwsNetworkPeeringConnectionSettings`

NewAwsNetworkPeeringConnectionSettings instantiates a new AwsNetworkPeeringConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsNetworkPeeringConnectionSettingsWithDefaults

`func NewAwsNetworkPeeringConnectionSettingsWithDefaults() *AwsNetworkPeeringConnectionSettings`

NewAwsNetworkPeeringConnectionSettingsWithDefaults instantiates a new AwsNetworkPeeringConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepterRegionName

`func (o *AwsNetworkPeeringConnectionSettings) GetAccepterRegionName() string`

GetAccepterRegionName returns the AccepterRegionName field if non-nil, zero value otherwise.

### GetAccepterRegionNameOk

`func (o *AwsNetworkPeeringConnectionSettings) GetAccepterRegionNameOk() (*string, bool)`

GetAccepterRegionNameOk returns a tuple with the AccepterRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterRegionName

`func (o *AwsNetworkPeeringConnectionSettings) SetAccepterRegionName(v string)`

SetAccepterRegionName sets AccepterRegionName field to given value.


### GetAwsAccountId

`func (o *AwsNetworkPeeringConnectionSettings) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *AwsNetworkPeeringConnectionSettings) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *AwsNetworkPeeringConnectionSettings) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.


### GetConnectionId

`func (o *AwsNetworkPeeringConnectionSettings) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AwsNetworkPeeringConnectionSettings) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AwsNetworkPeeringConnectionSettings) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AwsNetworkPeeringConnectionSettings) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetContainerId

`func (o *AwsNetworkPeeringConnectionSettings) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *AwsNetworkPeeringConnectionSettings) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *AwsNetworkPeeringConnectionSettings) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetErrorStateName

`func (o *AwsNetworkPeeringConnectionSettings) GetErrorStateName() string`

GetErrorStateName returns the ErrorStateName field if non-nil, zero value otherwise.

### GetErrorStateNameOk

`func (o *AwsNetworkPeeringConnectionSettings) GetErrorStateNameOk() (*string, bool)`

GetErrorStateNameOk returns a tuple with the ErrorStateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStateName

`func (o *AwsNetworkPeeringConnectionSettings) SetErrorStateName(v string)`

SetErrorStateName sets ErrorStateName field to given value.

### HasErrorStateName

`func (o *AwsNetworkPeeringConnectionSettings) HasErrorStateName() bool`

HasErrorStateName returns a boolean if a field has been set.

### GetId

`func (o *AwsNetworkPeeringConnectionSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsNetworkPeeringConnectionSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsNetworkPeeringConnectionSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsNetworkPeeringConnectionSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *AwsNetworkPeeringConnectionSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AwsNetworkPeeringConnectionSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AwsNetworkPeeringConnectionSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *AwsNetworkPeeringConnectionSettings) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetRouteTableCidrBlock

`func (o *AwsNetworkPeeringConnectionSettings) GetRouteTableCidrBlock() string`

GetRouteTableCidrBlock returns the RouteTableCidrBlock field if non-nil, zero value otherwise.

### GetRouteTableCidrBlockOk

`func (o *AwsNetworkPeeringConnectionSettings) GetRouteTableCidrBlockOk() (*string, bool)`

GetRouteTableCidrBlockOk returns a tuple with the RouteTableCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableCidrBlock

`func (o *AwsNetworkPeeringConnectionSettings) SetRouteTableCidrBlock(v string)`

SetRouteTableCidrBlock sets RouteTableCidrBlock field to given value.


### GetStatusName

`func (o *AwsNetworkPeeringConnectionSettings) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *AwsNetworkPeeringConnectionSettings) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *AwsNetworkPeeringConnectionSettings) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *AwsNetworkPeeringConnectionSettings) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetVpcId

`func (o *AwsNetworkPeeringConnectionSettings) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsNetworkPeeringConnectionSettings) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsNetworkPeeringConnectionSettings) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


