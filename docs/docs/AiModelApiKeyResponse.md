# AiModelApiKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **string** | Identifier used to reference this API key in admin API calls. | [optional] [readonly] 
**Cloud** | Pointer to **string** | Cloud provider scope for this API key. Use \&quot;ANY\&quot; for cloud-agnostic scope. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | UTC date when the API key was created. This parameter is formatted as an ISO 8601 timestamp. | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | Name of the user that created this API key. If no user name is available, the user ID is returned. | [optional] [readonly] 
**Endpoint** | Pointer to **string** | Server-computed endpoint hostname derived from cloud and geography. This field is read-only and must not be supplied in request bodies. | [optional] [readonly] 
**Geography** | Pointer to **string** | Geography scope for this API key. Use \&quot;ANY\&quot; for geography-agnostic scope. | [optional] [readonly] 
**GroupId** | Pointer to **string** | ID of the Atlas group this API key belongs to. | [optional] [readonly] 
**LastUsedAt** | Pointer to **string** | UTC date when the API key was last used. This parameter is formatted as an ISO 8601 timestamp. | [optional] [readonly] 
**MaskedSecret** | Pointer to **string** | A partially obfuscated version of the API key secret returned when the API key was created. | [optional] [readonly] 
**Name** | Pointer to **string** | Arbitrary string identifier assigned to this API key for convenient identification. | [optional] 
**Secret** | Pointer to **string** | The full API key secret used for interacting with the embedding / reranking service. Note: this will only be fully populated in the response to a create API key request. Responses to get, list, and update requests will not include the secret. | [optional] [readonly] 
**Status** | Pointer to **string** | A string describing the current status of the API key. | [optional] [readonly] 

## Methods

### NewAiModelApiKeyResponse

`func NewAiModelApiKeyResponse() *AiModelApiKeyResponse`

NewAiModelApiKeyResponse instantiates a new AiModelApiKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelApiKeyResponseWithDefaults

`func NewAiModelApiKeyResponseWithDefaults() *AiModelApiKeyResponse`

NewAiModelApiKeyResponseWithDefaults instantiates a new AiModelApiKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *AiModelApiKeyResponse) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *AiModelApiKeyResponse) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *AiModelApiKeyResponse) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *AiModelApiKeyResponse) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.
### GetCloud

`func (o *AiModelApiKeyResponse) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AiModelApiKeyResponse) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AiModelApiKeyResponse) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *AiModelApiKeyResponse) HasCloud() bool`

HasCloud returns a boolean if a field has been set.
### GetCreatedAt

`func (o *AiModelApiKeyResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiModelApiKeyResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiModelApiKeyResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AiModelApiKeyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetCreatedBy

`func (o *AiModelApiKeyResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AiModelApiKeyResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AiModelApiKeyResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AiModelApiKeyResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.
### GetEndpoint

`func (o *AiModelApiKeyResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AiModelApiKeyResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AiModelApiKeyResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AiModelApiKeyResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.
### GetGeography

`func (o *AiModelApiKeyResponse) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *AiModelApiKeyResponse) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *AiModelApiKeyResponse) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *AiModelApiKeyResponse) HasGeography() bool`

HasGeography returns a boolean if a field has been set.
### GetGroupId

`func (o *AiModelApiKeyResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AiModelApiKeyResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AiModelApiKeyResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AiModelApiKeyResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetLastUsedAt

`func (o *AiModelApiKeyResponse) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AiModelApiKeyResponse) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AiModelApiKeyResponse) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *AiModelApiKeyResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.
### GetMaskedSecret

`func (o *AiModelApiKeyResponse) GetMaskedSecret() string`

GetMaskedSecret returns the MaskedSecret field if non-nil, zero value otherwise.

### GetMaskedSecretOk

`func (o *AiModelApiKeyResponse) GetMaskedSecretOk() (*string, bool)`

GetMaskedSecretOk returns a tuple with the MaskedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedSecret

`func (o *AiModelApiKeyResponse) SetMaskedSecret(v string)`

SetMaskedSecret sets MaskedSecret field to given value.

### HasMaskedSecret

`func (o *AiModelApiKeyResponse) HasMaskedSecret() bool`

HasMaskedSecret returns a boolean if a field has been set.
### GetName

`func (o *AiModelApiKeyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiModelApiKeyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiModelApiKeyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AiModelApiKeyResponse) HasName() bool`

HasName returns a boolean if a field has been set.
### GetSecret

`func (o *AiModelApiKeyResponse) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AiModelApiKeyResponse) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AiModelApiKeyResponse) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AiModelApiKeyResponse) HasSecret() bool`

HasSecret returns a boolean if a field has been set.
### GetStatus

`func (o *AiModelApiKeyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiModelApiKeyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiModelApiKeyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AiModelApiKeyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


