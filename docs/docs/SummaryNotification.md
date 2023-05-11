# SummaryNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**EmailAddress** | Pointer to **string** | Email address to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;EMAIL\&quot;&#x60;. You don’t need to set this value to send emails to individual or groups of MongoDB Cloud users including:  - specific MongoDB Cloud users (&#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60;) - MongoDB Cloud users with specific project roles (&#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60;) - MongoDB Cloud users with specific organization roles (&#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60;) - MongoDB Cloud teams (&#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;TEAM\&quot;&#x60;)  To send emails to one MongoDB Cloud user or grouping of users, set the **notifications.[n].emailEnabled** parameter. | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 

## Methods

### NewSummaryNotification

`func NewSummaryNotification(typeName string, ) *SummaryNotification`

NewSummaryNotification instantiates a new SummaryNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryNotificationWithDefaults

`func NewSummaryNotificationWithDefaults() *SummaryNotification`

NewSummaryNotificationWithDefaults instantiates a new SummaryNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMin

`func (o *SummaryNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *SummaryNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *SummaryNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *SummaryNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetEmailAddress

`func (o *SummaryNotification) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SummaryNotification) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SummaryNotification) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SummaryNotification) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetIntervalMin

`func (o *SummaryNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *SummaryNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *SummaryNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *SummaryNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetTypeName

`func (o *SummaryNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *SummaryNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *SummaryNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


