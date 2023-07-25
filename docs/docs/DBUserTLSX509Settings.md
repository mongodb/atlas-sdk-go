# DBUserTLSX509Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cas** | Pointer to **string** | Concatenated list of customer certificate authority (CA) certificates needed to authenticate database users. MongoDB Cloud expects this as a PEM-formatted certificate. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewDBUserTLSX509Settings

`func NewDBUserTLSX509Settings() *DBUserTLSX509Settings`

NewDBUserTLSX509Settings instantiates a new DBUserTLSX509Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBUserTLSX509SettingsWithDefaults

`func NewDBUserTLSX509SettingsWithDefaults() *DBUserTLSX509Settings`

NewDBUserTLSX509SettingsWithDefaults instantiates a new DBUserTLSX509Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCas

`func (o *DBUserTLSX509Settings) GetCas() string`

GetCas returns the Cas field if non-nil, zero value otherwise.

### GetCasOk

`func (o *DBUserTLSX509Settings) GetCasOk() (*string, bool)`

GetCasOk returns a tuple with the Cas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCas

`func (o *DBUserTLSX509Settings) SetCas(v string)`

SetCas sets Cas field to given value.

### HasCas

`func (o *DBUserTLSX509Settings) HasCas() bool`

HasCas returns a boolean if a field has been set.
### GetLinks

`func (o *DBUserTLSX509Settings) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DBUserTLSX509Settings) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DBUserTLSX509Settings) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DBUserTLSX509Settings) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


