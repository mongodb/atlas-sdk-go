# TokenizeredgeGram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxGram** | **int** | Characters to include in the longest token that Atlas Search creates. | 
**MinGram** | **int** | Characters to include in the shortest token that Atlas Search creates. | 
**Type** | **string** | Human-readable label that identifies this tokenizer type. | 

## Methods

### NewTokenizeredgeGram

`func NewTokenizeredgeGram(maxGram int, minGram int, type_ string, ) *TokenizeredgeGram`

NewTokenizeredgeGram instantiates a new TokenizeredgeGram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizeredgeGramWithDefaults

`func NewTokenizeredgeGramWithDefaults() *TokenizeredgeGram`

NewTokenizeredgeGramWithDefaults instantiates a new TokenizeredgeGram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxGram

`func (o *TokenizeredgeGram) GetMaxGram() int`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *TokenizeredgeGram) GetMaxGramOk() (*int, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *TokenizeredgeGram) SetMaxGram(v int)`

SetMaxGram sets MaxGram field to given value.


### GetMinGram

`func (o *TokenizeredgeGram) GetMinGram() int`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *TokenizeredgeGram) GetMinGramOk() (*int, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *TokenizeredgeGram) SetMinGram(v int)`

SetMinGram sets MinGram field to given value.


### GetType

`func (o *TokenizeredgeGram) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizeredgeGram) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizeredgeGram) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


