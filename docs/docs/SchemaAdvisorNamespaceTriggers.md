# SchemaAdvisorNamespaceTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Namespace of the affected collection. Will be null for REDUCE_NUMBER_OF_NAMESPACE recommendation. | [optional] [readonly] 
**Triggers** | Pointer to [**[]SchemaAdvisorTriggerDetails**](SchemaAdvisorTriggerDetails.md) | List of triggers that specify why the collection activated the recommendation. | [optional] [readonly] 

## Methods

### NewSchemaAdvisorNamespaceTriggers

`func NewSchemaAdvisorNamespaceTriggers() *SchemaAdvisorNamespaceTriggers`

NewSchemaAdvisorNamespaceTriggers instantiates a new SchemaAdvisorNamespaceTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaAdvisorNamespaceTriggersWithDefaults

`func NewSchemaAdvisorNamespaceTriggersWithDefaults() *SchemaAdvisorNamespaceTriggers`

NewSchemaAdvisorNamespaceTriggersWithDefaults instantiates a new SchemaAdvisorNamespaceTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *SchemaAdvisorNamespaceTriggers) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SchemaAdvisorNamespaceTriggers) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SchemaAdvisorNamespaceTriggers) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SchemaAdvisorNamespaceTriggers) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.
### GetTriggers

`func (o *SchemaAdvisorNamespaceTriggers) GetTriggers() []SchemaAdvisorTriggerDetails`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *SchemaAdvisorNamespaceTriggers) GetTriggersOk() (*[]SchemaAdvisorTriggerDetails, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *SchemaAdvisorNamespaceTriggers) SetTriggers(v []SchemaAdvisorTriggerDetails)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *SchemaAdvisorNamespaceTriggers) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


