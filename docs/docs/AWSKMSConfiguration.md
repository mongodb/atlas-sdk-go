# AWSKMSConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyID** | Pointer to **string** | Unique alphanumeric string that identifies an Identity and Access Management (IAM) access key with permissions required to access your Amazon Web Services (AWS) Customer Master Key (CMK). | [optional] 
**CustomerMasterKeyID** | Pointer to **string** | Unique alphanumeric string that identifies the Amazon Web Services (AWS) Customer Master Key (CMK) you used to encrypt and decrypt the MongoDB master keys. | [optional] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled encryption at rest for the specified project through Amazon Web Services (AWS) Key Management Service (KMS). To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of &#x60;false&#x60;. | [optional] 
**Region** | Pointer to **string** | Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies an Amazon Web Services (AWS) Identity and Access Management (IAM) role. This IAM role has the permissions required to manage your AWS customer master key. | [optional] 
**SecretAccessKey** | Pointer to **string** | Human-readable label of the Identity and Access Management (IAM) secret access key with permissions required to access your Amazon Web Services (AWS) customer master key. | [optional] 
**Valid** | Pointer to **bool** | Flag that indicates whether the Amazon Web Services (AWS) Key Management Service (KMS) encryption key can encrypt and decrypt data. | [optional] [readonly] 

## Methods

### NewAWSKMSConfiguration

`func NewAWSKMSConfiguration() *AWSKMSConfiguration`

NewAWSKMSConfiguration instantiates a new AWSKMSConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSKMSConfigurationWithDefaults

`func NewAWSKMSConfigurationWithDefaults() *AWSKMSConfiguration`

NewAWSKMSConfigurationWithDefaults instantiates a new AWSKMSConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyID

`func (o *AWSKMSConfiguration) GetAccessKeyID() string`

GetAccessKeyID returns the AccessKeyID field if non-nil, zero value otherwise.

### GetAccessKeyIDOk

`func (o *AWSKMSConfiguration) GetAccessKeyIDOk() (*string, bool)`

GetAccessKeyIDOk returns a tuple with the AccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyID

`func (o *AWSKMSConfiguration) SetAccessKeyID(v string)`

SetAccessKeyID sets AccessKeyID field to given value.

### HasAccessKeyID

`func (o *AWSKMSConfiguration) HasAccessKeyID() bool`

HasAccessKeyID returns a boolean if a field has been set.
### GetCustomerMasterKeyID

`func (o *AWSKMSConfiguration) GetCustomerMasterKeyID() string`

GetCustomerMasterKeyID returns the CustomerMasterKeyID field if non-nil, zero value otherwise.

### GetCustomerMasterKeyIDOk

`func (o *AWSKMSConfiguration) GetCustomerMasterKeyIDOk() (*string, bool)`

GetCustomerMasterKeyIDOk returns a tuple with the CustomerMasterKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMasterKeyID

`func (o *AWSKMSConfiguration) SetCustomerMasterKeyID(v string)`

SetCustomerMasterKeyID sets CustomerMasterKeyID field to given value.

### HasCustomerMasterKeyID

`func (o *AWSKMSConfiguration) HasCustomerMasterKeyID() bool`

HasCustomerMasterKeyID returns a boolean if a field has been set.
### GetEnabled

`func (o *AWSKMSConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AWSKMSConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AWSKMSConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AWSKMSConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.
### GetRegion

`func (o *AWSKMSConfiguration) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSKMSConfiguration) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSKMSConfiguration) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AWSKMSConfiguration) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetRoleId

`func (o *AWSKMSConfiguration) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AWSKMSConfiguration) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AWSKMSConfiguration) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *AWSKMSConfiguration) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetSecretAccessKey

`func (o *AWSKMSConfiguration) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AWSKMSConfiguration) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AWSKMSConfiguration) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AWSKMSConfiguration) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.
### GetValid

`func (o *AWSKMSConfiguration) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AWSKMSConfiguration) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AWSKMSConfiguration) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AWSKMSConfiguration) HasValid() bool`

HasValid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


