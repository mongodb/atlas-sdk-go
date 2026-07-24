# ApiAtlasNonCompliantResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the organization the resource belongs to. | [optional] [readonly] 
**ResourceId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the non-compliant resource. | [optional] [readonly] 
**ResourceName** | Pointer to **string** | Unique human readable string that identifies the non-compliant resource. | [optional] [readonly] 
**ResourcePoliciesCausingNonCompliance** | Pointer to [**[]ApiAtlasResourcePolicyMetadata**](ApiAtlasResourcePolicyMetadata.md) | List of resource policies causing the resource to be considered non-compliant. | [optional] [readonly] 
**ResourceType** | Pointer to **string** | Human-readable label that displays the type of a resource. | [optional] [readonly] 

## Methods

### NewApiAtlasNonCompliantResource

`func NewApiAtlasNonCompliantResource() *ApiAtlasNonCompliantResource`

NewApiAtlasNonCompliantResource instantiates a new ApiAtlasNonCompliantResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasNonCompliantResourceWithDefaults

`func NewApiAtlasNonCompliantResourceWithDefaults() *ApiAtlasNonCompliantResource`

NewApiAtlasNonCompliantResourceWithDefaults instantiates a new ApiAtlasNonCompliantResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ApiAtlasNonCompliantResource) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiAtlasNonCompliantResource) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiAtlasNonCompliantResource) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ApiAtlasNonCompliantResource) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *ApiAtlasNonCompliantResource) SetOrgIdNil()`

SetOrgIdNil sets OrgId to an explicit JSON null when marshaled, overriding any value previously set with SetOrgId. Calling SetOrgId again clears the null override.

### GetResourceId

`func (o *ApiAtlasNonCompliantResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ApiAtlasNonCompliantResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ApiAtlasNonCompliantResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ApiAtlasNonCompliantResource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *ApiAtlasNonCompliantResource) SetResourceIdNil()`

SetResourceIdNil sets ResourceId to an explicit JSON null when marshaled, overriding any value previously set with SetResourceId. Calling SetResourceId again clears the null override.

### GetResourceName

`func (o *ApiAtlasNonCompliantResource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ApiAtlasNonCompliantResource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ApiAtlasNonCompliantResource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ApiAtlasNonCompliantResource) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *ApiAtlasNonCompliantResource) SetResourceNameNil()`

SetResourceNameNil sets ResourceName to an explicit JSON null when marshaled, overriding any value previously set with SetResourceName. Calling SetResourceName again clears the null override.

### GetResourcePoliciesCausingNonCompliance

`func (o *ApiAtlasNonCompliantResource) GetResourcePoliciesCausingNonCompliance() []ApiAtlasResourcePolicyMetadata`

GetResourcePoliciesCausingNonCompliance returns the ResourcePoliciesCausingNonCompliance field if non-nil, zero value otherwise.

### GetResourcePoliciesCausingNonComplianceOk

`func (o *ApiAtlasNonCompliantResource) GetResourcePoliciesCausingNonComplianceOk() (*[]ApiAtlasResourcePolicyMetadata, bool)`

GetResourcePoliciesCausingNonComplianceOk returns a tuple with the ResourcePoliciesCausingNonCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoliciesCausingNonCompliance

`func (o *ApiAtlasNonCompliantResource) SetResourcePoliciesCausingNonCompliance(v []ApiAtlasResourcePolicyMetadata)`

SetResourcePoliciesCausingNonCompliance sets ResourcePoliciesCausingNonCompliance field to given value.

### HasResourcePoliciesCausingNonCompliance

`func (o *ApiAtlasNonCompliantResource) HasResourcePoliciesCausingNonCompliance() bool`

HasResourcePoliciesCausingNonCompliance returns a boolean if a field has been set.

### SetResourcePoliciesCausingNonComplianceNil

`func (o *ApiAtlasNonCompliantResource) SetResourcePoliciesCausingNonComplianceNil()`

SetResourcePoliciesCausingNonComplianceNil sets ResourcePoliciesCausingNonCompliance to an explicit JSON null when marshaled, overriding any value previously set with SetResourcePoliciesCausingNonCompliance. Calling SetResourcePoliciesCausingNonCompliance again clears the null override.

### GetResourceType

`func (o *ApiAtlasNonCompliantResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ApiAtlasNonCompliantResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ApiAtlasNonCompliantResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ApiAtlasNonCompliantResource) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *ApiAtlasNonCompliantResource) SetResourceTypeNil()`

SetResourceTypeNil sets ResourceType to an explicit JSON null when marshaled, overriding any value previously set with SetResourceType. Calling SetResourceType again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


