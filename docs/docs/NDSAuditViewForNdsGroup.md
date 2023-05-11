# NDSAuditViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **userId** parameter. | [optional] [readonly] 
**Created** | **time.Time** | Date and time when this event occurred. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 
**EventTypeName** | [**NDSAuditTypeViewForNdsGroup**](NDSAuditTypeViewForNdsGroup.md) |  | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event. | [optional] [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the event. | [readonly] 
**IsGlobalAdmin** | Pointer to **bool** | Flag that indicates whether a MongoDB employee triggered the specified event. | [optional] [readonly] [default to false]
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to which these events apply. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **username** parameter. | [optional] [readonly] 
**Raw** | Pointer to [**Raw**](Raw.md) |  | [optional] 
**RemoteAddress** | Pointer to **string** | IPv4 or IPv6 address from which the user triggered this event. | [optional] [readonly] 
**UserId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn&#39;t return the **apiKeyId** parameter. | [optional] [readonly] 
**Username** | Pointer to **string** | Email address for the user who triggered this event. If this resource returns this parameter, it doesn&#39;t return the **publicApiKey** parameter. | [optional] [readonly] 
**WhitelistEntry** | Pointer to **string** | Entry in the list of source host addresses that the API key accepts and this event targets. | [optional] [readonly] 

## Methods

### NewNDSAuditViewForNdsGroup

`func NewNDSAuditViewForNdsGroup(created time.Time, eventTypeName NDSAuditTypeViewForNdsGroup, id string, ) *NDSAuditViewForNdsGroup`

NewNDSAuditViewForNdsGroup instantiates a new NDSAuditViewForNdsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNDSAuditViewForNdsGroupWithDefaults

`func NewNDSAuditViewForNdsGroupWithDefaults() *NDSAuditViewForNdsGroup`

NewNDSAuditViewForNdsGroupWithDefaults instantiates a new NDSAuditViewForNdsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *NDSAuditViewForNdsGroup) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *NDSAuditViewForNdsGroup) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *NDSAuditViewForNdsGroup) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *NDSAuditViewForNdsGroup) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetCreated

`func (o *NDSAuditViewForNdsGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NDSAuditViewForNdsGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NDSAuditViewForNdsGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEventTypeName

`func (o *NDSAuditViewForNdsGroup) GetEventTypeName() NDSAuditTypeViewForNdsGroup`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *NDSAuditViewForNdsGroup) GetEventTypeNameOk() (*NDSAuditTypeViewForNdsGroup, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *NDSAuditViewForNdsGroup) SetEventTypeName(v NDSAuditTypeViewForNdsGroup)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *NDSAuditViewForNdsGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NDSAuditViewForNdsGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NDSAuditViewForNdsGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NDSAuditViewForNdsGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *NDSAuditViewForNdsGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NDSAuditViewForNdsGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NDSAuditViewForNdsGroup) SetId(v string)`

SetId sets Id field to given value.


### GetIsGlobalAdmin

`func (o *NDSAuditViewForNdsGroup) GetIsGlobalAdmin() bool`

GetIsGlobalAdmin returns the IsGlobalAdmin field if non-nil, zero value otherwise.

### GetIsGlobalAdminOk

`func (o *NDSAuditViewForNdsGroup) GetIsGlobalAdminOk() (*bool, bool)`

GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalAdmin

`func (o *NDSAuditViewForNdsGroup) SetIsGlobalAdmin(v bool)`

SetIsGlobalAdmin sets IsGlobalAdmin field to given value.

### HasIsGlobalAdmin

`func (o *NDSAuditViewForNdsGroup) HasIsGlobalAdmin() bool`

HasIsGlobalAdmin returns a boolean if a field has been set.

### GetLinks

`func (o *NDSAuditViewForNdsGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NDSAuditViewForNdsGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NDSAuditViewForNdsGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NDSAuditViewForNdsGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrgId

`func (o *NDSAuditViewForNdsGroup) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NDSAuditViewForNdsGroup) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NDSAuditViewForNdsGroup) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NDSAuditViewForNdsGroup) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPublicKey

`func (o *NDSAuditViewForNdsGroup) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *NDSAuditViewForNdsGroup) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *NDSAuditViewForNdsGroup) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *NDSAuditViewForNdsGroup) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRaw

`func (o *NDSAuditViewForNdsGroup) GetRaw() Raw`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *NDSAuditViewForNdsGroup) GetRawOk() (*Raw, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *NDSAuditViewForNdsGroup) SetRaw(v Raw)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *NDSAuditViewForNdsGroup) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *NDSAuditViewForNdsGroup) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *NDSAuditViewForNdsGroup) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *NDSAuditViewForNdsGroup) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *NDSAuditViewForNdsGroup) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetUserId

`func (o *NDSAuditViewForNdsGroup) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *NDSAuditViewForNdsGroup) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *NDSAuditViewForNdsGroup) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *NDSAuditViewForNdsGroup) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *NDSAuditViewForNdsGroup) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NDSAuditViewForNdsGroup) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NDSAuditViewForNdsGroup) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NDSAuditViewForNdsGroup) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWhitelistEntry

`func (o *NDSAuditViewForNdsGroup) GetWhitelistEntry() string`

GetWhitelistEntry returns the WhitelistEntry field if non-nil, zero value otherwise.

### GetWhitelistEntryOk

`func (o *NDSAuditViewForNdsGroup) GetWhitelistEntryOk() (*string, bool)`

GetWhitelistEntryOk returns a tuple with the WhitelistEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEntry

`func (o *NDSAuditViewForNdsGroup) SetWhitelistEntry(v string)`

SetWhitelistEntry sets WhitelistEntry field to given value.

### HasWhitelistEntry

`func (o *NDSAuditViewForNdsGroup) HasWhitelistEntry() bool`

HasWhitelistEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


