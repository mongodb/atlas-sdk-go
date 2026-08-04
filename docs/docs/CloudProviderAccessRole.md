# CloudProviderAccessRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Human-readable label that identifies the cloud provider of the role. | 
**AtlasAWSAccountArn** | Pointer to **string** | Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role. | [optional] [readonly] 
**AtlasAssumedRoleExternalId** | Pointer to **string** | Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account. | [optional] [readonly] 
**AuthorizedDate** | Pointer to **time.Time** | Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** | Date and time when this Google Service Account was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FeatureUsages** | Pointer to [**[]CloudProviderAccessFeatureUsage**](CloudProviderAccessFeatureUsage.md) | List that contains application features associated with this Google Service Account. | [optional] [readonly] 
**IamAssumedRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**AtlasAzureAppId** | Pointer to **string** | Azure Active Directory Application ID of Atlas. This field is optional and will be derived from the Azure subscription if not provided. | [optional] 
**LastUpdatedDate** | Pointer to **time.Time** | Date and time when this Azure Service Principal was last updated. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**ServicePrincipalId** | Pointer to **string** | UUID string that identifies the Azure Service Principal. | [optional] 
**TenantId** | Pointer to **string** | UUID String that identifies the Azure Active Directory Tenant ID. | [optional] 
**GcpServiceAccountForAtlas** | Pointer to **string** | Email address for the Google Service Account created by Atlas. | [optional] 
**Status** | Pointer to **string** | Provision status of the service account. | [optional] [readonly] 

## Methods

### NewCloudProviderAccessRole

`func NewCloudProviderAccessRole(providerName string, ) *CloudProviderAccessRole`

NewCloudProviderAccessRole instantiates a new CloudProviderAccessRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessRoleWithDefaults

`func NewCloudProviderAccessRoleWithDefaults() *CloudProviderAccessRole`

NewCloudProviderAccessRoleWithDefaults instantiates a new CloudProviderAccessRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *CloudProviderAccessRole) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CloudProviderAccessRole) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CloudProviderAccessRole) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### GetAtlasAWSAccountArn

`func (o *CloudProviderAccessRole) GetAtlasAWSAccountArn() string`

GetAtlasAWSAccountArn returns the AtlasAWSAccountArn field if non-nil, zero value otherwise.

### GetAtlasAWSAccountArnOk

`func (o *CloudProviderAccessRole) GetAtlasAWSAccountArnOk() (*string, bool)`

GetAtlasAWSAccountArnOk returns a tuple with the AtlasAWSAccountArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAWSAccountArn

`func (o *CloudProviderAccessRole) SetAtlasAWSAccountArn(v string)`

SetAtlasAWSAccountArn sets AtlasAWSAccountArn field to given value.

### HasAtlasAWSAccountArn

`func (o *CloudProviderAccessRole) HasAtlasAWSAccountArn() bool`

HasAtlasAWSAccountArn returns a boolean if a field has been set.

### SetAtlasAWSAccountArnNil

`func (o *CloudProviderAccessRole) SetAtlasAWSAccountArnNil()`

SetAtlasAWSAccountArnNil sets AtlasAWSAccountArn to an explicit JSON null when marshaled, overriding any value previously set with SetAtlasAWSAccountArn. Calling SetAtlasAWSAccountArn again clears the null override.

### GetAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessRole) GetAtlasAssumedRoleExternalId() string`

GetAtlasAssumedRoleExternalId returns the AtlasAssumedRoleExternalId field if non-nil, zero value otherwise.

### GetAtlasAssumedRoleExternalIdOk

`func (o *CloudProviderAccessRole) GetAtlasAssumedRoleExternalIdOk() (*string, bool)`

GetAtlasAssumedRoleExternalIdOk returns a tuple with the AtlasAssumedRoleExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessRole) SetAtlasAssumedRoleExternalId(v string)`

SetAtlasAssumedRoleExternalId sets AtlasAssumedRoleExternalId field to given value.

### HasAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessRole) HasAtlasAssumedRoleExternalId() bool`

HasAtlasAssumedRoleExternalId returns a boolean if a field has been set.

### SetAtlasAssumedRoleExternalIdNil

`func (o *CloudProviderAccessRole) SetAtlasAssumedRoleExternalIdNil()`

SetAtlasAssumedRoleExternalIdNil sets AtlasAssumedRoleExternalId to an explicit JSON null when marshaled, overriding any value previously set with SetAtlasAssumedRoleExternalId. Calling SetAtlasAssumedRoleExternalId again clears the null override.

### GetAuthorizedDate

