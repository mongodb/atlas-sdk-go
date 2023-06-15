// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// checks if the BSONTimestamp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BSONTimestamp{}

// BSONTimestamp BSON timestamp that indicates when the checkpoint token entry in the oplog occurred.
type BSONTimestamp struct {
	// Date and time when the oplog recorded this database operation. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Date *time.Time `json:"date,omitempty"`
	// Order of the database operation that the oplog recorded at specific date and time.
	Increment *int `json:"increment,omitempty"`
}

// NewBSONTimestamp instantiates a new BSONTimestamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBSONTimestamp() *BSONTimestamp {
	this := BSONTimestamp{}
	return &this
}

// NewBSONTimestampWithDefaults instantiates a new BSONTimestamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBSONTimestampWithDefaults() *BSONTimestamp {
	this := BSONTimestamp{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BSONTimestamp) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BSONTimestamp) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BSONTimestamp) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *BSONTimestamp) SetDate(v time.Time) {
	o.Date = &v
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *BSONTimestamp) GetIncrement() int {
	if o == nil || IsNil(o.Increment) {
		var ret int
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BSONTimestamp) GetIncrementOk() (*int, bool) {
	if o == nil || IsNil(o.Increment) {
		return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *BSONTimestamp) HasIncrement() bool {
	if o != nil && !IsNil(o.Increment) {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given int and assigns it to the Increment field.
func (o *BSONTimestamp) SetIncrement(v int) {
	o.Increment = &v
}

func (o BSONTimestamp) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o BSONTimestamp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableBSONTimestamp struct {
	value *BSONTimestamp
	isSet bool
}

func (v NullableBSONTimestamp) Get() *BSONTimestamp {
	return v.value
}

func (v *NullableBSONTimestamp) Set(val *BSONTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableBSONTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableBSONTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBSONTimestamp(val *BSONTimestamp) *NullableBSONTimestamp {
	return &NullableBSONTimestamp{value: val, isSet: true}
}

func (v NullableBSONTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBSONTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
