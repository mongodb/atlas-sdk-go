/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// BillingEventTypeViewAlertableWithThreshold Event type that triggers an alert.
type BillingEventTypeViewAlertableWithThreshold string

// List of BillingEventTypeViewAlertableWithThreshold
const (
	BILLINGEVENTTYPEVIEWALERTABLEWITHTHRESHOLD_PENDING_INVOICE_OVER_THRESHOLD BillingEventTypeViewAlertableWithThreshold = "PENDING_INVOICE_OVER_THRESHOLD"
	BILLINGEVENTTYPEVIEWALERTABLEWITHTHRESHOLD_DAILY_BILL_OVER_THRESHOLD BillingEventTypeViewAlertableWithThreshold = "DAILY_BILL_OVER_THRESHOLD"
)

// All allowed values of BillingEventTypeViewAlertableWithThreshold enum
var AllowedBillingEventTypeViewAlertableWithThresholdEnumValues = []BillingEventTypeViewAlertableWithThreshold{
	"PENDING_INVOICE_OVER_THRESHOLD",
	"DAILY_BILL_OVER_THRESHOLD",
}

func (v *BillingEventTypeViewAlertableWithThreshold) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingEventTypeViewAlertableWithThreshold(value)
	for _, existing := range AllowedBillingEventTypeViewAlertableWithThresholdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingEventTypeViewAlertableWithThreshold", value)
}

// NewBillingEventTypeViewAlertableWithThresholdFromValue returns a pointer to a valid BillingEventTypeViewAlertableWithThreshold
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingEventTypeViewAlertableWithThresholdFromValue(v string) (*BillingEventTypeViewAlertableWithThreshold, error) {
	ev := BillingEventTypeViewAlertableWithThreshold(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingEventTypeViewAlertableWithThreshold: valid values are %v", v, AllowedBillingEventTypeViewAlertableWithThresholdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingEventTypeViewAlertableWithThreshold) IsValid() bool {
	for _, existing := range AllowedBillingEventTypeViewAlertableWithThresholdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingEventTypeViewAlertableWithThreshold value
func (v BillingEventTypeViewAlertableWithThreshold) Ptr() *BillingEventTypeViewAlertableWithThreshold {
	return &v
}


