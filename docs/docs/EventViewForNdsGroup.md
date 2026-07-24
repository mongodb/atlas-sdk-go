# EventViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the API Key that triggered the event. If this resource returns this parameter, it doesn&#39;t return the &#x60;userId&#x60; parameter. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when this event occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**EventTypeName** | Pointer to **string** | Unique identifier of event type. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The &#x60;eventId&#x60; identifies the specific event. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the event. | [optional] [readonly] 
**IsGlobalAdmin** | Pointer to **bool** | Flag that indicates whether a MongoDB employee triggered the specified event. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to which these events apply. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public part of the API key that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **username** parameter. | [optional] [readonly] 
**Raw** | Pointer to [**Raw**](Raw.md) |  | [optional] 
**RemoteAddress** | Pointer to **string** | IPv4 or IPv6 address from which the user triggered this event. | [optional] [readonly] 
**UserId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn&#39;t return the &#x60;apiKeyId&#x60; parameter. | [optional] [readonly] 
**Username** | Pointer to **string** | Email address for the user who triggered this event. If this resource returns this parameter, it doesn&#39;t return the &#x60;publicApiKey&#x60; parameter. | [optional] [readonly] 
**AlertId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert associated with the event. | [optional] [readonly] 
**AlertConfigId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert configuration associated with the &#x60;alertId&#x60;. | [optional] [readonly] 
**TargetPublicKey** | Pointer to **string** | Public part of the API key that this event targets. | [optional] [readonly] 
**WhitelistEntry** | Pointer to **string** | Entry in the list of source host addresses that the API key accepts and this event targets. | [optional] [readonly] 
**InvoiceId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies of the invoice associated with the event. | [optional] [readonly] 
**PaymentId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the invoice payment associated with this event. | [optional] [readonly] 
**ShardName** | Pointer to **string** | Human-readable label of the shard associated with the event. | [optional] [readonly] 
**Collection** | Pointer to **string** | Human-readable label of the collection on which the event occurred. The resource returns this parameter when the &#x60;eventTypeName&#x60; includes &#x60;DATA_EXPLORER&#x60;. | [optional] [readonly] 
**Database** | Pointer to **string** | Human-readable label of the database on which this incident occurred. The resource returns this parameter when &#x60;\&quot;eventTypeName\&quot; : \&quot;DATA_EXPLORER\&quot;&#x60; or &#x60;\&quot;eventTypeName\&quot; : \&quot;DATA_EXPLORER_CRUD\&quot;&#x60;. | [optional] [readonly] 
**OpType** | Pointer to **string** | Action that the database attempted to execute when the event triggered. The response returns this parameter when &#x60;eventTypeName\&quot; : \&quot;DATA_EXPLORER\&quot;&#x60;. | [optional] [readonly] 
**SessionId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Data Explorer session associated with the event. | [optional] [readonly] 
**DeskLocation** | Pointer to **string** | Desk location of MongoDB employee associated with the event. | [optional] [readonly] 
**EmployeeIdentifier** | Pointer to **string** | Identifier of MongoDB employee associated with the event. | [optional] [readonly] 
**Port** | Pointer to **int** | IANA port on which the MongoDB process listens for requests. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label of the replica set associated with the event. | [optional] [readonly] 
**CurrentValue** | Pointer to [**NumberMetricValue**](NumberMetricValue.md) |  | [optional] 
**MetricName** | Pointer to **string** | Human-readable label of the metric associated with the &#x60;alertId&#x60;. This field may change type of &#x60;currentValue&#x60; field. | [optional] [readonly] 
**DbUserUsername** | Pointer to **string** | The username of the MongoDB User that was created, deleted, or edited. | [optional] [readonly] 
**EndpointId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the endpoint associated with this event. | [optional] [readonly] 
**ProviderEndpointId** | Pointer to **string** | Unique identification string that the cloud provider uses to identify the private endpoint. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Fully qualified domain name (FQDN) of the host associated with the event. | [optional] [readonly] 
**TeamId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization team associated with this event. | [optional] [readonly] 
**TargetUsername** | Pointer to **string** | Email address for the console user that this event targets. The resource returns this parameter when &#x60;\&quot;eventTypeName\&quot; : \&quot;USER\&quot;&#x60;. | [optional] [readonly] 
**ResourceId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the resource associated with the event. | [optional] [readonly] 
**ResourceType** | Pointer to **string** | Unique identifier of resource type. | [optional] 
**InstanceName** | Pointer to **string** | Name of the stream processing workspace associated with the event. | [optional] [readonly] 
**ProcessorErrorMsg** | Pointer to **string** | Error message linked to the stream processor associated with the event. | [optional] [readonly] 
**ProcessorName** | Pointer to **string** | Name of the stream processor associated with the event. | [optional] [readonly] 
**ProcessorState** | Pointer to **string** | State of the stream processor associated with the event. | [optional] [readonly] 
**ResourcePolicyId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the resource policy. | [optional] [readonly] 
**ViolatedPolicies** | Pointer to **[]string** | String representation of the violated resource policy ids. | [optional] 
**ViolationAction** | Pointer to **string** | Resource policy action taken by the user and evaluated against the currently active policies. | [optional] [readonly] 

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

