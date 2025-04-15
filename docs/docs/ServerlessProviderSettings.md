# ServerlessProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackingProviderName** | **string** | Cloud service provider on which MongoDB Cloud provisioned the serverless instance. | 
**EffectiveDiskSizeGBLimit** | Pointer to **int** | Storage capacity of instance data volumes expressed in gigabytes. This value is not configurable for Serverless or effectively Flex clusters. | [optional] [readonly] 
**EffectiveInstanceSizeName** | Pointer to **string** | Instance size boundary to which your cluster can automatically scale. | [optional] [readonly] 
**EffectiveProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud effectively provisioned the serverless instance. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Human-readable label that identifies the cloud service provider. | [optional] [default to "SERVERLESS"]
**RegionName** | **string** | Human-readable label that identifies the geographic location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). | 

## Methods

### NewServerlessProviderSettings

`func NewServerlessProviderSettings(backingProviderName string, regionName string, ) *ServerlessProviderSettings`

NewServerlessProviderSettings instantiates a new ServerlessProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessProviderSettingsWithDefaults

`func NewServerlessProviderSettingsWithDefaults() *ServerlessProviderSettings`

NewServerlessProviderSettingsWithDefaults instantiates a new ServerlessProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackingProviderName

`func (o *ServerlessProviderSettings) GetBackingProviderName() string`

GetBackingProviderName returns the BackingProviderName field if non-nil, zero value otherwise.

### GetBackingProviderNameOk

`func (o *ServerlessProviderSettings) GetBackingProviderNameOk() (*string, bool)`

GetBackingProviderNameOk returns a tuple with the BackingProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingProviderName

`func (o *ServerlessProviderSettings) SetBackingProviderName(v string)`

SetBackingProviderName sets BackingProviderName field to given value.

### GetEffectiveDiskSizeGBLimit

`func (o *ServerlessProviderSettings) GetEffectiveDiskSizeGBLimit() int`

GetEffectiveDiskSizeGBLimit returns the EffectiveDiskSizeGBLimit field if non-nil, zero value otherwise.

### GetEffectiveDiskSizeGBLimitOk

`func (o *ServerlessProviderSettings) GetEffectiveDiskSizeGBLimitOk() (*int, bool)`

GetEffectiveDiskSizeGBLimitOk returns a tuple with the EffectiveDiskSizeGBLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDiskSizeGBLimit

`func (o *ServerlessProviderSettings) SetEffectiveDiskSizeGBLimit(v int)`

SetEffectiveDiskSizeGBLimit sets EffectiveDiskSizeGBLimit field to given value.

### HasEffectiveDiskSizeGBLimit

`func (o *ServerlessProviderSettings) HasEffectiveDiskSizeGBLimit() bool`

HasEffectiveDiskSizeGBLimit returns a boolean if a field has been set.
### GetEffectiveInstanceSizeName

`func (o *ServerlessProviderSettings) GetEffectiveInstanceSizeName() string`

GetEffectiveInstanceSizeName returns the EffectiveInstanceSizeName field if non-nil, zero value otherwise.

### GetEffectiveInstanceSizeNameOk

`func (o *ServerlessProviderSettings) GetEffectiveInstanceSizeNameOk() (*string, bool)`

GetEffectiveInstanceSizeNameOk returns a tuple with the EffectiveInstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveInstanceSizeName

`func (o *ServerlessProviderSettings) SetEffectiveInstanceSizeName(v string)`

SetEffectiveInstanceSizeName sets EffectiveInstanceSizeName field to given value.

### HasEffectiveInstanceSizeName

`func (o *ServerlessProviderSettings) HasEffectiveInstanceSizeName() bool`

HasEffectiveInstanceSizeName returns a boolean if a field has been set.
### GetEffectiveProviderName

`func (o *ServerlessProviderSettings) GetEffectiveProviderName() string`

GetEffectiveProviderName returns the EffectiveProviderName field if non-nil, zero value otherwise.

### GetEffectiveProviderNameOk

`func (o *ServerlessProviderSettings) GetEffectiveProviderNameOk() (*string, bool)`

GetEffectiveProviderNameOk returns a tuple with the EffectiveProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveProviderName

`func (o *ServerlessProviderSettings) SetEffectiveProviderName(v string)`

SetEffectiveProviderName sets EffectiveProviderName field to given value.

### HasEffectiveProviderName

`func (o *ServerlessProviderSettings) HasEffectiveProviderName() bool`

HasEffectiveProviderName returns a boolean if a field has been set.
### GetProviderName

`func (o *ServerlessProviderSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ServerlessProviderSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ServerlessProviderSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ServerlessProviderSettings) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.
### GetRegionName

`func (o *ServerlessProviderSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ServerlessProviderSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ServerlessProviderSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


