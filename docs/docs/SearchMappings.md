# SearchMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dynamic** | Pointer to **bool** | Flag that indicates whether the index uses dynamic or static mappings. Required if **mappings.fields** is omitted. | [optional] 
**Fields** | Pointer to [**map[string]any**](interface{}.md) | One or more field specifications for the Atlas Search index. Required if **mappings.dynamic** is omitted or set to **false**. | [optional] 

## Methods

### NewSearchMappings

`func NewSearchMappings() *SearchMappings`

NewSearchMappings instantiates a new SearchMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMappingsWithDefaults

`func NewSearchMappingsWithDefaults() *SearchMappings`

NewSearchMappingsWithDefaults instantiates a new SearchMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamic

`func (o *SearchMappings) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *SearchMappings) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *SearchMappings) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *SearchMappings) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.
### GetFields

`func (o *SearchMappings) GetFields() map[string]any`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchMappings) GetFieldsOk() (*map[string]any, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchMappings) SetFields(v map[string]any)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchMappings) HasFields() bool`

HasFields returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


