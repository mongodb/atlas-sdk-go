# PrometheusConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ListenAddress** | Pointer to **string** |  | [optional] 
**MetricsPath** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**RateLimitIntervalSeconds** | Pointer to **int** |  | [optional] 
**Scheme** | Pointer to **string** |  | [optional] 
**ServiceDiscovery** | Pointer to **string** |  | [optional] 
**TlsPemPassword** | Pointer to **string** |  | [optional] 
**TlsPemPath** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewPrometheusConfig

`func NewPrometheusConfig() *PrometheusConfig`

NewPrometheusConfig instantiates a new PrometheusConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusConfigWithDefaults

`func NewPrometheusConfigWithDefaults() *PrometheusConfig`

NewPrometheusConfigWithDefaults instantiates a new PrometheusConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PrometheusConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PrometheusConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PrometheusConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PrometheusConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.
### GetListenAddress

`func (o *PrometheusConfig) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *PrometheusConfig) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *PrometheusConfig) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *PrometheusConfig) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.
### GetMetricsPath

`func (o *PrometheusConfig) GetMetricsPath() string`

GetMetricsPath returns the MetricsPath field if non-nil, zero value otherwise.

### GetMetricsPathOk

`func (o *PrometheusConfig) GetMetricsPathOk() (*string, bool)`

GetMetricsPathOk returns a tuple with the MetricsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPath

`func (o *PrometheusConfig) SetMetricsPath(v string)`

SetMetricsPath sets MetricsPath field to given value.

### HasMetricsPath

`func (o *PrometheusConfig) HasMetricsPath() bool`

HasMetricsPath returns a boolean if a field has been set.
### GetMode

`func (o *PrometheusConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PrometheusConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PrometheusConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PrometheusConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.
### GetRateLimitIntervalSeconds

`func (o *PrometheusConfig) GetRateLimitIntervalSeconds() int`

GetRateLimitIntervalSeconds returns the RateLimitIntervalSeconds field if non-nil, zero value otherwise.

### GetRateLimitIntervalSecondsOk

`func (o *PrometheusConfig) GetRateLimitIntervalSecondsOk() (*int, bool)`

GetRateLimitIntervalSecondsOk returns a tuple with the RateLimitIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitIntervalSeconds

`func (o *PrometheusConfig) SetRateLimitIntervalSeconds(v int)`

SetRateLimitIntervalSeconds sets RateLimitIntervalSeconds field to given value.

### HasRateLimitIntervalSeconds

`func (o *PrometheusConfig) HasRateLimitIntervalSeconds() bool`

HasRateLimitIntervalSeconds returns a boolean if a field has been set.
### GetScheme

`func (o *PrometheusConfig) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PrometheusConfig) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PrometheusConfig) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PrometheusConfig) HasScheme() bool`

HasScheme returns a boolean if a field has been set.
### GetServiceDiscovery

`func (o *PrometheusConfig) GetServiceDiscovery() string`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *PrometheusConfig) GetServiceDiscoveryOk() (*string, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *PrometheusConfig) SetServiceDiscovery(v string)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *PrometheusConfig) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.
### GetTlsPemPassword

`func (o *PrometheusConfig) GetTlsPemPassword() string`

GetTlsPemPassword returns the TlsPemPassword field if non-nil, zero value otherwise.

### GetTlsPemPasswordOk

`func (o *PrometheusConfig) GetTlsPemPasswordOk() (*string, bool)`

GetTlsPemPasswordOk returns a tuple with the TlsPemPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPemPassword

`func (o *PrometheusConfig) SetTlsPemPassword(v string)`

SetTlsPemPassword sets TlsPemPassword field to given value.

### HasTlsPemPassword

`func (o *PrometheusConfig) HasTlsPemPassword() bool`

HasTlsPemPassword returns a boolean if a field has been set.
### GetTlsPemPath

`func (o *PrometheusConfig) GetTlsPemPath() string`

GetTlsPemPath returns the TlsPemPath field if non-nil, zero value otherwise.

### GetTlsPemPathOk

`func (o *PrometheusConfig) GetTlsPemPathOk() (*string, bool)`

GetTlsPemPathOk returns a tuple with the TlsPemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPemPath

`func (o *PrometheusConfig) SetTlsPemPath(v string)`

SetTlsPemPath sets TlsPemPath field to given value.

### HasTlsPemPath

`func (o *PrometheusConfig) HasTlsPemPath() bool`

HasTlsPemPath returns a boolean if a field has been set.
### GetUsername

`func (o *PrometheusConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PrometheusConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PrometheusConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PrometheusConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


