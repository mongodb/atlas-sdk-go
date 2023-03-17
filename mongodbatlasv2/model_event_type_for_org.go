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

// EventTypeForOrg the model 'EventTypeForOrg'
type EventTypeForOrg string

// List of EventTypeForOrg
const (
	EVENTTYPEFORORG_ALERT_ACKNOWLEDGED_AUDIT EventTypeForOrg = "ALERT_ACKNOWLEDGED_AUDIT"
	EVENTTYPEFORORG_ALERT_UNACKNOWLEDGED_AUDIT EventTypeForOrg = "ALERT_UNACKNOWLEDGED_AUDIT"
	EVENTTYPEFORORG_ALERT_CONFIG_DISABLED_AUDIT EventTypeForOrg = "ALERT_CONFIG_DISABLED_AUDIT"
	EVENTTYPEFORORG_ALERT_CONFIG_ENABLED_AUDIT EventTypeForOrg = "ALERT_CONFIG_ENABLED_AUDIT"
	EVENTTYPEFORORG_ALERT_CONFIG_ADDED_AUDIT EventTypeForOrg = "ALERT_CONFIG_ADDED_AUDIT"
	EVENTTYPEFORORG_ALERT_CONFIG_DELETED_AUDIT EventTypeForOrg = "ALERT_CONFIG_DELETED_AUDIT"
	EVENTTYPEFORORG_ALERT_CONFIG_CHANGED_AUDIT EventTypeForOrg = "ALERT_CONFIG_CHANGED_AUDIT"
	EVENTTYPEFORORG_API_KEY_CREATED EventTypeForOrg = "API_KEY_CREATED"
	EVENTTYPEFORORG_API_KEY_DELETED EventTypeForOrg = "API_KEY_DELETED"
	EVENTTYPEFORORG_API_KEY_ACCESS_LIST_ENTRY_ADDED EventTypeForOrg = "API_KEY_ACCESS_LIST_ENTRY_ADDED"
	EVENTTYPEFORORG_API_KEY_ACCESS_LIST_ENTRY_DELETED EventTypeForOrg = "API_KEY_ACCESS_LIST_ENTRY_DELETED"
	EVENTTYPEFORORG_API_KEY_ROLES_CHANGED EventTypeForOrg = "API_KEY_ROLES_CHANGED"
	EVENTTYPEFORORG_API_KEY_DESCRIPTION_CHANGED EventTypeForOrg = "API_KEY_DESCRIPTION_CHANGED"
	EVENTTYPEFORORG_API_KEY_ADDED_TO_GROUP EventTypeForOrg = "API_KEY_ADDED_TO_GROUP"
	EVENTTYPEFORORG_API_KEY_REMOVED_FROM_GROUP EventTypeForOrg = "API_KEY_REMOVED_FROM_GROUP"
	EVENTTYPEFORORG_CHARGE_SUCCEEDED EventTypeForOrg = "CHARGE_SUCCEEDED"
	EVENTTYPEFORORG_CHARGE_FAILED EventTypeForOrg = "CHARGE_FAILED"
	EVENTTYPEFORORG_CHARGE_PROCESSING EventTypeForOrg = "CHARGE_PROCESSING"
	EVENTTYPEFORORG_CHARGE_PENDING_REVERSAL EventTypeForOrg = "CHARGE_PENDING_REVERSAL"
	EVENTTYPEFORORG_BRAINTREE_CHARGE_FAILED EventTypeForOrg = "BRAINTREE_CHARGE_FAILED"
	EVENTTYPEFORORG_INVOICE_CLOSED EventTypeForOrg = "INVOICE_CLOSED"
	EVENTTYPEFORORG_CHECK_PAYMENT_RECEIVED EventTypeForOrg = "CHECK_PAYMENT_RECEIVED"
	EVENTTYPEFORORG_WIRE_TRANSFER_PAYMENT_RECEIVED EventTypeForOrg = "WIRE_TRANSFER_PAYMENT_RECEIVED"
	EVENTTYPEFORORG_DISCOUNT_APPLIED EventTypeForOrg = "DISCOUNT_APPLIED"
	EVENTTYPEFORORG_CREDIT_ISSUED EventTypeForOrg = "CREDIT_ISSUED"
	EVENTTYPEFORORG_CREDIT_PULLED_FWD EventTypeForOrg = "CREDIT_PULLED_FWD"
	EVENTTYPEFORORG_CREDIT_END_DATE_MODIFIED EventTypeForOrg = "CREDIT_END_DATE_MODIFIED"
	EVENTTYPEFORORG_PROMO_CODE_APPLIED EventTypeForOrg = "PROMO_CODE_APPLIED"
	EVENTTYPEFORORG_PAYMENT_FORGIVEN EventTypeForOrg = "PAYMENT_FORGIVEN"
	EVENTTYPEFORORG_REFUND_ISSUED EventTypeForOrg = "REFUND_ISSUED"
	EVENTTYPEFORORG_ACCOUNT_DOWNGRADED EventTypeForOrg = "ACCOUNT_DOWNGRADED"
	EVENTTYPEFORORG_ACCOUNT_UPGRADED EventTypeForOrg = "ACCOUNT_UPGRADED"
	EVENTTYPEFORORG_ACCOUNT_MODIFIED EventTypeForOrg = "ACCOUNT_MODIFIED"
	EVENTTYPEFORORG_SUPPORT_PLAN_ACTIVATED EventTypeForOrg = "SUPPORT_PLAN_ACTIVATED"
	EVENTTYPEFORORG_SUPPORT_PLAN_CANCELLED EventTypeForOrg = "SUPPORT_PLAN_CANCELLED"
	EVENTTYPEFORORG_SUPPORT_PLAN_CANCELLATION_SCHEDULED EventTypeForOrg = "SUPPORT_PLAN_CANCELLATION_SCHEDULED"
	EVENTTYPEFORORG_INITIATE_SALESFORCE_SERVICE_CLOUD_SYNC EventTypeForOrg = "INITIATE_SALESFORCE_SERVICE_CLOUD_SYNC"
	EVENTTYPEFORORG_INVOICE_ADDRESS_CHANGED EventTypeForOrg = "INVOICE_ADDRESS_CHANGED"
	EVENTTYPEFORORG_INVOICE_ADDRESS_ADDED EventTypeForOrg = "INVOICE_ADDRESS_ADDED"
	EVENTTYPEFORORG_PREPAID_PLAN_ACTIVATED EventTypeForOrg = "PREPAID_PLAN_ACTIVATED"
	EVENTTYPEFORORG_ELASTIC_INVOICING_MODE_ACTIVATED EventTypeForOrg = "ELASTIC_INVOICING_MODE_ACTIVATED"
	EVENTTYPEFORORG_ELASTIC_INVOICING_MODE_DEACTIVATED EventTypeForOrg = "ELASTIC_INVOICING_MODE_DEACTIVATED"
	EVENTTYPEFORORG_TERMINATE_PAID_SERVICES EventTypeForOrg = "TERMINATE_PAID_SERVICES"
	EVENTTYPEFORORG_BILLING_EMAIL_ADDRESS_ADDED EventTypeForOrg = "BILLING_EMAIL_ADDRESS_ADDED"
	EVENTTYPEFORORG_BILLING_EMAIL_ADDRESS_CHANGED EventTypeForOrg = "BILLING_EMAIL_ADDRESS_CHANGED"
	EVENTTYPEFORORG_BILLING_EMAIL_ADDRESS_REMOVED EventTypeForOrg = "BILLING_EMAIL_ADDRESS_REMOVED"
	EVENTTYPEFORORG_AWS_BILLING_ACCOUNT_CREDIT_ISSUED EventTypeForOrg = "AWS_BILLING_ACCOUNT_CREDIT_ISSUED"
	EVENTTYPEFORORG_GCP_BILLING_ACCOUNT_CREDIT_ISSUED EventTypeForOrg = "GCP_BILLING_ACCOUNT_CREDIT_ISSUED"
	EVENTTYPEFORORG_CREDIT_SFOLI_MODIFIED EventTypeForOrg = "CREDIT_SFOLI_MODIFIED"
	EVENTTYPEFORORG_CREDIT_SFOLID_MODIFIED EventTypeForOrg = "CREDIT_SFOLID_MODIFIED"
	EVENTTYPEFORORG_CREDIT_AMOUNT_MODIFIED EventTypeForOrg = "CREDIT_AMOUNT_MODIFIED"
	EVENTTYPEFORORG_PREPAID_PLAN_MODIFIED EventTypeForOrg = "PREPAID_PLAN_MODIFIED"
	EVENTTYPEFORORG_AWS_USAGE_REPORTED EventTypeForOrg = "AWS_USAGE_REPORTED"
	EVENTTYPEFORORG_AZURE_USAGE_REPORTED EventTypeForOrg = "AZURE_USAGE_REPORTED"
	EVENTTYPEFORORG_GCP_USAGE_REPORTED EventTypeForOrg = "GCP_USAGE_REPORTED"
	EVENTTYPEFORORG_BECAME_PAYING_ORG EventTypeForOrg = "BECAME_PAYING_ORG"
	EVENTTYPEFORORG_NEW_LINKED_ORG EventTypeForOrg = "NEW_LINKED_ORG"
	EVENTTYPEFORORG_UNLINKED_ORG EventTypeForOrg = "UNLINKED_ORG"
	EVENTTYPEFORORG_ORG_LINKED_TO_PAYING_ORG EventTypeForOrg = "ORG_LINKED_TO_PAYING_ORG"
	EVENTTYPEFORORG_ORG_UNLINKED_FROM_PAYING_ORG EventTypeForOrg = "ORG_UNLINKED_FROM_PAYING_ORG"
	EVENTTYPEFORORG_PAYMENT_UPDATED_THROUGH_API EventTypeForOrg = "PAYMENT_UPDATED_THROUGH_API"
	EVENTTYPEFORORG_AZURE_BILLING_ACCOUNT_CREDIT_ISSUED EventTypeForOrg = "AZURE_BILLING_ACCOUNT_CREDIT_ISSUED"
	EVENTTYPEFORORG_CREDIT_START_DATE_MODIFIED EventTypeForOrg = "CREDIT_START_DATE_MODIFIED"
	EVENTTYPEFORORG_CREDIT_ELASTIC_INVOICING_MODIFIED EventTypeForOrg = "CREDIT_ELASTIC_INVOICING_MODIFIED"
	EVENTTYPEFORORG_CREDIT_TYPE_MODIFIED EventTypeForOrg = "CREDIT_TYPE_MODIFIED"
	EVENTTYPEFORORG_CREDIT_AMOUNT_CENTS_MODIFIED EventTypeForOrg = "CREDIT_AMOUNT_CENTS_MODIFIED"
	EVENTTYPEFORORG_CREDIT_AMOUNT_REMAINING_CENTS_MODIFIED EventTypeForOrg = "CREDIT_AMOUNT_REMAINING_CENTS_MODIFIED"
	EVENTTYPEFORORG_CREDIT_TOTAL_BILLED_CENTS_MODIFIED EventTypeForOrg = "CREDIT_TOTAL_BILLED_CENTS_MODIFIED"
	EVENTTYPEFORORG_CREDIT_AWS_CUSTOMER_ID_MODIFIED EventTypeForOrg = "CREDIT_AWS_CUSTOMER_ID_MODIFIED"
	EVENTTYPEFORORG_CREDIT_AWS_PRODUCT_CODE_MODIFIED EventTypeForOrg = "CREDIT_AWS_PRODUCT_CODE_MODIFIED"
	EVENTTYPEFORORG_CREDIT_GCP_MARKETPLACE_ENTITLEMENT_ID_MODIFIED EventTypeForOrg = "CREDIT_GCP_MARKETPLACE_ENTITLEMENT_ID_MODIFIED"
	EVENTTYPEFORORG_CREDIT_AZURE_SUBSCRIPTION_ID_MODIFIED EventTypeForOrg = "CREDIT_AZURE_SUBSCRIPTION_ID_MODIFIED"
	EVENTTYPEFORORG_FEDERATION_SETTINGS_CREATED EventTypeForOrg = "FEDERATION_SETTINGS_CREATED"
	EVENTTYPEFORORG_FEDERATION_SETTINGS_DELETED EventTypeForOrg = "FEDERATION_SETTINGS_DELETED"
	EVENTTYPEFORORG_FEDERATION_SETTINGS_UPDATED EventTypeForOrg = "FEDERATION_SETTINGS_UPDATED"
	EVENTTYPEFORORG_IDENTITY_PROVIDER_CREATED EventTypeForOrg = "IDENTITY_PROVIDER_CREATED"
	EVENTTYPEFORORG_IDENTITY_PROVIDER_UPDATED EventTypeForOrg = "IDENTITY_PROVIDER_UPDATED"
	EVENTTYPEFORORG_IDENTITY_PROVIDER_DELETED EventTypeForOrg = "IDENTITY_PROVIDER_DELETED"
	EVENTTYPEFORORG_IDENTITY_PROVIDER_ACTIVATED EventTypeForOrg = "IDENTITY_PROVIDER_ACTIVATED"
	EVENTTYPEFORORG_IDENTITY_PROVIDER_DEACTIVATED EventTypeForOrg = "IDENTITY_PROVIDER_DEACTIVATED"
	EVENTTYPEFORORG_DOMAINS_ASSOCIATED EventTypeForOrg = "DOMAINS_ASSOCIATED"
	EVENTTYPEFORORG_DOMAIN_CREATED EventTypeForOrg = "DOMAIN_CREATED"
	EVENTTYPEFORORG_DOMAIN_DELETED EventTypeForOrg = "DOMAIN_DELETED"
	EVENTTYPEFORORG_DOMAIN_VERIFIED EventTypeForOrg = "DOMAIN_VERIFIED"
	EVENTTYPEFORORG_ORG_SETTINGS_CONFIGURED EventTypeForOrg = "ORG_SETTINGS_CONFIGURED"
	EVENTTYPEFORORG_ORG_SETTINGS_UPDATED EventTypeForOrg = "ORG_SETTINGS_UPDATED"
	EVENTTYPEFORORG_ORG_SETTINGS_DELETED EventTypeForOrg = "ORG_SETTINGS_DELETED"
	EVENTTYPEFORORG_RESTRICT_ORG_MEMBERSHIP_ENABLED EventTypeForOrg = "RESTRICT_ORG_MEMBERSHIP_ENABLED"
	EVENTTYPEFORORG_RESTRICT_ORG_MEMBERSHIP_DISABLED EventTypeForOrg = "RESTRICT_ORG_MEMBERSHIP_DISABLED"
	EVENTTYPEFORORG_ROLE_MAPPING_CREATED EventTypeForOrg = "ROLE_MAPPING_CREATED"
	EVENTTYPEFORORG_ROLE_MAPPING_UPDATED EventTypeForOrg = "ROLE_MAPPING_UPDATED"
	EVENTTYPEFORORG_ROLE_MAPPING_DELETED EventTypeForOrg = "ROLE_MAPPING_DELETED"
	EVENTTYPEFORORG_GROUP_DELETED EventTypeForOrg = "GROUP_DELETED"
	EVENTTYPEFORORG_GROUP_CREATED EventTypeForOrg = "GROUP_CREATED"
	EVENTTYPEFORORG_GROUP_MOVED EventTypeForOrg = "GROUP_MOVED"
	EVENTTYPEFORORG_MLAB_MIGRATION_COMPLETED EventTypeForOrg = "MLAB_MIGRATION_COMPLETED"
	EVENTTYPEFORORG_MLAB_MIGRATION_TARGET_CLUSTER_CREATED EventTypeForOrg = "MLAB_MIGRATION_TARGET_CLUSTER_CREATED"
	EVENTTYPEFORORG_MLAB_MIGRATION_DATABASE_USERS_IMPORTED EventTypeForOrg = "MLAB_MIGRATION_DATABASE_USERS_IMPORTED"
	EVENTTYPEFORORG_MLAB_MIGRATION_IP_WHITELIST_IMPORTED EventTypeForOrg = "MLAB_MIGRATION_IP_WHITELIST_IMPORTED"
	EVENTTYPEFORORG_MLAB_MIGRATION_TARGET_CLUSTER_SET EventTypeForOrg = "MLAB_MIGRATION_TARGET_CLUSTER_SET"
	EVENTTYPEFORORG_MLAB_MIGRATION_DATABASE_RENAMED EventTypeForOrg = "MLAB_MIGRATION_DATABASE_RENAMED"
	EVENTTYPEFORORG_MLAB_MIGRATION_LIVE_IMPORT_STARTED EventTypeForOrg = "MLAB_MIGRATION_LIVE_IMPORT_STARTED"
	EVENTTYPEFORORG_MLAB_MIGRATION_LIVE_IMPORT_READY_FOR_CUTOVER EventTypeForOrg = "MLAB_MIGRATION_LIVE_IMPORT_READY_FOR_CUTOVER"
	EVENTTYPEFORORG_MLAB_MIGRATION_LIVE_IMPORT_CUTOVER_COMPLETE EventTypeForOrg = "MLAB_MIGRATION_LIVE_IMPORT_CUTOVER_COMPLETE"
	EVENTTYPEFORORG_MLAB_MIGRATION_LIVE_IMPORT_ERROR EventTypeForOrg = "MLAB_MIGRATION_LIVE_IMPORT_ERROR"
	EVENTTYPEFORORG_MLAB_MIGRATION_LIVE_IMPORT_CANCELLED EventTypeForOrg = "MLAB_MIGRATION_LIVE_IMPORT_CANCELLED"
	EVENTTYPEFORORG_MLAB_MIGRATION_DUMP_AND_RESTORE_TEST_STARTED EventTypeForOrg = "MLAB_MIGRATION_DUMP_AND_RESTORE_TEST_STARTED"
	EVENTTYPEFORORG_MLAB_MIGRATION_DUMP_AND_RESTORE_TEST_SKIPPED EventTypeForOrg = "MLAB_MIGRATION_DUMP_AND_RESTORE_TEST_SKIPPED"
	EVENTTYPEFORORG_MLAB_MIGRATION_DUMP_AND_RESTORE_STARTED EventTypeForOrg = "MLAB_MIGRATION_DUMP_AND_RESTORE_STARTED"
	EVENTTYPEFORORG_MLAB_MIGRATION_SUPPORT_PLAN_SELECTED EventTypeForOrg = "MLAB_MIGRATION_SUPPORT_PLAN_SELECTED"
	EVENTTYPEFORORG_MLAB_MIGRATION_SUPPORT_PLAN_OPTED_OUT EventTypeForOrg = "MLAB_MIGRATION_SUPPORT_PLAN_OPTED_OUT"
	EVENTTYPEFORORG_ORG_LIMIT_UPDATED EventTypeForOrg = "ORG_LIMIT_UPDATED"
	EVENTTYPEFORORG_ORG_CREATED EventTypeForOrg = "ORG_CREATED"
	EVENTTYPEFORORG_ORG_CREDIT_CARD_ADDED EventTypeForOrg = "ORG_CREDIT_CARD_ADDED"
	EVENTTYPEFORORG_ORG_CREDIT_CARD_UPDATED EventTypeForOrg = "ORG_CREDIT_CARD_UPDATED"
	EVENTTYPEFORORG_ORG_CREDIT_CARD_CURRENT EventTypeForOrg = "ORG_CREDIT_CARD_CURRENT"
	EVENTTYPEFORORG_ORG_CREDIT_CARD_ABOUT_TO_EXPIRE EventTypeForOrg = "ORG_CREDIT_CARD_ABOUT_TO_EXPIRE"
	EVENTTYPEFORORG_ORG_PAYPAL_LINKED EventTypeForOrg = "ORG_PAYPAL_LINKED"
	EVENTTYPEFORORG_ORG_PAYPAL_UPDATED EventTypeForOrg = "ORG_PAYPAL_UPDATED"
	EVENTTYPEFORORG_ORG_PAYPAL_CANCELLED EventTypeForOrg = "ORG_PAYPAL_CANCELLED"
	EVENTTYPEFORORG_ORG_OVERRIDE_PAYMENT_METHOD_ADDED EventTypeForOrg = "ORG_OVERRIDE_PAYMENT_METHOD_ADDED"
	EVENTTYPEFORORG_ORG_ACTIVATED EventTypeForOrg = "ORG_ACTIVATED"
	EVENTTYPEFORORG_ORG_TEMPORARILY_ACTIVATED EventTypeForOrg = "ORG_TEMPORARILY_ACTIVATED"
	EVENTTYPEFORORG_ORG_SUSPENDED EventTypeForOrg = "ORG_SUSPENDED"
	EVENTTYPEFORORG_ORG_ADMIN_SUSPENDED EventTypeForOrg = "ORG_ADMIN_SUSPENDED"
	EVENTTYPEFORORG_ORG_ADMIN_LOCKED EventTypeForOrg = "ORG_ADMIN_LOCKED"
	EVENTTYPEFORORG_ORG_CLUSTERS_DELETED EventTypeForOrg = "ORG_CLUSTERS_DELETED"
	EVENTTYPEFORORG_ORG_CLUSTERS_PAUSED EventTypeForOrg = "ORG_CLUSTERS_PAUSED"
	EVENTTYPEFORORG_ORG_LOCKED EventTypeForOrg = "ORG_LOCKED"
	EVENTTYPEFORORG_ORG_RENAMED EventTypeForOrg = "ORG_RENAMED"
	EVENTTYPEFORORG_ALL_ORG_USERS_HAVE_MFA EventTypeForOrg = "ALL_ORG_USERS_HAVE_MFA"
	EVENTTYPEFORORG_ORG_USERS_WITHOUT_MFA EventTypeForOrg = "ORG_USERS_WITHOUT_MFA"
	EVENTTYPEFORORG_ORG_INVOICE_UNDER_THRESHOLD EventTypeForOrg = "ORG_INVOICE_UNDER_THRESHOLD"
	EVENTTYPEFORORG_ORG_INVOICE_OVER_THRESHOLD EventTypeForOrg = "ORG_INVOICE_OVER_THRESHOLD"
	EVENTTYPEFORORG_ORG_DAILY_BILL_UNDER_THRESHOLD EventTypeForOrg = "ORG_DAILY_BILL_UNDER_THRESHOLD"
	EVENTTYPEFORORG_ORG_DAILY_BILL_OVER_THRESHOLD EventTypeForOrg = "ORG_DAILY_BILL_OVER_THRESHOLD"
	EVENTTYPEFORORG_ORG_GROUP_CHARGES_UNDER_THRESHOLD EventTypeForOrg = "ORG_GROUP_CHARGES_UNDER_THRESHOLD"
	EVENTTYPEFORORG_ORG_GROUP_CHARGES_OVER_THRESHOLD EventTypeForOrg = "ORG_GROUP_CHARGES_OVER_THRESHOLD"
	EVENTTYPEFORORG_ORG_TWO_FACTOR_AUTH_REQUIRED EventTypeForOrg = "ORG_TWO_FACTOR_AUTH_REQUIRED"
	EVENTTYPEFORORG_ORG_TWO_FACTOR_AUTH_OPTIONAL EventTypeForOrg = "ORG_TWO_FACTOR_AUTH_OPTIONAL"
	EVENTTYPEFORORG_ORG_PUBLIC_API_ACCESS_LIST_REQUIRED EventTypeForOrg = "ORG_PUBLIC_API_ACCESS_LIST_REQUIRED"
	EVENTTYPEFORORG_ORG_PUBLIC_API_ACCESS_LIST_NOT_REQUIRED EventTypeForOrg = "ORG_PUBLIC_API_ACCESS_LIST_NOT_REQUIRED"
	EVENTTYPEFORORG_ORG_EMPLOYEE_ACCESS_RESTRICTED EventTypeForOrg = "ORG_EMPLOYEE_ACCESS_RESTRICTED"
	EVENTTYPEFORORG_ORG_EMPLOYEE_ACCESS_UNRESTRICTED EventTypeForOrg = "ORG_EMPLOYEE_ACCESS_UNRESTRICTED"
	EVENTTYPEFORORG_ORG_SFDC_ACCOUNT_ID_CHANGED EventTypeForOrg = "ORG_SFDC_ACCOUNT_ID_CHANGED"
	EVENTTYPEFORORG_ORG_CONNECTED_TO_MLAB EventTypeForOrg = "ORG_CONNECTED_TO_MLAB"
	EVENTTYPEFORORG_ORG_DISCONNECTED_FROM_MLAB EventTypeForOrg = "ORG_DISCONNECTED_FROM_MLAB"
	EVENTTYPEFORORG_ORG_IDP_CERTIFICATE_CURRENT EventTypeForOrg = "ORG_IDP_CERTIFICATE_CURRENT"
	EVENTTYPEFORORG_ORG_IDP_CERTIFICATE_ABOUT_TO_EXPIRE EventTypeForOrg = "ORG_IDP_CERTIFICATE_ABOUT_TO_EXPIRE"
	EVENTTYPEFORORG_ORG_CONNECTED_TO_VERCEL EventTypeForOrg = "ORG_CONNECTED_TO_VERCEL"
	EVENTTYPEFORORG_ORG_DISCONNECTED_FROM_VERCEL EventTypeForOrg = "ORG_DISCONNECTED_FROM_VERCEL"
	EVENTTYPEFORORG_ORG_DISCONNECTED_TO_VERCEL EventTypeForOrg = "ORG_DISCONNECTED_TO_VERCEL"
	EVENTTYPEFORORG_ORG_CONNECTION_UNINSTALLED_FROM_VERCEL EventTypeForOrg = "ORG_CONNECTION_UNINSTALLED_FROM_VERCEL"
	EVENTTYPEFORORG_ORG_UI_IP_ACCESS_LIST_ENABLED EventTypeForOrg = "ORG_UI_IP_ACCESS_LIST_ENABLED"
	EVENTTYPEFORORG_ORG_UI_IP_ACCESS_LIST_DISABLED EventTypeForOrg = "ORG_UI_IP_ACCESS_LIST_DISABLED"
	EVENTTYPEFORORG_ORG_EDITED_UI_IP_ACCESS_LIST_ENTRIES EventTypeForOrg = "ORG_EDITED_UI_IP_ACCESS_LIST_ENTRIES"
	EVENTTYPEFORORG_AWS_SELF_SERVE_ACCOUNT_LINKED EventTypeForOrg = "AWS_SELF_SERVE_ACCOUNT_LINKED"
	EVENTTYPEFORORG_AWS_SELF_SERVE_ACCOUNT_LINK_PENDING EventTypeForOrg = "AWS_SELF_SERVE_ACCOUNT_LINK_PENDING"
	EVENTTYPEFORORG_AWS_SELF_SERVE_ACCOUNT_CANCELLED EventTypeForOrg = "AWS_SELF_SERVE_ACCOUNT_CANCELLED"
	EVENTTYPEFORORG_AWS_SELF_SERVE_ACCOUNT_LINK_FAILED EventTypeForOrg = "AWS_SELF_SERVE_ACCOUNT_LINK_FAILED"
	EVENTTYPEFORORG_GCP_SELF_SERVE_ACCOUNT_LINK_PENDING EventTypeForOrg = "GCP_SELF_SERVE_ACCOUNT_LINK_PENDING"
	EVENTTYPEFORORG_GCP_SELF_SERVE_ACCOUNT_LINK_FAILED EventTypeForOrg = "GCP_SELF_SERVE_ACCOUNT_LINK_FAILED"
	EVENTTYPEFORORG_AZURE_SELF_SERVE_ACCOUNT_LINKED EventTypeForOrg = "AZURE_SELF_SERVE_ACCOUNT_LINKED"
	EVENTTYPEFORORG_AZURE_SELF_SERVE_ACCOUNT_LINK_PENDING EventTypeForOrg = "AZURE_SELF_SERVE_ACCOUNT_LINK_PENDING"
	EVENTTYPEFORORG_AZURE_SELF_SERVE_ACCOUNT_CANCELLED EventTypeForOrg = "AZURE_SELF_SERVE_ACCOUNT_CANCELLED"
	EVENTTYPEFORORG_AZURE_SELF_SERVE_ACCOUNT_LINK_FAILED EventTypeForOrg = "AZURE_SELF_SERVE_ACCOUNT_LINK_FAILED"
	EVENTTYPEFORORG_GCP_SELF_SERVE_ACCOUNT_LINKED EventTypeForOrg = "GCP_SELF_SERVE_ACCOUNT_LINKED"
	EVENTTYPEFORORG_GCP_SELF_SERVE_ACCOUNT_CANCELLED EventTypeForOrg = "GCP_SELF_SERVE_ACCOUNT_CANCELLED"
	EVENTTYPEFORORG_ORG_POLICY_CREATED EventTypeForOrg = "ORG_POLICY_CREATED"
	EVENTTYPEFORORG_ORG_POLICY_DELETED EventTypeForOrg = "ORG_POLICY_DELETED"
	EVENTTYPEFORORG_ORG_POLICY_EDITED EventTypeForOrg = "ORG_POLICY_EDITED"
	EVENTTYPEFORORG_ORG_POLICY_CLONED EventTypeForOrg = "ORG_POLICY_CLONED"
	EVENTTYPEFORORG_SUPPORT_EMAILS_SENT_SUCCESSFULLY EventTypeForOrg = "SUPPORT_EMAILS_SENT_SUCCESSFULLY"
	EVENTTYPEFORORG_SUPPORT_EMAILS_SENT_FAILURE EventTypeForOrg = "SUPPORT_EMAILS_SENT_FAILURE"
	EVENTTYPEFORORG_TEAM_CREATED EventTypeForOrg = "TEAM_CREATED"
	EVENTTYPEFORORG_TEAM_DELETED EventTypeForOrg = "TEAM_DELETED"
	EVENTTYPEFORORG_TEAM_UPDATED EventTypeForOrg = "TEAM_UPDATED"
	EVENTTYPEFORORG_TEAM_NAME_CHANGED EventTypeForOrg = "TEAM_NAME_CHANGED"
	EVENTTYPEFORORG_TEAM_ADDED_TO_GROUP EventTypeForOrg = "TEAM_ADDED_TO_GROUP"
	EVENTTYPEFORORG_TEAM_REMOVED_FROM_GROUP EventTypeForOrg = "TEAM_REMOVED_FROM_GROUP"
	EVENTTYPEFORORG_TEAM_ROLES_MODIFIED EventTypeForOrg = "TEAM_ROLES_MODIFIED"
	EVENTTYPEFORORG_JOINED_ORG EventTypeForOrg = "JOINED_ORG"
	EVENTTYPEFORORG_JOINED_TEAM EventTypeForOrg = "JOINED_TEAM"
	EVENTTYPEFORORG_INVITED_TO_ORG EventTypeForOrg = "INVITED_TO_ORG"
	EVENTTYPEFORORG_ORG_INVITATION_DELETED EventTypeForOrg = "ORG_INVITATION_DELETED"
	EVENTTYPEFORORG_REMOVED_FROM_ORG EventTypeForOrg = "REMOVED_FROM_ORG"
	EVENTTYPEFORORG_REMOVED_FROM_TEAM EventTypeForOrg = "REMOVED_FROM_TEAM"
	EVENTTYPEFORORG_USER_ROLES_CHANGED_AUDIT EventTypeForOrg = "USER_ROLES_CHANGED_AUDIT"
	EVENTTYPEFORORG_ORG_FLEX_CONSULTING_PURCHASED EventTypeForOrg = "ORG_FLEX_CONSULTING_PURCHASED"
	EVENTTYPEFORORG_ORG_FLEX_CONSULTING_PURCHASE_FAILED EventTypeForOrg = "ORG_FLEX_CONSULTING_PURCHASE_FAILED"
)

