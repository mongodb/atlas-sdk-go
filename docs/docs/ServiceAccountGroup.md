# ServiceAccountGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies your project. **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | [optional] [readonly] 

## Methods

### NewServiceAccountGroup

`func NewServiceAccountGroup() *ServiceAccountGroup`

NewServiceAccountGroup instantiates a new ServiceAccountGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountGroupWithDefaults

`func NewServiceAccountGroupWithDefaults() *ServiceAccountGroup`

NewServiceAccountGroupWithDefaults instantiates a new ServiceAccountGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ServiceAccountGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServiceAccountGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServiceAccountGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ServiceAccountGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


