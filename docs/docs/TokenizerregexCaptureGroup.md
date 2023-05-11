# TokenizerregexCaptureGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **int** | Index of the character group within the matching expression to extract into tokens. Use &#x60;0&#x60; to extract all character groups. | 
**Pattern** | **string** | Regular expression to match against. | 
**Type** | **string** | Human-readable label that identifies this tokenizer type. | 

## Methods

### NewTokenizerregexCaptureGroup

`func NewTokenizerregexCaptureGroup(group int, pattern string, type_ string, ) *TokenizerregexCaptureGroup`

NewTokenizerregexCaptureGroup instantiates a new TokenizerregexCaptureGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizerregexCaptureGroupWithDefaults

`func NewTokenizerregexCaptureGroupWithDefaults() *TokenizerregexCaptureGroup`

NewTokenizerregexCaptureGroupWithDefaults instantiates a new TokenizerregexCaptureGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *TokenizerregexCaptureGroup) GetGroup() int`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TokenizerregexCaptureGroup) GetGroupOk() (*int, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TokenizerregexCaptureGroup) SetGroup(v int)`

SetGroup sets Group field to given value.


### GetPattern

`func (o *TokenizerregexCaptureGroup) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *TokenizerregexCaptureGroup) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *TokenizerregexCaptureGroup) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetType

`func (o *TokenizerregexCaptureGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizerregexCaptureGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizerregexCaptureGroup) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


