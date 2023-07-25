# AlertsNotificationRootForGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatadogApiKey** | Pointer to **string** | Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog. You can find this API key in the Datadog dashboard. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;DATADOG\&quot;&#x60;.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**DatadogRegion** | Pointer to **string** | Datadog region that indicates which API Uniform Resource Locator (URL) to use. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;DATADOG\&quot;&#x60;.  To learn more about Datadog&#39;s regions, see &lt;a href&#x3D;\&quot;https://docs.datadoghq.com/getting_started/site/\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;Datadog Sites&lt;/a&gt;. | [optional] [default to "US"]
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**NotifierId** | Pointer to **string** | The notifierId is a system-generated unique identifier assigned to each notification method. This is needed when updating third-party notifications without requiring explicit authentication credentials. | [optional] 
**TypeName** | Pointer to **string** | Human-readable label that displays the alert notification type. | [optional] 
**EmailAddress** | Pointer to **string** | Email address to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;EMAIL\&quot;&#x60;. You don’t need to set this value to send emails to individual or groups of MongoDB Cloud users including:  - specific MongoDB Cloud users (&#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60;) - MongoDB Cloud users with specific project roles (&#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60;) - MongoDB Cloud users with specific organization roles (&#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60;) - MongoDB Cloud teams (&#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;TEAM\&quot;&#x60;)  To send emails to one MongoDB Cloud user or grouping of users, set the &#x60;notifications.[n].emailEnabled&#x60; parameter. | [optional] 
**EmailEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should send email notifications. The resource requires this parameter when one of the following values have been set:  - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60; | [optional] 
**Roles** | Pointer to **[]string** | List that contains the one or more [organization](https://dochub.mongodb.org/core/atlas-org-roles) or [project roles](https://dochub.mongodb.org/core/atlas-proj-roles) that receive the configured alert. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60; or &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60;. If you include this parameter, MongoDB Cloud sends alerts only to users assigned the roles you specify in the array. If you omit this parameter, MongoDB Cloud sends alerts to users assigned any role. | [optional] 
**SmsEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should send text message notifications. The resource requires this parameter when one of the following values have been set:  - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60; | [optional] 
**NotificationToken** | Pointer to **string** | HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;HIP_CHAT\&quot;&#x60;\&quot;. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**RoomName** | Pointer to **string** | HipChat API room name to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;HIP_CHAT\&quot;&#x60;\&quot;. | [optional] 
**MicrosoftTeamsWebhookUrl** | Pointer to **string** | Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;MICROSOFT_TEAMS\&quot;&#x60;. If the URL later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: When you view or edit the alert for a Microsoft Teams notification, the URL appears partially redacted. | [optional] 
**OpsGenieApiKey** | Pointer to **string** | API Key that MongoDB Cloud needs to send this notification via Opsgenie. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;OPS_GENIE\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**OpsGenieRegion** | Pointer to **string** | Opsgenie region that indicates which API Uniform Resource Locator (URL) to use. | [optional] [default to "US"]
**Region** | Pointer to **string** | PagerDuty region that indicates which API Uniform Resource Locator (URL) to use. | [optional] [default to "US"]
**ServiceKey** | Pointer to **string** | PagerDuty service key that MongoDB Cloud needs to send notifications via PagerDuty. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;PAGER_DUTY\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**ApiToken** | Pointer to **string** | Slack API token or Bot token that MongoDB Cloud needs to send alert notifications via Slack. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;SLACK\&quot;&#x60;. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes the token.   **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**ChannelName** | Pointer to **string** | Name of the Slack channel to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;SLACK\&quot;&#x60;. | [optional] 
**MobileNumber** | Pointer to **string** | Mobile phone number to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;SMS\&quot;&#x60;. | [optional] 
**TeamId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies one MongoDB Cloud team. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;TEAM\&quot;&#x60;. | [optional] 
**TeamName** | Pointer to **string** | Name of the MongoDB Cloud team that receives this notification. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;TEAM\&quot;&#x60;. | [optional] 
**Username** | Pointer to **string** | MongoDB Cloud username of the person to whom MongoDB Cloud sends notifications. Specify only MongoDB Cloud users who belong to the project that owns the alert configuration. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60;. | [optional] 
**VictorOpsApiKey** | Pointer to **string** | API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;VICTOR_OPS\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**VictorOpsRoutingKey** | Pointer to **string** | Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;VICTOR_OPS\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it. | [optional] 
**WebhookSecret** | Pointer to **string** | Authentication secret for a webhook-based alert.  Atlas returns this value if you set &#x60;\&quot;notifications.[n].typeName\&quot; :\&quot;WEBHOOK\&quot;&#x60; and either: * You set &#x60;notification.[n].webhookSecret&#x60; to a non-empty string * You set a default webhookSecret either on the [Integrations](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/#std-label-third-party-integrations) page, or with the [Integrations API](#tag/Third-Party-Service-Integrations/operation/createIntegration)  **NOTE**: When you view or edit the alert for a webhook notification, the secret appears completely redacted. | [optional] 
**WebhookUrl** | Pointer to **string** | Target URL for a webhook-based alert.  Atlas returns this value if you set &#x60;\&quot;notifications.[n].typeName\&quot; :\&quot;WEBHOOK\&quot;&#x60; and either: * You set &#x60;notification.[n].webhookURL&#x60; to a non-empty string * You set a default webhookUrl either on the [Integrations](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/#std-label-third-party-integrations) page, or with the [Integrations API](#tag/Third-Party-Service-Integrations/operation/createIntegration)  **NOTE**: When you view or edit the alert for a Webhook URL notification, the URL appears partially redacted. | [optional] 

## Methods

### NewAlertsNotificationRootForGroup

`func NewAlertsNotificationRootForGroup() *AlertsNotificationRootForGroup`

NewAlertsNotificationRootForGroup instantiates a new AlertsNotificationRootForGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsNotificationRootForGroupWithDefaults

`func NewAlertsNotificationRootForGroupWithDefaults() *AlertsNotificationRootForGroup`

NewAlertsNotificationRootForGroupWithDefaults instantiates a new AlertsNotificationRootForGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatadogApiKey

`func (o *AlertsNotificationRootForGroup) GetDatadogApiKey() string`

GetDatadogApiKey returns the DatadogApiKey field if non-nil, zero value otherwise.

### GetDatadogApiKeyOk

`func (o *AlertsNotificationRootForGroup) GetDatadogApiKeyOk() (*string, bool)`

GetDatadogApiKeyOk returns a tuple with the DatadogApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadogApiKey

`func (o *AlertsNotificationRootForGroup) SetDatadogApiKey(v string)`

SetDatadogApiKey sets DatadogApiKey field to given value.

### HasDatadogApiKey

`func (o *AlertsNotificationRootForGroup) HasDatadogApiKey() bool`

HasDatadogApiKey returns a boolean if a field has been set.
### GetDatadogRegion

`func (o *AlertsNotificationRootForGroup) GetDatadogRegion() string`

GetDatadogRegion returns the DatadogRegion field if non-nil, zero value otherwise.

### GetDatadogRegionOk

`func (o *AlertsNotificationRootForGroup) GetDatadogRegionOk() (*string, bool)`

GetDatadogRegionOk returns a tuple with the DatadogRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadogRegion

`func (o *AlertsNotificationRootForGroup) SetDatadogRegion(v string)`

SetDatadogRegion sets DatadogRegion field to given value.

### HasDatadogRegion

`func (o *AlertsNotificationRootForGroup) HasDatadogRegion() bool`

HasDatadogRegion returns a boolean if a field has been set.
### GetDelayMin

`func (o *AlertsNotificationRootForGroup) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *AlertsNotificationRootForGroup) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *AlertsNotificationRootForGroup) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *AlertsNotificationRootForGroup) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.
### GetIntervalMin

`func (o *AlertsNotificationRootForGroup) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *AlertsNotificationRootForGroup) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *AlertsNotificationRootForGroup) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *AlertsNotificationRootForGroup) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.
### GetNotifierId

