# LineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster that incurred the charge. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**DiscountCents** | Pointer to **int64** | Sum by which MongoDB discounted this line item. MongoDB Cloud expresses this value in cents (100ths of one US Dollar). The resource returns this parameter when a discount applies. | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | Date and time when when MongoDB Cloud finished charging for this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project associated to this line item. | [optional] [readonly] 
**GroupName** | Pointer to **string** | Human-readable label that identifies the project. | [optional] 
**Note** | Pointer to **string** | Comment that applies to this line item. | [optional] [readonly] 
**PercentDiscount** | Pointer to **float32** | Percentage by which MongoDB discounted this line item. The resource returns this parameter when a discount applies. | [optional] [readonly] 
**Quantity** | Pointer to **float64** | Number of units included for the line item. These can be expressions of storage (GB), time (hours), or other units. | [optional] [readonly] 
**Sku** | Pointer to **string** | Human-readable description of the service that this line item provided. This Stock Keeping Unit (SKU) could be the instance type, a support charge, advanced security, or another service. | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud began charging for this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**StitchAppName** | Pointer to **string** | Human-readable label that identifies the Atlas App Services application associated with this line item. | [optional] [readonly] 
**TierLowerBound** | Pointer to **float64** | Lower bound for usage amount range in current SKU tier.   **NOTE**: **lineItems[n].tierLowerBound** appears only if your **lineItems[n].sku** is tiered. | [optional] [readonly] 
**TierUpperBound** | Pointer to **float64** | Upper bound for usage amount range in current SKU tier.   **NOTE**: **lineItems[n].tierUpperBound** appears only if your **lineItems[n].sku** is tiered. | [optional] [readonly] 
**TotalPriceCents** | Pointer to **int64** | Sum of the cost set for this line item. MongoDB Cloud expresses this value in cents (100ths of one US Dollar) and calculates this value as **unitPriceDollars** × **quantity** × 100. | [optional] [readonly] 
**Unit** | Pointer to **string** | Element used to express what **quantity** this line item measures. This value can be elements of time, storage capacity, and the like. | [optional] [readonly] 
**UnitPriceDollars** | Pointer to **float64** | Value per **unit** for this line item expressed in US Dollars. | [optional] [readonly] 

## Methods

### NewLineItem

`func NewLineItem() *LineItem`

NewLineItem instantiates a new LineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *LineItem) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *LineItem) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *LineItem) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *LineItem) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreated

`func (o *LineItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LineItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LineItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LineItem) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiscountCents

`func (o *LineItem) GetDiscountCents() int64`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *LineItem) GetDiscountCentsOk() (*int64, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *LineItem) SetDiscountCents(v int64)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *LineItem) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetEndDate

`func (o *LineItem) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LineItem) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LineItem) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LineItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroupId

`func (o *LineItem) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LineItem) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LineItem) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LineItem) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *LineItem) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *LineItem) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *LineItem) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *LineItem) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetNote

`func (o *LineItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LineItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LineItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *LineItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPercentDiscount

`func (o *LineItem) GetPercentDiscount() float32`

GetPercentDiscount returns the PercentDiscount field if non-nil, zero value otherwise.

### GetPercentDiscountOk

`func (o *LineItem) GetPercentDiscountOk() (*float32, bool)`

GetPercentDiscountOk returns a tuple with the PercentDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentDiscount

`func (o *LineItem) SetPercentDiscount(v float32)`

SetPercentDiscount sets PercentDiscount field to given value.

### HasPercentDiscount

`func (o *LineItem) HasPercentDiscount() bool`

HasPercentDiscount returns a boolean if a field has been set.

### GetQuantity

`func (o *LineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSku

`func (o *LineItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *LineItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *LineItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *LineItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetStartDate

`func (o *LineItem) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LineItem) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LineItem) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LineItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStitchAppName

`func (o *LineItem) GetStitchAppName() string`

GetStitchAppName returns the StitchAppName field if non-nil, zero value otherwise.

### GetStitchAppNameOk

`func (o *LineItem) GetStitchAppNameOk() (*string, bool)`

GetStitchAppNameOk returns a tuple with the StitchAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStitchAppName

`func (o *LineItem) SetStitchAppName(v string)`

SetStitchAppName sets StitchAppName field to given value.

### HasStitchAppName

`func (o *LineItem) HasStitchAppName() bool`

HasStitchAppName returns a boolean if a field has been set.

### GetTierLowerBound

`func (o *LineItem) GetTierLowerBound() float64`

GetTierLowerBound returns the TierLowerBound field if non-nil, zero value otherwise.

### GetTierLowerBoundOk

`func (o *LineItem) GetTierLowerBoundOk() (*float64, bool)`

GetTierLowerBoundOk returns a tuple with the TierLowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierLowerBound

`func (o *LineItem) SetTierLowerBound(v float64)`

SetTierLowerBound sets TierLowerBound field to given value.

### HasTierLowerBound

`func (o *LineItem) HasTierLowerBound() bool`

HasTierLowerBound returns a boolean if a field has been set.

### GetTierUpperBound

`func (o *LineItem) GetTierUpperBound() float64`

GetTierUpperBound returns the TierUpperBound field if non-nil, zero value otherwise.

### GetTierUpperBoundOk

`func (o *LineItem) GetTierUpperBoundOk() (*float64, bool)`

GetTierUpperBoundOk returns a tuple with the TierUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierUpperBound

`func (o *LineItem) SetTierUpperBound(v float64)`

SetTierUpperBound sets TierUpperBound field to given value.

### HasTierUpperBound

`func (o *LineItem) HasTierUpperBound() bool`

HasTierUpperBound returns a boolean if a field has been set.

### GetTotalPriceCents

`func (o *LineItem) GetTotalPriceCents() int64`

GetTotalPriceCents returns the TotalPriceCents field if non-nil, zero value otherwise.

### GetTotalPriceCentsOk

`func (o *LineItem) GetTotalPriceCentsOk() (*int64, bool)`

GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceCents

`func (o *LineItem) SetTotalPriceCents(v int64)`

SetTotalPriceCents sets TotalPriceCents field to given value.

### HasTotalPriceCents

`func (o *LineItem) HasTotalPriceCents() bool`

HasTotalPriceCents returns a boolean if a field has been set.

### GetUnit

`func (o *LineItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LineItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LineItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *LineItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPriceDollars

`func (o *LineItem) GetUnitPriceDollars() float64`

GetUnitPriceDollars returns the UnitPriceDollars field if non-nil, zero value otherwise.

### GetUnitPriceDollarsOk

`func (o *LineItem) GetUnitPriceDollarsOk() (*float64, bool)`

GetUnitPriceDollarsOk returns a tuple with the UnitPriceDollars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceDollars

`func (o *LineItem) SetUnitPriceDollars(v float64)`

SetUnitPriceDollars sets UnitPriceDollars field to given value.

### HasUnitPriceDollars

`func (o *LineItem) HasUnitPriceDollars() bool`

HasUnitPriceDollars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


