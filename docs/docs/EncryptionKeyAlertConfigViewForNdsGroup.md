# EncryptionKeyAlertConfigViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled this alert configuration for the specified project. | [optional] [default to false]
**EventTypeName** | [**EncryptionKeyEventTypeViewAlertable**](EncryptionKeyEventTypeViewAlertable.md) |  | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this alert configuration. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Matchers** | Pointer to **[]map[string]interface{}** | No matchers are available for these alert types. The list is always empty. | [optional] [readonly] 
**Notifications** | Pointer to [**[]NotificationViewForNdsGroup**](NotificationViewForNdsGroup.md) | List that contains the targets that MongoDB Cloud sends notifications. | [optional] 
**Threshold** | Pointer to [**GreaterThanDaysThreshold**](GreaterThanDaysThreshold.md) |  | [optional] 
**Updated** | Pointer to **time.Time** | Date and time when someone last updated this alert configuration. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewEncryptionKeyAlertConfigViewForNdsGroup

`func NewEncryptionKeyAlertConfigViewForNdsGroup(eventTypeName EncryptionKeyEventTypeViewAlertable, ) *EncryptionKeyAlertConfigViewForNdsGroup`

NewEncryptionKeyAlertConfigViewForNdsGroup instantiates a new EncryptionKeyAlertConfigViewForNdsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionKeyAlertConfigViewForNdsGroupWithDefaults

`func NewEncryptionKeyAlertConfigViewForNdsGroupWithDefaults() *EncryptionKeyAlertConfigViewForNdsGroup`

NewEncryptionKeyAlertConfigViewForNdsGroupWithDefaults instantiates a new EncryptionKeyAlertConfigViewForNdsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypeName

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetEventTypeName() EncryptionKeyEventTypeViewAlertable`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*EncryptionKeyEventTypeViewAlertable, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetEventTypeName(v EncryptionKeyEventTypeViewAlertable)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMatchers

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetMatchers() []map[string]interface{}`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetMatchersOk() (*[]map[string]interface{}, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetMatchers(v []map[string]interface{})`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetNotifications

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetNotificationsOk() (*[]NotificationViewForNdsGroup, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetThreshold

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetThreshold() GreaterThanDaysThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetThresholdOk() (*GreaterThanDaysThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetThreshold(v GreaterThanDaysThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdated

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *EncryptionKeyAlertConfigViewForNdsGroup) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


