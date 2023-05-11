# ExportStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportedCollections** | Pointer to **int** | Number of collections on the replica set that MongoDB Cloud exported. | [optional] [readonly] 
**TotalCollections** | Pointer to **int** | Total number of collections on the replica set to export. | [optional] [readonly] 

## Methods

### NewExportStatus

`func NewExportStatus() *ExportStatus`

NewExportStatus instantiates a new ExportStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportStatusWithDefaults

`func NewExportStatusWithDefaults() *ExportStatus`

NewExportStatusWithDefaults instantiates a new ExportStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportedCollections

`func (o *ExportStatus) GetExportedCollections() int`

GetExportedCollections returns the ExportedCollections field if non-nil, zero value otherwise.

### GetExportedCollectionsOk

`func (o *ExportStatus) GetExportedCollectionsOk() (*int, bool)`

GetExportedCollectionsOk returns a tuple with the ExportedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedCollections

`func (o *ExportStatus) SetExportedCollections(v int)`

SetExportedCollections sets ExportedCollections field to given value.

### HasExportedCollections

`func (o *ExportStatus) HasExportedCollections() bool`

HasExportedCollections returns a boolean if a field has been set.

### GetTotalCollections

`func (o *ExportStatus) GetTotalCollections() int`

GetTotalCollections returns the TotalCollections field if non-nil, zero value otherwise.

### GetTotalCollectionsOk

`func (o *ExportStatus) GetTotalCollectionsOk() (*int, bool)`

GetTotalCollectionsOk returns a tuple with the TotalCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollections

`func (o *ExportStatus) SetTotalCollections(v int)`

SetTotalCollections sets TotalCollections field to given value.

### HasTotalCollections

`func (o *ExportStatus) HasTotalCollections() bool`

HasTotalCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


