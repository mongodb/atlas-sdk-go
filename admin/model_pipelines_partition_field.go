// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PipelinesPartitionField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelinesPartitionField{}

// PipelinesPartitionField Partition Field in the Data Lake Storage provider for a Data Lake Pipeline.
type PipelinesPartitionField struct {
	// Human-readable label that identifies the field name used to partition data.
	FieldName string `json:"fieldName"`
	// Sequence in which MongoDB Cloud slices the collection data to create partitions. The resource expresses this sequence starting with zero.
	Order int `json:"order"`
}

// NewPipelinesPartitionField instantiates a new PipelinesPartitionField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelinesPartitionField(fieldName string, order int) *PipelinesPartitionField {
	this := PipelinesPartitionField{}
	this.FieldName = fieldName
	this.Order = order
	return &this
}

// NewPipelinesPartitionFieldWithDefaults instantiates a new PipelinesPartitionField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelinesPartitionFieldWithDefaults() *PipelinesPartitionField {
	this := PipelinesPartitionField{}
	var order int = 0
	this.Order = order
	return &this
}

// GetFieldName returns the FieldName field value
func (o *PipelinesPartitionField) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *PipelinesPartitionField) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *PipelinesPartitionField) SetFieldName(v string) {
	o.FieldName = v
}

// GetOrder returns the Order field value
func (o *PipelinesPartitionField) GetOrder() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *PipelinesPartitionField) GetOrderOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *PipelinesPartitionField) SetOrder(v int) {
	o.Order = v
}

func (o PipelinesPartitionField) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PipelinesPartitionField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullablePipelinesPartitionField struct {
	value *PipelinesPartitionField
	isSet bool
}

func (v NullablePipelinesPartitionField) Get() *PipelinesPartitionField {
	return v.value
}

func (v *NullablePipelinesPartitionField) Set(val *PipelinesPartitionField) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelinesPartitionField) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelinesPartitionField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelinesPartitionField(val *PipelinesPartitionField) *NullablePipelinesPartitionField {
	return &NullablePipelinesPartitionField{value: val, isSet: true}
}

func (v NullablePipelinesPartitionField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelinesPartitionField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
