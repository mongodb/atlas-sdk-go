# CloudProviderAccessAzureServicePrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**AtlasAzureAppId** | Pointer to **string** | Azure Active Directory Application ID of Atlas. | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date and time when this Azure Service Principal was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FeatureUsages** | Pointer to [**[]CloudProviderAccessFeatureUsage**](CloudProviderAccessFeatureUsage.md) | List that contains application features associated with this Azure Service Principal. | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **time.Time** | Date and time when this Azure Service Principal was last updated. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**ServicePrincipalId** | Pointer to **string** | UUID string that identifies the Azure Service Principal. | [optional] 
**TenantId** | Pointer to **string** | UUID String that identifies the Azure Active Directory Tenant ID. | [optional] 
**ProviderName** | **string** | Human-readable label that identifies the cloud provider of the role. | 
**AtlasAWSAccountArn** | Pointer to **string** | Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role. | [optional] [readonly] 
**AtlasAssumedRoleExternalId** | Pointer to **string** | Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account. | [optional] [readonly] 
**AuthorizedDate** | Pointer to **time.Time** | Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**IamAssumedRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**GcpServiceAccountForAtlas** | Pointer to **string** | Email address for the Google Service Account created by Atlas. | [optional] 

## Methods

### NewCloudProviderAccessAzureServicePrincipal

`func NewCloudProviderAccessAzureServicePrincipal(providerName string, ) *CloudProviderAccessAzureServicePrincipal`

NewCloudProviderAccessAzureServicePrincipal instantiates a new CloudProviderAccessAzureServicePrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessAzureServicePrincipalWithDefaults

`func NewCloudProviderAccessAzureServicePrincipalWithDefaults() *CloudProviderAccessAzureServicePrincipal`

NewCloudProviderAccessAzureServicePrincipalWithDefaults instantiates a new CloudProviderAccessAzureServicePrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudProviderAccessAzureServicePrincipal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProviderAccessAzureServicePrincipal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProviderAccessAzureServicePrincipal) HasId() bool`

HasId returns a boolean if a field has been set.
### GetAtlasAzureAppId

`func (o *CloudProviderAccessAzureServicePrincipal) GetAtlasAzureAppId() string`

GetAtlasAzureAppId returns the AtlasAzureAppId field if non-nil, zero value otherwise.

### GetAtlasAzureAppIdOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetAtlasAzureAppIdOk() (*string, bool)`

GetAtlasAzureAppIdOk returns a tuple with the AtlasAzureAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAzureAppId

`func (o *CloudProviderAccessAzureServicePrincipal) SetAtlasAzureAppId(v string)`

SetAtlasAzureAppId sets AtlasAzureAppId field to given value.

### HasAtlasAzureAppId

`func (o *CloudProviderAccessAzureServicePrincipal) HasAtlasAzureAppId() bool`

HasAtlasAzureAppId returns a boolean if a field has been set.
### GetCreatedDate

`func (o *CloudProviderAccessAzureServicePrincipal) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CloudProviderAccessAzureServicePrincipal) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CloudProviderAccessAzureServicePrincipal) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.
### GetFeatureUsages

`func (o *CloudProviderAccessAzureServicePrincipal) GetFeatureUsages() []CloudProviderAccessFeatureUsage`

GetFeatureUsages returns the FeatureUsages field if non-nil, zero value otherwise.

### GetFeatureUsagesOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetFeatureUsagesOk() (*[]CloudProviderAccessFeatureUsage, bool)`

GetFeatureUsagesOk returns a tuple with the FeatureUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsages

`func (o *CloudProviderAccessAzureServicePrincipal) SetFeatureUsages(v []CloudProviderAccessFeatureUsage)`

SetFeatureUsages sets FeatureUsages field to given value.

### HasFeatureUsages

`func (o *CloudProviderAccessAzureServicePrincipal) HasFeatureUsages() bool`

HasFeatureUsages returns a boolean if a field has been set.
### GetLastUpdatedDate

`func (o *CloudProviderAccessAzureServicePrincipal) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *CloudProviderAccessAzureServicePrincipal) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *CloudProviderAccessAzureServicePrincipal) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.
### GetServicePrincipalId

`func (o *CloudProviderAccessAzureServicePrincipal) GetServicePrincipalId() string`

GetServicePrincipalId returns the ServicePrincipalId field if non-nil, zero value otherwise.

### GetServicePrincipalIdOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetServicePrincipalIdOk() (*string, bool)`

GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalId

`func (o *CloudProviderAccessAzureServicePrincipal) SetServicePrincipalId(v string)`

SetServicePrincipalId sets ServicePrincipalId field to given value.

### HasServicePrincipalId

`func (o *CloudProviderAccessAzureServicePrincipal) HasServicePrincipalId() bool`

HasServicePrincipalId returns a boolean if a field has been set.
### GetTenantId

