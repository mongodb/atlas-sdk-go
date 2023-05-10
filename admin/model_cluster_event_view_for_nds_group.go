// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"time"
)

// checks if the ClusterEventViewForNdsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterEventViewForNdsGroup{}

// ClusterEventViewForNdsGroup Cluster event identifies different activities about cluster of mongod hosts.
type ClusterEventViewForNdsGroup struct {
	// Date and time when this event occurred. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created       time.Time                       `json:"created"`
	EventTypeName ClusterEventTypeViewForNdsGroup `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	Id string `json:"id"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	OrgId *string `json:"orgId,omitempty"`
	Raw   *Raw    `json:"raw,omitempty"`
	// Human-readable label of the shard associated with the event.
	ShardName *string `json:"shardName,omitempty"`
}

// NewClusterEventViewForNdsGroup instantiates a new ClusterEventViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterEventViewForNdsGroup(created time.Time, eventTypeName ClusterEventTypeViewForNdsGroup, id string) *ClusterEventViewForNdsGroup {
	this := ClusterEventViewForNdsGroup{}
	this.Created = created
	this.EventTypeName = eventTypeName
	this.Id = id
	return &this
}

// NewClusterEventViewForNdsGroupWithDefaults instantiates a new ClusterEventViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterEventViewForNdsGroupWithDefaults() *ClusterEventViewForNdsGroup {
	this := ClusterEventViewForNdsGroup{}
	return &this
}

// GetCreated returns the Created field value
func (o *ClusterEventViewForNdsGroup) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ClusterEventViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ClusterEventViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = v
}

// GetEventTypeName returns the EventTypeName field value
func (o *ClusterEventViewForNdsGroup) GetEventTypeName() ClusterEventTypeViewForNdsGroup {
	if o == nil {
		var ret ClusterEventTypeViewForNdsGroup
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *ClusterEventViewForNdsGroup) GetEventTypeNameOk() (*ClusterEventTypeViewForNdsGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *ClusterEventViewForNdsGroup) SetEventTypeName(v ClusterEventTypeViewForNdsGroup) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ClusterEventViewForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEventViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ClusterEventViewForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ClusterEventViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *ClusterEventViewForNdsGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterEventViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClusterEventViewForNdsGroup) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ClusterEventViewForNdsGroup) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEventViewForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ClusterEventViewForNdsGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ClusterEventViewForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ClusterEventViewForNdsGroup) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEventViewForNdsGroup) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ClusterEventViewForNdsGroup) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ClusterEventViewForNdsGroup) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *ClusterEventViewForNdsGroup) GetRaw() Raw {
	if o == nil || IsNil(o.Raw) {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEventViewForNdsGroup) GetRawOk() (*Raw, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *ClusterEventViewForNdsGroup) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *ClusterEventViewForNdsGroup) SetRaw(v Raw) {
	o.Raw = &v
}

// GetShardName returns the ShardName field value if set, zero value otherwise.
func (o *ClusterEventViewForNdsGroup) GetShardName() string {
	if o == nil || IsNil(o.ShardName) {
		var ret string
		return ret
	}
	return *o.ShardName
}

// GetShardNameOk returns a tuple with the ShardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEventViewForNdsGroup) GetShardNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShardName) {
		return nil, false
	}
	return o.ShardName, true
}

// HasShardName returns a boolean if a field has been set.
func (o *ClusterEventViewForNdsGroup) HasShardName() bool {
	if o != nil && !IsNil(o.ShardName) {
		return true
	}

	return false
}

// SetShardName gets a reference to the given string and assigns it to the ShardName field.
func (o *ClusterEventViewForNdsGroup) SetShardName(v string) {
	o.ShardName = &v
}

func (o ClusterEventViewForNdsGroup) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterEventViewForNdsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventTypeName"] = o.EventTypeName
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	return toSerialize, nil
}

type NullableClusterEventViewForNdsGroup struct {
	value *ClusterEventViewForNdsGroup
	isSet bool
}

func (v NullableClusterEventViewForNdsGroup) Get() *ClusterEventViewForNdsGroup {
	return v.value
}

func (v *NullableClusterEventViewForNdsGroup) Set(val *ClusterEventViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterEventViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterEventViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterEventViewForNdsGroup(val *ClusterEventViewForNdsGroup) *NullableClusterEventViewForNdsGroup {
	return &NullableClusterEventViewForNdsGroup{value: val, isSet: true}
}

func (v NullableClusterEventViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterEventViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
