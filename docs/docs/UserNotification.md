# UserNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**EmailEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should send email notifications. The resource requires this parameter when one of the following values have been set:  - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60; | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**SmsEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should send text message notifications. The resource requires this parameter when one of the following values have been set:  - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60; | [optional] 
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 
**Username** | Pointer to **string** | MongoDB Cloud username of the person to whom MongoDB Cloud sends notifications. Specify only MongoDB Cloud users who belong to the project that owns the alert configuration. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60;. | [optional] 

## Methods

### NewUserNotification

`func NewUserNotification(typeName string, ) *UserNotification`

NewUserNotification instantiates a new UserNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNotificationWithDefaults

`func NewUserNotificationWithDefaults() *UserNotification`

NewUserNotificationWithDefaults instantiates a new UserNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMin

`func (o *UserNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *UserNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *UserNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *UserNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetEmailEnabled

`func (o *UserNotification) GetEmailEnabled() bool`

GetEmailEnabled returns the EmailEnabled field if non-nil, zero value otherwise.

### GetEmailEnabledOk

`func (o *UserNotification) GetEmailEnabledOk() (*bool, bool)`

GetEmailEnabledOk returns a tuple with the EmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEnabled

`func (o *UserNotification) SetEmailEnabled(v bool)`

SetEmailEnabled sets EmailEnabled field to given value.

### HasEmailEnabled

`func (o *UserNotification) HasEmailEnabled() bool`

HasEmailEnabled returns a boolean if a field has been set.

### GetIntervalMin

`func (o *UserNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *UserNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *UserNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *UserNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetSmsEnabled

`func (o *UserNotification) GetSmsEnabled() bool`

GetSmsEnabled returns the SmsEnabled field if non-nil, zero value otherwise.

### GetSmsEnabledOk

`func (o *UserNotification) GetSmsEnabledOk() (*bool, bool)`

GetSmsEnabledOk returns a tuple with the SmsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsEnabled

`func (o *UserNotification) SetSmsEnabled(v bool)`

SetSmsEnabled sets SmsEnabled field to given value.

### HasSmsEnabled

`func (o *UserNotification) HasSmsEnabled() bool`

HasSmsEnabled returns a boolean if a field has been set.

### GetTypeName

`func (o *UserNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *UserNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *UserNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetUsername

`func (o *UserNotification) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserNotification) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserNotification) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserNotification) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


