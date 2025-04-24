# SearchIndexDefinitionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The time at which this index definition was created. | [optional] 
**Version** | Pointer to **int64** | The version number associated with this index definition when it was created. | [optional] 

## Methods

### NewSearchIndexDefinitionVersion

`func NewSearchIndexDefinitionVersion() *SearchIndexDefinitionVersion`

NewSearchIndexDefinitionVersion instantiates a new SearchIndexDefinitionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexDefinitionVersionWithDefaults

`func NewSearchIndexDefinitionVersionWithDefaults() *SearchIndexDefinitionVersion`

NewSearchIndexDefinitionVersionWithDefaults instantiates a new SearchIndexDefinitionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SearchIndexDefinitionVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SearchIndexDefinitionVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SearchIndexDefinitionVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SearchIndexDefinitionVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetVersion

`func (o *SearchIndexDefinitionVersion) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SearchIndexDefinitionVersion) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SearchIndexDefinitionVersion) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SearchIndexDefinitionVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


