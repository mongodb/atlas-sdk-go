/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the EncryptionAtRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionAtRest{}

// EncryptionAtRest struct for EncryptionAtRest
type EncryptionAtRest struct {
	AwsKms *AWSKMS `json:"awsKms,omitempty"`
	AzureKeyVault *AzureKeyVault `json:"azureKeyVault,omitempty"`
	GoogleCloudKms *GoogleCloudKMS `json:"googleCloudKms,omitempty"`
}

// NewEncryptionAtRest instantiates a new EncryptionAtRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionAtRest() *EncryptionAtRest {
	this := EncryptionAtRest{}
	return &this
}

// NewEncryptionAtRestWithDefaults instantiates a new EncryptionAtRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionAtRestWithDefaults() *EncryptionAtRest {
	this := EncryptionAtRest{}
	return &this
}

// GetAwsKms returns the AwsKms field value if set, zero value otherwise.
func (o *EncryptionAtRest) GetAwsKms() AWSKMS {
	if o == nil || IsNil(o.AwsKms) {
		var ret AWSKMS
		return ret
	}
	return *o.AwsKms
}

// GetAwsKmsOk returns a tuple with the AwsKms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionAtRest) GetAwsKmsOk() (*AWSKMS, bool) {
	if o == nil || IsNil(o.AwsKms) {
		return nil, false
	}
	return o.AwsKms, true
}

// HasAwsKms returns a boolean if a field has been set.
func (o *EncryptionAtRest) HasAwsKms() bool {
	if o != nil && !IsNil(o.AwsKms) {
		return true
	}

	return false
}

// SetAwsKms gets a reference to the given AWSKMS and assigns it to the AwsKms field.
func (o *EncryptionAtRest) SetAwsKms(v AWSKMS) {
	o.AwsKms = &v
}

// GetAzureKeyVault returns the AzureKeyVault field value if set, zero value otherwise.
func (o *EncryptionAtRest) GetAzureKeyVault() AzureKeyVault {
	if o == nil || IsNil(o.AzureKeyVault) {
		var ret AzureKeyVault
		return ret
	}
	return *o.AzureKeyVault
}

// GetAzureKeyVaultOk returns a tuple with the AzureKeyVault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionAtRest) GetAzureKeyVaultOk() (*AzureKeyVault, bool) {
	if o == nil || IsNil(o.AzureKeyVault) {
		return nil, false
	}
	return o.AzureKeyVault, true
}

// HasAzureKeyVault returns a boolean if a field has been set.
func (o *EncryptionAtRest) HasAzureKeyVault() bool {
	if o != nil && !IsNil(o.AzureKeyVault) {
		return true
	}

	return false
}

// SetAzureKeyVault gets a reference to the given AzureKeyVault and assigns it to the AzureKeyVault field.
func (o *EncryptionAtRest) SetAzureKeyVault(v AzureKeyVault) {
	o.AzureKeyVault = &v
}

// GetGoogleCloudKms returns the GoogleCloudKms field value if set, zero value otherwise.
func (o *EncryptionAtRest) GetGoogleCloudKms() GoogleCloudKMS {
	if o == nil || IsNil(o.GoogleCloudKms) {
		var ret GoogleCloudKMS
		return ret
	}
	return *o.GoogleCloudKms
}

// GetGoogleCloudKmsOk returns a tuple with the GoogleCloudKms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionAtRest) GetGoogleCloudKmsOk() (*GoogleCloudKMS, bool) {
	if o == nil || IsNil(o.GoogleCloudKms) {
		return nil, false
	}
	return o.GoogleCloudKms, true
}

// HasGoogleCloudKms returns a boolean if a field has been set.
func (o *EncryptionAtRest) HasGoogleCloudKms() bool {
	if o != nil && !IsNil(o.GoogleCloudKms) {
		return true
	}

	return false
}

// SetGoogleCloudKms gets a reference to the given GoogleCloudKMS and assigns it to the GoogleCloudKms field.
func (o *EncryptionAtRest) SetGoogleCloudKms(v GoogleCloudKMS) {
	o.GoogleCloudKms = &v
}

func (o EncryptionAtRest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o EncryptionAtRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsKms) {
		toSerialize["awsKms"] = o.AwsKms
	}
	if !IsNil(o.AzureKeyVault) {
		toSerialize["azureKeyVault"] = o.AzureKeyVault
	}
	if !IsNil(o.GoogleCloudKms) {
		toSerialize["googleCloudKms"] = o.GoogleCloudKms
	}
	return toSerialize, nil
}

type NullableEncryptionAtRest struct {
	value *EncryptionAtRest
	isSet bool
}

func (v NullableEncryptionAtRest) Get() *EncryptionAtRest {
	return v.value
}

func (v *NullableEncryptionAtRest) Set(val *EncryptionAtRest) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionAtRest) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionAtRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionAtRest(val *EncryptionAtRest) *NullableEncryptionAtRest {
	return &NullableEncryptionAtRest{value: val, isSet: true}
}

func (v NullableEncryptionAtRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionAtRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


