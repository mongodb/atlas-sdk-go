# UpdateAtlasProjectApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Purpose or explanation provided when someone creates this project API key. | [optional] 
**Roles** | Pointer to **[]string** | List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project. | [optional] 

## Methods

### NewUpdateAtlasProjectApiKey

`func NewUpdateAtlasProjectApiKey() *UpdateAtlasProjectApiKey`

NewUpdateAtlasProjectApiKey instantiates a new UpdateAtlasProjectApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAtlasProjectApiKeyWithDefaults

`func NewUpdateAtlasProjectApiKeyWithDefaults() *UpdateAtlasProjectApiKey`

NewUpdateAtlasProjectApiKeyWithDefaults instantiates a new UpdateAtlasProjectApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *UpdateAtlasProjectApiKey) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *UpdateAtlasProjectApiKey) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *UpdateAtlasProjectApiKey) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *UpdateAtlasProjectApiKey) HasDesc() bool`

HasDesc returns a boolean if a field has been set.
### GetRoles

`func (o *UpdateAtlasProjectApiKey) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateAtlasProjectApiKey) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateAtlasProjectApiKey) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateAtlasProjectApiKey) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


