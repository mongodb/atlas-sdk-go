/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
)

// checks if the GoogleCloudKMS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleCloudKMS{}

// GoogleCloudKMS Details that define the configuration of Encryption at Rest using Google Cloud Key Management Service (KMS).
type GoogleCloudKMS struct {
	// Flag that indicates whether someone enabled encryption at rest for the specified  project. To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of `false`.
	Enabled *bool `json:"enabled,omitempty"`
	// Resource path that displays the key version resource ID for your Google Cloud KMS.
	KeyVersionResourceID *string `json:"keyVersionResourceID,omitempty"`
	// JavaScript Object Notation (JSON) object that contains the Google Cloud Key Management Service (KMS). Format the JSON as a string and not as an object.
	ServiceAccountKey *string `json:"serviceAccountKey,omitempty"`
	// Flag that indicates whether the Google Cloud Key Management Service (KMS) encryption key can encrypt and decrypt data.
	Valid *bool `json:"valid,omitempty"`
}

// NewGoogleCloudKMS instantiates a new GoogleCloudKMS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudKMS() *GoogleCloudKMS {
	this := GoogleCloudKMS{}
	return &this
}

// NewGoogleCloudKMSWithDefaults instantiates a new GoogleCloudKMS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKMSWithDefaults() *GoogleCloudKMS {
	this := GoogleCloudKMS{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GoogleCloudKMS) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudKMS) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GoogleCloudKMS) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GoogleCloudKMS) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetKeyVersionResourceID returns the KeyVersionResourceID field value if set, zero value otherwise.
func (o *GoogleCloudKMS) GetKeyVersionResourceID() string {
	if o == nil || IsNil(o.KeyVersionResourceID) {
		var ret string
		return ret
	}
	return *o.KeyVersionResourceID
}

// GetKeyVersionResourceIDOk returns a tuple with the KeyVersionResourceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudKMS) GetKeyVersionResourceIDOk() (*string, bool) {
	if o == nil || IsNil(o.KeyVersionResourceID) {
		return nil, false
	}
	return o.KeyVersionResourceID, true
}

// HasKeyVersionResourceID returns a boolean if a field has been set.
func (o *GoogleCloudKMS) HasKeyVersionResourceID() bool {
	if o != nil && !IsNil(o.KeyVersionResourceID) {
		return true
	}

	return false
}

// SetKeyVersionResourceID gets a reference to the given string and assigns it to the KeyVersionResourceID field.
func (o *GoogleCloudKMS) SetKeyVersionResourceID(v string) {
	o.KeyVersionResourceID = &v
}

// GetServiceAccountKey returns the ServiceAccountKey field value if set, zero value otherwise.
func (o *GoogleCloudKMS) GetServiceAccountKey() string {
	if o == nil || IsNil(o.ServiceAccountKey) {
		var ret string
		return ret
	}
	return *o.ServiceAccountKey
}

// GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudKMS) GetServiceAccountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountKey) {
		return nil, false
	}
	return o.ServiceAccountKey, true
}

// HasServiceAccountKey returns a boolean if a field has been set.
func (o *GoogleCloudKMS) HasServiceAccountKey() bool {
	if o != nil && !IsNil(o.ServiceAccountKey) {
		return true
	}

	return false
}

// SetServiceAccountKey gets a reference to the given string and assigns it to the ServiceAccountKey field.
func (o *GoogleCloudKMS) SetServiceAccountKey(v string) {
	o.ServiceAccountKey = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *GoogleCloudKMS) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudKMS) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *GoogleCloudKMS) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *GoogleCloudKMS) SetValid(v bool) {
	o.Valid = &v
}

func (o GoogleCloudKMS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleCloudKMS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.KeyVersionResourceID) {
		toSerialize["keyVersionResourceID"] = o.KeyVersionResourceID
	}
	if !IsNil(o.ServiceAccountKey) {
		toSerialize["serviceAccountKey"] = o.ServiceAccountKey
	}
	// skip: valid is readOnly
	return toSerialize, nil
}

type NullableGoogleCloudKMS struct {
	value *GoogleCloudKMS
	isSet bool
}

func (v NullableGoogleCloudKMS) Get() *GoogleCloudKMS {
	return v.value
}

func (v *NullableGoogleCloudKMS) Set(val *GoogleCloudKMS) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudKMS) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudKMS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudKMS(val *GoogleCloudKMS) *NullableGoogleCloudKMS {
	return &NullableGoogleCloudKMS{value: val, isSet: true}
}

func (v NullableGoogleCloudKMS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudKMS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


