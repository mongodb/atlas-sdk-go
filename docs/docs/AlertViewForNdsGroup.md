# AlertViewForNdsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedUntil** | Pointer to **time.Time** | Date and time until which this alert has been acknowledged. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.  - To acknowledge this alert forever, set the parameter value to 100 years in the future.  - To unacknowledge a previously acknowledged alert, do not set this parameter value. | [optional] 
**AcknowledgementComment** | Pointer to **string** | Comment that a MongoDB Cloud user submitted when acknowledging the alert. | [optional] 
**AcknowledgingUsername** | Pointer to **string** | MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert. | [optional] [readonly] 
**AlertConfigId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**EventTypeName** | Pointer to **string** | Incident that triggered this alert. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this alert. | [optional] [readonly] 
**LastNotified** | Pointer to **time.Time** | Date and time that any notifications were last sent for this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the organization that owns the project to which this alert applies. | [optional] [readonly] 
**Resolved** | Pointer to **time.Time** | Date and time that this alert changed to &#x60;\&quot;status\&quot; : \&quot;CLOSED\&quot;&#x60;. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter once &#x60;\&quot;status\&quot; : \&quot;CLOSED\&quot;&#x60;. | [optional] [readonly] 
**Status** | Pointer to **string** | State of this alert at the time you requested its details. TRACKING indicates the alert condition exists but has not persisted for the minimum notification delay. OPEN indicates the alert condition currently exists. CLOSED indicates the alert condition has been resolved. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | Date and time when someone last updated this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster to which this alert applies. This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters. | [optional] [readonly] 
**HostnameAndPort** | Pointer to **string** | Hostname and port of the host to which this alert applies. The resource returns this parameter for alerts of events impacting hosts or replica sets. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Name of the replica set to which this alert applies. The response returns this parameter for alerts of events impacting backups, hosts, or replica sets. | [optional] [readonly] 
**CurrentValue** | Pointer to [**NumberMetricValue**](NumberMetricValue.md) |  | [optional] 
**MetricName** | Pointer to **string** | Name of the metric against which Atlas checks the configured &#x60;metricThreshold.threshold&#x60;.  To learn more about the available metrics, see &lt;a href&#x3D;\&quot;https://www.mongodb.com/docs/atlas/reference/alert-host-metrics/#std-label-measurement-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Host Metrics&lt;/a&gt;.  **NOTE**: If you set eventTypeName to OUTSIDE_SERVERLESS_METRIC_THRESHOLD, you can specify only metrics available for serverless. To learn more, see &lt;a href&#x3D;\&quot;https://dochub.mongodb.org/core/alert-config-serverless-measurements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Serverless Measurements&lt;/a&gt;. | [optional] [readonly] 
**NonRunningHostIds** | Pointer to **[]string** | List of unique 24-hexadecimal character strings that identify the replica set members that are not in PRIMARY nor SECONDARY state. | [optional] [readonly] 
**ParentClusterId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the parent cluster to which this alert applies. The parent cluster contains the sharded nodes. MongoDB Cloud returns this parameter only for alerts of events impacting sharded clusters. | [optional] [readonly] 
**InstanceName** | Pointer to **string** | The name of the Stream Processing Instance to which this alert applies. The resource returns this parameter for alerts of events impacting Stream Processing Instances. | [optional] [readonly] 
**ProcessorErrorMsg** | Pointer to **string** | The error message associated with the Stream Processor to which this alert applies. | [optional] [readonly] 
**ProcessorName** | Pointer to **string** | The name of the Stream Processor to which this alert applies. The resource returns this parameter for alerts of events impacting Stream Processors. | [optional] [readonly] 
**ProcessorState** | Pointer to **string** | The state of the Stream Processor to which this alert applies. The resource returns this parameter for alerts of events impacting Stream Processors. | [optional] [readonly] 

## Methods

### NewAlertViewForNdsGroup

`func NewAlertViewForNdsGroup() *AlertViewForNdsGroup`

NewAlertViewForNdsGroup instantiates a new AlertViewForNdsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertViewForNdsGroupWithDefaults

`func NewAlertViewForNdsGroupWithDefaults() *AlertViewForNdsGroup`

NewAlertViewForNdsGroupWithDefaults instantiates a new AlertViewForNdsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedUntil

`func (o *AlertViewForNdsGroup) GetAcknowledgedUntil() time.Time`

GetAcknowledgedUntil returns the AcknowledgedUntil field if non-nil, zero value otherwise.

### GetAcknowledgedUntilOk

`func (o *AlertViewForNdsGroup) GetAcknowledgedUntilOk() (*time.Time, bool)`

GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedUntil

`func (o *AlertViewForNdsGroup) SetAcknowledgedUntil(v time.Time)`

SetAcknowledgedUntil sets AcknowledgedUntil field to given value.

