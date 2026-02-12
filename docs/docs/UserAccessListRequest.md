# UserAccessListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | Range of network addresses that you want to add to the access list for the API key. This parameter requires the range to be expressed in classless inter-domain routing (CIDR) notation of Internet Protocol version 4 or version 6 addresses. You can set a value for this parameter or &#x60;ipAddress&#x60; but not both in the same request. | [optional] 
**IpAddress** | Pointer to **string** | Network address that you want to add to the access list for the API key. This parameter requires the address to be expressed as one Internet Protocol version 4 or version 6 address. You can set a value for this parameter or &#x60;cidrBlock&#x60; but not both in the same request. | [optional] 

## Methods

### NewUserAccessListRequest

`func NewUserAccessListRequest() *UserAccessListRequest`

NewUserAccessListRequest instantiates a new UserAccessListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessListRequestWithDefaults

`func NewUserAccessListRequestWithDefaults() *UserAccessListRequest`

NewUserAccessListRequestWithDefaults instantiates a new UserAccessListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *UserAccessListRequest) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *UserAccessListRequest) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *UserAccessListRequest) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *UserAccessListRequest) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.
### GetIpAddress

`func (o *UserAccessListRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UserAccessListRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UserAccessListRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UserAccessListRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


