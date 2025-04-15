# BillingInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBilledCents** | Pointer to **int64** | Sum of services that the specified organization consumed in the period covered in this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**AmountPaidCents** | Pointer to **int64** | Sum that the specified organization paid toward this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this invoice. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CreditsCents** | Pointer to **int64** | Sum that MongoDB credited the specified organization toward this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud finished the billing period that this invoice covers. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | [optional] [readonly] 
**LineItems** | Pointer to [**[]InvoiceLineItem**](InvoiceLineItem.md) | List that contains individual services included in this invoice. | [optional] [readonly] 
**LinkedInvoices** | Pointer to [**[]BillingInvoice**](BillingInvoice.md) | List that contains the invoices for organizations linked to the paying organization. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization charged for services consumed from MongoDB Cloud. | [optional] [readonly] 
**Payments** | Pointer to [**[]BillingPayment**](BillingPayment.md) | List that contains funds transferred to MongoDB to cover the specified service noted in this invoice. | [optional] [readonly] 
**Refunds** | Pointer to [**[]BillingRefund**](BillingRefund.md) | List that contains payments that MongoDB returned to the organization for this invoice. | [optional] [readonly] 
**SalesTaxCents** | Pointer to **int64** | Sum of sales tax applied to this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud began the billing period that this invoice covers. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**StartingBalanceCents** | Pointer to **int64** | Sum that the specified organization owed to MongoDB when MongoDB issued this invoice. This parameter expresses its value in US Dollars. | [optional] [readonly] 
**StatusName** | Pointer to **string** | Phase of payment processing in which this invoice exists when you made this request. Accepted phases include:  - &#x60;CLOSED&#x60;: MongoDB finalized all charges in the billing cycle but has yet to charge the customer. - &#x60;FAILED&#x60;: MongoDB attempted to charge the provided credit card but charge for that amount failed. - &#x60;FORGIVEN&#x60;: Customer initiated payment which MongoDB later forgave. - &#x60;FREE&#x60;: All charges totalled zero so the customer won&#39;t be charged. - &#x60;INVOICED&#x60;: MongoDB handled these charges using elastic invoicing. - &#x60;PAID&#x60;: MongoDB succeeded in charging the provided credit card. - &#x60;PENDING&#x60;: Invoice includes charges for the current billing cycle. - &#x60;PREPAID&#x60;: Customer has a pre-paid plan so they won&#39;t be charged. | [optional] 
**SubtotalCents** | Pointer to **int64** | Sum of all positive invoice line items contained in this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | Date and time when MongoDB Cloud last updated the value of this payment. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewBillingInvoice

`func NewBillingInvoice() *BillingInvoice`

NewBillingInvoice instantiates a new BillingInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceWithDefaults

`func NewBillingInvoiceWithDefaults() *BillingInvoice`

NewBillingInvoiceWithDefaults instantiates a new BillingInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBilledCents

`func (o *BillingInvoice) GetAmountBilledCents() int64`

GetAmountBilledCents returns the AmountBilledCents field if non-nil, zero value otherwise.

### GetAmountBilledCentsOk

`func (o *BillingInvoice) GetAmountBilledCentsOk() (*int64, bool)`

GetAmountBilledCentsOk returns a tuple with the AmountBilledCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBilledCents

`func (o *BillingInvoice) SetAmountBilledCents(v int64)`

SetAmountBilledCents sets AmountBilledCents field to given value.

### HasAmountBilledCents

`func (o *BillingInvoice) HasAmountBilledCents() bool`

HasAmountBilledCents returns a boolean if a field has been set.
### GetAmountPaidCents

`func (o *BillingInvoice) GetAmountPaidCents() int64`

GetAmountPaidCents returns the AmountPaidCents field if non-nil, zero value otherwise.

### GetAmountPaidCentsOk

`func (o *BillingInvoice) GetAmountPaidCentsOk() (*int64, bool)`

GetAmountPaidCentsOk returns a tuple with the AmountPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidCents

`func (o *BillingInvoice) SetAmountPaidCents(v int64)`

SetAmountPaidCents sets AmountPaidCents field to given value.

### HasAmountPaidCents

`func (o *BillingInvoice) HasAmountPaidCents() bool`

HasAmountPaidCents returns a boolean if a field has been set.
### GetCreated

`func (o *BillingInvoice) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BillingInvoice) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BillingInvoice) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BillingInvoice) HasCreated() bool`

HasCreated returns a boolean if a field has been set.
### GetCreditsCents

