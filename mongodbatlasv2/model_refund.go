/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the Refund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Refund{}

// Refund One payment that MongoDB returned to the organization for this invoice.
type Refund struct {
	// Sum of the funds returned to the specified organization expressed in cents (100th of US Dollar).
	AmountCents *int64 `json:"amountCents,omitempty"`
	// Date and time when MongoDB Cloud created this refund. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the payment that the organization had made.
	PaymentId *string `json:"paymentId,omitempty"`
	// Justification that MongoDB accepted to return funds to the organization.
	Reason *string `json:"reason,omitempty"`
}

// NewRefund instantiates a new Refund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefund() *Refund {
	this := Refund{}
	return &this
}

// NewRefundWithDefaults instantiates a new Refund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundWithDefaults() *Refund {
	this := Refund{}
	return &this
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *Refund) GetAmountCents() int64 {
	if o == nil || IsNil(o.AmountCents) {
		var ret int64
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetAmountCentsOk() (*int64, bool) {
	if o == nil || IsNil(o.AmountCents) {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *Refund) HasAmountCents() bool {
	if o != nil && !IsNil(o.AmountCents) {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int64 and assigns it to the AmountCents field.
func (o *Refund) SetAmountCents(v int64) {
	o.AmountCents = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Refund) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Refund) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Refund) SetCreated(v time.Time) {
	o.Created = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *Refund) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *Refund) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *Refund) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Refund) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Refund) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Refund) SetReason(v string) {
	o.Reason = &v
}

func (o Refund) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Refund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableRefund struct {
	value *Refund
	isSet bool
}

func (v NullableRefund) Get() *Refund {
	return v.value
}

func (v *NullableRefund) Set(val *Refund) {
	v.value = val
	v.isSet = true
}

func (v NullableRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefund(val *Refund) *NullableRefund {
	return &NullableRefund{value: val, isSet: true}
}

func (v NullableRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


