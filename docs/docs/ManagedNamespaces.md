# ManagedNamespaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** | Human-readable label of the collection to manage for this Global Cluster. | 
**CustomShardKey** | **string** | Database parameter used to divide the *collection* into shards. Global clusters require a compound shard key. This compound shard key combines the location parameter and the user-selected custom key. | 
**Db** | **string** | Human-readable label of the database to manage for this Global Cluster. | 
**IsCustomShardKeyHashed** | Pointer to **bool** | Flag that indicates whether someone hashed the custom shard key for the specified collection. If you set this value to &#x60;false&#x60;, MongoDB Cloud uses ranged sharding. | [optional] [default to false]
**IsShardKeyUnique** | Pointer to **bool** | Flag that indicates whether someone [hashed](https://www.mongodb.com/docs/manual/reference/method/sh.shardCollection/#hashed-shard-keys) the custom shard key. If this parameter returns &#x60;false&#x60;, this cluster uses [ranged sharding](https://www.mongodb.com/docs/manual/core/ranged-sharding/). | [optional] [default to false]
**NumInitialChunks** | Pointer to **int64** | Minimum number of chunks to create initially when sharding an empty collection with a [hashed shard key](https://www.mongodb.com/docs/manual/core/hashed-sharding/). | [optional] 
**PresplitHashedZones** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should create and distribute initial chunks for an empty or non-existing collection. MongoDB Cloud distributes data based on the defined zones and zone ranges for the collection. | [optional] [default to false]

## Methods

### NewManagedNamespaces

`func NewManagedNamespaces(collection string, customShardKey string, db string, ) *ManagedNamespaces`

NewManagedNamespaces instantiates a new ManagedNamespaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedNamespacesWithDefaults

`func NewManagedNamespacesWithDefaults() *ManagedNamespaces`

NewManagedNamespacesWithDefaults instantiates a new ManagedNamespaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *ManagedNamespaces) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ManagedNamespaces) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ManagedNamespaces) SetCollection(v string)`

SetCollection sets Collection field to given value.

### GetCustomShardKey

`func (o *ManagedNamespaces) GetCustomShardKey() string`

GetCustomShardKey returns the CustomShardKey field if non-nil, zero value otherwise.

### GetCustomShardKeyOk

`func (o *ManagedNamespaces) GetCustomShardKeyOk() (*string, bool)`

GetCustomShardKeyOk returns a tuple with the CustomShardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShardKey

`func (o *ManagedNamespaces) SetCustomShardKey(v string)`

SetCustomShardKey sets CustomShardKey field to given value.

### GetDb

`func (o *ManagedNamespaces) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *ManagedNamespaces) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *ManagedNamespaces) SetDb(v string)`

SetDb sets Db field to given value.

### GetIsCustomShardKeyHashed

`func (o *ManagedNamespaces) GetIsCustomShardKeyHashed() bool`

GetIsCustomShardKeyHashed returns the IsCustomShardKeyHashed field if non-nil, zero value otherwise.

### GetIsCustomShardKeyHashedOk

`func (o *ManagedNamespaces) GetIsCustomShardKeyHashedOk() (*bool, bool)`

GetIsCustomShardKeyHashedOk returns a tuple with the IsCustomShardKeyHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomShardKeyHashed

`func (o *ManagedNamespaces) SetIsCustomShardKeyHashed(v bool)`

SetIsCustomShardKeyHashed sets IsCustomShardKeyHashed field to given value.

### HasIsCustomShardKeyHashed

`func (o *ManagedNamespaces) HasIsCustomShardKeyHashed() bool`

HasIsCustomShardKeyHashed returns a boolean if a field has been set.
### GetIsShardKeyUnique

`func (o *ManagedNamespaces) GetIsShardKeyUnique() bool`

GetIsShardKeyUnique returns the IsShardKeyUnique field if non-nil, zero value otherwise.

### GetIsShardKeyUniqueOk

`func (o *ManagedNamespaces) GetIsShardKeyUniqueOk() (*bool, bool)`

GetIsShardKeyUniqueOk returns a tuple with the IsShardKeyUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShardKeyUnique

`func (o *ManagedNamespaces) SetIsShardKeyUnique(v bool)`

SetIsShardKeyUnique sets IsShardKeyUnique field to given value.

### HasIsShardKeyUnique

`func (o *ManagedNamespaces) HasIsShardKeyUnique() bool`

HasIsShardKeyUnique returns a boolean if a field has been set.
### GetNumInitialChunks

`func (o *ManagedNamespaces) GetNumInitialChunks() int64`

GetNumInitialChunks returns the NumInitialChunks field if non-nil, zero value otherwise.

### GetNumInitialChunksOk

`func (o *ManagedNamespaces) GetNumInitialChunksOk() (*int64, bool)`

GetNumInitialChunksOk returns a tuple with the NumInitialChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInitialChunks

`func (o *ManagedNamespaces) SetNumInitialChunks(v int64)`

SetNumInitialChunks sets NumInitialChunks field to given value.

### HasNumInitialChunks

`func (o *ManagedNamespaces) HasNumInitialChunks() bool`

HasNumInitialChunks returns a boolean if a field has been set.
### GetPresplitHashedZones

`func (o *ManagedNamespaces) GetPresplitHashedZones() bool`

GetPresplitHashedZones returns the PresplitHashedZones field if non-nil, zero value otherwise.

### GetPresplitHashedZonesOk

`func (o *ManagedNamespaces) GetPresplitHashedZonesOk() (*bool, bool)`

GetPresplitHashedZonesOk returns a tuple with the PresplitHashedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresplitHashedZones

`func (o *ManagedNamespaces) SetPresplitHashedZones(v bool)`

SetPresplitHashedZones sets PresplitHashedZones field to given value.

### HasPresplitHashedZones

`func (o *ManagedNamespaces) HasPresplitHashedZones() bool`

HasPresplitHashedZones returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


