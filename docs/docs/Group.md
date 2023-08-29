# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAgentCount** | Pointer to **int** |  | [optional] 
**Chttl** | Pointer to **int64** |  | [optional] 
**Cre** | Pointer to **time.Time** |  | [optional] 
**DefaultTimeZoneDisplay** | Pointer to **string** |  | [optional] 
**DefaultTimeZoneDisplayShort** | Pointer to **string** |  | [optional] 
**DefaultTimeZoneId** | Pointer to **string** |  | [optional] 
**DisableDbstats** | Pointer to **bool** |  | [optional] 
**EnableAllHostLogs** | Pointer to **bool** |  | [optional] 
**EnableAllHostProfilers** | Pointer to **bool** |  | [optional] 
**EnableCurrentIpWarning** | Pointer to **bool** |  | [optional] 
**Experiments** | Pointer to [**[]ExperimentRegistration**](ExperimentRegistration.md) |  | [optional] 
**FeatureFlags** | Pointer to [**[]ToggleableFeatureFlag**](ToggleableFeatureFlag.md) |  | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**HasActiveBackups** | Pointer to **bool** |  | [optional] 
**HasActiveUI** | Pointer to **bool** |  | [optional] 
**HasAddedHosts** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Integrations** | Pointer to **map[string][]map[string]interface{}** |  | [optional] 
**IsIntercomEligible** | Pointer to **bool** |  | [optional] 
**IsProjectOverviewEnabled** | Pointer to **bool** |  | [optional] 
**LastClusterActiveSampleTime** | Pointer to **time.Time** |  | [optional] 
**Lmav** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewRelicInsightsAccountId** | Pointer to **string** |  | [optional] 
**NewRelicInsightsReadToken** | Pointer to **string** |  | [optional] 
**NewRelicInsightsWriteToken** | Pointer to **string** |  | [optional] 
**NewRelicLicenseKey** | Pointer to **string** |  | [optional] 
**PreferredHostnames** | Pointer to [**[]PreferredHostname**](PreferredHostname.md) |  | [optional] 
**Prom** | Pointer to [**PrometheusConfig**](PrometheusConfig.md) |  | [optional] 
**Status** | Pointer to [**GroupStatus**](GroupStatus.md) |  | [optional] 
**StorageConfig** | Pointer to [**GroupStorageConfig**](GroupStorageConfig.md) |  | [optional] 
**SummaryStatistics** | Pointer to [**GroupSummaryStatistics**](GroupSummaryStatistics.md) |  | [optional] 
**SuppressMongosAutoDiscovery** | Pointer to **bool** |  | [optional] 
**Teams** | Pointer to [**[]TeamRoleAssignment**](TeamRoleAssignment.md) |  | [optional] 
**UseCNRegionsOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAgentCount

`func (o *Group) GetActiveAgentCount() int`

GetActiveAgentCount returns the ActiveAgentCount field if non-nil, zero value otherwise.

### GetActiveAgentCountOk

`func (o *Group) GetActiveAgentCountOk() (*int, bool)`

GetActiveAgentCountOk returns a tuple with the ActiveAgentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAgentCount

`func (o *Group) SetActiveAgentCount(v int)`

SetActiveAgentCount sets ActiveAgentCount field to given value.

### HasActiveAgentCount

`func (o *Group) HasActiveAgentCount() bool`

HasActiveAgentCount returns a boolean if a field has been set.
### GetChttl

`func (o *Group) GetChttl() int64`

GetChttl returns the Chttl field if non-nil, zero value otherwise.

### GetChttlOk

`func (o *Group) GetChttlOk() (*int64, bool)`

GetChttlOk returns a tuple with the Chttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChttl

`func (o *Group) SetChttl(v int64)`

SetChttl sets Chttl field to given value.

### HasChttl

`func (o *Group) HasChttl() bool`

HasChttl returns a boolean if a field has been set.
### GetCre

