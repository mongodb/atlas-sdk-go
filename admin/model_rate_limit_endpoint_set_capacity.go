// Code based on the AtlasAPI V2 OpenAPI file

package admin

// RateLimitEndpointSetCapacity The rate limit capacity for the endpoint set.
type RateLimitEndpointSetCapacity struct {
	// The default request capacity of the endpoint set. Returned if there is a capacity override set for the requested entity.
	DefaultValue *int64 `json:"defaultValue,omitempty"`
	// The applied request capacity of the endpoint set.
	Value *int64 `json:"value,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *RateLimitEndpointSetCapacity) MarshalJSON() ([]byte, error) {
	type noMethod RateLimitEndpointSetCapacity
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewRateLimitEndpointSetCapacity instantiates a new RateLimitEndpointSetCapacity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitEndpointSetCapacity() *RateLimitEndpointSetCapacity {
	this := RateLimitEndpointSetCapacity{}
	return &this
}

// NewRateLimitEndpointSetCapacityWithDefaults instantiates a new RateLimitEndpointSetCapacity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitEndpointSetCapacityWithDefaults() *RateLimitEndpointSetCapacity {
	this := RateLimitEndpointSetCapacity{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise
func (o *RateLimitEndpointSetCapacity) GetDefaultValue() int64 {
	if o == nil || IsNil(o.DefaultValue) {
		var ret int64
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitEndpointSetCapacity) GetDefaultValueOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}

	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *RateLimitEndpointSetCapacity) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given int64 and assigns it to the DefaultValue field.
func (o *RateLimitEndpointSetCapacity) SetDefaultValue(v int64) {
	o.DefaultValue = &v
}

// SetDefaultValueNil sets DefaultValue to an explicit JSON null when marshaled.
func (o *RateLimitEndpointSetCapacity) SetDefaultValueNil() {
	o.DefaultValue = nil
	o.NullFields = append(o.NullFields, "DefaultValue")
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *RateLimitEndpointSetCapacity) GetValue() int64 {
	if o == nil || IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitEndpointSetCapacity) GetValueOk() (*int64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RateLimitEndpointSetCapacity) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *RateLimitEndpointSetCapacity) SetValue(v int64) {
	o.Value = &v
}

// SetValueNil sets Value to an explicit JSON null when marshaled.
func (o *RateLimitEndpointSetCapacity) SetValueNil() {
	o.Value = nil
	o.NullFields = append(o.NullFields, "Value")
}
