# CloudProviderAccessRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Human-readable label that identifies the cloud provider of the role. | 
**AtlasAWSAccountArn** | Pointer to **string** | Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role. | [optional] [readonly] 
**AtlasAssumedRoleExternalId** | Pointer to **string** | Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account. | [optional] [readonly] 
**AuthorizedDate** | Pointer to **time.Time** | Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** | Date and time when someone created this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.  Alternatively: Date and time when this Azure Service Principal was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC.  Alternatively: Date and time when this GCP Service Account was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FeatureUsages** | Pointer to [**[]CloudProviderAccessFeatureUsage**](CloudProviderAccessFeatureUsage.md) | List that contains application features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role.  Alternatively: List that contains application features associated with this Azure Service Principal.  Alternatively: List that contains application features associated with this GCP Service Account. | [optional] [readonly] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**AtlasAzureAppId** | Pointer to **string** | Azure Active Directory Application ID of Atlas. | [optional] 
**LastUpdatedDate** | Pointer to **time.Time** | Date and time when this Azure Service Principal was last updated. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**ServicePrincipalId** | Pointer to **string** | UUID string that identifies the Azure Service Principal. | [optional] 
**TenantId** | Pointer to **string** | UUID String that identifies the Azure Active Directory Tenant ID. | [optional] 
**GcpServiceAccountForAtlas** | Pointer to **string** | ID string that identifies the GCP Service Account used by Atlas. | [optional] [readonly] 

## Methods

### NewCloudProviderAccessRoleRequest

`func NewCloudProviderAccessRoleRequest(providerName string, ) *CloudProviderAccessRoleRequest`

NewCloudProviderAccessRoleRequest instantiates a new CloudProviderAccessRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessRoleRequestWithDefaults

`func NewCloudProviderAccessRoleRequestWithDefaults() *CloudProviderAccessRoleRequest`

NewCloudProviderAccessRoleRequestWithDefaults instantiates a new CloudProviderAccessRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *CloudProviderAccessRoleRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CloudProviderAccessRoleRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CloudProviderAccessRoleRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### GetAtlasAWSAccountArn

`func (o *CloudProviderAccessRoleRequest) GetAtlasAWSAccountArn() string`

GetAtlasAWSAccountArn returns the AtlasAWSAccountArn field if non-nil, zero value otherwise.

### GetAtlasAWSAccountArnOk

`func (o *CloudProviderAccessRoleRequest) GetAtlasAWSAccountArnOk() (*string, bool)`

GetAtlasAWSAccountArnOk returns a tuple with the AtlasAWSAccountArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAWSAccountArn

`func (o *CloudProviderAccessRoleRequest) SetAtlasAWSAccountArn(v string)`

SetAtlasAWSAccountArn sets AtlasAWSAccountArn field to given value.

### HasAtlasAWSAccountArn

`func (o *CloudProviderAccessRoleRequest) HasAtlasAWSAccountArn() bool`

HasAtlasAWSAccountArn returns a boolean if a field has been set.
### GetAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessRoleRequest) GetAtlasAssumedRoleExternalId() string`

GetAtlasAssumedRoleExternalId returns the AtlasAssumedRoleExternalId field if non-nil, zero value otherwise.

### GetAtlasAssumedRoleExternalIdOk

`func (o *CloudProviderAccessRoleRequest) GetAtlasAssumedRoleExternalIdOk() (*string, bool)`

GetAtlasAssumedRoleExternalIdOk returns a tuple with the AtlasAssumedRoleExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessRoleRequest) SetAtlasAssumedRoleExternalId(v string)`

SetAtlasAssumedRoleExternalId sets AtlasAssumedRoleExternalId field to given value.

### HasAtlasAssumedRoleExternalId

`func (o *CloudProviderAccessRoleRequest) HasAtlasAssumedRoleExternalId() bool`

HasAtlasAssumedRoleExternalId returns a boolean if a field has been set.
### GetAuthorizedDate

`func (o *CloudProviderAccessRoleRequest) GetAuthorizedDate() time.Time`

GetAuthorizedDate returns the AuthorizedDate field if non-nil, zero value otherwise.

### GetAuthorizedDateOk

`func (o *CloudProviderAccessRoleRequest) GetAuthorizedDateOk() (*time.Time, bool)`

GetAuthorizedDateOk returns a tuple with the AuthorizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDate

`func (o *CloudProviderAccessRoleRequest) SetAuthorizedDate(v time.Time)`

SetAuthorizedDate sets AuthorizedDate field to given value.

### HasAuthorizedDate

`func (o *CloudProviderAccessRoleRequest) HasAuthorizedDate() bool`

HasAuthorizedDate returns a boolean if a field has been set.
### GetCreatedDate

`func (o *CloudProviderAccessRoleRequest) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CloudProviderAccessRoleRequest) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CloudProviderAccessRoleRequest) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CloudProviderAccessRoleRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.
### GetFeatureUsages

