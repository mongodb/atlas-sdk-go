# ApiAtlasCollectionRestoreCollectionStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveTargetNamespace** | Pointer to **string** | Actual target namespace after restore (e.g. after conflict rename). | [optional] [readonly] 
**IndexStatus** | Pointer to [**ApiAtlasCollectionRestoreIndexStatus**](ApiAtlasCollectionRestoreIndexStatus.md) |  | [optional] 
**RestoredDocuments** | Pointer to **int64** | Number of documents restored so far. | [optional] [readonly] 
**SourceNamespace** | Pointer to **string** | Source namespace that was requested to restore. | [optional] [readonly] 
**State** | Pointer to **string** | Current state of this collection within the restore job. | [optional] [readonly] 
**TargetNamespace** | Pointer to **string** | Requested target namespace for the restored collection. | [optional] [readonly] 
**TotalDocuments** | Pointer to **int64** | Total document count for this collection. | [optional] [readonly] 

## Methods

### NewApiAtlasCollectionRestoreCollectionStateResponse

`func NewApiAtlasCollectionRestoreCollectionStateResponse() *ApiAtlasCollectionRestoreCollectionStateResponse`

NewApiAtlasCollectionRestoreCollectionStateResponse instantiates a new ApiAtlasCollectionRestoreCollectionStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCollectionRestoreCollectionStateResponseWithDefaults

`func NewApiAtlasCollectionRestoreCollectionStateResponseWithDefaults() *ApiAtlasCollectionRestoreCollectionStateResponse`

NewApiAtlasCollectionRestoreCollectionStateResponseWithDefaults instantiates a new ApiAtlasCollectionRestoreCollectionStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveTargetNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetEffectiveTargetNamespace() string`

GetEffectiveTargetNamespace returns the EffectiveTargetNamespace field if non-nil, zero value otherwise.

### GetEffectiveTargetNamespaceOk

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetEffectiveTargetNamespaceOk() (*string, bool)`

GetEffectiveTargetNamespaceOk returns a tuple with the EffectiveTargetNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTargetNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetEffectiveTargetNamespace(v string)`

SetEffectiveTargetNamespace sets EffectiveTargetNamespace field to given value.

### HasEffectiveTargetNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasEffectiveTargetNamespace() bool`

HasEffectiveTargetNamespace returns a boolean if a field has been set.
### GetIndexStatus

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetIndexStatus() ApiAtlasCollectionRestoreIndexStatus`

GetIndexStatus returns the IndexStatus field if non-nil, zero value otherwise.

### GetIndexStatusOk

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetIndexStatusOk() (*ApiAtlasCollectionRestoreIndexStatus, bool)`

GetIndexStatusOk returns a tuple with the IndexStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStatus

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetIndexStatus(v ApiAtlasCollectionRestoreIndexStatus)`

SetIndexStatus sets IndexStatus field to given value.

### HasIndexStatus

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasIndexStatus() bool`

HasIndexStatus returns a boolean if a field has been set.
### GetRestoredDocuments

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetRestoredDocuments() int64`

GetRestoredDocuments returns the RestoredDocuments field if non-nil, zero value otherwise.

### GetRestoredDocumentsOk

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetRestoredDocumentsOk() (*int64, bool)`

GetRestoredDocumentsOk returns a tuple with the RestoredDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredDocuments

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetRestoredDocuments(v int64)`

SetRestoredDocuments sets RestoredDocuments field to given value.

### HasRestoredDocuments

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasRestoredDocuments() bool`

HasRestoredDocuments returns a boolean if a field has been set.
### GetSourceNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetSourceNamespace() string`

GetSourceNamespace returns the SourceNamespace field if non-nil, zero value otherwise.

### GetSourceNamespaceOk

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetSourceNamespaceOk() (*string, bool)`

GetSourceNamespaceOk returns a tuple with the SourceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetSourceNamespace(v string)`

SetSourceNamespace sets SourceNamespace field to given value.

### HasSourceNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasSourceNamespace() bool`

HasSourceNamespace returns a boolean if a field has been set.
### GetState

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasState() bool`

HasState returns a boolean if a field has been set.
### GetTargetNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetTargetNamespace() string`

GetTargetNamespace returns the TargetNamespace field if non-nil, zero value otherwise.

### GetTargetNamespaceOk

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetTargetNamespaceOk() (*string, bool)`

GetTargetNamespaceOk returns a tuple with the TargetNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetTargetNamespace(v string)`

SetTargetNamespace sets TargetNamespace field to given value.

### HasTargetNamespace

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasTargetNamespace() bool`

HasTargetNamespace returns a boolean if a field has been set.
### GetTotalDocuments

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetTotalDocuments() int64`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) GetTotalDocumentsOk() (*int64, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) SetTotalDocuments(v int64)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *ApiAtlasCollectionRestoreCollectionStateResponse) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


