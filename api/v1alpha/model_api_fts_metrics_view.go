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

// checks if the ApiFTSMetricsView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiFTSMetricsView{}

// ApiFTSMetricsView struct for ApiFTSMetricsView
type ApiFTSMetricsView struct {
	// Unique 24-hexadecimal digit string that identifies the project.
	GroupId string `json:"groupId"`
	// List that contains all host compute, memory, and storage utilization dedicated to Atlas Search when MongoDB Atlas received this request.
	HardwareMetrics []ApiFTSMetricView `json:"hardwareMetrics,omitempty"`
	// List that contains all performance and utilization measurements that Atlas Search index performed by the time MongoDB Atlas received this request.
	IndexMetrics []ApiFTSMetricView `json:"indexMetrics,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Hostname and port that identifies the process.
	ProcessId string `json:"processId"`
	// List that contains all available Atlas Search status metrics when MongoDB Atlas received this request.
	StatusMetrics []ApiFTSMetricView `json:"statusMetrics,omitempty"`
}

// NewApiFTSMetricsView instantiates a new ApiFTSMetricsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFTSMetricsView() *ApiFTSMetricsView {
	this := ApiFTSMetricsView{}
	return &this
}

// NewApiFTSMetricsViewWithDefaults instantiates a new ApiFTSMetricsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFTSMetricsViewWithDefaults() *ApiFTSMetricsView {
	this := ApiFTSMetricsView{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ApiFTSMetricsView) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ApiFTSMetricsView) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ApiFTSMetricsView) SetGroupId(v string) {
	o.GroupId = v
}

// GetHardwareMetrics returns the HardwareMetrics field value if set, zero value otherwise.
func (o *ApiFTSMetricsView) GetHardwareMetrics() []ApiFTSMetricView {
	if o == nil || IsNil(o.HardwareMetrics) {
		var ret []ApiFTSMetricView
		return ret
	}
	return o.HardwareMetrics
}

// GetHardwareMetricsOk returns a tuple with the HardwareMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFTSMetricsView) GetHardwareMetricsOk() ([]ApiFTSMetricView, bool) {
	if o == nil || IsNil(o.HardwareMetrics) {
		return nil, false
	}
	return o.HardwareMetrics, true
}

// HasHardwareMetrics returns a boolean if a field has been set.
func (o *ApiFTSMetricsView) HasHardwareMetrics() bool {
	if o != nil && !IsNil(o.HardwareMetrics) {
		return true
	}

	return false
}

// SetHardwareMetrics gets a reference to the given []ApiFTSMetricView and assigns it to the HardwareMetrics field.
func (o *ApiFTSMetricsView) SetHardwareMetrics(v []ApiFTSMetricView) {
	o.HardwareMetrics = v
}

// GetIndexMetrics returns the IndexMetrics field value if set, zero value otherwise.
func (o *ApiFTSMetricsView) GetIndexMetrics() []ApiFTSMetricView {
	if o == nil || IsNil(o.IndexMetrics) {
		var ret []ApiFTSMetricView
		return ret
	}
	return o.IndexMetrics
}

// GetIndexMetricsOk returns a tuple with the IndexMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFTSMetricsView) GetIndexMetricsOk() ([]ApiFTSMetricView, bool) {
	if o == nil || IsNil(o.IndexMetrics) {
		return nil, false
	}
	return o.IndexMetrics, true
}

// HasIndexMetrics returns a boolean if a field has been set.
func (o *ApiFTSMetricsView) HasIndexMetrics() bool {
	if o != nil && !IsNil(o.IndexMetrics) {
		return true
	}

	return false
}

// SetIndexMetrics gets a reference to the given []ApiFTSMetricView and assigns it to the IndexMetrics field.
func (o *ApiFTSMetricsView) SetIndexMetrics(v []ApiFTSMetricView) {
	o.IndexMetrics = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiFTSMetricsView) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFTSMetricsView) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiFTSMetricsView) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ApiFTSMetricsView) SetLinks(v []Link) {
	o.Links = v
}

// GetProcessId returns the ProcessId field value
func (o *ApiFTSMetricsView) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *ApiFTSMetricsView) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *ApiFTSMetricsView) SetProcessId(v string) {
	o.ProcessId = v
}

// GetStatusMetrics returns the StatusMetrics field value if set, zero value otherwise.
func (o *ApiFTSMetricsView) GetStatusMetrics() []ApiFTSMetricView {
	if o == nil || IsNil(o.StatusMetrics) {
		var ret []ApiFTSMetricView
		return ret
	}
	return o.StatusMetrics
}

// GetStatusMetricsOk returns a tuple with the StatusMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFTSMetricsView) GetStatusMetricsOk() ([]ApiFTSMetricView, bool) {
	if o == nil || IsNil(o.StatusMetrics) {
		return nil, false
	}
	return o.StatusMetrics, true
}

// HasStatusMetrics returns a boolean if a field has been set.
func (o *ApiFTSMetricsView) HasStatusMetrics() bool {
	if o != nil && !IsNil(o.StatusMetrics) {
		return true
	}

	return false
}

// SetStatusMetrics gets a reference to the given []ApiFTSMetricView and assigns it to the StatusMetrics field.
func (o *ApiFTSMetricsView) SetStatusMetrics(v []ApiFTSMetricView) {
	o.StatusMetrics = v
}

func (o ApiFTSMetricsView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiFTSMetricsView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: groupId is readOnly
	// skip: hardwareMetrics is readOnly
	// skip: indexMetrics is readOnly
	// skip: links is readOnly
	// skip: processId is readOnly
	// skip: statusMetrics is readOnly
	return toSerialize, nil
}

type NullableApiFTSMetricsView struct {
	value *ApiFTSMetricsView
	isSet bool
}

func (v NullableApiFTSMetricsView) Get() *ApiFTSMetricsView {
	return v.value
}

func (v *NullableApiFTSMetricsView) Set(val *ApiFTSMetricsView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFTSMetricsView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFTSMetricsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFTSMetricsView(val *ApiFTSMetricsView) *NullableApiFTSMetricsView {
	return &NullableApiFTSMetricsView{value: val, isSet: true}
}

func (v NullableApiFTSMetricsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFTSMetricsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


