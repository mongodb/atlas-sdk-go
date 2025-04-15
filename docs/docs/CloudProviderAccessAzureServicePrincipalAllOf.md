# CloudProviderAccessAzureServicePrincipalAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**AtlasAzureAppId** | Pointer to **string** | Azure Active Directory Application ID of Atlas. | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date and time when someone created this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.  Alternatively: Date and time when this Azure Service Principal was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FeatureUsages** | Pointer to [**[]CloudProviderAccessFeatureUsage**](CloudProviderAccessFeatureUsage.md) | List that contains application features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role.  Alternatively: List that contains application features associated with this Azure Service Principal. | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **time.Time** | Date and time when this Azure Service Principal was last updated. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**ServicePrincipalId** | Pointer to **string** | UUID string that identifies the Azure Service Principal. | [optional] 
**TenantId** | Pointer to **string** | UUID String that identifies the Azure Active Directory Tenant ID. | [optional] 

## Methods

### NewCloudProviderAccessAzureServicePrincipalAllOf

`func NewCloudProviderAccessAzureServicePrincipalAllOf() *CloudProviderAccessAzureServicePrincipalAllOf`

NewCloudProviderAccessAzureServicePrincipalAllOf instantiates a new CloudProviderAccessAzureServicePrincipalAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessAzureServicePrincipalAllOfWithDefaults

`func NewCloudProviderAccessAzureServicePrincipalAllOfWithDefaults() *CloudProviderAccessAzureServicePrincipalAllOf`

NewCloudProviderAccessAzureServicePrincipalAllOfWithDefaults instantiates a new CloudProviderAccessAzureServicePrincipalAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasId() bool`

HasId returns a boolean if a field has been set.
### GetAtlasAzureAppId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetAtlasAzureAppId() string`

GetAtlasAzureAppId returns the AtlasAzureAppId field if non-nil, zero value otherwise.

### GetAtlasAzureAppIdOk

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetAtlasAzureAppIdOk() (*string, bool)`

GetAtlasAzureAppIdOk returns a tuple with the AtlasAzureAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAzureAppId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetAtlasAzureAppId(v string)`

SetAtlasAzureAppId sets AtlasAzureAppId field to given value.

### HasAtlasAzureAppId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasAtlasAzureAppId() bool`

HasAtlasAzureAppId returns a boolean if a field has been set.
### GetCreatedDate

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.
### GetFeatureUsages

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetFeatureUsages() []CloudProviderAccessFeatureUsage`

GetFeatureUsages returns the FeatureUsages field if non-nil, zero value otherwise.

### GetFeatureUsagesOk

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetFeatureUsagesOk() (*[]CloudProviderAccessFeatureUsage, bool)`

GetFeatureUsagesOk returns a tuple with the FeatureUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsages

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetFeatureUsages(v []CloudProviderAccessFeatureUsage)`

SetFeatureUsages sets FeatureUsages field to given value.

### HasFeatureUsages

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasFeatureUsages() bool`

HasFeatureUsages returns a boolean if a field has been set.
### GetLastUpdatedDate

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.
### GetServicePrincipalId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetServicePrincipalId() string`

GetServicePrincipalId returns the ServicePrincipalId field if non-nil, zero value otherwise.

### GetServicePrincipalIdOk

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetServicePrincipalIdOk() (*string, bool)`

GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetServicePrincipalId(v string)`

SetServicePrincipalId sets ServicePrincipalId field to given value.

### HasServicePrincipalId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasServicePrincipalId() bool`

HasServicePrincipalId returns a boolean if a field has been set.
### GetTenantId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


