# ClusterOutageSimulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster that undergoes outage simulation. | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud expires the outage simulation. This parameter expresses its value in the ISO 8601 timestamp format in UTC. If not provided, defaults to 3 days from the start date. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project that contains the cluster to undergo outage simulation. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the outage simulation. | [optional] [readonly] 
**OutageFilters** | Pointer to [**[]AtlasClusterOutageSimulationOutageFilter**](AtlasClusterOutageSimulationOutageFilter.md) | List of settings that specify the type of cluster outage simulation. | [optional] 
**StartRequestDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud started the regional outage simulation. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**State** | Pointer to **string** | Phase of the outage simulation.  | State       | Indication | |-------------|------------| | &#x60;START_REQUESTED&#x60;    | User has requested cluster outage simulation.| | &#x60;STARTING&#x60;           | MongoDB Cloud is starting cluster outage simulation.| | &#x60;SIMULATING&#x60;         | MongoDB Cloud is simulating cluster outage.| | &#x60;RECOVERY_REQUESTED&#x60; | User has requested recovery from the simulated outage.| | &#x60;RECOVERING&#x60;         | MongoDB Cloud is recovering the cluster from the simulated outage.| | &#x60;COMPLETE&#x60;           | MongoDB Cloud has completed the cluster outage simulation.| | [optional] [readonly] 

## Methods

### NewClusterOutageSimulation

`func NewClusterOutageSimulation() *ClusterOutageSimulation`

NewClusterOutageSimulation instantiates a new ClusterOutageSimulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOutageSimulationWithDefaults

`func NewClusterOutageSimulationWithDefaults() *ClusterOutageSimulation`

NewClusterOutageSimulationWithDefaults instantiates a new ClusterOutageSimulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ClusterOutageSimulation) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterOutageSimulation) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterOutageSimulation) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ClusterOutageSimulation) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetExpirationDate

`func (o *ClusterOutageSimulation) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ClusterOutageSimulation) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ClusterOutageSimulation) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ClusterOutageSimulation) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.
### GetGroupId

`func (o *ClusterOutageSimulation) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ClusterOutageSimulation) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ClusterOutageSimulation) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ClusterOutageSimulation) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *ClusterOutageSimulation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterOutageSimulation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterOutageSimulation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterOutageSimulation) HasId() bool`

HasId returns a boolean if a field has been set.
### GetOutageFilters

`func (o *ClusterOutageSimulation) GetOutageFilters() []AtlasClusterOutageSimulationOutageFilter`

GetOutageFilters returns the OutageFilters field if non-nil, zero value otherwise.

### GetOutageFiltersOk

`func (o *ClusterOutageSimulation) GetOutageFiltersOk() (*[]AtlasClusterOutageSimulationOutageFilter, bool)`

GetOutageFiltersOk returns a tuple with the OutageFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageFilters

`func (o *ClusterOutageSimulation) SetOutageFilters(v []AtlasClusterOutageSimulationOutageFilter)`

SetOutageFilters sets OutageFilters field to given value.

### HasOutageFilters

`func (o *ClusterOutageSimulation) HasOutageFilters() bool`

HasOutageFilters returns a boolean if a field has been set.
### GetStartRequestDate

`func (o *ClusterOutageSimulation) GetStartRequestDate() time.Time`

GetStartRequestDate returns the StartRequestDate field if non-nil, zero value otherwise.

### GetStartRequestDateOk

`func (o *ClusterOutageSimulation) GetStartRequestDateOk() (*time.Time, bool)`

GetStartRequestDateOk returns a tuple with the StartRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRequestDate

`func (o *ClusterOutageSimulation) SetStartRequestDate(v time.Time)`

SetStartRequestDate sets StartRequestDate field to given value.

### HasStartRequestDate

`func (o *ClusterOutageSimulation) HasStartRequestDate() bool`

HasStartRequestDate returns a boolean if a field has been set.
### GetState

`func (o *ClusterOutageSimulation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterOutageSimulation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterOutageSimulation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ClusterOutageSimulation) HasState() bool`

HasState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


