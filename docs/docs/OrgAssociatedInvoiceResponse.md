# OrgAssociatedInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedInvoices** | Pointer to [**[]AssociatedInvoice**](AssociatedInvoice.md) | List of invoices associated with the organization for the specified period. | [optional] [readonly] 
**Month** | Pointer to **string** | Two-digit number that represents the month of the associated invoices. | [optional] [readonly] 
**Year** | Pointer to **string** | Four-digit number that represents the year of the associated invoices. | [optional] [readonly] 

## Methods

### NewOrgAssociatedInvoiceResponse

`func NewOrgAssociatedInvoiceResponse() *OrgAssociatedInvoiceResponse`

NewOrgAssociatedInvoiceResponse instantiates a new OrgAssociatedInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgAssociatedInvoiceResponseWithDefaults

`func NewOrgAssociatedInvoiceResponseWithDefaults() *OrgAssociatedInvoiceResponse`

NewOrgAssociatedInvoiceResponseWithDefaults instantiates a new OrgAssociatedInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedInvoices

`func (o *OrgAssociatedInvoiceResponse) GetAssociatedInvoices() []AssociatedInvoice`

GetAssociatedInvoices returns the AssociatedInvoices field if non-nil, zero value otherwise.

### GetAssociatedInvoicesOk

`func (o *OrgAssociatedInvoiceResponse) GetAssociatedInvoicesOk() (*[]AssociatedInvoice, bool)`

GetAssociatedInvoicesOk returns a tuple with the AssociatedInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedInvoices

`func (o *OrgAssociatedInvoiceResponse) SetAssociatedInvoices(v []AssociatedInvoice)`

SetAssociatedInvoices sets AssociatedInvoices field to given value.

### HasAssociatedInvoices

`func (o *OrgAssociatedInvoiceResponse) HasAssociatedInvoices() bool`

HasAssociatedInvoices returns a boolean if a field has been set.

### SetAssociatedInvoicesNil

`func (o *OrgAssociatedInvoiceResponse) SetAssociatedInvoicesNil()`

SetAssociatedInvoicesNil sets AssociatedInvoices to an explicit JSON null when marshaled, overriding any value previously set with SetAssociatedInvoices. Calling SetAssociatedInvoices again clears the null override.

### GetMonth

`func (o *OrgAssociatedInvoiceResponse) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *OrgAssociatedInvoiceResponse) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *OrgAssociatedInvoiceResponse) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *OrgAssociatedInvoiceResponse) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### SetMonthNil

`func (o *OrgAssociatedInvoiceResponse) SetMonthNil()`

SetMonthNil sets Month to an explicit JSON null when marshaled, overriding any value previously set with SetMonth. Calling SetMonth again clears the null override.

### GetYear

`func (o *OrgAssociatedInvoiceResponse) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *OrgAssociatedInvoiceResponse) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *OrgAssociatedInvoiceResponse) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *OrgAssociatedInvoiceResponse) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *OrgAssociatedInvoiceResponse) SetYearNil()`

SetYearNil sets Year to an explicit JSON null when marshaled, overriding any value previously set with SetYear. Calling SetYear again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


