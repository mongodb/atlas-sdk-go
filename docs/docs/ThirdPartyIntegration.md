# ThirdPartyIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Integration id. | [optional] 
**Type** | Pointer to **string** | Integration type  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type. | [optional] 
**ApiKey** | Pointer to **string** | Key that allows MongoDB Cloud to access your Datadog account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  Alternatively: Key that allows MongoDB Cloud to access your Opsgenie account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  Alternatively: Key that allows MongoDB Cloud to access your VictorOps account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**Region** | Pointer to **string** | Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API.  To learn more about Datadog&#39;s regions, see &lt;a href&#x3D;\&quot;https://docs.datadoghq.com/getting_started/site/\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;Datadog Sites&lt;/a&gt;.  Alternatively: Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie API.  Alternatively: PagerDuty region that indicates the API Uniform Resource Locator (URL) to use. | [optional] 
**MicrosoftTeamsWebhookUrl** | Pointer to **string** | Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.  **NOTE**: When you view or edit the alert for a Microsoft Teams notification, the URL appears partially redacted. | [optional] 
**AccountId** | Pointer to **string** | Unique 40-hexadecimal digit string that identifies your New Relic account. | [optional] 
**LicenseKey** | Pointer to **string** | Unique 40-hexadecimal digit string that identifies your New Relic license.  **IMPORTANT**: Effective Wednesday, June 16th, 2021, New Relic no longer supports the plugin-based integration with MongoDB. We do not recommend that you sign up for the plugin-based integration. To learn more, see the &lt;a href&#x3D;\&quot;https://discuss.newrelic.com/t/new-relic-plugin-eol-wednesday-june-16th-2021/127267\&quot; target&#x3D;\&quot;_blank\&quot;&gt;New Relic Plugin EOL Statement&lt;/a&gt; Consider configuring an alternative monitoring integration before June 16th to maintain visibility into your MongoDB deployments. | [optional] 
**ReadToken** | Pointer to **string** | Query key used to access your New Relic account. | [optional] 
**WriteToken** | Pointer to **string** | Insert key associated with your New Relic account. | [optional] 
**ServiceKey** | Pointer to **string** | Service key associated with your PagerDuty account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone has activated the Prometheus integration. | [optional] 
**Password** | Pointer to **string** | Password needed to allow MongoDB Cloud to access your Prometheus account. | [optional] 
**ServiceDiscovery** | Pointer to **string** | Desired method to discover the Prometheus service. | [optional] 
**Username** | Pointer to **string** | Human-readable label that identifies your Prometheus incoming webhook. | [optional] 
**ApiToken** | Pointer to **string** | Key that allows MongoDB Cloud to access your Slack account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  **IMPORTANT**: Slack integrations now use the OAuth2 verification method and must  be initially configured, or updated from a legacy integration, through the Atlas  third-party service integrations page. Legacy tokens will soon no longer be  supported. | [optional] 
**ChannelName** | Pointer to **string** | Name of the Slack channel to which MongoDB Cloud sends alert notifications. | [optional] 
**TeamName** | Pointer to **string** | Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration. | [optional] 
**RoutingKey** | Pointer to **string** | Routing key associated with your Splunk On-Call account. | [optional] 
**Secret** | Pointer to **string** | An optional field returned if your webhook is configured with a secret.  **NOTE**: When you view or edit the alert for a webhook notification, the secret appears completely redacted. | [optional] 
**Url** | Pointer to **string** | Endpoint web address to which MongoDB Cloud sends notifications.  **NOTE**: When you view or edit the alert for a webhook notification, the URL appears partially redacted. | [optional] 

## Methods

### NewThirdPartyIntegration

`func NewThirdPartyIntegration() *ThirdPartyIntegration`

NewThirdPartyIntegration instantiates a new ThirdPartyIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyIntegrationWithDefaults

`func NewThirdPartyIntegrationWithDefaults() *ThirdPartyIntegration`

NewThirdPartyIntegrationWithDefaults instantiates a new ThirdPartyIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyIntegration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThirdPartyIntegration) HasId() bool`

HasId returns a boolean if a field has been set.
### GetType

`func (o *ThirdPartyIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdPartyIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdPartyIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThirdPartyIntegration) HasType() bool`

HasType returns a boolean if a field has been set.
### GetApiKey

`func (o *ThirdPartyIntegration) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ThirdPartyIntegration) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ThirdPartyIntegration) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ThirdPartyIntegration) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.
### GetRegion

