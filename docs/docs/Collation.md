# Collation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternate** | Pointer to **string** | Method to handle whitespace and punctuation as base characters for purposes of comparison. &#x60;\&quot;non-ignorable\&quot;&#x60; will evaluate Whitespace and Punctuation as Base Characters. &#x60;\&quot;shifted\&quot;&#x60; will not, MongoDB Cloud distinguishes these characters when &#x60;\&quot;strength\&quot; &gt; 3&#x60;. | [optional] [default to "non-ignorable"]
**Backwards** | Pointer to **bool** | Flag that indicates whether strings with diacritics sort from back of the string. Some French dictionary orders strings in this way. &#x60;true&#x60; will compare from back to front. &#x60;false&#x60; will compare from front to back. | [optional] [default to false]
**CaseFirst** | Pointer to **string** | Method to handle sort order of case differences during tertiary level comparisons. &#x60;\&quot;upper\&quot;&#x60; sorts Uppercase before lowercase. &#x60;\&quot;lower\&quot;&#x60; sorts Lowercase before uppercase. &#x60;\&quot;off\&quot;&#x60; is similar to \&quot;lower\&quot; with slight differences. | [optional] [default to "false"]
**CaseLevel** | Pointer to **bool** | Flag that indicates whether to include case comparison when &#x60;\&quot;strength\&quot; : 1&#x60; or &#x60;\&quot;strength\&quot; : 2&#x60;.  | Value | Compare case at level 1 or 2? | Strength Level | Comparisons Include |  |---|---|---|---|  | true | Yes | 1 | Base characters and case. |  |  |  | 2 | Base characters, diacritics (and possible other secondary differences),   and case. |  | false | No |  |  |  | [optional] [default to false]
**Locale** | **string** | International Components for Unicode (ICU) code that represents a localized language. To specify simple binary comparison, set &#x60;\&quot;locale\&quot; : \&quot;simple\&quot;&#x60;. | 
**MaxVariable** | Pointer to **string** | Field that indicates which characters can be ignored when &#x60;\&quot;alternate\&quot; : \&quot;shifted\&quot;&#x60;.&#x60;\&quot;punct\&quot;&#x60; ignores both whitespace and punctuation. &#x60;\&quot;space\&quot;&#x60; ignores whitespace. Thishas no affect if &#x60;\&quot;alternate\&quot; : \&quot;non-ignorable\&quot;&#x60;. | [optional] 
**Normalization** | Pointer to **bool** | Flag that indicates whether to check if the text requires normalization and then perform it. Most text doesn&#39;t require this normalization processing.  &#x60;true&#x60; will check if fully normalized and perform normalization to compare text. &#x60;false&#x60; will not check. | [optional] [default to false]
**NumericOrdering** | Pointer to **bool** | Flag that indicates whether to compare sequences of digits as numbers or as strings. &#x60;true&#x60; will compare as numbers, this results in &#x60;10 &gt; 2&#x60;. &#x60;false&#x60; will Compare as strings. This results in &#x60;\&quot;10\&quot; &lt; \&quot;2\&quot;&#x60;. | [optional] [default to false]
**Strength** | Pointer to **int** | Degree of comparison to perform when sorting words. MongoDB Cloud accepts the following values:  | Value | Comparison Level | Comparison Method | |---|---|---| | 1 | Primary | Compares the base characters only, ignoring other differences such as diacritics and case. | | 2 | Secondary | Compares base characters (primary) and diacritics (secondary). Primary differences take precedence over secondary differences. | | 3 | Tertiary | Compares base characters (primary), diacritics (secondary), and case and variants (tertiary). Differences between base characters takes precedence over secondary differences which take precedence over tertiary differences. | | 4 | Quaternary | Compares for the specific use case to consider punctuation when levels 1 through 3 ignore punctuation or for processing Japanese text. | | 5 | Identical | Compares for the specific use case of tie breaker. |  | [optional] [default to 3]

## Methods

### NewCollation

`func NewCollation(locale string, ) *Collation`

NewCollation instantiates a new Collation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollationWithDefaults

