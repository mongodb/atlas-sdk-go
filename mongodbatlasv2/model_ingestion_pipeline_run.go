/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the IngestionPipelineRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestionPipelineRun{}

// IngestionPipelineRun Run details of a Data Lake Pipeline.
type IngestionPipelineRun struct {
	// Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.
	Id *string `json:"_id,omitempty"`
	// Backup schedule interval of the Data Lake Pipeline.
	BackupFrequencyType *string `json:"backupFrequencyType,omitempty"`
	// Timestamp that indicates when the pipeline run was created.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Human-readable label that identifies the dataset that Atlas generates during this pipeline run. You can use this dataset as a `dataSource` in a Federated Database collection.
	DatasetName *string `json:"datasetName,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	GroupId *string `json:"groupId,omitempty"`
	// Timestamp that indicates the last time that the pipeline run was updated.
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
	// Processing phase of the Data Lake Pipeline.
	Phase *string `json:"phase,omitempty"`
	// Unique 24-hexadecimal character string that identifies a Data Lake Pipeline.
	PipelineId *string `json:"pipelineId,omitempty"`
	// Unique 24-hexadecimal character string that identifies the snapshot of a cluster.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// State of the pipeline run.
	State *string `json:"state,omitempty"`
	Stats *PipelineRunStats `json:"stats,omitempty"`
}

// NewIngestionPipelineRun instantiates a new IngestionPipelineRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionPipelineRun() *IngestionPipelineRun {
	this := IngestionPipelineRun{}
	return &this
}

// NewIngestionPipelineRunWithDefaults instantiates a new IngestionPipelineRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionPipelineRunWithDefaults() *IngestionPipelineRun {
	this := IngestionPipelineRun{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IngestionPipelineRun) SetId(v string) {
	o.Id = &v
}

// GetBackupFrequencyType returns the BackupFrequencyType field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetBackupFrequencyType() string {
	if o == nil || IsNil(o.BackupFrequencyType) {
		var ret string
		return ret
	}
	return *o.BackupFrequencyType
}

// GetBackupFrequencyTypeOk returns a tuple with the BackupFrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetBackupFrequencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BackupFrequencyType) {
		return nil, false
	}
	return o.BackupFrequencyType, true
}

// HasBackupFrequencyType returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasBackupFrequencyType() bool {
	if o != nil && !IsNil(o.BackupFrequencyType) {
		return true
	}

	return false
}

// SetBackupFrequencyType gets a reference to the given string and assigns it to the BackupFrequencyType field.
func (o *IngestionPipelineRun) SetBackupFrequencyType(v string) {
	o.BackupFrequencyType = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *IngestionPipelineRun) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDatasetName returns the DatasetName field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetDatasetName() string {
	if o == nil || IsNil(o.DatasetName) {
		var ret string
		return ret
	}
	return *o.DatasetName
}

// GetDatasetNameOk returns a tuple with the DatasetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetDatasetNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetName) {
		return nil, false
	}
	return o.DatasetName, true
}

// HasDatasetName returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasDatasetName() bool {
	if o != nil && !IsNil(o.DatasetName) {
		return true
	}

	return false
}

// SetDatasetName gets a reference to the given string and assigns it to the DatasetName field.
func (o *IngestionPipelineRun) SetDatasetName(v string) {
	o.DatasetName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *IngestionPipelineRun) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *IngestionPipelineRun) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetPhase() string {
	if o == nil || IsNil(o.Phase) {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetPhaseOk() (*string, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *IngestionPipelineRun) SetPhase(v string) {
	o.Phase = &v
}

// GetPipelineId returns the PipelineId field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetPipelineId() string {
	if o == nil || IsNil(o.PipelineId) {
		var ret string
		return ret
	}
	return *o.PipelineId
}

// GetPipelineIdOk returns a tuple with the PipelineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetPipelineIdOk() (*string, bool) {
	if o == nil || IsNil(o.PipelineId) {
		return nil, false
	}
	return o.PipelineId, true
}

// HasPipelineId returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasPipelineId() bool {
	if o != nil && !IsNil(o.PipelineId) {
		return true
	}

	return false
}

// SetPipelineId gets a reference to the given string and assigns it to the PipelineId field.
func (o *IngestionPipelineRun) SetPipelineId(v string) {
	o.PipelineId = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *IngestionPipelineRun) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IngestionPipelineRun) SetState(v string) {
	o.State = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetStats() PipelineRunStats {
	if o == nil || IsNil(o.Stats) {
		var ret PipelineRunStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetStatsOk() (*PipelineRunStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given PipelineRunStats and assigns it to the Stats field.
func (o *IngestionPipelineRun) SetStats(v PipelineRunStats) {
	o.Stats = &v
}

func (o IngestionPipelineRun) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o IngestionPipelineRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	return toSerialize, nil
}

type NullableIngestionPipelineRun struct {
	value *IngestionPipelineRun
	isSet bool
}

func (v NullableIngestionPipelineRun) Get() *IngestionPipelineRun {
	return v.value
}

func (v *NullableIngestionPipelineRun) Set(val *IngestionPipelineRun) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionPipelineRun) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionPipelineRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionPipelineRun(val *IngestionPipelineRun) *NullableIngestionPipelineRun {
	return &NullableIngestionPipelineRun{value: val, isSet: true}
}

func (v NullableIngestionPipelineRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionPipelineRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


