// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// OverloadProtectionSimulationResponse Overload protection simulation for a cluster.
type OverloadProtectionSimulationResponse struct {
	// Date and time when cancellation of the overload protection simulation was requested. This parameter is only present when a cancellation has been requested and expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	CancelRequestedAt *time.Time `json:"cancelRequestedAt,omitempty"`
	// Human-readable label that identifies the cluster on which the simulation is running.
	// Read only field.
	ClusterName string `json:"clusterName"`
	// Duration of the overload protection simulation in seconds.
	DurationSeconds int `json:"durationSeconds"`
	// Date and time when the overload protection simulation expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	ExpiresAt time.Time `json:"expiresAt"`
	// Unique 24-hexadecimal character string that identifies the project that contains the cluster.
	// Read only field.
	GroupId string `json:"groupId"`
	// Date and time when the overload protection simulation was requested. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	RequestDate time.Time `json:"requestDate"`
	// Unique identifier of the overload protection simulation.
	// Read only field.
	SimulationId string `json:"simulationId"`
	// Current state of the overload protection simulation.
	// Read only field.
	State string `json:"state"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OverloadProtectionSimulationResponse) MarshalJSON() ([]byte, error) {
	type noMethod OverloadProtectionSimulationResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOverloadProtectionSimulationResponse instantiates a new OverloadProtectionSimulationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverloadProtectionSimulationResponse(clusterName string, durationSeconds int, expiresAt time.Time, groupId string, requestDate time.Time, simulationId string, state string) *OverloadProtectionSimulationResponse {
	this := OverloadProtectionSimulationResponse{}
	this.ClusterName = clusterName
	this.DurationSeconds = durationSeconds
	this.ExpiresAt = expiresAt
	this.GroupId = groupId
	this.RequestDate = requestDate
	this.SimulationId = simulationId
	this.State = state
	return &this
}

// NewOverloadProtectionSimulationResponseWithDefaults instantiates a new OverloadProtectionSimulationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverloadProtectionSimulationResponseWithDefaults() *OverloadProtectionSimulationResponse {
	this := OverloadProtectionSimulationResponse{}
	return &this
}

// GetCancelRequestedAt returns the CancelRequestedAt field value if set, zero value otherwise
func (o *OverloadProtectionSimulationResponse) GetCancelRequestedAt() time.Time {
	if o == nil || IsNil(o.CancelRequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.CancelRequestedAt
}

// GetCancelRequestedAtOk returns a tuple with the CancelRequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationResponse) GetCancelRequestedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CancelRequestedAt) {
		return nil, false
	}

	return o.CancelRequestedAt, true
}

// HasCancelRequestedAt returns a boolean if a field has been set.
func (o *OverloadProtectionSimulationResponse) HasCancelRequestedAt() bool {
	if o != nil && !IsNil(o.CancelRequestedAt) {
		return true
	}

	return false
}

// SetCancelRequestedAt gets a reference to the given time.Time and assigns it to the CancelRequestedAt field.
func (o *OverloadProtectionSimulationResponse) SetCancelRequestedAt(v time.Time) {
	o.CancelRequestedAt = &v
	o.NullFields = removeNullField(o.NullFields, "CancelRequestedAt")
}

// SetCancelRequestedAtNil sets CancelRequestedAt to an explicit JSON null when marshaled.
func (o *OverloadProtectionSimulationResponse) SetCancelRequestedAtNil() {
	o.CancelRequestedAt = nil
	o.NullFields = addNullField(o.NullFields, "CancelRequestedAt")
}

// GetClusterName returns the ClusterName field value
func (o *OverloadProtectionSimulationResponse) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationResponse) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *OverloadProtectionSimulationResponse) SetClusterName(v string) {
	o.ClusterName = v
}

// GetDurationSeconds returns the DurationSeconds field value
func (o *OverloadProtectionSimulationResponse) GetDurationSeconds() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationResponse) GetDurationSecondsOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value
func (o *OverloadProtectionSimulationResponse) SetDurationSeconds(v int) {
	o.DurationSeconds = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *OverloadProtectionSimulationResponse) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *OverloadProtectionSimulationResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetGroupId returns the GroupId field value
func (o *OverloadProtectionSimulationResponse) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationResponse) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *OverloadProtectionSimulationResponse) SetGroupId(v string) {
	o.GroupId = v
}

// GetRequestDate returns the RequestDate field value
func (o *OverloadProtectionSimulationResponse) GetRequestDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RequestDate
}

// GetRequestDateOk returns a tuple with the RequestDate field value
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationResponse) GetRequestDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestDate, true
}

// SetRequestDate sets field value
func (o *OverloadProtectionSimulationResponse) SetRequestDate(v time.Time) {
	o.RequestDate = v
}

// GetSimulationId returns the SimulationId field value
func (o *OverloadProtectionSimulationResponse) GetSimulationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SimulationId
}

// GetSimulationIdOk returns a tuple with the SimulationId field value
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationResponse) GetSimulationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SimulationId, true
}

// SetSimulationId sets field value
func (o *OverloadProtectionSimulationResponse) SetSimulationId(v string) {
	o.SimulationId = v
}

// GetState returns the State field value
func (o *OverloadProtectionSimulationResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *OverloadProtectionSimulationResponse) SetState(v string) {
	o.State = v
}
