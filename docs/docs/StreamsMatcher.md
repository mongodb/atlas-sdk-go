# StreamsMatcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Name of the parameter in the target object that MongoDB Cloud checks. The parameter must match all rules for MongoDB Cloud to check for alert configurations. | 
**Operator** | **string** | Comparison operator to apply when checking the current metric value against **matcher[n].value**. | 
**Value** | **string** | Value to match or exceed using the specified **matchers.operator**. | 

## Methods

### NewStreamsMatcher

`func NewStreamsMatcher(fieldName string, operator string, value string, ) *StreamsMatcher`

NewStreamsMatcher instantiates a new StreamsMatcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsMatcherWithDefaults

`func NewStreamsMatcherWithDefaults() *StreamsMatcher`

NewStreamsMatcherWithDefaults instantiates a new StreamsMatcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *StreamsMatcher) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *StreamsMatcher) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *StreamsMatcher) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### GetOperator

`func (o *StreamsMatcher) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StreamsMatcher) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StreamsMatcher) SetOperator(v string)`

SetOperator sets Operator field to given value.

### GetValue

`func (o *StreamsMatcher) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StreamsMatcher) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StreamsMatcher) SetValue(v string)`

SetValue sets Value field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


