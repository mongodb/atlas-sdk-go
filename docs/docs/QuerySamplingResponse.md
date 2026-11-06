# QuerySamplingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Flag that indicates whether query sampling is observably running on the namespace. May still be &#x60;false&#x60; in the response to a successful create request; it becomes &#x60;true&#x60; asynchronously once the cluster applies the sampling configuration. | [readonly] 
**Namespace** | **string** | Human-readable label that identifies the namespace (&#x60;&lt;database&gt;.&lt;collection&gt;&#x60;) on which queries are sampled. | 
**SampledReadsBytes** | Pointer to **int64** | Total size in bytes of the read operations sampled on the namespace. | [optional] [readonly] 
**SampledReadsCount** | Pointer to **int64** | Number of read operations sampled on the namespace. | [optional] [readonly] 
**SampledWritesBytes** | Pointer to **int64** | Total size in bytes of the write operations sampled on the namespace. | [optional] [readonly] 
**SampledWritesCount** | Pointer to **int64** | Number of write operations sampled on the namespace. | [optional] [readonly] 
**SamplesPerSecond** | Pointer to **float64** | Maximum number of queries sampled per second on the namespace. MongoDB Cloud sets this rate. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | Date and time when query sampling started on the namespace. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewQuerySamplingResponse

`func NewQuerySamplingResponse(active bool, namespace string, ) *QuerySamplingResponse`

NewQuerySamplingResponse instantiates a new QuerySamplingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySamplingResponseWithDefaults

`func NewQuerySamplingResponseWithDefaults() *QuerySamplingResponse`

NewQuerySamplingResponseWithDefaults instantiates a new QuerySamplingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *QuerySamplingResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QuerySamplingResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QuerySamplingResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### GetNamespace

`func (o *QuerySamplingResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *QuerySamplingResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *QuerySamplingResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### GetSampledReadsBytes

`func (o *QuerySamplingResponse) GetSampledReadsBytes() int64`

GetSampledReadsBytes returns the SampledReadsBytes field if non-nil, zero value otherwise.

### GetSampledReadsBytesOk

`func (o *QuerySamplingResponse) GetSampledReadsBytesOk() (*int64, bool)`

GetSampledReadsBytesOk returns a tuple with the SampledReadsBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampledReadsBytes

`func (o *QuerySamplingResponse) SetSampledReadsBytes(v int64)`

SetSampledReadsBytes sets SampledReadsBytes field to given value.

### HasSampledReadsBytes

`func (o *QuerySamplingResponse) HasSampledReadsBytes() bool`

HasSampledReadsBytes returns a boolean if a field has been set.

### SetSampledReadsBytesNil

`func (o *QuerySamplingResponse) SetSampledReadsBytesNil()`

SetSampledReadsBytesNil sets SampledReadsBytes to an explicit JSON null when marshaled, overriding any value previously set with SetSampledReadsBytes. Calling SetSampledReadsBytes again clears the null override.

### GetSampledReadsCount

`func (o *QuerySamplingResponse) GetSampledReadsCount() int64`

GetSampledReadsCount returns the SampledReadsCount field if non-nil, zero value otherwise.

### GetSampledReadsCountOk

`func (o *QuerySamplingResponse) GetSampledReadsCountOk() (*int64, bool)`

GetSampledReadsCountOk returns a tuple with the SampledReadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampledReadsCount

`func (o *QuerySamplingResponse) SetSampledReadsCount(v int64)`

SetSampledReadsCount sets SampledReadsCount field to given value.

### HasSampledReadsCount

`func (o *QuerySamplingResponse) HasSampledReadsCount() bool`

HasSampledReadsCount returns a boolean if a field has been set.

### SetSampledReadsCountNil

`func (o *QuerySamplingResponse) SetSampledReadsCountNil()`

SetSampledReadsCountNil sets SampledReadsCount to an explicit JSON null when marshaled, overriding any value previously set with SetSampledReadsCount. Calling SetSampledReadsCount again clears the null override.

### GetSampledWritesBytes

`func (o *QuerySamplingResponse) GetSampledWritesBytes() int64`

GetSampledWritesBytes returns the SampledWritesBytes field if non-nil, zero value otherwise.

### GetSampledWritesBytesOk

`func (o *QuerySamplingResponse) GetSampledWritesBytesOk() (*int64, bool)`

GetSampledWritesBytesOk returns a tuple with the SampledWritesBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampledWritesBytes

`func (o *QuerySamplingResponse) SetSampledWritesBytes(v int64)`

SetSampledWritesBytes sets SampledWritesBytes field to given value.

### HasSampledWritesBytes

`func (o *QuerySamplingResponse) HasSampledWritesBytes() bool`

HasSampledWritesBytes returns a boolean if a field has been set.

### SetSampledWritesBytesNil

`func (o *QuerySamplingResponse) SetSampledWritesBytesNil()`

SetSampledWritesBytesNil sets SampledWritesBytes to an explicit JSON null when marshaled, overriding any value previously set with SetSampledWritesBytes. Calling SetSampledWritesBytes again clears the null override.

### GetSampledWritesCount

`func (o *QuerySamplingResponse) GetSampledWritesCount() int64`

GetSampledWritesCount returns the SampledWritesCount field if non-nil, zero value otherwise.

### GetSampledWritesCountOk

`func (o *QuerySamplingResponse) GetSampledWritesCountOk() (*int64, bool)`

GetSampledWritesCountOk returns a tuple with the SampledWritesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampledWritesCount

`func (o *QuerySamplingResponse) SetSampledWritesCount(v int64)`

SetSampledWritesCount sets SampledWritesCount field to given value.

### HasSampledWritesCount

`func (o *QuerySamplingResponse) HasSampledWritesCount() bool`

HasSampledWritesCount returns a boolean if a field has been set.

### SetSampledWritesCountNil

`func (o *QuerySamplingResponse) SetSampledWritesCountNil()`

SetSampledWritesCountNil sets SampledWritesCount to an explicit JSON null when marshaled, overriding any value previously set with SetSampledWritesCount. Calling SetSampledWritesCount again clears the null override.

### GetSamplesPerSecond

`func (o *QuerySamplingResponse) GetSamplesPerSecond() float64`

GetSamplesPerSecond returns the SamplesPerSecond field if non-nil, zero value otherwise.

### GetSamplesPerSecondOk

`func (o *QuerySamplingResponse) GetSamplesPerSecondOk() (*float64, bool)`

GetSamplesPerSecondOk returns a tuple with the SamplesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesPerSecond

`func (o *QuerySamplingResponse) SetSamplesPerSecond(v float64)`

SetSamplesPerSecond sets SamplesPerSecond field to given value.

### HasSamplesPerSecond

`func (o *QuerySamplingResponse) HasSamplesPerSecond() bool`

HasSamplesPerSecond returns a boolean if a field has been set.

### SetSamplesPerSecondNil

`func (o *QuerySamplingResponse) SetSamplesPerSecondNil()`

SetSamplesPerSecondNil sets SamplesPerSecond to an explicit JSON null when marshaled, overriding any value previously set with SetSamplesPerSecond. Calling SetSamplesPerSecond again clears the null override.

### GetStartTime

`func (o *QuerySamplingResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *QuerySamplingResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *QuerySamplingResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *QuerySamplingResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *QuerySamplingResponse) SetStartTimeNil()`

SetStartTimeNil sets StartTime to an explicit JSON null when marshaled, overriding any value previously set with SetStartTime. Calling SetStartTime again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


