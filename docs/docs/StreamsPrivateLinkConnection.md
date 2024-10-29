# StreamsPrivateLinkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Private Link connection. | [optional] 
**DnsDomain** | Pointer to **string** | Domain name of Confluent cluster. | [optional] 
**DnsSubDomain** | Pointer to **[]string** | Sub-Domain name of Confluent cluster. These are typically your availability zones. | [optional] 
**InterfaceEndpointId** | Pointer to **string** | Interface endpoint ID that is created from the service endpoint ID provided. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Provider** | Pointer to **string** | Provider where the Kafka cluster is deployed. | [optional] 
**Region** | Pointer to **string** | Domain name of Confluent cluster. | [optional] 
**ServiceEndpointId** | Pointer to **string** | AWS Service Endpoint ID. | [optional] 
**State** | Pointer to **string** | State the connection is in. | [optional] 
**Vendor** | Pointer to **string** | Vendor who manages the Kafka cluster. | [optional] 

## Methods

### NewStreamsPrivateLinkConnection

`func NewStreamsPrivateLinkConnection() *StreamsPrivateLinkConnection`

NewStreamsPrivateLinkConnection instantiates a new StreamsPrivateLinkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsPrivateLinkConnectionWithDefaults

`func NewStreamsPrivateLinkConnectionWithDefaults() *StreamsPrivateLinkConnection`

NewStreamsPrivateLinkConnectionWithDefaults instantiates a new StreamsPrivateLinkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamsPrivateLinkConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsPrivateLinkConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsPrivateLinkConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamsPrivateLinkConnection) HasId() bool`

HasId returns a boolean if a field has been set.
### GetDnsDomain

`func (o *StreamsPrivateLinkConnection) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *StreamsPrivateLinkConnection) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *StreamsPrivateLinkConnection) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *StreamsPrivateLinkConnection) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.
### GetDnsSubDomain

`func (o *StreamsPrivateLinkConnection) GetDnsSubDomain() []string`

GetDnsSubDomain returns the DnsSubDomain field if non-nil, zero value otherwise.

### GetDnsSubDomainOk

`func (o *StreamsPrivateLinkConnection) GetDnsSubDomainOk() (*[]string, bool)`

GetDnsSubDomainOk returns a tuple with the DnsSubDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubDomain

`func (o *StreamsPrivateLinkConnection) SetDnsSubDomain(v []string)`

SetDnsSubDomain sets DnsSubDomain field to given value.

### HasDnsSubDomain

`func (o *StreamsPrivateLinkConnection) HasDnsSubDomain() bool`

HasDnsSubDomain returns a boolean if a field has been set.
### GetInterfaceEndpointId

`func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointId() string`

GetInterfaceEndpointId returns the InterfaceEndpointId field if non-nil, zero value otherwise.

### GetInterfaceEndpointIdOk

`func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointIdOk() (*string, bool)`

GetInterfaceEndpointIdOk returns a tuple with the InterfaceEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpointId

`func (o *StreamsPrivateLinkConnection) SetInterfaceEndpointId(v string)`

SetInterfaceEndpointId sets InterfaceEndpointId field to given value.

### HasInterfaceEndpointId

`func (o *StreamsPrivateLinkConnection) HasInterfaceEndpointId() bool`

HasInterfaceEndpointId returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsPrivateLinkConnection) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsPrivateLinkConnection) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsPrivateLinkConnection) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsPrivateLinkConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetProvider

`func (o *StreamsPrivateLinkConnection) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *StreamsPrivateLinkConnection) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *StreamsPrivateLinkConnection) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *StreamsPrivateLinkConnection) HasProvider() bool`

HasProvider returns a boolean if a field has been set.
### GetRegion

`func (o *StreamsPrivateLinkConnection) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsPrivateLinkConnection) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsPrivateLinkConnection) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StreamsPrivateLinkConnection) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetServiceEndpointId

`func (o *StreamsPrivateLinkConnection) GetServiceEndpointId() string`

GetServiceEndpointId returns the ServiceEndpointId field if non-nil, zero value otherwise.

### GetServiceEndpointIdOk

`func (o *StreamsPrivateLinkConnection) GetServiceEndpointIdOk() (*string, bool)`

GetServiceEndpointIdOk returns a tuple with the ServiceEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpointId

`func (o *StreamsPrivateLinkConnection) SetServiceEndpointId(v string)`

SetServiceEndpointId sets ServiceEndpointId field to given value.

### HasServiceEndpointId

`func (o *StreamsPrivateLinkConnection) HasServiceEndpointId() bool`

HasServiceEndpointId returns a boolean if a field has been set.
### GetState

`func (o *StreamsPrivateLinkConnection) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamsPrivateLinkConnection) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamsPrivateLinkConnection) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StreamsPrivateLinkConnection) HasState() bool`

HasState returns a boolean if a field has been set.
### GetVendor

`func (o *StreamsPrivateLinkConnection) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StreamsPrivateLinkConnection) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StreamsPrivateLinkConnection) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StreamsPrivateLinkConnection) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