`func (o *AlertsNotificationRootForGroup) GetNotifierId() string`

GetNotifierId returns the NotifierId field if non-nil, zero value otherwise.

### GetNotifierIdOk

`func (o *AlertsNotificationRootForGroup) GetNotifierIdOk() (*string, bool)`

GetNotifierIdOk returns a tuple with the NotifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierId

`func (o *AlertsNotificationRootForGroup) SetNotifierId(v string)`

SetNotifierId sets NotifierId field to given value.

### HasNotifierId

`func (o *AlertsNotificationRootForGroup) HasNotifierId() bool`

HasNotifierId returns a boolean if a field has been set.
### GetTypeName

`func (o *AlertsNotificationRootForGroup) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AlertsNotificationRootForGroup) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AlertsNotificationRootForGroup) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *AlertsNotificationRootForGroup) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.
### GetEmailAddress

`func (o *AlertsNotificationRootForGroup) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *AlertsNotificationRootForGroup) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *AlertsNotificationRootForGroup) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *AlertsNotificationRootForGroup) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.
### GetEmailEnabled

`func (o *AlertsNotificationRootForGroup) GetEmailEnabled() bool`

GetEmailEnabled returns the EmailEnabled field if non-nil, zero value otherwise.

### GetEmailEnabledOk

