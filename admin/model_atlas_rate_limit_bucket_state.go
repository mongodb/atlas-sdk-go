// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AtlasRateLimitBucketState Configuration and current state of a single rate limit token bucket.
type AtlasRateLimitBucketState struct {
	// The capacity of the bucket.
	// Read only field.
	Capacity *int64 `json:"capacity,omitempty"`
	// The name of the bucket.
	// Read only field.
	Name *string `json:"name,omitempty"`
	// The remaining tokens of the bucket.
	// Read only field.
	Remaining *int64 `json:"remaining,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AtlasRateLimitBucketState) MarshalJSON() ([]byte, error) {
	type noMethod AtlasRateLimitBucketState
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAtlasRateLimitBucketState instantiates a new AtlasRateLimitBucketState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtlasRateLimitBucketState() *AtlasRateLimitBucketState {
	this := AtlasRateLimitBucketState{}
	return &this
}

// NewAtlasRateLimitBucketStateWithDefaults instantiates a new AtlasRateLimitBucketState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtlasRateLimitBucketStateWithDefaults() *AtlasRateLimitBucketState {
	this := AtlasRateLimitBucketState{}
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise
func (o *AtlasRateLimitBucketState) GetCapacity() int64 {
	if o == nil || IsNil(o.Capacity) {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasRateLimitBucketState) GetCapacityOk() (*int64, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}

	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *AtlasRateLimitBucketState) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *AtlasRateLimitBucketState) SetCapacity(v int64) {
	o.Capacity = &v
	o.NullFields = removeNullField(o.NullFields, "Capacity")
}

// SetCapacityNil sets Capacity to an explicit JSON null when marshaled.
func (o *AtlasRateLimitBucketState) SetCapacityNil() {
	o.Capacity = nil
	o.NullFields = addNullField(o.NullFields, "Capacity")
}

// GetName returns the Name field value if set, zero value otherwise
func (o *AtlasRateLimitBucketState) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasRateLimitBucketState) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AtlasRateLimitBucketState) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AtlasRateLimitBucketState) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *AtlasRateLimitBucketState) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}

// GetRemaining returns the Remaining field value if set, zero value otherwise
func (o *AtlasRateLimitBucketState) GetRemaining() int64 {
	if o == nil || IsNil(o.Remaining) {
		var ret int64
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasRateLimitBucketState) GetRemainingOk() (*int64, bool) {
	if o == nil || IsNil(o.Remaining) {
		return nil, false
	}

	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *AtlasRateLimitBucketState) HasRemaining() bool {
	if o != nil && !IsNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int64 and assigns it to the Remaining field.
func (o *AtlasRateLimitBucketState) SetRemaining(v int64) {
	o.Remaining = &v
	o.NullFields = removeNullField(o.NullFields, "Remaining")
}

// SetRemainingNil sets Remaining to an explicit JSON null when marshaled.
func (o *AtlasRateLimitBucketState) SetRemainingNil() {
	o.Remaining = nil
	o.NullFields = addNullField(o.NullFields, "Remaining")
}
