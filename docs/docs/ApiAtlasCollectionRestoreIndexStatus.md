# ApiAtlasCollectionRestoreIndexStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | Error message if index build failed. | [optional] 
**FailedIndexes** | Pointer to **[]map[string]any** | List of index specifications that failed to build (up to 64 items). | [optional] 
**State** | Pointer to **string** | Index build state indicating the status of index creation during or after a restore operation. | [optional] 

## Methods

### NewApiAtlasCollectionRestoreIndexStatus

`func NewApiAtlasCollectionRestoreIndexStatus() *ApiAtlasCollectionRestoreIndexStatus`

NewApiAtlasCollectionRestoreIndexStatus instantiates a new ApiAtlasCollectionRestoreIndexStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCollectionRestoreIndexStatusWithDefaults

`func NewApiAtlasCollectionRestoreIndexStatusWithDefaults() *ApiAtlasCollectionRestoreIndexStatus`

NewApiAtlasCollectionRestoreIndexStatusWithDefaults instantiates a new ApiAtlasCollectionRestoreIndexStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *ApiAtlasCollectionRestoreIndexStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasCollectionRestoreIndexStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasCollectionRestoreIndexStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasCollectionRestoreIndexStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetFailedIndexes

`func (o *ApiAtlasCollectionRestoreIndexStatus) GetFailedIndexes() []map[string]any`

GetFailedIndexes returns the FailedIndexes field if non-nil, zero value otherwise.

### GetFailedIndexesOk

`func (o *ApiAtlasCollectionRestoreIndexStatus) GetFailedIndexesOk() (*[]map[string]any, bool)`

GetFailedIndexesOk returns a tuple with the FailedIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedIndexes

`func (o *ApiAtlasCollectionRestoreIndexStatus) SetFailedIndexes(v []map[string]any)`

SetFailedIndexes sets FailedIndexes field to given value.

### HasFailedIndexes

`func (o *ApiAtlasCollectionRestoreIndexStatus) HasFailedIndexes() bool`

HasFailedIndexes returns a boolean if a field has been set.
### GetState

`func (o *ApiAtlasCollectionRestoreIndexStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiAtlasCollectionRestoreIndexStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiAtlasCollectionRestoreIndexStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiAtlasCollectionRestoreIndexStatus) HasState() bool`

HasState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


