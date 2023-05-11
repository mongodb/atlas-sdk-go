# OpsGenie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Key that allows MongoDB Cloud to access your Opsgenie account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | 
**Region** | Pointer to **string** | Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie API. | [optional] [default to "US"]
**Type** | Pointer to **string** | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type. | [optional] 

## Methods

### NewOpsGenie

`func NewOpsGenie(apiKey string, ) *OpsGenie`

NewOpsGenie instantiates a new OpsGenie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsGenieWithDefaults

`func NewOpsGenieWithDefaults() *OpsGenie`

NewOpsGenieWithDefaults instantiates a new OpsGenie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *OpsGenie) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OpsGenie) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OpsGenie) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRegion

`func (o *OpsGenie) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OpsGenie) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OpsGenie) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OpsGenie) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *OpsGenie) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpsGenie) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpsGenie) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpsGenie) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


