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
}
