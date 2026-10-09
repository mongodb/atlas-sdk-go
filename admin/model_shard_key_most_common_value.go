// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardKeyMostCommonValue Shard key value that occurs frequently in the sampled documents, with the number of sampled documents that carry it.
type ShardKeyMostCommonValue struct {
	// Number of sampled documents that carry this shard key value.
	// Read only field.
	Frequency *int64 `json:"frequency,omitempty"`
	// Flag that indicates whether MongoDB Cloud omitted `value` because it was too large to return. When this flag is `true`, `frequency` still describes a real shard key value even though the value itself is absent.
	// Read only field.
	Truncated *bool `json:"truncated,omitempty"`
	// Shard key value, given as one field-value pair per shard key field. This parameter is absent when `truncated` is `true`.
	// Read only field.
	Value any `json:"value,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardKeyMostCommonValue) MarshalJSON() ([]byte, error) {
	type noMethod ShardKeyMostCommonValue
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardKeyMostCommonValue instantiates a new ShardKeyMostCommonValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardKeyMostCommonValue() *ShardKeyMostCommonValue {
	this := ShardKeyMostCommonValue{}
	return &this
}

// NewShardKeyMostCommonValueWithDefaults instantiates a new ShardKeyMostCommonValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardKeyMostCommonValueWithDefaults() *ShardKeyMostCommonValue {
	this := ShardKeyMostCommonValue{}
	return &this
}

// GetFrequency returns the Frequency field value if set, zero value otherwise
func (o *ShardKeyMostCommonValue) GetFrequency() int64 {
	if o == nil || IsNil(o.Frequency) {
		var ret int64
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyMostCommonValue) GetFrequencyOk() (*int64, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}

	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *ShardKeyMostCommonValue) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given int64 and assigns it to the Frequency field.
func (o *ShardKeyMostCommonValue) SetFrequency(v int64) {
	o.Frequency = &v
	o.NullFields = removeNullField(o.NullFields, "Frequency")
}

// SetFrequencyNil sets Frequency to an explicit JSON null when marshaled.
func (o *ShardKeyMostCommonValue) SetFrequencyNil() {
	o.Frequency = nil
	o.NullFields = addNullField(o.NullFields, "Frequency")
}

// GetTruncated returns the Truncated field value if set, zero value otherwise
func (o *ShardKeyMostCommonValue) GetTruncated() bool {
	if o == nil || IsNil(o.Truncated) {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyMostCommonValue) GetTruncatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Truncated) {
		return nil, false
	}

	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *ShardKeyMostCommonValue) HasTruncated() bool {
	if o != nil && !IsNil(o.Truncated) {
		return true
	}

	return false
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *ShardKeyMostCommonValue) SetTruncated(v bool) {
	o.Truncated = &v
	o.NullFields = removeNullField(o.NullFields, "Truncated")
}

// SetTruncatedNil sets Truncated to an explicit JSON null when marshaled.
func (o *ShardKeyMostCommonValue) SetTruncatedNil() {
	o.Truncated = nil
	o.NullFields = addNullField(o.NullFields, "Truncated")
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *ShardKeyMostCommonValue) GetValue() any {
	if o == nil || IsNil(o.Value) {
		var ret any
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyMostCommonValue) GetValueOk() (any, bool) {
	if o == nil || IsNil(o.Value) {
		var ret any
		return ret, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ShardKeyMostCommonValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given any and assigns it to the Value field.
func (o *ShardKeyMostCommonValue) SetValue(v any) {
	o.Value = v
	o.NullFields = removeNullField(o.NullFields, "Value")
}

// SetValueNil sets Value to an explicit JSON null when marshaled.
func (o *ShardKeyMostCommonValue) SetValueNil() {
	o.Value = nil
	o.NullFields = addNullField(o.NullFields, "Value")
}