`func (o *ThirdPartyIntegration) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ThirdPartyIntegration) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ThirdPartyIntegration) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ThirdPartyIntegration) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetMicrosoftTeamsWebhookUrl

`func (o *ThirdPartyIntegration) GetMicrosoftTeamsWebhookUrl() string`

GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field if non-nil, zero value otherwise.

### GetMicrosoftTeamsWebhookUrlOk

`func (o *ThirdPartyIntegration) GetMicrosoftTeamsWebhookUrlOk() (*string, bool)`

GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftTeamsWebhookUrl

`func (o *ThirdPartyIntegration) SetMicrosoftTeamsWebhookUrl(v string)`

SetMicrosoftTeamsWebhookUrl sets MicrosoftTeamsWebhookUrl field to given value.

### HasMicrosoftTeamsWebhookUrl

`func (o *ThirdPartyIntegration) HasMicrosoftTeamsWebhookUrl() bool`

HasMicrosoftTeamsWebhookUrl returns a boolean if a field has been set.
### GetAccountId

`func (o *ThirdPartyIntegration) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ThirdPartyIntegration) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ThirdPartyIntegration) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ThirdPartyIntegration) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.
### GetLicenseKey

`func (o *ThirdPartyIntegration) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *ThirdPartyIntegration) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *ThirdPartyIntegration) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *ThirdPartyIntegration) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.
### GetReadToken

`func (o *ThirdPartyIntegration) GetReadToken() string`

GetReadToken returns the ReadToken field if non-nil, zero value otherwise.

### GetReadTokenOk

`func (o *ThirdPartyIntegration) GetReadTokenOk() (*string, bool)`

GetReadTokenOk returns a tuple with the ReadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadToken

`func (o *ThirdPartyIntegration) SetReadToken(v string)`

SetReadToken sets ReadToken field to given value.

### HasReadToken

`func (o *ThirdPartyIntegration) HasReadToken() bool`

HasReadToken returns a boolean if a field has been set.
### GetWriteToken

`func (o *ThirdPartyIntegration) GetWriteToken() string`

GetWriteToken returns the WriteToken field if non-nil, zero value otherwise.

### GetWriteTokenOk

`func (o *ThirdPartyIntegration) GetWriteTokenOk() (*string, bool)`

GetWriteTokenOk returns a tuple with the WriteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToken

`func (o *ThirdPartyIntegration) SetWriteToken(v string)`

SetWriteToken sets WriteToken field to given value.

### HasWriteToken

`func (o *ThirdPartyIntegration) HasWriteToken() bool`

HasWriteToken returns a boolean if a field has been set.
### GetServiceKey

`func (o *ThirdPartyIntegration) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *ThirdPartyIntegration) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *ThirdPartyIntegration) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *ThirdPartyIntegration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.
### GetEnabled

`func (o *ThirdPartyIntegration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyIntegration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyIntegration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ThirdPartyIntegration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.
### GetPassword

`func (o *ThirdPartyIntegration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ThirdPartyIntegration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ThirdPartyIntegration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ThirdPartyIntegration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.
### GetServiceDiscovery

`func (o *ThirdPartyIntegration) GetServiceDiscovery() string`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *ThirdPartyIntegration) GetServiceDiscoveryOk() (*string, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *ThirdPartyIntegration) SetServiceDiscovery(v string)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *ThirdPartyIntegration) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.
### GetUsername

`func (o *ThirdPartyIntegration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ThirdPartyIntegration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ThirdPartyIntegration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ThirdPartyIntegration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.
### GetApiToken

`func (o *ThirdPartyIntegration) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ThirdPartyIntegration) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ThirdPartyIntegration) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *ThirdPartyIntegration) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.
### GetChannelName

`func (o *ThirdPartyIntegration) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ThirdPartyIntegration) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ThirdPartyIntegration) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ThirdPartyIntegration) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.
### GetTeamName

`func (o *ThirdPartyIntegration) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ThirdPartyIntegration) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ThirdPartyIntegration) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *ThirdPartyIntegration) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.
### GetRoutingKey

`func (o *ThirdPartyIntegration) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *ThirdPartyIntegration) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *ThirdPartyIntegration) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *ThirdPartyIntegration) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.
### GetSecret

`func (o *ThirdPartyIntegration) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ThirdPartyIntegration) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ThirdPartyIntegration) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ThirdPartyIntegration) HasSecret() bool`

HasSecret returns a boolean if a field has been set.
### GetUrl

`func (o *ThirdPartyIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThirdPartyIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThirdPartyIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ThirdPartyIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


