# IndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collation** | Pointer to [**Collation**](Collation.md) |  | [optional] 
**Collection** | **string** | Human-readable label of the collection for which MongoDB Cloud creates an index. | 
**Db** | **string** | Human-readable label of the database that holds the collection on which MongoDB Cloud creates an index. | 
**Keys** | Pointer to **[]map[string]string** | List that contains one or more objects that describe the parameters that you want to index. | [optional] 
**Options** | Pointer to [**IndexOptions**](IndexOptions.md) |  | [optional] 

## Methods

### NewIndexRequest

`func NewIndexRequest(collection string, db string, ) *IndexRequest`

NewIndexRequest instantiates a new IndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexRequestWithDefaults

`func NewIndexRequestWithDefaults() *IndexRequest`

NewIndexRequestWithDefaults instantiates a new IndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollation

`func (o *IndexRequest) GetCollation() Collation`

GetCollation returns the Collation field if non-nil, zero value otherwise.

### GetCollationOk

`func (o *IndexRequest) GetCollationOk() (*Collation, bool)`

GetCollationOk returns a tuple with the Collation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollation

`func (o *IndexRequest) SetCollation(v Collation)`

SetCollation sets Collation field to given value.

### HasCollation

`func (o *IndexRequest) HasCollation() bool`

HasCollation returns a boolean if a field has been set.

### GetCollection

`func (o *IndexRequest) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *IndexRequest) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *IndexRequest) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetDb

`func (o *IndexRequest) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *IndexRequest) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *IndexRequest) SetDb(v string)`

SetDb sets Db field to given value.


### GetKeys

`func (o *IndexRequest) GetKeys() []map[string]string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *IndexRequest) GetKeysOk() (*[]map[string]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *IndexRequest) SetKeys(v []map[string]string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *IndexRequest) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetOptions

`func (o *IndexRequest) GetOptions() IndexOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IndexRequest) GetOptionsOk() (*IndexOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IndexRequest) SetOptions(v IndexOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IndexRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


