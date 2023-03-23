/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the EventViewForOrg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventViewForOrg{}

// EventViewForOrg struct for EventViewForOrg
type EventViewForOrg struct {
	// Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **userId** parameter.
	ApiKeyId *string `json:"apiKeyId,omitempty"`
	// Date and time when this event occurred. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	EventTypeName *UserEventTypeViewForOrg `json:"eventTypeName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	Id *string `json:"id,omitempty"`
	// Flag that indicates whether a MongoDB employee triggered the specified event.
	IsGlobalAdmin *bool `json:"isGlobalAdmin,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	OrgId *string `json:"orgId,omitempty"`
	// Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **username** parameter.
	PublicKey *string `json:"publicKey,omitempty"`
	Raw *Raw `json:"raw,omitempty"`
	// IPv4 or IPv6 address from which the user triggered this event.
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn't return the **apiKeyId** parameter.
	UserId *string `json:"userId,omitempty"`
	// Email address for the user who triggered this event. If this resource returns this parameter, it doesn't return the **publicApiKey** parameter.
	Username *string `json:"username,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert associated with the event.
	AlertId *string `json:"alertId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert configuration associated with the **alertId**.
	AlertConfigId *string `json:"alertConfigId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies of the invoice associated with the event.
	InvoiceId *string `json:"invoiceId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the invoice payment associated with this event.
	PaymentId *string `json:"paymentId,omitempty"`
	// Entry in the list of source host addresses that the API key accepts and this event targets.
	WhitelistEntry *string `json:"whitelistEntry,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization team associated with this event.
	TeamId *string `json:"teamId,omitempty"`
	// Email address for the console user that this event targets. The resource returns this parameter when `\"eventTypeName\" : \"USER\"`.
	TargetUsername *string `json:"targetUsername,omitempty"`
}

// NewEventViewForOrg instantiates a new EventViewForOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventViewForOrg() *EventViewForOrg {
	this := EventViewForOrg{}
	return &this
}

// NewEventViewForOrgWithDefaults instantiates a new EventViewForOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventViewForOrgWithDefaults() *EventViewForOrg {
	this := EventViewForOrg{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetApiKeyId() string {
	if o == nil || IsNil(o.ApiKeyId) {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetApiKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKeyId) {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasApiKeyId() bool {
	if o != nil && !IsNil(o.ApiKeyId) {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *EventViewForOrg) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EventViewForOrg) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EventViewForOrg) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *EventViewForOrg) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEventTypeName returns the EventTypeName field value if set, zero value otherwise.
func (o *EventViewForOrg) GetEventTypeName() UserEventTypeViewForOrg {
	if o == nil || IsNil(o.EventTypeName) {
		var ret UserEventTypeViewForOrg
		return ret
	}
	return *o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetEventTypeNameOk() (*UserEventTypeViewForOrg, bool) {
	if o == nil || IsNil(o.EventTypeName) {
		return nil, false
	}
	return o.EventTypeName, true
}

// HasEventTypeName returns a boolean if a field has been set.
func (o *EventViewForOrg) HasEventTypeName() bool {
	if o != nil && !IsNil(o.EventTypeName) {
		return true
	}

	return false
}

// SetEventTypeName gets a reference to the given UserEventTypeViewForOrg and assigns it to the EventTypeName field.
func (o *EventViewForOrg) SetEventTypeName(v UserEventTypeViewForOrg) {
	o.EventTypeName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *EventViewForOrg) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventViewForOrg) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventViewForOrg) SetId(v string) {
	o.Id = &v
}

// GetIsGlobalAdmin returns the IsGlobalAdmin field value if set, zero value otherwise.
func (o *EventViewForOrg) GetIsGlobalAdmin() bool {
	if o == nil || IsNil(o.IsGlobalAdmin) {
		var ret bool
		return ret
	}
	return *o.IsGlobalAdmin
}

// GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetIsGlobalAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalAdmin) {
		return nil, false
	}
	return o.IsGlobalAdmin, true
}

// HasIsGlobalAdmin returns a boolean if a field has been set.
func (o *EventViewForOrg) HasIsGlobalAdmin() bool {
	if o != nil && !IsNil(o.IsGlobalAdmin) {
		return true
	}

	return false
}

