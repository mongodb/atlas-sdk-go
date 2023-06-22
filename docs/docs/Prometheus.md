# Prometheus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Flag that indicates whether someone has activated the Prometheus integration. | 
**ListenAddress** | Pointer to **string** | Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics. | [optional] [default to ":9216"]
**Password** | Pointer to **string** |  | [optional] 
**RateLimitInterval** | Pointer to **int** |  | [optional] 
**Scheme** | **string** | Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud. | 
**ServiceDiscovery** | **string** | Desired method to discover the Prometheus service. | 
**TlsPemPath** | Pointer to **string** | Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host. | [optional] 
**Type** | Pointer to **string** | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type. | [optional] 
**Username** | **string** | Human-readable label that identifies your Prometheus incoming webhook. | 

## Methods

### NewPrometheus

`func NewPrometheus(enabled bool, scheme string, serviceDiscovery string, username string, ) *Prometheus`

NewPrometheus instantiates a new Prometheus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusWithDefaults

`func NewPrometheusWithDefaults() *Prometheus`

NewPrometheusWithDefaults instantiates a new Prometheus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Prometheus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Prometheus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Prometheus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetListenAddress

`func (o *Prometheus) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *Prometheus) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *Prometheus) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *Prometheus) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetPassword

`func (o *Prometheus) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Prometheus) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Prometheus) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Prometheus) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRateLimitInterval

`func (o *Prometheus) GetRateLimitInterval() int`

GetRateLimitInterval returns the RateLimitInterval field if non-nil, zero value otherwise.

### GetRateLimitIntervalOk

`func (o *Prometheus) GetRateLimitIntervalOk() (*int, bool)`

GetRateLimitIntervalOk returns a tuple with the RateLimitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitInterval

`func (o *Prometheus) SetRateLimitInterval(v int)`

SetRateLimitInterval sets RateLimitInterval field to given value.

### HasRateLimitInterval

`func (o *Prometheus) HasRateLimitInterval() bool`

HasRateLimitInterval returns a boolean if a field has been set.

### GetScheme

`func (o *Prometheus) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *Prometheus) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *Prometheus) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetServiceDiscovery

`func (o *Prometheus) GetServiceDiscovery() string`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *Prometheus) GetServiceDiscoveryOk() (*string, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *Prometheus) SetServiceDiscovery(v string)`

SetServiceDiscovery sets ServiceDiscovery field to given value.


### GetTlsPemPath

`func (o *Prometheus) GetTlsPemPath() string`

GetTlsPemPath returns the TlsPemPath field if non-nil, zero value otherwise.

### GetTlsPemPathOk

`func (o *Prometheus) GetTlsPemPathOk() (*string, bool)`

GetTlsPemPathOk returns a tuple with the TlsPemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPemPath

`func (o *Prometheus) SetTlsPemPath(v string)`

SetTlsPemPath sets TlsPemPath field to given value.

### HasTlsPemPath

`func (o *Prometheus) HasTlsPemPath() bool`

HasTlsPemPath returns a boolean if a field has been set.

### GetType

`func (o *Prometheus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Prometheus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Prometheus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Prometheus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *Prometheus) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Prometheus) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Prometheus) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


