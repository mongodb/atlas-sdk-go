# OrgServiceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Human readable description for the Service Account. | 
**Name** | **string** | Human-readable name for the Service Account. The name is modifiable and does not have to be unique. | 
**Roles** | **[]string** | A list of organization-level roles for the Service Account. | 
**SecretExpiresAfterHours** | **int** | The expiration time of the new Service Account secret. The expiration is provided in hours. | 

## Methods

### NewOrgServiceAccountRequest

`func NewOrgServiceAccountRequest(description string, name string, roles []string, secretExpiresAfterHours int, ) *OrgServiceAccountRequest`

NewOrgServiceAccountRequest instantiates a new OrgServiceAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgServiceAccountRequestWithDefaults

`func NewOrgServiceAccountRequestWithDefaults() *OrgServiceAccountRequest`

NewOrgServiceAccountRequestWithDefaults instantiates a new OrgServiceAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OrgServiceAccountRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrgServiceAccountRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrgServiceAccountRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetName

`func (o *OrgServiceAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgServiceAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgServiceAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### GetRoles

`func (o *OrgServiceAccountRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrgServiceAccountRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrgServiceAccountRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### GetSecretExpiresAfterHours

`func (o *OrgServiceAccountRequest) GetSecretExpiresAfterHours() int`

GetSecretExpiresAfterHours returns the SecretExpiresAfterHours field if non-nil, zero value otherwise.

### GetSecretExpiresAfterHoursOk

`func (o *OrgServiceAccountRequest) GetSecretExpiresAfterHoursOk() (*int, bool)`

GetSecretExpiresAfterHoursOk returns a tuple with the SecretExpiresAfterHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpiresAfterHours

`func (o *OrgServiceAccountRequest) SetSecretExpiresAfterHours(v int)`

SetSecretExpiresAfterHours sets SecretExpiresAfterHours field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


