# DataLakeStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Databases** | Pointer to [**[]DataLakeDatabase**](DataLakeDatabase.md) | Array that contains the queryable databases and collections for this data lake. | [optional] 
**Stores** | Pointer to [**[]DataLakeStore**](DataLakeStore.md) | Array that contains the data stores for the data lake. | [optional] 

## Methods

### NewDataLakeStorage

`func NewDataLakeStorage() *DataLakeStorage`

NewDataLakeStorage instantiates a new DataLakeStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeStorageWithDefaults

`func NewDataLakeStorageWithDefaults() *DataLakeStorage`

NewDataLakeStorageWithDefaults instantiates a new DataLakeStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabases

`func (o *DataLakeStorage) GetDatabases() []DataLakeDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *DataLakeStorage) GetDatabasesOk() (*[]DataLakeDatabase, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *DataLakeStorage) SetDatabases(v []DataLakeDatabase)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *DataLakeStorage) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetStores

`func (o *DataLakeStorage) GetStores() []DataLakeStore`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *DataLakeStorage) GetStoresOk() (*[]DataLakeStore, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *DataLakeStorage) SetStores(v []DataLakeStore)`

SetStores sets Stores field to given value.

### HasStores

`func (o *DataLakeStorage) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


