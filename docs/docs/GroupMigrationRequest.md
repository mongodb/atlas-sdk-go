# GroupMigrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationOrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to move the specified project to. | [optional] 
**DestinationOrgPrivateApiKey** | Pointer to **string** | Unique string that identifies the private part of the API Key used to verify access to the destination organization. This parameter is required only when you authenticate with Programmatic API Keys. | [optional] 
**DestinationOrgPublicApiKey** | Pointer to **string** | Unique string that identifies the public part of the API Key used to verify access to the destination organization. This parameter is required only when you authenticate with Programmatic API Keys. | [optional] 

## Methods

### NewGroupMigrationRequest

`func NewGroupMigrationRequest() *GroupMigrationRequest`

NewGroupMigrationRequest instantiates a new GroupMigrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMigrationRequestWithDefaults

`func NewGroupMigrationRequestWithDefaults() *GroupMigrationRequest`

NewGroupMigrationRequestWithDefaults instantiates a new GroupMigrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationOrgId

`func (o *GroupMigrationRequest) GetDestinationOrgId() string`

GetDestinationOrgId returns the DestinationOrgId field if non-nil, zero value otherwise.

### GetDestinationOrgIdOk

`func (o *GroupMigrationRequest) GetDestinationOrgIdOk() (*string, bool)`

GetDestinationOrgIdOk returns a tuple with the DestinationOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationOrgId

`func (o *GroupMigrationRequest) SetDestinationOrgId(v string)`

SetDestinationOrgId sets DestinationOrgId field to given value.

### HasDestinationOrgId

`func (o *GroupMigrationRequest) HasDestinationOrgId() bool`

HasDestinationOrgId returns a boolean if a field has been set.
### GetDestinationOrgPrivateApiKey

`func (o *GroupMigrationRequest) GetDestinationOrgPrivateApiKey() string`

GetDestinationOrgPrivateApiKey returns the DestinationOrgPrivateApiKey field if non-nil, zero value otherwise.

### GetDestinationOrgPrivateApiKeyOk

`func (o *GroupMigrationRequest) GetDestinationOrgPrivateApiKeyOk() (*string, bool)`

GetDestinationOrgPrivateApiKeyOk returns a tuple with the DestinationOrgPrivateApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationOrgPrivateApiKey

`func (o *GroupMigrationRequest) SetDestinationOrgPrivateApiKey(v string)`

SetDestinationOrgPrivateApiKey sets DestinationOrgPrivateApiKey field to given value.

### HasDestinationOrgPrivateApiKey

`func (o *GroupMigrationRequest) HasDestinationOrgPrivateApiKey() bool`

HasDestinationOrgPrivateApiKey returns a boolean if a field has been set.
### GetDestinationOrgPublicApiKey

`func (o *GroupMigrationRequest) GetDestinationOrgPublicApiKey() string`

GetDestinationOrgPublicApiKey returns the DestinationOrgPublicApiKey field if non-nil, zero value otherwise.

### GetDestinationOrgPublicApiKeyOk

`func (o *GroupMigrationRequest) GetDestinationOrgPublicApiKeyOk() (*string, bool)`

GetDestinationOrgPublicApiKeyOk returns a tuple with the DestinationOrgPublicApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationOrgPublicApiKey

`func (o *GroupMigrationRequest) SetDestinationOrgPublicApiKey(v string)`

SetDestinationOrgPublicApiKey sets DestinationOrgPublicApiKey field to given value.

### HasDestinationOrgPublicApiKey

`func (o *GroupMigrationRequest) HasDestinationOrgPublicApiKey() bool`

HasDestinationOrgPublicApiKey returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


