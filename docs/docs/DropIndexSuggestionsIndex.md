# DropIndexSuggestionsIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCount** | Pointer to **int64** | Usage count (since last restart) of index. | [optional] 
**Index** | Pointer to [**[]any**](any.md) | List that contains documents that specify a key in the index and its sort order. | [optional] 
**Name** | Pointer to **string** | Name of index. | [optional] 
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] 
**Shards** | Pointer to **[]string** | List that contains strings that specifies the shards where the index is found. | [optional] 
**Since** | Pointer to **time.Time** | Date of most recent usage of index. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] 
**SizeBytes** | Pointer to **int64** | Size of index. | [optional] 

## Methods

### NewDropIndexSuggestionsIndex

`func NewDropIndexSuggestionsIndex() *DropIndexSuggestionsIndex`

NewDropIndexSuggestionsIndex instantiates a new DropIndexSuggestionsIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropIndexSuggestionsIndexWithDefaults

`func NewDropIndexSuggestionsIndexWithDefaults() *DropIndexSuggestionsIndex`

NewDropIndexSuggestionsIndexWithDefaults instantiates a new DropIndexSuggestionsIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCount

`func (o *DropIndexSuggestionsIndex) GetAccessCount() int64`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *DropIndexSuggestionsIndex) GetAccessCountOk() (*int64, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *DropIndexSuggestionsIndex) SetAccessCount(v int64)`

SetAccessCount sets AccessCount field to given value.

### HasAccessCount

`func (o *DropIndexSuggestionsIndex) HasAccessCount() bool`

HasAccessCount returns a boolean if a field has been set.
### GetIndex

`func (o *DropIndexSuggestionsIndex) GetIndex() []any`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DropIndexSuggestionsIndex) GetIndexOk() (*[]any, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DropIndexSuggestionsIndex) SetIndex(v []any)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DropIndexSuggestionsIndex) HasIndex() bool`

HasIndex returns a boolean if a field has been set.
### GetName

`func (o *DropIndexSuggestionsIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DropIndexSuggestionsIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DropIndexSuggestionsIndex) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DropIndexSuggestionsIndex) HasName() bool`

HasName returns a boolean if a field has been set.
### GetNamespace

`func (o *DropIndexSuggestionsIndex) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DropIndexSuggestionsIndex) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DropIndexSuggestionsIndex) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DropIndexSuggestionsIndex) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.
### GetShards

`func (o *DropIndexSuggestionsIndex) GetShards() []string`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *DropIndexSuggestionsIndex) GetShardsOk() (*[]string, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *DropIndexSuggestionsIndex) SetShards(v []string)`

SetShards sets Shards field to given value.

### HasShards

`func (o *DropIndexSuggestionsIndex) HasShards() bool`

HasShards returns a boolean if a field has been set.
### GetSince

`func (o *DropIndexSuggestionsIndex) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *DropIndexSuggestionsIndex) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *DropIndexSuggestionsIndex) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *DropIndexSuggestionsIndex) HasSince() bool`

HasSince returns a boolean if a field has been set.
### GetSizeBytes

`func (o *DropIndexSuggestionsIndex) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *DropIndexSuggestionsIndex) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *DropIndexSuggestionsIndex) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *DropIndexSuggestionsIndex) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


