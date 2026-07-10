// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AtlasRateLimitInspectionResponse Rate limit inspection response containing bucket states for debugging and monitoring.
type AtlasRateLimitInspectionResponse struct {
	// List of bucket states.
	// Read only field.
	Limits *[]AtlasRateLimitBucketState `json:"limits,omitempty"`
	// The scope type of the rate limit.
	// Read only field.
	Scope *string `json:"scope,omitempty"`
	// The scope id associated to the bucket.
	// Read only field.
	ScopeId *string `json:"scopeId,omitempty"`
}

// NewAtlasRateLimitInspectionResponse instantiates a new AtlasRateLimitInspectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtlasRateLimitInspectionResponse() *AtlasRateLimitInspectionResponse {
	this := AtlasRateLimitInspectionResponse{}
	return &this
}

// NewAtlasRateLimitInspectionResponseWithDefaults instantiates a new AtlasRateLimitInspectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtlasRateLimitInspectionResponseWithDefaults() *AtlasRateLimitInspectionResponse {
	this := AtlasRateLimitInspectionResponse{}
	return &this
}

// GetLimits returns the Limits field value if set, zero value otherwise
func (o *AtlasRateLimitInspectionResponse) GetLimits() []AtlasRateLimitBucketState {
	if o == nil || IsNil(o.Limits) {
		var ret []AtlasRateLimitBucketState
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasRateLimitInspectionResponse) GetLimitsOk() (*[]AtlasRateLimitBucketState, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}

	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *AtlasRateLimitInspectionResponse) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []AtlasRateLimitBucketState and assigns it to the Limits field.
func (o *AtlasRateLimitInspectionResponse) SetLimits(v []AtlasRateLimitBucketState) {
	o.Limits = &v
}

// GetScope returns the Scope field value if set, zero value otherwise
func (o *AtlasRateLimitInspectionResponse) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasRateLimitInspectionResponse) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}

	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AtlasRateLimitInspectionResponse) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AtlasRateLimitInspectionResponse) SetScope(v string) {
	o.Scope = &v
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise
func (o *AtlasRateLimitInspectionResponse) GetScopeId() string {
	if o == nil || IsNil(o.ScopeId) {
		var ret string
		return ret
	}
	return *o.ScopeId
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasRateLimitInspectionResponse) GetScopeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeId) {
		return nil, false
	}

	return o.ScopeId, true
}

// HasScopeId returns a boolean if a field has been set.
func (o *AtlasRateLimitInspectionResponse) HasScopeId() bool {
	if o != nil && !IsNil(o.ScopeId) {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given string and assigns it to the ScopeId field.
func (o *AtlasRateLimitInspectionResponse) SetScopeId(v string) {
	o.ScopeId = &v
}
