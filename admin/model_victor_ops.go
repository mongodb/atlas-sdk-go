// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the VictorOps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VictorOps{}

// VictorOps Details to integrate one Splunk On-Call account with one MongoDB Cloud project.
type VictorOps struct {
	// Key that allows MongoDB Cloud to access your VictorOps account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ApiKey string `json:"apiKey"`
	// Routing key associated with your Splunk On-Call account.
	RoutingKey *string `json:"routingKey,omitempty"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewVictorOps instantiates a new VictorOps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVictorOps(apiKey string) *VictorOps {
	this := VictorOps{}
	this.ApiKey = apiKey
	return &this
}

// NewVictorOpsWithDefaults instantiates a new VictorOps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVictorOpsWithDefaults() *VictorOps {
	this := VictorOps{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *VictorOps) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *VictorOps) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *VictorOps) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRoutingKey returns the RoutingKey field value if set, zero value otherwise.
func (o *VictorOps) GetRoutingKey() string {
	if o == nil || IsNil(o.RoutingKey) {
		var ret string
		return ret
	}
	return *o.RoutingKey
}

// GetRoutingKeyOk returns a tuple with the RoutingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOps) GetRoutingKeyOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingKey) {
		return nil, false
	}
	return o.RoutingKey, true
}

// HasRoutingKey returns a boolean if a field has been set.
func (o *VictorOps) HasRoutingKey() bool {
	if o != nil && !IsNil(o.RoutingKey) {
		return true
	}

	return false
}

// SetRoutingKey gets a reference to the given string and assigns it to the RoutingKey field.
func (o *VictorOps) SetRoutingKey(v string) {
	o.RoutingKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VictorOps) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOps) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VictorOps) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VictorOps) SetType(v string) {
	o.Type = &v
}

func (o VictorOps) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o VictorOps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey
	if !IsNil(o.RoutingKey) {
		toSerialize["routingKey"] = o.RoutingKey
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableVictorOps struct {
	value *VictorOps
	isSet bool
}

func (v NullableVictorOps) Get() *VictorOps {
	return v.value
}

func (v *NullableVictorOps) Set(val *VictorOps) {
	v.value = val
	v.isSet = true
}

func (v NullableVictorOps) IsSet() bool {
	return v.isSet
}

func (v *NullableVictorOps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVictorOps(val *VictorOps) *NullableVictorOps {
	return &NullableVictorOps{value: val, isSet: true}
}

func (v NullableVictorOps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVictorOps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
