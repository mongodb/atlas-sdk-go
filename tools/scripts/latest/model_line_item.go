/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
	"time"
)

// checks if the LineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItem{}

// LineItem One service included in this invoice.
type LineItem struct {
	// Human-readable label that identifies the cluster that incurred the charge.
	ClusterName *string `json:"clusterName,omitempty"`
	// Date and time when MongoDB Cloud created this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Sum by which MongoDB discounted this line item. MongoDB Cloud expresses this value in cents (100ths of one US Dollar). The resource returns this parameter when a discount applies.
	DiscountCents *int64 `json:"discountCents,omitempty"`
	// Date and time when when MongoDB Cloud finished charging for this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	EndDate *time.Time `json:"endDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project associated to this line item.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the project.
	GroupName *string `json:"groupName,omitempty"`
	// Comment that applies to this line item.
	Note *string `json:"note,omitempty"`
	// Percentage by which MongoDB discounted this line item. The resource returns this parameter when a discount applies.
	PercentDiscount *float32 `json:"percentDiscount,omitempty"`
	// Number of units included for the line item. These can be expressions of storage (GB), time (hours), or other units.
	Quantity *float64 `json:"quantity,omitempty"`
	// Human-readable description of the service that this line item provided. This Stock Keeping Unit (SKU) could be the instance type, a support charge, advanced security, or another service.
	Sku *string `json:"sku,omitempty"`
	// Date and time when MongoDB Cloud began charging for this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Human-readable label that identifies the Atlas App Services application associated with this line item.
	StitchAppName *string `json:"stitchAppName,omitempty"`
	// Lower bound for usage amount range in current SKU tier.   **NOTE**: **lineItems[n].tierLowerBound** appears only if your **lineItems[n].sku** is tiered.
	TierLowerBound *float64 `json:"tierLowerBound,omitempty"`
	// Upper bound for usage amount range in current SKU tier.   **NOTE**: **lineItems[n].tierUpperBound** appears only if your **lineItems[n].sku** is tiered.
	TierUpperBound *float64 `json:"tierUpperBound,omitempty"`
	// Sum of the cost set for this line item. MongoDB Cloud expresses this value in cents (100ths of one US Dollar) and calculates this value as **unitPriceDollars** × **quantity** × 100.
	TotalPriceCents *int64 `json:"totalPriceCents,omitempty"`
	// Element used to express what **quantity** this line item measures. This value can be elements of time, storage capacity, and the like.
	Unit *string `json:"unit,omitempty"`
	// Value per **unit** for this line item expressed in US Dollars.
	UnitPriceDollars *float64 `json:"unitPriceDollars,omitempty"`
}

// NewLineItem instantiates a new LineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItem() *LineItem {
	this := LineItem{}
	return &this
}

// NewLineItemWithDefaults instantiates a new LineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemWithDefaults() *LineItem {
	this := LineItem{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *LineItem) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *LineItem) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *LineItem) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *LineItem) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *LineItem) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *LineItem) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDiscountCents returns the DiscountCents field value if set, zero value otherwise.
func (o *LineItem) GetDiscountCents() int64 {
	if o == nil || IsNil(o.DiscountCents) {
		var ret int64
		return ret
	}
	return *o.DiscountCents
}

// GetDiscountCentsOk returns a tuple with the DiscountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetDiscountCentsOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountCents) {
		return nil, false
	}
	return o.DiscountCents, true
}

// HasDiscountCents returns a boolean if a field has been set.
func (o *LineItem) HasDiscountCents() bool {
	if o != nil && !IsNil(o.DiscountCents) {
		return true
	}

	return false
}

// SetDiscountCents gets a reference to the given int64 and assigns it to the DiscountCents field.
func (o *LineItem) SetDiscountCents(v int64) {
	o.DiscountCents = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *LineItem) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *LineItem) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *LineItem) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *LineItem) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *LineItem) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *LineItem) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *LineItem) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *LineItem) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *LineItem) SetGroupName(v string) {
	o.GroupName = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *LineItem) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *LineItem) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *LineItem) SetNote(v string) {
	o.Note = &v
}

// GetPercentDiscount returns the PercentDiscount field value if set, zero value otherwise.
func (o *LineItem) GetPercentDiscount() float32 {
	if o == nil || IsNil(o.PercentDiscount) {
		var ret float32
		return ret
	}
	return *o.PercentDiscount
}

// GetPercentDiscountOk returns a tuple with the PercentDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetPercentDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentDiscount) {
		return nil, false
	}
	return o.PercentDiscount, true
}

// HasPercentDiscount returns a boolean if a field has been set.
func (o *LineItem) HasPercentDiscount() bool {
	if o != nil && !IsNil(o.PercentDiscount) {
		return true
	}

	return false
}

// SetPercentDiscount gets a reference to the given float32 and assigns it to the PercentDiscount field.
func (o *LineItem) SetPercentDiscount(v float32) {
	o.PercentDiscount = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *LineItem) GetQuantity() float64 {
	if o == nil || IsNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *LineItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *LineItem) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *LineItem) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *LineItem) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *LineItem) SetSku(v string) {
	o.Sku = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *LineItem) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *LineItem) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *LineItem) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStitchAppName returns the StitchAppName field value if set, zero value otherwise.
func (o *LineItem) GetStitchAppName() string {
	if o == nil || IsNil(o.StitchAppName) {
		var ret string
		return ret
	}
	return *o.StitchAppName
}

