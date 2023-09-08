// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// Group struct for Group
type Group struct {
	ActiveAgentCount            *int                                 `json:"activeAgentCount,omitempty"`
	Chttl                       *int64                               `json:"chttl,omitempty"`
	Cre                         *time.Time                           `json:"cre,omitempty"`
	DefaultTimeZoneDisplay      *string                              `json:"defaultTimeZoneDisplay,omitempty"`
	DefaultTimeZoneDisplayShort *string                              `json:"defaultTimeZoneDisplayShort,omitempty"`
	DefaultTimeZoneId           *string                              `json:"defaultTimeZoneId,omitempty"`
	DisableDbstats              *bool                                `json:"disableDbstats,omitempty"`
	EnableAllHostLogs           *bool                                `json:"enableAllHostLogs,omitempty"`
	EnableAllHostProfilers      *bool                                `json:"enableAllHostProfilers,omitempty"`
	EnableCurrentIpWarning      *bool                                `json:"enableCurrentIpWarning,omitempty"`
	Experiments                 []ExperimentRegistration             `json:"experiments,omitempty"`
	FeatureFlags                []ToggleableFeatureFlag              `json:"featureFlags,omitempty"`
	GroupType                   *string                              `json:"groupType,omitempty"`
	HasActiveBackups            *bool                                `json:"hasActiveBackups,omitempty"`
	HasActiveUI                 *bool                                `json:"hasActiveUI,omitempty"`
	HasAddedHosts               *bool                                `json:"hasAddedHosts,omitempty"`
	Id                          *string                              `json:"id,omitempty"`
	Integrations                *map[string][]map[string]interface{} `json:"integrations,omitempty"`
	IsIntercomEligible          *bool                                `json:"isIntercomEligible,omitempty"`
	IsProjectOverviewEnabled    *bool                                `json:"isProjectOverviewEnabled,omitempty"`
	LastClusterActiveSampleTime *time.Time                           `json:"lastClusterActiveSampleTime,omitempty"`
	Lmav                        *string                              `json:"lmav,omitempty"`
	Name                        *string                              `json:"name,omitempty"`
	NewRelicInsightsAccountId   *string                              `json:"newRelicInsightsAccountId,omitempty"`
	NewRelicInsightsReadToken   *string                              `json:"newRelicInsightsReadToken,omitempty"`
	NewRelicInsightsWriteToken  *string                              `json:"newRelicInsightsWriteToken,omitempty"`
	NewRelicLicenseKey          *string                              `json:"newRelicLicenseKey,omitempty"`
	PreferredHostnames          []PreferredHostname                  `json:"preferredHostnames,omitempty"`
	Prom                        *PrometheusConfig                    `json:"prom,omitempty"`
	Status                      *GroupStatus                         `json:"status,omitempty"`
	StorageConfig               *GroupStorageConfig                  `json:"storageConfig,omitempty"`
	SummaryStatistics           *GroupSummaryStatistics              `json:"summaryStatistics,omitempty"`
	SuppressMongosAutoDiscovery *bool                                `json:"suppressMongosAutoDiscovery,omitempty"`
	Teams                       []TeamRoleAssignment                 `json:"teams,omitempty"`
	UseCNRegionsOnly            *bool                                `json:"useCNRegionsOnly,omitempty"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetActiveAgentCount returns the ActiveAgentCount field value if set, zero value otherwise
func (o *Group) GetActiveAgentCount() int {
	if o == nil || IsNil(o.ActiveAgentCount) {
		var ret int
		return ret
	}
	return *o.ActiveAgentCount
}

// GetActiveAgentCountOk returns a tuple with the ActiveAgentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetActiveAgentCountOk() (*int, bool) {
	if o == nil || IsNil(o.ActiveAgentCount) {
		return nil, false
	}

	return o.ActiveAgentCount, true
}

// HasActiveAgentCount returns a boolean if a field has been set.
func (o *Group) HasActiveAgentCount() bool {
	if o != nil && !IsNil(o.ActiveAgentCount) {
		return true
	}

	return false
}

// SetActiveAgentCount gets a reference to the given int and assigns it to the ActiveAgentCount field.
func (o *Group) SetActiveAgentCount(v int) {
	o.ActiveAgentCount = &v
}

// GetChttl returns the Chttl field value if set, zero value otherwise
func (o *Group) GetChttl() int64 {
	if o == nil || IsNil(o.Chttl) {
		var ret int64
		return ret
	}
	return *o.Chttl
}

// GetChttlOk returns a tuple with the Chttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetChttlOk() (*int64, bool) {
	if o == nil || IsNil(o.Chttl) {
		return nil, false
	}

	return o.Chttl, true
}

// HasChttl returns a boolean if a field has been set.
func (o *Group) HasChttl() bool {
	if o != nil && !IsNil(o.Chttl) {
		return true
	}

	return false
}

// SetChttl gets a reference to the given int64 and assigns it to the Chttl field.
func (o *Group) SetChttl(v int64) {
	o.Chttl = &v
}

// GetCre returns the Cre field value if set, zero value otherwise
func (o *Group) GetCre() time.Time {
	if o == nil || IsNil(o.Cre) {
		var ret time.Time
		return ret
	}
	return *o.Cre
}

// GetCreOk returns a tuple with the Cre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetCreOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Cre) {
		return nil, false
	}

	return o.Cre, true
}

// HasCre returns a boolean if a field has been set.
func (o *Group) HasCre() bool {
	if o != nil && !IsNil(o.Cre) {
		return true
	}

	return false
}

// SetCre gets a reference to the given time.Time and assigns it to the Cre field.
func (o *Group) SetCre(v time.Time) {
	o.Cre = &v
}

// GetDefaultTimeZoneDisplay returns the DefaultTimeZoneDisplay field value if set, zero value otherwise
func (o *Group) GetDefaultTimeZoneDisplay() string {
	if o == nil || IsNil(o.DefaultTimeZoneDisplay) {
		var ret string
		return ret
	}
	return *o.DefaultTimeZoneDisplay
}

// GetDefaultTimeZoneDisplayOk returns a tuple with the DefaultTimeZoneDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDefaultTimeZoneDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTimeZoneDisplay) {
		return nil, false
	}

	return o.DefaultTimeZoneDisplay, true
}

// HasDefaultTimeZoneDisplay returns a boolean if a field has been set.
func (o *Group) HasDefaultTimeZoneDisplay() bool {
	if o != nil && !IsNil(o.DefaultTimeZoneDisplay) {
		return true
	}

	return false
}

// SetDefaultTimeZoneDisplay gets a reference to the given string and assigns it to the DefaultTimeZoneDisplay field.
func (o *Group) SetDefaultTimeZoneDisplay(v string) {
	o.DefaultTimeZoneDisplay = &v
}

// GetDefaultTimeZoneDisplayShort returns the DefaultTimeZoneDisplayShort field value if set, zero value otherwise
func (o *Group) GetDefaultTimeZoneDisplayShort() string {
	if o == nil || IsNil(o.DefaultTimeZoneDisplayShort) {
		var ret string
		return ret
	}
	return *o.DefaultTimeZoneDisplayShort
}

// GetDefaultTimeZoneDisplayShortOk returns a tuple with the DefaultTimeZoneDisplayShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDefaultTimeZoneDisplayShortOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTimeZoneDisplayShort) {
		return nil, false
	}

	return o.DefaultTimeZoneDisplayShort, true
}

// HasDefaultTimeZoneDisplayShort returns a boolean if a field has been set.
func (o *Group) HasDefaultTimeZoneDisplayShort() bool {
	if o != nil && !IsNil(o.DefaultTimeZoneDisplayShort) {
		return true
	}

	return false
}

// SetDefaultTimeZoneDisplayShort gets a reference to the given string and assigns it to the DefaultTimeZoneDisplayShort field.
func (o *Group) SetDefaultTimeZoneDisplayShort(v string) {
	o.DefaultTimeZoneDisplayShort = &v
}

// GetDefaultTimeZoneId returns the DefaultTimeZoneId field value if set, zero value otherwise
func (o *Group) GetDefaultTimeZoneId() string {
	if o == nil || IsNil(o.DefaultTimeZoneId) {
		var ret string
		return ret
	}
	return *o.DefaultTimeZoneId
}

// GetDefaultTimeZoneIdOk returns a tuple with the DefaultTimeZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDefaultTimeZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTimeZoneId) {
		return nil, false
	}

	return o.DefaultTimeZoneId, true
}

// HasDefaultTimeZoneId returns a boolean if a field has been set.
func (o *Group) HasDefaultTimeZoneId() bool {
	if o != nil && !IsNil(o.DefaultTimeZoneId) {
		return true
	}

	return false
}

// SetDefaultTimeZoneId gets a reference to the given string and assigns it to the DefaultTimeZoneId field.
func (o *Group) SetDefaultTimeZoneId(v string) {
	o.DefaultTimeZoneId = &v
}

// GetDisableDbstats returns the DisableDbstats field value if set, zero value otherwise
func (o *Group) GetDisableDbstats() bool {
	if o == nil || IsNil(o.DisableDbstats) {
		var ret bool
		return ret
	}
	return *o.DisableDbstats
}

// GetDisableDbstatsOk returns a tuple with the DisableDbstats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDisableDbstatsOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableDbstats) {
		return nil, false
	}

	return o.DisableDbstats, true
}

// HasDisableDbstats returns a boolean if a field has been set.
func (o *Group) HasDisableDbstats() bool {
	if o != nil && !IsNil(o.DisableDbstats) {
		return true
	}

	return false
}

// SetDisableDbstats gets a reference to the given bool and assigns it to the DisableDbstats field.
func (o *Group) SetDisableDbstats(v bool) {
	o.DisableDbstats = &v
}

// GetEnableAllHostLogs returns the EnableAllHostLogs field value if set, zero value otherwise
func (o *Group) GetEnableAllHostLogs() bool {
	if o == nil || IsNil(o.EnableAllHostLogs) {
		var ret bool
		return ret
	}
	return *o.EnableAllHostLogs
}

// GetEnableAllHostLogsOk returns a tuple with the EnableAllHostLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetEnableAllHostLogsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAllHostLogs) {
		return nil, false
	}

	return o.EnableAllHostLogs, true
}

// HasEnableAllHostLogs returns a boolean if a field has been set.
func (o *Group) HasEnableAllHostLogs() bool {
	if o != nil && !IsNil(o.EnableAllHostLogs) {
		return true
	}

	return false
}

// SetEnableAllHostLogs gets a reference to the given bool and assigns it to the EnableAllHostLogs field.
func (o *Group) SetEnableAllHostLogs(v bool) {
	o.EnableAllHostLogs = &v
}

// GetEnableAllHostProfilers returns the EnableAllHostProfilers field value if set, zero value otherwise
func (o *Group) GetEnableAllHostProfilers() bool {
	if o == nil || IsNil(o.EnableAllHostProfilers) {
		var ret bool
		return ret
	}
	return *o.EnableAllHostProfilers
}

// GetEnableAllHostProfilersOk returns a tuple with the EnableAllHostProfilers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetEnableAllHostProfilersOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAllHostProfilers) {
		return nil, false
	}

	return o.EnableAllHostProfilers, true
}

// HasEnableAllHostProfilers returns a boolean if a field has been set.
func (o *Group) HasEnableAllHostProfilers() bool {
	if o != nil && !IsNil(o.EnableAllHostProfilers) {
		return true
	}

	return false
}

// SetEnableAllHostProfilers gets a reference to the given bool and assigns it to the EnableAllHostProfilers field.
func (o *Group) SetEnableAllHostProfilers(v bool) {
	o.EnableAllHostProfilers = &v
}

// GetEnableCurrentIpWarning returns the EnableCurrentIpWarning field value if set, zero value otherwise
func (o *Group) GetEnableCurrentIpWarning() bool {
	if o == nil || IsNil(o.EnableCurrentIpWarning) {
		var ret bool
		return ret
	}
	return *o.EnableCurrentIpWarning
}

// GetEnableCurrentIpWarningOk returns a tuple with the EnableCurrentIpWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetEnableCurrentIpWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCurrentIpWarning) {
		return nil, false
	}

	return o.EnableCurrentIpWarning, true
}

// HasEnableCurrentIpWarning returns a boolean if a field has been set.
func (o *Group) HasEnableCurrentIpWarning() bool {
	if o != nil && !IsNil(o.EnableCurrentIpWarning) {
		return true
	}

	return false
}

// SetEnableCurrentIpWarning gets a reference to the given bool and assigns it to the EnableCurrentIpWarning field.
func (o *Group) SetEnableCurrentIpWarning(v bool) {
	o.EnableCurrentIpWarning = &v
}

// GetExperiments returns the Experiments field value if set, zero value otherwise
func (o *Group) GetExperiments() []ExperimentRegistration {
	if o == nil || IsNil(o.Experiments) {
		var ret []ExperimentRegistration
		return ret
	}
	return o.Experiments
}

// GetExperimentsOk returns a tuple with the Experiments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetExperimentsOk() ([]ExperimentRegistration, bool) {
	if o == nil || IsNil(o.Experiments) {
		return nil, false
	}

	return o.Experiments, true
}

// HasExperiments returns a boolean if a field has been set.
func (o *Group) HasExperiments() bool {
	if o != nil && !IsNil(o.Experiments) {
		return true
	}

	return false
}

// SetExperiments gets a reference to the given []ExperimentRegistration and assigns it to the Experiments field.
func (o *Group) SetExperiments(v []ExperimentRegistration) {
	o.Experiments = v
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise
func (o *Group) GetFeatureFlags() []ToggleableFeatureFlag {
	if o == nil || IsNil(o.FeatureFlags) {
		var ret []ToggleableFeatureFlag
		return ret
	}
	return o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetFeatureFlagsOk() ([]ToggleableFeatureFlag, bool) {
	if o == nil || IsNil(o.FeatureFlags) {
		return nil, false
	}

	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *Group) HasFeatureFlags() bool {
	if o != nil && !IsNil(o.FeatureFlags) {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given []ToggleableFeatureFlag and assigns it to the FeatureFlags field.
func (o *Group) SetFeatureFlags(v []ToggleableFeatureFlag) {
	o.FeatureFlags = v
}

// GetGroupType returns the GroupType field value if set, zero value otherwise
func (o *Group) GetGroupType() string {
	if o == nil || IsNil(o.GroupType) {
		var ret string
		return ret
	}
	return *o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetGroupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupType) {
		return nil, false
	}

	return o.GroupType, true
}

// HasGroupType returns a boolean if a field has been set.
func (o *Group) HasGroupType() bool {
	if o != nil && !IsNil(o.GroupType) {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given string and assigns it to the GroupType field.
func (o *Group) SetGroupType(v string) {
	o.GroupType = &v
}

// GetHasActiveBackups returns the HasActiveBackups field value if set, zero value otherwise
func (o *Group) GetHasActiveBackups() bool {
	if o == nil || IsNil(o.HasActiveBackups) {
		var ret bool
		return ret
	}
	return *o.HasActiveBackups
}

// GetHasActiveBackupsOk returns a tuple with the HasActiveBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetHasActiveBackupsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasActiveBackups) {
		return nil, false
	}

	return o.HasActiveBackups, true
}

// HasHasActiveBackups returns a boolean if a field has been set.
func (o *Group) HasHasActiveBackups() bool {
	if o != nil && !IsNil(o.HasActiveBackups) {
		return true
	}

	return false
}

// SetHasActiveBackups gets a reference to the given bool and assigns it to the HasActiveBackups field.
func (o *Group) SetHasActiveBackups(v bool) {
	o.HasActiveBackups = &v
}

// GetHasActiveUI returns the HasActiveUI field value if set, zero value otherwise
func (o *Group) GetHasActiveUI() bool {
	if o == nil || IsNil(o.HasActiveUI) {
		var ret bool
		return ret
	}
	return *o.HasActiveUI
}

// GetHasActiveUIOk returns a tuple with the HasActiveUI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetHasActiveUIOk() (*bool, bool) {
	if o == nil || IsNil(o.HasActiveUI) {
		return nil, false
	}

	return o.HasActiveUI, true
}

// HasHasActiveUI returns a boolean if a field has been set.
func (o *Group) HasHasActiveUI() bool {
	if o != nil && !IsNil(o.HasActiveUI) {
		return true
	}

	return false
}

// SetHasActiveUI gets a reference to the given bool and assigns it to the HasActiveUI field.
func (o *Group) SetHasActiveUI(v bool) {
	o.HasActiveUI = &v
}

// GetHasAddedHosts returns the HasAddedHosts field value if set, zero value otherwise
func (o *Group) GetHasAddedHosts() bool {
	if o == nil || IsNil(o.HasAddedHosts) {
		var ret bool
		return ret
	}
	return *o.HasAddedHosts
}

// GetHasAddedHostsOk returns a tuple with the HasAddedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetHasAddedHostsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAddedHosts) {
		return nil, false
	}

	return o.HasAddedHosts, true
}

// HasHasAddedHosts returns a boolean if a field has been set.
func (o *Group) HasHasAddedHosts() bool {
	if o != nil && !IsNil(o.HasAddedHosts) {
		return true
	}

	return false
}

// SetHasAddedHosts gets a reference to the given bool and assigns it to the HasAddedHosts field.
func (o *Group) SetHasAddedHosts(v bool) {
	o.HasAddedHosts = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *Group) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Group) SetId(v string) {
	o.Id = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise
func (o *Group) GetIntegrations() map[string][]map[string]interface{} {
	if o == nil || IsNil(o.Integrations) {
		var ret map[string][]map[string]interface{}
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIntegrationsOk() (*map[string][]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}

	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *Group) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given map[string][]map[string]interface{} and assigns it to the Integrations field.
func (o *Group) SetIntegrations(v map[string][]map[string]interface{}) {
	o.Integrations = &v
}

// GetIsIntercomEligible returns the IsIntercomEligible field value if set, zero value otherwise
func (o *Group) GetIsIntercomEligible() bool {
	if o == nil || IsNil(o.IsIntercomEligible) {
		var ret bool
		return ret
	}
	return *o.IsIntercomEligible
}

// GetIsIntercomEligibleOk returns a tuple with the IsIntercomEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIsIntercomEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIntercomEligible) {
		return nil, false
	}

	return o.IsIntercomEligible, true
}

// HasIsIntercomEligible returns a boolean if a field has been set.
func (o *Group) HasIsIntercomEligible() bool {
	if o != nil && !IsNil(o.IsIntercomEligible) {
		return true
	}

	return false
}

// SetIsIntercomEligible gets a reference to the given bool and assigns it to the IsIntercomEligible field.
func (o *Group) SetIsIntercomEligible(v bool) {
	o.IsIntercomEligible = &v
}

// GetIsProjectOverviewEnabled returns the IsProjectOverviewEnabled field value if set, zero value otherwise
func (o *Group) GetIsProjectOverviewEnabled() bool {
	if o == nil || IsNil(o.IsProjectOverviewEnabled) {
		var ret bool
		return ret
	}
	return *o.IsProjectOverviewEnabled
}

// GetIsProjectOverviewEnabledOk returns a tuple with the IsProjectOverviewEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIsProjectOverviewEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProjectOverviewEnabled) {
		return nil, false
	}

	return o.IsProjectOverviewEnabled, true
}

// HasIsProjectOverviewEnabled returns a boolean if a field has been set.
func (o *Group) HasIsProjectOverviewEnabled() bool {
	if o != nil && !IsNil(o.IsProjectOverviewEnabled) {
		return true
	}

	return false
}

// SetIsProjectOverviewEnabled gets a reference to the given bool and assigns it to the IsProjectOverviewEnabled field.
func (o *Group) SetIsProjectOverviewEnabled(v bool) {
	o.IsProjectOverviewEnabled = &v
}

// GetLastClusterActiveSampleTime returns the LastClusterActiveSampleTime field value if set, zero value otherwise
func (o *Group) GetLastClusterActiveSampleTime() time.Time {
	if o == nil || IsNil(o.LastClusterActiveSampleTime) {
		var ret time.Time
		return ret
	}
	return *o.LastClusterActiveSampleTime
}

// GetLastClusterActiveSampleTimeOk returns a tuple with the LastClusterActiveSampleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLastClusterActiveSampleTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastClusterActiveSampleTime) {
		return nil, false
	}

	return o.LastClusterActiveSampleTime, true
}

// HasLastClusterActiveSampleTime returns a boolean if a field has been set.
func (o *Group) HasLastClusterActiveSampleTime() bool {
	if o != nil && !IsNil(o.LastClusterActiveSampleTime) {
		return true
	}

	return false
}

// SetLastClusterActiveSampleTime gets a reference to the given time.Time and assigns it to the LastClusterActiveSampleTime field.
func (o *Group) SetLastClusterActiveSampleTime(v time.Time) {
	o.LastClusterActiveSampleTime = &v
}

// GetLmav returns the Lmav field value if set, zero value otherwise
func (o *Group) GetLmav() string {
	if o == nil || IsNil(o.Lmav) {
		var ret string
		return ret
	}
	return *o.Lmav
}

// GetLmavOk returns a tuple with the Lmav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLmavOk() (*string, bool) {
	if o == nil || IsNil(o.Lmav) {
		return nil, false
	}

	return o.Lmav, true
}

// HasLmav returns a boolean if a field has been set.
func (o *Group) HasLmav() bool {
	if o != nil && !IsNil(o.Lmav) {
		return true
	}

	return false
}

// SetLmav gets a reference to the given string and assigns it to the Lmav field.
func (o *Group) SetLmav(v string) {
	o.Lmav = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *Group) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Group) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Group) SetName(v string) {
	o.Name = &v
}

// GetNewRelicInsightsAccountId returns the NewRelicInsightsAccountId field value if set, zero value otherwise
func (o *Group) GetNewRelicInsightsAccountId() string {
	if o == nil || IsNil(o.NewRelicInsightsAccountId) {
		var ret string
		return ret
	}
	return *o.NewRelicInsightsAccountId
}

// GetNewRelicInsightsAccountIdOk returns a tuple with the NewRelicInsightsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNewRelicInsightsAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewRelicInsightsAccountId) {
		return nil, false
	}

	return o.NewRelicInsightsAccountId, true
}

// HasNewRelicInsightsAccountId returns a boolean if a field has been set.
func (o *Group) HasNewRelicInsightsAccountId() bool {
	if o != nil && !IsNil(o.NewRelicInsightsAccountId) {
		return true
	}

	return false
}

// SetNewRelicInsightsAccountId gets a reference to the given string and assigns it to the NewRelicInsightsAccountId field.
func (o *Group) SetNewRelicInsightsAccountId(v string) {
	o.NewRelicInsightsAccountId = &v
}

// GetNewRelicInsightsReadToken returns the NewRelicInsightsReadToken field value if set, zero value otherwise
func (o *Group) GetNewRelicInsightsReadToken() string {
	if o == nil || IsNil(o.NewRelicInsightsReadToken) {
		var ret string
		return ret
	}
	return *o.NewRelicInsightsReadToken
}

// GetNewRelicInsightsReadTokenOk returns a tuple with the NewRelicInsightsReadToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNewRelicInsightsReadTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NewRelicInsightsReadToken) {
		return nil, false
	}

	return o.NewRelicInsightsReadToken, true
}

// HasNewRelicInsightsReadToken returns a boolean if a field has been set.
func (o *Group) HasNewRelicInsightsReadToken() bool {
	if o != nil && !IsNil(o.NewRelicInsightsReadToken) {
		return true
	}

	return false
}

// SetNewRelicInsightsReadToken gets a reference to the given string and assigns it to the NewRelicInsightsReadToken field.
func (o *Group) SetNewRelicInsightsReadToken(v string) {
	o.NewRelicInsightsReadToken = &v
}

// GetNewRelicInsightsWriteToken returns the NewRelicInsightsWriteToken field value if set, zero value otherwise
func (o *Group) GetNewRelicInsightsWriteToken() string {
	if o == nil || IsNil(o.NewRelicInsightsWriteToken) {
		var ret string
		return ret
	}
	return *o.NewRelicInsightsWriteToken
}

// GetNewRelicInsightsWriteTokenOk returns a tuple with the NewRelicInsightsWriteToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNewRelicInsightsWriteTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NewRelicInsightsWriteToken) {
		return nil, false
	}

	return o.NewRelicInsightsWriteToken, true
}

// HasNewRelicInsightsWriteToken returns a boolean if a field has been set.
func (o *Group) HasNewRelicInsightsWriteToken() bool {
	if o != nil && !IsNil(o.NewRelicInsightsWriteToken) {
		return true
	}

	return false
}

// SetNewRelicInsightsWriteToken gets a reference to the given string and assigns it to the NewRelicInsightsWriteToken field.
func (o *Group) SetNewRelicInsightsWriteToken(v string) {
	o.NewRelicInsightsWriteToken = &v
}

// GetNewRelicLicenseKey returns the NewRelicLicenseKey field value if set, zero value otherwise
func (o *Group) GetNewRelicLicenseKey() string {
	if o == nil || IsNil(o.NewRelicLicenseKey) {
		var ret string
		return ret
	}
	return *o.NewRelicLicenseKey
}

// GetNewRelicLicenseKeyOk returns a tuple with the NewRelicLicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNewRelicLicenseKeyOk() (*string, bool) {
	if o == nil || IsNil(o.NewRelicLicenseKey) {
		return nil, false
	}

	return o.NewRelicLicenseKey, true
}

// HasNewRelicLicenseKey returns a boolean if a field has been set.
func (o *Group) HasNewRelicLicenseKey() bool {
	if o != nil && !IsNil(o.NewRelicLicenseKey) {
		return true
	}

	return false
}

// SetNewRelicLicenseKey gets a reference to the given string and assigns it to the NewRelicLicenseKey field.
func (o *Group) SetNewRelicLicenseKey(v string) {
	o.NewRelicLicenseKey = &v
}

// GetPreferredHostnames returns the PreferredHostnames field value if set, zero value otherwise
func (o *Group) GetPreferredHostnames() []PreferredHostname {
	if o == nil || IsNil(o.PreferredHostnames) {
		var ret []PreferredHostname
		return ret
	}
	return o.PreferredHostnames
}

// GetPreferredHostnamesOk returns a tuple with the PreferredHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetPreferredHostnamesOk() ([]PreferredHostname, bool) {
	if o == nil || IsNil(o.PreferredHostnames) {
		return nil, false
	}

	return o.PreferredHostnames, true
}

// HasPreferredHostnames returns a boolean if a field has been set.
func (o *Group) HasPreferredHostnames() bool {
	if o != nil && !IsNil(o.PreferredHostnames) {
		return true
	}

	return false
}

// SetPreferredHostnames gets a reference to the given []PreferredHostname and assigns it to the PreferredHostnames field.
func (o *Group) SetPreferredHostnames(v []PreferredHostname) {
	o.PreferredHostnames = v
}

// GetProm returns the Prom field value if set, zero value otherwise
func (o *Group) GetProm() PrometheusConfig {
	if o == nil || IsNil(o.Prom) {
		var ret PrometheusConfig
		return ret
	}
	return *o.Prom
}

// GetPromOk returns a tuple with the Prom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetPromOk() (*PrometheusConfig, bool) {
	if o == nil || IsNil(o.Prom) {
		return nil, false
	}

	return o.Prom, true
}

// HasProm returns a boolean if a field has been set.
func (o *Group) HasProm() bool {
	if o != nil && !IsNil(o.Prom) {
		return true
	}

	return false
}

// SetProm gets a reference to the given PrometheusConfig and assigns it to the Prom field.
func (o *Group) SetProm(v PrometheusConfig) {
	o.Prom = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *Group) GetStatus() GroupStatus {
	if o == nil || IsNil(o.Status) {
		var ret GroupStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetStatusOk() (*GroupStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Group) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GroupStatus and assigns it to the Status field.
func (o *Group) SetStatus(v GroupStatus) {
	o.Status = &v
}

// GetStorageConfig returns the StorageConfig field value if set, zero value otherwise
func (o *Group) GetStorageConfig() GroupStorageConfig {
	if o == nil || IsNil(o.StorageConfig) {
		var ret GroupStorageConfig
		return ret
	}
	return *o.StorageConfig
}

// GetStorageConfigOk returns a tuple with the StorageConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetStorageConfigOk() (*GroupStorageConfig, bool) {
	if o == nil || IsNil(o.StorageConfig) {
		return nil, false
	}

	return o.StorageConfig, true
}

// HasStorageConfig returns a boolean if a field has been set.
func (o *Group) HasStorageConfig() bool {
	if o != nil && !IsNil(o.StorageConfig) {
		return true
	}

	return false
}

// SetStorageConfig gets a reference to the given GroupStorageConfig and assigns it to the StorageConfig field.
func (o *Group) SetStorageConfig(v GroupStorageConfig) {
	o.StorageConfig = &v
}

// GetSummaryStatistics returns the SummaryStatistics field value if set, zero value otherwise
func (o *Group) GetSummaryStatistics() GroupSummaryStatistics {
	if o == nil || IsNil(o.SummaryStatistics) {
		var ret GroupSummaryStatistics
		return ret
	}
	return *o.SummaryStatistics
}

// GetSummaryStatisticsOk returns a tuple with the SummaryStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSummaryStatisticsOk() (*GroupSummaryStatistics, bool) {
	if o == nil || IsNil(o.SummaryStatistics) {
		return nil, false
	}

	return o.SummaryStatistics, true
}

// HasSummaryStatistics returns a boolean if a field has been set.
func (o *Group) HasSummaryStatistics() bool {
	if o != nil && !IsNil(o.SummaryStatistics) {
		return true
	}

	return false
}

// SetSummaryStatistics gets a reference to the given GroupSummaryStatistics and assigns it to the SummaryStatistics field.
func (o *Group) SetSummaryStatistics(v GroupSummaryStatistics) {
	o.SummaryStatistics = &v
}

// GetSuppressMongosAutoDiscovery returns the SuppressMongosAutoDiscovery field value if set, zero value otherwise
func (o *Group) GetSuppressMongosAutoDiscovery() bool {
	if o == nil || IsNil(o.SuppressMongosAutoDiscovery) {
		var ret bool
		return ret
	}
	return *o.SuppressMongosAutoDiscovery
}

// GetSuppressMongosAutoDiscoveryOk returns a tuple with the SuppressMongosAutoDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSuppressMongosAutoDiscoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressMongosAutoDiscovery) {
		return nil, false
	}

	return o.SuppressMongosAutoDiscovery, true
}

// HasSuppressMongosAutoDiscovery returns a boolean if a field has been set.
func (o *Group) HasSuppressMongosAutoDiscovery() bool {
	if o != nil && !IsNil(o.SuppressMongosAutoDiscovery) {
		return true
	}

	return false
}

// SetSuppressMongosAutoDiscovery gets a reference to the given bool and assigns it to the SuppressMongosAutoDiscovery field.
func (o *Group) SetSuppressMongosAutoDiscovery(v bool) {
	o.SuppressMongosAutoDiscovery = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise
func (o *Group) GetTeams() []TeamRoleAssignment {
	if o == nil || IsNil(o.Teams) {
		var ret []TeamRoleAssignment
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetTeamsOk() ([]TeamRoleAssignment, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}

	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *Group) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []TeamRoleAssignment and assigns it to the Teams field.
func (o *Group) SetTeams(v []TeamRoleAssignment) {
	o.Teams = v
}

// GetUseCNRegionsOnly returns the UseCNRegionsOnly field value if set, zero value otherwise
func (o *Group) GetUseCNRegionsOnly() bool {
	if o == nil || IsNil(o.UseCNRegionsOnly) {
		var ret bool
		return ret
	}
	return *o.UseCNRegionsOnly
}

// GetUseCNRegionsOnlyOk returns a tuple with the UseCNRegionsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUseCNRegionsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCNRegionsOnly) {
		return nil, false
	}

	return o.UseCNRegionsOnly, true
}

// HasUseCNRegionsOnly returns a boolean if a field has been set.
func (o *Group) HasUseCNRegionsOnly() bool {
	if o != nil && !IsNil(o.UseCNRegionsOnly) {
		return true
	}

	return false
}

// SetUseCNRegionsOnly gets a reference to the given bool and assigns it to the UseCNRegionsOnly field.
func (o *Group) SetUseCNRegionsOnly(v bool) {
	o.UseCNRegionsOnly = &v
}

func (o Group) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Group) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveAgentCount) {
		toSerialize["activeAgentCount"] = o.ActiveAgentCount
	}
	if !IsNil(o.Chttl) {
		toSerialize["chttl"] = o.Chttl
	}
	if !IsNil(o.Cre) {
		toSerialize["cre"] = o.Cre
	}
	if !IsNil(o.DefaultTimeZoneDisplay) {
		toSerialize["defaultTimeZoneDisplay"] = o.DefaultTimeZoneDisplay
	}
	if !IsNil(o.DefaultTimeZoneDisplayShort) {
		toSerialize["defaultTimeZoneDisplayShort"] = o.DefaultTimeZoneDisplayShort
	}
	if !IsNil(o.DefaultTimeZoneId) {
		toSerialize["defaultTimeZoneId"] = o.DefaultTimeZoneId
	}
	if !IsNil(o.DisableDbstats) {
		toSerialize["disableDbstats"] = o.DisableDbstats
	}
	if !IsNil(o.EnableAllHostLogs) {
		toSerialize["enableAllHostLogs"] = o.EnableAllHostLogs
	}
	if !IsNil(o.EnableAllHostProfilers) {
		toSerialize["enableAllHostProfilers"] = o.EnableAllHostProfilers
	}
	if !IsNil(o.EnableCurrentIpWarning) {
		toSerialize["enableCurrentIpWarning"] = o.EnableCurrentIpWarning
	}
	if !IsNil(o.Experiments) {
		toSerialize["experiments"] = o.Experiments
	}
	if !IsNil(o.FeatureFlags) {
		toSerialize["featureFlags"] = o.FeatureFlags
	}
	if !IsNil(o.GroupType) {
		toSerialize["groupType"] = o.GroupType
	}
	if !IsNil(o.HasActiveBackups) {
		toSerialize["hasActiveBackups"] = o.HasActiveBackups
	}
	if !IsNil(o.HasActiveUI) {
		toSerialize["hasActiveUI"] = o.HasActiveUI
	}
	if !IsNil(o.HasAddedHosts) {
		toSerialize["hasAddedHosts"] = o.HasAddedHosts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Integrations) {
		toSerialize["integrations"] = o.Integrations
	}
	if !IsNil(o.IsIntercomEligible) {
		toSerialize["isIntercomEligible"] = o.IsIntercomEligible
	}
	if !IsNil(o.IsProjectOverviewEnabled) {
		toSerialize["isProjectOverviewEnabled"] = o.IsProjectOverviewEnabled
	}
	if !IsNil(o.LastClusterActiveSampleTime) {
		toSerialize["lastClusterActiveSampleTime"] = o.LastClusterActiveSampleTime
	}
	if !IsNil(o.Lmav) {
		toSerialize["lmav"] = o.Lmav
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NewRelicInsightsAccountId) {
		toSerialize["newRelicInsightsAccountId"] = o.NewRelicInsightsAccountId
	}
	if !IsNil(o.NewRelicInsightsReadToken) {
		toSerialize["newRelicInsightsReadToken"] = o.NewRelicInsightsReadToken
	}
	if !IsNil(o.NewRelicInsightsWriteToken) {
		toSerialize["newRelicInsightsWriteToken"] = o.NewRelicInsightsWriteToken
	}
	if !IsNil(o.NewRelicLicenseKey) {
		toSerialize["newRelicLicenseKey"] = o.NewRelicLicenseKey
	}
	if !IsNil(o.PreferredHostnames) {
		toSerialize["preferredHostnames"] = o.PreferredHostnames
	}
	if !IsNil(o.Prom) {
		toSerialize["prom"] = o.Prom
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StorageConfig) {
		toSerialize["storageConfig"] = o.StorageConfig
	}
	if !IsNil(o.SummaryStatistics) {
		toSerialize["summaryStatistics"] = o.SummaryStatistics
	}
	if !IsNil(o.SuppressMongosAutoDiscovery) {
		toSerialize["suppressMongosAutoDiscovery"] = o.SuppressMongosAutoDiscovery
	}
	if !IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	if !IsNil(o.UseCNRegionsOnly) {
		toSerialize["useCNRegionsOnly"] = o.UseCNRegionsOnly
	}
	return toSerialize, nil
}
