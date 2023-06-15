# MeasurementsNonIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **time.Time** | Date and time that specifies when to stop retrieving measurements. If you set **end**, you must set **start**. You can&#39;t set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Granularity** | Pointer to **string** | Duration that specifies the interval between measurement data points. The parameter expresses its value in ISO 8601 timestamp format in UTC. If you set this parameter, you must set either **period** or **start** and **end**. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project. The project contains MongoDB processes that you want to return. The MongoDB process can be either the &#x60;mongod&#x60; or &#x60;mongos&#x60;. | [optional] [readonly] 
**HardwareMeasurements** | Pointer to [**[]MetricsMeasurement**](MetricsMeasurement.md) | List that contains the Atlas Search hardware measurements. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**ProcessId** | Pointer to **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | [optional] [readonly] 
**Start** | Pointer to **time.Time** | Date and time that specifies when to start retrieving measurements. If you set **start**, you must set **end**. You can&#39;t set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**StatusMeasurements** | Pointer to [**[]MetricsMeasurement**](MetricsMeasurement.md) | List that contains the Atlas Search status measurements. | [optional] [readonly] 

## Methods

### NewMeasurementsNonIndex

`func NewMeasurementsNonIndex() *MeasurementsNonIndex`

NewMeasurementsNonIndex instantiates a new MeasurementsNonIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementsNonIndexWithDefaults

`func NewMeasurementsNonIndexWithDefaults() *MeasurementsNonIndex`

NewMeasurementsNonIndexWithDefaults instantiates a new MeasurementsNonIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *MeasurementsNonIndex) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MeasurementsNonIndex) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MeasurementsNonIndex) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MeasurementsNonIndex) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetGranularity

`func (o *MeasurementsNonIndex) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MeasurementsNonIndex) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MeasurementsNonIndex) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MeasurementsNonIndex) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetGroupId

`func (o *MeasurementsNonIndex) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MeasurementsNonIndex) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MeasurementsNonIndex) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MeasurementsNonIndex) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHardwareMeasurements

`func (o *MeasurementsNonIndex) GetHardwareMeasurements() []MetricsMeasurement`

GetHardwareMeasurements returns the HardwareMeasurements field if non-nil, zero value otherwise.

### GetHardwareMeasurementsOk

`func (o *MeasurementsNonIndex) GetHardwareMeasurementsOk() (*[]MetricsMeasurement, bool)`

GetHardwareMeasurementsOk returns a tuple with the HardwareMeasurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareMeasurements

`func (o *MeasurementsNonIndex) SetHardwareMeasurements(v []MetricsMeasurement)`

SetHardwareMeasurements sets HardwareMeasurements field to given value.

### HasHardwareMeasurements

`func (o *MeasurementsNonIndex) HasHardwareMeasurements() bool`

HasHardwareMeasurements returns a boolean if a field has been set.

### GetLinks

`func (o *MeasurementsNonIndex) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MeasurementsNonIndex) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MeasurementsNonIndex) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MeasurementsNonIndex) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProcessId

`func (o *MeasurementsNonIndex) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *MeasurementsNonIndex) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *MeasurementsNonIndex) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *MeasurementsNonIndex) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetStart

`func (o *MeasurementsNonIndex) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MeasurementsNonIndex) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MeasurementsNonIndex) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *MeasurementsNonIndex) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatusMeasurements

`func (o *MeasurementsNonIndex) GetStatusMeasurements() []MetricsMeasurement`

GetStatusMeasurements returns the StatusMeasurements field if non-nil, zero value otherwise.

### GetStatusMeasurementsOk

`func (o *MeasurementsNonIndex) GetStatusMeasurementsOk() (*[]MetricsMeasurement, bool)`

GetStatusMeasurementsOk returns a tuple with the StatusMeasurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMeasurements

`func (o *MeasurementsNonIndex) SetStatusMeasurements(v []MetricsMeasurement)`

SetStatusMeasurements sets StatusMeasurements field to given value.

### HasStatusMeasurements

`func (o *MeasurementsNonIndex) HasStatusMeasurements() bool`

HasStatusMeasurements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


