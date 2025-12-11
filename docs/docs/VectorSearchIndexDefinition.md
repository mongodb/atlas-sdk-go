# VectorSearchIndexDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]any**](any.md) | Settings that configure the fields, one per object, to index. You must define at least one \&quot;vector\&quot; type field. You can optionally define \&quot;filter\&quot; type fields also. | [optional] 
**NumPartitions** | Pointer to **int** | Number of index partitions. Allowed values are [1, 2, 4]. | [optional] [default to 1]
**StoredSource** | Pointer to [**any**](interface{}.md) | Flag that indicates whether to store all fields (true) on Atlas Search. By default, Atlas doesn&#39;t store (false) the fields on Atlas Search.  Alternatively, you can specify an object that only contains the list of fields to store (include) or not store (exclude) on Atlas Search. Note that storing all fields (true) is not allowed for vector search indexes. To learn more, see Stored Source Fields. | [optional] 

## Methods

### NewVectorSearchIndexDefinition

`func NewVectorSearchIndexDefinition() *VectorSearchIndexDefinition`

NewVectorSearchIndexDefinition instantiates a new VectorSearchIndexDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchIndexDefinitionWithDefaults

`func NewVectorSearchIndexDefinitionWithDefaults() *VectorSearchIndexDefinition`

NewVectorSearchIndexDefinitionWithDefaults instantiates a new VectorSearchIndexDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *VectorSearchIndexDefinition) GetFields() []any`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *VectorSearchIndexDefinition) GetFieldsOk() (*[]any, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *VectorSearchIndexDefinition) SetFields(v []any)`

SetFields sets Fields field to given value.

### HasFields

`func (o *VectorSearchIndexDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.
### GetNumPartitions

`func (o *VectorSearchIndexDefinition) GetNumPartitions() int`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *VectorSearchIndexDefinition) GetNumPartitionsOk() (*int, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *VectorSearchIndexDefinition) SetNumPartitions(v int)`

SetNumPartitions sets NumPartitions field to given value.

### HasNumPartitions

`func (o *VectorSearchIndexDefinition) HasNumPartitions() bool`

HasNumPartitions returns a boolean if a field has been set.
### GetStoredSource

`func (o *VectorSearchIndexDefinition) GetStoredSource() any`

GetStoredSource returns the StoredSource field if non-nil, zero value otherwise.

### GetStoredSourceOk

`func (o *VectorSearchIndexDefinition) GetStoredSourceOk() (*any, bool)`

GetStoredSourceOk returns a tuple with the StoredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredSource

`func (o *VectorSearchIndexDefinition) SetStoredSource(v any)`

SetStoredSource sets StoredSource field to given value.

### HasStoredSource

`func (o *VectorSearchIndexDefinition) HasStoredSource() bool`

HasStoredSource returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


