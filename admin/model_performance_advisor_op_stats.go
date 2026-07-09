// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PerformanceAdvisorOpStats Details that this resource returned about the specified query.
type PerformanceAdvisorOpStats struct {
	// Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds. This parameter relates to the **duration** query parameter.
	// Read only field.
	Ms *int64 `json:"ms,omitempty"`
	// Number of results that the query returns.
	// Read only field.
	NReturned *int64 `json:"nReturned,omitempty"`
	// Number of documents that the query read.
	// Read only field.
	NScanned *int64 `json:"nScanned,omitempty"`
	// Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. This parameter relates to the **since** query parameter.
	// Read only field.
	Ts *int64 `json:"ts,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PerformanceAdvisorOpStats) MarshalJSON() ([]byte, error) {
	type noMethod PerformanceAdvisorOpStats
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPerformanceAdvisorOpStats instantiates a new PerformanceAdvisorOpStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorOpStats() *PerformanceAdvisorOpStats {
	this := PerformanceAdvisorOpStats{}
	return &this
}

// NewPerformanceAdvisorOpStatsWithDefaults instantiates a new PerformanceAdvisorOpStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorOpStatsWithDefaults() *PerformanceAdvisorOpStats {
	this := PerformanceAdvisorOpStats{}
	return &this
}

// GetMs returns the Ms field value if set, zero value otherwise
func (o *PerformanceAdvisorOpStats) GetMs() int64 {
	if o == nil || IsNil(o.Ms) {
		var ret int64
		return ret
	}
	return *o.Ms
}

// GetMsOk returns a tuple with the Ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorOpStats) GetMsOk() (*int64, bool) {
	if o == nil || IsNil(o.Ms) {
		return nil, false
	}

	return o.Ms, true
}

// HasMs returns a boolean if a field has been set.
func (o *PerformanceAdvisorOpStats) HasMs() bool {
	if o != nil && !IsNil(o.Ms) {
		return true
	}

	return false
}

// SetMs gets a reference to the given int64 and assigns it to the Ms field.
func (o *PerformanceAdvisorOpStats) SetMs(v int64) {
	o.Ms = &v
}

// SetMsNil sets Ms to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorOpStats) SetMsNil() {
	o.Ms = nil
	o.NullFields = append(o.NullFields, "Ms")
}

// GetNReturned returns the NReturned field value if set, zero value otherwise
func (o *PerformanceAdvisorOpStats) GetNReturned() int64 {
	if o == nil || IsNil(o.NReturned) {
		var ret int64
		return ret
	}
	return *o.NReturned
}

// GetNReturnedOk returns a tuple with the NReturned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorOpStats) GetNReturnedOk() (*int64, bool) {
	if o == nil || IsNil(o.NReturned) {
		return nil, false
	}

	return o.NReturned, true
}

// HasNReturned returns a boolean if a field has been set.
func (o *PerformanceAdvisorOpStats) HasNReturned() bool {
	if o != nil && !IsNil(o.NReturned) {
		return true
	}

	return false
}

// SetNReturned gets a reference to the given int64 and assigns it to the NReturned field.
func (o *PerformanceAdvisorOpStats) SetNReturned(v int64) {
	o.NReturned = &v
}

// SetNReturnedNil sets NReturned to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorOpStats) SetNReturnedNil() {
	o.NReturned = nil
	o.NullFields = append(o.NullFields, "NReturned")
}

// GetNScanned returns the NScanned field value if set, zero value otherwise
func (o *PerformanceAdvisorOpStats) GetNScanned() int64 {
	if o == nil || IsNil(o.NScanned) {
		var ret int64
		return ret
	}
	return *o.NScanned
}

// GetNScannedOk returns a tuple with the NScanned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorOpStats) GetNScannedOk() (*int64, bool) {
	if o == nil || IsNil(o.NScanned) {
		return nil, false
	}

	return o.NScanned, true
}

// HasNScanned returns a boolean if a field has been set.
func (o *PerformanceAdvisorOpStats) HasNScanned() bool {
	if o != nil && !IsNil(o.NScanned) {
		return true
	}

	return false
}

// SetNScanned gets a reference to the given int64 and assigns it to the NScanned field.
func (o *PerformanceAdvisorOpStats) SetNScanned(v int64) {
	o.NScanned = &v
}

// SetNScannedNil sets NScanned to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorOpStats) SetNScannedNil() {
	o.NScanned = nil
	o.NullFields = append(o.NullFields, "NScanned")
}

// GetTs returns the Ts field value if set, zero value otherwise
func (o *PerformanceAdvisorOpStats) GetTs() int64 {
	if o == nil || IsNil(o.Ts) {
		var ret int64
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorOpStats) GetTsOk() (*int64, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}

	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *PerformanceAdvisorOpStats) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given int64 and assigns it to the Ts field.
func (o *PerformanceAdvisorOpStats) SetTs(v int64) {
	o.Ts = &v
}

// SetTsNil sets Ts to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorOpStats) SetTsNil() {
	o.Ts = nil
	o.NullFields = append(o.NullFields, "Ts")
}
