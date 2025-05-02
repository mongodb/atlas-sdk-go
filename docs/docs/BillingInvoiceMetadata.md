# BillingInvoiceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBilledCents** | Pointer to **int64** | Sum of services that the specified organization consumed in the period covered in this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**AmountPaidCents** | Pointer to **int64** | Sum that the specified organization paid toward this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this invoice. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CreditsCents** | Pointer to **int64** | Sum that MongoDB credited the specified organization toward this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud finished the billing period that this invoice covers. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | [optional] [readonly] 
**LinkedInvoices** | Pointer to [**[]BillingInvoiceMetadata**](BillingInvoiceMetadata.md) | List that contains the invoices for organizations linked to the paying organization. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization charged for services consumed from MongoDB Cloud. | [optional] [readonly] 
**SalesTaxCents** | Pointer to **int64** | Sum of sales tax applied to this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud began the billing period that this invoice covers. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**StartingBalanceCents** | Pointer to **int64** | Sum that the specified organization owed to MongoDB when MongoDB issued this invoice. This parameter expresses its value in US Dollars. | [optional] [readonly] 
**StatusName** | Pointer to **string** | Phase of payment processing in which this invoice exists when you made this request. Accepted phases include:  - &#x60;CLOSED&#x60;: MongoDB finalized all charges in the billing cycle but has yet to charge the customer. - &#x60;FAILED&#x60;: MongoDB attempted to charge the provided credit card but charge for that amount failed. - &#x60;FORGIVEN&#x60;: Customer initiated payment which MongoDB later forgave. - &#x60;FREE&#x60;: All charges totalled zero so the customer won&#39;t be charged. - &#x60;INVOICED&#x60;: MongoDB handled these charges using elastic invoicing. - &#x60;PAID&#x60;: MongoDB succeeded in charging the provided credit card. - &#x60;PENDING&#x60;: Invoice includes charges for the current billing cycle. - &#x60;PREPAID&#x60;: Customer has a pre-paid plan so they won&#39;t be charged. | [optional] 
**SubtotalCents** | Pointer to **int64** | Sum of all positive invoice line items contained in this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | Date and time when MongoDB Cloud last updated the value of this payment. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewBillingInvoiceMetadata

`func NewBillingInvoiceMetadata() *BillingInvoiceMetadata`

NewBillingInvoiceMetadata instantiates a new BillingInvoiceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceMetadataWithDefaults

`func NewBillingInvoiceMetadataWithDefaults() *BillingInvoiceMetadata`

NewBillingInvoiceMetadataWithDefaults instantiates a new BillingInvoiceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBilledCents

`func (o *BillingInvoiceMetadata) GetAmountBilledCents() int64`

GetAmountBilledCents returns the AmountBilledCents field if non-nil, zero value otherwise.

### GetAmountBilledCentsOk

`func (o *BillingInvoiceMetadata) GetAmountBilledCentsOk() (*int64, bool)`

GetAmountBilledCentsOk returns a tuple with the AmountBilledCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBilledCents

`func (o *BillingInvoiceMetadata) SetAmountBilledCents(v int64)`

SetAmountBilledCents sets AmountBilledCents field to given value.

### HasAmountBilledCents

`func (o *BillingInvoiceMetadata) HasAmountBilledCents() bool`

HasAmountBilledCents returns a boolean if a field has been set.
### GetAmountPaidCents

`func (o *BillingInvoiceMetadata) GetAmountPaidCents() int64`

GetAmountPaidCents returns the AmountPaidCents field if non-nil, zero value otherwise.

### GetAmountPaidCentsOk

`func (o *BillingInvoiceMetadata) GetAmountPaidCentsOk() (*int64, bool)`

GetAmountPaidCentsOk returns a tuple with the AmountPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidCents

`func (o *BillingInvoiceMetadata) SetAmountPaidCents(v int64)`

SetAmountPaidCents sets AmountPaidCents field to given value.

### HasAmountPaidCents

`func (o *BillingInvoiceMetadata) HasAmountPaidCents() bool`

HasAmountPaidCents returns a boolean if a field has been set.
### GetCreated

`func (o *BillingInvoiceMetadata) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BillingInvoiceMetadata) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BillingInvoiceMetadata) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BillingInvoiceMetadata) HasCreated() bool`

HasCreated returns a boolean if a field has been set.
### GetCreditsCents

