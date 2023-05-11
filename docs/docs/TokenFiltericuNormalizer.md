# TokenFiltericuNormalizer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NormalizationForm** | Pointer to **string** | Normalization form to apply. | [optional] [default to "nfc"]
**Type** | **string** | Human-readable label that identifies this token filter type. | 

## Methods

### NewTokenFiltericuNormalizer

`func NewTokenFiltericuNormalizer(type_ string, ) *TokenFiltericuNormalizer`

NewTokenFiltericuNormalizer instantiates a new TokenFiltericuNormalizer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenFiltericuNormalizerWithDefaults

`func NewTokenFiltericuNormalizerWithDefaults() *TokenFiltericuNormalizer`

NewTokenFiltericuNormalizerWithDefaults instantiates a new TokenFiltericuNormalizer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNormalizationForm

`func (o *TokenFiltericuNormalizer) GetNormalizationForm() string`

GetNormalizationForm returns the NormalizationForm field if non-nil, zero value otherwise.

### GetNormalizationFormOk

`func (o *TokenFiltericuNormalizer) GetNormalizationFormOk() (*string, bool)`

GetNormalizationFormOk returns a tuple with the NormalizationForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizationForm

`func (o *TokenFiltericuNormalizer) SetNormalizationForm(v string)`

SetNormalizationForm sets NormalizationForm field to given value.

### HasNormalizationForm

`func (o *TokenFiltericuNormalizer) HasNormalizationForm() bool`

HasNormalizationForm returns a boolean if a field has been set.

### GetType

`func (o *TokenFiltericuNormalizer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenFiltericuNormalizer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenFiltericuNormalizer) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


