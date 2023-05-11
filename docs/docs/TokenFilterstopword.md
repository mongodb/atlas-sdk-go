# TokenFilterstopword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreCase** | Pointer to **bool** | Flag that indicates whether to ignore the case of stop words when filtering the tokens to remove. | [optional] [default to true]
**Tokens** | **[]string** | The stop words that correspond to the tokens to remove. Value must be one or more stop words. | 
**Type** | **string** | Human-readable label that identifies this token filter type. | 

## Methods

### NewTokenFilterstopword

`func NewTokenFilterstopword(tokens []string, type_ string, ) *TokenFilterstopword`

NewTokenFilterstopword instantiates a new TokenFilterstopword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenFilterstopwordWithDefaults

`func NewTokenFilterstopwordWithDefaults() *TokenFilterstopword`

NewTokenFilterstopwordWithDefaults instantiates a new TokenFilterstopword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreCase

`func (o *TokenFilterstopword) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *TokenFilterstopword) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *TokenFilterstopword) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *TokenFilterstopword) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetTokens

`func (o *TokenFilterstopword) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TokenFilterstopword) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TokenFilterstopword) SetTokens(v []string)`

SetTokens sets Tokens field to given value.


### GetType

`func (o *TokenFilterstopword) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenFilterstopword) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenFilterstopword) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


