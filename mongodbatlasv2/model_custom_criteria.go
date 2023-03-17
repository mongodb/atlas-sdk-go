/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the CustomCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomCriteria{}

// CustomCriteria **CUSTOM criteria.type**.
type CustomCriteria struct {
	// MongoDB find query that selects documents to archive. The specified query follows the syntax of the `db.collection.find(query)` command. This query can't use the empty document (`{}`) to return all documents. Set this parameter when **\"criteria.type\" : \"CUSTOM\"**.
	Query string `json:"query"`
	// Means by which MongoDB Cloud selects data to archive. Data can be chosen using the age of the data or a MongoDB query. **DATE** selects documents to archive based on a date. **CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **CUSTOM** when `\"collectionType\": \"TIMESERIES\"`.
	Type *string `json:"type,omitempty"`
}

// NewCustomCriteria instantiates a new CustomCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomCriteria(query string) *CustomCriteria {
	this := CustomCriteria{}
	this.Query = query
	return &this
}

// NewCustomCriteriaWithDefaults instantiates a new CustomCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomCriteriaWithDefaults() *CustomCriteria {
	this := CustomCriteria{}
	return &this
}

// GetQuery returns the Query field value
func (o *CustomCriteria) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CustomCriteria) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CustomCriteria) SetQuery(v string) {
	o.Query = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomCriteria) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCriteria) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomCriteria) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomCriteria) SetType(v string) {
	o.Type = &v
}

func (o CustomCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCustomCriteria struct {
	value *CustomCriteria
	isSet bool
}

func (v NullableCustomCriteria) Get() *CustomCriteria {
	return v.value
}

func (v *NullableCustomCriteria) Set(val *CustomCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomCriteria(val *CustomCriteria) *NullableCustomCriteria {
	return &NullableCustomCriteria{value: val, isSet: true}
}

func (v NullableCustomCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


