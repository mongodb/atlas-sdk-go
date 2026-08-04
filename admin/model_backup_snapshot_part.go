// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// BackupSnapshotPart Characteristics that identify this snapshot.
type BackupSnapshotPart struct {
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return.
	// Read only field.
	ClusterId *string `json:"clusterId,omitempty"`
	// Date and time when the snapshot completed. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	CompletedTime *time.Time `json:"completedTime,omitempty"`
	// Human-readable label that identifies the method of compression for the snapshot.
	// Read only field.
	CompressionSetting *string `json:"compressionSetting,omitempty"`
	// Total size of the data stored on each node in the cluster. This parameter expresses its value in bytes.
	// Read only field.
	DataSizeBytes *int64 `json:"dataSizeBytes,omitempty"`
	// Flag that indicates whether someone encrypted this snapshot.
	// Read only field.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`
	// Number that indicates the feature compatibility version of MongoDB that the replica set primary ran when MongoDB Cloud created the snapshot.
	// Read only field.
	Fcv *string `json:"fcv,omitempty"`
	// Number that indicates the total size of the data files in bytes.
	// Read only field.
	FileSizeBytes *int64 `json:"fileSizeBytes,omitempty"`
	// Hostname and port that indicate the node on which MongoDB Cloud created the snapshot.
	// Read only field.
	MachineId *string `json:"machineId,omitempty"`
	// Unique string that identifies the Key Management Interoperability (KMIP) master key used to encrypt the snapshot data. The resource returns this parameter when `\"parts.encryptionEnabled\" : true`.
	// Read only field.
	MasterKeyUUID *string `json:"masterKeyUUID,omitempty"`
	// Number that indicates the version of MongoDB that the replica set primary ran when MongoDB Cloud created the snapshot.
	// Read only field.
	MongodVersion *string `json:"mongodVersion,omitempty"`
	// Human-readable label that identifies the replica set.
	// Read only field.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// The node's role at the time when snapshot process began.
	// Read only field.
	ReplicaState *string `json:"replicaState,omitempty"`
	// Number that indicates the total size of space allocated for document storage.
	// Read only field.
	StorageSizeBytes *int64 `json:"storageSizeBytes,omitempty"`
	// Human-readable label that identifies the type of server from which MongoDB Cloud took this snapshot.
	// Read only field.
	TypeName *string `json:"typeName,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *BackupSnapshotPart) MarshalJSON() ([]byte, error) {
	type noMethod BackupSnapshotPart
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewBackupSnapshotPart instantiates a new BackupSnapshotPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupSnapshotPart() *BackupSnapshotPart {
	this := BackupSnapshotPart{}
	return &this
}

// NewBackupSnapshotPartWithDefaults instantiates a new BackupSnapshotPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupSnapshotPartWithDefaults() *BackupSnapshotPart {
	this := BackupSnapshotPart{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *BackupSnapshotPart) SetClusterId(v string) {
	o.ClusterId = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterId")
}

// SetClusterIdNil sets ClusterId to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetClusterIdNil() {
	o.ClusterId = nil
	o.NullFields = addNullField(o.NullFields, "ClusterId")
}

// GetCompletedTime returns the CompletedTime field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetCompletedTime() time.Time {
	if o == nil || IsNil(o.CompletedTime) {
		var ret time.Time
		return ret
	}
	return *o.CompletedTime
}

// GetCompletedTimeOk returns a tuple with the CompletedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetCompletedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedTime) {
		return nil, false
	}

	return o.CompletedTime, true
}

// HasCompletedTime returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasCompletedTime() bool {
	if o != nil && !IsNil(o.CompletedTime) {
		return true
	}

	return false
}

// SetCompletedTime gets a reference to the given time.Time and assigns it to the CompletedTime field.
func (o *BackupSnapshotPart) SetCompletedTime(v time.Time) {
	o.CompletedTime = &v
	o.NullFields = removeNullField(o.NullFields, "CompletedTime")
}

// SetCompletedTimeNil sets CompletedTime to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetCompletedTimeNil() {
	o.CompletedTime = nil
	o.NullFields = addNullField(o.NullFields, "CompletedTime")
}

// GetCompressionSetting returns the CompressionSetting field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetCompressionSetting() string {
	if o == nil || IsNil(o.CompressionSetting) {
		var ret string
		return ret
	}
	return *o.CompressionSetting
}

// GetCompressionSettingOk returns a tuple with the CompressionSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetCompressionSettingOk() (*string, bool) {
	if o == nil || IsNil(o.CompressionSetting) {
		return nil, false
	}

	return o.CompressionSetting, true
}

// HasCompressionSetting returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasCompressionSetting() bool {
	if o != nil && !IsNil(o.CompressionSetting) {
		return true
	}

	return false
}

// SetCompressionSetting gets a reference to the given string and assigns it to the CompressionSetting field.
func (o *BackupSnapshotPart) SetCompressionSetting(v string) {
	o.CompressionSetting = &v
	o.NullFields = removeNullField(o.NullFields, "CompressionSetting")
}

// SetCompressionSettingNil sets CompressionSetting to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetCompressionSettingNil() {
	o.CompressionSetting = nil
	o.NullFields = addNullField(o.NullFields, "CompressionSetting")
}

// GetDataSizeBytes returns the DataSizeBytes field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetDataSizeBytes() int64 {
	if o == nil || IsNil(o.DataSizeBytes) {
		var ret int64
		return ret
	}
	return *o.DataSizeBytes
}

// GetDataSizeBytesOk returns a tuple with the DataSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetDataSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DataSizeBytes) {
		return nil, false
	}

	return o.DataSizeBytes, true
}

// HasDataSizeBytes returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasDataSizeBytes() bool {
	if o != nil && !IsNil(o.DataSizeBytes) {
		return true
	}

	return false
}

// SetDataSizeBytes gets a reference to the given int64 and assigns it to the DataSizeBytes field.
func (o *BackupSnapshotPart) SetDataSizeBytes(v int64) {
	o.DataSizeBytes = &v
	o.NullFields = removeNullField(o.NullFields, "DataSizeBytes")
}

// SetDataSizeBytesNil sets DataSizeBytes to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetDataSizeBytesNil() {
	o.DataSizeBytes = nil
	o.NullFields = addNullField(o.NullFields, "DataSizeBytes")
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetEncryptionEnabled() bool {
	if o == nil || IsNil(o.EncryptionEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptionEnabled) {
		return nil, false
	}

	return o.EncryptionEnabled, true
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasEncryptionEnabled() bool {
	if o != nil && !IsNil(o.EncryptionEnabled) {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given bool and assigns it to the EncryptionEnabled field.
func (o *BackupSnapshotPart) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "EncryptionEnabled")
}

// SetEncryptionEnabledNil sets EncryptionEnabled to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetEncryptionEnabledNil() {
	o.EncryptionEnabled = nil
	o.NullFields = addNullField(o.NullFields, "EncryptionEnabled")
}

// GetFcv returns the Fcv field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetFcv() string {
	if o == nil || IsNil(o.Fcv) {
		var ret string
		return ret
	}
	return *o.Fcv
}

// GetFcvOk returns a tuple with the Fcv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetFcvOk() (*string, bool) {
	if o == nil || IsNil(o.Fcv) {
		return nil, false
	}

	return o.Fcv, true
}

// HasFcv returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasFcv() bool {
	if o != nil && !IsNil(o.Fcv) {
		return true
	}

	return false
}

// SetFcv gets a reference to the given string and assigns it to the Fcv field.
func (o *BackupSnapshotPart) SetFcv(v string) {
	o.Fcv = &v
	o.NullFields = removeNullField(o.NullFields, "Fcv")
}

// SetFcvNil sets Fcv to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetFcvNil() {
	o.Fcv = nil
	o.NullFields = addNullField(o.NullFields, "Fcv")
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetFileSizeBytes() int64 {
	if o == nil || IsNil(o.FileSizeBytes) {
		var ret int64
		return ret
	}
	return *o.FileSizeBytes
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetFileSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSizeBytes) {
		return nil, false
	}

	return o.FileSizeBytes, true
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasFileSizeBytes() bool {
	if o != nil && !IsNil(o.FileSizeBytes) {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given int64 and assigns it to the FileSizeBytes field.
func (o *BackupSnapshotPart) SetFileSizeBytes(v int64) {
	o.FileSizeBytes = &v
	o.NullFields = removeNullField(o.NullFields, "FileSizeBytes")
}

// SetFileSizeBytesNil sets FileSizeBytes to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetFileSizeBytesNil() {
	o.FileSizeBytes = nil
	o.NullFields = addNullField(o.NullFields, "FileSizeBytes")
}

// GetMachineId returns the MachineId field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetMachineId() string {
	if o == nil || IsNil(o.MachineId) {
		var ret string
		return ret
	}
	return *o.MachineId
}

// GetMachineIdOk returns a tuple with the MachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetMachineIdOk() (*string, bool) {
	if o == nil || IsNil(o.MachineId) {
		return nil, false
	}

	return o.MachineId, true
}

// HasMachineId returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasMachineId() bool {
	if o != nil && !IsNil(o.MachineId) {
		return true
	}

	return false
}

// SetMachineId gets a reference to the given string and assigns it to the MachineId field.
func (o *BackupSnapshotPart) SetMachineId(v string) {
	o.MachineId = &v
	o.NullFields = removeNullField(o.NullFields, "MachineId")
}

// SetMachineIdNil sets MachineId to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetMachineIdNil() {
	o.MachineId = nil
	o.NullFields = addNullField(o.NullFields, "MachineId")
}

// GetMasterKeyUUID returns the MasterKeyUUID field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetMasterKeyUUID() string {
	if o == nil || IsNil(o.MasterKeyUUID) {
		var ret string
		return ret
	}
	return *o.MasterKeyUUID
}

// GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetMasterKeyUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.MasterKeyUUID) {
		return nil, false
	}

	return o.MasterKeyUUID, true
}

// HasMasterKeyUUID returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasMasterKeyUUID() bool {
	if o != nil && !IsNil(o.MasterKeyUUID) {
		return true
	}

	return false
}

// SetMasterKeyUUID gets a reference to the given string and assigns it to the MasterKeyUUID field.
func (o *BackupSnapshotPart) SetMasterKeyUUID(v string) {
	o.MasterKeyUUID = &v
	o.NullFields = removeNullField(o.NullFields, "MasterKeyUUID")
}

// SetMasterKeyUUIDNil sets MasterKeyUUID to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetMasterKeyUUIDNil() {
	o.MasterKeyUUID = nil
	o.NullFields = addNullField(o.NullFields, "MasterKeyUUID")
}

// GetMongodVersion returns the MongodVersion field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetMongodVersion() string {
	if o == nil || IsNil(o.MongodVersion) {
		var ret string
		return ret
	}
	return *o.MongodVersion
}

// GetMongodVersionOk returns a tuple with the MongodVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetMongodVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongodVersion) {
		return nil, false
	}

	return o.MongodVersion, true
}

// HasMongodVersion returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasMongodVersion() bool {
	if o != nil && !IsNil(o.MongodVersion) {
		return true
	}

	return false
}

// SetMongodVersion gets a reference to the given string and assigns it to the MongodVersion field.
func (o *BackupSnapshotPart) SetMongodVersion(v string) {
	o.MongodVersion = &v
	o.NullFields = removeNullField(o.NullFields, "MongodVersion")
}

// SetMongodVersionNil sets MongodVersion to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetMongodVersionNil() {
	o.MongodVersion = nil
	o.NullFields = addNullField(o.NullFields, "MongodVersion")
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}

	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *BackupSnapshotPart) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
	o.NullFields = removeNullField(o.NullFields, "ReplicaSetName")
}

// SetReplicaSetNameNil sets ReplicaSetName to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetReplicaSetNameNil() {
	o.ReplicaSetName = nil
	o.NullFields = addNullField(o.NullFields, "ReplicaSetName")
}

// GetReplicaState returns the ReplicaState field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetReplicaState() string {
	if o == nil || IsNil(o.ReplicaState) {
		var ret string
		return ret
	}
	return *o.ReplicaState
}

// GetReplicaStateOk returns a tuple with the ReplicaState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetReplicaStateOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaState) {
		return nil, false
	}

	return o.ReplicaState, true
}

// HasReplicaState returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasReplicaState() bool {
	if o != nil && !IsNil(o.ReplicaState) {
		return true
	}

	return false
}

// SetReplicaState gets a reference to the given string and assigns it to the ReplicaState field.
func (o *BackupSnapshotPart) SetReplicaState(v string) {
	o.ReplicaState = &v
	o.NullFields = removeNullField(o.NullFields, "ReplicaState")
}

// SetReplicaStateNil sets ReplicaState to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetReplicaStateNil() {
	o.ReplicaState = nil
	o.NullFields = addNullField(o.NullFields, "ReplicaState")
}

// GetStorageSizeBytes returns the StorageSizeBytes field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetStorageSizeBytes() int64 {
	if o == nil || IsNil(o.StorageSizeBytes) {
		var ret int64
		return ret
	}
	return *o.StorageSizeBytes
}

// GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetStorageSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.StorageSizeBytes) {
		return nil, false
	}

	return o.StorageSizeBytes, true
}

// HasStorageSizeBytes returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasStorageSizeBytes() bool {
	if o != nil && !IsNil(o.StorageSizeBytes) {
		return true
	}

	return false
}

// SetStorageSizeBytes gets a reference to the given int64 and assigns it to the StorageSizeBytes field.
func (o *BackupSnapshotPart) SetStorageSizeBytes(v int64) {
	o.StorageSizeBytes = &v
	o.NullFields = removeNullField(o.NullFields, "StorageSizeBytes")
}

// SetStorageSizeBytesNil sets StorageSizeBytes to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetStorageSizeBytesNil() {
	o.StorageSizeBytes = nil
	o.NullFields = addNullField(o.NullFields, "StorageSizeBytes")
}

// GetTypeName returns the TypeName field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetTypeName() string {
	if o == nil || IsNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TypeName) {
		return nil, false
	}

	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasTypeName() bool {
	if o != nil && !IsNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *BackupSnapshotPart) SetTypeName(v string) {
	o.TypeName = &v
	o.NullFields = removeNullField(o.NullFields, "TypeName")
}

// SetTypeNameNil sets TypeName to an explicit JSON null when marshaled.
func (o *BackupSnapshotPart) SetTypeNameNil() {
	o.TypeName = nil
	o.NullFields = addNullField(o.NullFields, "TypeName")
}
