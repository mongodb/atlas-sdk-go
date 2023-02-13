/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// ApiMeasurementsIndexesView struct for ApiMeasurementsIndexesView
type ApiMeasurementsIndexesView struct {
	// Human-readable label that identifies the collection.
	CollectionName *string `json:"collectionName,omitempty"`
	// Human-readable label that identifies the database that the specified MongoDB process serves.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Date and time that specifies when to stop retrieving measurements. If you set **end**, you must set **start**. You can't set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	End *time.Time `json:"end,omitempty"`
	// Duration that specifies the interval between measurement data points. The parameter expresses its value in ISO 8601 timestamp format in UTC. If you set this parameter, you must set either **period** or **start** and **end**.
	Granularity *string `json:"granularity,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project. The project contains MongoDB processes that you want to return. The MongoDB process can be either the `mongod` or `mongos`.
	GroupId *string `json:"groupId,omitempty"`
	// List that contains the Atlas Search index identifiers.
	IndexIds []string `json:"indexIds,omitempty"`
	// List that contains the Atlas Search index stats measurements.
	IndexStatsMeasurements []ApiMeasurementView `json:"indexStatsMeasurements,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	ProcessId *string `json:"processId,omitempty"`
	// Date and time that specifies when to start retrieving measurements. If you set **start**, you must set **end**. You can't set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Start *time.Time `json:"start,omitempty"`
}

// NewApiMeasurementsIndexesView instantiates a new ApiMeasurementsIndexesView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMeasurementsIndexesView() *ApiMeasurementsIndexesView {
	this := ApiMeasurementsIndexesView{}
	return &this
}

// NewApiMeasurementsIndexesViewWithDefaults instantiates a new ApiMeasurementsIndexesView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMeasurementsIndexesViewWithDefaults() *ApiMeasurementsIndexesView {
	this := ApiMeasurementsIndexesView{}
	return &this
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetCollectionName() string {
	if o == nil || o.CollectionName == nil {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetCollectionNameOk() (*string, bool) {
	if o == nil || o.CollectionName == nil {
		return nil, false
	}
	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasCollectionName() bool {
	if o != nil && o.CollectionName != nil {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *ApiMeasurementsIndexesView) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasDatabaseName() bool {
	if o != nil && o.DatabaseName != nil {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *ApiMeasurementsIndexesView) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetEnd() time.Time {
	if o == nil || o.End == nil {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetEndOk() (*time.Time, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *ApiMeasurementsIndexesView) SetEnd(v time.Time) {
	o.End = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetGranularity() string {
	if o == nil || o.Granularity == nil {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetGranularityOk() (*string, bool) {
	if o == nil || o.Granularity == nil {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasGranularity() bool {
	if o != nil && o.Granularity != nil {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *ApiMeasurementsIndexesView) SetGranularity(v string) {
	o.Granularity = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiMeasurementsIndexesView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetIndexIds returns the IndexIds field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetIndexIds() []string {
	if o == nil || o.IndexIds == nil {
		var ret []string
		return ret
	}
	return o.IndexIds
}

// GetIndexIdsOk returns a tuple with the IndexIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetIndexIdsOk() ([]string, bool) {
	if o == nil || o.IndexIds == nil {
		return nil, false
	}
	return o.IndexIds, true
}

// HasIndexIds returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasIndexIds() bool {
	if o != nil && o.IndexIds != nil {
		return true
	}

	return false
}

// SetIndexIds gets a reference to the given []string and assigns it to the IndexIds field.
func (o *ApiMeasurementsIndexesView) SetIndexIds(v []string) {
	o.IndexIds = v
}

// GetIndexStatsMeasurements returns the IndexStatsMeasurements field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetIndexStatsMeasurements() []ApiMeasurementView {
	if o == nil || o.IndexStatsMeasurements == nil {
		var ret []ApiMeasurementView
		return ret
	}
	return o.IndexStatsMeasurements
}

// GetIndexStatsMeasurementsOk returns a tuple with the IndexStatsMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetIndexStatsMeasurementsOk() ([]ApiMeasurementView, bool) {
	if o == nil || o.IndexStatsMeasurements == nil {
		return nil, false
	}
	return o.IndexStatsMeasurements, true
}

// HasIndexStatsMeasurements returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasIndexStatsMeasurements() bool {
	if o != nil && o.IndexStatsMeasurements != nil {
		return true
	}

	return false
}

// SetIndexStatsMeasurements gets a reference to the given []ApiMeasurementView and assigns it to the IndexStatsMeasurements field.
func (o *ApiMeasurementsIndexesView) SetIndexStatsMeasurements(v []ApiMeasurementView) {
	o.IndexStatsMeasurements = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ApiMeasurementsIndexesView) SetLinks(v []Link) {
	o.Links = v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetProcessId() string {
	if o == nil || o.ProcessId == nil {
		var ret string
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetProcessIdOk() (*string, bool) {
	if o == nil || o.ProcessId == nil {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasProcessId() bool {
	if o != nil && o.ProcessId != nil {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given string and assigns it to the ProcessId field.
func (o *ApiMeasurementsIndexesView) SetProcessId(v string) {
	o.ProcessId = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ApiMeasurementsIndexesView) GetStart() time.Time {
	if o == nil || o.Start == nil {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsIndexesView) GetStartOk() (*time.Time, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ApiMeasurementsIndexesView) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *ApiMeasurementsIndexesView) SetStart(v time.Time) {
	o.Start = &v
}

func (o ApiMeasurementsIndexesView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollectionName != nil {
		toSerialize["collectionName"] = o.CollectionName
	}
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Granularity != nil {
		toSerialize["granularity"] = o.Granularity
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.IndexIds != nil {
		toSerialize["indexIds"] = o.IndexIds
	}
	if o.IndexStatsMeasurements != nil {
		toSerialize["indexStatsMeasurements"] = o.IndexStatsMeasurements
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.ProcessId != nil {
		toSerialize["processId"] = o.ProcessId
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableApiMeasurementsIndexesView struct {
	value *ApiMeasurementsIndexesView
	isSet bool
}

func (v NullableApiMeasurementsIndexesView) Get() *ApiMeasurementsIndexesView {
	return v.value
}

func (v *NullableApiMeasurementsIndexesView) Set(val *ApiMeasurementsIndexesView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMeasurementsIndexesView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMeasurementsIndexesView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMeasurementsIndexesView(val *ApiMeasurementsIndexesView) *NullableApiMeasurementsIndexesView {
	return &NullableApiMeasurementsIndexesView{value: val, isSet: true}
}

func (v NullableApiMeasurementsIndexesView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMeasurementsIndexesView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

