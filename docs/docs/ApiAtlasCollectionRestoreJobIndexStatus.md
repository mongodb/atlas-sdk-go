# ApiAtlasCollectionRestoreJobIndexStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedCollectionCount** | Pointer to **int** | Number of collections that failed to build indexes. | [optional] 
**State** | Pointer to **string** | Index build state indicating the status of index creation during or after a restore operation. | [optional] 

## Methods

### NewApiAtlasCollectionRestoreJobIndexStatus

`func NewApiAtlasCollectionRestoreJobIndexStatus() *ApiAtlasCollectionRestoreJobIndexStatus`

NewApiAtlasCollectionRestoreJobIndexStatus instantiates a new ApiAtlasCollectionRestoreJobIndexStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCollectionRestoreJobIndexStatusWithDefaults

`func NewApiAtlasCollectionRestoreJobIndexStatusWithDefaults() *ApiAtlasCollectionRestoreJobIndexStatus`

NewApiAtlasCollectionRestoreJobIndexStatusWithDefaults instantiates a new ApiAtlasCollectionRestoreJobIndexStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedCollectionCount

`func (o *ApiAtlasCollectionRestoreJobIndexStatus) GetFailedCollectionCount() int`

GetFailedCollectionCount returns the FailedCollectionCount field if non-nil, zero value otherwise.

### GetFailedCollectionCountOk

`func (o *ApiAtlasCollectionRestoreJobIndexStatus) GetFailedCollectionCountOk() (*int, bool)`

GetFailedCollectionCountOk returns a tuple with the FailedCollectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCollectionCount

`func (o *ApiAtlasCollectionRestoreJobIndexStatus) SetFailedCollectionCount(v int)`

SetFailedCollectionCount sets FailedCollectionCount field to given value.

### HasFailedCollectionCount

`func (o *ApiAtlasCollectionRestoreJobIndexStatus) HasFailedCollectionCount() bool`

HasFailedCollectionCount returns a boolean if a field has been set.
### GetState

`func (o *ApiAtlasCollectionRestoreJobIndexStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiAtlasCollectionRestoreJobIndexStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiAtlasCollectionRestoreJobIndexStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiAtlasCollectionRestoreJobIndexStatus) HasState() bool`

HasState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


