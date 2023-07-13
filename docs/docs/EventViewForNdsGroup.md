# EventViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **userId** parameter. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when this event occurred. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 
**EventTypeName** | Pointer to **string** | Unique identifier of event type. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the event. | [optional] [readonly] 
**IsGlobalAdmin** | Pointer to **bool** | Flag that indicates whether a MongoDB employee triggered the specified event. | [optional] [readonly] [default to false]
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to which these events apply. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **username** parameter. | [optional] [readonly] 
**Raw** | Pointer to [**Raw**](Raw.md) |  | [optional] 
**RemoteAddress** | Pointer to **string** | IPv4 or IPv6 address from which the user triggered this event. | [optional] [readonly] 
**UserId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn&#39;t return the **apiKeyId** parameter. | [optional] [readonly] 
**Username** | Pointer to **string** | Email address for the user who triggered this event. If this resource returns this parameter, it doesn&#39;t return the **publicApiKey** parameter. | [optional] [readonly] 
**AlertId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert associated with the event. | [optional] [readonly] 
**AlertConfigId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert configuration associated with the **alertId**. | [optional] [readonly] 
**InvoiceId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies of the invoice associated with the event. | [optional] [readonly] 
**PaymentId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the invoice payment associated with this event. | [optional] [readonly] 
**ShardName** | Pointer to **string** | Human-readable label of the shard associated with the event. | [optional] [readonly] 
**Collection** | Pointer to **string** | Human-readable label of the collection on which the event occurred. The resource returns this parameter when the **eventTypeName** includes &#x60;DATA_EXPLORER&#x60;. | [optional] [readonly] 
**Database** | Pointer to **string** | Human-readable label of the database on which this incident occurred. The resource returns this parameter when &#x60;\&quot;eventTypeName\&quot; : \&quot;DATA_EXPLORER\&quot;&#x60; or &#x60;\&quot;eventTypeName\&quot; : \&quot;DATA_EXPLORER_CRUD\&quot;&#x60;. | [optional] [readonly] 
**OpType** | Pointer to **string** | Action that the database attempted to execute when the event triggered. The response returns this parameter when &#x60;eventTypeName\&quot; : \&quot;DATA_EXPLORER\&quot;&#x60;. | [optional] [readonly] 
**Port** | Pointer to **int** | IANA port on which the MongoDB process listens for requests. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label of the replica set associated with the event. | [optional] [readonly] 
**CurrentValue** | Pointer to [**NumberMetricValue**](NumberMetricValue.md) |  | [optional] 
**MetricName** | Pointer to **string** | Human-readable label of the metric associated with the **alertId**. This field may change type of **currentValue** field. | [optional] [readonly] 
**WhitelistEntry** | Pointer to **string** | Entry in the list of source host addresses that the API key accepts and this event targets. | [optional] [readonly] 
**EndpointId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the endpoint associated with this event. | [optional] [readonly] 
**ProviderEndpointId** | Pointer to **string** | Unique identification string that the cloud provider uses to identify the private endpoint. | [optional] [readonly] 
**TeamId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization team associated with this event. | [optional] [readonly] 
**TargetUsername** | Pointer to **string** | Email address for the console user that this event targets. The resource returns this parameter when &#x60;\&quot;eventTypeName\&quot; : \&quot;USER\&quot;&#x60;. | [optional] [readonly] 
**ResourceId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the resource associated with the event. | [optional] [readonly] 
**ResourceType** | Pointer to **string** | Unique identifier of resource type. | [optional] 

## Methods

### NewEventViewForNdsGroup

`func NewEventViewForNdsGroup() *EventViewForNdsGroup`

NewEventViewForNdsGroup instantiates a new EventViewForNdsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventViewForNdsGroupWithDefaults

`func NewEventViewForNdsGroupWithDefaults() *EventViewForNdsGroup`

NewEventViewForNdsGroupWithDefaults instantiates a new EventViewForNdsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *EventViewForNdsGroup) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *EventViewForNdsGroup) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *EventViewForNdsGroup) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *EventViewForNdsGroup) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.
### GetCreated

`func (o *EventViewForNdsGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EventViewForNdsGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EventViewForNdsGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EventViewForNdsGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.
### GetEventTypeName

`func (o *EventViewForNdsGroup) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *EventViewForNdsGroup) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *EventViewForNdsGroup) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.

### HasEventTypeName

`func (o *EventViewForNdsGroup) HasEventTypeName() bool`

