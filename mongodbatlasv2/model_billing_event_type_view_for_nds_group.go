/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// BillingEventTypeViewForNdsGroup Unique identifier of event type.
type BillingEventTypeViewForNdsGroup string

// List of BillingEventTypeViewForNdsGroup
const (
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_CARD_CURRENT BillingEventTypeViewForNdsGroup = "CREDIT_CARD_CURRENT"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_CARD_ABOUT_TO_EXPIRE BillingEventTypeViewForNdsGroup = "CREDIT_CARD_ABOUT_TO_EXPIRE"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_PENDING_INVOICE_UNDER_THRESHOLD BillingEventTypeViewForNdsGroup = "PENDING_INVOICE_UNDER_THRESHOLD"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_PENDING_INVOICE_OVER_THRESHOLD BillingEventTypeViewForNdsGroup = "PENDING_INVOICE_OVER_THRESHOLD"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_DAILY_BILL_UNDER_THRESHOLD BillingEventTypeViewForNdsGroup = "DAILY_BILL_UNDER_THRESHOLD"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_DAILY_BILL_OVER_THRESHOLD BillingEventTypeViewForNdsGroup = "DAILY_BILL_OVER_THRESHOLD"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CHARGE_SUCCEEDED BillingEventTypeViewForNdsGroup = "CHARGE_SUCCEEDED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CHARGE_FAILED BillingEventTypeViewForNdsGroup = "CHARGE_FAILED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CHARGE_PROCESSING BillingEventTypeViewForNdsGroup = "CHARGE_PROCESSING"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CHARGE_PENDING_REVERSAL BillingEventTypeViewForNdsGroup = "CHARGE_PENDING_REVERSAL"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_BRAINTREE_CHARGE_FAILED BillingEventTypeViewForNdsGroup = "BRAINTREE_CHARGE_FAILED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_INVOICE_CLOSED BillingEventTypeViewForNdsGroup = "INVOICE_CLOSED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CHECK_PAYMENT_RECEIVED BillingEventTypeViewForNdsGroup = "CHECK_PAYMENT_RECEIVED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_WIRE_TRANSFER_PAYMENT_RECEIVED BillingEventTypeViewForNdsGroup = "WIRE_TRANSFER_PAYMENT_RECEIVED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_DISCOUNT_APPLIED BillingEventTypeViewForNdsGroup = "DISCOUNT_APPLIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_ISSUED BillingEventTypeViewForNdsGroup = "CREDIT_ISSUED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_PULLED_FWD BillingEventTypeViewForNdsGroup = "CREDIT_PULLED_FWD"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_START_DATE_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_START_DATE_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_END_DATE_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_END_DATE_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_ELASTIC_INVOICING_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_ELASTIC_INVOICING_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_TYPE_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_TYPE_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_AMOUNT_CENTS_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_AMOUNT_CENTS_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_AMOUNT_REMAINING_CENTS_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_AMOUNT_REMAINING_CENTS_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_TOTAL_BILLED_CENTS_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_TOTAL_BILLED_CENTS_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_AWS_CUSTOMER_ID_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_AWS_CUSTOMER_ID_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_AWS_PRODUCT_CODE_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_AWS_PRODUCT_CODE_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_GCP_MARKETPLACE_ENTITLEMENT_ID_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_GCP_MARKETPLACE_ENTITLEMENT_ID_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_CREDIT_AZURE_SUBSCRIPTION_ID_MODIFIED BillingEventTypeViewForNdsGroup = "CREDIT_AZURE_SUBSCRIPTION_ID_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_PROMO_CODE_APPLIED BillingEventTypeViewForNdsGroup = "PROMO_CODE_APPLIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_PAYMENT_FORGIVEN BillingEventTypeViewForNdsGroup = "PAYMENT_FORGIVEN"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_REFUND_ISSUED BillingEventTypeViewForNdsGroup = "REFUND_ISSUED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_ACCOUNT_DOWNGRADED BillingEventTypeViewForNdsGroup = "ACCOUNT_DOWNGRADED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_ACCOUNT_UPGRADED BillingEventTypeViewForNdsGroup = "ACCOUNT_UPGRADED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_ACCOUNT_MODIFIED BillingEventTypeViewForNdsGroup = "ACCOUNT_MODIFIED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_SUPPORT_PLAN_ACTIVATED BillingEventTypeViewForNdsGroup = "SUPPORT_PLAN_ACTIVATED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_SUPPORT_PLAN_CANCELLED BillingEventTypeViewForNdsGroup = "SUPPORT_PLAN_CANCELLED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_SUPPORT_PLAN_CANCELLATION_SCHEDULED BillingEventTypeViewForNdsGroup = "SUPPORT_PLAN_CANCELLATION_SCHEDULED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_INITIATE_SALESFORCE_SERVICE_CLOUD_SYNC BillingEventTypeViewForNdsGroup = "INITIATE_SALESFORCE_SERVICE_CLOUD_SYNC"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_INVOICE_ADDRESS_CHANGED BillingEventTypeViewForNdsGroup = "INVOICE_ADDRESS_CHANGED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_INVOICE_ADDRESS_ADDED BillingEventTypeViewForNdsGroup = "INVOICE_ADDRESS_ADDED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_PREPAID_PLAN_ACTIVATED BillingEventTypeViewForNdsGroup = "PREPAID_PLAN_ACTIVATED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_ELASTIC_INVOICING_MODE_ACTIVATED BillingEventTypeViewForNdsGroup = "ELASTIC_INVOICING_MODE_ACTIVATED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_ELASTIC_INVOICING_MODE_DEACTIVATED BillingEventTypeViewForNdsGroup = "ELASTIC_INVOICING_MODE_DEACTIVATED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_TERMINATE_PAID_SERVICES BillingEventTypeViewForNdsGroup = "TERMINATE_PAID_SERVICES"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_BILLING_EMAIL_ADDRESS_ADDED BillingEventTypeViewForNdsGroup = "BILLING_EMAIL_ADDRESS_ADDED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_BILLING_EMAIL_ADDRESS_CHANGED BillingEventTypeViewForNdsGroup = "BILLING_EMAIL_ADDRESS_CHANGED"
	BILLINGEVENTTYPEVIEWFORNDSGROUP_BILLING_EMAIL_ADDRESS_REMOVED BillingEventTypeViewForNdsGroup = "BILLING_EMAIL_ADDRESS_REMOVED"
)

