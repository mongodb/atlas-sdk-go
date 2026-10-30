// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardMeasurementSeries Time series of one measurement on the shard named by the enclosing response.
type ShardMeasurementSeries struct {
	// Data points of the measurement at `granularity` intervals, starting at `start` and not passing `end`, in ascending time order.
	// Read only field.
	DataPoints []ShardMeasurementDataPoint `json:"dataPoints"`
	// Measurement that the data points quantify.
	// Read only field.
	Name string `json:"name"`
	// Unit in which the data points express the measurement.
	// Read only field.
	Units string `json:"units"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardMeasurementSeries) MarshalJSON() ([]byte, error) {
	type noMethod ShardMeasurementSeries
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardMeasurementSeries instantiates a new ShardMeasurementSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardMeasurementSeries(dataPoints []ShardMeasurementDataPoint, name string, units string) *ShardMeasurementSeries {
	this := ShardMeasurementSeries{}
	this.DataPoints = dataPoints
	this.Name = name
	this.Units = units
	return &this
}

// NewShardMeasurementSeriesWithDefaults instantiates a new ShardMeasurementSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardMeasurementSeriesWithDefaults() *ShardMeasurementSeries {
	this := ShardMeasurementSeries{}
	return &this
}

// GetDataPoints returns the DataPoints field value
func (o *ShardMeasurementSeries) GetDataPoints() []ShardMeasurementDataPoint {
	if o == nil {
		var ret []ShardMeasurementDataPoint
		return ret
	}

	return o.DataPoints
}

// GetDataPointsOk returns a tuple with the DataPoints field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementSeries) GetDataPointsOk() (*[]ShardMeasurementDataPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataPoints, true
}

// SetDataPoints sets field value
func (o *ShardMeasurementSeries) SetDataPoints(v []ShardMeasurementDataPoint) {
	o.DataPoints = v
}

// GetName returns the Name field value
func (o *ShardMeasurementSeries) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementSeries) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShardMeasurementSeries) SetName(v string) {
	o.Name = v
}

// GetUnits returns the Units field value
func (o *ShardMeasurementSeries) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementSeries) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *ShardMeasurementSeries) SetUnits(v string) {
	o.Units = v
}
