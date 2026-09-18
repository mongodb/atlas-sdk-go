// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// QuerySamplingResponse Query sampling state for one namespace of one cluster. Sampled queries feed the shard key analysis for the namespace.
type QuerySamplingResponse struct {
	// Flag that indicates whether query sampling is observably running on the namespace. May still be `false` in the response to a successful create request; it becomes `true` asynchronously once the cluster applies the sampling configuration.
	// Read only field.
	Active bool `json:"active"`
	// Human-readable label that identifies the namespace (`<database>.<collection>`) on which queries are sampled.
	Namespace string `json:"namespace"`
	// Total size in bytes of the read operations sampled on the namespace.
	// Read only field.
	SampledReadsBytes *int64 `json:"sampledReadsBytes,omitempty"`
	// Number of read operations sampled on the namespace.
	// Read only field.
	SampledReadsCount *int64 `json:"sampledReadsCount,omitempty"`
	// Total size in bytes of the write operations sampled on the namespace.
	// Read only field.
	SampledWritesBytes *int64 `json:"sampledWritesBytes,omitempty"`
	// Number of write operations sampled on the namespace.
	// Read only field.
	SampledWritesCount *int64 `json:"sampledWritesCount,omitempty"`
	// Maximum number of queries sampled per second on the namespace. MongoDB Cloud sets this rate.
	// Read only field.
	SamplesPerSecond *float64 `json:"samplesPerSecond,omitempty"`
	// Date and time when query sampling started on the namespace. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	StartTime *time.Time `json:"startTime,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *QuerySamplingResponse) MarshalJSON() ([]byte, error) {
	type noMethod QuerySamplingResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewQuerySamplingResponse instantiates a new QuerySamplingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySamplingResponse(active bool, namespace string) *QuerySamplingResponse {
	this := QuerySamplingResponse{}
	this.Active = active
	this.Namespace = namespace
	return &this
}

// NewQuerySamplingResponseWithDefaults instantiates a new QuerySamplingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySamplingResponseWithDefaults() *QuerySamplingResponse {
	this := QuerySamplingResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *QuerySamplingResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *QuerySamplingResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *QuerySamplingResponse) SetActive(v bool) {
	o.Active = v
}

// GetNamespace returns the Namespace field value
func (o *QuerySamplingResponse) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *QuerySamplingResponse) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *QuerySamplingResponse) SetNamespace(v string) {
	o.Namespace = v
}

// GetSampledReadsBytes returns the SampledReadsBytes field value if set, zero value otherwise
func (o *QuerySamplingResponse) GetSampledReadsBytes() int64 {
	if o == nil || IsNil(o.SampledReadsBytes) {
		var ret int64
		return ret
	}
	return *o.SampledReadsBytes
}

// GetSampledReadsBytesOk returns a tuple with the SampledReadsBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySamplingResponse) GetSampledReadsBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.SampledReadsBytes) {
		return nil, false
	}

	return o.SampledReadsBytes, true
}

// HasSampledReadsBytes returns a boolean if a field has been set.
func (o *QuerySamplingResponse) HasSampledReadsBytes() bool {
	if o != nil && !IsNil(o.SampledReadsBytes) {
		return true
	}

	return false
}

// SetSampledReadsBytes gets a reference to the given int64 and assigns it to the SampledReadsBytes field.
func (o *QuerySamplingResponse) SetSampledReadsBytes(v int64) {
	o.SampledReadsBytes = &v
	o.NullFields = removeNullField(o.NullFields, "SampledReadsBytes")
}

// SetSampledReadsBytesNil sets SampledReadsBytes to an explicit JSON null when marshaled.
func (o *QuerySamplingResponse) SetSampledReadsBytesNil() {
	o.SampledReadsBytes = nil
	o.NullFields = addNullField(o.NullFields, "SampledReadsBytes")
}

// GetSampledReadsCount returns the SampledReadsCount field value if set, zero value otherwise
func (o *QuerySamplingResponse) GetSampledReadsCount() int64 {
	if o == nil || IsNil(o.SampledReadsCount) {
		var ret int64
		return ret
	}
	return *o.SampledReadsCount
}

// GetSampledReadsCountOk returns a tuple with the SampledReadsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySamplingResponse) GetSampledReadsCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SampledReadsCount) {
		return nil, false
	}

	return o.SampledReadsCount, true
}

// HasSampledReadsCount returns a boolean if a field has been set.
func (o *QuerySamplingResponse) HasSampledReadsCount() bool {
	if o != nil && !IsNil(o.SampledReadsCount) {
		return true
	}

	return false
}

// SetSampledReadsCount gets a reference to the given int64 and assigns it to the SampledReadsCount field.
func (o *QuerySamplingResponse) SetSampledReadsCount(v int64) {
	o.SampledReadsCount = &v
	o.NullFields = removeNullField(o.NullFields, "SampledReadsCount")
}

