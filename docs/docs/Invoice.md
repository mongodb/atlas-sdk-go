# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBilledCents** | Pointer to **int64** | Sum of services that the specified organization consumed in the period covered in this invoice. This parameter expresses its value in cents (100ths of one US Dollar) and calculates its value as **subtotalCents** + **salesTaxCents** - **startingBalanceCents**. | [optional] [readonly] 
**AmountPaidCents** | Pointer to **int64** | Sum that the specified organization paid toward this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this invoice. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CreditsCents** | Pointer to **int64** | Sum that MongoDB credited the specified organization toward this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud finished the billing period that this invoice covers. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project associated to this invoice. This identifying string doesn&#39;t appear on all invoices. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | [optional] [readonly] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | List that contains individual services included in this invoice. | [optional] [readonly] 
**LinkedInvoices** | Pointer to [**[]Invoice**](Invoice.md) | List that contains the invoices for organizations linked to the paying organization. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization charged for services consumed from MongoDB Cloud. | [optional] [readonly] 
**Payments** | Pointer to [**[]Payment**](Payment.md) | List that contains funds transferred to MongoDB to cover the specified service noted in this invoice. | [optional] [readonly] 
**Refunds** | Pointer to [**[]Refund**](Refund.md) | List that contains payments that MongoDB returned to the organization for this invoice. | [optional] [readonly] 
**SalesTaxCents** | Pointer to **int64** | Sum of sales tax applied to this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud began the billing period that this invoice covers. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**StartingBalanceCents** | Pointer to **int64** | Sum that the specified organization owed to MongoDB when MongoDB issued this invoice. This parameter expresses its value in US Dollars. | [optional] [readonly] 
**StatusName** | Pointer to **string** | Phase of payment processing in which this invoice exists when you made this request. Accepted phases include:  | Phase Value | Reason | |---|---| | CLOSED | MongoDB finalized all charges in the billing cycle but has yet to charge the customer. | | FAILED | MongoDB attempted to charge the provided credit card but charge for that amount failed. | | FORGIVEN | Customer initiated payment which MongoDB later forgave. | | FREE | All charges totalled zero so the customer won&#39;t be charged. | | INVOICED | MongoDB handled these charges using elastic invoicing. | | PAID | MongoDB succeeded in charging the provided credit card. | | PENDING | Invoice includes charges for the current billing cycle. | | PREPAID | Customer has a pre-paid plan so they won&#39;t be charged. |  | [optional] 
**SubtotalCents** | Pointer to **int64** | Sum of all positive invoice line items contained in this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | Date and time when MongoDB Cloud last updated the value of this payment. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBilledCents

`func (o *Invoice) GetAmountBilledCents() int64`

GetAmountBilledCents returns the AmountBilledCents field if non-nil, zero value otherwise.

### GetAmountBilledCentsOk

`func (o *Invoice) GetAmountBilledCentsOk() (*int64, bool)`

GetAmountBilledCentsOk returns a tuple with the AmountBilledCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBilledCents

`func (o *Invoice) SetAmountBilledCents(v int64)`

SetAmountBilledCents sets AmountBilledCents field to given value.

### HasAmountBilledCents

`func (o *Invoice) HasAmountBilledCents() bool`

HasAmountBilledCents returns a boolean if a field has been set.

### GetAmountPaidCents

`func (o *Invoice) GetAmountPaidCents() int64`

GetAmountPaidCents returns the AmountPaidCents field if non-nil, zero value otherwise.

### GetAmountPaidCentsOk

`func (o *Invoice) GetAmountPaidCentsOk() (*int64, bool)`

GetAmountPaidCentsOk returns a tuple with the AmountPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidCents

`func (o *Invoice) SetAmountPaidCents(v int64)`

SetAmountPaidCents sets AmountPaidCents field to given value.

### HasAmountPaidCents

`func (o *Invoice) HasAmountPaidCents() bool`

HasAmountPaidCents returns a boolean if a field has been set.

### GetCreated

