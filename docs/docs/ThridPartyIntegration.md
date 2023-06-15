# ThridPartyIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Key that allows MongoDB Cloud to access your VictorOps account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | 
**Region** | Pointer to **string** | PagerDuty region that indicates the API Uniform Resource Locator (URL) to use. | [optional] 
**Type** | Pointer to **string** | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type. | [optional] 
**MicrosoftTeamsWebhookUrl** | **string** | Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.  **NOTE**: When you view or edit the alert for a Microsoft Teams notification, the URL appears partially redacted. | 
**AccountId** | **string** | Unique 40-hexadecimal digit string that identifies your New Relic account. | 
**LicenseKey** | **string** | Unique 40-hexadecimal digit string that identifies your New Relic license.  **IMPORTANT**: Effective Wednesday, June 16th, 2021, New Relic no longer supports the plugin-based integration with MongoDB. We do not recommend that you sign up for the plugin-based integration. To learn more, see the &lt;a href&#x3D;\&quot;https://discuss.newrelic.com/t/new-relic-plugin-eol-wednesday-june-16th-2021/127267\&quot; target&#x3D;\&quot;_blank\&quot;&gt;New Relic Plugin EOL Statement&lt;/a&gt; Consider configuring an alternative monitoring integration before June 16th to maintain visibility into your MongoDB deployments. | 
**ReadToken** | **string** | Query key used to access your New Relic account. | 
**WriteToken** | **string** | Insert key associated with your New Relic account. | 
**ServiceKey** | **string** | Service key associated with your PagerDuty account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | 
**Enabled** | **bool** | Flag that indicates whether someone has activated the Prometheus integration. | 
**ListenAddress** | Pointer to **string** | Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics. | [optional] [default to ":9216"]
**Password** | Pointer to **string** |  | [optional] 
**RateLimitInterval** | Pointer to **int** |  | [optional] 
**Scheme** | **string** | Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud. | 
**ServiceDiscovery** | **string** | Desired method to discover the Prometheus service. | 
**TlsPemPath** | Pointer to **string** | Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host. | [optional] 
**Username** | **string** | Human-readable label that identifies your Prometheus incoming webhook. | 
**ApiToken** | **string** | Key that allows MongoDB Cloud to access your Slack account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  **IMPORTANT**: Slack integrations now use the OAuth2 verification method and must  be initially configured, or updated from a legacy integration, through the Atlas  third-party service integrations page. Legacy tokens will soon no longer be  supported. | 
**ChannelName** | **NullableString** | Name of the Slack channel to which MongoDB Cloud sends alert notifications. | 
**TeamName** | Pointer to **string** | Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration. | [optional] 
**RoutingKey** | Pointer to **string** | Routing key associated with your Splunk On-Call account. | [optional] 
**Secret** | Pointer to **string** | An optional field returned if your webhook is configured with a secret.  **NOTE**: When you view or edit the alert for a webhook notification, the secret appears completely redacted. | [optional] 
**Url** | **string** | Endpoint web address to which MongoDB Cloud sends notifications.  **NOTE**: When you view or edit the alert for a webhook notification, the URL appears partially redacted. | 

## Methods

### NewThridPartyIntegration

`func NewThridPartyIntegration(apiKey string, microsoftTeamsWebhookUrl string, accountId string, licenseKey string, readToken string, writeToken string, serviceKey string, enabled bool, scheme string, serviceDiscovery string, username string, apiToken string, channelName NullableString, url string, ) *ThridPartyIntegration`

NewThridPartyIntegration instantiates a new ThridPartyIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThridPartyIntegrationWithDefaults

`func NewThridPartyIntegrationWithDefaults() *ThridPartyIntegration`

NewThridPartyIntegrationWithDefaults instantiates a new ThridPartyIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *ThridPartyIntegration) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ThridPartyIntegration) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ThridPartyIntegration) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRegion

`func (o *ThridPartyIntegration) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ThridPartyIntegration) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ThridPartyIntegration) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ThridPartyIntegration) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *ThridPartyIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThridPartyIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThridPartyIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThridPartyIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMicrosoftTeamsWebhookUrl

`func (o *ThridPartyIntegration) GetMicrosoftTeamsWebhookUrl() string`

GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field if non-nil, zero value otherwise.

### GetMicrosoftTeamsWebhookUrlOk

`func (o *ThridPartyIntegration) GetMicrosoftTeamsWebhookUrlOk() (*string, bool)`

GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftTeamsWebhookUrl

`func (o *ThridPartyIntegration) SetMicrosoftTeamsWebhookUrl(v string)`

SetMicrosoftTeamsWebhookUrl sets MicrosoftTeamsWebhookUrl field to given value.


### GetAccountId

`func (o *ThridPartyIntegration) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ThridPartyIntegration) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ThridPartyIntegration) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetLicenseKey

