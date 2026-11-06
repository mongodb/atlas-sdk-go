# ShardOperations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Average** | **float64** | Average operations per second across the shard&#39;s nodes over the last hour. | [readonly] 
**Maximum** | **float64** | Highest operations per second that any one of the shard&#39;s nodes reached over the last hour. | [readonly] 

## Methods

### NewShardOperations

`func NewShardOperations(average float64, maximum float64, ) *ShardOperations`

NewShardOperations instantiates a new ShardOperations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardOperationsWithDefaults

`func NewShardOperationsWithDefaults() *ShardOperations`

NewShardOperationsWithDefaults instantiates a new ShardOperations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverage

`func (o *ShardOperations) GetAverage() float64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *ShardOperations) GetAverageOk() (*float64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *ShardOperations) SetAverage(v float64)`

SetAverage sets Average field to given value.

### GetMaximum

`func (o *ShardOperations) GetMaximum() float64`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *ShardOperations) GetMaximumOk() (*float64, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *ShardOperations) SetMaximum(v float64)`

SetMaximum sets Maximum field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


