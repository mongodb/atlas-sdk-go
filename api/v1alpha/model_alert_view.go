/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// checks if the AlertView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertView{}

// AlertView Alert represents a notice of warning, a threat or a problem in the system. It can reflect a certain event or condition in the system. An Alert can be acknowledged by the user, but stays open until alert condition is resolved in the system.
type AlertView struct {
	// Date and time until which this alert has been acknowledged. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.  - To acknowledge this alert forever, set the parameter value to 100 years in the future.  - To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past.
	AcknowledgedUntil time.Time `json:"acknowledgedUntil"`
	// Comment that a MongoDB Cloud user submitted when acknowledging the alert.
	AcknowledgementComment *string `json:"acknowledgementComment,omitempty"`
	// MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert.
	AcknowledgingUsername *string `json:"acknowledgingUsername,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert.
	AlertConfigId string `json:"alertConfigId"`
	// Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created time.Time `json:"created"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert.
	Id string `json:"id"`
	// Date and time that any notifications were last sent for this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert.
	LastNotified *time.Time `json:"lastNotified,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Unique 24-hexadecimal character string that identifies the organization that owns the project to which this alert applies.
	OrgId *string `json:"orgId,omitempty"`
	// Date and time that this alert changed to `\"status\" : \"CLOSED\"`. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter once `\"status\" : \"CLOSED\"`.
	Resolved *time.Time `json:"resolved,omitempty"`
	// State of this alert at the time you requested its details.
	Status string `json:"status"`
	// Date and time when someone last updated this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated time.Time `json:"updated"`
}

// NewAlertView instantiates a new AlertView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertView() *AlertView {
	this := AlertView{}
	return &this
}

// NewAlertViewWithDefaults instantiates a new AlertView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertViewWithDefaults() *AlertView {
	this := AlertView{}
	return &this
}

// GetAcknowledgedUntil returns the AcknowledgedUntil field value
func (o *AlertView) GetAcknowledgedUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AcknowledgedUntil
}

// GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field value
// and a boolean to check if the value has been set.
func (o *AlertView) GetAcknowledgedUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgedUntil, true
}

// SetAcknowledgedUntil sets field value
func (o *AlertView) SetAcknowledgedUntil(v time.Time) {
	o.AcknowledgedUntil = v
}

// GetAcknowledgementComment returns the AcknowledgementComment field value if set, zero value otherwise.
func (o *AlertView) GetAcknowledgementComment() string {
	if o == nil || IsNil(o.AcknowledgementComment) {
		var ret string
		return ret
	}
	return *o.AcknowledgementComment
}

// GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertView) GetAcknowledgementCommentOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgementComment) {
		return nil, false
	}
	return o.AcknowledgementComment, true
}

// HasAcknowledgementComment returns a boolean if a field has been set.
func (o *AlertView) HasAcknowledgementComment() bool {
	if o != nil && !IsNil(o.AcknowledgementComment) {
		return true
	}

	return false
}

// SetAcknowledgementComment gets a reference to the given string and assigns it to the AcknowledgementComment field.
func (o *AlertView) SetAcknowledgementComment(v string) {
	o.AcknowledgementComment = &v
}

// GetAcknowledgingUsername returns the AcknowledgingUsername field value if set, zero value otherwise.
func (o *AlertView) GetAcknowledgingUsername() string {
	if o == nil || IsNil(o.AcknowledgingUsername) {
		var ret string
		return ret
	}
	return *o.AcknowledgingUsername
}

// GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertView) GetAcknowledgingUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgingUsername) {
		return nil, false
	}
	return o.AcknowledgingUsername, true
}

// HasAcknowledgingUsername returns a boolean if a field has been set.
func (o *AlertView) HasAcknowledgingUsername() bool {
	if o != nil && !IsNil(o.AcknowledgingUsername) {
		return true
	}

	return false
}

// SetAcknowledgingUsername gets a reference to the given string and assigns it to the AcknowledgingUsername field.
func (o *AlertView) SetAcknowledgingUsername(v string) {
	o.AcknowledgingUsername = &v
}

// GetAlertConfigId returns the AlertConfigId field value
func (o *AlertView) GetAlertConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertConfigId
}

// GetAlertConfigIdOk returns a tuple with the AlertConfigId field value
// and a boolean to check if the value has been set.
func (o *AlertView) GetAlertConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertConfigId, true
}

// SetAlertConfigId sets field value
func (o *AlertView) SetAlertConfigId(v string) {
	o.AlertConfigId = v
}

// GetCreated returns the Created field value
func (o *AlertView) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AlertView) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AlertView) SetCreated(v time.Time) {
	o.Created = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AlertView) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertView) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AlertView) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AlertView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *AlertView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AlertView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AlertView) SetId(v string) {
	o.Id = v
}

// GetLastNotified returns the LastNotified field value if set, zero value otherwise.
func (o *AlertView) GetLastNotified() time.Time {
	if o == nil || IsNil(o.LastNotified) {
		var ret time.Time
		return ret
	}
	return *o.LastNotified
}

// GetLastNotifiedOk returns a tuple with the LastNotified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertView) GetLastNotifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastNotified) {
		return nil, false
	}
	return o.LastNotified, true
}

// HasLastNotified returns a boolean if a field has been set.
func (o *AlertView) HasLastNotified() bool {
	if o != nil && !IsNil(o.LastNotified) {
		return true
	}

	return false
}

// SetLastNotified gets a reference to the given time.Time and assigns it to the LastNotified field.
func (o *AlertView) SetLastNotified(v time.Time) {
	o.LastNotified = &v
}

// GetLinks returns the Links field value
func (o *AlertView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AlertView) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *AlertView) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *AlertView) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertView) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *AlertView) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *AlertView) SetOrgId(v string) {
	o.OrgId = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *AlertView) GetResolved() time.Time {
	if o == nil || IsNil(o.Resolved) {
		var ret time.Time
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertView) GetResolvedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *AlertView) HasResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given time.Time and assigns it to the Resolved field.
func (o *AlertView) SetResolved(v time.Time) {
	o.Resolved = &v
}

// GetStatus returns the Status field value
func (o *AlertView) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AlertView) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AlertView) SetStatus(v string) {
	o.Status = v
}

// GetUpdated returns the Updated field value
func (o *AlertView) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *AlertView) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *AlertView) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o AlertView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acknowledgedUntil"] = o.AcknowledgedUntil
	if !IsNil(o.AcknowledgementComment) {
		toSerialize["acknowledgementComment"] = o.AcknowledgementComment
	}
	// skip: acknowledgingUsername is readOnly
	// skip: alertConfigId is readOnly
	// skip: created is readOnly
	// skip: groupId is readOnly
	// skip: id is readOnly
	// skip: lastNotified is readOnly
	// skip: links is readOnly
	// skip: orgId is readOnly
	// skip: resolved is readOnly
	// skip: status is readOnly
	// skip: updated is readOnly
	return toSerialize, nil
}

type NullableAlertView struct {
	value *AlertView
	isSet bool
}

func (v NullableAlertView) Get() *AlertView {
	return v.value
}

func (v *NullableAlertView) Set(val *AlertView) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertView) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertView(val *AlertView) *NullableAlertView {
	return &NullableAlertView{value: val, isSet: true}
}

func (v NullableAlertView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


