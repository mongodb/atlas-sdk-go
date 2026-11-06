# SchemaAdvisorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Explanation of Performance Advisor availability limitations on the current tier. | [optional] [readonly] 
**Recommendations** | Pointer to [**[]SchemaAdvisorItemRecommendation**](SchemaAdvisorItemRecommendation.md) | List that contains the documents with information about the schema advice that Performance Advisor suggests. | [optional] [readonly] 
**Scale** | Pointer to [**ApiClusterScale**](ApiClusterScale.md) |  | [optional] 

## Methods

### NewSchemaAdvisorResponse

`func NewSchemaAdvisorResponse() *SchemaAdvisorResponse`

NewSchemaAdvisorResponse instantiates a new SchemaAdvisorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaAdvisorResponseWithDefaults

`func NewSchemaAdvisorResponseWithDefaults() *SchemaAdvisorResponse`

NewSchemaAdvisorResponseWithDefaults instantiates a new SchemaAdvisorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SchemaAdvisorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SchemaAdvisorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SchemaAdvisorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SchemaAdvisorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SchemaAdvisorResponse) SetMessageNil()`

SetMessageNil sets Message to an explicit JSON null when marshaled, overriding any value previously set with SetMessage. Calling SetMessage again clears the null override.

### GetRecommendations

`func (o *SchemaAdvisorResponse) GetRecommendations() []SchemaAdvisorItemRecommendation`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *SchemaAdvisorResponse) GetRecommendationsOk() (*[]SchemaAdvisorItemRecommendation, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *SchemaAdvisorResponse) SetRecommendations(v []SchemaAdvisorItemRecommendation)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *SchemaAdvisorResponse) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### SetRecommendationsNil

`func (o *SchemaAdvisorResponse) SetRecommendationsNil()`

SetRecommendationsNil sets Recommendations to an explicit JSON null when marshaled, overriding any value previously set with SetRecommendations. Calling SetRecommendations again clears the null override.

### GetScale

`func (o *SchemaAdvisorResponse) GetScale() ApiClusterScale`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *SchemaAdvisorResponse) GetScaleOk() (*ApiClusterScale, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *SchemaAdvisorResponse) SetScale(v ApiClusterScale)`

SetScale sets Scale field to given value.

### HasScale

`func (o *SchemaAdvisorResponse) HasScale() bool`

HasScale returns a boolean if a field has been set.

### SetScaleNil

`func (o *SchemaAdvisorResponse) SetScaleNil()`

SetScaleNil sets Scale to an explicit JSON null when marshaled, overriding any value previously set with SetScale. Calling SetScale again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


