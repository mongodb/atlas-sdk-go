/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
)

// checks if the DiskBackupShardedClusterSnapshotMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskBackupShardedClusterSnapshotMember{}

// DiskBackupShardedClusterSnapshotMember List that includes the snapshots and the cloud provider that stores the snapshots. The resource returns this parameter when `\"type\" : \"SHARDED_CLUSTER\"`.
type DiskBackupShardedClusterSnapshotMember struct {
	// Human-readable label that identifies the cloud provider that stores this snapshot. The resource returns this parameter when `\"type\": \"replicaSet\"`.
	CloudProvider string `json:"cloudProvider"`
	// Unique 24-hexadecimal digit string that identifies the snapshot.
	Id string `json:"id"`
	// Human-readable label that identifies the shard or config host from which MongoDB Cloud took this snapshot.
	ReplicaSetName string `json:"replicaSetName"`
}

// NewDiskBackupShardedClusterSnapshotMember instantiates a new DiskBackupShardedClusterSnapshotMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupShardedClusterSnapshotMember(cloudProvider string, id string, replicaSetName string) *DiskBackupShardedClusterSnapshotMember {
	this := DiskBackupShardedClusterSnapshotMember{}
	this.CloudProvider = cloudProvider
	this.Id = id
	this.ReplicaSetName = replicaSetName
	return &this
}

// NewDiskBackupShardedClusterSnapshotMemberWithDefaults instantiates a new DiskBackupShardedClusterSnapshotMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupShardedClusterSnapshotMemberWithDefaults() *DiskBackupShardedClusterSnapshotMember {
	this := DiskBackupShardedClusterSnapshotMember{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *DiskBackupShardedClusterSnapshotMember) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *DiskBackupShardedClusterSnapshotMember) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *DiskBackupShardedClusterSnapshotMember) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetId returns the Id field value
func (o *DiskBackupShardedClusterSnapshotMember) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiskBackupShardedClusterSnapshotMember) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiskBackupShardedClusterSnapshotMember) SetId(v string) {
	o.Id = v
}

// GetReplicaSetName returns the ReplicaSetName field value
func (o *DiskBackupShardedClusterSnapshotMember) GetReplicaSetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value
// and a boolean to check if the value has been set.
func (o *DiskBackupShardedClusterSnapshotMember) GetReplicaSetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicaSetName, true
}

// SetReplicaSetName sets field value
func (o *DiskBackupShardedClusterSnapshotMember) SetReplicaSetName(v string) {
	o.ReplicaSetName = v
}

func (o DiskBackupShardedClusterSnapshotMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskBackupShardedClusterSnapshotMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: cloudProvider is readOnly
	// skip: id is readOnly
	// skip: replicaSetName is readOnly
	return toSerialize, nil
}

type NullableDiskBackupShardedClusterSnapshotMember struct {
	value *DiskBackupShardedClusterSnapshotMember
	isSet bool
}

func (v NullableDiskBackupShardedClusterSnapshotMember) Get() *DiskBackupShardedClusterSnapshotMember {
	return v.value
}

func (v *NullableDiskBackupShardedClusterSnapshotMember) Set(val *DiskBackupShardedClusterSnapshotMember) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskBackupShardedClusterSnapshotMember) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskBackupShardedClusterSnapshotMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskBackupShardedClusterSnapshotMember(val *DiskBackupShardedClusterSnapshotMember) *NullableDiskBackupShardedClusterSnapshotMember {
	return &NullableDiskBackupShardedClusterSnapshotMember{value: val, isSet: true}
}

func (v NullableDiskBackupShardedClusterSnapshotMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskBackupShardedClusterSnapshotMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


