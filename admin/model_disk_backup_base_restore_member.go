// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DiskBackupBaseRestoreMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskBackupBaseRestoreMember{}

// DiskBackupBaseRestoreMember struct for DiskBackupBaseRestoreMember
type DiskBackupBaseRestoreMember struct {
	// Human-readable label that identifies the replica set on the sharded cluster.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
}

// NewDiskBackupBaseRestoreMember instantiates a new DiskBackupBaseRestoreMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupBaseRestoreMember() *DiskBackupBaseRestoreMember {
	this := DiskBackupBaseRestoreMember{}
	return &this
}

// NewDiskBackupBaseRestoreMemberWithDefaults instantiates a new DiskBackupBaseRestoreMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupBaseRestoreMemberWithDefaults() *DiskBackupBaseRestoreMember {
	this := DiskBackupBaseRestoreMember{}
	return &this
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *DiskBackupBaseRestoreMember) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupBaseRestoreMember) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *DiskBackupBaseRestoreMember) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *DiskBackupBaseRestoreMember) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

func (o DiskBackupBaseRestoreMember) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DiskBackupBaseRestoreMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableDiskBackupBaseRestoreMember struct {
	value *DiskBackupBaseRestoreMember
	isSet bool
}

func (v NullableDiskBackupBaseRestoreMember) Get() *DiskBackupBaseRestoreMember {
	return v.value
}

func (v *NullableDiskBackupBaseRestoreMember) Set(val *DiskBackupBaseRestoreMember) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskBackupBaseRestoreMember) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskBackupBaseRestoreMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskBackupBaseRestoreMember(val *DiskBackupBaseRestoreMember) *NullableDiskBackupBaseRestoreMember {
	return &NullableDiskBackupBaseRestoreMember{value: val, isSet: true}
}

func (v NullableDiskBackupBaseRestoreMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskBackupBaseRestoreMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
