# GeoSharding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomZoneMapping** | Pointer to **map[string]string** | List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.  This parameter returns an empty object if no custom zones exist. | [optional] [readonly] 
**ManagedNamespaces** | Pointer to [**[]ManagedNamespaces**](ManagedNamespaces.md) | List that contains a namespace for a Global Cluster. MongoDB Cloud manages this cluster. | [optional] [readonly] 

## Methods

### NewGeoSharding

`func NewGeoSharding() *GeoSharding`

NewGeoSharding instantiates a new GeoSharding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoShardingWithDefaults

`func NewGeoShardingWithDefaults() *GeoSharding`

NewGeoShardingWithDefaults instantiates a new GeoSharding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomZoneMapping

`func (o *GeoSharding) GetCustomZoneMapping() map[string]string`

GetCustomZoneMapping returns the CustomZoneMapping field if non-nil, zero value otherwise.

### GetCustomZoneMappingOk

`func (o *GeoSharding) GetCustomZoneMappingOk() (*map[string]string, bool)`

GetCustomZoneMappingOk returns a tuple with the CustomZoneMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomZoneMapping

`func (o *GeoSharding) SetCustomZoneMapping(v map[string]string)`

SetCustomZoneMapping sets CustomZoneMapping field to given value.

### HasCustomZoneMapping

`func (o *GeoSharding) HasCustomZoneMapping() bool`

HasCustomZoneMapping returns a boolean if a field has been set.
### GetManagedNamespaces

`func (o *GeoSharding) GetManagedNamespaces() []ManagedNamespaces`

GetManagedNamespaces returns the ManagedNamespaces field if non-nil, zero value otherwise.

### GetManagedNamespacesOk

`func (o *GeoSharding) GetManagedNamespacesOk() (*[]ManagedNamespaces, bool)`

GetManagedNamespacesOk returns a tuple with the ManagedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNamespaces

`func (o *GeoSharding) SetManagedNamespaces(v []ManagedNamespaces)`

SetManagedNamespaces sets ManagedNamespaces field to given value.

### HasManagedNamespaces

`func (o *GeoSharding) HasManagedNamespaces() bool`

HasManagedNamespaces returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


