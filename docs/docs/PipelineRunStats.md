# PipelineRunStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesExported** | Pointer to **int64** | Total data size in bytes exported for this pipeline run. | [optional] [readonly] 
**NumDocs** | Pointer to **int64** | Number of docs ingested for a this pipeline run. | [optional] [readonly] 

## Methods

### NewPipelineRunStats

`func NewPipelineRunStats() *PipelineRunStats`

NewPipelineRunStats instantiates a new PipelineRunStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineRunStatsWithDefaults

`func NewPipelineRunStatsWithDefaults() *PipelineRunStats`

NewPipelineRunStatsWithDefaults instantiates a new PipelineRunStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesExported

`func (o *PipelineRunStats) GetBytesExported() int64`

GetBytesExported returns the BytesExported field if non-nil, zero value otherwise.

### GetBytesExportedOk

`func (o *PipelineRunStats) GetBytesExportedOk() (*int64, bool)`

GetBytesExportedOk returns a tuple with the BytesExported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesExported

`func (o *PipelineRunStats) SetBytesExported(v int64)`

SetBytesExported sets BytesExported field to given value.

### HasBytesExported

`func (o *PipelineRunStats) HasBytesExported() bool`

HasBytesExported returns a boolean if a field has been set.

### GetNumDocs

`func (o *PipelineRunStats) GetNumDocs() int64`

GetNumDocs returns the NumDocs field if non-nil, zero value otherwise.

### GetNumDocsOk

`func (o *PipelineRunStats) GetNumDocsOk() (*int64, bool)`

GetNumDocsOk returns a tuple with the NumDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDocs

`func (o *PipelineRunStats) SetNumDocs(v int64)`

SetNumDocs sets NumDocs field to given value.

### HasNumDocs

`func (o *PipelineRunStats) HasNumDocs() bool`

HasNumDocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


