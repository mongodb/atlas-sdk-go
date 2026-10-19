// Code based on the AtlasAPI V2 OpenAPI file

package admin

// QueryStatsSummary A summary of execution statistics for a given query shape.
type QueryStatsSummary struct {
	// Average time in microseconds from the beginning of query processing to the first server response, per execution of queries with the given query shape. Calculated as `totalTimeToResponseMicros` divided by `execCount` over the requested time range. As a duration it can be longer or shorter than `avgWorkingMillis` (reported in milliseconds, so convert before comparing): longer when executions spent time blocked on locks or flow control rather than working, shorter when they kept working after the first batch to return later `getMore` batches.
	AvgTimeToResponseMicros *float64 `json:"avgTimeToResponseMicros,omitempty"`
	// Average active execution time in milliseconds per execution of queries with the given query shape, calculated as `totalWorkingMillis` divided by `execCount` over the requested time range. Excludes intentional pauses such as time waiting on locks or flow control.
	AvgWorkingMillis *float64 `json:"avgWorkingMillis,omitempty"`
	// The number of bytes read by the given query shape from the disk to the cache.
	BytesRead *float64 `json:"bytesRead,omitempty"`
	// The MongoDB command issued for this query shape. The insert, update, and delete commands appear only for clusters running MongoDB 9.0 or later.
	Command *string `json:"command,omitempty"`
	// Total CPU time in nanoseconds consumed by queries with the given query shape. Available for MDB 8.2 and higher.
	CpuTime *float64 `json:"cpuTime,omitempty"`
	// Total number of documents examined by queries with the given query shape.
	DocsExamined *float64 `json:"docsExamined,omitempty"`
	// Ratio of documents examined to documents returned by queries with the given query shape.
	DocsExaminedRatio *float64 `json:"docsExaminedRatio,omitempty"`
	// Total number of documents returned by queries with the given query shape.
	DocsReturned *float64 `json:"docsReturned,omitempty"`
	// Total number of times that queries with the given query shape have been executed.
	ExecCount *float64 `json:"execCount,omitempty"`
	// Total number of index keys deleted by queries with the given query shape. Available for MongoDB 9.0+ write commands.
	KeysDeleted *float64 `json:"keysDeleted,omitempty"`
	// Total number of in-bounds and out-of-bounds index keys examined by queries with the given query shape.
	KeysExamined *float64 `json:"keysExamined,omitempty"`
	// Ratio of in-bounds and out-of-bounds index keys examined to indexes containing documents returned by queries with the given query shape.
	KeysExaminedRatio *float64 `json:"keysExaminedRatio,omitempty"`
	// Total number of index keys inserted by queries with the given query shape. Available for MongoDB 9.0+ write commands.
	KeysInserted *float64 `json:"keysInserted,omitempty"`
	// Execution runtime in microseconds for the most recent query with the given query shape.
	LastExecMicros *float64 `json:"lastExecMicros,omitempty"`
	// Total number of documents deleted by queries with the given query shape. Available for MongoDB 9.0+ write commands.
	NDeleted *float64 `json:"nDeleted,omitempty"`
	// Total number of documents inserted by queries with the given query shape. Available for MongoDB 9.0+ write commands.
	NInserted *float64 `json:"nInserted,omitempty"`
	// Total number of documents matched by queries with the given query shape. Available for MongoDB 9.0+ write commands.
	NMatched *float64 `json:"nMatched,omitempty"`
	// Total number of documents modified by queries with the given query shape. Available for MongoDB 9.0+ write commands.
	NModified *float64 `json:"nModified,omitempty"`
	// Total number of documents upserted by queries with the given query shape. Available for MongoDB 9.0+ write commands.
	NUpserted *float64 `json:"nUpserted,omitempty"`
	// Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as `<database>.<collection>`.
	Namespace *string `json:"namespace,omitempty"`
	// The 50th percentile value of execution time in microseconds. This field is deprecated as the values it reports may be inaccurate. It will be removed in a future release.
	// Deprecated
	P50ExecMicros *float64 `json:"p50ExecMicros,omitempty"`
	// The 90th percentile value of execution time in microseconds. This field is deprecated as the values it reports may be inaccurate. It will be removed in a future release.
	// Deprecated
	P90ExecMicros *float64 `json:"p90ExecMicros,omitempty"`
	// The 99th percentile value of execution time in microseconds. This field is deprecated as the values it reports may be inaccurate. It will be removed in a future release.
	// Deprecated
	P99ExecMicros *float64 `json:"p99ExecMicros,omitempty"`
	// A query shape is a set of specifications that group similar queries together. Specifications can include filters, sorts, projections, aggregation pipeline stages, a namespace, and others. Queries that have similar specifications have the same query shape.
	QueryShape *string `json:"queryShape,omitempty"`
	// A hexadecimal string that represents the hash of a MongoDB query shape.
	QueryShapeHash *string `json:"queryShapeHash,omitempty"`
	// Indicates whether this query shape represents a system-initiated query.
	SystemQuery *bool `json:"systemQuery,omitempty"`
	// Total time in microseconds from the beginning of query processing to the first server response, summed across executions of queries with the given query shape. Includes time spent waiting on locks or flow control. Stops at the first batch of results, so time spent serving later `getMore` batches is not counted.
	TotalTimeToResponseMicros *float64 `json:"totalTimeToResponseMicros,omitempty"`
	// Total time in milliseconds that queries with the given query shape spent actively executing, summed across executions, including time spent processing `getMore` requests. Excludes intentional pauses such as time waiting on locks or flow control, and time spent waiting for the client.
	TotalWorkingMillis *float64 `json:"totalWorkingMillis,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *QueryStatsSummary) MarshalJSON() ([]byte, error) {
	type noMethod QueryStatsSummary
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewQueryStatsSummary instantiates a new QueryStatsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryStatsSummary() *QueryStatsSummary {
	this := QueryStatsSummary{}
	return &this
}

// NewQueryStatsSummaryWithDefaults instantiates a new QueryStatsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryStatsSummaryWithDefaults() *QueryStatsSummary {
	this := QueryStatsSummary{}
	return &this
}

// GetAvgTimeToResponseMicros returns the AvgTimeToResponseMicros field value if set, zero value otherwise
func (o *QueryStatsSummary) GetAvgTimeToResponseMicros() float64 {
	if o == nil || IsNil(o.AvgTimeToResponseMicros) {
		var ret float64
		return ret
	}
	return *o.AvgTimeToResponseMicros
}

// GetAvgTimeToResponseMicrosOk returns a tuple with the AvgTimeToResponseMicros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetAvgTimeToResponseMicrosOk() (*float64, bool) {
	if o == nil || IsNil(o.AvgTimeToResponseMicros) {
		return nil, false
	}

	return o.AvgTimeToResponseMicros, true
}

// HasAvgTimeToResponseMicros returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasAvgTimeToResponseMicros() bool {
	if o != nil && !IsNil(o.AvgTimeToResponseMicros) {
		return true
	}

	return false
}

// SetAvgTimeToResponseMicros gets a reference to the given float64 and assigns it to the AvgTimeToResponseMicros field.
func (o *QueryStatsSummary) SetAvgTimeToResponseMicros(v float64) {
	o.AvgTimeToResponseMicros = &v
	o.NullFields = removeNullField(o.NullFields, "AvgTimeToResponseMicros")
}

// SetAvgTimeToResponseMicrosNil sets AvgTimeToResponseMicros to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetAvgTimeToResponseMicrosNil() {
	o.AvgTimeToResponseMicros = nil
	o.NullFields = addNullField(o.NullFields, "AvgTimeToResponseMicros")
}

// GetAvgWorkingMillis returns the AvgWorkingMillis field value if set, zero value otherwise
func (o *QueryStatsSummary) GetAvgWorkingMillis() float64 {
	if o == nil || IsNil(o.AvgWorkingMillis) {
		var ret float64
		return ret
	}
	return *o.AvgWorkingMillis
}

// GetAvgWorkingMillisOk returns a tuple with the AvgWorkingMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetAvgWorkingMillisOk() (*float64, bool) {
	if o == nil || IsNil(o.AvgWorkingMillis) {
		return nil, false
	}

	return o.AvgWorkingMillis, true
}

// HasAvgWorkingMillis returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasAvgWorkingMillis() bool {
	if o != nil && !IsNil(o.AvgWorkingMillis) {
		return true
	}

	return false
}

// SetAvgWorkingMillis gets a reference to the given float64 and assigns it to the AvgWorkingMillis field.
func (o *QueryStatsSummary) SetAvgWorkingMillis(v float64) {
	o.AvgWorkingMillis = &v
	o.NullFields = removeNullField(o.NullFields, "AvgWorkingMillis")
}

// SetAvgWorkingMillisNil sets AvgWorkingMillis to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetAvgWorkingMillisNil() {
	o.AvgWorkingMillis = nil
	o.NullFields = addNullField(o.NullFields, "AvgWorkingMillis")
}

// GetBytesRead returns the BytesRead field value if set, zero value otherwise
func (o *QueryStatsSummary) GetBytesRead() float64 {
	if o == nil || IsNil(o.BytesRead) {
		var ret float64
		return ret
	}
	return *o.BytesRead
}

// GetBytesReadOk returns a tuple with the BytesRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetBytesReadOk() (*float64, bool) {
	if o == nil || IsNil(o.BytesRead) {
		return nil, false
	}

	return o.BytesRead, true
}

// HasBytesRead returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasBytesRead() bool {
	if o != nil && !IsNil(o.BytesRead) {
		return true
	}

	return false
}

// SetBytesRead gets a reference to the given float64 and assigns it to the BytesRead field.
func (o *QueryStatsSummary) SetBytesRead(v float64) {
	o.BytesRead = &v
	o.NullFields = removeNullField(o.NullFields, "BytesRead")
}

// SetBytesReadNil sets BytesRead to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetBytesReadNil() {
	o.BytesRead = nil
	o.NullFields = addNullField(o.NullFields, "BytesRead")
}

// GetCommand returns the Command field value if set, zero value otherwise
func (o *QueryStatsSummary) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}

	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *QueryStatsSummary) SetCommand(v string) {
	o.Command = &v
	o.NullFields = removeNullField(o.NullFields, "Command")
}

// SetCommandNil sets Command to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetCommandNil() {
	o.Command = nil
	o.NullFields = addNullField(o.NullFields, "Command")
}

// GetCpuTime returns the CpuTime field value if set, zero value otherwise
func (o *QueryStatsSummary) GetCpuTime() float64 {
	if o == nil || IsNil(o.CpuTime) {
		var ret float64
		return ret
	}
	return *o.CpuTime
}

// GetCpuTimeOk returns a tuple with the CpuTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetCpuTimeOk() (*float64, bool) {
	if o == nil || IsNil(o.CpuTime) {
		return nil, false
	}

	return o.CpuTime, true
}

// HasCpuTime returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasCpuTime() bool {
	if o != nil && !IsNil(o.CpuTime) {
		return true
	}

	return false
}

// SetCpuTime gets a reference to the given float64 and assigns it to the CpuTime field.
func (o *QueryStatsSummary) SetCpuTime(v float64) {
	o.CpuTime = &v
	o.NullFields = removeNullField(o.NullFields, "CpuTime")
}

// SetCpuTimeNil sets CpuTime to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetCpuTimeNil() {
	o.CpuTime = nil
	o.NullFields = addNullField(o.NullFields, "CpuTime")
}

// GetDocsExamined returns the DocsExamined field value if set, zero value otherwise
func (o *QueryStatsSummary) GetDocsExamined() float64 {
	if o == nil || IsNil(o.DocsExamined) {
		var ret float64
		return ret
	}
	return *o.DocsExamined
}

// GetDocsExaminedOk returns a tuple with the DocsExamined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetDocsExaminedOk() (*float64, bool) {
	if o == nil || IsNil(o.DocsExamined) {
		return nil, false
	}

	return o.DocsExamined, true
}

// HasDocsExamined returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasDocsExamined() bool {
	if o != nil && !IsNil(o.DocsExamined) {
		return true
	}

	return false
}

// SetDocsExamined gets a reference to the given float64 and assigns it to the DocsExamined field.
func (o *QueryStatsSummary) SetDocsExamined(v float64) {
	o.DocsExamined = &v
	o.NullFields = removeNullField(o.NullFields, "DocsExamined")
}

// SetDocsExaminedNil sets DocsExamined to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetDocsExaminedNil() {
	o.DocsExamined = nil
	o.NullFields = addNullField(o.NullFields, "DocsExamined")
}

// GetDocsExaminedRatio returns the DocsExaminedRatio field value if set, zero value otherwise
func (o *QueryStatsSummary) GetDocsExaminedRatio() float64 {
	if o == nil || IsNil(o.DocsExaminedRatio) {
		var ret float64
		return ret
	}
	return *o.DocsExaminedRatio
}

// GetDocsExaminedRatioOk returns a tuple with the DocsExaminedRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetDocsExaminedRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.DocsExaminedRatio) {
		return nil, false
	}

	return o.DocsExaminedRatio, true
}

// HasDocsExaminedRatio returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasDocsExaminedRatio() bool {
	if o != nil && !IsNil(o.DocsExaminedRatio) {
		return true
	}

	return false
}

// SetDocsExaminedRatio gets a reference to the given float64 and assigns it to the DocsExaminedRatio field.
func (o *QueryStatsSummary) SetDocsExaminedRatio(v float64) {
	o.DocsExaminedRatio = &v
	o.NullFields = removeNullField(o.NullFields, "DocsExaminedRatio")
}

// SetDocsExaminedRatioNil sets DocsExaminedRatio to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetDocsExaminedRatioNil() {
	o.DocsExaminedRatio = nil
	o.NullFields = addNullField(o.NullFields, "DocsExaminedRatio")
}

// GetDocsReturned returns the DocsReturned field value if set, zero value otherwise
func (o *QueryStatsSummary) GetDocsReturned() float64 {
	if o == nil || IsNil(o.DocsReturned) {
		var ret float64
		return ret
	}
	return *o.DocsReturned
}

// GetDocsReturnedOk returns a tuple with the DocsReturned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetDocsReturnedOk() (*float64, bool) {
	if o == nil || IsNil(o.DocsReturned) {
		return nil, false
	}

	return o.DocsReturned, true
}

// HasDocsReturned returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasDocsReturned() bool {
	if o != nil && !IsNil(o.DocsReturned) {
		return true
	}

	return false
}

// SetDocsReturned gets a reference to the given float64 and assigns it to the DocsReturned field.
func (o *QueryStatsSummary) SetDocsReturned(v float64) {
	o.DocsReturned = &v
	o.NullFields = removeNullField(o.NullFields, "DocsReturned")
}

// SetDocsReturnedNil sets DocsReturned to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetDocsReturnedNil() {
	o.DocsReturned = nil
	o.NullFields = addNullField(o.NullFields, "DocsReturned")
}

// GetExecCount returns the ExecCount field value if set, zero value otherwise
func (o *QueryStatsSummary) GetExecCount() float64 {
	if o == nil || IsNil(o.ExecCount) {
		var ret float64
		return ret
	}
	return *o.ExecCount
}

// GetExecCountOk returns a tuple with the ExecCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetExecCountOk() (*float64, bool) {
	if o == nil || IsNil(o.ExecCount) {
		return nil, false
	}

	return o.ExecCount, true
}

// HasExecCount returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasExecCount() bool {
	if o != nil && !IsNil(o.ExecCount) {
		return true
	}

	return false
}

// SetExecCount gets a reference to the given float64 and assigns it to the ExecCount field.
func (o *QueryStatsSummary) SetExecCount(v float64) {
	o.ExecCount = &v
	o.NullFields = removeNullField(o.NullFields, "ExecCount")
}

// SetExecCountNil sets ExecCount to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetExecCountNil() {
	o.ExecCount = nil
	o.NullFields = addNullField(o.NullFields, "ExecCount")
}

// GetKeysDeleted returns the KeysDeleted field value if set, zero value otherwise
func (o *QueryStatsSummary) GetKeysDeleted() float64 {
	if o == nil || IsNil(o.KeysDeleted) {
		var ret float64
		return ret
	}
	return *o.KeysDeleted
}

// GetKeysDeletedOk returns a tuple with the KeysDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetKeysDeletedOk() (*float64, bool) {
	if o == nil || IsNil(o.KeysDeleted) {
		return nil, false
	}

	return o.KeysDeleted, true
}

// HasKeysDeleted returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasKeysDeleted() bool {
	if o != nil && !IsNil(o.KeysDeleted) {
		return true
	}

	return false
}

// SetKeysDeleted gets a reference to the given float64 and assigns it to the KeysDeleted field.
func (o *QueryStatsSummary) SetKeysDeleted(v float64) {
	o.KeysDeleted = &v
	o.NullFields = removeNullField(o.NullFields, "KeysDeleted")
}

// SetKeysDeletedNil sets KeysDeleted to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetKeysDeletedNil() {
	o.KeysDeleted = nil
	o.NullFields = addNullField(o.NullFields, "KeysDeleted")
}

// GetKeysExamined returns the KeysExamined field value if set, zero value otherwise
func (o *QueryStatsSummary) GetKeysExamined() float64 {
	if o == nil || IsNil(o.KeysExamined) {
		var ret float64
		return ret
	}
	return *o.KeysExamined
}

// GetKeysExaminedOk returns a tuple with the KeysExamined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetKeysExaminedOk() (*float64, bool) {
	if o == nil || IsNil(o.KeysExamined) {
		return nil, false
	}

	return o.KeysExamined, true
}

// HasKeysExamined returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasKeysExamined() bool {
	if o != nil && !IsNil(o.KeysExamined) {
		return true
	}

	return false
}

// SetKeysExamined gets a reference to the given float64 and assigns it to the KeysExamined field.
func (o *QueryStatsSummary) SetKeysExamined(v float64) {
	o.KeysExamined = &v
	o.NullFields = removeNullField(o.NullFields, "KeysExamined")
}

// SetKeysExaminedNil sets KeysExamined to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetKeysExaminedNil() {
	o.KeysExamined = nil
	o.NullFields = addNullField(o.NullFields, "KeysExamined")
}

// GetKeysExaminedRatio returns the KeysExaminedRatio field value if set, zero value otherwise
func (o *QueryStatsSummary) GetKeysExaminedRatio() float64 {
	if o == nil || IsNil(o.KeysExaminedRatio) {
		var ret float64
		return ret
	}
	return *o.KeysExaminedRatio
}

// GetKeysExaminedRatioOk returns a tuple with the KeysExaminedRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetKeysExaminedRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.KeysExaminedRatio) {
		return nil, false
	}

	return o.KeysExaminedRatio, true
}

// HasKeysExaminedRatio returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasKeysExaminedRatio() bool {
	if o != nil && !IsNil(o.KeysExaminedRatio) {
		return true
	}

	return false
}

// SetKeysExaminedRatio gets a reference to the given float64 and assigns it to the KeysExaminedRatio field.
func (o *QueryStatsSummary) SetKeysExaminedRatio(v float64) {
	o.KeysExaminedRatio = &v
	o.NullFields = removeNullField(o.NullFields, "KeysExaminedRatio")
}

// SetKeysExaminedRatioNil sets KeysExaminedRatio to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetKeysExaminedRatioNil() {
	o.KeysExaminedRatio = nil
	o.NullFields = addNullField(o.NullFields, "KeysExaminedRatio")
}

// GetKeysInserted returns the KeysInserted field value if set, zero value otherwise
func (o *QueryStatsSummary) GetKeysInserted() float64 {
	if o == nil || IsNil(o.KeysInserted) {
		var ret float64
		return ret
	}
	return *o.KeysInserted
}

// GetKeysInsertedOk returns a tuple with the KeysInserted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetKeysInsertedOk() (*float64, bool) {
	if o == nil || IsNil(o.KeysInserted) {
		return nil, false
	}

	return o.KeysInserted, true
}

// HasKeysInserted returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasKeysInserted() bool {
	if o != nil && !IsNil(o.KeysInserted) {
		return true
	}

	return false
}

// SetKeysInserted gets a reference to the given float64 and assigns it to the KeysInserted field.
func (o *QueryStatsSummary) SetKeysInserted(v float64) {
	o.KeysInserted = &v
	o.NullFields = removeNullField(o.NullFields, "KeysInserted")
}

// SetKeysInsertedNil sets KeysInserted to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetKeysInsertedNil() {
	o.KeysInserted = nil
	o.NullFields = addNullField(o.NullFields, "KeysInserted")
}

// GetLastExecMicros returns the LastExecMicros field value if set, zero value otherwise
func (o *QueryStatsSummary) GetLastExecMicros() float64 {
	if o == nil || IsNil(o.LastExecMicros) {
		var ret float64
		return ret
	}
	return *o.LastExecMicros
}

// GetLastExecMicrosOk returns a tuple with the LastExecMicros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetLastExecMicrosOk() (*float64, bool) {
	if o == nil || IsNil(o.LastExecMicros) {
		return nil, false
	}

	return o.LastExecMicros, true
}

// HasLastExecMicros returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasLastExecMicros() bool {
	if o != nil && !IsNil(o.LastExecMicros) {
		return true
	}

	return false
}

// SetLastExecMicros gets a reference to the given float64 and assigns it to the LastExecMicros field.
func (o *QueryStatsSummary) SetLastExecMicros(v float64) {
	o.LastExecMicros = &v
	o.NullFields = removeNullField(o.NullFields, "LastExecMicros")
}

// SetLastExecMicrosNil sets LastExecMicros to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetLastExecMicrosNil() {
	o.LastExecMicros = nil
	o.NullFields = addNullField(o.NullFields, "LastExecMicros")
}

// GetNDeleted returns the NDeleted field value if set, zero value otherwise
func (o *QueryStatsSummary) GetNDeleted() float64 {
	if o == nil || IsNil(o.NDeleted) {
		var ret float64
		return ret
	}
	return *o.NDeleted
}

// GetNDeletedOk returns a tuple with the NDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetNDeletedOk() (*float64, bool) {
	if o == nil || IsNil(o.NDeleted) {
		return nil, false
	}

	return o.NDeleted, true
}

// HasNDeleted returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasNDeleted() bool {
	if o != nil && !IsNil(o.NDeleted) {
		return true
	}

	return false
}

// SetNDeleted gets a reference to the given float64 and assigns it to the NDeleted field.
func (o *QueryStatsSummary) SetNDeleted(v float64) {
	o.NDeleted = &v
	o.NullFields = removeNullField(o.NullFields, "NDeleted")
}

// SetNDeletedNil sets NDeleted to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetNDeletedNil() {
	o.NDeleted = nil
	o.NullFields = addNullField(o.NullFields, "NDeleted")
}

// GetNInserted returns the NInserted field value if set, zero value otherwise
func (o *QueryStatsSummary) GetNInserted() float64 {
	if o == nil || IsNil(o.NInserted) {
		var ret float64
		return ret
	}
	return *o.NInserted
}

// GetNInsertedOk returns a tuple with the NInserted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetNInsertedOk() (*float64, bool) {
	if o == nil || IsNil(o.NInserted) {
		return nil, false
	}

	return o.NInserted, true
}

// HasNInserted returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasNInserted() bool {
	if o != nil && !IsNil(o.NInserted) {
		return true
	}

	return false
}

// SetNInserted gets a reference to the given float64 and assigns it to the NInserted field.
func (o *QueryStatsSummary) SetNInserted(v float64) {
	o.NInserted = &v
	o.NullFields = removeNullField(o.NullFields, "NInserted")
}

// SetNInsertedNil sets NInserted to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetNInsertedNil() {
	o.NInserted = nil
	o.NullFields = addNullField(o.NullFields, "NInserted")
}

// GetNMatched returns the NMatched field value if set, zero value otherwise
func (o *QueryStatsSummary) GetNMatched() float64 {
	if o == nil || IsNil(o.NMatched) {
		var ret float64
		return ret
	}
	return *o.NMatched
}

// GetNMatchedOk returns a tuple with the NMatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetNMatchedOk() (*float64, bool) {
	if o == nil || IsNil(o.NMatched) {
		return nil, false
	}

	return o.NMatched, true
}

// HasNMatched returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasNMatched() bool {
	if o != nil && !IsNil(o.NMatched) {
		return true
	}

	return false
}

// SetNMatched gets a reference to the given float64 and assigns it to the NMatched field.
func (o *QueryStatsSummary) SetNMatched(v float64) {
	o.NMatched = &v
	o.NullFields = removeNullField(o.NullFields, "NMatched")
}

// SetNMatchedNil sets NMatched to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetNMatchedNil() {
	o.NMatched = nil
	o.NullFields = addNullField(o.NullFields, "NMatched")
}

// GetNModified returns the NModified field value if set, zero value otherwise
func (o *QueryStatsSummary) GetNModified() float64 {
	if o == nil || IsNil(o.NModified) {
		var ret float64
		return ret
	}
	return *o.NModified
}

// GetNModifiedOk returns a tuple with the NModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetNModifiedOk() (*float64, bool) {
	if o == nil || IsNil(o.NModified) {
		return nil, false
	}

	return o.NModified, true
}

// HasNModified returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasNModified() bool {
	if o != nil && !IsNil(o.NModified) {
		return true
	}

	return false
}

// SetNModified gets a reference to the given float64 and assigns it to the NModified field.
func (o *QueryStatsSummary) SetNModified(v float64) {
	o.NModified = &v
	o.NullFields = removeNullField(o.NullFields, "NModified")
}

// SetNModifiedNil sets NModified to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetNModifiedNil() {
	o.NModified = nil
	o.NullFields = addNullField(o.NullFields, "NModified")
}

// GetNUpserted returns the NUpserted field value if set, zero value otherwise
func (o *QueryStatsSummary) GetNUpserted() float64 {
	if o == nil || IsNil(o.NUpserted) {
		var ret float64
		return ret
	}
	return *o.NUpserted
}

// GetNUpsertedOk returns a tuple with the NUpserted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetNUpsertedOk() (*float64, bool) {
	if o == nil || IsNil(o.NUpserted) {
		return nil, false
	}

	return o.NUpserted, true
}

// HasNUpserted returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasNUpserted() bool {
	if o != nil && !IsNil(o.NUpserted) {
		return true
	}

	return false
}

// SetNUpserted gets a reference to the given float64 and assigns it to the NUpserted field.
func (o *QueryStatsSummary) SetNUpserted(v float64) {
	o.NUpserted = &v
	o.NullFields = removeNullField(o.NullFields, "NUpserted")
}

// SetNUpsertedNil sets NUpserted to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetNUpsertedNil() {
	o.NUpserted = nil
	o.NullFields = addNullField(o.NullFields, "NUpserted")
}

// GetNamespace returns the Namespace field value if set, zero value otherwise
func (o *QueryStatsSummary) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}

	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *QueryStatsSummary) SetNamespace(v string) {
	o.Namespace = &v
	o.NullFields = removeNullField(o.NullFields, "Namespace")
}

// SetNamespaceNil sets Namespace to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetNamespaceNil() {
	o.Namespace = nil
	o.NullFields = addNullField(o.NullFields, "Namespace")
}

// GetP50ExecMicros returns the P50ExecMicros field value if set, zero value otherwise
// Deprecated
func (o *QueryStatsSummary) GetP50ExecMicros() float64 {
	if o == nil || IsNil(o.P50ExecMicros) {
		var ret float64
		return ret
	}
	return *o.P50ExecMicros
}

// GetP50ExecMicrosOk returns a tuple with the P50ExecMicros field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryStatsSummary) GetP50ExecMicrosOk() (*float64, bool) {
	if o == nil || IsNil(o.P50ExecMicros) {
		return nil, false
	}

	return o.P50ExecMicros, true
}

// HasP50ExecMicros returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasP50ExecMicros() bool {
	if o != nil && !IsNil(o.P50ExecMicros) {
		return true
	}

	return false
}

// SetP50ExecMicros gets a reference to the given float64 and assigns it to the P50ExecMicros field.
// Deprecated
func (o *QueryStatsSummary) SetP50ExecMicros(v float64) {
	o.P50ExecMicros = &v
	o.NullFields = removeNullField(o.NullFields, "P50ExecMicros")
}

// SetP50ExecMicrosNil sets P50ExecMicros to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetP50ExecMicrosNil() {
	o.P50ExecMicros = nil
	o.NullFields = addNullField(o.NullFields, "P50ExecMicros")
}

// GetP90ExecMicros returns the P90ExecMicros field value if set, zero value otherwise
// Deprecated
func (o *QueryStatsSummary) GetP90ExecMicros() float64 {
	if o == nil || IsNil(o.P90ExecMicros) {
		var ret float64
		return ret
	}
	return *o.P90ExecMicros
}

// GetP90ExecMicrosOk returns a tuple with the P90ExecMicros field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryStatsSummary) GetP90ExecMicrosOk() (*float64, bool) {
	if o == nil || IsNil(o.P90ExecMicros) {
		return nil, false
	}

	return o.P90ExecMicros, true
}

// HasP90ExecMicros returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasP90ExecMicros() bool {
	if o != nil && !IsNil(o.P90ExecMicros) {
		return true
	}

	return false
}

// SetP90ExecMicros gets a reference to the given float64 and assigns it to the P90ExecMicros field.
// Deprecated
func (o *QueryStatsSummary) SetP90ExecMicros(v float64) {
	o.P90ExecMicros = &v
	o.NullFields = removeNullField(o.NullFields, "P90ExecMicros")
}

// SetP90ExecMicrosNil sets P90ExecMicros to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetP90ExecMicrosNil() {
	o.P90ExecMicros = nil
	o.NullFields = addNullField(o.NullFields, "P90ExecMicros")
}

// GetP99ExecMicros returns the P99ExecMicros field value if set, zero value otherwise
// Deprecated
func (o *QueryStatsSummary) GetP99ExecMicros() float64 {
	if o == nil || IsNil(o.P99ExecMicros) {
		var ret float64
		return ret
	}
	return *o.P99ExecMicros
}

// GetP99ExecMicrosOk returns a tuple with the P99ExecMicros field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryStatsSummary) GetP99ExecMicrosOk() (*float64, bool) {
	if o == nil || IsNil(o.P99ExecMicros) {
		return nil, false
	}

	return o.P99ExecMicros, true
}

// HasP99ExecMicros returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasP99ExecMicros() bool {
	if o != nil && !IsNil(o.P99ExecMicros) {
		return true
	}

	return false
}

// SetP99ExecMicros gets a reference to the given float64 and assigns it to the P99ExecMicros field.
// Deprecated
func (o *QueryStatsSummary) SetP99ExecMicros(v float64) {
	o.P99ExecMicros = &v
	o.NullFields = removeNullField(o.NullFields, "P99ExecMicros")
}

// SetP99ExecMicrosNil sets P99ExecMicros to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetP99ExecMicrosNil() {
	o.P99ExecMicros = nil
	o.NullFields = addNullField(o.NullFields, "P99ExecMicros")
}

// GetQueryShape returns the QueryShape field value if set, zero value otherwise
func (o *QueryStatsSummary) GetQueryShape() string {
	if o == nil || IsNil(o.QueryShape) {
		var ret string
		return ret
	}
	return *o.QueryShape
}

// GetQueryShapeOk returns a tuple with the QueryShape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetQueryShapeOk() (*string, bool) {
	if o == nil || IsNil(o.QueryShape) {
		return nil, false
	}

	return o.QueryShape, true
}

// HasQueryShape returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasQueryShape() bool {
	if o != nil && !IsNil(o.QueryShape) {
		return true
	}

	return false
}

// SetQueryShape gets a reference to the given string and assigns it to the QueryShape field.
func (o *QueryStatsSummary) SetQueryShape(v string) {
	o.QueryShape = &v
	o.NullFields = removeNullField(o.NullFields, "QueryShape")
}

// SetQueryShapeNil sets QueryShape to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetQueryShapeNil() {
	o.QueryShape = nil
	o.NullFields = addNullField(o.NullFields, "QueryShape")
}

// GetQueryShapeHash returns the QueryShapeHash field value if set, zero value otherwise
func (o *QueryStatsSummary) GetQueryShapeHash() string {
	if o == nil || IsNil(o.QueryShapeHash) {
		var ret string
		return ret
	}
	return *o.QueryShapeHash
}

// GetQueryShapeHashOk returns a tuple with the QueryShapeHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetQueryShapeHashOk() (*string, bool) {
	if o == nil || IsNil(o.QueryShapeHash) {
		return nil, false
	}

	return o.QueryShapeHash, true
}

// HasQueryShapeHash returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasQueryShapeHash() bool {
	if o != nil && !IsNil(o.QueryShapeHash) {
		return true
	}

	return false
}

// SetQueryShapeHash gets a reference to the given string and assigns it to the QueryShapeHash field.
func (o *QueryStatsSummary) SetQueryShapeHash(v string) {
	o.QueryShapeHash = &v
	o.NullFields = removeNullField(o.NullFields, "QueryShapeHash")
}

// SetQueryShapeHashNil sets QueryShapeHash to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetQueryShapeHashNil() {
	o.QueryShapeHash = nil
	o.NullFields = addNullField(o.NullFields, "QueryShapeHash")
}

// GetSystemQuery returns the SystemQuery field value if set, zero value otherwise
func (o *QueryStatsSummary) GetSystemQuery() bool {
	if o == nil || IsNil(o.SystemQuery) {
		var ret bool
		return ret
	}
	return *o.SystemQuery
}

// GetSystemQueryOk returns a tuple with the SystemQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetSystemQueryOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemQuery) {
		return nil, false
	}

	return o.SystemQuery, true
}

// HasSystemQuery returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasSystemQuery() bool {
	if o != nil && !IsNil(o.SystemQuery) {
		return true
	}

	return false
}

// SetSystemQuery gets a reference to the given bool and assigns it to the SystemQuery field.
func (o *QueryStatsSummary) SetSystemQuery(v bool) {
	o.SystemQuery = &v
	o.NullFields = removeNullField(o.NullFields, "SystemQuery")
}

// SetSystemQueryNil sets SystemQuery to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetSystemQueryNil() {
	o.SystemQuery = nil
	o.NullFields = addNullField(o.NullFields, "SystemQuery")
}

// GetTotalTimeToResponseMicros returns the TotalTimeToResponseMicros field value if set, zero value otherwise
func (o *QueryStatsSummary) GetTotalTimeToResponseMicros() float64 {
	if o == nil || IsNil(o.TotalTimeToResponseMicros) {
		var ret float64
		return ret
	}
	return *o.TotalTimeToResponseMicros
}

// GetTotalTimeToResponseMicrosOk returns a tuple with the TotalTimeToResponseMicros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetTotalTimeToResponseMicrosOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalTimeToResponseMicros) {
		return nil, false
	}

	return o.TotalTimeToResponseMicros, true
}

// HasTotalTimeToResponseMicros returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasTotalTimeToResponseMicros() bool {
	if o != nil && !IsNil(o.TotalTimeToResponseMicros) {
		return true
	}

	return false
}

// SetTotalTimeToResponseMicros gets a reference to the given float64 and assigns it to the TotalTimeToResponseMicros field.
func (o *QueryStatsSummary) SetTotalTimeToResponseMicros(v float64) {
	o.TotalTimeToResponseMicros = &v
	o.NullFields = removeNullField(o.NullFields, "TotalTimeToResponseMicros")
}

// SetTotalTimeToResponseMicrosNil sets TotalTimeToResponseMicros to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetTotalTimeToResponseMicrosNil() {
	o.TotalTimeToResponseMicros = nil
	o.NullFields = addNullField(o.NullFields, "TotalTimeToResponseMicros")
}

// GetTotalWorkingMillis returns the TotalWorkingMillis field value if set, zero value otherwise
func (o *QueryStatsSummary) GetTotalWorkingMillis() float64 {
	if o == nil || IsNil(o.TotalWorkingMillis) {
		var ret float64
		return ret
	}
	return *o.TotalWorkingMillis
}

// GetTotalWorkingMillisOk returns a tuple with the TotalWorkingMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummary) GetTotalWorkingMillisOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalWorkingMillis) {
		return nil, false
	}

	return o.TotalWorkingMillis, true
}

// HasTotalWorkingMillis returns a boolean if a field has been set.
func (o *QueryStatsSummary) HasTotalWorkingMillis() bool {
	if o != nil && !IsNil(o.TotalWorkingMillis) {
		return true
	}

	return false
}

// SetTotalWorkingMillis gets a reference to the given float64 and assigns it to the TotalWorkingMillis field.
func (o *QueryStatsSummary) SetTotalWorkingMillis(v float64) {
	o.TotalWorkingMillis = &v
	o.NullFields = removeNullField(o.NullFields, "TotalWorkingMillis")
}

// SetTotalWorkingMillisNil sets TotalWorkingMillis to an explicit JSON null when marshaled.
func (o *QueryStatsSummary) SetTotalWorkingMillisNil() {
	o.TotalWorkingMillis = nil
	o.NullFields = addNullField(o.NullFields, "TotalWorkingMillis")
}
