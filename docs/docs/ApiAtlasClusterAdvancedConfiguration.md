# ApiAtlasClusterAdvancedConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomOpensslCipherConfigTls12** | Pointer to **[]string** | The custom OpenSSL cipher suite list for TLS 1.2. This field is only valid when &#x60;tlsCipherConfigMode&#x60; is set to &#x60;CUSTOM&#x60;. | [optional] 
**CustomOpensslCipherConfigTls13** | Pointer to **[]string** | The custom OpenSSL cipher suite list for TLS 1.3. This field is only valid when &#x60;tlsCipherConfigMode&#x60; is set to &#x60;CUSTOM&#x60;. | [optional] 
**MinimumEnabledTlsProtocol** | Pointer to **string** | Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections. Clusters using TLS 1.0 or 1.1 should consider setting TLS 1.2 as the minimum TLS protocol version. | [optional] 
**TlsCipherConfigMode** | Pointer to **string** | The TLS cipher suite configuration mode. The default mode uses the default cipher suites. The custom mode allows you to specify custom cipher suites for both TLS 1.2 and TLS 1.3. | [optional] 

## Methods

### NewApiAtlasClusterAdvancedConfiguration

`func NewApiAtlasClusterAdvancedConfiguration() *ApiAtlasClusterAdvancedConfiguration`

NewApiAtlasClusterAdvancedConfiguration instantiates a new ApiAtlasClusterAdvancedConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterAdvancedConfigurationWithDefaults

`func NewApiAtlasClusterAdvancedConfigurationWithDefaults() *ApiAtlasClusterAdvancedConfiguration`

NewApiAtlasClusterAdvancedConfigurationWithDefaults instantiates a new ApiAtlasClusterAdvancedConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomOpensslCipherConfigTls12

`func (o *ApiAtlasClusterAdvancedConfiguration) GetCustomOpensslCipherConfigTls12() []string`

GetCustomOpensslCipherConfigTls12 returns the CustomOpensslCipherConfigTls12 field if non-nil, zero value otherwise.

### GetCustomOpensslCipherConfigTls12Ok

`func (o *ApiAtlasClusterAdvancedConfiguration) GetCustomOpensslCipherConfigTls12Ok() (*[]string, bool)`

GetCustomOpensslCipherConfigTls12Ok returns a tuple with the CustomOpensslCipherConfigTls12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOpensslCipherConfigTls12

`func (o *ApiAtlasClusterAdvancedConfiguration) SetCustomOpensslCipherConfigTls12(v []string)`

SetCustomOpensslCipherConfigTls12 sets CustomOpensslCipherConfigTls12 field to given value.

### HasCustomOpensslCipherConfigTls12

`func (o *ApiAtlasClusterAdvancedConfiguration) HasCustomOpensslCipherConfigTls12() bool`

HasCustomOpensslCipherConfigTls12 returns a boolean if a field has been set.
### GetCustomOpensslCipherConfigTls13

`func (o *ApiAtlasClusterAdvancedConfiguration) GetCustomOpensslCipherConfigTls13() []string`

GetCustomOpensslCipherConfigTls13 returns the CustomOpensslCipherConfigTls13 field if non-nil, zero value otherwise.

### GetCustomOpensslCipherConfigTls13Ok

`func (o *ApiAtlasClusterAdvancedConfiguration) GetCustomOpensslCipherConfigTls13Ok() (*[]string, bool)`

GetCustomOpensslCipherConfigTls13Ok returns a tuple with the CustomOpensslCipherConfigTls13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOpensslCipherConfigTls13

`func (o *ApiAtlasClusterAdvancedConfiguration) SetCustomOpensslCipherConfigTls13(v []string)`

SetCustomOpensslCipherConfigTls13 sets CustomOpensslCipherConfigTls13 field to given value.

### HasCustomOpensslCipherConfigTls13

`func (o *ApiAtlasClusterAdvancedConfiguration) HasCustomOpensslCipherConfigTls13() bool`

HasCustomOpensslCipherConfigTls13 returns a boolean if a field has been set.
### GetMinimumEnabledTlsProtocol

`func (o *ApiAtlasClusterAdvancedConfiguration) GetMinimumEnabledTlsProtocol() string`

GetMinimumEnabledTlsProtocol returns the MinimumEnabledTlsProtocol field if non-nil, zero value otherwise.

### GetMinimumEnabledTlsProtocolOk

`func (o *ApiAtlasClusterAdvancedConfiguration) GetMinimumEnabledTlsProtocolOk() (*string, bool)`

GetMinimumEnabledTlsProtocolOk returns a tuple with the MinimumEnabledTlsProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumEnabledTlsProtocol

`func (o *ApiAtlasClusterAdvancedConfiguration) SetMinimumEnabledTlsProtocol(v string)`

SetMinimumEnabledTlsProtocol sets MinimumEnabledTlsProtocol field to given value.

### HasMinimumEnabledTlsProtocol

`func (o *ApiAtlasClusterAdvancedConfiguration) HasMinimumEnabledTlsProtocol() bool`

HasMinimumEnabledTlsProtocol returns a boolean if a field has been set.
### GetTlsCipherConfigMode

`func (o *ApiAtlasClusterAdvancedConfiguration) GetTlsCipherConfigMode() string`

GetTlsCipherConfigMode returns the TlsCipherConfigMode field if non-nil, zero value otherwise.

### GetTlsCipherConfigModeOk

`func (o *ApiAtlasClusterAdvancedConfiguration) GetTlsCipherConfigModeOk() (*string, bool)`

GetTlsCipherConfigModeOk returns a tuple with the TlsCipherConfigMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherConfigMode

`func (o *ApiAtlasClusterAdvancedConfiguration) SetTlsCipherConfigMode(v string)`

SetTlsCipherConfigMode sets TlsCipherConfigMode field to given value.

### HasTlsCipherConfigMode

`func (o *ApiAtlasClusterAdvancedConfiguration) HasTlsCipherConfigMode() bool`

HasTlsCipherConfigMode returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