`func (o *AlertsNotificationRootForGroup) GetEmailEnabledOk() (*bool, bool)`

GetEmailEnabledOk returns a tuple with the EmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEnabled

`func (o *AlertsNotificationRootForGroup) SetEmailEnabled(v bool)`

SetEmailEnabled sets EmailEnabled field to given value.

### HasEmailEnabled

`func (o *AlertsNotificationRootForGroup) HasEmailEnabled() bool`

HasEmailEnabled returns a boolean if a field has been set.
### GetRoles

`func (o *AlertsNotificationRootForGroup) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AlertsNotificationRootForGroup) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AlertsNotificationRootForGroup) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AlertsNotificationRootForGroup) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetSmsEnabled

`func (o *AlertsNotificationRootForGroup) GetSmsEnabled() bool`

GetSmsEnabled returns the SmsEnabled field if non-nil, zero value otherwise.

### GetSmsEnabledOk

`func (o *AlertsNotificationRootForGroup) GetSmsEnabledOk() (*bool, bool)`

GetSmsEnabledOk returns a tuple with the SmsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsEnabled

`func (o *AlertsNotificationRootForGroup) SetSmsEnabled(v bool)`

SetSmsEnabled sets SmsEnabled field to given value.

### HasSmsEnabled

`func (o *AlertsNotificationRootForGroup) HasSmsEnabled() bool`

HasSmsEnabled returns a boolean if a field has been set.
### GetNotificationToken

`func (o *AlertsNotificationRootForGroup) GetNotificationToken() string`

GetNotificationToken returns the NotificationToken field if non-nil, zero value otherwise.

### GetNotificationTokenOk

`func (o *AlertsNotificationRootForGroup) GetNotificationTokenOk() (*string, bool)`

GetNotificationTokenOk returns a tuple with the NotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationToken

`func (o *AlertsNotificationRootForGroup) SetNotificationToken(v string)`

SetNotificationToken sets NotificationToken field to given value.

### HasNotificationToken

`func (o *AlertsNotificationRootForGroup) HasNotificationToken() bool`

HasNotificationToken returns a boolean if a field has been set.
### GetRoomName

`func (o *AlertsNotificationRootForGroup) GetRoomName() string`

GetRoomName returns the RoomName field if non-nil, zero value otherwise.

### GetRoomNameOk

`func (o *AlertsNotificationRootForGroup) GetRoomNameOk() (*string, bool)`

GetRoomNameOk returns a tuple with the RoomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomName

`func (o *AlertsNotificationRootForGroup) SetRoomName(v string)`

SetRoomName sets RoomName field to given value.

### HasRoomName

`func (o *AlertsNotificationRootForGroup) HasRoomName() bool`

HasRoomName returns a boolean if a field has been set.
### GetMicrosoftTeamsWebhookUrl

`func (o *AlertsNotificationRootForGroup) GetMicrosoftTeamsWebhookUrl() string`

GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field if non-nil, zero value otherwise.

### GetMicrosoftTeamsWebhookUrlOk

`func (o *AlertsNotificationRootForGroup) GetMicrosoftTeamsWebhookUrlOk() (*string, bool)`

GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftTeamsWebhookUrl

`func (o *AlertsNotificationRootForGroup) SetMicrosoftTeamsWebhookUrl(v string)`

SetMicrosoftTeamsWebhookUrl sets MicrosoftTeamsWebhookUrl field to given value.

### HasMicrosoftTeamsWebhookUrl

`func (o *AlertsNotificationRootForGroup) HasMicrosoftTeamsWebhookUrl() bool`

HasMicrosoftTeamsWebhookUrl returns a boolean if a field has been set.
### GetOpsGenieApiKey

`func (o *AlertsNotificationRootForGroup) GetOpsGenieApiKey() string`

GetOpsGenieApiKey returns the OpsGenieApiKey field if non-nil, zero value otherwise.

### GetOpsGenieApiKeyOk

`func (o *AlertsNotificationRootForGroup) GetOpsGenieApiKeyOk() (*string, bool)`

GetOpsGenieApiKeyOk returns a tuple with the OpsGenieApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsGenieApiKey

`func (o *AlertsNotificationRootForGroup) SetOpsGenieApiKey(v string)`

SetOpsGenieApiKey sets OpsGenieApiKey field to given value.

### HasOpsGenieApiKey

`func (o *AlertsNotificationRootForGroup) HasOpsGenieApiKey() bool`

HasOpsGenieApiKey returns a boolean if a field has been set.
### GetOpsGenieRegion

`func (o *AlertsNotificationRootForGroup) GetOpsGenieRegion() string`

GetOpsGenieRegion returns the OpsGenieRegion field if non-nil, zero value otherwise.

### GetOpsGenieRegionOk

`func (o *AlertsNotificationRootForGroup) GetOpsGenieRegionOk() (*string, bool)`

GetOpsGenieRegionOk returns a tuple with the OpsGenieRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsGenieRegion

`func (o *AlertsNotificationRootForGroup) SetOpsGenieRegion(v string)`

SetOpsGenieRegion sets OpsGenieRegion field to given value.

### HasOpsGenieRegion

`func (o *AlertsNotificationRootForGroup) HasOpsGenieRegion() bool`

HasOpsGenieRegion returns a boolean if a field has been set.
### GetRegion

`func (o *AlertsNotificationRootForGroup) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AlertsNotificationRootForGroup) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AlertsNotificationRootForGroup) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AlertsNotificationRootForGroup) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetServiceKey

