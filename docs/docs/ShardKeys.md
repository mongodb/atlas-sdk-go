# ShardKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **[]map[string]interface{}** | List of fields to use for the shard key. | [optional] 

## Methods

### NewShardKeys

`func NewShardKeys() *ShardKeys`

NewShardKeys instantiates a new ShardKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardKeysWithDefaults

`func NewShardKeysWithDefaults() *ShardKeys`

NewShardKeysWithDefaults instantiates a new ShardKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ShardKeys) GetKey() []map[string]interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ShardKeys) GetKeyOk() (*[]map[string]interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ShardKeys) SetKey(v []map[string]interface{})`

SetKey sets Key field to given value.

### HasKey

`func (o *ShardKeys) HasKey() bool`

HasKey returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


