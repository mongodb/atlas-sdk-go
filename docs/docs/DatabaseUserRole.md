# DatabaseUserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | Pointer to **string** | Collection on which this role applies. | [optional] 
**DatabaseName** | **string** | Database to which the user is granted access privileges. | 
**RoleName** | **string** | Human-readable label that identifies a group of privileges assigned to a database user. This value can either be a built-in role or a custom role. | 

## Methods

### NewDatabaseUserRole

`func NewDatabaseUserRole(databaseName string, roleName string, ) *DatabaseUserRole`

NewDatabaseUserRole instantiates a new DatabaseUserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseUserRoleWithDefaults

`func NewDatabaseUserRoleWithDefaults() *DatabaseUserRole`

NewDatabaseUserRoleWithDefaults instantiates a new DatabaseUserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *DatabaseUserRole) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *DatabaseUserRole) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *DatabaseUserRole) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *DatabaseUserRole) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.
### GetDatabaseName

`func (o *DatabaseUserRole) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *DatabaseUserRole) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *DatabaseUserRole) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### GetRoleName

`func (o *DatabaseUserRole) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *DatabaseUserRole) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *DatabaseUserRole) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


