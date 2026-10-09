// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardCpuUtilization CPU utilization of the shard's nodes over the last hour, as a percentage.
type ShardCpuUtilization struct {
	// Average CPU utilization across the shard's nodes over the last hour.
	// Read only field.
	AveragePercent float64 `json:"averagePercent"`
	// Highest CPU utilization that any one of the shard's nodes reached over the last hour.
	// Read only field.
	MaximumPercent float64 `json:"maximumPercent"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardCpuUtilization) MarshalJSON() ([]byte, error) {
	type noMethod ShardCpuUtilization
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardCpuUtilization instantiates a new ShardCpuUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardCpuUtilization(averagePercent float64, maximumPercent float64) *ShardCpuUtilization {
	this := ShardCpuUtilization{}
	this.AveragePercent = averagePercent
	this.MaximumPercent = maximumPercent
	return &this
}

// NewShardCpuUtilizationWithDefaults instantiates a new ShardCpuUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardCpuUtilizationWithDefaults() *ShardCpuUtilization {
	this := ShardCpuUtilization{}
	return &this
}

// GetAveragePercent returns the AveragePercent field value
func (o *ShardCpuUtilization) GetAveragePercent() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AveragePercent
}

// GetAveragePercentOk returns a tuple with the AveragePercent field value
// and a boolean to check if the value has been set.
func (o *ShardCpuUtilization) GetAveragePercentOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AveragePercent, true
}

// SetAveragePercent sets field value
func (o *ShardCpuUtilization) SetAveragePercent(v float64) {
	o.AveragePercent = v
}

// GetMaximumPercent returns the MaximumPercent field value
func (o *ShardCpuUtilization) GetMaximumPercent() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MaximumPercent
}

// GetMaximumPercentOk returns a tuple with the MaximumPercent field value
// and a boolean to check if the value has been set.
func (o *ShardCpuUtilization) GetMaximumPercentOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumPercent, true
}

// SetMaximumPercent sets field value
func (o *ShardCpuUtilization) SetMaximumPercent(v float64) {
	o.MaximumPercent = v
}
