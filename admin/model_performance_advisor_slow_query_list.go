// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PerformanceAdvisorSlowQueryList struct for PerformanceAdvisorSlowQueryList
type PerformanceAdvisorSlowQueryList struct {
	// List of operations that the Performance Advisor detected that took longer to execute than a specified threshold.
	// Read only field.
	SlowQueries *[]PerformanceAdvisorSlowQuery `json:"slowQueries,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PerformanceAdvisorSlowQueryList) MarshalJSON() ([]byte, error) {
	type noMethod PerformanceAdvisorSlowQueryList
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPerformanceAdvisorSlowQueryList instantiates a new PerformanceAdvisorSlowQueryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorSlowQueryList() *PerformanceAdvisorSlowQueryList {
	this := PerformanceAdvisorSlowQueryList{}
	return &this
}

// NewPerformanceAdvisorSlowQueryListWithDefaults instantiates a new PerformanceAdvisorSlowQueryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorSlowQueryListWithDefaults() *PerformanceAdvisorSlowQueryList {
	this := PerformanceAdvisorSlowQueryList{}
	return &this
}

// GetSlowQueries returns the SlowQueries field value if set, zero value otherwise
func (o *PerformanceAdvisorSlowQueryList) GetSlowQueries() []PerformanceAdvisorSlowQuery {
	if o == nil || IsNil(o.SlowQueries) {
		var ret []PerformanceAdvisorSlowQuery
		return ret
	}
	return *o.SlowQueries
}

// GetSlowQueriesOk returns a tuple with the SlowQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQueryList) GetSlowQueriesOk() (*[]PerformanceAdvisorSlowQuery, bool) {
	if o == nil || IsNil(o.SlowQueries) {
		return nil, false
	}

	return o.SlowQueries, true
}

// HasSlowQueries returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQueryList) HasSlowQueries() bool {
	if o != nil && !IsNil(o.SlowQueries) {
		return true
	}

	return false
}

// SetSlowQueries gets a reference to the given []PerformanceAdvisorSlowQuery and assigns it to the SlowQueries field.
func (o *PerformanceAdvisorSlowQueryList) SetSlowQueries(v []PerformanceAdvisorSlowQuery) {
	o.SlowQueries = &v
	o.NullFields = removeNullField(o.NullFields, "SlowQueries")
}

// SetSlowQueriesNil sets SlowQueries to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorSlowQueryList) SetSlowQueriesNil() {
	o.SlowQueries = nil
	o.NullFields = addNullField(o.NullFields, "SlowQueries")
}
