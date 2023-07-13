# ManagedNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to **string** | Human-readable label of the collection to manage for this Global Cluster. | [optional] 
**CustomShardKey** | Pointer to **string** | Database parameter used to divide the *collection* into shards. Global clusters require a compound shard key. This compound shard key combines the location parameter and the user-selected custom key. | [optional] 
**Db** | Pointer to **string** | Human-readable label of the database to manage for this Global Cluster. | [optional] 
**IsCustomShardKeyHashed** | Pointer to **bool** | Flag that indicates whether someone hashed the custom shard key. If this parameter returns &#x60;false&#x60;, this cluster uses ranged sharding. | [optional] [default to false]
**IsShardKeyUnique** | Pointer to **bool** | Flag that indicates whether the underlying index enforces unique values. | [optional] [default to false]
**NumInitialChunks** | Pointer to **int64** | Minimum number of chunks to create initially when sharding an empty collection with a hashed shard key. | [optional] 
**PresplitHashedZones** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should create and distribute initial chunks for an empty or non-existing collection. MongoDB Cloud distributes data based on the defined zones and zone ranges for the collection. | [optional] [default to false]

## Methods

### NewManagedNamespace

`func NewManagedNamespace() *ManagedNamespace`

NewManagedNamespace instantiates a new ManagedNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedNamespaceWithDefaults

`func NewManagedNamespaceWithDefaults() *ManagedNamespace`

NewManagedNamespaceWithDefaults instantiates a new ManagedNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *ManagedNamespace) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ManagedNamespace) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ManagedNamespace) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ManagedNamespace) HasCollection() bool`

HasCollection returns a boolean if a field has been set.
### GetCustomShardKey

`func (o *ManagedNamespace) GetCustomShardKey() string`

GetCustomShardKey returns the CustomShardKey field if non-nil, zero value otherwise.

### GetCustomShardKeyOk

`func (o *ManagedNamespace) GetCustomShardKeyOk() (*string, bool)`

GetCustomShardKeyOk returns a tuple with the CustomShardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShardKey

`func (o *ManagedNamespace) SetCustomShardKey(v string)`

SetCustomShardKey sets CustomShardKey field to given value.

### HasCustomShardKey

`func (o *ManagedNamespace) HasCustomShardKey() bool`

HasCustomShardKey returns a boolean if a field has been set.
### GetDb

`func (o *ManagedNamespace) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *ManagedNamespace) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *ManagedNamespace) SetDb(v string)`

SetDb sets Db field to given value.

### HasDb

`func (o *ManagedNamespace) HasDb() bool`

HasDb returns a boolean if a field has been set.
### GetIsCustomShardKeyHashed

`func (o *ManagedNamespace) GetIsCustomShardKeyHashed() bool`

GetIsCustomShardKeyHashed returns the IsCustomShardKeyHashed field if non-nil, zero value otherwise.

### GetIsCustomShardKeyHashedOk

`func (o *ManagedNamespace) GetIsCustomShardKeyHashedOk() (*bool, bool)`

GetIsCustomShardKeyHashedOk returns a tuple with the IsCustomShardKeyHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomShardKeyHashed

`func (o *ManagedNamespace) SetIsCustomShardKeyHashed(v bool)`

SetIsCustomShardKeyHashed sets IsCustomShardKeyHashed field to given value.

### HasIsCustomShardKeyHashed

`func (o *ManagedNamespace) HasIsCustomShardKeyHashed() bool`

HasIsCustomShardKeyHashed returns a boolean if a field has been set.
### GetIsShardKeyUnique

`func (o *ManagedNamespace) GetIsShardKeyUnique() bool`

GetIsShardKeyUnique returns the IsShardKeyUnique field if non-nil, zero value otherwise.

### GetIsShardKeyUniqueOk

`func (o *ManagedNamespace) GetIsShardKeyUniqueOk() (*bool, bool)`

GetIsShardKeyUniqueOk returns a tuple with the IsShardKeyUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShardKeyUnique

`func (o *ManagedNamespace) SetIsShardKeyUnique(v bool)`

SetIsShardKeyUnique sets IsShardKeyUnique field to given value.

### HasIsShardKeyUnique

`func (o *ManagedNamespace) HasIsShardKeyUnique() bool`

HasIsShardKeyUnique returns a boolean if a field has been set.
### GetNumInitialChunks

`func (o *ManagedNamespace) GetNumInitialChunks() int64`

GetNumInitialChunks returns the NumInitialChunks field if non-nil, zero value otherwise.

### GetNumInitialChunksOk

`func (o *ManagedNamespace) GetNumInitialChunksOk() (*int64, bool)`

GetNumInitialChunksOk returns a tuple with the NumInitialChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInitialChunks

`func (o *ManagedNamespace) SetNumInitialChunks(v int64)`

SetNumInitialChunks sets NumInitialChunks field to given value.

### HasNumInitialChunks

`func (o *ManagedNamespace) HasNumInitialChunks() bool`

HasNumInitialChunks returns a boolean if a field has been set.
### GetPresplitHashedZones

`func (o *ManagedNamespace) GetPresplitHashedZones() bool`

GetPresplitHashedZones returns the PresplitHashedZones field if non-nil, zero value otherwise.

### GetPresplitHashedZonesOk

`func (o *ManagedNamespace) GetPresplitHashedZonesOk() (*bool, bool)`

GetPresplitHashedZonesOk returns a tuple with the PresplitHashedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresplitHashedZones

`func (o *ManagedNamespace) SetPresplitHashedZones(v bool)`

SetPresplitHashedZones sets PresplitHashedZones field to given value.

### HasPresplitHashedZones

`func (o *ManagedNamespace) HasPresplitHashedZones() bool`

HasPresplitHashedZones returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


