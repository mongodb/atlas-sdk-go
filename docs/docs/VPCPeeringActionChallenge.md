# VPCPeeringActionChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**RequesterAccountId** | Pointer to **string** | The AWS requester account ID. | [optional] 
**RequesterVpcId** | Pointer to **string** | The AWS requester VPC ID. | [optional] 

## Methods

### NewVPCPeeringActionChallenge

`func NewVPCPeeringActionChallenge() *VPCPeeringActionChallenge`

NewVPCPeeringActionChallenge instantiates a new VPCPeeringActionChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPCPeeringActionChallengeWithDefaults

`func NewVPCPeeringActionChallengeWithDefaults() *VPCPeeringActionChallenge`

NewVPCPeeringActionChallengeWithDefaults instantiates a new VPCPeeringActionChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *VPCPeeringActionChallenge) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VPCPeeringActionChallenge) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VPCPeeringActionChallenge) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VPCPeeringActionChallenge) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRequesterAccountId

`func (o *VPCPeeringActionChallenge) GetRequesterAccountId() string`

GetRequesterAccountId returns the RequesterAccountId field if non-nil, zero value otherwise.

### GetRequesterAccountIdOk

`func (o *VPCPeeringActionChallenge) GetRequesterAccountIdOk() (*string, bool)`

GetRequesterAccountIdOk returns a tuple with the RequesterAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterAccountId

`func (o *VPCPeeringActionChallenge) SetRequesterAccountId(v string)`

SetRequesterAccountId sets RequesterAccountId field to given value.

### HasRequesterAccountId

`func (o *VPCPeeringActionChallenge) HasRequesterAccountId() bool`

HasRequesterAccountId returns a boolean if a field has been set.
### GetRequesterVpcId

`func (o *VPCPeeringActionChallenge) GetRequesterVpcId() string`

GetRequesterVpcId returns the RequesterVpcId field if non-nil, zero value otherwise.

### GetRequesterVpcIdOk

`func (o *VPCPeeringActionChallenge) GetRequesterVpcIdOk() (*string, bool)`

GetRequesterVpcIdOk returns a tuple with the RequesterVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterVpcId

`func (o *VPCPeeringActionChallenge) SetRequesterVpcId(v string)`

SetRequesterVpcId sets RequesterVpcId field to given value.

### HasRequesterVpcId

`func (o *VPCPeeringActionChallenge) HasRequesterVpcId() bool`

HasRequesterVpcId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