### SetApiKeyIdNil

`func (o *EventViewForNdsGroup) SetApiKeyIdNil()`

SetApiKeyIdNil sets ApiKeyId to an explicit JSON null when marshaled, overriding any value previously set with SetApiKeyId. Calling SetApiKeyId again clears the null override.

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

### SetCreatedNil

`func (o *EventViewForNdsGroup) SetCreatedNil()`

SetCreatedNil sets Created to an explicit JSON null when marshaled, overriding any value previously set with SetCreated. Calling SetCreated again clears the null override.

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

### SetEventTypeNameNil

`func (o *EventViewForNdsGroup) SetEventTypeNameNil()`

SetEventTypeNameNil sets EventTypeName to an explicit JSON null when marshaled, overriding any value previously set with SetEventTypeName. Calling SetEventTypeName again clears the null override.

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

### SetGroupIdNil

`func (o *EventViewForNdsGroup) SetGroupIdNil()`

SetGroupIdNil sets GroupId to an explicit JSON null when marshaled, overriding any value previously set with SetGroupId. Calling SetGroupId again clears the null override.

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

### SetIdNil

`func (o *EventViewForNdsGroup) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

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

### SetIsGlobalAdminNil

`func (o *EventViewForNdsGroup) SetIsGlobalAdminNil()`

SetIsGlobalAdminNil sets IsGlobalAdmin to an explicit JSON null when marshaled, overriding any value previously set with SetIsGlobalAdmin. Calling SetIsGlobalAdmin again clears the null override.

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

### SetLinksNil

`func (o *EventViewForNdsGroup) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

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

### SetOrgIdNil

`func (o *EventViewForNdsGroup) SetOrgIdNil()`

SetOrgIdNil sets OrgId to an explicit JSON null when marshaled, overriding any value previously set with SetOrgId. Calling SetOrgId again clears the null override.

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

### SetPublicKeyNil

`func (o *EventViewForNdsGroup) SetPublicKeyNil()`

SetPublicKeyNil sets PublicKey to an explicit JSON null when marshaled, overriding any value previously set with SetPublicKey. Calling SetPublicKey again clears the null override.

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

### SetRawNil

`func (o *EventViewForNdsGroup) SetRawNil()`

SetRawNil sets Raw to an explicit JSON null when marshaled, overriding any value previously set with SetRaw. Calling SetRaw again clears the null override.

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

### SetRemoteAddressNil

`func (o *EventViewForNdsGroup) SetRemoteAddressNil()`

SetRemoteAddressNil sets RemoteAddress to an explicit JSON null when marshaled, overriding any value previously set with SetRemoteAddress. Calling SetRemoteAddress again clears the null override.

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

### SetUserIdNil

`func (o *EventViewForNdsGroup) SetUserIdNil()`

SetUserIdNil sets UserId to an explicit JSON null when marshaled, overriding any value previously set with SetUserId. Calling SetUserId again clears the null override.

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

### SetUsernameNil

`func (o *EventViewForNdsGroup) SetUsernameNil()`

SetUsernameNil sets Username to an explicit JSON null when marshaled, overriding any value previously set with SetUsername. Calling SetUsername again clears the null override.

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

