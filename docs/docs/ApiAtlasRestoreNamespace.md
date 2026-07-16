# ApiAtlasRestoreNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceNamespace** | **string** | Namespace requested to restore (e.g. database name or &#x60;database.collection&#x60;). | 
**TargetNamespace** | Pointer to **string** | Requested target namespace for the restored data; if empty, source namespace is used. | [optional] 

## Methods

### NewApiAtlasRestoreNamespace

`func NewApiAtlasRestoreNamespace(sourceNamespace string, ) *ApiAtlasRestoreNamespace`

NewApiAtlasRestoreNamespace instantiates a new ApiAtlasRestoreNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasRestoreNamespaceWithDefaults

`func NewApiAtlasRestoreNamespaceWithDefaults() *ApiAtlasRestoreNamespace`

NewApiAtlasRestoreNamespaceWithDefaults instantiates a new ApiAtlasRestoreNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceNamespace

`func (o *ApiAtlasRestoreNamespace) GetSourceNamespace() string`

GetSourceNamespace returns the SourceNamespace field if non-nil, zero value otherwise.

### GetSourceNamespaceOk

`func (o *ApiAtlasRestoreNamespace) GetSourceNamespaceOk() (*string, bool)`

GetSourceNamespaceOk returns a tuple with the SourceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespace

`func (o *ApiAtlasRestoreNamespace) SetSourceNamespace(v string)`

SetSourceNamespace sets SourceNamespace field to given value.

### GetTargetNamespace

`func (o *ApiAtlasRestoreNamespace) GetTargetNamespace() string`

GetTargetNamespace returns the TargetNamespace field if non-nil, zero value otherwise.

### GetTargetNamespaceOk

`func (o *ApiAtlasRestoreNamespace) GetTargetNamespaceOk() (*string, bool)`

GetTargetNamespaceOk returns a tuple with the TargetNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespace

`func (o *ApiAtlasRestoreNamespace) SetTargetNamespace(v string)`

SetTargetNamespace sets TargetNamespace field to given value.

### HasTargetNamespace

`func (o *ApiAtlasRestoreNamespace) HasTargetNamespace() bool`

HasTargetNamespace returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


