// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the SnapshotPart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotPart{}

// SnapshotPart struct for SnapshotPart
type SnapshotPart struct {
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return.
	ClusterId *string `json:"clusterId,omitempty"`
	// Human-readable label that identifies the method of compression for the snapshot.
	CompressionSetting *string `json:"compressionSetting,omitempty"`
	// Total size of the data stored on each node in the cluster. This parameter expresses its value in bytes.
	DataSizeBytes *int64 `json:"dataSizeBytes,omitempty"`
	// Flag that indicates whether someone encrypted this snapshot.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`
	// Number that indicates the total size of the data files in bytes.
	FileSizeBytes *int64 `json:"fileSizeBytes,omitempty"`
	// Unique string that identifies the Key Management Interoperability (KMIP) master key used to encrypt the snapshot data. The resource returns this parameter when `\"parts.encryptionEnabled\" : true`.
	MasterKeyUUID *string `json:"masterKeyUUID,omitempty"`
	// Number that indicates the version of MongoDB that the replica set primary ran when MongoDB Cloud created the snapshot.
	MongodVersion *string `json:"mongodVersion,omitempty"`
	// Human-readable label that identifies the replica set.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Number that indicates the total size of space allocated for document storage.
	StorageSizeBytes *int64 `json:"storageSizeBytes,omitempty"`
	// Human-readable label that identifies the type of server from which MongoDB Cloud took this snapshot.
	TypeName *string `json:"typeName,omitempty"`
}

// NewSnapshotPart instantiates a new SnapshotPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotPart() *SnapshotPart {
	this := SnapshotPart{}
	return &this
}

// NewSnapshotPartWithDefaults instantiates a new SnapshotPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotPartWithDefaults() *SnapshotPart {
	this := SnapshotPart{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *SnapshotPart) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *SnapshotPart) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *SnapshotPart) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetCompressionSetting returns the CompressionSetting field value if set, zero value otherwise.
func (o *SnapshotPart) GetCompressionSetting() string {
	if o == nil || IsNil(o.CompressionSetting) {
		var ret string
		return ret
	}
	return *o.CompressionSetting
}

// GetCompressionSettingOk returns a tuple with the CompressionSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetCompressionSettingOk() (*string, bool) {
	if o == nil || IsNil(o.CompressionSetting) {
		return nil, false
	}
	return o.CompressionSetting, true
}

// HasCompressionSetting returns a boolean if a field has been set.
func (o *SnapshotPart) HasCompressionSetting() bool {
	if o != nil && !IsNil(o.CompressionSetting) {
		return true
	}

	return false
}

// SetCompressionSetting gets a reference to the given string and assigns it to the CompressionSetting field.
func (o *SnapshotPart) SetCompressionSetting(v string) {
	o.CompressionSetting = &v
}

// GetDataSizeBytes returns the DataSizeBytes field value if set, zero value otherwise.
func (o *SnapshotPart) GetDataSizeBytes() int64 {
	if o == nil || IsNil(o.DataSizeBytes) {
		var ret int64
		return ret
	}
	return *o.DataSizeBytes
}

// GetDataSizeBytesOk returns a tuple with the DataSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetDataSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DataSizeBytes) {
		return nil, false
	}
	return o.DataSizeBytes, true
}

// HasDataSizeBytes returns a boolean if a field has been set.
func (o *SnapshotPart) HasDataSizeBytes() bool {
	if o != nil && !IsNil(o.DataSizeBytes) {
		return true
	}

	return false
}

