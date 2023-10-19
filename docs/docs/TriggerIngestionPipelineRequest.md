# TriggerIngestionPipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetRetentionPolicy** | Pointer to [**DatasetRetentionPolicy**](DatasetRetentionPolicy.md) |  | [optional] 
**SnapshotId** | **string** | Unique 24-hexadecimal character string that identifies the snapshot. | 

## Methods

### NewTriggerIngestionPipelineRequest

`func NewTriggerIngestionPipelineRequest(snapshotId string, ) *TriggerIngestionPipelineRequest`

NewTriggerIngestionPipelineRequest instantiates a new TriggerIngestionPipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerIngestionPipelineRequestWithDefaults

`func NewTriggerIngestionPipelineRequestWithDefaults() *TriggerIngestionPipelineRequest`

NewTriggerIngestionPipelineRequestWithDefaults instantiates a new TriggerIngestionPipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetRetentionPolicy

`func (o *TriggerIngestionPipelineRequest) GetDatasetRetentionPolicy() DatasetRetentionPolicy`

GetDatasetRetentionPolicy returns the DatasetRetentionPolicy field if non-nil, zero value otherwise.

### GetDatasetRetentionPolicyOk

`func (o *TriggerIngestionPipelineRequest) GetDatasetRetentionPolicyOk() (*DatasetRetentionPolicy, bool)`

GetDatasetRetentionPolicyOk returns a tuple with the DatasetRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetRetentionPolicy

`func (o *TriggerIngestionPipelineRequest) SetDatasetRetentionPolicy(v DatasetRetentionPolicy)`

SetDatasetRetentionPolicy sets DatasetRetentionPolicy field to given value.

### HasDatasetRetentionPolicy

`func (o *TriggerIngestionPipelineRequest) HasDatasetRetentionPolicy() bool`

HasDatasetRetentionPolicy returns a boolean if a field has been set.
### GetSnapshotId

`func (o *TriggerIngestionPipelineRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *TriggerIngestionPipelineRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *TriggerIngestionPipelineRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


