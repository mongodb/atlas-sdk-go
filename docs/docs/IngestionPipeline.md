# IngestionPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Data Lake Pipeline. | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** | Timestamp that indicates when the Data Lake Pipeline was created. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the group. | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **time.Time** | Timestamp that indicates the last time that the Data Lake Pipeline was updated. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of this Data Lake Pipeline. | [optional] 
**Sink** | Pointer to [**IngestionSink**](IngestionSink.md) |  | [optional] 
**Source** | Pointer to [**IngestionSource**](IngestionSource.md) |  | [optional] 
**State** | Pointer to **string** | State of this Data Lake Pipeline. | [optional] [readonly] 
**Transformations** | Pointer to [**[]FieldTransformation**](FieldTransformation.md) | Fields to be excluded for this Data Lake Pipeline. | [optional] 

## Methods

### NewIngestionPipeline

`func NewIngestionPipeline() *IngestionPipeline`

NewIngestionPipeline instantiates a new IngestionPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionPipelineWithDefaults

`func NewIngestionPipelineWithDefaults() *IngestionPipeline`

NewIngestionPipelineWithDefaults instantiates a new IngestionPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestionPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestionPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestionPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IngestionPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *IngestionPipeline) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *IngestionPipeline) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *IngestionPipeline) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *IngestionPipeline) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetGroupId

`func (o *IngestionPipeline) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IngestionPipeline) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IngestionPipeline) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *IngestionPipeline) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *IngestionPipeline) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *IngestionPipeline) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *IngestionPipeline) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *IngestionPipeline) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetName

`func (o *IngestionPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestionPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestionPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IngestionPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSink

`func (o *IngestionPipeline) GetSink() IngestionSink`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *IngestionPipeline) GetSinkOk() (*IngestionSink, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *IngestionPipeline) SetSink(v IngestionSink)`

SetSink sets Sink field to given value.

### HasSink

`func (o *IngestionPipeline) HasSink() bool`

HasSink returns a boolean if a field has been set.

### GetSource

`func (o *IngestionPipeline) GetSource() IngestionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IngestionPipeline) GetSourceOk() (*IngestionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IngestionPipeline) SetSource(v IngestionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *IngestionPipeline) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *IngestionPipeline) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IngestionPipeline) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IngestionPipeline) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IngestionPipeline) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTransformations

`func (o *IngestionPipeline) GetTransformations() []FieldTransformation`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *IngestionPipeline) GetTransformationsOk() (*[]FieldTransformation, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *IngestionPipeline) SetTransformations(v []FieldTransformation)`

SetTransformations sets Transformations field to given value.

### HasTransformations

`func (o *IngestionPipeline) HasTransformations() bool`

HasTransformations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


