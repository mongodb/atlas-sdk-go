// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendationSummary Summary of the recommended change. The type field identifies which recommendation type the summary describes and which fields are present.
type ShardingRecommendationSummary struct {
	// Recommendation type that this summary describes. This value determines which properties are present.
	// Read only field.
	Type string `json:"type"`
	// Average disk space used across the cluster's shards, in bytes. This parameter is absent when no measurement was available.
	// Read only field.
	AvgDiskSpaceUsedBytes *int64 `json:"avgDiskSpaceUsedBytes,omitempty"`
	// Flag that indicates whether the average disk space used breached its warning threshold.
	// Read only field.
	AvgDiskSpaceWarning *bool `json:"avgDiskSpaceWarning,omitempty"`
	// Largest collection size per shard in the cluster, in bytes. This parameter is absent when no measurement was available.
	// Read only field.
	CollectionSizePerShardBytes *int64 `json:"collectionSizePerShardBytes,omitempty"`
	// Flag that indicates whether the collection size per shard breached its warning threshold.
	// Read only field.
	CollectionSizeWarning *bool `json:"collectionSizeWarning,omitempty"`
	// Number of shards in the cluster when the recommendation was generated.
	// Read only field.
	CurrentShardCount *int `json:"currentShardCount,omitempty"`
	// Recommended total number of shards for the cluster.
	// Read only field.
	RecommendedShardCount *int `json:"recommendedShardCount,omitempty"`
	// Disk space that the collection and its indexes occupy, in bytes. This parameter is absent when no measurement was available.
	// Read only field.
	CollectionDiskFootprintBytes *int64 `json:"collectionDiskFootprintBytes,omitempty"`
	// Name of the shard that hosts the collection.
	// Read only field.
	CurrentShard *string `json:"currentShard,omitempty"`
	// Flag that indicates whether the destination shard's free disk space breached its headroom threshold.
	// Read only field.
	HeadroomWarning *bool `json:"headroomWarning,omitempty"`
	// Name of the recommended destination shard. This parameter is absent when no destination shard has enough free disk space.
	// Read only field.
	ProposedShard *string `json:"proposedShard,omitempty"`
	// Free disk space on the recommended destination shard, in bytes.
	// Read only field.
	ProposedShardFreeBytes *int64 `json:"proposedShardFreeBytes,omitempty"`
	// Flag that indicates whether the cluster uses Queryable Encryption, which the `moveCollection` operation doesn't support.
	// Read only field.
	QueryableEncryptionWarning *bool `json:"queryableEncryptionWarning,omitempty"`
	// Free disk space that a destination shard requires to accept the collection, in bytes.
	// Read only field.
	RequiredFreeBytes *int64 `json:"requiredFreeBytes,omitempty"`
	// Uncompressed size of the collection's data, in bytes. This parameter is absent when no measurement was available.
	// Read only field.
	CollectionDataSizeBytes *int64 `json:"collectionDataSizeBytes,omitempty"`
	// Peak growth of the collection's data size over the evaluation window, as a percentage. This parameter is absent when growth did not contribute to the recommendation.
	// Read only field.
	CollectionDataSizeGrowthPercent *float64 `json:"collectionDataSizeGrowthPercent,omitempty"`
	// Flag that indicates whether the collection data size breached its warning threshold.
	// Read only field.
	CollectionDataSizeWarning *bool `json:"collectionDataSizeWarning,omitempty"`
	// Number of documents in the collection. This parameter is absent when no measurement was available.
	// Read only field.
	DocumentCount *int64 `json:"documentCount,omitempty"`
	// Flag that indicates whether the document count breached its warning threshold.
	// Read only field.
	DocumentCountWarning *bool `json:"documentCountWarning,omitempty"`
	// Flag that indicates whether the disk space used breached its warning threshold.
	// Read only field.
	DiskSpaceByteWarning *bool `json:"diskSpaceByteWarning,omitempty"`
	// Flag that indicates whether the percentage of disk space used breached its warning threshold.
	// Read only field.
	DiskSpacePercentWarning *bool `json:"diskSpacePercentWarning,omitempty"`
	// Largest amount of disk space used across the cluster's hosts, in bytes. This parameter is absent when no measurement was available.
	// Read only field.
	MaxDiskSpaceUsedBytes *int64 `json:"maxDiskSpaceUsedBytes,omitempty"`
	// Largest percentage of disk space used across the cluster's hosts. This parameter is absent when no measurement was available.
	// Read only field.
	MaxDiskSpaceUsedPercent *float64 `json:"maxDiskSpaceUsedPercent,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingRecommendationSummary) MarshalJSON() ([]byte, error) {
	type noMethod ShardingRecommendationSummary
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingRecommendationSummary instantiates a new ShardingRecommendationSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingRecommendationSummary(type_ string) *ShardingRecommendationSummary {
	this := ShardingRecommendationSummary{}
	this.Type = type_
	return &this
}

// NewShardingRecommendationSummaryWithDefaults instantiates a new ShardingRecommendationSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingRecommendationSummaryWithDefaults() *ShardingRecommendationSummary {
	this := ShardingRecommendationSummary{}
	return &this
}

// GetType returns the Type field value
func (o *ShardingRecommendationSummary) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShardingRecommendationSummary) SetType(v string) {
	o.Type = v
}

// GetAvgDiskSpaceUsedBytes returns the AvgDiskSpaceUsedBytes field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetAvgDiskSpaceUsedBytes() int64 {
	if o == nil || IsNil(o.AvgDiskSpaceUsedBytes) {
		var ret int64
		return ret
	}
	return *o.AvgDiskSpaceUsedBytes
}

// GetAvgDiskSpaceUsedBytesOk returns a tuple with the AvgDiskSpaceUsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetAvgDiskSpaceUsedBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.AvgDiskSpaceUsedBytes) {
		return nil, false
	}

	return o.AvgDiskSpaceUsedBytes, true
}

// HasAvgDiskSpaceUsedBytes returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasAvgDiskSpaceUsedBytes() bool {
	if o != nil && !IsNil(o.AvgDiskSpaceUsedBytes) {
		return true
	}

	return false
}

// SetAvgDiskSpaceUsedBytes gets a reference to the given int64 and assigns it to the AvgDiskSpaceUsedBytes field.
func (o *ShardingRecommendationSummary) SetAvgDiskSpaceUsedBytes(v int64) {
	o.AvgDiskSpaceUsedBytes = &v
	o.NullFields = removeNullField(o.NullFields, "AvgDiskSpaceUsedBytes")
}

// SetAvgDiskSpaceUsedBytesNil sets AvgDiskSpaceUsedBytes to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetAvgDiskSpaceUsedBytesNil() {
	o.AvgDiskSpaceUsedBytes = nil
	o.NullFields = addNullField(o.NullFields, "AvgDiskSpaceUsedBytes")
}

// GetAvgDiskSpaceWarning returns the AvgDiskSpaceWarning field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetAvgDiskSpaceWarning() bool {
	if o == nil || IsNil(o.AvgDiskSpaceWarning) {
		var ret bool
		return ret
	}
	return *o.AvgDiskSpaceWarning
}

// GetAvgDiskSpaceWarningOk returns a tuple with the AvgDiskSpaceWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetAvgDiskSpaceWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.AvgDiskSpaceWarning) {
		return nil, false
	}

	return o.AvgDiskSpaceWarning, true
}

// HasAvgDiskSpaceWarning returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasAvgDiskSpaceWarning() bool {
	if o != nil && !IsNil(o.AvgDiskSpaceWarning) {
		return true
	}

	return false
}

// SetAvgDiskSpaceWarning gets a reference to the given bool and assigns it to the AvgDiskSpaceWarning field.
func (o *ShardingRecommendationSummary) SetAvgDiskSpaceWarning(v bool) {
	o.AvgDiskSpaceWarning = &v
	o.NullFields = removeNullField(o.NullFields, "AvgDiskSpaceWarning")
}

// SetAvgDiskSpaceWarningNil sets AvgDiskSpaceWarning to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetAvgDiskSpaceWarningNil() {
	o.AvgDiskSpaceWarning = nil
	o.NullFields = addNullField(o.NullFields, "AvgDiskSpaceWarning")
}

// GetCollectionSizePerShardBytes returns the CollectionSizePerShardBytes field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetCollectionSizePerShardBytes() int64 {
	if o == nil || IsNil(o.CollectionSizePerShardBytes) {
		var ret int64
		return ret
	}
	return *o.CollectionSizePerShardBytes
}

// GetCollectionSizePerShardBytesOk returns a tuple with the CollectionSizePerShardBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetCollectionSizePerShardBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.CollectionSizePerShardBytes) {
		return nil, false
	}

	return o.CollectionSizePerShardBytes, true
}

// HasCollectionSizePerShardBytes returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasCollectionSizePerShardBytes() bool {
	if o != nil && !IsNil(o.CollectionSizePerShardBytes) {
		return true
	}

	return false
}

// SetCollectionSizePerShardBytes gets a reference to the given int64 and assigns it to the CollectionSizePerShardBytes field.
func (o *ShardingRecommendationSummary) SetCollectionSizePerShardBytes(v int64) {
	o.CollectionSizePerShardBytes = &v
	o.NullFields = removeNullField(o.NullFields, "CollectionSizePerShardBytes")
}

// SetCollectionSizePerShardBytesNil sets CollectionSizePerShardBytes to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetCollectionSizePerShardBytesNil() {
	o.CollectionSizePerShardBytes = nil
	o.NullFields = addNullField(o.NullFields, "CollectionSizePerShardBytes")
}

// GetCollectionSizeWarning returns the CollectionSizeWarning field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetCollectionSizeWarning() bool {
	if o == nil || IsNil(o.CollectionSizeWarning) {
		var ret bool
		return ret
	}
	return *o.CollectionSizeWarning
}

// GetCollectionSizeWarningOk returns a tuple with the CollectionSizeWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetCollectionSizeWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.CollectionSizeWarning) {
		return nil, false
	}

	return o.CollectionSizeWarning, true
}

// HasCollectionSizeWarning returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasCollectionSizeWarning() bool {
	if o != nil && !IsNil(o.CollectionSizeWarning) {
		return true
	}

	return false
}

// SetCollectionSizeWarning gets a reference to the given bool and assigns it to the CollectionSizeWarning field.
func (o *ShardingRecommendationSummary) SetCollectionSizeWarning(v bool) {
	o.CollectionSizeWarning = &v
	o.NullFields = removeNullField(o.NullFields, "CollectionSizeWarning")
}

// SetCollectionSizeWarningNil sets CollectionSizeWarning to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetCollectionSizeWarningNil() {
	o.CollectionSizeWarning = nil
	o.NullFields = addNullField(o.NullFields, "CollectionSizeWarning")
}

// GetCurrentShardCount returns the CurrentShardCount field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetCurrentShardCount() int {
	if o == nil || IsNil(o.CurrentShardCount) {
		var ret int
		return ret
	}
	return *o.CurrentShardCount
}

// GetCurrentShardCountOk returns a tuple with the CurrentShardCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetCurrentShardCountOk() (*int, bool) {
	if o == nil || IsNil(o.CurrentShardCount) {
		return nil, false
	}

	return o.CurrentShardCount, true
}

// HasCurrentShardCount returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasCurrentShardCount() bool {
	if o != nil && !IsNil(o.CurrentShardCount) {
		return true
	}

	return false
}

// SetCurrentShardCount gets a reference to the given int and assigns it to the CurrentShardCount field.
func (o *ShardingRecommendationSummary) SetCurrentShardCount(v int) {
	o.CurrentShardCount = &v
	o.NullFields = removeNullField(o.NullFields, "CurrentShardCount")
}

// SetCurrentShardCountNil sets CurrentShardCount to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetCurrentShardCountNil() {
	o.CurrentShardCount = nil
	o.NullFields = addNullField(o.NullFields, "CurrentShardCount")
}

// GetRecommendedShardCount returns the RecommendedShardCount field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetRecommendedShardCount() int {
	if o == nil || IsNil(o.RecommendedShardCount) {
		var ret int
		return ret
	}
	return *o.RecommendedShardCount
}

// GetRecommendedShardCountOk returns a tuple with the RecommendedShardCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetRecommendedShardCountOk() (*int, bool) {
	if o == nil || IsNil(o.RecommendedShardCount) {
		return nil, false
	}

	return o.RecommendedShardCount, true
}

// HasRecommendedShardCount returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasRecommendedShardCount() bool {
	if o != nil && !IsNil(o.RecommendedShardCount) {
		return true
	}

	return false
}

// SetRecommendedShardCount gets a reference to the given int and assigns it to the RecommendedShardCount field.
func (o *ShardingRecommendationSummary) SetRecommendedShardCount(v int) {
	o.RecommendedShardCount = &v
	o.NullFields = removeNullField(o.NullFields, "RecommendedShardCount")
}

// SetRecommendedShardCountNil sets RecommendedShardCount to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetRecommendedShardCountNil() {
	o.RecommendedShardCount = nil
	o.NullFields = addNullField(o.NullFields, "RecommendedShardCount")
}

// GetCollectionDiskFootprintBytes returns the CollectionDiskFootprintBytes field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetCollectionDiskFootprintBytes() int64 {
	if o == nil || IsNil(o.CollectionDiskFootprintBytes) {
		var ret int64
		return ret
	}
	return *o.CollectionDiskFootprintBytes
}

// GetCollectionDiskFootprintBytesOk returns a tuple with the CollectionDiskFootprintBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetCollectionDiskFootprintBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.CollectionDiskFootprintBytes) {
		return nil, false
	}

	return o.CollectionDiskFootprintBytes, true
}

// HasCollectionDiskFootprintBytes returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasCollectionDiskFootprintBytes() bool {
	if o != nil && !IsNil(o.CollectionDiskFootprintBytes) {
		return true
	}

	return false
}

// SetCollectionDiskFootprintBytes gets a reference to the given int64 and assigns it to the CollectionDiskFootprintBytes field.
func (o *ShardingRecommendationSummary) SetCollectionDiskFootprintBytes(v int64) {
	o.CollectionDiskFootprintBytes = &v
	o.NullFields = removeNullField(o.NullFields, "CollectionDiskFootprintBytes")
}

// SetCollectionDiskFootprintBytesNil sets CollectionDiskFootprintBytes to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetCollectionDiskFootprintBytesNil() {
	o.CollectionDiskFootprintBytes = nil
	o.NullFields = addNullField(o.NullFields, "CollectionDiskFootprintBytes")
}

// GetCurrentShard returns the CurrentShard field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetCurrentShard() string {
	if o == nil || IsNil(o.CurrentShard) {
		var ret string
		return ret
	}
	return *o.CurrentShard
}

// GetCurrentShardOk returns a tuple with the CurrentShard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetCurrentShardOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentShard) {
		return nil, false
	}

	return o.CurrentShard, true
}

// HasCurrentShard returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasCurrentShard() bool {
	if o != nil && !IsNil(o.CurrentShard) {
		return true
	}

	return false
}

// SetCurrentShard gets a reference to the given string and assigns it to the CurrentShard field.
func (o *ShardingRecommendationSummary) SetCurrentShard(v string) {
	o.CurrentShard = &v
	o.NullFields = removeNullField(o.NullFields, "CurrentShard")
}

// SetCurrentShardNil sets CurrentShard to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetCurrentShardNil() {
	o.CurrentShard = nil
	o.NullFields = addNullField(o.NullFields, "CurrentShard")
}

// GetHeadroomWarning returns the HeadroomWarning field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetHeadroomWarning() bool {
	if o == nil || IsNil(o.HeadroomWarning) {
		var ret bool
		return ret
	}
	return *o.HeadroomWarning
}

// GetHeadroomWarningOk returns a tuple with the HeadroomWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetHeadroomWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.HeadroomWarning) {
		return nil, false
	}

	return o.HeadroomWarning, true
}

// HasHeadroomWarning returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasHeadroomWarning() bool {
	if o != nil && !IsNil(o.HeadroomWarning) {
		return true
	}

	return false
}

// SetHeadroomWarning gets a reference to the given bool and assigns it to the HeadroomWarning field.
func (o *ShardingRecommendationSummary) SetHeadroomWarning(v bool) {
	o.HeadroomWarning = &v
	o.NullFields = removeNullField(o.NullFields, "HeadroomWarning")
}

// SetHeadroomWarningNil sets HeadroomWarning to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetHeadroomWarningNil() {
	o.HeadroomWarning = nil
	o.NullFields = addNullField(o.NullFields, "HeadroomWarning")
}

// GetProposedShard returns the ProposedShard field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetProposedShard() string {
	if o == nil || IsNil(o.ProposedShard) {
		var ret string
		return ret
	}
	return *o.ProposedShard
}

// GetProposedShardOk returns a tuple with the ProposedShard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetProposedShardOk() (*string, bool) {
	if o == nil || IsNil(o.ProposedShard) {
		return nil, false
	}

	return o.ProposedShard, true
}

// HasProposedShard returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasProposedShard() bool {
	if o != nil && !IsNil(o.ProposedShard) {
		return true
	}

	return false
}

// SetProposedShard gets a reference to the given string and assigns it to the ProposedShard field.
func (o *ShardingRecommendationSummary) SetProposedShard(v string) {
	o.ProposedShard = &v
	o.NullFields = removeNullField(o.NullFields, "ProposedShard")
}

// SetProposedShardNil sets ProposedShard to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetProposedShardNil() {
	o.ProposedShard = nil
	o.NullFields = addNullField(o.NullFields, "ProposedShard")
}

// GetProposedShardFreeBytes returns the ProposedShardFreeBytes field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetProposedShardFreeBytes() int64 {
	if o == nil || IsNil(o.ProposedShardFreeBytes) {
		var ret int64
		return ret
	}
	return *o.ProposedShardFreeBytes
}

// GetProposedShardFreeBytesOk returns a tuple with the ProposedShardFreeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetProposedShardFreeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.ProposedShardFreeBytes) {
		return nil, false
	}

	return o.ProposedShardFreeBytes, true
}

// HasProposedShardFreeBytes returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasProposedShardFreeBytes() bool {
	if o != nil && !IsNil(o.ProposedShardFreeBytes) {
		return true
	}

	return false
}

// SetProposedShardFreeBytes gets a reference to the given int64 and assigns it to the ProposedShardFreeBytes field.
func (o *ShardingRecommendationSummary) SetProposedShardFreeBytes(v int64) {
	o.ProposedShardFreeBytes = &v
	o.NullFields = removeNullField(o.NullFields, "ProposedShardFreeBytes")
}

// SetProposedShardFreeBytesNil sets ProposedShardFreeBytes to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetProposedShardFreeBytesNil() {
	o.ProposedShardFreeBytes = nil
	o.NullFields = addNullField(o.NullFields, "ProposedShardFreeBytes")
}

// GetQueryableEncryptionWarning returns the QueryableEncryptionWarning field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetQueryableEncryptionWarning() bool {
	if o == nil || IsNil(o.QueryableEncryptionWarning) {
		var ret bool
		return ret
	}
	return *o.QueryableEncryptionWarning
}

// GetQueryableEncryptionWarningOk returns a tuple with the QueryableEncryptionWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetQueryableEncryptionWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.QueryableEncryptionWarning) {
		return nil, false
	}

	return o.QueryableEncryptionWarning, true
}

// HasQueryableEncryptionWarning returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasQueryableEncryptionWarning() bool {
	if o != nil && !IsNil(o.QueryableEncryptionWarning) {
		return true
	}

	return false
}

// SetQueryableEncryptionWarning gets a reference to the given bool and assigns it to the QueryableEncryptionWarning field.
func (o *ShardingRecommendationSummary) SetQueryableEncryptionWarning(v bool) {
	o.QueryableEncryptionWarning = &v
	o.NullFields = removeNullField(o.NullFields, "QueryableEncryptionWarning")
}

// SetQueryableEncryptionWarningNil sets QueryableEncryptionWarning to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetQueryableEncryptionWarningNil() {
	o.QueryableEncryptionWarning = nil
	o.NullFields = addNullField(o.NullFields, "QueryableEncryptionWarning")
}

// GetRequiredFreeBytes returns the RequiredFreeBytes field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetRequiredFreeBytes() int64 {
	if o == nil || IsNil(o.RequiredFreeBytes) {
		var ret int64
		return ret
	}
	return *o.RequiredFreeBytes
}

// GetRequiredFreeBytesOk returns a tuple with the RequiredFreeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetRequiredFreeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.RequiredFreeBytes) {
		return nil, false
	}

	return o.RequiredFreeBytes, true
}

// HasRequiredFreeBytes returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasRequiredFreeBytes() bool {
	if o != nil && !IsNil(o.RequiredFreeBytes) {
		return true
	}

	return false
}

// SetRequiredFreeBytes gets a reference to the given int64 and assigns it to the RequiredFreeBytes field.
func (o *ShardingRecommendationSummary) SetRequiredFreeBytes(v int64) {
	o.RequiredFreeBytes = &v
	o.NullFields = removeNullField(o.NullFields, "RequiredFreeBytes")
}

// SetRequiredFreeBytesNil sets RequiredFreeBytes to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetRequiredFreeBytesNil() {
	o.RequiredFreeBytes = nil
	o.NullFields = addNullField(o.NullFields, "RequiredFreeBytes")
}

// GetCollectionDataSizeBytes returns the CollectionDataSizeBytes field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetCollectionDataSizeBytes() int64 {
	if o == nil || IsNil(o.CollectionDataSizeBytes) {
		var ret int64
		return ret
	}
	return *o.CollectionDataSizeBytes
}

// GetCollectionDataSizeBytesOk returns a tuple with the CollectionDataSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetCollectionDataSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.CollectionDataSizeBytes) {
		return nil, false
	}

	return o.CollectionDataSizeBytes, true
}

// HasCollectionDataSizeBytes returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasCollectionDataSizeBytes() bool {
	if o != nil && !IsNil(o.CollectionDataSizeBytes) {
		return true
	}

	return false
}

// SetCollectionDataSizeBytes gets a reference to the given int64 and assigns it to the CollectionDataSizeBytes field.
func (o *ShardingRecommendationSummary) SetCollectionDataSizeBytes(v int64) {
	o.CollectionDataSizeBytes = &v
	o.NullFields = removeNullField(o.NullFields, "CollectionDataSizeBytes")
}

// SetCollectionDataSizeBytesNil sets CollectionDataSizeBytes to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetCollectionDataSizeBytesNil() {
	o.CollectionDataSizeBytes = nil
	o.NullFields = addNullField(o.NullFields, "CollectionDataSizeBytes")
}

// GetCollectionDataSizeGrowthPercent returns the CollectionDataSizeGrowthPercent field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetCollectionDataSizeGrowthPercent() float64 {
	if o == nil || IsNil(o.CollectionDataSizeGrowthPercent) {
		var ret float64
		return ret
	}
	return *o.CollectionDataSizeGrowthPercent
}

// GetCollectionDataSizeGrowthPercentOk returns a tuple with the CollectionDataSizeGrowthPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetCollectionDataSizeGrowthPercentOk() (*float64, bool) {
	if o == nil || IsNil(o.CollectionDataSizeGrowthPercent) {
		return nil, false
	}

	return o.CollectionDataSizeGrowthPercent, true
}

// HasCollectionDataSizeGrowthPercent returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasCollectionDataSizeGrowthPercent() bool {
	if o != nil && !IsNil(o.CollectionDataSizeGrowthPercent) {
		return true
	}

	return false
}

// SetCollectionDataSizeGrowthPercent gets a reference to the given float64 and assigns it to the CollectionDataSizeGrowthPercent field.
func (o *ShardingRecommendationSummary) SetCollectionDataSizeGrowthPercent(v float64) {
	o.CollectionDataSizeGrowthPercent = &v
	o.NullFields = removeNullField(o.NullFields, "CollectionDataSizeGrowthPercent")
}

// SetCollectionDataSizeGrowthPercentNil sets CollectionDataSizeGrowthPercent to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetCollectionDataSizeGrowthPercentNil() {
	o.CollectionDataSizeGrowthPercent = nil
	o.NullFields = addNullField(o.NullFields, "CollectionDataSizeGrowthPercent")
}

// GetCollectionDataSizeWarning returns the CollectionDataSizeWarning field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetCollectionDataSizeWarning() bool {
	if o == nil || IsNil(o.CollectionDataSizeWarning) {
		var ret bool
		return ret
	}
	return *o.CollectionDataSizeWarning
}

// GetCollectionDataSizeWarningOk returns a tuple with the CollectionDataSizeWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetCollectionDataSizeWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.CollectionDataSizeWarning) {
		return nil, false
	}

	return o.CollectionDataSizeWarning, true
}

// HasCollectionDataSizeWarning returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasCollectionDataSizeWarning() bool {
	if o != nil && !IsNil(o.CollectionDataSizeWarning) {
		return true
	}

	return false
}

// SetCollectionDataSizeWarning gets a reference to the given bool and assigns it to the CollectionDataSizeWarning field.
func (o *ShardingRecommendationSummary) SetCollectionDataSizeWarning(v bool) {
	o.CollectionDataSizeWarning = &v
	o.NullFields = removeNullField(o.NullFields, "CollectionDataSizeWarning")
}

// SetCollectionDataSizeWarningNil sets CollectionDataSizeWarning to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetCollectionDataSizeWarningNil() {
	o.CollectionDataSizeWarning = nil
	o.NullFields = addNullField(o.NullFields, "CollectionDataSizeWarning")
}

// GetDocumentCount returns the DocumentCount field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetDocumentCount() int64 {
	if o == nil || IsNil(o.DocumentCount) {
		var ret int64
		return ret
	}
	return *o.DocumentCount
}

// GetDocumentCountOk returns a tuple with the DocumentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetDocumentCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DocumentCount) {
		return nil, false
	}

	return o.DocumentCount, true
}

// HasDocumentCount returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasDocumentCount() bool {
	if o != nil && !IsNil(o.DocumentCount) {
		return true
	}

	return false
}

// SetDocumentCount gets a reference to the given int64 and assigns it to the DocumentCount field.
func (o *ShardingRecommendationSummary) SetDocumentCount(v int64) {
	o.DocumentCount = &v
	o.NullFields = removeNullField(o.NullFields, "DocumentCount")
}

// SetDocumentCountNil sets DocumentCount to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetDocumentCountNil() {
	o.DocumentCount = nil
	o.NullFields = addNullField(o.NullFields, "DocumentCount")
}

// GetDocumentCountWarning returns the DocumentCountWarning field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetDocumentCountWarning() bool {
	if o == nil || IsNil(o.DocumentCountWarning) {
		var ret bool
		return ret
	}
	return *o.DocumentCountWarning
}

// GetDocumentCountWarningOk returns a tuple with the DocumentCountWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetDocumentCountWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.DocumentCountWarning) {
		return nil, false
	}

	return o.DocumentCountWarning, true
}

// HasDocumentCountWarning returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasDocumentCountWarning() bool {
	if o != nil && !IsNil(o.DocumentCountWarning) {
		return true
	}

	return false
}

// SetDocumentCountWarning gets a reference to the given bool and assigns it to the DocumentCountWarning field.
func (o *ShardingRecommendationSummary) SetDocumentCountWarning(v bool) {
	o.DocumentCountWarning = &v
	o.NullFields = removeNullField(o.NullFields, "DocumentCountWarning")
}

// SetDocumentCountWarningNil sets DocumentCountWarning to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetDocumentCountWarningNil() {
	o.DocumentCountWarning = nil
	o.NullFields = addNullField(o.NullFields, "DocumentCountWarning")
}

// GetDiskSpaceByteWarning returns the DiskSpaceByteWarning field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetDiskSpaceByteWarning() bool {
	if o == nil || IsNil(o.DiskSpaceByteWarning) {
		var ret bool
		return ret
	}
	return *o.DiskSpaceByteWarning
}

// GetDiskSpaceByteWarningOk returns a tuple with the DiskSpaceByteWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetDiskSpaceByteWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.DiskSpaceByteWarning) {
		return nil, false
	}

	return o.DiskSpaceByteWarning, true
}

// HasDiskSpaceByteWarning returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasDiskSpaceByteWarning() bool {
	if o != nil && !IsNil(o.DiskSpaceByteWarning) {
		return true
	}

	return false
}

// SetDiskSpaceByteWarning gets a reference to the given bool and assigns it to the DiskSpaceByteWarning field.
func (o *ShardingRecommendationSummary) SetDiskSpaceByteWarning(v bool) {
	o.DiskSpaceByteWarning = &v
	o.NullFields = removeNullField(o.NullFields, "DiskSpaceByteWarning")
}

// SetDiskSpaceByteWarningNil sets DiskSpaceByteWarning to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetDiskSpaceByteWarningNil() {
	o.DiskSpaceByteWarning = nil
	o.NullFields = addNullField(o.NullFields, "DiskSpaceByteWarning")
}

// GetDiskSpacePercentWarning returns the DiskSpacePercentWarning field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetDiskSpacePercentWarning() bool {
	if o == nil || IsNil(o.DiskSpacePercentWarning) {
		var ret bool
		return ret
	}
	return *o.DiskSpacePercentWarning
}

// GetDiskSpacePercentWarningOk returns a tuple with the DiskSpacePercentWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetDiskSpacePercentWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.DiskSpacePercentWarning) {
		return nil, false
	}

	return o.DiskSpacePercentWarning, true
}

// HasDiskSpacePercentWarning returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasDiskSpacePercentWarning() bool {
	if o != nil && !IsNil(o.DiskSpacePercentWarning) {
		return true
	}

	return false
}

// SetDiskSpacePercentWarning gets a reference to the given bool and assigns it to the DiskSpacePercentWarning field.
func (o *ShardingRecommendationSummary) SetDiskSpacePercentWarning(v bool) {
	o.DiskSpacePercentWarning = &v
	o.NullFields = removeNullField(o.NullFields, "DiskSpacePercentWarning")
}

// SetDiskSpacePercentWarningNil sets DiskSpacePercentWarning to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetDiskSpacePercentWarningNil() {
	o.DiskSpacePercentWarning = nil
	o.NullFields = addNullField(o.NullFields, "DiskSpacePercentWarning")
}

// GetMaxDiskSpaceUsedBytes returns the MaxDiskSpaceUsedBytes field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetMaxDiskSpaceUsedBytes() int64 {
	if o == nil || IsNil(o.MaxDiskSpaceUsedBytes) {
		var ret int64
		return ret
	}
	return *o.MaxDiskSpaceUsedBytes
}

// GetMaxDiskSpaceUsedBytesOk returns a tuple with the MaxDiskSpaceUsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetMaxDiskSpaceUsedBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxDiskSpaceUsedBytes) {
		return nil, false
	}

	return o.MaxDiskSpaceUsedBytes, true
}

// HasMaxDiskSpaceUsedBytes returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasMaxDiskSpaceUsedBytes() bool {
	if o != nil && !IsNil(o.MaxDiskSpaceUsedBytes) {
		return true
	}

	return false
}

// SetMaxDiskSpaceUsedBytes gets a reference to the given int64 and assigns it to the MaxDiskSpaceUsedBytes field.
func (o *ShardingRecommendationSummary) SetMaxDiskSpaceUsedBytes(v int64) {
	o.MaxDiskSpaceUsedBytes = &v
	o.NullFields = removeNullField(o.NullFields, "MaxDiskSpaceUsedBytes")
}

// SetMaxDiskSpaceUsedBytesNil sets MaxDiskSpaceUsedBytes to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetMaxDiskSpaceUsedBytesNil() {
	o.MaxDiskSpaceUsedBytes = nil
	o.NullFields = addNullField(o.NullFields, "MaxDiskSpaceUsedBytes")
}

// GetMaxDiskSpaceUsedPercent returns the MaxDiskSpaceUsedPercent field value if set, zero value otherwise
func (o *ShardingRecommendationSummary) GetMaxDiskSpaceUsedPercent() float64 {
	if o == nil || IsNil(o.MaxDiskSpaceUsedPercent) {
		var ret float64
		return ret
	}
	return *o.MaxDiskSpaceUsedPercent
}

// GetMaxDiskSpaceUsedPercentOk returns a tuple with the MaxDiskSpaceUsedPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationSummary) GetMaxDiskSpaceUsedPercentOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxDiskSpaceUsedPercent) {
		return nil, false
	}

	return o.MaxDiskSpaceUsedPercent, true
}

// HasMaxDiskSpaceUsedPercent returns a boolean if a field has been set.
func (o *ShardingRecommendationSummary) HasMaxDiskSpaceUsedPercent() bool {
	if o != nil && !IsNil(o.MaxDiskSpaceUsedPercent) {
		return true
	}

	return false
}

// SetMaxDiskSpaceUsedPercent gets a reference to the given float64 and assigns it to the MaxDiskSpaceUsedPercent field.
func (o *ShardingRecommendationSummary) SetMaxDiskSpaceUsedPercent(v float64) {
	o.MaxDiskSpaceUsedPercent = &v
	o.NullFields = removeNullField(o.NullFields, "MaxDiskSpaceUsedPercent")
}

// SetMaxDiskSpaceUsedPercentNil sets MaxDiskSpaceUsedPercent to an explicit JSON null when marshaled.
func (o *ShardingRecommendationSummary) SetMaxDiskSpaceUsedPercentNil() {
	o.MaxDiskSpaceUsedPercent = nil
	o.NullFields = addNullField(o.NullFields, "MaxDiskSpaceUsedPercent")
}
