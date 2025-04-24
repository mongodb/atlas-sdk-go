# GroupServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The Client ID of the Service Account. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date that the Service Account was created on. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] 
**Description** | Pointer to **string** | Human readable description for the Service Account. | [optional] 
**Name** | Pointer to **string** | Human-readable name for the Service Account. | [optional] 
**Roles** | Pointer to **[]string** | A list of Project roles associated with the Service Account. | [optional] 
**Secrets** | Pointer to [**[]ServiceAccountSecret**](ServiceAccountSecret.md) | A list of secrets associated with the specified Service Account. | [optional] 

## Methods

### NewGroupServiceAccount

`func NewGroupServiceAccount() *GroupServiceAccount`

NewGroupServiceAccount instantiates a new GroupServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupServiceAccountWithDefaults

`func NewGroupServiceAccountWithDefaults() *GroupServiceAccount`

NewGroupServiceAccountWithDefaults instantiates a new GroupServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *GroupServiceAccount) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GroupServiceAccount) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GroupServiceAccount) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GroupServiceAccount) HasClientId() bool`

HasClientId returns a boolean if a field has been set.
### GetCreatedAt

`func (o *GroupServiceAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupServiceAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupServiceAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupServiceAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetDescription

`func (o *GroupServiceAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupServiceAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupServiceAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupServiceAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetName

`func (o *GroupServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupServiceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupServiceAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupServiceAccount) HasName() bool`

HasName returns a boolean if a field has been set.
### GetRoles

`func (o *GroupServiceAccount) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupServiceAccount) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupServiceAccount) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GroupServiceAccount) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetSecrets

`func (o *GroupServiceAccount) GetSecrets() []ServiceAccountSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *GroupServiceAccount) GetSecretsOk() (*[]ServiceAccountSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *GroupServiceAccount) SetSecrets(v []ServiceAccountSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *GroupServiceAccount) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


