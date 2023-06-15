// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the RollingIndexRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollingIndexRequest{}

// RollingIndexRequest struct for RollingIndexRequest
type RollingIndexRequest struct {
	Collation *Collation `json:"collation,omitempty"`
	// Human-readable label of the collection for which MongoDB Cloud creates an index.
	Collection string `json:"collection"`
	// Human-readable label of the database that holds the collection on which MongoDB Cloud creates an index.
	Db string `json:"db"`
	// List that contains one or more objects that describe the parameters that you want to index.
	Keys    []map[string]string `json:"keys,omitempty"`
	Options *IndexOptions       `json:"options,omitempty"`
}

// NewRollingIndexRequest instantiates a new RollingIndexRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollingIndexRequest(collection string, db string) *RollingIndexRequest {
	this := RollingIndexRequest{}
	this.Collection = collection
	this.Db = db
	return &this
}

// NewRollingIndexRequestWithDefaults instantiates a new RollingIndexRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollingIndexRequestWithDefaults() *RollingIndexRequest {
	this := RollingIndexRequest{}
	return &this
}

// GetCollation returns the Collation field value if set, zero value otherwise.
func (o *RollingIndexRequest) GetCollation() Collation {
	if o == nil || IsNil(o.Collation) {
		var ret Collation
		return ret
	}
	return *o.Collation
}

// GetCollationOk returns a tuple with the Collation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollingIndexRequest) GetCollationOk() (*Collation, bool) {
	if o == nil || IsNil(o.Collation) {
		return nil, false
	}
	return o.Collation, true
}

// HasCollation returns a boolean if a field has been set.
func (o *RollingIndexRequest) HasCollation() bool {
	if o != nil && !IsNil(o.Collation) {
		return true
	}

	return false
}

// SetCollation gets a reference to the given Collation and assigns it to the Collation field.
func (o *RollingIndexRequest) SetCollation(v Collation) {
	o.Collation = &v
}

// GetCollection returns the Collection field value
func (o *RollingIndexRequest) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *RollingIndexRequest) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *RollingIndexRequest) SetCollection(v string) {
	o.Collection = v
}

// GetDb returns the Db field value
func (o *RollingIndexRequest) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *RollingIndexRequest) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value
func (o *RollingIndexRequest) SetDb(v string) {
	o.Db = v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *RollingIndexRequest) GetKeys() []map[string]string {
	if o == nil || IsNil(o.Keys) {
		var ret []map[string]string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollingIndexRequest) GetKeysOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *RollingIndexRequest) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []map[string]string and assigns it to the Keys field.
func (o *RollingIndexRequest) SetKeys(v []map[string]string) {
	o.Keys = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *RollingIndexRequest) GetOptions() IndexOptions {
	if o == nil || IsNil(o.Options) {
		var ret IndexOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollingIndexRequest) GetOptionsOk() (*IndexOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *RollingIndexRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given IndexOptions and assigns it to the Options field.
func (o *RollingIndexRequest) SetOptions(v IndexOptions) {
	o.Options = &v
}

func (o RollingIndexRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o RollingIndexRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collation) {
		toSerialize["collation"] = o.Collation
	}
	toSerialize["collection"] = o.Collection
	toSerialize["db"] = o.Db
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableRollingIndexRequest struct {
	value *RollingIndexRequest
	isSet bool
}

func (v NullableRollingIndexRequest) Get() *RollingIndexRequest {
	return v.value
}

func (v *NullableRollingIndexRequest) Set(val *RollingIndexRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRollingIndexRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRollingIndexRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollingIndexRequest(val *RollingIndexRequest) *NullableRollingIndexRequest {
	return &NullableRollingIndexRequest{value: val, isSet: true}
}

func (v NullableRollingIndexRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollingIndexRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
