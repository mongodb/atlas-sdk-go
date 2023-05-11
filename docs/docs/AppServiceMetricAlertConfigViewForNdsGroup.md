# AppServiceMetricAlertConfigViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled this alert configuration for the specified project. | [optional] [default to false]
**EventTypeName** | [**AppServiceEventTypeViewAlertableWithThreshold**](AppServiceEventTypeViewAlertableWithThreshold.md) |  | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this alert configuration. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Matchers** | Pointer to [**[]AppServiceMetricMatcher**](AppServiceMetricMatcher.md) | List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster. | [optional] 
**MetricThreshold** | Pointer to [**AppServiceMetricThreshold**](AppServiceMetricThreshold.md) |  | [optional] 
**Notifications** | Pointer to [**[]NotificationViewForNdsGroup**](NotificationViewForNdsGroup.md) | List that contains the targets that MongoDB Cloud sends notifications. | [optional] 
**Updated** | Pointer to **time.Time** | Date and time when someone last updated this alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewAppServiceMetricAlertConfigViewForNdsGroup

`func NewAppServiceMetricAlertConfigViewForNdsGroup(eventTypeName AppServiceEventTypeViewAlertableWithThreshold, ) *AppServiceMetricAlertConfigViewForNdsGroup`

NewAppServiceMetricAlertConfigViewForNdsGroup instantiates a new AppServiceMetricAlertConfigViewForNdsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceMetricAlertConfigViewForNdsGroupWithDefaults

`func NewAppServiceMetricAlertConfigViewForNdsGroupWithDefaults() *AppServiceMetricAlertConfigViewForNdsGroup`

NewAppServiceMetricAlertConfigViewForNdsGroupWithDefaults instantiates a new AppServiceMetricAlertConfigViewForNdsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypeName

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetEventTypeName() AppServiceEventTypeViewAlertableWithThreshold`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*AppServiceEventTypeViewAlertableWithThreshold, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetEventTypeName(v AppServiceEventTypeViewAlertableWithThreshold)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMatchers

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetMatchers() []AppServiceMetricMatcher`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetMatchersOk() (*[]AppServiceMetricMatcher, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetMatchers(v []AppServiceMetricMatcher)`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetMetricThreshold

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetMetricThreshold() AppServiceMetricThreshold`

GetMetricThreshold returns the MetricThreshold field if non-nil, zero value otherwise.

### GetMetricThresholdOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetMetricThresholdOk() (*AppServiceMetricThreshold, bool)`

GetMetricThresholdOk returns a tuple with the MetricThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricThreshold

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetMetricThreshold(v AppServiceMetricThreshold)`

SetMetricThreshold sets MetricThreshold field to given value.

### HasMetricThreshold

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasMetricThreshold() bool`

HasMetricThreshold returns a boolean if a field has been set.

### GetNotifications

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetNotificationsOk() (*[]NotificationViewForNdsGroup, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetUpdated

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AppServiceMetricAlertConfigViewForNdsGroup) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


