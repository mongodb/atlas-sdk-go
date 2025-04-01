# ApiAtlasResourcePolicyEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the atlas resource policy. | [optional] 
**Name** | Pointer to **string** | Human-readable label that describes the atlas resource policy. | [optional] 
**Policies** | Pointer to [**[]ApiAtlasPolicyCreate**](ApiAtlasPolicyCreate.md) | List of policies that make up the atlas resource policy. | [optional] 

## Methods

### NewApiAtlasResourcePolicyEdit

`func NewApiAtlasResourcePolicyEdit() *ApiAtlasResourcePolicyEdit`

NewApiAtlasResourcePolicyEdit instantiates a new ApiAtlasResourcePolicyEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasResourcePolicyEditWithDefaults

`func NewApiAtlasResourcePolicyEditWithDefaults() *ApiAtlasResourcePolicyEdit`

NewApiAtlasResourcePolicyEditWithDefaults instantiates a new ApiAtlasResourcePolicyEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiAtlasResourcePolicyEdit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiAtlasResourcePolicyEdit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiAtlasResourcePolicyEdit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiAtlasResourcePolicyEdit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetName

`func (o *ApiAtlasResourcePolicyEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasResourcePolicyEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasResourcePolicyEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasResourcePolicyEdit) HasName() bool`

HasName returns a boolean if a field has been set.
### GetPolicies

`func (o *ApiAtlasResourcePolicyEdit) GetPolicies() []ApiAtlasPolicyCreate`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ApiAtlasResourcePolicyEdit) GetPoliciesOk() (*[]ApiAtlasPolicyCreate, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ApiAtlasResourcePolicyEdit) SetPolicies(v []ApiAtlasPolicyCreate)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ApiAtlasResourcePolicyEdit) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


