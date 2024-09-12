# ApiAtlasInvalidPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | A string that defines the permissions for the policy. The syntax used is the Cedar Policy language. | [optional] [readonly] 
**Errors** | Pointer to [**[]ApiAtlasInvalidPolicyErrorDetail**](ApiAtlasInvalidPolicyErrorDetail.md) | List of validation errors. | [optional] [readonly] 

## Methods

### NewApiAtlasInvalidPolicy

`func NewApiAtlasInvalidPolicy() *ApiAtlasInvalidPolicy`

NewApiAtlasInvalidPolicy instantiates a new ApiAtlasInvalidPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasInvalidPolicyWithDefaults

`func NewApiAtlasInvalidPolicyWithDefaults() *ApiAtlasInvalidPolicy`

NewApiAtlasInvalidPolicyWithDefaults instantiates a new ApiAtlasInvalidPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ApiAtlasInvalidPolicy) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ApiAtlasInvalidPolicy) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ApiAtlasInvalidPolicy) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ApiAtlasInvalidPolicy) HasBody() bool`

HasBody returns a boolean if a field has been set.
### GetErrors

`func (o *ApiAtlasInvalidPolicy) GetErrors() []ApiAtlasInvalidPolicyErrorDetail`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiAtlasInvalidPolicy) GetErrorsOk() (*[]ApiAtlasInvalidPolicyErrorDetail, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiAtlasInvalidPolicy) SetErrors(v []ApiAtlasInvalidPolicyErrorDetail)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ApiAtlasInvalidPolicy) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


