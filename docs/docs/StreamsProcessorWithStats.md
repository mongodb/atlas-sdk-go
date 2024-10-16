# StreamsProcessorWithStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-hexadecimal character string that identifies the stream processor. | [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | **string** | Human-readable name of the stream processor. | [readonly] 
**Options** | Pointer to [**StreamsOptions**](StreamsOptions.md) |  | [optional] 
**Pipeline** | [**[]any**](any.md) | Stream aggregation pipeline you want to apply to your streaming data. | [readonly] 
**State** | **string** | The state of the stream processor. Commonly occurring states are &#39;CREATED&#39;, &#39;STARTED&#39;, &#39;STOPPED&#39; and &#39;FAILED&#39;. | [readonly] 
**Stats** | Pointer to [**any**](interface{}.md) | The stats associated with the stream processor. [MongoDB Atlas Docs](https://www.mongodb.com/docs/atlas/atlas-stream-processing/manage-stream-processor/#view-statistics-of-a-stream-processor) contain more information. | [optional] [readonly] 

## Methods

### NewStreamsProcessorWithStats

`func NewStreamsProcessorWithStats(id string, name string, pipeline []any, state string, ) *StreamsProcessorWithStats`

NewStreamsProcessorWithStats instantiates a new StreamsProcessorWithStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsProcessorWithStatsWithDefaults

`func NewStreamsProcessorWithStatsWithDefaults() *StreamsProcessorWithStats`

NewStreamsProcessorWithStatsWithDefaults instantiates a new StreamsProcessorWithStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamsProcessorWithStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsProcessorWithStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsProcessorWithStats) SetId(v string)`

SetId sets Id field to given value.

### GetLinks

`func (o *StreamsProcessorWithStats) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsProcessorWithStats) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsProcessorWithStats) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsProcessorWithStats) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetName

`func (o *StreamsProcessorWithStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsProcessorWithStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsProcessorWithStats) SetName(v string)`

SetName sets Name field to given value.

### GetOptions

`func (o *StreamsProcessorWithStats) GetOptions() StreamsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StreamsProcessorWithStats) GetOptionsOk() (*StreamsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StreamsProcessorWithStats) SetOptions(v StreamsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StreamsProcessorWithStats) HasOptions() bool`

HasOptions returns a boolean if a field has been set.
### GetPipeline

`func (o *StreamsProcessorWithStats) GetPipeline() []any`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *StreamsProcessorWithStats) GetPipelineOk() (*[]any, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *StreamsProcessorWithStats) SetPipeline(v []any)`

SetPipeline sets Pipeline field to given value.

### GetState

`func (o *StreamsProcessorWithStats) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamsProcessorWithStats) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamsProcessorWithStats) SetState(v string)`

SetState sets State field to given value.

### GetStats

`func (o *StreamsProcessorWithStats) GetStats() any`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StreamsProcessorWithStats) GetStatsOk() (*any, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StreamsProcessorWithStats) SetStats(v any)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StreamsProcessorWithStats) HasStats() bool`

HasStats returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


