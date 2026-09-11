# GroupServiceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Human readable description for the Service Account. | 
**Name** | **string** | Human-readable name for the Service Account. The name is modifiable and does not have to be unique. | 
**Roles** | **[]string** | A list of project-level roles for the Service Account. | 
**SecretExpiresAfterHours** | Pointer to **int** | The expiration time of the new Service Account secret, provided in hours. The minimum and maximum allowed expiration times are subject to change and are controlled by the organization&#39;s settings. Required unless &#x60;withoutInitialSecret&#x60; is true. | [optional] 
**WithoutInitialSecret** | Pointer to **bool** | If true, creates the Service Account without generating an initial secret. &#x60;secretExpiresAfterHours&#x60; must not be set when this is true. Defaults to false, which preserves existing behavior: a secret is generated and returned in the response. Use the &#x60;CreateGroupServiceAccountSecret&#x60; endpoint to add a secret later. | [optional] [default to false]

## Methods

### NewGroupServiceAccountRequest

`func NewGroupServiceAccountRequest(description string, name string, roles []string, ) *GroupServiceAccountRequest`

NewGroupServiceAccountRequest instantiates a new GroupServiceAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupServiceAccountRequestWithDefaults

`func NewGroupServiceAccountRequestWithDefaults() *GroupServiceAccountRequest`

NewGroupServiceAccountRequestWithDefaults instantiates a new GroupServiceAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupServiceAccountRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupServiceAccountRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupServiceAccountRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetName

`func (o *GroupServiceAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupServiceAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupServiceAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### GetRoles

`func (o *GroupServiceAccountRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupServiceAccountRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupServiceAccountRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### GetSecretExpiresAfterHours

`func (o *GroupServiceAccountRequest) GetSecretExpiresAfterHours() int`

GetSecretExpiresAfterHours returns the SecretExpiresAfterHours field if non-nil, zero value otherwise.

### GetSecretExpiresAfterHoursOk

`func (o *GroupServiceAccountRequest) GetSecretExpiresAfterHoursOk() (*int, bool)`

GetSecretExpiresAfterHoursOk returns a tuple with the SecretExpiresAfterHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpiresAfterHours

`func (o *GroupServiceAccountRequest) SetSecretExpiresAfterHours(v int)`

SetSecretExpiresAfterHours sets SecretExpiresAfterHours field to given value.

### HasSecretExpiresAfterHours

`func (o *GroupServiceAccountRequest) HasSecretExpiresAfterHours() bool`

HasSecretExpiresAfterHours returns a boolean if a field has been set.

### SetSecretExpiresAfterHoursNil

`func (o *GroupServiceAccountRequest) SetSecretExpiresAfterHoursNil()`

SetSecretExpiresAfterHoursNil sets SecretExpiresAfterHours to an explicit JSON null when marshaled, overriding any value previously set with SetSecretExpiresAfterHours. Calling SetSecretExpiresAfterHours again clears the null override.

### GetWithoutInitialSecret

`func (o *GroupServiceAccountRequest) GetWithoutInitialSecret() bool`

GetWithoutInitialSecret returns the WithoutInitialSecret field if non-nil, zero value otherwise.

### GetWithoutInitialSecretOk

`func (o *GroupServiceAccountRequest) GetWithoutInitialSecretOk() (*bool, bool)`

GetWithoutInitialSecretOk returns a tuple with the WithoutInitialSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutInitialSecret

`func (o *GroupServiceAccountRequest) SetWithoutInitialSecret(v bool)`

SetWithoutInitialSecret sets WithoutInitialSecret field to given value.

### HasWithoutInitialSecret

`func (o *GroupServiceAccountRequest) HasWithoutInitialSecret() bool`

HasWithoutInitialSecret returns a boolean if a field has been set.

### SetWithoutInitialSecretNil

`func (o *GroupServiceAccountRequest) SetWithoutInitialSecretNil()`

SetWithoutInitialSecretNil sets WithoutInitialSecret to an explicit JSON null when marshaled, overriding any value previously set with SetWithoutInitialSecret. Calling SetWithoutInitialSecret again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


