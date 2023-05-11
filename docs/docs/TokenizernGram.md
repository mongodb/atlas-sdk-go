# TokenizernGram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxGram** | **int** | Characters to include in the longest token that Atlas Search creates. | 
**MinGram** | **int** | Characters to include in the shortest token that Atlas Search creates. | 
**Type** | **string** | Human-readable label that identifies this tokenizer type. | 

## Methods

### NewTokenizernGram

`func NewTokenizernGram(maxGram int, minGram int, type_ string, ) *TokenizernGram`

NewTokenizernGram instantiates a new TokenizernGram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizernGramWithDefaults

`func NewTokenizernGramWithDefaults() *TokenizernGram`

NewTokenizernGramWithDefaults instantiates a new TokenizernGram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxGram

`func (o *TokenizernGram) GetMaxGram() int`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *TokenizernGram) GetMaxGramOk() (*int, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *TokenizernGram) SetMaxGram(v int)`

SetMaxGram sets MaxGram field to given value.


### GetMinGram

`func (o *TokenizernGram) GetMinGram() int`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *TokenizernGram) GetMinGramOk() (*int, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *TokenizernGram) SetMinGram(v int)`

SetMinGram sets MinGram field to given value.


### GetType

`func (o *TokenizernGram) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizernGram) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizernGram) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


