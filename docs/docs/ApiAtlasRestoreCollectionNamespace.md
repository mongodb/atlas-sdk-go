# ApiAtlasRestoreCollectionNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceNamespace** | **string** | Collection requested to restore, as &#x60;database.collection&#x60;. | 
**TargetNamespace** | Pointer to **string** | Requested target collection as &#x60;database.collection&#x60;; if empty, source namespace is used. | [optional] 

## Methods

### NewApiAtlasRestoreCollectionNamespace

`func NewApiAtlasRestoreCollectionNamespace(sourceNamespace string, ) *ApiAtlasRestoreCollectionNamespace`

NewApiAtlasRestoreCollectionNamespace instantiates a new ApiAtlasRestoreCollectionNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasRestoreCollectionNamespaceWithDefaults

`func NewApiAtlasRestoreCollectionNamespaceWithDefaults() *ApiAtlasRestoreCollectionNamespace`

NewApiAtlasRestoreCollectionNamespaceWithDefaults instantiates a new ApiAtlasRestoreCollectionNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceNamespace

`func (o *ApiAtlasRestoreCollectionNamespace) GetSourceNamespace() string`

GetSourceNamespace returns the SourceNamespace field if non-nil, zero value otherwise.

### GetSourceNamespaceOk

`func (o *ApiAtlasRestoreCollectionNamespace) GetSourceNamespaceOk() (*string, bool)`

GetSourceNamespaceOk returns a tuple with the SourceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespace

`func (o *ApiAtlasRestoreCollectionNamespace) SetSourceNamespace(v string)`

SetSourceNamespace sets SourceNamespace field to given value.

### GetTargetNamespace

`func (o *ApiAtlasRestoreCollectionNamespace) GetTargetNamespace() string`

GetTargetNamespace returns the TargetNamespace field if non-nil, zero value otherwise.

### GetTargetNamespaceOk

`func (o *ApiAtlasRestoreCollectionNamespace) GetTargetNamespaceOk() (*string, bool)`

GetTargetNamespaceOk returns a tuple with the TargetNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespace

`func (o *ApiAtlasRestoreCollectionNamespace) SetTargetNamespace(v string)`

SetTargetNamespace sets TargetNamespace field to given value.

### HasTargetNamespace

`func (o *ApiAtlasRestoreCollectionNamespace) HasTargetNamespace() bool`

HasTargetNamespace returns a boolean if a field has been set.

### SetTargetNamespaceNil

`func (o *ApiAtlasRestoreCollectionNamespace) SetTargetNamespaceNil()`

SetTargetNamespaceNil sets TargetNamespace to an explicit JSON null when marshaled, overriding any value previously set with SetTargetNamespace. Calling SetTargetNamespace again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


