# GroupSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsClusterAiAssistantEnabled** | Pointer to **bool** | Flag that indicates whether the AI Cluster Assistant is enabled for the specified project. | [optional] 
**IsCollectDatabaseSpecificsStatisticsEnabled** | Pointer to **bool** | Flag that indicates whether to collect database-specific metrics for the specified project. | [optional] 
**IsDataExplorerEnabled** | Pointer to **bool** | Flag that indicates whether to enable the Data Explorer for the specified project. | [optional] 
**IsDataExplorerGenAIFeaturesEnabled** | Pointer to **bool** | Flag that indicates whether to enable the use of generative AI features which make requests to 3rd party services in Data Explorer for the specified project. | [optional] 
**IsDataExplorerGenAISampleDocumentPassingEnabled** | Pointer to **bool** | Flag that indicates whether to enable the passing of sample field values with the use of generative AI features in the Data Explorer for the specified project. | [optional] [default to false]
**IsExtendedStorageSizesEnabled** | Pointer to **bool** | Flag that indicates whether to enable extended storage sizes for the specified project. | [optional] 
**IsNativeRerankingEnabled** | Pointer to **bool** | Flag that indicates whether to enable Native Reranking with Voyage AI models in the Aggregation Pipeline for the specified project. | [optional] 
**IsPerformanceAdvisorEnabled** | Pointer to **bool** | Flag that indicates whether to enable the Performance Advisor and Profiler for the specified project. | [optional] 
**IsRealtimePerformancePanelEnabled** | Pointer to **bool** | Flag that indicates whether to enable the Real Time Performance Panel for the specified project. | [optional] 
**IsSchemaAdvisorEnabled** | Pointer to **bool** | Flag that indicates whether to enable the Schema Advisor for the specified project. | [optional] 

## Methods

### NewGroupSettings

`func NewGroupSettings() *GroupSettings`

NewGroupSettings instantiates a new GroupSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSettingsWithDefaults

`func NewGroupSettingsWithDefaults() *GroupSettings`

NewGroupSettingsWithDefaults instantiates a new GroupSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsClusterAiAssistantEnabled

`func (o *GroupSettings) GetIsClusterAiAssistantEnabled() bool`

GetIsClusterAiAssistantEnabled returns the IsClusterAiAssistantEnabled field if non-nil, zero value otherwise.

### GetIsClusterAiAssistantEnabledOk

`func (o *GroupSettings) GetIsClusterAiAssistantEnabledOk() (*bool, bool)`

GetIsClusterAiAssistantEnabledOk returns a tuple with the IsClusterAiAssistantEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClusterAiAssistantEnabled

`func (o *GroupSettings) SetIsClusterAiAssistantEnabled(v bool)`

SetIsClusterAiAssistantEnabled sets IsClusterAiAssistantEnabled field to given value.

### HasIsClusterAiAssistantEnabled

`func (o *GroupSettings) HasIsClusterAiAssistantEnabled() bool`

HasIsClusterAiAssistantEnabled returns a boolean if a field has been set.
### GetIsCollectDatabaseSpecificsStatisticsEnabled

`func (o *GroupSettings) GetIsCollectDatabaseSpecificsStatisticsEnabled() bool`

GetIsCollectDatabaseSpecificsStatisticsEnabled returns the IsCollectDatabaseSpecificsStatisticsEnabled field if non-nil, zero value otherwise.

### GetIsCollectDatabaseSpecificsStatisticsEnabledOk

`func (o *GroupSettings) GetIsCollectDatabaseSpecificsStatisticsEnabledOk() (*bool, bool)`

GetIsCollectDatabaseSpecificsStatisticsEnabledOk returns a tuple with the IsCollectDatabaseSpecificsStatisticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectDatabaseSpecificsStatisticsEnabled

`func (o *GroupSettings) SetIsCollectDatabaseSpecificsStatisticsEnabled(v bool)`

SetIsCollectDatabaseSpecificsStatisticsEnabled sets IsCollectDatabaseSpecificsStatisticsEnabled field to given value.

### HasIsCollectDatabaseSpecificsStatisticsEnabled

`func (o *GroupSettings) HasIsCollectDatabaseSpecificsStatisticsEnabled() bool`

