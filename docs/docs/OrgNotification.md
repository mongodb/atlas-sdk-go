# OrgNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMin** | Pointer to **int** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**EmailEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should send email notifications. The resource requires this parameter when one of the following values have been set:  - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60; | [optional] 
**IntervalMin** | Pointer to **int** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**Roles** | Pointer to **[]string** | List that contains the one or more [organization](https://dochub.mongodb.org/core/atlas-org-roles) or [project roles](https://dochub.mongodb.org/core/atlas-proj-roles) that receive the configured alert. The resource requires this parameter when &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60; or &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60;. If you include this parameter, MongoDB Cloud sends alerts only to users assigned the roles you specify in the array. If you omit this parameter, MongoDB Cloud sends alerts to users assigned any role. | [optional] 
**SmsEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should send text message notifications. The resource requires this parameter when one of the following values have been set:  - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;ORG\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;GROUP\&quot;&#x60; - &#x60;\&quot;notifications.[n].typeName\&quot; : \&quot;USER\&quot;&#x60; | [optional] 
**TypeName** | **string** | Human-readable label that displays the alert notification type. | 

## Methods

### NewOrgNotification

`func NewOrgNotification(typeName string, ) *OrgNotification`

NewOrgNotification instantiates a new OrgNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgNotificationWithDefaults

`func NewOrgNotificationWithDefaults() *OrgNotification`

NewOrgNotificationWithDefaults instantiates a new OrgNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMin

`func (o *OrgNotification) GetDelayMin() int`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *OrgNotification) GetDelayMinOk() (*int, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *OrgNotification) SetDelayMin(v int)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *OrgNotification) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetEmailEnabled

`func (o *OrgNotification) GetEmailEnabled() bool`

GetEmailEnabled returns the EmailEnabled field if non-nil, zero value otherwise.

### GetEmailEnabledOk

`func (o *OrgNotification) GetEmailEnabledOk() (*bool, bool)`

GetEmailEnabledOk returns a tuple with the EmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEnabled

`func (o *OrgNotification) SetEmailEnabled(v bool)`

SetEmailEnabled sets EmailEnabled field to given value.

### HasEmailEnabled

`func (o *OrgNotification) HasEmailEnabled() bool`

HasEmailEnabled returns a boolean if a field has been set.

### GetIntervalMin

`func (o *OrgNotification) GetIntervalMin() int`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *OrgNotification) GetIntervalMinOk() (*int, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *OrgNotification) SetIntervalMin(v int)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *OrgNotification) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetRoles

`func (o *OrgNotification) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrgNotification) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrgNotification) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrgNotification) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSmsEnabled

`func (o *OrgNotification) GetSmsEnabled() bool`

GetSmsEnabled returns the SmsEnabled field if non-nil, zero value otherwise.

### GetSmsEnabledOk

`func (o *OrgNotification) GetSmsEnabledOk() (*bool, bool)`

GetSmsEnabledOk returns a tuple with the SmsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsEnabled

`func (o *OrgNotification) SetSmsEnabled(v bool)`

SetSmsEnabled sets SmsEnabled field to given value.

### HasSmsEnabled

`func (o *OrgNotification) HasSmsEnabled() bool`

HasSmsEnabled returns a boolean if a field has been set.

### GetTypeName

`func (o *OrgNotification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *OrgNotification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *OrgNotification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


