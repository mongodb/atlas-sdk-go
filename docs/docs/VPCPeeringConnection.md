# VPCPeeringConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Internal VPC Peering Connection ID. | [optional] 
**AccepterAccountId** | Pointer to **string** | The account ID responsible for accepting the request. | [optional] 
**AccepterCidr** | Pointer to **string** | The CIDR block for the accepter VPC. | [optional] 
**AccepterVpcId** | Pointer to **string** | The VPC ID accepting the request. | [optional] 
**CloudStatus** | Pointer to **string** | The status in the cloud provider for this connection. | [optional] 
**ExpirationTime** | Pointer to **time.Time** | The time when the VPC Peering Connection request expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] 
**GroupId** | Pointer to **string** | The internal project ID. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**LocalStatus** | Pointer to **string** | The local status of the VPC Peering Connection. | [optional] 
**Name** | Pointer to **string** | Unique VPC Peering Connection name. | [optional] 
**RequesterAccountId** | Pointer to **string** | The account ID requesting the VPC Peering connection. | [optional] 
**RequesterCidr** | Pointer to **string** | The CIDR block for the requesting VPC. | [optional] 
**RequesterVpcId** | Pointer to **string** | The VPC ID requesting the VPC Peering connection. | [optional] 
**StatusMessage** | Pointer to **string** | A status message. | [optional] 

## Methods

### NewVPCPeeringConnection

`func NewVPCPeeringConnection() *VPCPeeringConnection`

NewVPCPeeringConnection instantiates a new VPCPeeringConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPCPeeringConnectionWithDefaults

`func NewVPCPeeringConnectionWithDefaults() *VPCPeeringConnection`

NewVPCPeeringConnectionWithDefaults instantiates a new VPCPeeringConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPCPeeringConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPCPeeringConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPCPeeringConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPCPeeringConnection) HasId() bool`

HasId returns a boolean if a field has been set.
### GetAccepterAccountId

`func (o *VPCPeeringConnection) GetAccepterAccountId() string`

GetAccepterAccountId returns the AccepterAccountId field if non-nil, zero value otherwise.

### GetAccepterAccountIdOk

`func (o *VPCPeeringConnection) GetAccepterAccountIdOk() (*string, bool)`

GetAccepterAccountIdOk returns a tuple with the AccepterAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterAccountId

`func (o *VPCPeeringConnection) SetAccepterAccountId(v string)`

SetAccepterAccountId sets AccepterAccountId field to given value.

### HasAccepterAccountId

`func (o *VPCPeeringConnection) HasAccepterAccountId() bool`

HasAccepterAccountId returns a boolean if a field has been set.
### GetAccepterCidr

`func (o *VPCPeeringConnection) GetAccepterCidr() string`

GetAccepterCidr returns the AccepterCidr field if non-nil, zero value otherwise.

### GetAccepterCidrOk

`func (o *VPCPeeringConnection) GetAccepterCidrOk() (*string, bool)`

GetAccepterCidrOk returns a tuple with the AccepterCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterCidr

`func (o *VPCPeeringConnection) SetAccepterCidr(v string)`

SetAccepterCidr sets AccepterCidr field to given value.

### HasAccepterCidr

`func (o *VPCPeeringConnection) HasAccepterCidr() bool`

HasAccepterCidr returns a boolean if a field has been set.
### GetAccepterVpcId

`func (o *VPCPeeringConnection) GetAccepterVpcId() string`

GetAccepterVpcId returns the AccepterVpcId field if non-nil, zero value otherwise.

### GetAccepterVpcIdOk

`func (o *VPCPeeringConnection) GetAccepterVpcIdOk() (*string, bool)`

GetAccepterVpcIdOk returns a tuple with the AccepterVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterVpcId

`func (o *VPCPeeringConnection) SetAccepterVpcId(v string)`

SetAccepterVpcId sets AccepterVpcId field to given value.

### HasAccepterVpcId

`func (o *VPCPeeringConnection) HasAccepterVpcId() bool`

HasAccepterVpcId returns a boolean if a field has been set.
### GetCloudStatus

`func (o *VPCPeeringConnection) GetCloudStatus() string`

GetCloudStatus returns the CloudStatus field if non-nil, zero value otherwise.

### GetCloudStatusOk

`func (o *VPCPeeringConnection) GetCloudStatusOk() (*string, bool)`

GetCloudStatusOk returns a tuple with the CloudStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStatus

`func (o *VPCPeeringConnection) SetCloudStatus(v string)`

SetCloudStatus sets CloudStatus field to given value.

### HasCloudStatus

`func (o *VPCPeeringConnection) HasCloudStatus() bool`

HasCloudStatus returns a boolean if a field has been set.
### GetExpirationTime

`func (o *VPCPeeringConnection) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *VPCPeeringConnection) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *VPCPeeringConnection) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *VPCPeeringConnection) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.
### GetGroupId

