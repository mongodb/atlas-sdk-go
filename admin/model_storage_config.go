// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StorageConfig struct for StorageConfig
type StorageConfig struct {
	// Available in Public Preview: Per-shard data-size limit that MongoDB Cloud enforces for this cluster, expressed in gigabytes. This is read-only. When you set `shardSizeLimitGB`, this value usually matches it; otherwise it reflects the default limit that MongoDB Cloud assigns when it creates or updates the cluster. This value may differ from `shardSizeLimitGB` due to system-managed changes. MongoDB Cloud returns this only for Atlas INFINITE clusters.
	// Read only field.
	EffectiveShardSizeLimitGB *int `json:"effectiveShardSizeLimitGB,omitempty"`
	// Available in Public Preview: Maximum data size that MongoDB Cloud allows each shard of this cluster to reach, expressed in gigabytes. MongoDB Cloud rejects writes to a shard that reaches the enforced limit (`effectiveShardSizeLimitGB`). This limit applies to every shard of the cluster; set the same value on each region configuration's `autoScaling`, as MongoDB Cloud rejects requests that specify differing values. You can set this only on Atlas INFINITE clusters. Set this to `null` to remove a previously configured limit; omit it to leave the limit unchanged.
	ShardSizeLimitGB *int `json:"shardSizeLimitGB,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StorageConfig) MarshalJSON() ([]byte, error) {
	type noMethod StorageConfig
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStorageConfig instantiates a new StorageConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageConfig() *StorageConfig {
	this := StorageConfig{}
	return &this
}

// NewStorageConfigWithDefaults instantiates a new StorageConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageConfigWithDefaults() *StorageConfig {
	this := StorageConfig{}
	return &this
}

// GetEffectiveShardSizeLimitGB returns the EffectiveShardSizeLimitGB field value if set, zero value otherwise
func (o *StorageConfig) GetEffectiveShardSizeLimitGB() int {
	if o == nil || IsNil(o.EffectiveShardSizeLimitGB) {
		var ret int
		return ret
	}
	return *o.EffectiveShardSizeLimitGB
}

// GetEffectiveShardSizeLimitGBOk returns a tuple with the EffectiveShardSizeLimitGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageConfig) GetEffectiveShardSizeLimitGBOk() (*int, bool) {
	if o == nil || IsNil(o.EffectiveShardSizeLimitGB) {
		return nil, false
	}

	return o.EffectiveShardSizeLimitGB, true
}

// HasEffectiveShardSizeLimitGB returns a boolean if a field has been set.
func (o *StorageConfig) HasEffectiveShardSizeLimitGB() bool {
	if o != nil && !IsNil(o.EffectiveShardSizeLimitGB) {
		return true
	}

	return false
}

// SetEffectiveShardSizeLimitGB gets a reference to the given int and assigns it to the EffectiveShardSizeLimitGB field.
func (o *StorageConfig) SetEffectiveShardSizeLimitGB(v int) {
	o.EffectiveShardSizeLimitGB = &v
	o.NullFields = removeNullField(o.NullFields, "EffectiveShardSizeLimitGB")
}

// SetEffectiveShardSizeLimitGBNil sets EffectiveShardSizeLimitGB to an explicit JSON null when marshaled.
func (o *StorageConfig) SetEffectiveShardSizeLimitGBNil() {
	o.EffectiveShardSizeLimitGB = nil
	o.NullFields = addNullField(o.NullFields, "EffectiveShardSizeLimitGB")
}

// GetShardSizeLimitGB returns the ShardSizeLimitGB field value if set, zero value otherwise
func (o *StorageConfig) GetShardSizeLimitGB() int {
	if o == nil || IsNil(o.ShardSizeLimitGB) {
		var ret int
		return ret
	}
	return *o.ShardSizeLimitGB
}

// GetShardSizeLimitGBOk returns a tuple with the ShardSizeLimitGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageConfig) GetShardSizeLimitGBOk() (*int, bool) {
	if o == nil || IsNil(o.ShardSizeLimitGB) {
		return nil, false
	}

	return o.ShardSizeLimitGB, true
}

// HasShardSizeLimitGB returns a boolean if a field has been set.
func (o *StorageConfig) HasShardSizeLimitGB() bool {
	if o != nil && !IsNil(o.ShardSizeLimitGB) {
		return true
	}

	return false
}

// SetShardSizeLimitGB gets a reference to the given int and assigns it to the ShardSizeLimitGB field.
func (o *StorageConfig) SetShardSizeLimitGB(v int) {
	o.ShardSizeLimitGB = &v
	o.NullFields = removeNullField(o.NullFields, "ShardSizeLimitGB")
}

// SetShardSizeLimitGBNil sets ShardSizeLimitGB to an explicit JSON null when marshaled.
func (o *StorageConfig) SetShardSizeLimitGBNil() {
	o.ShardSizeLimitGB = nil
	o.NullFields = addNullField(o.NullFields, "ShardSizeLimitGB")
}
