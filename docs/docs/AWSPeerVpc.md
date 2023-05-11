# AWSPeerVpc

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

### NewAWSPeerVpc

`func NewAWSPeerVpc(accepterRegionName string, awsAccountId string, containerId string, routeTableCidrBlock string, vpcId string, ) *AWSPeerVpc`

NewAWSPeerVpc instantiates a new AWSPeerVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSPeerVpcWithDefaults

`func NewAWSPeerVpcWithDefaults() *AWSPeerVpc`

NewAWSPeerVpcWithDefaults instantiates a new AWSPeerVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepterRegionName

`func (o *AWSPeerVpc) GetAccepterRegionName() string`

GetAccepterRegionName returns the AccepterRegionName field if non-nil, zero value otherwise.

### GetAccepterRegionNameOk

`func (o *AWSPeerVpc) GetAccepterRegionNameOk() (*string, bool)`

GetAccepterRegionNameOk returns a tuple with the AccepterRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterRegionName

`func (o *AWSPeerVpc) SetAccepterRegionName(v string)`

SetAccepterRegionName sets AccepterRegionName field to given value.


### GetAwsAccountId

`func (o *AWSPeerVpc) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *AWSPeerVpc) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *AWSPeerVpc) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.


### GetConnectionId

`func (o *AWSPeerVpc) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AWSPeerVpc) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AWSPeerVpc) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AWSPeerVpc) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetContainerId

`func (o *AWSPeerVpc) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *AWSPeerVpc) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *AWSPeerVpc) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetErrorStateName

`func (o *AWSPeerVpc) GetErrorStateName() string`

GetErrorStateName returns the ErrorStateName field if non-nil, zero value otherwise.

### GetErrorStateNameOk

`func (o *AWSPeerVpc) GetErrorStateNameOk() (*string, bool)`

GetErrorStateNameOk returns a tuple with the ErrorStateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStateName

`func (o *AWSPeerVpc) SetErrorStateName(v string)`

SetErrorStateName sets ErrorStateName field to given value.

### HasErrorStateName

`func (o *AWSPeerVpc) HasErrorStateName() bool`

HasErrorStateName returns a boolean if a field has been set.

### GetId

`func (o *AWSPeerVpc) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSPeerVpc) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSPeerVpc) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AWSPeerVpc) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *AWSPeerVpc) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AWSPeerVpc) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AWSPeerVpc) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *AWSPeerVpc) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetRouteTableCidrBlock

`func (o *AWSPeerVpc) GetRouteTableCidrBlock() string`

GetRouteTableCidrBlock returns the RouteTableCidrBlock field if non-nil, zero value otherwise.

### GetRouteTableCidrBlockOk

`func (o *AWSPeerVpc) GetRouteTableCidrBlockOk() (*string, bool)`

GetRouteTableCidrBlockOk returns a tuple with the RouteTableCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableCidrBlock

`func (o *AWSPeerVpc) SetRouteTableCidrBlock(v string)`

SetRouteTableCidrBlock sets RouteTableCidrBlock field to given value.


### GetStatusName

`func (o *AWSPeerVpc) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *AWSPeerVpc) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *AWSPeerVpc) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *AWSPeerVpc) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetVpcId

`func (o *AWSPeerVpc) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AWSPeerVpc) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AWSPeerVpc) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


