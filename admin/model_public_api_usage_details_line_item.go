// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// PublicApiUsageDetailsLineItem struct for PublicApiUsageDetailsLineItem
type PublicApiUsageDetailsLineItem struct {
	AdditionalData *AdditionalData `json:"additionalData,omitempty"`
	// Billing date of the line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	BillDate *time.Time `json:"billDate,omitempty"`
	// Cluster associated with the line item.
	ClusterName *string `json:"clusterName,omitempty"`
	// Description of the line item, which can include SKU name and other identifying information.
	Description *string `json:"description,omitempty"`
	// Group id associated with the line item.
	GroupId *string `json:"groupId,omitempty"`
	// Quantity of line item in units associated with SKU.
	// Read only field.
	Quantity *float64 `json:"quantity,omitempty"`
	// Price * quantity in applicable units, expressed as an integral number of cents.
	TotalPriceCents *int64 `json:"totalPriceCents,omitempty"`
	// Price in units associated with the SKU for the line item.
	UnitPriceDollars *float64 `json:"unitPriceDollars,omitempty"`
	// Usage date of the line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	UsageDate *time.Time `json:"usageDate,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PublicApiUsageDetailsLineItem) MarshalJSON() ([]byte, error) {
	type noMethod PublicApiUsageDetailsLineItem
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPublicApiUsageDetailsLineItem instantiates a new PublicApiUsageDetailsLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicApiUsageDetailsLineItem() *PublicApiUsageDetailsLineItem {
	this := PublicApiUsageDetailsLineItem{}
	return &this
}

// NewPublicApiUsageDetailsLineItemWithDefaults instantiates a new PublicApiUsageDetailsLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicApiUsageDetailsLineItemWithDefaults() *PublicApiUsageDetailsLineItem {
	this := PublicApiUsageDetailsLineItem{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetAdditionalData() AdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret AdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetAdditionalDataOk() (*AdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}

	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given AdditionalData and assigns it to the AdditionalData field.
func (o *PublicApiUsageDetailsLineItem) SetAdditionalData(v AdditionalData) {
	o.AdditionalData = &v
	o.NullFields = removeNullField(o.NullFields, "AdditionalData")
}

// SetAdditionalDataNil sets AdditionalData to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetAdditionalDataNil() {
	o.AdditionalData = nil
	o.NullFields = addNullField(o.NullFields, "AdditionalData")
}

// GetBillDate returns the BillDate field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetBillDate() time.Time {
	if o == nil || IsNil(o.BillDate) {
		var ret time.Time
		return ret
	}
	return *o.BillDate
}

// GetBillDateOk returns a tuple with the BillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetBillDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BillDate) {
		return nil, false
	}

	return o.BillDate, true
}

// HasBillDate returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasBillDate() bool {
	if o != nil && !IsNil(o.BillDate) {
		return true
	}

	return false
}

// SetBillDate gets a reference to the given time.Time and assigns it to the BillDate field.
func (o *PublicApiUsageDetailsLineItem) SetBillDate(v time.Time) {
	o.BillDate = &v
	o.NullFields = removeNullField(o.NullFields, "BillDate")
}

// SetBillDateNil sets BillDate to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetBillDateNil() {
	o.BillDate = nil
	o.NullFields = addNullField(o.NullFields, "BillDate")
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *PublicApiUsageDetailsLineItem) SetClusterName(v string) {
	o.ClusterName = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterName")
}

// SetClusterNameNil sets ClusterName to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetClusterNameNil() {
	o.ClusterName = nil
	o.NullFields = addNullField(o.NullFields, "ClusterName")
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PublicApiUsageDetailsLineItem) SetDescription(v string) {
	o.Description = &v
	o.NullFields = removeNullField(o.NullFields, "Description")
}

// SetDescriptionNil sets Description to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetDescriptionNil() {
	o.Description = nil
	o.NullFields = addNullField(o.NullFields, "Description")
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PublicApiUsageDetailsLineItem) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetQuantity returns the Quantity field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetQuantity() float64 {
	if o == nil || IsNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}

	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *PublicApiUsageDetailsLineItem) SetQuantity(v float64) {
	o.Quantity = &v
	o.NullFields = removeNullField(o.NullFields, "Quantity")
}

