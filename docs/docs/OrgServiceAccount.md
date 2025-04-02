# OrgServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The Client ID of the Service Account. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date that the Service Account was created on. | [optional] 
**Description** | Pointer to **string** | Human readable description for the Service Account. | [optional] 
**Name** | Pointer to **string** | Human-readable name for the Service Account. | [optional] 
**Roles** | Pointer to **[]string** | A list of Organization roles associated with the Service Account. | [optional] 
**Secrets** | Pointer to [**[]ServiceAccountSecret**](ServiceAccountSecret.md) | A list of secrets associated with the specified Service Account. | [optional] 

## Methods

### NewOrgServiceAccount

`func NewOrgServiceAccount() *OrgServiceAccount`

NewOrgServiceAccount instantiates a new OrgServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgServiceAccountWithDefaults

`func NewOrgServiceAccountWithDefaults() *OrgServiceAccount`

NewOrgServiceAccountWithDefaults instantiates a new OrgServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OrgServiceAccount) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OrgServiceAccount) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OrgServiceAccount) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OrgServiceAccount) HasClientId() bool`

HasClientId returns a boolean if a field has been set.
### GetCreatedAt

`func (o *OrgServiceAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgServiceAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgServiceAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrgServiceAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetDescription

`func (o *OrgServiceAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrgServiceAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrgServiceAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrgServiceAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetName

`func (o *OrgServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgServiceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgServiceAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgServiceAccount) HasName() bool`

HasName returns a boolean if a field has been set.
### GetRoles

`func (o *OrgServiceAccount) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrgServiceAccount) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrgServiceAccount) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrgServiceAccount) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetSecrets

`func (o *OrgServiceAccount) GetSecrets() []ServiceAccountSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *OrgServiceAccount) GetSecretsOk() (*[]ServiceAccountSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *OrgServiceAccount) SetSecrets(v []ServiceAccountSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *OrgServiceAccount) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


