# StreamsKafkaAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | OIDC client identifier for authentication to the Kafka cluster. | [optional] 
**ClientSecret** | Pointer to **string** | OIDC client secret for authentication to the Kafka cluster. | [optional] 
**HttpsCaPem** | Pointer to **string** | HTTPS CA certificate in PEM format for SSL/TLS verification. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Mechanism** | Pointer to **string** | Style of authentication. Can be one of PLAIN, SCRAM-256, SCRAM-512, or OAUTHBEARER. | [optional] 
**Password** | Pointer to **string** | Password of the account to connect to the Kafka cluster. | [optional] 
**SaslOauthbearerExtensions** | Pointer to **string** | SASL OAUTHBEARER extensions parameter for additional OAuth2 configuration. | [optional] 
**Scope** | Pointer to **string** | OIDC scope parameter defining the access permissions requested. | [optional] 
**SslCertificate** | Pointer to **string** | SSL certificate for client authentication to Kafka. | [optional] 
**SslKey** | Pointer to **string** | SSL key for client authentication to Kafka. | [optional] 
**SslKeyPassword** | Pointer to **string** | Password for the SSL key, if it is password protected. | [optional] 
**TokenEndpointUrl** | Pointer to **string** | OIDC token endpoint URL for obtaining access tokens. | [optional] 
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

### GetClientId

`func (o *StreamsKafkaAuthentication) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StreamsKafkaAuthentication) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StreamsKafkaAuthentication) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *StreamsKafkaAuthentication) HasClientId() bool`

HasClientId returns a boolean if a field has been set.
### GetClientSecret

`func (o *StreamsKafkaAuthentication) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *StreamsKafkaAuthentication) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *StreamsKafkaAuthentication) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *StreamsKafkaAuthentication) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.
### GetHttpsCaPem

`func (o *StreamsKafkaAuthentication) GetHttpsCaPem() string`

GetHttpsCaPem returns the HttpsCaPem field if non-nil, zero value otherwise.

### GetHttpsCaPemOk

`func (o *StreamsKafkaAuthentication) GetHttpsCaPemOk() (*string, bool)`

GetHttpsCaPemOk returns a tuple with the HttpsCaPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsCaPem

`func (o *StreamsKafkaAuthentication) SetHttpsCaPem(v string)`

SetHttpsCaPem sets HttpsCaPem field to given value.

### HasHttpsCaPem

`func (o *StreamsKafkaAuthentication) HasHttpsCaPem() bool`

HasHttpsCaPem returns a boolean if a field has been set.
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
### GetSaslOauthbearerExtensions

`func (o *StreamsKafkaAuthentication) GetSaslOauthbearerExtensions() string`

GetSaslOauthbearerExtensions returns the SaslOauthbearerExtensions field if non-nil, zero value otherwise.

### GetSaslOauthbearerExtensionsOk

`func (o *StreamsKafkaAuthentication) GetSaslOauthbearerExtensionsOk() (*string, bool)`

GetSaslOauthbearerExtensionsOk returns a tuple with the SaslOauthbearerExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaslOauthbearerExtensions

`func (o *StreamsKafkaAuthentication) SetSaslOauthbearerExtensions(v string)`

SetSaslOauthbearerExtensions sets SaslOauthbearerExtensions field to given value.

### HasSaslOauthbearerExtensions

`func (o *StreamsKafkaAuthentication) HasSaslOauthbearerExtensions() bool`

HasSaslOauthbearerExtensions returns a boolean if a field has been set.
### GetScope

`func (o *StreamsKafkaAuthentication) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *StreamsKafkaAuthentication) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *StreamsKafkaAuthentication) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *StreamsKafkaAuthentication) HasScope() bool`

HasScope returns a boolean if a field has been set.
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
### GetTokenEndpointUrl

`func (o *StreamsKafkaAuthentication) GetTokenEndpointUrl() string`

GetTokenEndpointUrl returns the TokenEndpointUrl field if non-nil, zero value otherwise.

### GetTokenEndpointUrlOk

`func (o *StreamsKafkaAuthentication) GetTokenEndpointUrlOk() (*string, bool)`

GetTokenEndpointUrlOk returns a tuple with the TokenEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointUrl

`func (o *StreamsKafkaAuthentication) SetTokenEndpointUrl(v string)`

SetTokenEndpointUrl sets TokenEndpointUrl field to given value.

### HasTokenEndpointUrl

`func (o *StreamsKafkaAuthentication) HasTokenEndpointUrl() bool`

HasTokenEndpointUrl returns a boolean if a field has been set.
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


