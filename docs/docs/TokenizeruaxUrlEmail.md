# TokenizeruaxUrlEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTokenLength** | Pointer to **int** | Maximum number of characters in a single token. Tokens greater than this length are split at this length into multiple tokens. | [optional] [default to 255]
**Type** | **string** | Human-readable label that identifies this tokenizer type. | 

## Methods

### NewTokenizeruaxUrlEmail

`func NewTokenizeruaxUrlEmail(type_ string, ) *TokenizeruaxUrlEmail`

NewTokenizeruaxUrlEmail instantiates a new TokenizeruaxUrlEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizeruaxUrlEmailWithDefaults

`func NewTokenizeruaxUrlEmailWithDefaults() *TokenizeruaxUrlEmail`

NewTokenizeruaxUrlEmailWithDefaults instantiates a new TokenizeruaxUrlEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTokenLength

`func (o *TokenizeruaxUrlEmail) GetMaxTokenLength() int`

GetMaxTokenLength returns the MaxTokenLength field if non-nil, zero value otherwise.

### GetMaxTokenLengthOk

`func (o *TokenizeruaxUrlEmail) GetMaxTokenLengthOk() (*int, bool)`

GetMaxTokenLengthOk returns a tuple with the MaxTokenLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenLength

`func (o *TokenizeruaxUrlEmail) SetMaxTokenLength(v int)`

SetMaxTokenLength sets MaxTokenLength field to given value.

### HasMaxTokenLength

`func (o *TokenizeruaxUrlEmail) HasMaxTokenLength() bool`

HasMaxTokenLength returns a boolean if a field has been set.

### GetType

`func (o *TokenizeruaxUrlEmail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizeruaxUrlEmail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizeruaxUrlEmail) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