`func (o *Group) GetCre() time.Time`

GetCre returns the Cre field if non-nil, zero value otherwise.

### GetCreOk

`func (o *Group) GetCreOk() (*time.Time, bool)`

GetCreOk returns a tuple with the Cre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCre

`func (o *Group) SetCre(v time.Time)`

SetCre sets Cre field to given value.

### HasCre

`func (o *Group) HasCre() bool`

HasCre returns a boolean if a field has been set.
### GetDefaultTimeZoneDisplay

`func (o *Group) GetDefaultTimeZoneDisplay() string`

GetDefaultTimeZoneDisplay returns the DefaultTimeZoneDisplay field if non-nil, zero value otherwise.

### GetDefaultTimeZoneDisplayOk

`func (o *Group) GetDefaultTimeZoneDisplayOk() (*string, bool)`

GetDefaultTimeZoneDisplayOk returns a tuple with the DefaultTimeZoneDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZoneDisplay

`func (o *Group) SetDefaultTimeZoneDisplay(v string)`

SetDefaultTimeZoneDisplay sets DefaultTimeZoneDisplay field to given value.

### HasDefaultTimeZoneDisplay

`func (o *Group) HasDefaultTimeZoneDisplay() bool`

HasDefaultTimeZoneDisplay returns a boolean if a field has been set.
### GetDefaultTimeZoneDisplayShort

`func (o *Group) GetDefaultTimeZoneDisplayShort() string`

GetDefaultTimeZoneDisplayShort returns the DefaultTimeZoneDisplayShort field if non-nil, zero value otherwise.

### GetDefaultTimeZoneDisplayShortOk

`func (o *Group) GetDefaultTimeZoneDisplayShortOk() (*string, bool)`

GetDefaultTimeZoneDisplayShortOk returns a tuple with the DefaultTimeZoneDisplayShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZoneDisplayShort

`func (o *Group) SetDefaultTimeZoneDisplayShort(v string)`

SetDefaultTimeZoneDisplayShort sets DefaultTimeZoneDisplayShort field to given value.

### HasDefaultTimeZoneDisplayShort

`func (o *Group) HasDefaultTimeZoneDisplayShort() bool`

HasDefaultTimeZoneDisplayShort returns a boolean if a field has been set.
### GetDefaultTimeZoneId

`func (o *Group) GetDefaultTimeZoneId() string`

GetDefaultTimeZoneId returns the DefaultTimeZoneId field if non-nil, zero value otherwise.

### GetDefaultTimeZoneIdOk

`func (o *Group) GetDefaultTimeZoneIdOk() (*string, bool)`

GetDefaultTimeZoneIdOk returns a tuple with the DefaultTimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZoneId

`func (o *Group) SetDefaultTimeZoneId(v string)`

SetDefaultTimeZoneId sets DefaultTimeZoneId field to given value.

### HasDefaultTimeZoneId

`func (o *Group) HasDefaultTimeZoneId() bool`

HasDefaultTimeZoneId returns a boolean if a field has been set.
### GetDisableDbstats

`func (o *Group) GetDisableDbstats() bool`

GetDisableDbstats returns the DisableDbstats field if non-nil, zero value otherwise.

### GetDisableDbstatsOk

`func (o *Group) GetDisableDbstatsOk() (*bool, bool)`

GetDisableDbstatsOk returns a tuple with the DisableDbstats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDbstats

`func (o *Group) SetDisableDbstats(v bool)`

SetDisableDbstats sets DisableDbstats field to given value.

### HasDisableDbstats

`func (o *Group) HasDisableDbstats() bool`

HasDisableDbstats returns a boolean if a field has been set.
### GetEnableAllHostLogs

`func (o *Group) GetEnableAllHostLogs() bool`

GetEnableAllHostLogs returns the EnableAllHostLogs field if non-nil, zero value otherwise.

### GetEnableAllHostLogsOk