HasEventTypeName returns a boolean if a field has been set.
### GetGroupId

`func (o *EventViewForNdsGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *EventViewForNdsGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *EventViewForNdsGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *EventViewForNdsGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *EventViewForNdsGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventViewForNdsGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventViewForNdsGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventViewForNdsGroup) HasId() bool`

HasId returns a boolean if a field has been set.
### GetIsGlobalAdmin

`func (o *EventViewForNdsGroup) GetIsGlobalAdmin() bool`

GetIsGlobalAdmin returns the IsGlobalAdmin field if non-nil, zero value otherwise.

### GetIsGlobalAdminOk

`func (o *EventViewForNdsGroup) GetIsGlobalAdminOk() (*bool, bool)`

GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalAdmin

`func (o *EventViewForNdsGroup) SetIsGlobalAdmin(v bool)`

SetIsGlobalAdmin sets IsGlobalAdmin field to given value.

### HasIsGlobalAdmin

`func (o *EventViewForNdsGroup) HasIsGlobalAdmin() bool`

HasIsGlobalAdmin returns a boolean if a field has been set.
### GetLinks

`func (o *EventViewForNdsGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventViewForNdsGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventViewForNdsGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventViewForNdsGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOrgId

`func (o *EventViewForNdsGroup) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EventViewForNdsGroup) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EventViewForNdsGroup) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EventViewForNdsGroup) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.
### GetPublicKey

`func (o *EventViewForNdsGroup) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *EventViewForNdsGroup) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *EventViewForNdsGroup) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *EventViewForNdsGroup) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.
### GetRaw

`func (o *EventViewForNdsGroup) GetRaw() Raw`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *EventViewForNdsGroup) GetRawOk() (*Raw, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *EventViewForNdsGroup) SetRaw(v Raw)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *EventViewForNdsGroup) HasRaw() bool`

HasRaw returns a boolean if a field has been set.
### GetRemoteAddress

`func (o *EventViewForNdsGroup) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *EventViewForNdsGroup) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *EventViewForNdsGroup) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *EventViewForNdsGroup) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.
### GetUserId

`func (o *EventViewForNdsGroup) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventViewForNdsGroup) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventViewForNdsGroup) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventViewForNdsGroup) HasUserId() bool`

HasUserId returns a boolean if a field has been set.
### GetUsername

`func (o *EventViewForNdsGroup) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EventViewForNdsGroup) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EventViewForNdsGroup) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EventViewForNdsGroup) HasUsername() bool`

HasUsername returns a boolean if a field has been set.
### GetAlertId

`func (o *EventViewForNdsGroup) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *EventViewForNdsGroup) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *EventViewForNdsGroup) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *EventViewForNdsGroup) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.
### GetAlertConfigId

`func (o *EventViewForNdsGroup) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *EventViewForNdsGroup) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *EventViewForNdsGroup) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.

### HasAlertConfigId

`func (o *EventViewForNdsGroup) HasAlertConfigId() bool`

HasAlertConfigId returns a boolean if a field has been set.
### GetInvoiceId

`func (o *EventViewForNdsGroup) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *EventViewForNdsGroup) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *EventViewForNdsGroup) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *EventViewForNdsGroup) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.
### GetPaymentId

`func (o *EventViewForNdsGroup) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *EventViewForNdsGroup) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *EventViewForNdsGroup) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *EventViewForNdsGroup) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.
### GetShardName

`func (o *EventViewForNdsGroup) GetShardName() string`

GetShardName returns the ShardName field if non-nil, zero value otherwise.

### GetShardNameOk

`func (o *EventViewForNdsGroup) GetShardNameOk() (*string, bool)`

GetShardNameOk returns a tuple with the ShardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardName

`func (o *EventViewForNdsGroup) SetShardName(v string)`

SetShardName sets ShardName field to given value.

### HasShardName

`func (o *EventViewForNdsGroup) HasShardName() bool`

HasShardName returns a boolean if a field has been set.
### GetCollection

`func (o *EventViewForNdsGroup) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *EventViewForNdsGroup) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *EventViewForNdsGroup) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *EventViewForNdsGroup) HasCollection() bool`

HasCollection returns a boolean if a field has been set.
### GetDatabase

`func (o *EventViewForNdsGroup) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *EventViewForNdsGroup) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *EventViewForNdsGroup) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *EventViewForNdsGroup) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.
### GetOpType

