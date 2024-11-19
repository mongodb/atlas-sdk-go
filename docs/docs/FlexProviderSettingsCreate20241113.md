# FlexProviderSettingsCreate20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackingProviderName** | **string** | Cloud service provider on which MongoDB Cloud provisioned the flex cluster. | 
**DiskSizeGB** | Pointer to **float64** | Storage capacity available to the flex cluster expressed in gigabytes. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Human-readable label that identifies the provider type. | [optional] [readonly] [default to "FLEX"]
**RegionName** | **string** | Human-readable label that identifies the geographic location of your MongoDB flex cluster. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). | 

## Methods

### NewFlexProviderSettingsCreate20241113

`func NewFlexProviderSettingsCreate20241113(backingProviderName string, regionName string, ) *FlexProviderSettingsCreate20241113`

NewFlexProviderSettingsCreate20241113 instantiates a new FlexProviderSettingsCreate20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexProviderSettingsCreate20241113WithDefaults

`func NewFlexProviderSettingsCreate20241113WithDefaults() *FlexProviderSettingsCreate20241113`

NewFlexProviderSettingsCreate20241113WithDefaults instantiates a new FlexProviderSettingsCreate20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackingProviderName

`func (o *FlexProviderSettingsCreate20241113) GetBackingProviderName() string`

GetBackingProviderName returns the BackingProviderName field if non-nil, zero value otherwise.

### GetBackingProviderNameOk

`func (o *FlexProviderSettingsCreate20241113) GetBackingProviderNameOk() (*string, bool)`

GetBackingProviderNameOk returns a tuple with the BackingProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingProviderName

`func (o *FlexProviderSettingsCreate20241113) SetBackingProviderName(v string)`

SetBackingProviderName sets BackingProviderName field to given value.

### GetDiskSizeGB

`func (o *FlexProviderSettingsCreate20241113) GetDiskSizeGB() float64`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *FlexProviderSettingsCreate20241113) GetDiskSizeGBOk() (*float64, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *FlexProviderSettingsCreate20241113) SetDiskSizeGB(v float64)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *FlexProviderSettingsCreate20241113) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.
### GetProviderName

`func (o *FlexProviderSettingsCreate20241113) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *FlexProviderSettingsCreate20241113) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *FlexProviderSettingsCreate20241113) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *FlexProviderSettingsCreate20241113) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.
### GetRegionName

`func (o *FlexProviderSettingsCreate20241113) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *FlexProviderSettingsCreate20241113) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *FlexProviderSettingsCreate20241113) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