// SetQuantityNil sets Quantity to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetQuantityNil() {
	o.Quantity = nil
	o.NullFields = addNullField(o.NullFields, "Quantity")
}

// GetTotalPriceCents returns the TotalPriceCents field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetTotalPriceCents() int64 {
	if o == nil || IsNil(o.TotalPriceCents) {
		var ret int64
		return ret
	}
	return *o.TotalPriceCents
}

// GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetTotalPriceCentsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalPriceCents) {
		return nil, false
	}

	return o.TotalPriceCents, true
}

// HasTotalPriceCents returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasTotalPriceCents() bool {
	if o != nil && !IsNil(o.TotalPriceCents) {
		return true
	}

	return false
}

// SetTotalPriceCents gets a reference to the given int64 and assigns it to the TotalPriceCents field.
func (o *PublicApiUsageDetailsLineItem) SetTotalPriceCents(v int64) {
	o.TotalPriceCents = &v
	o.NullFields = removeNullField(o.NullFields, "TotalPriceCents")
}

// SetTotalPriceCentsNil sets TotalPriceCents to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetTotalPriceCentsNil() {
	o.TotalPriceCents = nil
	o.NullFields = addNullField(o.NullFields, "TotalPriceCents")
}

// GetUnitPriceDollars returns the UnitPriceDollars field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetUnitPriceDollars() float64 {
	if o == nil || IsNil(o.UnitPriceDollars) {
		var ret float64
		return ret
	}
	return *o.UnitPriceDollars
}

// GetUnitPriceDollarsOk returns a tuple with the UnitPriceDollars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetUnitPriceDollarsOk() (*float64, bool) {
	if o == nil || IsNil(o.UnitPriceDollars) {
		return nil, false
	}

	return o.UnitPriceDollars, true
}

// HasUnitPriceDollars returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasUnitPriceDollars() bool {
	if o != nil && !IsNil(o.UnitPriceDollars) {
		return true
	}

	return false
}

// SetUnitPriceDollars gets a reference to the given float64 and assigns it to the UnitPriceDollars field.
func (o *PublicApiUsageDetailsLineItem) SetUnitPriceDollars(v float64) {
	o.UnitPriceDollars = &v
	o.NullFields = removeNullField(o.NullFields, "UnitPriceDollars")
}

// SetUnitPriceDollarsNil sets UnitPriceDollars to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetUnitPriceDollarsNil() {
	o.UnitPriceDollars = nil
	o.NullFields = addNullField(o.NullFields, "UnitPriceDollars")
}

// GetUsageDate returns the UsageDate field value if set, zero value otherwise
func (o *PublicApiUsageDetailsLineItem) GetUsageDate() time.Time {
	if o == nil || IsNil(o.UsageDate) {
		var ret time.Time
		return ret
	}
	return *o.UsageDate
}

// GetUsageDateOk returns a tuple with the UsageDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicApiUsageDetailsLineItem) GetUsageDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UsageDate) {
		return nil, false
	}

	return o.UsageDate, true
}

// HasUsageDate returns a boolean if a field has been set.
func (o *PublicApiUsageDetailsLineItem) HasUsageDate() bool {
	if o != nil && !IsNil(o.UsageDate) {
		return true
	}

	return false
}

// SetUsageDate gets a reference to the given time.Time and assigns it to the UsageDate field.
func (o *PublicApiUsageDetailsLineItem) SetUsageDate(v time.Time) {
	o.UsageDate = &v
	o.NullFields = removeNullField(o.NullFields, "UsageDate")
}

// SetUsageDateNil sets UsageDate to an explicit JSON null when marshaled.
func (o *PublicApiUsageDetailsLineItem) SetUsageDateNil() {
	o.UsageDate = nil
	o.NullFields = addNullField(o.NullFields, "UsageDate")
}