`func (o *Group) GetEnableAllHostLogsOk() (*bool, bool)`

GetEnableAllHostLogsOk returns a tuple with the EnableAllHostLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllHostLogs

`func (o *Group) SetEnableAllHostLogs(v bool)`

SetEnableAllHostLogs sets EnableAllHostLogs field to given value.

### HasEnableAllHostLogs

`func (o *Group) HasEnableAllHostLogs() bool`

HasEnableAllHostLogs returns a boolean if a field has been set.
### GetEnableAllHostProfilers

`func (o *Group) GetEnableAllHostProfilers() bool`

GetEnableAllHostProfilers returns the EnableAllHostProfilers field if non-nil, zero value otherwise.

### GetEnableAllHostProfilersOk

`func (o *Group) GetEnableAllHostProfilersOk() (*bool, bool)`

GetEnableAllHostProfilersOk returns a tuple with the EnableAllHostProfilers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllHostProfilers

`func (o *Group) SetEnableAllHostProfilers(v bool)`

SetEnableAllHostProfilers sets EnableAllHostProfilers field to given value.

### HasEnableAllHostProfilers

`func (o *Group) HasEnableAllHostProfilers() bool`

HasEnableAllHostProfilers returns a boolean if a field has been set.
### GetEnableCurrentIpWarning

`func (o *Group) GetEnableCurrentIpWarning() bool`

GetEnableCurrentIpWarning returns the EnableCurrentIpWarning field if non-nil, zero value otherwise.

### GetEnableCurrentIpWarningOk

`func (o *Group) GetEnableCurrentIpWarningOk() (*bool, bool)`

GetEnableCurrentIpWarningOk returns a tuple with the EnableCurrentIpWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCurrentIpWarning

`func (o *Group) SetEnableCurrentIpWarning(v bool)`

SetEnableCurrentIpWarning sets EnableCurrentIpWarning field to given value.

### HasEnableCurrentIpWarning

`func (o *Group) HasEnableCurrentIpWarning() bool`

HasEnableCurrentIpWarning returns a boolean if a field has been set.
### GetExperiments

`func (o *Group) GetExperiments() []ExperimentRegistration`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *Group) GetExperimentsOk() (*[]ExperimentRegistration, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *Group) SetExperiments(v []ExperimentRegistration)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *Group) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.
### GetFeatureFlags

`func (o *Group) GetFeatureFlags() []ToggleableFeatureFlag`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *Group) GetFeatureFlagsOk() (*[]ToggleableFeatureFlag, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *Group) SetFeatureFlags(v []ToggleableFeatureFlag)`

SetFeatureFlags sets FeatureFlags field to given value.

### HasFeatureFlags

`func (o *Group) HasFeatureFlags() bool`

HasFeatureFlags returns a boolean if a field has been set.
### GetGroupType

`func (o *Group) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *Group) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *Group) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *Group) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.
### GetHasActiveBackups

`func (o *Group) GetHasActiveBackups() bool`

GetHasActiveBackups returns the HasActiveBackups field if non-nil, zero value otherwise.

### GetHasActiveBackupsOk

`func (o *Group) GetHasActiveBackupsOk() (*bool, bool)`

GetHasActiveBackupsOk returns a tuple with the HasActiveBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveBackups

`func (o *Group) SetHasActiveBackups(v bool)`

SetHasActiveBackups sets HasActiveBackups field to given value.

### HasHasActiveBackups

`func (o *Group) HasHasActiveBackups() bool`

HasHasActiveBackups returns a boolean if a field has been set.
### GetHasActiveUI

`func (o *Group) GetHasActiveUI() bool`

GetHasActiveUI returns the HasActiveUI field if non-nil, zero value otherwise.

### GetHasActiveUIOk

`func (o *Group) GetHasActiveUIOk() (*bool, bool)`

GetHasActiveUIOk returns a tuple with the HasActiveUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveUI

`func (o *Group) SetHasActiveUI(v bool)`

SetHasActiveUI sets HasActiveUI field to given value.

### HasHasActiveUI

`func (o *Group) HasHasActiveUI() bool`

HasHasActiveUI returns a boolean if a field has been set.
### GetHasAddedHosts

`func (o *Group) GetHasAddedHosts() bool`

GetHasAddedHosts returns the HasAddedHosts field if non-nil, zero value otherwise.

### GetHasAddedHostsOk

`func (o *Group) GetHasAddedHostsOk() (*bool, bool)`

GetHasAddedHostsOk returns a tuple with the HasAddedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAddedHosts

`func (o *Group) SetHasAddedHosts(v bool)`

SetHasAddedHosts sets HasAddedHosts field to given value.

### HasHasAddedHosts

`func (o *Group) HasHasAddedHosts() bool`

HasHasAddedHosts returns a boolean if a field has been set.
### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.
### GetIntegrations

`func (o *Group) GetIntegrations() map[string][]map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *Group) GetIntegrationsOk() (*map[string][]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *Group) SetIntegrations(v map[string][]map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *Group) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.
### GetIsIntercomEligible

`func (o *Group) GetIsIntercomEligible() bool`

GetIsIntercomEligible returns the IsIntercomEligible field if non-nil, zero value otherwise.

### GetIsIntercomEligibleOk

`func (o *Group) GetIsIntercomEligibleOk() (*bool, bool)`

GetIsIntercomEligibleOk returns a tuple with the IsIntercomEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIntercomEligible

`func (o *Group) SetIsIntercomEligible(v bool)`

SetIsIntercomEligible sets IsIntercomEligible field to given value.

### HasIsIntercomEligible

`func (o *Group) HasIsIntercomEligible() bool`

HasIsIntercomEligible returns a boolean if a field has been set.
### GetIsProjectOverviewEnabled

`func (o *Group) GetIsProjectOverviewEnabled() bool`

GetIsProjectOverviewEnabled returns the IsProjectOverviewEnabled field if non-nil, zero value otherwise.

### GetIsProjectOverviewEnabledOk

`func (o *Group) GetIsProjectOverviewEnabledOk() (*bool, bool)`

GetIsProjectOverviewEnabledOk returns a tuple with the IsProjectOverviewEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProjectOverviewEnabled

`func (o *Group) SetIsProjectOverviewEnabled(v bool)`

SetIsProjectOverviewEnabled sets IsProjectOverviewEnabled field to given value.

### HasIsProjectOverviewEnabled

`func (o *Group) HasIsProjectOverviewEnabled() bool`

HasIsProjectOverviewEnabled returns a boolean if a field has been set.
### GetLastClusterActiveSampleTime

`func (o *Group) GetLastClusterActiveSampleTime() time.Time`

GetLastClusterActiveSampleTime returns the LastClusterActiveSampleTime field if non-nil, zero value otherwise.

### GetLastClusterActiveSampleTimeOk

`func (o *Group) GetLastClusterActiveSampleTimeOk() (*time.Time, bool)`

GetLastClusterActiveSampleTimeOk returns a tuple with the LastClusterActiveSampleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastClusterActiveSampleTime

`func (o *Group) SetLastClusterActiveSampleTime(v time.Time)`

SetLastClusterActiveSampleTime sets LastClusterActiveSampleTime field to given value.

### HasLastClusterActiveSampleTime

`func (o *Group) HasLastClusterActiveSampleTime() bool`

HasLastClusterActiveSampleTime returns a boolean if a field has been set.
### GetLmav

`func (o *Group) GetLmav() string`

GetLmav returns the Lmav field if non-nil, zero value otherwise.

### GetLmavOk

`func (o *Group) GetLmavOk() (*string, bool)`

GetLmavOk returns a tuple with the Lmav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmav

`func (o *Group) SetLmav(v string)`

SetLmav sets Lmav field to given value.

### HasLmav

`func (o *Group) HasLmav() bool`

HasLmav returns a boolean if a field has been set.
### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.
### GetNewRelicInsightsAccountId

`func (o *Group) GetNewRelicInsightsAccountId() string`

GetNewRelicInsightsAccountId returns the NewRelicInsightsAccountId field if non-nil, zero value otherwise.

### GetNewRelicInsightsAccountIdOk

`func (o *Group) GetNewRelicInsightsAccountIdOk() (*string, bool)`

GetNewRelicInsightsAccountIdOk returns a tuple with the NewRelicInsightsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelicInsightsAccountId

`func (o *Group) SetNewRelicInsightsAccountId(v string)`

SetNewRelicInsightsAccountId sets NewRelicInsightsAccountId field to given value.

### HasNewRelicInsightsAccountId

`func (o *Group) HasNewRelicInsightsAccountId() bool`

HasNewRelicInsightsAccountId returns a boolean if a field has been set.
### GetNewRelicInsightsReadToken

`func (o *Group) GetNewRelicInsightsReadToken() string`

GetNewRelicInsightsReadToken returns the NewRelicInsightsReadToken field if non-nil, zero value otherwise.

### GetNewRelicInsightsReadTokenOk

`func (o *Group) GetNewRelicInsightsReadTokenOk() (*string, bool)`

GetNewRelicInsightsReadTokenOk returns a tuple with the NewRelicInsightsReadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelicInsightsReadToken

`func (o *Group) SetNewRelicInsightsReadToken(v string)`

SetNewRelicInsightsReadToken sets NewRelicInsightsReadToken field to given value.

### HasNewRelicInsightsReadToken

`func (o *Group) HasNewRelicInsightsReadToken() bool`

HasNewRelicInsightsReadToken returns a boolean if a field has been set.
### GetNewRelicInsightsWriteToken

`func (o *Group) GetNewRelicInsightsWriteToken() string`

GetNewRelicInsightsWriteToken returns the NewRelicInsightsWriteToken field if non-nil, zero value otherwise.

### GetNewRelicInsightsWriteTokenOk

`func (o *Group) GetNewRelicInsightsWriteTokenOk() (*string, bool)`

GetNewRelicInsightsWriteTokenOk returns a tuple with the NewRelicInsightsWriteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelicInsightsWriteToken

`func (o *Group) SetNewRelicInsightsWriteToken(v string)`

SetNewRelicInsightsWriteToken sets NewRelicInsightsWriteToken field to given value.

### HasNewRelicInsightsWriteToken

`func (o *Group) HasNewRelicInsightsWriteToken() bool`

HasNewRelicInsightsWriteToken returns a boolean if a field has been set.
### GetNewRelicLicenseKey

`func (o *Group) GetNewRelicLicenseKey() string`

GetNewRelicLicenseKey returns the NewRelicLicenseKey field if non-nil, zero value otherwise.

### GetNewRelicLicenseKeyOk

`func (o *Group) GetNewRelicLicenseKeyOk() (*string, bool)`

GetNewRelicLicenseKeyOk returns a tuple with the NewRelicLicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelicLicenseKey

`func (o *Group) SetNewRelicLicenseKey(v string)`

SetNewRelicLicenseKey sets NewRelicLicenseKey field to given value.

### HasNewRelicLicenseKey

`func (o *Group) HasNewRelicLicenseKey() bool`

HasNewRelicLicenseKey returns a boolean if a field has been set.
### GetPreferredHostnames

`func (o *Group) GetPreferredHostnames() []PreferredHostname`

GetPreferredHostnames returns the PreferredHostnames field if non-nil, zero value otherwise.

### GetPreferredHostnamesOk

`func (o *Group) GetPreferredHostnamesOk() (*[]PreferredHostname, bool)`

GetPreferredHostnamesOk returns a tuple with the PreferredHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredHostnames

`func (o *Group) SetPreferredHostnames(v []PreferredHostname)`

SetPreferredHostnames sets PreferredHostnames field to given value.

### HasPreferredHostnames

`func (o *Group) HasPreferredHostnames() bool`

HasPreferredHostnames returns a boolean if a field has been set.
### GetProm

`func (o *Group) GetProm() PrometheusConfig`

GetProm returns the Prom field if non-nil, zero value otherwise.

### GetPromOk

`func (o *Group) GetPromOk() (*PrometheusConfig, bool)`

GetPromOk returns a tuple with the Prom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProm

`func (o *Group) SetProm(v PrometheusConfig)`

SetProm sets Prom field to given value.

### HasProm

`func (o *Group) HasProm() bool`

HasProm returns a boolean if a field has been set.
### GetStatus

`func (o *Group) GetStatus() GroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Group) GetStatusOk() (*GroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Group) SetStatus(v GroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Group) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetStorageConfig

