// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the FTSMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTSMetrics{}

// FTSMetrics struct for FTSMetrics
type FTSMetrics struct {
	// Unique 24-hexadecimal digit string that identifies the project.
	GroupId string `json:"groupId"`
	// List that contains all host compute, memory, and storage utilization dedicated to Atlas Search when MongoDB Atlas received this request.
	HardwareMetrics []FTSMetric `json:"hardwareMetrics,omitempty"`
	// List that contains all performance and utilization measurements that Atlas Search index performed by the time MongoDB Atlas received this request.
	IndexMetrics []FTSMetric `json:"indexMetrics,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Hostname and port that identifies the process.
	ProcessId string `json:"processId"`
	// List that contains all available Atlas Search status metrics when MongoDB Atlas received this request.
	StatusMetrics []FTSMetric `json:"statusMetrics,omitempty"`
}

// NewFTSMetrics instantiates a new FTSMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTSMetrics(groupId string, processId string) *FTSMetrics {
	this := FTSMetrics{}
	this.GroupId = groupId
	this.ProcessId = processId
	return &this
}

// NewFTSMetricsWithDefaults instantiates a new FTSMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTSMetricsWithDefaults() *FTSMetrics {
	this := FTSMetrics{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *FTSMetrics) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *FTSMetrics) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *FTSMetrics) SetGroupId(v string) {
	o.GroupId = v
}

// GetHardwareMetrics returns the HardwareMetrics field value if set, zero value otherwise.
func (o *FTSMetrics) GetHardwareMetrics() []FTSMetric {
	if o == nil || IsNil(o.HardwareMetrics) {
		var ret []FTSMetric
		return ret
	}
	return o.HardwareMetrics
}

// GetHardwareMetricsOk returns a tuple with the HardwareMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSMetrics) GetHardwareMetricsOk() ([]FTSMetric, bool) {
	if o == nil || IsNil(o.HardwareMetrics) {
		return nil, false
	}
	return o.HardwareMetrics, true
}

// HasHardwareMetrics returns a boolean if a field has been set.
func (o *FTSMetrics) HasHardwareMetrics() bool {
	if o != nil && !IsNil(o.HardwareMetrics) {
		return true
	}

	return false
}

// SetHardwareMetrics gets a reference to the given []FTSMetric and assigns it to the HardwareMetrics field.
func (o *FTSMetrics) SetHardwareMetrics(v []FTSMetric) {
	o.HardwareMetrics = v
}

// GetIndexMetrics returns the IndexMetrics field value if set, zero value otherwise.
func (o *FTSMetrics) GetIndexMetrics() []FTSMetric {
	if o == nil || IsNil(o.IndexMetrics) {
		var ret []FTSMetric
		return ret
	}
	return o.IndexMetrics
}

// GetIndexMetricsOk returns a tuple with the IndexMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSMetrics) GetIndexMetricsOk() ([]FTSMetric, bool) {
	if o == nil || IsNil(o.IndexMetrics) {
		return nil, false
	}
	return o.IndexMetrics, true
}

// HasIndexMetrics returns a boolean if a field has been set.
func (o *FTSMetrics) HasIndexMetrics() bool {
	if o != nil && !IsNil(o.IndexMetrics) {
		return true
	}

	return false
}

// SetIndexMetrics gets a reference to the given []FTSMetric and assigns it to the IndexMetrics field.
func (o *FTSMetrics) SetIndexMetrics(v []FTSMetric) {
	o.IndexMetrics = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FTSMetrics) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSMetrics) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FTSMetrics) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *FTSMetrics) SetLinks(v []Link) {
	o.Links = v
}

// GetProcessId returns the ProcessId field value
func (o *FTSMetrics) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *FTSMetrics) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *FTSMetrics) SetProcessId(v string) {
	o.ProcessId = v
}

// GetStatusMetrics returns the StatusMetrics field value if set, zero value otherwise.
func (o *FTSMetrics) GetStatusMetrics() []FTSMetric {
	if o == nil || IsNil(o.StatusMetrics) {
		var ret []FTSMetric
		return ret
	}
	return o.StatusMetrics
}

// GetStatusMetricsOk returns a tuple with the StatusMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSMetrics) GetStatusMetricsOk() ([]FTSMetric, bool) {
	if o == nil || IsNil(o.StatusMetrics) {
		return nil, false
	}
	return o.StatusMetrics, true
}

// HasStatusMetrics returns a boolean if a field has been set.
func (o *FTSMetrics) HasStatusMetrics() bool {
	if o != nil && !IsNil(o.StatusMetrics) {
		return true
	}

	return false
}

// SetStatusMetrics gets a reference to the given []FTSMetric and assigns it to the StatusMetrics field.
func (o *FTSMetrics) SetStatusMetrics(v []FTSMetric) {
	o.StatusMetrics = v
}

func (o FTSMetrics) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o FTSMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableFTSMetrics struct {
	value *FTSMetrics
	isSet bool
}

func (v NullableFTSMetrics) Get() *FTSMetrics {
	return v.value
}

func (v *NullableFTSMetrics) Set(val *FTSMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSMetrics(val *FTSMetrics) *NullableFTSMetrics {
	return &NullableFTSMetrics{value: val, isSet: true}
}

func (v NullableFTSMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
