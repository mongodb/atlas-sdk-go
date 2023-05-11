# CheckpointPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set to which this checkpoint applies. | [optional] [readonly] 
**ShardName** | Pointer to **string** | Human-readable label that identifies the shard to which this checkpoint applies. | [optional] [readonly] 
**TokenDiscovered** | Pointer to **bool** | Flag that indicates whether the token exists. | [optional] [readonly] 
**TokenTimestamp** | Pointer to [**BSONTimestamp**](BSONTimestamp.md) |  | [optional] 
**TypeName** | Pointer to **string** | Human-readable label that identifies the type of host that the part represents. | [optional] [readonly] 

## Methods

### NewCheckpointPart

`func NewCheckpointPart() *CheckpointPart`

NewCheckpointPart instantiates a new CheckpointPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointPartWithDefaults

`func NewCheckpointPartWithDefaults() *CheckpointPart`

NewCheckpointPartWithDefaults instantiates a new CheckpointPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicaSetName

`func (o *CheckpointPart) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *CheckpointPart) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *CheckpointPart) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *CheckpointPart) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

### GetShardName

`func (o *CheckpointPart) GetShardName() string`

GetShardName returns the ShardName field if non-nil, zero value otherwise.

### GetShardNameOk

`func (o *CheckpointPart) GetShardNameOk() (*string, bool)`

GetShardNameOk returns a tuple with the ShardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardName

`func (o *CheckpointPart) SetShardName(v string)`

SetShardName sets ShardName field to given value.

### HasShardName

`func (o *CheckpointPart) HasShardName() bool`

HasShardName returns a boolean if a field has been set.

### GetTokenDiscovered

`func (o *CheckpointPart) GetTokenDiscovered() bool`

GetTokenDiscovered returns the TokenDiscovered field if non-nil, zero value otherwise.

### GetTokenDiscoveredOk

`func (o *CheckpointPart) GetTokenDiscoveredOk() (*bool, bool)`

GetTokenDiscoveredOk returns a tuple with the TokenDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDiscovered

`func (o *CheckpointPart) SetTokenDiscovered(v bool)`

SetTokenDiscovered sets TokenDiscovered field to given value.

### HasTokenDiscovered

`func (o *CheckpointPart) HasTokenDiscovered() bool`

HasTokenDiscovered returns a boolean if a field has been set.

### GetTokenTimestamp

`func (o *CheckpointPart) GetTokenTimestamp() BSONTimestamp`

GetTokenTimestamp returns the TokenTimestamp field if non-nil, zero value otherwise.

### GetTokenTimestampOk

`func (o *CheckpointPart) GetTokenTimestampOk() (*BSONTimestamp, bool)`

GetTokenTimestampOk returns a tuple with the TokenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTimestamp

`func (o *CheckpointPart) SetTokenTimestamp(v BSONTimestamp)`

SetTokenTimestamp sets TokenTimestamp field to given value.

### HasTokenTimestamp

`func (o *CheckpointPart) HasTokenTimestamp() bool`

HasTokenTimestamp returns a boolean if a field has been set.

### GetTypeName

`func (o *CheckpointPart) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *CheckpointPart) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *CheckpointPart) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *CheckpointPart) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


