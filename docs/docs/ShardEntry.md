# ShardEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** | Human-readable label that identifies the collection to be sharded on the destination cluster. | 
**Database** | **string** | Human-readable label that identifies the database that contains the collection to be sharded on the destination cluster. | 
**ShardCollection** | [**ShardKeys**](ShardKeys.md) |  | 

## Methods

### NewShardEntry

`func NewShardEntry(collection string, database string, shardCollection ShardKeys, ) *ShardEntry`

NewShardEntry instantiates a new ShardEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardEntryWithDefaults

`func NewShardEntryWithDefaults() *ShardEntry`

NewShardEntryWithDefaults instantiates a new ShardEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *ShardEntry) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ShardEntry) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ShardEntry) SetCollection(v string)`

SetCollection sets Collection field to given value.

### GetDatabase

`func (o *ShardEntry) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ShardEntry) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ShardEntry) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### GetShardCollection

`func (o *ShardEntry) GetShardCollection() ShardKeys`

GetShardCollection returns the ShardCollection field if non-nil, zero value otherwise.

### GetShardCollectionOk

`func (o *ShardEntry) GetShardCollectionOk() (*ShardKeys, bool)`

GetShardCollectionOk returns a tuple with the ShardCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardCollection

`func (o *ShardEntry) SetShardCollection(v ShardKeys)`

SetShardCollection sets ShardCollection field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


