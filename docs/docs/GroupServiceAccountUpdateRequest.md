# GroupServiceAccountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Human readable description for the Service Account. | [optional] 
**Name** | Pointer to **string** | Human-readable name for the Service Account. The name is modifiable and does not have to be unique. | [optional] 
**Roles** | Pointer to **[]string** | A list of Project roles associated with the Service Account. | [optional] 

## Methods

### NewGroupServiceAccountUpdateRequest

`func NewGroupServiceAccountUpdateRequest() *GroupServiceAccountUpdateRequest`

NewGroupServiceAccountUpdateRequest instantiates a new GroupServiceAccountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupServiceAccountUpdateRequestWithDefaults

`func NewGroupServiceAccountUpdateRequestWithDefaults() *GroupServiceAccountUpdateRequest`

NewGroupServiceAccountUpdateRequestWithDefaults instantiates a new GroupServiceAccountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupServiceAccountUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupServiceAccountUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupServiceAccountUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupServiceAccountUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetName

`func (o *GroupServiceAccountUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupServiceAccountUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupServiceAccountUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupServiceAccountUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.
### GetRoles

`func (o *GroupServiceAccountUpdateRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupServiceAccountUpdateRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupServiceAccountUpdateRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GroupServiceAccountUpdateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


