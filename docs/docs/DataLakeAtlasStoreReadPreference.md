# DataLakeAtlasStoreReadPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxStalenessSeconds** | Pointer to **int** | Maximum replication lag, or **staleness**, for reads from secondaries. | [optional] 
**Mode** | Pointer to **string** | [Read preference mode](https://docs.mongodb.com/manual/core/read-preference/#read-preference-modes) that specifies to which replica set member to route the read requests. | [optional] 
**TagSets** | Pointer to [**[][]DataLakeAtlasStoreReadPreferenceTag**]([]DataLakeAtlasStoreReadPreferenceTag.md) | List that contains [tag sets](https://docs.mongodb.com/manual/core/read-preference-tags/) or tag specification documents. If specified, Atlas Data Lake routes read requests to replica set member or members that are associated with the specified tags. | [optional] 

## Methods

### NewDataLakeAtlasStoreReadPreference

`func NewDataLakeAtlasStoreReadPreference() *DataLakeAtlasStoreReadPreference`

NewDataLakeAtlasStoreReadPreference instantiates a new DataLakeAtlasStoreReadPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeAtlasStoreReadPreferenceWithDefaults

`func NewDataLakeAtlasStoreReadPreferenceWithDefaults() *DataLakeAtlasStoreReadPreference`

NewDataLakeAtlasStoreReadPreferenceWithDefaults instantiates a new DataLakeAtlasStoreReadPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxStalenessSeconds

`func (o *DataLakeAtlasStoreReadPreference) GetMaxStalenessSeconds() int`

GetMaxStalenessSeconds returns the MaxStalenessSeconds field if non-nil, zero value otherwise.

### GetMaxStalenessSecondsOk

`func (o *DataLakeAtlasStoreReadPreference) GetMaxStalenessSecondsOk() (*int, bool)`

GetMaxStalenessSecondsOk returns a tuple with the MaxStalenessSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStalenessSeconds

`func (o *DataLakeAtlasStoreReadPreference) SetMaxStalenessSeconds(v int)`

SetMaxStalenessSeconds sets MaxStalenessSeconds field to given value.

### HasMaxStalenessSeconds

`func (o *DataLakeAtlasStoreReadPreference) HasMaxStalenessSeconds() bool`

HasMaxStalenessSeconds returns a boolean if a field has been set.

### GetMode

`func (o *DataLakeAtlasStoreReadPreference) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DataLakeAtlasStoreReadPreference) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DataLakeAtlasStoreReadPreference) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DataLakeAtlasStoreReadPreference) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTagSets

`func (o *DataLakeAtlasStoreReadPreference) GetTagSets() [][]DataLakeAtlasStoreReadPreferenceTag`

GetTagSets returns the TagSets field if non-nil, zero value otherwise.

### GetTagSetsOk

`func (o *DataLakeAtlasStoreReadPreference) GetTagSetsOk() (*[][]DataLakeAtlasStoreReadPreferenceTag, bool)`

GetTagSetsOk returns a tuple with the TagSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSets

`func (o *DataLakeAtlasStoreReadPreference) SetTagSets(v [][]DataLakeAtlasStoreReadPreferenceTag)`

SetTagSets sets TagSets field to given value.

### HasTagSets

`func (o *DataLakeAtlasStoreReadPreference) HasTagSets() bool`

HasTagSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


