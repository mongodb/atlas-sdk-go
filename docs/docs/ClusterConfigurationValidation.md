# ClusterConfigurationValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedConfiguration** | Pointer to [**ClusterDescriptionProcessArgs20240805**](ClusterDescriptionProcessArgs20240805.md) |  | [optional] 
**ClusterDescription** | [**ClusterDescription20240805**](ClusterDescription20240805.md) |  | 
**ClusterName** | Pointer to **string** | Name of the existing cluster to validate against when editing is true. | [optional] 
**Editing** | **bool** | When true, validates the configuration as an update to an existing cluster. When false, validates as a new cluster create. | 
**SearchDeploymentSpec** | Pointer to [**ApiSearchDeploymentSpec**](ApiSearchDeploymentSpec.md) |  | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | Optional list of resource tags. | [optional] 

## Methods

### NewClusterConfigurationValidation

`func NewClusterConfigurationValidation(clusterDescription ClusterDescription20240805, editing bool, ) *ClusterConfigurationValidation`

NewClusterConfigurationValidation instantiates a new ClusterConfigurationValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationValidationWithDefaults

`func NewClusterConfigurationValidationWithDefaults() *ClusterConfigurationValidation`

NewClusterConfigurationValidationWithDefaults instantiates a new ClusterConfigurationValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedConfiguration

`func (o *ClusterConfigurationValidation) GetAdvancedConfiguration() ClusterDescriptionProcessArgs20240805`

GetAdvancedConfiguration returns the AdvancedConfiguration field if non-nil, zero value otherwise.

### GetAdvancedConfigurationOk

`func (o *ClusterConfigurationValidation) GetAdvancedConfigurationOk() (*ClusterDescriptionProcessArgs20240805, bool)`

GetAdvancedConfigurationOk returns a tuple with the AdvancedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfiguration

`func (o *ClusterConfigurationValidation) SetAdvancedConfiguration(v ClusterDescriptionProcessArgs20240805)`

SetAdvancedConfiguration sets AdvancedConfiguration field to given value.

### HasAdvancedConfiguration

`func (o *ClusterConfigurationValidation) HasAdvancedConfiguration() bool`

HasAdvancedConfiguration returns a boolean if a field has been set.

### SetAdvancedConfigurationNil

`func (o *ClusterConfigurationValidation) SetAdvancedConfigurationNil()`

SetAdvancedConfigurationNil sets AdvancedConfiguration to an explicit JSON null when marshaled, overriding any value previously set with SetAdvancedConfiguration. Calling SetAdvancedConfiguration again clears the null override.

### GetClusterDescription

`func (o *ClusterConfigurationValidation) GetClusterDescription() ClusterDescription20240805`

GetClusterDescription returns the ClusterDescription field if non-nil, zero value otherwise.

### GetClusterDescriptionOk

`func (o *ClusterConfigurationValidation) GetClusterDescriptionOk() (*ClusterDescription20240805, bool)`

GetClusterDescriptionOk returns a tuple with the ClusterDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDescription

`func (o *ClusterConfigurationValidation) SetClusterDescription(v ClusterDescription20240805)`

SetClusterDescription sets ClusterDescription field to given value.

### GetClusterName

`func (o *ClusterConfigurationValidation) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterConfigurationValidation) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterConfigurationValidation) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ClusterConfigurationValidation) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *ClusterConfigurationValidation) SetClusterNameNil()`

SetClusterNameNil sets ClusterName to an explicit JSON null when marshaled, overriding any value previously set with SetClusterName. Calling SetClusterName again clears the null override.

### GetEditing

`func (o *ClusterConfigurationValidation) GetEditing() bool`

GetEditing returns the Editing field if non-nil, zero value otherwise.

### GetEditingOk

`func (o *ClusterConfigurationValidation) GetEditingOk() (*bool, bool)`

GetEditingOk returns a tuple with the Editing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditing

`func (o *ClusterConfigurationValidation) SetEditing(v bool)`

SetEditing sets Editing field to given value.

### GetSearchDeploymentSpec

`func (o *ClusterConfigurationValidation) GetSearchDeploymentSpec() ApiSearchDeploymentSpec`

GetSearchDeploymentSpec returns the SearchDeploymentSpec field if non-nil, zero value otherwise.

### GetSearchDeploymentSpecOk

`func (o *ClusterConfigurationValidation) GetSearchDeploymentSpecOk() (*ApiSearchDeploymentSpec, bool)`

GetSearchDeploymentSpecOk returns a tuple with the SearchDeploymentSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDeploymentSpec

`func (o *ClusterConfigurationValidation) SetSearchDeploymentSpec(v ApiSearchDeploymentSpec)`

SetSearchDeploymentSpec sets SearchDeploymentSpec field to given value.

### HasSearchDeploymentSpec

`func (o *ClusterConfigurationValidation) HasSearchDeploymentSpec() bool`

HasSearchDeploymentSpec returns a boolean if a field has been set.

### SetSearchDeploymentSpecNil

`func (o *ClusterConfigurationValidation) SetSearchDeploymentSpecNil()`

SetSearchDeploymentSpecNil sets SearchDeploymentSpec to an explicit JSON null when marshaled, overriding any value previously set with SetSearchDeploymentSpec. Calling SetSearchDeploymentSpec again clears the null override.

### GetTags

`func (o *ClusterConfigurationValidation) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterConfigurationValidation) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterConfigurationValidation) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterConfigurationValidation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ClusterConfigurationValidation) SetTagsNil()`

SetTagsNil sets Tags to an explicit JSON null when marshaled, overriding any value previously set with SetTags. Calling SetTags again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


