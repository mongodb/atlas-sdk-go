# LiveImportAvailableProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | [**[]AvailableClustersDeployment**](AvailableClustersDeployment.md) | List of clusters that can be migrated to MongoDB Cloud. | 
**MigrationHosts** | **[]string** | Hostname of MongoDB Agent list that you configured to perform a migration. | 
**Name** | **string** | Human-readable label that identifies this project. | [readonly] 
**ProjectId** | **string** | Unique 24-hexadecimal digit string that identifies the project to be migrated. | [readonly] 

## Methods

### NewLiveImportAvailableProject

`func NewLiveImportAvailableProject(deployments []AvailableClustersDeployment, migrationHosts []string, name string, projectId string, ) *LiveImportAvailableProject`

NewLiveImportAvailableProject instantiates a new LiveImportAvailableProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveImportAvailableProjectWithDefaults

`func NewLiveImportAvailableProjectWithDefaults() *LiveImportAvailableProject`

NewLiveImportAvailableProjectWithDefaults instantiates a new LiveImportAvailableProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *LiveImportAvailableProject) GetDeployments() []AvailableClustersDeployment`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *LiveImportAvailableProject) GetDeploymentsOk() (*[]AvailableClustersDeployment, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *LiveImportAvailableProject) SetDeployments(v []AvailableClustersDeployment)`

SetDeployments sets Deployments field to given value.

### GetMigrationHosts

`func (o *LiveImportAvailableProject) GetMigrationHosts() []string`

GetMigrationHosts returns the MigrationHosts field if non-nil, zero value otherwise.

### GetMigrationHostsOk

`func (o *LiveImportAvailableProject) GetMigrationHostsOk() (*[]string, bool)`

GetMigrationHostsOk returns a tuple with the MigrationHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationHosts

`func (o *LiveImportAvailableProject) SetMigrationHosts(v []string)`

SetMigrationHosts sets MigrationHosts field to given value.

### GetName

`func (o *LiveImportAvailableProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LiveImportAvailableProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LiveImportAvailableProject) SetName(v string)`

SetName sets Name field to given value.

### GetProjectId

`func (o *LiveImportAvailableProject) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LiveImportAvailableProject) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LiveImportAvailableProject) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


