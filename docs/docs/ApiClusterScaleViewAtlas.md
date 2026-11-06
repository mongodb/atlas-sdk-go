# ApiClusterScaleViewAtlas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocsUrl** | Pointer to **string** | Documentation page describing feature availability by cluster tier. | [optional] [readonly] 
**MinimumTier** | Pointer to **string** | Lowest cluster tier that supports the requested feature. | [optional] [readonly] 
**ScaleUrl** | Pointer to **string** | Link to the Atlas UI page that scales the cluster, with the minimum supported tier pre-selected. | [optional] [readonly] 

## Methods

### NewApiClusterScaleViewAtlas

`func NewApiClusterScaleViewAtlas() *ApiClusterScaleViewAtlas`

NewApiClusterScaleViewAtlas instantiates a new ApiClusterScaleViewAtlas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClusterScaleViewAtlasWithDefaults

`func NewApiClusterScaleViewAtlasWithDefaults() *ApiClusterScaleViewAtlas`

NewApiClusterScaleViewAtlasWithDefaults instantiates a new ApiClusterScaleViewAtlas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocsUrl

`func (o *ApiClusterScaleViewAtlas) GetDocsUrl() string`

GetDocsUrl returns the DocsUrl field if non-nil, zero value otherwise.

### GetDocsUrlOk

`func (o *ApiClusterScaleViewAtlas) GetDocsUrlOk() (*string, bool)`

GetDocsUrlOk returns a tuple with the DocsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsUrl

`func (o *ApiClusterScaleViewAtlas) SetDocsUrl(v string)`

SetDocsUrl sets DocsUrl field to given value.

### HasDocsUrl

`func (o *ApiClusterScaleViewAtlas) HasDocsUrl() bool`

HasDocsUrl returns a boolean if a field has been set.

### SetDocsUrlNil

`func (o *ApiClusterScaleViewAtlas) SetDocsUrlNil()`

SetDocsUrlNil sets DocsUrl to an explicit JSON null when marshaled, overriding any value previously set with SetDocsUrl. Calling SetDocsUrl again clears the null override.

### GetMinimumTier

`func (o *ApiClusterScaleViewAtlas) GetMinimumTier() string`

GetMinimumTier returns the MinimumTier field if non-nil, zero value otherwise.

### GetMinimumTierOk

`func (o *ApiClusterScaleViewAtlas) GetMinimumTierOk() (*string, bool)`

GetMinimumTierOk returns a tuple with the MinimumTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTier

`func (o *ApiClusterScaleViewAtlas) SetMinimumTier(v string)`

SetMinimumTier sets MinimumTier field to given value.

### HasMinimumTier

`func (o *ApiClusterScaleViewAtlas) HasMinimumTier() bool`

HasMinimumTier returns a boolean if a field has been set.

### SetMinimumTierNil

`func (o *ApiClusterScaleViewAtlas) SetMinimumTierNil()`

SetMinimumTierNil sets MinimumTier to an explicit JSON null when marshaled, overriding any value previously set with SetMinimumTier. Calling SetMinimumTier again clears the null override.

### GetScaleUrl

`func (o *ApiClusterScaleViewAtlas) GetScaleUrl() string`

GetScaleUrl returns the ScaleUrl field if non-nil, zero value otherwise.

### GetScaleUrlOk

`func (o *ApiClusterScaleViewAtlas) GetScaleUrlOk() (*string, bool)`

GetScaleUrlOk returns a tuple with the ScaleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUrl

`func (o *ApiClusterScaleViewAtlas) SetScaleUrl(v string)`

SetScaleUrl sets ScaleUrl field to given value.

### HasScaleUrl

`func (o *ApiClusterScaleViewAtlas) HasScaleUrl() bool`

HasScaleUrl returns a boolean if a field has been set.

### SetScaleUrlNil

`func (o *ApiClusterScaleViewAtlas) SetScaleUrlNil()`

SetScaleUrlNil sets ScaleUrl to an explicit JSON null when marshaled, overriding any value previously set with SetScaleUrl. Calling SetScaleUrl again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


