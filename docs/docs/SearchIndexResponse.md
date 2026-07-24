# SearchIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | Pointer to **string** | Label that identifies the collection that contains one or more Atlas Search indexes. | [optional] 
**Database** | Pointer to **string** | Label that identifies the database that contains the collection with one or more Atlas Search indexes. | [optional] 
**IndexID** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this Atlas Search index. | [optional] 
**LatestDefinition** | Pointer to [**BaseSearchIndexResponseLatestDefinition**](BaseSearchIndexResponseLatestDefinition.md) |  | [optional] 
**LatestDefinitionVersion** | Pointer to [**SearchIndexDefinitionVersion**](SearchIndexDefinitionVersion.md) |  | [optional] 
**Name** | Pointer to **string** | Label that identifies this index. Within each namespace, the names of all indexes must be unique. | [optional] 
**Queryable** | Pointer to **bool** | Flag that indicates whether the index is queryable on all hosts. | [optional] 
**Status** | Pointer to **string** | Condition of the search index when you made this request.  - &#x60;DELETING&#x60;: The index is being deleted. - &#x60;FAILED&#x60; The index build failed. Indexes can enter the FAILED state due to an invalid index definition. - &#x60;STALE&#x60;: The index is queryable but has stopped replicating data from the indexed collection. Searches on the index may return out-of-date data. - &#x60;PENDING&#x60;: Atlas has not yet started building the index. - &#x60;BUILDING&#x60;: Atlas is building or re-building the index after an edit. - &#x60;READY&#x60;: The index is ready and can support queries. | [optional] 
**StatusDetail** | Pointer to [**[]VectorSearchHostStatusDetail**](VectorSearchHostStatusDetail.md) | List of documents detailing index status on each host. | [optional] 
**Type** | Pointer to **string** | Type of the index. The default type is search. | [optional] 
**SynonymMappingStatus** | Pointer to **string** | Status that describes this index&#39;s synonym mappings. This status appears only if the index has synonyms defined. | [optional] 
**SynonymMappingStatusDetail** | Pointer to [**[]map[string]SynonymMappingStatusDetail**](map[string]SynonymMappingStatusDetail.md) | A list of documents describing the status of the index&#39;s synonym mappings on each search host. Only appears if the index has synonyms defined. | [optional] 

## Methods

### NewSearchIndexResponse

`func NewSearchIndexResponse() *SearchIndexResponse`

NewSearchIndexResponse instantiates a new SearchIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexResponseWithDefaults

`func NewSearchIndexResponseWithDefaults() *SearchIndexResponse`

NewSearchIndexResponseWithDefaults instantiates a new SearchIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *SearchIndexResponse) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *SearchIndexResponse) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *SearchIndexResponse) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *SearchIndexResponse) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### SetCollectionNameNil

`func (o *SearchIndexResponse) SetCollectionNameNil()`

SetCollectionNameNil sets CollectionName to an explicit JSON null when marshaled, overriding any value previously set with SetCollectionName. Calling SetCollectionName again clears the null override.

### GetDatabase

