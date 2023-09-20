# StreamsKafkaSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerPublicCertificate** | Pointer to **string** | A trusted, public x509 certificate for connecting to Kafka over SSL. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Protocol** | Pointer to **string** | Describes the transport type. Can be either PLAINTEXT or SSL. | [optional] 

## Methods

### NewStreamsKafkaSecurity

`func NewStreamsKafkaSecurity() *StreamsKafkaSecurity`

NewStreamsKafkaSecurity instantiates a new StreamsKafkaSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsKafkaSecurityWithDefaults

`func NewStreamsKafkaSecurityWithDefaults() *StreamsKafkaSecurity`

NewStreamsKafkaSecurityWithDefaults instantiates a new StreamsKafkaSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerPublicCertificate

`func (o *StreamsKafkaSecurity) GetBrokerPublicCertificate() string`

GetBrokerPublicCertificate returns the BrokerPublicCertificate field if non-nil, zero value otherwise.

### GetBrokerPublicCertificateOk

`func (o *StreamsKafkaSecurity) GetBrokerPublicCertificateOk() (*string, bool)`

GetBrokerPublicCertificateOk returns a tuple with the BrokerPublicCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerPublicCertificate

`func (o *StreamsKafkaSecurity) SetBrokerPublicCertificate(v string)`

SetBrokerPublicCertificate sets BrokerPublicCertificate field to given value.

### HasBrokerPublicCertificate

`func (o *StreamsKafkaSecurity) HasBrokerPublicCertificate() bool`

HasBrokerPublicCertificate returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsKafkaSecurity) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsKafkaSecurity) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsKafkaSecurity) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsKafkaSecurity) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetProtocol

`func (o *StreamsKafkaSecurity) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StreamsKafkaSecurity) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StreamsKafkaSecurity) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StreamsKafkaSecurity) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


