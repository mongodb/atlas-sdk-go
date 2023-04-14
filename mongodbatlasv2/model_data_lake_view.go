/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the DataLakeView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeView{}

// DataLakeView An aggregation pipeline that applies to the collection.
type DataLakeView struct {
	// Human-readable label that identifies the view, which corresponds to an aggregation pipeline on a collection.
	Name *string `json:"name,omitempty"`
	// Aggregation pipeline stages to apply to the source collection.
	Pipeline *string `json:"pipeline,omitempty"`
	// Human-readable label that identifies the source collection for the view.
	Source *string `json:"source,omitempty"`
}

// NewDataLakeView instantiates a new DataLakeView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeView() *DataLakeView {
	this := DataLakeView{}
	return &this
}

// NewDataLakeViewWithDefaults instantiates a new DataLakeView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeViewWithDefaults() *DataLakeView {
	this := DataLakeView{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataLakeView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataLakeView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataLakeView) SetName(v string) {
	o.Name = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *DataLakeView) GetPipeline() string {
	if o == nil || IsNil(o.Pipeline) {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeView) GetPipelineOk() (*string, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *DataLakeView) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *DataLakeView) SetPipeline(v string) {
	o.Pipeline = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DataLakeView) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeView) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DataLakeView) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DataLakeView) SetSource(v string) {
	o.Source = &v
}

func (o DataLakeView) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pipeline) {
		toSerialize["pipeline"] = o.Pipeline
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableDataLakeView struct {
	value *DataLakeView
	isSet bool
}

func (v NullableDataLakeView) Get() *DataLakeView {
	return v.value
}

func (v *NullableDataLakeView) Set(val *DataLakeView) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeView) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeView(val *DataLakeView) *NullableDataLakeView {
	return &NullableDataLakeView{value: val, isSet: true}
}

func (v NullableDataLakeView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