`func (o *EventViewForNdsGroup) GetOpType() string`

GetOpType returns the OpType field if non-nil, zero value otherwise.

### GetOpTypeOk

`func (o *EventViewForNdsGroup) GetOpTypeOk() (*string, bool)`

GetOpTypeOk returns a tuple with the OpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpType

`func (o *EventViewForNdsGroup) SetOpType(v string)`

SetOpType sets OpType field to given value.

### HasOpType

`func (o *EventViewForNdsGroup) HasOpType() bool`

HasOpType returns a boolean if a field has been set.
### GetPort

`func (o *EventViewForNdsGroup) GetPort() int`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EventViewForNdsGroup) GetPortOk() (*int, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EventViewForNdsGroup) SetPort(v int)`

SetPort sets Port field to given value.

### HasPort

`func (o *EventViewForNdsGroup) HasPort() bool`

HasPort returns a boolean if a field has been set.
### GetReplicaSetName

`func (o *EventViewForNdsGroup) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *EventViewForNdsGroup) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *EventViewForNdsGroup) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *EventViewForNdsGroup) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.
### GetCurrentValue

`func (o *EventViewForNdsGroup) GetCurrentValue() NumberMetricValue`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *EventViewForNdsGroup) GetCurrentValueOk() (*NumberMetricValue, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *EventViewForNdsGroup) SetCurrentValue(v NumberMetricValue)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *EventViewForNdsGroup) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.
### GetMetricName

`func (o *EventViewForNdsGroup) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *EventViewForNdsGroup) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *EventViewForNdsGroup) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *EventViewForNdsGroup) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.
### GetWhitelistEntry

`func (o *EventViewForNdsGroup) GetWhitelistEntry() string`

GetWhitelistEntry returns the WhitelistEntry field if non-nil, zero value otherwise.

### GetWhitelistEntryOk

`func (o *EventViewForNdsGroup) GetWhitelistEntryOk() (*string, bool)`

GetWhitelistEntryOk returns a tuple with the WhitelistEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEntry

`func (o *EventViewForNdsGroup) SetWhitelistEntry(v string)`

SetWhitelistEntry sets WhitelistEntry field to given value.

### HasWhitelistEntry

`func (o *EventViewForNdsGroup) HasWhitelistEntry() bool`

HasWhitelistEntry returns a boolean if a field has been set.
### GetEndpointId

`func (o *EventViewForNdsGroup) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *EventViewForNdsGroup) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *EventViewForNdsGroup) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *EventViewForNdsGroup) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.
### GetProviderEndpointId

`func (o *EventViewForNdsGroup) GetProviderEndpointId() string`

GetProviderEndpointId returns the ProviderEndpointId field if non-nil, zero value otherwise.

### GetProviderEndpointIdOk

`func (o *EventViewForNdsGroup) GetProviderEndpointIdOk() (*string, bool)`

GetProviderEndpointIdOk returns a tuple with the ProviderEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderEndpointId

`func (o *EventViewForNdsGroup) SetProviderEndpointId(v string)`

SetProviderEndpointId sets ProviderEndpointId field to given value.

### HasProviderEndpointId

`func (o *EventViewForNdsGroup) HasProviderEndpointId() bool`

HasProviderEndpointId returns a boolean if a field has been set.
### GetTeamId

`func (o *EventViewForNdsGroup) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *EventViewForNdsGroup) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *EventViewForNdsGroup) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *EventViewForNdsGroup) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.
### GetTargetUsername

`func (o *EventViewForNdsGroup) GetTargetUsername() string`

GetTargetUsername returns the TargetUsername field if non-nil, zero value otherwise.

### GetTargetUsernameOk

`func (o *EventViewForNdsGroup) GetTargetUsernameOk() (*string, bool)`

GetTargetUsernameOk returns a tuple with the TargetUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUsername

`func (o *EventViewForNdsGroup) SetTargetUsername(v string)`

SetTargetUsername sets TargetUsername field to given value.

### HasTargetUsername

`func (o *EventViewForNdsGroup) HasTargetUsername() bool`

HasTargetUsername returns a boolean if a field has been set.
### GetResourceId

`func (o *EventViewForNdsGroup) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *EventViewForNdsGroup) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *EventViewForNdsGroup) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *EventViewForNdsGroup) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.
### GetResourceType

`func (o *EventViewForNdsGroup) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EventViewForNdsGroup) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EventViewForNdsGroup) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *EventViewForNdsGroup) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