// All allowed values of BillingEventTypeViewForNdsGroup enum
var AllowedBillingEventTypeViewForNdsGroupEnumValues = []BillingEventTypeViewForNdsGroup{
	"CREDIT_CARD_CURRENT",
	"CREDIT_CARD_ABOUT_TO_EXPIRE",
	"PENDING_INVOICE_UNDER_THRESHOLD",
	"PENDING_INVOICE_OVER_THRESHOLD",
	"DAILY_BILL_UNDER_THRESHOLD",
	"DAILY_BILL_OVER_THRESHOLD",
	"CHARGE_SUCCEEDED",
	"CHARGE_FAILED",
	"CHARGE_PROCESSING",
	"CHARGE_PENDING_REVERSAL",
	"BRAINTREE_CHARGE_FAILED",
	"INVOICE_CLOSED",
	"CHECK_PAYMENT_RECEIVED",
	"WIRE_TRANSFER_PAYMENT_RECEIVED",
	"DISCOUNT_APPLIED",
	"CREDIT_ISSUED",
	"CREDIT_PULLED_FWD",
	"CREDIT_START_DATE_MODIFIED",
	"CREDIT_END_DATE_MODIFIED",
	"CREDIT_ELASTIC_INVOICING_MODIFIED",
	"CREDIT_TYPE_MODIFIED",
	"CREDIT_AMOUNT_CENTS_MODIFIED",
	"CREDIT_AMOUNT_REMAINING_CENTS_MODIFIED",
	"CREDIT_TOTAL_BILLED_CENTS_MODIFIED",
	"CREDIT_AWS_CUSTOMER_ID_MODIFIED",
	"CREDIT_AWS_PRODUCT_CODE_MODIFIED",
	"CREDIT_GCP_MARKETPLACE_ENTITLEMENT_ID_MODIFIED",
	"CREDIT_AZURE_SUBSCRIPTION_ID_MODIFIED",
	"PROMO_CODE_APPLIED",
	"PAYMENT_FORGIVEN",
	"REFUND_ISSUED",
	"ACCOUNT_DOWNGRADED",
	"ACCOUNT_UPGRADED",
	"ACCOUNT_MODIFIED",
	"SUPPORT_PLAN_ACTIVATED",
	"SUPPORT_PLAN_CANCELLED",
	"SUPPORT_PLAN_CANCELLATION_SCHEDULED",
	"INITIATE_SALESFORCE_SERVICE_CLOUD_SYNC",
	"INVOICE_ADDRESS_CHANGED",
	"INVOICE_ADDRESS_ADDED",
	"PREPAID_PLAN_ACTIVATED",
	"ELASTIC_INVOICING_MODE_ACTIVATED",
	"ELASTIC_INVOICING_MODE_DEACTIVATED",
	"TERMINATE_PAID_SERVICES",
	"BILLING_EMAIL_ADDRESS_ADDED",
	"BILLING_EMAIL_ADDRESS_CHANGED",
	"BILLING_EMAIL_ADDRESS_REMOVED",
}

func (v *BillingEventTypeViewForNdsGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingEventTypeViewForNdsGroup(value)
	for _, existing := range AllowedBillingEventTypeViewForNdsGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingEventTypeViewForNdsGroup", value)
}

// NewBillingEventTypeViewForNdsGroupFromValue returns a pointer to a valid BillingEventTypeViewForNdsGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingEventTypeViewForNdsGroupFromValue(v string) (*BillingEventTypeViewForNdsGroup, error) {
	ev := BillingEventTypeViewForNdsGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingEventTypeViewForNdsGroup: valid values are %v", v, AllowedBillingEventTypeViewForNdsGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingEventTypeViewForNdsGroup) IsValid() bool {
	for _, existing := range AllowedBillingEventTypeViewForNdsGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingEventTypeViewForNdsGroup value
func (v BillingEventTypeViewForNdsGroup) Ptr() *BillingEventTypeViewForNdsGroup {
	return &v
}