`func (o *Group) GetStorageConfig() GroupStorageConfig`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *Group) GetStorageConfigOk() (*GroupStorageConfig, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *Group) SetStorageConfig(v GroupStorageConfig)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *Group) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.
### GetSummaryStatistics

`func (o *Group) GetSummaryStatistics() GroupSummaryStatistics`

GetSummaryStatistics returns the SummaryStatistics field if non-nil, zero value otherwise.

### GetSummaryStatisticsOk

`func (o *Group) GetSummaryStatisticsOk() (*GroupSummaryStatistics, bool)`

GetSummaryStatisticsOk returns a tuple with the SummaryStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStatistics

`func (o *Group) SetSummaryStatistics(v GroupSummaryStatistics)`

SetSummaryStatistics sets SummaryStatistics field to given value.

### HasSummaryStatistics

`func (o *Group) HasSummaryStatistics() bool`

HasSummaryStatistics returns a boolean if a field has been set.
### GetSuppressMongosAutoDiscovery

`func (o *Group) GetSuppressMongosAutoDiscovery() bool`

GetSuppressMongosAutoDiscovery returns the SuppressMongosAutoDiscovery field if non-nil, zero value otherwise.

### GetSuppressMongosAutoDiscoveryOk

`func (o *Group) GetSuppressMongosAutoDiscoveryOk() (*bool, bool)`

