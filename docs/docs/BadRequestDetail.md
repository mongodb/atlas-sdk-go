# BadRequestDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]FieldViolation**](FieldViolation.md) | Describes all violations in a client request. | [optional] 

## Methods

### NewBadRequestDetail

`func NewBadRequestDetail() *BadRequestDetail`

NewBadRequestDetail instantiates a new BadRequestDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestDetailWithDefaults

`func NewBadRequestDetailWithDefaults() *BadRequestDetail`

NewBadRequestDetailWithDefaults instantiates a new BadRequestDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *BadRequestDetail) GetFields() []FieldViolation`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BadRequestDetail) GetFieldsOk() (*[]FieldViolation, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BadRequestDetail) SetFields(v []FieldViolation)`

SetFields sets Fields field to given value.

### HasFields

`func (o *BadRequestDetail) HasFields() bool`

HasFields returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


