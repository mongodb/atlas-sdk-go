// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// LoadSheddingSimulationResponse Load shedding simulation for a cluster.
type LoadSheddingSimulationResponse struct {
	// Human-readable label that identifies the cluster on which the simulation is running.
	// Read only field.
	ClusterName *string `json:"clusterName,omitempty"`
	// Duration of the load shedding simulation in seconds.
	DurationSeconds *int `json:"durationSeconds,omitempty"`
	// Date and time when the load shedding simulation expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project that contains the cluster.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Date and time when the load shedding simulation was requested. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	RequestDate *time.Time `json:"requestDate,omitempty"`
	// Unique identifier of the load shedding simulation.
	// Read only field.
	SimulationId *string `json:"simulationId,omitempty"`
	// Current state of the load shedding simulation.
	// Read only field.
	State *string `json:"state,omitempty"`
}

// NewLoadSheddingSimulationResponse instantiates a new LoadSheddingSimulationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadSheddingSimulationResponse() *LoadSheddingSimulationResponse {
	this := LoadSheddingSimulationResponse{}
	return &this
}

// NewLoadSheddingSimulationResponseWithDefaults instantiates a new LoadSheddingSimulationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadSheddingSimulationResponseWithDefaults() *LoadSheddingSimulationResponse {
	this := LoadSheddingSimulationResponse{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *LoadSheddingSimulationResponse) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSheddingSimulationResponse) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *LoadSheddingSimulationResponse) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *LoadSheddingSimulationResponse) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise
func (o *LoadSheddingSimulationResponse) GetDurationSeconds() int {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret int
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSheddingSimulationResponse) GetDurationSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}

	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *LoadSheddingSimulationResponse) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given int and assigns it to the DurationSeconds field.
func (o *LoadSheddingSimulationResponse) SetDurationSeconds(v int) {
	o.DurationSeconds = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise
func (o *LoadSheddingSimulationResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSheddingSimulationResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}

	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *LoadSheddingSimulationResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *LoadSheddingSimulationResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *LoadSheddingSimulationResponse) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSheddingSimulationResponse) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *LoadSheddingSimulationResponse) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *LoadSheddingSimulationResponse) SetGroupId(v string) {
	o.GroupId = &v
}

// GetRequestDate returns the RequestDate field value if set, zero value otherwise
func (o *LoadSheddingSimulationResponse) GetRequestDate() time.Time {
	if o == nil || IsNil(o.RequestDate) {
		var ret time.Time
		return ret
	}
	return *o.RequestDate
}

// GetRequestDateOk returns a tuple with the RequestDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSheddingSimulationResponse) GetRequestDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestDate) {
		return nil, false
	}

	return o.RequestDate, true
}

// HasRequestDate returns a boolean if a field has been set.
func (o *LoadSheddingSimulationResponse) HasRequestDate() bool {
	if o != nil && !IsNil(o.RequestDate) {
		return true
	}

	return false
}

// SetRequestDate gets a reference to the given time.Time and assigns it to the RequestDate field.
func (o *LoadSheddingSimulationResponse) SetRequestDate(v time.Time) {
	o.RequestDate = &v
}

// GetSimulationId returns the SimulationId field value if set, zero value otherwise
func (o *LoadSheddingSimulationResponse) GetSimulationId() string {
	if o == nil || IsNil(o.SimulationId) {
		var ret string
		return ret
	}
	return *o.SimulationId
}

// GetSimulationIdOk returns a tuple with the SimulationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSheddingSimulationResponse) GetSimulationIdOk() (*string, bool) {
	if o == nil || IsNil(o.SimulationId) {
		return nil, false
	}

	return o.SimulationId, true
}

// HasSimulationId returns a boolean if a field has been set.
func (o *LoadSheddingSimulationResponse) HasSimulationId() bool {
	if o != nil && !IsNil(o.SimulationId) {
		return true
	}

	return false
}

// SetSimulationId gets a reference to the given string and assigns it to the SimulationId field.
func (o *LoadSheddingSimulationResponse) SetSimulationId(v string) {
	o.SimulationId = &v
}

// GetState returns the State field value if set, zero value otherwise
func (o *LoadSheddingSimulationResponse) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSheddingSimulationResponse) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LoadSheddingSimulationResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LoadSheddingSimulationResponse) SetState(v string) {
	o.State = &v
}
