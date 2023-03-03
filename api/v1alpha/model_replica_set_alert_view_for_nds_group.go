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

// checks if the ReplicaSetAlertViewForNdsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicaSetAlertViewForNdsGroup{}

// ReplicaSetAlertViewForNdsGroup ReplicaSet alert notifies about different activities on replica set of mongod instances.
type ReplicaSetAlertViewForNdsGroup struct {
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
	EventTypeName ReplicaSetEventTypeViewForNdsGroupAlertable `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert.
	GroupId *string `json:"groupId,omitempty"`
	// Hostname and port of the host to which this alert applies. The resource returns this parameter for alerts of events impacting hosts or replica sets.
	HostnameAndPort *string `json:"hostnameAndPort,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert.
	Id string `json:"id"`
	// Date and time that any notifications were last sent for this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert.
	LastNotified *time.Time `json:"lastNotified,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	NonRunningHostIds []string `json:"nonRunningHostIds,omitempty"`
	// Unique 24-hexadecimal character string that identifies the organization that owns the project to which this alert applies.
	OrgId *string `json:"orgId,omitempty"`
	// Unique 24-hexadecimal character string that identifies the parent cluster to which this alert applies. The parent cluster contains the sharded nodes. MongoDB Cloud returns this parameter only for alerts of events impacting sharded clusters.
	ParentClusterId *string `json:"parentClusterId,omitempty"`
	// Name of the replica set to which this alert applies. The response returns this parameter for alerts of events impacting backups, hosts, or replica sets.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Date and time that this alert changed to `\"status\" : \"CLOSED\"`. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter once `\"status\" : \"CLOSED\"`.
	Resolved *time.Time `json:"resolved,omitempty"`
	// State of this alert at the time you requested its details.
	Status string `json:"status"`
	// Date and time when someone last updated this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated time.Time `json:"updated"`
}

// NewReplicaSetAlertViewForNdsGroup instantiates a new ReplicaSetAlertViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaSetAlertViewForNdsGroup(acknowledgedUntil time.Time, alertConfigId string, created time.Time, eventTypeName ReplicaSetEventTypeViewForNdsGroupAlertable, id string, status string, updated time.Time) *ReplicaSetAlertViewForNdsGroup {
	this := ReplicaSetAlertViewForNdsGroup{}
	this.AcknowledgedUntil = acknowledgedUntil
	this.AlertConfigId = alertConfigId
	this.Created = created
	this.EventTypeName = eventTypeName
	this.Id = id
	this.Status = status
	this.Updated = updated
	return &this
}

// NewReplicaSetAlertViewForNdsGroupWithDefaults instantiates a new ReplicaSetAlertViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaSetAlertViewForNdsGroupWithDefaults() *ReplicaSetAlertViewForNdsGroup {
	this := ReplicaSetAlertViewForNdsGroup{}
	return &this
}

// GetAcknowledgedUntil returns the AcknowledgedUntil field value
func (o *ReplicaSetAlertViewForNdsGroup) GetAcknowledgedUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AcknowledgedUntil
}

// GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field value
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetAcknowledgedUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgedUntil, true
}

// SetAcknowledgedUntil sets field value
func (o *ReplicaSetAlertViewForNdsGroup) SetAcknowledgedUntil(v time.Time) {
	o.AcknowledgedUntil = v
}

// GetAcknowledgementComment returns the AcknowledgementComment field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetAcknowledgementComment() string {
	if o == nil || IsNil(o.AcknowledgementComment) {
		var ret string
		return ret
	}
	return *o.AcknowledgementComment
}

// GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetAcknowledgementCommentOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgementComment) {
		return nil, false
	}
	return o.AcknowledgementComment, true
}

// HasAcknowledgementComment returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasAcknowledgementComment() bool {
	if o != nil && !IsNil(o.AcknowledgementComment) {
		return true
	}

	return false
}

// SetAcknowledgementComment gets a reference to the given string and assigns it to the AcknowledgementComment field.
func (o *ReplicaSetAlertViewForNdsGroup) SetAcknowledgementComment(v string) {
	o.AcknowledgementComment = &v
}

