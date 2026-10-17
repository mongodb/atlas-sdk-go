# StreamsWorkspacePrivateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerDnsName** | **string** | DNS name of the customer&#39;s VPC endpoint. | 
**Id** | **string** | Unique identifier of the customer&#39;s VPC endpoint. | 
**Region** | **string** | Human-readable label identifying the region of the customer&#39;s VPC endpoint. | 

## Methods

### NewStreamsWorkspacePrivateEndpointRequest

`func NewStreamsWorkspacePrivateEndpointRequest(customerDnsName string, id string, region string, ) *StreamsWorkspacePrivateEndpointRequest`

NewStreamsWorkspacePrivateEndpointRequest instantiates a new StreamsWorkspacePrivateEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsWorkspacePrivateEndpointRequestWithDefaults

`func NewStreamsWorkspacePrivateEndpointRequestWithDefaults() *StreamsWorkspacePrivateEndpointRequest`

NewStreamsWorkspacePrivateEndpointRequestWithDefaults instantiates a new StreamsWorkspacePrivateEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerDnsName

`func (o *StreamsWorkspacePrivateEndpointRequest) GetCustomerDnsName() string`

GetCustomerDnsName returns the CustomerDnsName field if non-nil, zero value otherwise.

### GetCustomerDnsNameOk

`func (o *StreamsWorkspacePrivateEndpointRequest) GetCustomerDnsNameOk() (*string, bool)`

GetCustomerDnsNameOk returns a tuple with the CustomerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDnsName

`func (o *StreamsWorkspacePrivateEndpointRequest) SetCustomerDnsName(v string)`

SetCustomerDnsName sets CustomerDnsName field to given value.

### GetId

`func (o *StreamsWorkspacePrivateEndpointRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsWorkspacePrivateEndpointRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsWorkspacePrivateEndpointRequest) SetId(v string)`

SetId sets Id field to given value.

### GetRegion

`func (o *StreamsWorkspacePrivateEndpointRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsWorkspacePrivateEndpointRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsWorkspacePrivateEndpointRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