// GetStitchAppNameOk returns a tuple with the StitchAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetStitchAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.StitchAppName) {
		return nil, false
	}
	return o.StitchAppName, true
}

// HasStitchAppName returns a boolean if a field has been set.
func (o *LineItem) HasStitchAppName() bool {
	if o != nil && !IsNil(o.StitchAppName) {
		return true
	}

	return false
}

// SetStitchAppName gets a reference to the given string and assigns it to the StitchAppName field.
func (o *LineItem) SetStitchAppName(v string) {
	o.StitchAppName = &v
}

// GetTierLowerBound returns the TierLowerBound field value if set, zero value otherwise.
func (o *LineItem) GetTierLowerBound() float64 {
	if o == nil || IsNil(o.TierLowerBound) {
		var ret float64
		return ret
	}
	return *o.TierLowerBound
}

// GetTierLowerBoundOk returns a tuple with the TierLowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetTierLowerBoundOk() (*float64, bool) {
	if o == nil || IsNil(o.TierLowerBound) {
		return nil, false
	}
	return o.TierLowerBound, true
}

// HasTierLowerBound returns a boolean if a field has been set.
func (o *LineItem) HasTierLowerBound() bool {
	if o != nil && !IsNil(o.TierLowerBound) {
		return true
	}

	return false
}

// SetTierLowerBound gets a reference to the given float64 and assigns it to the TierLowerBound field.
func (o *LineItem) SetTierLowerBound(v float64) {
	o.TierLowerBound = &v
}

// GetTierUpperBound returns the TierUpperBound field value if set, zero value otherwise.
func (o *LineItem) GetTierUpperBound() float64 {
	if o == nil || IsNil(o.TierUpperBound) {
		var ret float64
		return ret
	}
	return *o.TierUpperBound
}

// GetTierUpperBoundOk returns a tuple with the TierUpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetTierUpperBoundOk() (*float64, bool) {
	if o == nil || IsNil(o.TierUpperBound) {
		return nil, false
	}
	return o.TierUpperBound, true
}

// HasTierUpperBound returns a boolean if a field has been set.
func (o *LineItem) HasTierUpperBound() bool {
	if o != nil && !IsNil(o.TierUpperBound) {
		return true
	}

	return false
}

// SetTierUpperBound gets a reference to the given float64 and assigns it to the TierUpperBound field.
func (o *LineItem) SetTierUpperBound(v float64) {
	o.TierUpperBound = &v
}

// GetTotalPriceCents returns the TotalPriceCents field value if set, zero value otherwise.
func (o *LineItem) GetTotalPriceCents() int64 {
	if o == nil || IsNil(o.TotalPriceCents) {
		var ret int64
		return ret
	}
	return *o.TotalPriceCents
}

// GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetTotalPriceCentsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalPriceCents) {
		return nil, false
	}
	return o.TotalPriceCents, true
}

// HasTotalPriceCents returns a boolean if a field has been set.
func (o *LineItem) HasTotalPriceCents() bool {
	if o != nil && !IsNil(o.TotalPriceCents) {
		return true
	}

	return false
}

// SetTotalPriceCents gets a reference to the given int64 and assigns it to the TotalPriceCents field.
func (o *LineItem) SetTotalPriceCents(v int64) {
	o.TotalPriceCents = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *LineItem) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *LineItem) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *LineItem) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitPriceDollars returns the UnitPriceDollars field value if set, zero value otherwise.
func (o *LineItem) GetUnitPriceDollars() float64 {
	if o == nil || IsNil(o.UnitPriceDollars) {
		var ret float64
		return ret
	}
	return *o.UnitPriceDollars
}

// GetUnitPriceDollarsOk returns a tuple with the UnitPriceDollars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetUnitPriceDollarsOk() (*float64, bool) {
	if o == nil || IsNil(o.UnitPriceDollars) {
		return nil, false
	}
	return o.UnitPriceDollars, true
}

// HasUnitPriceDollars returns a boolean if a field has been set.
func (o *LineItem) HasUnitPriceDollars() bool {
	if o != nil && !IsNil(o.UnitPriceDollars) {
		return true
	}

	return false
}

// SetUnitPriceDollars gets a reference to the given float64 and assigns it to the UnitPriceDollars field.
func (o *LineItem) SetUnitPriceDollars(v float64) {
	o.UnitPriceDollars = &v
}

func (o LineItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: clusterName is readOnly
	// skip: created is readOnly
	// skip: discountCents is readOnly
	// skip: endDate is readOnly
	// skip: groupId is readOnly
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	// skip: note is readOnly
	// skip: percentDiscount is readOnly
	// skip: quantity is readOnly
	// skip: sku is readOnly
	// skip: startDate is readOnly
	// skip: stitchAppName is readOnly
	// skip: tierLowerBound is readOnly
	// skip: tierUpperBound is readOnly
	// skip: totalPriceCents is readOnly
	// skip: unit is readOnly
	// skip: unitPriceDollars is readOnly
	return toSerialize, nil
}

type NullableLineItem struct {
	value *LineItem
	isSet bool
}

func (v NullableLineItem) Get() *LineItem {
	return v.value
}

func (v *NullableLineItem) Set(val *LineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItem(val *LineItem) *NullableLineItem {
	return &NullableLineItem{value: val, isSet: true}
}

func (v NullableLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


