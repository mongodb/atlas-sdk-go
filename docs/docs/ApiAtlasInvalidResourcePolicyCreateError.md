# ApiAtlasInvalidResourcePolicyCreateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | Pointer to **string** | Human-readable label that displays the type of an error. | [optional] 
**InvalidPolicies** | Pointer to [**[]ApiAtlasInvalidPolicy**](ApiAtlasInvalidPolicy.md) | List of invalid policies containing details of their validation errors. | [optional] [readonly] 

## Methods

### NewApiAtlasInvalidResourcePolicyCreateError

`func NewApiAtlasInvalidResourcePolicyCreateError() *ApiAtlasInvalidResourcePolicyCreateError`

NewApiAtlasInvalidResourcePolicyCreateError instantiates a new ApiAtlasInvalidResourcePolicyCreateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasInvalidResourcePolicyCreateErrorWithDefaults

`func NewApiAtlasInvalidResourcePolicyCreateErrorWithDefaults() *ApiAtlasInvalidResourcePolicyCreateError`

NewApiAtlasInvalidResourcePolicyCreateErrorWithDefaults instantiates a new ApiAtlasInvalidResourcePolicyCreateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *ApiAtlasInvalidResourcePolicyCreateError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ApiAtlasInvalidResourcePolicyCreateError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ApiAtlasInvalidResourcePolicyCreateError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *ApiAtlasInvalidResourcePolicyCreateError) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.
### GetInvalidPolicies

`func (o *ApiAtlasInvalidResourcePolicyCreateError) GetInvalidPolicies() []ApiAtlasInvalidPolicy`

GetInvalidPolicies returns the InvalidPolicies field if non-nil, zero value otherwise.

### GetInvalidPoliciesOk

`func (o *ApiAtlasInvalidResourcePolicyCreateError) GetInvalidPoliciesOk() (*[]ApiAtlasInvalidPolicy, bool)`

GetInvalidPoliciesOk returns a tuple with the InvalidPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPolicies

`func (o *ApiAtlasInvalidResourcePolicyCreateError) SetInvalidPolicies(v []ApiAtlasInvalidPolicy)`

SetInvalidPolicies sets InvalidPolicies field to given value.

### HasInvalidPolicies

`func (o *ApiAtlasInvalidResourcePolicyCreateError) HasInvalidPolicies() bool`

HasInvalidPolicies returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


