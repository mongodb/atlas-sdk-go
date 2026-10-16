// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// ShardMeasurementsResponse Measurement time series of one shard over the requested window.
type ShardMeasurementsResponse struct {
	// Human-readable label that identifies the cluster.
	// Read only field.
	ClusterName string `json:"clusterName"`
	// Date and time at which the window ends. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	End time.Time `json:"end"`
	// Interval between consecutive data points, in ISO 8601 duration format. Either the requested `granularity` or, when none was requested, the resolution MongoDB Cloud chose for the width of the window.
	// Read only field.
	Granularity string `json:"granularity"`
	// Unique 24-hexadecimal digit string that identifies the project.
	// Read only field.
	GroupId string `json:"groupId"`
	// One time series per requested measurement, in the order of the measurement names. A measurement MongoDB Cloud has no readings for is still listed, with `null` data point values.
	// Read only field.
	Measurements []ShardMeasurementSeries `json:"measurements"`
	// Human-readable label that identifies the shard, as the replica set name of its members.
	// Read only field.
	ReplicaSetName string `json:"replicaSetName"`
	// Date and time of the first data point. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	Start time.Time `json:"start"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardMeasurementsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ShardMeasurementsResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardMeasurementsResponse instantiates a new ShardMeasurementsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardMeasurementsResponse(clusterName string, end time.Time, granularity string, groupId string, measurements []ShardMeasurementSeries, replicaSetName string, start time.Time) *ShardMeasurementsResponse {
	this := ShardMeasurementsResponse{}
	this.ClusterName = clusterName
	this.End = end
	this.Granularity = granularity
	this.GroupId = groupId
	this.Measurements = measurements
	this.ReplicaSetName = replicaSetName
	this.Start = start
	return &this
}

// NewShardMeasurementsResponseWithDefaults instantiates a new ShardMeasurementsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardMeasurementsResponseWithDefaults() *ShardMeasurementsResponse {
	this := ShardMeasurementsResponse{}
	return &this
}

// GetClusterName returns the ClusterName field value
func (o *ShardMeasurementsResponse) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementsResponse) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *ShardMeasurementsResponse) SetClusterName(v string) {
	o.ClusterName = v
}

// GetEnd returns the End field value
func (o *ShardMeasurementsResponse) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementsResponse) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ShardMeasurementsResponse) SetEnd(v time.Time) {
	o.End = v
}

// GetGranularity returns the Granularity field value
func (o *ShardMeasurementsResponse) GetGranularity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementsResponse) GetGranularityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *ShardMeasurementsResponse) SetGranularity(v string) {
	o.Granularity = v
}

// GetGroupId returns the GroupId field value
func (o *ShardMeasurementsResponse) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementsResponse) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ShardMeasurementsResponse) SetGroupId(v string) {
	o.GroupId = v
}

// GetMeasurements returns the Measurements field value
func (o *ShardMeasurementsResponse) GetMeasurements() []ShardMeasurementSeries {
	if o == nil {
		var ret []ShardMeasurementSeries
		return ret
	}

	return o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementsResponse) GetMeasurementsOk() (*[]ShardMeasurementSeries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Measurements, true
}

// SetMeasurements sets field value
func (o *ShardMeasurementsResponse) SetMeasurements(v []ShardMeasurementSeries) {
	o.Measurements = v
}

// GetReplicaSetName returns the ReplicaSetName field value
func (o *ShardMeasurementsResponse) GetReplicaSetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementsResponse) GetReplicaSetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicaSetName, true
}

// SetReplicaSetName sets field value
func (o *ShardMeasurementsResponse) SetReplicaSetName(v string) {
	o.ReplicaSetName = v
}

// GetStart returns the Start field value
func (o *ShardMeasurementsResponse) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementsResponse) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ShardMeasurementsResponse) SetStart(v time.Time) {
	o.Start = v
}
