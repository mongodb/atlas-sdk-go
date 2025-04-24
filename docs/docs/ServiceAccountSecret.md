# ServiceAccountSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The date that the secret was created on. | [readonly] 
**ExpiresAt** | **time.Time** | The date for the expiration of the secret. | [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the secret. | [readonly] 
**LastUsedAt** | Pointer to **time.Time** | The last time the secret was used. | [optional] [readonly] 
**MaskedSecretValue** | Pointer to **string** | The masked Service Account secret. | [optional] [readonly] 
**Secret** | Pointer to **string** | The secret for the Service Account. It will be returned only the first time after creation. | [optional] [readonly] 

## Methods

### NewServiceAccountSecret

`func NewServiceAccountSecret(createdAt time.Time, expiresAt time.Time, id string, ) *ServiceAccountSecret`

NewServiceAccountSecret instantiates a new ServiceAccountSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountSecretWithDefaults

`func NewServiceAccountSecretWithDefaults() *ServiceAccountSecret`

NewServiceAccountSecretWithDefaults instantiates a new ServiceAccountSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceAccountSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceAccountSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceAccountSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetExpiresAt

`func (o *ServiceAccountSecret) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ServiceAccountSecret) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ServiceAccountSecret) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### GetId

`func (o *ServiceAccountSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountSecret) SetId(v string)`

SetId sets Id field to given value.

### GetLastUsedAt

`func (o *ServiceAccountSecret) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ServiceAccountSecret) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ServiceAccountSecret) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ServiceAccountSecret) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.
### GetMaskedSecretValue

`func (o *ServiceAccountSecret) GetMaskedSecretValue() string`

GetMaskedSecretValue returns the MaskedSecretValue field if non-nil, zero value otherwise.

### GetMaskedSecretValueOk

`func (o *ServiceAccountSecret) GetMaskedSecretValueOk() (*string, bool)`

GetMaskedSecretValueOk returns a tuple with the MaskedSecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedSecretValue

`func (o *ServiceAccountSecret) SetMaskedSecretValue(v string)`

SetMaskedSecretValue sets MaskedSecretValue field to given value.

### HasMaskedSecretValue

`func (o *ServiceAccountSecret) HasMaskedSecretValue() bool`

HasMaskedSecretValue returns a boolean if a field has been set.
### GetSecret

`func (o *ServiceAccountSecret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ServiceAccountSecret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ServiceAccountSecret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ServiceAccountSecret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


