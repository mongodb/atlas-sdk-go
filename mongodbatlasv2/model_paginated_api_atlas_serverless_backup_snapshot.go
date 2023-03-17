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

// checks if the PaginatedApiAtlasServerlessBackupSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedApiAtlasServerlessBackupSnapshot{}

// PaginatedApiAtlasServerlessBackupSnapshot struct for PaginatedApiAtlasServerlessBackupSnapshot
type PaginatedApiAtlasServerlessBackupSnapshot struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ServerlessBackupSnapshot `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedApiAtlasServerlessBackupSnapshot instantiates a new PaginatedApiAtlasServerlessBackupSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiAtlasServerlessBackupSnapshot() *PaginatedApiAtlasServerlessBackupSnapshot {
	this := PaginatedApiAtlasServerlessBackupSnapshot{}
	return &this
}

// NewPaginatedApiAtlasServerlessBackupSnapshotWithDefaults instantiates a new PaginatedApiAtlasServerlessBackupSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiAtlasServerlessBackupSnapshotWithDefaults() *PaginatedApiAtlasServerlessBackupSnapshot {
	this := PaginatedApiAtlasServerlessBackupSnapshot{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) GetResults() []ServerlessBackupSnapshot {
	if o == nil || IsNil(o.Results) {
		var ret []ServerlessBackupSnapshot
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) GetResultsOk() ([]ServerlessBackupSnapshot, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ServerlessBackupSnapshot and assigns it to the Results field.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) SetResults(v []ServerlessBackupSnapshot) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedApiAtlasServerlessBackupSnapshot) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedApiAtlasServerlessBackupSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedApiAtlasServerlessBackupSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: links is readOnly
	// skip: results is readOnly
	// skip: totalCount is readOnly
	return toSerialize, nil
}

type NullablePaginatedApiAtlasServerlessBackupSnapshot struct {
	value *PaginatedApiAtlasServerlessBackupSnapshot
	isSet bool
}

func (v NullablePaginatedApiAtlasServerlessBackupSnapshot) Get() *PaginatedApiAtlasServerlessBackupSnapshot {
	return v.value
}

func (v *NullablePaginatedApiAtlasServerlessBackupSnapshot) Set(val *PaginatedApiAtlasServerlessBackupSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApiAtlasServerlessBackupSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApiAtlasServerlessBackupSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApiAtlasServerlessBackupSnapshot(val *PaginatedApiAtlasServerlessBackupSnapshot) *NullablePaginatedApiAtlasServerlessBackupSnapshot {
	return &NullablePaginatedApiAtlasServerlessBackupSnapshot{value: val, isSet: true}
}

func (v NullablePaginatedApiAtlasServerlessBackupSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApiAtlasServerlessBackupSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


