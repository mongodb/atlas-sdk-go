// Code based on the AtlasAPI V2 OpenAPI file

package admin

// QueryStatsCollectionResponse How `queryStats` entries were collected on this cluster's processes. Reflects the collection mode of the cluster's processes regardless of any host or process type filters in the request. Omitted when the collection mode cannot be determined. Sample-based collection requires MongoDB 9.0 or later but may not be enabled on every such cluster; this reflects the mode in effect on the cluster.
type QueryStatsCollectionResponse struct {
	// Method used to collect `queryStats` entries on the cluster's processes.
	// Read only field.
	Mode string `json:"mode"`
	// Fraction of operations recorded. Dividing reported counts by this rate yields an unbiased estimate of the true totals.
	// Read only field.
	SampleRate *float64 `json:"sampleRate,omitempty"`
	// Maximum read queries recorded per second.
	// Read only field.
	RateLimitPerSecond *int `json:"rateLimitPerSecond,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *QueryStatsCollectionResponse) MarshalJSON() ([]byte, error) {
	type noMethod QueryStatsCollectionResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewQueryStatsCollectionResponse instantiates a new QueryStatsCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryStatsCollectionResponse(mode string) *QueryStatsCollectionResponse {
	this := QueryStatsCollectionResponse{}
	this.Mode = mode
	return &this
}

// NewQueryStatsCollectionResponseWithDefaults instantiates a new QueryStatsCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryStatsCollectionResponseWithDefaults() *QueryStatsCollectionResponse {
	this := QueryStatsCollectionResponse{}
	return &this
}

// GetMode returns the Mode field value
func (o *QueryStatsCollectionResponse) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *QueryStatsCollectionResponse) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *QueryStatsCollectionResponse) SetMode(v string) {
	o.Mode = v
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise
func (o *QueryStatsCollectionResponse) GetSampleRate() float64 {
	if o == nil || IsNil(o.SampleRate) {
		var ret float64
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsCollectionResponse) GetSampleRateOk() (*float64, bool) {
	if o == nil || IsNil(o.SampleRate) {
		return nil, false
	}

	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *QueryStatsCollectionResponse) HasSampleRate() bool {
	if o != nil && !IsNil(o.SampleRate) {
		return true
	}

	return false
}

// SetSampleRate gets a reference to the given float64 and assigns it to the SampleRate field.
func (o *QueryStatsCollectionResponse) SetSampleRate(v float64) {
	o.SampleRate = &v
	o.NullFields = removeNullField(o.NullFields, "SampleRate")
}

// SetSampleRateNil sets SampleRate to an explicit JSON null when marshaled.
func (o *QueryStatsCollectionResponse) SetSampleRateNil() {
	o.SampleRate = nil
	o.NullFields = addNullField(o.NullFields, "SampleRate")
}

// GetRateLimitPerSecond returns the RateLimitPerSecond field value if set, zero value otherwise
func (o *QueryStatsCollectionResponse) GetRateLimitPerSecond() int {
	if o == nil || IsNil(o.RateLimitPerSecond) {
		var ret int
		return ret
	}
	return *o.RateLimitPerSecond
}

// GetRateLimitPerSecondOk returns a tuple with the RateLimitPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsCollectionResponse) GetRateLimitPerSecondOk() (*int, bool) {
	if o == nil || IsNil(o.RateLimitPerSecond) {
		return nil, false
	}

	return o.RateLimitPerSecond, true
}

// HasRateLimitPerSecond returns a boolean if a field has been set.
func (o *QueryStatsCollectionResponse) HasRateLimitPerSecond() bool {
	if o != nil && !IsNil(o.RateLimitPerSecond) {
		return true
	}

	return false
}

// SetRateLimitPerSecond gets a reference to the given int and assigns it to the RateLimitPerSecond field.
func (o *QueryStatsCollectionResponse) SetRateLimitPerSecond(v int) {
	o.RateLimitPerSecond = &v
	o.NullFields = removeNullField(o.NullFields, "RateLimitPerSecond")
}

// SetRateLimitPerSecondNil sets RateLimitPerSecond to an explicit JSON null when marshaled.
func (o *QueryStatsCollectionResponse) SetRateLimitPerSecondNil() {
	o.RateLimitPerSecond = nil
	o.NullFields = addNullField(o.NullFields, "RateLimitPerSecond")
}
