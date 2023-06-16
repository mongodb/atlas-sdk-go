# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBilledCents** | Pointer to **int64** | Sum of services that the specified organization consumed in the period covered in this invoice. This parameter expresses its value in cents (100ths of one US Dollar) and calculates its value as **subtotalCents** + **salesTaxCents** - **startingBalanceCents**. | [optional] [readonly] 
**AmountPaidCents** | Pointer to **int64** | Sum that the specified organization paid toward the associated invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when the customer made this payment attempt. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this payment toward the associated invoice. | [optional] [readonly] 
**SalesTaxCents** | Pointer to **int64** | Sum of sales tax applied to this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**StatusName** | Pointer to **string** | Phase of payment processing for the associated invoice when you made this request.  These phases include:  | Phase Value | Reason | |---|---| | &#x60;CANCELLED&#x60; | Customer or MongoDB cancelled the payment. | | &#x60;ERROR&#x60; | Issue arose when attempting to complete payment. | | &#x60;FAILED&#x60; | MongoDB tried to charge the credit card without success. | | &#x60;FAILED_AUTHENTICATION&#x60; | Strong Customer Authentication has failed. Confirm that your payment method is authenticated. | | &#x60;FORGIVEN&#x60; | Customer initiated payment which MongoDB later forgave. | | &#x60;INVOICED&#x60; | MongoDB issued an invoice that included this line item. | | &#x60;NEW&#x60; | Customer provided a method of payment, but MongoDB hasn&#39;t tried to charge the credit card. | | &#x60;PAID&#x60; | Customer submitted a successful payment. | | &#x60;PARTIAL_PAID&#x60; | Customer paid for part of this line item. |  | [optional] 
**SubtotalCents** | Pointer to **int64** | Sum of all positive invoice line items contained in this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | Date and time when the customer made an update to this payment attempt. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBilledCents

`func (o *Payment) GetAmountBilledCents() int64`

GetAmountBilledCents returns the AmountBilledCents field if non-nil, zero value otherwise.

### GetAmountBilledCentsOk

`func (o *Payment) GetAmountBilledCentsOk() (*int64, bool)`

GetAmountBilledCentsOk returns a tuple with the AmountBilledCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBilledCents

`func (o *Payment) SetAmountBilledCents(v int64)`

SetAmountBilledCents sets AmountBilledCents field to given value.

### HasAmountBilledCents

`func (o *Payment) HasAmountBilledCents() bool`

HasAmountBilledCents returns a boolean if a field has been set.

### GetAmountPaidCents

`func (o *Payment) GetAmountPaidCents() int64`

GetAmountPaidCents returns the AmountPaidCents field if non-nil, zero value otherwise.

### GetAmountPaidCentsOk

`func (o *Payment) GetAmountPaidCentsOk() (*int64, bool)`

GetAmountPaidCentsOk returns a tuple with the AmountPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidCents

`func (o *Payment) SetAmountPaidCents(v int64)`

SetAmountPaidCents sets AmountPaidCents field to given value.

### HasAmountPaidCents

`func (o *Payment) HasAmountPaidCents() bool`

HasAmountPaidCents returns a boolean if a field has been set.

### GetCreated

`func (o *Payment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Payment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Payment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Payment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Payment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSalesTaxCents

`func (o *Payment) GetSalesTaxCents() int64`

GetSalesTaxCents returns the SalesTaxCents field if non-nil, zero value otherwise.

### GetSalesTaxCentsOk

`func (o *Payment) GetSalesTaxCentsOk() (*int64, bool)`

GetSalesTaxCentsOk returns a tuple with the SalesTaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCents

`func (o *Payment) SetSalesTaxCents(v int64)`

SetSalesTaxCents sets SalesTaxCents field to given value.

### HasSalesTaxCents

`func (o *Payment) HasSalesTaxCents() bool`

HasSalesTaxCents returns a boolean if a field has been set.

### GetStatusName

`func (o *Payment) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *Payment) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *Payment) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *Payment) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetSubtotalCents

`func (o *Payment) GetSubtotalCents() int64`

GetSubtotalCents returns the SubtotalCents field if non-nil, zero value otherwise.

### GetSubtotalCentsOk

`func (o *Payment) GetSubtotalCentsOk() (*int64, bool)`

GetSubtotalCentsOk returns a tuple with the SubtotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalCents

`func (o *Payment) SetSubtotalCents(v int64)`

SetSubtotalCents sets SubtotalCents field to given value.

### HasSubtotalCents

`func (o *Payment) HasSubtotalCents() bool`

HasSubtotalCents returns a boolean if a field has been set.

### GetUpdated

`func (o *Payment) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Payment) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Payment) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Payment) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


