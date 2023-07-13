# LiveMigrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the migration request. | [optional] [readonly] 
**Destination** | [**Destination**](Destination.md) |  | 
**DropEnabled** | **bool** | Flag that indicates whether the migration process drops all collections from the destination cluster before the migration starts. | 
**MigrationHosts** | Pointer to **[]string** | List of migration hosts used for this migration. | [optional] 
**Source** | [**Source**](Source.md) |  | 

## Methods

### NewLiveMigrationRequest

`func NewLiveMigrationRequest(destination Destination, dropEnabled bool, source Source, ) *LiveMigrationRequest`

NewLiveMigrationRequest instantiates a new LiveMigrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveMigrationRequestWithDefaults

`func NewLiveMigrationRequestWithDefaults() *LiveMigrationRequest`

NewLiveMigrationRequestWithDefaults instantiates a new LiveMigrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveMigrationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveMigrationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveMigrationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveMigrationRequest) HasId() bool`

HasId returns a boolean if a field has been set.
### GetDestination

`func (o *LiveMigrationRequest) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *LiveMigrationRequest) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *LiveMigrationRequest) SetDestination(v Destination)`

SetDestination sets Destination field to given value.

### GetDropEnabled

`func (o *LiveMigrationRequest) GetDropEnabled() bool`

GetDropEnabled returns the DropEnabled field if non-nil, zero value otherwise.

### GetDropEnabledOk

`func (o *LiveMigrationRequest) GetDropEnabledOk() (*bool, bool)`

GetDropEnabledOk returns a tuple with the DropEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropEnabled

`func (o *LiveMigrationRequest) SetDropEnabled(v bool)`

SetDropEnabled sets DropEnabled field to given value.

### GetMigrationHosts

`func (o *LiveMigrationRequest) GetMigrationHosts() []string`

GetMigrationHosts returns the MigrationHosts field if non-nil, zero value otherwise.

### GetMigrationHostsOk

`func (o *LiveMigrationRequest) GetMigrationHostsOk() (*[]string, bool)`

GetMigrationHostsOk returns a tuple with the MigrationHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationHosts

`func (o *LiveMigrationRequest) SetMigrationHosts(v []string)`

SetMigrationHosts sets MigrationHosts field to given value.

### HasMigrationHosts

`func (o *LiveMigrationRequest) HasMigrationHosts() bool`

HasMigrationHosts returns a boolean if a field has been set.
### GetSource

`func (o *LiveMigrationRequest) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LiveMigrationRequest) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LiveMigrationRequest) SetSource(v Source)`

SetSource sets Source field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