// GetAcknowledgingUsername returns the AcknowledgingUsername field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetAcknowledgingUsername() string {
	if o == nil || IsNil(o.AcknowledgingUsername) {
		var ret string
		return ret
	}
	return *o.AcknowledgingUsername
}

// GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetAcknowledgingUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgingUsername) {
		return nil, false
	}
	return o.AcknowledgingUsername, true
}

// HasAcknowledgingUsername returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasAcknowledgingUsername() bool {
	if o != nil && !IsNil(o.AcknowledgingUsername) {
		return true
	}

	return false
}

// SetAcknowledgingUsername gets a reference to the given string and assigns it to the AcknowledgingUsername field.
func (o *ReplicaSetAlertViewForNdsGroup) SetAcknowledgingUsername(v string) {
	o.AcknowledgingUsername = &v
}

// GetAlertConfigId returns the AlertConfigId field value
func (o *ReplicaSetAlertViewForNdsGroup) GetAlertConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertConfigId
}

// GetAlertConfigIdOk returns a tuple with the AlertConfigId field value
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetAlertConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertConfigId, true
}

// SetAlertConfigId sets field value
func (o *ReplicaSetAlertViewForNdsGroup) SetAlertConfigId(v string) {
	o.AlertConfigId = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ReplicaSetAlertViewForNdsGroup) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreated returns the Created field value
func (o *ReplicaSetAlertViewForNdsGroup) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ReplicaSetAlertViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = v
}

// GetEventTypeName returns the EventTypeName field value
func (o *ReplicaSetAlertViewForNdsGroup) GetEventTypeName() ReplicaSetEventTypeViewForNdsGroupAlertable {
	if o == nil {
		var ret ReplicaSetEventTypeViewForNdsGroupAlertable
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetEventTypeNameOk() (*ReplicaSetEventTypeViewForNdsGroupAlertable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *ReplicaSetAlertViewForNdsGroup) SetEventTypeName(v ReplicaSetEventTypeViewForNdsGroupAlertable) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ReplicaSetAlertViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHostnameAndPort returns the HostnameAndPort field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetHostnameAndPort() string {
	if o == nil || IsNil(o.HostnameAndPort) {
		var ret string
		return ret
	}
	return *o.HostnameAndPort
}

// GetHostnameAndPortOk returns a tuple with the HostnameAndPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetHostnameAndPortOk() (*string, bool) {
	if o == nil || IsNil(o.HostnameAndPort) {
		return nil, false
	}
	return o.HostnameAndPort, true
}

// HasHostnameAndPort returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasHostnameAndPort() bool {
	if o != nil && !IsNil(o.HostnameAndPort) {
		return true
	}

	return false
}

// SetHostnameAndPort gets a reference to the given string and assigns it to the HostnameAndPort field.
func (o *ReplicaSetAlertViewForNdsGroup) SetHostnameAndPort(v string) {
	o.HostnameAndPort = &v
}

// GetId returns the Id field value
func (o *ReplicaSetAlertViewForNdsGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReplicaSetAlertViewForNdsGroup) SetId(v string) {
	o.Id = v
}

// GetLastNotified returns the LastNotified field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetLastNotified() time.Time {
	if o == nil || IsNil(o.LastNotified) {
		var ret time.Time
		return ret
	}
	return *o.LastNotified
}

// GetLastNotifiedOk returns a tuple with the LastNotified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetLastNotifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastNotified) {
		return nil, false
	}
	return o.LastNotified, true
}

// HasLastNotified returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasLastNotified() bool {
	if o != nil && !IsNil(o.LastNotified) {
		return true
	}

	return false
}

// SetLastNotified gets a reference to the given time.Time and assigns it to the LastNotified field.
func (o *ReplicaSetAlertViewForNdsGroup) SetLastNotified(v time.Time) {
	o.LastNotified = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ReplicaSetAlertViewForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetNonRunningHostIds returns the NonRunningHostIds field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetNonRunningHostIds() []string {
	if o == nil || IsNil(o.NonRunningHostIds) {
		var ret []string
		return ret
	}
	return o.NonRunningHostIds
}

// GetNonRunningHostIdsOk returns a tuple with the NonRunningHostIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetNonRunningHostIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.NonRunningHostIds) {
		return nil, false
	}
	return o.NonRunningHostIds, true
}

