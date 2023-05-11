# AccessListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that found in this project&#39;s access list. | [optional] [readonly] 
**IpAddress** | **string** | IP address included in the API access list. | [readonly] 

## Methods

### NewAccessListItem

`func NewAccessListItem(ipAddress string, ) *AccessListItem`

NewAccessListItem instantiates a new AccessListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessListItemWithDefaults

`func NewAccessListItemWithDefaults() *AccessListItem`

NewAccessListItemWithDefaults instantiates a new AccessListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *AccessListItem) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AccessListItem) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AccessListItem) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AccessListItem) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetIpAddress

`func (o *AccessListItem) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AccessListItem) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AccessListItem) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