`func (o *SearchIndexResponse) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *SearchIndexResponse) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *SearchIndexResponse) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *SearchIndexResponse) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### SetDatabaseNil

`func (o *SearchIndexResponse) SetDatabaseNil()`

SetDatabaseNil sets Database to an explicit JSON null when marshaled, overriding any value previously set with SetDatabase. Calling SetDatabase again clears the null override.

### GetIndexID

`func (o *SearchIndexResponse) GetIndexID() string`

GetIndexID returns the IndexID field if non-nil, zero value otherwise.

### GetIndexIDOk

`func (o *SearchIndexResponse) GetIndexIDOk() (*string, bool)`

GetIndexIDOk returns a tuple with the IndexID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexID

`func (o *SearchIndexResponse) SetIndexID(v string)`

SetIndexID sets IndexID field to given value.

### HasIndexID

`func (o *SearchIndexResponse) HasIndexID() bool`

HasIndexID returns a boolean if a field has been set.

### SetIndexIDNil

`func (o *SearchIndexResponse) SetIndexIDNil()`

SetIndexIDNil sets IndexID to an explicit JSON null when marshaled, overriding any value previously set with SetIndexID. Calling SetIndexID again clears the null override.

### GetLatestDefinition

`func (o *SearchIndexResponse) GetLatestDefinition() BaseSearchIndexResponseLatestDefinition`

GetLatestDefinition returns the LatestDefinition field if non-nil, zero value otherwise.

### GetLatestDefinitionOk

`func (o *SearchIndexResponse) GetLatestDefinitionOk() (*BaseSearchIndexResponseLatestDefinition, bool)`

GetLatestDefinitionOk returns a tuple with the LatestDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDefinition

`func (o *SearchIndexResponse) SetLatestDefinition(v BaseSearchIndexResponseLatestDefinition)`

SetLatestDefinition sets LatestDefinition field to given value.

### HasLatestDefinition

`func (o *SearchIndexResponse) HasLatestDefinition() bool`

HasLatestDefinition returns a boolean if a field has been set.

### SetLatestDefinitionNil

`func (o *SearchIndexResponse) SetLatestDefinitionNil()`

SetLatestDefinitionNil sets LatestDefinition to an explicit JSON null when marshaled, overriding any value previously set with SetLatestDefinition. Calling SetLatestDefinition again clears the null override.

### GetLatestDefinitionVersion

`func (o *SearchIndexResponse) GetLatestDefinitionVersion() SearchIndexDefinitionVersion`

GetLatestDefinitionVersion returns the LatestDefinitionVersion field if non-nil, zero value otherwise.

### GetLatestDefinitionVersionOk

`func (o *SearchIndexResponse) GetLatestDefinitionVersionOk() (*SearchIndexDefinitionVersion, bool)`

GetLatestDefinitionVersionOk returns a tuple with the LatestDefinitionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDefinitionVersion

`func (o *SearchIndexResponse) SetLatestDefinitionVersion(v SearchIndexDefinitionVersion)`

SetLatestDefinitionVersion sets LatestDefinitionVersion field to given value.

### HasLatestDefinitionVersion

`func (o *SearchIndexResponse) HasLatestDefinitionVersion() bool`

HasLatestDefinitionVersion returns a boolean if a field has been set.

### SetLatestDefinitionVersionNil

`func (o *SearchIndexResponse) SetLatestDefinitionVersionNil()`

SetLatestDefinitionVersionNil sets LatestDefinitionVersion to an explicit JSON null when marshaled, overriding any value previously set with SetLatestDefinitionVersion. Calling SetLatestDefinitionVersion again clears the null override.

### GetName

`func (o *SearchIndexResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchIndexResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchIndexResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchIndexResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SearchIndexResponse) SetNameNil()`

SetNameNil sets Name to an explicit JSON null when marshaled, overriding any value previously set with SetName. Calling SetName again clears the null override.

### GetQueryable

`func (o *SearchIndexResponse) GetQueryable() bool`

GetQueryable returns the Queryable field if non-nil, zero value otherwise.

### GetQueryableOk

`func (o *SearchIndexResponse) GetQueryableOk() (*bool, bool)`

GetQueryableOk returns a tuple with the Queryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryable

`func (o *SearchIndexResponse) SetQueryable(v bool)`

SetQueryable sets Queryable field to given value.

### HasQueryable

`func (o *SearchIndexResponse) HasQueryable() bool`

HasQueryable returns a boolean if a field has been set.

### SetQueryableNil

`func (o *SearchIndexResponse) SetQueryableNil()`

SetQueryableNil sets Queryable to an explicit JSON null when marshaled, overriding any value previously set with SetQueryable. Calling SetQueryable again clears the null override.

### GetStatus

`func (o *SearchIndexResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchIndexResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchIndexResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchIndexResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SearchIndexResponse) SetStatusNil()`

SetStatusNil sets Status to an explicit JSON null when marshaled, overriding any value previously set with SetStatus. Calling SetStatus again clears the null override.

### GetStatusDetail

`func (o *SearchIndexResponse) GetStatusDetail() []VectorSearchHostStatusDetail`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *SearchIndexResponse) GetStatusDetailOk() (*[]VectorSearchHostStatusDetail, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *SearchIndexResponse) SetStatusDetail(v []VectorSearchHostStatusDetail)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *SearchIndexResponse) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### SetStatusDetailNil

`func (o *SearchIndexResponse) SetStatusDetailNil()`

SetStatusDetailNil sets StatusDetail to an explicit JSON null when marshaled, overriding any value previously set with SetStatusDetail. Calling SetStatusDetail again clears the null override.

### GetType

`func (o *SearchIndexResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchIndexResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchIndexResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchIndexResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SearchIndexResponse) SetTypeNil()`

SetTypeNil sets Type to an explicit JSON null when marshaled, overriding any value previously set with SetType. Calling SetType again clears the null override.

### GetSynonymMappingStatus

`func (o *SearchIndexResponse) GetSynonymMappingStatus() string`

GetSynonymMappingStatus returns the SynonymMappingStatus field if non-nil, zero value otherwise.

### GetSynonymMappingStatusOk

`func (o *SearchIndexResponse) GetSynonymMappingStatusOk() (*string, bool)`

GetSynonymMappingStatusOk returns a tuple with the SynonymMappingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonymMappingStatus

`func (o *SearchIndexResponse) SetSynonymMappingStatus(v string)`

SetSynonymMappingStatus sets SynonymMappingStatus field to given value.

### HasSynonymMappingStatus

`func (o *SearchIndexResponse) HasSynonymMappingStatus() bool`

HasSynonymMappingStatus returns a boolean if a field has been set.

### SetSynonymMappingStatusNil

`func (o *SearchIndexResponse) SetSynonymMappingStatusNil()`

SetSynonymMappingStatusNil sets SynonymMappingStatus to an explicit JSON null when marshaled, overriding any value previously set with SetSynonymMappingStatus. Calling SetSynonymMappingStatus again clears the null override.

### GetSynonymMappingStatusDetail

`func (o *SearchIndexResponse) GetSynonymMappingStatusDetail() []map[string]SynonymMappingStatusDetail`

GetSynonymMappingStatusDetail returns the SynonymMappingStatusDetail field if non-nil, zero value otherwise.

### GetSynonymMappingStatusDetailOk

`func (o *SearchIndexResponse) GetSynonymMappingStatusDetailOk() (*[]map[string]SynonymMappingStatusDetail, bool)`

GetSynonymMappingStatusDetailOk returns a tuple with the SynonymMappingStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonymMappingStatusDetail

`func (o *SearchIndexResponse) SetSynonymMappingStatusDetail(v []map[string]SynonymMappingStatusDetail)`

SetSynonymMappingStatusDetail sets SynonymMappingStatusDetail field to given value.

### HasSynonymMappingStatusDetail

`func (o *SearchIndexResponse) HasSynonymMappingStatusDetail() bool`

HasSynonymMappingStatusDetail returns a boolean if a field has been set.

### SetSynonymMappingStatusDetailNil

`func (o *SearchIndexResponse) SetSynonymMappingStatusDetailNil()`

SetSynonymMappingStatusDetailNil sets SynonymMappingStatusDetail to an explicit JSON null when marshaled, overriding any value previously set with SetSynonymMappingStatusDetail. Calling SetSynonymMappingStatusDetail again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


