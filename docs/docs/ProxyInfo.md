# ProxyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | Pointer to **string** | Authentication key for the proxy. | [optional] 
**DnsName** | Pointer to **string** | DNS name to use to reach the proxy/s. | [optional] 

## Methods

### NewProxyInfo

`func NewProxyInfo() *ProxyInfo`

NewProxyInfo instantiates a new ProxyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyInfoWithDefaults

`func NewProxyInfoWithDefaults() *ProxyInfo`

NewProxyInfoWithDefaults instantiates a new ProxyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKey

`func (o *ProxyInfo) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *ProxyInfo) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *ProxyInfo) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *ProxyInfo) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.
### GetDnsName

`func (o *ProxyInfo) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ProxyInfo) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ProxyInfo) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *ProxyInfo) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


