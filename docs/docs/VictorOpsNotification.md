# VictorOpsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 
**VictorOpsApiKey** | Pointer to **string** | API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;VICTOR_OPS\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**VictorOpsRoutingKey** | Pointer to **string** | Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;VICTOR_OPS\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it. | [optional] 

## Methods

### NewVictorOpsNotification

`func NewVictorOpsNotification(typeName string, ) *VictorOpsNotification`

NewVictorOpsNotification instantiates a new VictorOpsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVictorOpsNotificationWithDefaults

`func NewVictorOpsNotificationWithDefaults() *VictorOpsNotification`

NewVictorOpsNotificationWithDefaults instantiates a new VictorOpsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMin

`func (o *VictorOpsNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *VictorOpsNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *VictorOpsNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *VictorOpsNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetIntervalMin

`func (o *VictorOpsNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *VictorOpsNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *VictorOpsNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *VictorOpsNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetTypeName

`func (o *VictorOpsNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *VictorOpsNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *VictorOpsNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetVictorOpsApiKey

`func (o *VictorOpsNotification) GetVictorOpsApiKey() string`

GetVictorOpsApiKey returns the VictorOpsApiKey field if non-nil, zero value otherwise.

### GetVictorOpsApiKeyOk

`func (o *VictorOpsNotification) GetVictorOpsApiKeyOk() (*string, bool)`

GetVictorOpsApiKeyOk returns a tuple with the VictorOpsApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictorOpsApiKey

`func (o *VictorOpsNotification) SetVictorOpsApiKey(v string)`

SetVictorOpsApiKey sets VictorOpsApiKey field to given value.

### HasVictorOpsApiKey

`func (o *VictorOpsNotification) HasVictorOpsApiKey() bool`

HasVictorOpsApiKey returns a boolean if a field has been set.

### GetVictorOpsRoutingKey

`func (o *VictorOpsNotification) GetVictorOpsRoutingKey() string`

GetVictorOpsRoutingKey returns the VictorOpsRoutingKey field if non-nil, zero value otherwise.

### GetVictorOpsRoutingKeyOk

`func (o *VictorOpsNotification) GetVictorOpsRoutingKeyOk() (*string, bool)`

GetVictorOpsRoutingKeyOk returns a tuple with the VictorOpsRoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictorOpsRoutingKey

`func (o *VictorOpsNotification) SetVictorOpsRoutingKey(v string)`

SetVictorOpsRoutingKey sets VictorOpsRoutingKey field to given value.

### HasVictorOpsRoutingKey

`func (o *VictorOpsNotification) HasVictorOpsRoutingKey() bool`

HasVictorOpsRoutingKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


