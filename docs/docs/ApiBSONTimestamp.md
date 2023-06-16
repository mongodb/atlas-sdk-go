# ApiBSONTimestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | Date and time when the oplog recorded this database operation. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Increment** | Pointer to **int** | Order of the database operation that the oplog recorded at specific date and time. | [optional] [readonly] 

## Methods

### NewApiBSONTimestamp

`func NewApiBSONTimestamp() *ApiBSONTimestamp`

NewApiBSONTimestamp instantiates a new ApiBSONTimestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBSONTimestampWithDefaults

`func NewApiBSONTimestampWithDefaults() *ApiBSONTimestamp`

NewApiBSONTimestampWithDefaults instantiates a new ApiBSONTimestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ApiBSONTimestamp) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ApiBSONTimestamp) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ApiBSONTimestamp) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ApiBSONTimestamp) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetIncrement

`func (o *ApiBSONTimestamp) GetIncrement() int`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *ApiBSONTimestamp) GetIncrementOk() (*int, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *ApiBSONTimestamp) SetIncrement(v int)`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *ApiBSONTimestamp) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


