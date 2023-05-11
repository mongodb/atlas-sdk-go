# RawMetricAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedUntil** | **time.Time** | Date and time until which this alert has been acknowledged. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.  - To acknowledge this alert forever, set the parameter value to 100 years in the future.  - To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past. | 
**AcknowledgementComment** | Pointer to **string** | Comment that a MongoDB Cloud user submitted when acknowledging the alert. | [optional] 
**AcknowledgingUsername** | Pointer to **string** | MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert. | [optional] [readonly] 
**AlertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert. | [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster to which this alert applies. This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters. | [optional] [readonly] 
**Created** | **time.Time** | Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [readonly] 
**CurrentValue** | Pointer to [**RawMetricValue**](RawMetricValue.md) |  | [optional] 
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

### NewRawMetricAlert

`func NewRawMetricAlert(acknowledgedUntil time.Time, alertConfigId string, created time.Time, eventTypeName HostMetricEventTypeViewAlertable, id string, status string, updated time.Time, ) *RawMetricAlert`

NewRawMetricAlert instantiates a new RawMetricAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawMetricAlertWithDefaults

`func NewRawMetricAlertWithDefaults() *RawMetricAlert`

NewRawMetricAlertWithDefaults instantiates a new RawMetricAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedUntil

`func (o *RawMetricAlert) GetAcknowledgedUntil() time.Time`

GetAcknowledgedUntil returns the AcknowledgedUntil field if non-nil, zero value otherwise.

### GetAcknowledgedUntilOk

`func (o *RawMetricAlert) GetAcknowledgedUntilOk() (*time.Time, bool)`

GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedUntil

`func (o *RawMetricAlert) SetAcknowledgedUntil(v time.Time)`

SetAcknowledgedUntil sets AcknowledgedUntil field to given value.


### GetAcknowledgementComment

`func (o *RawMetricAlert) GetAcknowledgementComment() string`

GetAcknowledgementComment returns the AcknowledgementComment field if non-nil, zero value otherwise.

### GetAcknowledgementCommentOk

`func (o *RawMetricAlert) GetAcknowledgementCommentOk() (*string, bool)`

GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementComment

`func (o *RawMetricAlert) SetAcknowledgementComment(v string)`

SetAcknowledgementComment sets AcknowledgementComment field to given value.

### HasAcknowledgementComment

`func (o *RawMetricAlert) HasAcknowledgementComment() bool`

HasAcknowledgementComment returns a boolean if a field has been set.

### GetAcknowledgingUsername

`func (o *RawMetricAlert) GetAcknowledgingUsername() string`

GetAcknowledgingUsername returns the AcknowledgingUsername field if non-nil, zero value otherwise.

### GetAcknowledgingUsernameOk

`func (o *RawMetricAlert) GetAcknowledgingUsernameOk() (*string, bool)`

GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgingUsername

`func (o *RawMetricAlert) SetAcknowledgingUsername(v string)`

SetAcknowledgingUsername sets AcknowledgingUsername field to given value.

### HasAcknowledgingUsername

`func (o *RawMetricAlert) HasAcknowledgingUsername() bool`

HasAcknowledgingUsername returns a boolean if a field has been set.

### GetAlertConfigId

`func (o *RawMetricAlert) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *RawMetricAlert) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *RawMetricAlert) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.


### GetClusterName

`func (o *RawMetricAlert) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *RawMetricAlert) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *RawMetricAlert) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *RawMetricAlert) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreated

`func (o *RawMetricAlert) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RawMetricAlert) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RawMetricAlert) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCurrentValue

`func (o *RawMetricAlert) GetCurrentValue() RawMetricValue`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *RawMetricAlert) GetCurrentValueOk() (*RawMetricValue, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *RawMetricAlert) SetCurrentValue(v RawMetricValue)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *RawMetricAlert) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetEventTypeName

`func (o *RawMetricAlert) GetEventTypeName() HostMetricEventTypeViewAlertable`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *RawMetricAlert) GetEventTypeNameOk() (*HostMetricEventTypeViewAlertable, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *RawMetricAlert) SetEventTypeName(v HostMetricEventTypeViewAlertable)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *RawMetricAlert) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RawMetricAlert) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RawMetricAlert) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RawMetricAlert) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHostnameAndPort

`func (o *RawMetricAlert) GetHostnameAndPort() string`

GetHostnameAndPort returns the HostnameAndPort field if non-nil, zero value otherwise.

### GetHostnameAndPortOk

`func (o *RawMetricAlert) GetHostnameAndPortOk() (*string, bool)`

GetHostnameAndPortOk returns a tuple with the HostnameAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameAndPort

`func (o *RawMetricAlert) SetHostnameAndPort(v string)`

SetHostnameAndPort sets HostnameAndPort field to given value.

### HasHostnameAndPort

`func (o *RawMetricAlert) HasHostnameAndPort() bool`

HasHostnameAndPort returns a boolean if a field has been set.

### GetId

`func (o *RawMetricAlert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RawMetricAlert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RawMetricAlert) SetId(v string)`

SetId sets Id field to given value.


### GetLastNotified

`func (o *RawMetricAlert) GetLastNotified() time.Time`

GetLastNotified returns the LastNotified field if non-nil, zero value otherwise.

### GetLastNotifiedOk

`func (o *RawMetricAlert) GetLastNotifiedOk() (*time.Time, bool)`

GetLastNotifiedOk returns a tuple with the LastNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotified

`func (o *RawMetricAlert) SetLastNotified(v time.Time)`

SetLastNotified sets LastNotified field to given value.

### HasLastNotified

`func (o *RawMetricAlert) HasLastNotified() bool`

HasLastNotified returns a boolean if a field has been set.

### GetLinks

`func (o *RawMetricAlert) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RawMetricAlert) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RawMetricAlert) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RawMetricAlert) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetricName

`func (o *RawMetricAlert) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *RawMetricAlert) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *RawMetricAlert) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *RawMetricAlert) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetOrgId

`func (o *RawMetricAlert) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RawMetricAlert) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RawMetricAlert) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *RawMetricAlert) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetReplicaSetName

`func (o *RawMetricAlert) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *RawMetricAlert) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *RawMetricAlert) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *RawMetricAlert) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

### GetResolved

`func (o *RawMetricAlert) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *RawMetricAlert) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *RawMetricAlert) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *RawMetricAlert) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetStatus

`func (o *RawMetricAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RawMetricAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RawMetricAlert) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdated

`func (o *RawMetricAlert) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RawMetricAlert) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RawMetricAlert) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


