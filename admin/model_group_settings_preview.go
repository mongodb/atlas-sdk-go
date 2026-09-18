// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GroupSettingsPreview struct for GroupSettingsPreview
type GroupSettingsPreview struct {
	// Flag that indicates whether the MongoDB Assistant on the Atlas Home Page is enabled for the specified project.
	IsAtlasHomePageAiAssistantEnabled *bool `json:"isAtlasHomePageAiAssistantEnabled,omitempty"`
	// Flag that indicates whether the AI Cluster Assistant is enabled for the specified project.
	IsClusterAiAssistantEnabled *bool `json:"isClusterAiAssistantEnabled,omitempty"`
	// Flag that indicates whether to collect database-specific metrics for the specified project.
	IsCollectDatabaseSpecificsStatisticsEnabled *bool `json:"isCollectDatabaseSpecificsStatisticsEnabled,omitempty"`
	// Flag that indicates whether to enable the Data Explorer for the specified project.
	IsDataExplorerEnabled *bool `json:"isDataExplorerEnabled,omitempty"`
	// Flag that indicates whether to enable the use of generative AI features which make requests to 3rd party services in Data Explorer for the specified project.
	IsDataExplorerGenAIFeaturesEnabled *bool `json:"isDataExplorerGenAIFeaturesEnabled,omitempty"`
	// Flag that indicates whether to enable the passing of sample field values with the use of generative AI features in the Data Explorer for the specified project.
	IsDataExplorerGenAISampleDocumentPassingEnabled *bool `json:"isDataExplorerGenAISampleDocumentPassingEnabled,omitempty"`
	// Flag that indicates whether data validation is enabled for all clusters in the specified project.
	IsDataValidationEnabled *bool `json:"isDataValidationEnabled,omitempty"`
	// Flag that indicates whether to enable extended storage sizes for the specified project.
	IsExtendedStorageSizesEnabled *bool `json:"isExtendedStorageSizesEnabled,omitempty"`
	// Flag that indicates whether to enable Native Reranking with Voyage AI models in the Aggregation Pipeline for the specified project.
	IsNativeRerankingEnabled *bool `json:"isNativeRerankingEnabled,omitempty"`
	// Flag that indicates whether to enable the Performance Advisor and Profiler for the specified project.
	IsPerformanceAdvisorEnabled *bool `json:"isPerformanceAdvisorEnabled,omitempty"`
	// Flag that indicates whether to enable AI features in Query Insights for the specified project.
	IsQueryInsightsGenAiFeaturesEnabled *bool `json:"isQueryInsightsGenAiFeaturesEnabled,omitempty"`
	// Flag that indicates whether to enable the Real Time Performance Panel for the specified project.
	IsRealtimePerformancePanelEnabled *bool `json:"isRealtimePerformancePanelEnabled,omitempty"`
	// Flag that indicates whether to enable the Schema Advisor for the specified project.
	IsSchemaAdvisorEnabled *bool `json:"isSchemaAdvisorEnabled,omitempty"`
	// Flag that indicates whether the project uses the private endpoint connection strings resource.
	PrivateEndpointConnectionStringsEnabled *bool `json:"privateEndpointConnectionStringsEnabled,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *GroupSettingsPreview) MarshalJSON() ([]byte, error) {
	type noMethod GroupSettingsPreview
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGroupSettingsPreview instantiates a new GroupSettingsPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSettingsPreview() *GroupSettingsPreview {
	this := GroupSettingsPreview{}
	var isDataExplorerGenAISampleDocumentPassingEnabled bool = false
	this.IsDataExplorerGenAISampleDocumentPassingEnabled = &isDataExplorerGenAISampleDocumentPassingEnabled
	var privateEndpointConnectionStringsEnabled bool = false
	this.PrivateEndpointConnectionStringsEnabled = &privateEndpointConnectionStringsEnabled
	return &this
}

// NewGroupSettingsPreviewWithDefaults instantiates a new GroupSettingsPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSettingsPreviewWithDefaults() *GroupSettingsPreview {
	this := GroupSettingsPreview{}
	var isDataExplorerGenAISampleDocumentPassingEnabled bool = false
	this.IsDataExplorerGenAISampleDocumentPassingEnabled = &isDataExplorerGenAISampleDocumentPassingEnabled
	var privateEndpointConnectionStringsEnabled bool = false
	this.PrivateEndpointConnectionStringsEnabled = &privateEndpointConnectionStringsEnabled
	return &this
}

// GetIsAtlasHomePageAiAssistantEnabled returns the IsAtlasHomePageAiAssistantEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsAtlasHomePageAiAssistantEnabled() bool {
	if o == nil || IsNil(o.IsAtlasHomePageAiAssistantEnabled) {
		var ret bool
		return ret
	}
	return *o.IsAtlasHomePageAiAssistantEnabled
}

// GetIsAtlasHomePageAiAssistantEnabledOk returns a tuple with the IsAtlasHomePageAiAssistantEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsAtlasHomePageAiAssistantEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAtlasHomePageAiAssistantEnabled) {
		return nil, false
	}

	return o.IsAtlasHomePageAiAssistantEnabled, true
}

// HasIsAtlasHomePageAiAssistantEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsAtlasHomePageAiAssistantEnabled() bool {
	if o != nil && !IsNil(o.IsAtlasHomePageAiAssistantEnabled) {
		return true
	}

	return false
}

// SetIsAtlasHomePageAiAssistantEnabled gets a reference to the given bool and assigns it to the IsAtlasHomePageAiAssistantEnabled field.
func (o *GroupSettingsPreview) SetIsAtlasHomePageAiAssistantEnabled(v bool) {
	o.IsAtlasHomePageAiAssistantEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsAtlasHomePageAiAssistantEnabled")
}

// SetIsAtlasHomePageAiAssistantEnabledNil sets IsAtlasHomePageAiAssistantEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsAtlasHomePageAiAssistantEnabledNil() {
	o.IsAtlasHomePageAiAssistantEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsAtlasHomePageAiAssistantEnabled")
}

// GetIsClusterAiAssistantEnabled returns the IsClusterAiAssistantEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsClusterAiAssistantEnabled() bool {
	if o == nil || IsNil(o.IsClusterAiAssistantEnabled) {
		var ret bool
		return ret
	}
	return *o.IsClusterAiAssistantEnabled
}

// GetIsClusterAiAssistantEnabledOk returns a tuple with the IsClusterAiAssistantEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsClusterAiAssistantEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsClusterAiAssistantEnabled) {
		return nil, false
	}

	return o.IsClusterAiAssistantEnabled, true
}

// HasIsClusterAiAssistantEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsClusterAiAssistantEnabled() bool {
	if o != nil && !IsNil(o.IsClusterAiAssistantEnabled) {
		return true
	}

	return false
}

// SetIsClusterAiAssistantEnabled gets a reference to the given bool and assigns it to the IsClusterAiAssistantEnabled field.
func (o *GroupSettingsPreview) SetIsClusterAiAssistantEnabled(v bool) {
	o.IsClusterAiAssistantEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsClusterAiAssistantEnabled")
}

// SetIsClusterAiAssistantEnabledNil sets IsClusterAiAssistantEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsClusterAiAssistantEnabledNil() {
	o.IsClusterAiAssistantEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsClusterAiAssistantEnabled")
}

// GetIsCollectDatabaseSpecificsStatisticsEnabled returns the IsCollectDatabaseSpecificsStatisticsEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsCollectDatabaseSpecificsStatisticsEnabled() bool {
	if o == nil || IsNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsCollectDatabaseSpecificsStatisticsEnabled
}

// GetIsCollectDatabaseSpecificsStatisticsEnabledOk returns a tuple with the IsCollectDatabaseSpecificsStatisticsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsCollectDatabaseSpecificsStatisticsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		return nil, false
	}

	return o.IsCollectDatabaseSpecificsStatisticsEnabled, true
}

// HasIsCollectDatabaseSpecificsStatisticsEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsCollectDatabaseSpecificsStatisticsEnabled() bool {
	if o != nil && !IsNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		return true
	}

	return false
}

// SetIsCollectDatabaseSpecificsStatisticsEnabled gets a reference to the given bool and assigns it to the IsCollectDatabaseSpecificsStatisticsEnabled field.
func (o *GroupSettingsPreview) SetIsCollectDatabaseSpecificsStatisticsEnabled(v bool) {
	o.IsCollectDatabaseSpecificsStatisticsEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsCollectDatabaseSpecificsStatisticsEnabled")
}

// SetIsCollectDatabaseSpecificsStatisticsEnabledNil sets IsCollectDatabaseSpecificsStatisticsEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsCollectDatabaseSpecificsStatisticsEnabledNil() {
	o.IsCollectDatabaseSpecificsStatisticsEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsCollectDatabaseSpecificsStatisticsEnabled")
}

// GetIsDataExplorerEnabled returns the IsDataExplorerEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsDataExplorerEnabled() bool {
	if o == nil || IsNil(o.IsDataExplorerEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDataExplorerEnabled
}

// GetIsDataExplorerEnabledOk returns a tuple with the IsDataExplorerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsDataExplorerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDataExplorerEnabled) {
		return nil, false
	}

	return o.IsDataExplorerEnabled, true
}

// HasIsDataExplorerEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsDataExplorerEnabled() bool {
	if o != nil && !IsNil(o.IsDataExplorerEnabled) {
		return true
	}

	return false
}

// SetIsDataExplorerEnabled gets a reference to the given bool and assigns it to the IsDataExplorerEnabled field.
func (o *GroupSettingsPreview) SetIsDataExplorerEnabled(v bool) {
	o.IsDataExplorerEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsDataExplorerEnabled")
}

// SetIsDataExplorerEnabledNil sets IsDataExplorerEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsDataExplorerEnabledNil() {
	o.IsDataExplorerEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsDataExplorerEnabled")
}

// GetIsDataExplorerGenAIFeaturesEnabled returns the IsDataExplorerGenAIFeaturesEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsDataExplorerGenAIFeaturesEnabled() bool {
	if o == nil || IsNil(o.IsDataExplorerGenAIFeaturesEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDataExplorerGenAIFeaturesEnabled
}

// GetIsDataExplorerGenAIFeaturesEnabledOk returns a tuple with the IsDataExplorerGenAIFeaturesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsDataExplorerGenAIFeaturesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDataExplorerGenAIFeaturesEnabled) {
		return nil, false
	}

	return o.IsDataExplorerGenAIFeaturesEnabled, true
}

// HasIsDataExplorerGenAIFeaturesEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsDataExplorerGenAIFeaturesEnabled() bool {
	if o != nil && !IsNil(o.IsDataExplorerGenAIFeaturesEnabled) {
		return true
	}

	return false
}

// SetIsDataExplorerGenAIFeaturesEnabled gets a reference to the given bool and assigns it to the IsDataExplorerGenAIFeaturesEnabled field.
func (o *GroupSettingsPreview) SetIsDataExplorerGenAIFeaturesEnabled(v bool) {
	o.IsDataExplorerGenAIFeaturesEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsDataExplorerGenAIFeaturesEnabled")
}

// SetIsDataExplorerGenAIFeaturesEnabledNil sets IsDataExplorerGenAIFeaturesEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsDataExplorerGenAIFeaturesEnabledNil() {
	o.IsDataExplorerGenAIFeaturesEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsDataExplorerGenAIFeaturesEnabled")
}

// GetIsDataExplorerGenAISampleDocumentPassingEnabled returns the IsDataExplorerGenAISampleDocumentPassingEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsDataExplorerGenAISampleDocumentPassingEnabled() bool {
	if o == nil || IsNil(o.IsDataExplorerGenAISampleDocumentPassingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDataExplorerGenAISampleDocumentPassingEnabled
}

// GetIsDataExplorerGenAISampleDocumentPassingEnabledOk returns a tuple with the IsDataExplorerGenAISampleDocumentPassingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsDataExplorerGenAISampleDocumentPassingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDataExplorerGenAISampleDocumentPassingEnabled) {
		return nil, false
	}

	return o.IsDataExplorerGenAISampleDocumentPassingEnabled, true
}

// HasIsDataExplorerGenAISampleDocumentPassingEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsDataExplorerGenAISampleDocumentPassingEnabled() bool {
	if o != nil && !IsNil(o.IsDataExplorerGenAISampleDocumentPassingEnabled) {
		return true
	}

	return false
}

// SetIsDataExplorerGenAISampleDocumentPassingEnabled gets a reference to the given bool and assigns it to the IsDataExplorerGenAISampleDocumentPassingEnabled field.
func (o *GroupSettingsPreview) SetIsDataExplorerGenAISampleDocumentPassingEnabled(v bool) {
	o.IsDataExplorerGenAISampleDocumentPassingEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsDataExplorerGenAISampleDocumentPassingEnabled")
}

// SetIsDataExplorerGenAISampleDocumentPassingEnabledNil sets IsDataExplorerGenAISampleDocumentPassingEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsDataExplorerGenAISampleDocumentPassingEnabledNil() {
	o.IsDataExplorerGenAISampleDocumentPassingEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsDataExplorerGenAISampleDocumentPassingEnabled")
}

// GetIsDataValidationEnabled returns the IsDataValidationEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsDataValidationEnabled() bool {
	if o == nil || IsNil(o.IsDataValidationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDataValidationEnabled
}

// GetIsDataValidationEnabledOk returns a tuple with the IsDataValidationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsDataValidationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDataValidationEnabled) {
		return nil, false
	}

	return o.IsDataValidationEnabled, true
}

// HasIsDataValidationEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsDataValidationEnabled() bool {
	if o != nil && !IsNil(o.IsDataValidationEnabled) {
		return true
	}

	return false
}

// SetIsDataValidationEnabled gets a reference to the given bool and assigns it to the IsDataValidationEnabled field.
func (o *GroupSettingsPreview) SetIsDataValidationEnabled(v bool) {
	o.IsDataValidationEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsDataValidationEnabled")
}

// SetIsDataValidationEnabledNil sets IsDataValidationEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsDataValidationEnabledNil() {
	o.IsDataValidationEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsDataValidationEnabled")
}

// GetIsExtendedStorageSizesEnabled returns the IsExtendedStorageSizesEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsExtendedStorageSizesEnabled() bool {
	if o == nil || IsNil(o.IsExtendedStorageSizesEnabled) {
		var ret bool
		return ret
	}
	return *o.IsExtendedStorageSizesEnabled
}

// GetIsExtendedStorageSizesEnabledOk returns a tuple with the IsExtendedStorageSizesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsExtendedStorageSizesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExtendedStorageSizesEnabled) {
		return nil, false
	}

	return o.IsExtendedStorageSizesEnabled, true
}

// HasIsExtendedStorageSizesEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsExtendedStorageSizesEnabled() bool {
	if o != nil && !IsNil(o.IsExtendedStorageSizesEnabled) {
		return true
	}

	return false
}

// SetIsExtendedStorageSizesEnabled gets a reference to the given bool and assigns it to the IsExtendedStorageSizesEnabled field.
func (o *GroupSettingsPreview) SetIsExtendedStorageSizesEnabled(v bool) {
	o.IsExtendedStorageSizesEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsExtendedStorageSizesEnabled")
}

// SetIsExtendedStorageSizesEnabledNil sets IsExtendedStorageSizesEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsExtendedStorageSizesEnabledNil() {
	o.IsExtendedStorageSizesEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsExtendedStorageSizesEnabled")
}

// GetIsNativeRerankingEnabled returns the IsNativeRerankingEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsNativeRerankingEnabled() bool {
	if o == nil || IsNil(o.IsNativeRerankingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsNativeRerankingEnabled
}

// GetIsNativeRerankingEnabledOk returns a tuple with the IsNativeRerankingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsNativeRerankingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNativeRerankingEnabled) {
		return nil, false
	}

	return o.IsNativeRerankingEnabled, true
}

// HasIsNativeRerankingEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsNativeRerankingEnabled() bool {
	if o != nil && !IsNil(o.IsNativeRerankingEnabled) {
		return true
	}

	return false
}

// SetIsNativeRerankingEnabled gets a reference to the given bool and assigns it to the IsNativeRerankingEnabled field.
func (o *GroupSettingsPreview) SetIsNativeRerankingEnabled(v bool) {
	o.IsNativeRerankingEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsNativeRerankingEnabled")
}

// SetIsNativeRerankingEnabledNil sets IsNativeRerankingEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsNativeRerankingEnabledNil() {
	o.IsNativeRerankingEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsNativeRerankingEnabled")
}

// GetIsPerformanceAdvisorEnabled returns the IsPerformanceAdvisorEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsPerformanceAdvisorEnabled() bool {
	if o == nil || IsNil(o.IsPerformanceAdvisorEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPerformanceAdvisorEnabled
}

// GetIsPerformanceAdvisorEnabledOk returns a tuple with the IsPerformanceAdvisorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsPerformanceAdvisorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPerformanceAdvisorEnabled) {
		return nil, false
	}

	return o.IsPerformanceAdvisorEnabled, true
}

// HasIsPerformanceAdvisorEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsPerformanceAdvisorEnabled() bool {
	if o != nil && !IsNil(o.IsPerformanceAdvisorEnabled) {
		return true
	}

	return false
}

// SetIsPerformanceAdvisorEnabled gets a reference to the given bool and assigns it to the IsPerformanceAdvisorEnabled field.
func (o *GroupSettingsPreview) SetIsPerformanceAdvisorEnabled(v bool) {
	o.IsPerformanceAdvisorEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsPerformanceAdvisorEnabled")
}

// SetIsPerformanceAdvisorEnabledNil sets IsPerformanceAdvisorEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsPerformanceAdvisorEnabledNil() {
	o.IsPerformanceAdvisorEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsPerformanceAdvisorEnabled")
}

// GetIsQueryInsightsGenAiFeaturesEnabled returns the IsQueryInsightsGenAiFeaturesEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsQueryInsightsGenAiFeaturesEnabled() bool {
	if o == nil || IsNil(o.IsQueryInsightsGenAiFeaturesEnabled) {
		var ret bool
		return ret
	}
	return *o.IsQueryInsightsGenAiFeaturesEnabled
}

// GetIsQueryInsightsGenAiFeaturesEnabledOk returns a tuple with the IsQueryInsightsGenAiFeaturesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsQueryInsightsGenAiFeaturesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQueryInsightsGenAiFeaturesEnabled) {
		return nil, false
	}

	return o.IsQueryInsightsGenAiFeaturesEnabled, true
}

// HasIsQueryInsightsGenAiFeaturesEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsQueryInsightsGenAiFeaturesEnabled() bool {
	if o != nil && !IsNil(o.IsQueryInsightsGenAiFeaturesEnabled) {
		return true
	}

	return false
}

// SetIsQueryInsightsGenAiFeaturesEnabled gets a reference to the given bool and assigns it to the IsQueryInsightsGenAiFeaturesEnabled field.
func (o *GroupSettingsPreview) SetIsQueryInsightsGenAiFeaturesEnabled(v bool) {
	o.IsQueryInsightsGenAiFeaturesEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsQueryInsightsGenAiFeaturesEnabled")
}

// SetIsQueryInsightsGenAiFeaturesEnabledNil sets IsQueryInsightsGenAiFeaturesEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsQueryInsightsGenAiFeaturesEnabledNil() {
	o.IsQueryInsightsGenAiFeaturesEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsQueryInsightsGenAiFeaturesEnabled")
}

// GetIsRealtimePerformancePanelEnabled returns the IsRealtimePerformancePanelEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsRealtimePerformancePanelEnabled() bool {
	if o == nil || IsNil(o.IsRealtimePerformancePanelEnabled) {
		var ret bool
		return ret
	}
	return *o.IsRealtimePerformancePanelEnabled
}

// GetIsRealtimePerformancePanelEnabledOk returns a tuple with the IsRealtimePerformancePanelEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsRealtimePerformancePanelEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRealtimePerformancePanelEnabled) {
		return nil, false
	}

	return o.IsRealtimePerformancePanelEnabled, true
}

// HasIsRealtimePerformancePanelEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsRealtimePerformancePanelEnabled() bool {
	if o != nil && !IsNil(o.IsRealtimePerformancePanelEnabled) {
		return true
	}

	return false
}

// SetIsRealtimePerformancePanelEnabled gets a reference to the given bool and assigns it to the IsRealtimePerformancePanelEnabled field.
func (o *GroupSettingsPreview) SetIsRealtimePerformancePanelEnabled(v bool) {
	o.IsRealtimePerformancePanelEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsRealtimePerformancePanelEnabled")
}

// SetIsRealtimePerformancePanelEnabledNil sets IsRealtimePerformancePanelEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsRealtimePerformancePanelEnabledNil() {
	o.IsRealtimePerformancePanelEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsRealtimePerformancePanelEnabled")
}

// GetIsSchemaAdvisorEnabled returns the IsSchemaAdvisorEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetIsSchemaAdvisorEnabled() bool {
	if o == nil || IsNil(o.IsSchemaAdvisorEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSchemaAdvisorEnabled
}

// GetIsSchemaAdvisorEnabledOk returns a tuple with the IsSchemaAdvisorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetIsSchemaAdvisorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSchemaAdvisorEnabled) {
		return nil, false
	}

	return o.IsSchemaAdvisorEnabled, true
}

// HasIsSchemaAdvisorEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasIsSchemaAdvisorEnabled() bool {
	if o != nil && !IsNil(o.IsSchemaAdvisorEnabled) {
		return true
	}

	return false
}

// SetIsSchemaAdvisorEnabled gets a reference to the given bool and assigns it to the IsSchemaAdvisorEnabled field.
func (o *GroupSettingsPreview) SetIsSchemaAdvisorEnabled(v bool) {
	o.IsSchemaAdvisorEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "IsSchemaAdvisorEnabled")
}

// SetIsSchemaAdvisorEnabledNil sets IsSchemaAdvisorEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetIsSchemaAdvisorEnabledNil() {
	o.IsSchemaAdvisorEnabled = nil
	o.NullFields = addNullField(o.NullFields, "IsSchemaAdvisorEnabled")
}

// GetPrivateEndpointConnectionStringsEnabled returns the PrivateEndpointConnectionStringsEnabled field value if set, zero value otherwise
func (o *GroupSettingsPreview) GetPrivateEndpointConnectionStringsEnabled() bool {
	if o == nil || IsNil(o.PrivateEndpointConnectionStringsEnabled) {
		var ret bool
		return ret
	}
	return *o.PrivateEndpointConnectionStringsEnabled
}

// GetPrivateEndpointConnectionStringsEnabledOk returns a tuple with the PrivateEndpointConnectionStringsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettingsPreview) GetPrivateEndpointConnectionStringsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivateEndpointConnectionStringsEnabled) {
		return nil, false
	}

	return o.PrivateEndpointConnectionStringsEnabled, true
}

// HasPrivateEndpointConnectionStringsEnabled returns a boolean if a field has been set.
func (o *GroupSettingsPreview) HasPrivateEndpointConnectionStringsEnabled() bool {
	if o != nil && !IsNil(o.PrivateEndpointConnectionStringsEnabled) {
		return true
	}

	return false
}

// SetPrivateEndpointConnectionStringsEnabled gets a reference to the given bool and assigns it to the PrivateEndpointConnectionStringsEnabled field.
func (o *GroupSettingsPreview) SetPrivateEndpointConnectionStringsEnabled(v bool) {
	o.PrivateEndpointConnectionStringsEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "PrivateEndpointConnectionStringsEnabled")
}

// SetPrivateEndpointConnectionStringsEnabledNil sets PrivateEndpointConnectionStringsEnabled to an explicit JSON null when marshaled.
func (o *GroupSettingsPreview) SetPrivateEndpointConnectionStringsEnabledNil() {
	o.PrivateEndpointConnectionStringsEnabled = nil
	o.NullFields = addNullField(o.NullFields, "PrivateEndpointConnectionStringsEnabled")
}
