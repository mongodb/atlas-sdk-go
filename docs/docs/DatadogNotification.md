# DatadogNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatadogApiKey** | Pointer to **string** | Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog. You can find this API key in the Datadog dashboard. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;DATADOG\&quot;&#x60;.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | [optional] 
**DatadogRegion** | Pointer to **string** | Datadog region that indicates which API Uniform Resource Locator (URL) to use. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;DATADOG\&quot;&#x60;.  To learn more about Datadog&#39;s regions, see &lt;a href&#x3D;\&quot;https://docs.datadoghq.com/getting_started/site/\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;Datadog Sites&lt;/a&gt;. | [optional] [default to "US"]
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 

## Methods

### NewDatadogNotification

`func NewDatadogNotification(typeName string, ) *DatadogNotification`

NewDatadogNotification instantiates a new DatadogNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatadogNotificationWithDefaults

`func NewDatadogNotificationWithDefaults() *DatadogNotification`

NewDatadogNotificationWithDefaults instantiates a new DatadogNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatadogApiKey

`func (o *DatadogNotification) GetDatadogApiKey() string`

GetDatadogApiKey returns the DatadogApiKey field if non-nil, zero value otherwise.

### GetDatadogApiKeyOk

`func (o *DatadogNotification) GetDatadogApiKeyOk() (*string, bool)`

GetDatadogApiKeyOk returns a tuple with the DatadogApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadogApiKey

`func (o *DatadogNotification) SetDatadogApiKey(v string)`

SetDatadogApiKey sets DatadogApiKey field to given value.

### HasDatadogApiKey

`func (o *DatadogNotification) HasDatadogApiKey() bool`

HasDatadogApiKey returns a boolean if a field has been set.

### GetDatadogRegion

`func (o *DatadogNotification) GetDatadogRegion() string`

GetDatadogRegion returns the DatadogRegion field if non-nil, zero value otherwise.

### GetDatadogRegionOk

`func (o *DatadogNotification) GetDatadogRegionOk() (*string, bool)`

GetDatadogRegionOk returns a tuple with the DatadogRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadogRegion

`func (o *DatadogNotification) SetDatadogRegion(v string)`

SetDatadogRegion sets DatadogRegion field to given value.

### HasDatadogRegion

`func (o *DatadogNotification) HasDatadogRegion() bool`

HasDatadogRegion returns a boolean if a field has been set.

### GetDelayMin

`func (o *DatadogNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *DatadogNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *DatadogNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *DatadogNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetIntervalMin

`func (o *DatadogNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *DatadogNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *DatadogNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *DatadogNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetTypeName

`func (o *DatadogNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *DatadogNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *DatadogNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


