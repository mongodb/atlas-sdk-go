// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// ClusterShardMeasurementsResponse Measurement time series for every shard of one sharded cluster over the requested window, one series per shard and measurement.
type ClusterShardMeasurementsResponse struct {
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
	// One time series per requested measurement on each shard of the cluster, ordered by measurement and then by `replicaSetName`. A shard that MongoDB Cloud has no readings for is still listed, with `null` data point values.
	// Read only field.
	Measurements []ShardMeasurement `json:"measurements"`
	// Date and time of the first data point. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	Start time.Time `json:"start"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ClusterShardMeasurementsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ClusterShardMeasurementsResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewClusterShardMeasurementsResponse instantiates a new ClusterShardMeasurementsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterShardMeasurementsResponse(clusterName string, end time.Time, granularity string, groupId string, measurements []ShardMeasurement, start time.Time) *ClusterShardMeasurementsResponse {
	this := ClusterShardMeasurementsResponse{}
	this.ClusterName = clusterName
	this.End = end
	this.Granularity = granularity
	this.GroupId = groupId
	this.Measurements = measurements
	this.Start = start
	return &this
}

// NewClusterShardMeasurementsResponseWithDefaults instantiates a new ClusterShardMeasurementsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterShardMeasurementsResponseWithDefaults() *ClusterShardMeasurementsResponse {
	this := ClusterShardMeasurementsResponse{}
	return &this
}

// GetClusterName returns the ClusterName field value
func (o *ClusterShardMeasurementsResponse) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *ClusterShardMeasurementsResponse) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *ClusterShardMeasurementsResponse) SetClusterName(v string) {
	o.ClusterName = v
}

// GetEnd returns the End field value
func (o *ClusterShardMeasurementsResponse) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ClusterShardMeasurementsResponse) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ClusterShardMeasurementsResponse) SetEnd(v time.Time) {
	o.End = v
}

// GetGranularity returns the Granularity field value
func (o *ClusterShardMeasurementsResponse) GetGranularity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *ClusterShardMeasurementsResponse) GetGranularityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *ClusterShardMeasurementsResponse) SetGranularity(v string) {
	o.Granularity = v
}

// GetGroupId returns the GroupId field value
func (o *ClusterShardMeasurementsResponse) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ClusterShardMeasurementsResponse) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ClusterShardMeasurementsResponse) SetGroupId(v string) {
	o.GroupId = v
}

// GetMeasurements returns the Measurements field value
func (o *ClusterShardMeasurementsResponse) GetMeasurements() []ShardMeasurement {
	if o == nil {
		var ret []ShardMeasurement
		return ret
	}

	return o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value
// and a boolean to check if the value has been set.
func (o *ClusterShardMeasurementsResponse) GetMeasurementsOk() (*[]ShardMeasurement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Measurements, true
}

// SetMeasurements sets field value
func (o *ClusterShardMeasurementsResponse) SetMeasurements(v []ShardMeasurement) {
	o.Measurements = v
}

// GetStart returns the Start field value
func (o *ClusterShardMeasurementsResponse) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ClusterShardMeasurementsResponse) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ClusterShardMeasurementsResponse) SetStart(v time.Time) {
	o.Start = v
}
