# SynonymMappingStatusDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Optional message describing an error. | [optional] 
**Queryable** | Pointer to **bool** | Flag that indicates whether the synonym mapping is queryable on a host. | [optional] 
**Status** | Pointer to **string** | Status that describes this index&#39;s synonym mappings. This status appears only if the index has synonyms defined. | [optional] 

## Methods

### NewSynonymMappingStatusDetail

`func NewSynonymMappingStatusDetail() *SynonymMappingStatusDetail`

NewSynonymMappingStatusDetail instantiates a new SynonymMappingStatusDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSynonymMappingStatusDetailWithDefaults

`func NewSynonymMappingStatusDetailWithDefaults() *SynonymMappingStatusDetail`

NewSynonymMappingStatusDetailWithDefaults instantiates a new SynonymMappingStatusDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SynonymMappingStatusDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SynonymMappingStatusDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SynonymMappingStatusDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SynonymMappingStatusDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.
### GetQueryable

`func (o *SynonymMappingStatusDetail) GetQueryable() bool`

GetQueryable returns the Queryable field if non-nil, zero value otherwise.

### GetQueryableOk

`func (o *SynonymMappingStatusDetail) GetQueryableOk() (*bool, bool)`

GetQueryableOk returns a tuple with the Queryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryable

`func (o *SynonymMappingStatusDetail) SetQueryable(v bool)`

SetQueryable sets Queryable field to given value.

### HasQueryable

`func (o *SynonymMappingStatusDetail) HasQueryable() bool`

HasQueryable returns a boolean if a field has been set.
### GetStatus

`func (o *SynonymMappingStatusDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SynonymMappingStatusDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SynonymMappingStatusDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SynonymMappingStatusDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


