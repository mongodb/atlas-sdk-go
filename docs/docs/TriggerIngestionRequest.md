# TriggerIngestionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **string** | Unique 24-hexadecimal character string that identifies the snapshot. | 

## Methods

### NewTriggerIngestionRequest

`func NewTriggerIngestionRequest(snapshotId string, ) *TriggerIngestionRequest`

NewTriggerIngestionRequest instantiates a new TriggerIngestionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerIngestionRequestWithDefaults

`func NewTriggerIngestionRequestWithDefaults() *TriggerIngestionRequest`

NewTriggerIngestionRequestWithDefaults instantiates a new TriggerIngestionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *TriggerIngestionRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *TriggerIngestionRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *TriggerIngestionRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


