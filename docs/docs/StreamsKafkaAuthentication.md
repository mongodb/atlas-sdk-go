# StreamsKafkaAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Mechanism** | Pointer to **string** | Style of authentication. Can be one of PLAIN, SCRAM-256, or SCRAM-512. | [optional] 
**Password** | Pointer to **string** | Password of the account to connect to the Kafka cluster. | [optional] 
**SslCertificate** | Pointer to **string** | SSL certificate for client authentication to Kafka. | [optional] 
**SslKey** | Pointer to **string** | SSL key for client authentication to Kafka. | [optional] 
**SslKeyPassword** | Pointer to **string** | Password for the SSL key, if it is password protected. | [optional] 
**Username** | Pointer to **string** | Username of the account to connect to the Kafka cluster. | [optional] 

## Methods

### NewStreamsKafkaAuthentication

`func NewStreamsKafkaAuthentication() *StreamsKafkaAuthentication`

NewStreamsKafkaAuthentication instantiates a new StreamsKafkaAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsKafkaAuthenticationWithDefaults

`func NewStreamsKafkaAuthenticationWithDefaults() *StreamsKafkaAuthentication`

NewStreamsKafkaAuthenticationWithDefaults instantiates a new StreamsKafkaAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsKafkaAuthentication) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsKafkaAuthentication) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsKafkaAuthentication) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsKafkaAuthentication) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMechanism

`func (o *StreamsKafkaAuthentication) GetMechanism() string`

GetMechanism returns the Mechanism field if non-nil, zero value otherwise.

### GetMechanismOk

`func (o *StreamsKafkaAuthentication) GetMechanismOk() (*string, bool)`

GetMechanismOk returns a tuple with the Mechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanism

`func (o *StreamsKafkaAuthentication) SetMechanism(v string)`

SetMechanism sets Mechanism field to given value.

### HasMechanism

`func (o *StreamsKafkaAuthentication) HasMechanism() bool`

HasMechanism returns a boolean if a field has been set.
### GetPassword

`func (o *StreamsKafkaAuthentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StreamsKafkaAuthentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StreamsKafkaAuthentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StreamsKafkaAuthentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.
### GetSslCertificate

`func (o *StreamsKafkaAuthentication) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *StreamsKafkaAuthentication) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *StreamsKafkaAuthentication) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *StreamsKafkaAuthentication) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.
### GetSslKey

`func (o *StreamsKafkaAuthentication) GetSslKey() string`

GetSslKey returns the SslKey field if non-nil, zero value otherwise.

### GetSslKeyOk

`func (o *StreamsKafkaAuthentication) GetSslKeyOk() (*string, bool)`

GetSslKeyOk returns a tuple with the SslKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslKey

`func (o *StreamsKafkaAuthentication) SetSslKey(v string)`

SetSslKey sets SslKey field to given value.

### HasSslKey

`func (o *StreamsKafkaAuthentication) HasSslKey() bool`

HasSslKey returns a boolean if a field has been set.
### GetSslKeyPassword

`func (o *StreamsKafkaAuthentication) GetSslKeyPassword() string`

GetSslKeyPassword returns the SslKeyPassword field if non-nil, zero value otherwise.

### GetSslKeyPasswordOk

`func (o *StreamsKafkaAuthentication) GetSslKeyPasswordOk() (*string, bool)`

GetSslKeyPasswordOk returns a tuple with the SslKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslKeyPassword

`func (o *StreamsKafkaAuthentication) SetSslKeyPassword(v string)`

SetSslKeyPassword sets SslKeyPassword field to given value.

### HasSslKeyPassword

`func (o *StreamsKafkaAuthentication) HasSslKeyPassword() bool`

HasSslKeyPassword returns a boolean if a field has been set.
### GetUsername

`func (o *StreamsKafkaAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StreamsKafkaAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StreamsKafkaAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StreamsKafkaAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


