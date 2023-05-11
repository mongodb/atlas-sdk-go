# PagerDutyNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**Region** | Pointer to **string** | PagerDuty region that indicates which API Uniform Resource Locator (URL) to use. | [optional] [default to "US"]
**ServiceKey** | Pointer to **string** | PagerDuty service key that MongoDB Cloud needs to send notifications via PagerDuty. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;PAGER_DUTY\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 

## Methods

### NewPagerDutyNotification

`func NewPagerDutyNotification(typeName string, ) *PagerDutyNotification`

NewPagerDutyNotification instantiates a new PagerDutyNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagerDutyNotificationWithDefaults

`func NewPagerDutyNotificationWithDefaults() *PagerDutyNotification`

NewPagerDutyNotificationWithDefaults instantiates a new PagerDutyNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMin

`func (o *PagerDutyNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *PagerDutyNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *PagerDutyNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *PagerDutyNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetIntervalMin

`func (o *PagerDutyNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *PagerDutyNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *PagerDutyNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *PagerDutyNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetRegion

`func (o *PagerDutyNotification) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PagerDutyNotification) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PagerDutyNotification) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PagerDutyNotification) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServiceKey

`func (o *PagerDutyNotification) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *PagerDutyNotification) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *PagerDutyNotification) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *PagerDutyNotification) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetTypeName

`func (o *PagerDutyNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *PagerDutyNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *PagerDutyNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


