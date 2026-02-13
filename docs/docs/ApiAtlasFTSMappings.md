# ApiAtlasFTSMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dynamic** | Pointer to **bool** | Flag that indicates whether the index uses dynamic or static mappings. Required if &#x60;mappings.fields&#x60; is omitted. | [optional] [default to false]
**Fields** | Pointer to [**any**](interface{}.md) | One or more field specifications for the Atlas Search index. Required if &#x60;mappings.dynamic&#x60; is omitted or set to &#x60;false&#x60;. | [optional] 

## Methods

### NewApiAtlasFTSMappings

`func NewApiAtlasFTSMappings() *ApiAtlasFTSMappings`

NewApiAtlasFTSMappings instantiates a new ApiAtlasFTSMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasFTSMappingsWithDefaults

`func NewApiAtlasFTSMappingsWithDefaults() *ApiAtlasFTSMappings`

NewApiAtlasFTSMappingsWithDefaults instantiates a new ApiAtlasFTSMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamic

`func (o *ApiAtlasFTSMappings) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *ApiAtlasFTSMappings) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *ApiAtlasFTSMappings) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *ApiAtlasFTSMappings) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.
### GetFields

`func (o *ApiAtlasFTSMappings) GetFields() any`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ApiAtlasFTSMappings) GetFieldsOk() (*any, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ApiAtlasFTSMappings) SetFields(v any)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ApiAtlasFTSMappings) HasFields() bool`

HasFields returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


