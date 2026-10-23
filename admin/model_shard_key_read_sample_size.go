// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardKeyReadSampleSize Number of sampled read operations that MongoDB Cloud used to compute the read distribution, broken down by the command that issued them.
type ShardKeyReadSampleSize struct {
	// Number of sampled `aggregate` operations.
	// Read only field.
	Aggregate *int64 `json:"aggregate,omitempty"`
	// Number of sampled `count` operations.
	// Read only field.
	Count *int64 `json:"count,omitempty"`
	// Number of sampled `distinct` operations.
	// Read only field.
	Distinct *int64 `json:"distinct,omitempty"`
	// Number of sampled `find` operations.
	// Read only field.
	Find *int64 `json:"find,omitempty"`
	// Total number of sampled read operations.
	// Read only field.
	Total *int64 `json:"total,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardKeyReadSampleSize) MarshalJSON() ([]byte, error) {
	type noMethod ShardKeyReadSampleSize
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardKeyReadSampleSize instantiates a new ShardKeyReadSampleSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardKeyReadSampleSize() *ShardKeyReadSampleSize {
	this := ShardKeyReadSampleSize{}
	return &this
}

// NewShardKeyReadSampleSizeWithDefaults instantiates a new ShardKeyReadSampleSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardKeyReadSampleSizeWithDefaults() *ShardKeyReadSampleSize {
	this := ShardKeyReadSampleSize{}
	return &this
}

// GetAggregate returns the Aggregate field value if set, zero value otherwise
func (o *ShardKeyReadSampleSize) GetAggregate() int64 {
	if o == nil || IsNil(o.Aggregate) {
		var ret int64
		return ret
	}
	return *o.Aggregate
}

// GetAggregateOk returns a tuple with the Aggregate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyReadSampleSize) GetAggregateOk() (*int64, bool) {
	if o == nil || IsNil(o.Aggregate) {
		return nil, false
	}

	return o.Aggregate, true
}

// HasAggregate returns a boolean if a field has been set.
func (o *ShardKeyReadSampleSize) HasAggregate() bool {
	if o != nil && !IsNil(o.Aggregate) {
		return true
	}

	return false
}

// SetAggregate gets a reference to the given int64 and assigns it to the Aggregate field.
func (o *ShardKeyReadSampleSize) SetAggregate(v int64) {
	o.Aggregate = &v
	o.NullFields = removeNullField(o.NullFields, "Aggregate")
}

// SetAggregateNil sets Aggregate to an explicit JSON null when marshaled.
func (o *ShardKeyReadSampleSize) SetAggregateNil() {
	o.Aggregate = nil
	o.NullFields = addNullField(o.NullFields, "Aggregate")
}

// GetCount returns the Count field value if set, zero value otherwise
func (o *ShardKeyReadSampleSize) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyReadSampleSize) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}

	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ShardKeyReadSampleSize) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ShardKeyReadSampleSize) SetCount(v int64) {
	o.Count = &v
	o.NullFields = removeNullField(o.NullFields, "Count")
}

// SetCountNil sets Count to an explicit JSON null when marshaled.
func (o *ShardKeyReadSampleSize) SetCountNil() {
	o.Count = nil
	o.NullFields = addNullField(o.NullFields, "Count")
}

// GetDistinct returns the Distinct field value if set, zero value otherwise
func (o *ShardKeyReadSampleSize) GetDistinct() int64 {
	if o == nil || IsNil(o.Distinct) {
		var ret int64
		return ret
	}
	return *o.Distinct
}

// GetDistinctOk returns a tuple with the Distinct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyReadSampleSize) GetDistinctOk() (*int64, bool) {
	if o == nil || IsNil(o.Distinct) {
		return nil, false
	}

	return o.Distinct, true
}

// HasDistinct returns a boolean if a field has been set.
func (o *ShardKeyReadSampleSize) HasDistinct() bool {
	if o != nil && !IsNil(o.Distinct) {
		return true
	}

	return false
}

// SetDistinct gets a reference to the given int64 and assigns it to the Distinct field.
func (o *ShardKeyReadSampleSize) SetDistinct(v int64) {
	o.Distinct = &v
	o.NullFields = removeNullField(o.NullFields, "Distinct")
}

// SetDistinctNil sets Distinct to an explicit JSON null when marshaled.
func (o *ShardKeyReadSampleSize) SetDistinctNil() {
	o.Distinct = nil
	o.NullFields = addNullField(o.NullFields, "Distinct")
}

// GetFind returns the Find field value if set, zero value otherwise
func (o *ShardKeyReadSampleSize) GetFind() int64 {
	if o == nil || IsNil(o.Find) {
		var ret int64
		return ret
	}
	return *o.Find
}

// GetFindOk returns a tuple with the Find field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyReadSampleSize) GetFindOk() (*int64, bool) {
	if o == nil || IsNil(o.Find) {
		return nil, false
	}

	return o.Find, true
}

// HasFind returns a boolean if a field has been set.
func (o *ShardKeyReadSampleSize) HasFind() bool {
	if o != nil && !IsNil(o.Find) {
		return true
	}

	return false
}

// SetFind gets a reference to the given int64 and assigns it to the Find field.
func (o *ShardKeyReadSampleSize) SetFind(v int64) {
	o.Find = &v
	o.NullFields = removeNullField(o.NullFields, "Find")
}

// SetFindNil sets Find to an explicit JSON null when marshaled.
func (o *ShardKeyReadSampleSize) SetFindNil() {
	o.Find = nil
	o.NullFields = addNullField(o.NullFields, "Find")
}

// GetTotal returns the Total field value if set, zero value otherwise
func (o *ShardKeyReadSampleSize) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyReadSampleSize) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}

	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ShardKeyReadSampleSize) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *ShardKeyReadSampleSize) SetTotal(v int64) {
	o.Total = &v
	o.NullFields = removeNullField(o.NullFields, "Total")
}

// SetTotalNil sets Total to an explicit JSON null when marshaled.
func (o *ShardKeyReadSampleSize) SetTotalNil() {
	o.Total = nil
	o.NullFields = addNullField(o.NullFields, "Total")
}