### HasAcknowledgedUntil

`func (o *AlertViewForNdsGroup) HasAcknowledgedUntil() bool`

HasAcknowledgedUntil returns a boolean if a field has been set.
### GetAcknowledgementComment

`func (o *AlertViewForNdsGroup) GetAcknowledgementComment() string`

GetAcknowledgementComment returns the AcknowledgementComment field if non-nil, zero value otherwise.

### GetAcknowledgementCommentOk

`func (o *AlertViewForNdsGroup) GetAcknowledgementCommentOk() (*string, bool)`

GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementComment

`func (o *AlertViewForNdsGroup) SetAcknowledgementComment(v string)`

SetAcknowledgementComment sets AcknowledgementComment field to given value.

### HasAcknowledgementComment

`func (o *AlertViewForNdsGroup) HasAcknowledgementComment() bool`

HasAcknowledgementComment returns a boolean if a field has been set.
### GetAcknowledgingUsername

`func (o *AlertViewForNdsGroup) GetAcknowledgingUsername() string`

GetAcknowledgingUsername returns the AcknowledgingUsername field if non-nil, zero value otherwise.

### GetAcknowledgingUsernameOk

`func (o *AlertViewForNdsGroup) GetAcknowledgingUsernameOk() (*string, bool)`

GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgingUsername

`func (o *AlertViewForNdsGroup) SetAcknowledgingUsername(v string)`

SetAcknowledgingUsername sets AcknowledgingUsername field to given value.

### HasAcknowledgingUsername

`func (o *AlertViewForNdsGroup) HasAcknowledgingUsername() bool`

HasAcknowledgingUsername returns a boolean if a field has been set.
### GetAlertConfigId

`func (o *AlertViewForNdsGroup) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *AlertViewForNdsGroup) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *AlertViewForNdsGroup) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.

### HasAlertConfigId

`func (o *AlertViewForNdsGroup) HasAlertConfigId() bool`

HasAlertConfigId returns a boolean if a field has been set.
### GetCreated

`func (o *AlertViewForNdsGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AlertViewForNdsGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AlertViewForNdsGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AlertViewForNdsGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.
### GetEventTypeName

`func (o *AlertViewForNdsGroup) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *AlertViewForNdsGroup) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *AlertViewForNdsGroup) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.

### HasEventTypeName

`func (o *AlertViewForNdsGroup) HasEventTypeName() bool`

HasEventTypeName returns a boolean if a field has been set.
### GetGroupId

`func (o *AlertViewForNdsGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AlertViewForNdsGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AlertViewForNdsGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AlertViewForNdsGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *AlertViewForNdsGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertViewForNdsGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertViewForNdsGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertViewForNdsGroup) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLastNotified

`func (o *AlertViewForNdsGroup) GetLastNotified() time.Time`

GetLastNotified returns the LastNotified field if non-nil, zero value otherwise.

### GetLastNotifiedOk

`func (o *AlertViewForNdsGroup) GetLastNotifiedOk() (*time.Time, bool)`

GetLastNotifiedOk returns a tuple with the LastNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotified

`func (o *AlertViewForNdsGroup) SetLastNotified(v time.Time)`

SetLastNotified sets LastNotified field to given value.

### HasLastNotified

`func (o *AlertViewForNdsGroup) HasLastNotified() bool`

HasLastNotified returns a boolean if a field has been set.
### GetLinks

`func (o *AlertViewForNdsGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlertViewForNdsGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlertViewForNdsGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AlertViewForNdsGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOrgId

`func (o *AlertViewForNdsGroup) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AlertViewForNdsGroup) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AlertViewForNdsGroup) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AlertViewForNdsGroup) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.
### GetResolved

`func (o *AlertViewForNdsGroup) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *AlertViewForNdsGroup) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *AlertViewForNdsGroup) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *AlertViewForNdsGroup) HasResolved() bool`

HasResolved returns a boolean if a field has been set.
### GetStatus

`func (o *AlertViewForNdsGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertViewForNdsGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertViewForNdsGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertViewForNdsGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetUpdated

`func (o *AlertViewForNdsGroup) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AlertViewForNdsGroup) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AlertViewForNdsGroup) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AlertViewForNdsGroup) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.
### GetClusterName

`func (o *AlertViewForNdsGroup) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *AlertViewForNdsGroup) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *AlertViewForNdsGroup) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *AlertViewForNdsGroup) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetHostnameAndPort

`func (o *AlertViewForNdsGroup) GetHostnameAndPort() string`

GetHostnameAndPort returns the HostnameAndPort field if non-nil, zero value otherwise.

### GetHostnameAndPortOk

`func (o *AlertViewForNdsGroup) GetHostnameAndPortOk() (*string, bool)`

GetHostnameAndPortOk returns a tuple with the HostnameAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameAndPort

`func (o *AlertViewForNdsGroup) SetHostnameAndPort(v string)`

SetHostnameAndPort sets HostnameAndPort field to given value.

### HasHostnameAndPort

`func (o *AlertViewForNdsGroup) HasHostnameAndPort() bool`

HasHostnameAndPort returns a boolean if a field has been set.
### GetReplicaSetName

`func (o *AlertViewForNdsGroup) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *AlertViewForNdsGroup) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *AlertViewForNdsGroup) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *AlertViewForNdsGroup) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.
### GetCurrentValue

