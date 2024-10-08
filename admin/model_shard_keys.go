// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardKeys Document that configures the shard key on the destination cluster.
type ShardKeys struct {
	// List of fields to use for the shard key.
	// Write only field.
	Key *[]any `json:"key,omitempty"`
}

// NewShardKeys instantiates a new ShardKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardKeys() *ShardKeys {
	this := ShardKeys{}
	return &this
}

// NewShardKeysWithDefaults instantiates a new ShardKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardKeysWithDefaults() *ShardKeys {
	this := ShardKeys{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise
func (o *ShardKeys) GetKey() []any {
	if o == nil || IsNil(o.Key) {
		var ret []any
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeys) GetKeyOk() (*[]any, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}

	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ShardKeys) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given []any and assigns it to the Key field.
func (o *ShardKeys) SetKey(v []any) {
	o.Key = &v
}
