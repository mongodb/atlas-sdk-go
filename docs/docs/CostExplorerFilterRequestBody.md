# CostExplorerFilterRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to **[]string** | The list of clusters to be included in the Cost Explorer Query. | [optional] 
**EndDate** | **string** | The exclusive ending date for the Cost Explorer query. The date must be the start of a month. | 
**GroupBy** | Pointer to **string** | The dimension to group the returned usage results by. At least one filter value needs to be provided for a dimension to be used. | [optional] 
**IncludePartialMatches** | Pointer to **bool** | Flag to control whether usage that matches the filter criteria, but does not have values for all filter criteria is included in response. Default is false, which excludes the partially matching data. | [optional] 
**Organizations** | Pointer to **[]string** | The list of organizations to be included in the Cost Explorer Query. | [optional] 
**Projects** | Pointer to **[]string** | The list of projects to be included in the Cost Explorer Query. | [optional] 
**Services** | Pointer to **[]string** | The list of projects to be included in the Cost Explorer Query. | [optional] 
**StartDate** | **string** | The inclusive starting date for the Cost Explorer query. The date must be the start of a month. | 

## Methods

### NewCostExplorerFilterRequestBody

`func NewCostExplorerFilterRequestBody(endDate string, startDate string, ) *CostExplorerFilterRequestBody`

NewCostExplorerFilterRequestBody instantiates a new CostExplorerFilterRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostExplorerFilterRequestBodyWithDefaults

`func NewCostExplorerFilterRequestBodyWithDefaults() *CostExplorerFilterRequestBody`

NewCostExplorerFilterRequestBodyWithDefaults instantiates a new CostExplorerFilterRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *CostExplorerFilterRequestBody) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *CostExplorerFilterRequestBody) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *CostExplorerFilterRequestBody) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *CostExplorerFilterRequestBody) HasClusters() bool`

HasClusters returns a boolean if a field has been set.
### GetEndDate

`func (o *CostExplorerFilterRequestBody) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CostExplorerFilterRequestBody) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CostExplorerFilterRequestBody) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### GetGroupBy

`func (o *CostExplorerFilterRequestBody) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *CostExplorerFilterRequestBody) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *CostExplorerFilterRequestBody) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *CostExplorerFilterRequestBody) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.
### GetIncludePartialMatches

`func (o *CostExplorerFilterRequestBody) GetIncludePartialMatches() bool`

GetIncludePartialMatches returns the IncludePartialMatches field if non-nil, zero value otherwise.

### GetIncludePartialMatchesOk

`func (o *CostExplorerFilterRequestBody) GetIncludePartialMatchesOk() (*bool, bool)`

GetIncludePartialMatchesOk returns a tuple with the IncludePartialMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePartialMatches

`func (o *CostExplorerFilterRequestBody) SetIncludePartialMatches(v bool)`

SetIncludePartialMatches sets IncludePartialMatches field to given value.

### HasIncludePartialMatches

`func (o *CostExplorerFilterRequestBody) HasIncludePartialMatches() bool`

HasIncludePartialMatches returns a boolean if a field has been set.
### GetOrganizations

`func (o *CostExplorerFilterRequestBody) GetOrganizations() []string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *CostExplorerFilterRequestBody) GetOrganizationsOk() (*[]string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *CostExplorerFilterRequestBody) SetOrganizations(v []string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *CostExplorerFilterRequestBody) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.
### GetProjects

`func (o *CostExplorerFilterRequestBody) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CostExplorerFilterRequestBody) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CostExplorerFilterRequestBody) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CostExplorerFilterRequestBody) HasProjects() bool`

HasProjects returns a boolean if a field has been set.
### GetServices

`func (o *CostExplorerFilterRequestBody) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CostExplorerFilterRequestBody) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CostExplorerFilterRequestBody) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *CostExplorerFilterRequestBody) HasServices() bool`

HasServices returns a boolean if a field has been set.
### GetStartDate

`func (o *CostExplorerFilterRequestBody) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CostExplorerFilterRequestBody) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CostExplorerFilterRequestBody) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


