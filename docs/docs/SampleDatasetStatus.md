# SampleDatasetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies this sample dataset. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster into which you loaded the sample dataset. | [optional] [readonly] 
**CompleteDate** | Pointer to **time.Time** | Date and time when the sample dataset load job completed. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**CreateDate** | Pointer to **time.Time** | Date and time when you started the sample dataset load job. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Details of the error returned when MongoDB Cloud loads the sample dataset. This endpoint returns null if state has a value other than FAILED. | [optional] [readonly] 
**State** | Pointer to **string** | Status of the sample dataset load job. | [optional] [readonly] 

## Methods

### NewSampleDatasetStatus

`func NewSampleDatasetStatus() *SampleDatasetStatus`

NewSampleDatasetStatus instantiates a new SampleDatasetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleDatasetStatusWithDefaults

`func NewSampleDatasetStatusWithDefaults() *SampleDatasetStatus`

NewSampleDatasetStatusWithDefaults instantiates a new SampleDatasetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SampleDatasetStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SampleDatasetStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SampleDatasetStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SampleDatasetStatus) HasId() bool`

HasId returns a boolean if a field has been set.
### GetClusterName

`func (o *SampleDatasetStatus) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *SampleDatasetStatus) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *SampleDatasetStatus) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *SampleDatasetStatus) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetCompleteDate

`func (o *SampleDatasetStatus) GetCompleteDate() time.Time`

GetCompleteDate returns the CompleteDate field if non-nil, zero value otherwise.

### GetCompleteDateOk

`func (o *SampleDatasetStatus) GetCompleteDateOk() (*time.Time, bool)`

GetCompleteDateOk returns a tuple with the CompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteDate

`func (o *SampleDatasetStatus) SetCompleteDate(v time.Time)`

SetCompleteDate sets CompleteDate field to given value.

### HasCompleteDate

`func (o *SampleDatasetStatus) HasCompleteDate() bool`

HasCompleteDate returns a boolean if a field has been set.
### GetCreateDate

`func (o *SampleDatasetStatus) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *SampleDatasetStatus) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *SampleDatasetStatus) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *SampleDatasetStatus) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.
### GetErrorMessage

`func (o *SampleDatasetStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SampleDatasetStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SampleDatasetStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SampleDatasetStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetState

`func (o *SampleDatasetStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SampleDatasetStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SampleDatasetStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SampleDatasetStatus) HasState() bool`

HasState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


