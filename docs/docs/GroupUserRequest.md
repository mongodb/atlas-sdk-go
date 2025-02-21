# GroupUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | **[]string** | One or more project-level roles to assign the MongoDB Cloud user. | 
**Username** | **string** | Email address that represents the username of the MongoDB Cloud user. | 

## Methods

### NewGroupUserRequest

`func NewGroupUserRequest(roles []string, username string, ) *GroupUserRequest`

NewGroupUserRequest instantiates a new GroupUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserRequestWithDefaults

`func NewGroupUserRequestWithDefaults() *GroupUserRequest`

NewGroupUserRequestWithDefaults instantiates a new GroupUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *GroupUserRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupUserRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupUserRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### GetUsername

`func (o *GroupUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GroupUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GroupUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


