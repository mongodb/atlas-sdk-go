// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardMeasurement Time series of one measurement on one shard of the cluster.
type ShardMeasurement struct {
	// Data points of the measurement at `granularity` intervals, starting at `start` and not passing `end`, in ascending time order.
	// Read only field.
	DataPoints []ShardMeasurementDataPoint `json:"dataPoints"`
	// Measurement that the data points quantify.
	// Read only field.
	Name string `json:"name"`
	// Human-readable label that identifies the shard, as the replica set name of its members.
	// Read only field.
	ReplicaSetName string `json:"replicaSetName"`
	// Unit in which the data points express the measurement.
	// Read only field.
	Units string `json:"units"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardMeasurement) MarshalJSON() ([]byte, error) {
	type noMethod ShardMeasurement
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardMeasurement instantiates a new ShardMeasurement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardMeasurement(dataPoints []ShardMeasurementDataPoint, name string, replicaSetName string, units string) *ShardMeasurement {
	this := ShardMeasurement{}
	this.DataPoints = dataPoints
	this.Name = name
	this.ReplicaSetName = replicaSetName
	this.Units = units
	return &this
}

// NewShardMeasurementWithDefaults instantiates a new ShardMeasurement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardMeasurementWithDefaults() *ShardMeasurement {
	this := ShardMeasurement{}
	return &this
}

// GetDataPoints returns the DataPoints field value
func (o *ShardMeasurement) GetDataPoints() []ShardMeasurementDataPoint {
	if o == nil {
		var ret []ShardMeasurementDataPoint
		return ret
	}

	return o.DataPoints
}

// GetDataPointsOk returns a tuple with the DataPoints field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurement) GetDataPointsOk() (*[]ShardMeasurementDataPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataPoints, true
}

// SetDataPoints sets field value
func (o *ShardMeasurement) SetDataPoints(v []ShardMeasurementDataPoint) {
	o.DataPoints = v
}

// GetName returns the Name field value
func (o *ShardMeasurement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurement) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShardMeasurement) SetName(v string) {
	o.Name = v
}

// GetReplicaSetName returns the ReplicaSetName field value
func (o *ShardMeasurement) GetReplicaSetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurement) GetReplicaSetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicaSetName, true
}

// SetReplicaSetName sets field value
func (o *ShardMeasurement) SetReplicaSetName(v string) {
	o.ReplicaSetName = v
}

// GetUnits returns the Units field value
func (o *ShardMeasurement) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurement) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *ShardMeasurement) SetUnits(v string) {
	o.Units = v
}
