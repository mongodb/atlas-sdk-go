# AiModelRateLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | Pointer to **string** | Cloud provider scope for this rate limit. Use \&quot;any\&quot; for cloud-agnostic scope. | [optional] [readonly] 
**Endpoint** | Pointer to **string** | Server-computed endpoint hostname derived from cloud and geography. This field is read-only and must not be supplied in request bodies. | [optional] [readonly] 
**Geography** | Pointer to **string** | Geography scope for this rate limit. Use \&quot;any\&quot; for geography-agnostic scope. | [optional] [readonly] 
**ModelGroupName** | Pointer to **string** | Identifier used to reference this model group. | [optional] [readonly] 
**ModelNames** | Pointer to **[]string** | List of embedding model names included in this model group. | [optional] [readonly] 
**RequestsPerMinuteLimit** | Pointer to **int** | The number of requests per minute allowed for this model group. Must be a positive integer. Cannot be more than the organization level limit for this group model. | [optional] 
**TokensPerMinuteLimit** | Pointer to **int** | The number of tokens per minute allowed for this model group. Must be a positive integer. Cannot be more than the organization level limit for this group model. | [optional] 

## Methods

### NewAiModelRateLimitResponse

`func NewAiModelRateLimitResponse() *AiModelRateLimitResponse`

NewAiModelRateLimitResponse instantiates a new AiModelRateLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelRateLimitResponseWithDefaults

`func NewAiModelRateLimitResponseWithDefaults() *AiModelRateLimitResponse`

NewAiModelRateLimitResponseWithDefaults instantiates a new AiModelRateLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *AiModelRateLimitResponse) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AiModelRateLimitResponse) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AiModelRateLimitResponse) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *AiModelRateLimitResponse) HasCloud() bool`

HasCloud returns a boolean if a field has been set.
### GetEndpoint

`func (o *AiModelRateLimitResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AiModelRateLimitResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AiModelRateLimitResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AiModelRateLimitResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.
### GetGeography

`func (o *AiModelRateLimitResponse) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *AiModelRateLimitResponse) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *AiModelRateLimitResponse) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *AiModelRateLimitResponse) HasGeography() bool`

HasGeography returns a boolean if a field has been set.
### GetModelGroupName

`func (o *AiModelRateLimitResponse) GetModelGroupName() string`

GetModelGroupName returns the ModelGroupName field if non-nil, zero value otherwise.

### GetModelGroupNameOk

`func (o *AiModelRateLimitResponse) GetModelGroupNameOk() (*string, bool)`

GetModelGroupNameOk returns a tuple with the ModelGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelGroupName

`func (o *AiModelRateLimitResponse) SetModelGroupName(v string)`

SetModelGroupName sets ModelGroupName field to given value.

### HasModelGroupName

`func (o *AiModelRateLimitResponse) HasModelGroupName() bool`

HasModelGroupName returns a boolean if a field has been set.
### GetModelNames

`func (o *AiModelRateLimitResponse) GetModelNames() []string`

GetModelNames returns the ModelNames field if non-nil, zero value otherwise.

### GetModelNamesOk

`func (o *AiModelRateLimitResponse) GetModelNamesOk() (*[]string, bool)`

GetModelNamesOk returns a tuple with the ModelNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNames

`func (o *AiModelRateLimitResponse) SetModelNames(v []string)`

SetModelNames sets ModelNames field to given value.

### HasModelNames

`func (o *AiModelRateLimitResponse) HasModelNames() bool`

HasModelNames returns a boolean if a field has been set.
### GetRequestsPerMinuteLimit

`func (o *AiModelRateLimitResponse) GetRequestsPerMinuteLimit() int`

GetRequestsPerMinuteLimit returns the RequestsPerMinuteLimit field if non-nil, zero value otherwise.

### GetRequestsPerMinuteLimitOk

`func (o *AiModelRateLimitResponse) GetRequestsPerMinuteLimitOk() (*int, bool)`

GetRequestsPerMinuteLimitOk returns a tuple with the RequestsPerMinuteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerMinuteLimit

`func (o *AiModelRateLimitResponse) SetRequestsPerMinuteLimit(v int)`

SetRequestsPerMinuteLimit sets RequestsPerMinuteLimit field to given value.

### HasRequestsPerMinuteLimit

`func (o *AiModelRateLimitResponse) HasRequestsPerMinuteLimit() bool`

HasRequestsPerMinuteLimit returns a boolean if a field has been set.
### GetTokensPerMinuteLimit

`func (o *AiModelRateLimitResponse) GetTokensPerMinuteLimit() int`

GetTokensPerMinuteLimit returns the TokensPerMinuteLimit field if non-nil, zero value otherwise.

### GetTokensPerMinuteLimitOk

`func (o *AiModelRateLimitResponse) GetTokensPerMinuteLimitOk() (*int, bool)`

GetTokensPerMinuteLimitOk returns a tuple with the TokensPerMinuteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensPerMinuteLimit

`func (o *AiModelRateLimitResponse) SetTokensPerMinuteLimit(v int)`

SetTokensPerMinuteLimit sets TokensPerMinuteLimit field to given value.

### HasTokensPerMinuteLimit

`func (o *AiModelRateLimitResponse) HasTokensPerMinuteLimit() bool`

HasTokensPerMinuteLimit returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


