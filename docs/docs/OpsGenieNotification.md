# OpsGenieNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**OpsGenieApiKey** | Pointer to **string** | API Key that MongoDB Cloud needs to send this notification via Opsgenie. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;OPS_GENIE\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**OpsGenieRegion** | Pointer to **string** | Opsgenie region that indicates which API Uniform Resource Locator (URL) to use. | [optional] [default to "US"]
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 

## Methods

### NewOpsGenieNotification

`func NewOpsGenieNotification(typeName string, ) *OpsGenieNotification`

NewOpsGenieNotification instantiates a new OpsGenieNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsGenieNotificationWithDefaults

`func NewOpsGenieNotificationWithDefaults() *OpsGenieNotification`

NewOpsGenieNotificationWithDefaults instantiates a new OpsGenieNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMin

`func (o *OpsGenieNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *OpsGenieNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *OpsGenieNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *OpsGenieNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetIntervalMin

`func (o *OpsGenieNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *OpsGenieNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *OpsGenieNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *OpsGenieNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetOpsGenieApiKey

`func (o *OpsGenieNotification) GetOpsGenieApiKey() string`

GetOpsGenieApiKey returns the OpsGenieApiKey field if non-nil, zero value otherwise.

### GetOpsGenieApiKeyOk

`func (o *OpsGenieNotification) GetOpsGenieApiKeyOk() (*string, bool)`

GetOpsGenieApiKeyOk returns a tuple with the OpsGenieApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsGenieApiKey

`func (o *OpsGenieNotification) SetOpsGenieApiKey(v string)`

SetOpsGenieApiKey sets OpsGenieApiKey field to given value.

### HasOpsGenieApiKey

`func (o *OpsGenieNotification) HasOpsGenieApiKey() bool`

HasOpsGenieApiKey returns a boolean if a field has been set.

### GetOpsGenieRegion

`func (o *OpsGenieNotification) GetOpsGenieRegion() string`

GetOpsGenieRegion returns the OpsGenieRegion field if non-nil, zero value otherwise.

### GetOpsGenieRegionOk

`func (o *OpsGenieNotification) GetOpsGenieRegionOk() (*string, bool)`

GetOpsGenieRegionOk returns a tuple with the OpsGenieRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsGenieRegion

`func (o *OpsGenieNotification) SetOpsGenieRegion(v string)`

SetOpsGenieRegion sets OpsGenieRegion field to given value.

### HasOpsGenieRegion

`func (o *OpsGenieNotification) HasOpsGenieRegion() bool`

HasOpsGenieRegion returns a boolean if a field has been set.

### GetTypeName

`func (o *OpsGenieNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *OpsGenieNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *OpsGenieNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


