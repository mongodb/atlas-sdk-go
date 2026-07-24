// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PerformanceAdvisorSlowQuery Details of one slow query that the Performance Advisor detected.
type PerformanceAdvisorSlowQuery struct {
	// Text of the MongoDB log related to this slow query.
	// Read only field.
	Line    *string                             `json:"line,omitempty"`
	Metrics *PerformanceAdvisorSlowQueryMetrics `json:"metrics,omitempty"`
	// Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as `<database>.<collection>`.
	// Read only field.
	Namespace *string `json:"namespace,omitempty"`
	// Operation type (read/write/command) associated with this slow query log.
	// Read only field.
	OpType *string `json:"opType,omitempty"`
	// Replica state associated with this slow query log.
	// Read only field.
	ReplicaState *string `json:"replicaState,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PerformanceAdvisorSlowQuery) MarshalJSON() ([]byte, error) {
	type noMethod PerformanceAdvisorSlowQuery
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPerformanceAdvisorSlowQuery instantiates a new PerformanceAdvisorSlowQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorSlowQuery() *PerformanceAdvisorSlowQuery {
	this := PerformanceAdvisorSlowQuery{}
	return &this
}

// NewPerformanceAdvisorSlowQueryWithDefaults instantiates a new PerformanceAdvisorSlowQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorSlowQueryWithDefaults() *PerformanceAdvisorSlowQuery {
	this := PerformanceAdvisorSlowQuery{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise
func (o *PerformanceAdvisorSlowQuery) GetLine() string {
	if o == nil || IsNil(o.Line) {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQuery) GetLineOk() (*string, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}

	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQuery) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *PerformanceAdvisorSlowQuery) SetLine(v string) {
	o.Line = &v
	o.NullFields = removeNullField(o.NullFields, "Line")
}

// SetLineNil sets Line to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorSlowQuery) SetLineNil() {
	o.Line = nil
	o.NullFields = addNullField(o.NullFields, "Line")
}

// GetMetrics returns the Metrics field value if set, zero value otherwise
func (o *PerformanceAdvisorSlowQuery) GetMetrics() PerformanceAdvisorSlowQueryMetrics {
	if o == nil || IsNil(o.Metrics) {
		var ret PerformanceAdvisorSlowQueryMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQuery) GetMetricsOk() (*PerformanceAdvisorSlowQueryMetrics, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}

	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQuery) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given PerformanceAdvisorSlowQueryMetrics and assigns it to the Metrics field.
func (o *PerformanceAdvisorSlowQuery) SetMetrics(v PerformanceAdvisorSlowQueryMetrics) {
	o.Metrics = &v
	o.NullFields = removeNullField(o.NullFields, "Metrics")
}

// SetMetricsNil sets Metrics to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorSlowQuery) SetMetricsNil() {
	o.Metrics = nil
	o.NullFields = addNullField(o.NullFields, "Metrics")
}

// GetNamespace returns the Namespace field value if set, zero value otherwise
func (o *PerformanceAdvisorSlowQuery) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQuery) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}

	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQuery) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *PerformanceAdvisorSlowQuery) SetNamespace(v string) {
	o.Namespace = &v
	o.NullFields = removeNullField(o.NullFields, "Namespace")
}

// SetNamespaceNil sets Namespace to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorSlowQuery) SetNamespaceNil() {
	o.Namespace = nil
	o.NullFields = addNullField(o.NullFields, "Namespace")
}

// GetOpType returns the OpType field value if set, zero value otherwise
func (o *PerformanceAdvisorSlowQuery) GetOpType() string {
	if o == nil || IsNil(o.OpType) {
		var ret string
		return ret
	}
	return *o.OpType
}

// GetOpTypeOk returns a tuple with the OpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQuery) GetOpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OpType) {
		return nil, false
	}

	return o.OpType, true
}

// HasOpType returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQuery) HasOpType() bool {
	if o != nil && !IsNil(o.OpType) {
		return true
	}

	return false
}

// SetOpType gets a reference to the given string and assigns it to the OpType field.
func (o *PerformanceAdvisorSlowQuery) SetOpType(v string) {
	o.OpType = &v
	o.NullFields = removeNullField(o.NullFields, "OpType")
}

// SetOpTypeNil sets OpType to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorSlowQuery) SetOpTypeNil() {
	o.OpType = nil
	o.NullFields = addNullField(o.NullFields, "OpType")
}

// GetReplicaState returns the ReplicaState field value if set, zero value otherwise
func (o *PerformanceAdvisorSlowQuery) GetReplicaState() string {
	if o == nil || IsNil(o.ReplicaState) {
		var ret string
		return ret
	}
	return *o.ReplicaState
}

// GetReplicaStateOk returns a tuple with the ReplicaState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQuery) GetReplicaStateOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaState) {
		return nil, false
	}

	return o.ReplicaState, true
}

// HasReplicaState returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQuery) HasReplicaState() bool {
	if o != nil && !IsNil(o.ReplicaState) {
		return true
	}

	return false
}

// SetReplicaState gets a reference to the given string and assigns it to the ReplicaState field.
func (o *PerformanceAdvisorSlowQuery) SetReplicaState(v string) {
	o.ReplicaState = &v
	o.NullFields = removeNullField(o.NullFields, "ReplicaState")
}

// SetReplicaStateNil sets ReplicaState to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorSlowQuery) SetReplicaStateNil() {
	o.ReplicaState = nil
	o.NullFields = addNullField(o.NullFields, "ReplicaState")
}
