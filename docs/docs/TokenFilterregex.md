# TokenFilterregex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matches** | **string** | Value that indicates whether to replace only the first matching pattern or all matching patterns. | 
**Pattern** | **string** | Regular expression pattern to apply to each token. | 
**Replacement** | **string** | Replacement string to substitute wherever a matching pattern occurs. | 
**Type** | **string** | Human-readable label that identifies this token filter type. | 

## Methods

### NewTokenFilterregex

`func NewTokenFilterregex(matches string, pattern string, replacement string, type_ string, ) *TokenFilterregex`

NewTokenFilterregex instantiates a new TokenFilterregex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenFilterregexWithDefaults

`func NewTokenFilterregexWithDefaults() *TokenFilterregex`

NewTokenFilterregexWithDefaults instantiates a new TokenFilterregex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatches

`func (o *TokenFilterregex) GetMatches() string`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *TokenFilterregex) GetMatchesOk() (*string, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *TokenFilterregex) SetMatches(v string)`

SetMatches sets Matches field to given value.


### GetPattern

`func (o *TokenFilterregex) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *TokenFilterregex) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *TokenFilterregex) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetReplacement

`func (o *TokenFilterregex) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *TokenFilterregex) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *TokenFilterregex) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.


### GetType

`func (o *TokenFilterregex) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenFilterregex) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenFilterregex) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


