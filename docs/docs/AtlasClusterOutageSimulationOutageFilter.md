# AtlasClusterOutageSimulationOutageFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider of the region that undergoes the outage simulation. | [optional] 
**RegionName** | Pointer to **string** | The name of the region to undergo an outage simulation. | [optional] 
**Type** | Pointer to **string** | The type of cluster outage to simulate. &#x60;REGION&#x60; simulates a cluster outage for a region. | [optional] 

## Methods

### NewAtlasClusterOutageSimulationOutageFilter

`func NewAtlasClusterOutageSimulationOutageFilter() *AtlasClusterOutageSimulationOutageFilter`

NewAtlasClusterOutageSimulationOutageFilter instantiates a new AtlasClusterOutageSimulationOutageFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasClusterOutageSimulationOutageFilterWithDefaults

`func NewAtlasClusterOutageSimulationOutageFilterWithDefaults() *AtlasClusterOutageSimulationOutageFilter`

NewAtlasClusterOutageSimulationOutageFilterWithDefaults instantiates a new AtlasClusterOutageSimulationOutageFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *AtlasClusterOutageSimulationOutageFilter) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AtlasClusterOutageSimulationOutageFilter) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AtlasClusterOutageSimulationOutageFilter) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *AtlasClusterOutageSimulationOutageFilter) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetRegionName

`func (o *AtlasClusterOutageSimulationOutageFilter) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AtlasClusterOutageSimulationOutageFilter) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AtlasClusterOutageSimulationOutageFilter) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AtlasClusterOutageSimulationOutageFilter) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetType

`func (o *AtlasClusterOutageSimulationOutageFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AtlasClusterOutageSimulationOutageFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AtlasClusterOutageSimulationOutageFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AtlasClusterOutageSimulationOutageFilter) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


