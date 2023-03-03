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

// checks if the ApiPagerDutyView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPagerDutyView{}

// ApiPagerDutyView Details to integrate one PagerDuty account with one MongoDB Cloud project.
type ApiPagerDutyView struct {
	// PagerDuty region that indicates the API Uniform Resource Locator (URL) to use.
	Region *string `json:"region,omitempty"`
	// Service key associated with your PagerDuty account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ServiceKey string `json:"serviceKey"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewApiPagerDutyView instantiates a new ApiPagerDutyView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPagerDutyView(serviceKey string) *ApiPagerDutyView {
	this := ApiPagerDutyView{}
	this.ServiceKey = serviceKey
	return &this
}

// NewApiPagerDutyViewWithDefaults instantiates a new ApiPagerDutyView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPagerDutyViewWithDefaults() *ApiPagerDutyView {
	this := ApiPagerDutyView{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ApiPagerDutyView) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPagerDutyView) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ApiPagerDutyView) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ApiPagerDutyView) SetRegion(v string) {
	o.Region = &v
}

// GetServiceKey returns the ServiceKey field value
func (o *ApiPagerDutyView) GetServiceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value
// and a boolean to check if the value has been set.
func (o *ApiPagerDutyView) GetServiceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceKey, true
}

// SetServiceKey sets field value
func (o *ApiPagerDutyView) SetServiceKey(v string) {
	o.ServiceKey = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiPagerDutyView) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPagerDutyView) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiPagerDutyView) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiPagerDutyView) SetType(v string) {
	o.Type = &v
}

func (o ApiPagerDutyView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPagerDutyView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	toSerialize["serviceKey"] = o.ServiceKey
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableApiPagerDutyView struct {
	value *ApiPagerDutyView
	isSet bool
}

func (v NullableApiPagerDutyView) Get() *ApiPagerDutyView {
	return v.value
}

func (v *NullableApiPagerDutyView) Set(val *ApiPagerDutyView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPagerDutyView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPagerDutyView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPagerDutyView(val *ApiPagerDutyView) *NullableApiPagerDutyView {
	return &NullableApiPagerDutyView{value: val, isSet: true}
}

func (v NullableApiPagerDutyView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPagerDutyView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


