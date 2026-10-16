// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardKeyAnalysisRequest Shard key to analyze for one collection.
type ShardKeyAnalysisRequest struct {
	// Flag that indicates whether to return the cardinality, frequency and monotonicity of the candidate shard key.
	// Write only field.
	IncludeKeyCharacteristics *bool `json:"includeKeyCharacteristics,omitempty"`
	// Flag that indicates whether to return how reads and writes would be routed across shards under the candidate shard key. This requires query sampling to have collected operations for the namespace.
	// Write only field.
	IncludeReadWriteDistribution *bool `json:"includeReadWriteDistribution,omitempty"`
	// Combination of the database and collection to analyze, written as `<database>.<collection>`.
	Namespace string `json:"namespace"`
	// Candidate shard key, given as one object per shard key field in the order the fields form the key. Field order is significant: `[{\"a\": \"1\"}, {\"b\": \"1\"}]` and `[{\"b\": \"1\"}, {\"a\": \"1\"}]` are different shard keys. At most one field may be `hashed`.
	ShardKey []map[string]string `json:"shardKey"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardKeyAnalysisRequest) MarshalJSON() ([]byte, error) {
	type noMethod ShardKeyAnalysisRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardKeyAnalysisRequest instantiates a new ShardKeyAnalysisRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardKeyAnalysisRequest(namespace string, shardKey []map[string]string) *ShardKeyAnalysisRequest {
	this := ShardKeyAnalysisRequest{}
	var includeKeyCharacteristics bool = true
	this.IncludeKeyCharacteristics = &includeKeyCharacteristics
	var includeReadWriteDistribution bool = false
	this.IncludeReadWriteDistribution = &includeReadWriteDistribution
	this.Namespace = namespace
	this.ShardKey = shardKey
	return &this
}

// NewShardKeyAnalysisRequestWithDefaults instantiates a new ShardKeyAnalysisRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardKeyAnalysisRequestWithDefaults() *ShardKeyAnalysisRequest {
	this := ShardKeyAnalysisRequest{}
	var includeKeyCharacteristics bool = true
	this.IncludeKeyCharacteristics = &includeKeyCharacteristics
	var includeReadWriteDistribution bool = false
	this.IncludeReadWriteDistribution = &includeReadWriteDistribution
	return &this
}

// GetIncludeKeyCharacteristics returns the IncludeKeyCharacteristics field value if set, zero value otherwise
func (o *ShardKeyAnalysisRequest) GetIncludeKeyCharacteristics() bool {
	if o == nil || IsNil(o.IncludeKeyCharacteristics) {
		var ret bool
		return ret
	}
	return *o.IncludeKeyCharacteristics
}

// GetIncludeKeyCharacteristicsOk returns a tuple with the IncludeKeyCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisRequest) GetIncludeKeyCharacteristicsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeKeyCharacteristics) {
		return nil, false
	}

	return o.IncludeKeyCharacteristics, true
}

// HasIncludeKeyCharacteristics returns a boolean if a field has been set.
func (o *ShardKeyAnalysisRequest) HasIncludeKeyCharacteristics() bool {
	if o != nil && !IsNil(o.IncludeKeyCharacteristics) {
		return true
	}

	return false
}

// SetIncludeKeyCharacteristics gets a reference to the given bool and assigns it to the IncludeKeyCharacteristics field.
func (o *ShardKeyAnalysisRequest) SetIncludeKeyCharacteristics(v bool) {
	o.IncludeKeyCharacteristics = &v
	o.NullFields = removeNullField(o.NullFields, "IncludeKeyCharacteristics")
}

// SetIncludeKeyCharacteristicsNil sets IncludeKeyCharacteristics to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisRequest) SetIncludeKeyCharacteristicsNil() {
	o.IncludeKeyCharacteristics = nil
	o.NullFields = addNullField(o.NullFields, "IncludeKeyCharacteristics")
}

// GetIncludeReadWriteDistribution returns the IncludeReadWriteDistribution field value if set, zero value otherwise
func (o *ShardKeyAnalysisRequest) GetIncludeReadWriteDistribution() bool {
	if o == nil || IsNil(o.IncludeReadWriteDistribution) {
		var ret bool
		return ret
	}
	return *o.IncludeReadWriteDistribution
}

// GetIncludeReadWriteDistributionOk returns a tuple with the IncludeReadWriteDistribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisRequest) GetIncludeReadWriteDistributionOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeReadWriteDistribution) {
		return nil, false
	}

	return o.IncludeReadWriteDistribution, true
}

// HasIncludeReadWriteDistribution returns a boolean if a field has been set.
func (o *ShardKeyAnalysisRequest) HasIncludeReadWriteDistribution() bool {
	if o != nil && !IsNil(o.IncludeReadWriteDistribution) {
		return true
	}

	return false
}

// SetIncludeReadWriteDistribution gets a reference to the given bool and assigns it to the IncludeReadWriteDistribution field.
func (o *ShardKeyAnalysisRequest) SetIncludeReadWriteDistribution(v bool) {
	o.IncludeReadWriteDistribution = &v
	o.NullFields = removeNullField(o.NullFields, "IncludeReadWriteDistribution")
}

// SetIncludeReadWriteDistributionNil sets IncludeReadWriteDistribution to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisRequest) SetIncludeReadWriteDistributionNil() {
	o.IncludeReadWriteDistribution = nil
	o.NullFields = addNullField(o.NullFields, "IncludeReadWriteDistribution")
}

// GetNamespace returns the Namespace field value
func (o *ShardKeyAnalysisRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ShardKeyAnalysisRequest) SetNamespace(v string) {
	o.Namespace = v
}

// GetShardKey returns the ShardKey field value
func (o *ShardKeyAnalysisRequest) GetShardKey() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.ShardKey
}

// GetShardKeyOk returns a tuple with the ShardKey field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisRequest) GetShardKeyOk() (*[]map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShardKey, true
}

// SetShardKey sets field value
func (o *ShardKeyAnalysisRequest) SetShardKey(v []map[string]string) {
	o.ShardKey = v
}
