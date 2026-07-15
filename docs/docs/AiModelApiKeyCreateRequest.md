# AiModelApiKeyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | **string** | Cloud provider scope for this API key. Must be \&quot;ANY\&quot;. Additional cloud values will be supported in future API versions. | 
**Geography** | **string** | Geography scope for this API key. Must be \&quot;ANY\&quot;. Additional geography values will be supported in future API versions. | 
**Name** | **string** | A name for the new API key that will be created. | 

## Methods

### NewAiModelApiKeyCreateRequest

`func NewAiModelApiKeyCreateRequest(cloud string, geography string, name string, ) *AiModelApiKeyCreateRequest`

NewAiModelApiKeyCreateRequest instantiates a new AiModelApiKeyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelApiKeyCreateRequestWithDefaults

`func NewAiModelApiKeyCreateRequestWithDefaults() *AiModelApiKeyCreateRequest`

NewAiModelApiKeyCreateRequestWithDefaults instantiates a new AiModelApiKeyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *AiModelApiKeyCreateRequest) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AiModelApiKeyCreateRequest) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AiModelApiKeyCreateRequest) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### GetGeography

`func (o *AiModelApiKeyCreateRequest) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *AiModelApiKeyCreateRequest) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *AiModelApiKeyCreateRequest) SetGeography(v string)`

SetGeography sets Geography field to given value.

### GetName

`func (o *AiModelApiKeyCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiModelApiKeyCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiModelApiKeyCreateRequest) SetName(v string)`

SetName sets Name field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


