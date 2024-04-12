# MeasurementsCollStatsLatencyCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | Unique identifier for Clusters. | [optional] [readonly] 
**ClusterView** | Pointer to **string** | Cluster topology view. | [optional] [readonly] 
**CollectionName** | Pointer to **string** | Human-readable label that identifies the collection. | [optional] [readonly] 
**DatabaseName** | Pointer to **string** | Human-readable label that identifies the database that the specified MongoDB process serves. | [optional] [readonly] 
**End** | Pointer to **time.Time** | Date and time that specifies when to stop retrieving measurements. If you set **end**, you must set **start**. You can&#39;t set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Granularity** | Pointer to **string** | Duration that specifies the interval between measurement data points. The parameter expresses its value in ISO 8601 timestamp format in UTC. If you set this parameter, you must set either **period** or **start** and **end**. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project. The project contains MongoDB processes that you want to return. The MongoDB process can be either the &#x60;mongod&#x60; or &#x60;mongos&#x60;. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Measurements** | Pointer to [**[]MetricsMeasurement**](MetricsMeasurement.md) | List that contains measurements and their data points. | [optional] [readonly] 
**ProcessId** | Pointer to **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | [optional] [readonly] 
**Start** | Pointer to **time.Time** | Date and time that specifies when to start retrieving measurements. If you set **start**, you must set **end**. You can&#39;t set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewMeasurementsCollStatsLatencyCluster

`func NewMeasurementsCollStatsLatencyCluster() *MeasurementsCollStatsLatencyCluster`

NewMeasurementsCollStatsLatencyCluster instantiates a new MeasurementsCollStatsLatencyCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementsCollStatsLatencyClusterWithDefaults

`func NewMeasurementsCollStatsLatencyClusterWithDefaults() *MeasurementsCollStatsLatencyCluster`

NewMeasurementsCollStatsLatencyClusterWithDefaults instantiates a new MeasurementsCollStatsLatencyCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *MeasurementsCollStatsLatencyCluster) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *MeasurementsCollStatsLatencyCluster) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *MeasurementsCollStatsLatencyCluster) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *MeasurementsCollStatsLatencyCluster) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.
### GetClusterView

`func (o *MeasurementsCollStatsLatencyCluster) GetClusterView() string`

GetClusterView returns the ClusterView field if non-nil, zero value otherwise.

### GetClusterViewOk

`func (o *MeasurementsCollStatsLatencyCluster) GetClusterViewOk() (*string, bool)`

GetClusterViewOk returns a tuple with the ClusterView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterView

`func (o *MeasurementsCollStatsLatencyCluster) SetClusterView(v string)`

SetClusterView sets ClusterView field to given value.

### HasClusterView

`func (o *MeasurementsCollStatsLatencyCluster) HasClusterView() bool`

HasClusterView returns a boolean if a field has been set.
### GetCollectionName

`func (o *MeasurementsCollStatsLatencyCluster) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *MeasurementsCollStatsLatencyCluster) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *MeasurementsCollStatsLatencyCluster) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *MeasurementsCollStatsLatencyCluster) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.
### GetDatabaseName

`func (o *MeasurementsCollStatsLatencyCluster) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *MeasurementsCollStatsLatencyCluster) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *MeasurementsCollStatsLatencyCluster) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *MeasurementsCollStatsLatencyCluster) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.
### GetEnd

`func (o *MeasurementsCollStatsLatencyCluster) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MeasurementsCollStatsLatencyCluster) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MeasurementsCollStatsLatencyCluster) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MeasurementsCollStatsLatencyCluster) HasEnd() bool`

HasEnd returns a boolean if a field has been set.
### GetGranularity

`func (o *MeasurementsCollStatsLatencyCluster) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MeasurementsCollStatsLatencyCluster) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MeasurementsCollStatsLatencyCluster) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MeasurementsCollStatsLatencyCluster) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.
### GetGroupId

`func (o *MeasurementsCollStatsLatencyCluster) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MeasurementsCollStatsLatencyCluster) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MeasurementsCollStatsLatencyCluster) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MeasurementsCollStatsLatencyCluster) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetLinks

`func (o *MeasurementsCollStatsLatencyCluster) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MeasurementsCollStatsLatencyCluster) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MeasurementsCollStatsLatencyCluster) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MeasurementsCollStatsLatencyCluster) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMeasurements

`func (o *MeasurementsCollStatsLatencyCluster) GetMeasurements() []MetricsMeasurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *MeasurementsCollStatsLatencyCluster) GetMeasurementsOk() (*[]MetricsMeasurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *MeasurementsCollStatsLatencyCluster) SetMeasurements(v []MetricsMeasurement)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *MeasurementsCollStatsLatencyCluster) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.
### GetProcessId

`func (o *MeasurementsCollStatsLatencyCluster) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *MeasurementsCollStatsLatencyCluster) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *MeasurementsCollStatsLatencyCluster) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *MeasurementsCollStatsLatencyCluster) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.
### GetStart

`func (o *MeasurementsCollStatsLatencyCluster) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MeasurementsCollStatsLatencyCluster) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MeasurementsCollStatsLatencyCluster) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *MeasurementsCollStatsLatencyCluster) HasStart() bool`

HasStart returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