`func (o *AlertsNotificationRootForGroup) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *AlertsNotificationRootForGroup) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *AlertsNotificationRootForGroup) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *AlertsNotificationRootForGroup) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.
### GetApiToken

`func (o *AlertsNotificationRootForGroup) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *AlertsNotificationRootForGroup) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *AlertsNotificationRootForGroup) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *AlertsNotificationRootForGroup) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.
### GetChannelName

`func (o *AlertsNotificationRootForGroup) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *AlertsNotificationRootForGroup) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *AlertsNotificationRootForGroup) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *AlertsNotificationRootForGroup) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.
### GetMobileNumber

`func (o *AlertsNotificationRootForGroup) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *AlertsNotificationRootForGroup) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *AlertsNotificationRootForGroup) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *AlertsNotificationRootForGroup) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.
### GetTeamId

`func (o *AlertsNotificationRootForGroup) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AlertsNotificationRootForGroup) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AlertsNotificationRootForGroup) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *AlertsNotificationRootForGroup) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.
### GetTeamName

`func (o *AlertsNotificationRootForGroup) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *AlertsNotificationRootForGroup) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *AlertsNotificationRootForGroup) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *AlertsNotificationRootForGroup) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.
### GetUsername

`func (o *AlertsNotificationRootForGroup) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AlertsNotificationRootForGroup) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AlertsNotificationRootForGroup) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AlertsNotificationRootForGroup) HasUsername() bool`

HasUsername returns a boolean if a field has been set.
### GetVictorOpsApiKey

`func (o *AlertsNotificationRootForGroup) GetVictorOpsApiKey() string`

GetVictorOpsApiKey returns the VictorOpsApiKey field if non-nil, zero value otherwise.

### GetVictorOpsApiKeyOk

`func (o *AlertsNotificationRootForGroup) GetVictorOpsApiKeyOk() (*string, bool)`

GetVictorOpsApiKeyOk returns a tuple with the VictorOpsApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictorOpsApiKey

`func (o *AlertsNotificationRootForGroup) SetVictorOpsApiKey(v string)`

SetVictorOpsApiKey sets VictorOpsApiKey field to given value.

### HasVictorOpsApiKey

`func (o *AlertsNotificationRootForGroup) HasVictorOpsApiKey() bool`

HasVictorOpsApiKey returns a boolean if a field has been set.
### GetVictorOpsRoutingKey

`func (o *AlertsNotificationRootForGroup) GetVictorOpsRoutingKey() string`

GetVictorOpsRoutingKey returns the VictorOpsRoutingKey field if non-nil, zero value otherwise.

### GetVictorOpsRoutingKeyOk

`func (o *AlertsNotificationRootForGroup) GetVictorOpsRoutingKeyOk() (*string, bool)`

GetVictorOpsRoutingKeyOk returns a tuple with the VictorOpsRoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictorOpsRoutingKey

`func (o *AlertsNotificationRootForGroup) SetVictorOpsRoutingKey(v string)`

SetVictorOpsRoutingKey sets VictorOpsRoutingKey field to given value.

### HasVictorOpsRoutingKey

`func (o *AlertsNotificationRootForGroup) HasVictorOpsRoutingKey() bool`

HasVictorOpsRoutingKey returns a boolean if a field has been set.
### GetWebhookSecret

`func (o *AlertsNotificationRootForGroup) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *AlertsNotificationRootForGroup) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *AlertsNotificationRootForGroup) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *AlertsNotificationRootForGroup) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.
### GetWebhookUrl

`func (o *AlertsNotificationRootForGroup) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AlertsNotificationRootForGroup) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AlertsNotificationRootForGroup) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *AlertsNotificationRootForGroup) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