GetSuppressMongosAutoDiscoveryOk returns a tuple with the SuppressMongosAutoDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressMongosAutoDiscovery

`func (o *Group) SetSuppressMongosAutoDiscovery(v bool)`

SetSuppressMongosAutoDiscovery sets SuppressMongosAutoDiscovery field to given value.

### HasSuppressMongosAutoDiscovery

`func (o *Group) HasSuppressMongosAutoDiscovery() bool`

HasSuppressMongosAutoDiscovery returns a boolean if a field has been set.
### GetTeams

`func (o *Group) GetTeams() []TeamRoleAssignment`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Group) GetTeamsOk() (*[]TeamRoleAssignment, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Group) SetTeams(v []TeamRoleAssignment)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *Group) HasTeams() bool`

HasTeams returns a boolean if a field has been set.
### GetUseCNRegionsOnly

`func (o *Group) GetUseCNRegionsOnly() bool`

GetUseCNRegionsOnly returns the UseCNRegionsOnly field if non-nil, zero value otherwise.

### GetUseCNRegionsOnlyOk

`func (o *Group) GetUseCNRegionsOnlyOk() (*bool, bool)`

GetUseCNRegionsOnlyOk returns a tuple with the UseCNRegionsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCNRegionsOnly

`func (o *Group) SetUseCNRegionsOnly(v bool)`

SetUseCNRegionsOnly sets UseCNRegionsOnly field to given value.

### HasUseCNRegionsOnly

`func (o *Group) HasUseCNRegionsOnly() bool`

HasUseCNRegionsOnly returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