### SetAlertIdNil

`func (o *EventViewForNdsGroup) SetAlertIdNil()`

SetAlertIdNil sets AlertId to an explicit JSON null when marshaled, overriding any value previously set with SetAlertId. Calling SetAlertId again clears the null override.

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

### SetAlertConfigIdNil

`func (o *EventViewForNdsGroup) SetAlertConfigIdNil()`

SetAlertConfigIdNil sets AlertConfigId to an explicit JSON null when marshaled, overriding any value previously set with SetAlertConfigId. Calling SetAlertConfigId again clears the null override.

### GetTargetPublicKey

`func (o *EventViewForNdsGroup) GetTargetPublicKey() string`

GetTargetPublicKey returns the TargetPublicKey field if non-nil, zero value otherwise.

### GetTargetPublicKeyOk

`func (o *EventViewForNdsGroup) GetTargetPublicKeyOk() (*string, bool)`

GetTargetPublicKeyOk returns a tuple with the TargetPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPublicKey

`func (o *EventViewForNdsGroup) SetTargetPublicKey(v string)`

SetTargetPublicKey sets TargetPublicKey field to given value.

### HasTargetPublicKey

`func (o *EventViewForNdsGroup) HasTargetPublicKey() bool`

HasTargetPublicKey returns a boolean if a field has been set.

### SetTargetPublicKeyNil

`func (o *EventViewForNdsGroup) SetTargetPublicKeyNil()`

SetTargetPublicKeyNil sets TargetPublicKey to an explicit JSON null when marshaled, overriding any value previously set with SetTargetPublicKey. Calling SetTargetPublicKey again clears the null override.

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

### SetWhitelistEntryNil

`func (o *EventViewForNdsGroup) SetWhitelistEntryNil()`

SetWhitelistEntryNil sets WhitelistEntry to an explicit JSON null when marshaled, overriding any value previously set with SetWhitelistEntry. Calling SetWhitelistEntry again clears the null override.

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

### SetInvoiceIdNil

`func (o *EventViewForNdsGroup) SetInvoiceIdNil()`

SetInvoiceIdNil sets InvoiceId to an explicit JSON null when marshaled, overriding any value previously set with SetInvoiceId. Calling SetInvoiceId again clears the null override.

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

### SetPaymentIdNil

`func (o *EventViewForNdsGroup) SetPaymentIdNil()`

SetPaymentIdNil sets PaymentId to an explicit JSON null when marshaled, overriding any value previously set with SetPaymentId. Calling SetPaymentId again clears the null override.

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

### SetShardNameNil

`func (o *EventViewForNdsGroup) SetShardNameNil()`

SetShardNameNil sets ShardName to an explicit JSON null when marshaled, overriding any value previously set with SetShardName. Calling SetShardName again clears the null override.

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

### SetCollectionNil

`func (o *EventViewForNdsGroup) SetCollectionNil()`

SetCollectionNil sets Collection to an explicit JSON null when marshaled, overriding any value previously set with SetCollection. Calling SetCollection again clears the null override.

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

### SetDatabaseNil

`func (o *EventViewForNdsGroup) SetDatabaseNil()`

SetDatabaseNil sets Database to an explicit JSON null when marshaled, overriding any value previously set with SetDatabase. Calling SetDatabase again clears the null override.

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

### SetOpTypeNil

`func (o *EventViewForNdsGroup) SetOpTypeNil()`

SetOpTypeNil sets OpType to an explicit JSON null when marshaled, overriding any value previously set with SetOpType. Calling SetOpType again clears the null override.

### GetSessionId

`func (o *EventViewForNdsGroup) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EventViewForNdsGroup) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EventViewForNdsGroup) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *EventViewForNdsGroup) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *EventViewForNdsGroup) SetSessionIdNil()`

SetSessionIdNil sets SessionId to an explicit JSON null when marshaled, overriding any value previously set with SetSessionId. Calling SetSessionId again clears the null override.

### GetDeskLocation

`func (o *EventViewForNdsGroup) GetDeskLocation() string`

GetDeskLocation returns the DeskLocation field if non-nil, zero value otherwise.