`func NewCollationWithDefaults() *Collation`

NewCollationWithDefaults instantiates a new Collation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternate

`func (o *Collation) GetAlternate() string`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *Collation) GetAlternateOk() (*string, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *Collation) SetAlternate(v string)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *Collation) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.
### GetBackwards

`func (o *Collation) GetBackwards() bool`

GetBackwards returns the Backwards field if non-nil, zero value otherwise.

### GetBackwardsOk

`func (o *Collation) GetBackwardsOk() (*bool, bool)`

GetBackwardsOk returns a tuple with the Backwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackwards

`func (o *Collation) SetBackwards(v bool)`

SetBackwards sets Backwards field to given value.

### HasBackwards

`func (o *Collation) HasBackwards() bool`

HasBackwards returns a boolean if a field has been set.
### GetCaseFirst

`func (o *Collation) GetCaseFirst() string`

GetCaseFirst returns the CaseFirst field if non-nil, zero value otherwise.

### GetCaseFirstOk

`func (o *Collation) GetCaseFirstOk() (*string, bool)`

GetCaseFirstOk returns a tuple with the CaseFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseFirst

`func (o *Collation) SetCaseFirst(v string)`

SetCaseFirst sets CaseFirst field to given value.

### HasCaseFirst

`func (o *Collation) HasCaseFirst() bool`

HasCaseFirst returns a boolean if a field has been set.
### GetCaseLevel

`func (o *Collation) GetCaseLevel() bool`

GetCaseLevel returns the CaseLevel field if non-nil, zero value otherwise.

### GetCaseLevelOk

`func (o *Collation) GetCaseLevelOk() (*bool, bool)`

GetCaseLevelOk returns a tuple with the CaseLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseLevel

`func (o *Collation) SetCaseLevel(v bool)`

SetCaseLevel sets CaseLevel field to given value.

### HasCaseLevel

`func (o *Collation) HasCaseLevel() bool`

HasCaseLevel returns a boolean if a field has been set.
### GetLocale

`func (o *Collation) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Collation) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Collation) SetLocale(v string)`

SetLocale sets Locale field to given value.

### GetMaxVariable

`func (o *Collation) GetMaxVariable() string`

GetMaxVariable returns the MaxVariable field if non-nil, zero value otherwise.

### GetMaxVariableOk

`func (o *Collation) GetMaxVariableOk() (*string, bool)`

GetMaxVariableOk returns a tuple with the MaxVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVariable

`func (o *Collation) SetMaxVariable(v string)`

SetMaxVariable sets MaxVariable field to given value.

### HasMaxVariable

`func (o *Collation) HasMaxVariable() bool`

HasMaxVariable returns a boolean if a field has been set.
### GetNormalization

`func (o *Collation) GetNormalization() bool`

GetNormalization returns the Normalization field if non-nil, zero value otherwise.

### GetNormalizationOk

`func (o *Collation) GetNormalizationOk() (*bool, bool)`

GetNormalizationOk returns a tuple with the Normalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalization

`func (o *Collation) SetNormalization(v bool)`

SetNormalization sets Normalization field to given value.

### HasNormalization

`func (o *Collation) HasNormalization() bool`

HasNormalization returns a boolean if a field has been set.
### GetNumericOrdering

`func (o *Collation) GetNumericOrdering() bool`

GetNumericOrdering returns the NumericOrdering field if non-nil, zero value otherwise.

### GetNumericOrderingOk

`func (o *Collation) GetNumericOrderingOk() (*bool, bool)`

GetNumericOrderingOk returns a tuple with the NumericOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericOrdering

`func (o *Collation) SetNumericOrdering(v bool)`

SetNumericOrdering sets NumericOrdering field to given value.

### HasNumericOrdering

`func (o *Collation) HasNumericOrdering() bool`

HasNumericOrdering returns a boolean if a field has been set.
### GetStrength

`func (o *Collation) GetStrength() int`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *Collation) GetStrengthOk() (*int, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *Collation) SetStrength(v int)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *Collation) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


