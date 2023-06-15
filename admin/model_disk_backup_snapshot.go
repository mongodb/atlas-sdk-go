// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// DiskBackupSnapshot - struct for DiskBackupSnapshot
type DiskBackupSnapshot struct {
	DiskBackupReplicaSet             *DiskBackupReplicaSet
	DiskBackupShardedClusterSnapshot *DiskBackupShardedClusterSnapshot
}

// DiskBackupReplicaSetAsDiskBackupSnapshot is a convenience function that returns DiskBackupReplicaSet wrapped in DiskBackupSnapshot
func DiskBackupReplicaSetAsDiskBackupSnapshot(v *DiskBackupReplicaSet) DiskBackupSnapshot {
	return DiskBackupSnapshot{
		DiskBackupReplicaSet: v,
	}
}

// DiskBackupShardedClusterSnapshotAsDiskBackupSnapshot is a convenience function that returns DiskBackupShardedClusterSnapshot wrapped in DiskBackupSnapshot
func DiskBackupShardedClusterSnapshotAsDiskBackupSnapshot(v *DiskBackupShardedClusterSnapshot) DiskBackupSnapshot {
	return DiskBackupSnapshot{
		DiskBackupShardedClusterSnapshot: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DiskBackupSnapshot) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DiskBackupReplicaSet'
	if jsonDict["type"] == "DiskBackupReplicaSet" {
		// try to unmarshal JSON data into DiskBackupReplicaSet
		err = json.Unmarshal(data, &dst.DiskBackupReplicaSet)
		if err == nil {
			return nil // data stored in dst.DiskBackupReplicaSet, return on the first match
		} else {
			dst.DiskBackupReplicaSet = nil
			return fmt.Errorf("failed to unmarshal DiskBackupSnapshot as DiskBackupReplicaSet: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DiskBackupShardedClusterSnapshot'
	if jsonDict["type"] == "DiskBackupShardedClusterSnapshot" {
		// try to unmarshal JSON data into DiskBackupShardedClusterSnapshot
		err = json.Unmarshal(data, &dst.DiskBackupShardedClusterSnapshot)
		if err == nil {
			return nil // data stored in dst.DiskBackupShardedClusterSnapshot, return on the first match
		} else {
			dst.DiskBackupShardedClusterSnapshot = nil
			return fmt.Errorf("failed to unmarshal DiskBackupSnapshot as DiskBackupShardedClusterSnapshot: %s", err.Error())
		}
	}

	// check if the discriminator value is 'replicaSet'
	if jsonDict["type"] == "replicaSet" {
		// try to unmarshal JSON data into DiskBackupReplicaSet
		err = json.Unmarshal(data, &dst.DiskBackupReplicaSet)
		if err == nil {
			return nil // data stored in dst.DiskBackupReplicaSet, return on the first match
		} else {
			dst.DiskBackupReplicaSet = nil
			return fmt.Errorf("failed to unmarshal DiskBackupSnapshot as DiskBackupReplicaSet: %s", err.Error())
		}
	}

	// check if the discriminator value is 'shardedCluster'
	if jsonDict["type"] == "shardedCluster" {
		// try to unmarshal JSON data into DiskBackupShardedClusterSnapshot
		err = json.Unmarshal(data, &dst.DiskBackupShardedClusterSnapshot)
		if err == nil {
			return nil // data stored in dst.DiskBackupShardedClusterSnapshot, return on the first match
		} else {
			dst.DiskBackupShardedClusterSnapshot = nil
			return fmt.Errorf("failed to unmarshal DiskBackupSnapshot as DiskBackupShardedClusterSnapshot: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DiskBackupSnapshot) MarshalJSON() ([]byte, error) {
	if src.DiskBackupReplicaSet != nil {
		return json.Marshal(&src.DiskBackupReplicaSet)
	}

	if src.DiskBackupShardedClusterSnapshot != nil {
		return json.Marshal(&src.DiskBackupShardedClusterSnapshot)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DiskBackupSnapshot) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DiskBackupReplicaSet != nil {
		return obj.DiskBackupReplicaSet
	}

	if obj.DiskBackupShardedClusterSnapshot != nil {
		return obj.DiskBackupShardedClusterSnapshot
	}

	// all schemas are nil
	return nil
}

type NullableDiskBackupSnapshot struct {
	value *DiskBackupSnapshot
	isSet bool
}

func (v NullableDiskBackupSnapshot) Get() *DiskBackupSnapshot {
	return v.value
}

func (v *NullableDiskBackupSnapshot) Set(val *DiskBackupSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskBackupSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskBackupSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskBackupSnapshot(val *DiskBackupSnapshot) *NullableDiskBackupSnapshot {
	return &NullableDiskBackupSnapshot{value: val, isSet: true}
}

func (v NullableDiskBackupSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskBackupSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
