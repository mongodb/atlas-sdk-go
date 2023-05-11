# ApiAtlasPartitionFieldView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Human-readable label that identifies the field name used to partition data. | 
**Order** | **int** | Sequence in which MongoDB Cloud slices the collection data to create partitions. The resource expresses this sequence starting with zero. | [default to 0]

## Methods

### NewApiAtlasPartitionFieldView

`func NewApiAtlasPartitionFieldView(fieldName string, order int, ) *ApiAtlasPartitionFieldView`

NewApiAtlasPartitionFieldView instantiates a new ApiAtlasPartitionFieldView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasPartitionFieldViewWithDefaults

`func NewApiAtlasPartitionFieldViewWithDefaults() *ApiAtlasPartitionFieldView`

NewApiAtlasPartitionFieldViewWithDefaults instantiates a new ApiAtlasPartitionFieldView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ApiAtlasPartitionFieldView) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ApiAtlasPartitionFieldView) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ApiAtlasPartitionFieldView) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetOrder

`func (o *ApiAtlasPartitionFieldView) GetOrder() int`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApiAtlasPartitionFieldView) GetOrderOk() (*int, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApiAtlasPartitionFieldView) SetOrder(v int)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


