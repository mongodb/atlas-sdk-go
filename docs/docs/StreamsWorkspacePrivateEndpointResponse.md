# StreamsWorkspacePrivateEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerDnsName** | Pointer to **string** | DNS name of the customer&#39;s VPC endpoint. | [optional] 
**EndpointServiceName** | Pointer to **string** | Name of the Atlas endpoint service the customer&#39;s VPC endpoint connects to. Absent when no endpoint service is configured for the endpoint&#39;s region. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier of the customer&#39;s VPC endpoint. | [optional] 
**Provider** | Pointer to **string** | Cloud provider hosting the customer&#39;s VPC endpoint. | [optional] [readonly] 
**Region** | Pointer to **string** | Human-readable label identifying the region of the customer&#39;s VPC endpoint. | [optional] 
**Status** | Pointer to **string** | Status of the endpoint registration. | [optional] [readonly] 

## Methods

### NewStreamsWorkspacePrivateEndpointResponse

`func NewStreamsWorkspacePrivateEndpointResponse() *StreamsWorkspacePrivateEndpointResponse`

NewStreamsWorkspacePrivateEndpointResponse instantiates a new StreamsWorkspacePrivateEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsWorkspacePrivateEndpointResponseWithDefaults

`func NewStreamsWorkspacePrivateEndpointResponseWithDefaults() *StreamsWorkspacePrivateEndpointResponse`

NewStreamsWorkspacePrivateEndpointResponseWithDefaults instantiates a new StreamsWorkspacePrivateEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerDnsName

`func (o *StreamsWorkspacePrivateEndpointResponse) GetCustomerDnsName() string`

GetCustomerDnsName returns the CustomerDnsName field if non-nil, zero value otherwise.

### GetCustomerDnsNameOk

`func (o *StreamsWorkspacePrivateEndpointResponse) GetCustomerDnsNameOk() (*string, bool)`

GetCustomerDnsNameOk returns a tuple with the CustomerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDnsName

`func (o *StreamsWorkspacePrivateEndpointResponse) SetCustomerDnsName(v string)`

SetCustomerDnsName sets CustomerDnsName field to given value.

### HasCustomerDnsName

`func (o *StreamsWorkspacePrivateEndpointResponse) HasCustomerDnsName() bool`

HasCustomerDnsName returns a boolean if a field has been set.

### SetCustomerDnsNameNil

`func (o *StreamsWorkspacePrivateEndpointResponse) SetCustomerDnsNameNil()`

SetCustomerDnsNameNil sets CustomerDnsName to an explicit JSON null when marshaled, overriding any value previously set with SetCustomerDnsName. Calling SetCustomerDnsName again clears the null override.

### GetEndpointServiceName

`func (o *StreamsWorkspacePrivateEndpointResponse) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *StreamsWorkspacePrivateEndpointResponse) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *StreamsWorkspacePrivateEndpointResponse) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *StreamsWorkspacePrivateEndpointResponse) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### SetEndpointServiceNameNil

`func (o *StreamsWorkspacePrivateEndpointResponse) SetEndpointServiceNameNil()`

SetEndpointServiceNameNil sets EndpointServiceName to an explicit JSON null when marshaled, overriding any value previously set with SetEndpointServiceName. Calling SetEndpointServiceName again clears the null override.

### GetId

`func (o *StreamsWorkspacePrivateEndpointResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsWorkspacePrivateEndpointResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsWorkspacePrivateEndpointResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamsWorkspacePrivateEndpointResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StreamsWorkspacePrivateEndpointResponse) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetProvider

`func (o *StreamsWorkspacePrivateEndpointResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *StreamsWorkspacePrivateEndpointResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *StreamsWorkspacePrivateEndpointResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *StreamsWorkspacePrivateEndpointResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *StreamsWorkspacePrivateEndpointResponse) SetProviderNil()`

SetProviderNil sets Provider to an explicit JSON null when marshaled, overriding any value previously set with SetProvider. Calling SetProvider again clears the null override.

### GetRegion

`func (o *StreamsWorkspacePrivateEndpointResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsWorkspacePrivateEndpointResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsWorkspacePrivateEndpointResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StreamsWorkspacePrivateEndpointResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *StreamsWorkspacePrivateEndpointResponse) SetRegionNil()`

SetRegionNil sets Region to an explicit JSON null when marshaled, overriding any value previously set with SetRegion. Calling SetRegion again clears the null override.

### GetStatus

`func (o *StreamsWorkspacePrivateEndpointResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StreamsWorkspacePrivateEndpointResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StreamsWorkspacePrivateEndpointResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StreamsWorkspacePrivateEndpointResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StreamsWorkspacePrivateEndpointResponse) SetStatusNil()`

SetStatusNil sets Status to an explicit JSON null when marshaled, overriding any value previously set with SetStatus. Calling SetStatus again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


