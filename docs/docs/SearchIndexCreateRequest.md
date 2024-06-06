# SearchIndexCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | **string** | Label that identifies the collection to create an Atlas Search index in. | 
**Database** | **string** | Label that identifies the database that contains the collection to create an Atlas Search index in. | 
**Name** | **string** | Label that identifies this index. Within each namespace, names of all indexes in the namespace must be unique. | 
**Type** | Pointer to **string** | Type of the index. The default type is search. | [optional] 
**Definition** | Pointer to [**BaseSearchIndexCreateRequestDefinition**](BaseSearchIndexCreateRequestDefinition.md) |  | [optional] 

## Methods

### NewSearchIndexCreateRequest

`func NewSearchIndexCreateRequest(collectionName string, database string, name string, ) *SearchIndexCreateRequest`

NewSearchIndexCreateRequest instantiates a new SearchIndexCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexCreateRequestWithDefaults

`func NewSearchIndexCreateRequestWithDefaults() *SearchIndexCreateRequest`

NewSearchIndexCreateRequestWithDefaults instantiates a new SearchIndexCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *SearchIndexCreateRequest) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *SearchIndexCreateRequest) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *SearchIndexCreateRequest) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### GetDatabase

`func (o *SearchIndexCreateRequest) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *SearchIndexCreateRequest) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *SearchIndexCreateRequest) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### GetName

`func (o *SearchIndexCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchIndexCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchIndexCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### GetType

`func (o *SearchIndexCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchIndexCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchIndexCreateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchIndexCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.
### GetDefinition

`func (o *SearchIndexCreateRequest) GetDefinition() BaseSearchIndexCreateRequestDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *SearchIndexCreateRequest) GetDefinitionOk() (*BaseSearchIndexCreateRequestDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *SearchIndexCreateRequest) SetDefinition(v BaseSearchIndexCreateRequestDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *SearchIndexCreateRequest) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


