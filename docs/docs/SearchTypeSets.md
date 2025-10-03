# SearchTypeSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Label that identifies the type set name. Each **typeSets.name** must be unique within the same index definition. | 
**Types** | Pointer to [**[]any**](any.md) | List of types associated with the type set. Each type definition must include a \&quot;type\&quot; field specifying the search field type (\&quot;autocomplete\&quot;, \&quot;boolean\&quot;, \&quot;date\&quot;, \&quot;geo\&quot;, \&quot;number\&quot;, \&quot;objectId\&quot;, \&quot;string\&quot;, \&quot;token\&quot;, or \&quot;uuid\&quot;) and may include additional configuration properties specific to that type. | [optional] 

## Methods

### NewSearchTypeSets

`func NewSearchTypeSets(name string, ) *SearchTypeSets`

NewSearchTypeSets instantiates a new SearchTypeSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTypeSetsWithDefaults

`func NewSearchTypeSetsWithDefaults() *SearchTypeSets`

NewSearchTypeSetsWithDefaults instantiates a new SearchTypeSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SearchTypeSets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchTypeSets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchTypeSets) SetName(v string)`

SetName sets Name field to given value.

### GetTypes

`func (o *SearchTypeSets) GetTypes() []any`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchTypeSets) GetTypesOk() (*[]any, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchTypeSets) SetTypes(v []any)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchTypeSets) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


