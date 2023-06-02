// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the OpsGenie type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsGenie{}

// OpsGenie Details to integrate one Opsgenie account with one MongoDB Cloud project.
type OpsGenie struct {
	// Key that allows MongoDB Cloud to access your Opsgenie account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ApiKey string `json:"apiKey"`
	// Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie API.
	Region *string `json:"region,omitempty"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewOpsGenie instantiates a new OpsGenie object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsGenie(apiKey string) *OpsGenie {
	this := OpsGenie{}
	this.ApiKey = apiKey
	var region string = "US"
	this.Region = &region
	return &this
}

// NewOpsGenieWithDefaults instantiates a new OpsGenie object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsGenieWithDefaults() *OpsGenie {
	this := OpsGenie{}
	var region string = "US"
	this.Region = &region
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *OpsGenie) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *OpsGenie) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *OpsGenie) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *OpsGenie) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenie) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *OpsGenie) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *OpsGenie) SetRegion(v string) {
	o.Region = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OpsGenie) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenie) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OpsGenie) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OpsGenie) SetType(v string) {
	o.Type = &v
}

func (o OpsGenie) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o OpsGenie) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableOpsGenie struct {
	value *OpsGenie
	isSet bool
}

func (v NullableOpsGenie) Get() *OpsGenie {
	return v.value
}

func (v *NullableOpsGenie) Set(val *OpsGenie) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsGenie) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsGenie) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsGenie(val *OpsGenie) *NullableOpsGenie {
	return &NullableOpsGenie{value: val, isSet: true}
}

func (v NullableOpsGenie) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsGenie) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
