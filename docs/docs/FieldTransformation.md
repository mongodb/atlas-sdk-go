# FieldTransformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Key in the document. | [optional] 
**Type** | Pointer to **string** | Type of transformation applied during the export of the namespace in a Data Lake Pipeline. | [optional] 

## Methods

### NewFieldTransformation

`func NewFieldTransformation() *FieldTransformation`

NewFieldTransformation instantiates a new FieldTransformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldTransformationWithDefaults

`func NewFieldTransformationWithDefaults() *FieldTransformation`

NewFieldTransformationWithDefaults instantiates a new FieldTransformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldTransformation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldTransformation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldTransformation) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FieldTransformation) HasField() bool`

HasField returns a boolean if a field has been set.

### GetType

`func (o *FieldTransformation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldTransformation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldTransformation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FieldTransformation) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


