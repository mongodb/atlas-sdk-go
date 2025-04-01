# AzureKeyVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureEnvironment** | Pointer to **string** | Azure environment in which your account credentials reside. | [optional] 
**ClientID** | Pointer to **string** | Unique 36-hexadecimal character string that identifies an Azure application associated with your Azure Active Directory tenant. | [optional] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled encryption at rest for the specified  project. To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of &#x60;false&#x60;. | [optional] 
**KeyIdentifier** | Pointer to **string** | Web address with a unique key that identifies for your Azure Key Vault. | [optional] 
**KeyVaultName** | Pointer to **string** | Unique string that identifies the Azure Key Vault that contains your key. This field cannot be modified when you enable and set up private endpoint connections to your Azure Key Vault. | [optional] 
**RequirePrivateNetworking** | Pointer to **bool** | Enable connection to your Azure Key Vault over private networking. | [optional] 
**ResourceGroupName** | Pointer to **string** | Name of the Azure resource group that contains your Azure Key Vault. This field cannot be modified when you enable and set up private endpoint connections to your Azure Key Vault. | [optional] 
**Secret** | Pointer to **string** | Private data that you need secured and that belongs to the specified Azure Key Vault (AKV) tenant (**azureKeyVault.tenantID**). This data can include any type of sensitive data such as passwords, database connection strings, API keys, and the like. AKV stores this information as encrypted binary data. | [optional] 
**SubscriptionID** | Pointer to **string** | Unique 36-hexadecimal character string that identifies your Azure subscription. This field cannot be modified when you enable and set up private endpoint connections to your Azure Key Vault. | [optional] 
**TenantID** | Pointer to **string** | Unique 36-hexadecimal character string that identifies the Azure Active Directory tenant within your Azure subscription. | [optional] 
**Valid** | Pointer to **bool** | Flag that indicates whether the Azure encryption key can encrypt and decrypt data. | [optional] [readonly] 

## Methods

### NewAzureKeyVault

`func NewAzureKeyVault() *AzureKeyVault`

NewAzureKeyVault instantiates a new AzureKeyVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultWithDefaults

`func NewAzureKeyVaultWithDefaults() *AzureKeyVault`

NewAzureKeyVaultWithDefaults instantiates a new AzureKeyVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureEnvironment

`func (o *AzureKeyVault) GetAzureEnvironment() string`

GetAzureEnvironment returns the AzureEnvironment field if non-nil, zero value otherwise.

### GetAzureEnvironmentOk

`func (o *AzureKeyVault) GetAzureEnvironmentOk() (*string, bool)`

GetAzureEnvironmentOk returns a tuple with the AzureEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEnvironment

`func (o *AzureKeyVault) SetAzureEnvironment(v string)`

SetAzureEnvironment sets AzureEnvironment field to given value.

### HasAzureEnvironment

`func (o *AzureKeyVault) HasAzureEnvironment() bool`

HasAzureEnvironment returns a boolean if a field has been set.
### GetClientID

`func (o *AzureKeyVault) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AzureKeyVault) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AzureKeyVault) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *AzureKeyVault) HasClientID() bool`

HasClientID returns a boolean if a field has been set.
### GetEnabled

`func (o *AzureKeyVault) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureKeyVault) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureKeyVault) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AzureKeyVault) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.
### GetKeyIdentifier

`func (o *AzureKeyVault) GetKeyIdentifier() string`

GetKeyIdentifier returns the KeyIdentifier field if non-nil, zero value otherwise.

### GetKeyIdentifierOk

`func (o *AzureKeyVault) GetKeyIdentifierOk() (*string, bool)`

GetKeyIdentifierOk returns a tuple with the KeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIdentifier

`func (o *AzureKeyVault) SetKeyIdentifier(v string)`

SetKeyIdentifier sets KeyIdentifier field to given value.

### HasKeyIdentifier

`func (o *AzureKeyVault) HasKeyIdentifier() bool`

HasKeyIdentifier returns a boolean if a field has been set.
### GetKeyVaultName

`func (o *AzureKeyVault) GetKeyVaultName() string`

GetKeyVaultName returns the KeyVaultName field if non-nil, zero value otherwise.

### GetKeyVaultNameOk

`func (o *AzureKeyVault) GetKeyVaultNameOk() (*string, bool)`

GetKeyVaultNameOk returns a tuple with the KeyVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultName

`func (o *AzureKeyVault) SetKeyVaultName(v string)`

SetKeyVaultName sets KeyVaultName field to given value.

### HasKeyVaultName

`func (o *AzureKeyVault) HasKeyVaultName() bool`

HasKeyVaultName returns a boolean if a field has been set.
### GetRequirePrivateNetworking

`func (o *AzureKeyVault) GetRequirePrivateNetworking() bool`

GetRequirePrivateNetworking returns the RequirePrivateNetworking field if non-nil, zero value otherwise.

### GetRequirePrivateNetworkingOk

`func (o *AzureKeyVault) GetRequirePrivateNetworkingOk() (*bool, bool)`

GetRequirePrivateNetworkingOk returns a tuple with the RequirePrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePrivateNetworking

`func (o *AzureKeyVault) SetRequirePrivateNetworking(v bool)`

SetRequirePrivateNetworking sets RequirePrivateNetworking field to given value.

### HasRequirePrivateNetworking

`func (o *AzureKeyVault) HasRequirePrivateNetworking() bool`

HasRequirePrivateNetworking returns a boolean if a field has been set.
### GetResourceGroupName

`func (o *AzureKeyVault) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureKeyVault) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureKeyVault) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureKeyVault) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.
### GetSecret

`func (o *AzureKeyVault) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AzureKeyVault) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AzureKeyVault) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AzureKeyVault) HasSecret() bool`

HasSecret returns a boolean if a field has been set.
### GetSubscriptionID

`func (o *AzureKeyVault) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *AzureKeyVault) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *AzureKeyVault) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.

### HasSubscriptionID

`func (o *AzureKeyVault) HasSubscriptionID() bool`

HasSubscriptionID returns a boolean if a field has been set.
### GetTenantID

`func (o *AzureKeyVault) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *AzureKeyVault) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *AzureKeyVault) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *AzureKeyVault) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.
### GetValid

`func (o *AzureKeyVault) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AzureKeyVault) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AzureKeyVault) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AzureKeyVault) HasValid() bool`

HasValid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


