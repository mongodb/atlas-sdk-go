// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DataLakeAtlasStoreReadPreferenceTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeAtlasStoreReadPreferenceTag{}

// DataLakeAtlasStoreReadPreferenceTag List that contains [tag sets](https://docs.mongodb.com/manual/core/read-preference-tags/) or tag specification documents. If specified, Atlas Data Lake routes read requests to replica set member or members that are associated with the specified tags.
type DataLakeAtlasStoreReadPreferenceTag struct {
	// Human-readable label of the tag.
	Name *string `json:"name,omitempty"`
	// Value of the tag.
	Value *string `json:"value,omitempty"`
}

// NewDataLakeAtlasStoreReadPreferenceTag instantiates a new DataLakeAtlasStoreReadPreferenceTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeAtlasStoreReadPreferenceTag() *DataLakeAtlasStoreReadPreferenceTag {
	this := DataLakeAtlasStoreReadPreferenceTag{}
	return &this
}

// NewDataLakeAtlasStoreReadPreferenceTagWithDefaults instantiates a new DataLakeAtlasStoreReadPreferenceTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeAtlasStoreReadPreferenceTagWithDefaults() *DataLakeAtlasStoreReadPreferenceTag {
	this := DataLakeAtlasStoreReadPreferenceTag{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataLakeAtlasStoreReadPreferenceTag) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreReadPreferenceTag) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreReadPreferenceTag) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataLakeAtlasStoreReadPreferenceTag) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataLakeAtlasStoreReadPreferenceTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreReadPreferenceTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreReadPreferenceTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DataLakeAtlasStoreReadPreferenceTag) SetValue(v string) {
	o.Value = &v
}

func (o DataLakeAtlasStoreReadPreferenceTag) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeAtlasStoreReadPreferenceTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDataLakeAtlasStoreReadPreferenceTag struct {
	value *DataLakeAtlasStoreReadPreferenceTag
	isSet bool
}

func (v NullableDataLakeAtlasStoreReadPreferenceTag) Get() *DataLakeAtlasStoreReadPreferenceTag {
	return v.value
}

func (v *NullableDataLakeAtlasStoreReadPreferenceTag) Set(val *DataLakeAtlasStoreReadPreferenceTag) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeAtlasStoreReadPreferenceTag) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeAtlasStoreReadPreferenceTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeAtlasStoreReadPreferenceTag(val *DataLakeAtlasStoreReadPreferenceTag) *NullableDataLakeAtlasStoreReadPreferenceTag {
	return &NullableDataLakeAtlasStoreReadPreferenceTag{value: val, isSet: true}
}

func (v NullableDataLakeAtlasStoreReadPreferenceTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeAtlasStoreReadPreferenceTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
