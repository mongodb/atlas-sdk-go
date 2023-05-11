# TimeMetricAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedUntil** | **time.Time** | Date and time until which this alert has been acknowledged. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.  - To acknowledge this alert forever, set the parameter value to 100 years in the future.  - To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past. | 
**AcknowledgementComment** | Pointer to **string** | Comment that a MongoDB Cloud user submitted when acknowledging the alert. | [optional] 
**AcknowledgingUsername** | Pointer to **string** | MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert. | [optional] [readonly] 
**AlertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert. | [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster to which this alert applies. This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters. | [optional] [readonly] 
**Created** | **time.Time** | Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 
**CurrentValue** | Pointer to [**TimeMetricValue**](TimeMetricValue.md) |  | [optional] 
**EventTypeName** | [**HostMetricEventTypeViewAlertable**](HostMetricEventTypeViewAlertable.md) |  | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert. | [optional] [readonly] 
**HostnameAndPort** | Pointer to **string** | Hostname and port of the host to which this alert applies. The resource returns this parameter for alerts of events impacting hosts or replica sets. | [optional] [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies this alert. | [readonly] 
**LastNotified** | Pointer to **time.Time** | Date and time that any notifications were last sent for this alert. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MetricName** | Pointer to **string** | Name of the metric against which Atlas checks the configured &#x60;metricThreshold.threshold&#x60;.  To learn more about the available metrics, see &lt;a href&#x3D;\&quot;https://www.mongodb.com/docs/atlas/reference/alert-host-metrics/#std-label-measurement-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Host Metrics&lt;/a&gt;.  **NOTE**: If you set eventTypeName to OUTSIDE_SERVERLESS_METRIC_THRESHOLD, you can specify only metrics available for serverless. To learn more, see &lt;a href&#x3D;\&quot;https://dochub.mongodb.org/core/alert-config-serverless-measurements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Serverless Measurements&lt;/a&gt;. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the organization that owns the project to which this alert applies. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Name of the replica set to which this alert applies. The response returns this parameter for alerts of events impacting backups, hosts, or replica sets. | [optional] [readonly] 
**Resolved** | Pointer to **time.Time** | Date and time that this alert changed to &#x60;\&quot;status\&quot; : \&quot;CLOSED\&quot;&#x60;. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. The resource returns this parameter once &#x60;\&quot;status\&quot; : \&quot;CLOSED\&quot;&#x60;. | [optional] [readonly] 
**Status** | **string** | State of this alert at the time you requested its details. | [readonly] 
**Updated** | **time.Time** | Date and time when someone last updated this alert. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 

## Methods

### NewTimeMetricAlert

`func NewTimeMetricAlert(acknowledgedUntil time.Time, alertConfigId string, created time.Time, eventTypeName HostMetricEventTypeViewAlertable, id string, status string, updated time.Time, ) *TimeMetricAlert`

NewTimeMetricAlert instantiates a new TimeMetricAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeMetricAlertWithDefaults

`func NewTimeMetricAlertWithDefaults() *TimeMetricAlert`

NewTimeMetricAlertWithDefaults instantiates a new TimeMetricAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedUntil

`func (o *TimeMetricAlert) GetAcknowledgedUntil() time.Time`

GetAcknowledgedUntil returns the AcknowledgedUntil field if non-nil, zero value otherwise.

### GetAcknowledgedUntilOk

`func (o *TimeMetricAlert) GetAcknowledgedUntilOk() (*time.Time, bool)`

GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedUntil

`func (o *TimeMetricAlert) SetAcknowledgedUntil(v time.Time)`

SetAcknowledgedUntil sets AcknowledgedUntil field to given value.


### GetAcknowledgementComment

`func (o *TimeMetricAlert) GetAcknowledgementComment() string`

GetAcknowledgementComment returns the AcknowledgementComment field if non-nil, zero value otherwise.

### GetAcknowledgementCommentOk

`func (o *TimeMetricAlert) GetAcknowledgementCommentOk() (*string, bool)`

GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementComment

`func (o *TimeMetricAlert) SetAcknowledgementComment(v string)`

SetAcknowledgementComment sets AcknowledgementComment field to given value.

### HasAcknowledgementComment

`func (o *TimeMetricAlert) HasAcknowledgementComment() bool`

HasAcknowledgementComment returns a boolean if a field has been set.

### GetAcknowledgingUsername

`func (o *TimeMetricAlert) GetAcknowledgingUsername() string`

GetAcknowledgingUsername returns the AcknowledgingUsername field if non-nil, zero value otherwise.

### GetAcknowledgingUsernameOk

`func (o *TimeMetricAlert) GetAcknowledgingUsernameOk() (*string, bool)`

GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgingUsername

`func (o *TimeMetricAlert) SetAcknowledgingUsername(v string)`

SetAcknowledgingUsername sets AcknowledgingUsername field to given value.

### HasAcknowledgingUsername

`func (o *TimeMetricAlert) HasAcknowledgingUsername() bool`

HasAcknowledgingUsername returns a boolean if a field has been set.

### GetAlertConfigId

`func (o *TimeMetricAlert) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *TimeMetricAlert) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *TimeMetricAlert) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.


### GetClusterName

`func (o *TimeMetricAlert) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *TimeMetricAlert) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *TimeMetricAlert) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *TimeMetricAlert) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreated

