# ClusterCloudProviderInstanceSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableRegions** | Pointer to [**[]AvailableCloudProviderRegion**](AvailableCloudProviderRegion.md) | List of regions that this cloud provider supports for this instance size. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the instance size or cluster tier. | [optional] [readonly] 

## Methods

### NewClusterCloudProviderInstanceSize

`func NewClusterCloudProviderInstanceSize() *ClusterCloudProviderInstanceSize`

NewClusterCloudProviderInstanceSize instantiates a new ClusterCloudProviderInstanceSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCloudProviderInstanceSizeWithDefaults

`func NewClusterCloudProviderInstanceSizeWithDefaults() *ClusterCloudProviderInstanceSize`

NewClusterCloudProviderInstanceSizeWithDefaults instantiates a new ClusterCloudProviderInstanceSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableRegions

`func (o *ClusterCloudProviderInstanceSize) GetAvailableRegions() []AvailableCloudProviderRegion`

GetAvailableRegions returns the AvailableRegions field if non-nil, zero value otherwise.

### GetAvailableRegionsOk

`func (o *ClusterCloudProviderInstanceSize) GetAvailableRegionsOk() (*[]AvailableCloudProviderRegion, bool)`

GetAvailableRegionsOk returns a tuple with the AvailableRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRegions

`func (o *ClusterCloudProviderInstanceSize) SetAvailableRegions(v []AvailableCloudProviderRegion)`

SetAvailableRegions sets AvailableRegions field to given value.

### HasAvailableRegions

`func (o *ClusterCloudProviderInstanceSize) HasAvailableRegions() bool`

HasAvailableRegions returns a boolean if a field has been set.
### GetName

`func (o *ClusterCloudProviderInstanceSize) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCloudProviderInstanceSize) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCloudProviderInstanceSize) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterCloudProviderInstanceSize) HasName() bool`

HasName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


