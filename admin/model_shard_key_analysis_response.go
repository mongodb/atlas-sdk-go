// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// ShardKeyAnalysisResponse Completed analysis of one candidate shard key, describing how well the key would distribute the collection's documents and operations across shards.
type ShardKeyAnalysisResponse struct {
	// Average size of one document in the collection, in bytes.
	// Read only field.
	AvgDocSizeBytes *int64 `json:"avgDocSizeBytes,omitempty"`
	// Correlation between the candidate shard key and the order in which documents were inserted, from `-1` to `1`. Values near `1` or `-1` indicate a monotonically changing key, which directs all inserts to one shard. This parameter is absent when the shard key has no supporting index or the collection is clustered, in which case MongoDB Cloud does not compute it.
	// Read only field.
	CorrelationCoefficient *float64 `json:"correlationCoefficient,omitempty"`
	// Date and time when this shard key analysis was requested (ISO 8601 format in UTC). This is when the analysis was submitted, not when it completed; the completion time is the `updatedAt` of the operation that produced it.
	// Read only field.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time when MongoDB Cloud stops returning this shard key analysis (ISO 8601 format in UTC). Reads after this time return an HTTP 404 status.
	// Read only field.
	ExpiresAt time.Time `json:"expiresAt"`
	// Unique 24-hexadecimal digit string that identifies this shard key analysis. This is the same identifier as the `operationId` of the operation that produced it.
	// Read only field.
	Id string `json:"id"`
	// Flag that indicates whether the index supporting the candidate shard key enforces a uniqueness constraint.
	// Read only field.
	IsUnique *bool `json:"isUnique,omitempty"`
	// Whether the shard key value increases or decreases with the order in which documents were inserted. A monotonically changing shard key directs all inserts to one shard. `UNKNOWN` indicates that MongoDB Cloud could not determine monotonicity, which happens when the shard key has no supporting index or the collection is clustered.
	// Read only field.
	Monotonicity *string `json:"monotonicity,omitempty"`
	// Shard key values that occur most frequently in the sampled documents, with how often each occurs. MongoDB Cloud returns this parameter only to callers that can read the collection's data, because it carries values out of the collection rather than statistics about them.
	// Read only field.
	MostCommonValues *[]ShardKeyMostCommonValue `json:"mostCommonValues,omitempty"`
	// Combination of the database and collection that was analyzed, written as `<database>.<collection>`.
	Namespace string `json:"namespace"`
	// Guidance from MongoDB on how to interpret the metrics in this analysis. This parameter is absent when the analysis carries no guidance.
	// Read only field.
	Note *string `json:"note,omitempty"`
	// Number of distinct values that the candidate shard key takes across the sampled documents. A value close to `numDocsSampled` indicates high cardinality, which distributes documents well.
	// Read only field.
	NumDistinctValues *int64 `json:"numDistinctValues,omitempty"`
	// Number of documents that MongoDB Cloud sampled to compute these metrics.
	// Read only field.
	NumDocsSampled *int64 `json:"numDocsSampled,omitempty"`
	// Number of documents in the collection.
	// Read only field.
	NumDocsTotal *int64 `json:"numDocsTotal,omitempty"`
	// Number of orphaned documents in the collection. This parameter is absent for collections that have none.
	// Read only field.
	NumOrphanDocs *int64 `json:"numOrphanDocs,omitempty"`
	// Percentage of sampled reads that the candidate shard key would route to more than one shard but not to every shard.
	// Read only field.
	PercentageOfMultiShardReads *float64 `json:"percentageOfMultiShardReads,omitempty"`
	// Percentage of sampled writes that the candidate shard key would route to more than one shard but not to every shard.
	// Read only field.
	PercentageOfMultiShardWrites *float64 `json:"percentageOfMultiShardWrites,omitempty"`
	// Percentage of sampled reads that the candidate shard key would broadcast to every shard. Lower is better.
	// Read only field.
	PercentageOfScatterGatherReads *float64 `json:"percentageOfScatterGatherReads,omitempty"`
	// Percentage of sampled writes that the candidate shard key would broadcast to every shard. Lower is better.
	// Read only field.
	PercentageOfScatterGatherWrites *float64 `json:"percentageOfScatterGatherWrites,omitempty"`
	// Percentage of sampled reads that the candidate shard key would route to exactly one shard. Higher is better.
	// Read only field.
	PercentageOfSingleShardReads *float64 `json:"percentageOfSingleShardReads,omitempty"`
	// Percentage of sampled writes that the candidate shard key would route to exactly one shard. Higher is better.
	// Read only field.
	PercentageOfSingleShardWrites *float64 `json:"percentageOfSingleShardWrites,omitempty"`
	// Unmodified result of the `analyzeShardKey` command, as a JSON string. Its content tracks the MongoDB server version of the analyzed cluster rather than the version of this resource, so members outside the parameters above may appear or change without a new resource version. Do not parse it as a stable contract. `mostCommonValues` is removed from it for callers that cannot read the collection's data. This parameter is absent if MongoDB Cloud could not serialize the command result, which costs the raw view but not the analysis.
	// Read only field.
	RawOutput      *string                 `json:"rawOutput,omitempty"`
	ReadSampleSize *ShardKeyReadSampleSize `json:"readSampleSize,omitempty"`
	// Candidate shard key that was analyzed, given as one object per shard key field in the order the fields form the key. This parameter is absent when MongoDB Cloud has no shard key recorded for the analysis.
	ShardKey        *[]map[string]string     `json:"shardKey,omitempty"`
	WriteSampleSize *ShardKeyWriteSampleSize `json:"writeSampleSize,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardKeyAnalysisResponse) MarshalJSON() ([]byte, error) {
	type noMethod ShardKeyAnalysisResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardKeyAnalysisResponse instantiates a new ShardKeyAnalysisResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardKeyAnalysisResponse(createdAt time.Time, expiresAt time.Time, id string, namespace string) *ShardKeyAnalysisResponse {
	this := ShardKeyAnalysisResponse{}
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.Id = id
	this.Namespace = namespace
	return &this
}

// NewShardKeyAnalysisResponseWithDefaults instantiates a new ShardKeyAnalysisResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardKeyAnalysisResponseWithDefaults() *ShardKeyAnalysisResponse {
	this := ShardKeyAnalysisResponse{}
	return &this
}

// GetAvgDocSizeBytes returns the AvgDocSizeBytes field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetAvgDocSizeBytes() int64 {
	if o == nil || IsNil(o.AvgDocSizeBytes) {
		var ret int64
		return ret
	}
	return *o.AvgDocSizeBytes
}

// GetAvgDocSizeBytesOk returns a tuple with the AvgDocSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetAvgDocSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.AvgDocSizeBytes) {
		return nil, false
	}

	return o.AvgDocSizeBytes, true
}

// HasAvgDocSizeBytes returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasAvgDocSizeBytes() bool {
	if o != nil && !IsNil(o.AvgDocSizeBytes) {
		return true
	}

	return false
}

// SetAvgDocSizeBytes gets a reference to the given int64 and assigns it to the AvgDocSizeBytes field.
func (o *ShardKeyAnalysisResponse) SetAvgDocSizeBytes(v int64) {
	o.AvgDocSizeBytes = &v
	o.NullFields = removeNullField(o.NullFields, "AvgDocSizeBytes")
}

// SetAvgDocSizeBytesNil sets AvgDocSizeBytes to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetAvgDocSizeBytesNil() {
	o.AvgDocSizeBytes = nil
	o.NullFields = addNullField(o.NullFields, "AvgDocSizeBytes")
}

// GetCorrelationCoefficient returns the CorrelationCoefficient field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetCorrelationCoefficient() float64 {
	if o == nil || IsNil(o.CorrelationCoefficient) {
		var ret float64
		return ret
	}
	return *o.CorrelationCoefficient
}

// GetCorrelationCoefficientOk returns a tuple with the CorrelationCoefficient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetCorrelationCoefficientOk() (*float64, bool) {
	if o == nil || IsNil(o.CorrelationCoefficient) {
		return nil, false
	}

	return o.CorrelationCoefficient, true
}

// HasCorrelationCoefficient returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasCorrelationCoefficient() bool {
	if o != nil && !IsNil(o.CorrelationCoefficient) {
		return true
	}

	return false
}

// SetCorrelationCoefficient gets a reference to the given float64 and assigns it to the CorrelationCoefficient field.
func (o *ShardKeyAnalysisResponse) SetCorrelationCoefficient(v float64) {
	o.CorrelationCoefficient = &v
	o.NullFields = removeNullField(o.NullFields, "CorrelationCoefficient")
}

// SetCorrelationCoefficientNil sets CorrelationCoefficient to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetCorrelationCoefficientNil() {
	o.CorrelationCoefficient = nil
	o.NullFields = addNullField(o.NullFields, "CorrelationCoefficient")
}

// GetCreatedAt returns the CreatedAt field value
func (o *ShardKeyAnalysisResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ShardKeyAnalysisResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *ShardKeyAnalysisResponse) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *ShardKeyAnalysisResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *ShardKeyAnalysisResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShardKeyAnalysisResponse) SetId(v string) {
	o.Id = v
}

// GetIsUnique returns the IsUnique field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetIsUnique() bool {
	if o == nil || IsNil(o.IsUnique) {
		var ret bool
		return ret
	}
	return *o.IsUnique
}

// GetIsUniqueOk returns a tuple with the IsUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetIsUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnique) {
		return nil, false
	}

	return o.IsUnique, true
}

// HasIsUnique returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasIsUnique() bool {
	if o != nil && !IsNil(o.IsUnique) {
		return true
	}

	return false
}

// SetIsUnique gets a reference to the given bool and assigns it to the IsUnique field.
func (o *ShardKeyAnalysisResponse) SetIsUnique(v bool) {
	o.IsUnique = &v
	o.NullFields = removeNullField(o.NullFields, "IsUnique")
}

// SetIsUniqueNil sets IsUnique to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetIsUniqueNil() {
	o.IsUnique = nil
	o.NullFields = addNullField(o.NullFields, "IsUnique")
}

// GetMonotonicity returns the Monotonicity field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetMonotonicity() string {
	if o == nil || IsNil(o.Monotonicity) {
		var ret string
		return ret
	}
	return *o.Monotonicity
}

// GetMonotonicityOk returns a tuple with the Monotonicity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetMonotonicityOk() (*string, bool) {
	if o == nil || IsNil(o.Monotonicity) {
		return nil, false
	}

	return o.Monotonicity, true
}

// HasMonotonicity returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasMonotonicity() bool {
	if o != nil && !IsNil(o.Monotonicity) {
		return true
	}

	return false
}

// SetMonotonicity gets a reference to the given string and assigns it to the Monotonicity field.
func (o *ShardKeyAnalysisResponse) SetMonotonicity(v string) {
	o.Monotonicity = &v
	o.NullFields = removeNullField(o.NullFields, "Monotonicity")
}

// SetMonotonicityNil sets Monotonicity to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetMonotonicityNil() {
	o.Monotonicity = nil
	o.NullFields = addNullField(o.NullFields, "Monotonicity")
}

// GetMostCommonValues returns the MostCommonValues field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetMostCommonValues() []ShardKeyMostCommonValue {
	if o == nil || IsNil(o.MostCommonValues) {
		var ret []ShardKeyMostCommonValue
		return ret
	}
	return *o.MostCommonValues
}

// GetMostCommonValuesOk returns a tuple with the MostCommonValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetMostCommonValuesOk() (*[]ShardKeyMostCommonValue, bool) {
	if o == nil || IsNil(o.MostCommonValues) {
		return nil, false
	}

	return o.MostCommonValues, true
}

// HasMostCommonValues returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasMostCommonValues() bool {
	if o != nil && !IsNil(o.MostCommonValues) {
		return true
	}

	return false
}

// SetMostCommonValues gets a reference to the given []ShardKeyMostCommonValue and assigns it to the MostCommonValues field.
func (o *ShardKeyAnalysisResponse) SetMostCommonValues(v []ShardKeyMostCommonValue) {
	o.MostCommonValues = &v
	o.NullFields = removeNullField(o.NullFields, "MostCommonValues")
}

// SetMostCommonValuesNil sets MostCommonValues to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetMostCommonValuesNil() {
	o.MostCommonValues = nil
	o.NullFields = addNullField(o.NullFields, "MostCommonValues")
}

// GetNamespace returns the Namespace field value
func (o *ShardKeyAnalysisResponse) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ShardKeyAnalysisResponse) SetNamespace(v string) {
	o.Namespace = v
}

// GetNote returns the Note field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}

	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *ShardKeyAnalysisResponse) SetNote(v string) {
	o.Note = &v
	o.NullFields = removeNullField(o.NullFields, "Note")
}

// SetNoteNil sets Note to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetNoteNil() {
	o.Note = nil
	o.NullFields = addNullField(o.NullFields, "Note")
}

// GetNumDistinctValues returns the NumDistinctValues field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetNumDistinctValues() int64 {
	if o == nil || IsNil(o.NumDistinctValues) {
		var ret int64
		return ret
	}
	return *o.NumDistinctValues
}

// GetNumDistinctValuesOk returns a tuple with the NumDistinctValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetNumDistinctValuesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumDistinctValues) {
		return nil, false
	}

	return o.NumDistinctValues, true
}

// HasNumDistinctValues returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasNumDistinctValues() bool {
	if o != nil && !IsNil(o.NumDistinctValues) {
		return true
	}

	return false
}

// SetNumDistinctValues gets a reference to the given int64 and assigns it to the NumDistinctValues field.
func (o *ShardKeyAnalysisResponse) SetNumDistinctValues(v int64) {
	o.NumDistinctValues = &v
	o.NullFields = removeNullField(o.NullFields, "NumDistinctValues")
}

// SetNumDistinctValuesNil sets NumDistinctValues to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetNumDistinctValuesNil() {
	o.NumDistinctValues = nil
	o.NullFields = addNullField(o.NullFields, "NumDistinctValues")
}

// GetNumDocsSampled returns the NumDocsSampled field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetNumDocsSampled() int64 {
	if o == nil || IsNil(o.NumDocsSampled) {
		var ret int64
		return ret
	}
	return *o.NumDocsSampled
}

// GetNumDocsSampledOk returns a tuple with the NumDocsSampled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetNumDocsSampledOk() (*int64, bool) {
	if o == nil || IsNil(o.NumDocsSampled) {
		return nil, false
	}

	return o.NumDocsSampled, true
}

// HasNumDocsSampled returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasNumDocsSampled() bool {
	if o != nil && !IsNil(o.NumDocsSampled) {
		return true
	}

	return false
}

// SetNumDocsSampled gets a reference to the given int64 and assigns it to the NumDocsSampled field.
func (o *ShardKeyAnalysisResponse) SetNumDocsSampled(v int64) {
	o.NumDocsSampled = &v
	o.NullFields = removeNullField(o.NullFields, "NumDocsSampled")
}

// SetNumDocsSampledNil sets NumDocsSampled to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetNumDocsSampledNil() {
	o.NumDocsSampled = nil
	o.NullFields = addNullField(o.NullFields, "NumDocsSampled")
}

// GetNumDocsTotal returns the NumDocsTotal field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetNumDocsTotal() int64 {
	if o == nil || IsNil(o.NumDocsTotal) {
		var ret int64
		return ret
	}
	return *o.NumDocsTotal
}

// GetNumDocsTotalOk returns a tuple with the NumDocsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetNumDocsTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.NumDocsTotal) {
		return nil, false
	}

	return o.NumDocsTotal, true
}

// HasNumDocsTotal returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasNumDocsTotal() bool {
	if o != nil && !IsNil(o.NumDocsTotal) {
		return true
	}

	return false
}

// SetNumDocsTotal gets a reference to the given int64 and assigns it to the NumDocsTotal field.
func (o *ShardKeyAnalysisResponse) SetNumDocsTotal(v int64) {
	o.NumDocsTotal = &v
	o.NullFields = removeNullField(o.NullFields, "NumDocsTotal")
}

// SetNumDocsTotalNil sets NumDocsTotal to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetNumDocsTotalNil() {
	o.NumDocsTotal = nil
	o.NullFields = addNullField(o.NullFields, "NumDocsTotal")
}

// GetNumOrphanDocs returns the NumOrphanDocs field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetNumOrphanDocs() int64 {
	if o == nil || IsNil(o.NumOrphanDocs) {
		var ret int64
		return ret
	}
	return *o.NumOrphanDocs
}

// GetNumOrphanDocsOk returns a tuple with the NumOrphanDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetNumOrphanDocsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumOrphanDocs) {
		return nil, false
	}

	return o.NumOrphanDocs, true
}

// HasNumOrphanDocs returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasNumOrphanDocs() bool {
	if o != nil && !IsNil(o.NumOrphanDocs) {
		return true
	}

	return false
}

// SetNumOrphanDocs gets a reference to the given int64 and assigns it to the NumOrphanDocs field.
func (o *ShardKeyAnalysisResponse) SetNumOrphanDocs(v int64) {
	o.NumOrphanDocs = &v
	o.NullFields = removeNullField(o.NullFields, "NumOrphanDocs")
}

// SetNumOrphanDocsNil sets NumOrphanDocs to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetNumOrphanDocsNil() {
	o.NumOrphanDocs = nil
	o.NullFields = addNullField(o.NullFields, "NumOrphanDocs")
}

// GetPercentageOfMultiShardReads returns the PercentageOfMultiShardReads field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetPercentageOfMultiShardReads() float64 {
	if o == nil || IsNil(o.PercentageOfMultiShardReads) {
		var ret float64
		return ret
	}
	return *o.PercentageOfMultiShardReads
}

// GetPercentageOfMultiShardReadsOk returns a tuple with the PercentageOfMultiShardReads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetPercentageOfMultiShardReadsOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentageOfMultiShardReads) {
		return nil, false
	}

	return o.PercentageOfMultiShardReads, true
}

// HasPercentageOfMultiShardReads returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasPercentageOfMultiShardReads() bool {
	if o != nil && !IsNil(o.PercentageOfMultiShardReads) {
		return true
	}

	return false
}

// SetPercentageOfMultiShardReads gets a reference to the given float64 and assigns it to the PercentageOfMultiShardReads field.
func (o *ShardKeyAnalysisResponse) SetPercentageOfMultiShardReads(v float64) {
	o.PercentageOfMultiShardReads = &v
	o.NullFields = removeNullField(o.NullFields, "PercentageOfMultiShardReads")
}

// SetPercentageOfMultiShardReadsNil sets PercentageOfMultiShardReads to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetPercentageOfMultiShardReadsNil() {
	o.PercentageOfMultiShardReads = nil
	o.NullFields = addNullField(o.NullFields, "PercentageOfMultiShardReads")
}

// GetPercentageOfMultiShardWrites returns the PercentageOfMultiShardWrites field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetPercentageOfMultiShardWrites() float64 {
	if o == nil || IsNil(o.PercentageOfMultiShardWrites) {
		var ret float64
		return ret
	}
	return *o.PercentageOfMultiShardWrites
}

// GetPercentageOfMultiShardWritesOk returns a tuple with the PercentageOfMultiShardWrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetPercentageOfMultiShardWritesOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentageOfMultiShardWrites) {
		return nil, false
	}

	return o.PercentageOfMultiShardWrites, true
}

// HasPercentageOfMultiShardWrites returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasPercentageOfMultiShardWrites() bool {
	if o != nil && !IsNil(o.PercentageOfMultiShardWrites) {
		return true
	}

	return false
}

// SetPercentageOfMultiShardWrites gets a reference to the given float64 and assigns it to the PercentageOfMultiShardWrites field.
func (o *ShardKeyAnalysisResponse) SetPercentageOfMultiShardWrites(v float64) {
	o.PercentageOfMultiShardWrites = &v
	o.NullFields = removeNullField(o.NullFields, "PercentageOfMultiShardWrites")
}

// SetPercentageOfMultiShardWritesNil sets PercentageOfMultiShardWrites to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetPercentageOfMultiShardWritesNil() {
	o.PercentageOfMultiShardWrites = nil
	o.NullFields = addNullField(o.NullFields, "PercentageOfMultiShardWrites")
}

// GetPercentageOfScatterGatherReads returns the PercentageOfScatterGatherReads field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetPercentageOfScatterGatherReads() float64 {
	if o == nil || IsNil(o.PercentageOfScatterGatherReads) {
		var ret float64
		return ret
	}
	return *o.PercentageOfScatterGatherReads
}

// GetPercentageOfScatterGatherReadsOk returns a tuple with the PercentageOfScatterGatherReads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetPercentageOfScatterGatherReadsOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentageOfScatterGatherReads) {
		return nil, false
	}

	return o.PercentageOfScatterGatherReads, true
}

// HasPercentageOfScatterGatherReads returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasPercentageOfScatterGatherReads() bool {
	if o != nil && !IsNil(o.PercentageOfScatterGatherReads) {
		return true
	}

	return false
}

// SetPercentageOfScatterGatherReads gets a reference to the given float64 and assigns it to the PercentageOfScatterGatherReads field.
func (o *ShardKeyAnalysisResponse) SetPercentageOfScatterGatherReads(v float64) {
	o.PercentageOfScatterGatherReads = &v
	o.NullFields = removeNullField(o.NullFields, "PercentageOfScatterGatherReads")
}

// SetPercentageOfScatterGatherReadsNil sets PercentageOfScatterGatherReads to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetPercentageOfScatterGatherReadsNil() {
	o.PercentageOfScatterGatherReads = nil
	o.NullFields = addNullField(o.NullFields, "PercentageOfScatterGatherReads")
}

// GetPercentageOfScatterGatherWrites returns the PercentageOfScatterGatherWrites field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetPercentageOfScatterGatherWrites() float64 {
	if o == nil || IsNil(o.PercentageOfScatterGatherWrites) {
		var ret float64
		return ret
	}
	return *o.PercentageOfScatterGatherWrites
}

// GetPercentageOfScatterGatherWritesOk returns a tuple with the PercentageOfScatterGatherWrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetPercentageOfScatterGatherWritesOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentageOfScatterGatherWrites) {
		return nil, false
	}

	return o.PercentageOfScatterGatherWrites, true
}

// HasPercentageOfScatterGatherWrites returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasPercentageOfScatterGatherWrites() bool {
	if o != nil && !IsNil(o.PercentageOfScatterGatherWrites) {
		return true
	}

	return false
}

// SetPercentageOfScatterGatherWrites gets a reference to the given float64 and assigns it to the PercentageOfScatterGatherWrites field.
func (o *ShardKeyAnalysisResponse) SetPercentageOfScatterGatherWrites(v float64) {
	o.PercentageOfScatterGatherWrites = &v
	o.NullFields = removeNullField(o.NullFields, "PercentageOfScatterGatherWrites")
}

// SetPercentageOfScatterGatherWritesNil sets PercentageOfScatterGatherWrites to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetPercentageOfScatterGatherWritesNil() {
	o.PercentageOfScatterGatherWrites = nil
	o.NullFields = addNullField(o.NullFields, "PercentageOfScatterGatherWrites")
}

// GetPercentageOfSingleShardReads returns the PercentageOfSingleShardReads field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetPercentageOfSingleShardReads() float64 {
	if o == nil || IsNil(o.PercentageOfSingleShardReads) {
		var ret float64
		return ret
	}
	return *o.PercentageOfSingleShardReads
}

// GetPercentageOfSingleShardReadsOk returns a tuple with the PercentageOfSingleShardReads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetPercentageOfSingleShardReadsOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentageOfSingleShardReads) {
		return nil, false
	}

	return o.PercentageOfSingleShardReads, true
}

// HasPercentageOfSingleShardReads returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasPercentageOfSingleShardReads() bool {
	if o != nil && !IsNil(o.PercentageOfSingleShardReads) {
		return true
	}

	return false
}

// SetPercentageOfSingleShardReads gets a reference to the given float64 and assigns it to the PercentageOfSingleShardReads field.
func (o *ShardKeyAnalysisResponse) SetPercentageOfSingleShardReads(v float64) {
	o.PercentageOfSingleShardReads = &v
	o.NullFields = removeNullField(o.NullFields, "PercentageOfSingleShardReads")
}

// SetPercentageOfSingleShardReadsNil sets PercentageOfSingleShardReads to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetPercentageOfSingleShardReadsNil() {
	o.PercentageOfSingleShardReads = nil
	o.NullFields = addNullField(o.NullFields, "PercentageOfSingleShardReads")
}

// GetPercentageOfSingleShardWrites returns the PercentageOfSingleShardWrites field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetPercentageOfSingleShardWrites() float64 {
	if o == nil || IsNil(o.PercentageOfSingleShardWrites) {
		var ret float64
		return ret
	}
	return *o.PercentageOfSingleShardWrites
}

// GetPercentageOfSingleShardWritesOk returns a tuple with the PercentageOfSingleShardWrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetPercentageOfSingleShardWritesOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentageOfSingleShardWrites) {
		return nil, false
	}

	return o.PercentageOfSingleShardWrites, true
}

// HasPercentageOfSingleShardWrites returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasPercentageOfSingleShardWrites() bool {
	if o != nil && !IsNil(o.PercentageOfSingleShardWrites) {
		return true
	}

	return false
}

// SetPercentageOfSingleShardWrites gets a reference to the given float64 and assigns it to the PercentageOfSingleShardWrites field.
func (o *ShardKeyAnalysisResponse) SetPercentageOfSingleShardWrites(v float64) {
	o.PercentageOfSingleShardWrites = &v
	o.NullFields = removeNullField(o.NullFields, "PercentageOfSingleShardWrites")
}

// SetPercentageOfSingleShardWritesNil sets PercentageOfSingleShardWrites to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetPercentageOfSingleShardWritesNil() {
	o.PercentageOfSingleShardWrites = nil
	o.NullFields = addNullField(o.NullFields, "PercentageOfSingleShardWrites")
}

// GetRawOutput returns the RawOutput field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetRawOutput() string {
	if o == nil || IsNil(o.RawOutput) {
		var ret string
		return ret
	}
	return *o.RawOutput
}

// GetRawOutputOk returns a tuple with the RawOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetRawOutputOk() (*string, bool) {
	if o == nil || IsNil(o.RawOutput) {
		return nil, false
	}

	return o.RawOutput, true
}

// HasRawOutput returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasRawOutput() bool {
	if o != nil && !IsNil(o.RawOutput) {
		return true
	}

	return false
}

// SetRawOutput gets a reference to the given string and assigns it to the RawOutput field.
func (o *ShardKeyAnalysisResponse) SetRawOutput(v string) {
	o.RawOutput = &v
	o.NullFields = removeNullField(o.NullFields, "RawOutput")
}

// SetRawOutputNil sets RawOutput to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetRawOutputNil() {
	o.RawOutput = nil
	o.NullFields = addNullField(o.NullFields, "RawOutput")
}

// GetReadSampleSize returns the ReadSampleSize field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetReadSampleSize() ShardKeyReadSampleSize {
	if o == nil || IsNil(o.ReadSampleSize) {
		var ret ShardKeyReadSampleSize
		return ret
	}
	return *o.ReadSampleSize
}

// GetReadSampleSizeOk returns a tuple with the ReadSampleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetReadSampleSizeOk() (*ShardKeyReadSampleSize, bool) {
	if o == nil || IsNil(o.ReadSampleSize) {
		return nil, false
	}

	return o.ReadSampleSize, true
}

// HasReadSampleSize returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasReadSampleSize() bool {
	if o != nil && !IsNil(o.ReadSampleSize) {
		return true
	}

	return false
}

// SetReadSampleSize gets a reference to the given ShardKeyReadSampleSize and assigns it to the ReadSampleSize field.
func (o *ShardKeyAnalysisResponse) SetReadSampleSize(v ShardKeyReadSampleSize) {
	o.ReadSampleSize = &v
	o.NullFields = removeNullField(o.NullFields, "ReadSampleSize")
}

// SetReadSampleSizeNil sets ReadSampleSize to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetReadSampleSizeNil() {
	o.ReadSampleSize = nil
	o.NullFields = addNullField(o.NullFields, "ReadSampleSize")
}

// GetShardKey returns the ShardKey field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetShardKey() []map[string]string {
	if o == nil || IsNil(o.ShardKey) {
		var ret []map[string]string
		return ret
	}
	return *o.ShardKey
}

// GetShardKeyOk returns a tuple with the ShardKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetShardKeyOk() (*[]map[string]string, bool) {
	if o == nil || IsNil(o.ShardKey) {
		return nil, false
	}

	return o.ShardKey, true
}

// HasShardKey returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasShardKey() bool {
	if o != nil && !IsNil(o.ShardKey) {
		return true
	}

	return false
}

// SetShardKey gets a reference to the given []map[string]string and assigns it to the ShardKey field.
func (o *ShardKeyAnalysisResponse) SetShardKey(v []map[string]string) {
	o.ShardKey = &v
	o.NullFields = removeNullField(o.NullFields, "ShardKey")
}

// SetShardKeyNil sets ShardKey to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetShardKeyNil() {
	o.ShardKey = nil
	o.NullFields = addNullField(o.NullFields, "ShardKey")
}

// GetWriteSampleSize returns the WriteSampleSize field value if set, zero value otherwise
func (o *ShardKeyAnalysisResponse) GetWriteSampleSize() ShardKeyWriteSampleSize {
	if o == nil || IsNil(o.WriteSampleSize) {
		var ret ShardKeyWriteSampleSize
		return ret
	}
	return *o.WriteSampleSize
}

// GetWriteSampleSizeOk returns a tuple with the WriteSampleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisResponse) GetWriteSampleSizeOk() (*ShardKeyWriteSampleSize, bool) {
	if o == nil || IsNil(o.WriteSampleSize) {
		return nil, false
	}

	return o.WriteSampleSize, true
}

// HasWriteSampleSize returns a boolean if a field has been set.
func (o *ShardKeyAnalysisResponse) HasWriteSampleSize() bool {
	if o != nil && !IsNil(o.WriteSampleSize) {
		return true
	}

	return false
}

// SetWriteSampleSize gets a reference to the given ShardKeyWriteSampleSize and assigns it to the WriteSampleSize field.
func (o *ShardKeyAnalysisResponse) SetWriteSampleSize(v ShardKeyWriteSampleSize) {
	o.WriteSampleSize = &v
	o.NullFields = removeNullField(o.NullFields, "WriteSampleSize")
}

// SetWriteSampleSizeNil sets WriteSampleSize to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisResponse) SetWriteSampleSizeNil() {
	o.WriteSampleSize = nil
	o.NullFields = addNullField(o.NullFields, "WriteSampleSize")
}
