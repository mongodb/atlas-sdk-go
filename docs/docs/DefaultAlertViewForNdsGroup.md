# DefaultAlertViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedUntil** | **time.Time** | Date and time until which this alert has been acknowledged. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.  - To acknowledge this alert forever, set the parameter value to 100 years in the future.  - To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past. | 
**AcknowledgementComment** | Pointer to **string** | Comment that a MongoDB Cloud user submitted when acknowledging the alert. | [optional] 
**AcknowledgingUsername** | Pointer to **string** | MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert. | [optional] [readonly] 
**AlertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert. | [readonly] 
**Created** | **time.Time** | Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 
**EventTypeName** | **string** | Incident that triggered this alert. | [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert. | [optional] [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies this alert. | [readonly] 
**LastNotified** | Pointer to **time.Time** | Date and time that any notifications were last sent for this alert. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the organization that owns the project to which this alert applies. | [optional] [readonly] 
**Resolved** | Pointer to **time.Time** | Date and time that this alert changed to &#x60;\&quot;status\&quot; : \&quot;CLOSED\&quot;&#x60;. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. The resource returns this parameter once &#x60;\&quot;status\&quot; : \&quot;CLOSED\&quot;&#x60;. | [optional] [readonly] 
**Status** | **string** | State of this alert at the time you requested its details. | [readonly] 
**Updated** | **time.Time** | Date and time when someone last updated this alert. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 

## Methods

### NewDefaultAlertViewForNdsGroup

`func NewDefaultAlertViewForNdsGroup(acknowledgedUntil time.Time, alertConfigId string, created time.Time, eventTypeName string, id string, status string, updated time.Time, ) *DefaultAlertViewForNdsGroup`

NewDefaultAlertViewForNdsGroup instantiates a new DefaultAlertViewForNdsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultAlertViewForNdsGroupWithDefaults

`func NewDefaultAlertViewForNdsGroupWithDefaults() *DefaultAlertViewForNdsGroup`

NewDefaultAlertViewForNdsGroupWithDefaults instantiates a new DefaultAlertViewForNdsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedUntil

`func (o *DefaultAlertViewForNdsGroup) GetAcknowledgedUntil() time.Time`

GetAcknowledgedUntil returns the AcknowledgedUntil field if non-nil, zero value otherwise.

### GetAcknowledgedUntilOk

`func (o *DefaultAlertViewForNdsGroup) GetAcknowledgedUntilOk() (*time.Time, bool)`

GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedUntil

`func (o *DefaultAlertViewForNdsGroup) SetAcknowledgedUntil(v time.Time)`

SetAcknowledgedUntil sets AcknowledgedUntil field to given value.


### GetAcknowledgementComment

`func (o *DefaultAlertViewForNdsGroup) GetAcknowledgementComment() string`

GetAcknowledgementComment returns the AcknowledgementComment field if non-nil, zero value otherwise.

### GetAcknowledgementCommentOk

`func (o *DefaultAlertViewForNdsGroup) GetAcknowledgementCommentOk() (*string, bool)`

GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementComment

`func (o *DefaultAlertViewForNdsGroup) SetAcknowledgementComment(v string)`

SetAcknowledgementComment sets AcknowledgementComment field to given value.

### HasAcknowledgementComment

`func (o *DefaultAlertViewForNdsGroup) HasAcknowledgementComment() bool`

HasAcknowledgementComment returns a boolean if a field has been set.

### GetAcknowledgingUsername

`func (o *DefaultAlertViewForNdsGroup) GetAcknowledgingUsername() string`

GetAcknowledgingUsername returns the AcknowledgingUsername field if non-nil, zero value otherwise.

### GetAcknowledgingUsernameOk

`func (o *DefaultAlertViewForNdsGroup) GetAcknowledgingUsernameOk() (*string, bool)`

GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgingUsername

`func (o *DefaultAlertViewForNdsGroup) SetAcknowledgingUsername(v string)`

SetAcknowledgingUsername sets AcknowledgingUsername field to given value.

### HasAcknowledgingUsername

`func (o *DefaultAlertViewForNdsGroup) HasAcknowledgingUsername() bool`

HasAcknowledgingUsername returns a boolean if a field has been set.

### GetAlertConfigId

`func (o *DefaultAlertViewForNdsGroup) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *DefaultAlertViewForNdsGroup) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *DefaultAlertViewForNdsGroup) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.


### GetCreated

`func (o *DefaultAlertViewForNdsGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DefaultAlertViewForNdsGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DefaultAlertViewForNdsGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEventTypeName

`func (o *DefaultAlertViewForNdsGroup) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *DefaultAlertViewForNdsGroup) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *DefaultAlertViewForNdsGroup) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *DefaultAlertViewForNdsGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DefaultAlertViewForNdsGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DefaultAlertViewForNdsGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DefaultAlertViewForNdsGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *DefaultAlertViewForNdsGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefaultAlertViewForNdsGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefaultAlertViewForNdsGroup) SetId(v string)`

SetId sets Id field to given value.


### GetLastNotified

`func (o *DefaultAlertViewForNdsGroup) GetLastNotified() time.Time`

GetLastNotified returns the LastNotified field if non-nil, zero value otherwise.

### GetLastNotifiedOk

`func (o *DefaultAlertViewForNdsGroup) GetLastNotifiedOk() (*time.Time, bool)`

GetLastNotifiedOk returns a tuple with the LastNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotified

`func (o *DefaultAlertViewForNdsGroup) SetLastNotified(v time.Time)`

SetLastNotified sets LastNotified field to given value.

### HasLastNotified

`func (o *DefaultAlertViewForNdsGroup) HasLastNotified() bool`

HasLastNotified returns a boolean if a field has been set.

### GetLinks

`func (o *DefaultAlertViewForNdsGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DefaultAlertViewForNdsGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DefaultAlertViewForNdsGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DefaultAlertViewForNdsGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrgId

`func (o *DefaultAlertViewForNdsGroup) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DefaultAlertViewForNdsGroup) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DefaultAlertViewForNdsGroup) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DefaultAlertViewForNdsGroup) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetResolved

`func (o *DefaultAlertViewForNdsGroup) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *DefaultAlertViewForNdsGroup) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *DefaultAlertViewForNdsGroup) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *DefaultAlertViewForNdsGroup) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetStatus

`func (o *DefaultAlertViewForNdsGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DefaultAlertViewForNdsGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DefaultAlertViewForNdsGroup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdated

`func (o *DefaultAlertViewForNdsGroup) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DefaultAlertViewForNdsGroup) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DefaultAlertViewForNdsGroup) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


