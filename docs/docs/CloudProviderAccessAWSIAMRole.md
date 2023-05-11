# CloudProviderAccessAWSIAMRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasAWSAccountArn** | Pointer to **string** | Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role. | [optional] [readonly] 
**AtlasAssumedRoleExternalId** | Pointer to **string** | Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account. | [optional] [readonly] 
**AuthorizedDate** | Pointer to **time.Time** | Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** | Date and time when someone created this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FeatureUsages** | Pointer to [**[]CloudProviderAccessFeatureUsage**](CloudProviderAccessFeatureUsage.md) | List that contains application features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role. | [optional] [readonly] 
**IamAssumedRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**ProviderName** | **string** | Human-readable label that identifies the cloud provider of the role. | 

## Methods

### NewCloudProviderAccessAWSIAMRole

`func NewCloudProviderAccessAWSIAMRole(providerName string, ) *CloudProviderAccessAWSIAMRole`

NewCloudProviderAccessAWSIAMRole instantiates a new CloudProviderAccessAWSIAMRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessAWSIAMRoleWithDefaults

`func NewCloudProviderAccessAWSIAMRoleWithDefaults() *CloudProviderAccessAWSIAMRole`

NewCloudProviderAccessAWSIAMRoleWithDefaults instantiates a new CloudProviderAccessAWSIAMRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasAWSAccountArn

`func (o *CloudProviderAccessAWSIAMRole) GetAtlasAWSAccountArn() string`

GetAtlasAWSAccountArn returns the AtlasAWSAccountArn field if non-nil, zero value otherwise.

### GetAtlasAWSAccountArnOk

`func (o *CloudProviderAccessAWSIAMRole) GetAtlasAWSAccountArnOk() (*string, bool)`

GetAtlasAWSAccountArnOk returns a tuple with the AtlasAWSAccountArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAWSAccountArn

`func (o *CloudProviderAccessAWSIAMRole) SetAtlasAWSAccountArn(v string)`

SetAtlasAWSAccountArn sets AtlasAWSAccountArn field to given value.

### HasAtlasAWSAccountArn

`func (o *CloudProviderAccessAWSIAMRole) HasAtlasAWSAccountArn() bool`

HasAtlasAWSAccountArn returns a boolean if a field has been set.

### GetAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessAWSIAMRole) GetAtlasAssumedRoleExternalId() string`

GetAtlasAssumedRoleExternalId returns the AtlasAssumedRoleExternalId field if non-nil, zero value otherwise.

### GetAtlasAssumedRoleExternalIdOk

`func (o *CloudProviderAccessAWSIAMRole) GetAtlasAssumedRoleExternalIdOk() (*string, bool)`

GetAtlasAssumedRoleExternalIdOk returns a tuple with the AtlasAssumedRoleExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessAWSIAMRole) SetAtlasAssumedRoleExternalId(v string)`

SetAtlasAssumedRoleExternalId sets AtlasAssumedRoleExternalId field to given value.

### HasAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessAWSIAMRole) HasAtlasAssumedRoleExternalId() bool`

HasAtlasAssumedRoleExternalId returns a boolean if a field has been set.

### GetAuthorizedDate

`func (o *CloudProviderAccessAWSIAMRole) GetAuthorizedDate() time.Time`

GetAuthorizedDate returns the AuthorizedDate field if non-nil, zero value otherwise.

### GetAuthorizedDateOk

`func (o *CloudProviderAccessAWSIAMRole) GetAuthorizedDateOk() (*time.Time, bool)`

GetAuthorizedDateOk returns a tuple with the AuthorizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDate

`func (o *CloudProviderAccessAWSIAMRole) SetAuthorizedDate(v time.Time)`

SetAuthorizedDate sets AuthorizedDate field to given value.

### HasAuthorizedDate

`func (o *CloudProviderAccessAWSIAMRole) HasAuthorizedDate() bool`

HasAuthorizedDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *CloudProviderAccessAWSIAMRole) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CloudProviderAccessAWSIAMRole) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CloudProviderAccessAWSIAMRole) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CloudProviderAccessAWSIAMRole) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetFeatureUsages

`func (o *CloudProviderAccessAWSIAMRole) GetFeatureUsages() []CloudProviderAccessFeatureUsage`

GetFeatureUsages returns the FeatureUsages field if non-nil, zero value otherwise.

### GetFeatureUsagesOk

`func (o *CloudProviderAccessAWSIAMRole) GetFeatureUsagesOk() (*[]CloudProviderAccessFeatureUsage, bool)`

GetFeatureUsagesOk returns a tuple with the FeatureUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsages

`func (o *CloudProviderAccessAWSIAMRole) SetFeatureUsages(v []CloudProviderAccessFeatureUsage)`

SetFeatureUsages sets FeatureUsages field to given value.

### HasFeatureUsages

`func (o *CloudProviderAccessAWSIAMRole) HasFeatureUsages() bool`

HasFeatureUsages returns a boolean if a field has been set.

### GetIamAssumedRoleArn

`func (o *CloudProviderAccessAWSIAMRole) GetIamAssumedRoleArn() string`

GetIamAssumedRoleArn returns the IamAssumedRoleArn field if non-nil, zero value otherwise.

### GetIamAssumedRoleArnOk

`func (o *CloudProviderAccessAWSIAMRole) GetIamAssumedRoleArnOk() (*string, bool)`

GetIamAssumedRoleArnOk returns a tuple with the IamAssumedRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamAssumedRoleArn

`func (o *CloudProviderAccessAWSIAMRole) SetIamAssumedRoleArn(v string)`

SetIamAssumedRoleArn sets IamAssumedRoleArn field to given value.

### HasIamAssumedRoleArn

`func (o *CloudProviderAccessAWSIAMRole) HasIamAssumedRoleArn() bool`

HasIamAssumedRoleArn returns a boolean if a field has been set.

### GetRoleId

`func (o *CloudProviderAccessAWSIAMRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CloudProviderAccessAWSIAMRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CloudProviderAccessAWSIAMRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *CloudProviderAccessAWSIAMRole) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetProviderName

`func (o *CloudProviderAccessAWSIAMRole) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CloudProviderAccessAWSIAMRole) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CloudProviderAccessAWSIAMRole) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


