# TokenFilterlength

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | Pointer to **int** | Number that specifies the maximum length of a token. Value must be greater than or equal to **min**. | [optional] [default to 255]
**Min** | Pointer to **int** | Number that specifies the minimum length of a token. This value must be less than or equal to **max**. | [optional] [default to 0]
**Type** | **string** | Human-readable label that identifies this token filter type. | 

## Methods

### NewTokenFilterlength

`func NewTokenFilterlength(type_ string, ) *TokenFilterlength`

NewTokenFilterlength instantiates a new TokenFilterlength object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenFilterlengthWithDefaults

`func NewTokenFilterlengthWithDefaults() *TokenFilterlength`

NewTokenFilterlengthWithDefaults instantiates a new TokenFilterlength object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *TokenFilterlength) GetMax() int`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *TokenFilterlength) GetMaxOk() (*int, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *TokenFilterlength) SetMax(v int)`

SetMax sets Max field to given value.

### HasMax

`func (o *TokenFilterlength) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *TokenFilterlength) GetMin() int`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *TokenFilterlength) GetMinOk() (*int, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *TokenFilterlength) SetMin(v int)`

SetMin sets Min field to given value.

### HasMin

`func (o *TokenFilterlength) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetType

`func (o *TokenFilterlength) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenFilterlength) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenFilterlength) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


