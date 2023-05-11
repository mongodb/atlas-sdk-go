# NDSNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**Severity** | Pointer to **string** | Degree of seriousness given to this notification. | [optional] 
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 

## Methods

### NewNDSNotification

`func NewNDSNotification(typeName string, ) *NDSNotification`

NewNDSNotification instantiates a new NDSNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNDSNotificationWithDefaults

`func NewNDSNotificationWithDefaults() *NDSNotification`

NewNDSNotificationWithDefaults instantiates a new NDSNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMin

`func (o *NDSNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *NDSNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *NDSNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *NDSNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetIntervalMin

`func (o *NDSNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *NDSNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *NDSNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *NDSNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetSeverity

`func (o *NDSNotification) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NDSNotification) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NDSNotification) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NDSNotification) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTypeName

`func (o *NDSNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *NDSNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *NDSNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


