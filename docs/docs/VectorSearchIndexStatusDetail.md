# VectorSearchIndexStatusDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | Pointer to [**VectorSearchIndexDefinition**](VectorSearchIndexDefinition.md) |  | [optional] 
**DefinitionVersion** | Pointer to [**SearchIndexDefinitionVersion**](SearchIndexDefinitionVersion.md) |  | [optional] 
**Message** | Pointer to **string** | Optional message describing an error. | [optional] 
**Queryable** | Pointer to **bool** | Flag that indicates whether the index generation is queryable on the host. | [optional] 
**Status** | Pointer to **string** | Condition of the search index when you made this request.  - &#x60;DELETING&#x60;: The index is being deleted. - &#x60;FAILED&#x60; The index build failed. Indexes can enter the FAILED state due to an invalid index definition. - &#x60;STALE&#x60;: The index is queryable but has stopped replicating data from the indexed collection. Searches on the index may return out-of-date data. - &#x60;PENDING&#x60;: Atlas has not yet started building the index. - &#x60;BUILDING&#x60;: Atlas is building or re-building the index after an edit. - &#x60;READY&#x60;: The index is ready and can support queries. | [optional] 

## Methods

### NewVectorSearchIndexStatusDetail

`func NewVectorSearchIndexStatusDetail() *VectorSearchIndexStatusDetail`

NewVectorSearchIndexStatusDetail instantiates a new VectorSearchIndexStatusDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchIndexStatusDetailWithDefaults

`func NewVectorSearchIndexStatusDetailWithDefaults() *VectorSearchIndexStatusDetail`

NewVectorSearchIndexStatusDetailWithDefaults instantiates a new VectorSearchIndexStatusDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *VectorSearchIndexStatusDetail) GetDefinition() VectorSearchIndexDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *VectorSearchIndexStatusDetail) GetDefinitionOk() (*VectorSearchIndexDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *VectorSearchIndexStatusDetail) SetDefinition(v VectorSearchIndexDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *VectorSearchIndexStatusDetail) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.
### GetDefinitionVersion

`func (o *VectorSearchIndexStatusDetail) GetDefinitionVersion() SearchIndexDefinitionVersion`

GetDefinitionVersion returns the DefinitionVersion field if non-nil, zero value otherwise.

### GetDefinitionVersionOk

`func (o *VectorSearchIndexStatusDetail) GetDefinitionVersionOk() (*SearchIndexDefinitionVersion, bool)`

GetDefinitionVersionOk returns a tuple with the DefinitionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionVersion

`func (o *VectorSearchIndexStatusDetail) SetDefinitionVersion(v SearchIndexDefinitionVersion)`

SetDefinitionVersion sets DefinitionVersion field to given value.

### HasDefinitionVersion

`func (o *VectorSearchIndexStatusDetail) HasDefinitionVersion() bool`

HasDefinitionVersion returns a boolean if a field has been set.
### GetMessage

`func (o *VectorSearchIndexStatusDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VectorSearchIndexStatusDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VectorSearchIndexStatusDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VectorSearchIndexStatusDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.
### GetQueryable

`func (o *VectorSearchIndexStatusDetail) GetQueryable() bool`

GetQueryable returns the Queryable field if non-nil, zero value otherwise.

### GetQueryableOk

`func (o *VectorSearchIndexStatusDetail) GetQueryableOk() (*bool, bool)`

GetQueryableOk returns a tuple with the Queryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryable

`func (o *VectorSearchIndexStatusDetail) SetQueryable(v bool)`

SetQueryable sets Queryable field to given value.

### HasQueryable

`func (o *VectorSearchIndexStatusDetail) HasQueryable() bool`

HasQueryable returns a boolean if a field has been set.
### GetStatus

`func (o *VectorSearchIndexStatusDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VectorSearchIndexStatusDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VectorSearchIndexStatusDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VectorSearchIndexStatusDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