### GetDeskLocationOk

`func (o *EventViewForNdsGroup) GetDeskLocationOk() (*string, bool)`

GetDeskLocationOk returns a tuple with the DeskLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeskLocation

`func (o *EventViewForNdsGroup) SetDeskLocation(v string)`

SetDeskLocation sets DeskLocation field to given value.

### HasDeskLocation

`func (o *EventViewForNdsGroup) HasDeskLocation() bool`

HasDeskLocation returns a boolean if a field has been set.

### SetDeskLocationNil

`func (o *EventViewForNdsGroup) SetDeskLocationNil()`

SetDeskLocationNil sets DeskLocation to an explicit JSON null when marshaled, overriding any value previously set with SetDeskLocation. Calling SetDeskLocation again clears the null override.

### GetEmployeeIdentifier

`func (o *EventViewForNdsGroup) GetEmployeeIdentifier() string`

GetEmployeeIdentifier returns the EmployeeIdentifier field if non-nil, zero value otherwise.

### GetEmployeeIdentifierOk

`func (o *EventViewForNdsGroup) GetEmployeeIdentifierOk() (*string, bool)`

GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIdentifier

`func (o *EventViewForNdsGroup) SetEmployeeIdentifier(v string)`

SetEmployeeIdentifier sets EmployeeIdentifier field to given value.

### HasEmployeeIdentifier

`func (o *EventViewForNdsGroup) HasEmployeeIdentifier() bool`

HasEmployeeIdentifier returns a boolean if a field has been set.

### SetEmployeeIdentifierNil

`func (o *EventViewForNdsGroup) SetEmployeeIdentifierNil()`

SetEmployeeIdentifierNil sets EmployeeIdentifier to an explicit JSON null when marshaled, overriding any value previously set with SetEmployeeIdentifier. Calling SetEmployeeIdentifier again clears the null override.

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

### SetPortNil

`func (o *EventViewForNdsGroup) SetPortNil()`

SetPortNil sets Port to an explicit JSON null when marshaled, overriding any value previously set with SetPort. Calling SetPort again clears the null override.

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

### SetReplicaSetNameNil

`func (o *EventViewForNdsGroup) SetReplicaSetNameNil()`

SetReplicaSetNameNil sets ReplicaSetName to an explicit JSON null when marshaled, overriding any value previously set with SetReplicaSetName. Calling SetReplicaSetName again clears the null override.

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

### SetCurrentValueNil

`func (o *EventViewForNdsGroup) SetCurrentValueNil()`

SetCurrentValueNil sets CurrentValue to an explicit JSON null when marshaled, overriding any value previously set with SetCurrentValue. Calling SetCurrentValue again clears the null override.

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

### SetMetricNameNil

`func (o *EventViewForNdsGroup) SetMetricNameNil()`

SetMetricNameNil sets MetricName to an explicit JSON null when marshaled, overriding any value previously set with SetMetricName. Calling SetMetricName again clears the null override.

### GetDbUserUsername

`func (o *EventViewForNdsGroup) GetDbUserUsername() string`

GetDbUserUsername returns the DbUserUsername field if non-nil, zero value otherwise.

### GetDbUserUsernameOk

`func (o *EventViewForNdsGroup) GetDbUserUsernameOk() (*string, bool)`

GetDbUserUsernameOk returns a tuple with the DbUserUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUserUsername

`func (o *EventViewForNdsGroup) SetDbUserUsername(v string)`

SetDbUserUsername sets DbUserUsername field to given value.

### HasDbUserUsername

`func (o *EventViewForNdsGroup) HasDbUserUsername() bool`

HasDbUserUsername returns a boolean if a field has been set.

### SetDbUserUsernameNil

`func (o *EventViewForNdsGroup) SetDbUserUsernameNil()`

SetDbUserUsernameNil sets DbUserUsername to an explicit JSON null when marshaled, overriding any value previously set with SetDbUserUsername. Calling SetDbUserUsername again clears the null override.

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

### SetEndpointIdNil

`func (o *EventViewForNdsGroup) SetEndpointIdNil()`

SetEndpointIdNil sets EndpointId to an explicit JSON null when marshaled, overriding any value previously set with SetEndpointId. Calling SetEndpointId again clears the null override.

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