// SetIsGlobalAdmin gets a reference to the given bool and assigns it to the IsGlobalAdmin field.
func (o *EventViewForOrg) SetIsGlobalAdmin(v bool) {
	o.IsGlobalAdmin = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EventViewForOrg) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EventViewForOrg) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *EventViewForOrg) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *EventViewForOrg) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *EventViewForOrg) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *EventViewForOrg) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *EventViewForOrg) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *EventViewForOrg) GetRaw() Raw {
	if o == nil || IsNil(o.Raw) {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetRawOk() (*Raw, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *EventViewForOrg) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *EventViewForOrg) SetRaw(v Raw) {
	o.Raw = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *EventViewForOrg) GetRemoteAddress() string {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetRemoteAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *EventViewForOrg) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *EventViewForOrg) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EventViewForOrg) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EventViewForOrg) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EventViewForOrg) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EventViewForOrg) SetUsername(v string) {
	o.Username = &v
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetAlertId() string {
	if o == nil || IsNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetAlertIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasAlertId() bool {
	if o != nil && !IsNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *EventViewForOrg) SetAlertId(v string) {
	o.AlertId = &v
}

// GetAlertConfigId returns the AlertConfigId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetAlertConfigId() string {
	if o == nil || IsNil(o.AlertConfigId) {
		var ret string
		return ret
	}
	return *o.AlertConfigId
}

// GetAlertConfigIdOk returns a tuple with the AlertConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetAlertConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertConfigId) {
		return nil, false
	}
	return o.AlertConfigId, true
}

// HasAlertConfigId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasAlertConfigId() bool {
	if o != nil && !IsNil(o.AlertConfigId) {
		return true
	}

	return false
}

// SetAlertConfigId gets a reference to the given string and assigns it to the AlertConfigId field.
func (o *EventViewForOrg) SetAlertConfigId(v string) {
	o.AlertConfigId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *EventViewForOrg) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *EventViewForOrg) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetWhitelistEntry returns the WhitelistEntry field value if set, zero value otherwise.
func (o *EventViewForOrg) GetWhitelistEntry() string {
	if o == nil || IsNil(o.WhitelistEntry) {
		var ret string
		return ret
	}
	return *o.WhitelistEntry
}

// GetWhitelistEntryOk returns a tuple with the WhitelistEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetWhitelistEntryOk() (*string, bool) {
	if o == nil || IsNil(o.WhitelistEntry) {
		return nil, false
	}
	return o.WhitelistEntry, true
}

// HasWhitelistEntry returns a boolean if a field has been set.
func (o *EventViewForOrg) HasWhitelistEntry() bool {
	if o != nil && !IsNil(o.WhitelistEntry) {
		return true
	}

	return false
}

// SetWhitelistEntry gets a reference to the given string and assigns it to the WhitelistEntry field.
func (o *EventViewForOrg) SetWhitelistEntry(v string) {
	o.WhitelistEntry = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *EventViewForOrg) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *EventViewForOrg) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *EventViewForOrg) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTargetUsername returns the TargetUsername field value if set, zero value otherwise.
func (o *EventViewForOrg) GetTargetUsername() string {
	if o == nil || IsNil(o.TargetUsername) {
		var ret string
		return ret
	}
	return *o.TargetUsername
}

// GetTargetUsernameOk returns a tuple with the TargetUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForOrg) GetTargetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUsername) {
		return nil, false
	}
	return o.TargetUsername, true
}

// HasTargetUsername returns a boolean if a field has been set.
func (o *EventViewForOrg) HasTargetUsername() bool {
	if o != nil && !IsNil(o.TargetUsername) {
		return true
	}

	return false
}

// SetTargetUsername gets a reference to the given string and assigns it to the TargetUsername field.
func (o *EventViewForOrg) SetTargetUsername(v string) {
	o.TargetUsername = &v
}

func (o EventViewForOrg) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o EventViewForOrg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventTypeName) {
		toSerialize["eventTypeName"] = o.EventTypeName
	}
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	return toSerialize, nil
}

type NullableEventViewForOrg struct {
	value *EventViewForOrg
	isSet bool
}

func (v NullableEventViewForOrg) Get() *EventViewForOrg {
	return v.value
}

func (v *NullableEventViewForOrg) Set(val *EventViewForOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableEventViewForOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableEventViewForOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventViewForOrg(val *EventViewForOrg) *NullableEventViewForOrg {
	return &NullableEventViewForOrg{value: val, isSet: true}
}

func (v NullableEventViewForOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventViewForOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

