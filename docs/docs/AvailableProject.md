# AvailableProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | Pointer to [**[]AvailableDeployment**](AvailableDeployment.md) | List of clusters that can be migrated to MongoDB Cloud. | [optional] 
**MigrationHosts** | Pointer to **[]string** | Hostname of MongoDB Agent list that you configured to perform a migration. | [optional] 
**Name** | **string** | Human-readable label that identifies this project. | [readonly] 
**ProjectId** | **string** | Unique 24-hexadecimal digit string that identifies the project to be migrated. | [readonly] 

## Methods

### NewAvailableProject

`func NewAvailableProject(name string, projectId string, ) *AvailableProject`

NewAvailableProject instantiates a new AvailableProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableProjectWithDefaults

`func NewAvailableProjectWithDefaults() *AvailableProject`

NewAvailableProjectWithDefaults instantiates a new AvailableProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *AvailableProject) GetDeployments() []AvailableDeployment`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *AvailableProject) GetDeploymentsOk() (*[]AvailableDeployment, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *AvailableProject) SetDeployments(v []AvailableDeployment)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *AvailableProject) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetMigrationHosts

`func (o *AvailableProject) GetMigrationHosts() []string`

GetMigrationHosts returns the MigrationHosts field if non-nil, zero value otherwise.

### GetMigrationHostsOk

`func (o *AvailableProject) GetMigrationHostsOk() (*[]string, bool)`

GetMigrationHostsOk returns a tuple with the MigrationHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationHosts

`func (o *AvailableProject) SetMigrationHosts(v []string)`

SetMigrationHosts sets MigrationHosts field to given value.

### HasMigrationHosts

`func (o *AvailableProject) HasMigrationHosts() bool`

HasMigrationHosts returns a boolean if a field has been set.

### GetName

`func (o *AvailableProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableProject) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *AvailableProject) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AvailableProject) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AvailableProject) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


