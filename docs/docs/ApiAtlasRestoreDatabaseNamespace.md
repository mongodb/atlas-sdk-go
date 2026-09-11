# ApiAtlasRestoreDatabaseNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceNamespace** | **string** | Database name requested to restore. | 
**TargetNamespace** | Pointer to **string** | Requested target database name; if empty, source database name is used. | [optional] 

## Methods

### NewApiAtlasRestoreDatabaseNamespace

`func NewApiAtlasRestoreDatabaseNamespace(sourceNamespace string, ) *ApiAtlasRestoreDatabaseNamespace`

NewApiAtlasRestoreDatabaseNamespace instantiates a new ApiAtlasRestoreDatabaseNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasRestoreDatabaseNamespaceWithDefaults

`func NewApiAtlasRestoreDatabaseNamespaceWithDefaults() *ApiAtlasRestoreDatabaseNamespace`

NewApiAtlasRestoreDatabaseNamespaceWithDefaults instantiates a new ApiAtlasRestoreDatabaseNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceNamespace

`func (o *ApiAtlasRestoreDatabaseNamespace) GetSourceNamespace() string`

GetSourceNamespace returns the SourceNamespace field if non-nil, zero value otherwise.

### GetSourceNamespaceOk

`func (o *ApiAtlasRestoreDatabaseNamespace) GetSourceNamespaceOk() (*string, bool)`

GetSourceNamespaceOk returns a tuple with the SourceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespace

`func (o *ApiAtlasRestoreDatabaseNamespace) SetSourceNamespace(v string)`

SetSourceNamespace sets SourceNamespace field to given value.

### GetTargetNamespace

`func (o *ApiAtlasRestoreDatabaseNamespace) GetTargetNamespace() string`

GetTargetNamespace returns the TargetNamespace field if non-nil, zero value otherwise.

### GetTargetNamespaceOk

`func (o *ApiAtlasRestoreDatabaseNamespace) GetTargetNamespaceOk() (*string, bool)`

GetTargetNamespaceOk returns a tuple with the TargetNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespace

`func (o *ApiAtlasRestoreDatabaseNamespace) SetTargetNamespace(v string)`

SetTargetNamespace sets TargetNamespace field to given value.

### HasTargetNamespace

`func (o *ApiAtlasRestoreDatabaseNamespace) HasTargetNamespace() bool`

HasTargetNamespace returns a boolean if a field has been set.

### SetTargetNamespaceNil

`func (o *ApiAtlasRestoreDatabaseNamespace) SetTargetNamespaceNil()`

SetTargetNamespaceNil sets TargetNamespace to an explicit JSON null when marshaled, overriding any value previously set with SetTargetNamespace. Calling SetTargetNamespace again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


