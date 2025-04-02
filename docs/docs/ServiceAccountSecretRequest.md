# ServiceAccountSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretExpiresAfterHours** | **int** | The expiration time of the new Service Account secret, provided in hours. The minimum and maximum allowed expiration times are subject to change and are controlled by the organization&#39;s settings. | 

## Methods

### NewServiceAccountSecretRequest

`func NewServiceAccountSecretRequest(secretExpiresAfterHours int, ) *ServiceAccountSecretRequest`

NewServiceAccountSecretRequest instantiates a new ServiceAccountSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountSecretRequestWithDefaults

`func NewServiceAccountSecretRequestWithDefaults() *ServiceAccountSecretRequest`

NewServiceAccountSecretRequestWithDefaults instantiates a new ServiceAccountSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretExpiresAfterHours

`func (o *ServiceAccountSecretRequest) GetSecretExpiresAfterHours() int`

GetSecretExpiresAfterHours returns the SecretExpiresAfterHours field if non-nil, zero value otherwise.

### GetSecretExpiresAfterHoursOk

`func (o *ServiceAccountSecretRequest) GetSecretExpiresAfterHoursOk() (*int, bool)`

GetSecretExpiresAfterHoursOk returns a tuple with the SecretExpiresAfterHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpiresAfterHours

`func (o *ServiceAccountSecretRequest) SetSecretExpiresAfterHours(v int)`

SetSecretExpiresAfterHours sets SecretExpiresAfterHours field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


