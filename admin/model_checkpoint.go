// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// checks if the Checkpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Checkpoint{}

// Checkpoint struct for Checkpoint
type Checkpoint struct {
	// Unique 24-hexadecimal digit string that identifies the cluster that contains the checkpoint.
	ClusterId *string `json:"clusterId,omitempty"`
	// Date and time when the checkpoint completed and the balancer restarted. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Completed *time.Time `json:"completed,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that owns the checkpoints.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies checkpoint.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Metadata that describes the complete snapshot.  - For a replica set, this array contains a single document. - For a sharded cluster, this array contains one document for each shard plus one document for the config host.
	Parts []CheckpointPart `json:"parts,omitempty"`
	// Flag that indicates whether MongoDB Cloud can use the checkpoint for a restore.
	Restorable *bool `json:"restorable,omitempty"`
	// Date and time when the balancer stopped and began the checkpoint. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Started *time.Time `json:"started,omitempty"`
	// Date and time to which the checkpoint restores. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewCheckpoint instantiates a new Checkpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckpoint() *Checkpoint {
	this := Checkpoint{}
	return &this
}

// NewCheckpointWithDefaults instantiates a new Checkpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckpointWithDefaults() *Checkpoint {
	this := Checkpoint{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *Checkpoint) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *Checkpoint) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *Checkpoint) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *Checkpoint) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *Checkpoint) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *Checkpoint) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *Checkpoint) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *Checkpoint) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *Checkpoint) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Checkpoint) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Checkpoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Checkpoint) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Checkpoint) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Checkpoint) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Checkpoint) SetLinks(v []Link) {
	o.Links = v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *Checkpoint) GetParts() []CheckpointPart {
	if o == nil || IsNil(o.Parts) {
		var ret []CheckpointPart
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetPartsOk() ([]CheckpointPart, bool) {
	if o == nil || IsNil(o.Parts) {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *Checkpoint) HasParts() bool {
	if o != nil && !IsNil(o.Parts) {
		return true
	}

	return false
}

// SetParts gets a reference to the given []CheckpointPart and assigns it to the Parts field.
func (o *Checkpoint) SetParts(v []CheckpointPart) {
	o.Parts = v
}

// GetRestorable returns the Restorable field value if set, zero value otherwise.
func (o *Checkpoint) GetRestorable() bool {
	if o == nil || IsNil(o.Restorable) {
		var ret bool
		return ret
	}
	return *o.Restorable
}

// GetRestorableOk returns a tuple with the Restorable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetRestorableOk() (*bool, bool) {
	if o == nil || IsNil(o.Restorable) {
		return nil, false
	}
	return o.Restorable, true
}

// HasRestorable returns a boolean if a field has been set.
func (o *Checkpoint) HasRestorable() bool {
	if o != nil && !IsNil(o.Restorable) {
		return true
	}

	return false
}

// SetRestorable gets a reference to the given bool and assigns it to the Restorable field.
func (o *Checkpoint) SetRestorable(v bool) {
	o.Restorable = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *Checkpoint) GetStarted() time.Time {
	if o == nil || IsNil(o.Started) {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetStartedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *Checkpoint) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given time.Time and assigns it to the Started field.
func (o *Checkpoint) SetStarted(v time.Time) {
	o.Started = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Checkpoint) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Checkpoint) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Checkpoint) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o Checkpoint) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Checkpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableCheckpoint struct {
	value *Checkpoint
	isSet bool
}

func (v NullableCheckpoint) Get() *Checkpoint {
	return v.value
}

func (v *NullableCheckpoint) Set(val *Checkpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckpoint(val *Checkpoint) *NullableCheckpoint {
	return &NullableCheckpoint{value: val, isSet: true}
}

func (v NullableCheckpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