`func (o *TimeMetricAlert) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TimeMetricAlert) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TimeMetricAlert) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCurrentValue

`func (o *TimeMetricAlert) GetCurrentValue() TimeMetricValue`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *TimeMetricAlert) GetCurrentValueOk() (*TimeMetricValue, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *TimeMetricAlert) SetCurrentValue(v TimeMetricValue)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *TimeMetricAlert) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetEventTypeName

`func (o *TimeMetricAlert) GetEventTypeName() HostMetricEventTypeViewAlertable`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *TimeMetricAlert) GetEventTypeNameOk() (*HostMetricEventTypeViewAlertable, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *TimeMetricAlert) SetEventTypeName(v HostMetricEventTypeViewAlertable)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *TimeMetricAlert) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TimeMetricAlert) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TimeMetricAlert) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *TimeMetricAlert) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHostnameAndPort

`func (o *TimeMetricAlert) GetHostnameAndPort() string`

GetHostnameAndPort returns the HostnameAndPort field if non-nil, zero value otherwise.

### GetHostnameAndPortOk

`func (o *TimeMetricAlert) GetHostnameAndPortOk() (*string, bool)`

GetHostnameAndPortOk returns a tuple with the HostnameAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameAndPort

`func (o *TimeMetricAlert) SetHostnameAndPort(v string)`

SetHostnameAndPort sets HostnameAndPort field to given value.

### HasHostnameAndPort

`func (o *TimeMetricAlert) HasHostnameAndPort() bool`

HasHostnameAndPort returns a boolean if a field has been set.

### GetId

`func (o *TimeMetricAlert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeMetricAlert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeMetricAlert) SetId(v string)`

SetId sets Id field to given value.


### GetLastNotified

`func (o *TimeMetricAlert) GetLastNotified() time.Time`

GetLastNotified returns the LastNotified field if non-nil, zero value otherwise.

### GetLastNotifiedOk

`func (o *TimeMetricAlert) GetLastNotifiedOk() (*time.Time, bool)`

GetLastNotifiedOk returns a tuple with the LastNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotified

`func (o *TimeMetricAlert) SetLastNotified(v time.Time)`

SetLastNotified sets LastNotified field to given value.

### HasLastNotified

`func (o *TimeMetricAlert) HasLastNotified() bool`

HasLastNotified returns a boolean if a field has been set.

### GetLinks

`func (o *TimeMetricAlert) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TimeMetricAlert) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TimeMetricAlert) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TimeMetricAlert) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetricName

`func (o *TimeMetricAlert) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *TimeMetricAlert) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *TimeMetricAlert) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *TimeMetricAlert) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetOrgId

`func (o *TimeMetricAlert) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *TimeMetricAlert) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *TimeMetricAlert) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *TimeMetricAlert) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetReplicaSetName

`func (o *TimeMetricAlert) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *TimeMetricAlert) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *TimeMetricAlert) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *TimeMetricAlert) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

### GetResolved

`func (o *TimeMetricAlert) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *TimeMetricAlert) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *TimeMetricAlert) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *TimeMetricAlert) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetStatus

`func (o *TimeMetricAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimeMetricAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimeMetricAlert) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdated

`func (o *TimeMetricAlert) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TimeMetricAlert) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TimeMetricAlert) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


