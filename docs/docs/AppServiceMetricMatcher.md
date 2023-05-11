# AppServiceMetricMatcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to [**AppServiceMetricMatcherField**](AppServiceMetricMatcherField.md) |  | [optional] 
**Operator** | Pointer to **string** | Comparison operator to apply when checking the current metric value against **matcher[n].value**. | [optional] 
**Value** | Pointer to **string** | Value to match or exceed using the specified **matchers.operator**. | [optional] 

## Methods

### NewAppServiceMetricMatcher

`func NewAppServiceMetricMatcher() *AppServiceMetricMatcher`

NewAppServiceMetricMatcher instantiates a new AppServiceMetricMatcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceMetricMatcherWithDefaults

`func NewAppServiceMetricMatcherWithDefaults() *AppServiceMetricMatcher`

NewAppServiceMetricMatcherWithDefaults instantiates a new AppServiceMetricMatcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *AppServiceMetricMatcher) GetFieldName() AppServiceMetricMatcherField`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *AppServiceMetricMatcher) GetFieldNameOk() (*AppServiceMetricMatcherField, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *AppServiceMetricMatcher) SetFieldName(v AppServiceMetricMatcherField)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *AppServiceMetricMatcher) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetOperator

`func (o *AppServiceMetricMatcher) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppServiceMetricMatcher) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppServiceMetricMatcher) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AppServiceMetricMatcher) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *AppServiceMetricMatcher) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceMetricMatcher) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceMetricMatcher) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppServiceMetricMatcher) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


