# SchemaAdvisorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recommendations** | Pointer to [**[]SchemaAdvisorItemRecommendation**](SchemaAdvisorItemRecommendation.md) | List that contains the documents with information about the schema advice that Performance Advisor suggests. | [optional] [readonly] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


