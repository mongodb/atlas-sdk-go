// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TriggerIngestionPipelineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerIngestionPipelineRequest{}

// TriggerIngestionPipelineRequest struct for TriggerIngestionPipelineRequest
type TriggerIngestionPipelineRequest struct {
	// Unique 24-hexadecimal character string that identifies the snapshot.
	SnapshotId string `json:"snapshotId"`
}

// NewTriggerIngestionPipelineRequest instantiates a new TriggerIngestionPipelineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerIngestionPipelineRequest(snapshotId string) *TriggerIngestionPipelineRequest {
	this := TriggerIngestionPipelineRequest{}
	this.SnapshotId = snapshotId
	return &this
}

// NewTriggerIngestionPipelineRequestWithDefaults instantiates a new TriggerIngestionPipelineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerIngestionPipelineRequestWithDefaults() *TriggerIngestionPipelineRequest {
	this := TriggerIngestionPipelineRequest{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value
func (o *TriggerIngestionPipelineRequest) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *TriggerIngestionPipelineRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *TriggerIngestionPipelineRequest) SetSnapshotId(v string) {
	o.SnapshotId = v
}

func (o TriggerIngestionPipelineRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TriggerIngestionPipelineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snapshotId"] = o.SnapshotId
	return toSerialize, nil
}

type NullableTriggerIngestionPipelineRequest struct {
	value *TriggerIngestionPipelineRequest
	isSet bool
}

func (v NullableTriggerIngestionPipelineRequest) Get() *TriggerIngestionPipelineRequest {
	return v.value
}

func (v *NullableTriggerIngestionPipelineRequest) Set(val *TriggerIngestionPipelineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerIngestionPipelineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerIngestionPipelineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerIngestionPipelineRequest(val *TriggerIngestionPipelineRequest) *NullableTriggerIngestionPipelineRequest {
	return &NullableTriggerIngestionPipelineRequest{value: val, isSet: true}
}

func (v NullableTriggerIngestionPipelineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerIngestionPipelineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
