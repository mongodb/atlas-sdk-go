# PrivateNetworkEndpointIdEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Human-readable string to associate with this private endpoint. | [optional] 
**CustomerEndpointDNSName** | Pointer to **string** | Human-readable label to identify customer&#39;s VPC endpoint DNS name. | [optional] 
**EndpointId** | **string** | Unique 22-character alphanumeric string that identifies the private endpoint. | 
**Provider** | Pointer to **string** | Human-readable label that identifies the cloud service provider. Atlas Data Lake supports Amazon Web Services only. | [optional] [default to "AWS"]
**Region** | Pointer to **string** | Human-readable label to identify the region of customer&#39;s VPC endpoint. | [optional] 
**Type** | Pointer to **string** | Human-readable label that identifies the resource type associated with this private endpoint. | [optional] [default to "DATA_LAKE"]

## Methods

### NewPrivateNetworkEndpointIdEntry

`func NewPrivateNetworkEndpointIdEntry(endpointId string, ) *PrivateNetworkEndpointIdEntry`

NewPrivateNetworkEndpointIdEntry instantiates a new PrivateNetworkEndpointIdEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkEndpointIdEntryWithDefaults

`func NewPrivateNetworkEndpointIdEntryWithDefaults() *PrivateNetworkEndpointIdEntry`

NewPrivateNetworkEndpointIdEntryWithDefaults instantiates a new PrivateNetworkEndpointIdEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PrivateNetworkEndpointIdEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PrivateNetworkEndpointIdEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PrivateNetworkEndpointIdEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PrivateNetworkEndpointIdEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.
### GetCustomerEndpointDNSName

`func (o *PrivateNetworkEndpointIdEntry) GetCustomerEndpointDNSName() string`

GetCustomerEndpointDNSName returns the CustomerEndpointDNSName field if non-nil, zero value otherwise.

### GetCustomerEndpointDNSNameOk

`func (o *PrivateNetworkEndpointIdEntry) GetCustomerEndpointDNSNameOk() (*string, bool)`

GetCustomerEndpointDNSNameOk returns a tuple with the CustomerEndpointDNSName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEndpointDNSName

`func (o *PrivateNetworkEndpointIdEntry) SetCustomerEndpointDNSName(v string)`

SetCustomerEndpointDNSName sets CustomerEndpointDNSName field to given value.

### HasCustomerEndpointDNSName

`func (o *PrivateNetworkEndpointIdEntry) HasCustomerEndpointDNSName() bool`

HasCustomerEndpointDNSName returns a boolean if a field has been set.
### GetEndpointId

`func (o *PrivateNetworkEndpointIdEntry) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *PrivateNetworkEndpointIdEntry) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *PrivateNetworkEndpointIdEntry) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### GetProvider

`func (o *PrivateNetworkEndpointIdEntry) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PrivateNetworkEndpointIdEntry) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PrivateNetworkEndpointIdEntry) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PrivateNetworkEndpointIdEntry) HasProvider() bool`

HasProvider returns a boolean if a field has been set.
### GetRegion

`func (o *PrivateNetworkEndpointIdEntry) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateNetworkEndpointIdEntry) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateNetworkEndpointIdEntry) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateNetworkEndpointIdEntry) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetType

`func (o *PrivateNetworkEndpointIdEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateNetworkEndpointIdEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateNetworkEndpointIdEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PrivateNetworkEndpointIdEntry) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


