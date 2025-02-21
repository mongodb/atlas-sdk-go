# OrgUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**OrgUserRolesRequest**](OrgUserRolesRequest.md) |  | [optional] 
**TeamIds** | Pointer to **[]string** | List of unique 24-hexadecimal digit strings that identifies the teams to assign the MongoDB Cloud user. | [optional] 

## Methods

### NewOrgUserUpdateRequest

`func NewOrgUserUpdateRequest() *OrgUserUpdateRequest`

NewOrgUserUpdateRequest instantiates a new OrgUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUserUpdateRequestWithDefaults

`func NewOrgUserUpdateRequestWithDefaults() *OrgUserUpdateRequest`

NewOrgUserUpdateRequestWithDefaults instantiates a new OrgUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *OrgUserUpdateRequest) GetRoles() OrgUserRolesRequest`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrgUserUpdateRequest) GetRolesOk() (*OrgUserRolesRequest, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrgUserUpdateRequest) SetRoles(v OrgUserRolesRequest)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrgUserUpdateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetTeamIds

`func (o *OrgUserUpdateRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *OrgUserUpdateRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *OrgUserUpdateRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *OrgUserUpdateRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


