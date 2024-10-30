# GroupServiceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Human readable description for the Service Account. | 
**Name** | **string** | Human-readable name for the Service Account. The name is modifiable and does not have to be unique. | 
**Roles** | **[]string** | A list of project-level roles for the Service Account. | 
**SecretExpiresAfterHours** | **int** | The expiration time of the new Service Account secret. The expiration is provided in hours. | 

## Methods

### NewGroupServiceAccountRequest

`func NewGroupServiceAccountRequest(description string, name string, roles []string, secretExpiresAfterHours int, ) *GroupServiceAccountRequest`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