`func (o *CloudProviderAccessAzureServicePrincipal) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudProviderAccessAzureServicePrincipal) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CloudProviderAccessAzureServicePrincipal) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.
### GetProviderName

`func (o *CloudProviderAccessAzureServicePrincipal) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CloudProviderAccessAzureServicePrincipal) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### GetAtlasAWSAccountArn

`func (o *CloudProviderAccessAzureServicePrincipal) GetAtlasAWSAccountArn() string`

GetAtlasAWSAccountArn returns the AtlasAWSAccountArn field if non-nil, zero value otherwise.

### GetAtlasAWSAccountArnOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetAtlasAWSAccountArnOk() (*string, bool)`

GetAtlasAWSAccountArnOk returns a tuple with the AtlasAWSAccountArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAWSAccountArn

`func (o *CloudProviderAccessAzureServicePrincipal) SetAtlasAWSAccountArn(v string)`

SetAtlasAWSAccountArn sets AtlasAWSAccountArn field to given value.

### HasAtlasAWSAccountArn

`func (o *CloudProviderAccessAzureServicePrincipal) HasAtlasAWSAccountArn() bool`

HasAtlasAWSAccountArn returns a boolean if a field has been set.
### GetAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessAzureServicePrincipal) GetAtlasAssumedRoleExternalId() string`

GetAtlasAssumedRoleExternalId returns the AtlasAssumedRoleExternalId field if non-nil, zero value otherwise.

### GetAtlasAssumedRoleExternalIdOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetAtlasAssumedRoleExternalIdOk() (*string, bool)`

GetAtlasAssumedRoleExternalIdOk returns a tuple with the AtlasAssumedRoleExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessAzureServicePrincipal) SetAtlasAssumedRoleExternalId(v string)`

SetAtlasAssumedRoleExternalId sets AtlasAssumedRoleExternalId field to given value.

### HasAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessAzureServicePrincipal) HasAtlasAssumedRoleExternalId() bool`

HasAtlasAssumedRoleExternalId returns a boolean if a field has been set.
### GetAuthorizedDate

`func (o *CloudProviderAccessAzureServicePrincipal) GetAuthorizedDate() time.Time`

GetAuthorizedDate returns the AuthorizedDate field if non-nil, zero value otherwise.

### GetAuthorizedDateOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetAuthorizedDateOk() (*time.Time, bool)`

GetAuthorizedDateOk returns a tuple with the AuthorizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDate

`func (o *CloudProviderAccessAzureServicePrincipal) SetAuthorizedDate(v time.Time)`

SetAuthorizedDate sets AuthorizedDate field to given value.

### HasAuthorizedDate

`func (o *CloudProviderAccessAzureServicePrincipal) HasAuthorizedDate() bool`

HasAuthorizedDate returns a boolean if a field has been set.
### GetIamAssumedRoleArn

`func (o *CloudProviderAccessAzureServicePrincipal) GetIamAssumedRoleArn() string`

GetIamAssumedRoleArn returns the IamAssumedRoleArn field if non-nil, zero value otherwise.

### GetIamAssumedRoleArnOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetIamAssumedRoleArnOk() (*string, bool)`

GetIamAssumedRoleArnOk returns a tuple with the IamAssumedRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamAssumedRoleArn

`func (o *CloudProviderAccessAzureServicePrincipal) SetIamAssumedRoleArn(v string)`

SetIamAssumedRoleArn sets IamAssumedRoleArn field to given value.

### HasIamAssumedRoleArn

`func (o *CloudProviderAccessAzureServicePrincipal) HasIamAssumedRoleArn() bool`

HasIamAssumedRoleArn returns a boolean if a field has been set.
### GetRoleId

`func (o *CloudProviderAccessAzureServicePrincipal) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CloudProviderAccessAzureServicePrincipal) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *CloudProviderAccessAzureServicePrincipal) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetGcpServiceAccountForAtlas

`func (o *CloudProviderAccessAzureServicePrincipal) GetGcpServiceAccountForAtlas() string`

GetGcpServiceAccountForAtlas returns the GcpServiceAccountForAtlas field if non-nil, zero value otherwise.

### GetGcpServiceAccountForAtlasOk

`func (o *CloudProviderAccessAzureServicePrincipal) GetGcpServiceAccountForAtlasOk() (*string, bool)`

GetGcpServiceAccountForAtlasOk returns a tuple with the GcpServiceAccountForAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountForAtlas

`func (o *CloudProviderAccessAzureServicePrincipal) SetGcpServiceAccountForAtlas(v string)`

SetGcpServiceAccountForAtlas sets GcpServiceAccountForAtlas field to given value.

### HasGcpServiceAccountForAtlas

`func (o *CloudProviderAccessAzureServicePrincipal) HasGcpServiceAccountForAtlas() bool`

HasGcpServiceAccountForAtlas returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


