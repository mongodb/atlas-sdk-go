# OverloadProtectionSimulationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelRequestedAt** | Pointer to **time.Time** | Date and time when cancellation of the overload protection simulation was requested. This parameter is only present when a cancellation has been requested and expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**ClusterName** | **string** | Human-readable label that identifies the cluster on which the simulation is running. | [readonly] 
**DurationSeconds** | **int** | Duration of the overload protection simulation in seconds. | 
**ExpiresAt** | **time.Time** | Date and time when the overload protection simulation expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [readonly] 
**GroupId** | **string** | Unique 24-hexadecimal character string that identifies the project that contains the cluster. | [readonly] 
**RequestDate** | **time.Time** | Date and time when the overload protection simulation was requested. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [readonly] 
**SimulationId** | **string** | Unique identifier of the overload protection simulation. | [readonly] 
**State** | **string** | Current state of the overload protection simulation. | [readonly] 

## Methods

### NewOverloadProtectionSimulationResponse

`func NewOverloadProtectionSimulationResponse(clusterName string, durationSeconds int, expiresAt time.Time, groupId string, requestDate time.Time, simulationId string, state string, ) *OverloadProtectionSimulationResponse`

NewOverloadProtectionSimulationResponse instantiates a new OverloadProtectionSimulationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverloadProtectionSimulationResponseWithDefaults

`func NewOverloadProtectionSimulationResponseWithDefaults() *OverloadProtectionSimulationResponse`

NewOverloadProtectionSimulationResponseWithDefaults instantiates a new OverloadProtectionSimulationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelRequestedAt

`func (o *OverloadProtectionSimulationResponse) GetCancelRequestedAt() time.Time`

GetCancelRequestedAt returns the CancelRequestedAt field if non-nil, zero value otherwise.

### GetCancelRequestedAtOk

`func (o *OverloadProtectionSimulationResponse) GetCancelRequestedAtOk() (*time.Time, bool)`

GetCancelRequestedAtOk returns a tuple with the CancelRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRequestedAt

`func (o *OverloadProtectionSimulationResponse) SetCancelRequestedAt(v time.Time)`

SetCancelRequestedAt sets CancelRequestedAt field to given value.

### HasCancelRequestedAt

`func (o *OverloadProtectionSimulationResponse) HasCancelRequestedAt() bool`

HasCancelRequestedAt returns a boolean if a field has been set.

### SetCancelRequestedAtNil

`func (o *OverloadProtectionSimulationResponse) SetCancelRequestedAtNil()`

SetCancelRequestedAtNil sets CancelRequestedAt to an explicit JSON null when marshaled, overriding any value previously set with SetCancelRequestedAt. Calling SetCancelRequestedAt again clears the null override.

### GetClusterName

`func (o *OverloadProtectionSimulationResponse) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *OverloadProtectionSimulationResponse) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *OverloadProtectionSimulationResponse) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### GetDurationSeconds

`func (o *OverloadProtectionSimulationResponse) GetDurationSeconds() int`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *OverloadProtectionSimulationResponse) GetDurationSecondsOk() (*int, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *OverloadProtectionSimulationResponse) SetDurationSeconds(v int)`

SetDurationSeconds sets DurationSeconds field to given value.

### GetExpiresAt

`func (o *OverloadProtectionSimulationResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OverloadProtectionSimulationResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OverloadProtectionSimulationResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### GetGroupId

`func (o *OverloadProtectionSimulationResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OverloadProtectionSimulationResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OverloadProtectionSimulationResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### GetRequestDate

`func (o *OverloadProtectionSimulationResponse) GetRequestDate() time.Time`

GetRequestDate returns the RequestDate field if non-nil, zero value otherwise.

### GetRequestDateOk

`func (o *OverloadProtectionSimulationResponse) GetRequestDateOk() (*time.Time, bool)`

GetRequestDateOk returns a tuple with the RequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDate

`func (o *OverloadProtectionSimulationResponse) SetRequestDate(v time.Time)`

SetRequestDate sets RequestDate field to given value.

### GetSimulationId

`func (o *OverloadProtectionSimulationResponse) GetSimulationId() string`

GetSimulationId returns the SimulationId field if non-nil, zero value otherwise.

### GetSimulationIdOk

`func (o *OverloadProtectionSimulationResponse) GetSimulationIdOk() (*string, bool)`

GetSimulationIdOk returns a tuple with the SimulationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationId

`func (o *OverloadProtectionSimulationResponse) SetSimulationId(v string)`

SetSimulationId sets SimulationId field to given value.

### GetState

`func (o *OverloadProtectionSimulationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OverloadProtectionSimulationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OverloadProtectionSimulationResponse) SetState(v string)`

SetState sets State field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


