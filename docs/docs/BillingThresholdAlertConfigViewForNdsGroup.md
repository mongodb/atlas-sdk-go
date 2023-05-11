# BillingThresholdAlertConfigViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled this alert configuration for the specified project. | [optional] [default to false]
**EventTypeName** | [**BillingEventTypeViewAlertableWithThreshold**](BillingEventTypeViewAlertableWithThreshold.md) |  | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this alert configuration. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Matchers** | Pointer to **[]map[string]interface{}** | No matchers are available for these alert types. The list is always empty. | [optional] [readonly] 
**Notifications** | Pointer to [**[]NotificationViewForNdsGroup**](NotificationViewForNdsGroup.md) | List that contains the targets that MongoDB Cloud sends notifications. | [optional] 
**Threshold** | Pointer to [**GreaterThanRawThreshold**](GreaterThanRawThreshold.md) |  | [optional] 
**Updated** | Pointer to **time.Time** | Date and time when someone last updated this alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewBillingThresholdAlertConfigViewForNdsGroup

`func NewBillingThresholdAlertConfigViewForNdsGroup(eventTypeName BillingEventTypeViewAlertableWithThreshold, ) *BillingThresholdAlertConfigViewForNdsGroup`

NewBillingThresholdAlertConfigViewForNdsGroup instantiates a new BillingThresholdAlertConfigViewForNdsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingThresholdAlertConfigViewForNdsGroupWithDefaults

`func NewBillingThresholdAlertConfigViewForNdsGroupWithDefaults() *BillingThresholdAlertConfigViewForNdsGroup`

NewBillingThresholdAlertConfigViewForNdsGroupWithDefaults instantiates a new BillingThresholdAlertConfigViewForNdsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypeName

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetEventTypeName() BillingEventTypeViewAlertableWithThreshold`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*BillingEventTypeViewAlertableWithThreshold, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetEventTypeName(v BillingEventTypeViewAlertableWithThreshold)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMatchers

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetMatchers() []map[string]interface{}`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetMatchersOk() (*[]map[string]interface{}, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetMatchers(v []map[string]interface{})`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetNotifications

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetNotificationsOk() (*[]NotificationViewForNdsGroup, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetThreshold

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetThreshold() GreaterThanRawThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetThresholdOk() (*GreaterThanRawThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetThreshold(v GreaterThanRawThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdated

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BillingThresholdAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BillingThresholdAlertConfigViewForNdsGroup) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BillingThresholdAlertConfigViewForNdsGroup) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


