# LiveMigrationRequest20240530

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the migration request. | [optional] [readonly] 
**Destination** | [**Destination**](Destination.md) |  | 
**DropDestinationData** | Pointer to **bool** | Flag that indicates whether the migration process drops all collections from the destination cluster before the migration starts. | [optional] [default to false]
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MigrationHosts** | Pointer to **[]string** | List of migration hosts used for this migration. | [optional] 
**Sharding** | Pointer to [**ShardingRequest**](ShardingRequest.md) |  | [optional] 
**Source** | [**Source**](Source.md) |  | 

## Methods

### NewLiveMigrationRequest20240530

`func NewLiveMigrationRequest20240530(destination Destination, source Source, ) *LiveMigrationRequest20240530`

NewLiveMigrationRequest20240530 instantiates a new LiveMigrationRequest20240530 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveMigrationRequest20240530WithDefaults

`func NewLiveMigrationRequest20240530WithDefaults() *LiveMigrationRequest20240530`

NewLiveMigrationRequest20240530WithDefaults instantiates a new LiveMigrationRequest20240530 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveMigrationRequest20240530) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveMigrationRequest20240530) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveMigrationRequest20240530) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveMigrationRequest20240530) HasId() bool`

HasId returns a boolean if a field has been set.
### GetDestination

`func (o *LiveMigrationRequest20240530) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *LiveMigrationRequest20240530) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *LiveMigrationRequest20240530) SetDestination(v Destination)`

SetDestination sets Destination field to given value.

### GetDropDestinationData

`func (o *LiveMigrationRequest20240530) GetDropDestinationData() bool`

GetDropDestinationData returns the DropDestinationData field if non-nil, zero value otherwise.

### GetDropDestinationDataOk

`func (o *LiveMigrationRequest20240530) GetDropDestinationDataOk() (*bool, bool)`

GetDropDestinationDataOk returns a tuple with the DropDestinationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropDestinationData

`func (o *LiveMigrationRequest20240530) SetDropDestinationData(v bool)`

SetDropDestinationData sets DropDestinationData field to given value.

### HasDropDestinationData

`func (o *LiveMigrationRequest20240530) HasDropDestinationData() bool`

HasDropDestinationData returns a boolean if a field has been set.
### GetLinks

`func (o *LiveMigrationRequest20240530) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LiveMigrationRequest20240530) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LiveMigrationRequest20240530) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LiveMigrationRequest20240530) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMigrationHosts

`func (o *LiveMigrationRequest20240530) GetMigrationHosts() []string`

GetMigrationHosts returns the MigrationHosts field if non-nil, zero value otherwise.

### GetMigrationHostsOk

`func (o *LiveMigrationRequest20240530) GetMigrationHostsOk() (*[]string, bool)`

GetMigrationHostsOk returns a tuple with the MigrationHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationHosts

`func (o *LiveMigrationRequest20240530) SetMigrationHosts(v []string)`

SetMigrationHosts sets MigrationHosts field to given value.

### HasMigrationHosts

`func (o *LiveMigrationRequest20240530) HasMigrationHosts() bool`

HasMigrationHosts returns a boolean if a field has been set.
### GetSharding

`func (o *LiveMigrationRequest20240530) GetSharding() ShardingRequest`

GetSharding returns the Sharding field if non-nil, zero value otherwise.

### GetShardingOk

`func (o *LiveMigrationRequest20240530) GetShardingOk() (*ShardingRequest, bool)`

GetShardingOk returns a tuple with the Sharding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharding

`func (o *LiveMigrationRequest20240530) SetSharding(v ShardingRequest)`

SetSharding sets Sharding field to given value.

### HasSharding

`func (o *LiveMigrationRequest20240530) HasSharding() bool`

HasSharding returns a boolean if a field has been set.
### GetSource

`func (o *LiveMigrationRequest20240530) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LiveMigrationRequest20240530) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LiveMigrationRequest20240530) SetSource(v Source)`

SetSource sets Source field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


