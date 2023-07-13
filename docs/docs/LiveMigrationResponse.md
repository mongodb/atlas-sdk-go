# LiveMigrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the migration job. | [optional] [readonly] 
**LagTimeSeconds** | Pointer to **int64** | Replication lag between the source and destination clusters. Atlas returns this setting only during an active migration, before the cutover phase. | [optional] [readonly] 
**MigrationHosts** | Pointer to **[]string** | List of hosts running MongoDB Agents. These Agents can transfer your MongoDB data between one source and one target cluster. | [optional] [readonly] 
**ReadyForCutover** | Pointer to **bool** | Flag that indicates the migrated cluster can be cut over to MongoDB Atlas. | [optional] [readonly] 
**Status** | Pointer to **string** | Progress made in migrating one cluster to MongoDB Atlas.  | Status   | Explanation | |----------|-------------| | NEW      | Someone scheduled a local cluster migration to MongoDB Atlas. | | FAILED   | The cluster migration to MongoDB Atlas failed.                | | COMPLETE | The cluster migration to MongoDB Atlas succeeded.             | | EXPIRED  | MongoDB Atlas prepares to begin the cut over of the migrating cluster when source and target clusters have almost synchronized. If &#x60;\&quot;readyForCutover\&quot; : true&#x60;, this synchronization starts a timer of 120 hours. You can extend this timer. If the timer expires, MongoDB Atlas returns this status. | | WORKING  | The cluster migration to MongoDB Atlas is performing one of the following tasks:&lt;ul&gt;&lt;li&gt;Preparing connections to source and target clusters&lt;/li&gt;&lt;li&gt;Replicating data from source to target&lt;/li&gt;&lt;li&gt;Verifying MongoDB Atlas connection settings&lt;/li&gt;&lt;li&gt;Stopping replication after the cut over&lt;/li&gt;&lt;/ul&gt; |  | [optional] [readonly] 

## Methods

### NewLiveMigrationResponse

`func NewLiveMigrationResponse() *LiveMigrationResponse`

NewLiveMigrationResponse instantiates a new LiveMigrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveMigrationResponseWithDefaults

`func NewLiveMigrationResponseWithDefaults() *LiveMigrationResponse`

NewLiveMigrationResponseWithDefaults instantiates a new LiveMigrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveMigrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveMigrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveMigrationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveMigrationResponse) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLagTimeSeconds

`func (o *LiveMigrationResponse) GetLagTimeSeconds() int64`

GetLagTimeSeconds returns the LagTimeSeconds field if non-nil, zero value otherwise.

### GetLagTimeSecondsOk

`func (o *LiveMigrationResponse) GetLagTimeSecondsOk() (*int64, bool)`

GetLagTimeSecondsOk returns a tuple with the LagTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagTimeSeconds

`func (o *LiveMigrationResponse) SetLagTimeSeconds(v int64)`

SetLagTimeSeconds sets LagTimeSeconds field to given value.

### HasLagTimeSeconds

`func (o *LiveMigrationResponse) HasLagTimeSeconds() bool`

HasLagTimeSeconds returns a boolean if a field has been set.
### GetMigrationHosts

`func (o *LiveMigrationResponse) GetMigrationHosts() []string`

GetMigrationHosts returns the MigrationHosts field if non-nil, zero value otherwise.

### GetMigrationHostsOk

`func (o *LiveMigrationResponse) GetMigrationHostsOk() (*[]string, bool)`

GetMigrationHostsOk returns a tuple with the MigrationHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationHosts

`func (o *LiveMigrationResponse) SetMigrationHosts(v []string)`

SetMigrationHosts sets MigrationHosts field to given value.

### HasMigrationHosts

`func (o *LiveMigrationResponse) HasMigrationHosts() bool`

HasMigrationHosts returns a boolean if a field has been set.
### GetReadyForCutover

`func (o *LiveMigrationResponse) GetReadyForCutover() bool`

GetReadyForCutover returns the ReadyForCutover field if non-nil, zero value otherwise.

### GetReadyForCutoverOk

`func (o *LiveMigrationResponse) GetReadyForCutoverOk() (*bool, bool)`

GetReadyForCutoverOk returns a tuple with the ReadyForCutover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyForCutover

`func (o *LiveMigrationResponse) SetReadyForCutover(v bool)`

SetReadyForCutover sets ReadyForCutover field to given value.

### HasReadyForCutover

`func (o *LiveMigrationResponse) HasReadyForCutover() bool`

HasReadyForCutover returns a boolean if a field has been set.
### GetStatus

`func (o *LiveMigrationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LiveMigrationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LiveMigrationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LiveMigrationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


