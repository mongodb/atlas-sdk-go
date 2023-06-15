// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the Datadog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Datadog{}

// Datadog Details to integrate one Datadog account with one MongoDB Cloud project.
type Datadog struct {
	// Key that allows MongoDB Cloud to access your Datadog account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ApiKey string `json:"apiKey"`
	// Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API.  To learn more about Datadog's regions, see <a href=\"https://docs.datadoghq.com/getting_started/site/\" target=\"_blank\" rel=\"noopener noreferrer\">Datadog Sites</a>.
	Region *string `json:"region,omitempty"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewDatadog instantiates a new Datadog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadog(apiKey string) *Datadog {
	this := Datadog{}
	this.ApiKey = apiKey
	return &this
}

// NewDatadogWithDefaults instantiates a new Datadog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadogWithDefaults() *Datadog {
	this := Datadog{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *Datadog) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *Datadog) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *Datadog) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Datadog) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datadog) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Datadog) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Datadog) SetRegion(v string) {
	o.Region = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Datadog) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datadog) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Datadog) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Datadog) SetType(v string) {
	o.Type = &v
}

func (o Datadog) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Datadog) ToMap() (map[string]interface{}, error) {
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

type NullableDatadog struct {
	value *Datadog
	isSet bool
}

func (v NullableDatadog) Get() *Datadog {
	return v.value
}

func (v *NullableDatadog) Set(val *Datadog) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadog) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadog(val *Datadog) *NullableDatadog {
	return &NullableDatadog{value: val, isSet: true}
}

func (v NullableDatadog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
