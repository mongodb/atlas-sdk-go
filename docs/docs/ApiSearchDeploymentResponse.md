# ApiSearchDeploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveSpecs** | Pointer to [**[]ApiSearchDeploymentEffectiveSpec**](ApiSearchDeploymentEffectiveSpec.md) | List of settings that configure the Search Nodes for your cluster, with per-region detail including the region name and cloud provider. | [optional] [readonly] 
**EncryptionAtRestProvider** | Pointer to **string** | Cloud service provider that manages your customer keys to provide an additional layer of Encryption At Rest for the cluster. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the search deployment. | [optional] [readonly] 
**Specs** | Pointer to [**[]ApiSearchDeploymentSpec**](ApiSearchDeploymentSpec.md) | Deprecated. &#x60;specs&#x60; will be removed in a future release. We strongly recommend that you use &#x60;effectiveSpecs&#x60; instead. | [optional] [readonly] 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of this search deployment. | [optional] [readonly] 

## Methods

### NewApiSearchDeploymentResponse

`func NewApiSearchDeploymentResponse() *ApiSearchDeploymentResponse`

NewApiSearchDeploymentResponse instantiates a new ApiSearchDeploymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchDeploymentResponseWithDefaults

`func NewApiSearchDeploymentResponseWithDefaults() *ApiSearchDeploymentResponse`

NewApiSearchDeploymentResponseWithDefaults instantiates a new ApiSearchDeploymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveSpecs

`func (o *ApiSearchDeploymentResponse) GetEffectiveSpecs() []ApiSearchDeploymentEffectiveSpec`

GetEffectiveSpecs returns the EffectiveSpecs field if non-nil, zero value otherwise.

### GetEffectiveSpecsOk

`func (o *ApiSearchDeploymentResponse) GetEffectiveSpecsOk() (*[]ApiSearchDeploymentEffectiveSpec, bool)`

GetEffectiveSpecsOk returns a tuple with the EffectiveSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSpecs

`func (o *ApiSearchDeploymentResponse) SetEffectiveSpecs(v []ApiSearchDeploymentEffectiveSpec)`

SetEffectiveSpecs sets EffectiveSpecs field to given value.

### HasEffectiveSpecs

`func (o *ApiSearchDeploymentResponse) HasEffectiveSpecs() bool`

HasEffectiveSpecs returns a boolean if a field has been set.

### SetEffectiveSpecsNil

`func (o *ApiSearchDeploymentResponse) SetEffectiveSpecsNil()`

SetEffectiveSpecsNil sets EffectiveSpecs to an explicit JSON null when marshaled, overriding any value previously set with SetEffectiveSpecs. Calling SetEffectiveSpecs again clears the null override.

### GetEncryptionAtRestProvider

`func (o *ApiSearchDeploymentResponse) GetEncryptionAtRestProvider() string`

GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field if non-nil, zero value otherwise.

### GetEncryptionAtRestProviderOk

`func (o *ApiSearchDeploymentResponse) GetEncryptionAtRestProviderOk() (*string, bool)`

GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestProvider

`func (o *ApiSearchDeploymentResponse) SetEncryptionAtRestProvider(v string)`

SetEncryptionAtRestProvider sets EncryptionAtRestProvider field to given value.

### HasEncryptionAtRestProvider

`func (o *ApiSearchDeploymentResponse) HasEncryptionAtRestProvider() bool`

HasEncryptionAtRestProvider returns a boolean if a field has been set.

### SetEncryptionAtRestProviderNil

`func (o *ApiSearchDeploymentResponse) SetEncryptionAtRestProviderNil()`

SetEncryptionAtRestProviderNil sets EncryptionAtRestProvider to an explicit JSON null when marshaled, overriding any value previously set with SetEncryptionAtRestProvider. Calling SetEncryptionAtRestProvider again clears the null override.

### GetGroupId

`func (o *ApiSearchDeploymentResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiSearchDeploymentResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiSearchDeploymentResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiSearchDeploymentResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *ApiSearchDeploymentResponse) SetGroupIdNil()`

SetGroupIdNil sets GroupId to an explicit JSON null when marshaled, overriding any value previously set with SetGroupId. Calling SetGroupId again clears the null override.

### GetId

`func (o *ApiSearchDeploymentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiSearchDeploymentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiSearchDeploymentResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiSearchDeploymentResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ApiSearchDeploymentResponse) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetSpecs

`func (o *ApiSearchDeploymentResponse) GetSpecs() []ApiSearchDeploymentSpec`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *ApiSearchDeploymentResponse) GetSpecsOk() (*[]ApiSearchDeploymentSpec, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *ApiSearchDeploymentResponse) SetSpecs(v []ApiSearchDeploymentSpec)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *ApiSearchDeploymentResponse) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### SetSpecsNil

`func (o *ApiSearchDeploymentResponse) SetSpecsNil()`

SetSpecsNil sets Specs to an explicit JSON null when marshaled, overriding any value previously set with SetSpecs. Calling SetSpecs again clears the null override.

### GetStateName

`func (o *ApiSearchDeploymentResponse) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ApiSearchDeploymentResponse) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ApiSearchDeploymentResponse) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *ApiSearchDeploymentResponse) HasStateName() bool`

HasStateName returns a boolean if a field has been set.

### SetStateNameNil

`func (o *ApiSearchDeploymentResponse) SetStateNameNil()`

SetStateNameNil sets StateName to an explicit JSON null when marshaled, overriding any value previously set with SetStateName. Calling SetStateName again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


