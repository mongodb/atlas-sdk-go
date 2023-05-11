# TokenFilterdaitchMokotoffSoundex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalTokens** | Pointer to **string** | Value that indicates whether to include or omit the original tokens in the output of the token filter.  Choose &#x60;include&#x60; if you want to support queries on both the original tokens as well as the converted forms.   Choose &#x60;omit&#x60; if you want to query only on the converted forms of the original tokens. | [optional] [default to "include"]
**Type** | **string** | Human-readable label that identifies this token filter type. | 

## Methods

### NewTokenFilterdaitchMokotoffSoundex

`func NewTokenFilterdaitchMokotoffSoundex(type_ string, ) *TokenFilterdaitchMokotoffSoundex`

NewTokenFilterdaitchMokotoffSoundex instantiates a new TokenFilterdaitchMokotoffSoundex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenFilterdaitchMokotoffSoundexWithDefaults

`func NewTokenFilterdaitchMokotoffSoundexWithDefaults() *TokenFilterdaitchMokotoffSoundex`

NewTokenFilterdaitchMokotoffSoundexWithDefaults instantiates a new TokenFilterdaitchMokotoffSoundex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalTokens

`func (o *TokenFilterdaitchMokotoffSoundex) GetOriginalTokens() string`

GetOriginalTokens returns the OriginalTokens field if non-nil, zero value otherwise.

### GetOriginalTokensOk

`func (o *TokenFilterdaitchMokotoffSoundex) GetOriginalTokensOk() (*string, bool)`

GetOriginalTokensOk returns a tuple with the OriginalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTokens

`func (o *TokenFilterdaitchMokotoffSoundex) SetOriginalTokens(v string)`

SetOriginalTokens sets OriginalTokens field to given value.

### HasOriginalTokens

`func (o *TokenFilterdaitchMokotoffSoundex) HasOriginalTokens() bool`

HasOriginalTokens returns a boolean if a field has been set.

### GetType

`func (o *TokenFilterdaitchMokotoffSoundex) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenFilterdaitchMokotoffSoundex) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenFilterdaitchMokotoffSoundex) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


