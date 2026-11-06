// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardKeyWriteSampleSize Number of sampled write operations that MongoDB Cloud used to compute the write distribution, broken down by the command that issued them.
type ShardKeyWriteSampleSize struct {
	// Number of sampled `delete` operations.
	// Read only field.
	Delete *int64 `json:"delete,omitempty"`
	// Number of sampled `findAndModify` operations.
	// Read only field.
	FindAndModify *int64 `json:"findAndModify,omitempty"`
	// Total number of sampled write operations.
	// Read only field.
	Total *int64 `json:"total,omitempty"`
	// Number of sampled `update` operations.
	// Read only field.
	Update *int64 `json:"update,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardKeyWriteSampleSize) MarshalJSON() ([]byte, error) {
	type noMethod ShardKeyWriteSampleSize
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardKeyWriteSampleSize instantiates a new ShardKeyWriteSampleSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardKeyWriteSampleSize() *ShardKeyWriteSampleSize {
	this := ShardKeyWriteSampleSize{}
	return &this
}

// NewShardKeyWriteSampleSizeWithDefaults instantiates a new ShardKeyWriteSampleSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardKeyWriteSampleSizeWithDefaults() *ShardKeyWriteSampleSize {
	this := ShardKeyWriteSampleSize{}
	return &this
}

// GetDelete returns the Delete field value if set, zero value otherwise
func (o *ShardKeyWriteSampleSize) GetDelete() int64 {
	if o == nil || IsNil(o.Delete) {
		var ret int64
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyWriteSampleSize) GetDeleteOk() (*int64, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}

	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *ShardKeyWriteSampleSize) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given int64 and assigns it to the Delete field.
func (o *ShardKeyWriteSampleSize) SetDelete(v int64) {
	o.Delete = &v
	o.NullFields = removeNullField(o.NullFields, "Delete")
}

// SetDeleteNil sets Delete to an explicit JSON null when marshaled.
func (o *ShardKeyWriteSampleSize) SetDeleteNil() {
	o.Delete = nil
	o.NullFields = addNullField(o.NullFields, "Delete")
}

// GetFindAndModify returns the FindAndModify field value if set, zero value otherwise
func (o *ShardKeyWriteSampleSize) GetFindAndModify() int64 {
	if o == nil || IsNil(o.FindAndModify) {
		var ret int64
		return ret
	}
	return *o.FindAndModify
}

// GetFindAndModifyOk returns a tuple with the FindAndModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyWriteSampleSize) GetFindAndModifyOk() (*int64, bool) {
	if o == nil || IsNil(o.FindAndModify) {
		return nil, false
	}

	return o.FindAndModify, true
}

// HasFindAndModify returns a boolean if a field has been set.
func (o *ShardKeyWriteSampleSize) HasFindAndModify() bool {
	if o != nil && !IsNil(o.FindAndModify) {
		return true
	}

	return false
}

// SetFindAndModify gets a reference to the given int64 and assigns it to the FindAndModify field.
func (o *ShardKeyWriteSampleSize) SetFindAndModify(v int64) {
	o.FindAndModify = &v
	o.NullFields = removeNullField(o.NullFields, "FindAndModify")
}

// SetFindAndModifyNil sets FindAndModify to an explicit JSON null when marshaled.
func (o *ShardKeyWriteSampleSize) SetFindAndModifyNil() {
	o.FindAndModify = nil
	o.NullFields = addNullField(o.NullFields, "FindAndModify")
}

// GetTotal returns the Total field value if set, zero value otherwise
func (o *ShardKeyWriteSampleSize) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyWriteSampleSize) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}

	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ShardKeyWriteSampleSize) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *ShardKeyWriteSampleSize) SetTotal(v int64) {
	o.Total = &v
	o.NullFields = removeNullField(o.NullFields, "Total")
}

// SetTotalNil sets Total to an explicit JSON null when marshaled.
func (o *ShardKeyWriteSampleSize) SetTotalNil() {
	o.Total = nil
	o.NullFields = addNullField(o.NullFields, "Total")
}

// GetUpdate returns the Update field value if set, zero value otherwise
func (o *ShardKeyWriteSampleSize) GetUpdate() int64 {
	if o == nil || IsNil(o.Update) {
		var ret int64
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyWriteSampleSize) GetUpdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}

	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ShardKeyWriteSampleSize) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given int64 and assigns it to the Update field.
func (o *ShardKeyWriteSampleSize) SetUpdate(v int64) {
	o.Update = &v
	o.NullFields = removeNullField(o.NullFields, "Update")
}

// SetUpdateNil sets Update to an explicit JSON null when marshaled.
func (o *ShardKeyWriteSampleSize) SetUpdateNil() {
	o.Update = nil
	o.NullFields = addNullField(o.NullFields, "Update")
}
