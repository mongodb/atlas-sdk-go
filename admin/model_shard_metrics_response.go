// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardMetricsResponse Metrics for one shard of a sharded cluster, covering its size, its load, and any alerts open on its nodes.
type ShardMetricsResponse struct {
	// Alerts open on this shard's nodes, sorted by creation date with the most recently created first. MongoDB Cloud returns this property when you request a single shard, and returns at most 100 alerts.
	// Read only field.
	Alerts          *[]ShardAlert         `json:"alerts,omitempty"`
	CacheThroughput *ShardCacheThroughput `json:"cacheThroughput,omitempty"`
	CpuUtilization  ShardCpuUtilization   `json:"cpuUtilization"`
	// Total size of the user data stored on this shard. The resource expresses this value in bytes.
	// Read only field.
	DataSizeBytes int64 `json:"dataSizeBytes"`
	// Average number of range deletion tasks queued across the shard's nodes, measured at the time of the request.
	// Read only field.
	RangeDeleterTasks *float64 `json:"rangeDeleterTasks,omitempty"`
	// Average latency of the read operations this shard served over the last hour, in microseconds.
	// Read only field.
	ReadLatencyMicroseconds *float64        `json:"readLatencyMicroseconds,omitempty"`
	ReadOperationsPerSecond ShardOperations `json:"readOperationsPerSecond"`
	// Name of the shard's replica set. This is the name that identifies the same shard on the processes, alerts, events, and backup resources.
	// Read only field.
	ReplicaSetName string `json:"replicaSetName"`
	// Number of sharded collections with data on this shard.
	// Read only field.
	ShardedCollections int `json:"shardedCollections"`
	// Conditions that MongoDB Cloud observed on this shard relative to the rest of the cluster. `HOTSPOT` means this shard's average CPU utilization is more than twice the average across all shards of the cluster. This is a cluster-scoped comparison, so its value depends on how loaded the other shards are and not only on this shard.
	// Read only field.
	Warnings                 []string        `json:"warnings"`
	WriteOperationsPerSecond ShardOperations `json:"writeOperationsPerSecond"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardMetricsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ShardMetricsResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardMetricsResponse instantiates a new ShardMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardMetricsResponse(cpuUtilization ShardCpuUtilization, dataSizeBytes int64, readOperationsPerSecond ShardOperations, replicaSetName string, shardedCollections int, warnings []string, writeOperationsPerSecond ShardOperations) *ShardMetricsResponse {
	this := ShardMetricsResponse{}
	this.CpuUtilization = cpuUtilization
	this.DataSizeBytes = dataSizeBytes
	this.ReadOperationsPerSecond = readOperationsPerSecond
	this.ReplicaSetName = replicaSetName
	this.ShardedCollections = shardedCollections
	this.Warnings = warnings
	this.WriteOperationsPerSecond = writeOperationsPerSecond
	return &this
}

// NewShardMetricsResponseWithDefaults instantiates a new ShardMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardMetricsResponseWithDefaults() *ShardMetricsResponse {
	this := ShardMetricsResponse{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise
func (o *ShardMetricsResponse) GetAlerts() []ShardAlert {
	if o == nil || IsNil(o.Alerts) {
		var ret []ShardAlert
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetAlertsOk() (*[]ShardAlert, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}

	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *ShardMetricsResponse) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []ShardAlert and assigns it to the Alerts field.
func (o *ShardMetricsResponse) SetAlerts(v []ShardAlert) {
	o.Alerts = &v
	o.NullFields = removeNullField(o.NullFields, "Alerts")
}

// SetAlertsNil sets Alerts to an explicit JSON null when marshaled.
func (o *ShardMetricsResponse) SetAlertsNil() {
	o.Alerts = nil
	o.NullFields = addNullField(o.NullFields, "Alerts")
}

// GetCacheThroughput returns the CacheThroughput field value if set, zero value otherwise
func (o *ShardMetricsResponse) GetCacheThroughput() ShardCacheThroughput {
	if o == nil || IsNil(o.CacheThroughput) {
		var ret ShardCacheThroughput
		return ret
	}
	return *o.CacheThroughput
}

// GetCacheThroughputOk returns a tuple with the CacheThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetCacheThroughputOk() (*ShardCacheThroughput, bool) {
	if o == nil || IsNil(o.CacheThroughput) {
		return nil, false
	}

	return o.CacheThroughput, true
}

// HasCacheThroughput returns a boolean if a field has been set.
func (o *ShardMetricsResponse) HasCacheThroughput() bool {
	if o != nil && !IsNil(o.CacheThroughput) {
		return true
	}

	return false
}

// SetCacheThroughput gets a reference to the given ShardCacheThroughput and assigns it to the CacheThroughput field.
func (o *ShardMetricsResponse) SetCacheThroughput(v ShardCacheThroughput) {
	o.CacheThroughput = &v
	o.NullFields = removeNullField(o.NullFields, "CacheThroughput")
}

// SetCacheThroughputNil sets CacheThroughput to an explicit JSON null when marshaled.
func (o *ShardMetricsResponse) SetCacheThroughputNil() {
	o.CacheThroughput = nil
	o.NullFields = addNullField(o.NullFields, "CacheThroughput")
}

// GetCpuUtilization returns the CpuUtilization field value
func (o *ShardMetricsResponse) GetCpuUtilization() ShardCpuUtilization {
	if o == nil {
		var ret ShardCpuUtilization
		return ret
	}

	return o.CpuUtilization
}

// GetCpuUtilizationOk returns a tuple with the CpuUtilization field value
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetCpuUtilizationOk() (*ShardCpuUtilization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuUtilization, true
}

// SetCpuUtilization sets field value
func (o *ShardMetricsResponse) SetCpuUtilization(v ShardCpuUtilization) {
	o.CpuUtilization = v
}

// GetDataSizeBytes returns the DataSizeBytes field value
func (o *ShardMetricsResponse) GetDataSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DataSizeBytes
}

// GetDataSizeBytesOk returns a tuple with the DataSizeBytes field value
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetDataSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSizeBytes, true
}

// SetDataSizeBytes sets field value
func (o *ShardMetricsResponse) SetDataSizeBytes(v int64) {
	o.DataSizeBytes = v
}

// GetRangeDeleterTasks returns the RangeDeleterTasks field value if set, zero value otherwise
func (o *ShardMetricsResponse) GetRangeDeleterTasks() float64 {
	if o == nil || IsNil(o.RangeDeleterTasks) {
		var ret float64
		return ret
	}
	return *o.RangeDeleterTasks
}

// GetRangeDeleterTasksOk returns a tuple with the RangeDeleterTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetRangeDeleterTasksOk() (*float64, bool) {
	if o == nil || IsNil(o.RangeDeleterTasks) {
		return nil, false
	}

	return o.RangeDeleterTasks, true
}

// HasRangeDeleterTasks returns a boolean if a field has been set.
func (o *ShardMetricsResponse) HasRangeDeleterTasks() bool {
	if o != nil && !IsNil(o.RangeDeleterTasks) {
		return true
	}

	return false
}

// SetRangeDeleterTasks gets a reference to the given float64 and assigns it to the RangeDeleterTasks field.
func (o *ShardMetricsResponse) SetRangeDeleterTasks(v float64) {
	o.RangeDeleterTasks = &v
	o.NullFields = removeNullField(o.NullFields, "RangeDeleterTasks")
}

// SetRangeDeleterTasksNil sets RangeDeleterTasks to an explicit JSON null when marshaled.
func (o *ShardMetricsResponse) SetRangeDeleterTasksNil() {
	o.RangeDeleterTasks = nil
	o.NullFields = addNullField(o.NullFields, "RangeDeleterTasks")
}

// GetReadLatencyMicroseconds returns the ReadLatencyMicroseconds field value if set, zero value otherwise
func (o *ShardMetricsResponse) GetReadLatencyMicroseconds() float64 {
	if o == nil || IsNil(o.ReadLatencyMicroseconds) {
		var ret float64
		return ret
	}
	return *o.ReadLatencyMicroseconds
}

// GetReadLatencyMicrosecondsOk returns a tuple with the ReadLatencyMicroseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetReadLatencyMicrosecondsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReadLatencyMicroseconds) {
		return nil, false
	}

	return o.ReadLatencyMicroseconds, true
}

// HasReadLatencyMicroseconds returns a boolean if a field has been set.
func (o *ShardMetricsResponse) HasReadLatencyMicroseconds() bool {
	if o != nil && !IsNil(o.ReadLatencyMicroseconds) {
		return true
	}

	return false
}

// SetReadLatencyMicroseconds gets a reference to the given float64 and assigns it to the ReadLatencyMicroseconds field.
func (o *ShardMetricsResponse) SetReadLatencyMicroseconds(v float64) {
	o.ReadLatencyMicroseconds = &v
	o.NullFields = removeNullField(o.NullFields, "ReadLatencyMicroseconds")
}

// SetReadLatencyMicrosecondsNil sets ReadLatencyMicroseconds to an explicit JSON null when marshaled.
func (o *ShardMetricsResponse) SetReadLatencyMicrosecondsNil() {
	o.ReadLatencyMicroseconds = nil
	o.NullFields = addNullField(o.NullFields, "ReadLatencyMicroseconds")
}

// GetReadOperationsPerSecond returns the ReadOperationsPerSecond field value
func (o *ShardMetricsResponse) GetReadOperationsPerSecond() ShardOperations {
	if o == nil {
		var ret ShardOperations
		return ret
	}

	return o.ReadOperationsPerSecond
}

// GetReadOperationsPerSecondOk returns a tuple with the ReadOperationsPerSecond field value
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetReadOperationsPerSecondOk() (*ShardOperations, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOperationsPerSecond, true
}

// SetReadOperationsPerSecond sets field value
func (o *ShardMetricsResponse) SetReadOperationsPerSecond(v ShardOperations) {
	o.ReadOperationsPerSecond = v
}

// GetReplicaSetName returns the ReplicaSetName field value
func (o *ShardMetricsResponse) GetReplicaSetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetReplicaSetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicaSetName, true
}

// SetReplicaSetName sets field value
func (o *ShardMetricsResponse) SetReplicaSetName(v string) {
	o.ReplicaSetName = v
}

// GetShardedCollections returns the ShardedCollections field value
func (o *ShardMetricsResponse) GetShardedCollections() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.ShardedCollections
}

// GetShardedCollectionsOk returns a tuple with the ShardedCollections field value
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetShardedCollectionsOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShardedCollections, true
}

// SetShardedCollections sets field value
func (o *ShardMetricsResponse) SetShardedCollections(v int) {
	o.ShardedCollections = v
}

// GetWarnings returns the Warnings field value
func (o *ShardMetricsResponse) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value
func (o *ShardMetricsResponse) SetWarnings(v []string) {
	o.Warnings = v
}

// GetWriteOperationsPerSecond returns the WriteOperationsPerSecond field value
func (o *ShardMetricsResponse) GetWriteOperationsPerSecond() ShardOperations {
	if o == nil {
		var ret ShardOperations
		return ret
	}

	return o.WriteOperationsPerSecond
}

// GetWriteOperationsPerSecondOk returns a tuple with the WriteOperationsPerSecond field value
// and a boolean to check if the value has been set.
func (o *ShardMetricsResponse) GetWriteOperationsPerSecondOk() (*ShardOperations, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteOperationsPerSecond, true
}

// SetWriteOperationsPerSecond sets field value
func (o *ShardMetricsResponse) SetWriteOperationsPerSecond(v ShardOperations) {
	o.WriteOperationsPerSecond = v
}
