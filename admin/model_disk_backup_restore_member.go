// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// DiskBackupRestoreMember struct for DiskBackupRestoreMember
type DiskBackupRestoreMember struct {
	// One Uniform Resource Locator that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when `\"deliveryType\" : \"download\"`.
	// Read only field.
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// Human-readable label that identifies the replica set on the sharded cluster.
	// Read only field.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
}

// NewDiskBackupRestoreMember instantiates a new DiskBackupRestoreMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupRestoreMember() *DiskBackupRestoreMember {
	this := DiskBackupRestoreMember{}
	return &this
}

// NewDiskBackupRestoreMemberWithDefaults instantiates a new DiskBackupRestoreMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupRestoreMemberWithDefaults() *DiskBackupRestoreMember {
	this := DiskBackupRestoreMember{}
	return &this
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise
func (o *DiskBackupRestoreMember) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreMember) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}

	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *DiskBackupRestoreMember) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *DiskBackupRestoreMember) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise
func (o *DiskBackupRestoreMember) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreMember) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}

	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *DiskBackupRestoreMember) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *DiskBackupRestoreMember) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

func (o DiskBackupRestoreMember) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DiskBackupRestoreMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
