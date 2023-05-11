# Slack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | **string** | Key that allows MongoDB Cloud to access your Slack account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  **IMPORTANT**: Slack integrations now use the OAuth2 verification method and must  be initially configured, or updated from a legacy integration, through the Atlas  third-party service integrations page. Legacy tokens will soon no longer be  supported. | 
**ChannelName** | **NullableString** | Name of the Slack channel to which MongoDB Cloud sends alert notifications. | 
**TeamName** | Pointer to **string** | Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration. | [optional] 
**Type** | Pointer to **string** | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type. | [optional] 

## Methods

### NewSlack

`func NewSlack(apiToken string, channelName NullableString, ) *Slack`

NewSlack instantiates a new Slack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackWithDefaults

`func NewSlackWithDefaults() *Slack`

NewSlackWithDefaults instantiates a new Slack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *Slack) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *Slack) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *Slack) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetChannelName

`func (o *Slack) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *Slack) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *Slack) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### SetChannelNameNil

`func (o *Slack) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *Slack) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetTeamName

`func (o *Slack) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *Slack) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *Slack) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *Slack) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetType

`func (o *Slack) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Slack) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Slack) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Slack) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


