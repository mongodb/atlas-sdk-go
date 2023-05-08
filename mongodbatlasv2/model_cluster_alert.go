/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the ClusterAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterAlert{}

// ClusterAlert Cluster alert notifies different activities and conditions about cluster of mongod hosts.
type ClusterAlert struct {
	// Date and time until which this alert has been acknowledged. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.  - To acknowledge this alert forever, set the parameter value to 100 years in the future.  - To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past.
	AcknowledgedUntil time.Time `json:"acknowledgedUntil"`
	// Comment that a MongoDB Cloud user submitted when acknowledging the alert.
	AcknowledgementComment *string `json:"acknowledgementComment,omitempty"`
	// MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert.
	AcknowledgingUsername *string `json:"acknowledgingUsername,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert.
	AlertConfigId string `json:"alertConfigId"`
	// Human-readable label that identifies the cluster to which this alert applies. This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters.
	ClusterName *string `json:"clusterName,omitempty"`
	// Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created time.Time `json:"created"`
	EventTypeName ClusterEventTypeViewAlertable `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert.
	Id string `json:"id"`
	// Date and time that any notifications were last sent for this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert.
	LastNotified *time.Time `json:"lastNotified,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Unique 24-hexadecimal character string that identifies the organization that owns the project to which this alert applies.
	OrgId *string `json:"orgId,omitempty"`
	// Date and time that this alert changed to `\"status\" : \"CLOSED\"`. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter once `\"status\" : \"CLOSED\"`.
	Resolved *time.Time `json:"resolved,omitempty"`
	// State of this alert at the time you requested its details.
	Status string `json:"status"`
	// Date and time when someone last updated this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated time.Time `json:"updated"`
}

// NewClusterAlert instantiates a new ClusterAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAlert(acknowledgedUntil time.Time, alertConfigId string, created time.Time, eventTypeName ClusterEventTypeViewAlertable, id string, status string, updated time.Time) *ClusterAlert {
	this := ClusterAlert{}
	this.AcknowledgedUntil = acknowledgedUntil
	this.AlertConfigId = alertConfigId
	this.Created = created
	this.EventTypeName = eventTypeName
	this.Id = id
	this.Status = status
	this.Updated = updated
	return &this
}

// NewClusterAlertWithDefaults instantiates a new ClusterAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAlertWithDefaults() *ClusterAlert {
	this := ClusterAlert{}
	return &this
}

// GetAcknowledgedUntil returns the AcknowledgedUntil field value
func (o *ClusterAlert) GetAcknowledgedUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AcknowledgedUntil
}

// GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field value
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetAcknowledgedUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgedUntil, true
}

// SetAcknowledgedUntil sets field value
func (o *ClusterAlert) SetAcknowledgedUntil(v time.Time) {
	o.AcknowledgedUntil = v
}

// GetAcknowledgementComment returns the AcknowledgementComment field value if set, zero value otherwise.
func (o *ClusterAlert) GetAcknowledgementComment() string {
	if o == nil || IsNil(o.AcknowledgementComment) {
		var ret string
		return ret
	}
	return *o.AcknowledgementComment
}

// GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetAcknowledgementCommentOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgementComment) {
		return nil, false
	}
	return o.AcknowledgementComment, true
}

// HasAcknowledgementComment returns a boolean if a field has been set.
func (o *ClusterAlert) HasAcknowledgementComment() bool {
	if o != nil && !IsNil(o.AcknowledgementComment) {
		return true
	}

	return false
}

// SetAcknowledgementComment gets a reference to the given string and assigns it to the AcknowledgementComment field.
func (o *ClusterAlert) SetAcknowledgementComment(v string) {
	o.AcknowledgementComment = &v
}

// GetAcknowledgingUsername returns the AcknowledgingUsername field value if set, zero value otherwise.
func (o *ClusterAlert) GetAcknowledgingUsername() string {
	if o == nil || IsNil(o.AcknowledgingUsername) {
		var ret string
		return ret
	}
	return *o.AcknowledgingUsername
}

// GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetAcknowledgingUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgingUsername) {
		return nil, false
	}
	return o.AcknowledgingUsername, true
}

// HasAcknowledgingUsername returns a boolean if a field has been set.
func (o *ClusterAlert) HasAcknowledgingUsername() bool {
	if o != nil && !IsNil(o.AcknowledgingUsername) {
		return true
	}

	return false
}

// SetAcknowledgingUsername gets a reference to the given string and assigns it to the AcknowledgingUsername field.
func (o *ClusterAlert) SetAcknowledgingUsername(v string) {
	o.AcknowledgingUsername = &v
}

// GetAlertConfigId returns the AlertConfigId field value
func (o *ClusterAlert) GetAlertConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertConfigId
}

// GetAlertConfigIdOk returns a tuple with the AlertConfigId field value
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetAlertConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertConfigId, true
}

// SetAlertConfigId sets field value
func (o *ClusterAlert) SetAlertConfigId(v string) {
	o.AlertConfigId = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ClusterAlert) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ClusterAlert) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ClusterAlert) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreated returns the Created field value
func (o *ClusterAlert) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ClusterAlert) SetCreated(v time.Time) {
	o.Created = v
}

// GetEventTypeName returns the EventTypeName field value
func (o *ClusterAlert) GetEventTypeName() ClusterEventTypeViewAlertable {
	if o == nil {
		var ret ClusterEventTypeViewAlertable
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetEventTypeNameOk() (*ClusterEventTypeViewAlertable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *ClusterAlert) SetEventTypeName(v ClusterEventTypeViewAlertable) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ClusterAlert) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ClusterAlert) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ClusterAlert) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *ClusterAlert) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClusterAlert) SetId(v string) {
	o.Id = v
}

// GetLastNotified returns the LastNotified field value if set, zero value otherwise.
func (o *ClusterAlert) GetLastNotified() time.Time {
	if o == nil || IsNil(o.LastNotified) {
		var ret time.Time
		return ret
	}
	return *o.LastNotified
}

// GetLastNotifiedOk returns a tuple with the LastNotified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetLastNotifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastNotified) {
		return nil, false
	}
	return o.LastNotified, true
}

// HasLastNotified returns a boolean if a field has been set.
func (o *ClusterAlert) HasLastNotified() bool {
	if o != nil && !IsNil(o.LastNotified) {
		return true
	}

	return false
}

// SetLastNotified gets a reference to the given time.Time and assigns it to the LastNotified field.
func (o *ClusterAlert) SetLastNotified(v time.Time) {
	o.LastNotified = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ClusterAlert) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ClusterAlert) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ClusterAlert) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ClusterAlert) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ClusterAlert) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ClusterAlert) SetOrgId(v string) {
	o.OrgId = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *ClusterAlert) GetResolved() time.Time {
	if o == nil || IsNil(o.Resolved) {
		var ret time.Time
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetResolvedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *ClusterAlert) HasResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given time.Time and assigns it to the Resolved field.
func (o *ClusterAlert) SetResolved(v time.Time) {
	o.Resolved = &v
}

// GetStatus returns the Status field value
func (o *ClusterAlert) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ClusterAlert) SetStatus(v string) {
	o.Status = v
}

// GetUpdated returns the Updated field value
func (o *ClusterAlert) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *ClusterAlert) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *ClusterAlert) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o ClusterAlert) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acknowledgedUntil"] = o.AcknowledgedUntil
	if !IsNil(o.AcknowledgementComment) {
		toSerialize["acknowledgementComment"] = o.AcknowledgementComment
	}
	toSerialize["eventTypeName"] = o.EventTypeName
	return toSerialize, nil
}

type NullableClusterAlert struct {
	value *ClusterAlert
	isSet bool
}

func (v NullableClusterAlert) Get() *ClusterAlert {
	return v.value
}

func (v *NullableClusterAlert) Set(val *ClusterAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAlert(val *ClusterAlert) *NullableClusterAlert {
	return &NullableClusterAlert{value: val, isSet: true}
}

func (v NullableClusterAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


