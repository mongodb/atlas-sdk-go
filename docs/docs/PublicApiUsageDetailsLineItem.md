# PublicApiUsageDetailsLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillDate** | Pointer to **time.Time** | Billing date of the line item. | [optional] 
**ClusterName** | Pointer to **string** | Cluster associated with the line item. | [optional] 
**Description** | Pointer to **string** | Description of the line item, which can include SKU name and other identifying information. | [optional] 
**GroupId** | Pointer to **string** | Group id associated with the line item. | [optional] 
**Quantity** | Pointer to **float64** | Quantity of line item in units associated with SKU. | [optional] [readonly] 
**TotalPriceCents** | Pointer to **int64** | Price * quantity in applicable units, expressed as an integral number of cents. | [optional] 
**UnitPriceDollars** | Pointer to **float64** | Price in units associated with the SKU for the line item. | [optional] 
**UsageDate** | Pointer to **time.Time** | Usage date of the line item. | [optional] 

## Methods

### NewPublicApiUsageDetailsLineItem

`func NewPublicApiUsageDetailsLineItem() *PublicApiUsageDetailsLineItem`

NewPublicApiUsageDetailsLineItem instantiates a new PublicApiUsageDetailsLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicApiUsageDetailsLineItemWithDefaults

`func NewPublicApiUsageDetailsLineItemWithDefaults() *PublicApiUsageDetailsLineItem`

NewPublicApiUsageDetailsLineItemWithDefaults instantiates a new PublicApiUsageDetailsLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillDate

`func (o *PublicApiUsageDetailsLineItem) GetBillDate() time.Time`

GetBillDate returns the BillDate field if non-nil, zero value otherwise.

### GetBillDateOk

`func (o *PublicApiUsageDetailsLineItem) GetBillDateOk() (*time.Time, bool)`

GetBillDateOk returns a tuple with the BillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillDate

`func (o *PublicApiUsageDetailsLineItem) SetBillDate(v time.Time)`

SetBillDate sets BillDate field to given value.

### HasBillDate

`func (o *PublicApiUsageDetailsLineItem) HasBillDate() bool`

HasBillDate returns a boolean if a field has been set.
### GetClusterName

`func (o *PublicApiUsageDetailsLineItem) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *PublicApiUsageDetailsLineItem) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *PublicApiUsageDetailsLineItem) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *PublicApiUsageDetailsLineItem) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetDescription

`func (o *PublicApiUsageDetailsLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicApiUsageDetailsLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicApiUsageDetailsLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicApiUsageDetailsLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetGroupId

`func (o *PublicApiUsageDetailsLineItem) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PublicApiUsageDetailsLineItem) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PublicApiUsageDetailsLineItem) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PublicApiUsageDetailsLineItem) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetQuantity

`func (o *PublicApiUsageDetailsLineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PublicApiUsageDetailsLineItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PublicApiUsageDetailsLineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PublicApiUsageDetailsLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.
### GetTotalPriceCents

`func (o *PublicApiUsageDetailsLineItem) GetTotalPriceCents() int64`

GetTotalPriceCents returns the TotalPriceCents field if non-nil, zero value otherwise.

### GetTotalPriceCentsOk

`func (o *PublicApiUsageDetailsLineItem) GetTotalPriceCentsOk() (*int64, bool)`

GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceCents

`func (o *PublicApiUsageDetailsLineItem) SetTotalPriceCents(v int64)`

SetTotalPriceCents sets TotalPriceCents field to given value.

### HasTotalPriceCents

`func (o *PublicApiUsageDetailsLineItem) HasTotalPriceCents() bool`

HasTotalPriceCents returns a boolean if a field has been set.
### GetUnitPriceDollars

`func (o *PublicApiUsageDetailsLineItem) GetUnitPriceDollars() float64`

GetUnitPriceDollars returns the UnitPriceDollars field if non-nil, zero value otherwise.

### GetUnitPriceDollarsOk

`func (o *PublicApiUsageDetailsLineItem) GetUnitPriceDollarsOk() (*float64, bool)`

GetUnitPriceDollarsOk returns a tuple with the UnitPriceDollars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceDollars

`func (o *PublicApiUsageDetailsLineItem) SetUnitPriceDollars(v float64)`

SetUnitPriceDollars sets UnitPriceDollars field to given value.

### HasUnitPriceDollars

`func (o *PublicApiUsageDetailsLineItem) HasUnitPriceDollars() bool`

HasUnitPriceDollars returns a boolean if a field has been set.
### GetUsageDate

`func (o *PublicApiUsageDetailsLineItem) GetUsageDate() time.Time`

GetUsageDate returns the UsageDate field if non-nil, zero value otherwise.

### GetUsageDateOk

`func (o *PublicApiUsageDetailsLineItem) GetUsageDateOk() (*time.Time, bool)`

GetUsageDateOk returns a tuple with the UsageDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDate

`func (o *PublicApiUsageDetailsLineItem) SetUsageDate(v time.Time)`

SetUsageDate sets UsageDate field to given value.

### HasUsageDate

`func (o *PublicApiUsageDetailsLineItem) HasUsageDate() bool`

HasUsageDate returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