`func (o *CloudProviderAccessRoleRequest) GetFeatureUsages() []CloudProviderAccessFeatureUsage`

GetFeatureUsages returns the FeatureUsages field if non-nil, zero value otherwise.

### GetFeatureUsagesOk

`func (o *CloudProviderAccessRoleRequest) GetFeatureUsagesOk() (*[]CloudProviderAccessFeatureUsage, bool)`

GetFeatureUsagesOk returns a tuple with the FeatureUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsages

`func (o *CloudProviderAccessRoleRequest) SetFeatureUsages(v []CloudProviderAccessFeatureUsage)`

SetFeatureUsages sets FeatureUsages field to given value.

### HasFeatureUsages

`func (o *CloudProviderAccessRoleRequest) HasFeatureUsages() bool`

HasFeatureUsages returns a boolean if a field has been set.
### GetRoleId

`func (o *CloudProviderAccessRoleRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CloudProviderAccessRoleRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CloudProviderAccessRoleRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *CloudProviderAccessRoleRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetId

`func (o *CloudProviderAccessRoleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProviderAccessRoleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProviderAccessRoleRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProviderAccessRoleRequest) HasId() bool`

HasId returns a boolean if a field has been set.
### GetAtlasAzureAppId

`func (o *CloudProviderAccessRoleRequest) GetAtlasAzureAppId() string`

GetAtlasAzureAppId returns the AtlasAzureAppId field if non-nil, zero value otherwise.

### GetAtlasAzureAppIdOk

`func (o *CloudProviderAccessRoleRequest) GetAtlasAzureAppIdOk() (*string, bool)`

GetAtlasAzureAppIdOk returns a tuple with the AtlasAzureAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAzureAppId

`func (o *CloudProviderAccessRoleRequest) SetAtlasAzureAppId(v string)`

SetAtlasAzureAppId sets AtlasAzureAppId field to given value.

### HasAtlasAzureAppId

`func (o *CloudProviderAccessRoleRequest) HasAtlasAzureAppId() bool`

HasAtlasAzureAppId returns a boolean if a field has been set.
### GetLastUpdatedDate

`func (o *CloudProviderAccessRoleRequest) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *CloudProviderAccessRoleRequest) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *CloudProviderAccessRoleRequest) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *CloudProviderAccessRoleRequest) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.
### GetServicePrincipalId

`func (o *CloudProviderAccessRoleRequest) GetServicePrincipalId() string`

GetServicePrincipalId returns the ServicePrincipalId field if non-nil, zero value otherwise.

### GetServicePrincipalIdOk

`func (o *CloudProviderAccessRoleRequest) GetServicePrincipalIdOk() (*string, bool)`

GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalId

`func (o *CloudProviderAccessRoleRequest) SetServicePrincipalId(v string)`

SetServicePrincipalId sets ServicePrincipalId field to given value.

### HasServicePrincipalId

`func (o *CloudProviderAccessRoleRequest) HasServicePrincipalId() bool`

HasServicePrincipalId returns a boolean if a field has been set.
### GetTenantId

`func (o *CloudProviderAccessRoleRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudProviderAccessRoleRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudProviderAccessRoleRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CloudProviderAccessRoleRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.
### GetGcpServiceAccountForAtlas

`func (o *CloudProviderAccessRoleRequest) GetGcpServiceAccountForAtlas() string`

GetGcpServiceAccountForAtlas returns the GcpServiceAccountForAtlas field if non-nil, zero value otherwise.

### GetGcpServiceAccountForAtlasOk

`func (o *CloudProviderAccessRoleRequest) GetGcpServiceAccountForAtlasOk() (*string, bool)`

GetGcpServiceAccountForAtlasOk returns a tuple with the GcpServiceAccountForAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountForAtlas

`func (o *CloudProviderAccessRoleRequest) SetGcpServiceAccountForAtlas(v string)`

SetGcpServiceAccountForAtlas sets GcpServiceAccountForAtlas field to given value.

### HasGcpServiceAccountForAtlas

`func (o *CloudProviderAccessRoleRequest) HasGcpServiceAccountForAtlas() bool`

HasGcpServiceAccountForAtlas returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


