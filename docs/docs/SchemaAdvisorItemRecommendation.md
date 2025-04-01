# SchemaAdvisorItemRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedNamespaces** | Pointer to [**[]SchemaAdvisorNamespaceTriggers**](SchemaAdvisorNamespaceTriggers.md) | List that contains the namespaces and information on why those namespaces triggered the recommendation. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the specified recommendation. | [optional] [readonly] 
**Recommendation** | Pointer to **string** | Type of recommendation. | [optional] [readonly] 

## Methods

### NewSchemaAdvisorItemRecommendation

`func NewSchemaAdvisorItemRecommendation() *SchemaAdvisorItemRecommendation`

NewSchemaAdvisorItemRecommendation instantiates a new SchemaAdvisorItemRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaAdvisorItemRecommendationWithDefaults

`func NewSchemaAdvisorItemRecommendationWithDefaults() *SchemaAdvisorItemRecommendation`

NewSchemaAdvisorItemRecommendationWithDefaults instantiates a new SchemaAdvisorItemRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedNamespaces

`func (o *SchemaAdvisorItemRecommendation) GetAffectedNamespaces() []SchemaAdvisorNamespaceTriggers`

GetAffectedNamespaces returns the AffectedNamespaces field if non-nil, zero value otherwise.

### GetAffectedNamespacesOk

`func (o *SchemaAdvisorItemRecommendation) GetAffectedNamespacesOk() (*[]SchemaAdvisorNamespaceTriggers, bool)`

GetAffectedNamespacesOk returns a tuple with the AffectedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedNamespaces

`func (o *SchemaAdvisorItemRecommendation) SetAffectedNamespaces(v []SchemaAdvisorNamespaceTriggers)`

SetAffectedNamespaces sets AffectedNamespaces field to given value.

### HasAffectedNamespaces

`func (o *SchemaAdvisorItemRecommendation) HasAffectedNamespaces() bool`

HasAffectedNamespaces returns a boolean if a field has been set.
### GetDescription

`func (o *SchemaAdvisorItemRecommendation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaAdvisorItemRecommendation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaAdvisorItemRecommendation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaAdvisorItemRecommendation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetRecommendation

`func (o *SchemaAdvisorItemRecommendation) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *SchemaAdvisorItemRecommendation) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *SchemaAdvisorItemRecommendation) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *SchemaAdvisorItemRecommendation) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


