/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the IngestionPipeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestionPipeline{}

// IngestionPipeline Details of a Data Lake Pipeline.
type IngestionPipeline struct {
	// Unique 24-hexadecimal digit string that identifies the Data Lake Pipeline.
	Id *string `json:"_id,omitempty"`
	// Timestamp that indicates when the Data Lake Pipeline was created.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the group.
	GroupId *string `json:"groupId,omitempty"`
	// Timestamp that indicates the last time that the Data Lake Pipeline was updated.
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
	// Name of this Data Lake Pipeline.
	Name *string `json:"name,omitempty"`
	Sink *IngestionSink `json:"sink,omitempty"`
	Source *IngestionSource `json:"source,omitempty"`
	// State of this Data Lake Pipeline.
	State *string `json:"state,omitempty"`
	// Fields to be excluded for this Data Lake Pipeline.
	Transformations []FieldTransformation `json:"transformations,omitempty"`
}

// NewIngestionPipeline instantiates a new IngestionPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionPipeline() *IngestionPipeline {
	this := IngestionPipeline{}
	return &this
}

// NewIngestionPipelineWithDefaults instantiates a new IngestionPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionPipelineWithDefaults() *IngestionPipeline {
	this := IngestionPipeline{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IngestionPipeline) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IngestionPipeline) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IngestionPipeline) SetId(v string) {
	o.Id = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *IngestionPipeline) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *IngestionPipeline) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *IngestionPipeline) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *IngestionPipeline) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *IngestionPipeline) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *IngestionPipeline) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *IngestionPipeline) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *IngestionPipeline) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *IngestionPipeline) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IngestionPipeline) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IngestionPipeline) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IngestionPipeline) SetName(v string) {
	o.Name = &v
}

// GetSink returns the Sink field value if set, zero value otherwise.
func (o *IngestionPipeline) GetSink() IngestionSink {
	if o == nil || IsNil(o.Sink) {
		var ret IngestionSink
		return ret
	}
	return *o.Sink
}

// GetSinkOk returns a tuple with the Sink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetSinkOk() (*IngestionSink, bool) {
	if o == nil || IsNil(o.Sink) {
		return nil, false
	}
	return o.Sink, true
}

// HasSink returns a boolean if a field has been set.
func (o *IngestionPipeline) HasSink() bool {
	if o != nil && !IsNil(o.Sink) {
		return true
	}

	return false
}

// SetSink gets a reference to the given IngestionSink and assigns it to the Sink field.
func (o *IngestionPipeline) SetSink(v IngestionSink) {
	o.Sink = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *IngestionPipeline) GetSource() IngestionSource {
	if o == nil || IsNil(o.Source) {
		var ret IngestionSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetSourceOk() (*IngestionSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *IngestionPipeline) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given IngestionSource and assigns it to the Source field.
func (o *IngestionPipeline) SetSource(v IngestionSource) {
	o.Source = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IngestionPipeline) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IngestionPipeline) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IngestionPipeline) SetState(v string) {
	o.State = &v
}

// GetTransformations returns the Transformations field value if set, zero value otherwise.
func (o *IngestionPipeline) GetTransformations() []FieldTransformation {
	if o == nil || IsNil(o.Transformations) {
		var ret []FieldTransformation
		return ret
	}
	return o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipeline) GetTransformationsOk() ([]FieldTransformation, bool) {
	if o == nil || IsNil(o.Transformations) {
		return nil, false
	}
	return o.Transformations, true
}

// HasTransformations returns a boolean if a field has been set.
func (o *IngestionPipeline) HasTransformations() bool {
	if o != nil && !IsNil(o.Transformations) {
		return true
	}

	return false
}

// SetTransformations gets a reference to the given []FieldTransformation and assigns it to the Transformations field.
func (o *IngestionPipeline) SetTransformations(v []FieldTransformation) {
	o.Transformations = v
}

func (o IngestionPipeline) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o IngestionPipeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Sink) {
		toSerialize["sink"] = o.Sink
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Transformations) {
		toSerialize["transformations"] = o.Transformations
	}
	return toSerialize, nil
}

type NullableIngestionPipeline struct {
	value *IngestionPipeline
	isSet bool
}

func (v NullableIngestionPipeline) Get() *IngestionPipeline {
	return v.value
}

func (v *NullableIngestionPipeline) Set(val *IngestionPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionPipeline(val *IngestionPipeline) *NullableIngestionPipeline {
	return &NullableIngestionPipeline{value: val, isSet: true}
}

func (v NullableIngestionPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


