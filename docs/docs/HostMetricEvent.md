# HostMetricEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **userId** parameter. | [optional] [readonly] 
**Created** | **time.Time** | Date and time when this event occurred. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 
**CurrentValue** | Pointer to [**NumberMetricValue**](NumberMetricValue.md) |  | [optional] 
**EventTypeName** | [**HostMetricEventType**](HostMetricEventType.md) |  | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event. | [optional] [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the event. | [readonly] 
**IsGlobalAdmin** | Pointer to **bool** | Flag that indicates whether a MongoDB employee triggered the specified event. | [optional] [readonly] [default to false]
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MetricName** | Pointer to **string** | Human-readable label of the metric associated with the **alertId**. This field may change type of **currentValue** field. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to which these events apply. | [optional] [readonly] 
**Port** | Pointer to **int** | IANA port on which the MongoDB process listens for requests. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **username** parameter. | [optional] [readonly] 
**Raw** | Pointer to [**Raw**](Raw.md) |  | [optional] 
**RemoteAddress** | Pointer to **string** | IPv4 or IPv6 address from which the user triggered this event. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label of the replica set associated with the event. | [optional] [readonly] 
**ShardName** | Pointer to **string** | Human-readable label of the shard associated with the event. | [optional] [readonly] 
**UserId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn&#39;t return the **apiKeyId** parameter. | [optional] [readonly] 
**Username** | Pointer to **string** | Email address for the user who triggered this event. If this resource returns this parameter, it doesn&#39;t return the **publicApiKey** parameter. | [optional] [readonly] 

## Methods

### NewHostMetricEvent

`func NewHostMetricEvent(created time.Time, eventTypeName HostMetricEventType, id string, ) *HostMetricEvent`

NewHostMetricEvent instantiates a new HostMetricEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMetricEventWithDefaults

`func NewHostMetricEventWithDefaults() *HostMetricEvent`

NewHostMetricEventWithDefaults instantiates a new HostMetricEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *HostMetricEvent) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *HostMetricEvent) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *HostMetricEvent) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *HostMetricEvent) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetCreated

`func (o *HostMetricEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HostMetricEvent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HostMetricEvent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCurrentValue

`func (o *HostMetricEvent) GetCurrentValue() NumberMetricValue`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *HostMetricEvent) GetCurrentValueOk() (*NumberMetricValue, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *HostMetricEvent) SetCurrentValue(v NumberMetricValue)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *HostMetricEvent) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetEventTypeName

`func (o *HostMetricEvent) GetEventTypeName() HostMetricEventType`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *HostMetricEvent) GetEventTypeNameOk() (*HostMetricEventType, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *HostMetricEvent) SetEventTypeName(v HostMetricEventType)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *HostMetricEvent) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *HostMetricEvent) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *HostMetricEvent) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *HostMetricEvent) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *HostMetricEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostMetricEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostMetricEvent) SetId(v string)`

SetId sets Id field to given value.


### GetIsGlobalAdmin

`func (o *HostMetricEvent) GetIsGlobalAdmin() bool`

GetIsGlobalAdmin returns the IsGlobalAdmin field if non-nil, zero value otherwise.

### GetIsGlobalAdminOk

`func (o *HostMetricEvent) GetIsGlobalAdminOk() (*bool, bool)`

GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalAdmin

`func (o *HostMetricEvent) SetIsGlobalAdmin(v bool)`

SetIsGlobalAdmin sets IsGlobalAdmin field to given value.

### HasIsGlobalAdmin

`func (o *HostMetricEvent) HasIsGlobalAdmin() bool`

HasIsGlobalAdmin returns a boolean if a field has been set.

### GetLinks

`func (o *HostMetricEvent) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HostMetricEvent) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HostMetricEvent) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *HostMetricEvent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetricName

`func (o *HostMetricEvent) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *HostMetricEvent) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *HostMetricEvent) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *HostMetricEvent) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetOrgId

`func (o *HostMetricEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *HostMetricEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *HostMetricEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *HostMetricEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPort

`func (o *HostMetricEvent) GetPort() int`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HostMetricEvent) GetPortOk() (*int, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HostMetricEvent) SetPort(v int)`

SetPort sets Port field to given value.

### HasPort

`func (o *HostMetricEvent) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPublicKey

`func (o *HostMetricEvent) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *HostMetricEvent) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *HostMetricEvent) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *HostMetricEvent) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRaw

`func (o *HostMetricEvent) GetRaw() Raw`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *HostMetricEvent) GetRawOk() (*Raw, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *HostMetricEvent) SetRaw(v Raw)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *HostMetricEvent) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *HostMetricEvent) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *HostMetricEvent) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *HostMetricEvent) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *HostMetricEvent) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetReplicaSetName

`func (o *HostMetricEvent) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *HostMetricEvent) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *HostMetricEvent) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *HostMetricEvent) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

### GetShardName

`func (o *HostMetricEvent) GetShardName() string`

GetShardName returns the ShardName field if non-nil, zero value otherwise.

### GetShardNameOk

`func (o *HostMetricEvent) GetShardNameOk() (*string, bool)`

GetShardNameOk returns a tuple with the ShardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardName

`func (o *HostMetricEvent) SetShardName(v string)`

SetShardName sets ShardName field to given value.

### HasShardName

`func (o *HostMetricEvent) HasShardName() bool`

HasShardName returns a boolean if a field has been set.

### GetUserId

`func (o *HostMetricEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HostMetricEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HostMetricEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HostMetricEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *HostMetricEvent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HostMetricEvent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HostMetricEvent) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HostMetricEvent) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


