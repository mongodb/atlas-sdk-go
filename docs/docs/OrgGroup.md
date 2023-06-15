# OrgGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]CloudCluster**](CloudCluster.md) | Settings that describe the clusters in each project that the API key is authorized to view. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**GroupName** | Pointer to **string** | Human-readable label that identifies the project. | [optional] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the organization that contains the project. | [optional] [readonly] 
**OrgName** | Pointer to **string** | Human-readable label that identifies the organization that contains the project. | [optional] 
**PlanType** | Pointer to **string** | Human-readable label that indicates the plan type. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | List of human-readable labels that categorize the specified project. MongoDB Cloud returns an empty array. | [optional] [readonly] 

## Methods

### NewOrgGroup

`func NewOrgGroup() *OrgGroup`

NewOrgGroup instantiates a new OrgGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgGroupWithDefaults

`func NewOrgGroupWithDefaults() *OrgGroup`

NewOrgGroupWithDefaults instantiates a new OrgGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *OrgGroup) GetClusters() []CloudCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *OrgGroup) GetClustersOk() (*[]CloudCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *OrgGroup) SetClusters(v []CloudCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *OrgGroup) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetGroupId

`func (o *OrgGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OrgGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OrgGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *OrgGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *OrgGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *OrgGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *OrgGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *OrgGroup) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetOrgId

`func (o *OrgGroup) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgGroup) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgGroup) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgGroup) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *OrgGroup) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OrgGroup) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OrgGroup) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *OrgGroup) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPlanType

`func (o *OrgGroup) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *OrgGroup) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *OrgGroup) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *OrgGroup) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetTags

`func (o *OrgGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrgGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrgGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrgGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