`func (o *VPCPeeringConnection) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VPCPeeringConnection) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VPCPeeringConnection) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VPCPeeringConnection) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetLinks

`func (o *VPCPeeringConnection) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VPCPeeringConnection) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VPCPeeringConnection) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VPCPeeringConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetLocalStatus

`func (o *VPCPeeringConnection) GetLocalStatus() string`

GetLocalStatus returns the LocalStatus field if non-nil, zero value otherwise.

### GetLocalStatusOk

`func (o *VPCPeeringConnection) GetLocalStatusOk() (*string, bool)`

GetLocalStatusOk returns a tuple with the LocalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatus

`func (o *VPCPeeringConnection) SetLocalStatus(v string)`

SetLocalStatus sets LocalStatus field to given value.

### HasLocalStatus

`func (o *VPCPeeringConnection) HasLocalStatus() bool`

HasLocalStatus returns a boolean if a field has been set.
### GetName

`func (o *VPCPeeringConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPCPeeringConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPCPeeringConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VPCPeeringConnection) HasName() bool`

HasName returns a boolean if a field has been set.
### GetRequesterAccountId

`func (o *VPCPeeringConnection) GetRequesterAccountId() string`

GetRequesterAccountId returns the RequesterAccountId field if non-nil, zero value otherwise.

### GetRequesterAccountIdOk

`func (o *VPCPeeringConnection) GetRequesterAccountIdOk() (*string, bool)`

GetRequesterAccountIdOk returns a tuple with the RequesterAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterAccountId

`func (o *VPCPeeringConnection) SetRequesterAccountId(v string)`

SetRequesterAccountId sets RequesterAccountId field to given value.

### HasRequesterAccountId

`func (o *VPCPeeringConnection) HasRequesterAccountId() bool`

HasRequesterAccountId returns a boolean if a field has been set.
### GetRequesterCidr

`func (o *VPCPeeringConnection) GetRequesterCidr() string`

GetRequesterCidr returns the RequesterCidr field if non-nil, zero value otherwise.

### GetRequesterCidrOk

`func (o *VPCPeeringConnection) GetRequesterCidrOk() (*string, bool)`

GetRequesterCidrOk returns a tuple with the RequesterCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterCidr

`func (o *VPCPeeringConnection) SetRequesterCidr(v string)`

SetRequesterCidr sets RequesterCidr field to given value.

### HasRequesterCidr

`func (o *VPCPeeringConnection) HasRequesterCidr() bool`

HasRequesterCidr returns a boolean if a field has been set.
### GetRequesterVpcId

`func (o *VPCPeeringConnection) GetRequesterVpcId() string`

GetRequesterVpcId returns the RequesterVpcId field if non-nil, zero value otherwise.

### GetRequesterVpcIdOk

`func (o *VPCPeeringConnection) GetRequesterVpcIdOk() (*string, bool)`

GetRequesterVpcIdOk returns a tuple with the RequesterVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterVpcId

`func (o *VPCPeeringConnection) SetRequesterVpcId(v string)`

SetRequesterVpcId sets RequesterVpcId field to given value.

### HasRequesterVpcId

`func (o *VPCPeeringConnection) HasRequesterVpcId() bool`

HasRequesterVpcId returns a boolean if a field has been set.
### GetStatusMessage

`func (o *VPCPeeringConnection) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *VPCPeeringConnection) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *VPCPeeringConnection) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *VPCPeeringConnection) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


