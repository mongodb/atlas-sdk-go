# DataLakeDatabaseInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]DataLakeDatabaseCollection**](DataLakeDatabaseCollection.md) | Array of collections and data sources that map to a &#x60;&#x60;stores&#x60;&#x60; data store. | [optional] 
**MaxWildcardCollections** | Pointer to **int** | Maximum number of wildcard collections in the database. This only applies to S3 data sources. | [optional] [default to 100]
**Name** | Pointer to **string** | Human-readable label that identifies the database to which the data lake maps data. | [optional] 
**Views** | Pointer to [**[]DataLakeApiBase**](DataLakeApiBase.md) | Array of aggregation pipelines that apply to the collection. This only applies to S3 data sources. | [optional] 

## Methods

### NewDataLakeDatabaseInstance

`func NewDataLakeDatabaseInstance() *DataLakeDatabaseInstance`

NewDataLakeDatabaseInstance instantiates a new DataLakeDatabaseInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeDatabaseInstanceWithDefaults

`func NewDataLakeDatabaseInstanceWithDefaults() *DataLakeDatabaseInstance`

NewDataLakeDatabaseInstanceWithDefaults instantiates a new DataLakeDatabaseInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *DataLakeDatabaseInstance) GetCollections() []DataLakeDatabaseCollection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *DataLakeDatabaseInstance) GetCollectionsOk() (*[]DataLakeDatabaseCollection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *DataLakeDatabaseInstance) SetCollections(v []DataLakeDatabaseCollection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *DataLakeDatabaseInstance) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetMaxWildcardCollections

`func (o *DataLakeDatabaseInstance) GetMaxWildcardCollections() int`

GetMaxWildcardCollections returns the MaxWildcardCollections field if non-nil, zero value otherwise.

### GetMaxWildcardCollectionsOk

`func (o *DataLakeDatabaseInstance) GetMaxWildcardCollectionsOk() (*int, bool)`

GetMaxWildcardCollectionsOk returns a tuple with the MaxWildcardCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWildcardCollections

`func (o *DataLakeDatabaseInstance) SetMaxWildcardCollections(v int)`

SetMaxWildcardCollections sets MaxWildcardCollections field to given value.

### HasMaxWildcardCollections

`func (o *DataLakeDatabaseInstance) HasMaxWildcardCollections() bool`

HasMaxWildcardCollections returns a boolean if a field has been set.

### GetName

`func (o *DataLakeDatabaseInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLakeDatabaseInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLakeDatabaseInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataLakeDatabaseInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetViews

`func (o *DataLakeDatabaseInstance) GetViews() []DataLakeApiBase`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *DataLakeDatabaseInstance) GetViewsOk() (*[]DataLakeApiBase, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *DataLakeDatabaseInstance) SetViews(v []DataLakeApiBase)`

SetViews sets Views field to given value.

### HasViews

`func (o *DataLakeDatabaseInstance) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


