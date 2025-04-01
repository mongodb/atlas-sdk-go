# ClusterIPAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster. | [optional] [readonly] 
**FutureInbound** | Pointer to **[]string** | List of future inbound IP addresses associated with the cluster. If your network allows outbound HTTP requests only to specific IP addresses, you must allow access to the following IP addresses so that your application can connect to your Atlas cluster. | [optional] [readonly] 
**FutureOutbound** | Pointer to **[]string** | List of future outbound IP addresses associated with the cluster. If your network allows inbound HTTP requests only from specific IP addresses, you must allow access from the following IP addresses so that your Atlas cluster can communicate with your webhooks and KMS. | [optional] [readonly] 
**Inbound** | Pointer to **[]string** | List of inbound IP addresses associated with the cluster. If your network allows outbound HTTP requests only to specific IP addresses, you must allow access to the following IP addresses so that your application can connect to your Atlas cluster. | [optional] [readonly] 
**Outbound** | Pointer to **[]string** | List of outbound IP addresses associated with the cluster. If your network allows inbound HTTP requests only from specific IP addresses, you must allow access from the following IP addresses so that your Atlas cluster can communicate with your webhooks and KMS. | [optional] [readonly] 

## Methods

### NewClusterIPAddresses

`func NewClusterIPAddresses() *ClusterIPAddresses`

NewClusterIPAddresses instantiates a new ClusterIPAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterIPAddressesWithDefaults

`func NewClusterIPAddressesWithDefaults() *ClusterIPAddresses`

NewClusterIPAddressesWithDefaults instantiates a new ClusterIPAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ClusterIPAddresses) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterIPAddresses) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterIPAddresses) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ClusterIPAddresses) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetFutureInbound

`func (o *ClusterIPAddresses) GetFutureInbound() []string`

GetFutureInbound returns the FutureInbound field if non-nil, zero value otherwise.

### GetFutureInboundOk

`func (o *ClusterIPAddresses) GetFutureInboundOk() (*[]string, bool)`

GetFutureInboundOk returns a tuple with the FutureInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureInbound

`func (o *ClusterIPAddresses) SetFutureInbound(v []string)`

SetFutureInbound sets FutureInbound field to given value.

### HasFutureInbound

`func (o *ClusterIPAddresses) HasFutureInbound() bool`

HasFutureInbound returns a boolean if a field has been set.
### GetFutureOutbound

`func (o *ClusterIPAddresses) GetFutureOutbound() []string`

GetFutureOutbound returns the FutureOutbound field if non-nil, zero value otherwise.

### GetFutureOutboundOk

`func (o *ClusterIPAddresses) GetFutureOutboundOk() (*[]string, bool)`

GetFutureOutboundOk returns a tuple with the FutureOutbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureOutbound

`func (o *ClusterIPAddresses) SetFutureOutbound(v []string)`

SetFutureOutbound sets FutureOutbound field to given value.

### HasFutureOutbound

`func (o *ClusterIPAddresses) HasFutureOutbound() bool`

HasFutureOutbound returns a boolean if a field has been set.
### GetInbound

`func (o *ClusterIPAddresses) GetInbound() []string`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *ClusterIPAddresses) GetInboundOk() (*[]string, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *ClusterIPAddresses) SetInbound(v []string)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *ClusterIPAddresses) HasInbound() bool`

HasInbound returns a boolean if a field has been set.
### GetOutbound

`func (o *ClusterIPAddresses) GetOutbound() []string`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *ClusterIPAddresses) GetOutboundOk() (*[]string, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *ClusterIPAddresses) SetOutbound(v []string)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *ClusterIPAddresses) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