### SetProviderEndpointIdNil

`func (o *EventViewForNdsGroup) SetProviderEndpointIdNil()`

SetProviderEndpointIdNil sets ProviderEndpointId to an explicit JSON null when marshaled, overriding any value previously set with SetProviderEndpointId. Calling SetProviderEndpointId again clears the null override.

### GetHostname

`func (o *EventViewForNdsGroup) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EventViewForNdsGroup) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EventViewForNdsGroup) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EventViewForNdsGroup) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *EventViewForNdsGroup) SetHostnameNil()`

SetHostnameNil sets Hostname to an explicit JSON null when marshaled, overriding any value previously set with SetHostname. Calling SetHostname again clears the null override.

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

### SetTeamIdNil

`func (o *EventViewForNdsGroup) SetTeamIdNil()`

SetTeamIdNil sets TeamId to an explicit JSON null when marshaled, overriding any value previously set with SetTeamId. Calling SetTeamId again clears the null override.

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

### SetTargetUsernameNil

`func (o *EventViewForNdsGroup) SetTargetUsernameNil()`

SetTargetUsernameNil sets TargetUsername to an explicit JSON null when marshaled, overriding any value previously set with SetTargetUsername. Calling SetTargetUsername again clears the null override.

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

### SetResourceIdNil

`func (o *EventViewForNdsGroup) SetResourceIdNil()`

SetResourceIdNil sets ResourceId to an explicit JSON null when marshaled, overriding any value previously set with SetResourceId. Calling SetResourceId again clears the null override.

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

### SetResourceTypeNil

`func (o *EventViewForNdsGroup) SetResourceTypeNil()`

SetResourceTypeNil sets ResourceType to an explicit JSON null when marshaled, overriding any value previously set with SetResourceType. Calling SetResourceType again clears the null override.

### GetInstanceName

