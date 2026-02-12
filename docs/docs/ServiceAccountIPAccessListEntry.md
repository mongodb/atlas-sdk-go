# ServiceAccountIPAccessListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | Range of network addresses in the access list for the Service Account. This parameter requires the range to be expressed in Classless Inter-Domain Routing (CIDR) notation of Internet Protocol version 4 or version 6 addresses. You can set a value for this parameter or &#x60;ipAddress&#x60;, but not for both in the same request. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date MongoDB Cloud added the entry was added to the Access List. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | Network address in the access list for the Service Account. This parameter requires the address to be expressed as one Internet Protocol version 4 or version 6 address. You can set a value for this parameter or &#x60;cidrBlock&#x60;, but not for both in the same request. | [optional] 
**LastUsedAddress** | Pointer to **string** | Network address that issued the most recent request to the API. This parameter requires the address to be expressed as one Internet Protocol version 4 or version 6 address. The resource returns this parameter after this IP address makes at least one request. | [optional] [readonly] 
**LastUsedAt** | Pointer to **time.Time** | Date when MongoDB Cloud received the most recent request that originated from this Internet Protocol version 4 or version 6 address. The resource returns this parameter when at least one request originates from this IP address. MongoDB Cloud updates this parameter each time a client accesses the permitted resource, with a delay of up to 5 minutes. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**RequestCount** | Pointer to **int** | The number of requests that has originated from this network address. | [optional] [readonly] 

## Methods

### NewServiceAccountIPAccessListEntry

`func NewServiceAccountIPAccessListEntry() *ServiceAccountIPAccessListEntry`

NewServiceAccountIPAccessListEntry instantiates a new ServiceAccountIPAccessListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountIPAccessListEntryWithDefaults

`func NewServiceAccountIPAccessListEntryWithDefaults() *ServiceAccountIPAccessListEntry`

NewServiceAccountIPAccessListEntryWithDefaults instantiates a new ServiceAccountIPAccessListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *ServiceAccountIPAccessListEntry) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *ServiceAccountIPAccessListEntry) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *ServiceAccountIPAccessListEntry) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *ServiceAccountIPAccessListEntry) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.
### GetCreatedAt

`func (o *ServiceAccountIPAccessListEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceAccountIPAccessListEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceAccountIPAccessListEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceAccountIPAccessListEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetIpAddress

`func (o *ServiceAccountIPAccessListEntry) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ServiceAccountIPAccessListEntry) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ServiceAccountIPAccessListEntry) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ServiceAccountIPAccessListEntry) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.
### GetLastUsedAddress

`func (o *ServiceAccountIPAccessListEntry) GetLastUsedAddress() string`

GetLastUsedAddress returns the LastUsedAddress field if non-nil, zero value otherwise.

### GetLastUsedAddressOk

`func (o *ServiceAccountIPAccessListEntry) GetLastUsedAddressOk() (*string, bool)`

GetLastUsedAddressOk returns a tuple with the LastUsedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAddress

`func (o *ServiceAccountIPAccessListEntry) SetLastUsedAddress(v string)`

SetLastUsedAddress sets LastUsedAddress field to given value.

### HasLastUsedAddress

`func (o *ServiceAccountIPAccessListEntry) HasLastUsedAddress() bool`

HasLastUsedAddress returns a boolean if a field has been set.
### GetLastUsedAt

`func (o *ServiceAccountIPAccessListEntry) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ServiceAccountIPAccessListEntry) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ServiceAccountIPAccessListEntry) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ServiceAccountIPAccessListEntry) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.
### GetRequestCount

`func (o *ServiceAccountIPAccessListEntry) GetRequestCount() int`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *ServiceAccountIPAccessListEntry) GetRequestCountOk() (*int, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *ServiceAccountIPAccessListEntry) SetRequestCount(v int)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *ServiceAccountIPAccessListEntry) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


