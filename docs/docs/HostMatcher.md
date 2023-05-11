# HostMatcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to [**HostMatcherField**](HostMatcherField.md) |  | [optional] 
**Operator** | Pointer to **string** | Comparison operator to apply when checking the current metric value against **matcher[n].value**. | [optional] 
**Value** | Pointer to [**MatcherHostType**](MatcherHostType.md) |  | [optional] 

## Methods

### NewHostMatcher

`func NewHostMatcher() *HostMatcher`

NewHostMatcher instantiates a new HostMatcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMatcherWithDefaults

`func NewHostMatcherWithDefaults() *HostMatcher`

NewHostMatcherWithDefaults instantiates a new HostMatcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *HostMatcher) GetFieldName() HostMatcherField`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *HostMatcher) GetFieldNameOk() (*HostMatcherField, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *HostMatcher) SetFieldName(v HostMatcherField)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *HostMatcher) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetOperator

`func (o *HostMatcher) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *HostMatcher) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *HostMatcher) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *HostMatcher) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *HostMatcher) GetValue() MatcherHostType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HostMatcher) GetValueOk() (*MatcherHostType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HostMatcher) SetValue(v MatcherHostType)`

SetValue sets Value field to given value.

### HasValue

`func (o *HostMatcher) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