// SetSampledReadsCountNil sets SampledReadsCount to an explicit JSON null when marshaled.
func (o *QuerySamplingResponse) SetSampledReadsCountNil() {
	o.SampledReadsCount = nil
	o.NullFields = addNullField(o.NullFields, "SampledReadsCount")
}

// GetSampledWritesBytes returns the SampledWritesBytes field value if set, zero value otherwise
func (o *QuerySamplingResponse) GetSampledWritesBytes() int64 {
	if o == nil || IsNil(o.SampledWritesBytes) {
		var ret int64
		return ret
	}
	return *o.SampledWritesBytes
}

// GetSampledWritesBytesOk returns a tuple with the SampledWritesBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySamplingResponse) GetSampledWritesBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.SampledWritesBytes) {
		return nil, false
	}

	return o.SampledWritesBytes, true
}

// HasSampledWritesBytes returns a boolean if a field has been set.
func (o *QuerySamplingResponse) HasSampledWritesBytes() bool {
	if o != nil && !IsNil(o.SampledWritesBytes) {
		return true
	}

	return false
}

// SetSampledWritesBytes gets a reference to the given int64 and assigns it to the SampledWritesBytes field.
func (o *QuerySamplingResponse) SetSampledWritesBytes(v int64) {
	o.SampledWritesBytes = &v
	o.NullFields = removeNullField(o.NullFields, "SampledWritesBytes")
}

// SetSampledWritesBytesNil sets SampledWritesBytes to an explicit JSON null when marshaled.
func (o *QuerySamplingResponse) SetSampledWritesBytesNil() {
	o.SampledWritesBytes = nil
	o.NullFields = addNullField(o.NullFields, "SampledWritesBytes")
}

// GetSampledWritesCount returns the SampledWritesCount field value if set, zero value otherwise
func (o *QuerySamplingResponse) GetSampledWritesCount() int64 {
	if o == nil || IsNil(o.SampledWritesCount) {
		var ret int64
		return ret
	}
	return *o.SampledWritesCount
}

// GetSampledWritesCountOk returns a tuple with the SampledWritesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySamplingResponse) GetSampledWritesCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SampledWritesCount) {
		return nil, false
	}

	return o.SampledWritesCount, true
}

// HasSampledWritesCount returns a boolean if a field has been set.
func (o *QuerySamplingResponse) HasSampledWritesCount() bool {
	if o != nil && !IsNil(o.SampledWritesCount) {
		return true
	}

	return false
}

// SetSampledWritesCount gets a reference to the given int64 and assigns it to the SampledWritesCount field.
func (o *QuerySamplingResponse) SetSampledWritesCount(v int64) {
	o.SampledWritesCount = &v
	o.NullFields = removeNullField(o.NullFields, "SampledWritesCount")
}

// SetSampledWritesCountNil sets SampledWritesCount to an explicit JSON null when marshaled.
func (o *QuerySamplingResponse) SetSampledWritesCountNil() {
	o.SampledWritesCount = nil
	o.NullFields = addNullField(o.NullFields, "SampledWritesCount")
}

// GetSamplesPerSecond returns the SamplesPerSecond field value if set, zero value otherwise
func (o *QuerySamplingResponse) GetSamplesPerSecond() float64 {
	if o == nil || IsNil(o.SamplesPerSecond) {
		var ret float64
		return ret
	}
	return *o.SamplesPerSecond
}

// GetSamplesPerSecondOk returns a tuple with the SamplesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySamplingResponse) GetSamplesPerSecondOk() (*float64, bool) {
	if o == nil || IsNil(o.SamplesPerSecond) {
		return nil, false
	}

	return o.SamplesPerSecond, true
}

// HasSamplesPerSecond returns a boolean if a field has been set.
func (o *QuerySamplingResponse) HasSamplesPerSecond() bool {
	if o != nil && !IsNil(o.SamplesPerSecond) {
		return true
	}

	return false
}

// SetSamplesPerSecond gets a reference to the given float64 and assigns it to the SamplesPerSecond field.
func (o *QuerySamplingResponse) SetSamplesPerSecond(v float64) {
	o.SamplesPerSecond = &v
	o.NullFields = removeNullField(o.NullFields, "SamplesPerSecond")
}

// SetSamplesPerSecondNil sets SamplesPerSecond to an explicit JSON null when marshaled.
func (o *QuerySamplingResponse) SetSamplesPerSecondNil() {
	o.SamplesPerSecond = nil
	o.NullFields = addNullField(o.NullFields, "SamplesPerSecond")
}

// GetStartTime returns the StartTime field value if set, zero value otherwise
func (o *QuerySamplingResponse) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySamplingResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}

	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *QuerySamplingResponse) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *QuerySamplingResponse) SetStartTime(v time.Time) {
	o.StartTime = &v
	o.NullFields = removeNullField(o.NullFields, "StartTime")
}

// SetStartTimeNil sets StartTime to an explicit JSON null when marshaled.
func (o *QuerySamplingResponse) SetStartTimeNil() {
	o.StartTime = nil
	o.NullFields = addNullField(o.NullFields, "StartTime")
}