HasIsCollectDatabaseSpecificsStatisticsEnabled returns a boolean if a field has been set.
### GetIsDataExplorerEnabled

`func (o *GroupSettings) GetIsDataExplorerEnabled() bool`

GetIsDataExplorerEnabled returns the IsDataExplorerEnabled field if non-nil, zero value otherwise.

### GetIsDataExplorerEnabledOk

`func (o *GroupSettings) GetIsDataExplorerEnabledOk() (*bool, bool)`

GetIsDataExplorerEnabledOk returns a tuple with the IsDataExplorerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataExplorerEnabled

`func (o *GroupSettings) SetIsDataExplorerEnabled(v bool)`

SetIsDataExplorerEnabled sets IsDataExplorerEnabled field to given value.

### HasIsDataExplorerEnabled

`func (o *GroupSettings) HasIsDataExplorerEnabled() bool`

HasIsDataExplorerEnabled returns a boolean if a field has been set.
### GetIsDataExplorerGenAIFeaturesEnabled

`func (o *GroupSettings) GetIsDataExplorerGenAIFeaturesEnabled() bool`

GetIsDataExplorerGenAIFeaturesEnabled returns the IsDataExplorerGenAIFeaturesEnabled field if non-nil, zero value otherwise.

### GetIsDataExplorerGenAIFeaturesEnabledOk

`func (o *GroupSettings) GetIsDataExplorerGenAIFeaturesEnabledOk() (*bool, bool)`

GetIsDataExplorerGenAIFeaturesEnabledOk returns a tuple with the IsDataExplorerGenAIFeaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataExplorerGenAIFeaturesEnabled

`func (o *GroupSettings) SetIsDataExplorerGenAIFeaturesEnabled(v bool)`

SetIsDataExplorerGenAIFeaturesEnabled sets IsDataExplorerGenAIFeaturesEnabled field to given value.

### HasIsDataExplorerGenAIFeaturesEnabled

`func (o *GroupSettings) HasIsDataExplorerGenAIFeaturesEnabled() bool`

HasIsDataExplorerGenAIFeaturesEnabled returns a boolean if a field has been set.
### GetIsDataExplorerGenAISampleDocumentPassingEnabled

`func (o *GroupSettings) GetIsDataExplorerGenAISampleDocumentPassingEnabled() bool`

GetIsDataExplorerGenAISampleDocumentPassingEnabled returns the IsDataExplorerGenAISampleDocumentPassingEnabled field if non-nil, zero value otherwise.

### GetIsDataExplorerGenAISampleDocumentPassingEnabledOk

`func (o *GroupSettings) GetIsDataExplorerGenAISampleDocumentPassingEnabledOk() (*bool, bool)`

GetIsDataExplorerGenAISampleDocumentPassingEnabledOk returns a tuple with the IsDataExplorerGenAISampleDocumentPassingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataExplorerGenAISampleDocumentPassingEnabled

`func (o *GroupSettings) SetIsDataExplorerGenAISampleDocumentPassingEnabled(v bool)`

SetIsDataExplorerGenAISampleDocumentPassingEnabled sets IsDataExplorerGenAISampleDocumentPassingEnabled field to given value.

### HasIsDataExplorerGenAISampleDocumentPassingEnabled

`func (o *GroupSettings) HasIsDataExplorerGenAISampleDocumentPassingEnabled() bool`

HasIsDataExplorerGenAISampleDocumentPassingEnabled returns a boolean if a field has been set.
### GetIsExtendedStorageSizesEnabled

`func (o *GroupSettings) GetIsExtendedStorageSizesEnabled() bool`

GetIsExtendedStorageSizesEnabled returns the IsExtendedStorageSizesEnabled field if non-nil, zero value otherwise.

### GetIsExtendedStorageSizesEnabledOk

`func (o *GroupSettings) GetIsExtendedStorageSizesEnabledOk() (*bool, bool)`

GetIsExtendedStorageSizesEnabledOk returns a tuple with the IsExtendedStorageSizesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExtendedStorageSizesEnabled

`func (o *GroupSettings) SetIsExtendedStorageSizesEnabled(v bool)`

