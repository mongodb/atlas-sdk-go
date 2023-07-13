# GroupInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** | One or more organization or project level roles to assign to the MongoDB Cloud user. | [optional] 
**Username** | Pointer to **string** | Email address of the MongoDB Cloud user invited to the specified project. | [optional] 

## Methods

### NewGroupInvitationRequest

`func NewGroupInvitationRequest() *GroupInvitationRequest`

NewGroupInvitationRequest instantiates a new GroupInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInvitationRequestWithDefaults

`func NewGroupInvitationRequestWithDefaults() *GroupInvitationRequest`

NewGroupInvitationRequestWithDefaults instantiates a new GroupInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *GroupInvitationRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupInvitationRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupInvitationRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GroupInvitationRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetUsername

`func (o *GroupInvitationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GroupInvitationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GroupInvitationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GroupInvitationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


