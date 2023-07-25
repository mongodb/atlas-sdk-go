# ApiCheckpointPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set to which this checkpoint applies. | [optional] [readonly] 
**ShardName** | Pointer to **string** | Human-readable label that identifies the shard to which this checkpoint applies. | [optional] [readonly] 
**TokenDiscovered** | Pointer to **bool** | Flag that indicates whether the token exists. | [optional] [readonly] 
**TokenTimestamp** | Pointer to [**ApiBSONTimestamp**](ApiBSONTimestamp.md) |  | [optional] 
**TypeName** | Pointer to **string** | Human-readable label that identifies the type of host that the part represents. | [optional] [readonly] 

## Methods

### NewApiCheckpointPart

`func NewApiCheckpointPart() *ApiCheckpointPart`

NewApiCheckpointPart instantiates a new ApiCheckpointPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCheckpointPartWithDefaults

`func NewApiCheckpointPartWithDefaults() *ApiCheckpointPart`

NewApiCheckpointPartWithDefaults instantiates a new ApiCheckpointPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicaSetName

`func (o *ApiCheckpointPart) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *ApiCheckpointPart) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *ApiCheckpointPart) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *ApiCheckpointPart) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.
### GetShardName

`func (o *ApiCheckpointPart) GetShardName() string`

GetShardName returns the ShardName field if non-nil, zero value otherwise.

### GetShardNameOk

`func (o *ApiCheckpointPart) GetShardNameOk() (*string, bool)`

GetShardNameOk returns a tuple with the ShardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardName

`func (o *ApiCheckpointPart) SetShardName(v string)`

SetShardName sets ShardName field to given value.

### HasShardName

`func (o *ApiCheckpointPart) HasShardName() bool`

HasShardName returns a boolean if a field has been set.
### GetTokenDiscovered

`func (o *ApiCheckpointPart) GetTokenDiscovered() bool`

GetTokenDiscovered returns the TokenDiscovered field if non-nil, zero value otherwise.

### GetTokenDiscoveredOk

`func (o *ApiCheckpointPart) GetTokenDiscoveredOk() (*bool, bool)`

GetTokenDiscoveredOk returns a tuple with the TokenDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDiscovered

`func (o *ApiCheckpointPart) SetTokenDiscovered(v bool)`

SetTokenDiscovered sets TokenDiscovered field to given value.

### HasTokenDiscovered

`func (o *ApiCheckpointPart) HasTokenDiscovered() bool`

HasTokenDiscovered returns a boolean if a field has been set.
### GetTokenTimestamp

`func (o *ApiCheckpointPart) GetTokenTimestamp() ApiBSONTimestamp`

GetTokenTimestamp returns the TokenTimestamp field if non-nil, zero value otherwise.

### GetTokenTimestampOk

`func (o *ApiCheckpointPart) GetTokenTimestampOk() (*ApiBSONTimestamp, bool)`

GetTokenTimestampOk returns a tuple with the TokenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTimestamp

`func (o *ApiCheckpointPart) SetTokenTimestamp(v ApiBSONTimestamp)`

SetTokenTimestamp sets TokenTimestamp field to given value.

### HasTokenTimestamp

`func (o *ApiCheckpointPart) HasTokenTimestamp() bool`

HasTokenTimestamp returns a boolean if a field has been set.
### GetTypeName

`func (o *ApiCheckpointPart) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ApiCheckpointPart) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ApiCheckpointPart) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *ApiCheckpointPart) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