`func (o *CloudProviderAccessRole) GetAuthorizedDate() time.Time`

GetAuthorizedDate returns the AuthorizedDate field if non-nil, zero value otherwise.

### GetAuthorizedDateOk

`func (o *CloudProviderAccessRole) GetAuthorizedDateOk() (*time.Time, bool)`

GetAuthorizedDateOk returns a tuple with the AuthorizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDate

`func (o *CloudProviderAccessRole) SetAuthorizedDate(v time.Time)`

SetAuthorizedDate sets AuthorizedDate field to given value.

### HasAuthorizedDate

`func (o *CloudProviderAccessRole) HasAuthorizedDate() bool`

HasAuthorizedDate returns a boolean if a field has been set.

### SetAuthorizedDateNil

`func (o *CloudProviderAccessRole) SetAuthorizedDateNil()`

SetAuthorizedDateNil sets AuthorizedDate to an explicit JSON null when marshaled, overriding any value previously set with SetAuthorizedDate. Calling SetAuthorizedDate again clears the null override.

### GetCreatedDate

`func (o *CloudProviderAccessRole) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CloudProviderAccessRole) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CloudProviderAccessRole) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CloudProviderAccessRole) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *CloudProviderAccessRole) SetCreatedDateNil()`

SetCreatedDateNil sets CreatedDate to an explicit JSON null when marshaled, overriding any value previously set with SetCreatedDate. Calling SetCreatedDate again clears the null override.

### GetFeatureUsages

`func (o *CloudProviderAccessRole) GetFeatureUsages() []CloudProviderAccessFeatureUsage`

GetFeatureUsages returns the FeatureUsages field if non-nil, zero value otherwise.

### GetFeatureUsagesOk

`func (o *CloudProviderAccessRole) GetFeatureUsagesOk() (*[]CloudProviderAccessFeatureUsage, bool)`

GetFeatureUsagesOk returns a tuple with the FeatureUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsages

`func (o *CloudProviderAccessRole) SetFeatureUsages(v []CloudProviderAccessFeatureUsage)`

SetFeatureUsages sets FeatureUsages field to given value.

### HasFeatureUsages

`func (o *CloudProviderAccessRole) HasFeatureUsages() bool`

HasFeatureUsages returns a boolean if a field has been set.

### SetFeatureUsagesNil

`func (o *CloudProviderAccessRole) SetFeatureUsagesNil()`

SetFeatureUsagesNil sets FeatureUsages to an explicit JSON null when marshaled, overriding any value previously set with SetFeatureUsages. Calling SetFeatureUsages again clears the null override.

### GetIamAssumedRoleArn

`func (o *CloudProviderAccessRole) GetIamAssumedRoleArn() string`

GetIamAssumedRoleArn returns the IamAssumedRoleArn field if non-nil, zero value otherwise.

### GetIamAssumedRoleArnOk

`func (o *CloudProviderAccessRole) GetIamAssumedRoleArnOk() (*string, bool)`

GetIamAssumedRoleArnOk returns a tuple with the IamAssumedRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamAssumedRoleArn

`func (o *CloudProviderAccessRole) SetIamAssumedRoleArn(v string)`

SetIamAssumedRoleArn sets IamAssumedRoleArn field to given value.

### HasIamAssumedRoleArn

`func (o *CloudProviderAccessRole) HasIamAssumedRoleArn() bool`

HasIamAssumedRoleArn returns a boolean if a field has been set.

### SetIamAssumedRoleArnNil

`func (o *CloudProviderAccessRole) SetIamAssumedRoleArnNil()`

SetIamAssumedRoleArnNil sets IamAssumedRoleArn to an explicit JSON null when marshaled, overriding any value previously set with SetIamAssumedRoleArn. Calling SetIamAssumedRoleArn again clears the null override.

### GetRoleId

`func (o *CloudProviderAccessRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CloudProviderAccessRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CloudProviderAccessRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *CloudProviderAccessRole) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### SetRoleIdNil

`func (o *CloudProviderAccessRole) SetRoleIdNil()`

SetRoleIdNil sets RoleId to an explicit JSON null when marshaled, overriding any value previously set with SetRoleId. Calling SetRoleId again clears the null override.

### GetId

`func (o *CloudProviderAccessRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProviderAccessRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProviderAccessRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProviderAccessRole) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CloudProviderAccessRole) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetAtlasAzureAppId

`func (o *CloudProviderAccessRole) GetAtlasAzureAppId() string`

GetAtlasAzureAppId returns the AtlasAzureAppId field if non-nil, zero value otherwise.

