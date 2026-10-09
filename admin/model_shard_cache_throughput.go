// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardCacheThroughput Bytes per second that the shard moved in and out of the cache over the last hour.
type ShardCacheThroughput struct {
	// Average bytes per second read into the cache, across the shard's nodes over the last hour.
	// Read only field.
	ReadIntoBytesPerSecond float64 `json:"readIntoBytesPerSecond"`
	// Average bytes per second written from the cache, across the shard's nodes over the last hour.
	// Read only field.
	WrittenFromBytesPerSecond float64 `json:"writtenFromBytesPerSecond"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardCacheThroughput) MarshalJSON() ([]byte, error) {
	type noMethod ShardCacheThroughput
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardCacheThroughput instantiates a new ShardCacheThroughput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardCacheThroughput(readIntoBytesPerSecond float64, writtenFromBytesPerSecond float64) *ShardCacheThroughput {
	this := ShardCacheThroughput{}
	this.ReadIntoBytesPerSecond = readIntoBytesPerSecond
	this.WrittenFromBytesPerSecond = writtenFromBytesPerSecond
	return &this
}

// NewShardCacheThroughputWithDefaults instantiates a new ShardCacheThroughput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardCacheThroughputWithDefaults() *ShardCacheThroughput {
	this := ShardCacheThroughput{}
	return &this
}

// GetReadIntoBytesPerSecond returns the ReadIntoBytesPerSecond field value
func (o *ShardCacheThroughput) GetReadIntoBytesPerSecond() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ReadIntoBytesPerSecond
}

// GetReadIntoBytesPerSecondOk returns a tuple with the ReadIntoBytesPerSecond field value
// and a boolean to check if the value has been set.
func (o *ShardCacheThroughput) GetReadIntoBytesPerSecondOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadIntoBytesPerSecond, true
}

// SetReadIntoBytesPerSecond sets field value
func (o *ShardCacheThroughput) SetReadIntoBytesPerSecond(v float64) {
	o.ReadIntoBytesPerSecond = v
}

// GetWrittenFromBytesPerSecond returns the WrittenFromBytesPerSecond field value
func (o *ShardCacheThroughput) GetWrittenFromBytesPerSecond() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.WrittenFromBytesPerSecond
}

// GetWrittenFromBytesPerSecondOk returns a tuple with the WrittenFromBytesPerSecond field value
// and a boolean to check if the value has been set.
func (o *ShardCacheThroughput) GetWrittenFromBytesPerSecondOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WrittenFromBytesPerSecond, true
}

// SetWrittenFromBytesPerSecond sets field value
func (o *ShardCacheThroughput) SetWrittenFromBytesPerSecond(v float64) {
	o.WrittenFromBytesPerSecond = v
}
