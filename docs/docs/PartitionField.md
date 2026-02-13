# PartitionField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Human-readable label that identifies the parameter that MongoDB Cloud uses to partition data. To specify a nested parameter, use the dot notation. | 
**FieldType** | Pointer to **string** | Data type of the parameter that that MongoDB Cloud uses to partition data. Partition parameters of type UUID must be of binary subtype 4. MongoDB Cloud skips partition parameters of type UUID with subtype 3. | [optional] [readonly] 
**Order** | **int** | Sequence in which MongoDB Cloud slices the collection data to create partitions. The resource expresses this sequence starting with zero. The value of the &#x60;criteria.dateField&#x60; parameter defaults as the first item in the partition sequence. | [default to 0]

## Methods

### NewPartitionField

`func NewPartitionField(fieldName string, order int, ) *PartitionField`

NewPartitionField instantiates a new PartitionField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionFieldWithDefaults

`func NewPartitionFieldWithDefaults() *PartitionField`

NewPartitionFieldWithDefaults instantiates a new PartitionField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *PartitionField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *PartitionField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *PartitionField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### GetFieldType

`func (o *PartitionField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PartitionField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PartitionField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *PartitionField) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.
### GetOrder

`func (o *PartitionField) GetOrder() int`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PartitionField) GetOrderOk() (*int, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PartitionField) SetOrder(v int)`

SetOrder sets Order field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


