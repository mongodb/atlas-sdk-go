// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PagerDuty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagerDuty{}

// PagerDuty Details to integrate one PagerDuty account with one MongoDB Cloud project.
type PagerDuty struct {
	// PagerDuty region that indicates the API Uniform Resource Locator (URL) to use.
	Region *string `json:"region,omitempty"`
	// Service key associated with your PagerDuty account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ServiceKey string `json:"serviceKey"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewPagerDuty instantiates a new PagerDuty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagerDuty(serviceKey string) *PagerDuty {
	this := PagerDuty{}
	this.ServiceKey = serviceKey
	return &this
}

// NewPagerDutyWithDefaults instantiates a new PagerDuty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagerDutyWithDefaults() *PagerDuty {
	this := PagerDuty{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PagerDuty) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagerDuty) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PagerDuty) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PagerDuty) SetRegion(v string) {
	o.Region = &v
}

// GetServiceKey returns the ServiceKey field value
func (o *PagerDuty) GetServiceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value
// and a boolean to check if the value has been set.
func (o *PagerDuty) GetServiceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceKey, true
}

// SetServiceKey sets field value
func (o *PagerDuty) SetServiceKey(v string) {
	o.ServiceKey = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PagerDuty) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagerDuty) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PagerDuty) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PagerDuty) SetType(v string) {
	o.Type = &v
}

func (o PagerDuty) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PagerDuty) ToMap() (map[string]interface{}, error) {
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

type NullablePagerDuty struct {
	value *PagerDuty
	isSet bool
}

func (v NullablePagerDuty) Get() *PagerDuty {
	return v.value
}

func (v *NullablePagerDuty) Set(val *PagerDuty) {
	v.value = val
	v.isSet = true
}

func (v NullablePagerDuty) IsSet() bool {
	return v.isSet
}

func (v *NullablePagerDuty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagerDuty(val *PagerDuty) *NullablePagerDuty {
	return &NullablePagerDuty{value: val, isSet: true}
}

func (v NullablePagerDuty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagerDuty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
