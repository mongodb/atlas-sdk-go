# StreamsPrivateLinkConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsDomain** | Pointer to **string** | The domain hostname for the AWS Confluent Serverless Private Link connection. Allowed only when no domain is currently set, or when the connection is in &#x60;IDLE&#x60; state. | [optional] 

## Methods

### NewStreamsPrivateLinkConnectionRequest

`func NewStreamsPrivateLinkConnectionRequest() *StreamsPrivateLinkConnectionRequest`

NewStreamsPrivateLinkConnectionRequest instantiates a new StreamsPrivateLinkConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsPrivateLinkConnectionRequestWithDefaults

`func NewStreamsPrivateLinkConnectionRequestWithDefaults() *StreamsPrivateLinkConnectionRequest`

NewStreamsPrivateLinkConnectionRequestWithDefaults instantiates a new StreamsPrivateLinkConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsDomain

`func (o *StreamsPrivateLinkConnectionRequest) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *StreamsPrivateLinkConnectionRequest) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *StreamsPrivateLinkConnectionRequest) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *StreamsPrivateLinkConnectionRequest) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### SetDnsDomainNil

`func (o *StreamsPrivateLinkConnectionRequest) SetDnsDomainNil()`

SetDnsDomainNil sets DnsDomain to an explicit JSON null when marshaled, overriding any value previously set with SetDnsDomain. Calling SetDnsDomain again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


