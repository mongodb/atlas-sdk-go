# ApiSearchDeploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the search deployment. | [optional] [readonly] 
**Specs** | Pointer to [**[]ApiSearchDeploymentSpec**](ApiSearchDeploymentSpec.md) | List of settings that configure the Search Nodes for your cluster. | [optional] [readonly] 
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


