# CustomerX509

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cas** | Pointer to **string** | Concatenated list of customer certificate authority (CA) certificates needed to authenticate database users. MongoDB Cloud expects this as a PEM-formatted certificate. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewCustomerX509

`func NewCustomerX509() *CustomerX509`

NewCustomerX509 instantiates a new CustomerX509 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerX509WithDefaults

`func NewCustomerX509WithDefaults() *CustomerX509`

NewCustomerX509WithDefaults instantiates a new CustomerX509 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCas

`func (o *CustomerX509) GetCas() string`

GetCas returns the Cas field if non-nil, zero value otherwise.

### GetCasOk

`func (o *CustomerX509) GetCasOk() (*string, bool)`

GetCasOk returns a tuple with the Cas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCas

`func (o *CustomerX509) SetCas(v string)`

SetCas sets Cas field to given value.

### HasCas

`func (o *CustomerX509) HasCas() bool`

HasCas returns a boolean if a field has been set.

### GetLinks

`func (o *CustomerX509) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomerX509) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomerX509) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomerX509) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


