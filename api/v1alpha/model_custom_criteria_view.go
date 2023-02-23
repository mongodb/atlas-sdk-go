/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the CustomCriteriaView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomCriteriaView{}

// CustomCriteriaView **CUSTOM criteria.type**.
type CustomCriteriaView struct {
	// MongoDB find query that selects documents to archive. The specified query follows the syntax of the `db.collection.find(query)` command. This query can't use the empty document (`{}`) to return all documents. Set this parameter when **\"criteria.type\" : \"CUSTOM\"**.
	Query string `json:"query"`
}

// NewCustomCriteriaView instantiates a new CustomCriteriaView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomCriteriaView() *CustomCriteriaView {
	this := CustomCriteriaView{}
	return &this
}

// NewCustomCriteriaViewWithDefaults instantiates a new CustomCriteriaView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomCriteriaViewWithDefaults() *CustomCriteriaView {
	this := CustomCriteriaView{}
	return &this
}

// GetQuery returns the Query field value
func (o *CustomCriteriaView) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CustomCriteriaView) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CustomCriteriaView) SetQuery(v string) {
	o.Query = v
}

func (o CustomCriteriaView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomCriteriaView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

type NullableCustomCriteriaView struct {
	value *CustomCriteriaView
	isSet bool
}

func (v NullableCustomCriteriaView) Get() *CustomCriteriaView {
	return v.value
}

func (v *NullableCustomCriteriaView) Set(val *CustomCriteriaView) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomCriteriaView) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomCriteriaView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomCriteriaView(val *CustomCriteriaView) *NullableCustomCriteriaView {
	return &NullableCustomCriteriaView{value: val, isSet: true}
}

func (v NullableCustomCriteriaView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomCriteriaView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


