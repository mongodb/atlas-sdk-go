# DataFederationAzureCloudProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasAppId** | Pointer to **string** | The App ID generated by Atlas for the Service Principal&#39;s access policy. | [optional] [readonly] 
**RoleId** | **string** | Unique identifier of the role that Data Federation can use to access the data stores. Required if specifying cloudProviderConfig. | 
**ServicePrincipalId** | Pointer to **string** | The ID of the Service Principal for which there is an access policy for Atlas to access Azure resources. | [optional] [readonly] 
**TenantId** | Pointer to **string** | The Azure Active Directory / Entra ID tenant ID associated with the Service Principal. | [optional] [readonly] 

## Methods

### NewDataFederationAzureCloudProviderConfig

`func NewDataFederationAzureCloudProviderConfig(roleId string, ) *DataFederationAzureCloudProviderConfig`

NewDataFederationAzureCloudProviderConfig instantiates a new DataFederationAzureCloudProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataFederationAzureCloudProviderConfigWithDefaults

`func NewDataFederationAzureCloudProviderConfigWithDefaults() *DataFederationAzureCloudProviderConfig`

NewDataFederationAzureCloudProviderConfigWithDefaults instantiates a new DataFederationAzureCloudProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasAppId

`func (o *DataFederationAzureCloudProviderConfig) GetAtlasAppId() string`

GetAtlasAppId returns the AtlasAppId field if non-nil, zero value otherwise.

### GetAtlasAppIdOk

`func (o *DataFederationAzureCloudProviderConfig) GetAtlasAppIdOk() (*string, bool)`

GetAtlasAppIdOk returns a tuple with the AtlasAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAppId

`func (o *DataFederationAzureCloudProviderConfig) SetAtlasAppId(v string)`

SetAtlasAppId sets AtlasAppId field to given value.

### HasAtlasAppId

`func (o *DataFederationAzureCloudProviderConfig) HasAtlasAppId() bool`

HasAtlasAppId returns a boolean if a field has been set.
### GetRoleId

`func (o *DataFederationAzureCloudProviderConfig) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DataFederationAzureCloudProviderConfig) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DataFederationAzureCloudProviderConfig) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### GetServicePrincipalId

`func (o *DataFederationAzureCloudProviderConfig) GetServicePrincipalId() string`

GetServicePrincipalId returns the ServicePrincipalId field if non-nil, zero value otherwise.

### GetServicePrincipalIdOk

`func (o *DataFederationAzureCloudProviderConfig) GetServicePrincipalIdOk() (*string, bool)`

GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalId

`func (o *DataFederationAzureCloudProviderConfig) SetServicePrincipalId(v string)`

SetServicePrincipalId sets ServicePrincipalId field to given value.

### HasServicePrincipalId

`func (o *DataFederationAzureCloudProviderConfig) HasServicePrincipalId() bool`

HasServicePrincipalId returns a boolean if a field has been set.
### GetTenantId

`func (o *DataFederationAzureCloudProviderConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DataFederationAzureCloudProviderConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DataFederationAzureCloudProviderConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DataFederationAzureCloudProviderConfig) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