`func (o *BillingInvoiceMetadata) GetCreditsCents() int64`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *BillingInvoiceMetadata) GetCreditsCentsOk() (*int64, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *BillingInvoiceMetadata) SetCreditsCents(v int64)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *BillingInvoiceMetadata) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.
### GetEndDate

`func (o *BillingInvoiceMetadata) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingInvoiceMetadata) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingInvoiceMetadata) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingInvoiceMetadata) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.
### GetId

`func (o *BillingInvoiceMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoiceMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoiceMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInvoiceMetadata) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLinkedInvoices

`func (o *BillingInvoiceMetadata) GetLinkedInvoices() []BillingInvoiceMetadata`

GetLinkedInvoices returns the LinkedInvoices field if non-nil, zero value otherwise.

### GetLinkedInvoicesOk

`func (o *BillingInvoiceMetadata) GetLinkedInvoicesOk() (*[]BillingInvoiceMetadata, bool)`

GetLinkedInvoicesOk returns a tuple with the LinkedInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInvoices

`func (o *BillingInvoiceMetadata) SetLinkedInvoices(v []BillingInvoiceMetadata)`

SetLinkedInvoices sets LinkedInvoices field to given value.

### HasLinkedInvoices

`func (o *BillingInvoiceMetadata) HasLinkedInvoices() bool`

HasLinkedInvoices returns a boolean if a field has been set.
### GetLinks

`func (o *BillingInvoiceMetadata) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BillingInvoiceMetadata) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BillingInvoiceMetadata) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BillingInvoiceMetadata) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOrgId

`func (o *BillingInvoiceMetadata) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BillingInvoiceMetadata) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BillingInvoiceMetadata) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *BillingInvoiceMetadata) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.
### GetSalesTaxCents

`func (o *BillingInvoiceMetadata) GetSalesTaxCents() int64`

GetSalesTaxCents returns the SalesTaxCents field if non-nil, zero value otherwise.

### GetSalesTaxCentsOk

`func (o *BillingInvoiceMetadata) GetSalesTaxCentsOk() (*int64, bool)`

GetSalesTaxCentsOk returns a tuple with the SalesTaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCents

`func (o *BillingInvoiceMetadata) SetSalesTaxCents(v int64)`

SetSalesTaxCents sets SalesTaxCents field to given value.

### HasSalesTaxCents

`func (o *BillingInvoiceMetadata) HasSalesTaxCents() bool`

HasSalesTaxCents returns a boolean if a field has been set.
### GetStartDate

`func (o *BillingInvoiceMetadata) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingInvoiceMetadata) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingInvoiceMetadata) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingInvoiceMetadata) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.
### GetStartingBalanceCents

`func (o *BillingInvoiceMetadata) GetStartingBalanceCents() int64`

GetStartingBalanceCents returns the StartingBalanceCents field if non-nil, zero value otherwise.

### GetStartingBalanceCentsOk

`func (o *BillingInvoiceMetadata) GetStartingBalanceCentsOk() (*int64, bool)`

GetStartingBalanceCentsOk returns a tuple with the StartingBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBalanceCents

`func (o *BillingInvoiceMetadata) SetStartingBalanceCents(v int64)`

SetStartingBalanceCents sets StartingBalanceCents field to given value.

### HasStartingBalanceCents

`func (o *BillingInvoiceMetadata) HasStartingBalanceCents() bool`

HasStartingBalanceCents returns a boolean if a field has been set.
### GetStatusName

`func (o *BillingInvoiceMetadata) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *BillingInvoiceMetadata) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *BillingInvoiceMetadata) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *BillingInvoiceMetadata) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.
### GetSubtotalCents

`func (o *BillingInvoiceMetadata) GetSubtotalCents() int64`

GetSubtotalCents returns the SubtotalCents field if non-nil, zero value otherwise.

### GetSubtotalCentsOk

`func (o *BillingInvoiceMetadata) GetSubtotalCentsOk() (*int64, bool)`

GetSubtotalCentsOk returns a tuple with the SubtotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalCents

`func (o *BillingInvoiceMetadata) SetSubtotalCents(v int64)`

SetSubtotalCents sets SubtotalCents field to given value.

### HasSubtotalCents

`func (o *BillingInvoiceMetadata) HasSubtotalCents() bool`

HasSubtotalCents returns a boolean if a field has been set.
### GetUpdated

`func (o *BillingInvoiceMetadata) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BillingInvoiceMetadata) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BillingInvoiceMetadata) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BillingInvoiceMetadata) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


