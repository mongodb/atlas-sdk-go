# UserAccessListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation in the access list for the API key. | [optional] 
**Count** | Pointer to **int** | Total number of requests that have originated from the Internet Protocol (IP) address given as the value of the &#x60;lastUsedAddress&#x60; parameter. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when someone added the network addresses to the specified API access list. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | Network address in the access list for the API key. | [optional] 
**LastUsed** | Pointer to **time.Time** | Date and time when MongoDB Cloud received the most recent request that originated from this Internet Protocol version 4 or version 6 address. The resource returns this parameter when at least one request has originated from this IP address. MongoDB Cloud updates this parameter each time a client accesses the permitted resource. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**LastUsedAddress** | Pointer to **string** | Network address that issued the most recent request to the API. This parameter requires the address to be expressed as one Internet Protocol version 4 or version 6 address. The resource returns this parameter after this IP address made at least one request. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewUserAccessListResponse

`func NewUserAccessListResponse() *UserAccessListResponse`

NewUserAccessListResponse instantiates a new UserAccessListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessListResponseWithDefaults

`func NewUserAccessListResponseWithDefaults() *UserAccessListResponse`

NewUserAccessListResponseWithDefaults instantiates a new UserAccessListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *UserAccessListResponse) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *UserAccessListResponse) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *UserAccessListResponse) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *UserAccessListResponse) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.
### GetCount

`func (o *UserAccessListResponse) GetCount() int`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UserAccessListResponse) GetCountOk() (*int, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UserAccessListResponse) SetCount(v int)`

SetCount sets Count field to given value.

### HasCount

`func (o *UserAccessListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.
### GetCreated

`func (o *UserAccessListResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserAccessListResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserAccessListResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserAccessListResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.
### GetIpAddress

`func (o *UserAccessListResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UserAccessListResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UserAccessListResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UserAccessListResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.
### GetLastUsed

`func (o *UserAccessListResponse) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *UserAccessListResponse) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *UserAccessListResponse) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *UserAccessListResponse) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.
### GetLastUsedAddress

`func (o *UserAccessListResponse) GetLastUsedAddress() string`

GetLastUsedAddress returns the LastUsedAddress field if non-nil, zero value otherwise.

### GetLastUsedAddressOk

`func (o *UserAccessListResponse) GetLastUsedAddressOk() (*string, bool)`

GetLastUsedAddressOk returns a tuple with the LastUsedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAddress

`func (o *UserAccessListResponse) SetLastUsedAddress(v string)`

SetLastUsedAddress sets LastUsedAddress field to given value.

### HasLastUsedAddress

`func (o *UserAccessListResponse) HasLastUsedAddress() bool`

HasLastUsedAddress returns a boolean if a field has been set.
### GetLinks

`func (o *UserAccessListResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserAccessListResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserAccessListResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserAccessListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