// HasNonRunningHostIds returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasNonRunningHostIds() bool {
	if o != nil && !IsNil(o.NonRunningHostIds) {
		return true
	}

	return false
}

// SetNonRunningHostIds gets a reference to the given []string and assigns it to the NonRunningHostIds field.
func (o *ReplicaSetAlertViewForNdsGroup) SetNonRunningHostIds(v []string) {
	o.NonRunningHostIds = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ReplicaSetAlertViewForNdsGroup) SetOrgId(v string) {
	o.OrgId = &v
}

// GetParentClusterId returns the ParentClusterId field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetParentClusterId() string {
	if o == nil || IsNil(o.ParentClusterId) {
		var ret string
		return ret
	}
	return *o.ParentClusterId
}

// GetParentClusterIdOk returns a tuple with the ParentClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetParentClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentClusterId) {
		return nil, false
	}
	return o.ParentClusterId, true
}

// HasParentClusterId returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasParentClusterId() bool {
	if o != nil && !IsNil(o.ParentClusterId) {
		return true
	}

	return false
}

// SetParentClusterId gets a reference to the given string and assigns it to the ParentClusterId field.
func (o *ReplicaSetAlertViewForNdsGroup) SetParentClusterId(v string) {
	o.ParentClusterId = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *ReplicaSetAlertViewForNdsGroup) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *ReplicaSetAlertViewForNdsGroup) GetResolved() time.Time {
	if o == nil || IsNil(o.Resolved) {
		var ret time.Time
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetResolvedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *ReplicaSetAlertViewForNdsGroup) HasResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given time.Time and assigns it to the Resolved field.
func (o *ReplicaSetAlertViewForNdsGroup) SetResolved(v time.Time) {
	o.Resolved = &v
}

// GetStatus returns the Status field value
func (o *ReplicaSetAlertViewForNdsGroup) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReplicaSetAlertViewForNdsGroup) SetStatus(v string) {
	o.Status = v
}

// GetUpdated returns the Updated field value
func (o *ReplicaSetAlertViewForNdsGroup) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *ReplicaSetAlertViewForNdsGroup) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *ReplicaSetAlertViewForNdsGroup) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o ReplicaSetAlertViewForNdsGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicaSetAlertViewForNdsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acknowledgedUntil"] = o.AcknowledgedUntil
	if !IsNil(o.AcknowledgementComment) {
		toSerialize["acknowledgementComment"] = o.AcknowledgementComment
	}
	// skip: acknowledgingUsername is readOnly
	// skip: alertConfigId is readOnly
	// skip: clusterName is readOnly
	// skip: created is readOnly
	toSerialize["eventTypeName"] = o.EventTypeName
	// skip: groupId is readOnly
	// skip: hostnameAndPort is readOnly
	// skip: id is readOnly
	// skip: lastNotified is readOnly
	// skip: links is readOnly
	// skip: nonRunningHostIds is readOnly
	// skip: orgId is readOnly
	// skip: parentClusterId is readOnly
	// skip: replicaSetName is readOnly
	// skip: resolved is readOnly
	// skip: status is readOnly
	// skip: updated is readOnly
	return toSerialize, nil
}

type NullableReplicaSetAlertViewForNdsGroup struct {
	value *ReplicaSetAlertViewForNdsGroup
	isSet bool
}

func (v NullableReplicaSetAlertViewForNdsGroup) Get() *ReplicaSetAlertViewForNdsGroup {
	return v.value
}

func (v *NullableReplicaSetAlertViewForNdsGroup) Set(val *ReplicaSetAlertViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaSetAlertViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaSetAlertViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaSetAlertViewForNdsGroup(val *ReplicaSetAlertViewForNdsGroup) *NullableReplicaSetAlertViewForNdsGroup {
	return &NullableReplicaSetAlertViewForNdsGroup{value: val, isSet: true}
}

func (v NullableReplicaSetAlertViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaSetAlertViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


