// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardOperations Operations per second that the shard served over the last hour.
type ShardOperations struct {
	// Average operations per second across the shard's nodes over the last hour.
	// Read only field.
	Average float64 `json:"average"`
	// Highest operations per second that any one of the shard's nodes reached over the last hour.
	// Read only field.
	Maximum float64 `json:"maximum"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardOperations) MarshalJSON() ([]byte, error) {
	type noMethod ShardOperations
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardOperations instantiates a new ShardOperations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardOperations(average float64, maximum float64) *ShardOperations {
	this := ShardOperations{}
	this.Average = average
	this.Maximum = maximum
	return &this
}

// NewShardOperationsWithDefaults instantiates a new ShardOperations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardOperationsWithDefaults() *ShardOperations {
	this := ShardOperations{}
	return &this
}

// GetAverage returns the Average field value
func (o *ShardOperations) GetAverage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Average
}

// GetAverageOk returns a tuple with the Average field value
// and a boolean to check if the value has been set.
func (o *ShardOperations) GetAverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Average, true
}

// SetAverage sets field value
func (o *ShardOperations) SetAverage(v float64) {
	o.Average = v
}

// GetMaximum returns the Maximum field value
func (o *ShardOperations) GetMaximum() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value
// and a boolean to check if the value has been set.
func (o *ShardOperations) GetMaximumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maximum, true
}

// SetMaximum sets field value
func (o *ShardOperations) SetMaximum(v float64) {
	o.Maximum = v
}
