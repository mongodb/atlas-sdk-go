# AssociatedInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | Pointer to **string** | Unique 24-hexadecimal digit identifier for an invoice. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit identifier for an organization. | [optional] [readonly] 

## Methods

### NewAssociatedInvoice

`func NewAssociatedInvoice() *AssociatedInvoice`

NewAssociatedInvoice instantiates a new AssociatedInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociatedInvoiceWithDefaults

`func NewAssociatedInvoiceWithDefaults() *AssociatedInvoice`

NewAssociatedInvoiceWithDefaults instantiates a new AssociatedInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *AssociatedInvoice) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *AssociatedInvoice) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *AssociatedInvoice) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *AssociatedInvoice) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *AssociatedInvoice) SetInvoiceIdNil()`

SetInvoiceIdNil sets InvoiceId to an explicit JSON null when marshaled, overriding any value previously set with SetInvoiceId. Calling SetInvoiceId again clears the null override.

### GetOrgId

`func (o *AssociatedInvoice) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AssociatedInvoice) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AssociatedInvoice) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AssociatedInvoice) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *AssociatedInvoice) SetOrgIdNil()`

SetOrgIdNil sets OrgId to an explicit JSON null when marshaled, overriding any value previously set with SetOrgId. Calling SetOrgId again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


