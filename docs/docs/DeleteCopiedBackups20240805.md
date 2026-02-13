# DeleteCopiedBackups20240805

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete. | [optional] 
**RegionName** | Pointer to **string** | Target region for the deleted copy setting whose backup copies you want to delete. Please supply the &#39;Atlas Region&#39;. | [optional] 
**ZoneId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Zone Id, do a GET request to Return One Cluster from One Project and consult the &#x60;replicationSpecs&#x60; array. | [optional] 

## Methods

### NewDeleteCopiedBackups20240805

`func NewDeleteCopiedBackups20240805() *DeleteCopiedBackups20240805`

NewDeleteCopiedBackups20240805 instantiates a new DeleteCopiedBackups20240805 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCopiedBackups20240805WithDefaults

`func NewDeleteCopiedBackups20240805WithDefaults() *DeleteCopiedBackups20240805`

NewDeleteCopiedBackups20240805WithDefaults instantiates a new DeleteCopiedBackups20240805 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DeleteCopiedBackups20240805) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DeleteCopiedBackups20240805) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DeleteCopiedBackups20240805) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DeleteCopiedBackups20240805) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetRegionName

`func (o *DeleteCopiedBackups20240805) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *DeleteCopiedBackups20240805) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *DeleteCopiedBackups20240805) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *DeleteCopiedBackups20240805) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetZoneId

`func (o *DeleteCopiedBackups20240805) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DeleteCopiedBackups20240805) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DeleteCopiedBackups20240805) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DeleteCopiedBackups20240805) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


