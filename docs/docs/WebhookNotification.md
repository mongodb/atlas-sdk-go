# WebhookNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 
**WebhookSecret** | Pointer to **string** | Authentication secret for a webhook-based alert.  Atlas returns this value if you set &#x60;\&quot;notifications.[n].typeName\&quot; :\&quot;WEBHOOK\&quot;&#x60; and either: * You set &#x60;notification.[n].webhookSecret&#x60; to a non-empty string * You set a default webhookSecret either on the [Integrations](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/#std-label-third-party-integrations) page, or with the [Integrations API](#tag/Third-Party-Service-Integrations/operation/createIntegration)  **NOTE**: When you view or edit the alert for a webhook notification, the secret appears completely redacted. | [optional] 
**WebhookUrl** | Pointer to **string** | Target URL for a webhook-based alert.  Atlas returns this value if you set &#x60;\&quot;notifications.[n].typeName\&quot; :\&quot;WEBHOOK\&quot;&#x60; and either: * You set &#x60;notification.[n].webhookURL&#x60; to a non-empty string * You set a default webhookUrl either on the [Integrations](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/#std-label-third-party-integrations) page, or with the [Integrations API](#tag/Third-Party-Service-Integrations/operation/createIntegration)  **NOTE**: When you view or edit the alert for a Webhook URL notification, the URL appears partially redacted. | [optional] 

## Methods

### NewWebhookNotification

`func NewWebhookNotification(typeName string, ) *WebhookNotification`

NewWebhookNotification instantiates a new WebhookNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNotificationWithDefaults

`func NewWebhookNotificationWithDefaults() *WebhookNotification`

NewWebhookNotificationWithDefaults instantiates a new WebhookNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMin

`func (o *WebhookNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *WebhookNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *WebhookNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *WebhookNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetIntervalMin

`func (o *WebhookNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *WebhookNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *WebhookNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *WebhookNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetTypeName

`func (o *WebhookNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *WebhookNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *WebhookNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetWebhookSecret

`func (o *WebhookNotification) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *WebhookNotification) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *WebhookNotification) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *WebhookNotification) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *WebhookNotification) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *WebhookNotification) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *WebhookNotification) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *WebhookNotification) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