SetIsExtendedStorageSizesEnabled sets IsExtendedStorageSizesEnabled field to given value.

### HasIsExtendedStorageSizesEnabled

`func (o *GroupSettings) HasIsExtendedStorageSizesEnabled() bool`

HasIsExtendedStorageSizesEnabled returns a boolean if a field has been set.
### GetIsNativeRerankingEnabled

`func (o *GroupSettings) GetIsNativeRerankingEnabled() bool`

GetIsNativeRerankingEnabled returns the IsNativeRerankingEnabled field if non-nil, zero value otherwise.

### GetIsNativeRerankingEnabledOk

`func (o *GroupSettings) GetIsNativeRerankingEnabledOk() (*bool, bool)`

GetIsNativeRerankingEnabledOk returns a tuple with the IsNativeRerankingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeRerankingEnabled

`func (o *GroupSettings) SetIsNativeRerankingEnabled(v bool)`

SetIsNativeRerankingEnabled sets IsNativeRerankingEnabled field to given value.

### HasIsNativeRerankingEnabled

`func (o *GroupSettings) HasIsNativeRerankingEnabled() bool`

HasIsNativeRerankingEnabled returns a boolean if a field has been set.
### GetIsPerformanceAdvisorEnabled

`func (o *GroupSettings) GetIsPerformanceAdvisorEnabled() bool`

GetIsPerformanceAdvisorEnabled returns the IsPerformanceAdvisorEnabled field if non-nil, zero value otherwise.

### GetIsPerformanceAdvisorEnabledOk

`func (o *GroupSettings) GetIsPerformanceAdvisorEnabledOk() (*bool, bool)`

GetIsPerformanceAdvisorEnabledOk returns a tuple with the IsPerformanceAdvisorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerformanceAdvisorEnabled

`func (o *GroupSettings) SetIsPerformanceAdvisorEnabled(v bool)`

SetIsPerformanceAdvisorEnabled sets IsPerformanceAdvisorEnabled field to given value.

### HasIsPerformanceAdvisorEnabled

`func (o *GroupSettings) HasIsPerformanceAdvisorEnabled() bool`

HasIsPerformanceAdvisorEnabled returns a boolean if a field has been set.
### GetIsRealtimePerformancePanelEnabled

`func (o *GroupSettings) GetIsRealtimePerformancePanelEnabled() bool`

GetIsRealtimePerformancePanelEnabled returns the IsRealtimePerformancePanelEnabled field if non-nil, zero value otherwise.

### GetIsRealtimePerformancePanelEnabledOk

`func (o *GroupSettings) GetIsRealtimePerformancePanelEnabledOk() (*bool, bool)`

GetIsRealtimePerformancePanelEnabledOk returns a tuple with the IsRealtimePerformancePanelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealtimePerformancePanelEnabled

`func (o *GroupSettings) SetIsRealtimePerformancePanelEnabled(v bool)`

SetIsRealtimePerformancePanelEnabled sets IsRealtimePerformancePanelEnabled field to given value.

### HasIsRealtimePerformancePanelEnabled

`func (o *GroupSettings) HasIsRealtimePerformancePanelEnabled() bool`

HasIsRealtimePerformancePanelEnabled returns a boolean if a field has been set.
### GetIsSchemaAdvisorEnabled

`func (o *GroupSettings) GetIsSchemaAdvisorEnabled() bool`

GetIsSchemaAdvisorEnabled returns the IsSchemaAdvisorEnabled field if non-nil, zero value otherwise.

### GetIsSchemaAdvisorEnabledOk

`func (o *GroupSettings) GetIsSchemaAdvisorEnabledOk() (*bool, bool)`

GetIsSchemaAdvisorEnabledOk returns a tuple with the IsSchemaAdvisorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSchemaAdvisorEnabled

`func (o *GroupSettings) SetIsSchemaAdvisorEnabled(v bool)`

SetIsSchemaAdvisorEnabled sets IsSchemaAdvisorEnabled field to given value.

### HasIsSchemaAdvisorEnabled

`func (o *GroupSettings) HasIsSchemaAdvisorEnabled() bool`

HasIsSchemaAdvisorEnabled returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