// All allowed values of EventTypeForOrg enum
var AllowedEventTypeForOrgEnumValues = []EventTypeForOrg{
	"ALERT_ACKNOWLEDGED_AUDIT",
	"ALERT_UNACKNOWLEDGED_AUDIT",
	"ALERT_CONFIG_DISABLED_AUDIT",
	"ALERT_CONFIG_ENABLED_AUDIT",
	"ALERT_CONFIG_ADDED_AUDIT",
	"ALERT_CONFIG_DELETED_AUDIT",
	"ALERT_CONFIG_CHANGED_AUDIT",
	"API_KEY_CREATED",
	"API_KEY_DELETED",
	"API_KEY_ACCESS_LIST_ENTRY_ADDED",
	"API_KEY_ACCESS_LIST_ENTRY_DELETED",
	"API_KEY_ROLES_CHANGED",
	"API_KEY_DESCRIPTION_CHANGED",
	"API_KEY_ADDED_TO_GROUP",
	"API_KEY_REMOVED_FROM_GROUP",
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
	"CREDIT_END_DATE_MODIFIED",
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
	"AWS_BILLING_ACCOUNT_CREDIT_ISSUED",
	"GCP_BILLING_ACCOUNT_CREDIT_ISSUED",
	"CREDIT_SFOLI_MODIFIED",
	"CREDIT_SFOLID_MODIFIED",
	"CREDIT_AMOUNT_MODIFIED",
	"PREPAID_PLAN_MODIFIED",
	"AWS_USAGE_REPORTED",
	"AZURE_USAGE_REPORTED",
	"GCP_USAGE_REPORTED",
	"BECAME_PAYING_ORG",
	"NEW_LINKED_ORG",
	"UNLINKED_ORG",
	"ORG_LINKED_TO_PAYING_ORG",
	"ORG_UNLINKED_FROM_PAYING_ORG",
	"PAYMENT_UPDATED_THROUGH_API",
	"AZURE_BILLING_ACCOUNT_CREDIT_ISSUED",
	"CREDIT_START_DATE_MODIFIED",
	"CREDIT_ELASTIC_INVOICING_MODIFIED",
	"CREDIT_TYPE_MODIFIED",
	"CREDIT_AMOUNT_CENTS_MODIFIED",
	"CREDIT_AMOUNT_REMAINING_CENTS_MODIFIED",
	"CREDIT_TOTAL_BILLED_CENTS_MODIFIED",
	"CREDIT_AWS_CUSTOMER_ID_MODIFIED",
	"CREDIT_AWS_PRODUCT_CODE_MODIFIED",
	"CREDIT_GCP_MARKETPLACE_ENTITLEMENT_ID_MODIFIED",
	"CREDIT_AZURE_SUBSCRIPTION_ID_MODIFIED",
	"FEDERATION_SETTINGS_CREATED",
	"FEDERATION_SETTINGS_DELETED",
	"FEDERATION_SETTINGS_UPDATED",
	"IDENTITY_PROVIDER_CREATED",
	"IDENTITY_PROVIDER_UPDATED",
	"IDENTITY_PROVIDER_DELETED",
	"IDENTITY_PROVIDER_ACTIVATED",
	"IDENTITY_PROVIDER_DEACTIVATED",
	"DOMAINS_ASSOCIATED",
	"DOMAIN_CREATED",
	"DOMAIN_DELETED",
	"DOMAIN_VERIFIED",
	"ORG_SETTINGS_CONFIGURED",
	"ORG_SETTINGS_UPDATED",
	"ORG_SETTINGS_DELETED",
	"RESTRICT_ORG_MEMBERSHIP_ENABLED",
	"RESTRICT_ORG_MEMBERSHIP_DISABLED",
	"ROLE_MAPPING_CREATED",
	"ROLE_MAPPING_UPDATED",
	"ROLE_MAPPING_DELETED",
	"GROUP_DELETED",
	"GROUP_CREATED",
	"GROUP_MOVED",
	"MLAB_MIGRATION_COMPLETED",
	"MLAB_MIGRATION_TARGET_CLUSTER_CREATED",
	"MLAB_MIGRATION_DATABASE_USERS_IMPORTED",
	"MLAB_MIGRATION_IP_WHITELIST_IMPORTED",
	"MLAB_MIGRATION_TARGET_CLUSTER_SET",
	"MLAB_MIGRATION_DATABASE_RENAMED",
	"MLAB_MIGRATION_LIVE_IMPORT_STARTED",
	"MLAB_MIGRATION_LIVE_IMPORT_READY_FOR_CUTOVER",
	"MLAB_MIGRATION_LIVE_IMPORT_CUTOVER_COMPLETE",
	"MLAB_MIGRATION_LIVE_IMPORT_ERROR",
	"MLAB_MIGRATION_LIVE_IMPORT_CANCELLED",
	"MLAB_MIGRATION_DUMP_AND_RESTORE_TEST_STARTED",
	"MLAB_MIGRATION_DUMP_AND_RESTORE_TEST_SKIPPED",
	"MLAB_MIGRATION_DUMP_AND_RESTORE_STARTED",
	"MLAB_MIGRATION_SUPPORT_PLAN_SELECTED",
	"MLAB_MIGRATION_SUPPORT_PLAN_OPTED_OUT",
	"ORG_LIMIT_UPDATED",
	"ORG_CREATED",
	"ORG_CREDIT_CARD_ADDED",
	"ORG_CREDIT_CARD_UPDATED",
	"ORG_CREDIT_CARD_CURRENT",
	"ORG_CREDIT_CARD_ABOUT_TO_EXPIRE",
	"ORG_PAYPAL_LINKED",
	"ORG_PAYPAL_UPDATED",
	"ORG_PAYPAL_CANCELLED",
	"ORG_OVERRIDE_PAYMENT_METHOD_ADDED",
	"ORG_ACTIVATED",
	"ORG_TEMPORARILY_ACTIVATED",
	"ORG_SUSPENDED",
	"ORG_ADMIN_SUSPENDED",
	"ORG_ADMIN_LOCKED",
	"ORG_CLUSTERS_DELETED",
	"ORG_CLUSTERS_PAUSED",
	"ORG_LOCKED",
	"ORG_RENAMED",
	"ALL_ORG_USERS_HAVE_MFA",
	"ORG_USERS_WITHOUT_MFA",
	"ORG_INVOICE_UNDER_THRESHOLD",
	"ORG_INVOICE_OVER_THRESHOLD",
	"ORG_DAILY_BILL_UNDER_THRESHOLD",
	"ORG_DAILY_BILL_OVER_THRESHOLD",
	"ORG_GROUP_CHARGES_UNDER_THRESHOLD",
	"ORG_GROUP_CHARGES_OVER_THRESHOLD",
	"ORG_TWO_FACTOR_AUTH_REQUIRED",
	"ORG_TWO_FACTOR_AUTH_OPTIONAL",
	"ORG_PUBLIC_API_ACCESS_LIST_REQUIRED",
	"ORG_PUBLIC_API_ACCESS_LIST_NOT_REQUIRED",
	"ORG_EMPLOYEE_ACCESS_RESTRICTED",
	"ORG_EMPLOYEE_ACCESS_UNRESTRICTED",
	"ORG_SFDC_ACCOUNT_ID_CHANGED",
	"ORG_CONNECTED_TO_MLAB",
	"ORG_DISCONNECTED_FROM_MLAB",
	"ORG_IDP_CERTIFICATE_CURRENT",
	"ORG_IDP_CERTIFICATE_ABOUT_TO_EXPIRE",
	"ORG_CONNECTED_TO_VERCEL",
	"ORG_DISCONNECTED_FROM_VERCEL",
	"ORG_DISCONNECTED_TO_VERCEL",
	"ORG_CONNECTION_UNINSTALLED_FROM_VERCEL",
	"ORG_UI_IP_ACCESS_LIST_ENABLED",
	"ORG_UI_IP_ACCESS_LIST_DISABLED",
	"ORG_EDITED_UI_IP_ACCESS_LIST_ENTRIES",
	"AWS_SELF_SERVE_ACCOUNT_LINKED",
	"AWS_SELF_SERVE_ACCOUNT_LINK_PENDING",
	"AWS_SELF_SERVE_ACCOUNT_CANCELLED",
	"AWS_SELF_SERVE_ACCOUNT_LINK_FAILED",
	"GCP_SELF_SERVE_ACCOUNT_LINK_PENDING",
	"GCP_SELF_SERVE_ACCOUNT_LINK_FAILED",
	"AZURE_SELF_SERVE_ACCOUNT_LINKED",
	"AZURE_SELF_SERVE_ACCOUNT_LINK_PENDING",
	"AZURE_SELF_SERVE_ACCOUNT_CANCELLED",
	"AZURE_SELF_SERVE_ACCOUNT_LINK_FAILED",
	"GCP_SELF_SERVE_ACCOUNT_LINKED",
	"GCP_SELF_SERVE_ACCOUNT_CANCELLED",
	"ORG_POLICY_CREATED",
	"ORG_POLICY_DELETED",
	"ORG_POLICY_EDITED",
	"ORG_POLICY_CLONED",
	"SUPPORT_EMAILS_SENT_SUCCESSFULLY",
	"SUPPORT_EMAILS_SENT_FAILURE",
	"TEAM_CREATED",
	"TEAM_DELETED",
	"TEAM_UPDATED",
	"TEAM_NAME_CHANGED",
	"TEAM_ADDED_TO_GROUP",
	"TEAM_REMOVED_FROM_GROUP",
	"TEAM_ROLES_MODIFIED",
	"JOINED_ORG",
	"JOINED_TEAM",
	"INVITED_TO_ORG",
	"ORG_INVITATION_DELETED",
	"REMOVED_FROM_ORG",
	"REMOVED_FROM_TEAM",
	"USER_ROLES_CHANGED_AUDIT",
	"ORG_FLEX_CONSULTING_PURCHASED",
	"ORG_FLEX_CONSULTING_PURCHASE_FAILED",
}

func (v *EventTypeForOrg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTypeForOrg(value)
	for _, existing := range AllowedEventTypeForOrgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTypeForOrg", value)
}

// NewEventTypeForOrgFromValue returns a pointer to a valid EventTypeForOrg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeForOrgFromValue(v string) (*EventTypeForOrg, error) {
	ev := EventTypeForOrg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventTypeForOrg: valid values are %v", v, AllowedEventTypeForOrgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventTypeForOrg) IsValid() bool {
	for _, existing := range AllowedEventTypeForOrgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventTypeForOrg value
func (v EventTypeForOrg) Ptr() *EventTypeForOrg {
	return &v
}


