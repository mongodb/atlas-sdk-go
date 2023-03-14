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

// checks if the AWSKMS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSKMS{}

// AWSKMS Amazon Web Services (AWS) KMS configuration details and encryption at rest configuration set for the specified project.
type AWSKMS struct {
	// Unique alphanumeric string that identifies an Identity and Access Management (IAM) access key with permissions required to access your Amazon Web Services (AWS) Customer Master Key (CMK).
	AccessKeyID *string `json:"accessKeyID,omitempty"`
	// Unique alphanumeric string that identifies the Amazon Web Services (AWS) Customer Master Key (CMK) you used to encrypt and decrypt the MongoDB master keys.
	CustomerMasterKeyID *string `json:"customerMasterKeyID,omitempty"`
	// Flag that indicates whether someone enabled encryption at rest for the specified project through Amazon Web Services (AWS) Key Management Service (KMS). To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of `false`.
	Enabled *bool `json:"enabled,omitempty"`
	//  Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts.
	Region *string `json:"region,omitempty"`
	// Unique 24-hexadecimal digit string that identifies an Amazon Web Services (AWS) Identity and Access Management (IAM) role. This IAM role has the permissions required to manage your AWS customer master key.
	RoleId *string `json:"roleId,omitempty"`
	// Human-readable label of the Identity and Access Management (IAM) secret access key with permissions required to access your Amazon Web Services (AWS) customer master key.
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	// Flag that indicates whether the Amazon Web Services (AWS) Key Management Service (KMS) encryption key can encrypt and decrypt data.
	Valid *bool `json:"valid,omitempty"`
}

// NewAWSKMS instantiates a new AWSKMS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSKMS() *AWSKMS {
	this := AWSKMS{}
	return &this
}

// NewAWSKMSWithDefaults instantiates a new AWSKMS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSKMSWithDefaults() *AWSKMS {
	this := AWSKMS{}
	return &this
}

// GetAccessKeyID returns the AccessKeyID field value if set, zero value otherwise.
func (o *AWSKMS) GetAccessKeyID() string {
	if o == nil || IsNil(o.AccessKeyID) {
		var ret string
		return ret
	}
	return *o.AccessKeyID
}

// GetAccessKeyIDOk returns a tuple with the AccessKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSKMS) GetAccessKeyIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeyID) {
		return nil, false
	}
	return o.AccessKeyID, true
}

// HasAccessKeyID returns a boolean if a field has been set.
func (o *AWSKMS) HasAccessKeyID() bool {
	if o != nil && !IsNil(o.AccessKeyID) {
		return true
	}

	return false
}

// SetAccessKeyID gets a reference to the given string and assigns it to the AccessKeyID field.
func (o *AWSKMS) SetAccessKeyID(v string) {
	o.AccessKeyID = &v
}

// GetCustomerMasterKeyID returns the CustomerMasterKeyID field value if set, zero value otherwise.
func (o *AWSKMS) GetCustomerMasterKeyID() string {
	if o == nil || IsNil(o.CustomerMasterKeyID) {
		var ret string
		return ret
	}
	return *o.CustomerMasterKeyID
}

// GetCustomerMasterKeyIDOk returns a tuple with the CustomerMasterKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSKMS) GetCustomerMasterKeyIDOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerMasterKeyID) {
		return nil, false
	}
	return o.CustomerMasterKeyID, true
}

// HasCustomerMasterKeyID returns a boolean if a field has been set.
func (o *AWSKMS) HasCustomerMasterKeyID() bool {
	if o != nil && !IsNil(o.CustomerMasterKeyID) {
		return true
	}

	return false
}

// SetCustomerMasterKeyID gets a reference to the given string and assigns it to the CustomerMasterKeyID field.
func (o *AWSKMS) SetCustomerMasterKeyID(v string) {
	o.CustomerMasterKeyID = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AWSKMS) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSKMS) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AWSKMS) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AWSKMS) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AWSKMS) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSKMS) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AWSKMS) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AWSKMS) SetRegion(v string) {
	o.Region = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *AWSKMS) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSKMS) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *AWSKMS) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *AWSKMS) SetRoleId(v string) {
	o.RoleId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *AWSKMS) GetSecretAccessKey() string {
	if o == nil || IsNil(o.SecretAccessKey) {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSKMS) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretAccessKey) {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *AWSKMS) HasSecretAccessKey() bool {
	if o != nil && !IsNil(o.SecretAccessKey) {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *AWSKMS) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *AWSKMS) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSKMS) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *AWSKMS) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *AWSKMS) SetValid(v bool) {
	o.Valid = &v
}

func (o AWSKMS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSKMS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKeyID) {
		toSerialize["accessKeyID"] = o.AccessKeyID
	}
	if !IsNil(o.CustomerMasterKeyID) {
		toSerialize["customerMasterKeyID"] = o.CustomerMasterKeyID
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !IsNil(o.SecretAccessKey) {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	// skip: valid is readOnly
	return toSerialize, nil
}

type NullableAWSKMS struct {
	value *AWSKMS
	isSet bool
}

func (v NullableAWSKMS) Get() *AWSKMS {
	return v.value
}

func (v *NullableAWSKMS) Set(val *AWSKMS) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSKMS) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSKMS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSKMS(val *AWSKMS) *NullableAWSKMS {
	return &NullableAWSKMS{value: val, isSet: true}
}

func (v NullableAWSKMS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSKMS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


