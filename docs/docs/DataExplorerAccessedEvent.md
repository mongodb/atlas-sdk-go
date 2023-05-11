# DataExplorerAccessedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **userId** parameter. | [optional] [readonly] 
**Collection** | Pointer to **string** | Human-readable label of the collection on which the event occurred. The resource returns this parameter when the **eventTypeName** includes &#x60;DATA_EXPLORER&#x60;. | [optional] [readonly] 
**Created** | **time.Time** | Date and time when this event occurred. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 
**Database** | Pointer to **string** | Human-readable label of the database on which this incident occurred. The resource returns this parameter when &#x60;\&quot;eventTypeName\&quot; : \&quot;DATA_EXPLORER\&quot;&#x60; or &#x60;\&quot;eventTypeName\&quot; : \&quot;DATA_EXPLORER_CRUD\&quot;&#x60;. | [optional] [readonly] 
**EventTypeName** | [**DataExplorerAccessedEventType**](DataExplorerAccessedEventType.md) |  | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event. | [optional] [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the event. | [readonly] 
**IsGlobalAdmin** | Pointer to **bool** | Flag that indicates whether a MongoDB employee triggered the specified event. | [optional] [readonly] [default to false]
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OpType** | Pointer to **string** | Action that the database attempted to execute when the event triggered. The response returns this parameter when &#x60;eventTypeName\&quot; : \&quot;DATA_EXPLORER\&quot;&#x60;. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to which these events apply. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **username** parameter. | [optional] [readonly] 
**Raw** | Pointer to [**Raw**](Raw.md) |  | [optional] 
**RemoteAddress** | Pointer to **string** | IPv4 or IPv6 address from which the user triggered this event. | [optional] [readonly] 
**UserId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn&#39;t return the **apiKeyId** parameter. | [optional] [readonly] 
**Username** | Pointer to **string** | Email address for the user who triggered this event. If this resource returns this parameter, it doesn&#39;t return the **publicApiKey** parameter. | [optional] [readonly] 

## Methods

### NewDataExplorerAccessedEvent

`func NewDataExplorerAccessedEvent(created time.Time, eventTypeName DataExplorerAccessedEventType, id string, ) *DataExplorerAccessedEvent`

NewDataExplorerAccessedEvent instantiates a new DataExplorerAccessedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExplorerAccessedEventWithDefaults

`func NewDataExplorerAccessedEventWithDefaults() *DataExplorerAccessedEvent`

NewDataExplorerAccessedEventWithDefaults instantiates a new DataExplorerAccessedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *DataExplorerAccessedEvent) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *DataExplorerAccessedEvent) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *DataExplorerAccessedEvent) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *DataExplorerAccessedEvent) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetCollection

`func (o *DataExplorerAccessedEvent) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *DataExplorerAccessedEvent) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *DataExplorerAccessedEvent) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *DataExplorerAccessedEvent) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetCreated

`func (o *DataExplorerAccessedEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DataExplorerAccessedEvent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DataExplorerAccessedEvent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDatabase

`func (o *DataExplorerAccessedEvent) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DataExplorerAccessedEvent) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DataExplorerAccessedEvent) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *DataExplorerAccessedEvent) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetEventTypeName

`func (o *DataExplorerAccessedEvent) GetEventTypeName() DataExplorerAccessedEventType`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *DataExplorerAccessedEvent) GetEventTypeNameOk() (*DataExplorerAccessedEventType, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *DataExplorerAccessedEvent) SetEventTypeName(v DataExplorerAccessedEventType)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *DataExplorerAccessedEvent) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DataExplorerAccessedEvent) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DataExplorerAccessedEvent) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DataExplorerAccessedEvent) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *DataExplorerAccessedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataExplorerAccessedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataExplorerAccessedEvent) SetId(v string)`

SetId sets Id field to given value.


### GetIsGlobalAdmin

`func (o *DataExplorerAccessedEvent) GetIsGlobalAdmin() bool`

GetIsGlobalAdmin returns the IsGlobalAdmin field if non-nil, zero value otherwise.

### GetIsGlobalAdminOk

`func (o *DataExplorerAccessedEvent) GetIsGlobalAdminOk() (*bool, bool)`

GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalAdmin

`func (o *DataExplorerAccessedEvent) SetIsGlobalAdmin(v bool)`

SetIsGlobalAdmin sets IsGlobalAdmin field to given value.

### HasIsGlobalAdmin

`func (o *DataExplorerAccessedEvent) HasIsGlobalAdmin() bool`

HasIsGlobalAdmin returns a boolean if a field has been set.

### GetLinks

`func (o *DataExplorerAccessedEvent) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DataExplorerAccessedEvent) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DataExplorerAccessedEvent) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DataExplorerAccessedEvent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOpType

`func (o *DataExplorerAccessedEvent) GetOpType() string`

GetOpType returns the OpType field if non-nil, zero value otherwise.

### GetOpTypeOk

`func (o *DataExplorerAccessedEvent) GetOpTypeOk() (*string, bool)`

GetOpTypeOk returns a tuple with the OpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpType

`func (o *DataExplorerAccessedEvent) SetOpType(v string)`

SetOpType sets OpType field to given value.

### HasOpType

`func (o *DataExplorerAccessedEvent) HasOpType() bool`

HasOpType returns a boolean if a field has been set.

### GetOrgId

`func (o *DataExplorerAccessedEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DataExplorerAccessedEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DataExplorerAccessedEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DataExplorerAccessedEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPublicKey

`func (o *DataExplorerAccessedEvent) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DataExplorerAccessedEvent) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DataExplorerAccessedEvent) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DataExplorerAccessedEvent) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRaw

`func (o *DataExplorerAccessedEvent) GetRaw() Raw`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *DataExplorerAccessedEvent) GetRawOk() (*Raw, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *DataExplorerAccessedEvent) SetRaw(v Raw)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *DataExplorerAccessedEvent) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *DataExplorerAccessedEvent) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *DataExplorerAccessedEvent) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *DataExplorerAccessedEvent) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *DataExplorerAccessedEvent) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetUserId

`func (o *DataExplorerAccessedEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DataExplorerAccessedEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DataExplorerAccessedEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DataExplorerAccessedEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *DataExplorerAccessedEvent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DataExplorerAccessedEvent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DataExplorerAccessedEvent) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DataExplorerAccessedEvent) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


