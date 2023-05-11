# FTSIndexAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **userId** parameter. | [optional] [readonly] 
**Created** | **time.Time** | Date and time when this event occurred. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 
**EventTypeName** | [**FTSIndexAuditType**](FTSIndexAuditType.md) |  | 
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

## Methods

### NewFTSIndexAudit

`func NewFTSIndexAudit(created time.Time, eventTypeName FTSIndexAuditType, id string, ) *FTSIndexAudit`

NewFTSIndexAudit instantiates a new FTSIndexAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSIndexAuditWithDefaults

`func NewFTSIndexAuditWithDefaults() *FTSIndexAudit`

NewFTSIndexAuditWithDefaults instantiates a new FTSIndexAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *FTSIndexAudit) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *FTSIndexAudit) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *FTSIndexAudit) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *FTSIndexAudit) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetCreated

`func (o *FTSIndexAudit) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FTSIndexAudit) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FTSIndexAudit) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEventTypeName

`func (o *FTSIndexAudit) GetEventTypeName() FTSIndexAuditType`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *FTSIndexAudit) GetEventTypeNameOk() (*FTSIndexAuditType, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *FTSIndexAudit) SetEventTypeName(v FTSIndexAuditType)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *FTSIndexAudit) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FTSIndexAudit) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FTSIndexAudit) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *FTSIndexAudit) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *FTSIndexAudit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FTSIndexAudit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FTSIndexAudit) SetId(v string)`

SetId sets Id field to given value.


### GetIsGlobalAdmin

`func (o *FTSIndexAudit) GetIsGlobalAdmin() bool`

GetIsGlobalAdmin returns the IsGlobalAdmin field if non-nil, zero value otherwise.

### GetIsGlobalAdminOk

`func (o *FTSIndexAudit) GetIsGlobalAdminOk() (*bool, bool)`

GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalAdmin

`func (o *FTSIndexAudit) SetIsGlobalAdmin(v bool)`

SetIsGlobalAdmin sets IsGlobalAdmin field to given value.

### HasIsGlobalAdmin

`func (o *FTSIndexAudit) HasIsGlobalAdmin() bool`

HasIsGlobalAdmin returns a boolean if a field has been set.

### GetLinks

`func (o *FTSIndexAudit) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FTSIndexAudit) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FTSIndexAudit) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FTSIndexAudit) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrgId

`func (o *FTSIndexAudit) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FTSIndexAudit) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FTSIndexAudit) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *FTSIndexAudit) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPublicKey

`func (o *FTSIndexAudit) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *FTSIndexAudit) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *FTSIndexAudit) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *FTSIndexAudit) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRaw

`func (o *FTSIndexAudit) GetRaw() Raw`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *FTSIndexAudit) GetRawOk() (*Raw, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *FTSIndexAudit) SetRaw(v Raw)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *FTSIndexAudit) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *FTSIndexAudit) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *FTSIndexAudit) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *FTSIndexAudit) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *FTSIndexAudit) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetUserId

`func (o *FTSIndexAudit) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FTSIndexAudit) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FTSIndexAudit) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FTSIndexAudit) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *FTSIndexAudit) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FTSIndexAudit) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FTSIndexAudit) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FTSIndexAudit) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