`func (o *AlertViewForNdsGroup) GetCurrentValue() NumberMetricValue`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *AlertViewForNdsGroup) GetCurrentValueOk() (*NumberMetricValue, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *AlertViewForNdsGroup) SetCurrentValue(v NumberMetricValue)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *AlertViewForNdsGroup) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.
### GetMetricName

`func (o *AlertViewForNdsGroup) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *AlertViewForNdsGroup) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *AlertViewForNdsGroup) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *AlertViewForNdsGroup) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.
### GetNonRunningHostIds

`func (o *AlertViewForNdsGroup) GetNonRunningHostIds() []string`

GetNonRunningHostIds returns the NonRunningHostIds field if non-nil, zero value otherwise.

### GetNonRunningHostIdsOk

`func (o *AlertViewForNdsGroup) GetNonRunningHostIdsOk() (*[]string, bool)`

GetNonRunningHostIdsOk returns a tuple with the NonRunningHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonRunningHostIds

`func (o *AlertViewForNdsGroup) SetNonRunningHostIds(v []string)`

SetNonRunningHostIds sets NonRunningHostIds field to given value.

### HasNonRunningHostIds

`func (o *AlertViewForNdsGroup) HasNonRunningHostIds() bool`

HasNonRunningHostIds returns a boolean if a field has been set.
### GetParentClusterId

`func (o *AlertViewForNdsGroup) GetParentClusterId() string`

GetParentClusterId returns the ParentClusterId field if non-nil, zero value otherwise.

### GetParentClusterIdOk

`func (o *AlertViewForNdsGroup) GetParentClusterIdOk() (*string, bool)`

GetParentClusterIdOk returns a tuple with the ParentClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentClusterId

`func (o *AlertViewForNdsGroup) SetParentClusterId(v string)`

SetParentClusterId sets ParentClusterId field to given value.

### HasParentClusterId

`func (o *AlertViewForNdsGroup) HasParentClusterId() bool`

HasParentClusterId returns a boolean if a field has been set.
### GetInstanceName

`func (o *AlertViewForNdsGroup) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *AlertViewForNdsGroup) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *AlertViewForNdsGroup) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *AlertViewForNdsGroup) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.
### GetProcessorErrorMsg

`func (o *AlertViewForNdsGroup) GetProcessorErrorMsg() string`

GetProcessorErrorMsg returns the ProcessorErrorMsg field if non-nil, zero value otherwise.

### GetProcessorErrorMsgOk

`func (o *AlertViewForNdsGroup) GetProcessorErrorMsgOk() (*string, bool)`

GetProcessorErrorMsgOk returns a tuple with the ProcessorErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorErrorMsg

`func (o *AlertViewForNdsGroup) SetProcessorErrorMsg(v string)`

SetProcessorErrorMsg sets ProcessorErrorMsg field to given value.

### HasProcessorErrorMsg

`func (o *AlertViewForNdsGroup) HasProcessorErrorMsg() bool`

HasProcessorErrorMsg returns a boolean if a field has been set.
### GetProcessorName

`func (o *AlertViewForNdsGroup) GetProcessorName() string`

GetProcessorName returns the ProcessorName field if non-nil, zero value otherwise.

### GetProcessorNameOk

`func (o *AlertViewForNdsGroup) GetProcessorNameOk() (*string, bool)`

GetProcessorNameOk returns a tuple with the ProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorName

`func (o *AlertViewForNdsGroup) SetProcessorName(v string)`

SetProcessorName sets ProcessorName field to given value.

### HasProcessorName

`func (o *AlertViewForNdsGroup) HasProcessorName() bool`

HasProcessorName returns a boolean if a field has been set.
### GetProcessorState

`func (o *AlertViewForNdsGroup) GetProcessorState() string`

GetProcessorState returns the ProcessorState field if non-nil, zero value otherwise.

### GetProcessorStateOk

`func (o *AlertViewForNdsGroup) GetProcessorStateOk() (*string, bool)`

GetProcessorStateOk returns a tuple with the ProcessorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorState

`func (o *AlertViewForNdsGroup) SetProcessorState(v string)`

SetProcessorState sets ProcessorState field to given value.

### HasProcessorState

`func (o *AlertViewForNdsGroup) HasProcessorState() bool`

HasProcessorState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


