# FTSAnalyzersCharFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoredTags** | Pointer to **[]string** | The HTML tags that you want to exclude from filtering. | [optional] 
**Type** | Pointer to **string** | Human-readable label that identifies this character filter type. | [optional] 
**Mappings** | Pointer to [**FTSAnalyzersCharFiltersInnerMappings**](FTSAnalyzersCharFiltersInnerMappings.md) |  | [optional] 

## Methods

### NewFTSAnalyzersCharFiltersInner

`func NewFTSAnalyzersCharFiltersInner() *FTSAnalyzersCharFiltersInner`

NewFTSAnalyzersCharFiltersInner instantiates a new FTSAnalyzersCharFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSAnalyzersCharFiltersInnerWithDefaults

`func NewFTSAnalyzersCharFiltersInnerWithDefaults() *FTSAnalyzersCharFiltersInner`

NewFTSAnalyzersCharFiltersInnerWithDefaults instantiates a new FTSAnalyzersCharFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoredTags

`func (o *FTSAnalyzersCharFiltersInner) GetIgnoredTags() []string`

GetIgnoredTags returns the IgnoredTags field if non-nil, zero value otherwise.

### GetIgnoredTagsOk

`func (o *FTSAnalyzersCharFiltersInner) GetIgnoredTagsOk() (*[]string, bool)`

GetIgnoredTagsOk returns a tuple with the IgnoredTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredTags

`func (o *FTSAnalyzersCharFiltersInner) SetIgnoredTags(v []string)`

SetIgnoredTags sets IgnoredTags field to given value.

### HasIgnoredTags

`func (o *FTSAnalyzersCharFiltersInner) HasIgnoredTags() bool`

HasIgnoredTags returns a boolean if a field has been set.

### GetType

`func (o *FTSAnalyzersCharFiltersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FTSAnalyzersCharFiltersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FTSAnalyzersCharFiltersInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FTSAnalyzersCharFiltersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMappings

`func (o *FTSAnalyzersCharFiltersInner) GetMappings() FTSAnalyzersCharFiltersInnerMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *FTSAnalyzersCharFiltersInner) GetMappingsOk() (*FTSAnalyzersCharFiltersInnerMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *FTSAnalyzersCharFiltersInner) SetMappings(v FTSAnalyzersCharFiltersInnerMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *FTSAnalyzersCharFiltersInner) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


