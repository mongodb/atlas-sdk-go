# GroupAlertsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled this alert configuration for the specified project. | [optional] [default to false]
**EventTypeName** | Pointer to **string** | Event type that triggers an alert. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this alert configuration. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Matchers** | Pointer to **[]map[string]interface{}** | No matchers are available for these alert types. The list is always empty. | [optional] [readonly] 
**Notifications** | Pointer to [**[]AlertsNotificationRootForGroup**](AlertsNotificationRootForGroup.md) | List that contains the targets that MongoDB Cloud sends notifications. | [optional] 
**Updated** | Pointer to **time.Time** | Date and time when someone last updated this alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 
**MetricThreshold** | Pointer to [**ServerlessMetricThreshold**](ServerlessMetricThreshold.md) |  | [optional] 
**Threshold** | Pointer to [**AlertsThresholdInteger**](AlertsThresholdInteger.md) |  | [optional] 

## Methods

### NewGroupAlertsConfig

`func NewGroupAlertsConfig() *GroupAlertsConfig`

NewGroupAlertsConfig instantiates a new GroupAlertsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupAlertsConfigWithDefaults

`func NewGroupAlertsConfigWithDefaults() *GroupAlertsConfig`

NewGroupAlertsConfigWithDefaults instantiates a new GroupAlertsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GroupAlertsConfig) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupAlertsConfig) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupAlertsConfig) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GroupAlertsConfig) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *GroupAlertsConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroupAlertsConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroupAlertsConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GroupAlertsConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypeName

`func (o *GroupAlertsConfig) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *GroupAlertsConfig) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *GroupAlertsConfig) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.

### HasEventTypeName

`func (o *GroupAlertsConfig) HasEventTypeName() bool`

HasEventTypeName returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupAlertsConfig) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupAlertsConfig) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupAlertsConfig) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupAlertsConfig) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *GroupAlertsConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupAlertsConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupAlertsConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupAlertsConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *GroupAlertsConfig) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupAlertsConfig) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupAlertsConfig) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupAlertsConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMatchers

`func (o *GroupAlertsConfig) GetMatchers() []map[string]interface{}`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *GroupAlertsConfig) GetMatchersOk() (*[]map[string]interface{}, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *GroupAlertsConfig) SetMatchers(v []map[string]interface{})`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *GroupAlertsConfig) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetNotifications

`func (o *GroupAlertsConfig) GetNotifications() []AlertsNotificationRootForGroup`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *GroupAlertsConfig) GetNotificationsOk() (*[]AlertsNotificationRootForGroup, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *GroupAlertsConfig) SetNotifications(v []AlertsNotificationRootForGroup)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *GroupAlertsConfig) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetUpdated

`func (o *GroupAlertsConfig) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GroupAlertsConfig) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GroupAlertsConfig) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GroupAlertsConfig) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetMetricThreshold

`func (o *GroupAlertsConfig) GetMetricThreshold() ServerlessMetricThreshold`

GetMetricThreshold returns the MetricThreshold field if non-nil, zero value otherwise.

### GetMetricThresholdOk

`func (o *GroupAlertsConfig) GetMetricThresholdOk() (*ServerlessMetricThreshold, bool)`

GetMetricThresholdOk returns a tuple with the MetricThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricThreshold

`func (o *GroupAlertsConfig) SetMetricThreshold(v ServerlessMetricThreshold)`

SetMetricThreshold sets MetricThreshold field to given value.

### HasMetricThreshold

`func (o *GroupAlertsConfig) HasMetricThreshold() bool`

HasMetricThreshold returns a boolean if a field has been set.

### GetThreshold

`func (o *GroupAlertsConfig) GetThreshold() AlertsThresholdInteger`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GroupAlertsConfig) GetThresholdOk() (*AlertsThresholdInteger, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GroupAlertsConfig) SetThreshold(v AlertsThresholdInteger)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GroupAlertsConfig) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


