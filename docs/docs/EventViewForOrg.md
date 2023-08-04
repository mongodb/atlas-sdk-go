# EventViewForOrg

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
**WhitelistEntry** | Pointer to **string** | Entry in the list of source host addresses that the API key accepts and this event targets. | [optional] [readonly] 
**TeamId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization team associated with this event. | [optional] [readonly] 
**TargetUsername** | Pointer to **string** | Email address for the console user that this event targets. The resource returns this parameter when &#x60;\&quot;eventTypeName\&quot; : \&quot;USER\&quot;&#x60;. | [optional] [readonly] 
**ResourceId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the resource associated with the event. | [optional] [readonly] 
**ResourceType** | Pointer to **string** | Unique identifier of resource type. | [optional] 

## Methods

### NewEventViewForOrg

`func NewEventViewForOrg() *EventViewForOrg`

NewEventViewForOrg instantiates a new EventViewForOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventViewForOrgWithDefaults

`func NewEventViewForOrgWithDefaults() *EventViewForOrg`

NewEventViewForOrgWithDefaults instantiates a new EventViewForOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *EventViewForOrg) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *EventViewForOrg) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *EventViewForOrg) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *EventViewForOrg) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.
### GetCreated

`func (o *EventViewForOrg) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EventViewForOrg) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EventViewForOrg) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EventViewForOrg) HasCreated() bool`

HasCreated returns a boolean if a field has been set.
### GetEventTypeName

`func (o *EventViewForOrg) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *EventViewForOrg) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *EventViewForOrg) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.

### HasEventTypeName

`func (o *EventViewForOrg) HasEventTypeName() bool`

HasEventTypeName returns a boolean if a field has been set.
### GetGroupId

`func (o *EventViewForOrg) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *EventViewForOrg) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *EventViewForOrg) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *EventViewForOrg) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *EventViewForOrg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventViewForOrg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventViewForOrg) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventViewForOrg) HasId() bool`

HasId returns a boolean if a field has been set.
### GetIsGlobalAdmin

`func (o *EventViewForOrg) GetIsGlobalAdmin() bool`

GetIsGlobalAdmin returns the IsGlobalAdmin field if non-nil, zero value otherwise.

### GetIsGlobalAdminOk

`func (o *EventViewForOrg) GetIsGlobalAdminOk() (*bool, bool)`

GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalAdmin

`func (o *EventViewForOrg) SetIsGlobalAdmin(v bool)`

SetIsGlobalAdmin sets IsGlobalAdmin field to given value.

### HasIsGlobalAdmin

`func (o *EventViewForOrg) HasIsGlobalAdmin() bool`

HasIsGlobalAdmin returns a boolean if a field has been set.
### GetLinks

`func (o *EventViewForOrg) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventViewForOrg) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventViewForOrg) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventViewForOrg) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOrgId

`func (o *EventViewForOrg) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EventViewForOrg) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EventViewForOrg) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EventViewForOrg) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.
### GetPublicKey

`func (o *EventViewForOrg) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *EventViewForOrg) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *EventViewForOrg) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *EventViewForOrg) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.
### GetRaw

`func (o *EventViewForOrg) GetRaw() Raw`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *EventViewForOrg) GetRawOk() (*Raw, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *EventViewForOrg) SetRaw(v Raw)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *EventViewForOrg) HasRaw() bool`

HasRaw returns a boolean if a field has been set.
### GetRemoteAddress

`func (o *EventViewForOrg) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *EventViewForOrg) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *EventViewForOrg) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *EventViewForOrg) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.
### GetUserId

`func (o *EventViewForOrg) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventViewForOrg) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventViewForOrg) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventViewForOrg) HasUserId() bool`

HasUserId returns a boolean if a field has been set.
### GetUsername

`func (o *EventViewForOrg) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EventViewForOrg) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EventViewForOrg) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EventViewForOrg) HasUsername() bool`

HasUsername returns a boolean if a field has been set.
### GetAlertId

`func (o *EventViewForOrg) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *EventViewForOrg) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *EventViewForOrg) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *EventViewForOrg) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.
### GetAlertConfigId

`func (o *EventViewForOrg) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *EventViewForOrg) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *EventViewForOrg) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.

### HasAlertConfigId

`func (o *EventViewForOrg) HasAlertConfigId() bool`

HasAlertConfigId returns a boolean if a field has been set.
### GetInvoiceId

`func (o *EventViewForOrg) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *EventViewForOrg) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *EventViewForOrg) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *EventViewForOrg) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.
### GetPaymentId

`func (o *EventViewForOrg) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *EventViewForOrg) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *EventViewForOrg) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *EventViewForOrg) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.
### GetWhitelistEntry

`func (o *EventViewForOrg) GetWhitelistEntry() string`

GetWhitelistEntry returns the WhitelistEntry field if non-nil, zero value otherwise.

### GetWhitelistEntryOk

`func (o *EventViewForOrg) GetWhitelistEntryOk() (*string, bool)`

GetWhitelistEntryOk returns a tuple with the WhitelistEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEntry

`func (o *EventViewForOrg) SetWhitelistEntry(v string)`

SetWhitelistEntry sets WhitelistEntry field to given value.

### HasWhitelistEntry

`func (o *EventViewForOrg) HasWhitelistEntry() bool`

HasWhitelistEntry returns a boolean if a field has been set.
### GetTeamId

`func (o *EventViewForOrg) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *EventViewForOrg) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *EventViewForOrg) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *EventViewForOrg) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.
### GetTargetUsername

`func (o *EventViewForOrg) GetTargetUsername() string`

GetTargetUsername returns the TargetUsername field if non-nil, zero value otherwise.

### GetTargetUsernameOk

`func (o *EventViewForOrg) GetTargetUsernameOk() (*string, bool)`

GetTargetUsernameOk returns a tuple with the TargetUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUsername

`func (o *EventViewForOrg) SetTargetUsername(v string)`

SetTargetUsername sets TargetUsername field to given value.

### HasTargetUsername

`func (o *EventViewForOrg) HasTargetUsername() bool`

HasTargetUsername returns a boolean if a field has been set.
### GetResourceId

`func (o *EventViewForOrg) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *EventViewForOrg) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *EventViewForOrg) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *EventViewForOrg) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.
### GetResourceType

`func (o *EventViewForOrg) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EventViewForOrg) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EventViewForOrg) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *EventViewForOrg) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