`func (o *ThridPartyIntegration) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *ThridPartyIntegration) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *ThridPartyIntegration) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.


### GetReadToken

`func (o *ThridPartyIntegration) GetReadToken() string`

GetReadToken returns the ReadToken field if non-nil, zero value otherwise.

### GetReadTokenOk

`func (o *ThridPartyIntegration) GetReadTokenOk() (*string, bool)`

GetReadTokenOk returns a tuple with the ReadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadToken

`func (o *ThridPartyIntegration) SetReadToken(v string)`

SetReadToken sets ReadToken field to given value.


### GetWriteToken

`func (o *ThridPartyIntegration) GetWriteToken() string`

GetWriteToken returns the WriteToken field if non-nil, zero value otherwise.

### GetWriteTokenOk

`func (o *ThridPartyIntegration) GetWriteTokenOk() (*string, bool)`

GetWriteTokenOk returns a tuple with the WriteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToken

`func (o *ThridPartyIntegration) SetWriteToken(v string)`

SetWriteToken sets WriteToken field to given value.


### GetServiceKey

`func (o *ThridPartyIntegration) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *ThridPartyIntegration) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *ThridPartyIntegration) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetEnabled

`func (o *ThridPartyIntegration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThridPartyIntegration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThridPartyIntegration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetListenAddress

`func (o *ThridPartyIntegration) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *ThridPartyIntegration) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *ThridPartyIntegration) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *ThridPartyIntegration) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetPassword

`func (o *ThridPartyIntegration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ThridPartyIntegration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ThridPartyIntegration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ThridPartyIntegration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRateLimitInterval

`func (o *ThridPartyIntegration) GetRateLimitInterval() int`

GetRateLimitInterval returns the RateLimitInterval field if non-nil, zero value otherwise.

### GetRateLimitIntervalOk

`func (o *ThridPartyIntegration) GetRateLimitIntervalOk() (*int, bool)`

GetRateLimitIntervalOk returns a tuple with the RateLimitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitInterval

`func (o *ThridPartyIntegration) SetRateLimitInterval(v int)`

SetRateLimitInterval sets RateLimitInterval field to given value.

### HasRateLimitInterval

`func (o *ThridPartyIntegration) HasRateLimitInterval() bool`

HasRateLimitInterval returns a boolean if a field has been set.

### GetScheme

`func (o *ThridPartyIntegration) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ThridPartyIntegration) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ThridPartyIntegration) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetServiceDiscovery

`func (o *ThridPartyIntegration) GetServiceDiscovery() string`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *ThridPartyIntegration) GetServiceDiscoveryOk() (*string, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *ThridPartyIntegration) SetServiceDiscovery(v string)`

SetServiceDiscovery sets ServiceDiscovery field to given value.


### GetTlsPemPath

`func (o *ThridPartyIntegration) GetTlsPemPath() string`

GetTlsPemPath returns the TlsPemPath field if non-nil, zero value otherwise.

### GetTlsPemPathOk

`func (o *ThridPartyIntegration) GetTlsPemPathOk() (*string, bool)`

GetTlsPemPathOk returns a tuple with the TlsPemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPemPath

`func (o *ThridPartyIntegration) SetTlsPemPath(v string)`

SetTlsPemPath sets TlsPemPath field to given value.

### HasTlsPemPath

`func (o *ThridPartyIntegration) HasTlsPemPath() bool`

HasTlsPemPath returns a boolean if a field has been set.

### GetUsername

`func (o *ThridPartyIntegration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ThridPartyIntegration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ThridPartyIntegration) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetApiToken

`func (o *ThridPartyIntegration) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ThridPartyIntegration) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ThridPartyIntegration) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetChannelName

`func (o *ThridPartyIntegration) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ThridPartyIntegration) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ThridPartyIntegration) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### SetChannelNameNil

`func (o *ThridPartyIntegration) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *ThridPartyIntegration) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetTeamName

`func (o *ThridPartyIntegration) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ThridPartyIntegration) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ThridPartyIntegration) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *ThridPartyIntegration) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetRoutingKey

`func (o *ThridPartyIntegration) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *ThridPartyIntegration) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *ThridPartyIntegration) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *ThridPartyIntegration) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### GetSecret

`func (o *ThridPartyIntegration) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ThridPartyIntegration) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ThridPartyIntegration) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ThridPartyIntegration) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUrl

`func (o *ThridPartyIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThridPartyIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThridPartyIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


