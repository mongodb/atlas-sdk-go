# ApiAtlasResourcePolicyMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoliciesCausingNonCompliance** | Pointer to [**[]ApiAtlasPolicyMetadata**](ApiAtlasPolicyMetadata.md) | List of policies that are in conflict with the current state of the resource. | [optional] [readonly] 
**ResourcePolicyId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the atlas resource policy. | [optional] [readonly] 
**ResourcePolicyName** | Pointer to **string** | Human-readable label that describes the atlas resource policy. | [optional] [readonly] 

## Methods

### NewApiAtlasResourcePolicyMetadata

`func NewApiAtlasResourcePolicyMetadata() *ApiAtlasResourcePolicyMetadata`

NewApiAtlasResourcePolicyMetadata instantiates a new ApiAtlasResourcePolicyMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasResourcePolicyMetadataWithDefaults

`func NewApiAtlasResourcePolicyMetadataWithDefaults() *ApiAtlasResourcePolicyMetadata`

NewApiAtlasResourcePolicyMetadataWithDefaults instantiates a new ApiAtlasResourcePolicyMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoliciesCausingNonCompliance

`func (o *ApiAtlasResourcePolicyMetadata) GetPoliciesCausingNonCompliance() []ApiAtlasPolicyMetadata`

GetPoliciesCausingNonCompliance returns the PoliciesCausingNonCompliance field if non-nil, zero value otherwise.

### GetPoliciesCausingNonComplianceOk

`func (o *ApiAtlasResourcePolicyMetadata) GetPoliciesCausingNonComplianceOk() (*[]ApiAtlasPolicyMetadata, bool)`

GetPoliciesCausingNonComplianceOk returns a tuple with the PoliciesCausingNonCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesCausingNonCompliance

`func (o *ApiAtlasResourcePolicyMetadata) SetPoliciesCausingNonCompliance(v []ApiAtlasPolicyMetadata)`

SetPoliciesCausingNonCompliance sets PoliciesCausingNonCompliance field to given value.

### HasPoliciesCausingNonCompliance

`func (o *ApiAtlasResourcePolicyMetadata) HasPoliciesCausingNonCompliance() bool`

HasPoliciesCausingNonCompliance returns a boolean if a field has been set.
### GetResourcePolicyId

`func (o *ApiAtlasResourcePolicyMetadata) GetResourcePolicyId() string`

GetResourcePolicyId returns the ResourcePolicyId field if non-nil, zero value otherwise.

### GetResourcePolicyIdOk

`func (o *ApiAtlasResourcePolicyMetadata) GetResourcePolicyIdOk() (*string, bool)`

GetResourcePolicyIdOk returns a tuple with the ResourcePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicyId

`func (o *ApiAtlasResourcePolicyMetadata) SetResourcePolicyId(v string)`

SetResourcePolicyId sets ResourcePolicyId field to given value.

### HasResourcePolicyId

`func (o *ApiAtlasResourcePolicyMetadata) HasResourcePolicyId() bool`

HasResourcePolicyId returns a boolean if a field has been set.
### GetResourcePolicyName

`func (o *ApiAtlasResourcePolicyMetadata) GetResourcePolicyName() string`

GetResourcePolicyName returns the ResourcePolicyName field if non-nil, zero value otherwise.

### GetResourcePolicyNameOk

`func (o *ApiAtlasResourcePolicyMetadata) GetResourcePolicyNameOk() (*string, bool)`

GetResourcePolicyNameOk returns a tuple with the ResourcePolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicyName

`func (o *ApiAtlasResourcePolicyMetadata) SetResourcePolicyName(v string)`

SetResourcePolicyName sets ResourcePolicyName field to given value.

### HasResourcePolicyName

`func (o *ApiAtlasResourcePolicyMetadata) HasResourcePolicyName() bool`

HasResourcePolicyName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


