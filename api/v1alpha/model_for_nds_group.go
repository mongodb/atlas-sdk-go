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

// checks if the ForNdsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForNdsGroup{}

// ForNdsGroup ReplicaSet Event identifies different activities about replica set of mongod instances.
type ForNdsGroup struct {
	// Date and time when this event occurred. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created time.Time `json:"created"`
	EventTypeName ReplicaSetEventTypeViewForNdsGroup `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	Id string `json:"id"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	OrgId *string `json:"orgId,omitempty"`
	// IANA port on which the MongoDB process listens for requests.
	Port *int32 `json:"port,omitempty"`
	Raw *Raw `json:"raw,omitempty"`
	// Human-readable label of the replica set associated with the event.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Human-readable label of the shard associated with the event.
	ShardName *string `json:"shardName,omitempty"`
}

// NewForNdsGroup instantiates a new ForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForNdsGroup() *ForNdsGroup {
	this := ForNdsGroup{}
	return &this
}

// NewForNdsGroupWithDefaults instantiates a new ForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForNdsGroupWithDefaults() *ForNdsGroup {
	this := ForNdsGroup{}
	return &this
}

// GetCreated returns the Created field value
func (o *ForNdsGroup) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ForNdsGroup) SetCreated(v time.Time) {
	o.Created = v
}

// GetEventTypeName returns the EventTypeName field value
func (o *ForNdsGroup) GetEventTypeName() ReplicaSetEventTypeViewForNdsGroup {
	if o == nil {
		var ret ReplicaSetEventTypeViewForNdsGroup
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetEventTypeNameOk() (*ReplicaSetEventTypeViewForNdsGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *ForNdsGroup) SetEventTypeName(v ReplicaSetEventTypeViewForNdsGroup) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *ForNdsGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ForNdsGroup) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value
func (o *ForNdsGroup) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ForNdsGroup) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ForNdsGroup) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ForNdsGroup) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ForNdsGroup) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ForNdsGroup) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ForNdsGroup) SetPort(v int32) {
	o.Port = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *ForNdsGroup) GetRaw() Raw {
	if o == nil || IsNil(o.Raw) {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetRawOk() (*Raw, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *ForNdsGroup) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *ForNdsGroup) SetRaw(v Raw) {
	o.Raw = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *ForNdsGroup) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *ForNdsGroup) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *ForNdsGroup) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetShardName returns the ShardName field value if set, zero value otherwise.
func (o *ForNdsGroup) GetShardName() string {
	if o == nil || IsNil(o.ShardName) {
		var ret string
		return ret
	}
	return *o.ShardName
}

// GetShardNameOk returns a tuple with the ShardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForNdsGroup) GetShardNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShardName) {
		return nil, false
	}
	return o.ShardName, true
}

// HasShardName returns a boolean if a field has been set.
func (o *ForNdsGroup) HasShardName() bool {
	if o != nil && !IsNil(o.ShardName) {
		return true
	}

	return false
}

// SetShardName gets a reference to the given string and assigns it to the ShardName field.
func (o *ForNdsGroup) SetShardName(v string) {
	o.ShardName = &v
}

func (o ForNdsGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForNdsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created is readOnly
	toSerialize["eventTypeName"] = o.EventTypeName
	// skip: groupId is readOnly
	// skip: id is readOnly
	// skip: links is readOnly
	// skip: orgId is readOnly
	// skip: port is readOnly
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	// skip: replicaSetName is readOnly
	// skip: shardName is readOnly
	return toSerialize, nil
}

type NullableForNdsGroup struct {
	value *ForNdsGroup
	isSet bool
}

func (v NullableForNdsGroup) Get() *ForNdsGroup {
	return v.value
}

func (v *NullableForNdsGroup) Set(val *ForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForNdsGroup(val *ForNdsGroup) *NullableForNdsGroup {
	return &NullableForNdsGroup{value: val, isSet: true}
}

func (v NullableForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


