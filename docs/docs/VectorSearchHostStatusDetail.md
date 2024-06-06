# VectorSearchHostStatusDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | Hostname that corresponds to the status detail. | [optional] 
**MainIndex** | Pointer to [**VectorSearchIndexStatusDetail**](VectorSearchIndexStatusDetail.md) |  | [optional] 
**Queryable** | Pointer to **bool** | Flag that indicates whether the index is queryable on the host. | [optional] 
**StagedIndex** | Pointer to [**VectorSearchIndexStatusDetail**](VectorSearchIndexStatusDetail.md) |  | [optional] 
**Status** | Pointer to **string** | Condition of the search index when you made this request.   | Status | Index Condition |  |---|---|  | DELETING | The index is being deleted. |  | FAILED | The index build failed. Indexes can enter the FAILED state due to an invalid index definition. |  | STALE | The index is queryable but has stopped replicating data from the indexed collection. Searches on the index may return out-of-date data. |  | PENDING | Atlas has not yet started building the index. |  | BUILDING | Atlas is building or re-building the index after an edit. |  | READY | The index is ready and can support queries. |  | [optional] 

## Methods

### NewVectorSearchHostStatusDetail

`func NewVectorSearchHostStatusDetail() *VectorSearchHostStatusDetail`

NewVectorSearchHostStatusDetail instantiates a new VectorSearchHostStatusDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchHostStatusDetailWithDefaults

`func NewVectorSearchHostStatusDetailWithDefaults() *VectorSearchHostStatusDetail`

NewVectorSearchHostStatusDetailWithDefaults instantiates a new VectorSearchHostStatusDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VectorSearchHostStatusDetail) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VectorSearchHostStatusDetail) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VectorSearchHostStatusDetail) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VectorSearchHostStatusDetail) HasHostname() bool`

HasHostname returns a boolean if a field has been set.
### GetMainIndex

`func (o *VectorSearchHostStatusDetail) GetMainIndex() VectorSearchIndexStatusDetail`

GetMainIndex returns the MainIndex field if non-nil, zero value otherwise.

### GetMainIndexOk

`func (o *VectorSearchHostStatusDetail) GetMainIndexOk() (*VectorSearchIndexStatusDetail, bool)`

GetMainIndexOk returns a tuple with the MainIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainIndex

`func (o *VectorSearchHostStatusDetail) SetMainIndex(v VectorSearchIndexStatusDetail)`

SetMainIndex sets MainIndex field to given value.

### HasMainIndex

`func (o *VectorSearchHostStatusDetail) HasMainIndex() bool`

HasMainIndex returns a boolean if a field has been set.
### GetQueryable

`func (o *VectorSearchHostStatusDetail) GetQueryable() bool`

GetQueryable returns the Queryable field if non-nil, zero value otherwise.

### GetQueryableOk

`func (o *VectorSearchHostStatusDetail) GetQueryableOk() (*bool, bool)`

GetQueryableOk returns a tuple with the Queryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryable

`func (o *VectorSearchHostStatusDetail) SetQueryable(v bool)`

SetQueryable sets Queryable field to given value.

### HasQueryable

`func (o *VectorSearchHostStatusDetail) HasQueryable() bool`

HasQueryable returns a boolean if a field has been set.
### GetStagedIndex

`func (o *VectorSearchHostStatusDetail) GetStagedIndex() VectorSearchIndexStatusDetail`

GetStagedIndex returns the StagedIndex field if non-nil, zero value otherwise.

### GetStagedIndexOk

`func (o *VectorSearchHostStatusDetail) GetStagedIndexOk() (*VectorSearchIndexStatusDetail, bool)`

GetStagedIndexOk returns a tuple with the StagedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagedIndex

`func (o *VectorSearchHostStatusDetail) SetStagedIndex(v VectorSearchIndexStatusDetail)`

SetStagedIndex sets StagedIndex field to given value.

### HasStagedIndex

`func (o *VectorSearchHostStatusDetail) HasStagedIndex() bool`

HasStagedIndex returns a boolean if a field has been set.
### GetStatus

`func (o *VectorSearchHostStatusDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VectorSearchHostStatusDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VectorSearchHostStatusDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VectorSearchHostStatusDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