`func (o *EventViewForNdsGroup) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *EventViewForNdsGroup) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *EventViewForNdsGroup) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *EventViewForNdsGroup) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### SetInstanceNameNil

`func (o *EventViewForNdsGroup) SetInstanceNameNil()`

SetInstanceNameNil sets InstanceName to an explicit JSON null when marshaled, overriding any value previously set with SetInstanceName. Calling SetInstanceName again clears the null override.

### GetProcessorErrorMsg

`func (o *EventViewForNdsGroup) GetProcessorErrorMsg() string`

GetProcessorErrorMsg returns the ProcessorErrorMsg field if non-nil, zero value otherwise.

### GetProcessorErrorMsgOk

`func (o *EventViewForNdsGroup) GetProcessorErrorMsgOk() (*string, bool)`

GetProcessorErrorMsgOk returns a tuple with the ProcessorErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorErrorMsg

`func (o *EventViewForNdsGroup) SetProcessorErrorMsg(v string)`

SetProcessorErrorMsg sets ProcessorErrorMsg field to given value.

### HasProcessorErrorMsg

`func (o *EventViewForNdsGroup) HasProcessorErrorMsg() bool`

HasProcessorErrorMsg returns a boolean if a field has been set.

### SetProcessorErrorMsgNil

`func (o *EventViewForNdsGroup) SetProcessorErrorMsgNil()`

SetProcessorErrorMsgNil sets ProcessorErrorMsg to an explicit JSON null when marshaled, overriding any value previously set with SetProcessorErrorMsg. Calling SetProcessorErrorMsg again clears the null override.

### GetProcessorName

`func (o *EventViewForNdsGroup) GetProcessorName() string`

GetProcessorName returns the ProcessorName field if non-nil, zero value otherwise.

### GetProcessorNameOk

`func (o *EventViewForNdsGroup) GetProcessorNameOk() (*string, bool)`

GetProcessorNameOk returns a tuple with the ProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorName

`func (o *EventViewForNdsGroup) SetProcessorName(v string)`

SetProcessorName sets ProcessorName field to given value.

### HasProcessorName

`func (o *EventViewForNdsGroup) HasProcessorName() bool`

HasProcessorName returns a boolean if a field has been set.

### SetProcessorNameNil

`func (o *EventViewForNdsGroup) SetProcessorNameNil()`

SetProcessorNameNil sets ProcessorName to an explicit JSON null when marshaled, overriding any value previously set with SetProcessorName. Calling SetProcessorName again clears the null override.

### GetProcessorState

`func (o *EventViewForNdsGroup) GetProcessorState() string`

GetProcessorState returns the ProcessorState field if non-nil, zero value otherwise.

### GetProcessorStateOk

`func (o *EventViewForNdsGroup) GetProcessorStateOk() (*string, bool)`

GetProcessorStateOk returns a tuple with the ProcessorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorState

`func (o *EventViewForNdsGroup) SetProcessorState(v string)`

SetProcessorState sets ProcessorState field to given value.

### HasProcessorState

`func (o *EventViewForNdsGroup) HasProcessorState() bool`

HasProcessorState returns a boolean if a field has been set.

### SetProcessorStateNil

`func (o *EventViewForNdsGroup) SetProcessorStateNil()`

SetProcessorStateNil sets ProcessorState to an explicit JSON null when marshaled, overriding any value previously set with SetProcessorState. Calling SetProcessorState again clears the null override.

### GetResourcePolicyId

`func (o *EventViewForNdsGroup) GetResourcePolicyId() string`

GetResourcePolicyId returns the ResourcePolicyId field if non-nil, zero value otherwise.

### GetResourcePolicyIdOk

`func (o *EventViewForNdsGroup) GetResourcePolicyIdOk() (*string, bool)`

GetResourcePolicyIdOk returns a tuple with the ResourcePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicyId

`func (o *EventViewForNdsGroup) SetResourcePolicyId(v string)`

SetResourcePolicyId sets ResourcePolicyId field to given value.

### HasResourcePolicyId

`func (o *EventViewForNdsGroup) HasResourcePolicyId() bool`

HasResourcePolicyId returns a boolean if a field has been set.

### SetResourcePolicyIdNil

`func (o *EventViewForNdsGroup) SetResourcePolicyIdNil()`

SetResourcePolicyIdNil sets ResourcePolicyId to an explicit JSON null when marshaled, overriding any value previously set with SetResourcePolicyId. Calling SetResourcePolicyId again clears the null override.

### GetViolatedPolicies

`func (o *EventViewForNdsGroup) GetViolatedPolicies() []string`

GetViolatedPolicies returns the ViolatedPolicies field if non-nil, zero value otherwise.

### GetViolatedPoliciesOk

`func (o *EventViewForNdsGroup) GetViolatedPoliciesOk() (*[]string, bool)`

GetViolatedPoliciesOk returns a tuple with the ViolatedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatedPolicies

`func (o *EventViewForNdsGroup) SetViolatedPolicies(v []string)`

SetViolatedPolicies sets ViolatedPolicies field to given value.

### HasViolatedPolicies

`func (o *EventViewForNdsGroup) HasViolatedPolicies() bool`

HasViolatedPolicies returns a boolean if a field has been set.

### SetViolatedPoliciesNil

`func (o *EventViewForNdsGroup) SetViolatedPoliciesNil()`

SetViolatedPoliciesNil sets ViolatedPolicies to an explicit JSON null when marshaled, overriding any value previously set with SetViolatedPolicies. Calling SetViolatedPolicies again clears the null override.

### GetViolationAction

`func (o *EventViewForNdsGroup) GetViolationAction() string`

GetViolationAction returns the ViolationAction field if non-nil, zero value otherwise.

### GetViolationActionOk

`func (o *EventViewForNdsGroup) GetViolationActionOk() (*string, bool)`

GetViolationActionOk returns a tuple with the ViolationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationAction

`func (o *EventViewForNdsGroup) SetViolationAction(v string)`

SetViolationAction sets ViolationAction field to given value.

### HasViolationAction

`func (o *EventViewForNdsGroup) HasViolationAction() bool`

HasViolationAction returns a boolean if a field has been set.

### SetViolationActionNil

`func (o *EventViewForNdsGroup) SetViolationActionNil()`

SetViolationActionNil sets ViolationAction to an explicit JSON null when marshaled, overriding any value previously set with SetViolationAction. Calling SetViolationAction again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


