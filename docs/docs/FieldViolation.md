# FieldViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of why the request element is bad. | [optional] 
**Field** | Pointer to **string** | A path that leads to a field in the request body. | [optional] 

## Methods

### NewFieldViolation

`func NewFieldViolation() *FieldViolation`

NewFieldViolation instantiates a new FieldViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldViolationWithDefaults

`func NewFieldViolationWithDefaults() *FieldViolation`

NewFieldViolationWithDefaults instantiates a new FieldViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FieldViolation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldViolation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldViolation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldViolation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetField

`func (o *FieldViolation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldViolation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldViolation) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FieldViolation) HasField() bool`

HasField returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


