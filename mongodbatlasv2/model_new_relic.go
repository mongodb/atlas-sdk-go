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

// checks if the NewRelic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewRelic{}

// NewRelic Details to integrate one New Relic account with one MongoDB Cloud project.  ***IMPORTANT**: Effective Wednesday, June 16th, 2021, New Relic no longer supports the plugin-based integration with MongoDB. We do not recommend that you sign up for the plugin-based integration.  To learn more, see the <a href=\"https://discuss.newrelic.com/t/new-relic-plugin-eol-wednesday-june-16th-2021/127267\" target=\"_blank\">New Relic Plugin EOL Statement</a>. Consider configuring an alternative monitoring integration before June 16th to maintain visibility into your MongoDB deployments.
type NewRelic struct {
	// Unique 40-hexadecimal digit string that identifies your New Relic account.
	AccountId string `json:"accountId"`
	// Unique 40-hexadecimal digit string that identifies your New Relic license.  **IMPORTANT**: Effective Wednesday, June 16th, 2021, New Relic no longer supports the plugin-based integration with MongoDB. We do not recommend that you sign up for the plugin-based integration. To learn more, see the <a href=\"https://discuss.newrelic.com/t/new-relic-plugin-eol-wednesday-june-16th-2021/127267\" target=\"_blank\">New Relic Plugin EOL Statement</a> Consider configuring an alternative monitoring integration before June 16th to maintain visibility into your MongoDB deployments.
	LicenseKey string `json:"licenseKey"`
	// Query key used to access your New Relic account.
	ReadToken string `json:"readToken"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
	// Insert key associated with your New Relic account.
	WriteToken string `json:"writeToken"`
}

// NewNewRelic instantiates a new NewRelic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewRelic(accountId string, licenseKey string, readToken string, writeToken string) *NewRelic {
	this := NewRelic{}
	this.AccountId = accountId
	this.LicenseKey = licenseKey
	this.ReadToken = readToken
	this.WriteToken = writeToken
	return &this
}

// NewNewRelicWithDefaults instantiates a new NewRelic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewRelicWithDefaults() *NewRelic {
	this := NewRelic{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NewRelic) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NewRelic) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NewRelic) SetAccountId(v string) {
	o.AccountId = v
}

// GetLicenseKey returns the LicenseKey field value
func (o *NewRelic) GetLicenseKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value
// and a boolean to check if the value has been set.
func (o *NewRelic) GetLicenseKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseKey, true
}

// SetLicenseKey sets field value
func (o *NewRelic) SetLicenseKey(v string) {
	o.LicenseKey = v
}

// GetReadToken returns the ReadToken field value
func (o *NewRelic) GetReadToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReadToken
}

// GetReadTokenOk returns a tuple with the ReadToken field value
// and a boolean to check if the value has been set.
func (o *NewRelic) GetReadTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadToken, true
}

// SetReadToken sets field value
func (o *NewRelic) SetReadToken(v string) {
	o.ReadToken = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NewRelic) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRelic) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NewRelic) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NewRelic) SetType(v string) {
	o.Type = &v
}

// GetWriteToken returns the WriteToken field value
func (o *NewRelic) GetWriteToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WriteToken
}

// GetWriteTokenOk returns a tuple with the WriteToken field value
// and a boolean to check if the value has been set.
func (o *NewRelic) GetWriteTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteToken, true
}

// SetWriteToken sets field value
func (o *NewRelic) SetWriteToken(v string) {
	o.WriteToken = v
}

func (o NewRelic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewRelic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	toSerialize["licenseKey"] = o.LicenseKey
	toSerialize["readToken"] = o.ReadToken
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["writeToken"] = o.WriteToken
	return toSerialize, nil
}

type NullableNewRelic struct {
	value *NewRelic
	isSet bool
}

func (v NullableNewRelic) Get() *NewRelic {
	return v.value
}

func (v *NullableNewRelic) Set(val *NewRelic) {
	v.value = val
	v.isSet = true
}

func (v NullableNewRelic) IsSet() bool {
	return v.isSet
}

func (v *NullableNewRelic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewRelic(val *NewRelic) *NullableNewRelic {
	return &NullableNewRelic{value: val, isSet: true}
}

func (v NullableNewRelic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewRelic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


