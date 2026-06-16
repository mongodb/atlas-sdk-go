# ClusterConfigurationValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ClusterConfigurationValidationError**](ClusterConfigurationValidationError.md) | List of validation errors, present only when valid is false. | [optional] 
**Valid** | Pointer to **bool** | Whether the cluster configuration is valid. | [optional] 

## Methods

### NewClusterConfigurationValidationResult

`func NewClusterConfigurationValidationResult() *ClusterConfigurationValidationResult`

NewClusterConfigurationValidationResult instantiates a new ClusterConfigurationValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationValidationResultWithDefaults

`func NewClusterConfigurationValidationResultWithDefaults() *ClusterConfigurationValidationResult`

NewClusterConfigurationValidationResultWithDefaults instantiates a new ClusterConfigurationValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ClusterConfigurationValidationResult) GetErrors() []ClusterConfigurationValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ClusterConfigurationValidationResult) GetErrorsOk() (*[]ClusterConfigurationValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ClusterConfigurationValidationResult) SetErrors(v []ClusterConfigurationValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ClusterConfigurationValidationResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.
### GetValid

`func (o *ClusterConfigurationValidationResult) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ClusterConfigurationValidationResult) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ClusterConfigurationValidationResult) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ClusterConfigurationValidationResult) HasValid() bool`

HasValid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


