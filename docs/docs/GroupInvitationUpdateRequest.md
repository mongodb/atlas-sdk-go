# GroupInvitationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** | One or more project-level roles to assign to the MongoDB Cloud user. | [optional] 

## Methods

### NewGroupInvitationUpdateRequest

`func NewGroupInvitationUpdateRequest() *GroupInvitationUpdateRequest`

NewGroupInvitationUpdateRequest instantiates a new GroupInvitationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInvitationUpdateRequestWithDefaults

`func NewGroupInvitationUpdateRequestWithDefaults() *GroupInvitationUpdateRequest`

NewGroupInvitationUpdateRequestWithDefaults instantiates a new GroupInvitationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *GroupInvitationUpdateRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupInvitationUpdateRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupInvitationUpdateRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GroupInvitationUpdateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