### GetAtlasAzureAppIdOk

`func (o *CloudProviderAccessRole) GetAtlasAzureAppIdOk() (*string, bool)`

GetAtlasAzureAppIdOk returns a tuple with the AtlasAzureAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAzureAppId

`func (o *CloudProviderAccessRole) SetAtlasAzureAppId(v string)`

SetAtlasAzureAppId sets AtlasAzureAppId field to given value.

### HasAtlasAzureAppId

`func (o *CloudProviderAccessRole) HasAtlasAzureAppId() bool`

HasAtlasAzureAppId returns a boolean if a field has been set.

### SetAtlasAzureAppIdNil

`func (o *CloudProviderAccessRole) SetAtlasAzureAppIdNil()`

SetAtlasAzureAppIdNil sets AtlasAzureAppId to an explicit JSON null when marshaled, overriding any value previously set with SetAtlasAzureAppId. Calling SetAtlasAzureAppId again clears the null override.

### GetLastUpdatedDate

`func (o *CloudProviderAccessRole) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *CloudProviderAccessRole) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *CloudProviderAccessRole) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *CloudProviderAccessRole) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### SetLastUpdatedDateNil

`func (o *CloudProviderAccessRole) SetLastUpdatedDateNil()`

SetLastUpdatedDateNil sets LastUpdatedDate to an explicit JSON null when marshaled, overriding any value previously set with SetLastUpdatedDate. Calling SetLastUpdatedDate again clears the null override.

### GetServicePrincipalId

`func (o *CloudProviderAccessRole) GetServicePrincipalId() string`

GetServicePrincipalId returns the ServicePrincipalId field if non-nil, zero value otherwise.

### GetServicePrincipalIdOk

`func (o *CloudProviderAccessRole) GetServicePrincipalIdOk() (*string, bool)`

GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalId

`func (o *CloudProviderAccessRole) SetServicePrincipalId(v string)`

SetServicePrincipalId sets ServicePrincipalId field to given value.

### HasServicePrincipalId

`func (o *CloudProviderAccessRole) HasServicePrincipalId() bool`

HasServicePrincipalId returns a boolean if a field has been set.

### SetServicePrincipalIdNil

`func (o *CloudProviderAccessRole) SetServicePrincipalIdNil()`

SetServicePrincipalIdNil sets ServicePrincipalId to an explicit JSON null when marshaled, overriding any value previously set with SetServicePrincipalId. Calling SetServicePrincipalId again clears the null override.

### GetTenantId

`func (o *CloudProviderAccessRole) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudProviderAccessRole) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudProviderAccessRole) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CloudProviderAccessRole) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CloudProviderAccessRole) SetTenantIdNil()`

SetTenantIdNil sets TenantId to an explicit JSON null when marshaled, overriding any value previously set with SetTenantId. Calling SetTenantId again clears the null override.

### GetGcpServiceAccountForAtlas

`func (o *CloudProviderAccessRole) GetGcpServiceAccountForAtlas() string`

GetGcpServiceAccountForAtlas returns the GcpServiceAccountForAtlas field if non-nil, zero value otherwise.

### GetGcpServiceAccountForAtlasOk

`func (o *CloudProviderAccessRole) GetGcpServiceAccountForAtlasOk() (*string, bool)`

GetGcpServiceAccountForAtlasOk returns a tuple with the GcpServiceAccountForAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountForAtlas

`func (o *CloudProviderAccessRole) SetGcpServiceAccountForAtlas(v string)`

SetGcpServiceAccountForAtlas sets GcpServiceAccountForAtlas field to given value.

### HasGcpServiceAccountForAtlas

`func (o *CloudProviderAccessRole) HasGcpServiceAccountForAtlas() bool`

HasGcpServiceAccountForAtlas returns a boolean if a field has been set.

### SetGcpServiceAccountForAtlasNil

`func (o *CloudProviderAccessRole) SetGcpServiceAccountForAtlasNil()`

SetGcpServiceAccountForAtlasNil sets GcpServiceAccountForAtlas to an explicit JSON null when marshaled, overriding any value previously set with SetGcpServiceAccountForAtlas. Calling SetGcpServiceAccountForAtlas again clears the null override.

### GetStatus

`func (o *CloudProviderAccessRole) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudProviderAccessRole) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudProviderAccessRole) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudProviderAccessRole) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CloudProviderAccessRole) SetStatusNil()`

SetStatusNil sets Status to an explicit JSON null when marshaled, overriding any value previously set with SetStatus. Calling SetStatus again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