// SetDataSizeBytes gets a reference to the given int64 and assigns it to the DataSizeBytes field.
func (o *SnapshotPart) SetDataSizeBytes(v int64) {
	o.DataSizeBytes = &v
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise.
func (o *SnapshotPart) GetEncryptionEnabled() bool {
	if o == nil || IsNil(o.EncryptionEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptionEnabled) {
		return nil, false
	}
	return o.EncryptionEnabled, true
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *SnapshotPart) HasEncryptionEnabled() bool {
	if o != nil && !IsNil(o.EncryptionEnabled) {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given bool and assigns it to the EncryptionEnabled field.
func (o *SnapshotPart) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled = &v
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise.
func (o *SnapshotPart) GetFileSizeBytes() int64 {
	if o == nil || IsNil(o.FileSizeBytes) {
		var ret int64
		return ret
	}
	return *o.FileSizeBytes
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetFileSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSizeBytes) {
		return nil, false
	}
	return o.FileSizeBytes, true
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *SnapshotPart) HasFileSizeBytes() bool {
	if o != nil && !IsNil(o.FileSizeBytes) {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given int64 and assigns it to the FileSizeBytes field.
func (o *SnapshotPart) SetFileSizeBytes(v int64) {
	o.FileSizeBytes = &v
}

// GetMasterKeyUUID returns the MasterKeyUUID field value if set, zero value otherwise.
func (o *SnapshotPart) GetMasterKeyUUID() string {
	if o == nil || IsNil(o.MasterKeyUUID) {
		var ret string
		return ret
	}
	return *o.MasterKeyUUID
}

// GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetMasterKeyUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.MasterKeyUUID) {
		return nil, false
	}
	return o.MasterKeyUUID, true
}

// HasMasterKeyUUID returns a boolean if a field has been set.
func (o *SnapshotPart) HasMasterKeyUUID() bool {
	if o != nil && !IsNil(o.MasterKeyUUID) {
		return true
	}

	return false
}

// SetMasterKeyUUID gets a reference to the given string and assigns it to the MasterKeyUUID field.
func (o *SnapshotPart) SetMasterKeyUUID(v string) {
	o.MasterKeyUUID = &v
}

// GetMongodVersion returns the MongodVersion field value if set, zero value otherwise.
func (o *SnapshotPart) GetMongodVersion() string {
	if o == nil || IsNil(o.MongodVersion) {
		var ret string
		return ret
	}
	return *o.MongodVersion
}

// GetMongodVersionOk returns a tuple with the MongodVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetMongodVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongodVersion) {
		return nil, false
	}
	return o.MongodVersion, true
}

// HasMongodVersion returns a boolean if a field has been set.
func (o *SnapshotPart) HasMongodVersion() bool {
	if o != nil && !IsNil(o.MongodVersion) {
		return true
	}

	return false
}

// SetMongodVersion gets a reference to the given string and assigns it to the MongodVersion field.
func (o *SnapshotPart) SetMongodVersion(v string) {
	o.MongodVersion = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *SnapshotPart) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *SnapshotPart) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *SnapshotPart) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetStorageSizeBytes returns the StorageSizeBytes field value if set, zero value otherwise.
func (o *SnapshotPart) GetStorageSizeBytes() int64 {
	if o == nil || IsNil(o.StorageSizeBytes) {
		var ret int64
		return ret
	}
	return *o.StorageSizeBytes
}

// GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetStorageSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.StorageSizeBytes) {
		return nil, false
	}
	return o.StorageSizeBytes, true
}

// HasStorageSizeBytes returns a boolean if a field has been set.
func (o *SnapshotPart) HasStorageSizeBytes() bool {
	if o != nil && !IsNil(o.StorageSizeBytes) {
		return true
	}

	return false
}

// SetStorageSizeBytes gets a reference to the given int64 and assigns it to the StorageSizeBytes field.
func (o *SnapshotPart) SetStorageSizeBytes(v int64) {
	o.StorageSizeBytes = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *SnapshotPart) GetTypeName() string {
	if o == nil || IsNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPart) GetTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TypeName) {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *SnapshotPart) HasTypeName() bool {
	if o != nil && !IsNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *SnapshotPart) SetTypeName(v string) {
	o.TypeName = &v
}

func (o SnapshotPart) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o SnapshotPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableSnapshotPart struct {
	value *SnapshotPart
	isSet bool
}

func (v NullableSnapshotPart) Get() *SnapshotPart {
	return v.value
}

func (v *NullableSnapshotPart) Set(val *SnapshotPart) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotPart) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotPart(val *SnapshotPart) *NullableSnapshotPart {
	return &NullableSnapshotPart{value: val, isSet: true}
}

func (v NullableSnapshotPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
