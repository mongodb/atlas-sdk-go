// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the SystemStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemStatus{}

// SystemStatus struct for SystemStatus
type SystemStatus struct {
	ApiKey NullableKey `json:"apiKey"`
	// Human-readable label that identifies the service from which you requested this response.
	AppName string `json:"appName"`
	// Unique 40-hexadecimal digit hash that identifies the latest git commit merged for this application.
	Build string `json:"build"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Flag that indicates whether someone enabled throttling on this service.
	Throttling bool `json:"throttling"`
}

// NewSystemStatus instantiates a new SystemStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemStatus(apiKey NullableKey, appName string, build string, throttling bool) *SystemStatus {
	this := SystemStatus{}
	this.ApiKey = apiKey
	this.AppName = appName
	this.Build = build
	this.Throttling = throttling
	return &this
}

// NewSystemStatusWithDefaults instantiates a new SystemStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemStatusWithDefaults() *SystemStatus {
	this := SystemStatus{}
	return &this
}

// GetApiKey returns the ApiKey field value
// If the value is explicit nil, the zero value for Key will be returned
func (o *SystemStatus) GetApiKey() Key {
	if o == nil || o.ApiKey.Get() == nil {
		var ret Key
		return ret
	}

	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemStatus) GetApiKeyOk() (*Key, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// SetApiKey sets field value
func (o *SystemStatus) SetApiKey(v Key) {
	o.ApiKey.Set(&v)
}

// GetAppName returns the AppName field value
func (o *SystemStatus) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *SystemStatus) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *SystemStatus) SetAppName(v string) {
	o.AppName = v
}

// GetBuild returns the Build field value
func (o *SystemStatus) GetBuild() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Build
}

// GetBuildOk returns a tuple with the Build field value
// and a boolean to check if the value has been set.
func (o *SystemStatus) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Build, true
}

// SetBuild sets field value
func (o *SystemStatus) SetBuild(v string) {
	o.Build = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SystemStatus) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemStatus) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SystemStatus) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *SystemStatus) SetLinks(v []Link) {
	o.Links = v
}

// GetThrottling returns the Throttling field value
func (o *SystemStatus) GetThrottling() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Throttling
}

// GetThrottlingOk returns a tuple with the Throttling field value
// and a boolean to check if the value has been set.
func (o *SystemStatus) GetThrottlingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throttling, true
}

// SetThrottling sets field value
func (o *SystemStatus) SetThrottling(v bool) {
	o.Throttling = v
}

func (o SystemStatus) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o SystemStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey.Get()
	return toSerialize, nil
}

type NullableSystemStatus struct {
	value *SystemStatus
	isSet bool
}

func (v NullableSystemStatus) Get() *SystemStatus {
	return v.value
}

func (v *NullableSystemStatus) Set(val *SystemStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemStatus(val *SystemStatus) *NullableSystemStatus {
	return &NullableSystemStatus{value: val, isSet: true}
}

func (v NullableSystemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
