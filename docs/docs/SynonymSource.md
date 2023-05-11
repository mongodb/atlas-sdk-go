# SynonymSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** | Human-readable label that identifies the MongoDB collection that stores words and their applicable synonyms. | 

## Methods

### NewSynonymSource

`func NewSynonymSource(collection string, ) *SynonymSource`

NewSynonymSource instantiates a new SynonymSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSynonymSourceWithDefaults

`func NewSynonymSourceWithDefaults() *SynonymSource`

NewSynonymSourceWithDefaults instantiates a new SynonymSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *SynonymSource) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *SynonymSource) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *SynonymSource) SetCollection(v string)`

SetCollection sets Collection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


