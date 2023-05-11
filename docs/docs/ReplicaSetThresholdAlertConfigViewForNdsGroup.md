# ReplicaSetThresholdAlertConfigViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled this alert configuration for the specified project. | [optional] [default to false]
**EventTypeName** | [**ReplicaSetEventTypeViewForNdsGroupAlertableWithThreshold**](ReplicaSetEventTypeViewForNdsGroupAlertableWithThreshold.md) |  | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this alert configuration. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Matchers** | Pointer to [**[]ReplicaSetMatcher**](ReplicaSetMatcher.md) | List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster. | [optional] 
**Notifications** | Pointer to [**[]NotificationViewForNdsGroup**](NotificationViewForNdsGroup.md) | List that contains the targets that MongoDB Cloud sends notifications. | [optional] 
**Threshold** | Pointer to [**GreaterThanRawThreshold**](GreaterThanRawThreshold.md) |  | [optional] 
**Updated** | Pointer to **time.Time** | Date and time when someone last updated this alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewReplicaSetThresholdAlertConfigViewForNdsGroup

`func NewReplicaSetThresholdAlertConfigViewForNdsGroup(eventTypeName ReplicaSetEventTypeViewForNdsGroupAlertableWithThreshold, ) *ReplicaSetThresholdAlertConfigViewForNdsGroup`

NewReplicaSetThresholdAlertConfigViewForNdsGroup instantiates a new ReplicaSetThresholdAlertConfigViewForNdsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaSetThresholdAlertConfigViewForNdsGroupWithDefaults

`func NewReplicaSetThresholdAlertConfigViewForNdsGroupWithDefaults() *ReplicaSetThresholdAlertConfigViewForNdsGroup`

NewReplicaSetThresholdAlertConfigViewForNdsGroupWithDefaults instantiates a new ReplicaSetThresholdAlertConfigViewForNdsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypeName

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetEventTypeName() ReplicaSetEventTypeViewForNdsGroupAlertableWithThreshold`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*ReplicaSetEventTypeViewForNdsGroupAlertableWithThreshold, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetEventTypeName(v ReplicaSetEventTypeViewForNdsGroupAlertableWithThreshold)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMatchers

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetMatchers() []ReplicaSetMatcher`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetMatchersOk() (*[]ReplicaSetMatcher, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetMatchers(v []ReplicaSetMatcher)`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetNotifications

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetNotificationsOk() (*[]NotificationViewForNdsGroup, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetThreshold

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetThreshold() GreaterThanRawThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetThresholdOk() (*GreaterThanRawThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetThreshold(v GreaterThanRawThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdated

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ReplicaSetThresholdAlertConfigViewForNdsGroup) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


