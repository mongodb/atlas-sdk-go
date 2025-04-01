# ClusterDescriptionAutoScalingModeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScalingMode** | Pointer to **string** | Describes whether cluster nodes scale together across all shards or independently. | [optional] 

## Methods

### NewClusterDescriptionAutoScalingModeConfiguration

`func NewClusterDescriptionAutoScalingModeConfiguration() *ClusterDescriptionAutoScalingModeConfiguration`

NewClusterDescriptionAutoScalingModeConfiguration instantiates a new ClusterDescriptionAutoScalingModeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDescriptionAutoScalingModeConfigurationWithDefaults

`func NewClusterDescriptionAutoScalingModeConfigurationWithDefaults() *ClusterDescriptionAutoScalingModeConfiguration`

NewClusterDescriptionAutoScalingModeConfigurationWithDefaults instantiates a new ClusterDescriptionAutoScalingModeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScalingMode

`func (o *ClusterDescriptionAutoScalingModeConfiguration) GetAutoScalingMode() string`

GetAutoScalingMode returns the AutoScalingMode field if non-nil, zero value otherwise.

### GetAutoScalingModeOk

`func (o *ClusterDescriptionAutoScalingModeConfiguration) GetAutoScalingModeOk() (*string, bool)`

GetAutoScalingModeOk returns a tuple with the AutoScalingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingMode

`func (o *ClusterDescriptionAutoScalingModeConfiguration) SetAutoScalingMode(v string)`

SetAutoScalingMode sets AutoScalingMode field to given value.

### HasAutoScalingMode

`func (o *ClusterDescriptionAutoScalingModeConfiguration) HasAutoScalingMode() bool`

HasAutoScalingMode returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