`func (o *BillingInvoice) GetCreditsCents() int64`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *BillingInvoice) GetCreditsCentsOk() (*int64, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *BillingInvoice) SetCreditsCents(v int64)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *BillingInvoice) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.
### GetEndDate

`func (o *BillingInvoice) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingInvoice) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingInvoice) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingInvoice) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.
### GetId

`func (o *BillingInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInvoice) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLineItems

`func (o *BillingInvoice) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BillingInvoice) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BillingInvoice) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *BillingInvoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.
### GetLinkedInvoices

`func (o *BillingInvoice) GetLinkedInvoices() []BillingInvoice`

GetLinkedInvoices returns the LinkedInvoices field if non-nil, zero value otherwise.

### GetLinkedInvoicesOk

`func (o *BillingInvoice) GetLinkedInvoicesOk() (*[]BillingInvoice, bool)`

GetLinkedInvoicesOk returns a tuple with the LinkedInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInvoices

`func (o *BillingInvoice) SetLinkedInvoices(v []BillingInvoice)`

SetLinkedInvoices sets LinkedInvoices field to given value.

### HasLinkedInvoices

`func (o *BillingInvoice) HasLinkedInvoices() bool`

HasLinkedInvoices returns a boolean if a field has been set.
### GetLinks

`func (o *BillingInvoice) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BillingInvoice) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BillingInvoice) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BillingInvoice) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOrgId

`func (o *BillingInvoice) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BillingInvoice) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BillingInvoice) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *BillingInvoice) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.
### GetPayments

`func (o *BillingInvoice) GetPayments() []BillingPayment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *BillingInvoice) GetPaymentsOk() (*[]BillingPayment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *BillingInvoice) SetPayments(v []BillingPayment)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *BillingInvoice) HasPayments() bool`

HasPayments returns a boolean if a field has been set.
### GetRefunds

`func (o *BillingInvoice) GetRefunds() []BillingRefund`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *BillingInvoice) GetRefundsOk() (*[]BillingRefund, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *BillingInvoice) SetRefunds(v []BillingRefund)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *BillingInvoice) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.
### GetSalesTaxCents

`func (o *BillingInvoice) GetSalesTaxCents() int64`

GetSalesTaxCents returns the SalesTaxCents field if non-nil, zero value otherwise.

### GetSalesTaxCentsOk

`func (o *BillingInvoice) GetSalesTaxCentsOk() (*int64, bool)`

GetSalesTaxCentsOk returns a tuple with the SalesTaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCents

`func (o *BillingInvoice) SetSalesTaxCents(v int64)`

SetSalesTaxCents sets SalesTaxCents field to given value.

### HasSalesTaxCents

`func (o *BillingInvoice) HasSalesTaxCents() bool`

HasSalesTaxCents returns a boolean if a field has been set.
### GetStartDate

`func (o *BillingInvoice) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingInvoice) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingInvoice) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingInvoice) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.
### GetStartingBalanceCents

`func (o *BillingInvoice) GetStartingBalanceCents() int64`

GetStartingBalanceCents returns the StartingBalanceCents field if non-nil, zero value otherwise.

### GetStartingBalanceCentsOk

`func (o *BillingInvoice) GetStartingBalanceCentsOk() (*int64, bool)`

GetStartingBalanceCentsOk returns a tuple with the StartingBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBalanceCents

`func (o *BillingInvoice) SetStartingBalanceCents(v int64)`

SetStartingBalanceCents sets StartingBalanceCents field to given value.

### HasStartingBalanceCents

`func (o *BillingInvoice) HasStartingBalanceCents() bool`

HasStartingBalanceCents returns a boolean if a field has been set.
### GetStatusName

`func (o *BillingInvoice) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *BillingInvoice) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *BillingInvoice) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *BillingInvoice) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.
### GetSubtotalCents

`func (o *BillingInvoice) GetSubtotalCents() int64`

GetSubtotalCents returns the SubtotalCents field if non-nil, zero value otherwise.

### GetSubtotalCentsOk

`func (o *BillingInvoice) GetSubtotalCentsOk() (*int64, bool)`

GetSubtotalCentsOk returns a tuple with the SubtotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalCents

`func (o *BillingInvoice) SetSubtotalCents(v int64)`

SetSubtotalCents sets SubtotalCents field to given value.

### HasSubtotalCents

`func (o *BillingInvoice) HasSubtotalCents() bool`

HasSubtotalCents returns a boolean if a field has been set.
### GetUpdated

`func (o *BillingInvoice) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BillingInvoice) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BillingInvoice) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BillingInvoice) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


