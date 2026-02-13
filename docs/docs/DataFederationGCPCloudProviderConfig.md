# DataFederationGCPCloudProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcpServiceAccount** | Pointer to **string** | The email address of the Google Cloud Platform (GCP) service account created by Atlas which should be authorized to allow Atlas to access Google Cloud Storage. | [optional] [readonly] 
**RoleId** | **string** | Unique identifier of the role that Data Federation can use to access the data stores. Required if specifying &#x60;cloudProviderConfig&#x60;. | 

## Methods

### NewDataFederationGCPCloudProviderConfig

`func NewDataFederationGCPCloudProviderConfig(roleId string, ) *DataFederationGCPCloudProviderConfig`

NewDataFederationGCPCloudProviderConfig instantiates a new DataFederationGCPCloudProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataFederationGCPCloudProviderConfigWithDefaults

`func NewDataFederationGCPCloudProviderConfigWithDefaults() *DataFederationGCPCloudProviderConfig`

NewDataFederationGCPCloudProviderConfigWithDefaults instantiates a new DataFederationGCPCloudProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcpServiceAccount

`func (o *DataFederationGCPCloudProviderConfig) GetGcpServiceAccount() string`

GetGcpServiceAccount returns the GcpServiceAccount field if non-nil, zero value otherwise.

### GetGcpServiceAccountOk

`func (o *DataFederationGCPCloudProviderConfig) GetGcpServiceAccountOk() (*string, bool)`

GetGcpServiceAccountOk returns a tuple with the GcpServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccount

`func (o *DataFederationGCPCloudProviderConfig) SetGcpServiceAccount(v string)`

SetGcpServiceAccount sets GcpServiceAccount field to given value.

### HasGcpServiceAccount

`func (o *DataFederationGCPCloudProviderConfig) HasGcpServiceAccount() bool`

HasGcpServiceAccount returns a boolean if a field has been set.
### GetRoleId

`func (o *DataFederationGCPCloudProviderConfig) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DataFederationGCPCloudProviderConfig) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DataFederationGCPCloudProviderConfig) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


