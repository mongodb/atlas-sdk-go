# DataLakeApiBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable label that identifies the view, which corresponds to an aggregation pipeline on a collection. | [optional] 
**Pipeline** | Pointer to **string** | Aggregation pipeline stages to apply to the source collection. | [optional] 
**Source** | Pointer to **string** | Human-readable label that identifies the source collection for the view. | [optional] 

## Methods

### NewDataLakeApiBase

`func NewDataLakeApiBase() *DataLakeApiBase`

NewDataLakeApiBase instantiates a new DataLakeApiBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeApiBaseWithDefaults

`func NewDataLakeApiBaseWithDefaults() *DataLakeApiBase`

NewDataLakeApiBaseWithDefaults instantiates a new DataLakeApiBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataLakeApiBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLakeApiBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLakeApiBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataLakeApiBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPipeline

`func (o *DataLakeApiBase) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *DataLakeApiBase) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *DataLakeApiBase) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *DataLakeApiBase) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetSource

`func (o *DataLakeApiBase) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataLakeApiBase) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataLakeApiBase) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DataLakeApiBase) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