`func (o *Invoice) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Invoice) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Invoice) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Invoice) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreditsCents

`func (o *Invoice) GetCreditsCents() int64`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *Invoice) GetCreditsCentsOk() (*int64, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *Invoice) SetCreditsCents(v int64)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *Invoice) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetEndDate

`func (o *Invoice) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Invoice) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Invoice) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Invoice) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroupId

`func (o *Invoice) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Invoice) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Invoice) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Invoice) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLineItems

`func (o *Invoice) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Invoice) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Invoice) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *Invoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLinkedInvoices

`func (o *Invoice) GetLinkedInvoices() []Invoice`

GetLinkedInvoices returns the LinkedInvoices field if non-nil, zero value otherwise.

### GetLinkedInvoicesOk

`func (o *Invoice) GetLinkedInvoicesOk() (*[]Invoice, bool)`

GetLinkedInvoicesOk returns a tuple with the LinkedInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInvoices

`func (o *Invoice) SetLinkedInvoices(v []Invoice)`

SetLinkedInvoices sets LinkedInvoices field to given value.

### HasLinkedInvoices

`func (o *Invoice) HasLinkedInvoices() bool`

HasLinkedInvoices returns a boolean if a field has been set.

### GetLinks

`func (o *Invoice) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Invoice) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Invoice) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Invoice) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrgId

`func (o *Invoice) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Invoice) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Invoice) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Invoice) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPayments

`func (o *Invoice) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Invoice) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Invoice) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *Invoice) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetRefunds

`func (o *Invoice) GetRefunds() []Refund`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *Invoice) GetRefundsOk() (*[]Refund, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *Invoice) SetRefunds(v []Refund)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *Invoice) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetSalesTaxCents

`func (o *Invoice) GetSalesTaxCents() int64`

GetSalesTaxCents returns the SalesTaxCents field if non-nil, zero value otherwise.

### GetSalesTaxCentsOk

`func (o *Invoice) GetSalesTaxCentsOk() (*int64, bool)`

GetSalesTaxCentsOk returns a tuple with the SalesTaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCents

`func (o *Invoice) SetSalesTaxCents(v int64)`

SetSalesTaxCents sets SalesTaxCents field to given value.

### HasSalesTaxCents

`func (o *Invoice) HasSalesTaxCents() bool`

HasSalesTaxCents returns a boolean if a field has been set.

### GetStartDate

`func (o *Invoice) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Invoice) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Invoice) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Invoice) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartingBalanceCents

`func (o *Invoice) GetStartingBalanceCents() int64`

GetStartingBalanceCents returns the StartingBalanceCents field if non-nil, zero value otherwise.

### GetStartingBalanceCentsOk

`func (o *Invoice) GetStartingBalanceCentsOk() (*int64, bool)`

GetStartingBalanceCentsOk returns a tuple with the StartingBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBalanceCents

`func (o *Invoice) SetStartingBalanceCents(v int64)`

SetStartingBalanceCents sets StartingBalanceCents field to given value.

### HasStartingBalanceCents

`func (o *Invoice) HasStartingBalanceCents() bool`

HasStartingBalanceCents returns a boolean if a field has been set.

### GetStatusName

`func (o *Invoice) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *Invoice) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *Invoice) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *Invoice) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetSubtotalCents

`func (o *Invoice) GetSubtotalCents() int64`

GetSubtotalCents returns the SubtotalCents field if non-nil, zero value otherwise.

### GetSubtotalCentsOk

`func (o *Invoice) GetSubtotalCentsOk() (*int64, bool)`

GetSubtotalCentsOk returns a tuple with the SubtotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalCents

`func (o *Invoice) SetSubtotalCents(v int64)`

SetSubtotalCents sets SubtotalCents field to given value.

### HasSubtotalCents

`func (o *Invoice) HasSubtotalCents() bool`

HasSubtotalCents returns a boolean if a field has been set.

### GetUpdated

`func (o *Invoice) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Invoice) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Invoice) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Invoice) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


