// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// checks if the CloudProviderAccessAWSIAMRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessAWSIAMRole{}

// CloudProviderAccessAWSIAMRole Details that describe the features linked to the Amazon Web Services (AWS) Identity and Access Management (IAM) role.
type CloudProviderAccessAWSIAMRole struct {
	// Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role.
	AtlasAWSAccountArn *string `json:"atlasAWSAccountArn,omitempty"`
	// Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account.
	AtlasAssumedRoleExternalId *string `json:"atlasAssumedRoleExternalId,omitempty"`
	// Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	AuthorizedDate *time.Time `json:"authorizedDate,omitempty"`
	// Date and time when someone created this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// List that contains application features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
	FeatureUsages []CloudProviderAccessFeatureUsage `json:"featureUsages,omitempty"`
	// Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account.
	IamAssumedRoleArn *string `json:"iamAssumedRoleArn,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the role.
	RoleId *string `json:"roleId,omitempty"`
	// Human-readable label that identifies the cloud provider of the role.
	ProviderName string `json:"providerName"`
}

// NewCloudProviderAccessAWSIAMRole instantiates a new CloudProviderAccessAWSIAMRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessAWSIAMRole(providerName string) *CloudProviderAccessAWSIAMRole {
	this := CloudProviderAccessAWSIAMRole{}
	this.ProviderName = providerName
	return &this
}

// NewCloudProviderAccessAWSIAMRoleWithDefaults instantiates a new CloudProviderAccessAWSIAMRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessAWSIAMRoleWithDefaults() *CloudProviderAccessAWSIAMRole {
	this := CloudProviderAccessAWSIAMRole{}
	return &this
}

// GetAtlasAWSAccountArn returns the AtlasAWSAccountArn field value if set, zero value otherwise.
func (o *CloudProviderAccessAWSIAMRole) GetAtlasAWSAccountArn() string {
	if o == nil || IsNil(o.AtlasAWSAccountArn) {
		var ret string
		return ret
	}
	return *o.AtlasAWSAccountArn
}

// GetAtlasAWSAccountArnOk returns a tuple with the AtlasAWSAccountArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAWSIAMRole) GetAtlasAWSAccountArnOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasAWSAccountArn) {
		return nil, false
	}
	return o.AtlasAWSAccountArn, true
}

// HasAtlasAWSAccountArn returns a boolean if a field has been set.
func (o *CloudProviderAccessAWSIAMRole) HasAtlasAWSAccountArn() bool {
	if o != nil && !IsNil(o.AtlasAWSAccountArn) {
		return true
	}

	return false
}

// SetAtlasAWSAccountArn gets a reference to the given string and assigns it to the AtlasAWSAccountArn field.
func (o *CloudProviderAccessAWSIAMRole) SetAtlasAWSAccountArn(v string) {
	o.AtlasAWSAccountArn = &v
}

// GetAtlasAssumedRoleExternalId returns the AtlasAssumedRoleExternalId field value if set, zero value otherwise.
func (o *CloudProviderAccessAWSIAMRole) GetAtlasAssumedRoleExternalId() string {
	if o == nil || IsNil(o.AtlasAssumedRoleExternalId) {
		var ret string
		return ret
	}
	return *o.AtlasAssumedRoleExternalId
}

// GetAtlasAssumedRoleExternalIdOk returns a tuple with the AtlasAssumedRoleExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAWSIAMRole) GetAtlasAssumedRoleExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasAssumedRoleExternalId) {
		return nil, false
	}
	return o.AtlasAssumedRoleExternalId, true
}

// HasAtlasAssumedRoleExternalId returns a boolean if a field has been set.
func (o *CloudProviderAccessAWSIAMRole) HasAtlasAssumedRoleExternalId() bool {
	if o != nil && !IsNil(o.AtlasAssumedRoleExternalId) {
		return true
	}

	return false
}

// SetAtlasAssumedRoleExternalId gets a reference to the given string and assigns it to the AtlasAssumedRoleExternalId field.
func (o *CloudProviderAccessAWSIAMRole) SetAtlasAssumedRoleExternalId(v string) {
	o.AtlasAssumedRoleExternalId = &v
}

// GetAuthorizedDate returns the AuthorizedDate field value if set, zero value otherwise.
func (o *CloudProviderAccessAWSIAMRole) GetAuthorizedDate() time.Time {
	if o == nil || IsNil(o.AuthorizedDate) {
		var ret time.Time
		return ret
	}
	return *o.AuthorizedDate
}

// GetAuthorizedDateOk returns a tuple with the AuthorizedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAWSIAMRole) GetAuthorizedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AuthorizedDate) {
		return nil, false
	}
	return o.AuthorizedDate, true
}

// HasAuthorizedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessAWSIAMRole) HasAuthorizedDate() bool {
	if o != nil && !IsNil(o.AuthorizedDate) {
		return true
	}

	return false
}

// SetAuthorizedDate gets a reference to the given time.Time and assigns it to the AuthorizedDate field.
func (o *CloudProviderAccessAWSIAMRole) SetAuthorizedDate(v time.Time) {
	o.AuthorizedDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *CloudProviderAccessAWSIAMRole) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAWSIAMRole) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessAWSIAMRole) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *CloudProviderAccessAWSIAMRole) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetFeatureUsages returns the FeatureUsages field value if set, zero value otherwise.
func (o *CloudProviderAccessAWSIAMRole) GetFeatureUsages() []CloudProviderAccessFeatureUsage {
	if o == nil || IsNil(o.FeatureUsages) {
		var ret []CloudProviderAccessFeatureUsage
		return ret
	}
	return o.FeatureUsages
}

// GetFeatureUsagesOk returns a tuple with the FeatureUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAWSIAMRole) GetFeatureUsagesOk() ([]CloudProviderAccessFeatureUsage, bool) {
	if o == nil || IsNil(o.FeatureUsages) {
		return nil, false
	}
	return o.FeatureUsages, true
}

// HasFeatureUsages returns a boolean if a field has been set.
func (o *CloudProviderAccessAWSIAMRole) HasFeatureUsages() bool {
	if o != nil && !IsNil(o.FeatureUsages) {
		return true
	}

	return false
}

// SetFeatureUsages gets a reference to the given []CloudProviderAccessFeatureUsage and assigns it to the FeatureUsages field.
func (o *CloudProviderAccessAWSIAMRole) SetFeatureUsages(v []CloudProviderAccessFeatureUsage) {
	o.FeatureUsages = v
}

// GetIamAssumedRoleArn returns the IamAssumedRoleArn field value if set, zero value otherwise.
func (o *CloudProviderAccessAWSIAMRole) GetIamAssumedRoleArn() string {
	if o == nil || IsNil(o.IamAssumedRoleArn) {
		var ret string
		return ret
	}
	return *o.IamAssumedRoleArn
}

// GetIamAssumedRoleArnOk returns a tuple with the IamAssumedRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAWSIAMRole) GetIamAssumedRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.IamAssumedRoleArn) {
		return nil, false
	}
	return o.IamAssumedRoleArn, true
}

// HasIamAssumedRoleArn returns a boolean if a field has been set.
func (o *CloudProviderAccessAWSIAMRole) HasIamAssumedRoleArn() bool {
	if o != nil && !IsNil(o.IamAssumedRoleArn) {
		return true
	}

	return false
}

// SetIamAssumedRoleArn gets a reference to the given string and assigns it to the IamAssumedRoleArn field.
func (o *CloudProviderAccessAWSIAMRole) SetIamAssumedRoleArn(v string) {
	o.IamAssumedRoleArn = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *CloudProviderAccessAWSIAMRole) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAWSIAMRole) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *CloudProviderAccessAWSIAMRole) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *CloudProviderAccessAWSIAMRole) SetRoleId(v string) {
	o.RoleId = &v
}

// GetProviderName returns the ProviderName field value
func (o *CloudProviderAccessAWSIAMRole) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAWSIAMRole) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *CloudProviderAccessAWSIAMRole) SetProviderName(v string) {
	o.ProviderName = v
}

func (o CloudProviderAccessAWSIAMRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAccessAWSIAMRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IamAssumedRoleArn) {
		toSerialize["iamAssumedRoleArn"] = o.IamAssumedRoleArn
	}
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableCloudProviderAccessAWSIAMRole struct {
	value *CloudProviderAccessAWSIAMRole
	isSet bool
}

func (v NullableCloudProviderAccessAWSIAMRole) Get() *CloudProviderAccessAWSIAMRole {
	return v.value
}

func (v *NullableCloudProviderAccessAWSIAMRole) Set(val *CloudProviderAccessAWSIAMRole) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessAWSIAMRole) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessAWSIAMRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessAWSIAMRole(val *CloudProviderAccessAWSIAMRole) *NullableCloudProviderAccessAWSIAMRole {
	return &NullableCloudProviderAccessAWSIAMRole{value: val, isSet: true}
}

func (v NullableCloudProviderAccessAWSIAMRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessAWSIAMRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
